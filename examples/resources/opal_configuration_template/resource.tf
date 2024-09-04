resource "opal_configuration_template" "my_configurationtemplate" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            custom_request_notification = "Check your email to register your account."
            name = "Prod AWS Template"
            request_configurations = [
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "154552e3-0c04-48b8-82b5-ed5f0cf07c66",
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
                        "419b360e-dc1d-466b-a4bf-8cfb72161b24",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    ]
            require_mfa_to_approve = false
            require_mfa_to_connect = false
            visibility = {
        visibility = "GLOBAL"
        visibility_group_ids = {
            "5b24ee84-a3fb-4f49-8bcc-ae3a741f65f5",
        }
    }
        }