package store

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStoreSaveAndLoad(t *testing.T) {
	dbDir := os.TempDir()
	name := "test"
	r1 := NewReserve(dbDir, name)
	r2 := NewReserve(dbDir, name)

	Convey("A Store should be savable and loadable.", t, func() {
		So(r1.Save(newMockStore()), ShouldBeNil)
		So(fileExists(r1.path), ShouldBeTrue)
		So(r2.Load(), ShouldResemble, newMockStore())
		r1.Delete()
		So(fileExists(r1.path), ShouldBeFalse)
	})
	Convey("Reserve should handle invalid states", t, func() {
		os.RemoveAll(dbDir)
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
