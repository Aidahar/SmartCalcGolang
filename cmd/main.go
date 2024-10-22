package main

import (
	"calculator/internal/model"
	"calculator/internal/presenter"
	"calculator/internal/view"
	"fmt"
	"time"
)

func main() {
	data := time.Now()
	filePath := fmt.Sprintf("./log/log_%s.log", data.Format("02-01-06-15-04-05"))
	model := model.NewModel(filePath)
	view := view.InitApp()
	presenter := presenter.NewPresenter(model, view)

	presenter.View.Run()
}
