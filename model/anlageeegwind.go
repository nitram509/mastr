// Code generated by go mastr/model generator; DO NOT EDIT.

package main

import (
	"encoding/xml"
)

// AnlagenEegWind
type AnlagenEegWind struct {
	XMLName        xml.Name        `xml:"AnlagenEegWind"`
	AnlagenEegWind []AnlageEegWind `xml:"AnlageEegWind"`
}

type AnlageEegWind struct {
	XMLName                                       xml.Name `xml:"AnlageEegWind"`
	Registrierungsdatum                           string   `xml:"Registrierungsdatum"`
	DatumLetzteAktualisierung                     string   `xml:"DatumLetzteAktualisierung"`
	EegInbetriebnahmedatum                        string   `xml:"EegInbetriebnahmedatum"`
	EegMaStRNummer                                string   `xml:"EegMaStRNummer"`
	AnlagenkennzifferAnlagenregister              string   `xml:"AnlagenkennzifferAnlagenregister"`
	AnlagenkennzifferAnlagenregister_nv           string   `xml:"AnlagenkennzifferAnlagenregister_nv"`
	AnlagenschluesselEeg                          string   `xml:"AnlagenschluesselEeg"`
	PrototypAnlage                                string   `xml:"PrototypAnlage"`
	PilotAnlage                                   string   `xml:"PilotAnlage"`
	InstallierteLeistung                          float32  `xml:"InstallierteLeistung"`
	VerhaeltnisErtragsschaetzungReferenzertrag    float32  `xml:"VerhaeltnisErtragsschaetzungReferenzertrag"`
	VerhaeltnisErtragsschaetzungReferenzertrag_nv string   `xml:"VerhaeltnisErtragsschaetzungReferenzertrag_nv"`
	VerhaeltnisReferenzertragErtrag5Jahre         float32  `xml:"VerhaeltnisReferenzertragErtrag5Jahre"`
	VerhaeltnisReferenzertragErtrag5Jahre_nv      string   `xml:"VerhaeltnisReferenzertragErtrag5Jahre_nv"`
	VerhaeltnisReferenzertragErtrag10Jahre        float32  `xml:"VerhaeltnisReferenzertragErtrag10Jahre"`
	VerhaeltnisReferenzertragErtrag10Jahre_nv     string   `xml:"VerhaeltnisReferenzertragErtrag10Jahre_nv"`
	VerhaeltnisReferenzertragErtrag15Jahre        float32  `xml:"VerhaeltnisReferenzertragErtrag15Jahre"`
	VerhaeltnisReferenzertragErtrag15Jahre_nv     string   `xml:"VerhaeltnisReferenzertragErtrag15Jahre_nv"`
	AusschreibungZuschlag                         string   `xml:"AusschreibungZuschlag"`
	Zuschlagsnummer                               string   `xml:"Zuschlagsnummer"`
	AnlageBetriebsstatus                          uint     `xml:"AnlageBetriebsstatus"` // referenziert Katalogwert
	VerknuepfteEinheitenMaStRNummern              string   `xml:"VerknuepfteEinheitenMaStRNummern"`
}
