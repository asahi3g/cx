package opcodes

import (
	"github.com/skycoin/cx/cx/ast"
	"github.com/skycoin/cx/cx/constants"
	"github.com/skycoin/cx/cx/types"
)

//TODO: Rename opSliceLen
//TODO: Rework
func opLen(inputs []ast.CXValue, outputs []ast.CXValue) {
	elt := ast.GetAssignmentElement(inputs[0].Arg)

	var sliceLen types.Pointer
	if elt.IsSlice || elt.Type == constants.TYPE_AFF { //TODO: FIX
		sliceOffset := types.Read_ptr(ast.PROGRAM.Memory, inputs[0].Offset)
		if sliceOffset > 0 {
			sliceLen = ast.GetSliceLen(sliceOffset) // TODO: PTR remove hardcode 4
		} else if sliceOffset < 0 {
			panic(constants.CX_RUNTIME_ERROR)
		}

		// TODO: Had to add elt.Lengths to avoid doing this for arrays, but not entirely sure why
	} else if elt.Type == constants.TYPE_STR && elt.Lengths == nil {
		sliceLen = types.Read_str_size(ast.PROGRAM.Memory, inputs[0].Offset)
	} else {
		sliceLen = elt.Lengths[len(elt.Indexes)]
	}

	//inputs[0].Used = int8(inputs[0].Type) // TODO: Remove hacked type check
	outputs[0].Set_i32(types.Cast_ptr_to_i32(sliceLen)) // TODO: PTR remove hardcode i32
}

//TODO: Rename OpSliceAppend
//TODO: Rework
func opSliceAppend(inputs []ast.CXValue, outputs []ast.CXValue) {
	inp0, inp1, out0 := inputs[0].Arg, inputs[1].Arg, outputs[0].Arg

	eltInp0 := ast.GetAssignmentElement(inp0)
	eltOut0 := ast.GetAssignmentElement(out0)
	if inp0.Type != inp1.Type || inp0.Type != out0.Type || !eltInp0.IsSlice || !eltOut0.IsSlice {
		panic(constants.CX_RUNTIME_INVALID_ARGUMENT)
	}

	var inputSliceLen types.Pointer
	inputSliceOffset := types.Read_ptr(ast.PROGRAM.Memory, inputs[0].Offset)
	if inputSliceOffset != 0 {
		inputSliceLen = ast.GetSliceLen(inputSliceOffset)
	}

	// Preparing slice in case more memory is needed for the new element.
	outputSliceOffset := ast.SliceAppendResize(inputs[0].Frame, out0, inp0, inp1.Size)

	// We need to update the address of the output and input, as the final offsets
	// could be on the heap and they could have been moved by the GC.

	if inp1.Type == constants.TYPE_STR || inp1.Type == constants.TYPE_AFF {
		var obj [types.TYPE_POINTER_SIZE]byte
		types.Write_ptr(obj[:], 0, types.Read_ptr(ast.PROGRAM.Memory, inputs[1].Offset))
		ast.SliceAppendWrite(outputSliceOffset, obj[:], inputSliceLen)
	} else {
		obj := inputs[1].Get_bytes()
		ast.SliceAppendWrite(outputSliceOffset, obj, inputSliceLen)
	}

	//inputs[0].Used = int8(inputs[0].Type) // TODO: Remove hacked type check
	//inputs[1].Used = int8(inputs[1].Type) // TODO: Remove hacked type check
	outputs[0].Set_ptr(outputSliceOffset)
}

//TODO: Rename opSliceResize
//TODO: Rework
func opResize(inputs []ast.CXValue, outputs []ast.CXValue) {
	inp0, out0 := inputs[0].Arg, outputs[0].Arg
	fp := inputs[0].Frame

	if inp0.Type != out0.Type || !ast.GetAssignmentElement(inp0).IsSlice || !ast.GetAssignmentElement(out0).IsSlice {
		panic(constants.CX_RUNTIME_INVALID_ARGUMENT)
	}

	outputSliceOffset := ast.SliceResize(fp, out0, inp0, types.Cast_i32_to_ptr(inputs[1].Get_i32()), ast.GetAssignmentElement(inp0).TotalSize)

	//inputs[0].Used = int8(inputs[0].Type) // TODO: Remove hacked type check
	outputs[0].Set_ptr(outputSliceOffset)
}

