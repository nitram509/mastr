// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// AnlagenEegWasser
type AnlagenEegWasser struct {
	XMLName          xml.Name          `xml:"AnlagenEegWasser"`
	AnlagenEegWasser []AnlageEegWasser `xml:"AnlageEegWasser"`
}

type AnlageEegWasser struct {
	XMLName                             xml.Name `xml:"AnlageEegWasser"`
	Registrierungsdatum                 string   `xml:"Registrierungsdatum"`
	DatumLetzteAktualisierung           string   `xml:"DatumLetzteAktualisierung"`
	EegInbetriebnahmedatum              string   `xml:"EegInbetriebnahmedatum"`
	EegMaStRNummer                      string   `xml:"EegMaStRNummer"`
	AnlagenschluesselEeg                string   `xml:"AnlagenschluesselEeg"`
	AnlagenkennzifferAnlagenregister    string   `xml:"AnlagenkennzifferAnlagenregister"`
	AnlagenkennzifferAnlagenregister_nv string   `xml:"AnlagenkennzifferAnlagenregister_nv"`
	InstallierteLeistung                float32  `xml:"InstallierteLeistung"`
	AnlageBetriebsstatus                uint     `xml:"AnlageBetriebsstatus"` // referenziert Katalogwert
	ErtuechtigungIds                    string   `xml:"ErtuechtigungIds"`
	VerknuepfteEinheitenMaStRNummern    string   `xml:"VerknuepfteEinheitenMaStRNummern"`
}
