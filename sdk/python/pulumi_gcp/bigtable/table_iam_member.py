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

__all__ = ['TableIamMemberArgs', 'TableIamMember']

@pulumi.input_type
class TableIamMemberArgs:
    def __init__(__self__, *,
                 instance: pulumi.Input[str],
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 table: pulumi.Input[str],
                 condition: Optional[pulumi.Input['TableIamMemberConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TableIamMember resource.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance that owns the table.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `bigtable.TableIamPolicy` only:
        :param pulumi.Input[str] table: The name or relative resource id of the table to manage IAM policies for.
               
               For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] project: The project in which the table belongs. If it
               is not provided, this provider will use the provider default.
        """
        pulumi.set(__self__, "instance", instance)
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "table", table)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The name or relative resource id of the instance that owns the table.
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
        `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

        `bigtable.TableIamPolicy` only:
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def table(self) -> pulumi.Input[str]:
        """
        The name or relative resource id of the table to manage IAM policies for.

        For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "table")

    @table.setter
    def table(self, value: pulumi.Input[str]):
        pulumi.set(self, "table", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['TableIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['TableIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the table belongs. If it
        is not provided, this provider will use the provider default.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _TableIamMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['TableIamMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TableIamMember resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the tables's IAM policy.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance that owns the table.
        :param pulumi.Input[str] project: The project in which the table belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `bigtable.TableIamPolicy` only:
        :param pulumi.Input[str] table: The name or relative resource id of the table to manage IAM policies for.
               
               For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if table is not None:
            pulumi.set(__self__, "table", table)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['TableIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['TableIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the tables's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        The name or relative resource id of the instance that owns the table.
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
        The project in which the table belongs. If it
        is not provided, this provider will use the provider default.
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
        `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

        `bigtable.TableIamPolicy` only:
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def table(self) -> Optional[pulumi.Input[str]]:
        """
        The name or relative resource id of the table to manage IAM policies for.

        For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "table")

    @table.setter
    def table(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table", value)


class TableIamMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['TableIamMemberConditionArgs']]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:

        * `bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
        * `bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.

        > **Note:** `bigtable.TableIamPolicy` **cannot** be used in conjunction with `bigtable.TableIamBinding` and `bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `bigtable.TableIamPolicy` replaces the entire policy.

        > **Note:** `bigtable.TableIamBinding` resources **can be** used in conjunction with `bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_bigtable\\_table\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/bigtable.user",
            members=["user:jane@example.com"],
        )])
        editor = gcp.bigtable.TableIamPolicy("editor",
            project="your-project",
            instance="your-bigtable-instance",
            table="your-bigtable-table",
            policy_data=admin.policy_data)
        ```

        ## google\\_bigtable\\_table\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.TableIamBinding("editor",
            instance="your-bigtable-instance",
            members=["user:jane@example.com"],
            role="roles/bigtable.user",
            table="your-bigtable-table")
        ```

        ## google\\_bigtable\\_table\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.TableIamMember("editor",
            instance="your-bigtable-instance",
            member="user:jane@example.com",
            role="roles/bigtable.user",
            table="your-bigtable-table")
        ```

        ## Import

        ### Importing IAM policies IAM policy imports use the `table` identifier of the Bigtable Table resource only. For example* `"projects/{project}/tables/{table}"` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {

         id = "projects/{project}/tables/{table}"

         to = google_bigtable_table_iam_policy.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:bigtable/tableIamMember:TableIamMember default projects/{project}/tables/{table}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance that owns the table.
        :param pulumi.Input[str] project: The project in which the table belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `bigtable.TableIamPolicy` only:
        :param pulumi.Input[str] table: The name or relative resource id of the table to manage IAM policies for.
               
               For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TableIamMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:

        * `bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
        * `bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.

        > **Note:** `bigtable.TableIamPolicy` **cannot** be used in conjunction with `bigtable.TableIamBinding` and `bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `bigtable.TableIamPolicy` replaces the entire policy.

        > **Note:** `bigtable.TableIamBinding` resources **can be** used in conjunction with `bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_bigtable\\_table\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/bigtable.user",
            members=["user:jane@example.com"],
        )])
        editor = gcp.bigtable.TableIamPolicy("editor",
            project="your-project",
            instance="your-bigtable-instance",
            table="your-bigtable-table",
            policy_data=admin.policy_data)
        ```

        ## google\\_bigtable\\_table\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.TableIamBinding("editor",
            instance="your-bigtable-instance",
            members=["user:jane@example.com"],
            role="roles/bigtable.user",
            table="your-bigtable-table")
        ```

        ## google\\_bigtable\\_table\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.TableIamMember("editor",
            instance="your-bigtable-instance",
            member="user:jane@example.com",
            role="roles/bigtable.user",
            table="your-bigtable-table")
        ```

        ## Import

        ### Importing IAM policies IAM policy imports use the `table` identifier of the Bigtable Table resource only. For example* `"projects/{project}/tables/{table}"` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {

         id = "projects/{project}/tables/{table}"

         to = google_bigtable_table_iam_policy.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:bigtable/tableIamMember:TableIamMember default projects/{project}/tables/{table}
        ```

        :param str resource_name: The name of the resource.
        :param TableIamMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TableIamMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['TableIamMemberConditionArgs']]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TableIamMemberArgs.__new__(TableIamMemberArgs)

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
            if table is None and not opts.urn:
                raise TypeError("Missing required property 'table'")
            __props__.__dict__["table"] = table
            __props__.__dict__["etag"] = None
        super(TableIamMember, __self__).__init__(
            'gcp:bigtable/tableIamMember:TableIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['TableIamMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            table: Optional[pulumi.Input[str]] = None) -> 'TableIamMember':
        """
        Get an existing TableIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the tables's IAM policy.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance that owns the table.
        :param pulumi.Input[str] project: The project in which the table belongs. If it
               is not provided, this provider will use the provider default.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `bigtable.TableIamPolicy` only:
        :param pulumi.Input[str] table: The name or relative resource id of the table to manage IAM policies for.
               
               For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TableIamMemberState.__new__(_TableIamMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["instance"] = instance
        __props__.__dict__["member"] = member
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        __props__.__dict__["table"] = table
        return TableIamMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.TableIamMemberCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the tables's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name or relative resource id of the instance that owns the table.
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
        The project in which the table belongs. If it
        is not provided, this provider will use the provider default.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

        `bigtable.TableIamPolicy` only:
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def table(self) -> pulumi.Output[str]:
        """
        The name or relative resource id of the table to manage IAM policies for.

        For `bigtable.TableIamMember` or `bigtable.TableIamBinding`:

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "table")

