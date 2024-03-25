package services

import (
	"back/internal/database"
	"back/internal/models"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"math"
)

func BaseCalculations(personalCardId int64, calcDate string) ([]models.Need, error) {
	var need models.Need
	var needs []models.Need
	var countMonth float64
	var normMunitSp models.NormMunitSp
	var normMunitSps []models.NormMunitSp
	var addCond models.AdditionalCondition
	var addConds []models.AdditionalCondition
	// var personalCardNorm models.PersonalCardNorm
	var lkartNormSP models.LkartNormSP
	var lkartNormSPs []models.LkartNormSP // входящие остатки
	var personalCard models.PersonalCard
	var actualNormId, actualNormBeginDate, openingBalancesId string

	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		fmt.Println(connPostgres) // TODO: реализовать позднее
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			// получение актуальной нормы обеспечения
			row := connOracle.QueryRow(`SELECT ID, NORM_ID, BEGIN_DATE FROM parus.udo_personal_card_norm WHERE lkart_id = :1 AND END_DATE IS NULL`, personalCardId)
			err := row.Scan(&openingBalancesId, &actualNormId, &actualNormBeginDate)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			// получение данных из лицевой карточки
			row = connOracle.QueryRow(`SELECT SEX, DIVISION_ID, POSITION_ID, RANK_ID, CLIMATE FROM parus.udo_personal_card WHERE id = :1`, personalCardId)
			err = row.Scan(&personalCard.Sex, &personalCard.DivisionId, &personalCard.PositionId, &personalCard.RankId, &personalCard.Climate)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			// получение данных о имуществе из актуальной нормы
			rows, err := connOracle.Query(`SELECT MUNITION, PERIOD, QNT FROM parus.udo_norm_munitsp WHERE PRN = :1`, actualNormId)
			for rows.Next() {
				err = rows.Scan(&normMunitSp.Munition, &normMunitSp.Period, &normMunitSp.QNT)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				normMunitSps = append(normMunitSps, normMunitSp)
			}
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			// получение дополнительных условий к актуальной норме
			//rows, err = connOracle.Query(`SELECT MUNITION_ID, SEX, DIVISION_ID, POSITION_ID, RANK_ID, CLIMATE,
			//										   REPLACE_MUNITION_ID, QUANT, PERIOD
			//									FROM parus.udo_additional_condition WHERE NORM_MUNITION_ID = :1`, actualNormId)
			//for rows.Next() {
			//	err = rows.Scan(&addCond.MunitionId, &addCond.Sex, &addCond.DivisionId, &addCond.PositionId,
			//		&addCond.RankId, &addCond.Climate, &addCond.ReplaceMunitionId, &addCond.Quant, &addCond.Period)
			//	if err != nil {
			//		fmt.Println("Ошибка получения данных: ", err)
			//	}
			//	addConds = append(addConds, addCond)
			//}
			//if err != nil {
			//	fmt.Println("Ошибка получения данных: ", err)
			//}

			// row = connOracle.QueryRow(`SELECT ADD_MONTHS(to_Date('29.02.2024','DD.MM.YYYY'), 12 ) from dual`)

			// вычисляем разницу между датами (в месяцах)
			row = connOracle.QueryRow(`SELECT MONTHS_BETWEEN(to_Date(:1, 'YYYY-MM-DD'), to_Date(:2, 'YYYY-MM-DD') ) from dual`,
				calcDate, actualNormBeginDate[:10])
			err = row.Scan(&countMonth)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			fmt.Println(countMonth)
			for _, mun := range normMunitSps {
				need.CalculationByDate = calcDate
				need.Need = uint8(math.Ceil(countMonth/float64(mun.Period))) * mun.QNT
				need.Counterparty = mun.Munition
				needs = append(needs, need)
				need = models.Need{}
			}

			// получение дополнительных условий к актуальной норме
			rows, err = connOracle.Query(`SELECT a.MUNITION_ID, a.REPLACE_MUNITION_ID, a.QUANT, a.PERIOD
									FROM udo_personal_card c,
										 udo_personal_card_norm n,
										 udo_additional_condition a
									WHERE c.ID = :1
									  AND n.NORM_ID = a.NORM_MUNITION_ID
									  AND (c.DIVISION_ID = a.DIVISION_ID OR a.DIVISION_ID IS NULL)
									  AND (c.POSITION_ID = a.POSITION_ID  OR  a.POSITION_ID IS NULL)
									  AND (c.RANK_ID = a.RANK_ID  OR  a.RANK_ID IS NULL)
									  AND (c.CLIMATE = a.CLIMATE  OR  a.CLIMATE IS NULL)`, personalCardId)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}
			defer func(rows *sql.Rows) {
				err = rows.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(rows)
			for rows.Next() {
				err = rows.Scan(&addCond.MunitionId, &addCond.ReplaceMunitionId, &addCond.Quant, &addCond.Period)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				addConds = append(addConds, addCond)
			}

			// формирование потребностей с учетом дополнительных условий
			for _, mun := range addConds {
				if !mun.ReplaceMunitionId.Valid {
					need.CalculationByDate = "New munition"
					need.Need = uint8(math.Ceil(countMonth/float64(mun.Period))) * mun.Quant
					needs = append(needs, need)
					need = models.Need{}
				} else {
					// TODO: подумать???
					for i, replaceMun := range needs {
						if mun.ReplaceMunitionId.String == replaceMun.Counterparty {
							needs[i].CalculationByDate = "Replace"
							needs[i].Need = uint8(math.Ceil(countMonth/float64(mun.Period))) * mun.Quant
						}
					}
				}
			}

			// получение входящих остатков
			rows, err = connOracle.Query(`SELECT * FROM parus.udo_lkart_normsp WHERE prn = :1`, openingBalancesId)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}
			defer rows.Close()
			for rows.Next() {
				err = rows.Scan(&lkartNormSP.RN, &lkartNormSP.PRN, &lkartNormSP.Munition, &lkartNormSP.Period,
					&lkartNormSP.InQntNeed, &lkartNormSP.InQntNosPis, &lkartNormSP.BeginDate, &lkartNormSP.EndDate,
					&lkartNormSP.StopDate, &lkartNormSP.InQntTatters, &lkartNormSP.Note)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				lkartNormSPs = append(lkartNormSPs, lkartNormSP)
			}
			fmt.Println(lkartNormSPs)
		}
	}

	/*
		// вычисляем разницу между датами (в месяцах)
			t1, _ := time.Parse("2006-01-02", calcDate)
			t2, _ := time.Parse("2006-01-02", actualNormBeginDate[:10])
			countMonth := math.Ceil(t1.Sub(t2).Hours() / 24 / 30)
	*/

	fmt.Println(len(needs))
	return needs, nil
}
