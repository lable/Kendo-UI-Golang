package kendo

import "fmt"

func ExampleMethodEnum_String() {
  var m MethodEnum = METHOD_POST
  fmt.Print( m.String() )

  // Output:
  // POST
}

func ExampleMethodEnum_String2() {
  var m MethodEnum = METHOD_GET
  fmt.Print( m.String() )

  // Output:
  // GET
}

func ExampleMethodEnum_String3() {
  var m MethodEnum = METHOD_PUT
  fmt.Print( m.String() )

  // Output:
  // PUT
}

func ExampleMethodEnum_String4() {
  var m MethodEnum = METHOD_DELETE
  fmt.Print( m.String() )

  // Output:
  // DELETE
}