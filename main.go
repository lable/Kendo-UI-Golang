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

  type dataToJSon struct{
    Nome string
    Idade int
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
      {
        Aggregate: kendo.AGGREGATE_AVERAGE,
        Field: "name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Data: kendo.ComplexJavaScriptType{
      AsJSon: []dataToJSon{ { "José", 20 }, { "Maria", 30 } },
    },
  }
  fmt.Printf( "r: %s\n", ds.String() )
}