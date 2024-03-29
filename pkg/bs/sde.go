package bs

import (
	"github.com/RyoAsano/stochastic_calculus/pkg/bm"
	"github.com/RyoAsano/stochastic_calculus/pkg/grd"
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/randgen"
	"github.com/RyoAsano/stochastic_calculus/pkg/sde"
	"github.com/RyoAsano/stochastic_calculus/pkg/stchprc"
	"github.com/RyoAsano/stochastic_calculus/pkg/vecfld"
)

func NewSDE(
	grid grd.Grid,
	gen randgen.RandGenerator,
	initpt float64,
	mu float64,
	sigma float64,
) sde.SDE {
	return &blackScholes{
		grid:      grid,
		initPoint: point.New(initpt),
		drift:     vecfld.NewAffine(mu, point.New(0.0)),
		diffusion: vecfld.NewAffine(sigma, point.New(0.0)),
		intr:      bm.New(grid, 1, gen, true),
	}
}

type blackScholes struct {
	grid      grd.Grid
	initPoint point.Point
	drift     vecfld.VectorField
	diffusion vecfld.VectorField
	intr      stchprc.Process
}

var _ sde.SDE = (*blackScholes)(nil)

func (bs *blackScholes) Grid() grd.Grid {
	return bs.grid
}

func (bs *blackScholes) Dim() int {
	return 1
}

func (bs *blackScholes) InitPoint() point.Point {
	return bs.initPoint
}

func (bs *blackScholes) Integrand(dim int) (vecfld.VectorField, error) {
	if dim == 1 {
		return bs.drift, nil
	} else if dim == 2 {
		return bs.diffusion, nil
	} else {
		return nil, sde.DimOutOfRangeErr{SDE: bs, GivenDim: dim}
	}
}

func (bs *blackScholes) Integrator() stchprc.Process {
	return bs.intr
}
