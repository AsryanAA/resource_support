package models

type PKart struct {
	Id           string `json:"id"`
	FIO          string `json:"fio"`
	JobBeginDate string `json:"job_begin_date"`
	Sex          string `json:"sex"`
	DivisionId   string `json:"division_id"`
	PositionId   string `json:"position_id"`
	RankId       string `json:"rank_id"`
	Climate      string `json:"climate"`
}
