package handler

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/micro/go-micro/util/log"

	acl "github.com/ademuanthony/achibiti/acl/proto/acl"
	go_micro_srv_hr "github.com/ademuanthony/achibiti/hr/proto/hr"
)

type hr struct{
	dataSource DataSource
	aclService acl.AclService
}

func NewHr(dataSource DataSource, aclService acl.AclService) *hr {
	return &hr{dataSource: dataSource, aclService:aclService}
}

type DataSource interface {
	CreateDepartment(ctx context.Context, department go_micro_srv_hr.Department) error
	Departments(ctx context.Context, skipCount int32, maxResultCount int32) ([]*go_micro_srv_hr.Department, int64, error)
	UpdateDepartment(ctx context.Context, department go_micro_srv_hr.Department) error
	DeleteDepartment(ctx context.Context, id string) error
	CreateEmployeeType(ctx context.Context, employeeType go_micro_srv_hr.EmployeeType) error
	EmployeeTypes(ctx context.Context, skipCount int32, maxResultCount int32) ([]*go_micro_srv_hr.EmployeeType, int64, error)
	DeleteEmployeeType(ctx context.Context, id string) error
	EmployeeType(ctx context.Context, id string) (go_micro_srv_hr.EmployeeType, error)
	CreateEmployee(ctx context.Context, employee go_micro_srv_hr.Employee) error
	Employees(ctx context.Context, departmentId string, employeeTypeId string, skipCount int32,
		resultCount int32) ([]*go_micro_srv_hr.Employee, int64, error)
	Employee(ctx context.Context, id string) (*go_micro_srv_hr.Employee, error)
	UpdateEmployee(ctx context.Context, employee *go_micro_srv_hr.Employee) error
	DeleteEmployee(ctx context.Context, employeeId string) error
	
} 

func (h hr) CreateDepartment(ctx context.Context, req *go_micro_srv_hr.CreateDepartmentRequest, resp *go_micro_srv_hr.CreateDepartmentResponse) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("cannot create department, error in generating id, %s", err.Error())
	}

	department := go_micro_srv_hr.Department{
		Id:                   id.String(),
		Name:                 req.GetName(),
	}

	if err = h.dataSource.CreateDepartment(ctx, department); err != nil {
		return err
	}

	resp.Id = id.String()
	return nil
}

func (h hr) Departments(ctx context.Context, req *go_micro_srv_hr.DepartmentsRequest, resp *go_micro_srv_hr.DepartmentsResponse) error {
	departments, totalCount, err := h.dataSource.Departments(ctx, req.GetSkipCount(), req.GetMaxResultCount())
	if err != nil {
		return err
	}

	resp.Departments = departments
	resp.TotalCount = totalCount
	return nil
}

func (h hr) UpdateDepartment(ctx context.Context, req *go_micro_srv_hr.UpdateDepartmentRequest, _ *go_micro_srv_hr.EmptyMessage) error {
	department := go_micro_srv_hr.Department{
		Id:                   req.Id,
		Name:                 req.Name,
	}

	return h.dataSource.UpdateDepartment(ctx, department)
}

func (h hr) DeleteDepartment(ctx context.Context, req *go_micro_srv_hr.DeleteDepartmentRequest, _ *go_micro_srv_hr.EmptyMessage) error {
	return h.dataSource.DeleteDepartment(ctx, req.Id)
}

func (h hr) CreateEmployeeType(ctx context.Context, req *go_micro_srv_hr.CreateEmployeeTypeRequest, resp *go_micro_srv_hr.CreateEmployeeTypeResponse) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("error in generating id, %s", err.Error())
	}

	employeeType := go_micro_srv_hr.EmployeeType{
		Id:                   id.String(),
		Name:                 req.GetName(),
		DepartmentId:         req.GetDepartmentId(),
		CanLogin:             req.GetCanLogin(),
	}

	if err = h.dataSource.CreateEmployeeType(ctx, employeeType); err != nil {
		return err
	}

	resp.EmployeeTypeId = id.String()
	return nil
}

func (h hr) EmployeeTypes(ctx context.Context, req *go_micro_srv_hr.EmployeeTypesRequest, resp *go_micro_srv_hr.EmployeeTypesResponse) error {
	employeeTypes, totalCount, err := h.dataSource.EmployeeTypes(ctx, req.SkipCount, req.MaxResultCount)
	if err != nil {
		return err
	}

	resp.TotalCount = totalCount
	resp.EmployeeTypes = employeeTypes
	return nil
}

func (h hr) UpdateEmployeeType(ctx context.Context, req *go_micro_srv_hr.UpdateEmployeeTypeRequest, _ *go_micro_srv_hr.EmptyMessage) error {
	employeeType := go_micro_srv_hr.EmployeeType{
		Id:                   req.GetId(),
		Name:                 req.GetName(),
		DepartmentId:         req.GetDepartmentId(),
		CanLogin:             req.GetCanLogin(),
	}

	return h.dataSource.CreateEmployeeType(ctx, employeeType)
}

