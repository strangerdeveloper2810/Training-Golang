package main

import (
	"fmt"
	"sync"
)

type Student struct {
	firstName string
	lastName  string
	email     string
	age       int
}

//Receiver function

// Value receiver || pointer receiver(stunded *Student)
func (stunted *Student) getFullName() string {
	stunted.firstName = "Trinh"
	stunted.lastName = "Developer"
	return stunted.firstName + " " + stunted.lastName
}

//gorountine

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Đánh dấu goroutine đã hoàn thành khi kết thúc

	fmt.Printf("Worker %d is working\n", id)
}

func wokerDemo(id int, channel chan int) {
	// gửi giá trị id qua channel
	channel <- id
}

func runWoker() int {
	// khởi tạo channel có kiểu dữ liệu int
	channel := make(chan int)

	// khởi chạy 1 goroutine để thực thi công việc
	go wokerDemo(1, channel)

	// nhận giá trị từ channel
	result := <-channel

	return result

}

func wokerWithForLoop(id int, channel chan int) {
	// gửi giá trị id qua channel
	channel <- id
}

func runWokerWithForLoop() []int {
	// Khởi tạo channel có kiểu dữ liệu int
	channel := make(chan int)

	// Số lượng goroutine
	numberWorker := 5

	// Slice để lưu kết quả
	results := make([]int, numberWorker)

	// Khởi chạy các goroutine
	for index := 0; index < numberWorker; index++ {
		go wokerWithForLoop(index, channel)
	}

	// Nhận kết quả từ các goroutine và lưu vào slice
	for index := 0; index < numberWorker; index++ {
		results[index] = <-channel
	}
	return results
}
func main() {
	// biến, kiểu dữ liệu, function, cấu trúc dữ liệu, oop, go routine, channel

	//result := runWoker()
	//fmt.Println("Received from workerDemo", result)

	results := runWokerWithForLoop()
	fmt.Println("Received from wokerWithForLoop", results)
}

//function trong golang

func demoVariable() {
	// khai báo biến
	// cách 1
	var email string = "trinhDeveloper@gmail.com"
	fmt.Println(email)

	fmt.Printf("%T\n", email)
	fmt.Printf("%V", email)
	fmt.Println()

	// khai báo biến shorthand
	fullName := "Trinh FrontEnd Developer"
	fmt.Println(fullName)
	fmt.Printf("%T\n", fullName)

	age := 23
	fmt.Println(age)
	fmt.Printf("%T\n", age)

	float := 3.14
	fmt.Println(float)
	fmt.Printf("%T\n", float)
}

func demoFunctionInGolang() {
	// run function

	//demoVariable()

	//sayHello()

	total := sumTwoNumber(1, 2)
	fmt.Println("tổng 2 số:", total)
}

func sayHello() {
	fmt.Println("Hacked By Stranger Developer")
}

func sumTwoNumber(a int, b int) int {
	return a + b
}

func demoDataStructureInGolang() {
	//slice giống như list

	sliceVariable := make([]string, 0)

	sliceVariable = append(sliceVariable, "a")
	sliceVariable = append(sliceVariable, "b")
	sliceVariable = append(sliceVariable, "c")
	sliceVariable = append(sliceVariable, "d")
	sliceVariable = append(sliceVariable, "e")
	fmt.Println(sliceVariable)

	// map -> như object bên js key-value
	mapVariable := make(map[string]int)
	mapVariable["age"] = 23
	mapVariable["classMajor"] = 1
	fmt.Println(mapVariable["age"])
	fmt.Println(mapVariable["ages"])

	// array

	arrayVariable := [1]string{}
	arrayVariable[0] = "value 1"
	fmt.Println(arrayVariable)
}

func demoStruct() {
	student := Student{
		firstName: "James",
		lastName:  "May",
		email:     "james.ma@gmail.com",
		age:       25,
	}

	fmt.Println(student)
	fmt.Println("age student is", student.age)

	fullName := student.getFullName()
	fmt.Println(fullName)

	fullNameOrigin := student.firstName + " " + student.lastName
	fmt.Println(fullNameOrigin)
}

// go routine

func hello() {
	fmt.Println("Hello from Goroutine in side hello function!")
}

func demoGoroutine() {
	/*
		go hello()

		// anonymous function trong golang
		go func() {
			fmt.Println("Hello from Goroutine inside anonymous function!")
		}()

		// đợi 1 chút để đảm bảo goroutine có thời gian để chạy
		time.Sleep(1 * time.Second)
	*/

	var wg sync.WaitGroup
	for index := 0; index < 10; index++ {
		wg.Add(1)             // thêm 1 goroutine vào trong WaitGroup
		go worker(index, &wg) // khởi chạy goroutine worker
	}

	wg.Wait() // đợi cho tất cả các goroutine hoàn thành
	fmt.Println("All worker finished")

	name := "Trình is a Frontend Developer"
	fmt.Println(name)
}
