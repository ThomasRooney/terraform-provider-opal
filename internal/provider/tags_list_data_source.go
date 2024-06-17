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
var _ datasource.DataSource = &TagsListDataSource{}
var _ datasource.DataSourceWithConfigure = &TagsListDataSource{}

func NewTagsListDataSource() datasource.DataSource {
	return &TagsListDataSource{}
}

// TagsListDataSource is the data source implementation.
type TagsListDataSource struct {
	client *sdk.OpalAPI
}

// TagsListDataSourceModel describes the data model.
type TagsListDataSourceModel struct {
	Cursor   types.String  `tfsdk:"cursor"`
	Next     types.String  `tfsdk:"next"`
	PageSize types.Int64   `tfsdk:"page_size"`
	Previous types.String  `tfsdk:"previous"`
	Results  []tfTypes.Tag `tfsdk:"results"`
}

// Metadata returns the data source type name.
func (r *TagsListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tags_list"
}

// Schema defines the schema for the data source.
func (r *TagsListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "TagsList DataSource",

		Attributes: map[string]schema.Attribute{
			"cursor": schema.StringAttribute{
				Optional:    true,
				Description: `The pagination cursor value.`,
			},
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor with which to continue pagination if additional result pages exist.`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `Number of results to return per page. Default is 200.`,
			},
			"previous": schema.StringAttribute{
				Computed:    true,
				Description: `The cursor used to obtain the current result page.`,
			},
			"results": schema.ListNestedAttribute{
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
		},
	}
}

func (r *TagsListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *TagsListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *TagsListDataSourceModel
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

	cursor := new(string)
	if !data.Cursor.IsUnknown() && !data.Cursor.IsNull() {
		*cursor = data.Cursor.ValueString()
	} else {
		cursor = nil
	}
	pageSize := new(int64)
	if !data.PageSize.IsUnknown() && !data.PageSize.IsNull() {
		*pageSize = data.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	request := operations.GetTagsRequest{
		Cursor:   cursor,
		PageSize: pageSize,
	}
	res, err := r.client.Tags.GetTags(ctx, request)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.PaginatedTagsList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedPaginatedTagsList(res.PaginatedTagsList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
