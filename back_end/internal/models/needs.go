package models

type Need struct {
	CalculationByDate string  `json:"calculation_by_date"` // расчет на дату
	Counterparty      string  `json:"counterparty"`        // контрагент
	CounterpartyName  string  `json:"counterparty_name"`   // наименование контрагента
	CollateralNorm    string  `json:"collateral_norm"`     // норма обеспечения
	Property          string  `json:"property"`            // имущество
	Need              float64 `json:"need"`                // потребность
	WornOut           float64 `json:"worn_out"`            // изношено
	NotWrittenOff     float64 `json:"not_written_off"`     // не списано
	StartDate         string  `json:"start_date"`          // дата начала
	EndDate           string  `json:"end_date"`            // дата окончания
}
