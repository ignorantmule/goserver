// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

// ListOptions is the query options to a standard REST list call.
type ListOptions struct {

	// LabelSelector is used to find matching REST resources.
	LabelSelector string `json:"labelSelector,omitempty" form:"labelSelector"`

	// FieldSelector restricts the list of returned objects by their fields. Defaults to everything.
	FieldSelector string `json:"fieldSelector,omitempty" form:"fieldSelector"`

	// Offset specify the number of records to skip before starting to return the records.
	Offset *int64 `json:"offset,omitempty" form:"offset"`

	// Limit specify the number of records to be retrieved.
	Limit *int64 `json:"limit,omitempty" form:"limit"`
}

// GetOptions is the standard query options to the standard REST get call.
type GetOptions struct{}

// DeleteOptions may be provided when deleting an API object.
type DeleteOptions struct {
	// +optional
	Unscoped bool `json:"unscoped"`
}

// CreateOptions may be provided when creating an API object.
type CreateOptions struct {
	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`
}

// UpdateOptions may be provided when updating an API object.
// All fields in UpdateOptions should also be present in PatchOptions.
type UpdateOptions struct {
	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`
}

// ListMeta describes metadata that synthetic resources must have, including lists and
// various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}
