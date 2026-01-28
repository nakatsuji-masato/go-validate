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