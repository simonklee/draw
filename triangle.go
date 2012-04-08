package main

import (
    "github.com/banthar/gl"
    "unsafe"
)

var program gl.Program
var vboTri gl.Buffer
var data []gl.GLfloat

func initTriangleShaders() error {
    var v gl.GLfloat = 0.0
    data = []gl.GLfloat{
        0.0, 0.8,
        -0.8, -0.8,
        0.8, -0.8,
    }

    vboTri = gl.GenBuffer()
    // glGenBuffer(
    // specify the current target buffer
    vboTri.Bind(gl.ARRAY_BUFFER)

    // create and init a buffer data
    gl.BufferData(gl.ARRAY_BUFFER, len(data) * int(unsafe.Sizeof(v)), data, gl.STATIC_DRAW)

    p, e := createProgram("triangle1.v", "triangle1.f")

    if e != nil {
        return e
    }

    program = p
    return nil
}

func drawTriangle() {
    program.Use()
    coord2d := program.GetAttribLocation("coord2d")
    coord2d.EnableArray()

    //coord2d.AttribPointer(2, gl.FLOAT, false, 0, data)
    vboTri.Bind(gl.ARRAY_BUFFER)
    coord2d.AttribPointer(2, gl.FLOAT, false, 0, uintptr(0))
    gl.DrawArrays(gl.TRIANGLES, 0, 3)
    coord2d.DisableArray()
}
