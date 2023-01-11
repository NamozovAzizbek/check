package main
import "fmt"
func main(){
	var a, b int
	fmt.Scan(&a, &b)
	isMultiple(a, b)
}
func isMultiple(a, b int) {
for(i := 0; i < b; i++){
	if b % i == 0 && i + b%i == a{
		fmt.Println(i, b%i)
	}
}
}