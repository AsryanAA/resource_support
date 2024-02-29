package models

type PackModif struct {
	Code    string `json:"code"`
	Version string `json:"version"`
	CRN     string `json:"crn"`
	Name    string `json:"name"`
	RN      string `json:"rn"`
}

type PackModifSp struct {
	RN     string `json:"rn"`
	PRN    string `json:"prn"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Girth  string `json:"girth"`
	Neck   string `json:"neck"`
	Height string `json:"height"`
}