func (h hr) DeleteEmployeeType(ctx context.Context, req *go_micro_srv_hr.DeleteEmployeeTypeRequest, _ *go_micro_srv_hr.EmptyMessage) error {
	return h.dataSource.DeleteEmployeeType(ctx, req.GetId())
}

func (h hr) CreateEmployee(ctx context.Context, req *go_micro_srv_hr.CreateEmployeeRequest, resp *go_micro_srv_hr.CreateEmployeeResponse) error {
	employeeType, err := h.dataSource.EmployeeType(ctx, req.EmployeeTypeId)
	if err != nil {
		return fmt.Errorf("error in retreiving employee type, %s", err.Error())
	}

	if employeeType.CanLogin {
		createLoginReq := &acl.CreateUserRequest{
			Username:             req.GetUsername(),
			Password:             req.GetPassword(),
			Email:                req.GetEmail(),
			PhoneNumber:          req.GetPhoneNumber(),
			Name:                 req.GetName(),
			Role:                 employeeType.GetName(),
		}

		if _, err = h.aclService.CreateUser(ctx, createLoginReq); err != nil {
			return fmt.Errorf("error in creating user login, %s", err.Error())
		}
	}

	deleteUser := func() {
		if employeeType.CanLogin {
			if _, err = h.aclService.DeleteUser(ctx, &acl.DeleteUserRequest{Username: req.GetUsername()}); err != nil {
				log.Tracef("Error in deleting new user during failed employee creation, %s", err.Error())
			}
		}
	}

	id, err := uuid.NewV4()
	if err != nil {
		deleteUser()
		return fmt.Errorf("error in generating id, %s", err.Error())
	}

	employee := go_micro_srv_hr.Employee{
		Id:                   id.String(),
		Name:                 req.GetName(),
		DepartmentId:         req.GetDepartmentId(),
		EmployeeTypeId:       req.GetEmployeeTypeId(),
		Username:             req.GetUsername(),
		Email:                req.GetEmployeeTypeId(),
		PhoneNumber:          req.GetPhoneNumber(),
	}

	if err = h.dataSource.CreateEmployee(ctx, employee); err != nil {
		deleteUser()
		return err
	}

	resp.EmployeeId = id.String()
	return nil
}

func (h hr) Employees(ctx context.Context, req *go_micro_srv_hr.EmployeesRequest, resp *go_micro_srv_hr.EmployeesResponse) error {
	employess, totalCount, err := h.dataSource.Employees(ctx, req.GetDepartmentId(), req.GetEmployeeTypeId(), req.GetSkipCount(),
		req.GetMaxResultCount())
	if err != nil {
		return err
	}

	resp.Employees = employess
	resp.TotalCount = totalCount
	return nil
}

func (h hr) UpdateEmployee(ctx context.Context, req *go_micro_srv_hr.UpdateEmployeeRequest, resp *go_micro_srv_hr.EmptyMessage) error {
	employeeType, err := h.dataSource.EmployeeType(ctx, req.EmployeeTypeId)
	if err != nil {
		return fmt.Errorf("error in retreiving employee type, %s", err.Error())
	}

	employee, err := h.dataSource.Employee(ctx, req.GetId())
	if err != nil {
		return fmt.Errorf("error in retreiving old employee record, %s", err.Error())
	}

	if employeeType.CanLogin {
		updateUserReq := &acl.UpdateUserRequest{
			Username:             req.GetUsername(),
			Email:                req.GetEmail(),
			PhoneNumber:          req.GetPhoneNumber(),
			Name:                 req.GetName(),
		}

		if _, err = h.aclService.UpdateUser(ctx, updateUserReq); err != nil {
			return fmt.Errorf("error in updating user login, %s", err.Error())
		}
	}

	employee.Username = req.Username
	employee.EmployeeTypeId = req.GetEmployeeTypeId()
	employee.DepartmentId = req.GetDepartmentId()
	employee.Name = req.GetName()
	employee.PhoneNumber = req.GetPhoneNumber()
	employee.Email = req.GetEmail()

	return h.dataSource.UpdateEmployee(ctx, employee)
}

func (h hr) DeleteEmployee(ctx context.Context, req *go_micro_srv_hr.DeleteEmployeeRequest, resp *go_micro_srv_hr.EmptyMessage) error {
	employee, err := h.dataSource.Employee(ctx, req.GetEmployeeId())
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
			return err
		}
		return nil
	}

	if _, err := h.aclService.DeleteUser(ctx, &acl.DeleteUserRequest{Username:employee.Username}); err != nil {
		return fmt.Errorf("error in deleting user login, %s", err.Error())
	}

	return h.dataSource.DeleteEmployee(ctx, req.GetEmployeeId())
}



