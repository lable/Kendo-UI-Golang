package kendo

import "fmt"

func ExampleAutoCompleteFilterEnum_String1() {
  e := AUTOCOMPLETE_FILTER_CONTAINS

  fmt.Print( e.String() )

  // Output:
  // contains
}

func ExampleAutoCompleteFilterEnum_String2() {
  e := AUTOCOMPLETE_FILTER_ENDSWITH

  fmt.Print( e.String() )

  // Output:
  // endswith
}

func ExampleAutoCompleteFilterEnum_String3() {
  e := AUTOCOMPLETE_FILTER_STARTSWITH

  fmt.Print( e.String() )

  // Output:
  // startswith
}

func ExampleAutoCompleteFilterEnum_String4() {
  e := AUTOCOMPLETE_FILTER_NIL

  fmt.Print( e.String() )

  // Output:
  //
}