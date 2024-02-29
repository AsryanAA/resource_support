package models

import "database/sql"

type LkartNorm struct {
	RN        string         `json:"rn"`
	PRN       string         `json:"prn"`
	Status    string         `json:"status"`
	BeginDate string         `json:"begin_date"`
	EndDate   sql.NullString `json:"end_date"`
	NormMunit string         `json:"norm_munit"`
	IsProlong string         `json:"is_prolong"`
	Note      sql.NullString `json:"note"`
	CRN       string         `json:"crn"`
	Company   string         `json:"company"`
}

type LkartNormSP struct {
	RN           string         `json:"rn"`
	PRN          string         `json:"prn"`
	Munition     string         `json:"munition"`
	Period       string         `json:"period"`
	InQntNeed    string         `json:"in_qnt_need"`
	InQntNosPis  string         `json:"in_qnt_nos_pis"`
	BeginDate    string         `json:"begin_date"`
	EndDate      string         `json:"end_date"`
	StopDate     sql.NullString `json:"stop_date"`
	InQntTatters string         `json:"in_qnt_tatters"`
	Note         sql.NullString `json:"note"`
	MunitionName string         `json:"munition_name"`
	MunitionCode string         `json:"munition_code"`
}
