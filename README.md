# go-validate

A verification check library for the Go language

## Install as a Go library

Execute the following command in the package directory where ``go.mod`` is located to install it as an external library:

```bash
$ go get github.com/nakatsuji-masato/go-validate
```

## tutorial

First, define the Validate structure as shown in the code below.

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Login = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"user": {
			{
				Rule:    gv.Required,
				Message: "user is empty",
			},
			{
				Rule:    gv.LengthBetween,
				Data:    []interface{}{4, 30},
				Message: "user is 4-30 character length.",
			},
		},
		"pass": {
			{
				Rule:    gv.Required,
				Message: "pass is empty",
			},
			{
				Rule:    gv.LengthBetween,
				Data:    []interface{}{4, 30},
				Message: "pass is 4-30 character length.",
			},
		},
	},
}

```

Then, at the point where you actually want to perform the validation check,  
call the ``verify`` function from the variable in the Validate structure above.

```go
body := map[string]string{
    "user": "",
    "pass": "",
}
vres := Login.Verify(body)
fmt.Println(vres)
```

## Validate Rule

The main preset validations are as follows:

|Rule|overview|
|:--|:--|
|[Required](#required)|Required input.<br> An empty string or nil will be treated as an error.|
|[LengthMin](#lengthmin)|An error occurs if the length of the value (string) is less than the specified value.|
|[LengthMax](#lengthmax)|An error occurs if the length of the value (string) exceeds the specified value.|
|[LengthBetween](#lengthbetween)|An error occurs if the length of the value (string) is outside the specified range.|
|[ValueMin](#valuemin)|If the value (numeric value) is less than the specified value, an error occurs.|
|[ValueMax](#valuemax)|If the value (numeric value) exceeds the specified value, an error occurs.|
|[ValueBetween](#valuebetween)|If the value (numeric value) is outside the specified range, an error occurs.|
|[Regex](#regex)|An error occurs if the value (string) does not match the specified regular expression.|
|[Numeric](#numeric)|An error occurs if the value (string) contains any characters other than half-width numbers and the specified characters.|
|[AlphaNumeric](#alphanumeric)|An error occurs if the value (string) contains any characters other than half-width alphanumeric characters and the specified characters.|
|[Alpha](#alpha)|An error occurs if the value (string) contains any characters other than half-width English characters and the specified characters.|
|[TypeJSON](#typejson)|An error occurs if the value is not a JSON type.|
|[Selected](#selected)|If the value is not one of the options, an error occurs.|
|[SelectLengthMin](#selectlengthmin)|An error occurs if the number of values ​​(array values) is less than the specified number.|
|[SelectLengthMax](#selectlengthmax)|An error occurs if the number of values ​​(array values) exceeds the specified number.|
|[SelectLengthBetween](#selectlengthbetween)|An error occurs if the number of values ​​(array values) is outside the specified range.|
|[Custom](#custom)|For conditions other than the above presets, use custom validation.|

### Required

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Required,
			},
		},
	},
}``
```

### lengthMin

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.LengthMin,
				Data: []interface{}{4}, // Error if less than 4 characters
			},
		},
	},
}``
```

### LengthMax


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.LengthMax,
				Data: []interface{}{20}, // If it is more than 20 characters, an error occurs.
			},
		},
	},
}``
```

### LengthBetween


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.LengthBetween,
				Data: []interface{}{4, 20}, // If the number is outside the range of 4 to 20 characters, an error occurs.
			},
		},
	},
}``
```

### ValueMin

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.ValueMin,
				Data: []interface{}{6}, // If the value is less than 6, an error occurs.
			},
		},
	},
}``
```

### ValueMax


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.ValueMax,
				Data: []interface{}{255}, // If the value is 255 or greater, an error occurs.
			},
		},
	},
}``
```

### ValueBetween


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.ValueBetween,
				Data: []interface{}{6, 255}, // If the value is outside the range of 6 to 255, an error occurs..
			},
		},
	},
}``
```

### Regex


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Regex,
				Data: []interface{}{"`^[0-1]*$`"}, // If the value contains anything other than 0 and 1 characters, an error occurs.
			},
		},
	},
}``
```

### Numeric

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Numeric,
				// If the value contains anything other than half-width numbers, an error occurs.
			},
		},
	},
}``
```

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Numeric,
				Data: []interface{}{" _-"},
				// An error occurs if the value contains anything other than half-width numbers, half-width spaces, - and _ characters.
			},
		},
	},
}``
```

### AlphaNumeric


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.AlphaNumeric,
				// If the value contains anything other than half-width alpha numbers, an error occurs.
			},
		},
	},
}``
```

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.AlphaNumeric,
				Data: []interface{}{" _-"},
				// An error occurs if the value contains anything other than half-width alpha numbers, half-width spaces, - and _ characters.
			},
		},
	},
}``
```

### Alpha


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Alpha,
				// If the value contains anything other than half-width alpha, an error occurs.
			},
		},
	},
}``
```

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Alpha,
				Data: []interface{}{" _-"},
				// An error occurs if the value contains anything other than half-width alpha, half-width spaces, - and _ characters.
			},
		},
	},
}``
```

### TypeJson


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.TypeJson,
				// Error if the value is not a JSON string.
			},
		},
	},
}``
```

### Selected


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Selected,
				Data: []interface{"apple", "mango", "orange"},
				// If the value is anything other than apple, mango, or orange, an error occurs.
			},
		},
	},
}``
```

### SelectLengthMin


```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.SelectLengthMin,
				Data: []interface{3},
				// An error occurs if the number of values ​​(array values) is less than 3.
			},
		},
	},
}``
```

### SelectLengthMax

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.SelectLengthMax,
				Data: []interface{10},
				// An error occurs if the number of values ​​(array values) is 10 or more.
			},
		},
	},
}``
```

### SelectLengthBetween

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.SelectLengthBetween,
				Data: []interface{3, 10},
				// If the number of values ​​(array values) is outside the range of 3 to 10, an error occurs.
			},
		},
	},
}``
```

### Custom

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Custom,
				Data: []interface{func (value any) bool {

					if value.(string) == "hoge" {
						return true
					}

					return false
				}},
				// If the value is not "hoge", an error occurs.
			},
		},
	},
}``
```

## Error message settings

You can set an error message using Message as shown in the code below.

```go
import (
	gv "github.com/nakatsuji-masato/go-validate"
)

var Validate = gv.Validate{
	Rules: gv.ValidateRulesMaps{
		"value1": {
			{
				Rule: gv.Required,
				Message: "value1 is empty",
			},
		},
	},
}``
```
