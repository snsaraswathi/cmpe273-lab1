package main
import "fmt"
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if  n==1 {
		return 1
	} else if n<0 {
 		fmt.Println("Input cann not be negative")
		return -1	
	} else{
		return (fibonacci(n-2) + fibonacci(n-1))
	}
}

func main() {
	var input int
	var output int
	fmt.Println("Enter input")
	fmt.Scanf("%d",&input)
	output=fibonacci(input)
	fmt.Println(output);
}
