// +build ptr32

package types

type Pointer uint32
const TYPE_POINTER_SIZE = Pointer(4)
const InvalidPointer = Pointer(^uint32(0))

// Cast_int_to_ptr ...
func Cast_int_to_ptr(value int) Pointer {
    // TODO: assertions
    return Pointer(value)
}

func Cast_ptr_to_int(value Pointer) int {
    // TODO: assertions
    return int(value)
}

func Cast_ptr_to_i32(value Pointer) int32 {
    // TODO: assertions
	return int32(value)
}

func Cast_i32_to_ptr(value int32) Pointer {
    // TODO: assertions
    return Pointer(value)
}

func Cast_i64_to_ptr(value int64) Pointer {
    // TODO: assertions
    return Pointer(value)
}

func Cast_ui64_to_ptr(value uint64) Pointer {
    // TODO: assertions
    return Pointer(value)
}

func Cast_f32_to_ptr(value float32) Pointer {
    // TODO: assertions
    return Pointer(value)
}

func Cast_f64_to_ptr(value float64) Pointer {
    // TODO: assertions
    return Pointer(value)
}

// Cast_ptr_to_ui32 ...
func Cast_ptr_to_ui32(value Pointer) uint32 {
    // TODO: assertions
    return uint32(value)
}

// Cast_ptr_to_ui64 ...
func Cast_ptr_to_ui64(value Pointer) uint64 {
    // TODO: assertions
    return uint64(value)
}

// Deserialize_ptr ...
func Read_ptr(memory []byte, offset Pointer) Pointer {
    return Pointer(Read_ui32(memory, offset))
}

// Serialize_ptr ...
func Write_ptr(memory []byte, offset Pointer, value Pointer) {
    Write_ui32(memory, offset, Cast_ptr_to_ui32(value))
}

