package xkcdpwprovider

import (
	"context"
	"terraform-provider-xkcdpass/xkcdpwprovider/planmodifiers"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/martinhoefling/goxkcdpwgen/xkcdpwgen"
)

var _ resource.Resource = (*pwResource)(nil)

// NewPwResource a Helper to return pwResource
func NewPwResource() resource.Resource {
	return &pwResource{}
}

type pwResource struct{}

func (r *pwResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_generate"
}

func (r *pwResource) GetSchema(context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "generated password in result.\n" +
			"\n" +
			"This resource can be used in conjunction with resources that have the `create_before_destroy` " +
			"lifecycle flag set, to avoid conflicts with unique names during the brief period where both the old " +
			"and new resources exist concurrently.",
		Attributes: map[string]tfsdk.Attribute{
			"length": {
				Description: "The length (in words) of the passphrase. Defaults to 4",
				Type:        types.Int64Type,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					planmodifiers.DefaultValue(types.Int64Value(4)),
					planmodifiers.RequiresReplace(),
				},
			},
			"separator": {
				Description: "The character to separate words in the password. Defaults to \"-\"",
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					planmodifiers.DefaultValue(types.StringValue("-")),
					planmodifiers.RequiresReplace(),
				},
			},
			"capitalize": {
				Description: "Capitalize words, defaults to true",
				Type:        types.BoolType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					planmodifiers.DefaultValue(types.BoolValue(true)),
					planmodifiers.RequiresReplace(),
				},
			},
			"result": {
				Description: "The password.",
				Type:        types.StringType,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}, nil
}

func (r *pwResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan pwModelV0

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	length := plan.Length.ValueInt64()
	separator := plan.Separator.ValueString()
	capitalize := plan.Capitalize.ValueBool()

	pn := pwModelV0{
		Capitalize: types.BoolValue(capitalize),
		Length:     types.Int64Value(length),
		Separator:  types.StringValue(separator),
	}

	g := xkcdpwgen.NewGenerator()
	g.SetNumWords(int(length))
	g.SetCapitalize(capitalize)
	g.SetDelimiter(separator)
	password := g.GeneratePasswordString()

	pn.Result = types.StringValue(password)

	diags = resp.State.Set(ctx, pn)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read does not need to perform any operations as the state in ReadResourceResponse is already populated.
func (r *pwResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update ensures the plan value is copied to the state to complete the update.
func (r *pwResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model pwModelV0

	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

// Delete does not need to explicitly call resp.State.RemoveResource() as this is automatically handled by the
// [framework](https://github.com/hashicorp/terraform-plugin-framework/pull/301).
func (r *pwResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type pwModelV0 struct {
	Capitalize types.Bool   `tfsdk:"capitalize"`
	Result     types.String `tfsdk:"result"`
	Length     types.Int64  `tfsdk:"length"`
	Separator  types.String `tfsdk:"separator"`
}
