package models

import "database/sql"

type NormMunit struct {
	RN         string `json:"rn"`
	CRN        string `json:"crn"`
	Version    string `json:"version"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	NormType   string `json:"norm_type"`
	InventProp string `json:"invent_prop"` // только  Postgres
}

type NormMunitSp struct {
	RN          string         `json:"rn"`
	PRN         string         `json:"prn"`
	CRN         string         `json:"crn"`
	Version     string         `json:"version"`
	Munition    string         `json:"munition"`
	Period      uint8          `json:"period"`
	QNT         uint8          `json:"QNT"`
	PackACS     sql.NullString `json:"pack_acs"`
	ByDateFirst string         `json:"by_date_first"`
	ByDateFact  string         `json:"by_date_fact"`
	ForPay      string         `json:"for_pay"`
	ACSOneKompl string         `json:"acs_one_kompl"`
}

type NormMunitSpAlt struct {
	RN         string `json:"rn"`
	PRN        string `json:"prn"`
	CRN        string `json:"crn"`
	Version    string `json:"version"`
	GroupMunit string `json:"group_munit"`
}
