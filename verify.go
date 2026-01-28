package go_validate

import (
	"encoding/json"
)

// 検証チェック開始
func Verify(post interface{}, validateRUles ValidateRulesMaps) ValidateResponse {
	v_ := Validate{
		Rules: validateRUles,
	}
	result := v_.Verify(post)
	return result
}

// 検証チェック開始
func (v *Validate) Verify(post interface{}) ValidateResponse {

	result := ValidateResponse{
		Status:         true,
		ValidateErrors: make(map[string][]string),
	}

	postMap, _ := StructToMap(post)

	for name, rule := range v.Rules {
		val, ok := postMap[name]
		if !ok {
			val = nil
		}

		judge := true
		var vMsg []string
		for _, r_ := range rule {
			var judge_ bool

			switch r_.Rule {
			case Required:
				judge_ = ruleRequired(val)
			case LengthMin:
				judge_ = ruleLengthMin(val, r_.Data)
			case LengthMax:
				judge_ = ruleLengthMax(val, r_.Data)
			case LengthBetween:
				judge_ = ruleLengthBetween(val, r_.Data)
			case ValueMin:
				judge_ = ruleValueMin(val, r_.Data)
			case ValueMax:
				judge_ = ruleValueMax(val, r_.Data)
			case ValueBetween:
				judge_ = ruleValueBetween(val, r_.Data)
			case Regex:
				judge_ = ruleRegex(val, r_.Data)
			case Numeric:
				judge_ = ruleNumeric(val, r_.Data)
			case AlphaNumeric:
				judge_ = ruleAlphaNumeric(val, r_.Data)
			case Alpha:
				judge_ = ruleAlpha(val, r_.Data)
			case TypeJSON:
				judge_ = ruleTypeJSON(val)
			case Selected:
				judge_ = ruleSelected(val, r_.Data)
			case SelectLengthMin:
				judge_ = ruleSelectLengthMin(val, r_.Data)
			case SelectLengthMax:
				judge_ = ruleSelectLengthMax(val, r_.Data)
			case SelectLengthBetween:
				judge_ = ruleSelectLengthBetween(val, r_.Data)
			case Custom:
				judge_ = ruleCustom(val, r_.Data)
			}

			if judge_ == false {
				judge = judge_
				if r_.Message == "" {
					dataStr, _ := json.Marshal(r_.Data)
					vMsg = append(vMsg, "rule = "+string(r_.Rule)+", Data = "+string(dataStr))
				} else {
					vMsg = append(vMsg, r_.Message)
				}
			}
		}

		if judge == false {
			result.Status = false
			result.ValidateErrors[name] = vMsg
		}
	}

	return result
}

// Struct to Map Convert
func StructToMap(v interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	return m, err
}

// convert interface to int
func _toInt(v interface{}) (int, bool) {
	if i, ok := v.(int); ok {
		return i, true
	}
	if f, ok := v.(float64); ok {
		return int(f), true
	}
	return 0, false
}
