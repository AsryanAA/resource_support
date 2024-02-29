package services

import (
	"back/internal/database"
	"back/internal/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func CreateAccessor(accessor models.Accessor) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `	INSERT INTO parus.udo_accessor VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			accessor.RN, accessor.CRN, accessor.Version, accessor.Code, accessor.Name, accessor.DicNomNs, accessor.NomModif)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_accessor VALUES (:1, :2, :3, :4, :5, :6, :7)`,
				accessor.RN, accessor.CRN, accessor.Version, accessor.Code, accessor.Name, accessor.DicNomNs, accessor.NomModif)
			return err
		}
	}
	return nil
}

func ReadAccessors() ([]models.Accessor, error) {
	var accessors []models.Accessor
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_accessor`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var accessor models.Accessor
			err = rows.Scan(&accessor.RN, &accessor.CRN, &accessor.Version, &accessor.Code,
				&accessor.Name, &accessor.DicNomNs, &accessor.NomModif)
			if err != nil {
				fmt.Println("Ошибка получения данных: ", err)
			}
			accessors = append(accessors, accessor)
		}
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_accessor`)
			if err != nil {
				log.Fatal(err)
			}
			defer func(rows *sql.Rows) {
				err = rows.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(rows)
			for rows.Next() {
				var accessor models.Accessor
				err = rows.Scan(&accessor.RN, &accessor.CRN, &accessor.Version, &accessor.Code,
					&accessor.Name, &accessor.DicNomNs, &accessor.NomModif)
				if err != nil {
					fmt.Println("Ошибка получения данных: ", err)
				}
				accessors = append(accessors, accessor)
			}
		}
	}
	return accessors, nil
}

func UpdateAccessor(accessor models.Accessor) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `UPDATE parus.udo_accessor SET crn=$2, "VERSION"=$3, code=$4, "NAME"=$5, dicnomns=$6, nommodif=$7 WHERE rn=$1`,
			accessor.RN, accessor.CRN, accessor.Version, accessor.Code, accessor.Name, accessor.DicNomNs, accessor.NomModif)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_accessor SET crn=:2, "VERSION"=:3, code=:4, "NAME"=$5, dicnomns=:6, nommodif=:7 WHERE rn=:1`,
				accessor.RN, accessor.CRN, accessor.Version, accessor.Code, accessor.Name, accessor.DicNomNs, accessor.NomModif)
			return err
		}
	}
	return nil
}

func DeleteAccessor(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_accessor WHERE rn=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_accessor WHERE rn=:1`, id)
			return err
		}
	}
	return nil
}
