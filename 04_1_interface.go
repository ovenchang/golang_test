package main

import (
	"fmt"
)

func main() {

	st := Student{Human{"mike", 20, "0955444444"}, "MIT", 5445.01}
	te := Teacher{Human{"lee", 50, "0955444444"}, "ARK", 36588}

	var i Men
	i = st
	i.SayHi() //Student:Hi, I am....

    //這二個都是不同型別的元素，但是他們實現了 interface 同一個介面
	x := []Men{st, te}
	x[0].SayHi() //Student:Hi, I am....
	x[1].SayHi() //Teacher:Hi, I am....

}

// interface 值
// 定義了一個 interface 的變數，那麼這個變數裡面可以存"實現這個 interface 的任意型別的物件"
// 接收者不能用參考型別!!!!
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名欄位
	school string
	loan   float32
}

type Teacher struct {
	Human  //匿名欄位
	school string
	salary float32
}

func (h Human) SayHi() {
	fmt.Printf("Human:Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func (h Human) Sing(str string) {
	fmt.Printf("Human:Hi, I am singing %s\n", str)
}
func (s Student) SayHi() { //過載 Human 的 SayHi 方法
	fmt.Printf("Student:Hi, I am %s you can call me on %s\n", s.name, s.phone)
}
func (s Teacher) SayHi() { //過載 Human 的 SayHi 方法
	fmt.Printf("Teacher:Hi, I am %s you can call me on %s\n", s.name, s.phone)
}

// Interface Men 被 Human,Student Teacher 實現
// 因為這3個型別都實現了這兩個方法
// interface 就是一組抽象方法的集合，它必須由其他非 interface 型別實現，而不能自我實現，
// Go 透過 interface 實現了 duck-typing:即"當看到一隻鳥走起來像鴨子、游泳起來像鴨子、叫起來也像鴨子，那麼這隻鳥就可以被稱為鴨子"。
type Men interface {
	SayHi()
	Sing(str string)
}
