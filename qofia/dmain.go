package main

import (
	"log"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	// Create application
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Create main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	uiw := Ui_MainWindow_new()
	log.Println(uiw, window.Pointer())
	Ui_MainWindow_setupUi(uiw, window.Pointer())

	// Show the window
	window.Show()

	// Execute app
	app.Exec()
}
