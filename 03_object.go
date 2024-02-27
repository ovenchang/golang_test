package main

import (
	"fmt"
)

const (
	pi = 3.14159
)

func main() {
	c := Circle{radius: 10}
	fmt.Println(c.area("this is c:"))
	c.setX(50) // method 繼承 point的method
	fmt.Println(c)
	c.setRadius(32.6)//method 重寫 若point 已經有 setRadius， Circle 想要實現自己的 setRadius
	fmt.Println(c)

	map1 := mMap{"a": 1, "b": 36}
	map1.add1()
	fmt.Println(map1)
}

//函式當作 struct 的欄位一樣來處理
//函式的另一種形態，帶有接收者的函式，我們稱為method
//method 的語法如下：
//func (r ReceiverType) funcName(parameters) (results)

// 使用 method 的時候重要注意幾點
// 雖然 method 的名字一模一樣，但是如果接收者不一樣，那麼 method 就不一樣
// method 裡面可以存取接收者的欄位
// 呼叫 method 透過.存取，就像 struct 裡面存取欄位一樣
type point struct {
	x int
}

type Circle struct {
	point
	radius float32
	myArea float32
}

// 專屬於Circle的 area func
// 沒加* 是以值傳遞 無法更改c的欄位值
// 指標作為 receiver
// 如果一個 method 的 receiver 是 *T，
// 你可以在一個 T 型別的變數 V 上面呼叫這個 method，
// 而不需要 &V 去呼叫這個 method
func (c *Circle) area(str string) string {
	mArea := c.radius * c.radius * pi
	c.myArea = mArea
	return fmt.Sprintf("%s %f", str, mArea)
}

// method 繼承
// 如果匿名欄位實現了一個 method，
// 那麼包含這個匿名欄位的 struct 也能呼叫該 method
func (p *point) setX(n int) {
	p.x = n
}

// method 重寫
// 若point 已經有 setRadius， Circle 想要實現自己的 setRadius
func (p *point) setRadius(n int) {
	p.x = n
}
func (c *Circle) setRadius(n float32) {
	c.radius = n
}

// 自訂型別也可以有自已型別的專屬method
// map是參考型別所以裡面的值改變 map的值也會改
type mMap map[string]int

func (m mMap) add1() {
	for k, v := range m {
		m[k] = v + 1
	}
}
