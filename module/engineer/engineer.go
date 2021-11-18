package engineer

type Engineer struct {
	FirstName string
	LastName  string
	Salary    float64
	Position  string
}

func NewEngineer(firstName, lastName string) *Engineer {
	return &Engineer{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// params would be very long indicate code smells so we can use option struct approach here

type Option struct {
	Salary   float64
	Position string
}

func NewEngineerWithOptions(firstname, lastname string, opt Option) *Engineer {
	// Default or value pass by args
	engineer := &Engineer{
		FirstName: firstname,
		LastName:  lastname,
	}

	// Overwrite with the option or adding more field
	if opt.Salary != 0 {
		engineer.Salary = opt.Salary
	}

	if opt.Position != "" {
		engineer.Position = opt.Position
	}
	return engineer
}

// Or we can even make all option fields as a pointer so we don't have to concern
// about zero values
type OptionPointers struct {
	Salary   *float64
	Position *string
}

// However fallback on this on is you have to declare the variable since
// go doesn't support &"string" or &1 so you have to
// i := 1 then &i which make it very ugly
// You can see that AWS use this approach to as we have flightlog/helper.go
// to convert string or int to pointer when we using aws sdk
func NewEngineerWithOptionPointers(firstname, lastname string, opt OptionPointers) *Engineer {
	eng := &Engineer{
		FirstName: firstname,
		LastName:  lastname,
	}
	if opt.Salary != nil {
		eng.Salary = *opt.Salary
	}
	if opt.Position != nil {
		eng.Position = *opt.Position
	}
	return eng
}

// Last is functional options

// We define the new type for Optional function
type OptionFn func(eng *Engineer)

func WithSalary(s float64) OptionFn {
	return func(eng *Engineer) {
		eng.Salary = s
	}
}

func WithPosition(p string) OptionFn {
	return func(eng *Engineer) {
		eng.Position = p
	}
}

// Then for New method we can leave the not require field to optional function
func NewEngineerWithOptFn(firstname, lastname string, opts ...OptionFn) *Engineer {
	eng := &Engineer{
		FirstName: firstname,
		LastName:  lastname,
	}
	for _, applyOpt := range opts {
		applyOpt(eng)
	}

	return eng
}
