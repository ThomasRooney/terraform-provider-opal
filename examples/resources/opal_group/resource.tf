resource "opal_group" "my_group" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            app_id = "b5a5ca27-0ea3-4d86-9199-2126d57d1fbd"
            custom_request_notification = "Check your email to register your account."
            description = "This group represents Active Directory group \"Payments Production Admin\". We use this AD group to facilitate staging deployments and qualifying new releases."
            group_type = "OPAL_GROUP"
            message_channel_ids = {
        "1c0009dd-65b4-4e10-8428-93326c8d3b65",
    }
            name = "API Group"
            on_call_schedule_ids = {
        "3f385d17-ce27-4d58-b256-147d92ea2469",
    }
            request_configurations = {
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "24e146eb-05cd-406c-96b8-673554554108",
                }
                role_remote_ids = {
                    "...",
                }
            }
            max_duration = 120
            priority = 1
            recommended_duration = 120
            request_template_id = "06851574-e50d-40ca-8c78-f72ae6ab4304"
            require_mfa_to_request = false
            require_support_ticket = false
            reviewer_stages = [
                {
                    operator = "AND"
                    owner_ids = {
                        "5bf18869-ae72-46c0-8c01-8ec506c2a39b",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    }
            require_mfa_to_approve = false
            visibility = "GLOBAL"
        }