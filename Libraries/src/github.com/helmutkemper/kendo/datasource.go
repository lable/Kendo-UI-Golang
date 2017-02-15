package kendo

import "bytes"

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
//
// kendo.data.DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource
// Overview http://docs.telerik.com/kendo-ui/api/javascript/data/datasource
// See the DataSource Overview http://docs.telerik.com/kendo-ui/api/javascript/data/datasource and Basic Usage http://docs.telerik.com/kendo-ui/api/javascript/data/datasource for an introduction to the DataSource.
// Configuration http://docs.telerik.com/kendo-ui/api/javascript/data/datasource
type DataSource struct {

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
  //
  // Type: Array
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
  //
  //    The data source calculates aggregates client-side unless the 'serverAggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to 'true'.
  //
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
  AggregateLine    []AggregateLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-autoSync
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true' the data source would automatically save any changed data items by calling the 'sync' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-sync  method. By default, changes are not automatically saved.
  /*
      Example - enable auto sync
      <script>
      var dataSource = new kendo.data.DataSource({
        autoSync: true,
        transport: {
          read:  {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          update: {
            url: "http://demos.telerik.com/kendo-ui/service/products/update",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        },
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20); // auto-syncs and makes request to http://demos.telerik.com/kendo-ui/service/products/update
      });
      </script>
  */
  AutoSync    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-batch
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true', the data source will batch CRUD operation requests. For example, updating two data items would cause one HTTP request instead of two. By default, the data source
  // makes a HTTP request for every CRUD operation.
  //
  //    The changed data items are sent as 'models' by default. This can be changed via the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap option.
  //
  /*
      Example - enable the batch mode
      <script>
      var dataSource = new kendo.data.DataSource({
        batch: true,
        transport: {
          read:  {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          update: {
            url: "http://demos.telerik.com/kendo-ui/service/products/update",
            dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        },
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20);
        var anotherProduct = dataSource.at(1);
        anotherProduct.set("UnitPrice", 20);
        dataSource.sync(); // causes only one request to "http://demos.telerik.com/kendo-ui/service/products/update"
      });
      </script>
  */
  Batch    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-data
  //
  // Type: Array
  //
  // The array of data items which the data source contains. The data source will wrap those items as 'kendo.data.ObservableObject' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/observableobject  or 'kendo.data.Model' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/model  (if 'schema.model' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema-model  is set).
  //
  // Can be set to a string value if the 'schema.type' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type  option is set to "xml".
  //
  /*
      Example - set the data items of a data source
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ]
      });
      dataSource.fetch(function(){
        var janeDoe = dataSource.at(0);
        console.log(janeDoe.name); // displays "Jane Doe"
      });
      </script>


      Example - set the data items as an XML string
      <script>
      var dataSource = new kendo.data.DataSource({
        data: '<books><book id="1"><title>Secrets of the JavaScript Ninja</title></book></books>',
        schema: {
          // specify the the schema is XML
          type: "xml",
          // the XML element which represents a single data record
          data: "/books/book",
          // define the model - the object which will represent a single data record
          model: {
            // configure the fields of the object
            fields: {
              // the "title" field is mapped to the text of the "title" XML element
              title: "title/text()",
              // the "id" field is mapped to the "id" attribute of the "book" XML element
              id: "@id"
            }
          }
        }
      });
      dataSource.fetch(function() {
        var books = dataSource.data();
        console.log(books[0].title); // displays "Secrets of the JavaScript Ninja"
      });
      </script>
  */
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter
  //
  // Type: Array
  //
  // The filters which are applied over the data items. By default, no filter is applied.
  //
  //    The data source filters the data items client-side unless the 'serverFiltering' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering option is set to 'true'.
  //
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
  FilterLine    []FilterLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group
  //
  // Type: Array
  //
  // The grouping configuration of the data source. If set, the data items will be grouped when the data source is populated. By default, grouping is not applied.
  //
  //    The data source groups the data items client-side unless the 'serverGrouping' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping option is set to 'true'.
  //
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
  GroupLine    []GroupLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-offlineStorage
  //
  // Type: String
  //
  // The offline storage key or custom offline storage implementation.
  /*
      Example - set offline storage key
      <script>
      var dataSource = new kendo.data.DataSource({
          offlineStorage: "products-offline",
          transport: {
              read: {
                  url: "http://demos.telerik.com/kendo-ui/service/products",
                  type: "jsonp"
              }
          }
      });
      </script>


      Example - set custom offline storage implementation
      <script>
      var dataSource = new kendo.data.DataSource({
          // use sessionStorage instead of localStorage
          offlineStorage: {
              getItem: function() {
                  return JSON.parse(sessionStorage.getItem("products-key"));
              },
              setItem: function(item) {
                  sessionStorage.setItem("products-key", JSON.stringify(item));
              }
          },
          transport: {
              read: {
                  url: "http://demos.telerik.com/kendo-ui/service/products",
                  type: "jsonp"
              }
          }
      });
      </script>
  */
  OfflineStorage    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-page
  //
  // Type: Number
  //
  // The page of data which the data source will return when the 'view' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-view  method is invoked or request from the remote service.
  //
  //    The data source will page the data items client-side unless the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true'.
  //
  /*
      Example - set the current page
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages" },
          { name: "Coffee", category: "Beverages" },
          { name: "Ham", category: "Food" }
        ],
        // set the second page as the current page
        page: 2,
        pageSize: 2
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        console.log(view.length); // displays "1"
        console.log(view[0].name); // displays "Ham"
      });
      </script>
  */
  Page    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-pageSize
  //
  // Type: Number
  //
  // The number of data items per page. The property has no default value. That is why to use paging, make sure some 'pageSize' value is set.
  //
  //    The data source will page the data items client-side unless the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true'.
  //
  /*
      Example - set the page size
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages" },
          { name: "Coffee", category: "Beverages" },
          { name: "Ham", category: "Food" }
        ],
        page: 1,
        // a page of data contains two data items
        pageSize: 2
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        console.log(view.length); // displays "2"
        console.log(view[0].name); // displays "Tea"
        console.log(view[1].name); // displays "Coffee"
      });
      </script>
  */
  PageSize    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
  //
  // Type: Object
  //
  // The configuration used to parse the remote service response.
  /*
      Example - specify the schema of the remote service
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: { q: "html5" } // search for tweets that contain "html5"
          }
        },
        schema: {
          data: function(response) {
            return response.results; // twitter's response is { "results": [ /@ results @/ ] }
          }
        }
      });
      dataSource.fetch(function(){
        var data = this.data();
        console.log(data.length);
      });
      </script>
  */
  Schema    Schema

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true', the data source will leave the aggregate calculation to the remote service. By default, the data source calculates aggregates client-side.
  //
  //    Configure 'schema.aggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.aggregates if you set 'serverAggregates' to 'true'.
  //
  /*
      Example - enable server aggregates
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverAggregates: true,
        aggregate: [
          { field: "age", aggregate: "sum" }
        ],
        schema: {
          aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
        }
      });
      </script>
  */
  ServerAggregates    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true', the data source will leave the filtering implementation to the remote service. By default, the data source performs filtering client-side.
  //
  // By default, the 'filter' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the filter '{ logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }' is sent as:
  //
  // 'filter[logic]: and'
  // 'filter[filters][0][field]: name'
  // 'filter[filters][0][operator]: startswith'
  // 'filter[filters][0][value]: Jane'
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the filter option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  /*
      Example - enable server filtering
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverFiltering: true,
        filter: { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }
      });
      </script>
  */
  ServerFiltering    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true', the data source will leave the grouping implementation to the remote service. By default, the data source performs grouping client-side.
  //
  // By default, the 'group' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the group '{ field: "category", dir: "desc" }' is sent as:
  //
  // 'group[0][field]: category'
  // 'group[0][dir]: desc'
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the group option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  /*
      Example - enable server grouping
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverGrouping: true,
        group: { field: "category", dir: "desc" }
      });
      </script>
  */
  ServerGrouping    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true', the data source will leave the data item paging implementation to the remote service. By default, the data source performs paging client-side.
  //
  //    Configure 'schema.total' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.total if you set 'serverPaging' to 'true'. In addition, 'pageSize' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-pageSize should be set no matter if paging is performed client-side or server-side.
  //
  /*
      Example - enable server paging
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverPaging: true,
        schema: {
          total: "total" // total is returned in the "total" field of the response
        }
      });
      </script>
  */
  ServerPaging    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true', the data source will leave the data item sorting implementation to the remote service. By default, the data source performs sorting client-side.
  //
  // By default, the 'sort' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the sort '{ field: "age", dir: "desc" }' is sent as:
  //
  // 'sort[0][field]: age'
  // 'sort[0][dir]: desc'
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the paging options in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  /*
      Example - enable server sorting
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverSorting: true,
        sort: { field: "age", dir: "desc" }
      });
      </script>
  */
  ServerSorting    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort
  //
  // Type: Array
  //
  // The sort order which will be applied over the data items. By default the data items are not sorted.
  //
  //    The data source sorts the data items client-side unless the serverSorting option is set to 'true'.
  //
  /*
      Example - sort the data items
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        sort: { field: "age", dir: "desc" }
      });
      dataSource.fetch(function(){
        var data = dataSource.view();
        console.log(data[0].age); // displays "33"
      });
      </script>


      Example - sort the data items by multiple fields
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages" },
          { name: "Coffee", category: "Beverages" },
          { name: "Ham", category: "Food" }
        ],
        sort: [
          // sort by "category" in descending order and then by "name" in ascending order
          { field: "category", dir: "desc" },
          { field: "name", dir: "asc" }
        ]
      });
      dataSource.fetch(function(){
        var data = dataSource.view();
        console.log(data[1].name); // displays "Coffee"
      });
      </script>
  */
  SortLine    []SortLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport
  //
  // Type: Object
  //
  // The configuration used to load and save the data items. A data source is remote or local based on the way of it retrieves data items.
  //
  // Remote data sources load and save data items from and to a remote end-point (also known as remote service or server). The 'transport' option describes the remote service configuration - URL, HTTP verb, HTTP headers, and others. The 'transport' option can also be used to implement custom data loading and saving.
  //
  // Local data sources are bound to a JavaScript array via the 'data' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-data  option.
  /*
      Example - specify the remote service configuration
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        console.log(products[0].ProductName); // displays "Chai"
      });
      </script>
  */
  Transport    Transport

  GoTemplate    *GoTemplate
}

