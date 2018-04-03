// Code generated by go-swagger; DO NOT EDIT.

package job_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new job manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for job manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DescribeJobs describes jobs with filter
*/
func (a *Client) DescribeJobs(params *DescribeJobsParams) (*DescribeJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeJobsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeJobs",
		Method:             "GET",
		PathPattern:        "/v1/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeJobsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeJobsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
