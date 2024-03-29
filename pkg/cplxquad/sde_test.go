package cplxquad

import (
	"testing"

	"github.com/RyoAsano/stochastic_calculus/pkg/grd"
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/testutil"
)

var seed int64 = 99

func TestSDE(t *testing.T) {
	terminal := 1.0
	size := 10
	grid := grd.NewEquiDistGrid(size, terminal)
	x, y := 1.0, 1.0
	theSDE := NewSDE(grid, x, y, testutil.NewStdNorm(seed))

	testutil.Check2DimSDE(theSDE, point.New(x, y), driftChecker, diffusionChecker, seed, t)
}
