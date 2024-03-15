package models

type AdditionalCondition struct {
	Id                string `json:"id"`
	NormMunitionId    string `json:"norm_munition_id"`
	MunitionId        string `json:"munition_id"`
	Description       string `json:"description"`
	RankId            string `json:"rank_id"`
	PositionId        string `json:"position_id"`
	DivisionId        string `json:"division_id"`
	Sex               string `json:"sex"`
	Climate           string `json:"climate"`
	ReplaceMunitionId string `json:"replace_munition_id"`
}
