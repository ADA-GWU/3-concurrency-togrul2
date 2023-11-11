package gui

import (
	"fmt"
	"image"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
)

func RunGUIEventLoop(w *app.Window, img *image.Image, job func(chan<- image.Image, chan<- error)) error {
	var ops op.Ops
	imageChannel := make(chan image.Image, 10000)
	errorsChannel := make(chan error)
	go job(imageChannel, errorsChannel)

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				imgOp := paint.NewImageOp(*img)
				imgOp.Add(&ops)
				paint.PaintOp{}.Add(gtx.Ops)
				e.Frame(gtx.Ops)
			}
		case newImg := <-imageChannel:
			*img = newImg
			w.Invalidate()
			fmt.Println("Update received")
		case err := <-errorsChannel:
			return err
		}
	}
}
