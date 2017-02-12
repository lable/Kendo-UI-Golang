package kendo

import "fmt"

func ExampleTransport_String1() {
  e := Transport{
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}
