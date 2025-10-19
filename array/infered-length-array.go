package main
import ("fmt")

func main() {
	var arr1 = [...]int{1, 3, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}

	cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars, "\n")
	cars[2] = "Ferrari"
	fmt.Println(cars[2])

	fmt.Println(arr1)
	fmt.Println(arr2)
}
