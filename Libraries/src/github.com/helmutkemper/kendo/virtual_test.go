package kendo

import "fmt"

func ExampleVirtual_String1() {
  e := Virtual{
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleVirtual_String2() {
  e := Virtual{
    ItemHeight: 100,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // itemHeight: 100,
}

func ExampleVirtual_String3() {
  e := Virtual{
    MapValueTo: MAP_VALUE_TO_INDEX,
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // mapValueTo: 'index',
}

func ExampleVirtual_String4() {
  e := Virtual{
    ValueMapper: ComplexJavaScriptType{
      AsFunction: "$.ajax( { url: 'http://demos.telerik.com/kendo-ui/service/Orders/ValueMapper', type: 'GET', dataType: 'jsonp', data: convertValues( options.value ), success: function ( data ) { options.success( data ); } } )",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // valueMapper: $.ajax( { url: 'http://demos.telerik.com/kendo-ui/service/Orders/ValueMapper', type: 'GET', dataType: 'jsonp', data: convertValues( options.value ), success: function ( data ) { options.success( data ); } } ),
}
