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

__all__ = ['GroupMembershipArgs', 'GroupMembership']

@pulumi.input_type
class GroupMembershipArgs:
    def __init__(__self__, *,
                 group: pulumi.Input[str],
                 roles: pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]],
                 member_key: Optional[pulumi.Input['GroupMembershipMemberKeyArgs']] = None,
                 preferred_member_key: Optional[pulumi.Input['GroupMembershipPreferredMemberKeyArgs']] = None):
        """
        The set of arguments for constructing a GroupMembership resource.
        :param pulumi.Input[str] group: The name of the Group to create this membership in.
        :param pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]] roles: The MembershipRoles that apply to the Membership.
               Must not contain duplicate MembershipRoles with the same name.
               Structure is documented below.
        :param pulumi.Input['GroupMembershipMemberKeyArgs'] member_key: EntityKey of the member.
        :param pulumi.Input['GroupMembershipPreferredMemberKeyArgs'] preferred_member_key: EntityKey of the member.
               Structure is documented below.
        """
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "roles", roles)
        if member_key is not None:
            pulumi.set(__self__, "member_key", member_key)
        if preferred_member_key is not None:
            pulumi.set(__self__, "preferred_member_key", preferred_member_key)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input[str]:
        """
        The name of the Group to create this membership in.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input[str]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]]:
        """
        The MembershipRoles that apply to the Membership.
        Must not contain duplicate MembershipRoles with the same name.
        Structure is documented below.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="memberKey")
    def member_key(self) -> Optional[pulumi.Input['GroupMembershipMemberKeyArgs']]:
        """
        EntityKey of the member.
        """
        return pulumi.get(self, "member_key")

    @member_key.setter
    def member_key(self, value: Optional[pulumi.Input['GroupMembershipMemberKeyArgs']]):
        pulumi.set(self, "member_key", value)

    @property
    @pulumi.getter(name="preferredMemberKey")
    def preferred_member_key(self) -> Optional[pulumi.Input['GroupMembershipPreferredMemberKeyArgs']]:
        """
        EntityKey of the member.
        Structure is documented below.
        """
        return pulumi.get(self, "preferred_member_key")

    @preferred_member_key.setter
    def preferred_member_key(self, value: Optional[pulumi.Input['GroupMembershipPreferredMemberKeyArgs']]):
        pulumi.set(self, "preferred_member_key", value)


