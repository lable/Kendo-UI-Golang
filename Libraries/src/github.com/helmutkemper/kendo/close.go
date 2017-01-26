package kendo

import "bytes"

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
  EffectEnum    EffectEnum

  GoTemplate    *GoTemplate
}

func ( el Close ) getTemplate () string {
  return `{{if .Duration }}duration: {{.Duration}},{{end}}
{{if ne (string .EffectEnum) ""}}effects: '{{string .EffectEnum}}',{{end}}
`
}

func ( el Close ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "null" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Close ) String() string {
  out := el.Buffer()
  return out.String()
}
