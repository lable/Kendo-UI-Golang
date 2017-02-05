package kendo

import "fmt"

func ExampleTypeDataEnum_String1() {
  e := TYPE_DATA_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleTypeDataEnum_String2() {
  e := TYPE_DATA_XML

  fmt.Print( e.String() )

  // Output:
  // xml
}

func ExampleTypeDataEnum_String3() {
  e := TYPE_DATA_JSON

  fmt.Print( e.String() )

  // Output:
  // json
}
