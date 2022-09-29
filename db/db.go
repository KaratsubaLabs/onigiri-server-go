package db

import {
	"os"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
}

type Connection struct {
	DB		*pgx.Conn
	Connected	bool
}

func (r *Connection) Connect error {
	
	var err error
	r.DB, err = pgxpool.New(context.Background(), os.Getenv("ONIGIRI_DB_URL"))

	r.Connected = (err == nil)

	return err
}

func (r *Connection) Migrate error {

	var err [2]Error

	_, err[0] := r.DB.Exec(`CREATE TABLE Users()`)
	_, err[1] := r.DB.Exec(`CREATE TABLE Musics()`)

	for _, err := range err {
		if err != nil {
			return err
		}
	}

	r.DB.Exec(`ALTER TABLE Users ADD Username VARCHAR`);
	r.DB.Exec(`ALTTER TABLE Users ADD Pass
