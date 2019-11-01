package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"database/sql"
	"github.com/ademuanthony/achibiti/utils"
	"github.com/micro/go-micro/util/log"
	"time"
)

type PgDb struct {
	db *sql.DB
	queryTimeout time.Duration
}

func NewPgDb(host, port, user, pass, dbname string) (*PgDb, error) {
	db, err := utils.PgConnect(host, port, user, pass, dbname)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(5)
	return &PgDb{
		db: db,
		queryTimeout: time.Second * 30,
	}, nil
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}
