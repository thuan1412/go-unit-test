package shapes

import "testing"

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := Rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := rec.Area(); got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCirlce_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
    t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			circle := Cirlce{
				radius: tt.fields.radius,
			}
			if got := circle.Area(); got != tt.want {
				t.Errorf("Cirlce.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArea(t *testing.T) {
  type args struct {
    r Rectangle
  }
}
