package kendo

type EffectEnum int

const(
  EFFECTS_ZOOM_NULL EffectEnum = iota
  EFFECTS_ZOOM_IN
  EFFECTS_ZOOM_OUT
  EFFECTS_EXPAND_HORIZONTAL
  EFFECTS_EXPAND_VERTICAL
  EFFECTS_FADE_IN
  EFFECTS_FADE_OUT
  EFFECTS_FLIP_HORIZONTAL
  EFFECTS_FLIP_VERTICAL
  EFFECTS_TILE_LEFT
  EFFECTS_TILE_RIGHT
  EFFECTS_TILE_UP
  EFFECTS_TILE_DOWN

)

var EffectEnums = [...]string{
  "",
  "zoom:in",
  "zoom:out",
  "expand:horizontal",
  "expand:vertical",
  "fade:in",
  "fade:out",
  "flip:horizontal",
  "flip:vertical",
  "tile:left",
  "tile:right",
  "tile:up",
  "tile:down",
}

func (el EffectEnum) String() string {
  return EffectEnums[ el ]
}