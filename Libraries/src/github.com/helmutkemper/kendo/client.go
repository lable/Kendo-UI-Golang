package kendo

import "bytes"

type Client struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.create
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for creating data items.
  Create    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.destroy
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for destroying data items.
  Destroy    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.read
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for reading data items.
  Read    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.update
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for updating data items.
  Update    string

  GoTemplate    *GoTemplate
}

func ( el Client ) getTemplate () string {
  return `{{if ne (string .Create) "''"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "''"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .Read) "''"}}read: {{string .Read}},{{end}}
{{if ne (string .Update) "''"}}update: {{string .Update}},{{end}}
`
}

func ( el Client ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "null" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Client ) String() string {
  out := el.Buffer()
  return out.String()
}
