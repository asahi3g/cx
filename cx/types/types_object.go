package types

import (
	"fmt"
)

func Get_obj_size(data []byte) Pointer {
	return Cast_int_to_ptr(len(data)) + OBJECT_HEADER_SIZE
}

func Make_obj(data []byte) []byte {
	size := Get_obj_size(data)
	obj := make([]byte, size)	
	Write_obj_data(obj, 0, data)
	return obj
}

func AllocWrite_obj_data(memory []byte, obj []byte) Pointer {
	heapOffset := Allocator(Get_obj_size(obj))
	fmt.Printf("ALLOC_WRITE_OBJ_DATA HEAPOFFSET %d, OBJ '%v'\n", heapOffset, obj)
	Write_obj_data(memory, heapOffset, obj)
	return heapOffset
}

func Write_obj_data(memory []byte, offset Pointer, obj []byte) {
	size := Cast_int_to_ptr(len(obj))
	fmt.Printf("WRITE_OBJECT_DATA SIZE %d, OBJ `%v`\n", size, obj)
	Write_ptr(memory, offset+OBJECT_HEADER_SIZE-OBJECT_SIZE, size)
	WriteSlice_byte(memory, offset+OBJECT_HEADER_SIZE, obj)
}

func Read_obj_data(memory []byte, offset Pointer) []byte {
	size := Read_ptr(memory, offset+OBJECT_HEADER_SIZE-OBJECT_SIZE)
	obj := GetSlice_byte(memory, offset+OBJECT_HEADER_SIZE, size)
	fmt.Printf("READ_OBJECT_DATA OFFSET %d, SIZE %d OBJ `%v`\n", offset, size, obj)
	return obj
}

func Write_obj(memory []byte, offset Pointer, obj []byte) {
	heapOffset := AllocWrite_obj_data(memory, obj)
	fmt.Printf("WRITE_OBJ OFFSET %d, HEAPOFFSET %d, OBJ %v\n", offset, heapOffset, obj)
	Write_ptr(memory, offset, heapOffset)
}

func Read_obj(memory []byte, offset Pointer) []byte {
	panic("FUCK\n")
	heapOffset := Read_ptr(memory, offset)
	obj := Read_obj_data(memory, heapOffset)
	fmt.Printf("READ_OBJ OFFSET %d, HEAPOFFSET %d, OBJ `%v`\n", offset, heapOffset, obj)
	return obj
}
