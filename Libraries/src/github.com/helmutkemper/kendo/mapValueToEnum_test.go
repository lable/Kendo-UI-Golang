package kendo

import "fmt"

func ExampleMapValueToEnum_String1() {
  e := MAP_VALUE_TO_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleMapValueToEnum_String2() {
  e := MAP_VALUE_TO_DATA_ITEM

  fmt.Print( e.String() )

  // Output:
  // dataItem
}

func ExampleMapValueToEnum_String3() {
  e := MAP_VALUE_TO_INDEX

  fmt.Print( e.String() )

  // Output:
  // index
}
