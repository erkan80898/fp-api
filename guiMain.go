package main

import (
	"image/color"
	"log"
	"strconv"

	Action "flx/action"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Form Widget")

	quantityEntry := widget.NewEntry()
	skuToChange := widget.NewMultiLineEntry()
	outputText := canvas.NewText("Output", color.White)
	outputText.Alignment = fyne.TextAlignTrailing
	outputText.TextStyle = fyne.TextStyle{Italic: true}

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "New Quantity", Widget: quantityEntry},
			{Text: "Product Address - seperate by ,", Widget: skuToChange},
			{Text: "", Widget: outputText},
		},
		OnSubmit: func() {
			qty, err := strconv.Atoi(quantityEntry.Text)
			if err != nil {
				log.Println("Non-number Qty")
				return
			}

			resp := Action.Run(qty, skuToChange.Text)
			outputText.Text = resp
			myWindow.Close()
		},
	}

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
