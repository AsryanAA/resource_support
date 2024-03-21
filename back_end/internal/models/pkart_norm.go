package models

import "database/sql"

type PersonalCardNorm struct {
	Id        string         `json:"id"`
	LKartId   string         `json:"l_kart_id"`
	NormId    string         `json:"norm_id"`
	BeginDate string         `json:"begin_date"`
	EndDate   sql.NullString `json:"end_date"`
}
