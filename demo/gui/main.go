package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
)

func main() {
	t := &myTheme{}
	t.SetFonts("/Users/wanghongjie/GolandProjects/jim/demo/gui/hwkt.ttf", "")
	app := app.New()
	app.Settings().SetTheme(t)
	w := app.NewWindow("杰子学编程")
	parse, _ := url.Parse("https://julywhj.cn")
	content := widget.NewButton("click me", func() {
		log.Println("tapped")
	})

	img2 := canvas.NewImageFromFile("/Users/wanghongjie/GolandProjects/jim/demo/gui/clogo.png")

	container := fyne.NewContainerWithLayout(
		layout.NewGridWrapLayout(fyne.NewSize(150, 150)),
		img2, content, widget.NewHyperlink("杰子学编程", parse))
	w.SetContent(container)
	w.ShowAndRun()
}
