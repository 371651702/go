package main

import (
	_ "fmt"
	"fmt"
	"strconv"
)

const LINES int  = 10

func ShowYangHuiTriangle()  {
	nums := []int{}
	for i := 0; i < LINES; i++ {
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}

		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			}else {
				value = nums[length - i] + nums[length - i - 1]
			}
			nums = append(nums,value)
			fmt.Print(value," ")
		}
		fmt.Println("")
	}
}

func multiplicationTable()  {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			var ret string
			if i * j < 10 && j != 1 {
				ret = " " + strconv.Itoa(i * j)
			} else {
				ret = strconv.Itoa(i * j)
			}

			fmt.Print(j," * ",i," = ",ret,"  ")
		}
		fmt.Print("\n")
	}
}

func printBook( book Book ) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

const MAX int  = 3

type Book struct {
	title string
	author string
	subject string
	book_id int
}

func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func Factorial(u uint64)(result uint64)  {
	if (u > 0) {
		result = u * Factorial(u - 1)
		return result
	}
	return 1
}

type phone interface {
	call()
}
type NokiaPhone struct {
	
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i an nokia,i can call you!")
}

type IPhone struct {
	
}

func (iPhone IPhone) call()  {
	fmt.Println("I am iPhone, I can call you!")
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string  {
	strFormat := `
	Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
return fmt.Sprint(strFormat,de.dividee)
}

func Divide(varDividee int,varDivider int) (result int,errorMsg string) {
	if varDivider == 0{
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return

	} else {
		return varDividee /varDivider,""
	}
}

func main()  {
	if result,errorMsg := Divide(100,10); errorMsg == "" {
		fmt.Println("100/10 = ",result)
	}

	if _,errorMsg := Divide(100,0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
	//var phone phone
	//
	//phone = new(NokiaPhone)
	//phone.call()
	//
	//phone = new(IPhone)
	//phone.call()

	//var sum int = 17
	//var count int = 5
	//var mean float32
	//
	//mean = float32(sum)/float32(count)
	//fmt.Printf("mean 的值为: %f\n",mean)
	//var i int = 15
	//fmt.Printf("%d 的阶乘是 %d\n",i,Factorial(uint64(i)))
	//var countryCapitalMap map[string]string
	//countryCapitalMap = make(map[string]string)
	//
	///* map插入key - value对,各个国家对应的首都 */
	//countryCapitalMap [ "France" ] = "Paris"
	//countryCapitalMap [ "Italy" ] = "罗马"
	//countryCapitalMap [ "Japan" ] = "东京"
	//countryCapitalMap [ "India" ] = "新德里"
	//
	//for country := range countryCapitalMap{
	//	fmt.Println(country,"首都是:",countryCapitalMap[country])
	//}
	//
	//captial,ok := countryCapitalMap["India"]
	//if (ok){
	//	fmt.Println("美国的首都是", captial)
	//} else {
	//	fmt.Println("美国的首都不存在")
	//}

	//nums := []int{2,3,5}
	//
	//sum := 0
	//for _, num := range nums{
	//	sum += num
	//}
	//fmt.Println("sum: ",sum)
	//
	//for i, num := range nums{
	//	if num == 3 {
	//		fmt.Println("index: ",i)
	//	}
	//}
	//
	//kvs := map[string]string{"a":"apple","b":"banana"}
	//for k,v := range kvs{
	//	fmt.Println("%s -> %s\n",k,v)
	//}
	//
	//for i,c := range "go"{
	//	fmt.Println(i,c)
	//}
	//var numbers = make([]int,3,5)
	//printSlice(numbers)
	//create struct
	//fmt.Println(Book{"万历十五年","www.wdl.com","wwww",2222})
	//
	////use key => value
	//fmt.Println(Book{title:"万历十五年",author:"www.wdl.com",subject:"wdl",book_id:2222})

	//var Book1 Book
	////var Book2 Book
	//
	///* book 1 描述 */
	//Book1.title = "Go 语言"
	//Book1.author = "www.runoob.com"
	//Book1.subject = "Go 语言教程"
	//Book1.book_id = 6495407

	/* book 2 描述 */
	/*Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700*/

	/* 打印 Book1 信息 */
	//printBook(Book1)

	/* 打印 Book2 信息 */
	//printBook(Book2)

	//var struct_pointer *Book

	//print(struct_pointer)
	//print(struct_pointer.title)

	//var st = &Book1
	//print("\nwdl: ",st.title)


	//a := []int{10,20,50}
	//var i int
	//var ptr [MAX]*int;
	//
	//for i = 0; i < MAX; i++ {
	//	ptr[i] = &a[i]
	//
	//}
	//for i = 0; i < MAX; i++ {
	//	fmt.Printf("a[%d] = %d\n",i,*ptr[i])
	//}
	//var a int
	//var ptr *int
	//var pptr **int
	//
	//a = 3000
	///* 指针 ptr 地址 */
	//ptr = &a
	///* 指向指针 ptr 地址 */
	//pptr = &ptr
	//
	///* 获取 pptr 的值 */
	//fmt.Printf("变量 a = %d\n", a )
	//fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	//fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

}

