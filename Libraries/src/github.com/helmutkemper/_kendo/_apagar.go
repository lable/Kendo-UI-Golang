package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
//
// kendo.data.DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource
// Overview http://docs.telerik.com/kendo-ui/api/javascript/data/datasource
// See the DataSource Overview http://docs.telerik.com/kendo-ui/api/javascript/data/datasource and Basic Usage http://docs.telerik.com/kendo-ui/api/javascript/data/datasource for an introduction to the DataSource.
// Configuration http://docs.telerik.com/kendo-ui/api/javascript/data/datasource
//
type Datasource struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
  //
  // Type: Array
  //
  // The aggregates which are calculated when the data source populates with data.
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
  //
  Aggregate    []AggregateLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-autoSync
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true' the data source would automatically save any changed data items by calling the 'sync' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-sync  method. By default, changes are not automatically saved.
  //
  //
  //
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
  //
  AutoSync    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-batch
  //
  // Type: Boolean
  // Defalt: false
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
  //
  Batch    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-data
  //
  // Type: Array
  //
  // The array of data items which the data source contains. The data source will wrap those items as 'kendo.data.ObservableObject' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/observableobject  or 'kendo.data.Model' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/model  (if 'schema.model' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema-model  is set).
  //
  // Can be set to a string value if the 'schema.type' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type  option is set to "xml".
  //
  //
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
  //
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
  //
  Filter    []FilterLine

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
  //
  Group    []GroupLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-offlineStorage
  //
  // Type: String
  //
  // The offline storage key or custom offline storage implementation.
  //
  //
  //
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
  //
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
  //
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
  //
  PageSize    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
  //
  // Type: Object
  //
  // The configuration used to parse the remote service response.
  //
  //
  //
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
  //
  Schema    Schema

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates
  //
  // Type: Boolean
  // Defalt: false
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
  //
  ServerAggregates    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the filtering implementation to the remote service. By default, the data source performs filtering client-side.
  //
  // By default, the 'filter' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the filter '{ logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }' is sent as:
  //
  //
  // 'filter[logic]: and'
  //
  // 'filter[filters][0][field]: name'
  //
  // 'filter[filters][0][operator]: startswith'
  //
  // 'filter[filters][0][value]: Jane'
  //
  //
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the filter option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  //
  //
  //
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
  //
  ServerFiltering    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the grouping implementation to the remote service. By default, the data source performs grouping client-side.
  //
  // By default, the 'group' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the group '{ field: "category", dir: "desc" }' is sent as:
  //
  //
  // 'group[0][field]: category'
  //
  // 'group[0][dir]: desc'
  //
  //
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the group option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  //
  //
  //
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
  //
  ServerGrouping    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging
  //
  // Type: Boolean
  // Defalt: false
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
  //
  ServerPaging    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the data item sorting implementation to the remote service. By default, the data source performs sorting client-side.
  //
  // By default, the 'sort' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the sort '{ field: "age", dir: "desc" }' is sent as:
  //
  //
  // 'sort[0][field]: age'
  // 'sort[0][dir]: desc'
  //
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the paging options in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  //
  //
  //
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
  //
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
  //
  Sort    []SortLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport
  //
  // Type: Object
  //
  // The configuration used to load and save the data items. A data source is remote or local based on the way of it retrieves data items.
  //
  // Remote data sources load and save data items from and to a remote end-point (also known as remote service or server). The 'transport' option describes the remote service configuration - URL, HTTP verb, HTTP headers, and others. The 'transport' option can also be used to implement custom data loading and saving.
  //
  // Local data sources are bound to a JavaScript array via the 'data' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-data  option.
  //
  //
  //
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
  //
  Transport    Transport

  Template    *Template
}

