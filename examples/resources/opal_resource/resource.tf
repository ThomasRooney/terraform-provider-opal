resource "opal_resource" "my_resource" {
    admin_owner_id = "7c86c85d-0651-43e2-a748-d69d658418e8"
            app_id = "b5a5ca27-0ea3-4d86-9199-2126d57d1fbd"
            custom_request_notification = "Check your email to register your account."
            description = "Engineering team Okta role."
            name = "my-mongo-db"
            request_configurations = {
        {
            allow_requests = true
            auto_approval = false
            condition = {
                group_ids = {
                    "a4c7ec57-b816-4172-8cd3-687b3ad9a192",
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
                        "bdfcb3e7-464a-4e94-bc89-fbd03a92242f",
                    }
                    require_admin_approval = false
                    require_manager_approval = false
                },
            ]
        },
    }
            require_mfa_to_approve = false
            require_mfa_to_connect = false
            resource_type = "AWS_IAM_ROLE"
            visibility = "GLOBAL"
            visibility_group_ids = {
        "041a0d76-1a93-4988-b1ee-194539e33fdc",
    }
        }