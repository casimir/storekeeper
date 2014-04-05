package storage

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestItemAddAndRemove(t *testing.T) {
	Convey("When adding and removing several items", t, func() {
		bag := Bag{}

		Convey("The bag should contain the correct number of each item", func() {
			So(bag.Count("1"), ShouldBeZeroValue)
			bag.AddItem(14, "1")
			So(bag.Count("1"), ShouldEqual, 14)
			bag.AddItem(8, "1")
			So(bag.Count("1"), ShouldEqual, 22)
			bag.RemoveItem(11, "1")
			So(bag.Count("1"), ShouldEqual, 11)
			bag.RemoveItem(11, "1")
			So(len(bag.Items), ShouldBeZeroValue)
		})
		Convey("Errors should be returned on illegal actions", func() {
			bag.AddItem(1, "1")
			So(bag.Count("1"), ShouldEqual, 1)
			So(bag.RemoveItem(12, "1"), ShouldNotBeNil)
			So(bag.RemoveItem(1, "2"), ShouldNotBeNil)
		})
	})
}

func TestItemCount(t *testing.T) {
	Convey("Given several items IDs", t, func() {
		bag := makeTestingBag()

		Convey("It should return the number of items", func() {
			So(bag.Count("1"), ShouldEqual, 3)
		})
		Convey("When the item is not contained it should return 0", func() {
			So(bag.Count("2"), ShouldBeZeroValue)
		})
	})
}

func makeTestingBag() (b Bag) {
	b.Items = map[string]*Stack{
		"1": {Count: 3, Item: Item{"1", "Mocking bird"}},
	}
	return
}
