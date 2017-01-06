package kendo

import "fmt"

func ExampleOperatorEnum_String() {
  var o OperatorEnum = OPERATOR_EQ
  fmt.Print( o.String() )

  // Output:
  // eq
}

func ExampleOperatorEnum_String2() {
  var o OperatorEnum = OPERATOR_NEQ
  fmt.Print( o.String() )

  // Output:
  // neq
}

func ExampleOperatorEnum_String3() {
  var o OperatorEnum = OPERATOR_ISNULL
  fmt.Print( o.String() )

  // Output:
  // isnull
}

func ExampleOperatorEnum_String4() {
  var o OperatorEnum = OPERATOR_ISNOTNULL
  fmt.Print( o.String() )

  // Output:
  // isnotnull
}

func ExampleOperatorEnum_String5() {
  var o OperatorEnum = OPERATOR_LT
  fmt.Print( o.String() )

  // Output:
  // lt
}

func ExampleOperatorEnum_String6() {
  var o OperatorEnum = OPERATOR_LTE
  fmt.Print( o.String() )

  // Output:
  // lte
}

func ExampleOperatorEnum_String7() {
  var o OperatorEnum = OPERATOR_GT
  fmt.Print( o.String() )

  // Output:
  // gt
}

func ExampleOperatorEnum_String8() {
  var o OperatorEnum = OPERATOR_GTE
  fmt.Print( o.String() )

  // Output:
  // gte
}

func ExampleOperatorEnum_String9() {
  var o OperatorEnum = OPERATOR_STARTSWITH
  fmt.Print( o.String() )

  // Output:
  // startswith
}

func ExampleOperatorEnum_String10() {
  var o OperatorEnum = OPERATOR_ENDSWITH
  fmt.Print( o.String() )

  // Output:
  // endswith
}

func ExampleOperatorEnum_String11() {
  var o OperatorEnum = OPERATOR_CONTAINS
  fmt.Print( o.String() )

  // Output:
  // contains
}

func ExampleOperatorEnum_String12() {
  var o OperatorEnum = OPERATOR_DOESNOTCONTAIN
  fmt.Print( o.String() )

  // Output:
  // doesnotcontain
}

func ExampleOperatorEnum_String13() {
  var o OperatorEnum = OPERATOR_ISEMPTY
  fmt.Print( o.String() )

  // Output:
  // isempty
}

func ExampleOperatorEnum_String14() {
  var o OperatorEnum = OPERATOR_ISNOTEMPTY
  fmt.Print( o.String() )

  // Output:
  // isnotempty
}
