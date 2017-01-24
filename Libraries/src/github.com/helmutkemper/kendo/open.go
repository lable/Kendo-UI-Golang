package kendo

type Open struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.open.duration
  //
  // Type: Number
  // Default: 200
  //
  // The duration of the open animation in milliseconds.
  Duration    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.open.effects
  //
  // Type: String
  //
  // The effect(s) to use when playing the open animation. Multiple effects should be separated with a space.
  //
  // Complete list of available animations http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/effects/common
  Effects    EffectEnum

  Template    *Template
}

func ( el Open ) getTemplate () string {
  return `{{if .Duration }}duration: {{.Duration}},{{end}}
{{if ne (string .Effects) "null"}}effects: {{string .Effects}},{{end}}
`
}