func ( el Datasource ) getTemplate () string {
  return `{{if .ServerAggregates}}aggregate: [{{range $v := .Aggregate}}{{string $v}},{{end}}],{{end}}
{{if .AutoSync}}autoSync: true,{{end}}
{{if .Batch}}batch: true,{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .Filter) "null" and .ServerFiltering}}{{$length := len .[]FilterLine}}{{if le $length 1}}filter: { {{string .Filter}} },{{else}}filter: [ {{string .Filter}} ],{{end}},{{end}}
{{if ne (string .Group) "null" and .ServerGrouping}}{{$length := len .[]GroupLine}}{{if le $length 1}}group: { {{string .Group}} },{{else}}group: [ {{string .Group}} ],{{end}},{{end}}
{{if ne (string .OfflineStorage) "null"}}offlineStorage: {{string .OfflineStorage}},{{end}}
{{if .ServerPaging }}page: {{.Page}},{{end}}
{{if .ServerPaging }}pageSize: {{.PageSize}},{{end}}
{{if ne (string .Schema) "null"}}schema: {{string .Schema}},{{end}}
{{if .ServerAggregates}}serverAggregates: true,{{end}}
{{if .ServerFiltering}}serverFiltering: true,{{end}}
{{if .ServerGrouping}}serverGrouping: true,{{end}}
{{if .ServerPaging}}serverPaging: true,{{end}}
{{if .ServerSorting}}serverSorting: true,{{end}}
{{if ne (string .Sort) "null"}}{{$length := len .[]SortLine}}{{if le $length 1}}sort: { {{string .Sort}} },{{else}}sort: [ {{string .Sort}} ],{{end}},{{end}}
{{if ne (string .Transport) "null"}}transport: {{string .Transport}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
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
//
//
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
  //
  //
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














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
//
// The data item field to which the filter operator is applied.
//
//
//
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
//
type Filter struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
  //
  // Type: String
  //
  // The data item field to which the filter operator is applied.
  //
  //
  //
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
  //
  Field    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.filters
  //
  // Type: Array
  //
  // The nested filter expressions. Supports the same options as 'filter' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter . Filters can be nested indefinitely.
  //
  //
  //
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
  //
  Filters    Filters

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.logic
  //
  // Type: String
  //
  // The logical operation to use when the 'filter.filters' option is set.
  //
  // The supported values are:
  //
  //
  // *  "and"
  // *  "or"
  //
  //
  //
  //
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
  //
  Logic    LogicEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.operator
  //
  // Type: String
  //
  // The filter operator (comparison).
  //
  // The supported operators are:
  //
  //
  //
  // "eq" (equal to)
  //
  // "neq" (not equal to)
  //
  // "isnull" (is equal to null)
  //
  // "isnotnull" (is not equal to null)
  //
  // "lt" (less than)
  //
  // "lte" (less than or equal to)
  //
  // "gt" (greater than)
  //
  // "gte" (greater than or equal to)
  // *  "startswith"
  // *  "endswith"
  // *  "contains"
  // *  "doesnotcontain"
  // *  "isempty"
  //
  // "isnotempty"
  //
  // The last five are supported only for string fields.
  //
  //
  //
  //
  //
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
  //
  Operator    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.value
  //
  // Type: Object
  //
  // The value to which the 'field' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field  is compared. The value has to be of the same type as the field.
  //
  //    By design, the '"\n"' is removed from the filter before the filtering is performed. That is why an '"\n"' identifier from the filter will not match data items whose corresponding fields contain new lines.
  //
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
  //
  Value    Value

  Template    *Template
}

func ( el Filter ) getTemplate () string {
  return `{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
{{if ne (string .Filters) "null"}}filters: {{string .Filters}},{{end}}
{{if ne (string .Logic) "null"}}logic: {{string .Logic}},{{end}}
{{if ne (string .Operator) "null"}}operator: {{string .Operator}},{{end}}
{{if ne (string .Value) "null"}}value: {{string .Value}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
//
// The aggregates which are calculated during grouping.
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
//
//
//
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
*/
//
type Group struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
  //
  // Type: Array
  //
  // The aggregates which are calculated during grouping.
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
  //
  //
  //
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
  */
  //
  Aggregates    []AggregatesLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.dir
  //
  // Type: String
  // Defalt: asc
  //
  // The sort order of the group.
  //
  // The supported values are:
  //
  //
  // "asc" (ascending order)
  //
  // "desc" (descending order)
  //
  //
  // The default sort order is ascending.
  //
  //
  //
  /*
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
  */
  //
  Dir    DirEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.field
  //
  // Type: String
  //
  // The data item field to group by.
  //
  //
  //
  /*
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
  //
  Field    string

  Template    *Template
}

func ( el Group ) getTemplate () string {
  return `{{if .Aggregates}}aggregates: [{{range $v := .Aggregates}}{{string $v}},{{end}}],{{end}}
{{if ne (string .Dir) "null"}}dir: {{string .Dir}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.aggregate
//
// The name of the aggregate function. Specifies the aggregate function.
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
//
//
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
type Aggregates struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.aggregate
  //
  // Type: String
  //
  // The name of the aggregate function. Specifies the aggregate function.
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
  //
  //
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

func ( el Aggregates ) getTemplate () string {
  return `{{if ne (string .Aggregate) "null"}}aggregate: {{string .Aggregate}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.aggregates
//
// The field from the response which contains the aggregate results. Can be set to a function which is called to return the aggregate results from the response.
//
//    The 'aggregates' option is used only when the 'serverAggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to 'true'.
//
/*
    Example - set aggregates as a string
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        /@ transport configuration @/
      }
      serverAggregates: true,
      schema: {
        aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
      }
    });
    </script>


    Example - set aggregates as a function
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        /@ transport configuration @/
      }
      serverAggregates: true,
      schema: {
        aggregates: function(response) {
          return response.aggregates;
        }
      }
    });
    </script>
*/
//
type Schema struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.aggregates
  //
  // Type: Function
  //
  // The field from the response which contains the aggregate results. Can be set to a function which is called to return the aggregate results from the response.
  //
  //    The 'aggregates' option is used only when the 'serverAggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to 'true'.
  //
  /*
      Example - set aggregates as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        }
        serverAggregates: true,
        schema: {
          aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
        }
      });
      </script>


      Example - set aggregates as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        }
        serverAggregates: true,
        schema: {
          aggregates: function(response) {
            return response.aggregates;
          }
        }
      });
      </script>
  */
  //
  Aggregates    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.data
  //
  // Type: Function
  //
  // The field from the server response which contains the data items. Can be set to a function which is called to return the data items for the response.
  //
  //
  //
  /*
      Example - specify the field which contains the data items as a string
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
          data: "statuses" // twitter's response is { "statuses": [ /@ results @/ ] }
        }
      });
      dataSource.fetch(function(){
        var data = this.data();
        console.log(data.length);
      });
      </script>


      Example - specify the field which contains the data items as a function
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
            return response.statuses; // twitter's response is { "statuses": [ /@ results @/ ] }
          }
        }
      });
      dataSource.fetch(function(){
        var data = this.data();
        console.log(data.length);
      });
      </script>
  */
  //
  // 'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array  The data items from the response.
  //
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.errors
  //
  // Type: Function
  // Defalt: errors
  //
  // The field from the server response which contains server-side errors. Can be set to a function which is called to return the errors for response. If there are any errors, the 'error' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#events-error  event will be fired.
  //
  //
  //
  /*
      Example - specify the error field as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: { q: "#" }
          }
        },
        schema: {
          errors: "error" // twitter's response is { "error": "Invalid query" }
        },
        error: function(e) {
          console.log(e.errors); // displays "Invalid query"
        }
      });
      dataSource.fetch();
      </script>


      Example - specify the error field as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: { q: "#" }
          }
        },
        schema: {
          errors: function(response) {
            return response.error; // twitter's response is { "error": "Invalid query" }
          }
        },
        error: function(e) {
          console.log(e.errors); // displays "Invalid query"
        }
      });
      dataSource.fetch();
      </script>
  */
  //
  Errors    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.groups
  //
  // Type: Function
  //
  // The field from the server response which contains the groups. Can be set to a function which is called to return the groups from the response.
  //
  //    The 'groups' option is used only when the 'serverGrouping' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping option is set to 'true'.
  //
  /*
      Example - set groups as a string
      [{
        aggregates: {
          FIEL1DNAME: {
            FUNCTON1NAME: FUNCTION1VALUE,
            FUNCTON2NAME: FUNCTION2VALUE
          },
          FIELD2NAME: {
            FUNCTON1NAME: FUNCTION1VALUE
          }
        },
        field: FIELDNAME, // the field by which the data items are grouped
        hasSubgroups: true, // true if there are subgroups
        items: [
          // either the subgroups or the data items
          {
            aggregates: {
              //nested group aggregates
            },
            field: NESTEDGROUPFIELDNAME,
            hasSubgroups: false,
            items: [
            // data records
            ],
            value: NESTEDGROUPVALUE
          },
          //nestedgroup2, nestedgroup3, etc.
        ],
        value: VALUE // the group key
      } /@ other groups @/
      ]


      Example - set groups as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverGrouping: true,
        schema: {
          groups: "groups" // groups are returned in the "groups" field of the response
        }
      });
      </script>
  */
  //
  Groups    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.model
  //
  // Type: Object
  //
  // The data item (model) configuration.
  //
  // If set to an object, the 'Model.define' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/model#model.define  method will be used to initialize the data source model.
  //
  // If set to an existing 'kendo.data.Model' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/model  instance, the data source will use that instance and will not initialize a new one.
  //
  //
  //
  /*
      Example - set the model as a JavaScript object
      <script>
      var dataSource = new kendo.data.DataSource({
        schema: {
          model: {
            id: "ProductID",
            fields: {
              ProductID: {
                //this field will not be editable (default value is true)
                editable: false,
                // a defaultValue will not be assigned (default value is false)
                nullable: true
              },
              ProductName: {
                //set validation rules
                validation: { required: true }
              },
              UnitPrice: {
                //data type of the field {Number|String|Boolean|Date} default is String
                type: "number",
                // used when new model is created
                defaultValue: 42,
                validation: { required: true, min: 1 }
              }
            }
          }
        }
      });
      </script>


      Example - set the model as an existing <code>kendo.data.Model</code> instance
      <script>
      var Product = kendo.model.define({
        id: "ProductID",
        fields: {
          ProductID: {
            //this field will not be editable (default value is true)
            editable: false,
            // a defaultValue will not be assigned (default value is false)
            nullable: true
          },
          ProductName: {
            //set validation rules
            validation: { required: true }
          },
          UnitPrice: {
            //data type of the field {Number|String|Boolean|Date} default is String
            type: "number",
            // used when new model is created
            defaultValue: 42,
            validation: { required: true, min: 1 }
          }
        }
      });
      var dataSource = new kendo.data.DataSource({
        schema: {
          model: Product
        }
      });
      </script>
  */
  //
  Model    Model

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.parse
  //
  // Type: Function
  //
  // Executed before the server response is used. Use it to preprocess or parse the server response.
  //
  //
  //
  /*
      Example - data projection
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          }
        },
        schema: {
          parse: function(response) {
            var products = [];
            for (var i = 0; i < response.length; i++) {
              var product = {
                id: response[i].ProductID,
                name: response[i].ProductName
              };
              products.push(product);
            }
            return products;
          }
        }
      });
      dataSource.fetch(function(){
        var data = dataSource.data();
        var product = data[0];
        console.log(product.name); // displays "Chai"
      });
      </script>
  */
  //
  // Parameters:
  // response 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object 'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
  // The initially parsed server response that may need additional modifications.
  //
  // Return:
  // 'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array  The data items from the response.
  //
  Parse    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.total
  //
  // Type: Function
  //
  // The field from the server response which contains the total number of data items. Can be set to a function which is called to return the total number of data items for the response.
  //
  //    The 'schema.total' setting may be omitted when the Grid is bound to a plain 'Array' (that is, the data items' collection is not a value of a field in the server response). In this case, the 'length' of the response 'Array' will be used.
  //    The 'schema.total' must be set if the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true' or the 'schema.data' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.data option is used.
  //
  /*
      Example - set the total as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverGrouping: true,
        schema: {
          total: "total" // total is returned in the "total" field of the response
        }
      });
      </script>


      Example - set the total as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverGrouping: true,
        schema: {
          total: function(response) {
            return response.total; // total is returned in the "total" field of the response
          }
        }
      });
      </script>
  */
  //
  // 'Number' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number  The total number of data items.
  //
  Total    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type
  //
  // Type: String
  // Defalt: json
  //
  // The type of the response.
  //
  // The supported values are:
  //
  //
  // *  "xml"
  // *  "json"
  //
  //
  // By default, the schema interprets the server response as JSON.
  //
  //
  //
  /*
      Example - use XML data
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
              id: "@cover"
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
  //
  Type    TypeEnum

  Template    *Template
}

func ( el Schema ) getTemplate () string {
  return `{{if ne (string .Aggregates) "null"}}aggregates: {{string .Aggregates}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .Errors) "null"}}errors: {{string .Errors}},{{end}}
{{if ne (string .Groups) "null"}}groups: {{string .Groups}},{{end}}
{{if ne (string .Model) "null"}}model: {{string .Model}},{{end}}
{{if ne (string .Parse) "null"}}parse: {{string .Parse}},{{end}}
{{if ne (string .Total) "null"}}total: {{string .Total}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.dir
//
// The sort order (direction).
//
// The supported values are:
//
//
//
// "asc" (ascending order)
//
// "desc" (descending order)
//
//
//
//
/*
    Example - specify the sort order (direction)
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Jane Doe", age: 30 },
        { name: "John Doe", age: 33 }
      ],
      // order by "age" in descending order
      sort: { field: "age", dir: "desc" }
    });
    dataSource.fetch(function(){
      var data = dataSource.view();
      console.log(data[0].age); // displays "33"
    });
    </script>
*/
//
type Sort struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.dir
  //
  // Type: String
  //
  // The sort order (direction).
  //
  // The supported values are:
  //
  //
  //
  // "asc" (ascending order)
  //
  // "desc" (descending order)
  //
  //
  //
  //
  /*
      Example - specify the sort order (direction)
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        // order by "age" in descending order
        sort: { field: "age", dir: "desc" }
      });
      dataSource.fetch(function(){
        var data = dataSource.view();
        console.log(data[0].age); // displays "33"
      });
      </script>
  */
  //
  Dir    DirEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.field
  //
  // Type: String
  //
  // The field by which the data items are sorted.
  //
  //
  //
  /*
      Example - specify the sort field
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        // order by "age" in descending order
        sort: { field: "age", dir: "desc" }
      });
      dataSource.fetch(function(){
        var data = dataSource.view();
        console.log(data[0].age); // displays "33"
      });
      </script>
  */
  //
  Field    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.compare
  //
  // Type: Function
  //
  // Function which can be used for custom comparing of the DataSource items.
  //
  //
  //
  /*
      Example - use a custom compare function to compare items in the DataSource
      <script>
        var numbers = {
          "one"  : 1,
          "two"  : 2,
          "three": 3
        };

        var dataSource = new kendo.data.DataSource({
          data: [
            { id: 1, item: "two" },
            { id: 2, item: "one" },
            { id: 3, item: "three" }
          ],
          sort: { field: "item", dir: "asc", compare: function(a, b) {
            return numbers[a.item] - numbers[b.item];
          }
                }
        });

        $("#grid").kendoGrid({
          dataSource: dataSource,
          sortable: true,
          columns: [{
            field: "item",
            sortable: {
              compare: function(a, b) {
                return numbers[a.item] - numbers[b.item];
              }
            }
          }]
        });
      </script>
  */
  //
  Compare    ComplexJavaScriptType

  Template    *Template
}

func ( el Sort ) getTemplate () string {
  return `{{if ne (string .Dir) "null"}}dir: {{string .Dir}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
{{if ne (string .Compare) "null"}}compare: {{string .Compare}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create
//
// The configuration used when the data source saves newly created data items. Those are items added to the data source via the 'add' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-add  or 'insert' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-insert  methods.
//
//    The data source uses 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax/ to make a HTTP request to the remote service. The value configured via 'transport.create' is passed to 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jquery.ajax/#jQuery-ajax-settings. This means that you can set all options supported by 'jQuery.ajax' via 'transport.create' except the 'success' and 'error' callback functions which are used by the transport.
//
/*
    Example - set the create remote service
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/create
        create: {
          url: "http://demos.telerik.com/kendo-ui/service/products/create",
          dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
        },
        parameterMap: function(data, type) {
          if (type == "create") {
            // send the created data items as the "models" service parameter encoded in JSON
            return { models: kendo.stringify(data.models) };
          }
        }
      },
      batch: true,
      schema: {
        model: { id: "ProductID" }
      }
    });
    // create a new data item
    dataSource.add( { ProductName: "New Product" });
    // save the created data item
    dataSource.sync(); // server response is [{"ProductID":78,"ProductName":"New Product","UnitPrice":0,"UnitsInStock":0,"Discontinued":false}]
    </script>


    Example - set create as a function
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: function(options) {
          /@ implementation omitted for brevity @/
        },
        create: function(options) {
          // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/create
          $.ajax({
            url: "http://demos.telerik.com/kendo-ui/service/products/create",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            // send the created data items as the "models" service parameter encoded in JSON
            data: {
              models: kendo.stringify(options.data.models)
            },
            success: function(result) {
              // notify the data source that the request succeeded
              options.success(result);
            },
            error: function(result) {
              // notify the data source that the request failed
              options.error(result);
            }
          });
        }
      },
      batch: true,
      schema: {
        model: { id: "ProductID" }
      }
    });
    dataSource.add( { ProductName: "New Product" });
    dataSource.sync();
    </script>
*/
//
type Transport struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create
  //
  // Type: Object
  //
  // The configuration used when the data source saves newly created data items. Those are items added to the data source via the 'add' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-add  or 'insert' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-insert  methods.
  //
  //    The data source uses 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax/ to make a HTTP request to the remote service. The value configured via 'transport.create' is passed to 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jquery.ajax/#jQuery-ajax-settings. This means that you can set all options supported by 'jQuery.ajax' via 'transport.create' except the 'success' and 'error' callback functions which are used by the transport.
  //
  /*
      Example - set the create remote service
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/create
          create: {
            url: "http://demos.telerik.com/kendo-ui/service/products/create",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "create") {
              // send the created data items as the "models" service parameter encoded in JSON
              return { models: kendo.stringify(data.models) };
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      // create a new data item
      dataSource.add( { ProductName: "New Product" });
      // save the created data item
      dataSource.sync(); // server response is [{"ProductID":78,"ProductName":"New Product","UnitPrice":0,"UnitsInStock":0,"Discontinued":false}]
      </script>


      Example - set create as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: function(options) {
            /@ implementation omitted for brevity @/
          },
          create: function(options) {
            // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/create
            $.ajax({
              url: "http://demos.telerik.com/kendo-ui/service/products/create",
              dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
              // send the created data items as the "models" service parameter encoded in JSON
              data: {
                models: kendo.stringify(options.data.models)
              },
              success: function(result) {
                // notify the data source that the request succeeded
                options.success(result);
              },
              error: function(result) {
                // notify the data source that the request failed
                options.error(result);
              }
            });
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.add( { ProductName: "New Product" });
      dataSource.sync();
      </script>
  */
  //
  Create    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy
  //
  // Type: Object
  //
  // The configuration used when the data source destroys data items. Those are items removed from the data source via the 'remove' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-remove  method.
  //
  //    The data source uses 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax to make an HTTP request to the remote service. The value configured via 'transport.destroy' is passed to 'jQuery.ajax'. This means that you can set all options supported by 'jQuery.ajax' via 'transport.destroy' except the 'success' and 'error' callback functions which are used by the transport.
  //
  /*
      Example - set the destroy remote service
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          },
          // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/destroy
          destroy: {
            url: "http://demos.telerik.com/kendo-ui/service/products/destroy",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "destroy") {
              // send the destroyed data items as the "models" service parameter encoded in JSON
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        // remove the first data item
        dataSource.remove(products[0]);
        // send the destroyed data item to the remote service
        dataSource.sync();
      });
      </script>


      Example - set destroy as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: function(options) {
            $.ajax({
              url: "http://demos.telerik.com/kendo-ui/service/products",
              dataType: "jsonp",
              success: function(result) {
                options.success(result);
              }
            });
          },
          destroy: function (options) {
            // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/destroy
            $.ajax({
              url: "http://demos.telerik.com/kendo-ui/service/products/destroy",
              dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
              // send the destroyed data items as the "models" service parameter encoded in JSON
              data: {
                models: kendo.stringify(options.data.models)
              },
              success: function(result) {
                // notify the data source that the request succeeded
                options.success(result);
              },
              error: function(result) {
                // notify the data source that the request failed
                options.error(result);
              }
            });
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        dataSource.remove(products[0]);
        dataSource.sync();
      });
      </script>
  */
  //
  Destroy    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap
  //
  // Type: Function
  //
  // The function which converts the request parameters to a format suitable for the remote service. By default, the data source sends the parameters using jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  //    The 'parameterMap' method is often used to encode the parameters in JSON format.
  //    The 'parameterMap' function will not be called when using custom functions for the read, update, create, and destroy operations.
  //
  /*
      Example - convert data source request parameters
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders?$format=json",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            jsonp: "$callback",
            cache: true
          },
          parameterMap: function(data, type) {
            if (type == "read") {
              // send take as "$top" and skip as "$skip"
              return {
                $top: data.take,
                $skip: data.skip
              }
            }
          }
        },
        schema: {
          data: "d"
        },
        pageSize: 20,
        serverPaging: true // enable serverPaging so take and skip are sent as request parameters
      });
      dataSource.fetch(function() {
        console.log(dataSource.view().length); // displays "20"
      });
      </script>


      Example - send request parameters as JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            url: "http://demos.telerik.com/kendo-ui/service/products/create",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            return kendo.stringify(data);
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.add( { ProductName: "New Product" });
      dataSource.sync();
      </script>
  */
  //
  // Parameters:
  // data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  // The parameters which will be sent to the remote service. The value specified in the data field of the transport settings (create, read, update or destroy) is included as well. If 'Array' #batch-boolean-default">batch is set to false, the fields of the changed data items are also included.
  // data.aggregate aggregate option. Available if the serverAggregates option is set to true and the data source makes a "read" request.
  // data.group group option. Available if the serverGrouping option is set to true and the data source makes a "read" request.
  // data.filter filter option. Available if the serverFiltering option is set to true and the data source makes a "read" request.
  // data.models batch option is set to true.
  // data.page serverPaging option is set to true and the data source makes a "read" request.
  // data.pageSize pageSize option. Available if the serverPaging option is set to true and the data source makes a "read" request.
  // data.skip serverPaging option is set to true and the data source makes a "read" request.
  // data.sort sort option. Available if the serverSorting option is set to true and the data source makes a "read" request.
  // data.take serverPaging option is set to true and the data source makes a "read" request.
  // type
  //
  // Return:
  // 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object  the request parameters converted to a format required by the remote service.
  //
  ParameterMap    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.push
  //
  // Type: Function
  //
  // The function invoked during transport initialization which sets up push notifications. The data source will call this function only once and provide callbacks which will handle push notifications (data pushed from the server).
  //
  //
  //
  /*
      Example
      <script src="http://ajax.aspnetcdn.com/ajax/signalr/jquery.signalr-1.1.3.min.js"></script>
      <script>
      var hubUrl = "http://demos.telerik.com/kendo-ui/service/signalr/hubs";
      var connection = $.hubConnection(hubUrl, { useDefaultPath: false});
      var hub = connection.createHubProxy("productHub");
      var hubStart = connection.start({ jsonp: true });
      var dataSource = new kendo.data.DataSource({
      transport: {
        push: function(callbacks) {
          hub.on("create", function(result) {
            console.log("push create");
            callbacks.pushCreate(result);
          });
          hub.on("update", function(result) {
            console.log("push update");
            callbacks.pushUpdate(result);
          });
          hub.on("destroy", function(result) {
            console.log("push destroy");
            callbacks.pushDestroy(result);
          });
        }
      },
      schema: {
        model: {
          id: "ID",
          fields: {
            "ID": { editable: false },
            "CreatedAt": { type: "date" },
            "UnitPrice": { type: "number" }
          }
        }
      }
      });
      </script>
  */
  //
  // Parameters:
  // callbacks 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  // An object containing callbacks for notifying the data source of push notifications.
  // callbacks.pushCreate 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  // Function that should be invoked to notify the data source about newly created data items that are pushed from the server. Accepts a single argument - the object pushed from the server which should follow the schema.data configuration.
  // callbacks.pushDestroy 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  // Function that should be invoked to notify the data source about destroyed data items that are pushed from the server. Accepts a single argument - the object pushed from the server
  // which should follow the schema.data configuration.
  // callbacks.pushUpdate 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  // Function that should be invoked to notify the data source about updated data items that are pushed from the server. Accepts a single argument - the object pushed from the server
  // which should follow the schema.data configuration.
  //
  Push    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read
  //
  // Type: Object
  //
  // The configuration used when the data source loads data items from a remote service.
  //
  //    The data source uses 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax to make a HTTP request to the remote service. The value configured via 'transport.read' is passed to 'jQuery.ajax'. This means that you can set all options supported by 'jQuery.ajax' via 'transport.read' except the 'success' and 'error' callback functions which are used by the transport.
  //
  /*
      Example - set the read remote service
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          // make JSONP request to http://demos.telerik.com/kendo-ui/service/products
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        }
      });
      dataSource.fetch(function() {
        console.log(dataSource.view().length); // displays "77"
      });
      </script>


      Example - send additional parameters to the remote service
      <input value="html5" id="search" />
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: {
              q: $("#search").val() // send the value of the #search input to the remote service
            }
          }
        }
      });
      dataSource.fetch();
      </script>


      Example - set read as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: function(options) {
            // make JSONP request to http://demos.telerik.com/kendo-ui/service/products
            $.ajax({
              url: "http://demos.telerik.com/kendo-ui/service/products",
              dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
              success: function(result) {
                // notify the data source that the request succeeded
                options.success(result);
              },
              error: function(result) {
                // notify the data source that the request failed
                options.error(result);
              }
            });
          }
        }
      });
      dataSource.fetch(function() {
        console.log(dataSource.view().length); // displays "77"
      });
      </script>
  */
  //
  Read    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr
  //
  // Type: Object
  //
  // The configuration used when 'type' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-type  is set to "signalr". Configures the SignalR settings - hub, connection promise, server, and client hub methods.
  //
  // Live demo available at demos.telerik.com/kendo-ui http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://demos.telerik.com/kendo-ui/grid/signalr .
  //
  // It is recommended to get familiar with the SignalR JavaScript API http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://www.asp.net/signalr/overview/guide-to-the-api/hubs-api-guide-javascript-client .
  //
  //
  //
  /*
      Example
      <script src="http://ajax.aspnetcdn.com/ajax/signalr/jquery.signalr-1.1.3.min.js"></script>
      <script>
          var hubUrl = "http://demos.telerik.com/kendo-ui/service/signalr/hubs";
          var connection = $.hubConnection(hubUrl, { useDefaultPath: false});
          var hub = connection.createHubProxy("productHub");
          var hubStart = connection.start({ jsonp: true });

          var dataSource = new kendo.data.DataSource({
              type: "signalr",
              schema: {
                  model: {
                      id: "ID",
                      fields: {
                          "ID": { editable: false, nullable: true },
                          "CreatedAt": { type: "date" },
                          "UnitPrice": { type: "number" }
                      }
                  }
              },
              transport: {
                  signalr: {
                      promise: hubStart,
                      hub: hub,
                      server: {
                          read: "read",
                          update: "update",
                          destroy: "destroy",
                          create: "create"
                      },
                      client: {
                          read: "read",
                          update: "update",
                          destroy: "destroy",
                          create: "create"
                      }
                  }
              }
          });
      </script>
  */
  //
  Signalr    Signalr

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.submit
  //
  // Type: Function
  //
  // A function that will handle create, update and delete operations in a single batch.
  //
  // Typically, you have the 'transport.read' and 'transport.submit' operations defined together. The 'transport.create', 'transport.update' and 'transport.delete' operations will not be executed in this case.
  //
  //    This function will only be invoked when the DataSource is in batch mode.
  //

  //
  // Parameters:
  // e.data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  // An object containing the created (e.data.created), updated (e.data.updated), and destroyed (e.data.destroyed) items.
  // e.success 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  // A callback that should be called for each operation with two parameters - items and operation. See example below.
  // e.fail 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  // A callback that should be called in case of failure of any of the operations.
  // Example - declare transport submit function
  //
  Submit    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update
  //
  // Type: Object
  //
  // The configuration used when the data source saves updated data items. Those are data items whose fields have been updated.
  //
  //    The data source uses 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax to make a HTTP request to the remote service. The value configured via 'transport.update' is passed to 'jQuery.ajax'. This means that you can set all options supported by 'jQuery.ajax' via 'transport.update' except the 'success' and 'error' callback functions which are used by the transport.
  //
  /*
      Example - specify update as a string
      <script>
      var dataSource = new kendo.data.DataSource({
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
        product.set("UnitPrice", 20);
        dataSource.sync(); makes request to http://demos.telerik.com/kendo-ui/service/products/update
      });
      </script>


      Example - specify update as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: function(options) {
            /@ implementation omitted for brevity @/
          },
          update: function(options) {
            // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/update
            $.ajax({
              url: "http://demos.telerik.com/kendo-ui/service/products/update",
              dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
              // send the updated data items as the "models" service parameter encoded in JSON
              data: {
                models: kendo.stringify(options.data.models)
              },
              success: function(result) {
                // notify the data source that the request succeeded
                options.success(result);
              },
              error: function(result) {
                // notify the data source that the request failed
                options.error(result);
              }
            });
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20);
        dataSource.sync(); //makes request to http://demos.telerik.com/kendo-ui/service/products/update
      });
      </script>
  */
  //
  Update    ComplexJavaScriptType

  Template    *Template
}

func ( el Transport ) getTemplate () string {
  return `{{if ne (string .Create) "null"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "null"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .ParameterMap) "null"}}parameterMap: {{string .ParameterMap}},{{end}}
{{if ne (string .Push) "null"}}push: {{string .Push}},{{end}}
{{if ne (string .Read) "null"}}read: {{string .Read}},{{end}}
{{if ne (string .Signalr) "null"}}signalr: {{string .Signalr}},{{end}}
{{if ne (string .Submit) "null"}}submit: {{string .Submit}},{{end}}
{{if ne (string .Update) "null"}}update: {{string .Update}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further info.
//
//
//
/*
    Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        create: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Create struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further info.
  //
  //
  //
  /*
      Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set a content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.data
  //
  // Type: Object
  //
  // Additional parameters that are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            data: {
              name: "Jane Doe",
              age: 30
            }
          }
        }
      });
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            data: function() {
              return {
                name: "Jane Doe",
                age: 30
              }
            }
          }
        }
      });
      </script>
  */
  //
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.type
  //
  // Type: String
  // Defalt: GET
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
  //
  /*
      Example - set the HTTP verb of the request
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify the URL as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            url: "http://demos.telerik.com/kendo-ui/service/products/create",
            cache: true,
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "create") {
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.add( { ProductName: "New Product" });
      dataSource.sync();
      </script>


      Example - specify the URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            url: function(options) {
              return "http://demos.telerik.com/kendo-ui/service/products/create"
            },
            cache: true,
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "create") {
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.add( { ProductName: "New Product" });
      dataSource.sync();
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Create ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
//
//
//
/*
    Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        destroy: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Destroy struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the jQuery.ajax http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set the content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.data
  //
  // Type: Object
  //
  // Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            data: {
              name: "Jane Doe",
              age: 30
            }
          }
        }
      });
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            data: function() {
              return {
                name: "Jane Doe",
                age: 30
              }
            }
          }
        }
      });
      </script>
  */
  //
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.type
  //
  // Type: String
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
  //
  /*
      Example
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify the URL as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          },
          destroy: {
            url: "http://demos.telerik.com/kendo-ui/service/products/destroy",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "destroy") {
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        dataSource.remove(products[0]);
        dataSource.sync();
      });
      </script>


      Example - specify the URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          },
          destroy: {
            url: function (options) {
              return "http://demos.telerik.com/kendo-ui/service/products/destroy",
            },
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "destroy") {
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        dataSource.remove(products[0]);
        dataSource.sync();
      });
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Destroy ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting cache to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
//
//
//
/*
    Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Read struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting cache to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.data
  //
  // Type: Object
  //
  // Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: {
              q: "html5" // send "html5" as the "q" parameter
            }
          }
        }
      });
      dataSource.fetch();
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: function() {
              return {
                q: "html5" // send "html5" as the "q" parameter
              };
            }
          }
        }
      });
      dataSource.fetch();
      </script>
  */
  //
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.type
  //
  // Type: String
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to "jsonp". JSONP always uses GET requests.
  //
  /*
      Example - set the HTTP verb of the request
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify URL as a string
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
        console.log(dataSource.view().length); // displays "77"
      });
      </script>


      Example - specify URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: function(options) {
              return "http://demos.telerik.com/kendo-ui/service/products",
            }
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        }
      });
      dataSource.fetch(function() {
        console.log(dataSource.view().length); // displays "77"
      });
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Read ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client
//
// Specifies the client-side CRUD methods of the SignalR hub.
//
//
//

//
type Signalr struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client
  //
  // Type: Object
  //
  // Specifies the client-side CRUD methods of the SignalR hub.
  //
  //
  //

  //
  Client    Client

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.hub
  //
  // Type: Object
  //
  // The SignalR hub object returned by the 'createHubProxy' method. The 'hub' option is mandatory.
  //
  //
  //

  //
  Hub    Hub

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.promise
  //
  // Type: Object
  //
  // The promise returned by the 'start' method of the SignalR connection. The 'promise' option is mandatory.
  //
  //
  //

  //
  Promise    Promise

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server
  //
  // Type: Object
  //
  // Specifies the server-side CRUD methods of the SignalR hub.
  //
  //
  //

  //
  Server    Server

  Template    *Template
}

func ( el Signalr ) getTemplate () string {
  return `{{if ne (string .Client) "null"}}client: {{string .Client}},{{end}}
{{if ne (string .Hub) "null"}}hub: {{string .Hub}},{{end}}
{{if ne (string .Promise) "null"}}promise: {{string .Promise}},{{end}}
{{if ne (string .Server) "null"}}server: {{string .Server}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.create
//
// Specifies the name of the client-side method of the SignalR hub responsible for creating data items.
//
//
//

//
type Client struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.create
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for creating data items.
  //
  //
  //

  //
  Create    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.destroy
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for destroying data items.
  //
  //
  //

  //
  Destroy    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.read
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for reading data items.
  //
  //
  //

  //
  Read    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.update
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for updating data items.
  //
  //
  //

  //
  Update    string

  Template    *Template
}

func ( el Client ) getTemplate () string {
  return `{{if ne (string .Create) "null"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "null"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .Read) "null"}}read: {{string .Read}},{{end}}
{{if ne (string .Update) "null"}}update: {{string .Update}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.create
//
// Specifies the name of the server-side method of the SignalR hub responsible for creating data items.
//
//
//

//
type Server struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.create
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for creating data items.
  //
  //
  //

  //
  Create    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.destroy
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for destroying data items.
  //
  //
  //

  //
  Destroy    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.read
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for reading data items.
  //
  //
  //

  //
  Read    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.update
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for updating data items.
  //
  //
  //

  //
  Update    string

  Template    *Template
}

func ( el Server ) getTemplate () string {
  return `{{if ne (string .Create) "null"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "null"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .Read) "null"}}read: {{string .Read}},{{end}}
{{if ne (string .Update) "null"}}update: {{string .Update}},{{end}}
`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
//
//
//
/*
    Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        update: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Update struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. Default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON.
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.data
  //
  // Type: Object
  //
  // Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for
  // sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            data: {
              name: "Jane Doe",
              age: 30
            }
          }
        }
      });
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            data: function() {
              return {
                name: "Jane Doe",
                age: 30
              }
            }
          }
        }
      });
      </script>
  */
  //
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.type
  //
  // Type: String
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
  //
  /*
      Example - set the HTTP verb of the request
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify URL as a string
      <script>
      var dataSource = new kendo.data.DataSource({
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
        product.set("UnitPrice", 20);
        dataSource.sync();
      });
      </script>


      Example - specify URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read:  {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          update: {
            url: function(options) {
              return "http://demos.telerik.com/kendo-ui/service/products/update",
            },
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        },
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20);
        dataSource.sync();
      });
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Update ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}













// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
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
//
//
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


// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.logic
//
// The logical operation to use when the 'filter.filters' option is set.
//
// The supported values are:
//
//
// *  "and"
// *  "or"
//
//
//
//
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
//
type LogicEnum int

const(
  LOGIC_NIL LogicEnum   = iota
  LOGIC_AND
  LOGIC_OR
)

var LogicEnums  = [...]string{
  "",
  "and",
  "or",
}

func (el LogicEnum ) String() string {
  return LogicEnums[ el ]
}


// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.operator
//
// The filter operator (comparison).
//
// The supported operators are:
//
//
//
// "eq" (equal to)
//
// "neq" (not equal to)
//
// "isnull" (is equal to null)
//
// "isnotnull" (is not equal to null)
//
// "lt" (less than)
//
// "lte" (less than or equal to)
//
// "gt" (greater than)
//
// "gte" (greater than or equal to)
// *  "startswith"
// *  "endswith"
// *  "contains"
// *  "doesnotcontain"
// *  "isempty"
//
// "isnotempty"
//
// The last five are supported only for string fields.
//
//
//
//
//
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
//
type OperatorEnum int

const(
  OPERATOR_NIL OperatorEnum   = iota
  OPERATOR_EQ
  OPERATOR_NEQ
  OPERATOR_ISNULL
  OPERATOR_ISNOTNULL
  OPERATOR_LT
  OPERATOR_LTE
  OPERATOR_GT
  OPERATOR_GTE
  OPERATOR_STARTSWITH
  OPERATOR_ENDSWITH
  OPERATOR_CONTAINS
  OPERATOR_DOESNOTCONTAIN
  OPERATOR_ISEMPTY
)

var OperatorEnums  = [...]string{
  "",
  "eq",
  "neq",
  "isnull",
  "isnotnull",
  "lt",
  "lte",
  "gt",
  "gte",
  "startswith",
  "endswith",
  "contains",
  "doesnotcontain",
  "isempty",
}

func (el OperatorEnum ) String() string {
  return OperatorEnums[ el ]
}


// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
//
// The aggregates which are calculated during grouping.
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
//
//
//
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
*/
//
type AggregatesEnum int

const(
  AGGREGATES_NIL AggregatesEnum   = iota
  AGGREGATES_AVERAGE
  AGGREGATES_COUNT
  AGGREGATES_MAX
  AGGREGATES_MIN
  AGGREGATES_SUM
)

var AggregatesEnums  = [...]string{
  "",
  "average",
  "count",
  "max",
  "min",
  "sum",
}

func (el AggregatesEnum ) String() string {
  return AggregatesEnums[ el ]
}


// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.dir
//
// The sort order of the group.
//
// The supported values are:
//
//
// "asc" (ascending order)
//
// "desc" (descending order)
//
//
// The default sort order is ascending.
//
//
//
/*
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
*/
//
type DirEnum int

const(
  DIR_NIL DirEnum   = iota
  DIR_ASC
  DIR_DESC
)

var DirEnums  = [...]string{
  "",
  "asc",
  "desc",
}

func (el DirEnum ) String() string {
  return DirEnums[ el ]
}


// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type
//
// The type of the response.
//
// The supported values are:
//
//
// *  "xml"
// *  "json"
//
//
// By default, the schema interprets the server response as JSON.
//
//
//
/*
    Example - use XML data
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
            id: "@cover"
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
//
type TypeEnum int

const(
  TYPE_NIL TypeEnum   = iota
  TYPE_XML
  TYPE_JSON
)

var TypeEnums  = [...]string{
  "",
  "xml",
  "json",
}

func (el TypeEnum ) String() string {
  return TypeEnums[ el ]
}


// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap
//
// The function which converts the request parameters to a format suitable for the remote service. By default, the data source sends the parameters using jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
//
//    The 'parameterMap' method is often used to encode the parameters in JSON format.
//    The 'parameterMap' function will not be called when using custom functions for the read, update, create, and destroy operations.
//
/*
    Example - convert data source request parameters
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: {
          url: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders?$format=json",
          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          jsonp: "$callback",
          cache: true
        },
        parameterMap: function(data, type) {
          if (type == "read") {
            // send take as "$top" and skip as "$skip"
            return {
              $top: data.take,
              $skip: data.skip
            }
          }
        }
      },
      schema: {
        data: "d"
      },
      pageSize: 20,
      serverPaging: true // enable serverPaging so take and skip are sent as request parameters
    });
    dataSource.fetch(function() {
      console.log(dataSource.view().length); // displays "20"
    });
    </script>


    Example - send request parameters as JSON
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        create: {
          url: "http://demos.telerik.com/kendo-ui/service/products/create",
          dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
        },
        parameterMap: function(data, type) {
          return kendo.stringify(data);
        }
      },
      batch: true,
      schema: {
        model: { id: "ProductID" }
      }
    });
    dataSource.add( { ProductName: "New Product" });
    dataSource.sync();
    </script>
*/
//
// Parameters:
// data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
// The parameters which will be sent to the remote service. The value specified in the data field of the transport settings (create, read, update or destroy) is included as well. If 'Array' #batch-boolean-default">batch is set to false, the fields of the changed data items are also included.
// data.aggregate aggregate option. Available if the serverAggregates option is set to true and the data source makes a "read" request.
// data.group group option. Available if the serverGrouping option is set to true and the data source makes a "read" request.
// data.filter filter option. Available if the serverFiltering option is set to true and the data source makes a "read" request.
// data.models batch option is set to true.
// data.page serverPaging option is set to true and the data source makes a "read" request.
// data.pageSize pageSize option. Available if the serverPaging option is set to true and the data source makes a "read" request.
// data.skip serverPaging option is set to true and the data source makes a "read" request.
// data.sort sort option. Available if the serverSorting option is set to true and the data source makes a "read" request.
// data.take serverPaging option is set to true and the data source makes a "read" request.
// type
//
// Return:
// 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object  the request parameters converted to a format required by the remote service.
//
type ParameterMapEnum int

const(
  PARAMETERMAP_NIL ParameterMapEnum   = iota
  PARAMETERMAP_CREATE
  PARAMETERMAP_READ
  PARAMETERMAP_UPDATE
  PARAMETERMAP_DESTROY
)

var ParameterMapEnums  = [...]string{
  "",
  "create",
  "read",
  "update",
  "destroy",
}

func (el ParameterMapEnum ) String() string {
  return ParameterMapEnums[ el ]
}

