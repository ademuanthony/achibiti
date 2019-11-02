// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// EmployeeType is an object representing the database table.
type EmployeeType struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	DepartmentID string      `boil:"department_id" json:"department_id" toml:"department_id" yaml:"department_id"`
	Name         null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *employeeTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L employeeTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmployeeTypeColumns = struct {
	ID           string
	DepartmentID string
	Name         string
}{
	ID:           "id",
	DepartmentID: "department_id",
	Name:         "name",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var EmployeeTypeWhere = struct {
	ID           whereHelperstring
	DepartmentID whereHelperstring
	Name         whereHelpernull_String
}{
	ID:           whereHelperstring{field: "\"employee_type\".\"id\""},
	DepartmentID: whereHelperstring{field: "\"employee_type\".\"department_id\""},
	Name:         whereHelpernull_String{field: "\"employee_type\".\"name\""},
}

// EmployeeTypeRels is where relationship names are stored.
var EmployeeTypeRels = struct {
	Department string
	Employees  string
}{
	Department: "Department",
	Employees:  "Employees",
}

// employeeTypeR is where relationships are stored.
type employeeTypeR struct {
	Department *Department
	Employees  EmployeeSlice
}

// NewStruct creates a new relationship struct
func (*employeeTypeR) NewStruct() *employeeTypeR {
	return &employeeTypeR{}
}

// employeeTypeL is where Load methods for each relationship are stored.
type employeeTypeL struct{}

var (
	employeeTypeAllColumns            = []string{"id", "department_id", "name"}
	employeeTypeColumnsWithoutDefault = []string{"id", "department_id", "name"}
	employeeTypeColumnsWithDefault    = []string{}
	employeeTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// EmployeeTypeSlice is an alias for a slice of pointers to EmployeeType.
	// This should generally be used opposed to []EmployeeType.
	EmployeeTypeSlice []*EmployeeType

	employeeTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	employeeTypeType                 = reflect.TypeOf(&EmployeeType{})
	employeeTypeMapping              = queries.MakeStructMapping(employeeTypeType)
	employeeTypePrimaryKeyMapping, _ = queries.BindMapping(employeeTypeType, employeeTypeMapping, employeeTypePrimaryKeyColumns)
	employeeTypeInsertCacheMut       sync.RWMutex
	employeeTypeInsertCache          = make(map[string]insertCache)
	employeeTypeUpdateCacheMut       sync.RWMutex
	employeeTypeUpdateCache          = make(map[string]updateCache)
	employeeTypeUpsertCacheMut       sync.RWMutex
	employeeTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single employeeType record from the query.
func (q employeeTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EmployeeType, error) {
	o := &EmployeeType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for employee_type")
	}

	return o, nil
}

// All returns all EmployeeType records from the query.
func (q employeeTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (EmployeeTypeSlice, error) {
	var o []*EmployeeType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EmployeeType slice")
	}

	return o, nil
}

// Count returns the count of all EmployeeType records in the query.
func (q employeeTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count employee_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q employeeTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if employee_type exists")
	}

	return count > 0, nil
}

// Department pointed to by the foreign key.
func (o *EmployeeType) Department(mods ...qm.QueryMod) departmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DepartmentID),
	}

	queryMods = append(queryMods, mods...)

	query := Departments(queryMods...)
	queries.SetFrom(query.Query, "\"department\"")

	return query
}

// Employees retrieves all the employee's Employees with an executor.
func (o *EmployeeType) Employees(mods ...qm.QueryMod) employeeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"employee\".\"employee_type_id\"=?", o.ID),
	)

	query := Employees(queryMods...)
	queries.SetFrom(query.Query, "\"employee\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"employee\".*"})
	}

	return query
}

// LoadDepartment allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (employeeTypeL) LoadDepartment(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmployeeType interface{}, mods queries.Applicator) error {
	var slice []*EmployeeType
	var object *EmployeeType

	if singular {
		object = maybeEmployeeType.(*EmployeeType)
	} else {
		slice = *maybeEmployeeType.(*[]*EmployeeType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &employeeTypeR{}
		}
		args = append(args, object.DepartmentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &employeeTypeR{}
			}

			for _, a := range args {
				if a == obj.DepartmentID {
					continue Outer
				}
			}

			args = append(args, obj.DepartmentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`department`), qm.WhereIn(`department.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Department")
	}

	var resultSlice []*Department
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Department")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for department")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for department")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Department = foreign
		if foreign.R == nil {
			foreign.R = &departmentR{}
		}
		foreign.R.EmployeeTypes = append(foreign.R.EmployeeTypes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DepartmentID == foreign.ID {
				local.R.Department = foreign
				if foreign.R == nil {
					foreign.R = &departmentR{}
				}
				foreign.R.EmployeeTypes = append(foreign.R.EmployeeTypes, local)
				break
			}
		}
	}

	return nil
}

// LoadEmployees allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (employeeTypeL) LoadEmployees(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmployeeType interface{}, mods queries.Applicator) error {
	var slice []*EmployeeType
	var object *EmployeeType

	if singular {
		object = maybeEmployeeType.(*EmployeeType)
	} else {
		slice = *maybeEmployeeType.(*[]*EmployeeType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &employeeTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &employeeTypeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`employee`), qm.WhereIn(`employee.employee_type_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load employee")
	}

	var resultSlice []*Employee
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice employee")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on employee")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for employee")
	}

	if singular {
		object.R.Employees = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &employeeR{}
			}
			foreign.R.EmployeeType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.EmployeeTypeID {
				local.R.Employees = append(local.R.Employees, foreign)
				if foreign.R == nil {
					foreign.R = &employeeR{}
				}
				foreign.R.EmployeeType = local
				break
			}
		}
	}

	return nil
}

