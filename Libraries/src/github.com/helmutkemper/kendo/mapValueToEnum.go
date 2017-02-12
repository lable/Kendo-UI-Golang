package kendo

type MapValueToEnum int

const(
  MAP_VALUE_TO_NIL MapValueToEnum   = iota
  MAP_VALUE_TO_INDEX
  MAP_VALUE_TO_DATA_ITEM
)

var MapValueToEnums  = [...]string{
  "",
  "index",
  "dataItem",
}

func (el MapValueToEnum ) String() string {
  return MapValueToEnums[ el ]
}
