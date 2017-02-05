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
        GoTemplate: &t,
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
  // field: 'field_1',
  // filters: [ { field: 'field_2',
  // logic: 'and',
  // operator: 'eq',
  //
  //  }, ],
  // logic: 'and',
  // operator: 'contains',
  // value: new Date(1980, 1, 1),
}
