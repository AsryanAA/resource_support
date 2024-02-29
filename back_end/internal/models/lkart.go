package models

type LKart struct {
	RN              string `json:"rn"`
	NRN             string `json:"n_rn"`
	Company         string `json:"company"`
	NCompany        string `json:"n_company"`
	NCRN            string `json:"n_crn"`
	StabPref        string `json:"stab_pref"`
	StabNumb        string `json:"stab_numb"`
	NOwner          string `json:"n_owner"`
	SOwner          string `json:"s_owner"`
	NPersAgent      string `json:"n_pers_agent"`
	SAgnCode        string `json:"s_agn_code"`
	SAgnName        string `json:"s_agn_name"`
	DJobBeginDate   string `json:"d_job_begin_date"`
	DLastJobBegDate string `json:"d_last_job_beg_date"`
	NSex            string `json:"n_sex"`
	SPersNote       string `json:"s_pers_note"`
	NLastStatus     string `json:"n_last_status"`
	SLastNormMunit  string `json:"s_last_norm_munit"`
	SEmpPost        string `json:"s_emp_post"`
	SDep            string `json:"s_dep"`
	SLastRank       string `json:"s_last_rank"`
	DDismissDate    string `json:"d_dismiss_date"`
	SBabyInf        string `json:"s_baby_inf"`
}