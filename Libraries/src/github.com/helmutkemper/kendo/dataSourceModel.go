package kendo

import "bytes"

// http://docs.telerik.com/kendo-ui/api/javascript/data/model
//
// kendo.data.Model
//
// The Model inherits from the ObservableObject and extends it with the ability to define schema - fields and methods. The DataSource contains instances of the Model when the schema.model setting is specified.
//
// Configuration
//
// To define a new model use the Model.define method.
//
/*
    Example - Define a model
    <script>
    var Person = kendo.data.Model.define({
        id: "personId", // the identifier of the model
        fields: {
            "name": {
                type: "string"
            },
            "age": {
                type: "number"
            }
        }
    });

    var person = new Person({
        name: "John Doe",
        age: 42
    });

    console.log(person.get("name")); // outputs "John Doe"
    console.log(person.get("age")); // outputs 42
    </script>
 */
type DataSourceModel struct {

  // http://docs.telerik.com/kendo-ui/api/javascript/data/model#fields-id
  //
  // Type: String
  //
  // The value of the Model's ID. This field is available only if the id is defined in the Model configuration. See the example below.
  Id    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/model#fields-idField
  //
  // Type: String
  //
  // The name of the Model's ID field. This field is available <b>only</b> if the <b>id</b> is defined in the Model configuration.
  /*
      <script>
      var Person = kendo.data.Model.define({
          id: "personId",
          fields: {
              "name": {
                  type: "string"
              },
              "age": {
                  type: "number"
              }
          }
      });

      var person = new Person({
          personId: 1,
          name: "John Doe",
          age: 42
      });

      console.log(person.id); // outputs 1
      console.log(person.idField); // outputs "personId"
      </script>
  */
  IdField   string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/model#fields-uid
  //
  // Type: String
  //
  // The unique identifier of the <b>ObservableObject</b>.
  // The main benefit of uid's is to represent a link between data items (that may not have an ID of their own) and the corresponding rendered DOM elements (list items, table rows, etc). The uid's are generated randomly and they are not persisted across data or web page reloads.
  /*
      EXAMPLE - USING THE UID FIELD
      <script>
        var observable = new kendo.data.ObservableObject({ name: "John Doe" });
        console.log(observable.uid); // outputs "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" where "x" is a number or letter
      </script>
  */
  UId   string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/model#fields-dirty
  //
  // Type: Bool
  //
  // Indicates whether the model is modified.
  /*
      EXAMPLE - USING THE DIRTY FIELD
      <script>
        var model = new kendo.data.Model({
            name: "John Doe"
        });

        console.log(model.dirty); // outputs "false"
        model.set("name", "Jane Doe");
        console.log(model.dirty); // outputs "true"
      </script>
  */
  Dirty   bool

  /*
      <script>
      var Person = kendo.data.Model.define({
          id: "personId",
          fields: {
              "name": {
                  type: "string"
              },
              "age": {
                  type: "number"
              }
          }
      });

      var person = new Person({
          personId: 1,
          name: "John Doe",
          age: 42
      });

      console.log(person.id); // outputs 1
      console.log(person.idField); // outputs "personId"
      </script>
  */
  Fields    []FieldLine
  GoTemplate    *GoTemplate
}

func ( el DataSourceModel ) getTemplate () string {
  return `{{if ne (string .Id) "''"}}id: {{string .Id}}, {{end}}{{if ne (string .IdField) "''"}}idField: {{string .IdField}}, {{end}}{{if ne (string .UId) "''"}}uid: {{string .UId}}, {{end}}{{if .Dirty}}dirty: true,{{end}}{{if .Fields}}fields: {{range $v := .Fields}}{ {{string $v}} },{{end}}{{end}}`
}

func ( el DataSourceModel ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el DataSourceModel ) String() string {
  out := el.Buffer()
  return out.String()
}
