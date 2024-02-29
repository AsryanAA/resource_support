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

func CreateMunitionSew(munitionSew models.MunitionSew) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres
		return nil // TODO: реализовать позже
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_munition_sew VALUES (:1, :2, :3, :4, :5)`,
				munitionSew.RN, munitionSew.PRN, munitionSew.Munition, munitionSew.MunitionMod, munitionSew.ClothQnt)
			return err
		}
	}
	return nil
}

func ReadMunitionsSew() ([]models.MunitionSew, error) {
	var munitionsSew []models.MunitionSew
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_munition_sew`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			var munitionSew models.MunitionSew
			err = rows.Scan()
			if err != nil {
				return nil, err
			}
			munitionsSew = append(munitionsSew, munitionSew)
		}
		return munitionsSew, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_munition_sew`)
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
				var munitionSew models.MunitionSew
				err = rows.Scan(
					&munitionSew.RN, &munitionSew.PRN, &munitionSew.Munition, &munitionSew.MunitionMod, &munitionSew.ClothQnt)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				munitionsSew = append(munitionsSew, munitionSew)
			}
		}
		return munitionsSew, nil
	}
}

func UpdateMunitionSew(munitionSew models.MunitionSew) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres // TODO: реализовать позже
		return nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_munition_sew SET "VERSION"=:2, crn=:3, code=:4, "NAME"=:5
                              				 WHERE rn=:1`,
				munitionSew.RN, munitionSew.PRN, munitionSew.Munition, munitionSew.MunitionMod, munitionSew.ClothQnt)
			return err
		}
	}
	return nil
}

func DeleteMunitionSew(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_munition_sew WHERE rn=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_munition_sew WHERE rn=:1`, id)
			return err
		}
	}
	return nil
}