//TODO: Rename opSliceInsertElement
//TODO: Rework
func opInsert(inputs []ast.CXValue, outputs []ast.CXValue) {
	inp0, inp2, out0 := inputs[0].Arg, inputs[2].Arg, outputs[0].Arg
	fp := inputs[0].Frame

	if inp0.Type != inp2.Type || inp0.Type != out0.Type || !ast.GetAssignmentElement(inp0).IsSlice || !ast.GetAssignmentElement(out0).IsSlice {
		panic(constants.CX_RUNTIME_INVALID_ARGUMENT)
	}

	index := types.Cast_i32_to_ptr(inputs[1].Get_i32())
	var outputSliceOffset types.Pointer
	if inp2.Type == constants.TYPE_STR || inp2.Type == constants.TYPE_AFF {
		var obj [types.TYPE_POINTER_SIZE]byte
		types.Write_ptr(obj[:], 0, types.Read_ptr(ast.PROGRAM.Memory, inputs[2].Offset))
		outputSliceOffset = ast.SliceInsert(fp, out0, inp0, index, obj[:])
	} else {
		obj := inputs[2].Get_bytes()
		outputSliceOffset = ast.SliceInsert(fp, out0, inp0, index, obj)
	}

	//inputs[0].Used = int8(inputs[0].Type) // TODO: Remove hacked type check
	//inputs[2].Used = int8(inputs[2].Type) // TODO: Remove hacked type check
	outputs[0].Set_ptr(outputSliceOffset)
}

//TODO: Rename opSliceRemoveElement
//TODO: Rework
func opRemove(inputs []ast.CXValue, outputs []ast.CXValue) {
	inp0, out0 := inputs[0].Arg, outputs[0].Arg
	fp := inputs[0].Frame

	if inp0.Type != out0.Type || !ast.GetAssignmentElement(inp0).IsSlice || !ast.GetAssignmentElement(out0).IsSlice {
		panic(constants.CX_RUNTIME_INVALID_ARGUMENT)
	}

	outputSliceOffset := ast.SliceRemove(fp, out0, inp0, types.Cast_i32_to_ptr(inputs[1].Get_i32()), ast.GetAssignmentElement(inp0).TotalSize)

	//inputs[0].Used = int8(inputs[0].Type) // TODO: Remove hacked type check
	outputs[0].Set_ptr(outputSliceOffset)
}

//TODO: Rename opSliceCopy
func opCopy(inputs []ast.CXValue, outputs []ast.CXValue) {
	dstInput := inputs[0].Arg
	srcInput := inputs[1].Arg
	fp := inputs[0].Frame

	dstOffset := ast.GetSliceOffset(fp, dstInput)
	srcOffset := ast.GetSliceOffset(fp, srcInput)

	dstElem := ast.GetAssignmentElement(dstInput)
	srcElem := ast.GetAssignmentElement(srcInput)

	if dstInput.Type != srcInput.Type || !dstElem.IsSlice || !srcElem.IsSlice || dstElem.TotalSize != srcElem.TotalSize {
		panic(constants.CX_RUNTIME_INVALID_ARGUMENT)
	}

	var count types.Pointer
	if dstInput.Type == srcInput.Type && dstOffset >= 0 && srcOffset >= 0 {
		count = types.Cast_int_to_ptr(copy(ast.GetSliceData(dstOffset, dstElem.TotalSize), ast.GetSliceData(srcOffset, srcElem.TotalSize)))
		if count%dstElem.TotalSize != 0 {
			panic(constants.CX_RUNTIME_ERROR)
		}
	} else {
		panic(constants.CX_RUNTIME_INVALID_ARGUMENT)
	}

	//inputs[0].Used = int8(inputs[0].Type) // TODO: Remove hacked type check
	//inputs[1].Used = int8(inputs[1].Type) // TODO: Remove hacked type check
	outputs[0].Set_ptr(count / dstElem.TotalSize)
}
