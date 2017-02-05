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
type AggregateEnum int

const(
  AGGREGATE_NIL AggregateEnum   = iota
  AGGREGATE_AVERAGE
  AGGREGATE_COUNT
  AGGREGATE_MAX
  AGGREGATE_MIN
  AGGREGATE_SUM
)

var AggregateEnums  = [...]string{
  "",
  "average",
  "count",
  "max",
  "min",
  "sum",
}

func (el AggregateEnum ) String() string {
  return AggregateEnums[ el ]
}

