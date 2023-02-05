package cplxbsl

import (
	"bitbucket.org/AsanoRyo/experiments/pkg/bm"
	"bitbucket.org/AsanoRyo/experiments/pkg/grd"
	"bitbucket.org/AsanoRyo/experiments/pkg/point"
	"bitbucket.org/AsanoRyo/experiments/pkg/randgen"
	"bitbucket.org/AsanoRyo/experiments/pkg/sde"
	"bitbucket.org/AsanoRyo/experiments/pkg/stchprc"
	"bitbucket.org/AsanoRyo/experiments/pkg/vecfld"
)

func NewSDE(
	grid grd.Grid,
	gen randgen.RandGenerator,
	x float64,
	y float64,
) sde.SDE {
	drift, diffusion := NewVecFld()
	return &complexBessel{
		initPt:    point.New(x, y),
		grid:      grid,
		drift:     drift,
		diffusion: diffusion,
		intr:      bm.New(grid, 1, gen, true),
	}
}

type complexBessel struct {
	initPt    point.Point
	grid      grd.Grid
	drift     vecfld.VectorField
	diffusion vecfld.VectorField
	intr      stchprc.Process
}

var _ sde.SDE = (*complexBessel)(nil)

func (bsl *complexBessel) Grid() grd.Grid {
	return bsl.grid
}

func (bsl *complexBessel) Dim() int {
	return 2
}

func (bsl *complexBessel) InitPoint() point.Point {
	return bsl.initPt
}

func (bsl *complexBessel) Integrand(dim int) (vecfld.VectorField, error) {
	if dim == 1 {
		return bsl.drift, nil
	} else if dim == 2 {
		return bsl.diffusion, nil
	} else {
		return nil, sde.DimOutOfRangeErr{SDE: bsl, GivenDim: dim}
	}
}

func (bsl *complexBessel) Integrator() stchprc.Process {
	return bsl.intr
}