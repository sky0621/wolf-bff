// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboilermodel

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ContentText is an object representing the database table.
type ContentText struct {
	ID    int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	WHTID int64       `boil:"wht_id" json:"wht_id" toml:"wht_id" yaml:"wht_id"`
	Name  null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Text  string      `boil:"text" json:"text" toml:"text" yaml:"text"`

	R *contentTextR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L contentTextL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ContentTextColumns = struct {
	ID    string
	WHTID string
	Name  string
	Text  string
}{
	ID:    "id",
	WHTID: "wht_id",
	Name:  "name",
	Text:  "text",
}

// Generated where

var ContentTextWhere = struct {
	ID    whereHelperint64
	WHTID whereHelperint64
	Name  whereHelpernull_String
	Text  whereHelperstring
}{
	ID:    whereHelperint64{field: "\"content_text\".\"id\""},
	WHTID: whereHelperint64{field: "\"content_text\".\"wht_id\""},
	Name:  whereHelpernull_String{field: "\"content_text\".\"name\""},
	Text:  whereHelperstring{field: "\"content_text\".\"text\""},
}

// ContentTextRels is where relationship names are stored.
var ContentTextRels = struct {
	WHT string
}{
	WHT: "WHT",
}

// contentTextR is where relationships are stored.
type contentTextR struct {
	WHT *WHT `boil:"WHT" json:"WHT" toml:"WHT" yaml:"WHT"`
}

// NewStruct creates a new relationship struct
func (*contentTextR) NewStruct() *contentTextR {
	return &contentTextR{}
}

// contentTextL is where Load methods for each relationship are stored.
type contentTextL struct{}

var (
	contentTextAllColumns            = []string{"id", "wht_id", "name", "text"}
	contentTextColumnsWithoutDefault = []string{"wht_id", "name", "text"}
	contentTextColumnsWithDefault    = []string{"id"}
	contentTextPrimaryKeyColumns     = []string{"id"}
)

type (
	// ContentTextSlice is an alias for a slice of pointers to ContentText.
	// This should generally be used opposed to []ContentText.
	ContentTextSlice []*ContentText
	// ContentTextHook is the signature for custom ContentText hook methods
	ContentTextHook func(context.Context, boil.ContextExecutor, *ContentText) error

	contentTextQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	contentTextType                 = reflect.TypeOf(&ContentText{})
	contentTextMapping              = queries.MakeStructMapping(contentTextType)
	contentTextPrimaryKeyMapping, _ = queries.BindMapping(contentTextType, contentTextMapping, contentTextPrimaryKeyColumns)
	contentTextInsertCacheMut       sync.RWMutex
	contentTextInsertCache          = make(map[string]insertCache)
	contentTextUpdateCacheMut       sync.RWMutex
	contentTextUpdateCache          = make(map[string]updateCache)
	contentTextUpsertCacheMut       sync.RWMutex
	contentTextUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var contentTextBeforeInsertHooks []ContentTextHook
var contentTextBeforeUpdateHooks []ContentTextHook
var contentTextBeforeDeleteHooks []ContentTextHook
var contentTextBeforeUpsertHooks []ContentTextHook

var contentTextAfterInsertHooks []ContentTextHook
var contentTextAfterSelectHooks []ContentTextHook
var contentTextAfterUpdateHooks []ContentTextHook
var contentTextAfterDeleteHooks []ContentTextHook
var contentTextAfterUpsertHooks []ContentTextHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ContentText) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ContentText) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ContentText) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ContentText) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ContentText) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ContentText) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ContentText) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ContentText) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ContentText) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range contentTextAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddContentTextHook registers your hook function for all future operations.
func AddContentTextHook(hookPoint boil.HookPoint, contentTextHook ContentTextHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		contentTextBeforeInsertHooks = append(contentTextBeforeInsertHooks, contentTextHook)
	case boil.BeforeUpdateHook:
		contentTextBeforeUpdateHooks = append(contentTextBeforeUpdateHooks, contentTextHook)
	case boil.BeforeDeleteHook:
		contentTextBeforeDeleteHooks = append(contentTextBeforeDeleteHooks, contentTextHook)
	case boil.BeforeUpsertHook:
		contentTextBeforeUpsertHooks = append(contentTextBeforeUpsertHooks, contentTextHook)
	case boil.AfterInsertHook:
		contentTextAfterInsertHooks = append(contentTextAfterInsertHooks, contentTextHook)
	case boil.AfterSelectHook:
		contentTextAfterSelectHooks = append(contentTextAfterSelectHooks, contentTextHook)
	case boil.AfterUpdateHook:
		contentTextAfterUpdateHooks = append(contentTextAfterUpdateHooks, contentTextHook)
	case boil.AfterDeleteHook:
		contentTextAfterDeleteHooks = append(contentTextAfterDeleteHooks, contentTextHook)
	case boil.AfterUpsertHook:
		contentTextAfterUpsertHooks = append(contentTextAfterUpsertHooks, contentTextHook)
	}
}

// One returns a single contentText record from the query.
func (q contentTextQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ContentText, error) {
	o := &ContentText{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboilermodel: failed to execute a one query for content_text")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ContentText records from the query.
func (q contentTextQuery) All(ctx context.Context, exec boil.ContextExecutor) (ContentTextSlice, error) {
	var o []*ContentText

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "sqlboilermodel: failed to assign all query results to ContentText slice")
	}

	if len(contentTextAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ContentText records in the query.
func (q contentTextQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to count content_text rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q contentTextQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "sqlboilermodel: failed to check if content_text exists")
	}

	return count > 0, nil
}

