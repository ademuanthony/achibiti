package postgres

//go:generate sqlboiler --wipe psql --no-hooks --no-auto-timestamps

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ademuanthony/achibiti/hr/postgres/models"
	"github.com/ademuanthony/achibiti/hr/proto/hr"
	"github.com/ademuanthony/achibiti/utils"
	"github.com/micro/go-micro/util/log"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
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
	departmentModel := models.Department{
		ID:   department.Id,
		Name: department.Name,
	}

	return departmentModel.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) Departments(ctx context.Context, skipCount int32, maxResultCount int32) ([]*go_micro_srv_hr.Department, int64, error) {
	var queries []qm.QueryMod
	if maxResultCount > -1 {
		queries = append(queries, qm.Offset(int(skipCount)), qm.Limit(int(maxResultCount)))
	}

	totalCount, err := models.Departments().Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	departmentSlice, err := models.Departments(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var departments []*go_micro_srv_hr.Department
	for _, department := range departmentSlice {
		departments = append(departments, &go_micro_srv_hr.Department{
			Id:                   department.ID,
			Name:                 department.Name,
		})
	}

	return departments, totalCount, nil
}

func (pg *PgDb) UpdateDepartment(ctx context.Context, department go_micro_srv_hr.Department) error {
	departmentModel := models.Department{
		ID:   department.Id,
		Name: department.Name,
	}

	_, err := departmentModel.Update(ctx, pg.db, boil.Infer())
	return err
}

func (pg *PgDb) DeleteDepartment(ctx context.Context, id string) error {
	// todo delete associated records first
	departmentModel := models.Department{
		ID:   id,
	}

	_, err := departmentModel.Delete(ctx, pg.db)
	return err
}

func (pg *PgDb) CreateEmployeeType(ctx context.Context, employeeType go_micro_srv_hr.EmployeeType) error {
	employeeModel := models.EmployeeType{
		ID:           employeeType.Id,
		DepartmentID: employeeType.DepartmentId,
		Name:         employeeType.Name,
		CanLogin:     employeeType.CanLogin,
	}

	return employeeModel.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) EmployeeTypes(ctx context.Context, departmentId string, skipCount int32, maxResultCount int32) ([]*go_micro_srv_hr.EmployeeType, int64, error) {
	var queries = []qm.QueryMod{
		qm.Load(models.EmployeeTypeRels.Department),
	}
	if departmentId != "" {
		queries = append(queries, models.EmployeeTypeWhere.DepartmentID.EQ(departmentId))
	}

	totalCount, err := models.EmployeeTypes(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	if maxResultCount > 0 {
		queries = append(queries, qm.Offset(int(skipCount)), qm.Limit(int(maxResultCount)))
	}

	employeeTypeSlice, err := models.EmployeeTypes(queries...).All(ctx, pg.db)
	var employeeTypes []*go_micro_srv_hr.EmployeeType
	for _, employeeType := range employeeTypeSlice {
		employeeTypes = append(employeeTypes, &go_micro_srv_hr.EmployeeType{
			Id:                   employeeType.ID,
			Name:                 employeeType.Name,
			Department:           employeeType.R.Department.Name,
			DepartmentId:         employeeType.DepartmentID,
			CanLogin:             employeeType.CanLogin,
		})
	}

	return employeeTypes, totalCount, nil
}

func (pg *PgDb) DeleteEmployeeType(ctx context.Context, id string) error {
	// TODO delete related employees
	employeeType := models.EmployeeType{
		ID:           id,
	}

	_, err := employeeType.Delete(ctx, pg.db)
	return err
}

func (pg *PgDb) EmployeeType(ctx context.Context, id string) (*go_micro_srv_hr.EmployeeType, error) {
	employeeTypeModel, err := models.EmployeeTypes(models.EmployeeTypeWhere.ID.EQ(id)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	return &go_micro_srv_hr.EmployeeType{
		Id:                   employeeTypeModel.ID,
		Name:                 employeeTypeModel.Name,
		Department:           employeeTypeModel.R.Department.Name,
		DepartmentId:         employeeTypeModel.DepartmentID,
		CanLogin:             employeeTypeModel.CanLogin,
	}, nil
}

func (pg *PgDb) CreateEmployee(ctx context.Context, employee go_micro_srv_hr.Employee) error {
	// check for username availability
	exist, err := models.Employees(models.EmployeeWhere.Username.EQ(employee.GetUsername())).Exists(ctx, pg.db)
	if err != nil {
		return fmt.Errorf("error in validating username, %s", err.Error())
	}

	if exist {
		return fmt.Errorf("the selected username, %s has been taken", employee.GetUsername())
	}

	// check for email availability
	exist, err = models.Employees(models.EmployeeWhere.Email.EQ(employee.GetEmail())).Exists(ctx, pg.db)
	if err != nil {
		return fmt.Errorf("error in validating email, %s", err.Error())
	}

	if exist {
		return fmt.Errorf("the selected email, %s has been taken", employee.GetEmail())
	}

	// check for phone number availability
	exist, err = models.Employees(models.EmployeeWhere.PhoneNumber.EQ(employee.GetPhoneNumber())).Exists(ctx, pg.db)
	if err != nil {
		return fmt.Errorf("error in validating phone number, %s", err.Error())
	}

	if exist {
		return fmt.Errorf("the selected phone number, %s has been taken", employee.GetPhoneNumber())
	}

	employeeModel := models.Employee{
		ID:             employee.GetId(),
		EmployeeTypeID: employee.GetEmployeeTypeId(),
		DepartmentID:   employee.GetDepartmentId(),
		Name:           employee.GetName(),
		Username:       employee.GetUsername(),
		Email:          employee.GetEmail(),
		PhoneNumber:    employee.GetPhoneNumber(),
	}

	return employeeModel.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) Employees(ctx context.Context, departmentId string, employeeTypeId string, skipCount int32,
	resultCount int32) ([]*go_micro_srv_hr.Employee, int64, error) {

	var queries []qm.QueryMod
	if departmentId != "" {
		queries = append(queries, models.EmployeeWhere.DepartmentID.EQ(departmentId))
	}

	if employeeTypeId != "" {
		queries = append(queries, models.EmployeeWhere.EmployeeTypeID.EQ(employeeTypeId))
	}

	totalCount, err := models.Employees(queries...).Count(ctx, pg.db)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot get employee count, %s", err.Error())
	}

	queries = append(queries, qm.Offset(int(skipCount)), qm.Limit(int(resultCount)))
	employeeSlice, err := models.Employees(queries...).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var employees []*go_micro_srv_hr.Employee
	for _, employee := range employeeSlice {
		employees = append(employees, &go_micro_srv_hr.Employee{
			Id:                   employee.ID,
			Name:                 employee.Name,
			DepartmentId:         employee.DepartmentID,
			EmployeeTypeId:       employee.EmployeeTypeID,
			Department:           employee.R.Department.Name,
			EmployeeType:         employee.R.EmployeeType.Name,
			Username:             employee.Username,
			Email:                employee.Email,
			PhoneNumber:          employee.PhoneNumber,
		})
	}

	return employees, totalCount, nil
}

func (pg *PgDb) Employee(ctx context.Context, id string) (*go_micro_srv_hr.Employee, error) {
	employee, err := models.Employees(models.EmployeeWhere.ID.EQ(id)).One(ctx, pg.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no record found")
		}

		return nil, err
	}

	return &go_micro_srv_hr.Employee{
		Id:                   employee.ID,
		Name:                 employee.Name,
		DepartmentId:         employee.DepartmentID,
		EmployeeTypeId:       employee.EmployeeTypeID,
		Department:           employee.R.Department.Name,
		EmployeeType:         employee.R.EmployeeType.Name,
		Username:             employee.Username,
		Email:                employee.Email,
		PhoneNumber:          employee.PhoneNumber,
	}, nil
}

func (pg *PgDb) UpdateEmployee(ctx context.Context, employee *go_micro_srv_hr.Employee) error {
	employeeModel := models.Employee{
		ID:             employee.Id,
		EmployeeTypeID: employee.EmployeeTypeId,
		DepartmentID:   employee.DepartmentId,
		Name:           employee.Name,
		Username:       employee.Username,
		Email:          employee.Email,
		PhoneNumber:    employee.PhoneNumber,
	}

	_, err := employeeModel.Update(ctx, pg.db, boil.Infer())
	return err
}

func (pg *PgDb) DeleteEmployee(ctx context.Context, employeeId string) error {
	// todo delete associated records
	employee := models.Employee{ID:employeeId}
	_, err := employee.Delete(ctx, pg.db)
	return err
}

func (pg *PgDb) Close() error {
	log.Trace("Closing postgresql connection")
	return pg.db.Close()
}
