package kendo

import "fmt"

func ExampleTransport_String1() {
  e := Transport{
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleTransport_String2() {
  e := Transport{
    Read: Read{
      Url: ComplexJavaScriptType{
        AsString: "http://demos.telerik.com/kendo-ui/service/products",
      },
      DataType: TYPE_DATA_JSON_JSONP, // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // read: { dataType: 'jsonp',url: "http://demos.telerik.com/kendo-ui/service/products" },
}

