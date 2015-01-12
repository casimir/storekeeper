package store

import (
	"os"
	"testing"

	"github.com/casimir/storekeeper/storage"
	"github.com/casimir/storekeeper/util"
	"github.com/coopernurse/gorp"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReservePath(t *testing.T) {
	Convey("Given a store name", t, func() {
		Convey("It should return the DB path", func() {
			// Bad coupling but it's an utility package after all
			So(dbPath("name"), ShouldEqual, util.ApplicationPath()+"/db/name.db")
		})
	})
}

func TestReserveLifecycle(t *testing.T) {
	name := "test"
	Convey("Given a store name", t, func() {
		Convey("It should be able to create the reserve", func() {
			r := NewReserve(name)
			So(fileExists(dbPath(name)), ShouldBeTrue)
			r.dbm.Db.Close()
		})
		Convey("It should be able to delete the reserve", func() {
			DeleteReserve(name)
			So(fileExists(dbPath(name)), ShouldBeFalse)
		})
	})
}

func TestStoreSAveAndLoad(t *testing.T) {
	// TODO mockup fake store to test transactions
}

func TestDBConversion(t *testing.T) {
	jsonEmpty := "[]"
	sliceEmpty := []storage.Stack{}
	json := "[{\"Count\":1,\"Item\":{\"ID\":\"id1\",\"Name\":\"id1\"}},{\"Count\":2,\"Item\":{\"ID\":\"id2\",\"Name\":\"id2\"}}]"
	slice := []storage.Stack{{1, storage.Item{"id1", "id1"}}, {2, storage.Item{"id2", "id2"}}}

	resEmpty := stackToDb(sliceEmpty)
	res := stackToDb(slice)

	Convey("Given a slice of Stacks", t, func() {
		Convey("It should serialize the slice to store in DB", func() {
			So(typeToDb(1), ShouldEqual, 1)
			So(resEmpty, ShouldEqual, jsonEmpty)
			So(res, ShouldEqual, json)
		})
		Convey("It should deserialize the value in DB to a slice", func() {
			So(dbToType(1), ShouldResemble, gorp.CustomScanner{})
			So(dbToStack(resEmpty, sliceEmpty), ShouldBeEmpty)
			So(dbToStack(res, slice), ShouldResemble, slice)
		})
	})
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func dbToStack(json string, out []storage.Stack) []storage.Stack {
	c := converter{}
	scanner, ok := c.FromDb(&out)
	if !ok {
		return nil
	}
	if err := scanner.Binder(&json, &out); err != nil {
		return nil
	}
	return out
}

func dbToType(target interface{}) gorp.CustomScanner {
	c := converter{}
	scanner, _ := c.FromDb(target)
	return scanner
}

func stackToDb(s []storage.Stack) string {
	c := converter{}
	json, err := c.ToDb(s)
	if err != nil {
		return err.Error()
	}
	return json.(string)
}

func typeToDb(val interface{}) interface{} {
	c := converter{}
	ret, _ := c.ToDb(val)
	return ret
}
