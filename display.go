package main

import (
	"log"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"strings"
	"fmt"
	"strconv"
)

type display struct {
	width int
	height int
	rows int
	columns int
	vertexShaderSource string
	fragmentShaderSource string
	cells [][]*cell
        window *glfw.Window 
	program uint32
}

type cell struct {
	drawable uint32
	x int
	y int
	alive bool
}

func (c *cell) draw() {
	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square) / 3))
}


func (d *display) initialize() {
	d.width = 640
	d.height = 320
	d.rows = 32
	d.columns = 64
	d.vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"
	d.fragmentShaderSource = `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1);
		}
	` + "\x00"
}

func (d *display) start() {
	d.window = d.initGlfw()
	d.program = d.initOpenGL()
	d.cells = d.makeCells()
}

func (d *display) end() {
	glfw.Terminate()
}

func (d *display) makeCells() [][]*cell {
	cells := make([][]*cell, d.rows, d.columns)
	for x := 0; x < d.rows; x++ {
		for y := 0; y < d.columns; y++ {
			c := d.newCell(x, y)
			cells[x] = append(cells[x], c)
		}
	}
	return cells
}

var (
	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func (d *display) checkAnyKeyDown() uint8 {
	keys := []glfw.Key{glfw.Key0, glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4, glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9, glfw.KeyA, glfw.KeyB, glfw.KeyC, glfw.KeyD, glfw.KeyE, glfw.KeyF}
	selectedKey := uint8(0)
	f := false
	for !f  {
		if d.window.ShouldClose() {
			d.end()
		}
		for i := 0; i < len(keys); i++ {
			if d.window.GetKey(keys[i]) == glfw.Press || d.window.GetKey(keys[i])  == glfw.Repeat {
				selectedKey = uint8(i)
				f = true
					break
			}
		}
		glfw.PollEvents()
	}
	return selectedKey
}

func (d *display) checkKeyStatus(key uint8) bool {
	keys := []glfw.Key{glfw.Key0, glfw.Key1, glfw.Key2, glfw.Key3, glfw.Key4, glfw.Key5, glfw.Key6, glfw.Key7, glfw.Key8, glfw.Key9, glfw.KeyA, glfw.KeyB, glfw.KeyC, glfw.KeyD, glfw.KeyE, glfw.KeyF}
	s := false
	if d.window.GetKey(keys[key]) == glfw.Press || d.window.GetKey(keys[key]) == glfw.Repeat {
		s = true
	}
	return s
}

func (d *display) newCell(y, x int) *cell {
	points := make([]float32, len(square), len(square))
	copy(points, square)
	for i := 0; i < len(points); i++ {
		var position float32
		var size float32
		switch i % 3 {
		case 0:
			size = 1.0 / float32(d.columns)
			position = float32(x) * size
		case 1:
			size = 1.0 / float32(d.rows)
			position = float32(y) * size
		default: 
			continue
		}
		if points[i] < 0 {
			points[i] = (position * 2) - 1
		} else {
			points[i] = ((position + size) * 2) - 1
		}
	}
	return &cell{
		drawable: d.makeVao(points),
		x: x,
		y: y,
		alive: false,
	}
}

func (d *display) initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(d.width, d.height, "Chip8 Emulator", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}

func (d *display) initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))

	vertexShader, err := d.compileShader(d.vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := d.compileShader(d.fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	log.Println("OpenGL version", version)
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func (d *display) clear() {
	for i := 0; i < len(d.cells); i++ {
		for j := 0; j < len(d.cells[i]); j++ {
			d.cells[i][j].alive = false
		}
	}
}


func (d *display) updateScreenState(sprite uint8, x uint8, y uint8) bool {
	fmt.Println("UPDATE SS CALLED")
	mask := uint8(128)
	bitShift := 7
	collision := false
	fmt.Println(x)
	fmt.Println("Display byte", strconv.FormatInt(int64(sprite), 2), "starting at", x, 31-y)
	for bitShift != -1 {
		if x >= 64 {
			x = 0
		}
		pixel := (sprite & mask) >> bitShift
		if pixel == 1 && !d.cells[31-y][x].alive {
			d.cells[31-y][x].alive = true
			fmt.Println("Drawing pixel at:", x, 31-y)
		} else if pixel == 1 && d.cells[31-y][x].alive {
			d.cells[31-y][x].alive = false
			collision = true
			fmt.Println("Erasing pixel at:", x, 31-y)
		}
		x++
		mask /= 2
		bitShift--
	}
	return collision
}
func (d *display) draw(screenState [32][64]bool, cells [][]*cell, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if cells[31-i][j].alive {
				cells[31-i][j].draw()
			}
		}
	}
	
	glfw.PollEvents()
	window.SwapBuffers()
}

func (d *display) makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vao)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func (d *display) compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

