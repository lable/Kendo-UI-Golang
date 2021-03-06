package kendo

import "bytes"

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client
//
// Specifies the client-side CRUD methods of the SignalR hub.
//
//
//

//
type Signalr struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client
  //
  // Type: Object
  //
  // Specifies the client-side CRUD methods of the SignalR hub.
  //
  //
  //

  //
  Client    Client

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.hub
  //
  // Type: Object
  //
  // The SignalR hub object returned by the 'createHubProxy' method. The 'hub' option is mandatory.
  //
  //
  //

  //
  Hub    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.promise
  //
  // Type: Object
  //
  // The promise returned by the 'start' method of the SignalR connection. The 'promise' option is mandatory.
  Promise    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server
  //
  // Type: Object
  //
  // Specifies the server-side CRUD methods of the SignalR hub.
  //
  //
  //

  //
  Server    Server

  GoTemplate    *GoTemplate
}

func ( el Signalr ) getTemplate () string {
  return `{{if ne (string .Client) ""}}client: { {{string .Client}} },{{end}}{{if ne (string .Hub) "null"}}hub: {{string .Hub}},{{end}}{{if ne (string .Promise) "null"}}promise: { {{string .Promise}} },{{end}}{{if ne (string .Server) ""}}server: { {{string .Server}} },{{end}}`
}

func ( el Signalr ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Signalr ) String() string {
  out := el.Buffer()
  return out.String()
}
