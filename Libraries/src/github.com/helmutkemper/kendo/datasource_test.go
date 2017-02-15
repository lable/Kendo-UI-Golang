package kendo

import "fmt"

func ExampleDataSource_String() {
  e := DataSource{
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}

func ExampleDataSource_String1() {
  e := DataSource{
    AggregateLine: []AggregateLine{
      {
        Field: "field_1",
        AggregateEnum: AGGREGATE_COUNT,
        GoTemplate: &t,
      },
      {
        Field: "field_1",
        AggregateEnum: AGGREGATE_COUNT,
        GoTemplate: &t,
      },
    },
    ServerAggregates: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // aggregate: [ aggregate: 'count',field: 'field_1',aggregate: 'count',field: 'field_1', ], serverAggregates: true,
}

func ExampleDataSource_String2() {
  e := DataSource{
    AggregateLine: []AggregateLine{
      {
        Field: "field_1",
        AggregateEnum: AGGREGATE_COUNT,
        GoTemplate: &t,
      },
    },
    ServerAggregates: false,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}

func ExampleDataSource_String3() {
  e := DataSource{
    AutoSync: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // autoSync: true,
}

func ExampleDataSource_String4() {
  e := DataSource{
    Batch: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // batch: true,
}

func ExampleDataSource_String5() {
  e := DataSource{
    Data: ComplexJavaScriptType{
      AsFunction: "function ( e ){ trace.log( e ); }",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // data: function ( e ){ trace.log( e ); },
}


func ExampleDataSource_String6() {
  e := DataSource{
    FilterLine: []FilterLine{
      {
        Field: "field_1",
        Logic: LOGIC_AND,
        Filters: []FiltersLine{
          {
            Field: "field_2",
            Logic: LOGIC_AND,
            Operator: OPERATOR_EQ,
            GoTemplate: &t,
          },
        },
        Operator: OPERATOR_CONTAINS,
        Value: ComplexJavaScriptType{
          AsFunction: "new Date(1980, 1, 1)",
        },
        GoTemplate: &t,
      },
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // field: 'field_1',filters: [ { field: 'field_2',logic: 'and',operator: 'eq', }, ],logic: 'and',operator: 'contains',value: new Date(1980, 1, 1),
}
