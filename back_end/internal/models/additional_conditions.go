package models

import "database/sql"

type AdditionalCondition struct {
	Id                string         `json:"id"`
	NormMunitionId    string         `json:"norm_munition_id"`
	MunitionId        string         `json:"munition_id"`
	Description       string         `json:"description"`
	RankId            sql.NullString `json:"rank_id"`
	PositionId        sql.NullString `json:"position_id"`
	DivisionId        sql.NullString `json:"division_id"`
	Sex               string         `json:"sex"`
	Climate           sql.NullString `json:"climate"`
	ReplaceMunitionId sql.NullString `json:"replace_munition_id"`
	Quant             uint8          `json:"quant"`
	Period            uint8          `json:"period"`
}
