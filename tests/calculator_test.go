package tests

import (
	"calculator/internal/model"
	"errors"
	"testing"
)

func TestCalculatorOk(t *testing.T) {
	tests := []struct {
		name     string
		notation string
		x        float64
		want     float64
	}{
		{
			name:     "simple add",
			notation: "2+2*2",
			x:        0.0,
			want:     6.0,
		},
		{
			name:     "simple add with x",
			notation: "2+2*2+x",
			x:        2.0,
			want:     8.0,
		},
		{
			name:     "simple trigonometry",
			notation: "sin(2.0)",
			x:        0.0,
			want:     0.9092974268256817,
		},
		{
			name:     "simple trigonometry with x",
			notation: "ln(x)",
			x:        4.0,
			want:     1.3862943611198906,
		},
	}
	m := model.NewModel("log.log")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.Calculate(tt.notation, tt.x)
			if got != tt.want {
				t.Errorf("Got %f want %f", got, tt.want)
			} else if err != nil {
				t.Error(err)
			}
		})
	}
}

func TestCalculatorError(t *testing.T) {
	tests := []struct {
		name     string
		notation string
		x        float64
		want     error
	}{
		{
			name:     "error",
			notation: "2+(2*2))",
			x:        0.0,
			want:     errors.New("0"),
		},
		{
			name:     "error1",
			notation: "2+2*2+xs",
			x:        2.0,
			want:     errors.New("0"),
		},
		{
			name:     "error trigonometry",
			notation: "cos(",
			x:        0.0,
			want:     errors.New("0"),
		},
		{
			name:     "error trigonometry2",
			notation: "ln(x))",
			x:        4.0,
			want:     errors.New("0"),
		},
	}
	m := model.NewModel("log.log")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := m.Calculate(tt.notation, tt.x)
			if err.Error() != tt.want.Error() {
				t.Errorf("Got %s want %s", err, tt.want)
			}
		})
	}
}
