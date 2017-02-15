package kendo

import "fmt"

func ExampleClient_String() {
  e := Client{
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}

func ExampleClient_String1() {
  e := Client{
    Create: "create",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // create: 'create',
}

func ExampleClient_String2() {
  e := Client{
    Read: "read",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // read: 'read',
}

func ExampleClient_String3() {
  e := Client{
    Destroy: "destroy",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // destroy: 'destroy',
}

func ExampleClient_String4() {
  e := Client{
    Update: "update",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // update: 'update',
}

func ExampleClient_String5() {
  e := Client{
    Create: "create",
    Read: "read",
    Destroy: "destroy",
    Update: "update",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // create: 'create',destroy: 'destroy',read: 'read',update: 'update',
}
