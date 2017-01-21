package kendo

type OperatorEnum int

const(
  OPERATOR_NIL OperatorEnum   = iota
  OPERATOR_EQ
  OPERATOR_NEQ
  OPERATOR_ISNULL
  OPERATOR_ISNOTNULL
  OPERATOR_LT
  OPERATOR_LTE
  OPERATOR_GT
  OPERATOR_GTE
  OPERATOR_STARTSWITH
  OPERATOR_ENDSWITH
  OPERATOR_CONTAINS
  OPERATOR_DOESNOTCONTAIN
  OPERATOR_ISEMPTY
  OPERATOR_ISNOTEMPTY
)

var OperatorEnums  = [...]string{
  "",
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

func (el OperatorEnum ) String() string {
  return OperatorEnums[ el ]
}

