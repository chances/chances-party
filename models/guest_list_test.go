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

func testGuestLists(t *testing.T) {
	t.Parallel()

	query := GuestLists(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGuestListsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = guestList.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuestListsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = GuestLists(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuestListsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GuestListSlice{guestList}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGuestListsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GuestListExists(tx, guestList.ID)
	if err != nil {
		t.Errorf("Unable to check if GuestList exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GuestListExistsG to return true, but got false.")
	}
}
func testGuestListsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	guestListFound, err := FindGuestList(tx, guestList.ID)
	if err != nil {
		t.Error(err)
	}

	if guestListFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGuestListsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = GuestLists(tx).Bind(guestList); err != nil {
		t.Error(err)
	}
}

func testGuestListsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := GuestLists(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGuestListsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestListOne := &GuestList{}
	guestListTwo := &GuestList{}
	if err = randomize.Struct(seed, guestListOne, guestListDBTypes, false, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}
	if err = randomize.Struct(seed, guestListTwo, guestListDBTypes, false, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestListOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = guestListTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := GuestLists(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGuestListsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	guestListOne := &GuestList{}
	guestListTwo := &GuestList{}
	if err = randomize.Struct(seed, guestListOne, guestListDBTypes, false, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}
	if err = randomize.Struct(seed, guestListTwo, guestListDBTypes, false, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestListOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = guestListTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testGuestListsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuestListsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx, guestListColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuestListToManyGuestParties(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a GuestList
	var b, c Party

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, partyDBTypes, false, partyColumnsWithDefault...)
	randomize.Struct(seed, &c, partyDBTypes, false, partyColumnsWithDefault...)

	b.GuestsID = a.ID
	c.GuestsID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	party, err := a.GuestParties(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range party {
		if v.GuestsID == b.GuestsID {
			bFound = true
		}
		if v.GuestsID == c.GuestsID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GuestListSlice{&a}
	if err = a.L.LoadGuestParties(tx, false, (*[]*GuestList)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuestParties); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.GuestParties = nil
	if err = a.L.LoadGuestParties(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.GuestParties); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", party)
	}
}

func testGuestListToManyAddOpGuestParties(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a GuestList
	var b, c, d, e Party

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guestListDBTypes, false, strmangle.SetComplement(guestListPrimaryKeyColumns, guestListColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Party{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, partyDBTypes, false, strmangle.SetComplement(partyPrimaryKeyColumns, partyColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Party{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddGuestParties(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuestsID {
			t.Error("foreign key was wrong value", a.ID, first.GuestsID)
		}
		if a.ID != second.GuestsID {
			t.Error("foreign key was wrong value", a.ID, second.GuestsID)
		}

		if first.R.Guest != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guest != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.GuestParties[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.GuestParties[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.GuestParties(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testGuestListsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = guestList.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGuestListsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GuestListSlice{guestList}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGuestListsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := GuestLists(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	guestListDBTypes = map[string]string{`Data`: `json`, `ID`: `integer`}
	_                = bytes.MinRead
)

func testGuestListsUpdate(t *testing.T) {
	t.Parallel()

	if len(guestListColumns) == len(guestListPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	if err = guestList.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGuestListsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(guestListColumns) == len(guestListPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	guestList := &GuestList{}
	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, guestList, guestListDBTypes, true, guestListPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(guestListColumns, guestListPrimaryKeyColumns) {
		fields = guestListColumns
	} else {
		fields = strmangle.SetComplement(
			guestListColumns,
			guestListPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(guestList))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GuestListSlice{guestList}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGuestListsUpsert(t *testing.T) {
	t.Parallel()

	if len(guestListColumns) == len(guestListPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	guestList := GuestList{}
	if err = randomize.Struct(seed, &guestList, guestListDBTypes, true); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = guestList.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert GuestList: %s", err)
	}

	count, err := GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &guestList, guestListDBTypes, false, guestListPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize GuestList struct: %s", err)
	}

	if err = guestList.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert GuestList: %s", err)
	}

	count, err = GuestLists(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
