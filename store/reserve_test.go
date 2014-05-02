package store

import (
	"os"
	"testing"

	"github.com/casimir/storekeeper/util"
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
	// TODO mockup fake store to test transactions
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

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
