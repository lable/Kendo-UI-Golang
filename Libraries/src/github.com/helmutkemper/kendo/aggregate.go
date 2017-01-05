//Comentários Revisado em 20170105
package kendo

import "bytes"

// The aggregates which are calculated when the data source populates with data.
//
// The supported aggregates are: "average", "count", "max", "min" or "sum"
//
// The data source calculates aggregates client-side unless the 'serverAggregates' ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates ) option is set to 'true'.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
/*
    Example - specify aggregates
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
     ],
     aggregate: [
      { field: "age", aggregate: "sum" },
      { field: "age", aggregate: "min" },
      { field: "age", aggregate: "max" }
     ]
    });
    dataSource.fetch(function(){
     var results = dataSource.aggregates().age;
     console.log(results.sum, results.min, results.max); // displays "63 30 33"
    });
    </script>


*/
type AggregateEnum int

const(
  AGGREGATE_AVERAGE AggregateEnum = iota
  AGGREGATE_COUNT
  AGGREGATE_MAX
  AGGREGATE_MIN
  AGGREGATE_SUM
)

var AggregateEnums = [...]string{
  "average",
  "count",
  "max",
  "min",
  "sum",
}

func (el AggregateEnum) String() string {
  return AggregateEnums[ el ]
}



// The aggregates which are calculated when the data source populates with data.
//
// The supported aggregates are: "average", "count", "max", "min" or "sum"
//
// The data source calculates aggregates client-side unless the 'serverAggregates'
// ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates ) option is set to
// 'true'.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
/*
    Example - specify aggregates
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
     ],
     aggregate: [
      { field: "age", aggregate: "sum" },
      { field: "age", aggregate: "min" },
      { field: "age", aggregate: "max" }
     ]
    });
    dataSource.fetch(function(){
     var results = dataSource.aggregates().age;
     console.log(results.sum, results.min, results.max); // displays "63 30 33"
    });
    </script>
*/
type AggregateList struct {
  // The name of the aggregate function.
  //
  // The supported aggregates are: "average", "count", "max", "min" or "sum"
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
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
  Aggregate AggregateEnum

  // The data item field which will be used to calculate the aggregates.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.field
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
  Field     string

  Template *Template
}

func ( el AggregateList ) getTemplate () string {
  return `    { field: "{{.Field}}", aggregate: "{{.Aggregate}}" }{{nl}}`
}

func ( el AggregateList ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer
  el.Template.ParserString( el.getTemplate() )
  el.Template.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el AggregateList ) String() string {
  out := el.Buffer()
  return out.String()
}