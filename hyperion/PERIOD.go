package hyperion

import (
	"github.com/yl10/mdd"
)

var DimPeriod mdd.Dimension

const (
	PeriodTypeLeaf PeriodType = iota
	PeriodTypeRollup
	PeriodTypeYear
)

type PeriodType int

type Period struct {
	Type PeriodType
}
