package patterns

import (
	"fmt"
	"math"
)

type IShape interface {
	Accept(IVisitor)
	Draw()
	Area() float32
}

type IVisitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

type Visitor struct {
}

func (v *Visitor) VisitCircle(c *Circle) {
	c.Draw()
	fmt.Printf("Area: %.2f\n", c.Area())
}

func (v *Visitor) VisitRectangle(r *Rectangle) {
	r.Draw()
	fmt.Printf("Area: %.2f\n", r.Area())
}

type Circle struct {
	radius float32
}

func NewCircle(radius float32) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Accept(v IVisitor) {
	v.VisitCircle(c)
}

func (c *Circle) Draw() {
	fmt.Println("Draw Circle")
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	width  float32
	height float32
}

func NewRectangle(width float32, height float32) *Rectangle {
	return &Rectangle{
		width:  width,
		height: height,
	}
}

func (r *Rectangle) Accept(v IVisitor) {
	v.VisitRectangle(r)
}

func (r *Rectangle) Draw() {
	fmt.Println("Draw Rectangle")
}

func (r *Rectangle) Area() float32 {
	return r.width * r.height
}

func visitor() {
	shapes := []IShape{}
	shapes = append(shapes, NewCircle(5))
	shapes = append(shapes, NewRectangle(10, 5))
	for _, shape := range shapes {
		shape.Accept(&Visitor{})
	}
}
