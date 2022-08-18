// Code generated by SQLBoiler 4.12.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testGorpMigrations(t *testing.T) {
	t.Parallel()

	query := GorpMigrations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGorpMigrationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
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

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGorpMigrationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := GorpMigrations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGorpMigrationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GorpMigrationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGorpMigrationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GorpMigrationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if GorpMigration exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GorpMigrationExists to return true, but got false.")
	}
}

func testGorpMigrationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	gorpMigrationFound, err := FindGorpMigration(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if gorpMigrationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGorpMigrationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = GorpMigrations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGorpMigrationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := GorpMigrations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGorpMigrationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gorpMigrationOne := &GorpMigration{}
	gorpMigrationTwo := &GorpMigration{}
	if err = randomize.Struct(seed, gorpMigrationOne, gorpMigrationDBTypes, false, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, gorpMigrationTwo, gorpMigrationDBTypes, false, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = gorpMigrationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = gorpMigrationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GorpMigrations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGorpMigrationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	gorpMigrationOne := &GorpMigration{}
	gorpMigrationTwo := &GorpMigration{}
	if err = randomize.Struct(seed, gorpMigrationOne, gorpMigrationDBTypes, false, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, gorpMigrationTwo, gorpMigrationDBTypes, false, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = gorpMigrationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = gorpMigrationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testGorpMigrationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGorpMigrationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(gorpMigrationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGorpMigrationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
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

func testGorpMigrationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GorpMigrationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGorpMigrationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := GorpMigrations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	gorpMigrationDBTypes = map[string]string{`ID`: `text`, `AppliedAt`: `timestamp with time zone`}
	_                    = bytes.MinRead
)

func testGorpMigrationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(gorpMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(gorpMigrationAllColumns) == len(gorpMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGorpMigrationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(gorpMigrationAllColumns) == len(gorpMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &GorpMigration{}
	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, gorpMigrationDBTypes, true, gorpMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(gorpMigrationAllColumns, gorpMigrationPrimaryKeyColumns) {
		fields = gorpMigrationAllColumns
	} else {
		fields = strmangle.SetComplement(
			gorpMigrationAllColumns,
			gorpMigrationPrimaryKeyColumns,
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

	slice := GorpMigrationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGorpMigrationsUpsert(t *testing.T) {
	t.Parallel()

	if len(gorpMigrationAllColumns) == len(gorpMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := GorpMigration{}
	if err = randomize.Struct(seed, &o, gorpMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GorpMigration: %s", err)
	}

	count, err := GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, gorpMigrationDBTypes, false, gorpMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GorpMigration struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert GorpMigration: %s", err)
	}

	count, err = GorpMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
