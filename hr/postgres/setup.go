package postgres

import (
	"fmt"

	"github.com/micro/go-micro/util/log"
)

const (
	createDepartmentTable = `CREATE TABLE IF NOT EXISTS department (
			id VARCHAR(64) NOT NULL PRIMARY KEY,
			name VARCHAR(265) NOT NULL,
			UNIQUE(name)
		);`

	createEmployeeTypeTable = `CREATE TABLE IF NOT EXISTS employee_type (
			id VARCHAR(64) NOT NULL PRIMARY KEY,
			department_id VARCHAR(64) NOT NULL,
			name VARCHAR(256) NOT NULL,
			can_login BOOLEAN NOT NULL,
			FOREIGN KEY (department_id) REFERENCES department (id),
			UNIQUE(name)
		);`

	createEmployeeTable = `CREATE TABLE IF NOT EXISTS employee (
			id VARCHAR(64) NOT NULL PRIMARY KEY,
			employee_type_id VARCHAR(64) NOT NULL,
			department_id VARCHAR(64) NOT NULL,
			name VARCHAR(256) NOT NULL,
			username VARCHAR(64) NOT NULL,
			email VARCHAR(64) NOT NULL,
			phone_number VARCHAR(64) NOT NULL,
			FOREIGN KEY (department_id) REFERENCES department (id),
			FOREIGN KEY (employee_type_id) REFERENCES employee_type (id),
			UNIQUE(username),
			UNIQUE(email),
			UNIQUE(phone_number)
		);`
)

// department
func (pg *PgDb) CreateDepartmentTable() error {
	log.Trace("Creating user table")
	_, err := pg.db.Exec(createDepartmentTable)
	return err
}

func (pg *PgDb) DepartmentTableExists() bool {
	exists, _ := pg.tableExists("department")
	return exists
}

// employee_type
func (pg *PgDb) CreateEmployeeTypeTable() error {
	log.Trace("Creating employee_type table")
	_, err := pg.db.Exec(createEmployeeTypeTable)
	return err
}

func (pg *PgDb) EmployeeTypeTableExists() bool {
	exists, _ := pg.tableExists("employee_type")
	return exists
}

// employee
func (pg *PgDb) CreateEmployeeTable() error {
	log.Trace("Creating employee table")
	_, err := pg.db.Exec(createEmployeeTable)
	return err
}

func (pg *PgDb) EmployeeTableExists() bool {
	exists, _ := pg.tableExists("employee")
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
	if err := pg.dropIndex("employee"); err != nil {
		return err
	}

	if err := pg.dropIndex("employee_type"); err != nil {
		return err
	}

	if err := pg.dropIndex("department"); err != nil {
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

