# Release v1.26.0 (2025-05-08)
* Adds environments

# Release v1.25.0 (2025-04-09)
* Adds domains
* Adds manage_variables_by_yaml, manage_permissions_by_yaml to pipeline

# Release v1.24.0 (2025-02-04)
* [Breaking] - 'on' is removed from pipeline

# Release v1.23.0 (2025-01-21)
* [Breaking] - cron, startDate and delay in pipeline are moved to event
* [Breaking] - zone_id -> timezone in pipeline trigger condition


# Release v1.22.1 (2024-12-10)
* Adds cpu to pipeline

# Release v1.22.0 (2024-11-08)
* Adds project, pipeline, action to variable

# Release v1.21.0 (2024-10-14)
* Adds new properties to `Pipeline`
* concurrent_pipeline_runs
* description_required
* git_changeset_base
* filesystem_changeset_base

# Release v1.20.0 (2024-09-25)
* Adds new events to `Pipeline`

# Release v1.19.0 (2024-07-18)
* Adds `PauseOnRepeatedFailures` to `Pipeline`

# Release v1.18.0 (2024-07-08)
* Changes in Integration
* Remove scopes other than `WORKSPACE`, `PROJECT`
* Adds `permissions`
* Adds `allowed_pipelines`

# Release v1.17.0 (2024-04-19)
* Adds identifier to integration

# Release v1.16.0 (2024-04-12)
* Set Go version to 1.21
* Bump deps

# Release v1.15.0 (2024-02-21)
* Adds git configuration to pipeline

# Release v1.14.0 (2023-09-13)
* Adds new trigger conditions to pipeline

# Release v1.13.4 (2023-08-16)
* Adds configureable timeout to api client

# Release v1.13.3 (2023-08-16)
* Adds configureable timeout to api client

# Release v1.13.2 (2023-08-16)
* Adds configureable timeout to api client

# Release v1.13.1 (2023-07-19)
* Fixes OIDC support for Google

# Release v1.13.0 (2023-07-18)
* Adds OIDC support in integrations & login

# Release v1.12.0 (2023-07-17)
* Adds `STACK_HAWK` integration type

# Release v1.11.2 (2023-05-19)
* more logging fixes

# Release v1.11.1 (2023-05-19)
* fix logging system

# Release v1.11.0 (2023-05-16)
* Supported min GO version 1.19
* Upgraded deps
* Fix tests

# Release v1.10.3 (2023-04-28)
* Adds `GetMe` to `TokenService`

# Release v1.10.2 (2023-04-27)
* Adds token const

# Release v1.10.1 (2023-04-27)
* Adds token value on add

# Release v1.10.0 (2023-04-27)
* Adds token service

# Release v1.9.0 (2023-03-17)
* Adds SSO integration

# Release v1.8.2 (2023-01-11)
* Adds new Shopify integration type

# Release v1.8.1 (2023-01-10)
* Fix `permissions` in `Pipeline`

# Release v1.8.0 (2023-01-10)
* Adds `permissions` to `Pipeline`

# Release v1.7.0 (2022-12-08)
* Adds `without_repository` to `Project`

# Release v1.6.0 (2022-12-01)
* Remove `file_name` from `Variable`

# Release v1.5.0 (2022-10-15)
* Adds `allow_pull_requests, access, fetch_submodules, fetch_submodules_env_key` to `Project`
* Adds new scope `PRIVATE_IN_PROJECT` to `Integration`

# Release v1.4.0 (2022-07-20)
* Adds `update_default_branch_from_external` to `Project` integrated with external repository

# Release v1.3.0 (2022-06-29)
* Adds `status` to `GroupMember`

# Release v1.2.0 (2022-05-31)
* Adds `project_team_access_level` to `permission`

# Release v1.1.2 (2022-05-25)
* Adds `custom_repo_ssh_key_id` when creating custom project

# Release v1.1.1 (2022-04-19)
* Adds `GitHub` & `GitLab` token integration

# Release v1.1.0 (2022-04-14)
* Adds possibility to assign `Group` and `Member` to new projects by default using provided `Permission`