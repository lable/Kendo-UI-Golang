package kendo

















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

  The 'schema.total' must be set if the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true' or the 'schema.data' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.data option is used.
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
Type    string

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
  Dir    string

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
  The 'parameterMap' function will not be called when using custom functions for the read, update, create, and destroy operations.
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
  Create    Create

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.destroy
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for destroying data items.
  //
  //
  //

  //
  Destroy    Destroy

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.read
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for reading data items.
  //
  //
  //

  //
  Read    Read

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.update
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for updating data items.
  //
  //
  //

  //
  Update    Update

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
  Create    Create

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.destroy
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for destroying data items.
  //
  //
  //

  //
  Destroy    Destroy

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.read
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for reading data items.
  //
  //
  //

  //
  Read    Read

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.update
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for updating data items.
  //
  //
  //

  //
  Update    Update

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
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/%60jQuery.ajax%60  documentation for further information.
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













