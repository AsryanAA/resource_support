package models

type LkartDoc struct {
	RN        string `json:"rn"`
	DocType   string `json:"doc_type"`
	DocDate   string `json:"doc_date"`
	LkartNorm string `json:"lkart_norm"`
	OperType  string `json:"oper_type"`
	WorkDate  string `json:"work_date"`
	Store     string `json:"store"`
}

type LkartDocSP struct {
	// RN string `json:"rn"`
	// Company    string `json:"company"`
	// PRN        string `json:"prn"`
	MunitionRN string  `json:"munition_rn"`
	QntDoc     float64 `json:"qnt_doc"`
	QntFact    float64 `json:"qnt_fact"`
}
