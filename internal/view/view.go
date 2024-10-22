package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type FyneView struct {
	app    fyne.App
	Window fyne.Window

	// Calculator buttons
	Label    *widget.Label
	LabelX   *widget.Button
	EnterX   *widget.Entry
	Btn0     *widget.Button
	Btn00    *widget.Button
	Btn1     *widget.Button
	Btn2     *widget.Button
	Btn3     *widget.Button
	Btn4     *widget.Button
	Btn5     *widget.Button
	Btn6     *widget.Button
	Btn7     *widget.Button
	Btn8     *widget.Button
	Btn9     *widget.Button
	BtnAC    *widget.Button
	BtnBrctL *widget.Button
	BtnBrctR *widget.Button
	BtnC     *widget.Button
	BtnDiv   *widget.Button
	BtnDot   *widget.Button
	BtnMul   *widget.Button
	BtnPerc  *widget.Button
	BtnPow   *widget.Button
	BtnPlus  *widget.Button
	BtnSub   *widget.Button
	Atan     *widget.Button
	Acos     *widget.Button
	Asin     *widget.Button
	BtnSqrt  *widget.Button
	Sin      *widget.Button
	Cos      *widget.Button
	Tan      *widget.Button
	Lgr      *widget.Button
	Ln       *widget.Button
	BtnEqual *widget.Button
	Notation string

	// Credit buttons
	LabelMontPay    *widget.Label
	LabelMontPaySum *widget.Label
	LabelRub        *widget.Label
	LabelPayOver    *widget.Label
	LabelPayOverSum *widget.Label
	LabelTotal      *widget.Label
	LabelTotalSum   *widget.Label
	LabelSumCred    *widget.Label
	SumCred         *widget.Entry
	LabelTimeCred   *widget.Label
	TimeCred        *widget.Entry
	YearCred        *widget.Select
	LabelPercent    *widget.Label
	LabelP          *widget.Label
	Percent         *widget.Entry
	RadioButton     *widget.RadioGroup
	BtnCreditCalc   *widget.Button

	// Graphic btn
	Image    *fyne.Container
	BtnGraph *widget.Button

	// Info btn
}

