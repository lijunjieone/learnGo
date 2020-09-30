package main

import (
	"container/list"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"
)

func main() {
	// test()
	// printFlag()

	// testStringLen()

	// testType()
	// testAddMethod()

	// testStruct()

	// testSlice()

	// testMap()
	// testSyncMap()

	// testList()

	// testFor()

	// testChannel()

	// testFunc()

	// testFunc1()

	// testFunc2()

	// testFlag()
	// testInterface1()

	// testInterface2()

	// testClosure2()
	// testClosure3()

	// testPrintln()

	// testDefer()

	// testFileSize()

	// testCrash()

	// testStruct1()
	// testStruct2()
	// testStruct3()

	// testStruct4()
	// testStruct8()
	// testStruct7()
	// testStruct6()
	// testStruct5()

	testInterface3()

}

type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	println(data)
	return nil
}

func (d *file) CanWrite() bool {
	return true
}

func testInterface3() {
	f := new(file)
	f.WriteData("this is a file.")

	var m DataWriter

	m = f
	m.WriteData("this is a file by interface")

}

type BasicColor struct {
	R, G, B float32
}

type Color2 struct {
	BasicColor
	Alpha float32
}

func testStruct8() {

	var c Color2
	c.B = 1
	c.G = 2
	c.B = 3

	c.Alpha = 50.0

	println(c)
}

type class struct {
}

func (c *class) do(v int) {
	println("call method do:")
	println(v)
}

func do(v int) {
	println("call function do:")
	println(v)
}

func testStruct7() {
	var cc class

	cc.do(100)

	do(100)

	var delegate func(int)

	delegate = cc.do

	delegate(200)

	delegate = do

	delegate(200)
}

type MyInt int

func (m MyInt) isZero() bool {
	return m == 0
}

func (m MyInt) add(other int) int {
	return other + int(m)
}

func testStruct6() {
	var b MyInt

	b = 100

	println(b.isZero())

	println(b.add(20))
}

type Point struct {
	X int
	Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

func testStruct5() {
	p1 := Point{1, 2}
	p2 := Point{3, 4}

	result := p1.Add(p2)

	println(result)
}

type Cat struct {
	Color string
	Name  string
}

func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

func testStruct4() {
	var cat1 Cat
	cat1.Name = "cat1"
	cat1.Color = "black"

	println(cat1)

	var cat2 = NewCatByName("cat2")

	println(*cat2)

}

type Address struct {
	Province    string
	City        string
	ZipCode     string
	PhoneNumber string
}

func testStruct3() {
	addr := Address{
		"shanxi",
		"linfen",
		"002400",
		"15812345678",
	}

	println(addr)
}

type People struct {
	name  string
	child *People
}

func testStruct2() {
	relation := &People{
		name: "gradepa",
		child: &People{
			name: "dad",
			child: &People{
				name: "me",
			},
		},
	}

	printPeople(*relation)
}

func printPeople(p People) {
	println(p.name)
	if p.child != nil {
		printPeople(*p.child)
	}
}

type Color struct {
	R, G, B byte
}

func testStruct1() {
	var p Point
	p.X = 19
	p.Y = 20

	p.X = 30

	var p2 = new(Point)

	var p3 = &p

	print(typeof(p2))
	print(typeof(p))
	print(typeof(p3))

	println(p2.X)
	println(p.X)
	println(p3.X)

	println(p)
}
func testCrash() {

	defer println("do thing1")
	defer println("do thing2")
	panic("crash")
}
func testFileSize() {
	println(fileSize("first.go"))
}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)

	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()

	if err != nil {
		return 0
	}

	size := info.Size()

	return size
}
func testDefer() {
	println("defer begin")
	defer println(1)
	println(2)
	defer println(3)
	println("defer end")
}

