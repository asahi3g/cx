package ast

import (
	"github.com/skycoin/cx/cx/constants"
    "github.com/skycoin/cx/cx/types"


	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func DeserializeRaw(byts []byte, offset types.Pointer, size types.Pointer, item interface{}) {
	_, err := encoder.DeserializeRaw(byts[offset:offset+size], item)
	if err != nil {
		panic(err)
	}
}

func ReadMemory(offset types.Pointer, arg *CXArgument) []byte {
	size := GetSize(arg)
	return PROGRAM.Memory[offset : offset+size]
}

// ReadStr ...
func ReadStr(fp types.Pointer, inp *CXArgument) (out string) {
	off := GetFinalOffset(fp, inp)
	return ReadStrFromOffset(off, inp)
}

// ReadStrFromOffset ...
func ReadStrFromOffset(off types.Pointer, inp *CXArgument) (out string) {
	var offset types.Pointer
	if inp.ArgDetails.Name == "" {
		// Then it's a literal.
		offset = off
	} else {
		offset = types.Read_ptr(PROGRAM.Memory, off)
	}

	if offset == 0 {
		// Then it's nil string.
		out = ""
		return
	}

	// We need to check if the string lives on the data segment or on the
	// heap to know if we need to take into consideration the object header's size.
	if offset > PROGRAM.HeapStartsAt {
		size := types.Read_ptr(PROGRAM.Memory, offset+constants.OBJECT_HEADER_SIZE)
		DeserializeRaw(PROGRAM.Memory, offset+constants.OBJECT_HEADER_SIZE, constants.STR_HEADER_SIZE+size, &out)
	} else {
		size := types.Read_ptr(PROGRAM.Memory, offset)
		DeserializeRaw(PROGRAM.Memory, offset, constants.STR_HEADER_SIZE+size, &out)
	}

	return out
}

// ReadStringFromObject reads the string located at offset `off`.
func ReadStringFromObject(off types.Pointer) string {
	var plusOff types.Pointer
	if off > PROGRAM.HeapStartsAt {
		// Found in heap segment.
		plusOff += constants.OBJECT_HEADER_SIZE
	}

	size := types.Read_ptr(PROGRAM.Memory, off+plusOff)

	str := ""
	DeserializeRaw(PROGRAM.Memory, off+plusOff, constants.STR_HEADER_SIZE+size, &str)
	return str
}



// FromStr ...
func FromStr(in string) []byte {
	return encoder.Serialize(in)
}

// WriteObjectRef
// WARNING, is using heap variables?
//Is this "Write object ot heap?"
func WriteObjectData(obj []byte) types.Pointer {
	size := types.Cast_int_to_ptr(len(obj)) + constants.OBJECT_HEADER_SIZE
	heapOffset := AllocateSeq(size)
	types.Write_ptr(PROGRAM.Memory, heapOffset, size)
	types.WriteSlice_byte(PROGRAM.Memory, heapOffset+constants.OBJECT_HEADER_SIZE, obj)
	return heapOffset
}

// WriteObject ...
func WriteObject(offset types.Pointer, obj []byte) {
	heapOffset := WriteObjectData(obj)
	types.Write_ptr(PROGRAM.Memory, offset, heapOffset)
}

// WriteStringData writes `str` to the heap as an object and returns its absolute offset.
func WriteStringData(str string) types.Pointer {
	return WriteObjectData(encoder.Serialize(str))
}

// WriteString writes the string `str` on memory, starting at byte number `fp`.
func WriteString(fp types.Pointer, str string, out *CXArgument) {
	WriteObject(GetFinalOffset(fp, out), encoder.Serialize(str))
}

// GetStrOffset ...
func GetStrOffset(offset types.Pointer, name string) types.Pointer {
	if name != "" {
		// then it's not a literal
		return types.Read_ptr(PROGRAM.Memory, offset)
	}
	return offset
}

// ResizeMemory ...
func ResizeMemory(prgrm *CXProgram, newMemSize types.Pointer, isExpand bool) {
	// We can't expand memory to a value greater than `memLimit`.
	if newMemSize > constants.MAX_HEAP_SIZE {
		newMemSize = constants.MAX_HEAP_SIZE
	}

	if newMemSize == prgrm.HeapSize {
		// Then we're at the limit; we can't expand anymore.
		// We can only hope that the free memory is enough for the CX program to continue running.
		return
	}

	if isExpand {
		// Adding bytes to reach a heap equal to `newMemSize`.
		prgrm.Memory = append(prgrm.Memory, make([]byte, newMemSize-prgrm.HeapSize)...)
		prgrm.HeapSize = newMemSize
	} else {
		// Removing bytes to reach a heap equal to `newMemSize`.
		prgrm.Memory = append([]byte(nil), prgrm.Memory[:prgrm.HeapStartsAt+newMemSize]...)
		prgrm.HeapSize = newMemSize
	}
}

// AllocateSeq allocates memory in the heap
func AllocateSeq(size types.Pointer) (offset types.Pointer) {
	// Current object trying to be allocated would use this address.
	addr := PROGRAM.HeapPointer
	// Next object to be allocated will use this address.
	newFree := addr + size

	// Checking if we can allocate the entirety of the object in the current heap.
	if newFree > PROGRAM.HeapSize {
		// It does not fit, so calling garbage collector.
		MarkAndCompact(PROGRAM)
		// Heap pointer got moved by GC and recalculate these variables based on the new pointer.
		addr = PROGRAM.HeapPointer
		newFree = addr + size

		// If the new heap pointer exceeds `MAX_HEAP_SIZE`, there's nothing left to do.
		if newFree > constants.MAX_HEAP_SIZE {
			panic(constants.HEAP_EXHAUSTED_ERROR)
		}

		// According to MIN_HEAP_FREE_RATIO and MAX_HEAP_FREE_RATION we can either shrink
		// or expand the heap to maintain "healthy" heap sizes. The idea is that we don't want
		// to have an absurdly amount of free heap memory, as we would be wasting resources, and we
		// don't want to have a small amount of heap memory left as we'd be calling the garbage collector
		// too frequently.

		// Calculating free heap memory percentage.
		usedPerc := float32(newFree) / float32(PROGRAM.HeapSize)
		freeMemPerc := 1.0 - usedPerc

		// Then we have less than MIN_HEAP_FREE_RATIO memory left. Expand!
		if freeMemPerc < constants.MIN_HEAP_FREE_RATIO {
			// Calculating new heap size in order to reach MIN_HEAP_FREE_RATIO.
			newMemSize := types.Cast_f32_to_ptr(float32(newFree) / (1.0 - constants.MIN_HEAP_FREE_RATIO))
			ResizeMemory(PROGRAM, newMemSize, true)
		}

		// Then we have more than MAX_HEAP_FREE_RATIO memory left. Shrink!
		if freeMemPerc > constants.MAX_HEAP_FREE_RATIO {
			// Calculating new heap size in order to reach MAX_HEAP_FREE_RATIO.
			newMemSize := types.Cast_f32_to_ptr(float32(newFree) / (1.0 - constants.MAX_HEAP_FREE_RATIO))

			// This check guarantees that the CX program has always at least INIT_HEAP_SIZE bytes to work with.
			// A flag could be added later to remove this, as in some cases this mechanism could not be desired.
			if newMemSize > constants.INIT_HEAP_SIZE {
				ResizeMemory(PROGRAM, newMemSize, false)
			}
		}
	}

	PROGRAM.HeapPointer = newFree

	// Returning absolute memory address (not relative to where heap starts at).
	// Above this point we were performing all operations taking into
	// consideration only heap offsets.
	return addr + PROGRAM.HeapStartsAt
}
