package kendo

type Server struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.create
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for creating data items.
  Create    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.destroy
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for destroying data items.
  Destroy    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.read
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for reading data items.
  Read    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.update
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for updating data items.
  Update    string

  Template    *Template
}

func ( el Server ) getTemplate () string {
  return `{{if ne (string .Create) "null"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "null"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .Read) "null"}}read: {{string .Read}},{{end}}
{{if ne (string .Update) "null"}}update: {{string .Update}},{{end}}
`
}
