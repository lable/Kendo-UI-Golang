package kendo

import "fmt"

func ExampleClient_String1() {
  e := Client{
    Create: "create",
    Read: "read",
    Destroy: "destroy",
    Update: "update",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // create: 'create',
  // destroy: 'destroy',
  // read: 'read',
  // update: 'update',
}
