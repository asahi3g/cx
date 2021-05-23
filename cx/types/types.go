package types

import (
	"math"
	"fmt"
)


const (
	CX_SUCCESS = iota //zero can be success
	CX_COMPILATION_ERROR
	CX_PANIC // 2
	CX_INTERNAL_ERROR
	CX_ASSERT
	CX_RUNTIME_ERROR
	CX_RUNTIME_STACK_OVERFLOW_ERROR
	CX_RUNTIME_HEAP_EXHAUSTED_ERROR
	CX_RUNTIME_INVALID_ARGUMENT
	CX_RUNTIME_SLICE_INDEX_OUT_OF_RANGE
	CX_RUNTIME_NOT_IMPLEMENTED
	CX_RUNTIME_INVALID_CAST
)

const BOOL_SIZE = Pointer(1)

const I8_SIZE = Pointer(1)
const I16_SIZE = Pointer(2)
const I32_SIZE = Pointer(4)
const I64_SIZE = Pointer(8)

const UI8_SIZE = Pointer(1)
const UI16_SIZE = Pointer(2)
const UI32_SIZE = Pointer(4)
const UI64_SIZE = Pointer(8)

const F32_SIZE = Pointer(4)
const F64_SIZE = Pointer(8)

const STR_SIZE = TYPE_POINTER_SIZE

const MAX_INT = int(MAX_UINT >> 1)
const MAX_INT32 = int(MAX_UINT32 >> 1)
const MIN_INT32 = -MAX_INT32 - 1

const MAX_UINT = ^uint(0)
const MAX_UINT8 = ^uint8(0)
const MAX_UINT16 = ^uint16(0)
const MAX_UINT32 = ^uint32(0)
const MAX_UINT64 = ^uint64(0)


type AllocatorHandler func (Pointer) Pointer
var Allocator AllocatorHandler

func panicIf(condition bool, message string, error int) {
	if condition {
		fmt.Printf(message)
		panic(error)
	}
}

var DEBUG bool = true
func FMTDEBUG(message string) {
	if DEBUG {
	fmt.Printf(message)
}
}

func (pointer Pointer) IsValid() bool {
	return pointer != InvalidPointer
}

func Cast_sint_to_sptr(value []int) []Pointer {
    l := len(value)
    if l == 0 {
        return nil
    }

    sptr := make([]Pointer, l)
    for i, k := range value {
        sptr[i] = Cast_int_to_ptr(k)
    }
    return sptr
}

func Read_bool(memory []byte, offset Pointer) bool {
	memory = memory[offset:]
	panicIf(len(memory) < 1, "invalid memory len", CX_INTERNAL_ERROR)
	return memory[0] != 0
}

func Read_i8(memory []byte, offset Pointer) int8 {
	memory = memory[offset:]
	fmt.Printf("READ_I8 : OFFSET %d\n", offset)
	panicIf(len(memory) < 1, "invalid memory len", CX_INTERNAL_ERROR)
	return int8(memory[0])
}

func Read_i16(memory []byte, offset Pointer) int16 {
	memory = memory[offset:]
	panicIf(len(memory) < 2, "invalid memory len", CX_INTERNAL_ERROR)
	return int16(memory[0]) | int16(memory[1])<<8
}

func Read_i32(memory []byte, offset Pointer) int32 {
	memory = memory[offset:]
	panicIf(len(memory) < 4, "invalid memory len", CX_INTERNAL_ERROR)
	return int32(memory[0]) | int32(memory[1])<<8 | int32(memory[2])<<16 | int32(memory[3])<<24
}

func Read_i64(memory []byte, offset Pointer) int64 {
	memory = memory[offset:]
	panicIf(len(memory) < 8, "invalid memory len", CX_INTERNAL_ERROR)
	return int64(memory[0]) | int64(memory[1])<<8 | int64(memory[2])<<16 | int64(memory[3])<<24 |
		int64(memory[4])<<32 | int64(memory[5])<<40 | int64(memory[6])<<48 | int64(memory[7])<<56
}

func Read_ui8(memory []byte, offset Pointer) uint8 {
	memory = memory[offset:]
	panicIf(len(memory) < 1, "invalid memory len", CX_INTERNAL_ERROR)
	return uint8(memory[0])
}

func Read_ui16(memory []byte, offset Pointer) uint16 {
	memory = memory[offset:]
	panicIf(len(memory) < 2, "invalid memory len", CX_INTERNAL_ERROR)
	return uint16(memory[0]) | uint16(memory[1])<<8
}

func Read_ui32(memory []byte, offset Pointer) uint32 {
	memory = memory[offset:]
	panicIf(len(memory) < 4, "invalid memory len", CX_INTERNAL_ERROR)
	return uint32(memory[0]) | uint32(memory[1])<<8 | uint32(memory[2])<<16 | uint32(memory[3])<<24
}

func Read_ui64(memory []byte, offset Pointer) uint64 {
	memory = memory[offset:]
	panicIf(len(memory) < 8, "invalid memory len", CX_INTERNAL_ERROR)
	return uint64(memory[0]) | uint64(memory[1])<<8 | uint64(memory[2])<<16 | uint64(memory[3])<<24 |
		uint64(memory[4])<<32 | uint64(memory[5])<<40 | uint64(memory[6])<<48 | uint64(memory[7])<<56
}

func Read_f32(memory []byte, offset Pointer) float32 {
	return math.Float32frombits(Read_ui32(memory, offset))
}

func Read_f64(memory []byte, offset Pointer) float64 {
	return math.Float64frombits(Read_ui64(memory, offset))
}

func ReadSlice_i8(memory []byte, offset Pointer) (out []int8) {
	count := Cast_int_to_ptr(len(memory))
	if count > 0 {
		out = make([]int8, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_i8(memory, i)
		}
	}
	return
}

