package kendo

import "fmt"

func ExampleEffectEnum_String() {
  var e EffectEnum = EFFECTS_ZOOM_IN
  fmt.Print( e.String() )

  // Output:
  // zoom:in
}

func ExampleEffectEnum_String2() {
  var e EffectEnum = EFFECTS_ZOOM_OUT
  fmt.Print( e.String() )

  // Output:
  // zoom:out
}

func ExampleEffectEnum_String3() {
  var e EffectEnum = EFFECTS_EXPAND_HORIZONTAL
  fmt.Print( e.String() )

  // Output:
  // expand:horizontal
}

func ExampleEffectEnum_String4() {
  var e EffectEnum = EFFECTS_EXPAND_VERTICAL
  fmt.Print( e.String() )

  // Output:
  // expand:vertical
}

func ExampleEffectEnum_String5() {
  var e EffectEnum = EFFECTS_FADE_IN
  fmt.Print( e.String() )

  // Output:
  // fade:in
}

func ExampleEffectEnum_String6() {
  var e EffectEnum = EFFECTS_FADE_OUT
  fmt.Print( e.String() )

  // Output:
  // fade:out
}

func ExampleEffectEnum_String7() {
  var e EffectEnum = EFFECTS_FLIP_HORIZONTAL
  fmt.Print( e.String() )

  // Output:
  // flip:horizontal
}

func ExampleEffectEnum_String8() {
  var e EffectEnum = EFFECTS_FLIP_VERTICAL
  fmt.Print( e.String() )

  // Output:
  // flip:vertical
}

func ExampleEffectEnum_String9() {
  var e EffectEnum = EFFECTS_TILE_LEFT
  fmt.Print( e.String() )

  // Output:
  // tile:left
}

func ExampleEffectEnum_String10() {
  var e EffectEnum = EFFECTS_TILE_RIGHT
  fmt.Print( e.String() )

  // Output:
  // tile:right
}

func ExampleEffectEnum_String11() {
  var e EffectEnum = EFFECTS_TILE_UP
  fmt.Print( e.String() )

  // Output:
  // tile:up
}

func ExampleEffectEnum_String12() {
  var e EffectEnum = EFFECTS_TILE_DOWN
  fmt.Print( e.String() )

  // Output:
  // tile:down
}
