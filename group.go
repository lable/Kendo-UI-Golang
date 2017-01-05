package kendo

// The grouping configuration of the data source. If set, the data items will be grouped when the data source is
// populated. By default, grouping is not applied.
// The data source groups the data items client-side unless the serverGrouping option is set to true.
type Group struct {
  // The aggregates which are calculated during grouping.
  // The supported aggregates are: "average", "count", "max", "min", "sum"
  Aggregates      []AggregateList

  // The sort order of the group.
  // The supported values are: "asc" (ascending order) or "desc" (descending order)
  // default: "asc"
  Dir             DirEnum

  // The data item field to group by.
  Field           string
}