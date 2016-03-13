package datastructure

import "fmt"
import "errors"

var (
	Gv1 = 1
	Gv2 = 2
	Gv3 float32 = 3.4
	activated = true
)

const(
	GPi = 3.4
	enumX, enumYSameToX = iota, iota //1, second line
	enumZ = iota // 2

)

const(
	enumXX = iota //0
	enumYY = iota //1
	enumZZ        //2

)

func DefineVars(){
	var v1, v2, v3 int
	v1 = 1 + v3 // 1
	v2 = 2 + v1 //2
	v3 = 3 + v2 //6

	fmt.Printf("%d, %d, %d", v1, v2, v3)

	v4, v5 := 4, 5
	fmt.Printf("%d, %d", v4, v5)

	enabled := false
	fmt.Printf("%d", enabled)

	var i8 int8
	var i16 int16 = int16(i8) + 16  //force cast
	var i32 int32 = rune(i16) + 16  //rune is int32
	var i64 int64 = int64(i32) + 32
	fmt.Printf("%d, %d, %d", i8, i16, i64)

	var u8 uint8 = byte(0) //byte is uint8
	var u16 uint16 = uint16(u8) + 16
	var u32 uint32 = uint32(u16) + 16
	var u64 uint64 = uint64(u32) + 32
	fmt.Printf("%d, %d, %d", u8, u16, u64)

	str := "let me go"
	fmt.Printf("%s", str)

	fmt.Printf("%d, %d, %d, %d, %d, %d", enumX, enumYSameToX, enumZ, enumXX, enumYY, enumZZ)
}


func DemoErr(){
	err := errors.New("oh, dam")
	if err != nil {
		fmt.Println(err)
	}
}

func DemoSwap(a int, b int) (int, int){
	return b,a
}

func DemoArray(){
	var arr[10] int
	arr[0] = 0

	arr3 := [3]int{1,2,3}
	arr4 := [...]int{1,2,3,4} //should be [...]

	fmt.Println("len arr3:", cap(arr3))
	fmt.Println(cap(arr4))

	secondDArr := [2][3]int{{1,1,1},{2,2,2}}
	for i:=0;i<2;i++ {
		for j:=0;j<3;j++{
			fmt.Printf("%d", secondDArr[i][j])
		}
	}

	for i,v := range arr3 {
		fmt.Printf("%d->%d", i, v)
	}

	fmt.Println("len arr4:", len(secondDArr))
}

func DemoSlice(){
	slice := []byte{'a','b', 'c', 'd'}
	var a []byte = slice[1:2]
	var b []byte = slice[:2]
	var c []byte = slice[2:]
	var d []byte = slice[1:2:2]

	fmt.Println("len slice:", len(slice))
	fmt.Println("cap slice:", cap(slice))

	fmt.Println("len slice a:", len(a)) //index: 1
	fmt.Println("cap slice a:", cap(a)) //index: 1,2,3

	fmt.Println("len slice b:", len(b)) // index: 0,1
	fmt.Println("cap slice b:", cap(b))// index: 0, 1, 2,3

	fmt.Println("len slice c:", len(c)) // index: 2,3
	fmt.Println("cap slice c:", cap(c)) // index: 2,3

	fmt.Println("len slice d:", len(d)) // index: 1
	fmt.Println("cap slice d:", cap(d)) // index: 1

	e := append(d, 'e')
	fmt.Println("len slice d:", len(d)) // index: 1
	fmt.Println("cap slice d:", cap(d)) // index: 1

	fmt.Println("len slice e:", len(e)) // index: 2
	fmt.Println("cap slice e:", cap(e)) // capacity will be increased more..

	f := make([]string, 3, 5)
	fmt.Println("len slice f:", len(f)) // 3
	fmt.Println("cap slice f:", cap(f)) // 5

	g := make([]string, 3)
	v:= copy(g, f)
	fmt.Println("copy len:", v) // index: 3

	g = append(g, "a", "b", "c")
	fmt.Println("slice g:", g)
	fmt.Println("len slice g:", len(g)) // 6
	fmt.Println("cap slice g:", cap(g)) // 6

}

func DemoMap(){
	firstDict := map[string]int {"x":1, "y":2, "z":3}
	fmt.Println("firstDict:", firstDict)

	if _,ok := firstDict["t"]; ok{
		fmt.Println("include t.")
	}else{
		fmt.Println("not include t.")
	}

	delete(firstDict, "z")
	firstDict["z"]= 4
	fmt.Println("firstDict changed:", firstDict)

	secondDict := make(map[string]int)
	fmt.Println("secondDict :", secondDict)
	secondDict["x"] = 3
	for k, v := range secondDict {
		fmt.Printf("%s -> %d\n", k, v)
	}
}