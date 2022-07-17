package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	f, _ := os.Create("error.log")
	log.SetOutput(f)
	err := glfw.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "Testing"+time.Now().String(), nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}

	self, err := os.Executable()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(self)
	if err = cmd.Start(); err != nil {
		log.Fatalln(err)
	}
	if err = cmd.Process.Release(); err != nil {
		log.Fatalln(err)
	}
}
