package kendo

import (
  "fmt"
)

func ExampleAggregateList_String() {
  a := AggregateLine{
    Aggregate: AGGREGATE_MAX,
    Field: "field_name",
    Template: &t,
  }
  fmt.Print( a.String() )

  // Output:
  // { field: "field_name", aggregate: "max" }
}