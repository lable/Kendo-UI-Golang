package kendo

type DataSource  struct{
  // The aggregates which are calculated when the data source populates with data.
  //
  // The data source calculates aggregates client-side unless the serverAggregates
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to
  // true.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
  /*
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
  Aggregate       []AggregateList

  // If set to true the data source would automatically save any changed data items by calling the sync method. By
  // default, changes are not automatically saved.
  //
  // default: false
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-autoSync
  /*
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
  AutoSync        bool

  // If set to true, the data source will batch CRUD operation requests. For example, updating two data items would
  // cause one HTTP request instead of two. By default, the data source makes a HTTP request for every CRUD operation.
  //
  // default: false
  //
  // The changed data items are sent as models by default. This can be changed via the parameterMap option.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-batch
  /*
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
  Batch           bool

  // The JSon of data items which the data source contains. The data source will wrap those items as
  // kendo.data.ObservableObject or kendo.data.Model (if schema.model is set).
  //
  // the XML element which represents a single data record
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-data
  /*
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
  Data            ComplexJavaScriptType

  // The filters which are applied over the data items. By default, no filter is applied.
  //
  // The data source filters the data items client-side unless the serverFiltering option is set to true.
  //
  // Example: filter:
  // {
  //    field: "name", operator: "startswith", value: "Jane"
  // }
  //
  // Example: filter:
  // [
  //    leave data items which are "Beverage" and not "Coffee"
  //    { field: "category", operator: "eq", value: "Beverages" },
  //    { field: "name", operator: "neq", value: "Coffee" }
  // ]
  //
  // Example: filter:
  // {
  //    leave data items which are "Food" or "Tea"
  //    logic: "or",
  //    filters: [
  //      { field: "category", operator: "eq", value: "Food" },
  //      { field: "name", operator: "eq", value: "Tea" }
  //    ]
  // }
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter
  /*
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
  Filter          Filter

  // The grouping configuration of the data source. If set, the data items will be grouped when the data source is populated. By default, grouping is not applied.
  //
  //The data source groups the data items client-side unless the serverGrouping option is set to true.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group
  /*
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
  Group           Group

  // The offline storage unique key id
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-offlineStorage
  /*
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
  OfflineStorage  string

  // The page of data which the data source will return when the view method is invoked or request from the remote
  // service.
  //
  // The data source will page the data items client-side unless the serverPaging option is set to true.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-page
  /*
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
  Page            int64

  // The number of data items per page. The property has no default value. That is why to use paging, make sure some
  // pageSize value is set.
  //
  // The data source will page the data items client-side unless the serverPaging option is set to true.
  /*
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
  PageSize        int64

  // The configuration used to parse the remote service response.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
  /*
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
  Schema            Schema

  // If set to true, the data source will leave the aggregate calculation to the remote service. By default, the data
  // source calculates aggregates client-side.
  //
  // Configure schema.aggregates if you set serverAggregates to true.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the
  // DataSource.
  //
  // default: false
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates
  /*
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
  ServerAggregates  bool

  // If set to true, the data source will leave the filtering implementation to the remote service. By default, the data
  // source performs filtering client-side.
  //
  // By default, the filter is sent to the server following jQuery's conventions.
  //
  // For example, the filter { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] } is
  // sent as:
  //
  // *  filter[logic]: and
  // *  filter[filters][0][field]: name
  // *  filter[filters][0][operator]: startswith
  // *  filter[filters][0][value]: Jane
  //
  // Use the parameterMap option to send the filter option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the
  // DataSource.
  //
  // default: false
  /*
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
  ServerFiltering   bool

  // If set to true, the data source will leave the grouping implementation to the remote service. By default, the data
  // source performs grouping client-side.
  //
  // By default, the group is sent to the server following jQuery's conventions.
  //
  // For example, the group { field: "category", dir: "desc" } is sent as:
  //
  // *  group[0][field]: category
  // *  group[0][dir]: desc
  //
  // Use the parameterMap option to send the group option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the
  // DataSource.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping
  ServerGrouping    bool

  // If set to true, the data source will leave the data item paging implementation to the remote service. By default,
  // the data source performs paging client-side.
  //
  // Configure schema.total if you set serverPaging to true. In addition, pageSize should be set no matter if paging is
  // performed client-side or server-side.
  //
  // The following options are sent to the server when server paging is enabled:
  //
  // *  page - the page of data item to return (1 means the first page).
  // *  pageSize - the number of items to return.
  // *  skip - how many data items to skip.
  // *  take - the number of data items to return (the same as pageSize).
  //
  // Use the parameterMap option to send the paging options in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the
  // DataSource.
  //
  // default: false
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging
  /*
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
  ServerPaging      bool

  // If set to true, the data source will leave the data item sorting implementation to the remote service. By default,
  // the data source performs sorting client-side.
  //
  // By default, the sort is sent to the server following jQuery's conventions.
  //
  // For example, the sort { field: "age", dir: "desc" } is sent as:
  //
  // *  sort[0][field]: age
  // *  sort[0][dir]: desc
  //
  // Use the parameterMap option to send the paging options in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the
  // DataSource.
  //
  // default: false
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting
  /*
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
  ServerSorting     bool

  // The sort order which will be applied over the data items. By default the data items are not sorted.
  //
  // The data source sorts the data items client-side unless the serverSorting option is set to true.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort
  /*
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
  Sort              Sort

  // The configuration used to load and save the data items. A data source is remote or local based on the way of it
  // retrieves data items.
  //
  // Remote data sources load and save data items from and to a remote end-point (also known as remote service or server ).
  //
  // The transport option describes the remote service configuration - URL, HTTP verb, HTTP headers, and others. The
  // transport option can also be used to implement custom data loading and saving.
  //
  // Local data sources are bound to a JavaScript array via the data option.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport
  /*
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
  Transport         DataJSonOrJSonpEnum
}