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
  // autoWidth: true,
}
