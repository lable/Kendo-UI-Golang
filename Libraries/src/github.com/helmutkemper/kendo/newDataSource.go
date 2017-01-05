package kendo

func NewDataSource( t Template ) *DataSource {
  return &DataSource{
    Template: t,
  }
}
