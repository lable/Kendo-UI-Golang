package kendo

// The sort order (direction).
// The supported values are: "asc" (ascending order) or "desc" (descending order)
type DirEnum int

const(
  DIR_ASC DirEnum = iota
  DIR_DESC
)

var DirEnums = [...]string{
  "asc",
  "desc",
}

func (el DirEnum) String() string {
  return DirEnums[ el ]
}