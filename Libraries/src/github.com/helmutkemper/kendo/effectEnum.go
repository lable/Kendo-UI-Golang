package kendo

type EffectEnum int

const(
  EFFECT_NIL EffectEnum   = iota
  // Expends the element from zero to its regular size. Supported directions are 'horizontal' and 'vertical'. Playing
  // the effect in reverse will collapse the element to zero size and hide it.
  EFFECT_EXPAND_HORIZONTAL

  // Expends the element from zero to its regular size. Supported directions are 'horizontal' and 'vertical'. Playing
  // the effect in reverse will collapse the element to zero size and hide it.
  EFFECT_EXPAND_VERTICAL

  // Fades the element in or out. Supported directions are 'in' and 'out'.
  EFFECT_FADE_IN

  // Fades the element in or out. Supported directions are 'in' and 'out'.
  EFFECT_FADE_OUT

  // Flips the element around the axis specified by the axis parameter. Supported directions are 'horizontal' and
  // 'vertical'. The effect needs certain markup and styling in order to function properly. The element should be
  // positioned absolutely/relatively, and contain two child elements (face and back) with the same size as their parent,
  // positioned absolutely on top of each other.
  EFFECT_FLIP_HORIZONTAL

  // Flips the element around the axis specified by the axis parameter. Supported directions are 'horizontal' and
  // 'vertical'. The effect needs certain markup and styling in order to function properly. The element should be
  // positioned absolutely/relatively, and contain two child elements (face and back) with the same size as their parent,
  // positioned absolutely on top of each other.
  EFFECT_FLIP_VERTICAL

  // Slides the element to its original position in the specified direction, using the element width or height as an
  // offset. Playing the effect in reverse will slide the element out of its position.
  // The direction to which the sliding will occur. Supported directions are 'left', 'right', 'up' and 'down'.
  EFFECT_SLIDE_IN_LEFT

  // Slides the element to its original position in the specified direction, using the element width or height as an
  // offset. Playing the effect in reverse will slide the element out of its position.
  // The direction to which the sliding will occur. Supported directions are 'left', 'right', 'up' and 'down'.
  EFFECT_SLIDE_IN_RIGHT

  // Slides the element to its original position in the specified direction, using the element width or height as an
  // offset. Playing the effect in reverse will slide the element out of its position.
  // The direction to which the sliding will occur. Supported directions are 'left', 'right', 'up' and 'down'.
  EFFECT_SLIDE_IN_UP

  // Slides the element to its original position in the specified direction, using the element width or height as an
  // offset. Playing the effect in reverse will slide the element out of its position.
  // The direction to which the sliding will occur. Supported directions are 'left', 'right', 'up' and 'down'.
  EFFECT_SLIDE_IN_DOWN

  // Slides the element to its original position in the specified direction, while sliding out the element specified in
  // the previous parameter. Supported directions are 'left', 'right', 'up' and 'down'. For this effect to work as
  // expected, the elements should be positioned on top of each other.
  EFFECT_TILE_LEFT

  // Slides the element to its original position in the specified direction, while sliding out the element specified in
  // the previous parameter. Supported directions are 'left', 'right', 'up' and 'down'. For this effect to work as
  // expected, the elements should be positioned on top of each other.
  EFFECT_TILE_RIGHT

  // Slides the element to its original position in the specified direction, while sliding out the element specified in
  // the previous parameter. Supported directions are 'left', 'right', 'up' and 'down'. For this effect to work as
  // expected, the elements should be positioned on top of each other.
  EFFECT_TILE_UP

  // Slides the element to its original position in the specified direction, while sliding out the element specified in
  // the previous parameter. Supported directions are 'left', 'right', 'up' and 'down'. For this effect to work as
  // expected, the elements should be positioned on top of each other.
  EFFECT_TILE_DOWN

  // Zoom the element in or out. Supported directions are 'in' and 'out'.
  EFFECT_ZOOM_IN

  // Zoom the element in or out. Supported directions are 'in' and 'out'.
  EFFECT_ZOOM_OUT
)

var EffectEnums  = [...]string{
  "",
  "expand::horizontal",
  "expand::vertical",
  "fade::in",
  "fade::out",
  "flip::horizontal",
  "flip::vertical",
  "sladeIn::left",
  "sladeIn::right",
  "sladeIn::up",
  "sladeIn::down",
  "tile::left",
  "tile::right",
  "tile::up",
  "tile::down",
  "zoom::in",
  "zoom::out",
}

func (el EffectEnum ) String() string {
  return EffectEnums[ el ]
}

