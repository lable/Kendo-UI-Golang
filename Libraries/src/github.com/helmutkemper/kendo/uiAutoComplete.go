package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation
//
// kendo.ui.AutoComplete http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete
// Represents the Kendo UI AutoComplete widget. Inherits from Widget http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete.
// Important: The Kendo UI AutoComplete should be created from an input HTML element.
// Configuration http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete
//
type UIAutoComplete struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation
  //
  // Type: Boolean
  //
  // Configures the opening and closing animations of the suggestion popup. Setting the 'animation' option to 'false' will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
  //
  // 'animation:true' is not a valid configuration.
  //
  //
  //
  /*
      Example - disable open and close animations
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        animation: false
      });
      </script>


      Example - configure the animation
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        animation: {
         close: {
           effects: "fadeOut zoom:out",
           duration: 300
         },
         open: {
           effects: "fadeIn zoom:in",
           duration: 300
         }
        }
      });
      </script>
  */
  //
  Animation Animation

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-autoWidth
  //
  // Type: Boolean
  //
  // If set to 'true', the widget automatically adjusts the width of the popup element and does not wrap up the item label.
  //
  //    Note: Virtualized list doesn't support the auto-width functionality.
  //
  /*
      Example - enable autoWidth
      <input id="autocomplete" style="width: 100px;" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        autoWidth: true,
        dataSource: {
          data: ["Short item", "An item with really, really long text"]
        }
      });
      </script>
  */
  //
  AutoWidth    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-dataSource
  //
  // Type: Object
  //
  // The data source of the widget which is used to display suggestions for the current value. Can be a JavaScript object which represents a valid data source configuration, a JavaScript array or an existing kendo.data.DataSource http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/data/datasource
  // instance.
  //
  // If the 'dataSource' option is set to a JavaScript object or array the widget will initialize a new kendo.data.DataSource http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/data/datasource  instance using that value as data source configuration.
  //
  // If the 'dataSource' option is an existing kendo.data.DataSource http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/data/datasource  instance the widget will use that instance and will not initialize a new one.
  //
  //
  //
  /*
      Example - set dataSource as a JavaScript object
      <input id="autocomplete" />
      <script>
      $("#autoComplete").kendoAutoComplete({
        dataSource: {
          data: ["One", "Two"]
        }
      });
      </script>


      Example - set dataSource as a JavaScript array
      <input id="autocomplete" />
      <script>
      var data = ["One", "Two"];
      $("#autoComplete").kendoAutoComplete({
        dataSource: data
      });
      </script>


      Example - set dataSource as an existing kendo.data.DataSource instance
      <input id="autocomplete" />
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          }
        }
      });
      $("#autocomplete").kendoAutoComplete({
        dataSource: dataSource,
        dataTextField: "ProductName"
      });
      </script>
  */
  //
  DataSource    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-clearButton
  //
  // Type: Boolean
  // Default: true
  //
  // Unless this options is set to 'false', a button will appear when hovering the widget. Clicking that button will reset the widget's value and will trigger the change event.
  //
  //
  //
  /*
      Example - disable the clear button
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
          clearButton: false
      });
      </script>
  */
  //
  ClearButton    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-dataTextField
  //
  // Type: String
  // Default: null
  //
  // The field of the data item used when searching for suggestions.  This is the text that will be displayed in the list of matched results.
  //
  //
  //
  /*
      Example - set the dataTextField
      <input id="autocomplete" />
      <script>
      var data = [
        { id: 1, name: "Apples" },
        { id: 2, name: "Oranges" }
      ];
      $("#autocomplete").kendoAutoComplete({
        dataTextField: "name", // The widget is bound to the "name" field
        dataSource: data
      });
      </script>
  */
  //
  DataTextField    string

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-delay
  //
  // Type: Number
  // Default: 200
  //
  // The delay in milliseconds between a keystroke and when the widget displays the suggestion popup.
  //
  //
  //
  /*
      Example - set the delay
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        delay: 500
      });
      </script>
  */
  //
  Delay    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-enable
  //
  // Type: Boolean
  // Default: true
  //
  // If set to 'false' the widget will be disabled and will not allow user input. The widget is enabled by default and allows user input.
  //
  //
  //
  /*
      Example - disable the widget
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        enable: false
      });
      </script>
  */
  //
  Enable    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-enforceMinLength
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true' the widget will not show all items when the text of the search input cleared. By default the widget shows all items when the text of the search input is cleared. Works in conjunction with minLength http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-minLength .
  //
  //
  //
  /*
      Example - enforce minLength
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
          dataTextField: "ProductName",
          filter: "contains",
          minLength: 3,
          enforceMinLength: false,
          autoBind: false,
          dataSource: {
              type: "odata",
              serverFiltering: true,
              transport: {
                  read: "//demos.telerik.com/kendo-ui/service/Northwind.svc/Products"
              }
          }
      });
      </script>
  */
  //
  EnforceMinLength    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-filter
  //
  // Type: String
  // Default: startswith
  //
  // The filtering method used to determine the suggestions for the current value. The default filter is "startswith" -
  // all data items which begin with the current widget value are displayed in the suggestion popup. The supported 'filter' values are 'startswith', 'endswith' and 'contains'.
  //
  //
  //
  /*
      Example - set the filter
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        filter: "contains"
      });
      </script>
  */
  //
  Filter    string

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-fixedGroupTemplate
  //
  // Type: String
  //
  // The template http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/kendo#methods-template  used to render the fixed header group. By default the widget displays only the value of the current group.
  //
  //
  //

  //
  FixedGroupTemplate    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-footerTemplate
  //
  // Type: String
  //
  // The template http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/kendo#methods-template  used to render the footer template. The footer template receives the widget itself as a part of the data argument. Use the widget fields directly in the template.
  //
  //
  //
  /*
      Example - specify footerTemplate as a string
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        footerTemplate: 'Total <strong>#: instance.dataSource.total() #</strong> items found'
      });
      </script>
  */
  //
  // instance 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
  // The widget instance.
  //
  FooterTemplate    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-groupTemplate
  //
  // Type: String
  //
  // The template http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/kendo#methods-template  used to render the groups. By default the widget displays only the value of the group.
  //
  //
  //

  //
  GroupTemplate    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-height
  //
  // Type: Number
  // Default: 200
  //
  // The height of the suggestion popup in pixels. The default value is 200 pixels.
  //
  //
  //
  /*
      Example - set the height
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        height: 500
      });
      </script>
  */
  //
  Height    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-highlightFirst
  //
  // Type: Boolean
  // Default: true
  //
  // If set to 'true' the first suggestion will be automatically highlighted.
  //
  //
  //
  /*
      Example - set highlightFirst
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        highlightFirst: false
      });
      </script>
  */
  //
  HighlightFirst    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-ignoreCase
  //
  // Type: Boolean
  // Default: true
  //
  // If set to 'false' case-sensitive search will be performed to find suggestions. The widget performs case-insensitive searching by default.
  //
  //
  //
  /*
      Example - disable case-insensitive suggestions
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        ignoreCase: false
      });
      </script>
  */
  //
  IgnoreCase    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-minLength
  //
  // Type: Number
  // Default: 1
  //
  // The minimum number of characters the user must type before a search is performed. Set to higher value than '1' if the search could match a lot of items.
  //
  //    Widget will initiate a request when input value is cleared. If you would like to prevent this behavior please check the filtering event for more details.
  //
  /*
      Example - set minLength
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        minLength: 3
      });
      </script>
  */
  //
  MinLength    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-noDataTemplate
  //
  // Type: String
  // Default: NO DATA FOUND.
  //
  // The template http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/kendo#methods-template  used to render the "no data" template, which will be displayed if no results are found or the underlying data source is empty.
  // The noData template receives the widget itself as a part of the data argument. The template will be evaluated on every widget data bound.
  //
  //    Important The popup will open when 'noDataTemplate' is defined
  //
  /*
      Example - specify headerTemplate as a string
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        noDataTemplate: 'No Data!'
      });
      </script>
  */
  //
  NoDataTemplate    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-placeholder
  //
  // Type: String
  //
  // The hint displayed by the widget when it is empty. Not set by default.
  //
  //
  //
  /*
      Example - specify placeholder
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        placeholder: "Enter value ..."
      });
      </script>


      Example - use the placeholder HTML attribute
      <input id="autocomplete" placeholder="Enter value..." />
      <script>
      $("#autocomplete").kendoAutoComplete();
      </script>
  */
  //
  Placeholder    string

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-popup
  //
  // Type: Object
  //
  // The options that will be used for the popup initialization. For more details about the available options
  // refer to Popup http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/ui/popup  documentation.
  //
  //
  //
  /*
      Example - append the popup to a specific element
      <div id="container">
          <input id="autocomplete" />
      </div>
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        popup: {
          appendTo: $("#container")
        }
      });
      </script>
  */
  //
  Popup    Popup

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-separator
  //
  // Type: String
  //
  // The character used to separate multiple values. Empty by default.
  //
  //    As of Q3 2016 the Autocomplete widget supports multiple separators listed in an array. All separators will be replaced with the first array item, which acts as a default separator.
  //
  /*
      Example - set separator to allow multiple values
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        separator: ", "
      });
      </script>


      Example - set multiple separators
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        separator: [", ", "; "]
      });
      </script>
  */
  //
  Separator    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-suggest
  //
  // Type: Boolean
  // Default: false
  //
  // If set to 'true' the widget will automatically use the first suggestion as its value.
  //
  //
  //
  /*
      Example - enable automatic suggestion
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        suggest: true
      });
      </script>
  */
  //
  Suggest    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-headerTemplate
  //
  // Type: String
  //
  // Specifies a static HTML content, which will be rendered as a header of the popup element.
  //
  //    Important The header content should be wrapped with a HTML tag if it contains more than one element. This is applicable also when header content is just a string/text.
  //    Important Widget does not pass a model data to the header template. Use this option only with static HTML.
  //
  /*
      Example - specify headerTemplate as a string
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        headerTemplate: '<div><h2>Fruits</h2></div>'
      });
      </script>
  */
  //
  HeaderTemplate    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-template
  //
  // Type: String
  //
  // The template http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/api/javascript/kendo#methods-template  used to render the suggestions. By default the widget displays only the text of the suggestion (configured via 'dataTextField').
  //
  //
  //
  /*
      Example - specify template as a function
      <input id="autocomplete" />
      <script id="template" type="text/x-kendo-template">
        <span>
          <img src="/img/#: id #.png" alt="#: name #" />
          #: name #
        </span>
      </script>
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        template: kendo.template($("#template").html())
      });
      </script>


      Example - specify template as a string
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        template: '<span><img src="/img/#: id #.png" alt="#: name #" />#: name #</span>'
      });
      </script>
  */
  //
  Template    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-value
  //
  // Type: String
  //
  // The value of the widget.
  //
  //
  //
  /*
      Example - specify value of the widget
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { id: 1, name: "Apples" },
          { id: 2, name: "Oranges" }
        ],
        dataTextField: "name",
        value: "Oranges"
      });
      </script>
  */
  //
  Value    string

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-valuePrimitive
  //
  // Type: Boolean
  // Default: false
  //
  // Specifies the value binding http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/framework/mvvm/bindings/value  behavior for the widget when the initial model value is null. If set to true, the View-Model field will be updated with the selected item text field. If set to false, the View-Model field will be updated with the selected item.
  //
  //
  //
  /*
      Example - specify that the View-Model field should be updated with the selected item text
      <input id="autocomplete" data-bind="value: productName, source: products" />

      <script>
      $("#autocomplete").kendoAutoComplete({
        valuePrimitive: true,
        dataTextField: "name"
      });
      var viewModel = kendo.observable({
        productName: null,
        products: [
          { id: 1, name: "Coffee" },
          { id: 2, name: "Tea" },
          { id: 3, name: "Juice" }
        ]
      });

      kendo.bind($("#autocomplete"), viewModel);
      </script>
  */
  //
  ValuePrimitive    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual
  //
  // Type: Boolean
  // Default: false
  //
  // Enables the virtualization feature of the widget. The configuration can be set on an object, which contains two properties - 'itemHeight' and 'valueMapper'.
  //
  // For detailed information, refer to the article on virtualization http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/controls/editors/combobox/virtualization .
  //
  //
  //
  /*
      Example - AutoComplete with a virtualized list
      <input id="orders" style="width: 400px" />
      <script>
          $(document).ready(function() {
              $("#orders").kendoAutoComplete({
                  template: "#= OrderID # | For: #= ShipName #, #= ShipCountry #",
                  dataTextField: "ShipName",
                  virtual: true,
                  height: 520,
                  dataSource: {
                      type: "odata",
                      transport: {
                          read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
                      },
                      schema: {
                          model: {
                              fields: {
                                  OrderID: { type: "number" },
                                  Freight: { type: "number" },
                                  ShipName: { type: "string" },
                                  OrderDate: { type: "date" },
                                  ShipCity: { type: "string" }
                              }
                          }
                      },
                      pageSize: 80,
                      serverPaging: true,
                      serverFiltering: true
                  }
              });
          });
      </script>


      Example - AutoComplete widget with a declarative virtualization config
      <div class="demo-section k-header">
          <h4>Search for shipping name</h4>
          <input id="orders" style="width: 400px"
                 data-role="autocomplete"
                 data-bind="value: order, source: source"
                 data-text-field="ShipName"
                 data-virtual="{itemHeight:26,valueMapper:orderValueMapper}"
                 />
      </div>

      <script>
          $(document).ready(function() {
              var model = kendo.observable({
                order: "Hanari Carnes",
                source: new kendo.data.DataSource({
                  type: "odata",
                  transport: {
                    read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
                  },
                  schema: {
                    model: {
                      fields: {
                        OrderID: { type: "number" },
                        Freight: { type: "number" },
                        ShipName: { type: "string" },
                        OrderDate: { type: "date" },
                        ShipCity: { type: "string" }
                      }
                    }
                  },
                  pageSize: 80,
                  serverPaging: true,
                  serverFiltering: true
                })
              });

              kendo.bind($(document.body), model);
          });

          function orderValueMapper(options) {
              $.ajax({
                url: "http://demos.telerik.com/kendo-ui/service/Orders/ValueMapper",
                type: "GET",
                dataType: "jsonp",
                data: convertValues(options.value),
                success: function (data) {
                  options.success(data);
                }
              })
          }

          function convertValues(value) {
              var data = {};

              value = $.isArray(value) ? value : [value];

              for (var idx = 0; idx < value.length; idx++) {
                  data["values[" + idx + "]"] = value[idx];
              }

              return data;
          }
      </script>
  */
  //
  Virtual    bool

  Template    *Template
}

