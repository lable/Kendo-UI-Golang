package kendo

import "fmt"

func ExampleGroupLine_String1() {
  e := GroupLine{
    Aggregates: []AggregateLine{
      {
        AggregateEnum: AGGREGATE_AVERAGE,
        Field: "field_1",
        GoTemplate: &t,
      },
      {
        AggregateEnum: AGGREGATE_COUNT,
        Field: "field_2",
        GoTemplate: &t,
      },
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // aggregates: [{ aggregate: 'average',field: 'field_1' },{ aggregate: 'count',field: 'field_2' },],
}

func ExampleGroupLine_String2() {
  e := GroupLine{
    Dir: DIR_ASC,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // dir: 'asc',
}

func ExampleGroupLine_String3() {
  e := GroupLine{
    Field: "field_1",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // field: 'field_1',
}


func ExampleGroupLine_String4() {
  e := GroupLine{
    Aggregates: []AggregateLine{
      {
        AggregateEnum: AGGREGATE_AVERAGE,
        Field: "field_1",
        GoTemplate: &t,
      },
      {
        AggregateEnum: AGGREGATE_COUNT,
        Field: "field_2",
        GoTemplate: &t,
      },
    },
    Dir: DIR_ASC,
    Field: "field_1",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // aggregates: [{ aggregate: 'average',field: 'field_1' },{ aggregate: 'count',field: 'field_2' },],dir: 'asc',field: 'field_1',
}
