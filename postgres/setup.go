package postgres

import "fmt"

const (
	createExchangeTable = `CREATE TABLE IF NOT EXISTS exchange (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL);`
)

// CreateExchangeTable creates the exchange table
func (pg *PgDb) CreateExchangeTable() error {
	log.Trace("Creating exchange tick table")
	_, err := pg.db.Exec(createExchangeTable)
	return err
}

// ExchangeTableExits checks for the existence of exchange table
func (pg *PgDb) ExchangeTableExits() bool {
	exists, _ := pg.tableExists("exchange")
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


// DropAllTables drops all tables from the db
func (pg *PgDb) DropAllTables() error {
	// exchange
	if err := pg.dropTable("exchange"); err != nil {
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
