package kendo

type LogicEnum int

const(
  FILTER_LOGIC_AND LogicEnum = iota
  FILTER_LOGIC_OR
)

var LogicEnums = [...]string{
  "and",
  "or",
}

func (el LogicEnum) String() string {
  return LogicEnums[ el ]
}
