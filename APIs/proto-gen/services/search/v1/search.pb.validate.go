// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: search.proto

package search

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

// Validate checks the field values on ListItemsSearchResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListItemsSearchResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListItemsSearchResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListItemsSearchResponseMultiError, or nil if none found.
func (m *ListItemsSearchResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListItemsSearchResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListItemsSearchResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListItemsSearchResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListItemsSearchResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PageSize

	// no validation rules for PageNo

	// no validation rules for TotalItemsCount

	if len(errors) > 0 {
		return ListItemsSearchResponseMultiError(errors)
	}

	return nil
}

// ListItemsSearchResponseMultiError is an error wrapping multiple validation
// errors returned by ListItemsSearchResponse.ValidateAll() if the designated
// constraints aren't met.
type ListItemsSearchResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListItemsSearchResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListItemsSearchResponseMultiError) AllErrors() []error { return m }

// ListItemsSearchResponseValidationError is the validation error returned by
// ListItemsSearchResponse.Validate if the designated constraints aren't met.
type ListItemsSearchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListItemsSearchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListItemsSearchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListItemsSearchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListItemsSearchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListItemsSearchResponseValidationError) ErrorName() string {
	return "ListItemsSearchResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListItemsSearchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListItemsSearchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListItemsSearchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListItemsSearchResponseValidationError{}

// Validate checks the field values on Item with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Item with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ItemMultiError, or nil if none found.
func (m *Item) ValidateAll() error {
	return m.validate(true)
}

func (m *Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	// no validation rules for Description

	// no validation rules for Url

	if len(errors) > 0 {
		return ItemMultiError(errors)
	}

	return nil
}

// ItemMultiError is an error wrapping multiple validation errors returned by
// Item.ValidateAll() if the designated constraints aren't met.
type ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemMultiError) AllErrors() []error { return m }

// ItemValidationError is the validation error returned by Item.Validate if the
// designated constraints aren't met.
type ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemValidationError) ErrorName() string { return "ItemValidationError" }

// Error satisfies the builtin error interface
func (e ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemValidationError{}

// Validate checks the field values on ListItemsSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListItemsSearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListItemsSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListItemsSearchRequestMultiError, or nil if none found.
func (m *ListItemsSearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListItemsSearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _ListItemsSearchRequest_Q_NotInLookup[m.GetQ()]; ok {
		err := ListItemsSearchRequestValidationError{
			field:  "Q",
			reason: "value must not be in list []",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PageSize

	// no validation rules for PageNo

	if len(errors) > 0 {
		return ListItemsSearchRequestMultiError(errors)
	}

	return nil
}

// ListItemsSearchRequestMultiError is an error wrapping multiple validation
// errors returned by ListItemsSearchRequest.ValidateAll() if the designated
// constraints aren't met.
type ListItemsSearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListItemsSearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListItemsSearchRequestMultiError) AllErrors() []error { return m }

// ListItemsSearchRequestValidationError is the validation error returned by
// ListItemsSearchRequest.Validate if the designated constraints aren't met.
type ListItemsSearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListItemsSearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListItemsSearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListItemsSearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListItemsSearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListItemsSearchRequestValidationError) ErrorName() string {
	return "ListItemsSearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListItemsSearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListItemsSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListItemsSearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListItemsSearchRequestValidationError{}

var _ListItemsSearchRequest_Q_NotInLookup = map[string]struct{}{
	"": {},
}
