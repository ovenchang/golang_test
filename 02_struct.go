package main

import (
	"fmt"
)

//////struct/////
//宣告新的型別，作為其它型別的屬性或欄位的容器。稱之為struct

type person struct {
	name string
	age  int
}

// 比較誰年紀大
func compAge(p1, p2 person) (person, int) {
	temp := p1.age - p2.age
	if temp > 0 {
		return p1, temp
	} else {
		return p2, temp
	}
}

// 匿名欄位
// 當匿名欄位是一個 struct 的時候，
// 那麼這個 struct 所擁有的全部欄位都被隱含的引入了當前定義的這個 struct
// 若有同名  可用層級來存取  st.figure.phone  st.phone
type figure struct {
	height int
	weight int
	phone  string
}

type testInt []int

type student struct {
	figure  //student 有 figure的欄位 匿名欄位
	name    string
	age     int
	testInt // 自訂型別
	int     //內建型別
	phone   string
}

///

func main() {

	p1 := person{name: "edd", age: 20}
	p2 := person{name: "edd", age: 85}
	p3 := new(person) //透過 new 函式分配一個指標，此處 P 的型別為*person
	p3.name, p3.age = "tom", 25
	tperson, tAge := compAge(p1, p2)
	fmt.Println(tAge, tperson.name)
	fmt.Println(p2.name)

	//
	st := student{figure{52, 155, "0922"}, "edd", 20, testInt{}, 0, "0933"}
	st.figure = figure{65, 185, "0922"}
	st.testInt = testInt{1, 2}
	fmt.Println(st.height, st.weight, st.name, st.age, st.figure.height, st.testInt)
}
