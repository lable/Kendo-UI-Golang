package kendo

type OperatorEnum int

const(
  EQ  OperatorEnum = iota
  NEQ
  ISNULL
  ISNOTNULL
  LT
  LTE
  GT
  GTE
  STARTSWITH
  ENDSWITH
  CONTAINS
  DOESNOTCONTAIN
  ISEMPTY
  ISNOTEMPTY
)

var OperatorEnums = [...]string{
  "eq",
  "neq",
  "isnull",
  "isnotnull",
  "lt",
  "lte",
  "gt",
  "gte",
  "startswith",
  "endswith",
  "contains",
  "doesnotcontain",
  "isempty",
  "isnotempty",
}

func (el OperatorEnum) String() string {
  return OperatorEnums[ el ]
}
