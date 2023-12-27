package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func main() {

	a := app.New()
	w := a.NewWindow("test windows")

	w.SetContent(layout())

	w.Resize(fyne.NewSize(1200, 650))
	w.ShowAndRun()

}

func layout() fyne.CanvasObject {

	grid := widget.NewTextGridFromString("TextGridABABCCC\n\tContent\nZebra\nWhat are you talking about?")
	grid.SetStyleRange(0, 4, 0, 7,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})
	grid.SetRowStyle(1, &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})

	white := &widget.CustomTextGridStyle{FGColor: color.White, BGColor: color.Black}
	black := &widget.CustomTextGridStyle{FGColor: color.Black, BGColor: color.White}
	grid.Rows[2].Cells[0].Style = white
	grid.Rows[2].Cells[1].Style = black
	grid.Rows[2].Cells[2].Style = white
	grid.Rows[2].Cells[3].Style = black
	grid.Rows[2].Cells[4].Style = white

	grid.SetRowStyle(3, white)

	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	right := widget.NewTextGrid()
	right.SetText("what are you talking about?")

	return container.NewGridWithColumns(2, grid, right)
}
