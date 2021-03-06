// Scene.go : contain game components within nodes, and stucture them inside of scenes for rendering etc

package Graphics

import (
	"GT/Graphics/Components"
	"GT/Graphics/Opengl"
	"GT/Window"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// BaseScene : object to contain nodes for rendering. TODO: implement more scene types from this base type
type BaseScene struct {
	RootNode *Components.Node

	window *Window.Window
	// TODO: can put fps into struct fps counter
	fps       int
	timestart int32
	update    bool
}

// NewBasicScene : start a new simple scene of orthographic orientation
func NewBasicScene(window *Window.Window) (BaseScene, error) {

	s := BaseScene{RootNode: Components.NewNode("RootNode"), window: window}
	s.fps = 0
	s.timestart = int32(time.Now().Unix())
	s.update = true

	Opengl.SetOrthographic(window.Width, window.Height)

	return s, nil
}

// Start : initiate a scene's continous animation. While the scene is running, draw it and update our nodes
func (s *BaseScene) Start() {

	running := true
	s.update = true
	drawStart := int(time.Now().Unix())
	drawEnd := int(time.Now().Unix())

	// continue to render the scene unless signalled for termination
	for running {

		// poll any current events and handle accordingly
		// TODO: while implementing Issue #9 abstract this to an input handler
		var event sdl.Event
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				} else {
					// fmt.Printf("[%d ms] Keyboard\ttype:%d\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n",
					// t.Timestamp, t.Type, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
				}
			default:
				// fmt.Printf("Some event\n")
			}
		}

		s.window.Clear()
		drawEnd = int(time.Now().Unix())

		// TODO: implement or store this data. Implement this with the fps limiting mentioned below
		delta := float32(drawEnd - drawStart)
		s.RootNode.Update(delta)
		drawStart = int(time.Now().Unix())
		s.Draw()

		// sleep in between renders to smooth out the display
		// TODO: consider making this user defined (fps limit)
		time.Sleep(time.Nanosecond)
		s.window.Refresh()
	}
}

// Draw : wrap the openGl draw method and report fps info
func (s *BaseScene) Draw() {

	if s.update {
		fmt.Println("In update")

		// TODO: remove hard coded resolution
		// gl.BindTexture(gl.TEXTURE_2D, s.spriteSheet.textureId)
		Opengl.BindBuffers()
		s.update = false
	}

	Opengl.Draw()

	// calc fps
	if (int32(time.Now().Unix()) - s.timestart) >= 1 {
		fmt.Printf("FPS: %d \n", s.fps)
		s.fps = 0
		s.timestart = int32(time.Now().Unix())
	}

	s.fps = s.fps + 1
}
