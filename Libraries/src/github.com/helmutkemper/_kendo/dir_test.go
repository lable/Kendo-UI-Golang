package kendo

import "fmt"

func ExampleDirEnum_String() {
  var d DirEnum = DIR_ASC
  fmt.Print( d.String() )

  // Output:
  // asc
}

func ExampleDirEnum_String2() {
  var d DirEnum = DIR_DESC
  fmt.Print( d.String() )

  // Output:
  // desc
}
