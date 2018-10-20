package hambidgerender

import (
	htree "github.com/scisci/hambidgetree"
	"github.com/scisci/hambidgetree/algo"
	"github.com/scisci/hambidgetree/attributors"
)

const LeafFillKey = "fill"
const LeafFillDefaultFill = "#FF0000"

type LeafFillRenderer struct {
	offsetX float64
	offsetY float64
	scale   float64
	attrs   attributors.NodeAttributes
	snap    bool
}

func NewLeafFillRenderer(offsetX, offsetY, scale float64, attrs attributors.NodeAttributes) *LeafFillRenderer {
	return &LeafFillRenderer{
		offsetX: offsetX,
		offsetY: offsetY,
		scale:   scale,
		attrs:   attrs,
		snap:    false,
	}
}

func (renderer *LeafFillRenderer) Snap(snap bool) {
	renderer.snap = snap
}

func (renderer *LeafFillRenderer) Render(tree htree.Tree, gc GraphicsContext) error {
	leaves := algo.FindLeaves(tree)
	regionMap := htree.NewTreeRegionMap(tree, htree.NewVector(renderer.offsetX, renderer.offsetY, 0), renderer.scale)
	//nodeDimMap := htree.NewNodeDimensionMap(tree, htree.NewVector(renderer.offsetX, renderer.offsetY, 0), renderer.scale)

	for _, leaf := range leaves {
		fill, err := renderer.attrs.Attribute(leaf.ID(), LeafFillKey)
		if err != nil {
			fill = LeafFillDefaultFill
		}

		gc.Fill(fill)

		dim := regionMap[leaf.ID()].Dimension()
		gc.Rect(dim.Left(), dim.Top(), dim.Width(), dim.Height())
	}

	return nil
}