func ( el UIAutoComplete ) getTemplate () string {
  return `{{if .AutoWidth}}autoWidth: true,{{end}}
{{if .ClearButton}}clearButton: true,{{end}}
{{if ne (string .DataTextField) "null"}}dataTextField: {{string .DataTextField}},{{end}}
{{if .Delay }}delay: {{.Delay}},{{end}}
{{if .Enable}}enable: true,{{end}}
{{if .EnforceMinLength}}enforceMinLength: true,{{end}}
{{if ne (string .Filter) "null"}}filter: {{string .Filter}},{{end}}
{{if ne (string .FixedGroupTemplate) "null"}}fixedGroupTemplate: {{string .FixedGroupTemplate}},{{end}}
{{if ne (string .FooterTemplate) "null"}}footerTemplate: {{string .FooterTemplate}},{{end}}
{{if ne (string .GroupTemplate) "null"}}groupTemplate: {{string .GroupTemplate}},{{end}}
{{if .Height }}height: {{.Height}},{{end}}
{{if .HighlightFirst}}highlightFirst: true,{{end}}
{{if .IgnoreCase}}ignoreCase: true,{{end}}
{{if .MinLength }}minLength: {{.MinLength}},{{end}}
{{if ne (string .NoDataTemplate) "null"}}noDataTemplate: {{string .NoDataTemplate}},{{end}}
{{if ne (string .Placeholder) "null"}}placeholder: {{string .Placeholder}},{{end}}
{{if ne (string .Popup) "null"}}popup: {{string .Popup}},{{end}}
{{if ne (string .Separator) "null"}}separator: {{string .Separator}},{{end}}
{{if .Suggest}}suggest: true,{{end}}
{{if ne (string .HeaderTemplate) "null"}}headerTemplate: {{string .HeaderTemplate}},{{end}}
{{if ne (string .Template) "null"}}template: {{string .Template}},{{end}}
{{if ne (string .Value) "null"}}value: {{string .Value}},{{end}}
{{if .ValuePrimitive}}valuePrimitive: true,{{end}}
`
}

