package kendo

// The configuration used to parse the remote service response.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
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
       return response.results; // twitter's response is { "results": [ // results ] }
      }
     }
    });
    dataSource.fetch(function(){
     var data = this.data();
     console.log(data.length);
    });
    </script>
*/
type Schema struct {
  // The field from the response which contains the aggregate results. Can be set to a function which is called to return the aggregate results from the response.
  //
  // The 'aggregates' option is used only when the 'serverAggregates' ( http://docs.telerik.com#configuration-serverAggregates ) option is set to 'true'.
  //
  // The result of the function should be a JavaScript object which contains the aggregate results for every field in the following format:
  /*
      {
        Field1Name: {
          Function1Name: Function1Value,
          Function2Name: Function2Value
        },
        Field2Name: {
          Function1Name: Function1Value
        }
      }
  */
  //
  // For example, if the data source is configured like this:
  /*
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        }
        serverAggregates: true,
        aggregate: [
          { field: "unitPrice", aggregate: "max" },
          { field: "unitPrice", aggregate: "min" },
          { field: "ProductName", aggregate: "count" }
        ]
      });
  */
  //
  // The aggregate results should have the following format:
  /*
      {
        unitPrice: {
            max: 100,
            min: 1
        },
        productName: {
            count: 42
        }
      }
   */
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.aggregates
  /*
      Example - set aggregates as a string
      <script>
      var dataSource = new kendo.data.DataSource({
       transport: {
        // transport configuration
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
        // transport configuration
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
  Aggregates        []AggregateList

  // The field from the server response which contains the data items. Can be set to a function which is called to return the data items for the response.
  //
  // RETURNS
  // Array The data items from the response.
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
        data: "statuses" // twitter's response is { "statuses": [ // results ] }
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
         return response.statuses; // twitter's response is { "statuses": [ // results ] }
        }
       }
      });
      dataSource.fetch(function(){
       var data = this.data();
       console.log(data.length);
      });
      </script>
  */
  Data              ComplexJavaScriptType

  // default: "errors"
  //
  // The field from the server response which contains server-side errors. Can be set to a function which is called to return the errors for response. If there are any errors, the 'error' ( http://docs.telerik.comhttp://docs.telerik.com/kendo-ui/api/javascript/data/datasource#events-error ) event will be fired.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.errors
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
  Errors            ComplexJavaScriptType

  // The field from the server response which contains the groups. Can be set to a function which is called to return the groups from the response.
  //
  // The 'groups' option is used only when the 'serverGrouping' ( http://docs.telerik.com#configuration-serverGrouping ) option is set to 'true'.
  //
  // The result should have the following format:
  /*
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
        } // other groups
      ]
  */
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.groups
  /*
      Example - set groups as a string
      <script>
      var dataSource = new kendo.data.DataSource({
       transport: {
        // transport configuration
       },
       serverGrouping: true,
       schema: {
        groups: "groups" // groups are returned in the "groups" field of the response
       }
      });
      </script>


      Example - set groups as a function
      <script>
      var dataSource = new kendo.data.DataSource({
       transport: {
        // transport configuration
       },
       serverGrouping: true,
       schema: {
        groups: function(response) {
         return response.groups; // groups are returned in the "groups" field of the response
        }
       }
      });
      </script>
  */
  Groups            Group

  // The data item (model) configuration.
  //
  // If set to an object, the 'Model.define' ( http://docs.telerik.com/kendo-ui/api/javascript/data/model#model.define ) method will be used to initialize the data source model.
  //
  // If set to an existing 'kendo.data.Model' ( http://docs.telerik.com/kendo-ui/api/javascript/data/model ) instance, the data source will use that instance and will *not* initialize a new one.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.model
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


      Example - set the model as an existing kendo.data.Model instance
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
  Model           Model//fazerrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr

  // Executed before the server response is used. Use it to preprocess or parse the server response.
  //
  // The initially parsed server response that may need additional modifications.
  //
  // 'Array' The data items from the response.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.parse
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
  Parse             ComplexJavaScriptType

  // The field from the server response which contains the total number of data items. Can be set to a function which is called to return the total number of data items for the response.
  //
  // The 'schema.total' setting may be omitted when the Grid is bound to a plain 'Array' (that is, the data items' collection is not a value of a field in the server response). In this case, the 'length' of the response 'Array' will be used.
  //
  // The 'schema.total' must be set if the 'serverPaging' ( http://docs.telerik.comhttp://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging ) option is set to 'true' or the 'schema.data' ( http://docs.telerik.comhttp://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.data ) option is used.
  //
  // 'Number' The total number of data items.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.total
  /*
      Example - set the total as a string
      <script>
      var dataSource = new kendo.data.DataSource({
       transport: {
        // transport configuration
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
        // transport configuration
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
  Total             ComplexJavaScriptType

  // default: "json"
  //
  // The type of the response.
  //
  // The supported values are: "xml" OR "json"
  //
  // By default, the schema interprets the server response as JSON.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type
  /*
      xample - use XML data
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
  Type              DataXmlOrJSonEnum
}
