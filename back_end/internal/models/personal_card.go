package models

type PersonalCard struct {
	/*
		id
		sex
		phone_number
		lkart_id (ссылка на родительскую карточку lkart (вытаскиваем ФИО, дата приема на работу и т.д.))
		position_id
		division_id
		rank_id

		CREATE TABLE UDO_PERSONAL_CARDS (
			ID NUMBER(17, 0) NOT NULL,
			SEX NUMBER(1, 0),
			PHONE_NUMBER NUMBER(10, 0),
			LKART_ID NUMBER(17, 0),
			POSITION_ID NUMBER(17, 0),
			DIVISION_ID NUMBER(17, 0),
			RANK_ID NUMBER(17, 0)
		)
	*/

	Id          string `json:"id"`
	Sex         string `json:"sex"`
	PhoneNumber string `json:"phone_number"`
	LKartId     string `json:"l_kart_id"`
	PositionId  string `json:"position_id"`
	DivisionId  string `json:"division_id"`
	RankId      string `json:"rank_id"`
}
