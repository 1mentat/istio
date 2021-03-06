// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/admin/v2alpha/config_dump.proto

package envoy_admin_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on ConfigDump with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ConfigDump) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetConfigs() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ConfigDumpValidationError{
						field:  fmt.Sprintf("Configs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ConfigDumpValidationError is the validation error returned by
// ConfigDump.Validate if the designated constraints aren't met.
type ConfigDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigDumpValidationError) ErrorName() string { return "ConfigDumpValidationError" }

// Error satisfies the builtin error interface
func (e ConfigDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigDumpValidationError{}

// Validate checks the field values on BootstrapConfigDump with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *BootstrapConfigDump) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetBootstrap()

		if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapConfigDumpValidationError{
					field:  "Bootstrap",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapConfigDumpValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// BootstrapConfigDumpValidationError is the validation error returned by
// BootstrapConfigDump.Validate if the designated constraints aren't met.
type BootstrapConfigDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapConfigDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapConfigDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapConfigDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapConfigDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapConfigDumpValidationError) ErrorName() string {
	return "BootstrapConfigDumpValidationError"
}

// Error satisfies the builtin error interface
func (e BootstrapConfigDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrapConfigDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapConfigDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapConfigDumpValidationError{}

// Validate checks the field values on ListenersConfigDump with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListenersConfigDump) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	for idx, item := range m.GetStaticListeners() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenersConfigDumpValidationError{
						field:  fmt.Sprintf("StaticListeners[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetDynamicActiveListeners() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenersConfigDumpValidationError{
						field:  fmt.Sprintf("DynamicActiveListeners[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetDynamicWarmingListeners() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenersConfigDumpValidationError{
						field:  fmt.Sprintf("DynamicWarmingListeners[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetDynamicDrainingListeners() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListenersConfigDumpValidationError{
						field:  fmt.Sprintf("DynamicDrainingListeners[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ListenersConfigDumpValidationError is the validation error returned by
// ListenersConfigDump.Validate if the designated constraints aren't met.
type ListenersConfigDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenersConfigDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenersConfigDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenersConfigDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenersConfigDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenersConfigDumpValidationError) ErrorName() string {
	return "ListenersConfigDumpValidationError"
}

// Error satisfies the builtin error interface
func (e ListenersConfigDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenersConfigDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenersConfigDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenersConfigDumpValidationError{}

// Validate checks the field values on ClustersConfigDump with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClustersConfigDump) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	for idx, item := range m.GetStaticClusters() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClustersConfigDumpValidationError{
						field:  fmt.Sprintf("StaticClusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetDynamicActiveClusters() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClustersConfigDumpValidationError{
						field:  fmt.Sprintf("DynamicActiveClusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetDynamicWarmingClusters() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClustersConfigDumpValidationError{
						field:  fmt.Sprintf("DynamicWarmingClusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ClustersConfigDumpValidationError is the validation error returned by
// ClustersConfigDump.Validate if the designated constraints aren't met.
type ClustersConfigDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClustersConfigDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClustersConfigDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClustersConfigDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClustersConfigDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClustersConfigDumpValidationError) ErrorName() string {
	return "ClustersConfigDumpValidationError"
}

// Error satisfies the builtin error interface
func (e ClustersConfigDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClustersConfigDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClustersConfigDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClustersConfigDumpValidationError{}

// Validate checks the field values on RoutesConfigDump with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RoutesConfigDump) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStaticRouteConfigs() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RoutesConfigDumpValidationError{
						field:  fmt.Sprintf("StaticRouteConfigs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetDynamicRouteConfigs() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(&tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RoutesConfigDumpValidationError{
						field:  fmt.Sprintf("DynamicRouteConfigs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// RoutesConfigDumpValidationError is the validation error returned by
// RoutesConfigDump.Validate if the designated constraints aren't met.
type RoutesConfigDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoutesConfigDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoutesConfigDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoutesConfigDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoutesConfigDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoutesConfigDumpValidationError) ErrorName() string { return "RoutesConfigDumpValidationError" }

// Error satisfies the builtin error interface
func (e RoutesConfigDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoutesConfigDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoutesConfigDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoutesConfigDumpValidationError{}

// Validate checks the field values on ListenersConfigDump_StaticListener with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListenersConfigDump_StaticListener) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetListener()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenersConfigDump_StaticListenerValidationError{
					field:  "Listener",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenersConfigDump_StaticListenerValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListenersConfigDump_StaticListenerValidationError is the validation error
// returned by ListenersConfigDump_StaticListener.Validate if the designated
// constraints aren't met.
type ListenersConfigDump_StaticListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenersConfigDump_StaticListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenersConfigDump_StaticListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenersConfigDump_StaticListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenersConfigDump_StaticListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenersConfigDump_StaticListenerValidationError) ErrorName() string {
	return "ListenersConfigDump_StaticListenerValidationError"
}

// Error satisfies the builtin error interface
func (e ListenersConfigDump_StaticListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenersConfigDump_StaticListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenersConfigDump_StaticListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenersConfigDump_StaticListenerValidationError{}

// Validate checks the field values on ListenersConfigDump_DynamicListener with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ListenersConfigDump_DynamicListener) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	{
		tmp := m.GetListener()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenersConfigDump_DynamicListenerValidationError{
					field:  "Listener",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListenersConfigDump_DynamicListenerValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListenersConfigDump_DynamicListenerValidationError is the validation error
// returned by ListenersConfigDump_DynamicListener.Validate if the designated
// constraints aren't met.
type ListenersConfigDump_DynamicListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenersConfigDump_DynamicListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenersConfigDump_DynamicListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenersConfigDump_DynamicListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenersConfigDump_DynamicListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenersConfigDump_DynamicListenerValidationError) ErrorName() string {
	return "ListenersConfigDump_DynamicListenerValidationError"
}

// Error satisfies the builtin error interface
func (e ListenersConfigDump_DynamicListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenersConfigDump_DynamicListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenersConfigDump_DynamicListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenersConfigDump_DynamicListenerValidationError{}

// Validate checks the field values on ClustersConfigDump_StaticCluster with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ClustersConfigDump_StaticCluster) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetCluster()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClustersConfigDump_StaticClusterValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClustersConfigDump_StaticClusterValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClustersConfigDump_StaticClusterValidationError is the validation error
// returned by ClustersConfigDump_StaticCluster.Validate if the designated
// constraints aren't met.
type ClustersConfigDump_StaticClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClustersConfigDump_StaticClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClustersConfigDump_StaticClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClustersConfigDump_StaticClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClustersConfigDump_StaticClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClustersConfigDump_StaticClusterValidationError) ErrorName() string {
	return "ClustersConfigDump_StaticClusterValidationError"
}

// Error satisfies the builtin error interface
func (e ClustersConfigDump_StaticClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClustersConfigDump_StaticCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClustersConfigDump_StaticClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClustersConfigDump_StaticClusterValidationError{}

// Validate checks the field values on ClustersConfigDump_DynamicCluster with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ClustersConfigDump_DynamicCluster) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	{
		tmp := m.GetCluster()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClustersConfigDump_DynamicClusterValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClustersConfigDump_DynamicClusterValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClustersConfigDump_DynamicClusterValidationError is the validation error
// returned by ClustersConfigDump_DynamicCluster.Validate if the designated
// constraints aren't met.
type ClustersConfigDump_DynamicClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClustersConfigDump_DynamicClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClustersConfigDump_DynamicClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClustersConfigDump_DynamicClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClustersConfigDump_DynamicClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClustersConfigDump_DynamicClusterValidationError) ErrorName() string {
	return "ClustersConfigDump_DynamicClusterValidationError"
}

// Error satisfies the builtin error interface
func (e ClustersConfigDump_DynamicClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClustersConfigDump_DynamicCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClustersConfigDump_DynamicClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClustersConfigDump_DynamicClusterValidationError{}

// Validate checks the field values on RoutesConfigDump_StaticRouteConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *RoutesConfigDump_StaticRouteConfig) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetRouteConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RoutesConfigDump_StaticRouteConfigValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RoutesConfigDump_StaticRouteConfigValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RoutesConfigDump_StaticRouteConfigValidationError is the validation error
// returned by RoutesConfigDump_StaticRouteConfig.Validate if the designated
// constraints aren't met.
type RoutesConfigDump_StaticRouteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoutesConfigDump_StaticRouteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoutesConfigDump_StaticRouteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoutesConfigDump_StaticRouteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoutesConfigDump_StaticRouteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoutesConfigDump_StaticRouteConfigValidationError) ErrorName() string {
	return "RoutesConfigDump_StaticRouteConfigValidationError"
}

// Error satisfies the builtin error interface
func (e RoutesConfigDump_StaticRouteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoutesConfigDump_StaticRouteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoutesConfigDump_StaticRouteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoutesConfigDump_StaticRouteConfigValidationError{}

// Validate checks the field values on RoutesConfigDump_DynamicRouteConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *RoutesConfigDump_DynamicRouteConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VersionInfo

	{
		tmp := m.GetRouteConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RoutesConfigDump_DynamicRouteConfigValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLastUpdated()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RoutesConfigDump_DynamicRouteConfigValidationError{
					field:  "LastUpdated",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RoutesConfigDump_DynamicRouteConfigValidationError is the validation error
// returned by RoutesConfigDump_DynamicRouteConfig.Validate if the designated
// constraints aren't met.
type RoutesConfigDump_DynamicRouteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoutesConfigDump_DynamicRouteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoutesConfigDump_DynamicRouteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoutesConfigDump_DynamicRouteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoutesConfigDump_DynamicRouteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoutesConfigDump_DynamicRouteConfigValidationError) ErrorName() string {
	return "RoutesConfigDump_DynamicRouteConfigValidationError"
}

// Error satisfies the builtin error interface
func (e RoutesConfigDump_DynamicRouteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoutesConfigDump_DynamicRouteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoutesConfigDump_DynamicRouteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoutesConfigDump_DynamicRouteConfigValidationError{}
