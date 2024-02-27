package main

import (
	"fmt"
)

func main() {

	mike := Student{Human{"mike", 35, "0955444555"}, "MIT", 35212.02}
	//tom:=Student{Human:{"tom",45,"0922222222"},"abc",65442.14}
	//定義 Men 型別的變數 i
	var i Men
	i=mike
	fmt.Println(mike.school)
}

//什麼是 interface
//interface 是一組 method 簽名的組合，我們透過 interface 來定義物件的一組行為
//Student 實現了三個方法：SayHi、Sing、BorrowMoney；
//Employee 實現了 SayHi、Sing、SpendSalary。
//上面"這些方法的組合"稱為 interface(被物件 Student 和 Employee 實現)
//Employee 沒有實現這個 interface：SayHi、Sing 和 BorrowMoney，因為 Employee 沒有實現 BorrowMoney 這個方法。

//interface 型別
//interface 型別定義了一組方法，如果某個物件實現了某個介面的所有方法，則此物件就實現了此介面。
//struct 有一些method 把這些method 定義為一組interface

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

func (h *Human) sayHi() {
	fmt.Println("Human: i am %s ,my phone is %s", h.name, h.phone)
}
func (h *Human) sing(str string) {
	fmt.Println("Human: i am singing %s", str)
}

func (s *Student) sayHi() { //過載 Human 的 SayHi 方法
	fmt.Println("Student: i am %s ,my phone is %s", s.name, s.phone)
}

func (s *Student) borrow(amount float32) {
	s.loan += amount
}

// 定義 interface
// interface 可以被任意的物件實現。Men interface 被 Human、Student實現。
// 同理，一個物件可以實現任意多個 interface，例如上面的 Student 實現了 Men 和 Young 兩個 interface。
type Men interface { //Human 實現了此介面
	sayHi()
	sing(str string)
}
type Young interface { //Student 實現了此介面 也實現了Men介面
	sayHi(hi string)
	borrow(amount float32)
}

//interface 值
//定義了一個 interface 的變數，那麼這個變數裡面可以存實現這個 interface 的任意型別的物件
