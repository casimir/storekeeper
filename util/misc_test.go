package util

import (
	"os"
	"runtime"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestApplicationPath(t *testing.T) {
	// This test is designed to run only on travis
	if runtime.GOOS != "linux" {
		return
	}
	Convey("It should return the path", t, func() {
		So(ApplicationPath(), ShouldEqual, os.Getenv("HOME")+"/.storekeeper")
	})
}
