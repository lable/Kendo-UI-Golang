package kendo

import (
  "fmt"
)

func ExampleFilter_String() {
  f := Filter{
    Field: "field_name",
    Template: &t,
  }

  fmt.Print( f.String() )

  // Output:
  // filter: { field: "field_name", },
}

func ExampleFilter_String1() {
  f := Filter{
    Field: "field_name",
    Filters: []FilterLine{
      {
        Operator: OPERATOR_CONTAINS,
        Template: &t,
      },
      {
        Field: "field_name",
        Operator: OPERATOR_CONTAINS,
        Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
        Template: &t,
      },
    },
    Template: &t,
  }

  fmt.Print( f.String() )

  // Output:
  // field: "field_name",filters: [{ operator: &#34;contains&#34; },{ field: &#34;field_name&#34;, operator: &#34;contains&#34;, value: new Date( 1980, 1, 1 ) },],
}

func ExampleFilter_String2() {
  f := Filter{
    Field: "field_name",
    Filters: []FilterLine{
      {
        Operator: OPERATOR_CONTAINS,
        Template: &t,
      },
      {
        Field: "field_name",
        Operator: OPERATOR_CONTAINS,
        Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
        Template: &t,
      },
    },
    Logic: LOGIC_AND,
    Template: &t,
  }

  fmt.Print( f.String() )

  // Output:
  // field: "field_name",filters: [{ operator: &#34;contains&#34; },{ field: &#34;field_name&#34;, operator: &#34;contains&#34;, value: new Date( 1980, 1, 1 ) },],logic: "and",
}

func ExampleFilter_String3() {
  f := Filter{
    Field: "field_name",
    Filters: []FilterLine{
      {
        Operator: OPERATOR_CONTAINS,
        Template: &t,
      },
      {
        Field: "field_name",
        Operator: OPERATOR_CONTAINS,
        Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
        Template: &t,
      },
    },
    Logic: LOGIC_AND,
    Operator: OPERATOR_EQ,
    Template: &t,
  }

  fmt.Print( f.String() )

  // Output:
  // filter: { field: "field_name",filters: [{ operator: &#34;contains&#34; },{ field: &#34;field_name&#34;, operator: &#34;contains&#34;, value: new Date( 1980, 1, 1 ) },],logic: "and",operator: "eq", },
}