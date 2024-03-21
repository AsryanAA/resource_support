package services

import (
	"back/internal/database"
	"back/internal/models"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"math"
	"time"
)

func BaseCalculations(personalCardId int64, calcDate string) ([]models.Need, error) {
	var need models.Need
	var needs []models.Need
	var normMunitSp models.NormMunitSp
	var normMunitSps []models.NormMunitSp
	var addCond models.AdditionalCondition
	var addConds []models.AdditionalCondition
	// var personalCardNorm models.PersonalCardNorm
	var personalCard models.PersonalCard
	var actualNormId, actualNormBeginDate string

	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		fmt.Println(connPostgres) // TODO: реализовать позднее
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			// получение актуальной нормы обеспечения
			row := connOracle.QueryRow(`SELECT NORM_ID, BEGIN_DATE FROM parus.udo_personal_card_norm WHERE lkart_id = :1 AND END_DATE IS NULL`, personalCardId)
			err := row.Scan(&actualNormId, &actualNormBeginDate)
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
			rows, err = connOracle.Query(`SELECT MUNITION_ID, SEX, DIVISION_ID, POSITION_ID, RANK_ID, CLIMATE,
       												   REPLACE_MUNITION_ID, QUANT, PERIOD
												FROM parus.udo_additional_condition WHERE NORM_MUNITION_ID = :1`, actualNormId)
			for rows.Next() {
				err = rows.Scan(&addCond.MunitionId, &addCond.Sex, &addCond.DivisionId, &addCond.PositionId,
					&addCond.RankId, &addCond.Climate, &addCond.ReplaceMunitionId, &addCond.Quant, &addCond.Period)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				addConds = append(addConds, addCond)
			}
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}

			// row = connOracle.QueryRow(`SELECT ADD_MONTHS(to_Date('29.02.2024','DD.MM.YYYY'), 12 ) from dual`)
			fmt.Println(calcDate, actualNormBeginDate)
			row = connOracle.QueryRow(`SELECT MONTHS_BETWEEN(to_Date(:1,'DD.MM.YYYY'), to_Date(:2,'DD.MM.YYYY') ) from dual`,
				calcDate, actualNormBeginDate[:10])
			var countMonth string
			err = row.Scan(&countMonth)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}
			fmt.Println(countMonth)
		}
	}

	// вычисляем разницу между датами (в месяцах)
	t1, _ := time.Parse("2006-01-02", calcDate)
	t2, _ := time.Parse("2006-01-02", actualNormBeginDate[:10])
	countMonth := math.Ceil(t1.Sub(t2).Hours() / 24 / 30)

	for _, mun := range normMunitSps {
		need.CalculationByDate = calcDate
		need.Need = uint8(math.Ceil(countMonth/float64(mun.Period))) * mun.QNT
		needs = append(needs, need)
	}

	fmt.Println(len(needs))

	return needs, nil
}
