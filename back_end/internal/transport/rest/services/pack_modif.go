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

func CreatePackModif(packModif models.PackModif) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_pack_modif
			VALUES ($1, $2, $3, $4, $5)`, packModif.Code, packModif.Version, packModif.CRN, packModif.Name, packModif.RN)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_pack_modif
				VALUES (:1, :2, :3, :4, :5)`, packModif.Code, packModif.Version, packModif.CRN, packModif.Name, packModif.RN)
			return err
		}
	}
	return nil
}

func ReadPackModifs() ([]models.PackModif, error) {
	var packModif models.PackModif
	var packModifs []models.PackModif
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_pack_modif`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&packModif.Code, &packModif.Version, &packModif.CRN, &packModif.Name, &packModif.RN)
			if err != nil {
				return nil, err
			}
			packModifs = append(packModifs, packModif)
		}
		return packModifs, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_pack_modif`)
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
				err = rows.Scan(&packModif.Code, &packModif.Version, &packModif.CRN, &packModif.Name, &packModif.RN)
				if err != nil {
					return nil, err
				}
				packModifs = append(packModifs, packModif)
			}
		}
		return packModifs, nil
	}
}

func CreatePackModifSp(packModifSp models.PackModifSp) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_pack_modifsp
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, &packModifSp.RN, &packModifSp.PRN, &packModifSp.Code, &packModifSp.Name,
			&packModifSp.Girth, &packModifSp.Neck, &packModifSp.Height)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_norm_munitsp
			VALUES (:1, :2, :3, :4, :5, :6, :7)`, &packModifSp.RN, &packModifSp.PRN, &packModifSp.Code, &packModifSp.Name,
				&packModifSp.Girth, &packModifSp.Neck, &packModifSp.Height)
			return err
		}
	}
	return nil
}

func ReadPackModifSps() ([]models.PackModifSp, error) {
	var packModifSps []models.PackModifSp
	var packModifSp models.PackModifSp
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_pack_modifsp`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&packModifSp.RN, &packModifSp.PRN, &packModifSp.Code, &packModifSp.Name,
				&packModifSp.Girth, &packModifSp.Neck, &packModifSp.Height)
			if err != nil {
				return nil, err
			}
			packModifSps = append(packModifSps, packModifSp)
		}
		return packModifSps, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_pack_modifsp`)
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
				err = rows.Scan(&packModifSp.RN, &packModifSp.PRN, &packModifSp.Code, &packModifSp.Name,
					&packModifSp.Girth, &packModifSp.Neck, &packModifSp.Height)
				if err != nil {
					return nil, err
				}
				packModifSps = append(packModifSps, packModifSp)
			}
		}
		return packModifSps, nil
	}
}
