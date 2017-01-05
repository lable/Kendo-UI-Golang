package kendo

// The configuration used to load and save the data items. A data source is remote or local based on the way of it
// retrieves data items.
// Remote data sources load and save data items from and to a remote end-point (also known as remote service or server).
// The transport option describes the remote service configuration - URL, HTTP verb, HTTP headers, and others. The
// transport option can also be used to implement custom data loading and saving.
// Local data sources are bound to a JavaScript array via the data option.
type Transport struct{
  // The configuration used when the data source saves newly created data items. Those are items added to the data
  // source via the add or insert methods.
  // The data source uses jQuery.ajax to make a HTTP request to the remote service. The value configured via
  // transport.create is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via
  // transport.create except the success and error callback functions which are used by the transport.
  // If the value of transport.create is a function, the data source invokes that function instead of jQuery.ajax. Check
  // the jQuery documentation for more details on the provided argument.
  // If the value of transport.create is a string, the data source uses this string as the URL of the remote service.
  // The remote service must return the inserted data items and the data item field configured as the id must be set.
  // For example, if the id of the data item is ProductID, the "create" server response must be [{ "ProductID": 79 }].
  // All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as
  // objects. Mixing the different configuration alternatives is not possible.
  Create            TransportEvent

  // The configuration used when the data source destroys data items. Those are items removed from the data source via
  // the remove method.
  // The data source uses jQuery.ajax to make an HTTP request to the remote service. The value configured via
  // transport.destroy is passed to jQuery.ajax. This means that you can set all options supported by jQuery.ajax via
  // transport.destroy except the success and error callback functions which are used by the transport.
  // If the value of transport.destroy is a function, the data source invokes that function instead of jQuery.ajax.
  // If the value of transport.destroy is a string, the data source uses this string as the URL of the remote service.
  // All transport actions (read, update, create, destroy) must be defined in the same way, that is, as functions or as
  // objects. Mixing the different configuration alternatives is not possible.
  Destroy           TransportEvent

  //parei aqui http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap
}
