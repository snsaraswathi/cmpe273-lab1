package main
import "testing"
import "fmt"

func TestCirclePerimeter(t *testing.T) {
	c :=Circle{radius: 5}
        var shapeCir Shaper
        shapeCir=c
		
	result:= shapeCir.Perimeter()
	fmt.Println(result)
	if(result != 31.400000000000002){
		t.Errorf("Perimeter of a Circle: Expected result: 31.400000000000002.. Actual result: %f", result)
	}
}

func TestCircleArea(t *testing.T) {
        c :=Circle{radius: 5}
        var shapeCir Shaper
        shapeCir=c

        result:= shapeCir.Area()
	fmt.Println(result)
        if(result != 78.53981633974483){
                t.Errorf("Area of a Circle: Expected result: 78.53981633974483.. Actual result: %f", result)
        }
}

func TestRectPerimeter(t *testing.T) {
        r :=Rectangle{length:5, width: 5}
	var shapeRect Shaper
        shapeRect=r

        result:= shapeRect.Perimeter()
	fmt.Println(result)
        if(result != 20 ){
                t.Errorf("Perimeter of a Rectangle: Expected result- 20 .. Actual result- %d", result)
        }
}

func TestRectArea(t *testing.T) {
        r :=Rectangle{length:5, width: 5}
        var shapeRect Shaper
        shapeRect=r

        result:= shapeRect.Area()
	fmt.Println(result)
        if(result != 25 ){
                t.Errorf("Area of a Rectangle: Expected result:25 .. Actual result: %d", result)
        }
}
