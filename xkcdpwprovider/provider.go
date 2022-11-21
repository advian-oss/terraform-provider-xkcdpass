package xkcdpwprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &xkcdpwProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &xkcdpwProvider{}
}

// xkcdpwProvider is the provider implementation.
type xkcdpwProvider struct{}

// Metadata returns the provider type name.
func (p *xkcdpwProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "xkcdpass"
}

// GetSchema defines the provider-level schema for configuration data.
func (p *xkcdpwProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{}, nil
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *xkcdpwProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *xkcdpwProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *xkcdpwProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewPwResource,
	}
}
