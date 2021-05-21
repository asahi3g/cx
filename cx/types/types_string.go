package types

import (
	"fmt"
)

func AllocWrite_str_data(memory []byte, str string) Pointer {
	return AllocWrite_obj_data(memory, []byte(str))
}

func Write_str_data(memory []byte, offset Pointer, value string) {
	FMTDEBUG(fmt.Sprintf("WRITE_STRING_DATA `%s`\n", value))
	Write_obj_data(memory, offset, []byte(value))
}

func Read_str_data(memory []byte, offset Pointer) string {
	str := Read_obj_data(memory, offset)
	FMTDEBUG(fmt.Sprintf("READ_STRING_DATA `%s`\n", str))
	return string(str)
}

func Write_str(memory []byte, offset Pointer, str string) {
	FMTDEBUG(fmt.Sprintf("WRITE_STR OFFSET %d STR `%s`\n", offset, str))
	Write_obj(memory, offset, []byte(str))
}

func Read_str(memory []byte, offset Pointer) string {
	str := string(Read_obj(memory, offset))
	FMTDEBUG(fmt.Sprintf("READ_STR OFFSET %d STR `%s`\n", offset, str))
	return str
}

func Read_str_size(memory []byte, offset Pointer) Pointer {
	return Read_obj_size(memory, offset)
}