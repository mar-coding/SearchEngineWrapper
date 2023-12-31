// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: website.proto

package website

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Permission with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Permission) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Permission with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PermissionMultiError, or
// nil if none found.
func (m *Permission) ValidateAll() error {
	return m.validate(true)
}

func (m *Permission) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Optional

	// no validation rules for ValidatePermissions

	// no validation rules for Captcha

	if len(errors) > 0 {
		return PermissionMultiError(errors)
	}

	return nil
}

// PermissionMultiError is an error wrapping multiple validation errors
// returned by Permission.ValidateAll() if the designated constraints aren't met.
type PermissionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PermissionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PermissionMultiError) AllErrors() []error { return m }

// PermissionValidationError is the validation error returned by
// Permission.Validate if the designated constraints aren't met.
type PermissionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PermissionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PermissionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PermissionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PermissionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PermissionValidationError) ErrorName() string { return "PermissionValidationError" }

// Error satisfies the builtin error interface
func (e PermissionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPermission.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PermissionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PermissionValidationError{}

// Validate checks the field values on TestPostAPIRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TestPostAPIRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestPostAPIRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TestPostAPIRequestMultiError, or nil if none found.
func (m *TestPostAPIRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TestPostAPIRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAge() <= 10 {
		err := TestPostAPIRequestValidationError{
			field:  "Age",
			reason: "value must be greater than 10",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _TestPostAPIRequest_Name_NotInLookup[m.GetName()]; ok {
		err := TestPostAPIRequestValidationError{
			field:  "Name",
			reason: "value must not be in list []",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TestPostAPIRequestMultiError(errors)
	}

	return nil
}

// TestPostAPIRequestMultiError is an error wrapping multiple validation errors
// returned by TestPostAPIRequest.ValidateAll() if the designated constraints
// aren't met.
type TestPostAPIRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestPostAPIRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestPostAPIRequestMultiError) AllErrors() []error { return m }

// TestPostAPIRequestValidationError is the validation error returned by
// TestPostAPIRequest.Validate if the designated constraints aren't met.
type TestPostAPIRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestPostAPIRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestPostAPIRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestPostAPIRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestPostAPIRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestPostAPIRequestValidationError) ErrorName() string {
	return "TestPostAPIRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TestPostAPIRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestPostAPIRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestPostAPIRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestPostAPIRequestValidationError{}

var _TestPostAPIRequest_Name_NotInLookup = map[string]struct{}{
	"": {},
}

// Validate checks the field values on TestPostAPIResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TestPostAPIResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestPostAPIResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TestPostAPIResponseMultiError, or nil if none found.
func (m *TestPostAPIResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TestPostAPIResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Output

	if len(errors) > 0 {
		return TestPostAPIResponseMultiError(errors)
	}

	return nil
}

// TestPostAPIResponseMultiError is an error wrapping multiple validation
// errors returned by TestPostAPIResponse.ValidateAll() if the designated
// constraints aren't met.
type TestPostAPIResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestPostAPIResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestPostAPIResponseMultiError) AllErrors() []error { return m }

// TestPostAPIResponseValidationError is the validation error returned by
// TestPostAPIResponse.Validate if the designated constraints aren't met.
type TestPostAPIResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestPostAPIResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestPostAPIResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestPostAPIResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestPostAPIResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestPostAPIResponseValidationError) ErrorName() string {
	return "TestPostAPIResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TestPostAPIResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestPostAPIResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestPostAPIResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestPostAPIResponseValidationError{}

// Validate checks the field values on TestGetAPIRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TestGetAPIRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestGetAPIRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TestGetAPIRequestMultiError, or nil if none found.
func (m *TestGetAPIRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TestGetAPIRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return TestGetAPIRequestMultiError(errors)
	}

	return nil
}

// TestGetAPIRequestMultiError is an error wrapping multiple validation errors
// returned by TestGetAPIRequest.ValidateAll() if the designated constraints
// aren't met.
type TestGetAPIRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestGetAPIRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestGetAPIRequestMultiError) AllErrors() []error { return m }

// TestGetAPIRequestValidationError is the validation error returned by
// TestGetAPIRequest.Validate if the designated constraints aren't met.
type TestGetAPIRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestGetAPIRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestGetAPIRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestGetAPIRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestGetAPIRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestGetAPIRequestValidationError) ErrorName() string {
	return "TestGetAPIRequestValidationError"
}

// Error satisfies the builtin error interface
func (e TestGetAPIRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestGetAPIRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestGetAPIRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestGetAPIRequestValidationError{}

// Validate checks the field values on TestGetAPIResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TestGetAPIResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TestGetAPIResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TestGetAPIResponseMultiError, or nil if none found.
func (m *TestGetAPIResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *TestGetAPIResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Output

	// no validation rules for NumberOfNames

	if len(errors) > 0 {
		return TestGetAPIResponseMultiError(errors)
	}

	return nil
}

// TestGetAPIResponseMultiError is an error wrapping multiple validation errors
// returned by TestGetAPIResponse.ValidateAll() if the designated constraints
// aren't met.
type TestGetAPIResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TestGetAPIResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TestGetAPIResponseMultiError) AllErrors() []error { return m }

// TestGetAPIResponseValidationError is the validation error returned by
// TestGetAPIResponse.Validate if the designated constraints aren't met.
type TestGetAPIResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TestGetAPIResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TestGetAPIResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TestGetAPIResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TestGetAPIResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TestGetAPIResponseValidationError) ErrorName() string {
	return "TestGetAPIResponseValidationError"
}

// Error satisfies the builtin error interface
func (e TestGetAPIResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTestGetAPIResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TestGetAPIResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TestGetAPIResponseValidationError{}
