package main

import (
	"bytes"
	"container/list"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
	"unicode/utf8"

	"./p"
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

	// testInterface3()

	// testTypeTansfer()
	// testTypeTransfer2()

	// testState1()

	// testPackage1()

	// testChannel1()
	// testChannel2()
	// testChannel5()
	// testChannel4()
	// testChannel3()

	// testChannel8()

	// testChannel9()
	// testChannel7()

	// testChannel10()

	// testChannel11()

	// testReflect1()
	// testReflect2()
	// testReflect3()
	// testReflect4()
	// testReflect6()
	// testReflect5()
	// testReflect7()
	// testReflect8()

	// testReflect9()
	testReflect10()

}

func testReflect10() {
	type Skill struct {
		Name  string
		Level int
	}

	type Actor struct {
		Name string
		Age  int

		Skills []Skill
	}

	a := Actor{
		Name: "Go Lang",
		Age:  2,

		Skills: []Skill{
			{
				Name: "Web", Level: 1,
			},
			{
				Name: "Backend", Level: 2,
			},
			{
				Name: "Forend", Level: 3,
			},
		},
	}

	if result, err := MarshalJson(a); err == nil {
		println(result)
	} else {
		println(err)
	}

}

func MarshalJson(v interface{}) (string, error) {

	var b bytes.Buffer

	if err := writeAny(&b, reflect.ValueOf(v)); err == nil {
		return b.String(), nil
	} else {
		return "", nil
	}
}

func writeAny(buff *bytes.Buffer, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buff.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buff, value)
	case reflect.Struct:
		return writeStruct(buff, value)
	default:
		return errors.New("unsupport kind:" + value.Kind().String())
	}

	return nil
}

func writeSlice(buff *bytes.Buffer, value reflect.Value) error {

	buff.WriteString("[")

	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)

		writeAny(buff, sliceValue)

		if s < value.Len()-1 {
			buff.WriteString(",")
		}
	}

	buff.WriteString("]")

	return nil
}

func writeStruct(buff *bytes.Buffer, value reflect.Value) error {

	valueType := value.Type()

	buff.WriteString("{")

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)

		fieldType := valueType.Field(i)

		buff.WriteString("\"")
		buff.WriteString(fieldType.Name)
		buff.WriteString("\":")

		writeAny(buff, fieldValue)

		if i < value.NumField()-1 {
			buff.WriteString(",")
		}

	}

	buff.WriteString("}")

	return nil

}
func add(a, b int) int {
	return a + b
}

func testReflect9() {
	funcValue := reflect.ValueOf(add)

	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	retList := funcValue.Call(paramList)

	println(retList[0].Int())

}

func testReflect8() {
	type dog struct {
		LegCount int
	}

	d := dog{4}

	println(d.LegCount)

	valueOfDog := reflect.ValueOf(&d)

	valueOfDog = valueOfDog.Elem()

	vLegCount := valueOfDog.FieldByName("LegCount")

	vLegCount.SetInt(3)

	println(vLegCount.Int())

	println(d.LegCount)

	typeOfB := reflect.TypeOf(d)
	aIns := reflect.New(typeOfB)
	aIns = aIns.Elem()
	aInsCount := aIns.FieldByName("LegCount")

	println(d.LegCount)

	fmt.Println("type", aIns.Type(), aIns.Kind())
	aInsCount.SetInt(4)

	println(d.LegCount)

	println(aInsCount.Int())
}

func testReflect7() {
	var a int = 1024

	println(a)

	valueOfA := reflect.ValueOf(&a)

	valueOfA = valueOfA.Elem()
	valueOfA.SetInt(20)

	println(valueOfA.Int())
	println(a)
}

type dummy struct {
	a int
	b string

	float32
	bool

	next *dummy
}

