package kendo

import "fmt"

func ExampleRead_String1() {
  e := Read{
    Cache: false,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleRead_String2() {
  e := Read{
    Cache: true,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,
}

func ExampleRead_String3() {
  e := Read{
    ContentType: "application/json",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // contentType: 'application/json',
}

func ExampleRead_String4() {
  e := Read{
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // data: function() { return { name: 'Jane Doe', age: 30 } },
}

func ExampleRead_String5() {
  e := Read{
    DataType: TYPE_DATA_JSON_JSON,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // dataType: 'json',
}

func ExampleRead_String6() {
  e := Read{
    Method: METHOD_PUT,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // type: 'PUT',
}

func ExampleRead_String7() {
  e := Read{
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/Read' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/Read' },
}

func ExampleRead_String8() {
  e := Read{
    Cache: true,
    ContentType: "application/json",
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    DataType: TYPE_DATA_JSON_JSON,
    Method: METHOD_PUT,
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/read' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,contentType: 'application/json',data: function() { return { name: 'Jane Doe', age: 30 } },dataType: 'json',type: 'PUT',url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/read' },
}
