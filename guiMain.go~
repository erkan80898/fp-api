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
			{Text: "Product Address EX: A[1+2+21], B[1+22]", Widget: skuToChange},
			{Text: "", Widget: outputText},
		},
		OnSubmit: func() {
			qty, err := strconv.Atoi(quantityEntry.Text)
			if err != nil {
				log.Println("Non-number Qty")
				return
			}

			resp := Run(qty, skuToChange.Text)
			outputText.SetText(resp)
		},
	}
	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}

func Run(qty int, skusAsText string) string {

	parts := strings.Split(skusAsText, ",")
	res := []string{}
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		if strings.Index(parts[i], "[") != -1 {
			letter := string(parts[i][0])
			rest := string(parts[i][2 : len(parts[i])-1])
			nums := strings.Split(rest, "+")
			for _, v := range nums {
				res = append(res, "_"+letter+v+"_")
			}
		} else if len(parts[i]) <= 4 {
			res = append(res, "_"+parts[i]+"_")
		} else {
			res = append(res, parts[i])
		}
	}

	return Action.UpdateListingQty("fruitListingVariant.csv", res, qty)
}
