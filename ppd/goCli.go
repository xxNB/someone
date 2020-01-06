package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

func goCLi() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	//if dsds := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, newmovie); dsds != nil {
	//	log.Panicln(err)
	//}
	//if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, ); err != nil {
	//	log.Panicln(err)
	//}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Welcome Douban Cli", maxX/2-7, maxY/2, maxX/2+15, maxY/2+10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Welcome Douban Cli " + "\n" + "0: newmovie")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

//func esc(g *gocui.Gui, v *gocui.View) error {
//	return gocui.re
//}

func newmovie(g *gocui.Gui, v *gocui.View) []*DouMovie {
	//ss := &DouMovie{}
	//return ss.newHignMovieSelect()
	return nil
}

