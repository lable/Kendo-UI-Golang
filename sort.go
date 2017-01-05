package kendo

type Sort struct {
  // The sort order (direction).
  // The supported values are: "asc" (ascending order) or "desc" (descending order)
  Dir             DirEnum

  // The field by which the data items are sorted.
  Field           string

  // Function which can be used for custom comparing of the DataSource items.
  Compare         string
}
