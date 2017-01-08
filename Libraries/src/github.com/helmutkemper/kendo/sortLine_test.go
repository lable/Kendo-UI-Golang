package kendo

import "fmt"

func ExampleSort_String() {
  var s SortLine = SortLine{
    Dir:  DIR_DESC,
    Template: &t,
  }
  fmt.Print( s.String() )

  // Output:
  // dir: "desc",
}

func ExampleSort_String2() {
  var s SortLine = SortLine{
    Field: "field_name",
    Template: &t,
  }
  fmt.Print( s.String() )

  // Output:
  // field: "field_name",
}

func ExampleSort_String3() {
  var s SortLine = SortLine{
    Compare: ComplexJavaScriptType{
      AsFunction: `function(a, b) { return numbers[a.item] - numbers[b.item]; }`,
    },
    Template: &t,
  }
  fmt.Print( s.String() )

  // Output:
  // compare: function(a, b) { return numbers[a.item] - numbers[b.item]; }
}

func ExampleSort_String4() {
  var s SortLine = SortLine{
    Dir:  DIR_DESC,
    Field: "field_name",
    Template: &t,
  }
  fmt.Print( s.String() )

  // Output:
  // dir: "desc", field: "field_name",
}



func ExampleSort_String5() {
  var s SortLine = SortLine{
    Dir:  DIR_DESC,
    Field: "field_name",
    Compare: ComplexJavaScriptType{
      AsFunction: `function(a, b) { return numbers[a.item] - numbers[b.item]; }`,
    },
    Template: &t,
  }
  fmt.Print( s.String() )

  // Output:
  // dir: "desc", field: "field_name", compare: function(a, b) { return numbers[a.item] - numbers[b.item]; }
}

