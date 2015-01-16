package store

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStoreSaveAndLoad(t *testing.T) {
	dbDir = "/tmp/data"
	name := "test"
	r1 := NewReserve(name)
	r2 := NewReserve(name)

	Convey("A Store should be savable and loadable.", t, func() {
		So(r1.Save(newMockStore()), ShouldBeNil)
		So(fileExists(dbPath(name)), ShouldBeTrue)
		So(r2.Load(), ShouldResemble, newMockStore())
		DeleteReserve(name)
		So(fileExists(dbPath(name)), ShouldBeFalse)
	})
	Convey("Reserve should handle invalid states", t, func() {
		os.Remove(dbDir)
		So(r1.Save(newMockStore()), ShouldNotBeNil)
		So(r1.Load(), ShouldResemble, &Store{})
	})
}

func newMockStore() *Store {
	artisans := []Artisan{
		Artisan{ID: "id1", Label: "artisan1"},
		Artisan{ID: "id2", Label: "artisan2"},
	}
	return &Store{Artisans: artisans}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
