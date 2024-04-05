// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/opalsecurity/terraform-provider-opal/internal/provider/types"
	"github.com/opalsecurity/terraform-provider-opal/internal/sdk/models/shared"
)

func (r *GroupListDataSourceModel) RefreshFromSharedPaginatedGroupsList(resp *shared.PaginatedGroupsList) {
	if resp != nil {
		if len(r.Results) > len(resp.Results) {
			r.Results = r.Results[:len(resp.Results)]
		}
		for resultsCount, resultsItem := range resp.Results {
			var results1 tfTypes.Group
			results1.AdminOwnerID = types.StringPointerValue(resultsItem.AdminOwnerID)
			results1.AppID = types.StringPointerValue(resultsItem.AppID)
			results1.AutoApproval = types.BoolPointerValue(resultsItem.AutoApproval)
			results1.Description = types.StringPointerValue(resultsItem.Description)
			results1.GroupBindingID = types.StringPointerValue(resultsItem.GroupBindingID)
			if resultsItem.GroupType != nil {
				results1.GroupType = types.StringValue(string(*resultsItem.GroupType))
			} else {
				results1.GroupType = types.StringNull()
			}
			results1.ID = types.StringPointerValue(resultsItem.ID)
			results1.IsRequestable = types.BoolPointerValue(resultsItem.IsRequestable)
			results1.MaxDuration = types.Int64PointerValue(resultsItem.MaxDuration)
			results1.Name = types.StringPointerValue(resultsItem.Name)
			results1.RecommendedDuration = types.Int64PointerValue(resultsItem.RecommendedDuration)
			results1.RemoteID = types.StringPointerValue(resultsItem.RemoteID)
			if resultsItem.RemoteInfo == nil {
				results1.RemoteInfo = nil
			} else {
				results1.RemoteInfo = &tfTypes.GroupRemoteInfo{}
				if resultsItem.RemoteInfo.ActiveDirectoryGroup == nil {
					results1.RemoteInfo.ActiveDirectoryGroup = nil
				} else {
					results1.RemoteInfo.ActiveDirectoryGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.ActiveDirectoryGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.ActiveDirectoryGroup.GroupID)
				}
				if resultsItem.RemoteInfo.AzureAdMicrosoft365Group == nil {
					results1.RemoteInfo.AzureAdMicrosoft365Group = nil
				} else {
					results1.RemoteInfo.AzureAdMicrosoft365Group = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.AzureAdMicrosoft365Group.GroupID = types.StringValue(resultsItem.RemoteInfo.AzureAdMicrosoft365Group.GroupID)
				}
				if resultsItem.RemoteInfo.AzureAdSecurityGroup == nil {
					results1.RemoteInfo.AzureAdSecurityGroup = nil
				} else {
					results1.RemoteInfo.AzureAdSecurityGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.AzureAdSecurityGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.AzureAdSecurityGroup.GroupID)
				}
				if resultsItem.RemoteInfo.DuoGroup == nil {
					results1.RemoteInfo.DuoGroup = nil
				} else {
					results1.RemoteInfo.DuoGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.DuoGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.DuoGroup.GroupID)
				}
				if resultsItem.RemoteInfo.GithubTeam == nil {
					results1.RemoteInfo.GithubTeam = nil
				} else {
					results1.RemoteInfo.GithubTeam = &tfTypes.GithubTeam{}
					results1.RemoteInfo.GithubTeam.TeamSlug = types.StringValue(resultsItem.RemoteInfo.GithubTeam.TeamSlug)
				}
				if resultsItem.RemoteInfo.GitlabGroup == nil {
					results1.RemoteInfo.GitlabGroup = nil
				} else {
					results1.RemoteInfo.GitlabGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.GitlabGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.GitlabGroup.GroupID)
				}
				if resultsItem.RemoteInfo.GoogleGroup == nil {
					results1.RemoteInfo.GoogleGroup = nil
				} else {
					results1.RemoteInfo.GoogleGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.GoogleGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.GoogleGroup.GroupID)
				}
				if resultsItem.RemoteInfo.LdapGroup == nil {
					results1.RemoteInfo.LdapGroup = nil
				} else {
					results1.RemoteInfo.LdapGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.LdapGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.LdapGroup.GroupID)
				}
				if resultsItem.RemoteInfo.OktaGroup == nil {
					results1.RemoteInfo.OktaGroup = nil
				} else {
					results1.RemoteInfo.OktaGroup = &tfTypes.ActiveDirectoryGroup{}
					results1.RemoteInfo.OktaGroup.GroupID = types.StringValue(resultsItem.RemoteInfo.OktaGroup.GroupID)
				}
			}
			results1.RemoteName = types.StringPointerValue(resultsItem.RemoteName)
			for requestConfigurationListDataCount, requestConfigurationListDataItem := range resultsItem.RequestConfigurationListData {
				var requestConfigurationListData1 tfTypes.RequestConfiguration
				requestConfigurationListData1.AllowRequests = types.BoolValue(requestConfigurationListDataItem.AllowRequests)
				requestConfigurationListData1.AutoApproval = types.BoolValue(requestConfigurationListDataItem.AutoApproval)
				if requestConfigurationListDataItem.Condition == nil {
					requestConfigurationListData1.Condition = nil
				} else {
					requestConfigurationListData1.Condition = &tfTypes.Condition{}
					requestConfigurationListData1.Condition.GroupIds = []types.String{}
					for _, v := range requestConfigurationListDataItem.Condition.GroupIds {
						requestConfigurationListData1.Condition.GroupIds = append(requestConfigurationListData1.Condition.GroupIds, types.StringValue(v))
					}
					requestConfigurationListData1.Condition.RoleRemoteIds = []types.String{}
					for _, v := range requestConfigurationListDataItem.Condition.RoleRemoteIds {
						requestConfigurationListData1.Condition.RoleRemoteIds = append(requestConfigurationListData1.Condition.RoleRemoteIds, types.StringValue(v))
					}
				}
				requestConfigurationListData1.MaxDuration = types.Int64PointerValue(requestConfigurationListDataItem.MaxDuration)
				requestConfigurationListData1.Priority = types.Int64Value(requestConfigurationListDataItem.Priority)
				requestConfigurationListData1.RecommendedDuration = types.Int64PointerValue(requestConfigurationListDataItem.RecommendedDuration)
				requestConfigurationListData1.RequestTemplateID = types.StringPointerValue(requestConfigurationListDataItem.RequestTemplateID)
				requestConfigurationListData1.RequireMfaToRequest = types.BoolValue(requestConfigurationListDataItem.RequireMfaToRequest)
				requestConfigurationListData1.RequireSupportTicket = types.BoolValue(requestConfigurationListDataItem.RequireSupportTicket)
				for reviewerStagesCount, reviewerStagesItem := range requestConfigurationListDataItem.ReviewerStages {
					var reviewerStages1 tfTypes.ReviewerStage
					if reviewerStagesItem.Operator != nil {
						reviewerStages1.Operator = types.StringValue(string(*reviewerStagesItem.Operator))
					} else {
						reviewerStages1.Operator = types.StringNull()
					}
					reviewerStages1.OwnerIds = []types.String{}
					for _, v := range reviewerStagesItem.OwnerIds {
						reviewerStages1.OwnerIds = append(reviewerStages1.OwnerIds, types.StringValue(v))
					}
					reviewerStages1.RequireManagerApproval = types.BoolValue(reviewerStagesItem.RequireManagerApproval)
					if reviewerStagesCount+1 > len(requestConfigurationListData1.ReviewerStages) {
						requestConfigurationListData1.ReviewerStages = append(requestConfigurationListData1.ReviewerStages, reviewerStages1)
					} else {
						requestConfigurationListData1.ReviewerStages[reviewerStagesCount].Operator = reviewerStages1.Operator
						requestConfigurationListData1.ReviewerStages[reviewerStagesCount].OwnerIds = reviewerStages1.OwnerIds
						requestConfigurationListData1.ReviewerStages[reviewerStagesCount].RequireManagerApproval = reviewerStages1.RequireManagerApproval
					}
				}
				if requestConfigurationListDataCount+1 > len(results1.RequestConfigurationListData) {
					results1.RequestConfigurationListData = append(results1.RequestConfigurationListData, requestConfigurationListData1)
				} else {
					results1.RequestConfigurationListData[requestConfigurationListDataCount].AllowRequests = requestConfigurationListData1.AllowRequests
					results1.RequestConfigurationListData[requestConfigurationListDataCount].AutoApproval = requestConfigurationListData1.AutoApproval
					results1.RequestConfigurationListData[requestConfigurationListDataCount].Condition = requestConfigurationListData1.Condition
					results1.RequestConfigurationListData[requestConfigurationListDataCount].MaxDuration = requestConfigurationListData1.MaxDuration
					results1.RequestConfigurationListData[requestConfigurationListDataCount].Priority = requestConfigurationListData1.Priority
					results1.RequestConfigurationListData[requestConfigurationListDataCount].RecommendedDuration = requestConfigurationListData1.RecommendedDuration
					results1.RequestConfigurationListData[requestConfigurationListDataCount].RequestTemplateID = requestConfigurationListData1.RequestTemplateID
					results1.RequestConfigurationListData[requestConfigurationListDataCount].RequireMfaToRequest = requestConfigurationListData1.RequireMfaToRequest
					results1.RequestConfigurationListData[requestConfigurationListDataCount].RequireSupportTicket = requestConfigurationListData1.RequireSupportTicket
					results1.RequestConfigurationListData[requestConfigurationListDataCount].ReviewerStages = requestConfigurationListData1.ReviewerStages
				}
			}
			for requestConfigurationsCount, requestConfigurationsItem := range resultsItem.RequestConfigurations {
				var requestConfigurations1 tfTypes.RequestConfiguration
				requestConfigurations1.AllowRequests = types.BoolValue(requestConfigurationsItem.AllowRequests)
				requestConfigurations1.AutoApproval = types.BoolValue(requestConfigurationsItem.AutoApproval)
				if requestConfigurationsItem.Condition == nil {
					requestConfigurations1.Condition = nil
				} else {
					requestConfigurations1.Condition = &tfTypes.Condition{}
					requestConfigurations1.Condition.GroupIds = []types.String{}
					for _, v := range requestConfigurationsItem.Condition.GroupIds {
						requestConfigurations1.Condition.GroupIds = append(requestConfigurations1.Condition.GroupIds, types.StringValue(v))
					}
					requestConfigurations1.Condition.RoleRemoteIds = []types.String{}
					for _, v := range requestConfigurationsItem.Condition.RoleRemoteIds {
						requestConfigurations1.Condition.RoleRemoteIds = append(requestConfigurations1.Condition.RoleRemoteIds, types.StringValue(v))
					}
				}
				requestConfigurations1.MaxDuration = types.Int64PointerValue(requestConfigurationsItem.MaxDuration)
				requestConfigurations1.Priority = types.Int64Value(requestConfigurationsItem.Priority)
				requestConfigurations1.RecommendedDuration = types.Int64PointerValue(requestConfigurationsItem.RecommendedDuration)
				requestConfigurations1.RequestTemplateID = types.StringPointerValue(requestConfigurationsItem.RequestTemplateID)
				requestConfigurations1.RequireMfaToRequest = types.BoolValue(requestConfigurationsItem.RequireMfaToRequest)
				requestConfigurations1.RequireSupportTicket = types.BoolValue(requestConfigurationsItem.RequireSupportTicket)
				for reviewerStagesCount1, reviewerStagesItem1 := range requestConfigurationsItem.ReviewerStages {
					var reviewerStages3 tfTypes.ReviewerStage
					if reviewerStagesItem1.Operator != nil {
						reviewerStages3.Operator = types.StringValue(string(*reviewerStagesItem1.Operator))
					} else {
						reviewerStages3.Operator = types.StringNull()
					}
					reviewerStages3.OwnerIds = []types.String{}
					for _, v := range reviewerStagesItem1.OwnerIds {
						reviewerStages3.OwnerIds = append(reviewerStages3.OwnerIds, types.StringValue(v))
					}
					reviewerStages3.RequireManagerApproval = types.BoolValue(reviewerStagesItem1.RequireManagerApproval)
					if reviewerStagesCount1+1 > len(requestConfigurations1.ReviewerStages) {
						requestConfigurations1.ReviewerStages = append(requestConfigurations1.ReviewerStages, reviewerStages3)
					} else {
						requestConfigurations1.ReviewerStages[reviewerStagesCount1].Operator = reviewerStages3.Operator
						requestConfigurations1.ReviewerStages[reviewerStagesCount1].OwnerIds = reviewerStages3.OwnerIds
						requestConfigurations1.ReviewerStages[reviewerStagesCount1].RequireManagerApproval = reviewerStages3.RequireManagerApproval
					}
				}
				if requestConfigurationsCount+1 > len(results1.RequestConfigurations) {
					results1.RequestConfigurations = append(results1.RequestConfigurations, requestConfigurations1)
				} else {
					results1.RequestConfigurations[requestConfigurationsCount].AllowRequests = requestConfigurations1.AllowRequests
					results1.RequestConfigurations[requestConfigurationsCount].AutoApproval = requestConfigurations1.AutoApproval
					results1.RequestConfigurations[requestConfigurationsCount].Condition = requestConfigurations1.Condition
					results1.RequestConfigurations[requestConfigurationsCount].MaxDuration = requestConfigurations1.MaxDuration
					results1.RequestConfigurations[requestConfigurationsCount].Priority = requestConfigurations1.Priority
					results1.RequestConfigurations[requestConfigurationsCount].RecommendedDuration = requestConfigurations1.RecommendedDuration
					results1.RequestConfigurations[requestConfigurationsCount].RequestTemplateID = requestConfigurations1.RequestTemplateID
					results1.RequestConfigurations[requestConfigurationsCount].RequireMfaToRequest = requestConfigurations1.RequireMfaToRequest
					results1.RequestConfigurations[requestConfigurationsCount].RequireSupportTicket = requestConfigurations1.RequireSupportTicket
					results1.RequestConfigurations[requestConfigurationsCount].ReviewerStages = requestConfigurations1.ReviewerStages
				}
			}
			results1.RequestTemplateID = types.StringPointerValue(resultsItem.RequestTemplateID)
			results1.RequireMfaToApprove = types.BoolPointerValue(resultsItem.RequireMfaToApprove)
			results1.RequireMfaToRequest = types.BoolPointerValue(resultsItem.RequireMfaToRequest)
			results1.RequireSupportTicket = types.BoolPointerValue(resultsItem.RequireSupportTicket)
			if resultsCount+1 > len(r.Results) {
				r.Results = append(r.Results, results1)
			} else {
				r.Results[resultsCount].AdminOwnerID = results1.AdminOwnerID
				r.Results[resultsCount].AppID = results1.AppID
				r.Results[resultsCount].AutoApproval = results1.AutoApproval
				r.Results[resultsCount].Description = results1.Description
				r.Results[resultsCount].GroupBindingID = results1.GroupBindingID
				r.Results[resultsCount].GroupType = results1.GroupType
				r.Results[resultsCount].ID = results1.ID
				r.Results[resultsCount].IsRequestable = results1.IsRequestable
				r.Results[resultsCount].MaxDuration = results1.MaxDuration
				r.Results[resultsCount].Name = results1.Name
				r.Results[resultsCount].RecommendedDuration = results1.RecommendedDuration
				r.Results[resultsCount].RemoteID = results1.RemoteID
				r.Results[resultsCount].RemoteInfo = results1.RemoteInfo
				r.Results[resultsCount].RemoteName = results1.RemoteName
				r.Results[resultsCount].RequestConfigurationListData = results1.RequestConfigurationListData
				r.Results[resultsCount].RequestConfigurations = results1.RequestConfigurations
				r.Results[resultsCount].RequestTemplateID = results1.RequestTemplateID
				r.Results[resultsCount].RequireMfaToApprove = results1.RequireMfaToApprove
				r.Results[resultsCount].RequireMfaToRequest = results1.RequireMfaToRequest
				r.Results[resultsCount].RequireSupportTicket = results1.RequireSupportTicket
			}
		}
	}
}
