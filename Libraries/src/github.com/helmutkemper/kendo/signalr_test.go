package kendo

import "fmt"

func ExampleSignalr_String1() {
  e := Signalr{
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  //
}

func ExampleSignalr_String2() {
  e := Signalr{
    Client: Client{
      Create: "create",
      Read: "read",
      Destroy: "destroy",
      Update: "update",
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // client: { 'create: 'create',
  // destroy: 'destroy',
  // read: 'read',
  // update: 'update',
  // ' },
}

func ExampleSignalr_String3() {
  e := Signalr{
    Hub: ComplexJavaScriptType{
      AsFunction: "connection.createHubProxy('productHub')",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // hub: connection.createHubProxy('productHub'),
}

func ExampleSignalr_String4() {
  e := Signalr{
    Promise: ComplexJavaScriptType{
      AsFunction: "connection.start({ jsonp: true })",
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // promise: { connection.start({ jsonp: true }) },
}

func ExampleSignalr_String5() {
  e := Signalr{
    Server: Server{
      Create: "create",
      Read: "read",
      Destroy: "destroy",
      Update: "update",
      GoTemplate: &t,
    },
    GoTemplate: &t,
  }

  fmt.Print( e.String() )

  // Output:
  // server: { 'create: 'create',
  // destroy: 'destroy',
  // read: 'read',
  // update: 'update',
  // ' },
}