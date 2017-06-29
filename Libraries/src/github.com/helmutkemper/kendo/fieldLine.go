package kendo

import "bytes"

type FieldLine struct{
  Field     string
  Type      string

  GoTemplate    *GoTemplate
}

func ( el FieldLine ) getTemplate () string {
  return `{{if ne (string .Field) "''"}}{{string .Field}}: {{end}}{{if ne (string .Type) "''"}}{ type: {{string .Type}} }{{end}}`
}

func ( el FieldLine ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el FieldLine ) String() string {
  out := el.Buffer()
  return out.String()
}