func InitApp() *FyneView {
	app := app.New()
	w := app.NewWindow("Calculator")
	notation := ""
	label := widget.NewLabel("")
	labelX := widget.NewButton("x", nil)
	entrX := widget.NewEntry()
	btnAC := widget.NewButton("AC", nil)
	btnC := widget.NewButton("C", nil)
	btnBrctL := widget.NewButton("(", nil)
	btnBrctR := widget.NewButton(")", nil)
	btnPerc := widget.NewButton("%", nil)
	btnDiv := widget.NewButton("/", nil)
	btn1 := widget.NewButton("1", nil)
	btn2 := widget.NewButton("2", nil)
	btn3 := widget.NewButton("3", nil)
	btnMul := widget.NewButton("*", nil)
	btn4 := widget.NewButton("4", nil)
	btn5 := widget.NewButton("5", nil)
	btn6 := widget.NewButton("6", nil)
	btnSub := widget.NewButton("-", nil)
	btn7 := widget.NewButton("7", nil)
	btn8 := widget.NewButton("8", nil)
	btn9 := widget.NewButton("9", nil)
	btnPlus := widget.NewButton("+", nil)
	btnDot := widget.NewButton(".", nil)
	btnUnar := widget.NewButton("+", nil)
	btn00 := widget.NewButton("00", nil)
	btn0 := widget.NewButton("0", nil)
	atan := widget.NewButton("atan", nil)
	acos := widget.NewButton("acos", nil)
	asin := widget.NewButton("asin", nil)
	btnPow := widget.NewButton("^", nil)
	btnSqrt := widget.NewButton("Sqrt", nil)
	sin := widget.NewButton("sin", nil)
	cos := widget.NewButton("cos", nil)
	tan := widget.NewButton("tan", nil)
	lgr := widget.NewButton("log", nil)
	ln := widget.NewButton("ln", nil)
	btnEqual := widget.NewButton("=", nil)

	contentL := container.NewGridWithColumns(1, label)
	contentT := container.NewGridWithColumns(5, sin, cos, tan, lgr, ln)
	contentS1 := container.NewGridWithColumns(5, asin, btnAC, btnBrctL, btnBrctR, btnC)
	contentS2 := container.NewGridWithColumns(5, acos, btnPerc, btnPow, btnSqrt, btnDiv)
	content1 := container.NewGridWithColumns(5, atan, btn7, btn8, btn9, btnMul)
	content2 := container.NewGridWithColumns(5, entrX, btn4, btn5, btn6, btnSub)
	content3 := container.NewGridWithColumns(5, labelX, btn1, btn2, btn3, btnPlus)
	content4 := container.NewGridWithColumns(5, btnUnar, btn00, btn0, btnDot, btnEqual)
	content := container.NewGridWithRows(8, contentL, contentT, contentS1, contentS2, content1, content2, content3, content4)

	// Credit
	labelMontPay := widget.NewLabel("Ежемесячный платеж")
	labelMontPaySum := widget.NewLabel("")
	labelRub := widget.NewLabel("рублей")
	answerMonthContainer := container.NewGridWithColumns(3, labelMontPay, labelMontPaySum, labelRub)

	labelPayOver := widget.NewLabel("Переплата по кредиту")
	labelPayOverSum := widget.NewLabel("")
	answerOverContainer := container.NewGridWithColumns(3, labelPayOver, labelPayOverSum, labelRub)

	labelTotal := widget.NewLabel("Общая выплата")
	labelTotalSum := widget.NewLabel("")
	answerTotalContainer := container.NewGridWithColumns(3, labelTotal, labelTotalSum, labelRub)

	labelSumCred := widget.NewLabel("Сумма кредита")
	sumCred := widget.NewEntry()
	contentCredit1 := container.NewGridWithColumns(3, labelSumCred, sumCred, labelRub)

	labelTimeCred := widget.NewLabel("Срок кредита")
	timeCred := widget.NewEntry()
	yearCred := widget.NewSelect([]string{"лет", "месяцев"}, nil) 
	contentCredit2 := container.NewGridWithColumns(3, labelTimeCred, timeCred, yearCred)

	labelPercent := widget.NewLabel("Процентная ставка")
	labelP := widget.NewLabel("%")
	percent := widget.NewEntry()
	contentPercent := container.NewGridWithColumns(3, labelPercent, percent, labelP)
	radioButton := widget.NewRadioGroup([]string{"Annuitet", "Diff"}, nil)
	containerRadio := container.NewGridWithColumns(1, radioButton)
	btnCreditCalc := widget.NewButton("Рассчитать", nil) 

	contentCred := container.NewGridWithRows(8, answerMonthContainer, answerOverContainer, answerTotalContainer, contentCredit1, contentCredit2, contentPercent, containerRadio, btnCreditCalc)

	// Graphic
	contentGraphicImg := container.NewVBox()
	btnGraph := widget.NewButton("Draw graph", nil)
	contentGraphic := container.NewGridWithRows(3, container.NewVBox(), contentGraphicImg, btnGraph)

	// Menu
	menuCalc := fyne.NewMenuItem("Calculator", func() {
		w.SetContent(content)
	})
	menu1 := fyne.NewMenu("Calculator", menuCalc)

	menuCredit := fyne.NewMenuItem("Credit", func() {
		w.SetContent(contentCred)
	})
	menu2 := fyne.NewMenu("Credit", menuCredit)

	menuGraphic := fyne.NewMenuItem("Graphic", func() {
		w.SetContent(contentGraphic)
	})
	menu3 := fyne.NewMenu("Graphic", menuGraphic)

	menuHelp := fyne.NewMenuItem("About", func() {
		showHelp(w)
	})
	menu4 := fyne.NewMenu("About", menuHelp)

	mainMenu := fyne.NewMainMenu(menu1, menu2, menu3, menu4)

	w.SetContent(content)
	w.SetMainMenu(mainMenu)
	w.Resize(fyne.NewSize(500, 500))

	return &FyneView{
		app:    app,
		Window: w,

		// Calculator
		Label:    label,
		LabelX:   labelX,
		EnterX:   entrX,
		Notation: notation,
		Btn0:     btn0,
		Btn00:    btn00,
		Btn1:     btn1,
		Btn2:     btn2,
		Btn3:     btn3,
		Btn4:     btn4,
		Btn5:     btn5,
		Btn6:     btn6,
		Btn7:     btn7,
		Btn8:     btn8,
		Btn9:     btn9,
		BtnAC:    btnAC,
		BtnBrctL: btnBrctL,
		BtnBrctR: btnBrctR,
		BtnC:     btnC,
		BtnDiv:   btnDiv,
		BtnDot:   btnDot,
		BtnMul:   btnMul,
		BtnPerc:  btnPerc,
		BtnPow:   btnPow,
		BtnPlus:  btnPlus,
		BtnSub:   btnSub,
		BtnSqrt:  btnSqrt,
		Asin:     asin,
		Acos:     acos,
		Atan:     atan,
		Sin:      sin,
		Cos:      cos,
		Tan:      tan,
		Lgr:      lgr,
		Ln:       ln,
		BtnEqual: btnEqual,

		// Credit
		LabelMontPay:    labelMontPay,
		LabelMontPaySum: labelMontPaySum,
		LabelRub:        labelRub,
		LabelPayOver:    labelPayOver,
		LabelPayOverSum: labelPayOverSum,
		LabelTotal:      labelTotal,
		LabelTotalSum:   labelTotalSum,
		LabelSumCred:    labelSumCred,
		SumCred:         sumCred,
		LabelTimeCred:   labelTimeCred,
		TimeCred:        timeCred,
		YearCred:        yearCred,
		LabelPercent:    labelPercent,
		LabelP:          labelP,
		Percent:         percent,
		RadioButton:     radioButton,
		BtnCreditCalc:   btnCreditCalc,

		// Graphic
		Image:    contentGraphicImg,
		BtnGraph: btnGraph,
	}
}

func (v *FyneView) Run() {
	v.Window.ShowAndRun()
}

func (fv *FyneView) EqualPress(data string) {
	fv.Label.SetText(data)
}

func (fv *FyneView) BtnPress(data string) {
	fv.Notation += data
	fv.Label.SetText(fv.Notation)
}

func (fv *FyneView) PressAC() {
	fv.Notation = ""
	fv.Label.SetText(fv.Notation)
}

func (fv *FyneView) PressC() {
	if fv.Notation != "" {
		fv.Notation = fv.Notation[:len(fv.Notation)-1]
		fv.Label.SetText(fv.Notation)
	}
}

func showHelp(w fyne.Window) {
	helpDialog := dialog.NewInformation("Help", "This is my Calculator on golang\nCalculator - simple calculator. Backend on C and front on Golang\nCredit - calculation credit, annuitet and diff\nGraph - simple graphic on golang", w)
	helpDialog.Show()
}
