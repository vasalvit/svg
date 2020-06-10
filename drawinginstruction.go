package svg

// InstructionType tells our path drawing library which function it has
// to call
type InstructionType int

// These are instruction types that we use with our path drawing library
const (
	MoveInstruction InstructionType = iota
	CircleInstruction
	CurveInstruction
	LineInstruction
	CloseInstruction
	PaintInstruction
)

// CurvePoints are the points needed by a bezier curve.
type CurvePoints struct {
	C1 *Tuple
	C2 *Tuple
	T  *Tuple
}

// DrawingInstruction contains enough information that a simple drawing
// library can draw the shapes contained in an SVG file.
//
// The struct contains all necessary fields but only the ones needed (as
// indicated byt the InstructionType) will be non-nil.
type DrawingInstruction struct {
	Kind           InstructionType
	M              *Tuple
	CurvePoints    *CurvePoints
	Radius         *float64
	StrokeWidth    *float64
	Fill           *string
	Stroke         *string
	StrokeLineCap  *string
	StrokeLineJoin *string
}
