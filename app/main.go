package main

import (
	"os"

	"github.com/therecipe/qt/widgets"

)

const (
	APP_NAME = "NBS IPFS NODE"
	APP_WIN_WIDTH = 900
	APP_WIN_HEIGHT = 620
)

func main() {
	/**
	 *
	 */
	app :=  widgets.NewQApplication(len(os.Args),os.Args)

	window := widgets.NewQMainWindow(nil,0)

	window.SetMinimumSize2(APP_WIN_WIDTH,APP_WIN_HEIGHT)

	window.SetWindowTitle(APP_NAME)


	window.Show()

	app.Exec()
}
