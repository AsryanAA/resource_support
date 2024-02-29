package models

import "database/sql"

type Munition struct {
	RN            string `json:"rn"`
	CRN           string `json:"crn"`
	Version       string `json:"version"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	DicNomNs      string `json:"dicnomns"`
	NomModif      string `json:"nommodif"`
	PackModif     string `json:"packmodif"`
	KompSum       string `json:"kompsum"`
	Cloth         string `json:"cloth"`
	IgnoreSupply  string `json:"ignoresupply"`
	Sex           string `json:"sex"`
	UseFaktHeihth string `json:"usefaktheihth"`
	BegstepHeihth string `json:"begstepheihth"`
}

type MunitionSew struct {
	RN          string         `json:"rn"`
	PRN         string         `json:"prn"`
	Munition    string         `json:"munition"`
	MunitionMod sql.NullString `json:"munition_mod"`
	ClothQnt    string         `json:"cloth_qnt"`
}

type MunitionMod struct {
	RN          string `json:"rn"`
	PRN         string `json:"prn"`
	NomModif    string `json:"nom_modif"`
	PackModifSp string `json:"pack_modif_sp"`
	Code        string `json:"code"`
}

type MunitionSP struct {
	Code   string  `json:"code"`
	Qnt    float64 `json:"qnt"`
	Period float64 `json:"period"`
}