// WHT pointed to by the foreign key.
func (o *ContentText) WHT(mods ...qm.QueryMod) whtQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.WHTID),
	}

	queryMods = append(queryMods, mods...)

	query := WHTS(queryMods...)
	queries.SetFrom(query.Query, "\"wht\"")

	return query
}

// LoadWHT allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (contentTextL) LoadWHT(ctx context.Context, e boil.ContextExecutor, singular bool, maybeContentText interface{}, mods queries.Applicator) error {
	var slice []*ContentText
	var object *ContentText

	if singular {
		object = maybeContentText.(*ContentText)
	} else {
		slice = *maybeContentText.(*[]*ContentText)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &contentTextR{}
		}
		args = append(args, object.WHTID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &contentTextR{}
			}

			for _, a := range args {
				if a == obj.WHTID {
					continue Outer
				}
			}

			args = append(args, obj.WHTID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`wht`),
		qm.WhereIn(`wht.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load WHT")
	}

	var resultSlice []*WHT
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice WHT")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for wht")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for wht")
	}

	if len(contentTextAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.WHT = foreign
		if foreign.R == nil {
			foreign.R = &whtR{}
		}
		foreign.R.ContentTexts = append(foreign.R.ContentTexts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.WHTID == foreign.ID {
				local.R.WHT = foreign
				if foreign.R == nil {
					foreign.R = &whtR{}
				}
				foreign.R.ContentTexts = append(foreign.R.ContentTexts, local)
				break
			}
		}
	}

	return nil
}

// SetWHT of the contentText to the related item.
// Sets o.R.WHT to related.
// Adds o to related.R.ContentTexts.
func (o *ContentText) SetWHT(ctx context.Context, exec boil.ContextExecutor, insert bool, related *WHT) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"content_text\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"wht_id"}),
		strmangle.WhereClause("\"", "\"", 2, contentTextPrimaryKeyColumns),
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

	o.WHTID = related.ID
	if o.R == nil {
		o.R = &contentTextR{
			WHT: related,
		}
	} else {
		o.R.WHT = related
	}

	if related.R == nil {
		related.R = &whtR{
			ContentTexts: ContentTextSlice{o},
		}
	} else {
		related.R.ContentTexts = append(related.R.ContentTexts, o)
	}

	return nil
}

// ContentTexts retrieves all the records using an executor.
func ContentTexts(mods ...qm.QueryMod) contentTextQuery {
	mods = append(mods, qm.From("\"content_text\""))
	return contentTextQuery{NewQuery(mods...)}
}

// FindContentText retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindContentText(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*ContentText, error) {
	contentTextObj := &ContentText{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"content_text\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, contentTextObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "sqlboilermodel: unable to select from content_text")
	}

	return contentTextObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ContentText) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboilermodel: no content_text provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contentTextColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	contentTextInsertCacheMut.RLock()
	cache, cached := contentTextInsertCache[key]
	contentTextInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			contentTextAllColumns,
			contentTextColumnsWithDefault,
			contentTextColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(contentTextType, contentTextMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(contentTextType, contentTextMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"content_text\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"content_text\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "sqlboilermodel: unable to insert into content_text")
	}

	if !cached {
		contentTextInsertCacheMut.Lock()
		contentTextInsertCache[key] = cache
		contentTextInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ContentText.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ContentText) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	contentTextUpdateCacheMut.RLock()
	cache, cached := contentTextUpdateCache[key]
	contentTextUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			contentTextAllColumns,
			contentTextPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("sqlboilermodel: unable to update content_text, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"content_text\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, contentTextPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(contentTextType, contentTextMapping, append(wl, contentTextPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "sqlboilermodel: unable to update content_text row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by update for content_text")
	}

	if !cached {
		contentTextUpdateCacheMut.Lock()
		contentTextUpdateCache[key] = cache
		contentTextUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q contentTextQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to update all for content_text")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to retrieve rows affected for content_text")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ContentTextSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("sqlboilermodel: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contentTextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"content_text\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, contentTextPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to update all in contentText slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to retrieve rows affected all in update all contentText")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ContentText) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("sqlboilermodel: no content_text provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(contentTextColumnsWithDefault, o)

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

	contentTextUpsertCacheMut.RLock()
	cache, cached := contentTextUpsertCache[key]
	contentTextUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			contentTextAllColumns,
			contentTextColumnsWithDefault,
			contentTextColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			contentTextAllColumns,
			contentTextPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("sqlboilermodel: unable to upsert content_text, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(contentTextPrimaryKeyColumns))
			copy(conflict, contentTextPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"content_text\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(contentTextType, contentTextMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(contentTextType, contentTextMapping, ret)
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
		return errors.Wrap(err, "sqlboilermodel: unable to upsert content_text")
	}

	if !cached {
		contentTextUpsertCacheMut.Lock()
		contentTextUpsertCache[key] = cache
		contentTextUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ContentText record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ContentText) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("sqlboilermodel: no ContentText provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), contentTextPrimaryKeyMapping)
	sql := "DELETE FROM \"content_text\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to delete from content_text")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by delete for content_text")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q contentTextQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("sqlboilermodel: no contentTextQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to delete all from content_text")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by deleteall for content_text")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ContentTextSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(contentTextBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contentTextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"content_text\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, contentTextPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: unable to delete all from contentText slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "sqlboilermodel: failed to get rows affected by deleteall for content_text")
	}

	if len(contentTextAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ContentText) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindContentText(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ContentTextSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ContentTextSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), contentTextPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"content_text\".* FROM \"content_text\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, contentTextPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "sqlboilermodel: unable to reload all in ContentTextSlice")
	}

	*o = slice

	return nil
}

// ContentTextExists checks if the ContentText row exists.
func ContentTextExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"content_text\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "sqlboilermodel: unable to check if content_text exists")
	}

	return exists, nil
}
