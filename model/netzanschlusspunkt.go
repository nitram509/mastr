// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// Netzanschlusspunkte
type Netzanschlusspunkte struct {
	XMLName             xml.Name             `xml:"Netzanschlusspunkte"`
	Netzanschlusspunkte []Netzanschlusspunkt `xml:"Netzanschlusspunkt"`
}

type Netzanschlusspunkt struct {
	XMLName                                xml.Name `xml:"Netzanschlusspunkt"`
	NetzanschlusspunktMastrNummer          string   `xml:"NetzanschlusspunktMastrNummer"`
	NetzanschlusspunktBezeichnung          string   `xml:"NetzanschlusspunktBezeichnung"`
	LetzteAenderung                        string   `xml:"LetzteAenderung"`
	LokationMaStRNummer                    string   `xml:"LokationMaStRNummer"` // referenziert Lokation
	NameDerTechnischenLokation             string   `xml:"NameDerTechnischenLokation"`
	Lokationtyp                            uint     `xml:"Lokationtyp"`
	Messlokation                           string   `xml:"Messlokation"`
	Spannungsebene                         uint     `xml:"Spannungsebene"` // referenziert Katalogwert
	Nettoengpassleistung                   float32  `xml:"Nettoengpassleistung"`
	BilanzierungsgebietNetzanschlusspunkId uint     `xml:"BilanzierungsgebietNetzanschlusspunkId"` // referenziert Bilanzierungsgebiet
	Netzanschlusskapazitaet                float32  `xml:"Netzanschlusskapazitaet"`
	Marktgebiet                            uint     `xml:"Marktgebiet"` // referenziert Katalogwert
	MaximaleEinspeiseleistung              float32  `xml:"MaximaleEinspeiseleistung"`
	MaximaleAusspeiseleistung              float32  `xml:"MaximaleAusspeiseleistung"`
	Gasqualitaet                           uint     `xml:"Gasqualitaet"`                // referenziert Katalogwert
	RegelzoneNetzanschlusspunkt            uint     `xml:"RegelzoneNetzanschlusspunkt"` // referenziert Katalogwert
	NetzMaStRNummer                        string   `xml:"NetzMaStRNummer"`             // referenziert Netz

}