func testReflect6() {

	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})

	fmt.Println("NumField", d.NumField())

	floatField := d.Field(2)

	fmt.Println("Field", floatField.Type())

	fmt.Println("Field By Name(\"b\").Type", d.FieldByName("b").Type())

	fmt.Println("Field By Index([]int{4,0}).Type()", d.FieldByIndex([]int{4, 0}).Type())

}
func testReflect5() {
	var a int = 1024

	valueOfA := reflect.ValueOf(a)

	var getA int = valueOfA.Interface().(int)

	var getA2 int = int(valueOfA.Int())

	println(getA)
	println(getA2)
}
func testReflect4() {

	type cat struct {
		Name string

		Type int `json:"type" id:"100"`
	}

	ins := cat{Name: "goLang", Type: 1}

	typeOfCat := reflect.TypeOf(ins)

	for i := 0; i < typeOfCat.NumField(); i++ {

		fieldType := typeOfCat.Field(i)

		fmt.Printf("name:%v tag:'%v' \n", fieldType.Name, fieldType.Tag)

	}

	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

type Enum int

const (
	Zero Enum = 0
)

func testReflect3() {

	type cat struct {
	}

	ins := &cat{}

	typeOfCat := reflect.TypeOf(ins)

	println(typeOfCat.Name())
	println(typeOfCat.Kind())

	typeOfA := typeOfCat.Elem()

	println(typeOfA.Name())
	println(typeOfA.Kind())

}

func testReflect2() {

	type cat struct {
	}

	typeOfCat := reflect.TypeOf(cat{})

	println(typeOfCat.Name())
	println(typeOfCat.Kind())

	typeOfA := reflect.TypeOf(Zero)

	println(typeOfA.Name())
	println(typeOfA.Kind())

}

func testReflect1() {

	var a int

	var b float32
	typeOfA := reflect.TypeOf(a)

	typeOfB := reflect.TypeOf(b)

	println(typeOfA.Name())
	println(typeOfA.Kind())

	println("B")

	println(typeOfB.Name())
	println(typeOfB.Kind())
}
func testChannel11() {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.github.com",
		"https://www.qiniu.com",
		"https://www.golangtc.com",
	}

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			content, err := http.Get(url)
			println(content)

			fmt.Println(url, err)
		}(url)
	}

	wg.Wait()

	println("over")
}

var seq int64

var count int
var countGuard sync.Mutex

func GetCount() int {
	countGuard.Lock()

	defer countGuard.Unlock()

	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()

}

func testChannel10() {
	SetCount(10)
	println(GetCount())
}
func GenId() int64 {
	return atomic.AddInt64(&seq, 1)
	// println(seq)

	// return seq
}

func testChannel9() {

	for i := 0; i < 20; i++ {
		go GenId()
	}

	fmt.Println("result:", GenId())

	time.Sleep(time.Second * 10)
}
func testChannel7() {
	ticker := time.NewTicker(time.Millisecond * 500)

	stopper := time.NewTimer(time.Second * 2)

	var i int

	for {

		select {

		case <-stopper.C:
			println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}
	}

StopHere:
	println("done")
}

func testChannel5() {

	ch := make(chan int, 3)

	println(len(ch))

	ch <- 1
	ch <- 2
	ch <- 3

	data := <-ch
	println(data)
	data2 := <-ch
	println(data2)
	ch <- 4

	println(len(ch))
}

func RPCClient(ch chan string, req string) (string, error) {

	ch <- req

	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {

		data := <-ch

		fmt.Println("server received:", data)

		time.Sleep(time.Second * 2)
		ch <- "roger"
	}
}

func testChannel8() {

	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")

	if err != nil {
		println(err)
	} else {

		fmt.Println("client received", recv)
	}

}

func testChannel4() {
	c := make(chan int)

	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i

		time.Sleep(time.Second)
	}

	c <- 0

	<-c

}

func printer(c chan int) {

	for {
		data := <-c

		if data == 0 {
			break
		}

		println(data)
	}

	c <- 0
}
func testChannel3() {

	ch := make(chan int)

	go func() {
		for i := 3; i >= 0; i-- {
			ch <- i

			time.Sleep(time.Second)
		}

	}()

	for data := range ch {
		println(data)

		if data == 0 {
			break
		}
	}

	println("all done")
}

func testChannel2() {

	ch := make(chan int)

	go func() {
		println("start go running")
		ch <- 0

		println("exit go running")
	}()

	println("wait go")

	println("all done")
}

func testChannel1() {
	go running()

	var input string
	fmt.Scanln(&input)
}

