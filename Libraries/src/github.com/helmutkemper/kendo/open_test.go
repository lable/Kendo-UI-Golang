package kendo

import "fmt"

func ExampleOpen_String1() {
  e := Open{
    Duration: 300,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // duration: 300
}

func ExampleOpen_String2() {
  e := Open{
    EffectEnum: EFFECT_ZOOM_IN,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // effects: 'zoom::in'
}

func ExampleOpen_String3() {
  e := Open{
    Duration: 300,
    EffectEnum: EFFECT_ZOOM_IN,
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // duration: 300,effects: 'zoom::in'
}