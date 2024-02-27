package main

import (
	"fmt"
)

// //////定義變數////////
// var name int = 1 //一般用 var 方式來定義全域性變數
// var name = 1
// name := 1 //只能用在函式內部
//_, b := 34, 35 // _ 丟棄值

//////常數////
//const abc string = '123' //型別可省略
/*const (
    abc=1
)*/

// ///內建基礎型別////
//boolean
// var abc boolean=true
//數值
// var abcd byte = 'd' //100
//string
//abc:=`as //它沒有字元轉義，換行也將原樣輸出
//ds`
//字串是不可變的
//abc:="aa"
//abc[0] 報錯 要用切片
//s := "hello"
//s = "c" + s[1:] // 字串雖不能更改，但可進行切片(slice)操作
//使用+運算子來連線兩個字串

//字串與數值相合
/*num := 1
str := "string"
numberStr := strconv.Itoa(num) //第一種
result := numberStr + " " + str
result = fmt.Sprintf("%d and %s", num, str) //第二種
*/
/*
//字串取代
originalString := "Hello, world! Hello, Golang!"
// 将 "Hello" 替换为 "Hi"
newString := strings.Replace(originalString, "Hello", "Hi", -1)
//不会修改原始字符串，而是返回一个新的字符串，因为Go中的字符串是不可变的。
*/


//錯誤型別
//err := errors.New("emit macho dwarf: elf header corrupted")
//if err != nil {
//    fmt.Print(err)
//}

////////分組宣告//////////
//import(
//    "fmt"
//    "os"
//)
//const  var.....

// /////iota 列舉//////
// const 中 const 中"每增加一行"加 1
/*const (
	x = iota // x == 0
	w        // 常數宣告省略值時，預設和之前一個值的字面相同 w=1
	z = "b"
	k        //b
	m = iota //4
)
const v = iota // 每遇到一個 const 關鍵字，iota 就會重置，此時 v == 0
*/

////////Go 程式設計的一些規則///////
//Go 之所以會那麼簡潔，是因為它有一些預設的行為：
//大寫字母開頭的變數是可匯出的，也就是其它套件可以讀取的，是公有變數；小寫字母開頭的就是不可匯出的，是私有變數。
//大寫字母開頭的函式也是一樣，相當於 class 中的帶 public 關鍵詞的公有函式；小寫字母開頭的就是有 private 關鍵詞的私有函式。

// ////////array slice map //////////////////
// array  array 不是參考型別
//var arr [10]int // 宣告了一個 int 型別的陣列
//arr[0]=21
//b := [10]int{1, 2, 3} // 宣告了一個長度為 10 的 int 陣列，其中前三個元素初始化為 1、2、3，其它預設為 0
//長度也是陣列型別的一部分，因此 [3]int 與[4]int是不同的型別，陣列也就不能改變長度
//func arrayCheck(f [10]int) {
//	fmt.Println(f)
//}
//二維
//easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

//改變全域 array的值 array 不是參考型別
/*var iSlice [5]int
func main() {
	iSlice = [5]int{123, 123}
	change(&iSlice)
}
func change(s *[5]int) {
	s[0] = 200
	fmt.Println(iSlice)
}
*/

//slice
//slice並不是真正意義上的動態陣列，而是一個參考型別。slice總是指向一個底層array
//var fslice []int
//fslice := []byte{'a', 'b', 'c', 'd'}
//func slicecheck(f []byte) {
//fmt.Println(f)
//}
//fslice[:] fslice[1:2] //return slice


//slice是參考型別，所以當參考改變其中元素的值時，其它的所有參考都會改變該值
//  arr1 := [3]int{1, 2, 3}
//	slice1 := arr1[:3]
//	slice1[0] = 9
//	fmt.Println(arr1) //[9,2,3]
//slice2 := make([]int, len(slice1))
//copy(slice2, slice1)
//fmt.Println(slice2)//[9,2,3]
//slice3 := []int{}
//slice3 = append(slice3, slice1...)
//slice3[2] = 100
//fmt.Println(slice3)//[9,2,100]
//fmt.Println(slice2)//[9,2,3]

//len    取得 slice 的長度
//cap    取得 slice 的最大容量
//append 向 slice 裡面追加一個或者多個元素，然後回傳一個和 slice 一樣型別的slice
//copy   函式 copy 從源 slice 的src中複製元素到目標dst，並且回傳複製的元素的個數

///////////map//////////
//var numbers map[string]int //key string  value int
//numbers = make(map[string]int)
//另一種初始化
//rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }
// value,ok:=rating["c"] // map 有兩個回傳值，第二個回傳值，如果存在 ok 為 true

//map1 := map[string]string{"a": "aa", "b": "bb"}
//	for k, v := range map1 {
//		fmt.Println(k, v)
//	}
//	for _, v := range map1 {
//		fmt.Println(v)
//	}

/**map的長度是不固定的，也就是和 slice 一樣，也是一種參考型別  array 不是參考型別  
var mMap map[string]string

func main() {

	mMap = map[string]string{"a": "aa", "b": "bb"}
	changeMap(mMap)

}
func changeMap(s map[string]string) {
	s["a"] = "200"
	fmt.Println(mMap) //map[a:200,b:bb]
}
**/

