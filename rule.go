package go_validate

import (
	"encoding/json"
	"regexp"
)

// Rule Required
func ruleRequired(value any) bool {
	if value == nil || value == "" {
		return false
	}
	return true
}

// Rule Length Min
func ruleLengthMin(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}

	var min int = data[0].(int)
	if len(value.(string)) < min {
		return false
	}
	return true
}

// Rule Length Max
func ruleLengthMax(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}

	max := data[0].(int)
	if len(value.(string)) > max {
		return false
	}
	return true
}

// ruleLengthBetween
func ruleLengthBetween(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	judge1 := ruleLengthMin(value, []interface{}{data[0]})
	judge2 := ruleLengthMax(value, []interface{}{data[1]})

	if judge1 == false || judge2 == false {
		return false
	}

	return true
}

// Rule Value Min
func ruleValueMin(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	val, ok := _toInt(value)
	if !ok {
		return false
	}
	min, ok := _toInt(data[0])
	if !ok {
		return false
	}
	if val < min {
		return false
	}
	return true
}

// Rule Value Max
func ruleValueMax(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	val, ok := _toInt(value)
	if !ok {
		return false
	}
	max, ok := _toInt(data[0])
	if !ok {
		return false
	}
	if val > max {
		return false
	}
	return true
}

// Rule Value Between
func ruleValueBetween(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	judge1 := ruleValueMin(value, []interface{}{data[0]})
	judge2 := ruleValueMax(value, []interface{}{data[1]})

	if judge1 == false || judge2 == false {
		return false
	}
	return true
}

// Rule Regex
func ruleRegex(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	target := data[0].(string)
	re := regexp.MustCompile(target)
	if re.MatchString(value.(string)) == false {
		return false
	}
	return true
}

// Rule Numeric
func ruleNumeric(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	target := `^[0-9]*$`
	if len(data) >= 1 {
		addChars := data[0].(string)
		target = `^[0-9` + addChars + `]*$`
	}
	return ruleRegex(value, []interface{}{target})
}

// Rule Alpha Numeric
func ruleAlphaNumeric(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	target := `^[A-Za-z0-9]*$`
	if len(data) >= 1 {
		addChars := data[0].(string)
		target = `^[A-Za-z0-9` + addChars + `]*$`
	}
	return ruleRegex(value, []interface{}{target})
}

// Rule Alpha
func ruleAlpha(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	target := `^[A-Za-z]*$`
	if len(data) >= 1 {
		addChars := data[0].(string)
		target = `^[A-Za-z` + addChars + `]*$`
	}
	return ruleRegex(value, []interface{}{target})
}

// Rule Type Json
func ruleTypeJSON(value any) bool {
	if value == nil || value == "" {
		return true
	}
	return json.Valid([]byte(value.(string)))
}

// Rule Selected
func ruleSelected(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	judge := false
	for _, name := range data {
		if name == value {
			judge = true
			break
		}
	}
	return judge
}

// Rule Select Length Min
func ruleSelectLengthMin(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	values, ok := value.([]any)
	if !ok {
		return true
	}
	min := data[0].(int)
	if len(values) < min {
		return false
	}
	return true
}

// Rule Select Length Max
func ruleSelectLengthMax(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	values, ok := value.([]any)
	if !ok {
		return true
	}
	max := data[0].(int)
	if len(values) > max {
		return false
	}
	return true
}

// Rule Select Length Between
func ruleSelectLengthBetween(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	judge1 := ruleSelectLengthMin(value, []interface{}{data[0]})
	judge2 := ruleSelectLengthMax(value, []interface{}{data[1]})
	if judge1 == false || judge2 == false {
		return false
	}
	return true
}

// Rule Custom
func ruleCustom(value any, data []interface{}) bool {
	if value == nil || value == "" {
		return true
	}
	fn := data[0].(func(any) bool)
	if fn(value) == false {
		return false
	}
	return true
}
