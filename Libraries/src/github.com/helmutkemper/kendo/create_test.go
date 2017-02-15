package kendo

import "fmt"

func ExampleCreate_String1() {
  e := Create{
    Cache: false,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleCreate_String2() {
  e := Create{
    Cache: true,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,
}

func ExampleCreate_String3() {
  e := Create{
    ContentType: "application/json",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // contentType: 'application/json',
}

func ExampleCreate_String4() {
  e := Create{
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // data: function() { return { name: 'Jane Doe', age: 30 } },
}

func ExampleCreate_String5() {
  e := Create{
    DataType: TYPE_DATA_JSON_JSON,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // dataType: 'json',
}

func ExampleCreate_String6() {
  e := Create{
    Method: METHOD_PUT,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // type: 'PUT',
}

func ExampleCreate_String7() {
  e := Create{
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' },
}

func ExampleCreate_String8() {
  e := Create{
    Cache: true,
    ContentType: "application/json",
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    DataType: TYPE_DATA_JSON_JSON,
    Method: METHOD_PUT,
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,contentType: 'application/json',data: function() { return { name: 'Jane Doe', age: 30 } },dataType: 'json',type: 'PUT',url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' },
}
