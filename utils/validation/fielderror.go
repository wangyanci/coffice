package validation

import (
	"fmt"
	"sort"
	"strings"
)

const CurrentField = ""

type FieldError struct {
	Message string
	Paths   []string
	Details string
	errors  []FieldError
}

var _ error = (*FieldError)(nil)

func ErrMissingField(fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: "missing field(s)",
		Paths:   fieldPaths,
	}
}

func ErrDisallowedFields(fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: "must not set the field(s)",
		Paths:   fieldPaths,
	}
}

func ErrDisallowedUpdateDeprecatedFields(fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: "must not update deprecated field(s)",
		Paths:   fieldPaths,
	}
}

func ErrDisallowedUpdateFields(fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: "must not update field(s)",
		Paths:   fieldPaths,
	}
}

func ErrInvalidArrayValue(value interface{}, field string, index int) *FieldError {
	return ErrInvalidValue(value, CurrentField).ViaFieldIndex(field, index)
}

func ErrInvalidValue(value interface{}, fieldPath string) *FieldError {
	return &FieldError{
		Message: fmt.Sprintf("invalid value: %v", value),
		Paths:   []string{fieldPath},
	}
}

func ErrGeneric(diagnostic string, fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: diagnostic,
		Paths:   fieldPaths,
	}
}

func ErrMissingOneOf(fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: "expected exactly one, got neither",
		Paths:   fieldPaths,
	}
}

func ErrMultipleOneOf(fieldPaths ...string) *FieldError {
	return &FieldError{
		Message: "expected exactly one, got both",
		Paths:   fieldPaths,
	}
}

func ErrInvalidKeyName(key, fieldPath string, details ...string) *FieldError {
	return &FieldError{
		Message: fmt.Sprintf("invalid key name %q", key),
		Paths:   []string{fieldPath},
		Details: strings.Join(details, ", "),
	}
}

func ErrOutOfBoundsValue(value, lower, upper interface{}, fieldPath string) *FieldError {
	return &FieldError{
		Message: fmt.Sprintf("expected %v <= %v <= %v", lower, value, upper),
		Paths:   []string{fieldPath},
	}
}

func (fe *FieldError) SetDetails(format string, a ...interface{}) *FieldError {
	fe.Details = fmt.Sprintf(format, a...)
	return fe
}

func (fe *FieldError) ViaField(prefix ...string) *FieldError {
	if fe == nil {
		return nil
	}

	newErr := &FieldError{
		Message: fe.Message,
		Details: fe.Details,
	}

	newPaths := make([]string, 0, len(fe.Paths))
	for _, oldPath := range fe.Paths {
		newPaths = append(newPaths, flatten(append(prefix, oldPath)))
	}
	newErr.Paths = newPaths
	for _, e := range fe.errors {
		newErr = newErr.Also(e.ViaField(prefix...))
	}
	return newErr
}

func (fe *FieldError) ViaIndex(index int) *FieldError {
	return fe.ViaField(asIndex(index))
}

func (fe *FieldError) ViaFieldIndex(field string, index int) *FieldError {
	return fe.ViaIndex(index).ViaField(field)
}

func (fe *FieldError) ViaKey(key interface{}) *FieldError {
	return fe.ViaField(asKey(key))
}

// ViaFieldKey is the short way to chain: err.ViaKey(bar).ViaField(foo)
func (fe *FieldError) ViaFieldKey(field string, key interface{}) *FieldError {
	return fe.ViaKey(key).ViaField(field)
}

// Also collects errors, returns a new collection of existing errors and new errors.
func (fe *FieldError) Also(errs ...*FieldError) *FieldError {
	var newErr *FieldError
	// collect the current objects errors, if it has any
	if !fe.isEmpty() {
		newErr = fe.DeepCopy()
	} else {
		newErr = &FieldError{}
	}
	// and then collect the passed in errors
	for _, e := range errs {
		if !e.isEmpty() {
			newErr.errors = append(newErr.errors, *e)
		}
	}
	if newErr.isEmpty() {
		return nil
	}
	return newErr
}

func (in *FieldError) DeepCopy() *FieldError {
	if in == nil {
		return nil
	}
	out := new(FieldError)
	in.DeepCopyInto(out)
	return out
}

func (in *FieldError) DeepCopyInto(out *FieldError) {
	*out = *in
	if in.Paths != nil {
		in, out := &in.Paths, &out.Paths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.errors != nil {
		in, out := &in.errors, &out.errors
		*out = make([]FieldError, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (fe *FieldError) isEmpty() bool {
	if fe == nil {
		return true
	}
	return fe.Message == "" && fe.Details == "" && len(fe.errors) == 0 && len(fe.Paths) == 0
}

func (fe *FieldError) normalized() []*FieldError {
	if fe == nil {
		return []*FieldError(nil)
	}

	errors := make([]*FieldError, 0, len(fe.errors)+1)
	if fe.Message != "" {
		errors = append(errors, &FieldError{
			Message: fe.Message,
			Paths:   fe.Paths,
			Details: fe.Details,
		})
	}
	for _, e := range fe.errors {
		errors = append(errors, e.normalized()...)
	}
	return errors
}

func (fe *FieldError) Error() string {
	normedErrors := merge(fe.normalized())
	errs := make([]string, 0, len(normedErrors))
	for _, e := range normedErrors {
		if e.Details == "" {
			errs = append(errs, fmt.Sprintf("%v: %v", e.Message, strings.Join(e.Paths, ", ")))
		} else {
			errs = append(errs, fmt.Sprintf("%v: %v(%v)", e.Message, strings.Join(e.Paths, ", "), e.Details))
		}
	}
	//return strings.Join(errs, "\n")
	return strings.Join(errs, "; ")
}

func asIndex(index int) string {
	return fmt.Sprintf("[%d]", index)
}

func isIndex(part string) bool {
	return strings.HasPrefix(part, "[") && strings.HasSuffix(part, "]")
}

func asKey(key interface{}) string {
	return fmt.Sprintf("[%+v]", key)
}

func flatten(path []string) string {
	var newPath []string
	for _, part := range path {
		for _, p := range strings.Split(part, ".") {
			if p == CurrentField {
				continue
			} else if len(newPath) > 0 && isIndex(p) {
				newPath[len(newPath)-1] += p
			} else {
				newPath = append(newPath, p)
			}
		}
	}
	return strings.Join(newPath, ".")
}

func mergePaths(a, b []string) []string {
	newPaths := make([]string, 0, len(a)+len(b))
	newPaths = append(newPaths, a...)
	for _, bi := range b {
		if !containsString(newPaths, bi) {
			newPaths = append(newPaths, bi)
		}
	}
	return newPaths
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func merge(errs []*FieldError) []*FieldError {
	m := make(map[string]*FieldError, len(errs))

	for _, e := range errs {
		k := key(e)
		if v, ok := m[k]; ok {
			v.Paths = mergePaths(v.Paths, e.Paths)
		} else {
			m[k] = e
		}
	}

	newErrs := make([]*FieldError, 0, len(m))
	for _, v := range m {
		sort.Slice(v.Paths, func(i, j int) bool { return v.Paths[i] < v.Paths[j] })
		newErrs = append(newErrs, v)
	}

	sort.Slice(newErrs, func(i, j int) bool {
		if newErrs[i].Message == newErrs[j].Message {
			return newErrs[i].Details < newErrs[j].Details
		}
		return newErrs[i].Message < newErrs[j].Message
	})

	return newErrs
}

func key(err *FieldError) string {
	return fmt.Sprintf("%s-%s", err.Message, err.Details)
}


