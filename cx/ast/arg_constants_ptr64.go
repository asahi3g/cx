// +build ptr64

package ast

import "github.com/skycoin/cx/cx/constants"

// Pointer type, CXArg Constants
var ConstCxArg_PTR = NewCXArgument(constants.TYPE_I32) // TODO: PTR use right type when we have ptr alias in cx
