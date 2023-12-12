package main

import (
	"log"
	"strconv"

	Action "flx/action"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Inventory Updater")

	quantityEntry := widget.NewEntry()
	skuToChange := widget.NewMultiLineEntry()
	outputText := widget.NewTextGridFromString("")

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
			outputText.SetText(resp)
		},
	}

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
