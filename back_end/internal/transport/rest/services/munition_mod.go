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

func CreateMunitionMod(munitionMod models.MunitionMod) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres
		return nil // TODO: реализовать позже
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_munition_mod VALUES (:1, :2, :3, :4, :5)`,
				munitionMod.RN, munitionMod.PRN, munitionMod.NomModif, munitionMod.PackModifSp, munitionMod.Code)
			return err
		}
	}
	return nil
}

func ReadMunitionsMod() ([]models.MunitionMod, error) {
	var munitionsMod []models.MunitionMod
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_munition_mod`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			var munitionMod models.MunitionMod
			err = rows.Scan()
			if err != nil {
				return nil, err
			}
			munitionsMod = append(munitionsMod, munitionMod)
		}
		return munitionsMod, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_munition_mod`)
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
				var munitionMod models.MunitionMod
				err = rows.Scan(
					&munitionMod.RN, &munitionMod.PRN, &munitionMod.NomModif, &munitionMod.PackModifSp, &munitionMod.Code)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				munitionsMod = append(munitionsMod, munitionMod)
			}
		}
		return munitionsMod, nil
	}
}

func UpdateMunitionMod(munitionMod models.MunitionMod) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres // TODO: реализовать позже
		return nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_munition_mod SET "VERSION"=:2, crn=:3, code=:4, "NAME"=:5
                              				 WHERE rn=:1`,
				munitionMod.RN, munitionMod.PRN, munitionMod.NomModif, munitionMod.PackModifSp, munitionMod.Code)
			return err
		}
	}
	return nil
}

func DeleteMunitionMod(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_munition_mod WHERE rn=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_munition_mod WHERE rn=:1`, id)
			return err
		}
	}
	return nil
}
