package vo

import "regexp"

var (
	regexGolangCodeGenerated = regexp.MustCompile(`^// Code generated .*; DO NOT EDIT.\n`)
)

// Option Option Value Object
type Option struct {
	isForce bool
}

// Option Option Value Object Constructor
func NewOption(isForce bool) *Option {
	o := &Option{isForce}
	return o
}

// IsForce force option enable/disable decision
func (o Option) IsForce() bool {
	return o.isForce
}

// IsOverride override enable/disable decision
func (o Option) IsOverride(source string) bool {
	if o.IsForce() {
		return true
	}
	return regexGolangCodeGenerated.MatchString(source)
}
