package kendo

import "fmt"

func ExampleAnimation_String() {
  e := Animation{
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}

func ExampleAnimation_String1() {
  e := Animation{
    Close: Close{
      Duration: 300,
      EffectEnum: EFFECT_ZOOM_IN,
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // close: { duration: 300,effects: 'zoom::in', }
}

func ExampleAnimation_String2() {
  e := Animation{
    Open: Open{
      Duration: 300,
      EffectEnum: EFFECT_ZOOM_IN,
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // open: { duration: 300,effects: 'zoom::in' },
}

func ExampleAnimation_String3() {
  e := Animation{
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
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // open: { duration: 300,effects: 'zoom::in' },close: { duration: 300,effects: 'zoom::in', }
}