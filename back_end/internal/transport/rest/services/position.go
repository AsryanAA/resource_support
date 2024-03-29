package services

import (
	"back/internal/database"
	"back/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func CreatePosition(position models.Position) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `INSERT INTO parus.udo_position
			VALUES ($1, $2, $3, $4)`, position.Id, position.Name, position.Description)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`INSERT INTO parus.udo_position
				VALUES (:1, :2, :3)`, position.Id, position.Name, position.Description)
			if err != nil {
				textError := err.Error()
				if textError[11:28] == models.DataBaseErrors[1] {
					return errors.New(models.ErrorUnique)
				}
				return err
			}
		}
	}
	return nil
}

func ReadPositions() ([]models.Position, error) {
	var positions []models.Position
	var position models.Position
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		rows, err := connPostgres.Query(context.Background(), `SELECT * FROM parus.udo_position`)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan()
			if err != nil {
				return nil, err
			}
			positions = append(positions, position)
		}
		return positions, nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			rows, err := connOracle.Query(`SELECT * FROM parus.udo_position`)
			if err != nil {
				return nil, err
			}
			defer func(rows *sql.Rows) {
				err = rows.Close()
				if err != nil {
					log.Fatal(err)
				}
			}(rows)
			for rows.Next() {
				err = rows.Scan(
					&position.Id, &position.Name, &position.Description)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}
				positions = append(positions, position)
			}
		}
		return positions, nil
	}
}

func UpdatePosition(position models.Position) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_ = connPostgres // TODO: реализовать позже
		return nil
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`UPDATE parus.udo_position SET "name"=:2, description=:3 WHERE id=:1`,
				position.Name, position.Description, position.Id)
			return err
		}
	}
	return nil
}

func DeletePosition(id int64) error {
	connPostgres, okPostgres := database.DB.(*pgx.Conn)
	if okPostgres {
		_, err := connPostgres.Exec(context.Background(), `DELETE FROM parus.udo_position WHERE id=$1`, id)
		return err
	} else {
		connOracle, okOracle := database.DB.(*sql.DB)
		if okOracle {
			_, err := connOracle.Exec(`DELETE FROM parus.udo_position WHERE id=:1`, id)
			return err
		}
	}
	return nil
}
