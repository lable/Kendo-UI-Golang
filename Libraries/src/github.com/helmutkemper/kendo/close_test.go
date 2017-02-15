package kendo

import "fmt"

func ExampleClose_String1() {
  e := Close{
    Duration: 300,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // duration: 300,
}

func ExampleClose_String2() {
  e := Close{
    EffectEnum: EFFECT_ZOOM_IN,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // effects: 'zoom::in',
}

func ExampleClose_String3() {
  e := Close{
    Duration: 300,
    EffectEnum: EFFECT_ZOOM_IN,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // duration: 300,effects: 'zoom::in',
}