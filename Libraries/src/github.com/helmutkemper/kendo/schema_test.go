package kendo

import "fmt"

func ExampleSchema_String() {
  var s Schema = Schema{
    Data: ComplexJavaScriptType{
      AsFunction: `function(response) { return response.results; }`,
    },
    Template: &t,
  }
  fmt.Print( s.String() )

  // Output:
  // { schema: data: function(response) { return response.results; } },
}
