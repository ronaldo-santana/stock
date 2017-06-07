// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testTaxClasses(t *testing.T) {
	t.Parallel()

	query := TaxClasses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTaxClassesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = taxClass.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaxClassesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TaxClasses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaxClassesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TaxClassSlice{taxClass}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTaxClassesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TaxClassExists(tx, taxClass.ID)
	if err != nil {
		t.Errorf("Unable to check if TaxClass exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaxClassExistsG to return true, but got false.")
	}
}
func testTaxClassesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	taxClassFound, err := FindTaxClass(tx, taxClass.ID)
	if err != nil {
		t.Error(err)
	}

	if taxClassFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTaxClassesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TaxClasses(tx).Bind(taxClass); err != nil {
		t.Error(err)
	}
}

func testTaxClassesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := TaxClasses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTaxClassesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClassOne := &TaxClass{}
	taxClassTwo := &TaxClass{}
	if err = randomize.Struct(seed, taxClassOne, taxClassDBTypes, false, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}
	if err = randomize.Struct(seed, taxClassTwo, taxClassDBTypes, false, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClassOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = taxClassTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TaxClasses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTaxClassesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taxClassOne := &TaxClass{}
	taxClassTwo := &TaxClass{}
	if err = randomize.Struct(seed, taxClassOne, taxClassDBTypes, false, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}
	if err = randomize.Struct(seed, taxClassTwo, taxClassDBTypes, false, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClassOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = taxClassTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func taxClassBeforeInsertHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassAfterInsertHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassAfterSelectHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassBeforeUpdateHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassAfterUpdateHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassBeforeDeleteHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassAfterDeleteHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassBeforeUpsertHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func taxClassAfterUpsertHook(e boil.Executor, o *TaxClass) error {
	*o = TaxClass{}
	return nil
}

func testTaxClassesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &TaxClass{}
	o := &TaxClass{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taxClassDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TaxClass object: %s", err)
	}

	AddTaxClassHook(boil.BeforeInsertHook, taxClassBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taxClassBeforeInsertHooks = []TaxClassHook{}

	AddTaxClassHook(boil.AfterInsertHook, taxClassAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taxClassAfterInsertHooks = []TaxClassHook{}

	AddTaxClassHook(boil.AfterSelectHook, taxClassAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taxClassAfterSelectHooks = []TaxClassHook{}

	AddTaxClassHook(boil.BeforeUpdateHook, taxClassBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taxClassBeforeUpdateHooks = []TaxClassHook{}

	AddTaxClassHook(boil.AfterUpdateHook, taxClassAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taxClassAfterUpdateHooks = []TaxClassHook{}

	AddTaxClassHook(boil.BeforeDeleteHook, taxClassBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taxClassBeforeDeleteHooks = []TaxClassHook{}

	AddTaxClassHook(boil.AfterDeleteHook, taxClassAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taxClassAfterDeleteHooks = []TaxClassHook{}

	AddTaxClassHook(boil.BeforeUpsertHook, taxClassBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taxClassBeforeUpsertHooks = []TaxClassHook{}

	AddTaxClassHook(boil.AfterUpsertHook, taxClassAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taxClassAfterUpsertHooks = []TaxClassHook{}
}
func testTaxClassesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaxClassesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx, taxClassColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaxClassToManyProducts(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TaxClass
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, productDBTypes, false, productColumnsWithDefault...)
	randomize.Struct(seed, &c, productDBTypes, false, productColumnsWithDefault...)

	b.TaxClassID = a.ID
	c.TaxClassID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	product, err := a.Products(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range product {
		if v.TaxClassID == b.TaxClassID {
			bFound = true
		}
		if v.TaxClassID == c.TaxClassID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TaxClassSlice{&a}
	if err = a.L.LoadProducts(tx, false, (*[]*TaxClass)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Products); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Products = nil
	if err = a.L.LoadProducts(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Products); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", product)
	}
}

func testTaxClassToManyAddOpProducts(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a TaxClass
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taxClassDBTypes, false, strmangle.SetComplement(taxClassPrimaryKeyColumns, taxClassColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Product{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProducts(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TaxClassID {
			t.Error("foreign key was wrong value", a.ID, first.TaxClassID)
		}
		if a.ID != second.TaxClassID {
			t.Error("foreign key was wrong value", a.ID, second.TaxClassID)
		}

		if first.R.TaxClass != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.TaxClass != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Products[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Products[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Products(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTaxClassesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = taxClass.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTaxClassesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TaxClassSlice{taxClass}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTaxClassesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TaxClasses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taxClassDBTypes = map[string]string{`CreatedAt`: `timestamp`, `Description`: `varchar`, `ID`: `int`, `Title`: `varchar`, `UpdatedAt`: `timestamp`}
	_               = bytes.MinRead
)

func testTaxClassesUpdate(t *testing.T) {
	t.Parallel()

	if len(taxClassColumns) == len(taxClassPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	if err = taxClass.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTaxClassesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taxClassColumns) == len(taxClassPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	taxClass := &TaxClass{}
	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, taxClass, taxClassDBTypes, true, taxClassPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taxClassColumns, taxClassPrimaryKeyColumns) {
		fields = taxClassColumns
	} else {
		fields = strmangle.SetComplement(
			taxClassColumns,
			taxClassPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(taxClass))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TaxClassSlice{taxClass}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTaxClassesUpsert(t *testing.T) {
	t.Parallel()

	if len(taxClassColumns) == len(taxClassPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	taxClass := TaxClass{}
	if err = randomize.Struct(seed, &taxClass, taxClassDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taxClass.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert TaxClass: %s", err)
	}

	count, err := TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &taxClass, taxClassDBTypes, false, taxClassPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaxClass struct: %s", err)
	}

	if err = taxClass.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert TaxClass: %s", err)
	}

	count, err = TaxClasses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}