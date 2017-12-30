package hambidgerender

import (
	htree "github.com/scisci/hambidgetree"
	"testing"
)

var strokeTests = []struct {
	Tree  *htree.Tree
	Calls []GraphicsContextCall
}{
	{
		Tree: htree.NewGridTree2D(1),
		Calls: []GraphicsContextCall{
			&GraphicsContextLine{0.5, 0.0, 0.5, 1.0},
		},
	},
	{
		Tree: htree.NewGridTree2D(2),
		Calls: []GraphicsContextCall{
			&GraphicsContextLine{0.5, 0.0, 0.5, 1.0},
			&GraphicsContextLine{0.0, 0.5, 0.5, 0.5},
			&GraphicsContextLine{0.5, 0.5, 1.0, 0.5},
		},
	},
}

func TestTreeStrokeRenderer(t *testing.T) {
	for i, test := range strokeTests {
		// Generates a 4x4 grid
		renderer := NewTreeStrokeRenderer(0, 0, 1)
		gc := NewGraphicsContextRecorder()

		renderer.Render(test.Tree, gc)

		if len(gc.Calls) != len(test.Calls) {
			t.Errorf("Tree stroke test %d failed, lengths don't match, expected %d, got %d", i, len(test.Calls), len(gc.Calls))
			continue
		}

		for c := 0; c < len(gc.Calls); c++ {
			if !test.Calls[c].Equals(gc.Calls[c]) {
				t.Errorf("Tree stroke test %d failed, call %d doesn't match, expected %v, got %v", i, c, test.Calls[c], gc.Calls[c])
			}
		}

	}

}
