package kendo

import "bytes"

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
//
// The aggregates which are calculated when the data source populates with data.
//
// The supported aggregates are:
//
// *  "average"
// *  "count"
// *  "max"
// *  "min"
// *  "sum"
/*
    Example - specify aggregate function
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Jane Doe", age: 30 },
        { name: "John Doe", age: 33 }
      ],
      aggregate: [
        { field: "age", aggregate: "sum" }
      ]
    });
    dataSource.fetch(function(){
      var results = dataSource.aggregates().age;
      console.log(results.sum); // displays "63"
    });
    </script>
*/
//
type AggregateLine struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
  //
  // Type: String
  //
  // The name of the aggregate function.
  //
  // The supported aggregates are:
  //
  // *  "average"
  // *  "count"
  // *  "max"
  // *  "min"
  // *  "sum"
  /*
      Example - specify aggregate function
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        aggregate: [
          { field: "age", aggregate: "sum" }
        ]
      });
      dataSource.fetch(function(){
        var results = dataSource.aggregates().age;
        console.log(results.sum); // displays "63"
      });
      </script>
  */
  //
  AggregateEnum    AggregateEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.field
  //
  // Type: String
  //
  // The data item field which will be used to calculate the aggregates.
  /*
      Example - specify aggregate field
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        aggregate: [
          { field: "age", aggregate: "sum" }
        ]
      });
      dataSource.fetch(function(){
        var results = dataSource.aggregates().age;
        console.log(results.sum); // displays "63"
      });
      </script>
  */
  //
  Field    string

  Template    *Template
}

func ( el AggregateLine ) getTemplate () string {
  return `{{if ne (string .Aggregate) "null"}}aggregate: {{string .Aggregate}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
`
}

func ( el AggregateLine ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.Template == nil {
    buffer.WriteString( "null" )
    return buffer
  }

  el.Template.ParserString( el.getTemplate() )
  el.Template.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el AggregateLine ) String() string {
  out := el.Buffer()
  return out.String()
}
