package kendo

type TransportEvent struct{
  // If set to false, the request result will not be cached by the browser. Setting cache to false will only work
  // correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default,
  // "jsonp" requests are not cached.
  // Refer to the jQuery.ajax documentation for further info.
  Cache           bool

  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use
  // "application/json" if the content is JSON. Refer to the jQuery.ajax documentation for further information.
  ContentType     string

  // Additional parameters that are sent to the remote service. The parameter names must not match reserved words, which
  // are used by the Kendo UI DataSource for sorting, filtering, paging, and grouping.
  // Refer to the jQuery.ajax documentation for further information.
  Data            ComplexJavaScriptType

  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  // Refer to the jQuery.ajax documentation for further information.
  DataType        DataJSonOrJSonpEnum

  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  // The type option is ignored if dataType is set to "jsonp". JSONP always uses GET requests.
  // Refer to the jQuery.ajax documentation for further information.
  Type            MethodEnum

  // The URL to which the request is sent.
  // If set to function, the data source will invoke it and use the result as the URL.
  Url             ComplexJavaScriptType
}
