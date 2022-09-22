package main

import (
	"fmt"
)

// import (
// 	"fmt"
// 	"time"

// 	"github.com/google/go-cmp/cmp"
// 	"github.com/jimuyang/lets-go/subpack"
// )

// // 注意下面2个type的区别
// // 这里只是别名？
// type (
// 	str = string
// 	s   = str
// )

// // type definition 新的类型定义
// type strname string

// func main() {

// 	fmt.Println("Hello World")
// 	time.Sleep(3 * time.Millisecond)
// 	// fmt.Println("Hello World")
// 	fmt.Println(`"hello go\`)
// 	fmt.Println("\"hello go\\")
// 	arr := []int{1, 2, 3}
// 	subpack.ReverseInt(arr)
// 	fmt.Println(arr)

// 	fmt.Println(cmp.Diff("hello", "hi"))

// 	var name string
// 	var strn strname = "strname" // 这里可以赋值的原因是因为strn的underlying type是string
// 	var name1 str = "name1"
// 	name = name1
// 	// name = strn   // x's type V and T have identical underlying types and at least one of V or T is not a defined type
// 	// strn = name   // 虽然name的underlying type也是string 但至少得有一个是基本类型
// 	fmt.Println(name)
// 	fmt.Println(strn)
// }

// boolean literal
type boolean bool

var b boolean = true

// integer literal
type integer int

var i integer = 1

// slice literal
type strarr []string

//var strings []string = []string{"sfw"}
//var strs strarr = strings

// pointer literal
type stringptr *string

var string1 string = "11"
var ptr stringptr = &string1

func main111() {

	fmt.Println(fmt.Sprintf("1111", 222))



	//now := time.Now()
	//currentDay := now.Unix() / 3600 / 24
	//println(fmt.Sprintf("date = '%v' and bigint(poi_id) %% %v = %v",
	//	now.AddDate(0, 0, -1).Format("20060102"), 14, currentDay%14))
	//
	//spew.Dump(struct {
	//	name string
	//	age  int
	//}{"nihao", 11})
	//
	//spew.Dump(fmt.Sprintf("sortedCount: %v, totalCount: %v", 19009, 13443))
	//
	//arr := []int{1, 2, 3}
	//fmt.Println(arr[:len(arr)-1])
	//
	//var r []int
	//r = append(r, 1)
	//fmt.Println(r)
	//fmt.Println(r == nil)
	//fmt.Println(len(r))
	//
	//r = make([]int, 0)
	//fmt.Println(r)
	//fmt.Println(r == nil)
	//fmt.Println(len(r))
	//
	//fmt.Printf("%v > %v", "id", 1)
	//
	//// strings := []string{"string"}
	//// var strarr1 strarr
	//// strarr1 = strings
	//
	//// 	my1 := mystruct{"str"}
	//// 	var your1 yourstruct
	//// 	your1 = my1
	//for i := 0; i < 0; i++ {
	//	fmt.Println(1)
	//}
}
