package mathutil

import (
	"math"

	"github.com/RyoAsano/stochastic_calculus/pkg/consts"
)

func CloseBtw(a float64, b float64) bool {
	return math.Abs(a-b) < consts.CloseUpTo
}
