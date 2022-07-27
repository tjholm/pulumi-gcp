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

__all__ = ['SubscriptionIAMMemberArgs', 'SubscriptionIAMMember']

@pulumi.input_type
class SubscriptionIAMMemberArgs:
    def __init__(__self__, *,
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 subscription: pulumi.Input[str],
                 condition: Optional[pulumi.Input['SubscriptionIAMMemberConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubscriptionIAMMember resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subscription: The subscription name or id to bind to attach IAM policy to.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "subscription", subscription)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project is not None:
            pulumi.set(__self__, "project", project)

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
        `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def subscription(self) -> pulumi.Input[str]:
        """
        The subscription name or id to bind to attach IAM policy to.
        """
        return pulumi.get(self, "subscription")

    @subscription.setter
    def subscription(self, value: pulumi.Input[str]):
        pulumi.set(self, "subscription", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['SubscriptionIAMMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['SubscriptionIAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _SubscriptionIAMMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['SubscriptionIAMMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 subscription: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SubscriptionIAMMember resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the subscription's IAM policy.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subscription: The subscription name or id to bind to attach IAM policy to.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if subscription is not None:
            pulumi.set(__self__, "subscription", subscription)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['SubscriptionIAMMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['SubscriptionIAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the subscription's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def subscription(self) -> Optional[pulumi.Input[str]]:
        """
        The subscription name or id to bind to attach IAM policy to.
        """
        return pulumi.get(self, "subscription")

    @subscription.setter
    def subscription(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription", value)


class SubscriptionIAMMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['SubscriptionIAMMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 subscription: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:

        * `pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
        * `pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
        * `pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.

        > **Note:** `pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `pubsub.SubscriptionIAMBinding` and `pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.

        > **Note:** `pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_pubsub\\_subscription\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        editor = gcp.pubsub.SubscriptionIAMPolicy("editor",
            subscription="your-subscription-name",
            policy_data=admin.policy_data)
        ```

        ## google\\_pubsub\\_subscription\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.pubsub.SubscriptionIAMBinding("editor",
            members=["user:jane@example.com"],
            role="roles/editor",
            subscription="your-subscription-name")
        ```

        ## google\\_pubsub\\_subscription\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.pubsub.SubscriptionIAMMember("editor",
            member="user:jane@example.com",
            role="roles/editor",
            subscription="your-subscription-name")
        ```

        ## Import

        Pubsub subscription IAM resources can be imported using the project, subscription name, role and member.

        ```sh
         $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor projects/{your-project-id}/subscriptions/{your-subscription-name}
        ```

        ```sh
         $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor"
        ```

        ```sh
         $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subscription: The subscription name or id to bind to attach IAM policy to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubscriptionIAMMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for pubsub subscription. Each of these resources serves a different use case:

        * `pubsub.SubscriptionIAMPolicy`: Authoritative. Sets the IAM policy for the subscription and replaces any existing policy already attached.
        * `pubsub.SubscriptionIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subscription are preserved.
        * `pubsub.SubscriptionIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subscription are preserved.

        > **Note:** `pubsub.SubscriptionIAMPolicy` **cannot** be used in conjunction with `pubsub.SubscriptionIAMBinding` and `pubsub.SubscriptionIAMMember` or they will fight over what your policy should be.

        > **Note:** `pubsub.SubscriptionIAMBinding` resources **can be** used in conjunction with `pubsub.SubscriptionIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_pubsub\\_subscription\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        editor = gcp.pubsub.SubscriptionIAMPolicy("editor",
            subscription="your-subscription-name",
            policy_data=admin.policy_data)
        ```

        ## google\\_pubsub\\_subscription\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.pubsub.SubscriptionIAMBinding("editor",
            members=["user:jane@example.com"],
            role="roles/editor",
            subscription="your-subscription-name")
        ```

        ## google\\_pubsub\\_subscription\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.pubsub.SubscriptionIAMMember("editor",
            member="user:jane@example.com",
            role="roles/editor",
            subscription="your-subscription-name")
        ```

        ## Import

        Pubsub subscription IAM resources can be imported using the project, subscription name, role and member.

        ```sh
         $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor projects/{your-project-id}/subscriptions/{your-subscription-name}
        ```

        ```sh
         $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor"
        ```

        ```sh
         $ pulumi import gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param SubscriptionIAMMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubscriptionIAMMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['SubscriptionIAMMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 subscription: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubscriptionIAMMemberArgs.__new__(SubscriptionIAMMemberArgs)

            __props__.__dict__["condition"] = condition
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if subscription is None and not opts.urn:
                raise TypeError("Missing required property 'subscription'")
            __props__.__dict__["subscription"] = subscription
            __props__.__dict__["etag"] = None
        super(SubscriptionIAMMember, __self__).__init__(
            'gcp:pubsub/subscriptionIAMMember:SubscriptionIAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['SubscriptionIAMMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            subscription: Optional[pulumi.Input[str]] = None) -> 'SubscriptionIAMMember':
        """
        Get an existing SubscriptionIAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the subscription's IAM policy.
        :param pulumi.Input[str] project: The project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subscription: The subscription name or id to bind to attach IAM policy to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubscriptionIAMMemberState.__new__(_SubscriptionIAMMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["member"] = member
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        __props__.__dict__["subscription"] = subscription
        return SubscriptionIAMMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.SubscriptionIAMMemberCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the subscription's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `pubsub.SubscriptionIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def subscription(self) -> pulumi.Output[str]:
        """
        The subscription name or id to bind to attach IAM policy to.
        """
        return pulumi.get(self, "subscription")

