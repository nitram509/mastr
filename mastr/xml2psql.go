package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"golang.org/x/text/encoding/unicode"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	location         = time.UTC
	errMissingOption = errors.New("missing mandatory argument")
)

func main() {
	err := insert()
	if errors.Is(err, errMissingOption) {
		flag.PrintDefaults()
		os.Exit(64)
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func insert() error {
	const defaultOption = "<undefined>"
	exportFileName := flag.String("export", defaultOption, "file name of the export zip file")
	descriptorFileName := flag.String("descriptor", defaultOption, "file name of the table descriptor")
	filePrefix := flag.String("prefix", defaultOption, "prefix of xml files to extract")
	databaseUrl := flag.String("database", defaultOption, "postgres database URL")
	forceCreate := flag.Bool("force-create", false, "use CREATE instead of CREATE IF NOT EXISTS")
	flag.Parse()

	// Ensure mandatory flags are set.
	for _, arg := range []string{
		*exportFileName,
		*descriptorFileName,
		*filePrefix,
		*databaseUrl,
	} {
		if arg == defaultOption {
			return errMissingOption
		}
	}

	var err error
	location, err = time.LoadLocation("Europe/Berlin")
	if err != nil {
		return fmt.Errorf("failed to load location data: %w", err)
	}

	td, err := decodeDescriptor(*descriptorFileName)
	if err != nil {
		return fmt.Errorf("failed to decode descriptor: %w", err)
	}

	r, err := zip.OpenReader(*exportFileName)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Printf("%v", err)
		}
	}()

	// Connect to the database.
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, *databaseUrl)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("%v", err)
		}
	}()

	// Insert XML files one by one.
	dec := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	for _, xmlFile := range r.File {
		if !strings.HasPrefix(xmlFile.FileHeader.Name, *filePrefix) {
			continue
		}
		if err = func() error {
			start := time.Now()
			f, err := xmlFile.Open()
			if err != nil {
				return fmt.Errorf("failed to open xml file: %w", err)
			}
			defer func() {
				if err := f.Close(); err != nil {
					log.Printf("%v", err)
				}
			}()
			i, err := insertFromXml(dec.Reader(f), conn, ctx, td, *forceCreate)
			if err != nil {
				return fmt.Errorf("failed to insert from xml file: %w", err)
			}
			elapsed := time.Since(start).Seconds()
			log.Printf("%s\t%.f entries/second", xmlFile.FileHeader.Name, float64(i)/elapsed)
			return nil
		}(); err != nil {
			return fmt.Errorf("failed to process xml file: %w", err)
		}
	}
	return nil
}

type reference struct {
	Table  string `yaml:"table"`
	Column string `yaml:"column"`
}

type fieldDescriptor struct {
	Name       string     `yaml:"name"`
	Mandatory  bool       `yaml:"mandatory"`
	Xsd        string     `yaml:"xsd"`
	Sqlite     string     `yaml:"sqlite"`
	Psql       string     `yaml:"psql"`
	References *reference `yaml:"references,omitempty"`
}

type tableDescriptor struct {
	Root    string            `yaml:"root"`
	Element string            `yaml:"element"`
	Primary string            `yaml:"primary"`
	Fields  []fieldDescriptor `yaml:"fields"`
}

type fields struct {
	fields map[string]uint
	psqlty map[string]string
}

const (
	startRoot = iota
	startItemOrEndRoot
	startFieldOrEndItem
	fieldValueOrEndField
	finished
)

func decodeDescriptor(descriptorFileName string) (*tableDescriptor, error) {
	var tableDescriptor tableDescriptor
	f, err := os.Open(descriptorFileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("%v", err)
		}
	}()
	d := yaml.NewDecoder(f)
	err = d.Decode(&tableDescriptor)
	if err != nil {
		return nil, err
	}
	return &tableDescriptor, nil
}

func newFields(fieldDescriptors []fieldDescriptor) *fields {
	f := make(map[string]uint)
	t := make(map[string]string)
	for i, fieldDescriptor := range fieldDescriptors {
		f[fieldDescriptor.Name] = uint(i)
		t[fieldDescriptor.Name] = fieldDescriptor.Psql
	}
	return &fields{fields: f, psqlty: t}
}

func (f *fields) header() []string {
	n := len(f.fields)
	result := make([]string, n, n)
	for name, i := range f.fields {
		result[i] = name
	}
	return result
}

func (f *fields) record(item map[string]string) ([]interface{}, error) {
	n := len(f.fields)
	result := make([]interface{}, n, n)
	for name, value := range item {
		switch f.psqlty[name] {
		case "boolean":
			v := &pgtype.Bool{}
			if err := v.Set(value); err != nil {
				return result, err
			}
			result[f.fields[name]] = v
		case "date":
			ts, err := time.ParseInLocation("2006-01-02", value, location)
			if err != nil {
				return result, err
			}
			v := &pgtype.Date{}
			if err := v.Set(ts); err != nil {
				return result, err
			}
			result[f.fields[name]] = v
		case "integer":
			v := &pgtype.Int4{}
			if err := v.Set(value); err != nil {
				return result, err
			}
			result[f.fields[name]] = v
		case "real":
			v := &pgtype.Float4{}
			if err := v.Set(value); err != nil {
				return result, err
			}
			result[f.fields[name]] = v
		case "text", "":
			v := &pgtype.Text{}
			if err := v.Set(value); err != nil {
				return result, err
			}
			result[f.fields[name]] = v
		case "timestamp":
			ts, err := time.ParseInLocation("2006-01-02T15:04:05.9999999", value, location)
			if err != nil {
				return result, err
			}
			v := &pgtype.Timestamp{}
			if err := v.Set(ts); err != nil {
				return result, err
			}
			result[f.fields[name]] = v
		default:
			return nil, fmt.Errorf("unknown PostgreSQL type: %s", f.psqlty[name])
		}
	}
	return result, nil
}

