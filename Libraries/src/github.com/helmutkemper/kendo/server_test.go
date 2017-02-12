package kendo

import "fmt"

func ExampleServer_String() {
  e := Server{
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  //
}

func ExampleServer_String1() {
  e := Server{
    Create: "create",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // create: 'create',
}

func ExampleServer_String2() {
  e := Server{
    Read: "read",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // read: 'read',
}

func ExampleServer_String3() {
  e := Server{
    Destroy: "destroy",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // destroy: 'destroy',
}

func ExampleServer_String4() {
  e := Server{
    Update: "update",
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // update: 'update',
}

func ExampleServer_String5() {
  e := Server{
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
