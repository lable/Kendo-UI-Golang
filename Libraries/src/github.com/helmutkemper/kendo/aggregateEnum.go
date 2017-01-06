package kendo

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
  AGGREGATE_NULL AggregateEnum = iota
  AGGREGATE_AVERAGE
  AGGREGATE_COUNT
  AGGREGATE_MAX
  AGGREGATE_MIN
  AGGREGATE_SUM
)

var AggregateEnums = [...]string{
  "",
  "average",
  "count",
  "max",
  "min",
  "sum",
}

func (el AggregateEnum) String() string {
  return AggregateEnums[ el ]
}
