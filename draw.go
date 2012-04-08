package main

import (
    "fmt"
    "github.com/banthar/gl"
    "github.com/jteeuwen/glfw"
    "log"
    "os"
    "time"
)

var (
    esc    = true
    now    = time.Now()
    logger = log.New(os.Stdout, "", 0)
)

func escHandler(k, state int) {
    if k == glfw.KeyEsc && state == glfw.KeyPress {
        esc = false
    }
}

func timeSince() {
    time.Sleep(1e+6)
    s := time.Now()
    glfw.SetWindowTitle(fmt.Sprintf("%.4f s", s.Sub(now).Seconds()))
    now = s
}

func loop() {
    e := initTriangleShaders()

    if e != nil {
        log.Fatalln(e.Error())
    }

    for esc {
        gl.ClearColor(0.0, 0.0, 0.0, 0.0)
        gl.Clear(gl.COLOR_BUFFER_BIT)
        drawTriangle()
        glfw.SwapBuffers()
        timeSince()
    }
}

func main() {
    if e := glfw.Init(); e != nil {
        panic(e.Error())
    }

    defer glfw.Terminate()

    if e := glfw.OpenWindow(468, 320, 0, 0, 0, 0, 0, 0, glfw.Windowed); e != nil {
        panic(e.Error())
    }

    if gl.Init() != 0 {
        panic("GL Err")
    }

    glfw.SetKeyCallback(escHandler)

    loop()
}
