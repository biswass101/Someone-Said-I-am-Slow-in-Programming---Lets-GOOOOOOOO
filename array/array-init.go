package main
import ("fmt")

func main() {
	arr1 := [5]int{} // not init
	arr2 := [5]int{1, 2} // partially init
	arr3 := [5]int{1,2,3,4,5} // fully init
	arr4 := [5]int{1:3, 4:3} // specific init

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

}
