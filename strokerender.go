package hambidgerender

import (
	htree "github.com/scisci/hambidgetree"
	"math"
)

const RenderEpsilon = 0.0000001

type TreeStrokeRenderer struct {
	offsetX float64
	offsetY float64
	scale   float64
	snap    bool
}

func NewTreeStrokeRenderer(offsetX, offsetY, scale float64) *TreeStrokeRenderer {
	return &TreeStrokeRenderer{
		offsetX: offsetX,
		offsetY: offsetY,
		scale:   scale,
		snap:    false,
	}
}

func (renderer *TreeStrokeRenderer) Snap(snap bool) {
	renderer.snap = snap
}

func (renderer *TreeStrokeRenderer) Render(tree *htree.Tree, gc GraphicsContext) error {
	it := htree.NewDimensionalIterator(tree, htree.NewVector(renderer.offsetX, renderer.offsetY, 0), renderer.scale)

	var container *htree.DimensionalNode

	for it.HasNext() {
		node := it.Next()

		if container == nil {
			container = node
		}
		// Draw the stroke
		if !node.IsLeaf() {
			nodeRatio := tree.Ratio(tree.RatioIndex(node.Node, htree.RatioPlaneXY))
			nodeLeftRatio := node.Node.Left().Ratio()
			if node.Split().IsHorizontal() {
				y := node.Dimension.Top() + node.Dimension.Height()*nodeRatio/nodeLeftRatio
				if renderer.snap {
					y = math.Floor(y + 0.5)
				}

				gc.Line(node.Dimension.Left(), y, node.Dimension.Right(), y)
			} else {
				x := node.Dimension.Left() + node.Dimension.Width()*nodeLeftRatio/nodeRatio
				if renderer.snap {
					x = math.Floor(x + 0.5)
				}
				gc.Line(x, node.Dimension.Top(), x, node.Dimension.Bottom())
			}
		}
	}

	return nil
}
