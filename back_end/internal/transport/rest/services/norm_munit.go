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

func CreateNormMunit(normMunit models.NormMunit) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_norm_munit
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, normMunit.RN, normMunit.CRN, normMunit.Version, normMunit.Code, normMunit.Name,
			normMunit.NormType, normMunit.InventProp)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_norm_munit
				VALUES (:1, :2, :3, :4, :5, :6)`, normMunit.RN, normMunit.CRN, normMunit.Version, normMunit.Code, normMunit.Name, normMunit.NormType)
			return err
		}
	}
	return nil
}

func ReadNormMunits() ([]models.NormMunit, error) {
	var normMunits []models.NormMunit
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_norm_munit`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			var normMunit models.NormMunit
			err = rows.Scan(&normMunit.RN, &normMunit.CRN, &normMunit.Version, &normMunit.Code, &normMunit.Name, &normMunit.NormType, &normMunit.InventProp)
			if err != nil {
				return nil, err
			}
			normMunits = append(normMunits, normMunit)
		}
		return normMunits, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_norm_munit`)
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
				var normMunit models.NormMunit
				err = rows.Scan(&normMunit.RN, &normMunit.CRN, &normMunit.Version, &normMunit.Code, &normMunit.Name, &normMunit.NormType)
				if err != nil {
					return nil, err
				}
				normMunits = append(normMunits, normMunit)
			}
		}
		return normMunits, nil
	}
}

func CreateNormMunitSp(normMunitSp models.NormMunitSp) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_norm_munitsp
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
			&normMunitSp.RN, &normMunitSp.PRN, &normMunitSp.CRN, &normMunitSp.Version, &normMunitSp.Munition,
			&normMunitSp.Period, &normMunitSp.QNT, &normMunitSp.PackACS, &normMunitSp.ByDateFirst, &normMunitSp.ByDateFact,
			&normMunitSp.ForPay, &normMunitSp.ACSOneKompl)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_norm_munitsp
			VALUES (:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12)`,
				&normMunitSp.RN, &normMunitSp.PRN, &normMunitSp.CRN, &normMunitSp.Version, &normMunitSp.Munition,
				&normMunitSp.Period, &normMunitSp.QNT, &normMunitSp.PackACS, &normMunitSp.ByDateFirst, &normMunitSp.ByDateFact,
				&normMunitSp.ForPay, &normMunitSp.ACSOneKompl)
			return err
		}
	}
	return nil
}

func ReadNormMunitsSp() ([]models.NormMunitSp, error) {
	var normMunitsSp []models.NormMunitSp
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_norm_munitsp`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			var normMunitSp models.NormMunitSp
			err = rows.Scan(&normMunitSp.RN, &normMunitSp.PRN, &normMunitSp.CRN, &normMunitSp.Version, &normMunitSp.Munition,
				&normMunitSp.Period, &normMunitSp.QNT, &normMunitSp.PackACS, &normMunitSp.ByDateFirst, &normMunitSp.ByDateFact,
				&normMunitSp.ForPay, &normMunitSp.ACSOneKompl)
			if err != nil {
				return nil, err
			}
			normMunitsSp = append(normMunitsSp, normMunitSp)
		}
		return normMunitsSp, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_norm_munitsp`)
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
				var normMunitSp models.NormMunitSp
				err = rows.Scan(&normMunitSp.RN, &normMunitSp.PRN, &normMunitSp.CRN, &normMunitSp.Version, &normMunitSp.Munition,
					&normMunitSp.Period, &normMunitSp.QNT, &normMunitSp.PackACS, &normMunitSp.ByDateFirst, &normMunitSp.ByDateFact,
					&normMunitSp.ForPay, &normMunitSp.ACSOneKompl)
				if err != nil {
					return nil, err
				}
				normMunitsSp = append(normMunitsSp, normMunitSp)
			}
		}
		return normMunitsSp, nil
	}
}

func ReadNormMunitsSpAlt() ([]models.NormMunitSpAlt, error) {
	var normMunitsSpAlt []models.NormMunitSpAlt
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_norm_munitsp_alt`)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка получения данных: %v\n", err)
			os.Exit(1)
		}
		defer rows.Close()
		for rows.Next() {
			var normMunitSpAlt models.NormMunitSpAlt
			err = rows.Scan(&normMunitSpAlt.RN, &normMunitSpAlt.PRN, &normMunitSpAlt.CRN, &normMunitSpAlt.Version, &normMunitSpAlt.GroupMunit)
			if err != nil {
				return nil, err
			}
			normMunitsSpAlt = append(normMunitsSpAlt, normMunitSpAlt)
		}
		return normMunitsSpAlt, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_norm_munitsp_alt`)
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
				var normMunitSpAlt models.NormMunitSpAlt
				err = rows.Scan(&normMunitSpAlt.RN, &normMunitSpAlt.PRN, &normMunitSpAlt.CRN, &normMunitSpAlt.Version, &normMunitSpAlt.GroupMunit)
				if err != nil {
					return nil, err
				}
				normMunitsSpAlt = append(normMunitsSpAlt, normMunitSpAlt)
			}
		}
		return normMunitsSpAlt, nil
	}
}
