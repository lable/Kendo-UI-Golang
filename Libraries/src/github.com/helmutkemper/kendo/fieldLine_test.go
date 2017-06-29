package kendo

import "fmt"

func ExampleFieldLine_String() {

  e := FieldLine{
    Field: "name",
    Type: "string",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // 'name': { type: 'string' }
}
