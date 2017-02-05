package kendo

import "fmt"

func ExampleMethodEnum_String1() {
  e := METHOD_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleMethodEnum_String2() {
  e := METHOD_DELETE

  fmt.Print( e.String() )

  // Output:
  // DELETE
}

func ExampleMethodEnum_String3() {
  e := METHOD_PUT

  fmt.Print( e.String() )

  // Output:
  // PUT
}

func ExampleMethodEnum_String4() {
  e := METHOD_GET

  fmt.Print( e.String() )

  // Output:
  // GET
}

func ExampleMethodEnum_String5() {
  e := METHOD_POST

  fmt.Print( e.String() )

  // Output:
  // POST
}
