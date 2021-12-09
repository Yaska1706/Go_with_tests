package sminter

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 4.0}
	got := rectangle.Perimeter()
	want := 12.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{2.0, 4.0}, hasArea: 8.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{6.0, 8.0}, hasArea: 24.0},
		{name: "Square", shape: Square{5.0}, hasArea: 25.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
