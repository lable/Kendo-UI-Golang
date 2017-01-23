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
  Aggregates    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.data
  //
  // Type: Function
  //
  // The field from the server response which contains the data items. Can be set to a function which is called to return the data items for the response.
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
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.errors
  //
  // Type: Function
  // Default: errors
  //
  // The field from the server response which contains server-side errors. Can be set to a function which is called to return the errors for response. If there are any errors, the 'error' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#events-error  event will be fired.
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
  // todo http://docs.telerik.com/kendo-ui/api/javascript/data/model
  //Model    Model

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.parse
  //
  // Type: Function
  //
  // Executed before the server response is used. Use it to preprocess or parse the server response.
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
  //   response 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object 'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
  //   The initially parsed server response that may need additional modifications.
  //
  // Return:
  //   'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array  The data items from the response.
  //
  Parse    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.total
  //
  // Type: Function
  //
  // The field from the server response which contains the total number of data items. Can be set to a function which is called to return the total number of data items for the response.
  //
  //    The 'schema.total' setting may be omitted when the Grid is bound to a plain 'Array' (that is, the data items' collection is not a value of a field in the server response). In this case, the 'length' of the response 'Array' will be used.
  //
  // The 'schema.total' must be set if the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true' or the 'schema.data' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.data option is used.
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
  // 'Number' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number  The total number of data items.
  Total    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type
  //
  // Type: String
  // Default: json
  //
  // The type of the response.
  //
  // The supported values are:
  //
  // *  "xml"
  // *  "json"
  //
  // By default, the schema interprets the server response as JSON.
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
  Type    TypeDataEnum

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
