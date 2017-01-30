package kendo

import "fmt"

func ExampleFilterLine_String1() {
  e := FilterLine{
    Field: "field_1",

    Logic: LOGIC_AND,
    Filters: []FiltersLine{
      {
        Field: "field_2",
        Logic: LOGIC_AND,
        Operator: OPERATOR_EQ,
      },
    },
    Operator: OPERATOR_CONTAINS,
    Value: ComplexJavaScriptType{
      AsFunction: "new Date(1980, 1, 1)",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // data: function ( e ){ trace.log( e ); },
}
