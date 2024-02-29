package services

import (
	"back/internal/database"
	"back/internal/models"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"math"
	"time"
)

func RequirementCalculations(lkartId int64, calcDate string) ([]models.Need, error) {
	var lkart models.LKart
	var lkartNorm models.LkartNorm
	var lkartNormSP models.LkartNormSP
	var lkartNormSPs []models.LkartNormSP
	var need models.Need
	var needs []models.Need
	var munitionSP models.MunitionSP
	var doc models.LkartDocSP
	var docs []models.LkartDocSP
	// var lkartNormSPS []models.LkartNormSP
	// var lkartDocs []models.LkartDoc

	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		fmt.Println(connPostgres) // TODO: реализовать позднее
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			// получение актуальной нормы обеспечения
			row := connOracle.QueryRow(`SELECT * FROM parus.udo_lkart_norm WHERE prn = :1 AND DATE_END IS NULL`, lkartId)
			err := row.Scan(&lkartNorm.RN, &lkartNorm.PRN, &lkartNorm.Status, &lkartNorm.BeginDate, &lkartNorm.EndDate,
				&lkartNorm.NormMunit, &lkartNorm.IsProlong, &lkartNorm.Note, &lkartNorm.CRN, &lkartNorm.Company)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			// получение контрагента
			row = connOracle.QueryRow(`SELECT SAGNCODE, SAGNNAME FROM parus.udov_lkart WHERE rn = :1`, lkartId)
			err = row.Scan(&lkart.SAgnCode, &lkart.SAgnName)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			if lkartNorm.NormMunit == "" {
				fmt.Println("Норма обеспечения не найдена")
			}

			// получаем сроки носки исходя из актуальной нормы обеспечения
			rows, err := connOracle.Query(`SELECT m.code, unmsp.qnt, unmsp.period FROM parus.udo_norm_munit unm
												 JOIN parus.udo_norm_munitsp unmsp ON unm.rn = unmsp.prn
												 JOIN parus.udo_munition m ON unmsp.munition = m.rn
												 WHERE unm.rn = :1`, lkartNorm.NormMunit)
			defer func(rows *sql.Rows) {
				err = rows.Close()
				if err != nil {
					fmt.Println("Ошибка при закрытии", err)
				}
			}(rows)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}
			var munitionsSP []models.MunitionSP
			for rows.Next() {
				err = rows.Scan(&munitionSP.Code, &munitionSP.Qnt, &munitionSP.Period)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				munitionsSP = append(munitionsSP, munitionSP)
			}

			// формируем предварительный рассчет потребности
			layout := "2006-01-02"
			t, err := time.Parse(layout, lkartNorm.BeginDate[:10])
			t1, err := time.Parse(layout, calcDate)
			difference := differenceInMonths(t, t1)
			for _, msp := range munitionsSP {
				need.CalculationByDate = calcDate
				need.CounterpartyName = lkart.SAgnName
				need.Counterparty = lkart.SAgnCode
				need.CollateralNorm = msp.Code
				need.Property = msp.Code
				if difference/msp.Period > 0 {
					// need.Need += (difference/msp.Period + 1.0) * msp.Qnt
					need.Need = math.Ceil((difference / msp.Period) * msp.Qnt)
				}
				need.WornOut = 0
				need.NotWrittenOff = 0
				need.StartDate = t.String() // TODO: подумать!!!
				need.EndDate = (t.AddDate(0, int(msp.Period), 0)).String()
				needs = append(needs, need)
				need = models.Need{} // обнуляем объект
			}

			// находим документы по конртагенту
			date, err := time.Parse("2006-01-02", calcDate)
			if err != nil {
				log.Fatal(err)
			}
			rows, err = connOracle.Query(`SELECT m.CODE, lkdsp.QNT_DOC, lkdsp.QNT_FACT
														--lkdsp.QNT_CHDOC, lkdsp.QNT_CHFACT
														FROM parus.udo_lkart_doc lkd
														JOIN parus.udo_lkart_docsp lkdsp ON lkd.rn = lkdsp.prn
														JOIN parus.udo_munition m ON lkdsp.munition_rn = m.rn
														WHERE lkd.lkart_rn = :1 AND lkd.doc_date < :2`, lkartId, date)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", rows)
			}
			for rows.Next() {
				err = rows.Scan(&doc.MunitionRN, &doc.QntDoc, &doc.QntFact)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				docs = append(docs, doc)
			}

			if len(docs) > 0 {
				for i, n := range needs {
					for _, d := range docs {
						if n.Property == d.MunitionRN {
							// предмет, сколько требуется? сколько изношено? сколько не списано? кол по документу, кол фактически
							// fmt.Println("Предмет", n.Property, n.Need, n.WornOut, n.NotWrittenOff, d.QntDoc, d.QntFact)
							// n.Need = math.Ceil(n.Need) - d.QntFact
							// n.NotWrittenOff = d.QntFact
							// needs[i] = n
							_ = i
						}
					}
				}
			}

			rows, err = connOracle.Query(`SELECT lknsp.*, m.CODE FROM parus.UDO_LKART_NORMSP lknsp
         										JOIN parus.UDO_MUNITION m ON lknsp.munition = m.rn
         										WHERE prn = :1`, lkartNorm.RN)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", rows)
			}
			for rows.Next() {
				err = rows.Scan(&lkartNormSP.RN, &lkartNormSP.PRN, &lkartNormSP.Munition, &lkartNormSP.Period,
					&lkartNormSP.InQntNeed, &lkartNormSP.InQntNosPis, &lkartNormSP.BeginDate, &lkartNormSP.EndDate,
					&lkartNormSP.StopDate, &lkartNormSP.InQntTatters, &lkartNormSP.Note, &lkartNormSP.MunitionCode)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				lkartNormSPs = append(lkartNormSPs, lkartNormSP)
			}

			if len(lkartNormSPs) > 0 {
				for _, n := range needs {
					for _, l := range lkartNormSPs {
						// fmt.Println("1", l.Munition, "2", l.MunitionCode, "3", l.MunitionName, "4", n.Property)
						if n.Property == l.MunitionCode && l.EndDate < calcDate {
							fmt.Println("1", n)
							fmt.Println("2", l)
						}
					}
				}
			}
		}

		// получаем документы лицевой карточки
		//rows, err := connOracle.Query(`SELECT * FROM udo_lkart_doc WHERE lkart_rn = :1`, lkartId) // TODO: доделать все столбцы
		//if err != nil {
		//	fmt.Println("Ошибка получения данных: ", err)
		//}
		//defer rows.Close()
		//for rows.Next() {
		//	var lkartDoc models.LkartDoc
		//	err = rows.Scan(&lkartDoc.RN, &lkartDoc.DocType, &lkartDoc.DocDate, &lkartDoc.LkartNorm, &lkartDoc.OperType,
		//		&lkartDoc.WorkDate, &lkartDoc.Store)
		//	if err != nil {
		//		fmt.Println("Ошибка получения данных: ", err)
		//	}
		//	lkartDocs = append(lkartDocs, lkartDoc)
		//}
		//fmt.Println(lkartDocs)
		//
		// получаем входящие остатки
		//if lkartNorm.RN != "" {
		//	rows, err := connOracle.Query(`SELECT * FROM parus.udo_lkart_normsp WHERE prn = :1`, lkartNorm.RN)
		//	if err != nil {
		//		fmt.Println("Ошибка получения данных: ", err)
		//	}
		//	defer rows.Close()
		//	for rows.Next() {
		//		var lkartNormSP models.LkartNormSP
		//		err = rows.Scan(&lkartNormSP.RN, &lkartNormSP.PRN, &lkartNormSP.Munition, &lkartNormSP.Period,
		//			&lkartNormSP.InQntNeed, &lkartNormSP.InQntNosPis, &lkartNormSP.BeginDate, &lkartNormSP.EndDate,
		//			&lkartNormSP.StopDate, &lkartNormSP.InQntTatters, &lkartNormSP.Note)
		//		if err != nil {
		//			fmt.Println("Ошибка получения данных: ", err)
		//		}
		//		lkartNormSPS = append(lkartNormSPS, lkartNormSP)
		//	}
		//}
	}

	return needs, nil
}

func millisecondsToMonths(milliseconds int64) float64 {
	// Определяем количество миллисекунд в месяце
	millisecondsInMonth := float64(30 * 24 * 60 * 60 * 1000)

	// Выполняем преобразование
	months := float64(milliseconds) / millisecondsInMonth

	return months
}

func differenceInMonths(start, end time.Time) float64 {
	// Определяем длительность между датами
	duration := end.Sub(start)

	// Преобразуем длительность в месяцы (округляем вниз)
	months := float64(duration.Hours() / 24 / 30)

	return months
}
