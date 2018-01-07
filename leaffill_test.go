package hambidgerender

import (
	htree "github.com/scisci/hambidgetree"
	"testing"
)

func TestLeafFillRenderer(t *testing.T) {
	tree := htree.NewGridTree2D(1)
	leaves := tree.Leaves()
	attributer := htree.NewNodeAttributer()
	attributer.SetAttribute(leaves[0].ID(), LeafFillKey, "#00FF00")
	attributer.SetAttribute(leaves[1].ID(), LeafFillKey, "#0000FF")
	renderer := NewLeafFillRenderer(0, 0, 1, attributer)
	gc := NewGraphicsContextRecorder()
	err := renderer.Render(tree, gc)
	if err != nil {
		t.Errorf("Error rendering leaf fill (%v)", err)
	}

	expectedCalls := []GraphicsContextCall{
		&GraphicsContextFill{"#00FF00"},
		&GraphicsContextRect{0.0, 0.0, 0.5, 1.0},
		&GraphicsContextFill{"#0000FF"},
		&GraphicsContextRect{0.5, 0.0, 0.5, 1.0},
	}

	if len(expectedCalls) != len(gc.Calls) {
		t.Errorf("expected %d calls got %d", len(expectedCalls), len(gc.Calls))
	}

	for i, expectedCall := range expectedCalls {
		if !expectedCall.Equals(gc.Calls[i]) {
			t.Errorf("Call %d doesn't match expected %v, got %v", i, expectedCall, gc.Calls[i])
		}
	}
}
