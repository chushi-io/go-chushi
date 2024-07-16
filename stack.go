// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfe

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// Stacks describes all the stacks-related methods that the HCP Terraform API supports.
// NOTE WELL: This is a beta feature and is subject to change until noted otherwise in the
// release notes.
type Stacks interface {
	// List returns a list of stacks, optionally filtered by project.
	List(ctx context.Context, organization string, options *StackListOptions) (*StackList, error)

	// Read returns a stack by its ID.
	Read(ctx context.Context, stackID string) (*Stack, error)

	// Create creates a new stack.
	Create(ctx context.Context, options StackCreateOptions) (*Stack, error)

	// Update updates a stack.
	Update(ctx context.Context, stackID string, options StackUpdateOptions) (*Stack, error)

	// Delete deletes a stack.
	Delete(ctx context.Context, stackID string) error

	// UpdateConfiguration updates the configuration of a stack, triggering stack preparation.
	UpdateConfiguration(ctx context.Context, stackID string) (*Stack, error)
}

// stacks implements Stacks.
type stacks struct {
	client *Client
}

var _ Stacks = &stacks{}

// StackSortColumn represents a string that can be used to sort items when using
// the List method.
type StackSortColumn string

const (
	// StackSortByName sorts by the name attribute.
	StackSortByName StackSortColumn = "name"

	// StackSortByUpdatedAt sorts by the updated-at attribute.
	StackSortByUpdatedAt StackSortColumn = "updated-at"

	// StackSortByNameDesc sorts by the name attribute in descending order.
	StackSortByNameDesc StackSortColumn = "-name"

	// StackSortByUpdatedAtDesc sorts by the updated-at attribute in descending order.
	StackSortByUpdatedAtDesc StackSortColumn = "-updated-at"
)

// StackList represents a list of stacks.
type StackList struct {
	*Pagination
	Items []*Stack
}

// StackVCSRepo represents the version control system repository for a stack.
type StackVCSRepo struct {
	Identifier        string `jsonapi:"attr,identifier"`
	Branch            string `jsonapi:"attr,branch"`
	GHAInstallationID string `jsonapi:"attr,github-app-installation-id"`
	OAuthTokenID      string `jsonapi:"attr,oauth-token-id"`
}

// Stack represents a stack.
type Stack struct {
	ID              string        `jsonapi:"primary,stacks"`
	Name            string        `jsonapi:"attr,name"`
	Description     string        `jsonapi:"attr,description"`
	DeploymentNames []string      `jsonapi:"attr,deployment-names"`
	VCSRepo         *StackVCSRepo `jsonapi:"attr,vcs-repo"`
	ErrorsCount     int           `jsonapi:"attr,errors-count"`
	WarningsCount   int           `jsonapi:"attr,warnings-count"`
	CreatedAt       time.Time     `jsonapi:"attr,created-at,iso8601"`
	UpdatedAt       time.Time     `jsonapi:"attr,updated-at,iso8601"`

	// Relationships
	Project                  *Project            `jsonapi:"relation,project"`
	LatestStackConfiguration *StackConfiguration `jsonapi:"relation,latest-stack-configuration"`
}

type StackConfigurationStatusTimestamps struct {
	QueuedAt     *time.Time `jsonapi:"attr,queued-at,omitempty,rfc3339"`
	CompletedAt  *time.Time `jsonapi:"attr,completed-at,omitempty,rfc3339"`
	PreparingAt  *time.Time `jsonapi:"attr,preparing-at,omitempty,rfc3339"`
	EnqueueingAt *time.Time `jsonapi:"attr,enqueueing-at,omitempty,rfc3339"`
	CanceledAt   *time.Time `jsonapi:"attr,canceled-at,omitempty,rfc3339"`
	ErroredAt    *time.Time `jsonapi:"attr,errored-at,omitempty,rfc3339"`
}

type StackComponent struct {
	Name       string `json:"name"`
	Correlator string `json:"correlator"`
	Expanded   bool   `json:"expanded"`
}

type StackConfiguration struct {
	// Attributes
	ID                   string                              `jsonapi:"primary,stack-configurations"`
	Status               string                              `jsonapi:"attr,status"`
	StatusTimestamps     *StackConfigurationStatusTimestamps `jsonapi:"attr,status-timestamps"`
	SequenceNumber       int                                 `jsonapi:"attr,sequence-number"`
	DeploymentNames      []string                            `jsonapi:"attr,deployment-names"`
	ConvergedDeployments []string                            `jsonapi:"attr,converged-deployments"`
	Components           []*StackComponent                   `jsonapi:"attr,components"`
	ErrorMessage         *string                             `jsonapi:"attr,error-message"`
	EventStreamURL       string                              `jsonapi:"attr,event-stream-url"`
}

