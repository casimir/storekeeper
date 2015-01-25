package store

import (
	"os"
	"testing"

	"github.com/casimir/doable"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStoreSaveAndLoad(t *testing.T) {
	dbDir := os.TempDir() + "/storekeeper_test/"
	name := "test"
	expected := newMockStore()
	r1 := NewReserve(dbDir, name)
	r2 := NewReserve(dbDir, name)

	Convey("A Store could be saved and loaded.", t, func() {
		So(r1.Save(newMockStore()), ShouldBeNil)
		So(fileExists(r1.path), ShouldBeTrue)

		got := r2.Load()
		So(got.Artisans, ShouldResemble, expected.Artisans)
		So(got.Book, ShouldResemble, expected.Book)
		So(got.Catalog, ShouldResemble, expected.Catalog)

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
	artisans := []*Artisan{
		&Artisan{ID: "id1", Label: "artisan1"},
		&Artisan{ID: "id2", Label: "artisan2"},
	}
	s := NewStore()
	s.Artisans = artisans
	s.Catalog.Add(&mockItem{ID: "a", Name: "ab"})
	return s
}

type mockItem struct {
	ID   string
	Name string
}

func (i mockItem) Match(o doable.Item) bool {
	return false
}
func (i mockItem) UID() string {
	return i.ID
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
