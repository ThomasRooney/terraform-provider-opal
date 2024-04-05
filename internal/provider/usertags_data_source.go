// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UserTagsDataSource{}
var _ datasource.DataSourceWithConfigure = &UserTagsDataSource{}

func NewUserTagsDataSource() datasource.DataSource {
	return &UserTagsDataSource{}
}

// UserTagsDataSource is the data source implementation.
type UserTagsDataSource struct {
	client *sdk.OpalAPI
}

// UserTagsDataSourceModel describes the data model.
type UserTagsDataSourceModel struct {
	Tags   []tfTypes.Tag `tfsdk:"tags"`
	UserID types.String  `tfsdk:"user_id"`
}

// Metadata returns the data source type name.
func (r *UserTagsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_tags"
}

// Schema defines the schema for the data source.
func (r *UserTagsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "UserTags DataSource",

		Attributes: map[string]schema.Attribute{
			"tags": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: `The date the tag was created.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the tag.`,
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: `The key of the tag.`,
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: `The date the tag was last updated.`,
						},
						"user_creator_id": schema.StringAttribute{
							Computed:    true,
							Description: `The ID of the user that created the tag.`,
						},
						"value": schema.StringAttribute{
							Computed:    true,
							Description: `The value of the tag.`,
						},
					},
				},
			},
			"user_id": schema.StringAttribute{
				Required:    true,
				Description: `The ID of the user whose tags to return.`,
			},
		},
	}
}

func (r *UserTagsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.OpalAPI)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.OpalAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *UserTagsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UserTagsDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	userID := data.UserID.ValueString()
	request := operations.GetUserTagsRequest{
		UserID: userID,
	}
	res, err := r.client.Users.GetUserTags(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.TagsList == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedTagsList(res.TagsList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
