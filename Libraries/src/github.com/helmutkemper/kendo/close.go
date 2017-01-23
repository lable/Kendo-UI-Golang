package kendo

type Close struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close.duration
  //
  // Type: Number
  // Default: 100
  //
  // The duration of the close animation in milliseconds.
  Duration    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close.effects
  //
  // Type: String
  //
  // The effect(s) to use when playing the close animation. Multiple effects should be separated with a space.
  //
  // Complete list of available animations http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/effects/common
  Effects    string

  Template    *Template
}

func ( el Close ) getTemplate () string {
  return `{{if .Duration }}duration: {{.Duration}},{{end}}
{{if ne (string .Effects) "null"}}effects: {{string .Effects}},{{end}}
`
}
