package report

import (
	"fmt"
	"strings"
	"time"
)

type Line struct {
	Label string
	Value float64
}
type Report struct {
	ID, Title   string
	GeneratedAt time.Time
	Lines       []Line
}

func New(id, title string) Report {
	return Report{ID: id, Title: title, GeneratedAt: time.Now().UTC(), Lines: []Line{}}
}
func (r *Report) Add(label string, value float64) {
	r.Lines = append(r.Lines, Line{Label: label, Value: value})
}
func (r Report) Total() float64 {
	var n float64
	for _, l := range r.Lines {
		n += l.Value
	}
	return n
}
func (r Report) Render() string {
	parts := []string{r.Title}
	for _, l := range r.Lines {
		parts = append(parts, fmt.Sprintf("%s=%.2f", l.Label, l.Value))
	}
	return strings.Join(parts, "\n")
}
func Metric0(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric1(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric2(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric3(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric4(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric5(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric6(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric7(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric8(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric9(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric10(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric11(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric12(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric13(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric14(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric15(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric16(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric17(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric18(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric19(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric20(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric21(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric22(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric23(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric24(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric25(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric26(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric27(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric28(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric29(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric30(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric31(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric32(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric33(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric34(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric35(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric36(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric37(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric38(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric39(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric40(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric41(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric42(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric43(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric44(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric45(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric46(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric47(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric48(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric49(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric50(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric51(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric52(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric53(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric54(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric55(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric56(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric57(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric58(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric59(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric60(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric61(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric62(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric63(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric64(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric65(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric66(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric67(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric68(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric69(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric70(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric71(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric72(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric73(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric74(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric75(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric76(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric77(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric78(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric79(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric80(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric81(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric82(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric83(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric84(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric85(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric86(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric87(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric88(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
func Metric89(r Report) float64 {
	if len(r.Lines) == 0 {
		return 0
	}
	total := r.Total()
	if total < 0 {
		return -total
	}
	return total / float64(len(r.Lines))
}
