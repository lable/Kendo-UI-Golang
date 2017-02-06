package kendo

import "fmt"

func ExampleDestroy_String1() {
  e := Destroy{
    Cache: false,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleDestroy_String2() {
  e := Destroy{
    Cache: true,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // cache: true,
}

func ExampleDestroy_String3() {
  e := Destroy{
    ContentType: "application/json",
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // contentType: 'application/json',
}

func ExampleDestroy_String4() {
  e := Destroy{
    Data: ComplexJavaScriptType{
      AsFunction: "function() { return { name: 'Jane Doe', age: 30 } }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // data: function() { return { name: 'Jane Doe', age: 30 } },
}

func ExampleDestroy_String5() {
  e := Destroy{
    DataType: TYPE_DATA_JSON_JSON,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // dataType: 'json',
}

func ExampleDestroy_String6() {
  e := Destroy{
    Method: METHOD_PUT,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // type: 'PUT',
}

func ExampleDestroy_String7() {
  e := Destroy{
    Url: ComplexJavaScriptType{
      AsFunction: "function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' }",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' },
}

func ExampleDestroy_String8() {
  e := Destroy{
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
  // cache: true,
  // contentType: 'application/json',
  // data: function() { return { name: 'Jane Doe', age: 30 } },
  // dataType: 'json',
  // type: 'PUT',
  // url: function(options) { return 'http://demos.telerik.com/kendo-ui/service/products/create' },
}
