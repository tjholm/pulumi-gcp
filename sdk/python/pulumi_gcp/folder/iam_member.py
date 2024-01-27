# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['IAMMemberArgs', 'IAMMember']

@pulumi.input_type
class IAMMemberArgs:
    def __init__(__self__, *,
                 folder: pulumi.Input[str],
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['IAMMemberConditionArgs']] = None):
        """
        The set of arguments for constructing a IAMMember resource.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
               `organizations/{{org_id}}/roles/{{role_id}}`.
        :param pulumi.Input['IAMMemberConditionArgs'] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        """
        pulumi.set(__self__, "folder", folder)
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Input[str]:
        """
        The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def member(self) -> pulumi.Input[str]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: pulumi.Input[str]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
        `organizations/{{org_id}}/roles/{{role_id}}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['IAMMemberConditionArgs']]:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['IAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _IAMMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['IAMMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMMember resources.
        :param pulumi.Input['IAMMemberConditionArgs'] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
               `organizations/{{org_id}}/roles/{{role_id}}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['IAMMemberConditionArgs']]:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['IAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the folder's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
        `organizations/{{org_id}}/roles/{{role_id}}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class IAMMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['IAMMemberConditionArgs']]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Four different resources help you manage your IAM policy for a folder. Each of these resources serves a different use case:

        * `folder.IAMPolicy`: Authoritative. Sets the IAM policy for the folder and replaces any existing policy already attached.
        * `folder.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the folder are preserved.
        * `folder.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the folder are preserved.
        * `folder.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.

        > **Note:** `folder.IAMPolicy` **cannot** be used in conjunction with `folder.IAMBinding`, `folder.IAMMember`, or `folder.IamAuditConfig` or they will fight over what your policy should be.

        > **Note:** `folder.IAMBinding` resources **can be** used in conjunction with `folder.IAMMember` resources **only if** they do not grant privilege to the same role.

        > **Note:** The underlying API method `projects.setIamPolicy` has constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
           IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning a 400 error code so please review these if you encounter errors with this resource.

        ## google\\_folder\\_iam\\_policy

        !> **Be careful!** You can accidentally lock yourself out of your folder
           using this resource. Deleting a `folder.IAMPolicy` removes access
           from anyone without permissions on its parent folder/organization. Proceed with caution.
           It's not recommended to use `folder.IAMPolicy` with your provider folder
           to avoid locking yourself out, and it should generally only be used with folders
           fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
           applying the change.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        folder = gcp.folder.IAMPolicy("folder",
            folder="folders/1234567",
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
            role="roles/compute.admin",
        )])
        folder = gcp.folder.IAMPolicy("folder",
            folder="folders/1234567",
            policy_data=admin.policy_data)
        ```

        ## google\\_folder\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMBinding("folder",
            folder="folders/1234567",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMBinding("folder",
            condition=gcp.folder.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            folder="folders/1234567",
            members=["user:jane@example.com"],
            role="roles/container.admin")
        ```

        ## google\\_folder\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMMember("folder",
            folder="folders/1234567",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMMember("folder",
            condition=gcp.folder.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            folder="folders/1234567",
            member="user:jane@example.com",
            role="roles/firebase.admin")
        ```

        ## google\\_folder\\_iam\\_audit\\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IamAuditConfig("folder",
            audit_log_configs=[
                gcp.folder.IamAuditConfigAuditLogConfigArgs(
                    log_type="ADMIN_READ",
                ),
                gcp.folder.IamAuditConfigAuditLogConfigArgs(
                    exempted_members=["user:joebloggs@hashicorp.com"],
                    log_type="DATA_READ",
                ),
            ],
            folder="folders/1234567",
            service="allServices")
        ```

        ## Import

        ### Importing Audit Configs An audit config can be imported into a `google_folder_iam_audit_config` resource using the resource's `folder_id` and the `service`, e.g* `"folder/{{folder_id}} foo.googleapis.com"` An `import` block (Terraform v1.5.0 and later) can be used to import audit configstf import {

         id = "folder/{{folder_id}} foo.googleapis.com"

         to = google_folder_iam_audit_config.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:folder/iAMMember:IAMMember default "folder/{{folder_id}} foo.googleapis.com"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['IAMMemberConditionArgs']] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
               `organizations/{{org_id}}/roles/{{role_id}}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Four different resources help you manage your IAM policy for a folder. Each of these resources serves a different use case:

        * `folder.IAMPolicy`: Authoritative. Sets the IAM policy for the folder and replaces any existing policy already attached.
        * `folder.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the folder are preserved.
        * `folder.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the folder are preserved.
        * `folder.IamAuditConfig`: Authoritative for a given service. Updates the IAM policy to enable audit logging for the given service.

        > **Note:** `folder.IAMPolicy` **cannot** be used in conjunction with `folder.IAMBinding`, `folder.IAMMember`, or `folder.IamAuditConfig` or they will fight over what your policy should be.

        > **Note:** `folder.IAMBinding` resources **can be** used in conjunction with `folder.IAMMember` resources **only if** they do not grant privilege to the same role.

        > **Note:** The underlying API method `projects.setIamPolicy` has constraints which are documented [here](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setIamPolicy). In addition to these constraints,
           IAM Conditions cannot be used with Basic Roles such as Owner. Violating these constraints will result in the API returning a 400 error code so please review these if you encounter errors with this resource.

        ## google\\_folder\\_iam\\_policy

        !> **Be careful!** You can accidentally lock yourself out of your folder
           using this resource. Deleting a `folder.IAMPolicy` removes access
           from anyone without permissions on its parent folder/organization. Proceed with caution.
           It's not recommended to use `folder.IAMPolicy` with your provider folder
           to avoid locking yourself out, and it should generally only be used with folders
           fully managed by this provider. If you do use this resource, it is recommended to **import** the policy before
           applying the change.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        folder = gcp.folder.IAMPolicy("folder",
            folder="folders/1234567",
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
            role="roles/compute.admin",
        )])
        folder = gcp.folder.IAMPolicy("folder",
            folder="folders/1234567",
            policy_data=admin.policy_data)
        ```

        ## google\\_folder\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMBinding("folder",
            folder="folders/1234567",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMBinding("folder",
            condition=gcp.folder.IAMBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            folder="folders/1234567",
            members=["user:jane@example.com"],
            role="roles/container.admin")
        ```

        ## google\\_folder\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMMember("folder",
            folder="folders/1234567",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IAMMember("folder",
            condition=gcp.folder.IAMMemberConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
                title="expires_after_2019_12_31",
            ),
            folder="folders/1234567",
            member="user:jane@example.com",
            role="roles/firebase.admin")
        ```

        ## google\\_folder\\_iam\\_audit\\_config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        folder = gcp.folder.IamAuditConfig("folder",
            audit_log_configs=[
                gcp.folder.IamAuditConfigAuditLogConfigArgs(
                    log_type="ADMIN_READ",
                ),
                gcp.folder.IamAuditConfigAuditLogConfigArgs(
                    exempted_members=["user:joebloggs@hashicorp.com"],
                    log_type="DATA_READ",
                ),
            ],
            folder="folders/1234567",
            service="allServices")
        ```

        ## Import

        ### Importing Audit Configs An audit config can be imported into a `google_folder_iam_audit_config` resource using the resource's `folder_id` and the `service`, e.g* `"folder/{{folder_id}} foo.googleapis.com"` An `import` block (Terraform v1.5.0 and later) can be used to import audit configstf import {

         id = "folder/{{folder_id}} foo.googleapis.com"

         to = google_folder_iam_audit_config.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:folder/iAMMember:IAMMember default "folder/{{folder_id}} foo.googleapis.com"
        ```

        :param str resource_name: The name of the resource.
        :param IAMMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['IAMMemberConditionArgs']]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMMemberArgs.__new__(IAMMemberArgs)

            __props__.__dict__["condition"] = condition
            if folder is None and not opts.urn:
                raise TypeError("Missing required property 'folder'")
            __props__.__dict__["folder"] = folder
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(IAMMember, __self__).__init__(
            'gcp:folder/iAMMember:IAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['IAMMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'IAMMember':
        """
        Get an existing IAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['IAMMemberConditionArgs']] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
               `organizations/{{org_id}}/roles/{{role_id}}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMMemberState.__new__(_IAMMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["folder"] = folder
        __props__.__dict__["member"] = member
        __props__.__dict__["role"] = role
        return IAMMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.IAMMemberCondition']]:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the folder's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `folder.IAMBinding` can be used per role. Note that custom roles must be of the format
        `organizations/{{org_id}}/roles/{{role_id}}`.
        """
        return pulumi.get(self, "role")

