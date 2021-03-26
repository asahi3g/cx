// CAUTION: Generated by yy - DO NOT EDIT.

// Copyright 2015 The cxgo Authors. All rights reserved.  Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// This is a derived work base on the original at
//
// http://pubs.opengroup.org/onlinepubs/009695399/utilities/yacc.html
//
// The original work is
//
// Copyright © 2001-2004 The IEEE and The Open Group, All Rights reserved.
//
// Grammar for the input to yacc.

package parser // import "github.com/skycoin/cx/cmd/goyacc/cxgo/yacc"

import (
	"fmt"
)

func ExampleAction() {
	fmt.Println(exampleAST(2, "%% a: { }"))
	// Output:
	// &cxgo.Action{
	// · Values: []*cxgo.ActionValue{ // len 1
	// · · 0: &cxgo.ActionValue{
	// · · · Pos: 7,
	// · · · Src: "{ }",
	// · · },
	// · },
	// · Token: example.y:1:7: '{' "{", Comments: [],
	// · Token2: example.y:1:9: '}' "}", Comments: [],
	// }
}

func ExampleDefinition() {
	fmt.Println(exampleAST(3, "%start a %error-verbose"))
	// Output:
	// &cxgo.Definition{
	// · Token: example.y:1:1: START "%start", Comments: [],
	// · Token2: example.y:1:8: IDENTIFIER "a", Comments: [],
	// }
}

func ExampleDefinition_case1() {
	fmt.Println(exampleAST(4, "%union{int i} %%"))
	// Output:
	// &cxgo.Definition{
	// · Value: "{int i}",
	// · Case: 1,
	// · Token: example.y:1:1: UNION "%union", Comments: [],
	// }
}

func ExampleDefinition_case2() {
	fmt.Println(exampleAST(6, "%{ %} %error-verbose"))
	// Output:
	// &cxgo.Definition{
	// · Value: " ",
	// · Case: 2,
	// · Token: example.y:1:1: LCURL "%{", Comments: [],
	// · Token2: example.y:1:4: RCURL "%}", Comments: [],
	// }
}

func ExampleDefinition_case3() {
	fmt.Println(exampleAST(7, "%left a %error-verbose"))
	// Output:
	// &cxgo.Definition{
	// · Nlist: []*cxgo.Name{ // len 1
	// · · 0: &cxgo.Name{
	// · · · Identifier: "a",
	// · · · Number: -1,
	// · · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · · },
	// · },
	// · Case: 3,
	// · NameList: &cxgo.NameList{
	// · · Name: &cxgo.Name{ /* recursive/repetitive pointee not shown */ },
	// · },
	// · ReservedWord: &cxgo.ReservedWord{
	// · · Case: 1,
	// · · Token: example.y:1:1: LEFT "%left", Comments: [],
	// · },
	// }
}

func ExampleDefinition_case4() {
	fmt.Println(exampleAST(8, "%left %error-verbose"))
	// Output:
	// &cxgo.Definition{
	// · Case: 4,
	// · ReservedWord: &cxgo.ReservedWord{
	// · · Case: 1,
	// · · Token: example.y:1:1: LEFT "%left", Comments: [],
	// · },
	// }
}

func ExampleDefinition_case5() {
	fmt.Println(exampleAST(9, "%error-verbose %error-verbose"))
	// Output:
	// &cxgo.Definition{
	// · Case: 5,
	// · Token: example.y:1:1: ERROR_VERBOSE "%error-verbose", Comments: [],
	// }
}

func ExampleDefinitionList() {
	fmt.Println(exampleAST(10, "%error-verbose") == (*DefinitionList)(nil))
	// Output:
	// true
}

func ExampleDefinitionList_case1() {
	fmt.Println(exampleAST(11, "%error-verbose %error-verbose"))
	// Output:
	// &cxgo.DefinitionList{
	// · Definition: &cxgo.Definition{
	// · · Case: 5,
	// · · Token: example.y:1:1: ERROR_VERBOSE "%error-verbose", Comments: [],
	// · },
	// }
}

func ExampleLiteralStringOpt() {
	fmt.Println(exampleAST(12, "%left a ,") == (*LiteralStringOpt)(nil))
	// Output:
	// true
}

func ExampleLiteralStringOpt_case1() {
	fmt.Println(exampleAST(13, "%left a \"b\" ,"))
	// Output:
	// &cxgo.LiteralStringOpt{
	// · Token: example.y:1:9: STRING_LITERAL "\"b\"", Comments: [],
	// }
}

func ExampleName() {
	fmt.Println(exampleAST(14, "%left a ,"))
	// Output:
	// &cxgo.Name{
	// · Identifier: "a",
	// · Number: -1,
	// · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// }
}

