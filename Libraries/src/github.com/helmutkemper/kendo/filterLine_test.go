package kendo

import (
  "fmt"
)

func ExampleFilterLine_String() {
  f := FilterLine{
    Field: "field_name",
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  // { field: "field_name" }
}

func ExampleFilterLine_String2() {
  f := FilterLine{
    Operator: OPERATOR_CONTAINS,
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  // { operator: "contains" }
}

func ExampleFilterLine_String3() {
  f := FilterLine{
    Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  // { value: new Date( 1980, 1, 1 ) }
}

func ExampleFilterLine_String4() {
  f := FilterLine{
    Field: "field_name",
    Operator: OPERATOR_CONTAINS,
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  // { field: "field_name", operator: "contains" }
}

func ExampleFilterLine_String5() {
  f := FilterLine{
    Operator: OPERATOR_CONTAINS,
    Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  // { operator: "contains", value: new Date( 1980, 1, 1 ) }
}

func ExampleFilterLine_String6() {
  f := FilterLine{
    Field: "field_name",
    Operator: OPERATOR_CONTAINS,
    Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  // { field: "field_name", operator: "contains", value: new Date( 1980, 1, 1 ) }
}