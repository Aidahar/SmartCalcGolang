package tests

import (
	"calculator/internal/model"
	"testing"
)

func TestAnnuitetCredit(t *testing.T) {
	m := model.NewModel("log.log")
	tests := []struct {
		name  string
		tel   float64
		proc  float64
		month int
		wantT string
		wantP string
		wantM string
	}{
		{
			name:  "annuitet",
			tel:   10000.0,
			proc:  22,
			month: 36,
			wantT: "381.90",
			wantP: "3748.40",
			wantM: "13748.40",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, gotP, gotM := m.CreditAnnuitet(tt.tel, tt.proc, tt.month)
			if gotT != tt.wantT || gotP != tt.wantP || gotM != tt.wantM {
				t.Errorf("GotT %s wantT %s, gotP %s wantP %s, gotM %s wantM %s", gotT, tt.wantT, gotP, tt.wantP, gotM, tt.wantM)
			}
		})
	}
}

func TestDifCredit(t *testing.T) {
	m := model.NewModel("log.log")
	tests := []struct {
		name   string
		tel    float64
		proc   float64
		month  int
		wantTl string
		wantTr string
		wantP  string
		wantM  string
	}{
		{
			name:   "annuitet",
			tel:    10000.0,
			proc:   22,
			month:  36,
			wantTl: "461.11",
			wantTr: "282.87",
			wantP:  "3391.67",
			wantM:  "13391.67",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTl, gotTr, gotP, gotM := m.CreditDff(tt.tel, tt.proc, tt.month)
			if gotTl != tt.wantTl || gotTr != tt.wantTr || gotP != tt.wantP || gotM != tt.wantM {
				t.Errorf("GotTl %s wantTl %s, GotTr %s wantTr %s, gotP %s wantP %s, gotM %s wantM %s", gotTl, tt.wantTl, gotTr, tt.wantTr, gotP, tt.wantP, gotM, tt.wantM)
			}
		})
	}
}
