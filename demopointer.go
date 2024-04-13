package main

import "fmt"

func main() {
	// dùng con trỏ trỏ đến vùng nhớ(địa chỉ)

	x := 10

	// địa chỉ của biến x
	pointerX := &x

	// thay đổi giá trị của biến x
	y := x
	y = 20

	fmt.Println("pointerX =", pointerX)
	fmt.Println("value X =", x)
	fmt.Println("value Y Before=", y)
	pointerY := &y
	*pointerY = 30
	fmt.Println("value Y After =", y)
}