@pulumi.input_type
class _GroupMembershipState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 member_key: Optional[pulumi.Input['GroupMembershipMemberKeyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 preferred_member_key: Optional[pulumi.Input['GroupMembershipPreferredMemberKeyArgs']] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupMembership resources.
        :param pulumi.Input[str] create_time: The time when the Membership was created.
        :param pulumi.Input[str] group: The name of the Group to create this membership in.
        :param pulumi.Input['GroupMembershipMemberKeyArgs'] member_key: EntityKey of the member.
        :param pulumi.Input[str] name: The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
               Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        :param pulumi.Input['GroupMembershipPreferredMemberKeyArgs'] preferred_member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]] roles: The MembershipRoles that apply to the Membership.
               Must not contain duplicate MembershipRoles with the same name.
               Structure is documented below.
        :param pulumi.Input[str] type: The type of the membership.
        :param pulumi.Input[str] update_time: The time when the Membership was last updated.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if member_key is not None:
            pulumi.set(__self__, "member_key", member_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if preferred_member_key is not None:
            pulumi.set(__self__, "preferred_member_key", preferred_member_key)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the Membership was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Group to create this membership in.
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter(name="memberKey")
    def member_key(self) -> Optional[pulumi.Input['GroupMembershipMemberKeyArgs']]:
        """
        EntityKey of the member.
        """
        return pulumi.get(self, "member_key")

    @member_key.setter
    def member_key(self, value: Optional[pulumi.Input['GroupMembershipMemberKeyArgs']]):
        pulumi.set(self, "member_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
        Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="preferredMemberKey")
    def preferred_member_key(self) -> Optional[pulumi.Input['GroupMembershipPreferredMemberKeyArgs']]:
        """
        EntityKey of the member.
        Structure is documented below.
        """
        return pulumi.get(self, "preferred_member_key")

    @preferred_member_key.setter
    def preferred_member_key(self, value: Optional[pulumi.Input['GroupMembershipPreferredMemberKeyArgs']]):
        pulumi.set(self, "preferred_member_key", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]]]:
        """
        The MembershipRoles that apply to the Membership.
        Must not contain duplicate MembershipRoles with the same name.
        Structure is documented below.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GroupMembershipRoleArgs']]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the membership.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the Membership was last updated.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class GroupMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 member_key: Optional[pulumi.Input[pulumi.InputType['GroupMembershipMemberKeyArgs']]] = None,
                 preferred_member_key: Optional[pulumi.Input[pulumi.InputType['GroupMembershipPreferredMemberKeyArgs']]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupMembershipRoleArgs']]]]] = None,
                 __props__=None):
        """
        A Membership defines a relationship between a Group and an entity belonging to that Group, referred to as a "member".

        To get more information about GroupMembership, see:

        * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/identity/docs/how-to/memberships-google-groups)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Cloud Identity Group Membership

        ```python
        import pulumi
        import pulumi_gcp as gcp

        group = gcp.cloudidentity.Group("group",
            display_name="my-identity-group",
            parent="customers/A01b123xz",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            })
        child_group = gcp.cloudidentity.Group("child-group",
            display_name="my-identity-group-child",
            parent="customers/A01b123xz",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group-child@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            })
        cloud_identity_group_membership_basic = gcp.cloudidentity.GroupMembership("cloudIdentityGroupMembershipBasic",
            group=group.id,
            preferred_member_key=gcp.cloudidentity.GroupMembershipPreferredMemberKeyArgs(
                id=child_group.group_key.id,
            ),
            roles=[gcp.cloudidentity.GroupMembershipRoleArgs(
                name="MEMBER",
            )])
        ```
        ### Cloud Identity Group Membership User

        ```python
        import pulumi
        import pulumi_gcp as gcp

        group = gcp.cloudidentity.Group("group",
            display_name="my-identity-group",
            parent="customers/A01b123xz",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            })
        cloud_identity_group_membership_basic = gcp.cloudidentity.GroupMembership("cloudIdentityGroupMembershipBasic",
            group=group.id,
            preferred_member_key=gcp.cloudidentity.GroupMembershipPreferredMemberKeyArgs(
                id="cloud_identity_user@example.com",
            ),
            roles=[
                gcp.cloudidentity.GroupMembershipRoleArgs(
                    name="MEMBER",
                ),
                gcp.cloudidentity.GroupMembershipRoleArgs(
                    name="MANAGER",
                ),
            ])
        ```

        ## Import

        GroupMembership can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudidentity/groupMembership:GroupMembership default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group: The name of the Group to create this membership in.
        :param pulumi.Input[pulumi.InputType['GroupMembershipMemberKeyArgs']] member_key: EntityKey of the member.
        :param pulumi.Input[pulumi.InputType['GroupMembershipPreferredMemberKeyArgs']] preferred_member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupMembershipRoleArgs']]]] roles: The MembershipRoles that apply to the Membership.
               Must not contain duplicate MembershipRoles with the same name.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Membership defines a relationship between a Group and an entity belonging to that Group, referred to as a "member".

        To get more information about GroupMembership, see:

        * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/identity/docs/how-to/memberships-google-groups)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Cloud Identity Group Membership

        ```python
        import pulumi
        import pulumi_gcp as gcp

        group = gcp.cloudidentity.Group("group",
            display_name="my-identity-group",
            parent="customers/A01b123xz",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            })
        child_group = gcp.cloudidentity.Group("child-group",
            display_name="my-identity-group-child",
            parent="customers/A01b123xz",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group-child@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            })
        cloud_identity_group_membership_basic = gcp.cloudidentity.GroupMembership("cloudIdentityGroupMembershipBasic",
            group=group.id,
            preferred_member_key=gcp.cloudidentity.GroupMembershipPreferredMemberKeyArgs(
                id=child_group.group_key.id,
            ),
            roles=[gcp.cloudidentity.GroupMembershipRoleArgs(
                name="MEMBER",
            )])
        ```
        ### Cloud Identity Group Membership User

        ```python
        import pulumi
        import pulumi_gcp as gcp

        group = gcp.cloudidentity.Group("group",
            display_name="my-identity-group",
            parent="customers/A01b123xz",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            })
        cloud_identity_group_membership_basic = gcp.cloudidentity.GroupMembership("cloudIdentityGroupMembershipBasic",
            group=group.id,
            preferred_member_key=gcp.cloudidentity.GroupMembershipPreferredMemberKeyArgs(
                id="cloud_identity_user@example.com",
            ),
            roles=[
                gcp.cloudidentity.GroupMembershipRoleArgs(
                    name="MEMBER",
                ),
                gcp.cloudidentity.GroupMembershipRoleArgs(
                    name="MANAGER",
                ),
            ])
        ```

        ## Import

        GroupMembership can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudidentity/groupMembership:GroupMembership default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param GroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group: Optional[pulumi.Input[str]] = None,
                 member_key: Optional[pulumi.Input[pulumi.InputType['GroupMembershipMemberKeyArgs']]] = None,
                 preferred_member_key: Optional[pulumi.Input[pulumi.InputType['GroupMembershipPreferredMemberKeyArgs']]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupMembershipRoleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupMembershipArgs.__new__(GroupMembershipArgs)

            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["member_key"] = member_key
            __props__.__dict__["preferred_member_key"] = preferred_member_key
            if roles is None and not opts.urn:
                raise TypeError("Missing required property 'roles'")
            __props__.__dict__["roles"] = roles
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["update_time"] = None
        super(GroupMembership, __self__).__init__(
            'gcp:cloudidentity/groupMembership:GroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            group: Optional[pulumi.Input[str]] = None,
            member_key: Optional[pulumi.Input[pulumi.InputType['GroupMembershipMemberKeyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            preferred_member_key: Optional[pulumi.Input[pulumi.InputType['GroupMembershipPreferredMemberKeyArgs']]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupMembershipRoleArgs']]]]] = None,
            type: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'GroupMembership':
        """
        Get an existing GroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time when the Membership was created.
        :param pulumi.Input[str] group: The name of the Group to create this membership in.
        :param pulumi.Input[pulumi.InputType['GroupMembershipMemberKeyArgs']] member_key: EntityKey of the member.
        :param pulumi.Input[str] name: The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
               Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        :param pulumi.Input[pulumi.InputType['GroupMembershipPreferredMemberKeyArgs']] preferred_member_key: EntityKey of the member.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GroupMembershipRoleArgs']]]] roles: The MembershipRoles that apply to the Membership.
               Must not contain duplicate MembershipRoles with the same name.
               Structure is documented below.
        :param pulumi.Input[str] type: The type of the membership.
        :param pulumi.Input[str] update_time: The time when the Membership was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupMembershipState.__new__(_GroupMembershipState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["group"] = group
        __props__.__dict__["member_key"] = member_key
        __props__.__dict__["name"] = name
        __props__.__dict__["preferred_member_key"] = preferred_member_key
        __props__.__dict__["roles"] = roles
        __props__.__dict__["type"] = type
        __props__.__dict__["update_time"] = update_time
        return GroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the Membership was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output[str]:
        """
        The name of the Group to create this membership in.
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter(name="memberKey")
    def member_key(self) -> pulumi.Output['outputs.GroupMembershipMemberKey']:
        """
        EntityKey of the member.
        """
        return pulumi.get(self, "member_key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
        Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="preferredMemberKey")
    def preferred_member_key(self) -> pulumi.Output['outputs.GroupMembershipPreferredMemberKey']:
        """
        EntityKey of the member.
        Structure is documented below.
        """
        return pulumi.get(self, "preferred_member_key")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence['outputs.GroupMembershipRole']]:
        """
        The MembershipRoles that apply to the Membership.
        Must not contain duplicate MembershipRoles with the same name.
        Structure is documented below.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the membership.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the Membership was last updated.
        """
        return pulumi.get(self, "update_time")

