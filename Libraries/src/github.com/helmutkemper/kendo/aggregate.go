package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
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
//
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
type Aggregate struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
  //
  // Type: String
  //
  // The name of the aggregate function.
  //
  // The supported aggregates are:
  //
  //
  // *  "average"
  // *  "count"
  // *  "max"
  // *  "min"
  // *  "sum"
  //
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
  Aggregate    AggregateEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.field
  //
  // Type: String
  //
  // The data item field which will be used to calculate the aggregates.
  //
  //
  //
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

func ( el Aggregate ) getTemplate () string {
  return `{{if ne (string .Aggregate) "null"}}aggregate: {{string .Aggregate}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
`
}
