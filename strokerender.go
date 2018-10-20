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

func (renderer *TreeStrokeRenderer) Render(tree htree.Tree, gc GraphicsContext) error {
	it := htree.NewRegionIterator(tree, htree.NewVector(renderer.offsetX, renderer.offsetY, 0), renderer.scale)

	//it := htree.NewDimensionalIterator(tree, htree.NewVector(renderer.offsetX, renderer.offsetY, 0), renderer.scale)

	var container htree.NodeRegion

	for it.HasNext() {
		node := it.Next()

		if container == nil {
			container = node
		}
		// Draw the stroke
		branch := node.Node().Branch()
		ratios := tree.RatioSource().Ratios()
		if branch != nil {
			nodeRatio := ratios[node.Region().RatioIndexXY()] // tree.Ratio(tree.RatioIndex(node.Node, htree.RatioPlaneXY))
			nodeLeftRatio := ratios[branch.LeftIndex()]       // node.Node.Left().Ratio()
			dim := node.Region().Dimension()

			if branch.SplitType() == htree.SplitTypeHorizontal {
				y := dim.Top() + dim.Height()*htree.RatioNormalHeight(nodeRatio, nodeLeftRatio)
				if renderer.snap {
					y = math.Floor(y + 0.5)
				}

				gc.Line(dim.Left(), y, dim.Right(), y)
			} else {
				x := dim.Left() + dim.Width()*htree.RatioNormalWidth(nodeRatio, nodeLeftRatio)
				if renderer.snap {
					x = math.Floor(x + 0.5)
				}
				gc.Line(x, dim.Top(), x, dim.Bottom())
			}
		}
	}

	return nil
}
