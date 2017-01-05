package kendo

// The animation played when the suggestion popup is open/closed.
type AnimationAction struct {
  // default: 100
  // The duration of the open/close animation in milliseconds.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close.duration
  Duration  int64

  // The effect(s) to use when playing the open/close animation. Multiple effects should be separated with a space.
  //
  // Complete list of available animations ( http://docs.telerik.com/kendo-ui/api/javascript/effects/common )
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close.effects
  Effects   string
}

// Configures the opening and closing animations of the suggestion popup. Setting the 'animation' option to 'false' will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
//
// 'animation:true' is not a valid configuration.
//
// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation
/*
      Example - disable open and close animations
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       animation: false
      });
      </script>
*/
type Animation struct {
  // Configures the opening and closing animations of the suggestion popup. Setting the 'animation' option to 'false' will disable the opening and closing animations. As a result the suggestion popup will open and close instantly.
  //
  // 'animation:true' is not a valid configuration.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation
  /*
        Example - disable open and close animations
        <input id="autocomplete" />
        <script>
        $("#autocomplete").kendoAutoComplete({
         animation: false
        });
        </script>
  */
  AnimationEnable       bool

  // The animation played when the suggestion popup is closed.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close
  /*
        Example - configure the close animation
        <input id="autocomplete" />
        <script>
        $("#autocomplete").kendoAutoComplete({
         animation: {
          close: {
           effects: "zoom:out",
           duration: 300
          }
         }
        });
        </script>
  */
  Close       AnimationAction

  // The animation played when the suggestion popup is opened.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.open
  /*
        Example - configure the open animation
        <input id="autocomplete" />
        <script>
        $("#autocomplete").kendoAutoComplete({
         animation: {
          open: {
           effects: "zoom:in",
           duration: 300
          }
         }
        });
        </script>
  */
  Open        AnimationAction

  // If set to 'true', the widget automatically adjusts the width of the popup element and does not wrap up the item label.
  //
  // Note: Virtualized list doesn't support the auto-width functionality.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-autoWidth
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
  AutoWidth   bool

  // The data source of the widget which is used to display suggestions for the current value. Can be a JavaScript object which represents a valid data source configuration, a JavaScript array or an existing kendo.data.DataSource ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource ) instance.
  //
  // If the 'dataSource' option is set to a JavaScript object or array the widget will initialize a new kendo.data.DataSource ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource ) instance using that value as data source configuration.
  //
  // If the 'dataSource' option is an existing kendo.data.DataSource ( http://docs.telerik.com/kendo-ui/api/javascript/data/datasource ) instance the widget will use that instance and will *not* initialize a new one.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-dataSource
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
  */
  DataSource  kendo.data.DataSource

  // default: true
  // Unless this options is set to 'false', a button will appear when hovering the widget. Clicking that button will reset the widget's value and will trigger the change event.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-clearButton
  /*
      Example - disable the clear button
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        clearButton: false
      });
      </script>
  */
  ClearButton  bool

  // default: null
  // The field of the data item used when searching for suggestions. This is the text that will be displayed in the list of matched results.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-dataTextField
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
  DataTextField  string

  // default: 200
  // The delay in milliseconds between a keystroke and when the widget displays the suggestion popup.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-delay
  /*
      Example - set the delay
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       delay: 500
      });
      </script>
  */
  Delay  Number

  // default: true
  // If set to 'false' the widget will be disabled and will not allow user input. The widget is enabled by default and allows user input.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-enable
  /*
      Example - disable the widget
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       enable: false
      });
      </script>
  */
  Enable  bool

  // default: false
  // If set to 'true' the widget will not show all items when the text of the search input cleared. By default the widget shows all items when the text of the search input is cleared. Works in conjunction with minLength ( http://docs.telerik.comhttp://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-minLength ).
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-enforceMinLength
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
  EnforceMinLength  bool

  // default: "startswith"
  // The filtering method used to determine the suggestions for the current value. The default filter is "startswith" - all data items which begin with the current widget value are displayed in the suggestion popup. The supported 'filter' values are 'startswith', 'endswith' and 'contains'.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-filter
  /*
      Example - set the filter
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       filter: "contains"
      });
      </script>
  */
  Filter  string

  // The template ( http://docs.telerik.com/kendo-ui/api/javascript/kendo#methods-template ) used to render the fixed header group. By default the widget displays only the value of the current group.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-fixedGroupTemplate
  /*

      <input id="customers" style="width: 400px" />
      <script>
        $(document).ready(function() {
          $("#customers").kendoAutoComplete({
            dataTextField: "ContactName",
            fixedGroupTemplate: "Fixed group: #:data#",
            height: 400,
            dataSource: {
              type: "odata",
              transport: {
                read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Customers"
              },
              group: { field: "Country" }
            }
          });
        });
      </script>
  */
  FixedGroupTemplate  String - Function

  // The template ( http://docs.telerik.com/kendo-ui/api/javascript/kendo#methods-template ) used to render the footer template. The footer template receives the widget itself as a part of the data argument. Use the widget fields directly in the template.
  //
  // The widget instance.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-footerTemplate
  /*
      Parameters

  instance Object


  The widget instance.

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
  FooterTemplate  String - Function

  // The template ( http://docs.telerik.com/kendo-ui/api/javascript/kendo#methods-template ) used to render the groups. By default the widget displays only the value of the group.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-groupTemplate
  /*

      <input id="customers" style="width: 400px" />
      <script>
        $(document).ready(function() {
          $("#customers").kendoAutoComplete({
            dataTextField: "ContactName",
            groupTemplate: "Group: #:data#",
            height: 400,
            dataSource: {
              type: "odata",
              transport: {
                read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Customers"
              },
              group: { field: "Country" }
            }
          });
        });
      </script>
  */
  GroupTemplate  String - Function

  // default: 200
  // The height of the suggestion popup in pixels. The default value is 200 pixels.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-height
  /*
      Example - set the height
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       height: 500
      });
      </script>
  */
  Height  int64

  // default: true
  // If set to 'true' the first suggestion will be automatically highlighted.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-highlightFirst
  /*
      Example - set highlightFirst
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       highlightFirst: false
      });
      </script>
  */
  HighlightFirst  bool

  // default: true
  // If set to 'false' case-sensitive search will be performed to find suggestions. The widget performs case-insensitive searching by default.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-ignoreCase
  /*
      Example - disable case-insensitive suggestions
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       ignoreCase: false
      });
      </script>
  */
  IgnoreCase  bool

  // default: 1
  // The minimum number of characters the user must type before a search is performed. Set to higher value than '1' if the search could match a lot of items.
  //
  // Widget will initiate a request when input value is cleared. If you would like to prevent this behavior please check the filtering ( http://docs.telerik.com#events-filtering ) event for more details.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-minLength
  /*
      Example - set minLength
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       minLength: 3
      });
      </script>
  */
  MinLength  int64

  // default: "NO DATA FOUND."
  // The template ( http://docs.telerik.com/kendo-ui/api/javascript/kendo#methods-template ) used to render the "no data" template, which will be displayed if no results are found or the underlying data source is empty. The noData template receives the widget itself as a part of the data argument. The template will be evaluated on every widget data bound.
  //
  // *Important* The popup will open when 'noDataTemplate' is defined
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-noDataTemplate
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
  NoDataTemplate  String - Function

  // default: ""
  // The hint displayed by the widget when it is empty. Not set by default.
  //
  // The Kendo UI AutoComplete widget could also use the value of the 'placeholder' HTML attribute as hint.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-placeholder
  /*
      Example - specify placeholder
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       placeholder: "Enter value ..."
      });
      </script>
  */
  Placeholder  string

  // The options that will be used for the popup initialization. For more details about the available options refer to Popup ( http://docs.telerik.com/kendo-ui/api/javascript/ui/popup ) documentation.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-popup
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
  Popup  Object

  // default: ""
  // The character used to separate multiple values. Empty by default.
  //
  // As of Q3 2016 the Autocomplete widget supports multiple separators listed in an array. All separators will be replaced with the first array item, which acts as a default separator.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-separator
  /*
      Example - set separator to allow multiple values
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       separator: ", "
      });
      </script>
  */
  Separator  String - Array

  // default: false
  // If set to 'true' the widget will automatically use the first suggestion as its value.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-suggest
  /*
      Example - enable automatic suggestion
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
       suggest: true
      });
      </script>
  */
  Suggest  bool

  // Specifies a static HTML content, which will be rendered as a header of the popup element.
  //
  // *Important* The header content *should be wrapped* with a HTML tag if it contains more than one element. This is applicable also when header content is just a string/text.
  //
  // *Important* Widget does not pass a model data to the header template. Use this option only with static HTML.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-headerTemplate
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
  FeaderTemplate  String - Function

  // The template ( http://docs.telerik.com/kendo-ui/api/javascript/kendo#methods-template ) used to render the suggestions. By default the widget displays only the text of the suggestion (configured via 'dataTextField').
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-template
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
  */
  Template  String - Function

  // default: ""
  // The value of the widget.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-value
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
  Value  string

  // default: false
  // Specifies the value binding ( http://docs.telerik.com/kendo-ui/framework/mvvm/bindings/value ) behavior for the widget when the initial model value is null. If set to true, the View-Model field will be updated with the selected item text field. If set to false, the View-Model field will be updated with the selected item.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-valuePrimitive
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
  ValuePrimitive  bool
}














