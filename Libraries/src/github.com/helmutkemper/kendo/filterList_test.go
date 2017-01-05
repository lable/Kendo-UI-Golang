package kendo

import (
  "fmt"
)

func ExampleFilterLine_String() {
  var t Template = Template{}

  f := FilterLine{
    Field: "field_name",
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  //
}