func running() {
	var times int

	for {
		times++
		print("times:")
		println(times)
		time.Sleep(time.Second)
	}

}

func init() {
	println("first.go init")
}

func testPackage1() {
	c := p.AddSum(2, 3)
	println(c)
}

type State interface {
	Name() string
	EnableSameTransit() bool
	OnBegin()
	OnEnd()
	CanTransitTo(name string) bool
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}

	return reflect.TypeOf(s).Elem().Name()
}

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}

func (s *StateInfo) setName(name string) {
	s.name = name
}

func (s *StateInfo) EnableSameTransit() bool {
	return false
}

func (s *StateInfo) OnBegin() {

}

func (s *StateInfo) OnEnd() {

}

func (s *StateInfo) CanTransitTo(name string) bool {
	return true
}

type StateManager struct {
	stateByName map[string]State

	OnChange func(fro, to State)

	curr State
}

func (sm *StateManager) Add(s State) {
	name := StateName(s)

	s.(interface {
		setName(name string)
	}).setName(name)

	if sm.Get(name) != nil {
		panic("duplicate state:" + name)
	}

	sm.stateByName[name] = s

}

func (sm *StateManager) Get(name string) State {

	if v, ok := sm.stateByName[name]; ok {
		return v
	}

	return nil
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateByName: make(map[string]State),
	}
}

var ErrStateNotFound = errors.New("state not found")
var ErrForbidSameStateTransit = errors.New("forbid same state transit")

var ErrCannotTransitToState = errors.New("cannot transit to state")

func (sm *StateManager) CurrState() State {
	return sm.curr
}

func (sm *StateManager) CanCurrTransitTo(name string) bool {

	if sm.curr == nil {
		return true
	}

	if sm.curr.Name() == name && !sm.curr.EnableSameTransit() {
		return false
	}

	return sm.curr.CanTransitTo(name)
}

func (sm *StateManager) Transit(name string) error {

	next := sm.Get(name)

	if next == nil {
		return ErrStateNotFound
	}

	pre := sm.curr

	if sm.curr != nil {
		if !sm.curr.CanTransitTo(name) {
			return ErrCannotTransitToState
		}

		sm.curr.OnEnd()
	}

	sm.curr = next

	sm.curr.OnBegin()

	if sm.OnChange != nil {
		sm.OnChange(pre, sm.curr)
	}

	return nil
}

type IdleState struct {
	StateInfo
}

func (i *IdleState) OnBegin() {
	println("Idle State Begin")
}

func (i *IdleState) OnEnd() {
	println("Idle State End")
}

type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	println("Move State Begin")
}

func (m *MoveState) OnEnd() {
	println("Move State End")
}

func (m *MoveState) EnableSameTransit() bool {
	return true
}

type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	println("Jump State Begin")
}

func (j *JumpState) OnEnd() {
	println("Jump State End")
}

func (j *JumpState) CanTransitTo(name string) bool {
	return name != "Move State"
}

func testState1() {

	sm := NewStateManager()
	sm.OnChange = func(from, to State) {
		// print("from")
		// print(StateName(from))
		// print("------>")
		// print(StateName(to))
		// println("")
		fmt.Printf("%s ------> %s", StateName(from), StateName(to))
		println("")
	}

	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")

}

func transitAndReport(sm *StateManager, target string) {

	if err := sm.Transit(target); err != nil {
		fmt.Printf("fail! %s -----> %s, %s \n\n", sm.CurrState().Name(), target, err.Error())
	}
}

func testTypeTransfer2() {

	var a int = 1

	var i interface{} = a

	var b int = i.(int)

	var c float32 = float32(b)

	println(b)

	println(c)
}

type Flyer interface {
	Fly()
}

type Walker interface {
	Walker()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	println("bird: walk")
}

type pig struct {
}

func (p *pig) Walker() {
	println("pig:walk")
}

func testTypeTansfer() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {
		f, isFlayer := obj.(Flyer)
		w, isWalker := obj.(Walker)

		print(name)
		println(isFlayer)
		println(isWalker)

		if isFlayer {
			f.Fly()
		}

		if isWalker {
			w.Walker()
		}

	}

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
