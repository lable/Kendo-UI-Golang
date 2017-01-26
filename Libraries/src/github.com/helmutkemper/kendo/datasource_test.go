package kendo

import "fmt"

func ExampleDataSource_String() {
  e := DataSource{
    AggregateLine: []AggregateLine{
      {
        Field: "field_1",
        AggregateEnum: AGGREGATE_COUNT,
      },
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}
