// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type GcpBigQueryTable struct {
	DatasetID types.String `tfsdk:"dataset_id"`
	ProjectID types.String `tfsdk:"project_id"`
	TableID   types.String `tfsdk:"table_id"`
}
