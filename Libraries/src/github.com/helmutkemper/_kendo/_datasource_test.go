package kendo

import "fmt"

func ExampleDataSource_String() {
  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  // aggregate: [{ field: &#34;fild_name&#34;, aggregate: &#34;max&#34; },],
}

func ExampleDataSource_String2() {
  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}

func ExampleDataSource_String3() {
  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}

func ExampleDataSource_String4() {
  type data struct {
    Field string
    Value int
  }
  var dataJs []data = []data{
    {
      Field: "field_1",
      Value: 1,
    },
    {
      Field: "field_2",
      Value: 2,
    },
  }

  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Data: ComplexJavaScriptType{
      AsJSon: dataJs,
    },
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}

func ExampleDataSource_String5() {
  type data struct {
    Field string
    Value int
  }
  var dataJs []data = []data{
    {
      Field: "field_1",
      Value: 1,
    },
    {
      Field: "field_2",
      Value: 2,
    },
  }

  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Data: ComplexJavaScriptType{
      AsJSon: dataJs,
    },
    Filter: Filter{
      Field: "field_name",
      Filters: []FilterLine{
        {
          Operator: OPERATOR_CONTAINS,
          Template: &t,
        },
        {
          Field: "field_name",
          Operator: OPERATOR_CONTAINS,
          Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
          Template: &t,
        },
      },
      Logic: LOGIC_AND,
      Template: &t,
    },
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}

func ExampleDataSource_String6() {
  type data struct {
    Field string
    Value int
  }
  var dataJs []data = []data{
    {
      Field: "field_1",
      Value: 1,
    },
    {
      Field: "field_2",
      Value: 2,
    },
  }

  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Data: ComplexJavaScriptType{
      AsJSon: dataJs,
    },
    Filter: Filter{
      Field: "field_name",
      Filters: []FilterLine{
        {
          Operator: OPERATOR_CONTAINS,
          Template: &t,
        },
        {
          Field: "field_name",
          Operator: OPERATOR_CONTAINS,
          Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
          Template: &t,
        },
      },
      Logic: LOGIC_AND,
      Template: &t,
    },
    Group: Group{
      Aggregates: []AggregateLine{
        {
          Field: "field_name",
          Aggregate: AGGREGATE_MAX,
          Template: &t,
        },
        {
          Field: "field_name_2",
          Aggregate: AGGREGATE_COUNT,
          Template: &t,
        },
        {
          Field: "field_name_3",
          Aggregate: AGGREGATE_MIN,
          Template: &t,
        },
      },
      Template: &t,
    },
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}

func ExampleDataSource_String7() {
  type data struct {
    Field string
    Value int
  }
  var dataJs []data = []data{
    {
      Field: "field_1",
      Value: 1,
    },
    {
      Field: "field_2",
      Value: 2,
    },
  }

  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Data: ComplexJavaScriptType{
      AsJSon: dataJs,
    },
    Filter: Filter{
      Field: "field_name",
      Filters: []FilterLine{
        {
          Operator: OPERATOR_CONTAINS,
          Template: &t,
        },
        {
          Field: "field_name",
          Operator: OPERATOR_CONTAINS,
          Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
          Template: &t,
        },
      },
      Logic: LOGIC_AND,
      Template: &t,
    },
    Group: Group{
      Aggregates: []AggregateLine{
        {
          Field: "field_name",
          Aggregate: AGGREGATE_MAX,
          Template: &t,
        },
        {
          Field: "field_name_2",
          Aggregate: AGGREGATE_COUNT,
          Template: &t,
        },
        {
          Field: "field_name_3",
          Aggregate: AGGREGATE_MIN,
          Template: &t,
        },
      },
      Template: &t,
    },
    OfflineStorage: ComplexJavaScriptType{
      AsString: "products-offline",
    },
    Page: 0,
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}
func ExampleDataSource_String8() {
  type data struct {
    Field string
    Value int
  }
  var dataJs []data = []data{
    {
      Field: "field_1",
      Value: 1,
    },
    {
      Field: "field_2",
      Value: 2,
    },
  }

  var d DataSource = DataSource{
    Aggregate: []AggregateLine{
      {
        Aggregate: AGGREGATE_MAX,
        Field: "fild_name",
        Template: &t,
      },
    },
    AutoSync: true,
    Batch: true,
    Data: ComplexJavaScriptType{
      AsJSon: dataJs,
    },
    Filter: Filter{
      Field: "field_name",
      Filters: []FilterLine{
        {
          Operator: OPERATOR_CONTAINS,
          Template: &t,
        },
        {
          Field: "field_name",
          Operator: OPERATOR_CONTAINS,
          Value: ComplexJavaScriptType{ AsFunction: "new Date( 1980, 1, 1 )" },
          Template: &t,
        },
      },
      Logic: LOGIC_AND,
      Template: &t,
    },
    Group: Group{
      Aggregates: []AggregateLine{
        {
          Field: "field_name",
          Aggregate: AGGREGATE_MAX,
          Template: &t,
        },
        {
          Field: "field_name_2",
          Aggregate: AGGREGATE_COUNT,
          Template: &t,
        },
        {
          Field: "field_name_3",
          Aggregate: AGGREGATE_MIN,
          Template: &t,
        },
      },
      Template: &t,
    },
    OfflineStorage: ComplexJavaScriptType{
      AsString: "products-offline",
    },
    Page: 0,
    PageSize: 5,
    Schema: Schema{
      Data: ComplexJavaScriptType{
        AsFunction: `function(response) { return response.results; }`,
      },
      Template: &t,
    },
    Sort: []SortLine{
      {
        Dir:  DIR_DESC,
        Field: "field_name",
        Compare: ComplexJavaScriptType{
          AsFunction: `function(a, b) { return numbers[a.item] - numbers[b.item]; }`,
        },
        Template: &t,
      },
    },
    Template: &t,
  }
  fmt.Print( d.String() )

  // Output:
  //
}