func ExampleName_case1() {
	fmt.Println(exampleAST(15, "%left a 98 ,"))
	// Output:
	// &cxgo.Name{
	// · Identifier: "a",
	// · Number: 98,
	// · Case: 1,
	// · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · Token2: example.y:1:9: NUMBER "98", Comments: [],
	// }
}

func ExampleNameList() {
	fmt.Println(exampleAST(16, "%left a ,"))
	// Output:
	// &cxgo.NameList{
	// · Name: &cxgo.Name{
	// · · Identifier: "a",
	// · · Number: -1,
	// · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · },
	// }
}

func ExampleNameList_case1() {
	fmt.Println(exampleAST(17, "%left a b ,"))
	// Output:
	// &cxgo.NameList{
	// · Name: &cxgo.Name{
	// · · Identifier: "a",
	// · · Number: -1,
	// · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · },
	// · NameList: &cxgo.NameList{
	// · · Case: 1,
	// · · Name: &cxgo.Name{
	// · · · Identifier: "b",
	// · · · Number: -1,
	// · · · Token: example.y:1:9: IDENTIFIER "b", Comments: [],
	// · · },
	// · },
	// }
}

func ExampleNameList_case2() {
	fmt.Println(exampleAST(18, "%left a , b ,"))
	// Output:
	// &cxgo.NameList{
	// · Name: &cxgo.Name{
	// · · Identifier: "a",
	// · · Number: -1,
	// · · Token: example.y:1:7: IDENTIFIER "a", Comments: [],
	// · },
	// · NameList: &cxgo.NameList{
	// · · Case: 2,
	// · · Name: &cxgo.Name{
	// · · · Identifier: "b",
	// · · · Number: -1,
	// · · · Token: example.y:1:11: IDENTIFIER "b", Comments: [],
	// · · },
	// · · Token: example.y:1:9: ',' ",", Comments: [],
	// · },
	// }
}

func ExamplePrecedence() {
	fmt.Println(exampleAST(19, "%% a: |") == (*Precedence)(nil))
	// Output:
	// true
}

func ExamplePrecedence_case1() {
	fmt.Println(exampleAST(20, "%% a: %prec b"))
	// Output:
	// &cxgo.Precedence{
	// · Identifier: "b",
	// · Case: 1,
	// · Token: example.y:1:7: PREC "%prec", Comments: [],
	// · Token2: example.y:1:13: IDENTIFIER "b", Comments: [],
	// }
}

func ExamplePrecedence_case2() {
	fmt.Println(exampleAST(21, "%% a: %prec b { }"))
	// Output:
	// &cxgo.Precedence{
	// · Identifier: "b",
	// · Action: &cxgo.Action{
	// · · Values: []*cxgo.ActionValue{ // len 1
	// · · · 0: &cxgo.ActionValue{
	// · · · · Pos: 15,
	// · · · · Src: "{ }",
	// · · · },
	// · · },
	// · · Token: example.y:1:15: '{' "{", Comments: [],
	// · · Token2: example.y:1:17: '}' "}", Comments: [],
	// · },
	// · Case: 2,
	// · Token: example.y:1:7: PREC "%prec", Comments: [],
	// · Token2: example.y:1:13: IDENTIFIER "b", Comments: [],
	// }
}

func ExamplePrecedence_case3() {
	fmt.Println(exampleAST(22, "%% a: ;"))
	// Output:
	// &cxgo.Precedence{
	// · Case: 3,
	// · Token: example.y:1:7: ';' ";", Comments: [],
	// }
}

func ExampleReservedWord() {
	fmt.Println(exampleAST(23, "%token <"))
	// Output:
	// &cxgo.ReservedWord{
	// · Token: example.y:1:1: TOKEN "%token", Comments: [],
	// }
}

func ExampleReservedWord_case1() {
	fmt.Println(exampleAST(24, "%left <"))
	// Output:
	// &cxgo.ReservedWord{
	// · Case: 1,
	// · Token: example.y:1:1: LEFT "%left", Comments: [],
	// }
}

func ExampleReservedWord_case2() {
	fmt.Println(exampleAST(25, "%right <"))
	// Output:
	// &cxgo.ReservedWord{
	// · Case: 2,
	// · Token: example.y:1:1: RIGHT "%right", Comments: [],
	// }
}

func ExampleReservedWord_case3() {
	fmt.Println(exampleAST(26, "%nonassoc <"))
	// Output:
	// &cxgo.ReservedWord{
	// · Case: 3,
	// · Token: example.y:1:1: NONASSOC "%nonassoc", Comments: [],
	// }
}

