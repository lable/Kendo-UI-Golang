package kendo

import "fmt"

func ExampleLogicEnum_String() {
  var l LogicEnum = LOGIC_AND
  fmt.Print( l.String() )

  // Output:
  // and
}

func ExampleLogicEnum_String2() {
  var l LogicEnum = LOGIC_OR
  fmt.Print( l.String() )

  // Output:
  // or
}

func ExampleLogicEnum_String3() {
  var l LogicEnum = LOGIC_NIL
  fmt.Print( l.String() )

  // Output:
  //
}