# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IAMPolicyArgs', 'IAMPolicy']

@pulumi.input_type
class IAMPolicyArgs:
    def __init__(__self__, *,
                 org_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a IAMPolicy resource.
        :param pulumi.Input[str] org_id: The organization id of the target organization.
        :param pulumi.Input[str] policy_data: The `organizations_get_iam_policy` data source that represents
               the IAM policy that will be applied to the organization. The policy will be
               merged with any existing policy applied to the organization.
        """
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        The organization id of the target organization.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The `organizations_get_iam_policy` data source that represents
        the IAM policy that will be applied to the organization. The policy will be
        merged with any existing policy applied to the organization.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _IAMPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the organization's IAM policy.
        :param pulumi.Input[str] org_id: The organization id of the target organization.
        :param pulumi.Input[str] policy_data: The `organizations_get_iam_policy` data source that represents
               the IAM policy that will be applied to the organization. The policy will be
               merged with any existing policy applied to the organization.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the organization's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The organization id of the target organization.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The `organizations_get_iam_policy` data source that represents
        the IAM policy that will be applied to the organization. The policy will be
        merged with any existing policy applied to the organization.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)


class IAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Four different resources help you manage your IAM policy for a organization. Each of these resources serves a different use case:

        * `organizations.IAMPolicy`: Authoritative. Sets the IAM policy for the organization and replaces any existing policy already attached.
        * `organizations.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the organization are preserved.
        * `organizations.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the organization are preserved.
        * `organizations.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.

        > **Note:** `organizations.IAMPolicy` **cannot** be used in conjunction with `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig` or they will fight over what your policy should be.

        > **Note:** `organizations.IAMBinding` resources **can be** used in conjunction with `organizations.IAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_organization\\_iam\\_policy

        !> **Warning:** New organizations have several default policies which will,
           without extreme caution, be **overwritten** by use of this resource.
           The safest alternative is to use multiple `organizations.IAMBinding`
           resources. This resource makes it easy to remove your own access to
           an organization, which will require a call to Google Support to have
           fixed, and can take multiple days to resolve.

           In general, this resource should only be used with organizations
           fully managed by this provider.I f you do use this resource,
           the best way to be sure that you are not making dangerous changes is to start
           by **importing** your existing policy, and examining the diff very closely.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        organization = gcp.organizations.IAMPolicy("organization",
            org_id="1234567890",
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            role="roles/editor",
        )])
        organization = gcp.organizations.IAMPolicy("organization",
            org_id="1234567890",
            policy_data=admin.policy_data)
        ```

        ## google\\_organization\\_iam\\_binding

        > **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your organization.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMBinding("organization",
            members=["user:jane@example.com"],
            org_id="1234567890",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMBinding("organization",
            condition=gcp.organizations.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            org_id="1234567890",
            role="roles/editor")
        ```

        ## google\\_organization\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMMember("organization",
            member="user:jane@example.com",
            org_id="1234567890",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMMember("organization",
            condition=gcp.organizations.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            member="user:jane@example.com",
            org_id="1234567890",
            role="roles/editor")
        ```

        ## google\\_organization\\_iam\\_audit\\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IamAuditConfig("organization",
            audit_log_configs=[
                gcp.organizations.IamAuditConfigAuditLogConfigArgs(
                    log_type="ADMIN_READ",
                ),
                gcp.organizations.IamAuditConfigAuditLogConfigArgs(
                    exempted_members=["user:joebloggs@hashicorp.com"],
                    log_type="DATA_READ",
                ),
            ],
            org_id="1234567890",
            service="allServices")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `org_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization "your-orgid roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `org_id` and role, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization "your-org-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `org_id`.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization your-org-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization "your-organization-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_organization_iam_binding.my_organization "your-org-id roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org_id: The organization id of the target organization.
        :param pulumi.Input[str] policy_data: The `organizations_get_iam_policy` data source that represents
               the IAM policy that will be applied to the organization. The policy will be
               merged with any existing policy applied to the organization.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Four different resources help you manage your IAM policy for a organization. Each of these resources serves a different use case:

        * `organizations.IAMPolicy`: Authoritative. Sets the IAM policy for the organization and replaces any existing policy already attached.
        * `organizations.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the organization are preserved.
        * `organizations.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the organization are preserved.
        * `organizations.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.

        > **Note:** `organizations.IAMPolicy` **cannot** be used in conjunction with `organizations.IAMBinding`, `organizations.IAMMember`, or `organizations.IamAuditConfig` or they will fight over what your policy should be.

        > **Note:** `organizations.IAMBinding` resources **can be** used in conjunction with `organizations.IAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_organization\\_iam\\_policy

        !> **Warning:** New organizations have several default policies which will,
           without extreme caution, be **overwritten** by use of this resource.
           The safest alternative is to use multiple `organizations.IAMBinding`
           resources. This resource makes it easy to remove your own access to
           an organization, which will require a call to Google Support to have
           fixed, and can take multiple days to resolve.

           In general, this resource should only be used with organizations
           fully managed by this provider.I f you do use this resource,
           the best way to be sure that you are not making dangerous changes is to start
           by **importing** your existing policy, and examining the diff very closely.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        organization = gcp.organizations.IAMPolicy("organization",
            org_id="1234567890",
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            role="roles/editor",
        )])
        organization = gcp.organizations.IAMPolicy("organization",
            org_id="1234567890",
            policy_data=admin.policy_data)
        ```

        ## google\\_organization\\_iam\\_binding

        > **Note:** If `role` is set to `roles/owner` and you don't specify a user or service account you have access to in `members`, you can lock yourself out of your organization.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMBinding("organization",
            members=["user:jane@example.com"],
            org_id="1234567890",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMBinding("organization",
            condition=gcp.organizations.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            org_id="1234567890",
            role="roles/editor")
        ```

        ## google\\_organization\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMMember("organization",
            member="user:jane@example.com",
            org_id="1234567890",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IAMMember("organization",
            condition=gcp.organizations.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            member="user:jane@example.com",
            org_id="1234567890",
            role="roles/editor")
        ```

        ## google\\_organization\\_iam\\_audit\\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        organization = gcp.organizations.IamAuditConfig("organization",
            audit_log_configs=[
                gcp.organizations.IamAuditConfigAuditLogConfigArgs(
                    log_type="ADMIN_READ",
                ),
                gcp.organizations.IamAuditConfigAuditLogConfigArgs(
                    exempted_members=["user:joebloggs@hashicorp.com"],
                    log_type="DATA_READ",
                ),
            ],
            org_id="1234567890",
            service="allServices")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `org_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization "your-orgid roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `org_id` and role, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization "your-org-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `org_id`.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization your-org-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy my_organization "your-organization-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:organizations/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_organization_iam_binding.my_organization "your-org-id roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param IAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMPolicyArgs.__new__(IAMPolicyArgs)

            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["etag"] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:organizations/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'IAMPolicy':
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the organization's IAM policy.
        :param pulumi.Input[str] org_id: The organization id of the target organization.
        :param pulumi.Input[str] policy_data: The `organizations_get_iam_policy` data source that represents
               the IAM policy that will be applied to the organization. The policy will be
               merged with any existing policy applied to the organization.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMPolicyState.__new__(_IAMPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["policy_data"] = policy_data
        return IAMPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the organization's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The organization id of the target organization.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The `organizations_get_iam_policy` data source that represents
        the IAM policy that will be applied to the organization. The policy will be
        merged with any existing policy applied to the organization.
        """
        return pulumi.get(self, "policy_data")

