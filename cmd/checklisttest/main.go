package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func run(window *app.Window) error {
	theme := material.NewTheme()
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			var ops op.Ops
			// This graphics context is used for managing the rendering state.
			graphicContext := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")
			subtitle := material.H2(theme, "Hello my baby")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle
			title.MaxLines = 1

			// Draw the label to the graphics context.
			dims := title.Layout(graphicContext)
			{
				log.Println("DDims", dims)
				defer op.Offset(image.Point{0, dims.Size.Y}).Push(graphicContext.Ops).Pop()
				subtitle.Layout(graphicContext)
			}

			// Pass the drawing operations to the GPU.
			e.Frame(graphicContext.Ops)
		}
	}
}

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
