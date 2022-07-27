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
                 folder: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a IAMPolicy resource.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        pulumi.set(__self__, "folder", folder)
        pulumi.set(__self__, "policy_data", policy_data)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The `organizations.get_iam_policy` data source that represents
        the IAM policy that will be applied to the folder. The policy will be
        merged with any existing policy applied to the folder.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _IAMPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The `organizations.get_iam_policy` data source that represents
        the IAM policy that will be applied to the folder. The policy will be
        merged with any existing policy applied to the folder.
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
                 folder: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
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

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `folder`, role, and member e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `folder` and role, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `folder`.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder folder
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMPolicyArgs,
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

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `folder`, role, and member e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `folder` and role, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `folder`.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder folder
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
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
                 folder: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMPolicyArgs.__new__(IAMPolicyArgs)

            if folder is None and not opts.urn:
                raise TypeError("Missing required property 'folder'")
            __props__.__dict__["folder"] = folder
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["etag"] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:folder/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'IAMPolicy':
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMPolicyState.__new__(_IAMPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["folder"] = folder
        __props__.__dict__["policy_data"] = policy_data
        return IAMPolicy(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The `organizations.get_iam_policy` data source that represents
        the IAM policy that will be applied to the folder. The policy will be
        merged with any existing policy applied to the folder.
        """
        return pulumi.get(self, "policy_data")

