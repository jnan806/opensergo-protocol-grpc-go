// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: fault_tolerance/v1/fault_tolerance.proto

package v1

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

// Validate checks the field values on FaultToleranceRule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FaultToleranceRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FaultToleranceRule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FaultToleranceRuleMultiError, or nil if none found.
func (m *FaultToleranceRule) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultToleranceRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTargets() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("Targets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("Targets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultToleranceRuleValidationError{
					field:  fmt.Sprintf("Targets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetStrategies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("Strategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("Strategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultToleranceRuleValidationError{
					field:  fmt.Sprintf("Strategies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetAction()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FaultToleranceRuleValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FaultToleranceRuleValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultToleranceRuleValidationError{
				field:  "Action",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FaultToleranceRuleMultiError(errors)
	}

	return nil
}

// FaultToleranceRuleMultiError is an error wrapping multiple validation errors
// returned by FaultToleranceRule.ValidateAll() if the designated constraints
// aren't met.
type FaultToleranceRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultToleranceRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultToleranceRuleMultiError) AllErrors() []error { return m }

// FaultToleranceRuleValidationError is the validation error returned by
// FaultToleranceRule.Validate if the designated constraints aren't met.
type FaultToleranceRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultToleranceRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultToleranceRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultToleranceRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultToleranceRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultToleranceRuleValidationError) ErrorName() string {
	return "FaultToleranceRuleValidationError"
}

// Error satisfies the builtin error interface
func (e FaultToleranceRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultToleranceRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultToleranceRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultToleranceRuleValidationError{}

// Validate checks the field values on RateLimitStrategy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RateLimitStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitStrategyMultiError, or nil if none found.
func (m *RateLimitStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for MetricType

	// no validation rules for LimitMode

	if m.GetThreshold() < 0 {
		err := RateLimitStrategyValidationError{
			field:  "Threshold",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStatDuration() <= 0 {
		err := RateLimitStrategyValidationError{
			field:  "StatDuration",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for StatDurationTimeUnit

	if len(errors) > 0 {
		return RateLimitStrategyMultiError(errors)
	}

	return nil
}

// RateLimitStrategyMultiError is an error wrapping multiple validation errors
// returned by RateLimitStrategy.ValidateAll() if the designated constraints
// aren't met.
type RateLimitStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitStrategyMultiError) AllErrors() []error { return m }

// RateLimitStrategyValidationError is the validation error returned by
// RateLimitStrategy.Validate if the designated constraints aren't met.
type RateLimitStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitStrategyValidationError) ErrorName() string {
	return "RateLimitStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitStrategyValidationError{}

// Validate checks the field values on ThrottlingStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ThrottlingStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ThrottlingStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ThrottlingStrategyMultiError, or nil if none found.
func (m *ThrottlingStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *ThrottlingStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for MinIntervalMillisOfRequests

	// no validation rules for QueueTimeoutMillis

	if len(errors) > 0 {
		return ThrottlingStrategyMultiError(errors)
	}

	return nil
}

// ThrottlingStrategyMultiError is an error wrapping multiple validation errors
// returned by ThrottlingStrategy.ValidateAll() if the designated constraints
// aren't met.
type ThrottlingStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThrottlingStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThrottlingStrategyMultiError) AllErrors() []error { return m }

// ThrottlingStrategyValidationError is the validation error returned by
// ThrottlingStrategy.Validate if the designated constraints aren't met.
type ThrottlingStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThrottlingStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThrottlingStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThrottlingStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThrottlingStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThrottlingStrategyValidationError) ErrorName() string {
	return "ThrottlingStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e ThrottlingStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThrottlingStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThrottlingStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThrottlingStrategyValidationError{}

// Validate checks the field values on ConcurrencyLimitStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConcurrencyLimitStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConcurrencyLimitStrategy with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConcurrencyLimitStrategyMultiError, or nil if none found.
func (m *ConcurrencyLimitStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *ConcurrencyLimitStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for LimitMode

	// no validation rules for MaxConcurrency

	if len(errors) > 0 {
		return ConcurrencyLimitStrategyMultiError(errors)
	}

	return nil
}

// ConcurrencyLimitStrategyMultiError is an error wrapping multiple validation
// errors returned by ConcurrencyLimitStrategy.ValidateAll() if the designated
// constraints aren't met.
type ConcurrencyLimitStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConcurrencyLimitStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConcurrencyLimitStrategyMultiError) AllErrors() []error { return m }

// ConcurrencyLimitStrategyValidationError is the validation error returned by
// ConcurrencyLimitStrategy.Validate if the designated constraints aren't met.
type ConcurrencyLimitStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcurrencyLimitStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcurrencyLimitStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcurrencyLimitStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcurrencyLimitStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcurrencyLimitStrategyValidationError) ErrorName() string {
	return "ConcurrencyLimitStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e ConcurrencyLimitStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcurrencyLimitStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcurrencyLimitStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcurrencyLimitStrategyValidationError{}

// Validate checks the field values on CircuitBreakerStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CircuitBreakerStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CircuitBreakerStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CircuitBreakerStrategyMultiError, or nil if none found.
func (m *CircuitBreakerStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) > 1024 {
		err := CircuitBreakerStrategyValidationError{
			field:  "Name",
			reason: "value length must be at most 1024 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Strategy

	if val := m.GetTriggerRatio(); val < 0 || val > 1 {
		err := CircuitBreakerStrategyValidationError{
			field:  "TriggerRatio",
			reason: "value must be inside range [0, 1]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStatDuration() <= 0 {
		err := CircuitBreakerStrategyValidationError{
			field:  "StatDuration",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for StatDurationTimeUnit

	if m.GetRecoveryTimeout() <= 0 {
		err := CircuitBreakerStrategyValidationError{
			field:  "RecoveryTimeout",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for RecoveryTimeoutTimeUnit

	if m.GetMinRequestAmount() <= 0 {
		err := CircuitBreakerStrategyValidationError{
			field:  "MinRequestAmount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSlowCondition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "SlowCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "SlowCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSlowCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakerStrategyValidationError{
				field:  "SlowCondition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetErrorCondition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "ErrorCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "ErrorCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetErrorCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakerStrategyValidationError{
				field:  "ErrorCondition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CircuitBreakerStrategyMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategyMultiError is an error wrapping multiple validation
// errors returned by CircuitBreakerStrategy.ValidateAll() if the designated
// constraints aren't met.
type CircuitBreakerStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategyMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategyValidationError is the validation error returned by
// CircuitBreakerStrategy.Validate if the designated constraints aren't met.
type CircuitBreakerStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakerStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakerStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakerStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategyValidationError) ErrorName() string {
	return "CircuitBreakerStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategyValidationError{}

// Validate checks the field values on
// FaultToleranceRule_FaultToleranceRuleTargetRef with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FaultToleranceRule_FaultToleranceRuleTargetRef) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// FaultToleranceRule_FaultToleranceRuleTargetRef with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// FaultToleranceRule_FaultToleranceRuleTargetRefMultiError, or nil if none found.
func (m *FaultToleranceRule_FaultToleranceRuleTargetRef) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultToleranceRule_FaultToleranceRuleTargetRef) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TargetResourceName

	if len(errors) > 0 {
		return FaultToleranceRule_FaultToleranceRuleTargetRefMultiError(errors)
	}

	return nil
}

// FaultToleranceRule_FaultToleranceRuleTargetRefMultiError is an error
// wrapping multiple validation errors returned by
// FaultToleranceRule_FaultToleranceRuleTargetRef.ValidateAll() if the
// designated constraints aren't met.
type FaultToleranceRule_FaultToleranceRuleTargetRefMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultToleranceRule_FaultToleranceRuleTargetRefMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultToleranceRule_FaultToleranceRuleTargetRefMultiError) AllErrors() []error { return m }

// FaultToleranceRule_FaultToleranceRuleTargetRefValidationError is the
// validation error returned by
// FaultToleranceRule_FaultToleranceRuleTargetRef.Validate if the designated
// constraints aren't met.
type FaultToleranceRule_FaultToleranceRuleTargetRefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultToleranceRule_FaultToleranceRuleTargetRefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultToleranceRule_FaultToleranceRuleTargetRefValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e FaultToleranceRule_FaultToleranceRuleTargetRefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultToleranceRule_FaultToleranceRuleTargetRefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultToleranceRule_FaultToleranceRuleTargetRefValidationError) ErrorName() string {
	return "FaultToleranceRule_FaultToleranceRuleTargetRefValidationError"
}

// Error satisfies the builtin error interface
func (e FaultToleranceRule_FaultToleranceRuleTargetRefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultToleranceRule_FaultToleranceRuleTargetRef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultToleranceRule_FaultToleranceRuleTargetRefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultToleranceRule_FaultToleranceRuleTargetRefValidationError{}

// Validate checks the field values on
// FaultToleranceRule_FaultToleranceStrategyRef with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FaultToleranceRule_FaultToleranceStrategyRef) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// FaultToleranceRule_FaultToleranceStrategyRef with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// FaultToleranceRule_FaultToleranceStrategyRefMultiError, or nil if none found.
func (m *FaultToleranceRule_FaultToleranceStrategyRef) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultToleranceRule_FaultToleranceStrategyRef) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Kind

	if len(errors) > 0 {
		return FaultToleranceRule_FaultToleranceStrategyRefMultiError(errors)
	}

	return nil
}

// FaultToleranceRule_FaultToleranceStrategyRefMultiError is an error wrapping
// multiple validation errors returned by
// FaultToleranceRule_FaultToleranceStrategyRef.ValidateAll() if the
// designated constraints aren't met.
type FaultToleranceRule_FaultToleranceStrategyRefMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultToleranceRule_FaultToleranceStrategyRefMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultToleranceRule_FaultToleranceStrategyRefMultiError) AllErrors() []error { return m }

// FaultToleranceRule_FaultToleranceStrategyRefValidationError is the
// validation error returned by
// FaultToleranceRule_FaultToleranceStrategyRef.Validate if the designated
// constraints aren't met.
type FaultToleranceRule_FaultToleranceStrategyRefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultToleranceRule_FaultToleranceStrategyRefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultToleranceRule_FaultToleranceStrategyRefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultToleranceRule_FaultToleranceStrategyRefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultToleranceRule_FaultToleranceStrategyRefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultToleranceRule_FaultToleranceStrategyRefValidationError) ErrorName() string {
	return "FaultToleranceRule_FaultToleranceStrategyRefValidationError"
}

// Error satisfies the builtin error interface
func (e FaultToleranceRule_FaultToleranceStrategyRefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultToleranceRule_FaultToleranceStrategyRef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultToleranceRule_FaultToleranceStrategyRefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultToleranceRule_FaultToleranceStrategyRefValidationError{}

// Validate checks the field values on
// FaultToleranceRule_FaultToleranceActionRef with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FaultToleranceRule_FaultToleranceActionRef) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// FaultToleranceRule_FaultToleranceActionRef with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// FaultToleranceRule_FaultToleranceActionRefMultiError, or nil if none found.
func (m *FaultToleranceRule_FaultToleranceActionRef) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultToleranceRule_FaultToleranceActionRef) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Kind

	if len(errors) > 0 {
		return FaultToleranceRule_FaultToleranceActionRefMultiError(errors)
	}

	return nil
}

// FaultToleranceRule_FaultToleranceActionRefMultiError is an error wrapping
// multiple validation errors returned by
// FaultToleranceRule_FaultToleranceActionRef.ValidateAll() if the designated
// constraints aren't met.
type FaultToleranceRule_FaultToleranceActionRefMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultToleranceRule_FaultToleranceActionRefMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultToleranceRule_FaultToleranceActionRefMultiError) AllErrors() []error { return m }

// FaultToleranceRule_FaultToleranceActionRefValidationError is the validation
// error returned by FaultToleranceRule_FaultToleranceActionRef.Validate if
// the designated constraints aren't met.
type FaultToleranceRule_FaultToleranceActionRefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultToleranceRule_FaultToleranceActionRefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultToleranceRule_FaultToleranceActionRefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultToleranceRule_FaultToleranceActionRefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultToleranceRule_FaultToleranceActionRefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultToleranceRule_FaultToleranceActionRefValidationError) ErrorName() string {
	return "FaultToleranceRule_FaultToleranceActionRefValidationError"
}

// Error satisfies the builtin error interface
func (e FaultToleranceRule_FaultToleranceActionRefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultToleranceRule_FaultToleranceActionRef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultToleranceRule_FaultToleranceActionRefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultToleranceRule_FaultToleranceActionRefValidationError{}

// Validate checks the field values on
// CircuitBreakerStrategy_CircuitBreakerSlowCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CircuitBreakerStrategy_CircuitBreakerSlowCondition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CircuitBreakerStrategy_CircuitBreakerSlowCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError, or nil if
// none found.
func (m *CircuitBreakerStrategy_CircuitBreakerSlowCondition) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy_CircuitBreakerSlowCondition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MaxAllowedRtMillis

	if len(errors) > 0 {
		return CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError is an error
// wrapping multiple validation errors returned by
// CircuitBreakerStrategy_CircuitBreakerSlowCondition.ValidateAll() if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError is the
// validation error returned by
// CircuitBreakerStrategy_CircuitBreakerSlowCondition.Validate if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) ErrorName() string {
	return "CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy_CircuitBreakerSlowCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError{}

// Validate checks the field values on
// CircuitBreakerStrategy_CircuitBreakerErrorCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CircuitBreakerStrategy_CircuitBreakerErrorCondition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CircuitBreakerStrategy_CircuitBreakerErrorCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError, or nil if
// none found.
func (m *CircuitBreakerStrategy_CircuitBreakerErrorCondition) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy_CircuitBreakerErrorCondition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError is an error
// wrapping multiple validation errors returned by
// CircuitBreakerStrategy_CircuitBreakerErrorCondition.ValidateAll() if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError is the
// validation error returned by
// CircuitBreakerStrategy_CircuitBreakerErrorCondition.Validate if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) ErrorName() string {
	return "CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy_CircuitBreakerErrorCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError{}