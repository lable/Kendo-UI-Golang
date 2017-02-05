package kendo

import "fmt"

func ExampleTypeDataJSonEnum_String1() {
  e := TYPE_DATA_JSON_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleTypeDataJSonEnum_String2() {
  e := TYPE_DATA_JSON_JSONP

  fmt.Print( e.String() )

  // Output:
  // jsonp
}

func ExampleTypeDataJSonEnum_String3() {
  e := TYPE_DATA_JSON_JSON

  fmt.Print( e.String() )

  // Output:
  // json
}
