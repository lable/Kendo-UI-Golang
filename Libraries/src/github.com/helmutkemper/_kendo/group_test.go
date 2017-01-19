package kendo

import "fmt"

func ExampleGroup_String() {
  g := Group{
    Aggregates: []AggregateLine{
      {
        Field: "field_name",
        Aggregate: AGGREGATE_MAX,
        Template: &t,
      },
      {
        Field: "field_name_2",
        Aggregate: AGGREGATE_COUNT,
        Template: &t,
      },
      {
        Field: "field_name_3",
        Aggregate: AGGREGATE_MIN,
        Template: &t,
      },
    },
    Template: &t,
  }

  fmt.Print( g.String() )

  // Output:
  // aggregates: [{ field: &#34;field_name&#34;, aggregate: &#34;max&#34; },{ field: &#34;field_name_2&#34;, aggregate: &#34;count&#34; },{ field: &#34;field_name_3&#34;, aggregate: &#34;min&#34; },],
}

func ExampleGroup_String2() {
  g := Group{
    Dir: DIR_ASC,
    Template: &t,
  }

  fmt.Print( g.String() )

  // Output:
  // dir: "asc",
}


func ExampleGroup_String3() {
  g := Group{
    Field: "field_name",
    Template: &t,
  }

  fmt.Print( g.String() )

  // Output:
  // field: "field_name",
}
