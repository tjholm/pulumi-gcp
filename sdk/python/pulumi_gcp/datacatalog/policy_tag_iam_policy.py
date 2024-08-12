# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['PolicyTagIamPolicyArgs', 'PolicyTagIamPolicy']

@pulumi.input_type
class PolicyTagIamPolicyArgs:
    def __init__(__self__, *,
                 policy_data: pulumi.Input[str],
                 policy_tag: pulumi.Input[str]):
        """
        The set of arguments for constructing a PolicyTagIamPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        """
        pulumi.set(__self__, "policy_data", policy_data)
        pulumi.set(__self__, "policy_tag", policy_tag)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter(name="policyTag")
    def policy_tag(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "policy_tag")

    @policy_tag.setter
    def policy_tag(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_tag", value)


@pulumi.input_type
class _PolicyTagIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 policy_tag: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PolicyTagIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if policy_tag is not None:
            pulumi.set(__self__, "policy_tag", policy_tag)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter(name="policyTag")
    def policy_tag(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "policy_tag")

    @policy_tag.setter
    def policy_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_tag", value)


class PolicyTagIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 policy_tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:

        * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
        * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
        * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `datacatalog.PolicyTagIamPolicy`: Retrieves the IAM policy for the policytag

        > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.

        ## datacatalog.PolicyTagIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.datacatalog.PolicyTagIamPolicy("policy",
            policy_tag=basic_policy_tag["name"],
            policy_data=admin.policy_data)
        ```

        ## datacatalog.PolicyTagIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.PolicyTagIamBinding("binding",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## datacatalog.PolicyTagIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.PolicyTagIamMember("member",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## > **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
        ---

        # IAM policy for Data catalog PolicyTag
        Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:

        * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
        * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
        * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `datacatalog.PolicyTagIamPolicy`: Retrieves the IAM policy for the policytag

        > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.

        ## datacatalog.PolicyTagIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.datacatalog.PolicyTagIamPolicy("policy",
            policy_tag=basic_policy_tag["name"],
            policy_data=admin.policy_data)
        ```

        ## datacatalog.PolicyTagIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.PolicyTagIamBinding("binding",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## datacatalog.PolicyTagIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.PolicyTagIamMember("member",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * {{policy_tag}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor {{policy_tag}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyTagIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:

        * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
        * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
        * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `datacatalog.PolicyTagIamPolicy`: Retrieves the IAM policy for the policytag

        > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.

        ## datacatalog.PolicyTagIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.datacatalog.PolicyTagIamPolicy("policy",
            policy_tag=basic_policy_tag["name"],
            policy_data=admin.policy_data)
        ```

        ## datacatalog.PolicyTagIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.PolicyTagIamBinding("binding",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## datacatalog.PolicyTagIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.PolicyTagIamMember("member",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## > **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
        ---

        # IAM policy for Data catalog PolicyTag
        Three different resources help you manage your IAM policy for Data catalog PolicyTag. Each of these resources serves a different use case:

        * `datacatalog.PolicyTagIamPolicy`: Authoritative. Sets the IAM policy for the policytag and replaces any existing policy already attached.
        * `datacatalog.PolicyTagIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the policytag are preserved.
        * `datacatalog.PolicyTagIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the policytag are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `datacatalog.PolicyTagIamPolicy`: Retrieves the IAM policy for the policytag

        > **Note:** `datacatalog.PolicyTagIamPolicy` **cannot** be used in conjunction with `datacatalog.PolicyTagIamBinding` and `datacatalog.PolicyTagIamMember` or they will fight over what your policy should be.

        > **Note:** `datacatalog.PolicyTagIamBinding` resources **can be** used in conjunction with `datacatalog.PolicyTagIamMember` resources **only if** they do not grant privilege to the same role.

        ## datacatalog.PolicyTagIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.datacatalog.PolicyTagIamPolicy("policy",
            policy_tag=basic_policy_tag["name"],
            policy_data=admin.policy_data)
        ```

        ## datacatalog.PolicyTagIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.datacatalog.PolicyTagIamBinding("binding",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## datacatalog.PolicyTagIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.datacatalog.PolicyTagIamMember("member",
            policy_tag=basic_policy_tag["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * {{policy_tag}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor "{{policy_tag}} roles/viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy editor {{policy_tag}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param PolicyTagIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyTagIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 policy_tag: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyTagIamPolicyArgs.__new__(PolicyTagIamPolicyArgs)

            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            if policy_tag is None and not opts.urn:
                raise TypeError("Missing required property 'policy_tag'")
            __props__.__dict__["policy_tag"] = policy_tag
            __props__.__dict__["etag"] = None
        super(PolicyTagIamPolicy, __self__).__init__(
            'gcp:datacatalog/policyTagIamPolicy:PolicyTagIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            policy_tag: Optional[pulumi.Input[str]] = None) -> 'PolicyTagIamPolicy':
        """
        Get an existing PolicyTagIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] policy_tag: Used to find the parent resource to bind the IAM policy to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyTagIamPolicyState.__new__(_PolicyTagIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["policy_tag"] = policy_tag
        return PolicyTagIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter(name="policyTag")
    def policy_tag(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "policy_tag")

