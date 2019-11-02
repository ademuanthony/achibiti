package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"context"
	"database/sql"
	"github.com/ademuanthony/achibiti/hr/proto/hr"
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

func (pg *PgDb) CreateDepartment(ctx context.Context, department go_micro_srv_hr.Department) error {
	panic("implement me")
}

func (pg *PgDb) Departments(ctx context.Context, skipCount int32, maxResultCount int32) ([]*go_micro_srv_hr.Department, int64, error) {
	panic("implement me")
}

func (pg *PgDb) UpdateDepartment(ctx context.Context, department go_micro_srv_hr.Department) error {
	panic("implement me")
}

func (pg *PgDb) DeleteDepartment(ctx context.Context, id string) error {
	panic("implement me")
}

func (pg *PgDb) CreateEmployeeType(ctx context.Context, employeeType go_micro_srv_hr.EmployeeType) error {
	panic("implement me")
}

func (pg *PgDb) EmployeeTypes(ctx context.Context, skipCount int32, maxResultCount int32) ([]*go_micro_srv_hr.EmployeeType, int64, error) {
	panic("implement me")
}

func (pg *PgDb) DeleteEmployeeType(ctx context.Context, id string) error {
	panic("implement me")
}

func (pg *PgDb) EmployeeType(ctx context.Context, id string) (go_micro_srv_hr.EmployeeType, error) {
	panic("implement me")
}

func (pg *PgDb) CreateEmployee(ctx context.Context, employee go_micro_srv_hr.Employee) error {
	panic("implement me")
}

func (pg *PgDb) Employees(ctx context.Context, departmentId string, employeeTypeId string, skipCount int32,
	resultCount int32) ([]*go_micro_srv_hr.Employee, int64, error) {
	panic("implement me")
}

func (pg *PgDb) Employee(ctx context.Context, id string) (*go_micro_srv_hr.Employee, error) {
	panic("implement me")
}

func (pg *PgDb) UpdateEmployee(ctx context.Context, employee *go_micro_srv_hr.Employee) error {
	panic("implement me")
}

func (pg *PgDb) DeleteEmployee(ctx context.Context, employeeId string) error {
	panic("implement me")
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}
