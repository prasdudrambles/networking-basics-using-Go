package main
import "fmt"
func main(){

	fmt.Println("hello")
	array()
}





func array(){
	arr := [...]int{3, 1, 5, 10, 15}
	i:=0
	for true{
		fmt.Println(arr[i])
		i++
	}

}




func beyondHello(){
	var x int
	x = 3
	y :=4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum: ", sum, "prod: ", prod)
}

func learnMultiple(x, y int)(sum, prod int){
	return x+y, x*y
}


