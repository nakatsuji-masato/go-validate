package go_validate

import (
	"fmt"
)

var TestValidate = Validate{
	Rules: ValidateRulesMaps{
		"username": {
			{
				Rule:    Required,
				Message: "username is empty",
			},
			{
				Rule:    LengthMin,
				Data:    []interface{}{4},
				Message: "username under 4 character length.",
			},
			{
				Rule:    LengthMax,
				Data:    []interface{}{30},
				Message: "username over 30 character length.",
			},
		},
		"password": {
			{
				Rule:    Required,
				Message: "password is empty",
			},
			{
				Rule:    LengthMin,
				Data:    []interface{}{4},
				Message: "password under 4 character length.",
			},
			{
				Rule:    LengthMax,
				Data:    []interface{}{30},
				Message: "password over 30 character length.",
			},
		},
		"email": {
			{
				Rule:    Required,
				Message: "Email is empty",
			},
			{
				Rule: LengthBetween,
				Data: []interface{}{4, 20},
			},
		},
		"age": {
			{
				Rule:    ValueMin,
				Data:    []interface{}{5},
				Message: "age is under 5",
			},
			{
				Rule:    ValueMax,
				Data:    []interface{}{20},
				Message: "age is over 20",
			},
		},
		"number": {
			{

				Rule:    Numeric,
				Data:    []interface{}{"abc"},
				Message: "numeric error",
			},
		},
		"json": {
			{
				Rule:    TypeJSON,
				Message: "JSONではない",
			},
		},
		"select": {
			{
				Rule:    Selected,
				Data:    []interface{}{"mikan", "ringo", "kiwi"},
				Message: "選択肢以外の値が入力されています。",
			},
		},
		"selectMin": {
			{
				Rule:    SelectLengthMin,
				Data:    []interface{}{3},
				Message: "最低でも3つは選択してください",
			},
		},
		"selectMax": {
			{
				Rule:    SelectLengthMax,
				Data:    []interface{}{6},
				Message: "最大6つまでしか選択できません",
			},
		},
		"selectBetween": {
			{
				Rule:    SelectLengthBetween,
				Data:    []interface{}{2, 4},
				Message: "2-4の範囲で選択してください",
			},
		},
		"custom": {
			{
				Rule: Custom,
				Data: []interface{}{func(value any) bool {
					if value.(string) == "custom message" {
						return true
					}
					return false
				}},
				Message: "カスタムエラーメッセージ",
			},
		},
	},
}

func VerifyTest() {

	post := map[string]any{
		/*
				"username": "",
				"password": "aveee",
				"email":    "xxxx xxxx xxxx xxxx x",
				"age":      20,
				"number":   "1234e",
				"json":     `{"type":1}aa`,
			"select":  "kiwi2",
		*/
		"selectMin":     []string{"aaa", "bbb", "ccc"},
		"selectMax":     []string{"aaa", "bbb", "ccc", "ddd", "eee", "fff"},
		"selectBetween": []string{"aa", "bbb", "cc", "dd"},
		"custom":        "custom message...",
	}

	// 検証チェック
	vres := TestValidate.Verify(post)

	fmt.Println(vres.Status)
	fmt.Println(vres.ValidateErrors)
}
