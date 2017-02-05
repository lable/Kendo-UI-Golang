package kendo

import "fmt"

func ExampleAggregateEnum_String1() {
  e := AGGREGATE_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleAggregateEnum_String2() {
  e := AGGREGATE_AVERAGE

  fmt.Print( e.String() )

  // Output:
  // average
}

func ExampleAggregateEnum_String3() {
  e := AGGREGATE_COUNT

  fmt.Print( e.String() )

  // Output:
  // count
}

func ExampleAggregateEnum_String4() {
  e := AGGREGATE_MAX

  fmt.Print( e.String() )

  // Output:
  // max
}

func ExampleAggregateEnum_String5() {
  e := AGGREGATE_MIN

  fmt.Print( e.String() )

  // Output:
  // min
}

func ExampleAggregateEnum_String6() {
  e := AGGREGATE_SUM

  fmt.Print( e.String() )

  // Output:
  // sum
}
