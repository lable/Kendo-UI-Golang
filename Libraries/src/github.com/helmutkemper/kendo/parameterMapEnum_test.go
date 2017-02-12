package kendo

import "fmt"

func ExampleParameterMapEnum_String1() {
  e := PARAMETERMAP_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleParameterMapEnum_String2() {
  e := PARAMETERMAP_CREATE

  fmt.Print( e.String() )

  // Output:
  // create
}

func ExampleParameterMapEnum_String3() {
  e := PARAMETERMAP_READ

  fmt.Print( e.String() )

  // Output:
  // read
}

func ExampleParameterMapEnum_String4() {
  e := PARAMETERMAP_UPDATE

  fmt.Print( e.String() )

  // Output:
  // update
}

func ExampleParameterMapEnum_String5() {
  e := PARAMETERMAP_DESTROY

  fmt.Print( e.String() )

  // Output:
  // destroy
}
