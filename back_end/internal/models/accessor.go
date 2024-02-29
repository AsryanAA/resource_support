package models

type Accessor struct {
	RN       string `json:"rn"`
	CRN      string `json:"crn"`
	Version  string `json:"version"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	DicNomNs string `json:"dicnomns"`
	NomModif string `json:"nommodif"`
}