func testPrintln() {
	println("a", "b", 10, []int{1, 2, 34})
	fmt.Println()
	println("a")
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func println(a ...interface{}) (n int, err error) {
	// for _, s := range a {
	// 	fmt.Println(typeof(s))
	// 	// fmt.Printf("%d %T \n", s, s)
	// }
	return fmt.Println(a)
}

func playerGen(name string) func() (string, int) {
	hp := 150

	return func() (string, int) {
		return name, hp
	}
}

func testClosure3() {
	generator := playerGen("BigDog")

	name, hp := generator()

	fmt.Println(name, hp)
}
func Accumulate(defaultValue int) func() int {
	return func() int {
		defaultValue++

		return defaultValue
	}
}

func testClosure2() {
	accu := Accumulate(1)

	println(accu())
	println(accu())

	fmt.Printf("%p \n", accu)

	accu2 := Accumulate(20)

	println(accu2())
	println(accu2())

	fmt.Printf("%p \n", accu2)

}

var str = "hello world"

var foo = func() {
	str = "hello func"
	println(str)
}

func testClosure1() {
	println(str)
	foo()
	println(str)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func testInterface2() {
	var invoker Invoker

	invoker = FuncCaller(func(v interface{}) {
		println(v)
	})

	invoker.Call("Hello")
}

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	// fmt.Println("from struct", p)
	println(p)
}

func testInterface1() {
	var invoker Invoker
	s := new(Struct)
	invoker = s
	invoker.Call("Hello")
}

// func println(p interface{}) {
// 	fmt.Println(p)
// }

var skillParam = flag.String("skill", "", "skill to perform")

func testFlag() {

	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chickren fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

}

func testFunc2() {
	var p func()

	p = testFunc1

	p()

	//匿名函数
	func(data int) {
		fmt.Println("hello ", data)
	}(100)

	//匿名函数赋值
	f := func(data int) {
		fmt.Println("hello ", data)
	}

	f(200)

	var x []int
	x[0] = 0
	x[1] = 1

	visit(x, f)

}

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func testFunc1() {
	var in Data
	var innerData InnerData

	innerData.a = 20

	in.instance = innerData

	var x = []int{1, 2, 3, 4}

	in.complex = x

	in.ptr = &in.instance

	fmt.Println(in)

	y := passByValue(in)

	fmt.Println(y)

	in2 := Data{
		complex: []int{3, 4, 5},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	fmt.Println(in2)
}

func passByValue(inFunc Data) Data {
	fmt.Printf("inFunc value:%+v \n", inFunc)
	fmt.Printf("inFunc ptr: %p\n", inFunc)

	return inFunc
}

type InnerData struct {
	a int
}

type Data struct {
	complex  []int
	instance InnerData

	ptr *InnerData
}

func testFunc() {
	b := func1("testFunc", 10)
	fmt.Println(b)

	c, d := func2()

	fmt.Println(c, d)
}

func func1(a string, b int) int {
	print(a)
	return b
}

func func2() (int, int) {
	return 10, 20
}
func testFor() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d * %d = %d   ", x, y, x*y)
		}

		fmt.Println()
	}
}

func testChannel() {
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3

	}()

	for v := range c {
		fmt.Println(v)
	}
}

func testList() {
	l := list.New()

	l.PushBack("canon")
	l.PushFront(67)

	element := l.PushBack("first")

	l.InsertAfter("high", element)
	l.InsertBefore("noon", element)

	fmt.Println(l)
	l.Remove(element)
	fmt.Println(l)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}

func testSyncMap() {

	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	var m = 200
	go func() {
		for {
			// k, v := scene.Load("greece")
			m = m + 1
			scene.Store("greece", m)
		}
	}()

	fmt.Println(scene)

	go func() {
		for {
			k, v := scene.Load("greece")
			fmt.Println(k, v)
		}
	}()

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterator:", k, v)
		return true
	})

	for {

	}
}

