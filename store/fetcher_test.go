package store

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	badUrl  = "http://nothere.nowhere.com"
	goodUrl = "http://ip.jsontest.com"
)

func TestFetch(t *testing.T) {
	Convey("Given several URLs", t, func() {
		f := Fetcher{}
		Convey("It should fetch every url", func() {
			in := []string{goodUrl, badUrl, goodUrl, badUrl}
			out := f.Fetch(in)
			So(len(out), ShouldEqual, len(in))
		})
	})
}

func TestRequest(t *testing.T) {
	Convey("Given an URL", t, func() {
		f := Fetcher{}
		Convey("It should fetch it", func() {
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
