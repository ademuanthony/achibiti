// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEmployeeTypes(t *testing.T) {
	t.Parallel()

	query := EmployeeTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmployeeTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmployeeTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EmployeeTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmployeeTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmployeeTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmployeeTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmployeeTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EmployeeType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmployeeTypeExists to return true, but got false.")
	}
}

func testEmployeeTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	employeeTypeFound, err := FindEmployeeType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if employeeTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmployeeTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EmployeeTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmployeeTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EmployeeTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmployeeTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	employeeTypeOne := &EmployeeType{}
	employeeTypeTwo := &EmployeeType{}
	if err = randomize.Struct(seed, employeeTypeOne, employeeTypeDBTypes, false, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}
	if err = randomize.Struct(seed, employeeTypeTwo, employeeTypeDBTypes, false, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = employeeTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = employeeTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmployeeTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmployeeTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	employeeTypeOne := &EmployeeType{}
	employeeTypeTwo := &EmployeeType{}
	if err = randomize.Struct(seed, employeeTypeOne, employeeTypeDBTypes, false, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}
	if err = randomize.Struct(seed, employeeTypeTwo, employeeTypeDBTypes, false, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = employeeTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = employeeTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testEmployeeTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmployeeTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(employeeTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmployeeTypeToManyEmployees(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EmployeeType
	var b, c Employee

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.EmployeeTypeID = a.ID
	c.EmployeeTypeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Employees().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.EmployeeTypeID == b.EmployeeTypeID {
			bFound = true
		}
		if v.EmployeeTypeID == c.EmployeeTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EmployeeTypeSlice{&a}
	if err = a.L.LoadEmployees(ctx, tx, false, (*[]*EmployeeType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Employees); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Employees = nil
	if err = a.L.LoadEmployees(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Employees); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testEmployeeTypeToManyAddOpEmployees(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EmployeeType
	var b, c, d, e Employee

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, employeeTypeDBTypes, false, strmangle.SetComplement(employeeTypePrimaryKeyColumns, employeeTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Employee{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, employeeDBTypes, false, strmangle.SetComplement(employeePrimaryKeyColumns, employeeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Employee{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEmployees(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.EmployeeTypeID {
			t.Error("foreign key was wrong value", a.ID, first.EmployeeTypeID)
		}
		if a.ID != second.EmployeeTypeID {
			t.Error("foreign key was wrong value", a.ID, second.EmployeeTypeID)
		}

		if first.R.EmployeeType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.EmployeeType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Employees[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Employees[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Employees().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testEmployeeTypeToOneDepartmentUsingDepartment(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EmployeeType
	var foreign Department

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, employeeTypeDBTypes, false, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, departmentDBTypes, false, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DepartmentID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Department().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := EmployeeTypeSlice{&local}
	if err = local.L.LoadDepartment(ctx, tx, false, (*[]*EmployeeType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Department == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Department = nil
	if err = local.L.LoadDepartment(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Department == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testEmployeeTypeToOneSetOpDepartmentUsingDepartment(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EmployeeType
	var b, c Department

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, employeeTypeDBTypes, false, strmangle.SetComplement(employeeTypePrimaryKeyColumns, employeeTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, departmentDBTypes, false, strmangle.SetComplement(departmentPrimaryKeyColumns, departmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, departmentDBTypes, false, strmangle.SetComplement(departmentPrimaryKeyColumns, departmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Department{&b, &c} {
		err = a.SetDepartment(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Department != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EmployeeTypes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DepartmentID != x.ID {
			t.Error("foreign key was wrong value", a.DepartmentID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DepartmentID))
		reflect.Indirect(reflect.ValueOf(&a.DepartmentID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DepartmentID != x.ID {
			t.Error("foreign key was wrong value", a.DepartmentID, x.ID)
		}
	}
}

func testEmployeeTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmployeeTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmployeeTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmployeeTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EmployeeTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	employeeTypeDBTypes = map[string]string{`ID`: `character varying`, `DepartmentID`: `character varying`, `Name`: `character varying`}
	_                   = bytes.MinRead
)

func testEmployeeTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(employeeTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(employeeTypeAllColumns) == len(employeeTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmployeeTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(employeeTypeAllColumns) == len(employeeTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EmployeeType{}
	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, employeeTypeDBTypes, true, employeeTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(employeeTypeAllColumns, employeeTypePrimaryKeyColumns) {
		fields = employeeTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			employeeTypeAllColumns,
			employeeTypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EmployeeTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmployeeTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(employeeTypeAllColumns) == len(employeeTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EmployeeType{}
	if err = randomize.Struct(seed, &o, employeeTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmployeeType: %s", err)
	}

	count, err := EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, employeeTypeDBTypes, false, employeeTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EmployeeType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EmployeeType: %s", err)
	}

	count, err = EmployeeTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