// SetDepartment of the employeeType to the related item.
// Sets o.R.Department to related.
// Adds o to related.R.EmployeeTypes.
func (o *EmployeeType) SetDepartment(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Department) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"employee_type\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"department_id"}),
		strmangle.WhereClause("\"", "\"", 2, employeeTypePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DepartmentID = related.ID
	if o.R == nil {
		o.R = &employeeTypeR{
			Department: related,
		}
	} else {
		o.R.Department = related
	}

	if related.R == nil {
		related.R = &departmentR{
			EmployeeTypes: EmployeeTypeSlice{o},
		}
	} else {
		related.R.EmployeeTypes = append(related.R.EmployeeTypes, o)
	}

	return nil
}

// AddEmployees adds the given related objects to the existing relationships
// of the employee_type, optionally inserting them as new records.
// Appends related to o.R.Employees.
// Sets related.R.EmployeeType appropriately.
func (o *EmployeeType) AddEmployees(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Employee) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EmployeeTypeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"employee\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"employee_type_id"}),
				strmangle.WhereClause("\"", "\"", 2, employeePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EmployeeTypeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &employeeTypeR{
			Employees: related,
		}
	} else {
		o.R.Employees = append(o.R.Employees, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &employeeR{
				EmployeeType: o,
			}
		} else {
			rel.R.EmployeeType = o
		}
	}
	return nil
}

// EmployeeTypes retrieves all the records using an executor.
func EmployeeTypes(mods ...qm.QueryMod) employeeTypeQuery {
	mods = append(mods, qm.From("\"employee_type\""))
	return employeeTypeQuery{NewQuery(mods...)}
}

// FindEmployeeType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmployeeType(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*EmployeeType, error) {
	employeeTypeObj := &EmployeeType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"employee_type\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, employeeTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from employee_type")
	}

	return employeeTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EmployeeType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no employee_type provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(employeeTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	employeeTypeInsertCacheMut.RLock()
	cache, cached := employeeTypeInsertCache[key]
	employeeTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			employeeTypeAllColumns,
			employeeTypeColumnsWithDefault,
			employeeTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(employeeTypeType, employeeTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(employeeTypeType, employeeTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"employee_type\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"employee_type\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into employee_type")
	}

	if !cached {
		employeeTypeInsertCacheMut.Lock()
		employeeTypeInsertCache[key] = cache
		employeeTypeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the EmployeeType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EmployeeType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	employeeTypeUpdateCacheMut.RLock()
	cache, cached := employeeTypeUpdateCache[key]
	employeeTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			employeeTypeAllColumns,
			employeeTypePrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update employee_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"employee_type\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, employeeTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(employeeTypeType, employeeTypeMapping, append(wl, employeeTypePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update employee_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for employee_type")
	}

	if !cached {
		employeeTypeUpdateCacheMut.Lock()
		employeeTypeUpdateCache[key] = cache
		employeeTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q employeeTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for employee_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for employee_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmployeeTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), employeeTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"employee_type\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, employeeTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in employeeType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all employeeType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EmployeeType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no employee_type provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(employeeTypeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	employeeTypeUpsertCacheMut.RLock()
	cache, cached := employeeTypeUpsertCache[key]
	employeeTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			employeeTypeAllColumns,
			employeeTypeColumnsWithDefault,
			employeeTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			employeeTypeAllColumns,
			employeeTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert employee_type, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(employeeTypePrimaryKeyColumns))
			copy(conflict, employeeTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"employee_type\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(employeeTypeType, employeeTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(employeeTypeType, employeeTypeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert employee_type")
	}

	if !cached {
		employeeTypeUpsertCacheMut.Lock()
		employeeTypeUpsertCache[key] = cache
		employeeTypeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single EmployeeType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EmployeeType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EmployeeType provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), employeeTypePrimaryKeyMapping)
	sql := "DELETE FROM \"employee_type\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from employee_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for employee_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q employeeTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no employeeTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from employee_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for employee_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmployeeTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), employeeTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"employee_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, employeeTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from employeeType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for employee_type")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EmployeeType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEmployeeType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmployeeTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmployeeTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), employeeTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"employee_type\".* FROM \"employee_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, employeeTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EmployeeTypeSlice")
	}

	*o = slice

	return nil
}

// EmployeeTypeExists checks if the EmployeeType row exists.
func EmployeeTypeExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"employee_type\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if employee_type exists")
	}

	return exists, nil
}
