package main

import (
    "github.com/banthar/gl"
    "errors"
    "io/ioutil"
    "strings"
)

func loadShader(typ gl.GLenum, filename string) (gl.Shader, error) {
    if !strings.HasPrefix(filename, ".glsl") {
        filename = filename + ".glsl"
    }

    src, e := ioutil.ReadFile(filename)

    if e != nil {
        logger.Println("read file:", e.Error())
        return 0, e
    }

    //logger.Print(filename,"\n", string(src))
    s := gl.CreateShader(typ)
    s.Source(string(src))
    s.Compile()

    if s.Get(gl.COMPILE_STATUS) != gl.TRUE {
        e := errors.New(s.GetInfoLog())
        s.Delete()
        return 0, e
    }

    return s, nil
}

func createProgram(vname, fname string) (gl.Program, error) {
    vs, e := loadShader(gl.VERTEX_SHADER, vname)

    if e != nil {
        return 0, e
    }

    fs, e := loadShader(gl.FRAGMENT_SHADER, fname)

    if e != nil {
        return 0, e
    }

    p := gl.CreateProgram()
    p.AttachShader(vs)
    p.AttachShader(fs)
    p.Link()

    if p.Get(gl.LINK_STATUS) != gl.TRUE {
        return 0, errors.New(p.GetInfoLog())
    }

    return p, nil
}

