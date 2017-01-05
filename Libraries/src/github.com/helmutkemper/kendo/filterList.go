//Comentários Revisado em 20170105
package kendo

// The filters which are applied over the data items. By default, no filter is applied.
//
// The data source filters the data items client-side unless the 'serverFiltering'
// ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering ) option is set to
// 'true'.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter
/*
    Example - set a single filter
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe" },
      { name: "John Doe" }
     ],
     filter: { field: "name", operator: "startswith", value: "Jane" }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     console.log(view[0].name); // displays "Jane Doe"
    });
    </script>


    Example - set filter as conjunction (and)
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     filter: [
      // leave data items which are "Beverage" and not "Coffee"
      { field: "category", operator: "eq", value: "Beverages" },
      { field: "name", operator: "neq", value: "Coffee" }
     ]
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     console.log(view[0].name); // displays "Tea"
    });
    </script>


    Example - set filter as disjunction (or)
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     filter: {
      // leave data items which are "Food" or "Tea"
      logic: "or",
      filters: [
       { field: "category", operator: "eq", value: "Food" },
       { field: "name", operator: "eq", value: "Tea" }
      ]
     }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "2"
     console.log(view[0].name); // displays "Tea"
     console.log(view[1].name); // displays "Ham"
    });
    </script>
*/
type FilterLine struct {
  // The data item field to which the filter operator is applied.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
  /*
      Example - set the filter field
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe" },
        { name: "John Doe" }
       ],
       filter: { field: "name", operator: "startswith", value: "Jane" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "1"
       console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  Field     string

  // The filter operator (comparison).
  //
  // The supported operators are:
  //
  // '"isnotempty"'
  //
  // The last five are supported only for string fields.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.operator
  /*
      Example - set the filter operator
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe" },
        { name: "John Doe" }
       ],
       filter: { field: "name", operator: "startswith", value: "Jane" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "1"
       console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  Operator  OperatorEnum

  // The value to which the 'field'
  // ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field ) is compared. The
  // value has to be of the same type as the field.
  //
  // By design, the '"\n"' is removed from the filter before the filtering is performed. That is why an '"\n"'
  // identifier from the filter will not match data items whose corresponding fields contain new lines.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.value
  /*
      Example - specify the filter value
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe", birthday: new Date(1983, 1, 1) },
        { name: "John Doe", birthday: new Date(1980, 1, 1)}
       ],
       filter: { field: "birthday", operator: "gt", value: new Date(1980, 1, 1) }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "1"
       console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  Value     ComplexJavaScriptType
}

// The nested filter expressions. Supports the same options as 'filter'
// ( http://docs.telerik.comhttp://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter ).
// Filters can be nested indefinitely.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.filters
/*
    Example - set the filter field
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe" },
      { name: "John Doe" }
     ],
     filter: { field: "name", operator: "startswith", value: "Jane" }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     console.log(view[0].name); // displays "Jane Doe"
    });
    </script>


    Example - nested filters
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     filter: {
      // leave data items which are "Food" or "Tea"
      logic: "or",
      filters: [
       { field: "category", operator: "eq", value: "Food" },
       { field: "name", operator: "eq", value: "Tea" }
      ]
     }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "2"
     console.log(view[0].name); // displays "Tea"
     console.log(view[1].name); // displays "Ham"
    });
    </script>


    Example - set the filter logic
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     filter: {
      // leave data items which are "Food" or "Tea"
      logic: "or",
      filters: [
       { field: "category", operator: "eq", value: "Food" },
       { field: "name", operator: "eq", value: "Tea" }
      ]
     }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "2"
     console.log(view[0].name); // displays "Tea"
     console.log(view[1].name); // displays "Ham"
    });
    </script>


    Example - set the filter operator
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe" },
      { name: "John Doe" }
     ],
     filter: { field: "name", operator: "startswith", value: "Jane" }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     console.log(view[0].name); // displays "Jane Doe"
    });
    </script>


    Example - specify the filter value
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe", birthday: new Date(1983, 1, 1) },
      { name: "John Doe", birthday: new Date(1980, 1, 1)}
     ],
     filter: { field: "birthday", operator: "gt", value: new Date(1980, 1, 1) }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     console.log(view[0].name); // displays "Jane Doe"
    });
    </script>
*/
type Filter struct {
  // The data item field to which the filter operator is applied.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
  /*
      Example - set the filter field
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe" },
        { name: "John Doe" }
       ],
       filter: { field: "name", operator: "startswith", value: "Jane" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "1"
       console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  Field     string

  // The nested filter expressions. Supports the same options as 'filter'
  // ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter ). Filters can be nested
  // indefinitely.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.filters
  /*
      Example - nested filters
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages" },
        { name: "Coffee", category: "Beverages" },
        { name: "Ham", category: "Food" }
       ],
       filter: {
        // leave data items which are "Food" or "Tea"
        logic: "or",
        filters: [
         { field: "category", operator: "eq", value: "Food" },
         { field: "name", operator: "eq", value: "Tea" }
        ]
       }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "2"
       console.log(view[0].name); // displays "Tea"
       console.log(view[1].name); // displays "Ham"
      });
      </script>
  */
  Filters   []FilterLine

  // The logical operation to use when the 'filter.filters' option is set.
  //
  // The supported values are:
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.logic
  /*
      Example - set the filter logic
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Tea", category: "Beverages" },
        { name: "Coffee", category: "Beverages" },
        { name: "Ham", category: "Food" }
       ],
       filter: {
        // leave data items which are "Food" or "Tea"
        logic: "or",
        filters: [
         { field: "category", operator: "eq", value: "Food" },
         { field: "name", operator: "eq", value: "Tea" }
        ]
       }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "2"
       console.log(view[0].name); // displays "Tea"
       console.log(view[1].name); // displays "Ham"
      });
      </script>
  */
  Logic     LogicEnum

  // The filter operator (comparison).
  //
  // The supported operators are:
  //
  // '"isnotempty"'
  //
  // The last five are supported only for string fields.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.operator
  /*
      Example - set the filter operator
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe" },
        { name: "John Doe" }
       ],
       filter: { field: "name", operator: "startswith", value: "Jane" }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "1"
       console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  Operator  OperatorEnum

  // The value to which the 'field'
  // ( hhttp://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field ) is compared. The
  // value has to be of the same type as the field.
  //
  // By design, the '"\n"' is removed from the filter before the filtering is performed. That is why an '"\n"'
  // identifier from the filter will not match data items whose corresponding fields contain new lines.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.value
  /*
      Example - specify the filter value
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe", birthday: new Date(1983, 1, 1) },
        { name: "John Doe", birthday: new Date(1980, 1, 1)}
       ],
       filter: { field: "birthday", operator: "gt", value: new Date(1980, 1, 1) }
      });
      dataSource.fetch(function(){
       var view = dataSource.view();
       console.log(view.length); // displays "1"
       console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  Value     ComplexJavaScriptType
}
