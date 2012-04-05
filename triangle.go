package main

import (
    "github.com/banthar/gl"
    "errors"
)

var program gl.Program
var attrLoc gl.AttribLocation

func initTriangleShaders() error {
    p, e := createProgram("triangle1.v", "triangle1.f")

    if e != nil {
        return e
    }

    a := p.GetAttribLocation("coord2d")

    if a == -1 {
        return errors.New("attrib loc")
    }

    program = p
    attrLoc = a
    return nil
}

func drawTriangle() {
    program.Use()
    attrLoc.EnableArray()

    p := []gl.GLfloat{
        0.0,  0.8,
       -0.8, -0.8,
        0.8, -0.8,
    }

    attrLoc.AttribPointer(2, false, 0, p)
    gl.DrawArrays(gl.TRIANGLES, 0, 3)
    attrLoc.DisableArray()
}
