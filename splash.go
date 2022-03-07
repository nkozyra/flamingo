package flamingo

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func drawSplash() {

}

func splashHandler(w *sdl.Window, s *sdl.Surface) gameState {
	var g gameState
	// initLow := 0xff000011
	// initHigh := 0xff000033
	// finishLow := 0xff666699
	// finishHigh := 0xff99eeff

	// gradient := struct {
	// 	startBottom color
	// 	startTop    color
	// 	endBottom   color
	// 	endTop      color
	// }{
	// 	startBottom: color{r: 0, g: 20, b: 255},
	// 	startTop:    color{r: 10, g: 200, b: 255},
	// 	endBottom:   color{r: 15, g: 25, b: 255},
	// 	endTop:      color{r: 25, g: 210, b: 255},
	// }

	g.window = w
	g.surface = s
	g.name = "splash"
	g.screenState = struct {
		finished    bool
		initialized bool
		attributes  []interface{}
	}{
		finished:    false,
		initialized: false,
	}
	g.ping = func() {
		if !g.screenState.initialized {
			g.screenState.initialized = true
		}
		ww, wh := g.window.GetSize()
		rect := sdl.Rect{0, 0, ww, wh}
		g.surface.FillRect(&rect, 0xff6f97ff)
		g.window.UpdateSurface()
		if !g.screenState.finished {
			g.frames++
		}
		if g.frames > 599 {
			g.screenState.finished = true
		}
	}
	g.click = func(t *sdl.MouseButtonEvent) gameState {
		if g.screenState.finished {
			log.Println("changing handler")
			return gameSelectHandler(w, s)
		}
		return g
	}

	return g
}

func gameSelectHandler(w *sdl.Window, s *sdl.Surface) gameState {
	var g gameState
	g.window = w
	g.surface = s
	g.frames = 0
	g.name = "gameSelect"
	g.screenState = screenState{
		finished: false,
		attributes: []interface {
		}{},
	}
	g.ping = func() {
		ww, wh := g.window.GetSize()
		rect := sdl.Rect{0, 0, ww, wh}
		g.surface.FillRect(&rect, 0xffffff00)
		g.window.UpdateSurface()
	}
	g.click = func(t *sdl.MouseButtonEvent) gameState {
		return g
	}
	return g
}
