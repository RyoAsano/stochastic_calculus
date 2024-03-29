package mthd

import (
	"github.com/RyoAsano/stochastic_calculus/pkg/point"
	"github.com/RyoAsano/stochastic_calculus/pkg/sde"
	"github.com/RyoAsano/stochastic_calculus/pkg/stchprc"
)

type eulMaru struct{}

func NewEulerMaruyama() Method {
	return &eulMaru{}
}

var _ Method = (*eulMaru)(nil)

func (m *eulMaru) Apply(sde sde.SDE) stchprc.Process {
	return &mthdAppliedPrc{mthd: m, sde: sde}
}

func (m *eulMaru) To(from point.Point, dx DX) (point.Point, error) {
	// V1(x)dx1+...+Vn(x)dxn
	totalDir := point.Origin(from.Dim())

	for vecFld, incr := range dx {
		vec, err := vecFld.At(from)
		if err != nil {
			return nil, err
		}
		dir := point.Mul(vec, incr) // V(x)*dx

		totalDir, err = point.Add(totalDir, dir)
		if err != nil {
			return nil, err
		}
	}
	return point.Add(from, totalDir)
}

func (m *eulMaru) Modify(point point.Point) point.Point {
	return point
}
