package kendo

import "bytes"

// The grouping configuration of the data source. If set, the data items will be grouped when the data source is populated. By default, grouping is not applied.
//
// The data source groups the data items client-side unless the 'serverGrouping' ( http://docs.telerik.com#configuration-serverGrouping ) option is set to 'true'.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group
/*
    Example - set a group as an object
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     // group by the "category" field
     group: { field: "category" }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "2"
     var beverages = view[0];
     console.log(beverages.value); // displays "Beverages"
     console.log(beverages.items[0].name); // displays "Tea"
     console.log(beverages.items[1].name); // displays "Coffee"
     var food = view[1];
     console.log(food.value); // displays "Food"
     console.log(food.items[0].name); // displays "Ham"
    });
    </script>


    Example - set a group as an array (subgroups)
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Pork", category: "Food", subcategory: "Meat" },
      { name: "Pepper", category: "Food", subcategory: "Vegetables" },
      { name: "Beef", category: "Food", subcategory: "Meat" }
     ],
     group: [
      // group by "category" and then by "subcategory"
      { field: "category" },
      { field: "subcategory" },
     ]
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     var food = view[0];
     console.log(food.value); // displays "Food"
     var meat = food.items[0];
     console.log(meat.value); // displays "Meat"
     console.log(meat.items.length); // displays "2"
     console.log(meat.items[0].name); // displays "Pork"
     console.log(meat.items[1].name); // displays "Beef"
     var vegetables = food.items[1];
     console.log(vegetables.value); // displays "Vegetables"
     console.log(vegetables.items.length); // displays "1"
     console.log(vegetables.items[0].name); // displays "Pepper"
    });
    </script>
*/
type Group struct {
  // The aggregates which are calculated during grouping.
  //
  // The supported aggregates are:
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
  /*
      Example - set group aggregates
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages", price: 1 },
        { name: "Coffee", category: "Beverages", price: 2 },
        { name: "Ham", category: "Food", price: 3 },
       ],
       group: {
        field: "category",
        aggregates: [
         { field: "price", aggregate: "max" },
         { field: "price", aggregate: "min" }
        ]
       }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var beverages = view[0];
       console.log(beverages.aggregates.price.max); // displays "2"
       console.log(beverages.aggregates.price.min); // displays "1"
       var food = view[1];
       console.log(food.aggregates.price.max); // displays "3"
       console.log(food.aggregates.price.min); // displays "3"
      });
      </script>


      Example - sort the groups in descending order
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages"},
        { name: "Ham", category: "Food"},
       ],
       // group by "category" in descending order
       group: { field: "category", dir: "desc" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var food = view[0];
       console.log(food.value); // displays "Food"
       var beverages = view[1];
       console.log(beverages.value); // displays "Beverages"
      });
      </script>


      Example - set the field
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages" },
        { name: "Coffee", category: "Beverages" },
        { name: "Ham", category: "Food" }
       ],
       // group by the "category" field
       group: { field: "category" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var beverages = view[0];
       console.log(beverages.items[0].name); // displays "Tea"
       console.log(beverages.items[1].name); // displays "Coffee"
       var food = view[1];
       console.log(food.items[0].name); // displays "Ham"
      });
      </script>
  */
  Aggregates      []AggregateList

  // default: "asc"
  //
  // The sort order of the group.
  //
  // The supported values are:
  //
  // The default sort order is ascending.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.dir
  /*
      Example - set group aggregates
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages", price: 1 },
        { name: "Coffee", category: "Beverages", price: 2 },
        { name: "Ham", category: "Food", price: 3 },
       ],
       group: {
        field: "category",
        aggregates: [
         { field: "price", aggregate: "max" },
         { field: "price", aggregate: "min" }
        ]
       }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var beverages = view[0];
       console.log(beverages.aggregates.price.max); // displays "2"
       console.log(beverages.aggregates.price.min); // displays "1"
       var food = view[1];
       console.log(food.aggregates.price.max); // displays "3"
       console.log(food.aggregates.price.min); // displays "3"
      });
      </script>


      Example - sort the groups in descending order
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages"},
        { name: "Ham", category: "Food"},
       ],
       // group by "category" in descending order
       group: { field: "category", dir: "desc" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var food = view[0];
       console.log(food.value); // displays "Food"
       var beverages = view[1];
       console.log(beverages.value); // displays "Beverages"
      });
      </script>


      Example - set the field
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages" },
        { name: "Coffee", category: "Beverages" },
        { name: "Ham", category: "Food" }
       ],
       // group by the "category" field
       group: { field: "category" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var beverages = view[0];
       console.log(beverages.items[0].name); // displays "Tea"
       console.log(beverages.items[1].name); // displays "Coffee"
       var food = view[1];
       console.log(food.items[0].name); // displays "Ham"
      });
      </script>
  */
  Dir             DirEnum

  // The data item field to group by.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.field
  /*
      Example - set group aggregates
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages", price: 1 },
        { name: "Coffee", category: "Beverages", price: 2 },
        { name: "Ham", category: "Food", price: 3 },
       ],
       group: {
        field: "category",
        aggregates: [
         { field: "price", aggregate: "max" },
         { field: "price", aggregate: "min" }
        ]
       }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var beverages = view[0];
       console.log(beverages.aggregates.price.max); // displays "2"
       console.log(beverages.aggregates.price.min); // displays "1"
       var food = view[1];
       console.log(food.aggregates.price.max); // displays "3"
       console.log(food.aggregates.price.min); // displays "3"
      });
      </script>


      Example - sort the groups in descending order
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages"},
        { name: "Ham", category: "Food"},
       ],
       // group by "category" in descending order
       group: { field: "category", dir: "desc" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var food = view[0];
       console.log(food.value); // displays "Food"
       var beverages = view[1];
       console.log(beverages.value); // displays "Beverages"
      });
      </script>


      Example - set the field
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages" },
        { name: "Coffee", category: "Beverages" },
        { name: "Ham", category: "Food" }
       ],
       // group by the "category" field
       group: { field: "category" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       var beverages = view[0];
       console.log(beverages.items[0].name); // displays "Tea"
       console.log(beverages.items[1].name); // displays "Coffee"
       var food = view[1];
       console.log(food.items[0].name); // displays "Ham"
      });
      </script>
  */
  Field           string

  Template          Template
}

func ( el Group ) getTemplate () string {
  return `group: { {{if .Aggregates}}aggregates: [{{range $v := .Aggregates}}{{string $v}},{{end}}],{{end}}{{if .Dir}}dir: "{{.Dir}}",{{end}}{{if .Field}}field: "{{.Field}}",{{end}} }`
}

func ( el Group ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer
  el.Template.ParserString( el.getTemplate() )
  el.Template.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Group ) String() string {
  var buffer bytes.Buffer
  el.Template.ParserString( el.getTemplate() )
  el.Template.ExecuteTemplate( &buffer, "", el )

  return buffer.String()
}