type StackDeployment struct {
	// Attributes
	ID            string    `jsonapi:"primary,stack-deployments"`
	Name          string    `jsonapi:"attr,name"`
	Status        string    `jsonapi:"attr,status"`
	DeployedAt    time.Time `jsonapi:"attr,deployed-at,iso8601"`
	ErrorsCount   int       `jsonapi:"attr,errors-count"`
	WarningsCount int       `jsonapi:"attr,warnings-count"`
	PausedCount   int       `jsonapi:"attr,paused-count"`

	// Relationships
	CurrentStackState *StackState `jsonapi:"relation,current-stack-state"`
}

type StackState struct {
	// Attributes
	ID string `jsonapi:"primary,stack-states"`
}

// StackListOptions represents the options for listing stacks.
type StackListOptions struct {
	ListOptions
	ProjectID    string          `url:"filter[project[id]],omitempty"`
	Sort         StackSortColumn `url:"sort,omitempty"`
	SearchByName string          `url:"search[name],omitempty"`
}

// StackCreateOptions represents the options for creating a stack. The project
// relation is required.
type StackCreateOptions struct {
	Type        string        `jsonapi:"primary,stacks"`
	Name        string        `jsonapi:"attr,name"`
	Description *string       `jsonapi:"attr,description,omitempty"`
	VCSRepo     *StackVCSRepo `jsonapi:"attr,vcs-repo"`
	Project     *Project      `jsonapi:"relation,project"`
}

// StackUpdateOptions represents the options for updating a stack.
type StackUpdateOptions struct {
	Name        *string       `jsonapi:"attr,name,omitempty"`
	Description *string       `jsonapi:"attr,description,omitempty"`
	VCSRepo     *StackVCSRepo `jsonapi:"attr,vcs-repo,omitempty"`
}

func (s stacks) UpdateConfiguration(ctx context.Context, stackID string) (*Stack, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("stacks/%s/actions/update-configuration", url.PathEscape(stackID)), nil)
	if err != nil {
		return nil, err
	}

	stack := &Stack{}
	err = req.Do(ctx, stack)
	if err != nil {
		return nil, err
	}

	return stack, nil
}

func (s stacks) List(ctx context.Context, organization string, options *StackListOptions) (*StackList, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", fmt.Sprintf("organizations/%s/stacks", organization), options)
	if err != nil {
		return nil, err
	}

	sl := &StackList{}
	err = req.Do(ctx, sl)
	if err != nil {
		return nil, err
	}

	return sl, nil
}

func (s stacks) Read(ctx context.Context, stackID string) (*Stack, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("stacks/%s", url.PathEscape(stackID)), nil)
	if err != nil {
		return nil, err
	}

	stack := &Stack{}
	err = req.Do(ctx, stack)
	if err != nil {
		return nil, err
	}

	return stack, nil
}

func (s stacks) Create(ctx context.Context, options StackCreateOptions) (*Stack, error) {
	if err := options.valid(); err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", "stacks", &options)
	if err != nil {
		return nil, err
	}

	stack := &Stack{}
	err = req.Do(ctx, stack)
	if err != nil {
		return nil, err
	}

	return stack, nil
}

func (s stacks) Update(ctx context.Context, stackID string, options StackUpdateOptions) (*Stack, error) {
	req, err := s.client.NewRequest("PATCH", fmt.Sprintf("stacks/%s", url.PathEscape(stackID)), &options)
	if err != nil {
		return nil, err
	}

	stack := &Stack{}
	err = req.Do(ctx, stack)
	if err != nil {
		return nil, err
	}

	return stack, nil
}

func (s stacks) Delete(ctx context.Context, stackID string) error {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("stacks/%s/delete", url.PathEscape(stackID)), nil)
	if err != nil {
		return err
	}

	return req.Do(ctx, nil)
}

func (s *StackListOptions) valid() error {
	return nil
}

func (s StackCreateOptions) valid() error {
	if s.Name == "" {
		return ErrRequiredName
	}

	if s.Project.ID == "" {
		return ErrRequiredProject
	}

	return s.VCSRepo.valid()
}

func (s StackVCSRepo) valid() error {
	if s.Identifier == "" {
		return ErrRequiredVCSRepo
	}

	return nil
}
