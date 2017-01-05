package main

import (
  "github.com/helmutkemper/kendo"
  "os"
  "fmt"
)

func main() {
  var t kendo.Template = kendo.Template{
    Out: os.Stdout,
  }

  //ds := kendo.NewDataSource( t )

  a := kendo.AggregateList{
    Aggregate: kendo.AGGREGATE_AVERAGE,
    Field: "name",
    Template: &t,
  }
  fmt.Print( a.String() )

  ds := kendo.DataSource{
    Aggregate: []kendo.AggregateList{
      {
        Aggregate: kendo.AGGREGATE_AVERAGE,
        Field: "name",
        Template: &t,
      },
      {
        Aggregate: kendo.AGGREGATE_AVERAGE,
        Field: "name",
        Template: &t,
      },
    },
  }
  fmt.Printf( "r: %s\n", ds.String() )
}