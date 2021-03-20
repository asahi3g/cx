package main

func runTestCases(t *tester) {
	// tests
	t.Run("test-i8.cx", CxSuccess, "i32")
	t.Run("test-i16.cx", CxSuccess, "i32")
	t.Run("test-i32.cx", CxSuccess, "i32")
	t.Run("test-i64.cx", CxSuccess, "i64")
	t.Run("test-ui8.cx", CxSuccess, "i32")
	t.Run("test-ui16.cx", CxSuccess, "i32")
	t.Run("test-ui32.cx", CxSuccess, "i32")
	t.Run("test-ui64.cx", CxSuccess, "i64")
	t.Run("test-f32.cx", CxSuccess, "f32")
	t.Run("test-f64.cx", CxSuccess, "f64")
	t.Run("test-bool.cx", CxSuccess, "bool")
	t.Run("test-array.cx", CxSuccess, "array")
	t.Run("test-function.cx", CxSuccess, "function")
	t.Run("test-control-flow.cx", CxSuccess, "control floow")
	t.Run("test-utils.cx test-struct.cx", CxSuccess, "struct")
	t.Run("test-str.cx", CxSuccess, "str")
	t.Run("test-utils.cx test-pointers.cx", CxSuccess, "pointers")
	t.Run("test-slices.cx", CxSuccess, "slices")

	t.Run("--cxpath test-workspace test-workspace-a.cx", CxSuccess, "Testing if CX can set a workspace and then import a library, taking that workspace as the new relative path.")
	t.Run("--cxpath test-workspace test-workspace-b.cx", CxSuccess, "Testing if CX can set a workspace and then import a nested library, taking that workspace as the new relative path.")
	t.Run("--cxpath test-workspace test-workspace-c.cx test-workspace-d.cx", CxSuccess, "Testing if files supplied to the CLI override libraries in the workspace.")
	t.Run("test-slices-index-out-of-range-a.cx", CxRuntimeSliceIndexOutOfRange, "Test index < 0")
	t.Run("test-slices-index-out-of-range-b.cx", CxRuntimeSliceIndexOutOfRange, "Test index >= len")
	t.Run("test-slices-resize-out-of-range-a.cx", CxRuntimeSliceIndexOutOfRange, "Test out of range after resize")
	t.Run("test-slices-resize-out-of-range-b.cx", CxRuntimeSliceIndexOutOfRange, "Test resize with count < 0")
	t.Run("test-slices-insert-out-of-range-a.cx", CxRuntimeSliceIndexOutOfRange, "Test insert with index > len")
	t.Run("test-slices-insert-out-of-range-b.cx", CxRuntimeSliceIndexOutOfRange, "Test insert with index < 0")
	t.Run("test-slices-remove-out-of-range-a.cx", CxRuntimeSliceIndexOutOfRange, "Test remove with index < 0")
	t.Run("test-slices-remove-out-of-range-b.cx", CxRuntimeSliceIndexOutOfRange, "Test remove with index >= len")
	t.Run("test-slices-remove-out-of-range-c.cx", CxRuntimeSliceIndexOutOfRange, "Test remove with index == 0 && len == 0")
	t.Run("test-short-declarations.cx", CxSuccess, "short declarations")
	t.Run("test-parse.cx", CxSuccess, "parse")
	t.Run("test-collection-functions.cx", CxSuccess, "collection functions")
	t.Run("test-scopes.cx", CxSuccess, "Error in scopes.")
	t.RunEx("-heap-initial 0 test-gc.cx", CxSuccess, "Stress-testing the garbage collector", TestIssue, 0)
	t.Run("../lib/json.cx test-json.cx", CxSuccess, "Error in json lib.")
	t.Run("../lib/args.cx test-args.cx", CxSuccess, "Error in args lib.")
	t.Run("test-regexp-must-compile-fail.cx", CxRuntimeError, "Error in regexp lib - MustCompile should have thrown an error.")
	t.Run("test-regexp-compile-fail.cx", CxSuccess, "Error in regexp lib - error thrown by regexp.Compile does not matches expected error.")
	t.Run("test-regexp.cx", CxSuccess, "Error in regexp lib.")
	t.Run("test-cipher.cx", CxSuccess, "Error in cipher lib.")
	// t.RunEx("test-regexp.cx", CxCompilationError, "Panic when calling gl.BindBuffer with only one argument.", TestGui|TestStable, 0)

	// issues
	t.Run("issue-207.cx", CxCompilationError, "Type casting error not reported.")
	t.RunEx("issue-208.cx", CxCompilationError, "Panic if return value is not used.", TestGui|TestStable, 0)
	t.Run("issue-214.cx", CxSuccess, "String not working across packages")
	t.Run("issue-215.cx issue-215a.cx", CxSuccess, "Order of files matters for structs")
	t.Run("issue-215a.cx issue-215.cx", CxSuccess, "Order of files matters for structs")
	t.RunEx("issue-216.cx", CxCompilationError, "Panic when calling gl.BindBuffer with only one argument.", TestGui|TestStable, 0)
	t.RunEx("issue-217.cx", CxSuccess, "Panic when giving []f32 argument to gl.BufferData", TestGui, 0)
	t.Run("issue-218.cx", CxSuccess, "Struct field crushed")
	t.Run("issue-219.cx", CxSuccess, "Failed to modify value in an array")
	t.Run("issue-220.cx", CxSuccess, "Panic when trying to index (using a var) an array, member of a struct passed as a function argument")
	t.Run("issue-27.cx", CxSuccess, "Failed to use shorthand operator-assign (+=, etc.) for arithmetic statements")
	t.Run("issue-221.cx", CxSuccess, "Can't call method from package")
	t.Run("issue-222.cx", CxSuccess, "Can't call method if it has a parameter")
	t.Run("issue-223.cx", CxSuccess, "Panic when using arithmetic to index an array field of a struct")
	t.Run("issue-224.cx", CxSuccess, "Panic if return value is used in an expression")
	t.Run("issue-225.cx", CxSuccess, "Using a variable to store the return boolean value of a function doesnt work with an if statement")
	t.Run("issue-226.cx", CxSuccess, "Panic when accessing property of struct array passed in as argument to func")
	t.Run("issue-227.cx", CxSuccess, "Unexpected results when accessing arrays of structs in a struct")
	t.Run("issue-230.cx", CxSuccess, "Inline initializations and arrays")
	t.Run("issue-231.cx", CxSuccess, "Slice keeps growing though it's cleared inside the loop")
	t.Run("issue-232.cx", CxSuccess, "Scope not working in loops")
	t.Run("issue-233.cx", CxSuccess, "Interdependant Structs")
	t.Run("issue-234.cx", CxCompilationError, "Panic when trying to access an invalid field.")
	t.Run("issue-235.cx", CxCompilationError, "No compilation error when using an using an invalid identifier")
	t.Run("issue-236a.cx issue-236.cx", CxSuccess, "Silent name clash between packages")
	t.Run("issue-236.cx issue-236a.cx", CxSuccess, "Silent name clash between packages")
	t.Run("issue-237.cx", CxCompilationError, "Invalid implicit cast.")
	t.Run("issue-238.cx", CxCompilationError, "Panic when using +* in an expression")
	t.Run("issue-239.cx", CxCompilationError, "No compilation error when defining a struct with duplicate fields.")
	t.Run("issue-240.cx", CxSuccess, "Can't define struct with a single character identifier.")
	t.Run("issue-241.cx", CxSuccess, "Panic when variable used in if statement without parenthesis.")
	t.Run("issue-242.cx", CxSuccess, "Struct field stomped")
	t.Run("issue-243.cx", CxCompilationError, "No compilation error when indexing an array with a non integral var.")
	t.Run("issue-244a.cx", CxSuccess, "Panic when a field of a struct returned by a function is used in an expression")
	t.Run("issue-244b.cx", CxSuccess, "Panic when a field of a struct returned by a function is used in an expression")
	t.Run("issue-245a.cx issue-245.cx", CxCompilationError, "No compilation error when using var without package qualification.")
	t.Run("issue-246.cx", CxSuccess, "No compilation error when passing *i32 as an i32 arg and conversely")
	t.Run("issue-246a.cx", CxCompilationError, "No compilation error when passing *i32 as an i32 arg and conversely")
	t.Run("issue-247.cx", CxCompilationError, "No compilation error when dereferencing an i32 var.")
	t.Run("issue-248.cx", CxSuccess, "Wrong pointer behaviour.")
	t.Run("issue-249.cx", CxSuccess, "Return from a function doesnt work")
	t.Run("issue-249b.cx", CxCompilationError, "Mismatched number of returning arguments is not throwing an error")
	t.Run("issue-250.cx", CxCompilationError, "No compilation error when var is accessed outside of its declaring scope")
	t.Run("issue-251.cx", CxCompilationError, "Panic when a str var is shadowed by a struct var in another scope")
	t.RunEx("issue-252.cx", CxSuccess, "glfw.GetCursorPos() throws error", TestGui|TestStable, 0)
	t.Run("issue-253.cx", CxSuccess, "Inline field and index 'dereferences' to function calls' outputs")
	t.Run("issue-254.cx", CxCompilationError, "No compilation error when redeclaring a variable")
	t.Run("issue-255.cx", CxSuccess, "Multi-dimensional slices don't work")
	t.Run("issue-256.cx", CxSuccess, "can't prefix a (f32) variable with minus to flip it's signedness")
	t.Run("issue-257.cx", CxCompilationError, "Using int literal 0 where 0.0 was needed gave no error")
	t.Run("issue-258.cx", CxSuccess, "error with sending references of structs to functions")
	t.Run("issue-258b.cx", CxSuccess, "error with references to struct literals")
	t.Run("issue-259.cx", CxCompilationError, "struct identifier (when initializing fields) can be with or without a '&' prefix, with no CX error")
	t.Run("issue-260.cx", CxCompilationError, "can assign to previously undeclared vars with just '='")
	t.Run("issue-261.cx", CxSuccess, "empty code blocks (even if they contain commented-out lines) crash like this")
	t.Run("issue-262.cx", CxSuccess, "increment operator ++ does not work")
	t.Run("issue-263.cx", CxSuccess, "Method does not work")
	t.Run("issue-264.cx", CxSuccess, "Cannot use bool variable in if expression")
	t.Run("issue-265.cx", CxSuccess, "CX Parser does not recognize method")
	t.Run("issue-266.cx", CxSuccess, "Goto not working on windows")
	t.Run("issue-267.cx", CxSuccess, "Methods with pointer receivers don't work")
	t.RunEx("issue-268.cx", CxSuccess, "when using 2 f32 out parameters, only the value of the 2nd gets through", TestGui, 0)
	t.Run("issue-269.cx", CxCompilationError, "Variable redeclaration should not be allowed")
	t.Run("issue-270.cx", CxSuccess, "Short variable declarations are not working with calls to methods or functions")
	t.Run("issue-271.cx", CxCompilationError, "Panic when using equality operator between a bool and an i32")
	t.Run("issue-272.cx", CxSuccess, "String concatenation using the + operator doesn't work")
	t.Run("issue-273.cx", CxSuccess, "Argument list is not parsed correctly")
	t.Run("issue-274.cx", CxSuccess, "Dubious error message when indexing an array with a substraction expression")
	t.Run("issue-275.cx", CxSuccess, "Dubious error message when inline initializing a slice")
	t.Run("issue-276a.cx issue-276.cx", CxSuccess, "Troubles when accessing a global var from another package")
	t.Run("issue-277.cx", CxSuccess, "same func names (but in different packages) collide")
	t.Run("issue-278.cx", CxCompilationError, "can use vars from other packages without a 'packageName.' prefix")
	t.Run("issue-279.cx", CxSuccess, "False positive when detecting variable redeclaration.")
	t.RunEx("issue-279a.cx", CxSuccess, "False positive when detecting variable redeclaration.", TestIssue, 0)
	t.Run("issue-280.cx", CxSuccess, "Problem with struct literals in short variable declarations")
	t.Run("issue-281.cx", CxSuccess, "Panic when using the return value of a function in a short declaration")
	t.Run("issue-282a.cx", CxCompilationError, "Panic when inserting a new line in a string literal")
	t.Run("issue-282b.cx", CxSuccess, "Panic when inserting a new line in a string literal")
	t.Run("issue-283.cx", CxCompilationError, "Panic when declaring a variable of an unknown type")
	t.Run("issue-284.cx", CxCompilationError, "No compilation error when using arithmetic operators on struct instances")
	t.RunEx("issue-285.cx", CxSuccess, "Parser gets confused with `2 -2`", TestStable, 0)
	t.Run("issue-286.cx", CxSuccess, "Panic in when assigning an empty initializer list to a []i32 variable")
	t.Run("issue-287.cx", CxSuccess, "Cx stack overflow when appending to a slice passed by address")
	t.Run("issue-288.cx", CxCompilationError, "Panic when trying to assign return value of a function returning void")
	t.Run("issue-289.cx", CxCompilationError, "Panic when using a function declared in another package without importing the package")
	t.Run("issue-290.cx", CxSuccess, "Cx memory stomped")
	t.Run("issue-291.cx", CxSuccess, "Invalid offset calculation of non literal strings when appended to a slice")
	t.Run("issue-292.cx", CxCompilationError, "Panic when calling a function from another package where the package name alias a local variable name")
	t.Run("issue-293.cx", CxSuccess, "Garbage memory when passing the address of slice element to a function")
	t.Run("issue-294.cx", CxSuccess, "Type deduction of struct field fails")
	t.Run("issue-295.cx", CxCompilationError, "No compilation error when assigning a i32 value to a []i32 variable")
	t.Run("issue-296.cx", CxCompilationError, "No compilation error when comparing value of different types")
	t.Run("-stack-size 30 issue-297a.cx", CxRuntimeStackOverflowError, "No stack overflow error")
	t.Run("-heap-initial 100 -heap-max 110 issue-297b.cx", CxRuntimeHeapExhaustedError, "No heap exhausted error")
	t.Run("issue-298.cx", CxSuccess, "Argument type deduction failed when passing address of an i32 struct field to a function accepting *i32 argument.")
	t.Run("issue-299.cx", CxCompilationError, "Type checking is not working with receiving variables of unexpected types")
	t.Run("issue-300.cx", CxSuccess, "Crash when using a constant expression in a slice literal expression")
	t.Run("issue-301-a.cx", CxCompilationError, "Can redeclare variables if they are inline initialized")
	t.Run("issue-301-b.cx", CxCompilationError, "Can redeclare variables if they are inline initialized")
	t.Run("issue-302.cx", CxSuccess, "Trying to determine the length of a slice of struct instances throws an error.")
	t.RunEx("issue-68.cx", CxSuccess, "Wrong sprintf behaviour when passing increment expression as argument", TestIssue, 0)
	t.RunEx("issue-67.cx", CxCompilationError, "Panic when using void return value of a function in a for loop expression", TestIssue, 0)
	t.RunEx("issue-66.cx", CxSuccess, "Wrong sprintf behaviour when printing boolean values with %v", TestIssue, 0)
	t.RunEx("issue-65.cx", CxSuccess, "for true {} loop scope is not executed", TestIssue, 0)
	t.RunEx("issue-64.cx", CxSuccess, "for loop using boolean value is not compiling", TestIssue, 0)
	t.Run("issue-303.cx", CxSuccess, "Concatenation of str variables with + operator doesn't work")
	t.Run("issue-304.cx", CxSuccess, "Short declaration doesn't compile with opcode return value")
	t.Run("issue-305.cx", CxSuccess, "Compilation error when struct field is named 'input' or 'output'")
	t.RunEx("issue-63.cx", CxCompilationError, "No compilation error when using empty argument list after function call", TestIssue, 0)
	t.Run("issue-306.cx", CxCompilationError, "No compilation error when using float value in place of boolean expression")
	t.Run("issue-308.cx", CxCompilationError, "Panic when package contains duplicate function signature")
	t.RunEx("issue-62.cx", CxCompilationError, "Left hand side of , is not compiled", TestIssue, 0)
	t.RunEx("issue-61-a.cx", CxSuccess, "Compilation error when using return value of a member method in a expression", TestIssue, 0)
	t.RunEx("issue-61-b.cx", CxSuccess, "Compilation error when using return value of a member method in a expression", TestIssue, 0)
	t.Run("issue-309.cx", CxSuccess, "Compilation error when left hand side of an assignment expression is a struct field")
	t.RunEx("issue-60.cx", CxSuccess, "Crash when INIT_HEAP_SIZE limit is reached ", TestIssue, 0)
	t.RunEx("issue-59-a.cx", CxSuccess, "Crash in garbage collector when heap is resized", TestIssue, 0)
	t.RunEx("issue-59-b.cx", CxSuccess, "Crash in garbage collector when heap is resized", TestIssue, 0)
	t.RunEx("issue-53-a.cx", CxSuccess, "Issues with slice of type T where sizeof T is different than 4 ", TestStable, 0)
	t.RunEx("issue-53-b.cx", CxSuccess, "Issues with slice of type T where sizeof T is different than 4 ", TestStable, 0)
	t.RunEx("issue-53-c.cx", CxSuccess, "Issues with slice of type T where sizeof T is different than 4 ", TestStable, 0)
	t.RunEx("issue-51.cx", CxCompilationError, "No compilation error when global variable is redeclared at local scope", TestIssue, 0)
	t.RunEx("issue-49.cx", CxCompilationError, "No compilation error when assigning an literal which overflow the receiving type", TestIssue, 0)
	t.RunEx("issue-48.cx", CxSuccess, "Cx is not supporting short-circuit evaluation", TestIssue, 0)
	t.RunEx("issue-39.cx", CxSuccess, "func defined with no arguments, called WITH arguments causes PANIC", TestIssue, 0)
	t.Run("issue-2.cx", CxSuccess, "multi-dimensional arrays are not working")
	t.Run("issue-1.cx", CxSuccess, "multi-dimensional slices are not working")
	t.RunEx("issue-120-a.cx", CxCompilationError, "Invalid implicit cast when assigning the result of a math operator to a variable.", TestIssue, 0)
	t.RunEx("issue-120-b.cx", CxCompilationError, "Invalid implicit cast when assigning the result of a math operator to a variable.", TestIssue, 0)
	t.RunEx("issue-120-c.cx", CxCompilationError, "Invalid implicit cast when assigning the result of a math operator to a variable.", TestIssue, 0)
	t.RunEx("issue-121.cx", CxSuccess, "Compilation error when using unary negative operator on a function call", TestIssue, 0)
	t.RunEx("issue-131.cx", CxSuccess, "Panic when using arithmetic operations.", TestIssue, 0)
	t.RunEx("test-ar-1.cx", CxSuccess, "Panic when using string to pointer array.", TestIssue, 0)
	t.RunEx("issue-157.cx", CxSuccess, "expected either 'i32' or 'i64', got 'ident'", TestIssue, 0)

	// We need to fix serialization and deserialization as user-callable functions
	// t.RunEx("issue-309.cx", CxSuccess, "Serialization is not taking into account non-default stack sizes.", TestIssue, 0)
	// t.RunEx("issue-310.cx", CxSuccess, "Splitting a serialized program into its blockchain and transaction parts.", TestIssue, 0)
	// t.RunEx("issue-311.cx", CxSuccess, "`CurrentFunction` and `CurrentStruct` are causing errors in programs with more than 1 package.", TestIssue, 0)
	// t.RunEx("issue-312.cx", CxSuccess, "Deserialization is not setting correctly the sizes for the CallStack, HeapStartsAt and StackSize fields of the CXProgram structure.", TestIssue, 0)

	t.Run("issue-28-f32.cx", CxSuccess, "f32")
	t.Run("issue-28-f64.cx", CxSuccess, "f64")
	t.Run("issue-28-i8.cx", CxSuccess, "i8")
	t.Run("issue-28-i16.cx", CxSuccess, "i16")
	t.Run("issue-28-i32.cx", CxSuccess, "i32")
	t.Run("issue-28-i64.cx", CxSuccess, "i64")
	t.Run("issue-28-ui8.cx", CxSuccess, "ui8")
	t.Run("issue-28-ui16.cx", CxSuccess, "ui16")
	t.Run("issue-28-ui32.cx", CxSuccess, "ui32")
	t.Run("issue-28-ui64.cx", CxSuccess, "ui64")
}
