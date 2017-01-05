package kendo

type AnimationAction struct {
  // The effect(s) to use when playing the close animation. Multiple effects should be separated with a space.
  effects   []EffectEnum
  // The duration of the open/close animation in milliseconds.
  // default: 100
  duration  int
}

type Animation struct {
  // The animation played when the suggestion popup is opened.
  Open    AnimationAction
  // The animation played when the suggestion popup is closed.
  Close   AnimationAction
}