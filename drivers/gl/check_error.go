// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"fmt"

	"github.com/go-gl/gl/v3.2-core/gl"
)

func CheckError() {
	if v := gl.GetError(); v != 0 {
		switch v {
		case gl.INVALID_ENUM:
			panic("GL returned error GL_INVALID_ENUM")
		case gl.INVALID_FRAMEBUFFER_OPERATION:
			panic("GL returned error GL_INVALID_FRAMEBUFFER_OPERATION")
		case gl.INVALID_INDEX:
			panic("GL returned error GL_INVALID_INDEX")
		case gl.INVALID_OPERATION:
			panic("GL returned error GL_INVALID_OPERATION")
		case gl.INVALID_VALUE:
			panic("GL returned error GL_INVALID_VALUE")
		default:
			panic(fmt.Errorf("GL returned error 0x%.4x", v))
		}
	}
}