func testMap() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["route2"] = 76
	scene["route8"] = 86
	scene["route5"] = 56

	fmt.Println(scene)

	fmt.Println(scene["2"])

	var sceneList []string

	for k := range scene {
		sceneList = append(sceneList, k)
	}

	sort.Strings(sceneList)

	fmt.Println(sceneList)

	delete(scene, "route5")

	fmt.Println(scene)

}

func testSlice() {
	var team [3]string
	team[0] = "cow"
	team[1] = "horse"
	team[2] = "mouse"

	fmt.Println(team)

	var team2 = [3]string{team[0], team[2], team[1]}
	fmt.Println(team2)
	var team3 = [...]string{team[0], team[2], team[1], team2[1]}
	fmt.Println(team3)

	for i, v := range team3 {
		fmt.Println(i, v)
	}

	var team4 = team3[2:4]

	fmt.Println(team4)

	var team5 = team3[:]

	fmt.Println(team5)

	var team6 = team3[:3]

	fmt.Println(team6)

	var team7 = team3[2:]

	fmt.Println(team7)

	var name []string

	name = team7

	fmt.Println(name)

	a := make([]string, 3)

	fmt.Println(a)

	a = append(a, "test")
	a = append(a, "test2")

	fmt.Println(a)
}

func testAddMethod() {

	var a MyDuration

	a.EasySet("test")
}

type MyDuration time.Duration

func (m MyDuration) EasySet(a string) {
	print(a)
}

func testStruct() {
	var t Brand
	t.show("testStruct")

	var t2 fakeBrand
	t2.show("fakeBrand")

	var a Vehicle
	ta := reflect.TypeOf(a)

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)

		fmt.Printf("Field Name:%v ,Field Type:%v \n", f.Name, f.Type.Name())
	}
}

type fakeBrand = Brand
type Brand struct {
}

func (t Brand) show(s string) {
	print(s)
}

type Vehicle struct {
	Brand
	fakeBrand
}

func testType() {
	type NewInt int

	// type Int Alias = int

	type IntAlias = int

	var a NewInt = 20
	var b IntAlias = 30

	fmt.Printf("%d %d %T %T ", a, b, a, b)
}

func print(message string) {
	fmt.Println(message)
}

func testStringLen() {
	var a = "忍者"

	fmt.Println(len(a))

	var b = utf8.RuneCountInString(a)

	fmt.Println(b)

	for i, v := range a {
		fmt.Printf(" %d %c %d ", i, v, v)
	}

}

// func println(message T) {
// 	fmt.Println(message)
// }
func test() {
	c := "test4"
	print("test")
	print(b)
	print(c)

	a := 10
	b := 20

	a, b = swap(a, b)

	print(strconv.Itoa(a))

	createImage()

	c = `
	one line
	two line
	`

	print(c)

	var cat int = 1
	var str string = "banana"

	fmt.Println("%p %p", &cat, &str)

	var aa = 10
	var bb = 20

	swap2(&aa, &bb)

	fmt.Println("%d, %d", aa, bb)

	var aa2 = 20
	var bb2 = &aa2
	var cc2 = *bb2
	fmt.Println("%d %d %d", aa2, bb2, cc2)
	cc2 = 30
	fmt.Println("%d %d %d", aa2, bb2, cc2)

	*bb2 = 40
	fmt.Println("%d %d %d", aa2, bb2, cc2)

	printFlag()
}

var mode = flag.String("mode", "", "process mode")

func printFlag() {
	flag.Parse()

	fmt.Println(*mode)
}

func swap2(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
func createImage() {
	size := 300

	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / float64(size)

		y := float64(size)/2 - math.Sin(s)*float64(size)/2
		pic.SetGray(x, int(y), color.Gray{0})

	}

	file, err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file, pic)

	file.Close()

}

func swap(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

var b string = "test3"

// func f1() {
// 	var a int
// 	var b string
// 	var c []float32
// 	var d func() bool

// 	var e struct {
// 		x int
// 	}

// 	print(b)
// }
