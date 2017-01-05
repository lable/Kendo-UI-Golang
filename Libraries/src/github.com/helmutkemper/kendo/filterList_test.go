package kendo

import (
  "fmt"
)

var t Template

func init(){
  t = Template{}
}

func ExampleFilterLine_String() {
  f := FilterLine{
    Field: "field_name",
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  //
}

func ExampleFilterLine_String2() {
  f := FilterLine{
    Operator: FILTER_OPERATOR_CONTAINS,
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  //
}

func ExampleFilterLine_String3() {
  f := FilterLine{
    Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
    Template: &t,
  }
  fmt.Print( f.String() );

  // Output:
  //
}