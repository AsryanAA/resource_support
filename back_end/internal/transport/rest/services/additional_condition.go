package services

import (
	"back/internal/database"
	"back/internal/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func CreateAdditionalCondition(addCond models.AdditionalCondition) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_additional_condition
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, addCond.Id, addCond.NormId, addCond.MunitionId, addCond.RankId, addCond.PositionId,
			addCond.DivisionId, addCond.Sex)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_additional_condition
				VALUES (:1, :2, :3, :4, :5, :6, :7)`, addCond.Id, addCond.NormId, addCond.MunitionId, addCond.RankId, addCond.PositionId,
				addCond.DivisionId, addCond.Sex)
			return err
		}
	}
	return nil
}

func ReadAdditionalConditions() ([]models.AdditionalCondition, error) {
	var addConds []models.AdditionalCondition
	var addCond models.AdditionalCondition
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_additional_condition`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan()
			if err != nil {
				return nil, err
			}
			addConds = append(addConds, addCond)
		}
		return addConds, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_additional_condition`)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
				os.Exit(1)
			}
			defer func(rows *sql.Rows) {
				err = rows.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(rows)
			for rows.Next() {
				err = rows.Scan(
					&addCond.Id, &addCond.NormId, &addCond.MunitionId, &addCond.RankId, &addCond.PositionId,
					&addCond.DivisionId, &addCond.Sex)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				addConds = append(addConds, addCond)
			}
		}
		return addConds, nil
	}
}

func UpdateAdditionalCondition(addCond models.AdditionalCondition) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres // TODO: реализовать позже
		return nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_additional_condition SET norm_id=:2, munition_id=:3,
                                      rank_id=:4, position_id=:5, division_id=:6, sex=:7 WHERE id=:1`,
				addCond.Id, addCond.NormId, addCond.MunitionId, addCond.RankId, addCond.PositionId,
				addCond.DivisionId, addCond.Sex)
			return err
		}
	}
	return nil
}

func DeleteAdditionalCondition(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_additional_condition WHERE id=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_additional_condition WHERE id=:1`, id)
			return err
		}
	}
	return nil
}
