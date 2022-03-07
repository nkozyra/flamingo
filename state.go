package flamingo

import "github.com/veandco/go-sdl2/sdl"

type gameState struct {
	window      *sdl.Window
	surface     *sdl.Surface
	name        string
	screenState screenState
	frames      int
	ping        func()
	click       func(t *sdl.MouseButtonEvent) gameState
}

type screenState struct {
	finished    bool
	initialized bool
	attributes  []interface{}
}
