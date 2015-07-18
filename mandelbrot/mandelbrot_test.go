package mandelbrot

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMin(t *testing.T) {
	Convey("It should return lower values.", t, func() {
		So(min(200, 1), ShouldEqual, 1)
		So(min(20, 300), ShouldEqual, 20)
	})
}
