// +build ptr32

package types

import (
    "fmt"
)

type Pointer uint32
const TYPE_POINTER_SIZE = UI32_SIZE
const InvalidPointer = Pointer(MAX_UINT32)

func Cast_ptr_to_int(value Pointer) int {
    panicIf(int64(value) > int64(MAX_INT), fmt.Sprintf("Invalid Cast_ptr_to_int %v\n", value), CX_RUNTIME_INVALID_CAST)    
    return int(value)
}

func Cast_ptr_to_i32(value Pointer) int32 {
    panicIf(value < 0 || int64(value) > int64(MAX_INT32), fmt.Sprintf("Invalid Cast_ptr_to_i32 %v\n", value), CX_RUNTIME_INVALID_CAST)
	return int32(value)
}

func Cast_ptr_to_ui32(value Pointer) uint32 {
    return uint32(value)
}

func Cast_ptr_to_ui64(value Pointer) uint64 {
    return uint64(value)
}

func Cast_int_to_ptr(value int) Pointer {
    panicIf(value < 0 || int64(value) > int64(InvalidPointer), fmt.Sprintf("Invalid Cast_int_to_ptr %v\n", value), CX_RUNTIME_INVALID_CAST)
    return Pointer(value)
}

func Cast_i32_to_ptr(value int32) Pointer {
    panicIf(value < 0, fmt.Sprintf("Invalid Cast_i32_to_ptr %v\n", value), CX_RUNTIME_INVALID_CAST)
    return Pointer(value)
}

func Cast_i64_to_ptr(value int64) Pointer {
    panicIf(value < 0 || value > int64(InvalidPointer), fmt.Sprintf("Invalid Cast_i64_to_ptr %v\n", value), CX_RUNTIME_INVALID_CAST)
    return Pointer(value)
}

func Cast_ui64_to_ptr(value uint64) Pointer {
    panicIf(value > uint64(InvalidPointer), fmt.Sprintf("Invalid Cast_ui64_to_ptr %v\n", value), CX_RUNTIME_INVALID_CAST)
    return Pointer(value)
}

func Cast_f32_to_ptr(value float32) Pointer {
    panicIf(value < 0 || uint64(value) > uint64(InvalidPointer), fmt.Sprintf("Invalid Cast_f32_to_ptr %v\n", value), CX_RUNTIME_INVALID_CAST)
    return Pointer(value)
}

func Cast_f64_to_ptr(value float64) Pointer {
    panicIf(value < 0 || uint64(value) > uint64(InvalidPointer), fmt.Sprintf("Invalid Cast_f64_to_ptr %v\n", value), CX_RUNTIME_INVALID_CAST)
    return Pointer(value)
}

// Deserialize_ptr ...
func Read_ptr(memory []byte, offset Pointer) Pointer {
    return Pointer(Read_ui32(memory, offset))
}

// Serialize_ptr ...
func Write_ptr(memory []byte, offset Pointer, value Pointer) {
    Write_ui32(memory, offset, Cast_ptr_to_ui32(value))
}

