package types

import (
	"fmt"
)

const MARK_SIZE = Pointer(1)
const FORWARDING_ADDRESS_SIZE = TYPE_POINTER_SIZE
const OBJECT_GC_HEADER_SIZE = MARK_SIZE + FORWARDING_ADDRESS_SIZE

const OBJECT_SIZE = TYPE_POINTER_SIZE
const OBJECT_HEADER_SIZE = OBJECT_GC_HEADER_SIZE + OBJECT_SIZE


func Get_obj_header(memory []byte, offset Pointer) []byte {
	return memory[offset : offset+OBJECT_HEADER_SIZE]
}

func Get_obj_data(memory []byte, offset Pointer, size Pointer) []byte {
	offset+=OBJECT_HEADER_SIZE
	return memory[offset: offset+size]
}

func Compute_obj_size(data []byte) Pointer {
	return Cast_int_to_ptr(len(data)) + OBJECT_HEADER_SIZE
}

func Read_obj_forwarding_address(memory []byte, offset Pointer) Pointer {
	return Read_ptr(memory, offset+MARK_SIZE)
}

func Read_obj_size(memory []byte, offset Pointer) Pointer {
	return Read_ptr(memory, offset+OBJECT_HEADER_SIZE-OBJECT_SIZE)
}

func Write_obj_size(memory []byte, offset Pointer, size Pointer) {
	Write_ptr(memory, offset+OBJECT_HEADER_SIZE-OBJECT_SIZE, size)
}

func Write_obj_mark(memory []byte, offset Pointer, mark byte) {
	Write_ui8(memory, offset, mark)
}

func Write_obj_forwarding_address(memory []byte, offset Pointer, address Pointer) {
	Write_ptr(memory, offset+MARK_SIZE, address)
}

func Make_obj(data []byte) []byte {
	size := Compute_obj_size(data)
	obj := make([]byte, size)	
	Write_obj_data(obj, 0, data)
	return obj
}

func AllocWrite_obj_data(memory []byte, obj []byte) Pointer {
	heapOffset := Allocator(Compute_obj_size(obj))
	FMTDEBUG(fmt.Sprintf("ALLOC_WRITE_OBJ_DATA HEAPOFFSET %d, OBJ '%v'\n", heapOffset, obj))
	Write_obj_data(memory, heapOffset, obj)
	return heapOffset
}

func Write_obj_data(memory []byte, offset Pointer, obj []byte) {
	size := Cast_int_to_ptr(len(obj))
	FMTDEBUG(fmt.Sprintf("WRITE_OBJECT_DATA SIZE %d, OBJ `%v`\n", size, obj))
	Write_obj_size(memory, offset, size)
	WriteSlice_byte(memory, offset+OBJECT_HEADER_SIZE, obj)
}

func Read_obj_data(memory []byte, offset Pointer) []byte {
	size := Read_obj_size(memory, offset)
	obj := GetSlice_byte(memory, offset+OBJECT_HEADER_SIZE, size)
	FMTDEBUG(fmt.Sprintf("READ_OBJECT_DATA OFFSET %d, SIZE %d OBJ `%v`\n", offset, size, obj))
	return obj
}

func Write_obj(memory []byte, offset Pointer, obj []byte) {
	heapOffset := AllocWrite_obj_data(memory, obj)
	FMTDEBUG(fmt.Sprintf("WRITE_OBJ OFFSET %d, HEAPOFFSET %d, OBJ %v\n", offset, heapOffset, obj))
	Write_ptr(memory, offset, heapOffset)
}

func Read_obj(memory []byte, offset Pointer) []byte {
	//panic("FUCK\n")
	heapOffset := Read_ptr(memory, offset)
	obj := Read_obj_data(memory, heapOffset)
	FMTDEBUG(fmt.Sprintf("READ_OBJ OFFSET %d, HEAPOFFSET %d, OBJ `%v`\n", offset, heapOffset, obj))
	return obj
}
