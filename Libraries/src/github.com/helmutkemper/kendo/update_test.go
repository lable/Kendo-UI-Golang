package kendo

import "fmt"

func ExampleUpdate_String1() {
  e := Update{
    Cache: false,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleUpdate_String2() {
  e := Update{
    Cache: true,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,
}

func ExampleUpdate_String3() {
  e := Update{
    ContentType: "application/json",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // contentType: 'application/json',
}

func ExampleUpdate_String4() {
  e := Update{
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // data: function() { return { name: 'Jane Doe', age: 30 } },
}

func ExampleUpdate_String5() {
  e := Update{
    DataType: TYPE_DATA_JSON_JSON,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // dataType: 'json',
}

func ExampleUpdate_String6() {
  e := Update{
    Method: METHOD_PUT,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // type: 'PUT',
}

func ExampleUpdate_String7() {
  e := Update{
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/update' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/update' },
}

func ExampleUpdate_String8() {
  e := Update{
    Cache: true,
    ContentType: "application/json",
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    DataType: TYPE_DATA_JSON_JSON,
    Method: METHOD_PUT,
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/update' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,
  // contentType: 'application/json',
  // data: function() { return { name: 'Jane Doe', age: 30 } },
  // dataType: 'json',
  // type: 'PUT',
  // url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/update' },
}
