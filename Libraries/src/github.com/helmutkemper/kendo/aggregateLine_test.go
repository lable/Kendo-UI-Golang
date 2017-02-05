package kendo

import "fmt"

func ExampleAggregateLine_String1() {
  e := AggregateLine{
    AggregateEnum: AGGREGATE_AVERAGE,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // aggregate: 'average',
}

func ExampleAggregateLine_String2() {
  e := AggregateLine{
    Field: "field_1",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // field: 'field_1',
}

func ExampleAggregateLine_String3() {
  e := AggregateLine{
    AggregateEnum: AGGREGATE_AVERAGE,
    Field: "field_1",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // aggregate: 'average',
  // field: 'field_1',
}
