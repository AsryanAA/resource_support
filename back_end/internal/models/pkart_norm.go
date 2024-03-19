package models

type PersonalCardNorm struct {
	Id        string `json:"id"`
	LKartId   string `json:"l_kart_id"`
	NormId    string `json:"norm_id"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}
