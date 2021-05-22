package opcodes

import (
	"fmt"
	"github.com/skycoin/cx/cx/ast"
	"github.com/skycoin/cx/cx/constants"
	"github.com/skycoin/cx/cx/types"
)

var assertSuccess = true

// AssertFailed ...
func AssertFailed() bool {
	return !assertSuccess
}

//TODO: Rework
func assert(inputs []ast.CXValue, outputs []ast.CXValue) (same bool) {
	var byts1, byts2 []byte

	types.FMTDEBUG(fmt.Sprintf("BYTES1 %d\n", inputs[0].Offset))
	types.FMTDEBUG(fmt.Sprintf("BYTES2 %d\n", inputs[1].Offset))

	if inputs[0].Arg.Type == constants.TYPE_STR {
		//s1 := inputs[0].Get_str()
		//s2 := inputs[1].Get_str()

		//fmt.Printf("TESTING '%s' VS '%s'\n", s1, s2)
		byts1 = []byte(inputs[0].Get_str())
		byts2 = []byte(inputs[1].Get_str())
		//fmt.Printf("LEN %d, %d\n", len(byts1), len(byts2))
	} else {
	//	fmt.Printf("V0 %v AT OFFSET %v\n", inputs[0].Get_bytes(), inputs[0].Offset)
	//	fmt.Printf("V1 %v AT OFFSET %v\n", inputs[1].Get_bytes(), inputs[1].Offset)
		byts1 = inputs[0].Get_bytes()
		byts2 = inputs[1].Get_bytes()
	}

	same = true

	if len(byts1) != len(byts2) {
		same = false
		fmt.Println("byts1", byts1)
		fmt.Println("byts2", byts2)
	}

	if same {
		for i, byt := range byts1 {
			if byt != byts2[i] {
				same = false
				fmt.Println("byts1", byts1)
				fmt.Println("byts2", byts2)
				break
			}
		}
	}

	message := inputs[2].Get_str()

	if !same {
	    call := ast.PROGRAM.GetCurrentCall()
    	expr := call.Operator.Expressions[call.Line]
		if message != "" {
			fmt.Printf("%s: %d: result was not equal to the expected value; %s\n", expr.FileName, expr.FileLine, message)
		} else {
			fmt.Printf("%s: %d: result was not equal to the expected value\n", expr.FileName, expr.FileLine)
		}
	}

	assertSuccess = assertSuccess && same
	return same
}

func opAssertValue(inputs []ast.CXValue, outputs []ast.CXValue) {
	same := assert(inputs, outputs)
    outputs[0].Set_bool(same)
}

func opTest(inputs []ast.CXValue, outputs []ast.CXValue) {
	assert(inputs, outputs)
}

func opPanic(inputs []ast.CXValue, outputs []ast.CXValue) {
	if !assert(inputs, outputs) {
		panic(constants.CX_ASSERT)
	}
}

// panicIf/panicIfNot implementation
func panicIf(inputs []ast.CXValue, outputs []ast.CXValue, condition bool) {
    str := inputs[1].Get_str()
	if inputs[0].Get_bool() == condition {
	    call := ast.PROGRAM.GetCurrentCall()
    	expr := call.Operator.Expressions[call.Line]
		fmt.Printf("%s : %d, %s\n", expr.FileName, expr.FileLine, str)
		panic(constants.CX_ASSERT)
	}
}

// panic with CX_ASSERT exit code if condition is true
func opPanicIf(inputs []ast.CXValue, outputs []ast.CXValue) {
	panicIf(inputs, outputs, true)
}

// panic with CX_ASSERT exit code if condition is false
func opPanicIfNot(inputs []ast.CXValue, outputs []ast.CXValue) {
	panicIf(inputs, outputs, false)
}

func opStrError(inputs []ast.CXValue, outputs []ast.CXValue) {
    outputs[0].Set_str(ast.ErrorString(int(inputs[0].Get_i32())))
}
