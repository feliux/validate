Functional Go implementation for validating params. Useful for http handlers!

**Usage**

```go
type AccountSetupParams struct {
	Username string
}

type AccountSetupErrors struct {
	Username string
}

func HandleAccountSetup(w http.ResponseWriter, r *http.Request) error {
	var errors auth.AccountSetupErrors
	params := auth.AccountSetupParams{
		Username: r.FormValue("username"),
	}
	ok := validations.New(&params, validations.Fields{
		"Username": validations.Rules(validations.Min(2), validations.Max(50)),
	}).Validate(&errors)
	if !ok {
		// handle !ok
	}
    // handle ok ...
}
```

```go
type SignupParams struct {
	Email           string
	Password        string
	ConfirmPassword string
}

type SignupErrors struct {
	Email           string
	Password        string
	ConfirmPassword string
}

func HandleSignupForm(w http.ResponseWriter, r *http.Request) error {
	signupParams := auth.SignupParams{
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirmPassword"),
	}
	errors := auth.SignupErrors{}
	if ok := validations.New(&signupParams, validations.Fields{
		"Email":    validations.Rules(validations.Email),
		"Password": validations.Rules(validations.Password),
		"ConfirmPassword": validations.Rules(
			validations.Equal(signupParams.Password),
			validations.Message("Passwords do not match"),
		),
	}).Validate(&errors); !ok {
        // handle !ok
	}
    // handle ok ...
}
```
