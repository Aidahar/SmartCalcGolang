package model

/*
	#cgo CFLAGS: -I.
	#cgo LDFLAGS: -L. -lm
	#include <stdlib.h>
	#include "parse_string.h"
	#include "s21_helper.h"
	#include "math.h"
	#include "stack.h"
	#include "calculate.h"
*/
import "C"

import (
	"errors"
	"log"
	"math"
	"os"
	"strconv"
	"unsafe"
)

type Model struct {
	logFile *os.File
}

func NewModel(logPath string) *Model {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Ошибка открытия файла", err)
	}
	return &Model{
		logFile: logFile,
	}
}

func (m *Model) Calculate(n string, fl float64) (float64, error) {
	cData := C.CString(n)
	defer C.free(unsafe.Pointer(cData))
	var notation string
	cNotation := C.CString(notation)
	defer C.free(unsafe.Pointer(cNotation))
	err := C.parse_string(cData, cNotation)
	if err == 0 {
		return 0.0, errors.New("0")
	}
	x := C.double(fl)
	a := C.calculate(cNotation, x)
	an := float64(a)
	m.logFile.WriteString("Calculate: " + n + "\n")
	return an, nil
}

func (m *Model) CreditAnnuitet(tel, proc float64, month int) (string, string, string) {
	ansAnu := m.calcAnuitet(tel, proc, month)
	percAnu := m.calcAnuPerc(tel, proc, month)
	totalAnu := m.calcAnuDolg(tel, proc, month)
	anu := strconv.FormatFloat(ansAnu, 'f', 2, 64)
	anuPerc := strconv.FormatFloat(percAnu, 'f', 2, 64)
	anuTotal := strconv.FormatFloat(totalAnu, 'f', 2, 64)
	return anu, anuPerc, anuTotal
}

func (m *Model) CreditDff(tel, proc float64, month int) (string, string, string, string) {
	ansR := m.calcDifMMin(tel, proc, month)
	ansL := m.calcDifMMax(tel, proc, month)
	percDiff := m.calcDifPere(tel, proc, month)
	totalDiff := m.calcDifTotal(tel, proc, month)
	aR := strconv.FormatFloat(ansR, 'f', 2, 64)
	aL := strconv.FormatFloat(ansL, 'f', 2, 64)
	pD := strconv.FormatFloat(percDiff, 'f', 2, 64)
	tD := strconv.FormatFloat(totalDiff, 'f', 2, 64)
	return aR, aL, pD, tD
}

func (m *Model) calcAns(ans float64) float64 {
	return math.Floor(ans*100.0+0.5) / 100.0
}

func (m *Model) calcAnuitet(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	ps := proc / (100 * 12)
	ans = tel * (ps / (1 - math.Pow((1+ps), float64(-month))))
	ans = m.calcAns(ans)
	return ans
}

func (m *Model) calcAnuPerc(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	anu := m.calcAnuitet(tel, proc, month)
	ans = (anu * float64(month)) - tel
	ans = m.calcAns(ans)
	return ans
}

func (m *Model) calcAnuDolg(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	ap := m.calcAnuPerc(tel, proc, month)
	ans = tel + ap
	ans = m.calcAns(ans)
	return ans
}

func (m *Model) calcDifMMin(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	ans = tel / float64(month)
	i := proc / 100.0
	ans = m.calcAns(ans)
	ans += (tel - ans*0) * i / 12
	return ans
}

func (m *Model) calcDifMMax(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	ans = tel / float64(month)
	i := proc / 100.0
	ans += (tel - ans*(float64(month-1))) * i / 12
	return ans
}

func (m *Model) calcDifPere(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	ep := tel / float64(month)
	i := proc / 100
	for idx := 0; idx < month; idx++ {
		ans += m.calcDifAns(ans, ep, tel, i, idx)
	}
	ans = ans - tel
	return ans
}

func (m *Model) calcDifTotal(tel, proc float64, month int) float64 {
	var ans float64 = 0.0
	ep := tel / float64(month)
	i := proc / 100
	for idx := 0; idx < month; idx++ {
		ans += m.calcDifAns(ans, ep, tel, i, idx)
	}
	return ans
}

func (m *Model) calcDifAns(ans, ep, tel, i float64, idx int) float64 {
	return ans + (ep + (tel-ep*float64(idx))*i/12)
}
