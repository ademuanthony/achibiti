package handler

import (
	"context"
	"github.com/ademuanthony/achibiti/hr/proto/hr"
)

type hr struct{
	dataSource dataSource
}

func NewHr(dataSource dataSource) *hr {
	return &hr{dataSource: dataSource}
}

type dataSource interface {
	
} 

func (h hr) CreateDepartment(context.Context, *go_micro_srv_hr.CreateDepartmentRequest, *go_micro_srv_hr.CreateDepartmentResponse) error {
	panic("implement me")
}

func (h hr) Departments(context.Context, *go_micro_srv_hr.DepartmentsRequest, *go_micro_srv_hr.DepartmentsResponse) error {
	panic("implement me")
}

func (h hr) UpdateDepartment(context.Context, *go_micro_srv_hr.UpdateDepartmentRequest, *go_micro_srv_hr.EmptyMessage) error {
	panic("implement me")
}

func (h hr) DeleteDepartment(context.Context, *go_micro_srv_hr.DeleteDepartmentRequest, *go_micro_srv_hr.EmptyMessage) error {
	panic("implement me")
}

func (h hr) CreateEmployeeType(context.Context, *go_micro_srv_hr.CreateEmployeeTypeRequest, *go_micro_srv_hr.CreateEmployeeTypeResponse) error {
	panic("implement me")
}

func (h hr) EmployeeTypes(context.Context, *go_micro_srv_hr.EmployeeTypesRequest, *go_micro_srv_hr.EmployeeTypesResponse) error {
	panic("implement me")
}

func (h hr) UpdateEmployeeType(context.Context, *go_micro_srv_hr.UpdateEmployeeTypeRequest, *go_micro_srv_hr.EmptyMessage) error {
	panic("implement me")
}

func (h hr) DeleteEmployeeType(context.Context, *go_micro_srv_hr.DeleteEmployeeTypeRequest, *go_micro_srv_hr.EmptyMessage) error {
	panic("implement me")
}

func (h hr) CreateEmployee(context.Context, *go_micro_srv_hr.CreateEmployeeRequest, *go_micro_srv_hr.CreateEmployeeResponse) error {
	panic("implement me")
}

func (h hr) Employees(context.Context, *go_micro_srv_hr.EmployeesRequest, *go_micro_srv_hr.EmployeesResponse) error {
	panic("implement me")
}

func (h hr) UpdateEmployee(context.Context, *go_micro_srv_hr.UpdateEmployeeRequest, *go_micro_srv_hr.EmptyMessage) error {
	panic("implement me")
}

func (h hr) DeleteEmployee(context.Context, *go_micro_srv_hr.DeleteEmployeeRequest, *go_micro_srv_hr.EmptyMessage) error {
	panic("implement me")
}



