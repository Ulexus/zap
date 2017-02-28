// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
Package coloring adds coloring functionality.
*/
package coloring

import "fmt"

const (
	fgBlackCode colorCode = iota + 30
	fgRedCode
	fgGreenCode
	fgYellowCode
	fgBlueCode
	fgMagentaCode
	fgCyanCode
	fgWhiteCode
)

var (
	// FGBlack is a Color with a black foreground.
	FGBlack = &Color{fgBlackCode}
	// FGRed is a Color with a red foreground.
	FGRed = &Color{fgRedCode}
	// FGGreen is a Color with a green foreground.
	FGGreen = &Color{fgGreenCode}
	// FGYellow is a Color with a yellow foreground.
	FGYellow = &Color{fgYellowCode}
	// FGBlue is a Color with a blue foreground.
	FGBlue = &Color{fgBlueCode}
	// FGMagenta is a Color with a magenta foreground.
	FGMagenta = &Color{fgMagentaCode}
	// FGCyan is a Color with a cyan foreground.
	FGCyan = &Color{fgCyanCode}
	// FGWhite is a Color with a white foreground.
	FGWhite = &Color{fgWhiteCode}
)

type colorCode int

// Color represents a color.
type Color struct {
	colorCode colorCode
}

// Add adds the coloring to the given string.
func (c *Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c.colorCode, s)
}
