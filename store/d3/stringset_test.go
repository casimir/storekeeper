package d3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Given several string", t, func() {
		set := &StringSet{}
		l := []string{
			"some string",
			"another string",
			"one last string",
			"zorro string",
		}

		Convey("It should add it to the set if not present", func() {
			So(set.Len(), ShouldBeZeroValue)
			for _, it := range l {
				So(set.Add(it), ShouldBeTrue)
			}
			So(set.Len(), ShouldEqual, 4)
			for _, it := range l {
				So(set.Add(it), ShouldBeFalse)
			}
			So(set.Len(), ShouldEqual, 4)
		})
	})
}
