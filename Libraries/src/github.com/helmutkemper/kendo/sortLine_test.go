package kendo

import "fmt"

func ExampleSort_String() {
  var s Sort = Sort{
    Dir:  DIR_DESC,
  }
  fmt.Print( s.String() )

  // Output:
  //
}