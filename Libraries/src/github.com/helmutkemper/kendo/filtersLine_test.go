package kendo

import "fmt"

func ExampleFiltersLine_String1() {
  e := FiltersLine{
    Field: "field_2",
    Logic: LOGIC_AND,
    Operator: OPERATOR_EQ,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // field: 'field_2',
  // logic: 'and',
  // operator: 'eq',
}
