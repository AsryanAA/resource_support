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

func CreatePKart(pKart models.PKart) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_pkart
				VALUES (:1, :2, :3, :4, :5, :6, :7)`, pKart.Id, pKart.FIO, pKart.JobBeginDate, pKart.Sex, pKart.DivisionId,
				pKart.PositionId, pKart.RankId, pKart.Climate)
			return err
		}
	}
	return nil
}

func ReadPKarts() ([]models.PKart, error) {
	var pKarts []models.PKart
	var pKart models.PKart
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_pkart`)
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
			pKarts = append(pKarts, pKart)
		}
		return pKarts, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_pkart`)
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
				err = rows.Scan(&pKart.Id, &pKart.FIO, &pKart.JobBeginDate, &pKart.Sex, &pKart.DivisionId, &pKart.PositionId,
					&pKart.RankId, &pKart.Climate)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				pKarts = append(pKarts, pKart)
			}
		}
		return pKarts, nil
	}
}
