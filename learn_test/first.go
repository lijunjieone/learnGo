package main

import (
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
	testSyncMap()

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
