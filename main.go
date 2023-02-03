package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
	user_funcs "goproject.com/srcFiles"
)

func main() {

	rdb := user_funcs.ConnectDb()
	user_funcs.CreateTable(rdb)
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create a new window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Quizlet")
	window.SetMinimumSize2(800, 500)

	//function to handle the UI.
	user_funcs.HandleFrontend(window, rdb)

	window.Show()
	app.Exec()
}