// Implements CopyFromSource
type xmlSource struct {
	root    string
	element string
	fields  *fields
	state   int
	d       *xml.Decoder
	values  []interface{}
	err     error
}

func newXmlSource(td *tableDescriptor, d *xml.Decoder, fields *fields) xmlSource {
	return xmlSource{
		root:    td.Root,
		element: td.Element,
		fields:  fields,
		state:   startRoot,
		d:       d,
		values:  nil,
		err:     nil,
	}
}

// Next() implements pgx.CopyFromSource.
func (s *xmlSource) Next() bool {
	values, err := s.nextValues()
	if errors.Is(err, io.EOF) {
		return false
	}
	if err != nil {
		s.err = err
		return false
	}
	s.values = values
	return true
}

// Values() implements pgx.CopyFromSource.
func (s *xmlSource) Values() ([]interface{}, error) {
	return s.values, s.err
}

// Err() implements pgx.CopyFromSource.
func (s *xmlSource) Err() error {
	return s.err
}

func (s *xmlSource) nextValues() ([]interface{}, error) {
	d := s.d
	root := s.root
	element := s.element
	fields := s.fields

	item := make(map[string]string)
	var fieldName string
	var fieldValue []byte

	for {
		tok, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch s.state {
		case startRoot:
			switch t := tok.(type) {
			case xml.StartElement:
				name := xml.StartElement(t).Name.Local
				if name != root {
					return nil, fmt.Errorf("[%d] expected start of %s, got %s", s.state, root, name)
				}
				s.state = startItemOrEndRoot
			default: // ignore
			}
		case startItemOrEndRoot:
			switch t := tok.(type) {
			case xml.StartElement:
				name := xml.StartElement(t).Name.Local
				if name != element {
					return nil, fmt.Errorf("[%d] expected start of %s, got %s", s.state, element, name)
				}
				s.state = startFieldOrEndItem
			case xml.EndElement:
				name := xml.EndElement(t).Name.Local
				if name != root {
					return nil, fmt.Errorf("[%d] expected start of %s, got %s", s.state, root, name)
				}
				s.state = finished
			default: // ignore
			}
		case startFieldOrEndItem:
			switch t := tok.(type) {
			case xml.StartElement:
				name := xml.StartElement(t).Name.Local
				fieldName = name
				s.state = fieldValueOrEndField
			case xml.EndElement:
				name := xml.EndElement(t).Name.Local
				if name != element {
					return nil, fmt.Errorf("[%d] expected end of %s, got %s", s.state, element, name)
				}
				s.state = startItemOrEndRoot
				return fields.record(item)
			default: // ignore
			}
		case fieldValueOrEndField:
			switch t := tok.(type) {
			case xml.StartElement:
				name := xml.StartElement(t).Name.Local
				return nil, fmt.Errorf("[%d] expected end of %s, got start of %s", s.state, fieldName, name)
			case xml.EndElement:
				name := xml.EndElement(t).Name.Local
				if name != fieldName {
					return nil, fmt.Errorf("[%d] expected end of %s, got %s", s.state, fieldName, name)
				}
				item[fieldName] = string(fieldValue)
				fieldValue = []byte{}
				s.state = startFieldOrEndItem
			case xml.CharData:
				fieldValue = append(fieldValue, []byte(xml.CharData(t))...)
			default: // ignore
			}
		case finished:
			switch tok.(type) {
			case xml.CharData: // ignore
			default:
				return nil, fmt.Errorf("[%d] parsing finished, but got %v", s.state, tok)
			}
		}
	}
}

func createTable(tx pgx.Tx, ctx context.Context, td *tableDescriptor, force bool) error {
	// Generate "create table" statement.
	tmpl := template.Must(template.New("create").Parse(`
create unlogged table {{if .Force}}{{else}}if not exists{{end}}
{{- with .Descriptor}}"{{.Element}}" (
	{{range .Fields -}}
		"{{.Name}}"
		{{- with .Psql}} {{.}}{{else}} text{{end}}
		{{- if .Mandatory}} not null{{end}}
		{{- with .References}} references "{{.Table}}"("{{.Column}}") deferrable initially deferred{{end}},
	{{end -}}
	primary key ("{{.Primary}}")
) with (autovacuum_enabled=false);{{end}}
	`))
	var stmt bytes.Buffer
	if err := tmpl.Execute(&stmt, struct {
		Force      bool
		Descriptor *tableDescriptor
	}{force, td}); err != nil {
		return fmt.Errorf("failed to execute sql template: %w", err)
	}

	// Create the table.
	_, err := tx.Exec(ctx, stmt.String())
	if err != nil {
		return fmt.Errorf("failed to execute create table statement: %w", err)
	}
	return nil
}

func insertFromXml(f io.Reader, conn *pgx.Conn, ctx context.Context, td *tableDescriptor, force bool) (int64, error) {
	// Construct the buffered XML reader.
	const bufSize = 4096 * 1024
	br := xml.NewDecoder(bufio.NewReaderSize(f, bufSize))

	// Begin the transaction.
	tx, err := conn.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && err != pgx.ErrTxClosed {
			log.Printf("%v", err)
		}
	}()

	// Create the table.
	if err := createTable(tx, ctx, td, force); err != nil {
		return 0, err
	}

	// Copy data into the table.
	fields := newFields(td.Fields)
	s := newXmlSource(td, br, fields)
	i, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{td.Element},
		fields.header(),
		&s,
	)
	if err != nil {
		return i, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return i, err
	}
	return i, nil
}
