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

func CreateMunition(munition models.Munition) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `	INSERT INTO parus.udo_munition VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
			munition.RN, munition.Version, munition.CRN, munition.Code, munition.Name, munition.DicNomNs,
			munition.PackModif, munition.KompSum, munition.Cloth, munition.IgnoreSupply, munition.Sex,
			munition.UseFaktHeihth, munition.BegstepHeihth, munition.NomModif)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`	INSERT INTO parus.udo_munition VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14)`,
				munition.RN, munition.Version, munition.CRN, munition.Code, munition.Name, munition.DicNomNs,
				munition.PackModif, munition.KompSum, munition.Cloth, munition.IgnoreSupply, munition.Sex,
				munition.UseFaktHeihth, munition.BegstepHeihth, munition.NomModif)
			return err
		}
	}
	return nil
}

func ReadMunitions() ([]models.Munition, error) {
	var munitions []models.Munition
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_munition`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			var munition models.Munition
			err = rows.Scan(
				&munition.RN, &munition.Version, &munition.CRN, &munition.Code, &munition.Name, &munition.DicNomNs,
				&munition.PackModif, &munition.KompSum, &munition.Cloth, &munition.IgnoreSupply, &munition.Sex,
				&munition.UseFaktHeihth, &munition.BegstepHeihth, &munition.NomModif,
			)
			if err != nil {
				return nil, err
			}
			munitions = append(munitions, munition)
		}
		return munitions, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_munition`)
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
				var munition models.Munition
				err = rows.Scan(
					&munition.RN, &munition.Version, &munition.CRN, &munition.Code, &munition.Name, &munition.DicNomNs,
					&munition.PackModif, &munition.KompSum, &munition.Cloth, &munition.IgnoreSupply, &munition.Sex,
					&munition.UseFaktHeihth, &munition.BegstepHeihth, &munition.NomModif,
				)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				munitions = append(munitions, munition)
			}
		}
		return munitions, nil
	}
}

func UpdateMunition(munition models.Munition) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `UPDATE parus.udo_munition SET "VERSION"=$2, crn=$3,
		                     code=$4, "NAME"=$5, dicnomns=$6, pack_modif=$7, komp_sum=$8, nommodif=$9, cloth=$10,
		                     ignore_supply=$11, sex=$12, use_fakt_heihth=$13, begstep_heihth=$14 WHERE rn=$1`,
			munition.RN, munition.Version, munition.CRN, munition.Code, munition.Name, munition.DicNomNs,
			munition.PackModif, munition.KompSum, munition.Cloth, munition.IgnoreSupply, munition.Sex,
			munition.UseFaktHeihth, munition.BegstepHeihth, munition.NomModif,
		)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_munition SET "VERSION"=:2, crn=:3,
                              code=:4, "NAME"=:5, dicnomns=:6, pack_modif=:7, komp_sum=:8, nommodif=:9, cloth=:10,
                              ignore_supply=:11, sex=:12, use_fakt_heihth=:13, begstep_heihth=:14 WHERE rn=:1`,
				munition.RN, munition.Version, munition.CRN, munition.Code, munition.Name, munition.DicNomNs,
				munition.PackModif, munition.KompSum, munition.Cloth, munition.IgnoreSupply, munition.Sex,
				munition.UseFaktHeihth, munition.BegstepHeihth, munition.NomModif,
			)
			return err
		}
	}
	return nil
}

func DeleteMunition(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_munition WHERE rn=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_munition WHERE rn=:1`, id)
			return err
		}
	}
	return nil
}
