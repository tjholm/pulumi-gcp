# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TagValueIamBindingArgs', 'TagValueIamBinding']

@pulumi.input_type
class TagValueIamBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 tag_value: pulumi.Input[str],
                 condition: Optional[pulumi.Input['TagValueIamBindingConditionArgs']] = None):
        """
        The set of arguments for constructing a TagValueIamBinding resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] tag_value: Used to find the parent resource to bind the IAM policy to
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "tag_value", tag_value)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "tag_value")

    @tag_value.setter
    def tag_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "tag_value", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['TagValueIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['TagValueIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _TagValueIamBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['TagValueIamBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tag_value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TagValueIamBinding resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] tag_value: Used to find the parent resource to bind the IAM policy to
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if tag_value is not None:
            pulumi.set(__self__, "tag_value", tag_value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['TagValueIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['TagValueIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

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
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "tag_value")

    @tag_value.setter
    def tag_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_value", value)


class TagValueIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['TagValueIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tag_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Tags TagValue. Each of these resources serves a different use case:

        * `tags.TagValueIamPolicy`: Authoritative. Sets the IAM policy for the tagvalue and replaces any existing policy already attached.
        * `tags.TagValueIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tagvalue are preserved.
        * `tags.TagValueIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tagvalue are preserved.

        > **Note:** `tags.TagValueIamPolicy` **cannot** be used in conjunction with `tags.TagValueIamBinding` and `tags.TagValueIamMember` or they will fight over what your policy should be.

        > **Note:** `tags.TagValueIamBinding` resources **can be** used in conjunction with `tags.TagValueIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_tags\_tag\_value\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.tags.TagValueIamPolicy("policy",
            tag_value=google_tags_tag_value["value"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_tags\_tag\_value\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.tags.TagValueIamBinding("binding",
            tag_value=google_tags_tag_value["value"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_tags\_tag\_value\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.tags.TagValueIamMember("member",
            tag_value=google_tags_tag_value["value"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* tagValues/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Tags tagvalue IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:tags/tagValueIamBinding:TagValueIamBinding editor "tagValues/{{tag_value}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:tags/tagValueIamBinding:TagValueIamBinding editor "tagValues/{{tag_value}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:tags/tagValueIamBinding:TagValueIamBinding editor tagValues/{{tag_value}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] tag_value: Used to find the parent resource to bind the IAM policy to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TagValueIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Tags TagValue. Each of these resources serves a different use case:

        * `tags.TagValueIamPolicy`: Authoritative. Sets the IAM policy for the tagvalue and replaces any existing policy already attached.
        * `tags.TagValueIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tagvalue are preserved.
        * `tags.TagValueIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tagvalue are preserved.

        > **Note:** `tags.TagValueIamPolicy` **cannot** be used in conjunction with `tags.TagValueIamBinding` and `tags.TagValueIamMember` or they will fight over what your policy should be.

        > **Note:** `tags.TagValueIamBinding` resources **can be** used in conjunction with `tags.TagValueIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_tags\_tag\_value\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.tags.TagValueIamPolicy("policy",
            tag_value=google_tags_tag_value["value"]["name"],
            policy_data=admin.policy_data)
        ```

        ## google\_tags\_tag\_value\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.tags.TagValueIamBinding("binding",
            tag_value=google_tags_tag_value["value"]["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_tags\_tag\_value\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.tags.TagValueIamMember("member",
            tag_value=google_tags_tag_value["value"]["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* tagValues/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Tags tagvalue IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:tags/tagValueIamBinding:TagValueIamBinding editor "tagValues/{{tag_value}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:tags/tagValueIamBinding:TagValueIamBinding editor "tagValues/{{tag_value}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:tags/tagValueIamBinding:TagValueIamBinding editor tagValues/{{tag_value}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param TagValueIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TagValueIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['TagValueIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 tag_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TagValueIamBindingArgs.__new__(TagValueIamBindingArgs)

            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if tag_value is None and not opts.urn:
                raise TypeError("Missing required property 'tag_value'")
            __props__.__dict__["tag_value"] = tag_value
            __props__.__dict__["etag"] = None
        super(TagValueIamBinding, __self__).__init__(
            'gcp:tags/tagValueIamBinding:TagValueIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['TagValueIamBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None,
            tag_value: Optional[pulumi.Input[str]] = None) -> 'TagValueIamBinding':
        """
        Get an existing TagValueIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] tag_value: Used to find the parent resource to bind the IAM policy to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TagValueIamBindingState.__new__(_TagValueIamBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["role"] = role
        __props__.__dict__["tag_value"] = tag_value
        return TagValueIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.TagValueIamBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `tags.TagValueIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="tagValue")
    def tag_value(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "tag_value")

