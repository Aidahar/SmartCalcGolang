package presenter

import (
	"calculator/internal/model"
	"calculator/internal/view"
	"image/color"
	"log"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Presenter struct {
	model *model.Model
	View  *view.FyneView
}

func NewPresenter(model *model.Model, view *view.FyneView) *Presenter {
	p := &Presenter{
		model: model,
		View:  view,
	}
	//Credit buttons event
	p.View.Btn0.OnTapped = func() {
		p.press("0")
	}
	p.View.Btn00.OnTapped = func() {
		p.press("00")
	}
	p.View.Btn1.OnTapped = func() {
		p.press("1")
	}
	p.View.Btn2.OnTapped = func() {
		p.press("2")
	}
	p.View.Btn3.OnTapped = func() {
		p.press("3")
	}
	p.View.Btn4.OnTapped = func() {
		p.press("4")
	}
	p.View.Btn5.OnTapped = func() {
		p.press("5")
	}
	p.View.Btn6.OnTapped = func() {
		p.press("6")
	}
	p.View.Btn7.OnTapped = func() {
		p.press("7")
	}
	p.View.Btn8.OnTapped = func() {
		p.press("8")
	}
	p.View.Btn9.OnTapped = func() {
		p.press("9")
	}
	p.View.BtnAC.OnTapped = func() {
		p.pressAC()
	}
	p.View.BtnC.OnTapped = func() {
		p.pressC()
	}
	p.View.LabelX.OnTapped = func() {
		p.press("x")
	}
	p.View.BtnBrctL.OnTapped = func() {
		p.press("(")
	}
	p.View.BtnBrctR.OnTapped = func() {
		p.press(")")
	}
	p.View.BtnDiv.OnTapped = func() {
		p.press("/")
	}
	p.View.BtnDot.OnTapped = func() {
		p.press(".")
	}
	p.View.BtnMul.OnTapped = func() {
		p.press("*")
	}
	p.View.BtnPerc.OnTapped = func() {
		p.press("%")
	}
	p.View.BtnPlus.OnTapped = func() {
		p.press("+")
	}
	p.View.BtnPow.OnTapped = func() {
		p.press("^")
	}
	p.View.BtnSub.OnTapped = func() {
		p.press("-")
	}
	p.View.Atan.OnTapped = func() {
		p.press("atan(")
	}
	p.View.Asin.OnTapped = func() {
		p.press("asin(")
	}
	p.View.Acos.OnTapped = func() {
		p.press("acos(")
	}
	p.View.Cos.OnTapped = func() {
		p.press("cos(")
	}
	p.View.Sin.OnTapped = func() {
		p.press("sin(")
	}
	p.View.Tan.OnTapped = func() {
		p.press("tan(")
	}
	p.View.Lgr.OnTapped = func() {
		p.press("log(")
	}
	p.View.Ln.OnTapped = func() {
		p.press("ln(")
	}
	p.View.BtnSqrt.OnTapped = func() {
		p.press("sqrt(")
	}
	p.View.BtnEqual.OnTapped = func() {
		p.Calculate()
	}

	// Credit events
	p.View.BtnCreditCalc.OnTapped = func() {
		p.Credit()
	}

	// Graphic
	p.View.BtnGraph.OnTapped = func() {
		p.Graphic()
	}
	return p
}

func (p *Presenter) press(data string) {
	p.View.BtnPress(data)
}

func (p *Presenter) pressAC() {
	p.View.PressAC()
}

func (p *Presenter) pressC() {
	p.View.PressC()
}

func (p *Presenter) Calculate() {
	var fl float64
	var err error
	if p.View.EnterX.Text != "" {
		fl, err = strconv.ParseFloat(p.View.EnterX.Text, 64)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fl = 0.0
	}
	ans, err := p.model.Calculate(p.View.Notation, fl)
	if err != nil {
		p.View.Notation = "Incorrect input"
	} else {
		an := strconv.FormatFloat(ans, 'f', -1, 64)
		p.View.Notation = an
	}
	p.View.EqualPress(p.View.Notation)
}

func (p *Presenter) Credit() {
	anu := p.View.RadioButton.Selected
	year := p.View.YearCred.Selected
	if p.View.SumCred.Text != "" && p.View.Percent.Text != "" && p.View.TimeCred.Text != "" {
		var ansF, ansL, ansM, ansT string
		sC, err := strconv.ParseFloat(p.View.SumCred.Text, 64)
		if err != nil {
			log.Fatal(err)
		}
		pr, err := strconv.ParseFloat(p.View.Percent.Text, 64)
		if err != nil {
			log.Fatal(err)
		}
		tC, err := strconv.Atoi(p.View.TimeCred.Text)
		if year == "лет" {
			tC *= 12
		}
		if err != nil {
			log.Fatal(err)
		}
		if anu == "Annuitet" {
			ansF, ansM, ansT = p.model.CreditAnnuitet(sC, pr, tC)
			ansF += ansL
			p.View.LabelMontPaySum.SetText(ansF)
			p.View.LabelPayOverSum.SetText(ansM)
			p.View.LabelTotalSum.SetText(ansT)
		} else if anu == "Diff" {
			ansF, ansL, ansM, ansT = p.model.CreditDff(sC, pr, tC)
			answerDiff := ansF + "..." + ansL
			p.View.LabelMontPaySum.SetText(answerDiff)
			p.View.LabelPayOverSum.SetText(ansM)
			p.View.LabelTotalSum.SetText(ansT)
		}
	}
}

func (p *Presenter) Graphic() {
	p.View.Image.RemoveAll()

	if p.View.Notation != "" {
		var a, b, h = -10.0, 10.0, 0.1
		N := (b-a)/h + 1
		pts := make(plotter.XYs, int(N))
		i := 0
		for x := a; x <= b; x = a + h*float64(i) {
			ans, _ := p.model.Calculate(p.View.Notation, x)
			if !math.IsNaN(ans) && !math.IsInf(ans, 0) {
				pts[i].X = x

				pts[i].Y = ans
			}
			i++
		}

		i = 0
		for math.IsInf(pts[i].Y, 0) || math.IsNaN(pts[i].Y) {
			i++
		}
		minY, maxY := pts[i].Y, pts[i].Y
		for idx := i + 1.0; i < int(N); i++ {
			if math.IsNaN(pts[idx].Y) && math.IsInf(pts[idx].Y, 0) {
				if pts[idx].Y <= minY {
					minY = pts[idx].Y
				}
				if pts[idx].Y >= maxY {
					maxY = pts[idx].Y
				}

			}
		}
		pl := plot.New()
		pl.Title.Text = "Graphic"
		pl.X.Label.Text = "X"
		pl.Y.Label.Text = "Y"
		pl.Y.Min = minY
		pl.Y.Max = maxY
		pl.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(1)
		pl.Add(s)

		err = pl.Save(350, 350, "scatter.png")
		if err != nil {
			log.Fatal(err)
		}
		graph := canvas.NewImageFromFile("scatter.png")
		graph.SetMinSize(fyne.Size{Width: 300, Height: 300})
		//	graph.FillMode = canvas.ImageFillOriginal
		p.View.Image.Add(graph)
	}

}
