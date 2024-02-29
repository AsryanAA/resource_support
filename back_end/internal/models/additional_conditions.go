package models

type AdditionalCondition struct {
	Id         string `json:"id"`
	NormId     string `json:"norm_id"`
	MunitionId string `json:"munition_id"`
	RankId     string `json:"rank_id"`
	PositionId string `json:"position_id"`
	DivisionId string `json:"division_id"`
	Sex        string `json:"sex"`
}

/*
CREATE TABLE UDO_ADDITIONAL_CONDITIONS (
    ID NUMBER(17, 0) NOT NULL,
    NORM_ID NUMBER(17, 0)NOT NULL,
    MUNTION_ID NUMBER(17, 0),
    RANK_ID NUMBER(17, 0),
    POSITION_ID NUMBER(17, 0),
    DIVISION_ID NUMBER(17, 0),
    SEX NUMBER(1, 0)
)
*/
