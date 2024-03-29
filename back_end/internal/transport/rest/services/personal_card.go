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

func CreatePersonalCard(personalCard models.PersonalCard) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_personal_card
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, personalCard.Id, personalCard.FIO, personalCard.JobBeginDate, personalCard.Sex,
			personalCard.DivisionId, personalCard.PositionId, personalCard.RankId, personalCard.Climate)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_personal_card
				VALUES (:1, :2, :3, :4, :5, :6, :7)`, personalCard.Id, personalCard.FIO, personalCard.JobBeginDate, personalCard.Sex,
				personalCard.DivisionId, personalCard.PositionId, personalCard.RankId, personalCard.Climate)
			return err
		}
	}
	return nil
}

func ReadPersonalCards() ([]models.PersonalCard, error) {
	var personalCards []models.PersonalCard
	var personalCard models.PersonalCard
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_personal_card`)
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
			personalCards = append(personalCards, personalCard)
		}
		return personalCards, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_personal_card`)
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
					&personalCard.Id, &personalCard.FIO, &personalCard.JobBeginDate, &personalCard.Sex, &personalCard.DivisionId,
					&personalCard.PositionId, &personalCard.RankId, &personalCard.Climate)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				personalCards = append(personalCards, personalCard)
			}
		}
		return personalCards, nil
	}
}

func UpdatePersonalCard(personalCard models.PersonalCard) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres // TODO: реализовать позже
		return nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_ = connOracle
		}
	}
	return nil
}

func DeletePersonalCard(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_personal_card WHERE id=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_personal_card WHERE id=:1`, id)
			return err
		}
	}
	return nil
}
