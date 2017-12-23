package hambidgerender

import (
	"math"
	"strconv"
)

type GraphicsContextCall interface {
	Name() string
	String() string
	Equals(other GraphicsContextCall) bool
}

type GraphicsContextLine struct {
	x1, y1, x2, y2 float64
}

func (m *GraphicsContextLine) Name() string {
	return "Line"
}

func (m *GraphicsContextLine) Equals(other GraphicsContextCall) bool {
	if l, ok := other.(*GraphicsContextLine); ok {
		return math.Abs(l.x1-m.x1) < RenderEpsilon &&
			math.Abs(l.y1-m.y1) < RenderEpsilon &&
			math.Abs(l.x2-m.x2) < RenderEpsilon &&
			math.Abs(l.y2-m.y2) < RenderEpsilon
	}

	return false
}

func (m *GraphicsContextLine) String() string {
	return "Line{" +
		strconv.FormatFloat(m.x1, 'f', -1, 64) + ", " +
		strconv.FormatFloat(m.y1, 'f', -1, 64) + ", " +
		strconv.FormatFloat(m.x2, 'f', -1, 64) + ", " +
		strconv.FormatFloat(m.y2, 'f', -1, 64) + "}"
}

type GraphicsContextRect struct {
	x, y, width, height float64
}

func (m *GraphicsContextRect) Name() string {
	return "Rect"
}

func (m *GraphicsContextRect) String() string {
	return "Rect{" +
		strconv.FormatFloat(m.x, 'f', -1, 64) + ", " +
		strconv.FormatFloat(m.y, 'f', -1, 64) + ", " +
		strconv.FormatFloat(m.width, 'f', -1, 64) + ", " +
		strconv.FormatFloat(m.height, 'f', -1, 64) + "}"
}

func (m *GraphicsContextRect) Equals(other GraphicsContextCall) bool {
	if r, ok := other.(*GraphicsContextRect); ok {
		return math.Abs(r.x-m.x) < RenderEpsilon && math.Abs(r.y-m.y) < RenderEpsilon && math.Abs(r.width-m.width) < RenderEpsilon && math.Abs(r.height-m.height) < RenderEpsilon
	}

	return false
}

type GraphicsContext interface {
	Line(x1, y1, x2, y2 float64)
	Rect(x, y, w, h float64)
}

type GraphicsContextRecorder struct {
	Calls []GraphicsContextCall
}

func NewGraphicsContextRecorder() *GraphicsContextRecorder {
	return &GraphicsContextRecorder{}
}

func (gc *GraphicsContextRecorder) Line(x1, y1, x2, y2 float64) {
	gc.AddCall(&GraphicsContextLine{x1, y1, x2, y2})
}

func (gc *GraphicsContextRecorder) Rect(x, y, width, height float64) {
	gc.AddCall(&GraphicsContextRect{x, y, width, height})
}

func (gc *GraphicsContextRecorder) AddCall(call GraphicsContextCall) {
	gc.Calls = append(gc.Calls, call)
}