func ( el DataSource ) getTemplate () string {
  return `{{if .ServerAggregates}}aggregate: [ {{range $v := .AggregateLine}}{{string $v}},{{end}} ],{{end}} {{if .ServerAggregates}}serverAggregates: true,{{end}}{{if .AutoSync}}autoSync: true,{{end}}{{if .Batch}}batch: true,{{end}}{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}{{if ne (string .FilterLine) ""}}{{string .FilterLine}}{{end}}{{if (ne (string .GroupLine) "") and .ServerGrouping}}{{$length := len .GroupLine}}{{if le $length 1}}group: { {{range $v := .GroupLine}}{{string $v}}{{end}}group: [ {{range $v := .GroupLine}}{{string $v}},{{end}} ],{{end}},{{end}}{{if ne (string .OfflineStorage) "null"}}offlineStorage: {{string .OfflineStorage}},{{end}}{{if .ServerPaging}}page: {{.Page}},{{end}}{{if .ServerPaging}}pageSize: {{.PageSize}},{{end}}{{if ne (string .Schema) ""}}schema: {{string .Schema}},{{end}}{{if .ServerFiltering}}serverFiltering: true,{{end}}{{if .ServerGrouping}}serverGrouping: true,{{end}}{{if .ServerPaging}}serverPaging: true,{{end}}{{if .ServerSorting}}serverSorting: true,{{end}}{{if ne (string .SortLine) ""}}{{$length := len .SortLine}}{{if le $length 1}}sort: { {{range $v := .SortLine}}{{string $v}}{{end}} },{{else}}sort: [ {{range $v := .SortLine}}{{string $v}},{{end}} ],{{end}},{{end}}{{if ne (string .Transport) ""}}transport: {{string .Transport}},{{end}}`
}

func ( el DataSource ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el DataSource ) String() string {
  out := el.Buffer()
  return out.String()
}
