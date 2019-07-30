package main

import (
	"fmt"
	"reflect"
)

/*import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string
}

type Admin struct {
	User
	title string
}

func main() {
	var u Admin
	t := reflect.TypeOf(u)
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}*/

/*
type Student struct {
	Name string
	Age  int
}

type Class struct {
	Name    string
	Student *Student
	Grade   int
	School  string
}

func main() {
	s:=Student{"LiLei",20}


	fmt.Println("*****reflect.TypeOf****")
	typ:=reflect.TypeOf(s)
	fmt.Println(typ)
	fmt.Println(typ.Name())
	fmt.Println(typ.Kind())

	fmt.Println("*****reflect.ValueOf****")
	val:=reflect.ValueOf(s)
	fmt.Println(val.Kind())//basic type

	fmt.Println("*****change value****")
	//newValue:=reflect.New(typ)

	fmt.Printf("the type is %s\n",typ.String())
	fmt.Printf("the name is %s\n",val.FieldByName("Name").String())


	a:=10
	//a_val:=reflect.TypeOf(&a)
	//fmt.Println(a_val.Name())
	//fmt.Println(a_val.Kind())
	p:=reflect.ValueOf(&a)
	fmt.Println("can set of p:",p.CanSet())
	v:=p.Elem()
	fmt.Println("can set of v:",v.CanSet())


	t:=reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t)

}*/

//type User struct {
//	UserName string
//	Age      int
//}
//
//type Admin struct {
//	User
//	title string
//}
//
//func main() {
//	admin := &Admin{User{"Jack", 23}, "NT"}
//
//	/********type*********/
//	t := reflect.TypeOf(admin)
//	fmt.Println("t type is ",t)
//	t=t.Elem()
//	fmt.Println("t elem is ",t)
//
//	/********value*********/
//	v := reflect.ValueOf(admin)
//	fmt.Println("v value is ",v)
//	v=v.Elem()
//	fmt.Println("v elem is ",v)
//
//
//	/**********tag***********/
//	var u User
//	 tag:=reflect.TypeOf(u)
//	 f,_:=tag.FieldByName("UserName")
//
//	 fmt.Println(f.Tag)
//	 fmt.Println(f.Tag.Get("field"))
//	 fmt.Println(f.Tag.Get("type"))
//
//
//
//
//	/*fmt.Println(v.FieldByName("title").String())
//	fmt.Println(v.FieldByName("Age").Int())
//	fmt.Println(v.FieldByIndex([]int{0, 1}).Int())
//
//	user := User{"jack", 23}
//	r:=reflect.ValueOf(user)
//
//	fmt.Println(r.FieldByName("UserName").Interface())
//	fmt.Println(r.FieldByName("Age").Interface())*/
//}

//type data struct {
//	b byte
//	x int32
//	y int64
//}
//
//type action interface {
//	say(string)
//	eat(string)
//}
//
//type animal struct {
//	name string
//	age  int
//}
//
//type cat struct {
//	animal
//}
//
//func (c *cat) say(s string) {
//	fmt.Println("my name is:", s)
//}
//
//func (c *cat) eat(s string) {
//	fmt.Println("i like eat:", s)
//}
//
//func (c *cat) Foot(s string) {
//	fmt.Println("i havr four foot!")
//}
//
//type dog struct {
//	animal
//}
//
//func (d *dog) say(s string) {
//	fmt.Println("my name is:", s)
//}
//
//func (d *dog) eat(s string) {
//	fmt.Println("i like eat:", s)
//}
//
//func main() {
//	var d data
//	t := reflect.TypeOf(d)
//	fmt.Println(t.Size(), t.Align(), t.FieldAlign())
//
//	var c *cat
//	ct := reflect.TypeOf(c)
//	fmt.Println(ct.Size(), ct.Align(), ct.FieldAlign())
//	fmt.Println(ct.NumMethod())
//	//fmt.Println(*ct.Method(0))
//
//	fmt.Println("name:", ct.Elem().Name())
//
//	fmt.Println(ct.MethodByName("Foot"))
//
//	fmt.Println("path:", ct.Elem().PkgPath())
//
//	fmt.Println("String():", ct.String())
//
//	fmt.Println("kind:", ct.Kind())
//
//	//fmt.Println("implememts:",ct.Implements())
//	fmt.Println(ct.Elem().Field(0))
//
//	s := []int{0}
//	fmt.Println(ct.Elem().FieldByIndex(s))
//
//	fmt.Println(ct.Elem().FieldByName("animal"))
//
//	fmt.Println("********value*********")
//	dd := dog{animal{name: "ErHa", age: 3}}
//	dt := reflect.TypeOf(dd)
//
//	dv := reflect.ValueOf(dd)
//	fmt.Println("dt:", dt)
//	fmt.Println("cv:", dv)
//	fmt.Println("cv.CanAddr()", dv.CanAddr())
//
//	fmt.Println("fieldByName:", dv.FieldByName("animal"))
//	fmt.Println(dt.FieldByName("animal"))
//
//	fmt.Println(dv.Type())
//	fmt.Println(reflect.Indirect(dv).Type())
//	//fmt.Println(ct.Elem().In(0))
//}

func main() {

	var a = []int{10, 11, 12}
	t := reflect.TypeOf(a)
	fmt.Println("t:", t.Elem())
	fmt.Println("t kind:", t.Kind().String())

	// 根据反射创建array,数组类型位[10]int
	tArray := reflect.ArrayOf(10, reflect.TypeOf(10))
	fmt.Println("t_array;", tArray)

	//b := new([10]int)
	var b [10]int
	vv := reflect.ValueOf(&b)
	fmt.Println("vv:", vv)
	fmt.Println("vv type:",reflect.TypeOf(b))
	fmt.Println("vv kind:", vv.Kind().String())
	fmt.Println("vv Elem kind:", vv.Elem().Kind())
	fmt.Println("vv Elem type:", vv.Elem().Type())
	//reflect.SliceOf(vv.Elem().Type())
	//switch vv.Elem().Kind() {
	//
	//}

	v := reflect.MakeSlice(t, 3, 3)
	fmt.Println(v)
	//fmt.Println(reflect.TypeOf(10).Elem())
}