func ExampleReservedWord_case4() {
	fmt.Println(exampleAST(27, "%type <"))
	// Output:
	// &cxgo.ReservedWord{
	// · Case: 4,
	// · Token: example.y:1:1: TYPE "%type", Comments: [],
	// }
}

func ExampleReservedWord_case5() {
	fmt.Println(exampleAST(28, "%precedence <"))
	// Output:
	// &cxgo.ReservedWord{
	// · Case: 5,
	// · Token: example.y:1:1: PRECEDENCE "%precedence", Comments: [],
	// }
}

func ExampleRule() {
	fmt.Println(exampleAST(29, "%% a: b:"))
	// Output:
	// &cxgo.Rule{
	// · Name: example.y:1:7: C_IDENTIFIER "b", Comments: [],
	// · Token: example.y:1:7: C_IDENTIFIER "b", Comments: [],
	// }
}

func ExampleRule_case1() {
	fmt.Println(exampleAST(30, "%% a: |"))
	// Output:
	// &cxgo.Rule{
	// · Name: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · Case: 1,
	// · Token: example.y:1:7: '|' "|", Comments: [],
	// }
}

func ExampleRuleItemList() {
	fmt.Println(exampleAST(31, "%% a:") == (*RuleItemList)(nil))
	// Output:
	// true
}

func ExampleRuleItemList_case1() {
	fmt.Println(exampleAST(32, "%% a: b"))
	// Output:
	// &cxgo.RuleItemList{
	// · Case: 1,
	// · Token: example.y:1:7: IDENTIFIER "b", Comments: [],
	// }
}

func ExampleRuleItemList_case2() {
	fmt.Println(exampleAST(33, "%% a: { }"))
	// Output:
	// &cxgo.RuleItemList{
	// · Action: &cxgo.Action{
	// · · Values: []*cxgo.ActionValue{ // len 1
	// · · · 0: &cxgo.ActionValue{
	// · · · · Pos: 7,
	// · · · · Src: "{ }",
	// · · · },
	// · · },
	// · · Token: example.y:1:7: '{' "{", Comments: [],
	// · · Token2: example.y:1:9: '}' "}", Comments: [],
	// · },
	// · Case: 2,
	// }
}

func ExampleRuleItemList_case3() {
	fmt.Println(exampleAST(34, "%% a: \"b\""))
	// Output:
	// &cxgo.RuleItemList{
	// · Case: 3,
	// · Token: example.y:1:7: STRING_LITERAL "\"b\"", Comments: [],
	// }
}

func ExampleRuleList() {
	fmt.Println(exampleAST(35, "%% a:"))
	// Output:
	// &cxgo.RuleList{
	// · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// }
}

func ExampleRuleList_case1() {
	fmt.Println(exampleAST(36, "%% a: |"))
	// Output:
	// &cxgo.RuleList{
	// · RuleList: &cxgo.RuleList{
	// · · Case: 1,
	// · · Rule: &cxgo.Rule{
	// · · · Name: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · · · Case: 1,
	// · · · Token: example.y:1:7: '|' "|", Comments: [],
	// · · },
	// · },
	// · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// }
}

func ExampleSpecification() {
	fmt.Println(exampleAST(37, "%% a:"))
	// Output:
	// &cxgo.Specification{
	// · Rules: []*cxgo.Rule{ // len 1
	// · · 0: &cxgo.Rule{
	// · · · Name: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · · · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · · },
	// · },
	// · RuleList: &cxgo.RuleList{
	// · · Token: example.y:1:4: C_IDENTIFIER "a", Comments: [],
	// · },
	// · Token: example.y:1:1: MARK "%%", Comments: [],
	// }
}

func ExampleTag() {
	fmt.Println(exampleAST(38, "%left %error-verbose") == (*Tag)(nil))
	// Output:
	// true
}

func ExampleTag_case1() {
	fmt.Println(exampleAST(39, "%left < a > %error-verbose"))
	// Output:
	// &cxgo.Tag{
	// · Token: example.y:1:7: '<' "<", Comments: [],
	// · Token2: example.y:1:9: IDENTIFIER "a", Comments: [],
	// · Token3: example.y:1:11: '>' ">", Comments: [],
	// }
}

func ExampleTail() {
	fmt.Println(exampleAST(40, "%% a: %%"))
	// Output:
	// &cxgo.Tail{
	// · Token: example.y:1:7: MARK "%%", Comments: [],
	// }
}

func ExampleTail_case1() {
	fmt.Println(exampleAST(41, "%% a:") == (*Tail)(nil))
	// Output:
	// true
}