# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['IAMBindingArgs', 'IAMBinding']

@pulumi.input_type
class IAMBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 org_id: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['IAMBindingConditionArgs']] = None):
        """
        The set of arguments for constructing a IAMBinding resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        IAMBindingArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            members=members,
            org_id=org_id,
            role=role,
            condition=condition,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             condition: Optional[pulumi.Input['IAMBindingConditionArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if members is None:
            raise TypeError("Missing 'members' argument")
        if org_id is None and 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if org_id is None:
            raise TypeError("Missing 'org_id' argument")
        if role is None:
            raise TypeError("Missing 'role' argument")

        _setter("members", members)
        _setter("org_id", org_id)
        _setter("role", role)
        if condition is not None:
            _setter("condition", condition)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        The numeric ID of the organization in which you want to create a custom role.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['IAMBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['IAMBindingConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _IAMBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['IAMBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMBinding resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the organization's IAM policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        _IAMBindingState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            condition=condition,
            etag=etag,
            members=members,
            org_id=org_id,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             condition: Optional[pulumi.Input['IAMBindingConditionArgs']] = None,
             etag: Optional[pulumi.Input[str]] = None,
             members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if org_id is None and 'orgId' in kwargs:
            org_id = kwargs['orgId']

        if condition is not None:
            _setter("condition", condition)
        if etag is not None:
            _setter("etag", etag)
        if members is not None:
            _setter("members", members)
        if org_id is not None:
            _setter("org_id", org_id)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['IAMBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['IAMBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

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
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The numeric ID of the organization in which you want to create a custom role.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class IAMBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows creation and management of a single binding within IAM policy for
        an existing Google Cloud Platform Organization.

        > **Note:** This resource __must not__ be used in conjunction with
           `organizations.IAMMember` for the __same role__ or they will fight over
           what your policy should be.

        > **Note:** On create, this resource will overwrite members of any existing roles.
            Use `pulumi import` and inspect the `output to ensure
            your existing members are preserved.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.organizations.IAMBinding("binding",
            members=["user:alice@gmail.com"],
            org_id="123456789",
            role="roles/browser")
        ```

        ## Import

        IAM binding imports use space-delimited identifiers; first the resource in question and then the role.

        These bindings can be imported using the `org_id` and role, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMBinding:IAMBinding my_org "your-org-id roles/viewer"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows creation and management of a single binding within IAM policy for
        an existing Google Cloud Platform Organization.

        > **Note:** This resource __must not__ be used in conjunction with
           `organizations.IAMMember` for the __same role__ or they will fight over
           what your policy should be.

        > **Note:** On create, this resource will overwrite members of any existing roles.
            Use `pulumi import` and inspect the `output to ensure
            your existing members are preserved.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.organizations.IAMBinding("binding",
            members=["user:alice@gmail.com"],
            org_id="123456789",
            role="roles/browser")
        ```

        ## Import

        IAM binding imports use space-delimited identifiers; first the resource in question and then the role.

        These bindings can be imported using the `org_id` and role, e.g.

        ```sh
         $ pulumi import gcp:organizations/iAMBinding:IAMBinding my_org "your-org-id roles/viewer"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param IAMBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            IAMBindingArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMBindingArgs.__new__(IAMBindingArgs)

            condition = _utilities.configure(condition, IAMBindingConditionArgs, True)
            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(IAMBinding, __self__).__init__(
            'gcp:organizations/iAMBinding:IAMBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'IAMBinding':
        """
        Get an existing IAMBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the organization's IAM policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        :param pulumi.Input[str] org_id: The numeric ID of the organization in which you want to create a custom role.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMBindingState.__new__(_IAMBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["role"] = role
        return IAMBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.IAMBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the organization's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The numeric ID of the organization in which you want to create a custom role.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `organizations.IAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

