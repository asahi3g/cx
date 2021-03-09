package cxcore

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// The built-in str function returns the base 10 string representation of operand 1.
func opF32ToStr(expr *CXExpression, fp int) {
	outB0 := FromStr(strconv.FormatFloat(float64(ReadF32(fp, expr.Inputs[0])), 'f', -1, 32))
	WriteObject(GetOffset_str(fp, expr.Outputs[0]), outB0)
}

// The built-in i8 function returns operand 1 casted from type f32 to type i8.
func opF32ToI8(expr *CXExpression, fp int) {
	outV0 := int8(ReadF32(fp, expr.Inputs[0]))
	WriteI8(GetOffset_i8(fp, expr.Outputs[0]), outV0)
}

// The built-in i16 function returns operand 1 casted from type f32 to type i16.
func opF32ToI16(expr *CXExpression, fp int) {
	outV0 := int16(ReadF32(fp, expr.Inputs[0]))
	WriteI16(GetOffset_i16(fp, expr.Outputs[0]), outV0)
}

// The built-in i32 function return operand 1 casted from type f32 to type i32.
func opF32ToI32(expr *CXExpression, fp int) {
	outV0 := int32(ReadF32(fp, expr.Inputs[0]))
	WriteI32(GetOffset_i32(fp, expr.Outputs[0]), outV0)
}

// The built-in i64 function returns operand 1 casted from type f32 to type i64.
func opF32ToI64(expr *CXExpression, fp int) {
	outV0 := int64(ReadF32(fp, expr.Inputs[0]))
	WriteI64(GetOffset_i64(fp, expr.Outputs[0]), outV0)
}

// The built-in ui8 function returns operand 1 casted from type f32 to type ui8.
func opF32ToUI8(expr *CXExpression, fp int) {
	outV0 := uint8(ReadF32(fp, expr.Inputs[0]))
	WriteUI8(GetOffset_ui8(fp, expr.Outputs[0]), outV0)
}

// The built-in ui16 function returns the operand 1 casted from type f32 to type ui16.
func opF32ToUI16(expr *CXExpression, fp int) {
	outV0 := uint16(ReadF32(fp, expr.Inputs[0]))
	WriteUI16(GetOffset_ui16(fp, expr.Outputs[0]), outV0)
}

// The built-in ui32 function returns the operand 1 casted from type f32 to type ui32.
func opF32ToUI32(expr *CXExpression, fp int) {
	outV0 := uint32(ReadF32(fp, expr.Inputs[0]))
	WriteUI32(GetOffset_ui32(fp, expr.Outputs[0]), outV0)
}

// The built-in ui64 function returns the operand 1 casted from type f32 to type ui64.
func opF32ToUI64(expr *CXExpression, fp int) {
	outV0 := uint64(ReadF32(fp, expr.Inputs[0]))
	WriteUI64(GetOffset_ui64(fp, expr.Outputs[0]), outV0)
}

// The built-in f64 function returns operand 1 casted from type f32 to type f64.
func opF32ToF64(expr *CXExpression, fp int) {
	outV0 := float64(ReadF32(fp, expr.Inputs[0]))
	WriteF64(GetOffset_f64(fp, expr.Outputs[0]), outV0)
}

