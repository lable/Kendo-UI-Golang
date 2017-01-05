package kendo

// The nested filter expressions. Supports the same options as filter. Filters can be nested indefinitely.
type FilterLine struct {
  // The data item field to which the filter operator is applied.
  Field     string
  // The filter operator (comparison).
  // The supported operators are: "eq" (equal to), "neq" (not equal to), "isnull" (is equal to null), "isnotnull" (is
  // not equal to null), "lt" (less than), "lte" (less than or equal to), "gt" (greater than), "gte" (greater than or
  // equal to), "startswith", "endswith", "contains", "doesnotcontain", "isempty", "isnotempty"
  // The last five are supported only for string fields.
  Operator  OperatorEnum
  // The value to which the field is compared. The value has to be of the same type as the field.
  // By design, the "\n" is removed from the filter before the filtering is performed. That is why an "\n" identifier
  // from the filter will not match data items whose corresponding fields contain new lines.
  Value     JSonString
}

// The filters which are applied over the data items. By default, no filter is applied.
// The data source filters the data items client-side unless the serverFiltering option is set to true.
type Filter struct {
  // The data item field to which the filter operator is applied.
  Field     string
  // The nested filter expressions. Supports the same options as filter. Filters can be nested indefinitely.
  Filters   []FilterLine
  // The logical operation to use when the filter.filters option is set.
  // The supported values are: "and" or "or"
  Logic     LogicEnum
  // The filter operator (comparison).
  // The supported operators are: "eq" (equal to), "neq" (not equal to), "isnull" (is equal to null), "isnotnull" (is
  // not equal to null), "lt" (less than), "lte" (less than or equal to), "gt" (greater than), "gte" (greater than or
  // equal to), "startswith", "endswith", "contains", "doesnotcontain", "isempty", "isnotempty"
  // The last five are supported only for string fields.
  Operator  OperatorEnum
  // The value to which the field is compared. The value has to be of the same type as the field.
  // By design, the "\n" is removed from the filter before the filtering is performed. That is why an "\n" identifier
  // from the filter will not match data items whose corresponding fields contain new lines.
  Value     JSonString
}