func ReadSlice_i16(memory []byte, offset Pointer) (out []int16) {
	count := Cast_int_to_ptr(len(memory) / 2)
	if count > 0 {
		out = make([]int16, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_i16(memory, i*2)
		}
	}
	return
}

func ReadSlice_i32(memory []byte, offset Pointer) (out []int32) {
	count := Cast_int_to_ptr(len(memory) / 4)
	if count > 0 {
		out = make([]int32, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_i32(memory, i*4)
		}
	}
	return
}

func ReadSlice_i64(memory []byte, offset Pointer) (out []int64) {
	count := Cast_int_to_ptr(len(memory) / 8)
	if count > 0 {
		out = make([]int64, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_i64(memory, i*8)
		}
	}
	return
}

func ReadSlice_ui8(memory []byte, offset Pointer) (out []uint8) {
	count := Cast_int_to_ptr(len(memory))
	if count > 0 {
		out = make([]uint8, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_ui8(memory, i)
		}
	}
	return
}

func ReadSlice_ui16(memory []byte, offset Pointer) (out []uint16) {
	count := Cast_int_to_ptr(len(memory) / 2)
	if count > 0 {
		out = make([]uint16, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_ui16(memory, i*2)
		}
	}
	return
}

func ReadSlice_ui32(memory []byte, offset Pointer) (out []uint32) {
	count := Cast_int_to_ptr(len(memory) / 4)
	if count > 0 {
		out = make([]uint32, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_ui32(memory, i*4)
		}
	}
	return
}

func ReadSlice_ui64(memory []byte, offset Pointer) (out []uint64) {
	count := Cast_int_to_ptr(len(memory) / 8)
	if count > 0 {
		out = make([]uint64, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_ui64(memory, i*8)
		}
	}
	return
}

func ReadSlice_f32(memory []byte, offset Pointer) (out []float32) {
	count := Cast_int_to_ptr(len(memory) / 4)
	if count > 0 {
		out = make([]float32, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_f32(memory, i*4)
		}
	}
	return
}

func ReadSlice_f64(memory []byte, offset Pointer) (out []float64) {
	count := Cast_int_to_ptr(len(memory) / 8)
	if count > 0 {
		out = make([]float64, count)
		for i := Pointer(0); i < count; i++ {
			out[i] = Read_f64(memory, i*8)
		}
	}
	return
}

func GetSlice_byte(memory []byte, offset Pointer, size Pointer) []byte {
	return memory[offset : offset + size]
}

func Write_bool(memory []byte, offset Pointer, value bool) {
	if value {
		memory[offset] = 1
	} else {
		memory[offset] = 0		
	}
}

func Write_i8(mem []byte, offset Pointer, v int8) {
	mem[offset] = byte(v)
}

func Write_i16(mem []byte, offset Pointer, v int16) {
	mem[offset] = byte(v)
	mem[offset+1] = byte(v >> 8)
}

func Write_i32(mem []byte, offset Pointer, v int32) {
	mem[offset] = byte(v)
	mem[offset+1] = byte(v >> 8)
	mem[offset+2] = byte(v >> 16)
	mem[offset+3] = byte(v >> 24)
}

func Write_i64(mem []byte, offset Pointer, v int64) {
	mem[offset] = byte(v)
	mem[offset+1] = byte(v >> 8)
	mem[offset+2] = byte(v >> 16)
	mem[offset+3] = byte(v >> 24)
	mem[offset+4] = byte(v >> 32)
	mem[offset+5] = byte(v >> 40)
	mem[offset+6] = byte(v >> 48)
	mem[offset+7] = byte(v >> 56)
}

func Write_ui8(mem []byte, offset Pointer, v uint8) {
	mem[offset] = v
}

func Write_ui16(mem []byte, offset Pointer, v uint16) {
	mem[offset] = byte(v)
	mem[offset+1] = byte(v >> 8)
}

func Write_ui32(memory []byte, offset Pointer, value uint32) {
	memory[offset] = byte(value)
	memory[offset+1] = byte(value >> 8)
	memory[offset+2] = byte(value >> 16)
	memory[offset+3] = byte(value >> 24)
}

func Write_ui64(memory []byte, offset Pointer, value uint64) {
	memory[offset] = byte(value)
	memory[offset+1] = byte(value >> 8)
	memory[offset+2] = byte(value >> 16)
	memory[offset+3] = byte(value >> 24)
	memory[offset+4] = byte(value >> 32)
	memory[offset+5] = byte(value >> 40)
	memory[offset+6] = byte(value >> 48)
	memory[offset+7] = byte(value >> 56)
}

func Write_f32(mem []byte, offset Pointer, f float32) {
	v := math.Float32bits(f)
	mem[offset] = byte(v)
	mem[offset+1] = byte(v >> 8)
	mem[offset+2] = byte(v >> 16)
	mem[offset+3] = byte(v >> 24)
}

func Write_f64(mem []byte, offset Pointer, f float64) {
	v := math.Float64bits(f)
	mem[offset] = byte(v)
	mem[offset+1] = byte(v >> 8)
	mem[offset+2] = byte(v >> 16)
	mem[offset+3] = byte(v >> 24)
	mem[offset+4] = byte(v >> 32)
	mem[offset+5] = byte(v >> 40)
	mem[offset+6] = byte(v >> 48)
	mem[offset+7] = byte(v >> 56)
}

func WriteSlice_byte(memory []byte, offset Pointer, byts []byte) {
	// TODO: PTR use copy()
    count := Cast_int_to_ptr(len(byts))
	for c := Pointer(0); c < count; c++ {
		memory[offset+c] = byts[c]
	}
}
