package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.aggregate
//
// The name of the aggregate function. Specifies the aggregate function.
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
    Example - specify an aggregate function
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Tea", category: "Beverages", price: 1 },
        { name: "Coffee", category: "Beverages", price: 2 },
        { name: "Ham", category: "Food", price: 3 }
      ],
      group: {
        field: "category",
        aggregates: [
          // calculate max price
          { field: "price", aggregate: "max" }
        ]
      }
    });
    dataSource.fetch(function(){
      var view = dataSource.view();
      var beverages = view[0];
      console.log(beverages.aggregates.price.max); // displays "2"
    });
    </script>
*/
//
type AggregateLine struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.aggregate
  //
  // Type: String
  //
  // The name of the aggregate function. Specifies the aggregate function.
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
      Example - specify an aggregate function
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages", price: 1 },
          { name: "Coffee", category: "Beverages", price: 2 },
          { name: "Ham", category: "Food", price: 3 }
        ],
        group: {
          field: "category",
          aggregates: [
            // calculate max price
            { field: "price", aggregate: "max" }
          ]
        }
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        var beverages = view[0];
        console.log(beverages.aggregates.price.max); // displays "2"
      });
      </script>
  */
  //
  Aggregate    AggregateEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.field
  //
  // Type: String
  //
  // The data item field which will be used to calculate the aggregates.
  //
  //
  //
  /*
      Example - specify an aggregate field
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages", price: 1 },
          { name: "Coffee", category: "Beverages", price: 2 },
          { name: "Ham", category: "Food", price: 3 }
        ],
        group: {
          field: "category",
          aggregates: [
            // calculate max price
            { field: "price", aggregate: "max" }
          ]
        }
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        var beverages = view[0];
        console.log(beverages.aggregates.price.max); // displays "2"
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
