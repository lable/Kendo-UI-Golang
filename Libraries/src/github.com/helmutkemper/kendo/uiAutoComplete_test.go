package kendo

import "fmt"

func ExampleUIAutoComplete_String1() {
  e := UIAutoComplete{
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}

func ExampleUIAutoComplete_String2() {
  e := UIAutoComplete{
    AnimationDisable: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // animation: false,
}

func ExampleUIAutoComplete_String3() {
  e := UIAutoComplete{
    Animation: Animation{
      Open: Open{
        Duration: 300,
        EffectEnum: EFFECT_ZOOM_IN,
        GoTemplate: &t,
      },
      Close: Close{
        Duration: 300,
        EffectEnum: EFFECT_ZOOM_IN,
        GoTemplate: &t,
      },
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // animation: open: { duration: 300,effects: 'zoom::in' },close: { duration: 300,effects: 'zoom::in', },
}

func ExampleUIAutoComplete_String4() {
  e := UIAutoComplete{
    AutoWidth: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // autoWidth: true,
}

func ExampleUIAutoComplete_String5() {
  e := UIAutoComplete{
    DataSource: DataSource{
      FilterLine: []FilterLine{
        {
          Field: "field_1",
          Logic: LOGIC_AND,
          Filters: []FiltersLine{
            {
              Field: "field_2",
              Logic: LOGIC_AND,
              Operator: OPERATOR_EQ,
              GoTemplate: &t,
            },
          },
          Operator: OPERATOR_CONTAINS,
          Value: ComplexJavaScriptType{
            AsFunction: "new Date(1980, 1, 1)",
          },
          GoTemplate: &t,
        },
      },
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // dataSource: {  field: 'field_1',filters: [ { field: 'field_2',logic: 'and',operator: 'eq', }, ],logic: 'and',operator: 'contains',value: new Date(1980, 1, 1), },
}

func ExampleUIAutoComplete_String6() {
  e := UIAutoComplete{
    DataSource: DataSource{
      Transport: Transport{
        Read: Read{
          Url: ComplexJavaScriptType{
            AsString: "http://demos.telerik.com/kendo-ui/service/products",
          },
          DataType: TYPE_DATA_JSON_JSONP, // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          GoTemplate: &t,
        },
        GoTemplate: &t,
      },
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // dataSource: {  transport: { read: { dataType: 'jsonp',url: "http://demos.telerik.com/kendo-ui/service/products" }, }, },
}

func ExampleUIAutoComplete_String7() {
  e := UIAutoComplete{
    ClearButton: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // clearButton: true,
}

func ExampleUIAutoComplete_String8() {
  e := UIAutoComplete{
    DataTextField: "name",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // dataTextField: 'name',
}

func ExampleUIAutoComplete_String9() {
  e := UIAutoComplete{
    Delay: 300,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // delay: 300,
}

func ExampleUIAutoComplete_String10() {
  e := UIAutoComplete{
    Enable: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // enable: true,
}

func ExampleUIAutoComplete_String11() {
  e := UIAutoComplete{
    EnforceMinLength: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // enforceMinLength: true,
}

func ExampleUIAutoComplete_String12() {
  e := UIAutoComplete{
    AutoCompleteFilterEnum: AUTOCOMPLETE_FILTER_STARTSWITH,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // filter: 'startswith',
}
/*
func ExampleUIAutoComplete_String13() {
  e := UIAutoComplete{
    FixedGroupTemplate: ,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // filter: 'startswith',
}
*/
func ExampleUIAutoComplete_String14() {
  e := UIAutoComplete{
    FooterTemplate: ComplexJavaScriptType{
      AsString: "Total <strong>#: instance.dataSource.total() #</strong> items found",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // footerTemplate: "Total <strong>#: instance.dataSource.total() #</strong> items found",
}

func ExampleUIAutoComplete_String15() {
  e := UIAutoComplete{
    NoDataTemplate: ComplexJavaScriptType{
      AsString: "No Data!",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // noDataTemplate: "No Data!",
}

func ExampleUIAutoComplete_String16() {
  e := UIAutoComplete{
    Placeholder: "Enter value...",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // placeholder: 'Enter value...',
}

func ExampleUIAutoComplete_String17() {
  e := UIAutoComplete{
    Separator: []string{
      ", ", "; ",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // separator: ', ','; ',
}

func ExampleUIAutoComplete_String18() {
  e := UIAutoComplete{
    HeaderTemplate: ComplexJavaScriptType{
      AsString: "<div><h2>Fruits</h2></div>",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // headerTemplate: "<div><h2>Fruits</h2></div>",
}

func ExampleUIAutoComplete_String19() {
  e := UIAutoComplete{
    Template: ComplexJavaScriptType{
      AsString: "#= OrderID # | For: #= ShipName #, #= ShipCountry #",
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // template: "#= OrderID # | For: #= ShipName #, #= ShipCountry #",
}

func ExampleUIAutoComplete_String20() {
  e := UIAutoComplete{
    Value: "Oranges",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // value: 'Oranges',
}

func ExampleUIAutoComplete_String21() {
  e := UIAutoComplete{
    ValuePrimitive: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // valuePrimitive: true,
}

func ExampleUIAutoComplete_String22() {
  e := UIAutoComplete{
    Virtual: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // virtual: true,
}

func ExampleUIAutoComplete_String_Completo() {
  e := UIAutoComplete{
    Animation: Animation{
      Open: Open{
        Duration: 300,
        EffectEnum: EFFECT_ZOOM_IN,
        GoTemplate: &t,
      },
      Close: Close{
        Duration: 300,
        EffectEnum: EFFECT_ZOOM_IN,
        GoTemplate: &t,
      },
      GoTemplate: &t,
    },
    AutoWidth: true,
    DataSource: DataSource{
      Transport: Transport{
        Read: Read{
          Url: ComplexJavaScriptType{
            AsString: "http://demos.telerik.com/kendo-ui/service/products",
          },
          DataType: TYPE_DATA_JSON_JSONP, // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          GoTemplate: &t,
        },
        GoTemplate: &t,
      },
      GoTemplate: &t,
    },
    ClearButton: true,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // animation: open: { duration: 300,effects: 'zoom::in' },close: { duration: 300,effects: 'zoom::in', },
  // autoWidth: true,
  // dataSource: {  transport: { read: { dataType: 'jsonp',url: "http://demos.telerik.com/kendo-ui/service/products" }, }, },
  // clearButton: true,
}