package kendo

type Schema struct {
  // The field from the response which contains the aggregate results. Can be set to a function which is called to
  // return the aggregate results from the response.
  // The aggregates option is used only when the serverAggregates option is set to true.
  Aggregates        []AggregateList

  // The field from the server response which contains the data items. Can be set to a function which is called to
  // return the data items for the response.
  Data              ComplexJavaScriptType

  // The field from the server response which contains server-side errors. Can be set to a function which is called to
  // return the errors for response. If there are any errors, the error event will be fired.
  Errors            ComplexJavaScriptType

  // The field from the server response which contains the groups. Can be set to a function which is called to return
  // the groups from the response.
  // The groups option is used only when the serverGrouping option is set to true.
  Groups            Group

  //>>>>>>Model

  // Executed before the server response is used. Use it to preprocess or parse the server response.
  Parse             ComplexJavaScriptType

  // The field from the server response which contains the total number of data items. Can be set to a function which
  // is called to return the total number of data items for the response.
  Total             ComplexJavaScriptType

  // The type of the response.
  // The supported values are: "xml" OR "json"
  // By default, the schema interprets the server response as JSON.
  Type              DataXmlOrJSonEnum
}
