package main

import (
	"math"

	"github.com/tfriedel6/canvas/sdlcanvas"
)

func main() {
	wnd, cv, err := sdlcanvas.CreateWindow(1280, 720, "Hello")
	if err != nil {
		panic(err)
	}
	defer wnd.Destroy()

	wnd.MainLoop(func() {
		w, h := float64(cv.Width()), float64(cv.Height())
		cv.SetFillStyle("#000")
		cv.FillRect(0, 0, w, h)

		for r := 0.0; r < math.Pi*2; r += math.Pi * 0.1 {
			cv.SetFillStyle(int(r*10), int(r*20), int(r*40))
			cv.BeginPath()
			cv.MoveTo(w*0.5, h*0.5)
			cv.Arc(w*0.5, h*0.5, math.Min(w, h)*0.4, r, r+0.1*math.Pi, false)
			cv.ClosePath()
			cv.Fill()
		}

		cv.SetStrokeStyle("#FFF")
		cv.SetLineWidth(10)
		cv.BeginPath()
		cv.Arc(w*0.5, h*0.5, math.Min(w, h)*0.4, 0, math.Pi*2, false)
		cv.Stroke()
	})
}