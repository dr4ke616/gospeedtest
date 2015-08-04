// An app that displays the current interent speed for the device.
//
// Note: This demo is an early preview of Go 1.5. In order to build this
// program as an Android APK using the gomobile tool.
//
// See http://godoc.org/golang.org/x/mobile/cmd/gomobile to install gomobile.
//
// Get the basic example and use gomobile to build or install it on your device.
//
//   $ go get -d github.com/dr4ke616/gospeedtest
//   $ cd github.com/dr4ke616/gospeedtest
//   $ gomobile build . # will build an APK
//
//   # plug your Android device to your computer or start an Android emulator.
//   # if you have adb installed on your machine, use gomobile install to
//   # build and deploy the APK to an Android target.
//   $ gomobile install golang.org/x/mobile/example/basic
//
// Switch to your device or emulator to start the Basic application from
// the launcher.
// You can also run the application on your desktop by running the command
// below. (Note: It currently doesn't work on Windows.)
//   $ go install golang.org/x/mobile/example/basic && basic
package main

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/config"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/geom"
	"golang.org/x/mobile/gl"

	"github.com/dr4ke616/gospeedtest/nw_speedtest"
)

var (
	program      gl.Program
	position     gl.Attrib
	offset       gl.Uniform
	screen_color gl.Uniform

	speed_rate = make(chan int)
)

func main() {
	// check network speed runs only once when the app first loads.
	go func() {
		st := nw_speedtest.Speedtest{
			FileLocation: "http://download.thinkbroadband.com/10MB.zip",
			Verbos:       true,
		}
		result, _ := st.Start()
		speed_rate <- result
	}()

	app.Main(func(a app.App) {
		var c config.Event
		for e := range a.Events() {
			switch e := app.Filter(e).(type) {
			case config.Event:
				c = e
			case paint.Event:
				onDraw(c)
				a.EndPaint()
			}
		}
	})
}

func onDraw(c config.Event) {
	gl.ClearColor(1, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	debug.DrawFPS(c)
	DrawResult(c, <-speed_rate)
}

func DrawResult(c config.Event, rate int) {
	const imgW, imgH = 8*(fontWidth+1) + 1, fontHeight + 2

	image_result := glutil.NewImage(imgW, imgH)

	display := [8]byte{
		4: 'M',
		5: 'B',
		6: 'P',
		7: 'S',
	}

	display[2] = '0' + byte((rate/1e0)%10)
	display[1] = '0' + byte((rate/1e1)%10)
	display[0] = '0' + byte((rate/1e2)%10)

	draw.Draw(image_result.RGBA, image_result.RGBA.Bounds(), image.White, image.Point{}, draw.Src)
	for i, c := range display {
		glyph := glyphs[c]
		if len(glyph) != fontWidth*fontHeight {
			continue
		}
		for y := 0; y < fontHeight; y++ {
			for x := 0; x < fontWidth; x++ {
				if glyph[fontWidth*y+x] == ' ' {
					continue
				}
				image_result.RGBA.SetRGBA((fontWidth+1)*i+x+1, y+1, color.RGBA{A: 0xff})
			}
		}
	}

	image_result.Upload()
	image_result.Draw(
		c,
		geom.Point{0, c.HeightPt - imgH},
		geom.Point{imgW, c.HeightPt - imgH},
		geom.Point{0, c.HeightPt},
		image_result.RGBA.Bounds(),
	)
}

const (
	fontWidth  = 5
	fontHeight = 7
)

// glyphs comes from the 6x10 fixed font from the plan9port:
// https://github.com/9fans/plan9port/tree/master/font/fixed
//
// 6x10 becomes 5x7 because each glyph has a 1-pixel margin plus space for
// descenders.
//
// Its README file says that those fonts were converted from XFree86, and are
// in the public domain.
var glyphs = [256]string{
	'0': "" +
		"  X  " +
		" X X " +
		"X   X" +
		"X   X" +
		"X   X" +
		" X X " +
		"  X  ",
	'1': "" +
		"  X  " +
		" XX  " +
		"X X  " +
		"  X  " +
		"  X  " +
		"  X  " +
		"XXXXX",
	'2': "" +
		" XXX " +
		"X   X" +
		"    X" +
		"  XX " +
		" X   " +
		"X    " +
		"XXXXX",
	'3': "" +
		"XXXXX" +
		"    X" +
		"   X " +
		"  XX " +
		"    X" +
		"X   X" +
		" XXX ",
	'4': "" +
		"   X " +
		"  XX " +
		" X X " +
		"X  X " +
		"XXXXX" +
		"   X " +
		"   X ",
	'5': "" +
		"XXXXX" +
		"X    " +
		"X XX " +
		"XX  X" +
		"    X" +
		"X   X" +
		" XXX ",
	'6': "" +
		"  XX " +
		" X   " +
		"X    " +
		"X XX " +
		"XX  X" +
		"X   X" +
		" XXX ",
	'7': "" +
		"XXXXX" +
		"    X" +
		"   X " +
		"   X " +
		"  X  " +
		" X   " +
		" X   ",
	'8': "" +
		" XXX " +
		"X   X" +
		"X   X" +
		" XXX " +
		"X   X" +
		"X   X" +
		" XXX ",
	'9': "" +
		" XXX " +
		"X   X" +
		"X  XX" +
		" XX X" +
		"    X" +
		"   X " +
		" XX  ",
	'M': "" +
		"XXXXX" +
		"X X X" +
		"X X X" +
		"X X X" +
		"X X X" +
		"X X X" +
		"X   X",
	'B': "" +
		"XXXX " +
		"X   X" +
		"X   X" +
		"XXXX " +
		"X   X" +
		"X   X" +
		"XXXX ",
	'P': "" +
		"XXXX " +
		"X   X" +
		"X   X" +
		"XXXX " +
		"X    " +
		"X    " +
		"X    ",
	'S': "" +
		" XXX " +
		"X   X" +
		"X    " +
		" XXX " +
		"    X" +
		"X   X" +
		" XXX ",
}
