package kendo

import "fmt"

func ExampleEffectEnum_String1() {
  e := EFFECT_NIL

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleEffectEnum_String2() {
  e := EFFECT_EXPAND_HORIZONTAL

  fmt.Print( e.String() )

  // Output:
  // expand::horizontal
}

func ExampleEffectEnum_String3() {
  e := EFFECT_EXPAND_VERTICAL

  fmt.Print( e.String() )

  // Output:
  // expand::vertical
}

func ExampleEffectEnum_String4() {
  e := EFFECT_FADE_IN

  fmt.Print( e.String() )

  // Output:
  // fade::in
}

func ExampleEffectEnum_String5() {
  e := EFFECT_FADE_OUT

  fmt.Print( e.String() )

  // Output:
  // fade::out
}

func ExampleEffectEnum_String6() {
  e := EFFECT_FLIP_HORIZONTAL

  fmt.Print( e.String() )

  // Output:
  // flip::horizontal
}

func ExampleEffectEnum_String7() {
  e := EFFECT_FLIP_VERTICAL

  fmt.Print( e.String() )

  // Output:
  // flip::vertical
}

func ExampleEffectEnum_String8() {
  e := EFFECT_SLIDE_IN_LEFT

  fmt.Print( e.String() )

  // Output:
  // sladeIn::left
}

func ExampleEffectEnum_String9() {
  e := EFFECT_SLIDE_IN_RIGHT

  fmt.Print( e.String() )

  // Output:
  // sladeIn::right
}

func ExampleEffectEnum_String10() {
  e := EFFECT_SLIDE_IN_UP

  fmt.Print( e.String() )

  // Output:
  // sladeIn::up
}

func ExampleEffectEnum_String11() {
  e := EFFECT_SLIDE_IN_DOWN

  fmt.Print( e.String() )

  // Output:
  // sladeIn::down
}

func ExampleEffectEnum_String12() {
  e := EFFECT_TILE_LEFT

  fmt.Print( e.String() )

  // Output:
  // tile::left
}

func ExampleEffectEnum_String13() {
  e := EFFECT_TILE_RIGHT

  fmt.Print( e.String() )

  // Output:
  // tile::right
}

func ExampleEffectEnum_String14() {
  e := EFFECT_TILE_UP

  fmt.Print( e.String() )

  // Output:
  // tile::up
}

func ExampleEffectEnum_String15() {
  e := EFFECT_TILE_DOWN

  fmt.Print( e.String() )

  // Output:
  // tile::down
}

func ExampleEffectEnum_String16() {
  e := EFFECT_ZOOM_IN

  fmt.Print( e.String() )

  // Output:
  // zoom::in
}

func ExampleEffectEnum_String17() {
  e := EFFECT_ZOOM_OUT

  fmt.Print( e.String() )

  // Output:
  // zoom::out
}
