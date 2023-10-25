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

__all__ = ['InstanceIAMMemberArgs', 'InstanceIAMMember']

@pulumi.input_type
class InstanceIAMMemberArgs:
    def __init__(__self__, *,
                 instance: pulumi.Input[str],
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['InstanceIAMMemberConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceIAMMember resource.
        :param pulumi.Input[str] instance: The name of the instance.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        InstanceIAMMemberArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            instance=instance,
            member=member,
            role=role,
            condition=condition,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             instance: Optional[pulumi.Input[str]] = None,
             member: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             condition: Optional[pulumi.Input['InstanceIAMMemberConditionArgs']] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if instance is None:
            raise TypeError("Missing 'instance' argument")
        if member is None:
            raise TypeError("Missing 'member' argument")
        if role is None:
            raise TypeError("Missing 'role' argument")

        _setter("instance", instance)
        _setter("member", member)
        _setter("role", role)
        if condition is not None:
            _setter("condition", condition)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The name of the instance.

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

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
        `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['InstanceIAMMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['InstanceIAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _InstanceIAMMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['InstanceIAMMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceIAMMember resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the instance's IAM policy.
        :param pulumi.Input[str] instance: The name of the instance.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        _InstanceIAMMemberState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            condition=condition,
            etag=etag,
            instance=instance,
            member=member,
            project=project,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             condition: Optional[pulumi.Input['InstanceIAMMemberConditionArgs']] = None,
             etag: Optional[pulumi.Input[str]] = None,
             instance: Optional[pulumi.Input[str]] = None,
             member: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if condition is not None:
            _setter("condition", condition)
        if etag is not None:
            _setter("etag", etag)
        if instance is not None:
            _setter("instance", instance)
        if member is not None:
            _setter("member", member)
        if project is not None:
            _setter("project", project)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['InstanceIAMMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['InstanceIAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the instance's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the instance.

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

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
        The ID of the project in which the resource belongs. If it
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
        `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class InstanceIAMMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['InstanceIAMMemberConditionArgs']]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:

        * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.

        > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.

        * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.

        > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_spanner\\_instance\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        instance = gcp.spanner.InstanceIAMPolicy("instance",
            instance="your-instance-name",
            policy_data=admin.policy_data)
        ```

        ## google\\_spanner\\_instance\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMBinding("instance",
            instance="your-instance-name",
            members=["user:jane@example.com"],
            role="roles/spanner.databaseAdmin")
        ```

        ## google\\_spanner\\_instance\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMMember("instance",
            instance="your-instance-name",
            member="user:jane@example.com",
            role="roles/spanner.databaseAdmin")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{name}} * {{name}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the account, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance "project-name/instance-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance "project-name/instance-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance project-name/instance-name
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name of the instance.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceIAMMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for a Spanner instance. Each of these resources serves a different use case:

        * `spanner.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.

        > **Warning:** It's entirely possibly to lock yourself out of your instance using `spanner.InstanceIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.

        * `spanner.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `spanner.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `spanner.InstanceIAMPolicy` **cannot** be used in conjunction with `spanner.InstanceIAMBinding` and `spanner.InstanceIAMMember` or they will fight over what your policy should be.

        > **Note:** `spanner.InstanceIAMBinding` resources **can be** used in conjunction with `spanner.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_spanner\\_instance\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        instance = gcp.spanner.InstanceIAMPolicy("instance",
            instance="your-instance-name",
            policy_data=admin.policy_data)
        ```

        ## google\\_spanner\\_instance\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMBinding("instance",
            instance="your-instance-name",
            members=["user:jane@example.com"],
            role="roles/spanner.databaseAdmin")
        ```

        ## google\\_spanner\\_instance\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        instance = gcp.spanner.InstanceIAMMember("instance",
            instance="your-instance-name",
            member="user:jane@example.com",
            role="roles/spanner.databaseAdmin")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{name}} * {{name}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the account, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance "project-name/instance-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance "project-name/instance-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:spanner/instanceIAMMember:InstanceIAMMember instance project-name/instance-name
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param InstanceIAMMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIAMMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            InstanceIAMMemberArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['InstanceIAMMemberConditionArgs']]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceIAMMemberArgs.__new__(InstanceIAMMemberArgs)

            condition = _utilities.configure(condition, InstanceIAMMemberConditionArgs, True)
            __props__.__dict__["condition"] = condition
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__.__dict__["instance"] = instance
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(InstanceIAMMember, __self__).__init__(
            'gcp:spanner/instanceIAMMember:InstanceIAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['InstanceIAMMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'InstanceIAMMember':
        """
        Get an existing InstanceIAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the instance's IAM policy.
        :param pulumi.Input[str] instance: The name of the instance.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceIAMMemberState.__new__(_InstanceIAMMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["instance"] = instance
        __props__.__dict__["member"] = member
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        return InstanceIAMMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.InstanceIAMMemberCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the instance's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name of the instance.

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `spanner.InstanceIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

