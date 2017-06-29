package kendo

import "fmt"

func ExampleDataSourceModel_String() {

  e := DataSourceModel{
    Id: "personId",
    Fields: []FieldLine{
      {
        Field: "name",
        Type: "string",
        GoTemplate: &t,
      },
      {
        Field: "age",
        Type: "number",
        GoTemplate: &t,
      },
    },
    GoTemplate: &t,
  }

  fmt.Printf( "%v", e.String() )

  // Output:
  // transport: { read: { dataType: 'jsonp',url: "http://demos.telerik.com/kendo-ui/service/products" }, },
}
