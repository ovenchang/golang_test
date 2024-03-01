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

/*************************************************/
//空 interface
//空 interface(interface{})不包含任何的 method，
//正因為如此，所有的型別都實現了空 interface。
//空 interface 對於描述起不到任何的作用(因為它不包含任何的 method），
//但是空 interface 在我們需要"儲存任意型別的數值"的時候相當有用，因為它可以儲存任意型別的數值。

var  a interface{}
var i int = 1
s:="hello"
a = i
a = s 

//一個函式把 interface{} 作為參數，那麼他可以接受任意型別的值作為參數，
//如果一個函式回傳 interface{}，那麼也就可以回傳任意型別的值。

func a1(i interface{}) interface{} {
    return i
}
fmt.Println(a1(123))
fmt.Println(a1("helloo"))


//interface 函式參數
//interface 的變數可以持有任意實現該 interface 型別的物件，
//這給我們編寫函式(包括 method)提供了一些額外的思考，
//我們是不是可以透過定義 interface 參數，讓函式接受各種型別的參數。

// fmt.Println 是我們常用的一個函式，但是你是否注意到它可以接受任意型別的資料。
// 開啟 fmt 的原始碼檔案，你會看到這樣一個定義:
type Stringer interface {
	String() string
}

//也就是說，任何實現了 String 方法的型別都能作為參數被 fmt.Println 呼叫

type Human struct {
	name string
}

func (h Human) String() string { //過載 fmt.Stringer 的 String() 不然會依預設方式秀出
	return "i am " + h.name + "......."
}

func main() {
	Bob := Human{"Bob"}
	fmt.Println(Bob)
}



/*************************************************/
//interface 變數儲存的型別
//我們知道 interface 的變數裡面可以儲存任意型別的數值(該型別實現了 interface)。
//那麼我們怎麼反向知道這個變數裡面實際儲存了的是哪個型別的物件呢

//Comma-ok 斷言
//直接判斷是否是該型別的變數： value, ok = element.(T)，
//這裡 value 就是變數的值，ok 是一個 bool 型別，element 是 interface 變數，T 是斷言的型別。
//如果 element 裡面確實儲存了 T 型別的數值，那麼 ok 回傳 true，否則回傳 false

type Element interface{} //空interface
type List []Element
type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "i am " + p.name + " age " + strconv.Itoa(p.age)
}

func main() {
	list := List{1, "Hello", Person{"bob", 3}}
	
    //第一種 Comma-ok 斷言
    value, ok := list[0].(int)
	fmt.Println(value, ok) //1 true

	if value, ok := list[2].(Element); ok { //if 裡面允許初始化變數
		fmt.Println(value, ok) //i am.....  true
	}

    //第二種 switch
    //`element.(type)`語法不能在 switch 外的任何邏輯裡面使用，
    //如果你要在 switch 外面判斷一個型別就使用`comma-ok`。
    for index,element := range list{
        switch value:=element.(type){
                case int:
                    fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
                case Person:
                    fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
                default:
                    fmt.Println("list[%d] is of a different type", index)
        }
    }
}

/*****************************/
//嵌入 interface
//如果一個 interface1 作為 interface2 的一個嵌入欄位，
//那麼 interface2 隱式的包含了 interface1 裡面的 method

/**********/
//反射
//反射就是能檢查程式在執行時的狀態。我們一般用到的套件是 reflect 套件
//使用 reflect 一般分成三步
//1.要去反射是一個型別的值(這些值都實現了空 interface)，首先需要把它轉化成 reflect 物件
t := reflect.TypeOf(i)    //得到型別的 Meta 資料，透過 t 我們能取得型別定義裡面的所有元素
v := reflect.ValueOf(i)   //得到"實際的值"，透過 v 我們取得儲存在裡面的值，還可以去改變值

//2.轉化為 reflect 物件之後我們就可以進行一些操作了，也就是將 reflect 物件轉化成相應的值，例如
tag := t.Elem().Field(0).Tag  //取得定義在 struct 裡面的標籤
name := v.Elem().Field(0).String()  //取得儲存在第一個欄位裡面的值

//取得反射值能回傳相應的型別和數值/var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())

///3.那麼反射的欄位必須是可修改的，我們前面學習過傳值和傳參考，
//這個裡面也是一樣的道理。反射的欄位必須是可讀寫的意思是
//如果要修改相應的值，必須這樣寫
var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)