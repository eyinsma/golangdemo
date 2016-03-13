package main

import "fmt"
import ds "github.com/golangdemo/datastructure"
import fc "github.com/golangdemo/flowcontrol"

func main() {
	//ds
	ds.DefineVars()
	fmt.Printf("%d, %d", ds.Gv1, ds.Gv2)
	ds.DemoErr()
	ds.DemoArray()
	ds.DemoSlice()
	ds.DemoMap()

	ds.DemoStruct()
	ds.DemoMethod()

	ds.DemoInterface()

	//fc
	fc.DemoIf()
	fc.DemoGoto()
	fc.DemoFor()
	fc.DemoSwitch()
	fc.DemoVarPara(1,2)
	fc.DemoDefer()
	fc.DemoFuncPara()
}
