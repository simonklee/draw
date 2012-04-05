package main

import (
    "github.com/banthar/gl"
    "errors"
)

var vsSrc = `
#version 120

attribute vec2 coord2d;

void main(void) {
    gl_Position = vec4(coord2d, 0.0, 1.0);
}
`
var fsSrc = `
#version 120

void main(void) {
    gl_FragColor[0] = 0.8;
    gl_FragColor[1] = gl_FragCoord.y/320.0;
    gl_FragColor[2] = gl_FragCoord.x/468.0;
}
`

var program gl.Program
var attrLoc gl.AttribLocation

func loadShader(typ gl.GLenum, src string) (gl.Shader, error) {
    vs := gl.CreateShader(typ)
    vs.Source(src)
    vs.Compile()

    if vs.Get(gl.COMPILE_STATUS) != gl.TRUE {
        return 0, errors.New("shader compile err")
    }

    return vs, nil
}

func initTriangleShaders() error {
    vs, e := loadShader(gl.VERTEX_SHADER, vsSrc)

    if e != nil {
        return e
    }

    fs, e := loadShader(gl.FRAGMENT_SHADER, fsSrc)

    if e != nil {
        return e
    }

    p := gl.CreateProgram()
    p.AttachShader(vs)
    p.AttachShader(fs)
    p.Link()

    if p.Get(gl.LINK_STATUS) != gl.TRUE {
        return errors.New("program link")
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
