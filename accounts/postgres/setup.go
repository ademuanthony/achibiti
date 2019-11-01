package postgres

import (
	"fmt"

	"github.com/micro/go-micro/util/log"
)

const (
	createUserTable = `CREATE TABLE IF NOT EXISTS account_user (
		id VARCHAR(64) NOT NULL PRIMARY KEY,
		name VARCHAR(64) NOT NULL,
		username VARCHAR(265) NOT NULL,
		password_hash VARCHAR(265) NOT NULL,
		email VARCHAR(265) NOT NULL,
		phone_number VARCHAR(265) NOT NULL,
		role VARCHAR(265) NOT NULL);`
)

func (pg *PgDb) CreateUserTable() error {
	log.Trace("Creating user table")
	_, err := pg.db.Exec(createUserTable)
	return err
}

func (pg *PgDb) UserTableExists() bool {
	exists, _ := pg.tableExists("account_user")
	return exists
}


func (pg *PgDb) tableExists(name string) (bool, error) {
	rows, err := pg.db.Query(`SELECT relname FROM pg_class WHERE relname = $1`, name)
	if err == nil {
		defer func() {
			if e := rows.Close(); e != nil {
				log.Error("Close of Query failed: ", e)
			}
		}()
		return rows.Next(), nil
	}
	return false, err
}

func (pg *PgDb) DropAllTables() error {
	if err := pg.dropIndex("account_user"); err != nil {
		return err
	}

	return nil
}

func (pg *PgDb) dropTable(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, name))
	return err
}

func (pg *PgDb) dropIndex(name string) error {
	log.Tracef("Dropping table %s", name)
	_, err := pg.db.Exec(fmt.Sprintf(`DROP INDEX IF EXISTS %s;`, name))
	return err
}