//map是無序的
//map的長度是不固定的，也就是和 slice 一樣，也是一種參考型別
//len(map) delete(rating["c"])
//map和其他基本型別不同，它不是 thread-safe，在多個 go-routine 存取時，必須使用 mutex lock 機制


//自訂型別  實際上只是一個定義了一個別名
//type ages int
//type money float32
//type months map[string]int


// 函式作為值、型別
// 設定要傳遞的func 型別

////make、new 操作 ////
/*
make 用於內建型別（map、slice 和channel）的記憶體分配。
new  用於各種型別的記憶體分配。

new(T) 分配了零值填充的 T 型別的記憶體空間，並且回傳其地址，即一個*T型別的值
make(T, args) 回傳一個有初始值(非零)的 T 型別，而不是*T

new回傳指標。
make回傳初始化後的（非零）值。

導致這三個型別有所不同的原因是指向資料結構的參考在使用前必須被初始化。
例如，一個slice，是一個包含指向資料（內部array）的指標、長度和容量的三項描述符；
在這些專案被初始化之前，slice為nil。
對於slice、map和 channel 來說，make初始化了內部的資料結構，填充適當的值。


var mMap map[string]string
var fslice []int
type point struct {
	x int
}

func main() {
	fslice = make([]int, 3) //[0,0,0]  make([]T, len, cap)T 元素类型。 len 元素的个数） cap 切片容量可省略
	fslice[0] = 12
	mMap = make(map[string]string) //print 出來map[]
	mMap["a"] = "aa"
	fmt.Println(mMap) //map[a:aa]
	tp := new(point) //print 出來 &{0} 指標
	tp.x = 1
	fmt.Println(tp) //&{1} 指標
}

*/




type testInt func(int, int) int
type testBool func() bool

func main() {
rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }


	map1 := map[string]string{"a": "aa", "b": "bb"}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	for _, v := range map1 {
		fmt.Println(v)
	}

	x := 1
	if x > 0 && x < 10 {
		fmt.Println("ds")
	}

a:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				continue a
			}
			fmt.Print(i)
		}

	}

	if x := checkMax(6, 8); x < 10 {
		fmt.Println("x=", x)
	}

	switch x {
	case 1, 2:
		fmt.Println("x=1")
		fallthrough //強制執行後面程式碼
	case 3:
		fmt.Println("x=3")
		fallthrough //強制執行後面程式碼
	default:
		fmt.Println("x other")
	}

	fmt.Println(map1)
	fmt.Println(map1)
	fmt.Println(map1)
	fmt.Println(map1)
	ss, cc := addAndmultiple(12, 25)
	fmt.Println("addAndmultiple", ss, cc)
	ss, cc = addAndmultiple2(14, 33)
	fmt.Println("addAndmultiple2", ss, cc)
	fmt.Println("someArgs", someArgs(1, 2, 3, 4, 5, 6, 7, 8, 9))

	//指標
	pt := 1
	changePtVal(&pt) //傳pt的記憶體位址
	fmt.Println("changePtVal", pt)

	//defer
	deferFunc()

	//函式作為值、型別
	//設定要傳遞的func 型別
	//如果沒有帶入參數 回傳值  就是func()類型
	type testInt func(int, int) int
	fmt.Println("calSum", calSum(3, 45, addCal))

	//
	throwPanic(checkPanic)

}

func checkMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addAndmultiple(a, b int) (int, int) {
	return a + b, a * b
}

func addAndmultiple2(a, b int) (add2 int, multiple2 int) {
	add2 = a + b
	multiple2 = a * b
	return
}

func someArgs(arg ...int) (sum int) {
	for _, v := range arg {
		sum += v
	}
	return
}

func changePtVal(x *int) int { //*int 指標型別
	*x = *x + 1
	return *x
}

// defer defer 後指定的函式會在函式退出前呼叫
// defer 是採用後進先出模式
func deferFunc() {
	fmt.Println("deferFunc start")
	defer fmt.Println("deferFunc exit1")
	defer fmt.Println("deferFunc exit2")
	fmt.Println("deferFunc 2nd")
	fmt.Println("deferFunc 3rd")
}

// 函式作為值、型別
func calSum(a int, b int, f testInt) int {
	return f(a, b)
}

func addCal(a, b int) int {
	return a + b
}
func multipleCal(a, b int) int {
	return a * b
}

// //////////////////////////
// panic 中斷原有的控制流程
// panic可以直接呼叫 panic 產生。也可以由執行時錯誤產生，例如存取越界的陣列。
func checkPanic() bool {
	defer fmt.Println("exit") //延遲函式會正常執行
	panic("123")
	return true
}

// 可以讓進入 panic 狀態的 goroutine 恢復過來
// 在延遲函式中有效
// 正常的執行過程中，呼叫 recover 會回傳nil
func throwPanic(f testBool) {
	fmt.Println("throwPanic")
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recoverd")
		}
	}()
	f()
}

func arrayCheck(f [10]int) {
	fmt.Println(f)
}