// The built-in isnan function returns true if operand is nan value.
func opF32Isnan(expr *CXExpression, fp int) {
	outV0 := math.IsNaN(float64(ReadF32(fp, expr.Inputs[0])))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The print built-in function formats its arguments and prints them.
func opF32Print(expr *CXExpression, fp int) {
	fmt.Println(ReadF32(fp, expr.Inputs[0]))
}

// The built-in add function returns the sum of the two operands.
func opF32Add(expr *CXExpression, fp int) {
	outV0 := ReadF32(fp, expr.Inputs[0]) + ReadF32(fp, expr.Inputs[1])
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in sub function returns the difference between the two operands.
func opF32Sub(expr *CXExpression, fp int) {
	outV0 := ReadF32(fp, expr.Inputs[0]) - ReadF32(fp, expr.Inputs[1])
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in neg function returns the opposite of operand 1.
func opF32Neg(expr *CXExpression, fp int) {
	outV0 := -ReadF32(fp, expr.Inputs[0])
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in mul function returns the product of the two operands.
func opF32Mul(expr *CXExpression, fp int) {
	outV0 := ReadF32(fp, expr.Inputs[0]) * ReadF32(fp, expr.Inputs[1])
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in div function returns the quotient between the two operands.
func opF32Div(expr *CXExpression, fp int) {
	outV0 := ReadF32(fp, expr.Inputs[0]) / ReadF32(fp, expr.Inputs[1])
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in mod function return the floating-point remainder of operand 1 divided by operand 2.
func opF32Mod(expr *CXExpression, fp int) {
	outV0 := float32(math.Mod(float64(ReadF32(fp, expr.Inputs[0])), float64(ReadF32(fp, expr.Inputs[1]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in abs function returns the absolute value of the operand.
func opF32Abs(expr *CXExpression, fp int) {
	outV0 := float32(math.Abs(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in pow function returns x**n for n>0 otherwise 1.
func opF32Pow(expr *CXExpression, fp int) {
	outV0 := float32(math.Pow(float64(ReadF32(fp, expr.Inputs[0])), float64(ReadF32(fp, expr.Inputs[1]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in gt function returns true if operand 1 is greater than operand 2.
func opF32Gt(expr *CXExpression, fp int) {
	var outV0 bool = (ReadF32(fp, expr.Inputs[0]) > ReadF32(fp, expr.Inputs[1]))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The built-in gteq function returns true if the operand 1 is greater than or
// equal to operand 2.
func opF32Gteq(expr *CXExpression, fp int) {
	var outV0 bool = (ReadF32(fp, expr.Inputs[0]) >= ReadF32(fp, expr.Inputs[1]))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The built-in lt function returns true if operand 1 is less than operand 2.
func opF32Lt(expr *CXExpression, fp int) {
	var outV0 bool = (ReadF32(fp, expr.Inputs[0]) < ReadF32(fp, expr.Inputs[1]))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The built-in lteq function returns true if operand 1 is less than or
// equal to operand 2.
func opF32Lteq(expr *CXExpression, fp int) {
	var outV0 bool = (ReadF32(fp, expr.Inputs[0]) <= ReadF32(fp, expr.Inputs[1]))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The built-in eq function returns true if operand 1 is equal to operand 2.
func opF32Eq(expr *CXExpression, fp int) {
	var outV0 bool = (ReadF32(fp, expr.Inputs[0]) == ReadF32(fp, expr.Inputs[1]))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The built-in uneq function returns true operand1 is different from operand 2.
func opF32Uneq(expr *CXExpression, fp int) {
	var outV0 bool = (ReadF32(fp, expr.Inputs[0]) != ReadF32(fp, expr.Inputs[1]))
	WriteBool(GetOffset_bool(fp, expr.Outputs[0]), outV0)
}

// The built-in rand function returns a pseudo-random number in [0.0,1.0) from the default Source
func opF32Rand(expr *CXExpression, fp int) {
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), rand.Float32())
}

// The built-in acos function returns the arc cosine of the operand.
func opF32Acos(expr *CXExpression, fp int) {
	outV0 := float32(math.Acos(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in cos function returns the cosine of the operand.
func opF32Cos(expr *CXExpression, fp int) {
	outV0 := float32(math.Cos(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in asin function returns the arc sine of the operand.
func opF32Asin(expr *CXExpression, fp int) {
	outV0 := float32(math.Asin(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in sin function returns the sine of the operand.
func opF32Sin(expr *CXExpression, fp int) {
	outV0 := float32(math.Sin(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in sqrt function returns the square root of the operand.
func opF32Sqrt(expr *CXExpression, fp int) {
	outV0 := float32(math.Sqrt(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in log function returns the natural logarithm of the operand.
func opF32Log(expr *CXExpression, fp int) {
	outV0 := float32(math.Log(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in log2 function returns the 2-logarithm of the operand.
func opF32Log2(expr *CXExpression, fp int) {
	outV0 := float32(math.Log2(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in log10 function returns the 10-logarithm of the operand.
func opF32Log10(expr *CXExpression, fp int) {
	outV0 := float32(math.Log10(float64(ReadF32(fp, expr.Inputs[0]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in max function returns the largest value of the two operands.
func opF32Max(expr *CXExpression, fp int) {
	outV0 := float32(math.Max(float64(ReadF32(fp, expr.Inputs[0])), float64(ReadF32(fp, expr.Inputs[1]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}

// The built-in min function returns the smallest value of the two operands.
func opF32Min(expr *CXExpression, fp int) {
	outV0 := float32(math.Min(float64(ReadF32(fp, expr.Inputs[0])), float64(ReadF32(fp, expr.Inputs[1]))))
	WriteF32(GetOffset_f32(fp, expr.Outputs[0]), outV0)
}
