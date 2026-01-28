package go_validate

type ValidateRule string

const (
	Required            ValidateRule = "require"
	LengthMax           ValidateRule = "lengthMax"
	LengthMin           ValidateRule = "lengthMin"
	LengthBetween       ValidateRule = "lengthBetween"
	ValueMin            ValidateRule = "valueMin"
	ValueMax            ValidateRule = "valueMax"
	ValueBetween        ValidateRule = "valueBetween"
	Regex               ValidateRule = "regex"
	Numeric             ValidateRule = "numeric"
	AlphaNumeric        ValidateRule = "alphaNumeric"
	Alpha               ValidateRule = "alpha"
	TypeJSON            ValidateRule = "typeJSON"
	Selected            ValidateRule = "selected"
	SelectLengthMin     ValidateRule = "selectLengthMin"
	SelectLengthMax     ValidateRule = "selectLengthMax"
	SelectLengthBetween ValidateRule = "selectLengthBetween"
	Custom              ValidateRule = "custom"
)

type ValidateRulesMap struct {
	Rule    ValidateRule
	Message string
	Data    []interface{}
}

type ValidateRulesMaps map[string][]ValidateRulesMap

type Validate struct {
	Rules ValidateRulesMaps
}

type ValidateResponse struct {
	Status         bool
	ValidateErrors map[string][]string
}
