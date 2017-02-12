package kendo

import "bytes"

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
  Create    Create

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
  Destroy    Destroy

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
  //
  //   data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  //   The parameters which will be sent to the remote service. The value specified in the data field of the transport settings (create, read, update or destroy) is included as well. If 'Array' #batch-boolean-default">batch is set to false, the fields of the changed data items are also included.
  //   data.aggregate aggregate option. Available if the serverAggregates option is set to true and the data source makes a "read" request.
  //   data.group group option. Available if the serverGrouping option is set to true and the data source makes a "read" request.
  //   data.filter filter option. Available if the serverFiltering option is set to true and the data source makes a "read" request.
  //   data.models batch option is set to true.
  //   data.page serverPaging option is set to true and the data source makes a "read" request.
  //   data.pageSize pageSize option. Available if the serverPaging option is set to true and the data source makes a "read" request.
  //   data.skip serverPaging option is set to true and the data source makes a "read" request.
  //   data.sort sort option. Available if the serverSorting option is set to true and the data source makes a "read" request.
  //   data.take serverPaging option is set to true and the data source makes a "read" request.
  //   type
  //
  // Return:
  //
  //   'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object  the request parameters converted to a format required by the remote service.
  //
  ParameterMap    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.push
  //
  // Type: Function
  //
  // The function invoked during transport initialization which sets up push notifications. The data source will call this function only once and provide callbacks which will handle push notifications (data pushed from the server).
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
  //
  //   callbacks 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  //   An object containing callbacks for notifying the data source of push notifications.
  //   callbacks.pushCreate 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  //   Function that should be invoked to notify the data source about newly created data items that are pushed from the server. Accepts a single argument - the object pushed from the server which should follow the schema.data configuration.
  //   callbacks.pushDestroy 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  //   Function that should be invoked to notify the data source about destroyed data items that are pushed from the server. Accepts a single argument - the object pushed from the server
  //   which should follow the schema.data configuration.
  //   callbacks.pushUpdate 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  //   Function that should be invoked to notify the data source about updated data items that are pushed from the server. Accepts a single argument - the object pushed from the server
  //   which should follow the schema.data configuration.
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
  Read    Read

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr
  //
  // Type: Object
  //
  // The configuration used when 'type' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-type  is set to "signalr". Configures the SignalR settings - hub, connection promise, server, and client hub methods.
  //
  // Live demo available at demos.telerik.com/kendo-ui http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://demos.telerik.com/kendo-ui/grid/signalr .
  //
  // It is recommended to get familiar with the SignalR JavaScript API http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://www.asp.net/signalr/overview/guide-to-the-api/hubs-api-guide-javascript-client .
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
  // Parameters:
  //
  //   e.data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  //   An object containing the created (e.data.created), updated (e.data.updated), and destroyed (e.data.destroyed) items.
  //   e.success 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  //   A callback that should be called for each operation with two parameters - items and operation. See example below.
  //   e.fail 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
  //   A callback that should be called in case of failure of any of the operations.
  /*
        <h5>Example - declare transport submit function</h5>
        <script>
            var dataSource = new kendo.data.DataSource({
                batch: true,
                transport: {
                  submit: function(e) {
                    // Check out the network tab on "Save Changes"
                    $.ajaxBatch({
                      url: "<your service URL>",
                      data: e.data
                    })
                    .done(function() {
                      e.success(e.data.created, "create");
                      e.success([], "update");
                      e.success([], "destroy");
                    })
                    .fail(function() {
                      e.error();
                    });
                  },
                  read: function(e) {
                    $.ajax({
                      url: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Customers",
                      dataType: "jsonp",
                      data: data,
                      jsonp: "$callback"
                    })
                    .done(e.success)
                    .fail(e.error);
                  }
                }
              });
        </script>
   */
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
  Update    Update

  GoTemplate    *GoTemplate
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

func ( el Transport ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Transport ) String() string {
  out := el.Buffer()
  return out.String()
}
