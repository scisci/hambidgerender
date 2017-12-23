package hambidgerender

import (
	htree "github.com/scisci/hambidgetree"
)

const LeafFillKey = "fill"
const LeafFillDefaultFill = "#FF0000"

type LeafFillRenderer struct {
	offsetX float64
	offsetY float64
	scale   float64
	attrs   htree.NodeAttributes
	snap    bool
}

func NewLeafFillRenderer(offsetX, offsetY, scale float64, attrs htree.NodeAttributes) *LeafFillRenderer {
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

func (renderer *LeafFillRenderer) Render(tree *htree.Tree, gc GraphicsContext) error {
	leaves := tree.Leaves()

	nodeDimMap := htree.NewNodeDimensionMap(tree.Root(), renderer.offsetX, renderer.offsetY, renderer.scale)

	for _, leaf := range leaves {
		fill, err := renderer.attrs.Attribute(leaf, LeafFillKey)
		if err != nil {
			fill = LeafFillDefaultFill
		}

		gc.Fill(fill)

		dim, err := nodeDimMap.Dimension(leaf)
		if err != nil {
			return err
		}

		gc.Rect(dim.Left(), dim.Top(), dim.Width(), dim.Height())
	}

	return nil
}
