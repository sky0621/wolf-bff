// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboilermodel

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testContentImages(t *testing.T) {
	t.Parallel()

	query := ContentImages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testContentImagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
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

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentImagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ContentImages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentImagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContentImageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContentImagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ContentImageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ContentImage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContentImageExists to return true, but got false.")
	}
}

func testContentImagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	contentImageFound, err := FindContentImage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if contentImageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testContentImagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ContentImages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testContentImagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ContentImages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContentImagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	contentImageOne := &ContentImage{}
	contentImageTwo := &ContentImage{}
	if err = randomize.Struct(seed, contentImageOne, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}
	if err = randomize.Struct(seed, contentImageTwo, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contentImageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contentImageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ContentImages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContentImagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	contentImageOne := &ContentImage{}
	contentImageTwo := &ContentImage{}
	if err = randomize.Struct(seed, contentImageOne, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}
	if err = randomize.Struct(seed, contentImageTwo, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = contentImageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = contentImageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func contentImageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func contentImageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ContentImage) error {
	*o = ContentImage{}
	return nil
}

func testContentImagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ContentImage{}
	o := &ContentImage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, contentImageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ContentImage object: %s", err)
	}

	AddContentImageHook(boil.BeforeInsertHook, contentImageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	contentImageBeforeInsertHooks = []ContentImageHook{}

	AddContentImageHook(boil.AfterInsertHook, contentImageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	contentImageAfterInsertHooks = []ContentImageHook{}

	AddContentImageHook(boil.AfterSelectHook, contentImageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	contentImageAfterSelectHooks = []ContentImageHook{}

	AddContentImageHook(boil.BeforeUpdateHook, contentImageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	contentImageBeforeUpdateHooks = []ContentImageHook{}

	AddContentImageHook(boil.AfterUpdateHook, contentImageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	contentImageAfterUpdateHooks = []ContentImageHook{}

	AddContentImageHook(boil.BeforeDeleteHook, contentImageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	contentImageBeforeDeleteHooks = []ContentImageHook{}

	AddContentImageHook(boil.AfterDeleteHook, contentImageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	contentImageAfterDeleteHooks = []ContentImageHook{}

	AddContentImageHook(boil.BeforeUpsertHook, contentImageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	contentImageBeforeUpsertHooks = []ContentImageHook{}

	AddContentImageHook(boil.AfterUpsertHook, contentImageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	contentImageAfterUpsertHooks = []ContentImageHook{}
}

func testContentImagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentImagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(contentImageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContentImageToOneWHTUsingWHT(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ContentImage
	var foreign WHT

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, contentImageDBTypes, false, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, whtDBTypes, false, whtColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WHT struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.WHTID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.WHT().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ContentImageSlice{&local}
	if err = local.L.LoadWHT(ctx, tx, false, (*[]*ContentImage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WHT == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.WHT = nil
	if err = local.L.LoadWHT(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.WHT == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testContentImageToOneSetOpWHTUsingWHT(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ContentImage
	var b, c WHT

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, contentImageDBTypes, false, strmangle.SetComplement(contentImagePrimaryKeyColumns, contentImageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, whtDBTypes, false, strmangle.SetComplement(whtPrimaryKeyColumns, whtColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, whtDBTypes, false, strmangle.SetComplement(whtPrimaryKeyColumns, whtColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*WHT{&b, &c} {
		err = a.SetWHT(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.WHT != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ContentImages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.WHTID != x.ID {
			t.Error("foreign key was wrong value", a.WHTID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.WHTID))
		reflect.Indirect(reflect.ValueOf(&a.WHTID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.WHTID != x.ID {
			t.Error("foreign key was wrong value", a.WHTID, x.ID)
		}
	}
}

func testContentImagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
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

func testContentImagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ContentImageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testContentImagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ContentImages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	contentImageDBTypes = map[string]string{`ID`: `bigint`, `WHTID`: `bigint`, `Name`: `character varying`, `Path`: `character varying`, `CreatedBy`: `character varying`, `CreatedAt`: `timestamp with time zone`, `UpdatedBy`: `character varying`, `UpdatedAt`: `timestamp with time zone`}
	_                   = bytes.MinRead
)

func testContentImagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(contentImagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(contentImageAllColumns) == len(contentImagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testContentImagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(contentImageAllColumns) == len(contentImagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ContentImage{}
	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, contentImageDBTypes, true, contentImagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(contentImageAllColumns, contentImagePrimaryKeyColumns) {
		fields = contentImageAllColumns
	} else {
		fields = strmangle.SetComplement(
			contentImageAllColumns,
			contentImagePrimaryKeyColumns,
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

	slice := ContentImageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testContentImagesUpsert(t *testing.T) {
	t.Parallel()

	if len(contentImageAllColumns) == len(contentImagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ContentImage{}
	if err = randomize.Struct(seed, &o, contentImageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ContentImage: %s", err)
	}

	count, err := ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, contentImageDBTypes, false, contentImagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ContentImage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ContentImage: %s", err)
	}

	count, err = ContentImages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
