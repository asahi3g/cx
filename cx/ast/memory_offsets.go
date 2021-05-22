package ast

import (
	"github.com/skycoin/cx/cx/constants"
    "github.com/skycoin/cx/cx/types"
	"fmt"
)

// GetDerefSize ...
func GetDerefSize(arg *CXArgument) types.Pointer {
	if arg.CustomType != nil {
		return arg.CustomType.Size //TODO: WTF is a custom type?
	}
	return arg.Size
}

func PrintArg(name string, arg *CXArgument) {
	fmt.Sprintf("%s %s, SIZE %d, SLICE %v, TYPE %s\n", name, arg.ArgDetails.Name, arg.Size, arg.IsSlice, constants.TypeNames[arg.Type])
}

//TODO: Delete this eventually
var BEFORE_COUNT int
func GetFinalOffset(fp types.Pointer, arg *CXArgument) types.Pointer {
	finalOffset := arg.Offset

	//Todo: find way to eliminate this check
	types.FMTDEBUG(fmt.Sprintf("NAME %s, FINAL_OFFSET %d, PROGRAM.Stack %d, ARG_OFF %d\n", arg.ArgDetails.Name, finalOffset, PROGRAM.StackSize, arg.Offset))
	if finalOffset < PROGRAM.StackSize {
		// Then it's in the stack, not in data or heap and we need to consider the frame pointer.
		types.FMTDEBUG(fmt.Sprintf("ADDING FPFP %d + %d = %d\n", finalOffset, fp, finalOffset + fp))
		finalOffset += fp
	}

	// elt = arg
	//TODO: Eliminate all op codes with more than one return type
	//TODO: Eliminate this loop
	//Q: How can CalculateDereferences change offset?
	//Why is finalOffset fed in as a pointer?

	types.FMTDEBUG(fmt.Sprintf("BEFORECALCULATE %d\n",finalOffset))
/*	if finalOffset == 0 {
		BEFORE_COUNT++
		if BEFORE_COUNT == 2 {
			panic("FUCK")
		}
	}*/
	finalOffset = CalculateDereferences(arg, finalOffset, fp)
	types.FMTDEBUG(fmt.Sprintf("FIELD COUNT %d, FINAL %v\n", len(arg.Fields), finalOffset))
	for _, fld := range arg.Fields {
		// elt = fld
		types.FMTDEBUG(fmt.Sprintf("FINAL %d, FIELD %d\n", finalOffset, fld.Offset))
		finalOffset += fld.Offset
		finalOffset = CalculateDereferences(fld, finalOffset, fp)
		types.FMTDEBUG(fmt.Sprintf("FINAL %v\n", finalOffset))
	}

	return finalOffset
}

func CalculateDereferences(arg *CXArgument, finalOffset types.Pointer, fp types.Pointer) types.Pointer {
	//fmt.Printf("CALCULATE_DEREF\n")
	var isPointer bool

	var baseOffset types.Pointer
	var sizeofElement types.Pointer
	types.FMTDEBUG(fmt.Sprintf("CALCULATE_DEREFERENCES %d, %d, %s\n", finalOffset, fp, arg.ArgDetails.Name))
	types.FMTDEBUG(fmt.Sprintf("DEREF_COUNT %d\n", len(arg.DereferenceOperations)))
	idxCounter := 0
	for _, op := range arg.DereferenceOperations {
		switch op {
		case constants.DEREF_SLICE: //TODO: Move to CalculateDereference_slice
			types.FMTDEBUG(fmt.Sprintf("DEREF_SLICE\n"))
			if len(arg.Indexes) == 0 {
				continue
			}

			isPointer = false
			finalOffset = types.Read_ptr(PROGRAM.Memory, finalOffset)
			baseOffset = finalOffset

			finalOffset += types.OBJECT_HEADER_SIZE
			finalOffset += constants.SLICE_HEADER_SIZE

			//TODO: delete
			sizeToUse := GetDerefSize(arg) //TODO: is always arg.Size unless arg.CustomType != nil
			PrintArg("DEREF_SLICE: ", arg)
			finalOffset += types.Read_ptr(PROGRAM.Memory, GetFinalOffset(fp, arg.Indexes[idxCounter])) * sizeToUse
			//fmt.Printf("BASE_OFFSET %d, FINAL_OFFSET %d, SIZE_TO_USE %d\n",
			//	baseOffset, finalOffset, sizeToUse)
			if !IsValidSliceIndex(baseOffset, finalOffset, sizeToUse) {
				panic(constants.CX_RUNTIME_SLICE_INDEX_OUT_OF_RANGE)
			}

			idxCounter++

		case constants.DEREF_ARRAY: //TODO: Move to CalculateDereference_array
			types.FMTDEBUG(fmt.Sprintf("DEREF_ARRAY\n"))
			if len(arg.Indexes) == 0 {
				continue
			}
			var subSize = types.Pointer(1) // TODO: PTR remove hardcode 1
			for _, len := range arg.Lengths[idxCounter+1:] { // TODO: PTR remove hardcode 1
				subSize *= len
			}

			//TODO: Delete
			sizeToUse := GetDerefSize(arg) //TODO: is always arg.Size unless arg.CustomType != nil

			baseOffset = finalOffset
			sizeofElement = subSize * sizeToUse
			//tmpOO := types.Read_ptr(PROGRAM.Memory, GetFinalOffset(fp, arg.Indexes[idxCounter]))
			//tmpVV := types.Read_ptr(PROGRAM.Memory, GetFinalOffset(fp, arg.Indexes[idxCounter])) * sizeofElement
			//fmt.Printf("SIZEOF_ELEMENT %d, OFFSET %d, INDEX_OFFSET %d, INDEX_VALUE %v\n",
			//	sizeofElement, GetFinalOffset(fp, arg.Indexes[idxCounter]), tmpOO, tmpVV)
			PrintArg("DEREF_ARRAY: ", arg)
			finalOffset += types.Read_ptr(PROGRAM.Memory, GetFinalOffset(fp, arg.Indexes[idxCounter])) * sizeofElement
			idxCounter++
		case constants.DEREF_POINTER: //TODO: Move to CalculateDereference_ptr
			types.FMTDEBUG(fmt.Sprintf("DEREF_POINTER\n"))
			isPointer = true
			finalOffset = types.Read_ptr(PROGRAM.Memory, finalOffset)
		}
	}

	// if finalOffset >= PROGRAM.HeapStartsAt {
	if finalOffset.IsValid() && finalOffset >= PROGRAM.HeapStartsAt && isPointer {
		// then it's an object
		fmt.Printf("PROGRAM HEAP_START FINAL_OFFSET %d\n", finalOffset)
		finalOffset += types.OBJECT_HEADER_SIZE
		if arg.IsSlice {
			finalOffset += constants.SLICE_HEADER_SIZE
			if !IsValidSliceIndex(baseOffset, finalOffset, sizeofElement) {
				panic(constants.CX_RUNTIME_SLICE_INDEX_OUT_OF_RANGE)
			}
		}
	}

	return finalOffset
}