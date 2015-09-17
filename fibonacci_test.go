package main
import "fmt"
import "testing"


func TestInputChar(t *testing.T) {
        result := fibonacci(-1)

        if result != -1 {
        }
}

func TestInputZero(t *testing.T) {
        result := fibonacci(0)
	fmt.Println(result);
        if result != 0 {
		t.Errorf("Expected result: 0.. Actual result: %d", result)
        }
}
func TestInputOne(t *testing.T) {
	result := fibonacci(1)
	fmt.Println(result);
	if result != 1 {
		t.Errorf("Expected result: 1.. Actual result: %d", result)
	}
}
func TestInputTwo(t *testing.T) {
	result := fibonacci(2)
	fmt.Println(result);
	if result != 1{
		t.Error("Expected result: 1.. Actual result: %d", result)
	}
}	
func TestInputNine(t *testing.T) {
        result := fibonacci(9)
	fmt.Println(result);
        if result != 34{
                t.Error("Expected result: 34.. Actual result: %d", result)
        }
}

