package collision

import (
	"github.com/teomat/mater/aabb"
	"github.com/teomat/mater/transform"
	"log"

	"github.com/teomat/mater/vect"
)

// Base shape data.
// Holds data all shapetypes have in common.
type Shape struct {
	// The body this shape belongs to.
	Body        *Body
	Restitution float64
	Friction    float64
	AABB        aabb.AABB
	// The actual implementation of the shape.
	ShapeClass

	// If the shape is a sensor, collisions are reported but not resolved.
	IsSensor bool
}

type ShapeType int

const (
	ShapeType_Circle  = 0
	ShapeType_Segment = 1
	ShapeType_Polygon = 2
	ShapeType_Box     = 3
	numShapes         = iota
)

func (st ShapeType) ToString() string {
	switch st {
	case ShapeType_Circle:
		return "Circle"
	case ShapeType_Segment:
		return "Segment"
	case ShapeType_Polygon:
		return "Polygon"
	case ShapeType_Box:
		return "Box"
	default:
		return "Unknown"
	}
	panic("never reached")
}

type ShapeClass interface {
	ShapeType() ShapeType
	// Update the shape with the new transform and compute the AABB.
	update(xf transform.Transform) aabb.AABB
	// Returns if the given point is located inside the shape.
	TestPoint(point vect.Vect) bool

	marshalShape(shape *Shape) ([]byte, error)
	unmarshalShape(shape *Shape, data []byte) error
}

// Calls ShapeClass.Update and sets the new AABB.
func (shape *Shape) Update() {
	if shape.Body == nil {
		log.Printf("Error: uninitialized shape")
		return
	}

	shape.AABB = shape.ShapeClass.update(shape.Body.Transform)
}

// Returns shape.ShapeClass as CircleShape or nil.
func (shape *Shape) GetAsCircle() *CircleShape {
	if circle, ok := shape.ShapeClass.(*CircleShape); ok {
		return circle
	}

	return nil
}

// Returns shape.ShapeClass as PolygonShape or nil.
func (shape *Shape) GetAsPolygon() *PolygonShape {
	if poly, ok := shape.ShapeClass.(*PolygonShape); ok {
		return poly
	}

	return nil
}

// Returns shape.ShapeClass as SegmentShape or nil.
func (shape *Shape) GetAsSegment() *SegmentShape {
	if seg, ok := shape.ShapeClass.(*SegmentShape); ok {
		return seg
	}

	return nil
}

// Returns shape.ShapeClass as BoxShape or nil.
func (shape *Shape) GetAsBox() *BoxShape {
	if box, ok := shape.ShapeClass.(*BoxShape); ok {
		return box
	}

	return nil
}
