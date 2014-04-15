package util

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdd(t *testing.T) {
	Convey("Given several string", t, func() {
		set := StringSet{}
		So(set.Len(), ShouldBeZeroValue)
		Convey("It should add it to the set", func() {
			So(set.Add("some string"), ShouldBeTrue)
			So(set.Add("another string"), ShouldBeTrue)
			So(set.Add("one last string"), ShouldBeTrue)
			So(set.Add("zorro string"), ShouldBeTrue)
			So(set.Len(), ShouldEqual, 4)
		})
		Convey("It should not add it to the set if already in", func() {
			So(set.Add("some string"), ShouldBeFalse)
			So(set.Add("another string"), ShouldBeFalse)
			So(set.Add("one last string"), ShouldBeFalse)
			So(set.Add("zorro string"), ShouldBeFalse)
			So(set.Len(), ShouldEqual, 4)
		})
	})
}
