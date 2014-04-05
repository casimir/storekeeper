package store

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	badUrl  = "http://nothere.nowhere.com"
	goodUrl = "http://ip.jsontest.com"
)

func TestFetch(t *testing.T) {
	Convey("Given a fetcher", t, func() {
		f := Fetcher{}
		Convey("It should fetch every url", func() {
			in := []string{goodUrl, badUrl}

			out := f.Fetch(in)
			So(len(out), ShouldEqual, 2)
		})
	})
}
func TestRequest(t *testing.T) {
	Convey("Given a fetcher", t, func() {
		f := Fetcher{}
		Convey("It should fetch an url", func() {
			out := f.Request(goodUrl)
			So(out.Body, ShouldNotBeEmpty)
			So(out.Err, ShouldBeNil)
		})
		Convey("It should provide an error for a bad url", func() {
			out := f.Request(badUrl)
			So(out.Body, ShouldBeEmpty)
			So(out.Err, ShouldNotBeNil)
		})
	})
}
