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

func CreatePKartNorm(pKartNorm models.PersonalCardNorm) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_pkart_norm
				VALUES (:1, :2, :3, :4, :5)`, pKartNorm.Id, pKartNorm.LKartId, pKartNorm.NormId, pKartNorm.BeginDate, pKartNorm.EndDate)
			return err
		}
	}
	return nil
}

func ReadPKartNorms() ([]models.PersonalCardNorm, error) {
	var pKartNorms []models.PersonalCardNorm
	var pKartNorm models.PersonalCardNorm
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_pkart_norm`)
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
			pKartNorms = append(pKartNorms, pKartNorm)
		}
		return pKartNorms, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_pkart_norm`)
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
				err = rows.Scan(&pKartNorm.Id, &pKartNorm.LKartId, &pKartNorm.NormId, &pKartNorm.BeginDate, &pKartNorm.EndDate)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				pKartNorms = append(pKartNorms, pKartNorm)
			}
		}
		return pKartNorms, nil
	}
}
