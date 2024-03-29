// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Departments", testDepartments)
	t.Run("Employees", testEmployees)
	t.Run("EmployeeTypes", testEmployeeTypes)
}

func TestDelete(t *testing.T) {
	t.Run("Departments", testDepartmentsDelete)
	t.Run("Employees", testEmployeesDelete)
	t.Run("EmployeeTypes", testEmployeeTypesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Departments", testDepartmentsQueryDeleteAll)
	t.Run("Employees", testEmployeesQueryDeleteAll)
	t.Run("EmployeeTypes", testEmployeeTypesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Departments", testDepartmentsSliceDeleteAll)
	t.Run("Employees", testEmployeesSliceDeleteAll)
	t.Run("EmployeeTypes", testEmployeeTypesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Departments", testDepartmentsExists)
	t.Run("Employees", testEmployeesExists)
	t.Run("EmployeeTypes", testEmployeeTypesExists)
}

func TestFind(t *testing.T) {
	t.Run("Departments", testDepartmentsFind)
	t.Run("Employees", testEmployeesFind)
	t.Run("EmployeeTypes", testEmployeeTypesFind)
}

func TestBind(t *testing.T) {
	t.Run("Departments", testDepartmentsBind)
	t.Run("Employees", testEmployeesBind)
	t.Run("EmployeeTypes", testEmployeeTypesBind)
}

func TestOne(t *testing.T) {
	t.Run("Departments", testDepartmentsOne)
	t.Run("Employees", testEmployeesOne)
	t.Run("EmployeeTypes", testEmployeeTypesOne)
}

func TestAll(t *testing.T) {
	t.Run("Departments", testDepartmentsAll)
	t.Run("Employees", testEmployeesAll)
	t.Run("EmployeeTypes", testEmployeeTypesAll)
}

func TestCount(t *testing.T) {
	t.Run("Departments", testDepartmentsCount)
	t.Run("Employees", testEmployeesCount)
	t.Run("EmployeeTypes", testEmployeeTypesCount)
}

func TestInsert(t *testing.T) {
	t.Run("Departments", testDepartmentsInsert)
	t.Run("Departments", testDepartmentsInsertWhitelist)
	t.Run("Employees", testEmployeesInsert)
	t.Run("Employees", testEmployeesInsertWhitelist)
	t.Run("EmployeeTypes", testEmployeeTypesInsert)
	t.Run("EmployeeTypes", testEmployeeTypesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("EmployeeToDepartmentUsingDepartment", testEmployeeToOneDepartmentUsingDepartment)
	t.Run("EmployeeToEmployeeTypeUsingEmployeeType", testEmployeeToOneEmployeeTypeUsingEmployeeType)
	t.Run("EmployeeTypeToDepartmentUsingDepartment", testEmployeeTypeToOneDepartmentUsingDepartment)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DepartmentToEmployees", testDepartmentToManyEmployees)
	t.Run("DepartmentToEmployeeTypes", testDepartmentToManyEmployeeTypes)
	t.Run("EmployeeTypeToEmployees", testEmployeeTypeToManyEmployees)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("EmployeeToDepartmentUsingEmployees", testEmployeeToOneSetOpDepartmentUsingDepartment)
	t.Run("EmployeeToEmployeeTypeUsingEmployees", testEmployeeToOneSetOpEmployeeTypeUsingEmployeeType)
	t.Run("EmployeeTypeToDepartmentUsingEmployeeTypes", testEmployeeTypeToOneSetOpDepartmentUsingDepartment)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("DepartmentToEmployees", testDepartmentToManyAddOpEmployees)
	t.Run("DepartmentToEmployeeTypes", testDepartmentToManyAddOpEmployeeTypes)
	t.Run("EmployeeTypeToEmployees", testEmployeeTypeToManyAddOpEmployees)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Departments", testDepartmentsReload)
	t.Run("Employees", testEmployeesReload)
	t.Run("EmployeeTypes", testEmployeeTypesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Departments", testDepartmentsReloadAll)
	t.Run("Employees", testEmployeesReloadAll)
	t.Run("EmployeeTypes", testEmployeeTypesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Departments", testDepartmentsSelect)
	t.Run("Employees", testEmployeesSelect)
	t.Run("EmployeeTypes", testEmployeeTypesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Departments", testDepartmentsUpdate)
	t.Run("Employees", testEmployeesUpdate)
	t.Run("EmployeeTypes", testEmployeeTypesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Departments", testDepartmentsSliceUpdateAll)
	t.Run("Employees", testEmployeesSliceUpdateAll)
	t.Run("EmployeeTypes", testEmployeeTypesSliceUpdateAll)
}
