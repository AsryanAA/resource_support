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

func CreateRank(rank models.Rank) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_rank
			VALUES ($1, $2, $3)`, rank.Id, rank.Name, rank.Description)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_rank
				VALUES (:1, :2, :3)`, rank.Id, rank.Name, rank.Description)
			return err
		}
	}
	return nil
}

func ReadRanks() ([]models.Rank, error) {
	var ranks []models.Rank
	var rank models.Rank
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_rank`)
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
			ranks = append(ranks, rank)
		}
		return ranks, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_rank`)
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
					&rank.Id, &rank.Name, &rank.Description)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				ranks = append(ranks, rank)
			}
		}
		return ranks, nil
	}
}

func UpdateRank(rank models.Rank) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres // TODO: реализовать позже
		return nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_rank SET "name"=:2, description=:3
                              				 WHERE id=:1`,
				rank.Name, rank.Description, rank.Id)
			return err
		}
	}
	return nil
}

func DeleteRank(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_rank WHERE id=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_rank WHERE id=:1`, id)
			return err
		}
	}
	return nil
}
