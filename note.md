```go
package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	email     string
	age       int
}

//Receiver function

func (stunted Student) getFullName() string {
	stunted.firstName = "Trinh"
	stunted.lastName = "Developer"
	return stunted.firstName + " " + stunted.lastName
}

func main() {
	// biến, kiểu dữ liệu, function, cấu trúc dữ liệu, oop, go routine, channel

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

```

###### h5 Giá trị của biến fullName sẽ trả về là **Trinh Developer**, còn fullNameOrigin sẽ là **James May**.

###### h5 Vì trong receiver function getFullName(), chúng ta **đang thực hiện trên bản sao** của đối tượng **Stundent**. Khi chúng ta truyền tham số với từ khoá **stunded** vào receiver function, một bản sao của đối tượng gốc được tạo ra và các thay đổi được thực hiện trên bản sao này. Điều này không ảnh hướng đến đối tượng gốc mà bạn đã khởi tạo cho hàm main().

###### h5 Để thực sự thay đổi giá trị của đối tượng gốc, chúng ta cần truyền đối trượng **dưới dạng con trỏ(\*)**. Khi chúng ta truyền con trỏ, receiver function sẽ tham chiếu đến đối tượng gốc và các thay đổi sẽ được áp dụng trực tiếp lên đối tượng đó.

###### h5 goroutine khá giống với async await của javascript. Goroutine thường kết hợp với các cơ chế như sync.WaitGroup, channel để đồng bộ hoá các goroutine, còn async await thường kết hợp với promise, Trong Go, quản lý bộ nhớ của goroutine là trách nhiệm của runtime, trong khi trong JavaScript, bạn cần phải quản lý bộ nhớ một cách cẩn thận để tránh memory leaks khi sử dụng async/await trong Node.js hoặc trong các ứng dụng web.
