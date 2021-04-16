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

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 group_key: pulumi.Input['GroupGroupKeyArgs'],
                 labels: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 parent: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input['GroupGroupKeyArgs'] group_key: EntityKey of the Group.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels that apply to the Group.
               Must not contain more than one entry. Must contain the entry
               'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
               'system/groups/external': '' if the Group is an external-identity-mapped group.
        :param pulumi.Input[str] parent: The resource name of the entity under which this Group resides in the
               Cloud Identity resource hierarchy.
               Must be of the form identitysources/{identity_source_id} for external-identity-mapped
               groups or customers/{customer_id} for Google Groups.
        :param pulumi.Input[str] description: An extended description to help users determine the purpose of a Group.
               Must not be longer than 4,096 characters.
        :param pulumi.Input[str] display_name: The display name of the Group.
        """
        pulumi.set(__self__, "group_key", group_key)
        pulumi.set(__self__, "labels", labels)
        pulumi.set(__self__, "parent", parent)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)

    @property
    @pulumi.getter(name="groupKey")
    def group_key(self) -> pulumi.Input['GroupGroupKeyArgs']:
        """
        EntityKey of the Group.
        Structure is documented below.
        """
        return pulumi.get(self, "group_key")

    @group_key.setter
    def group_key(self, value: pulumi.Input['GroupGroupKeyArgs']):
        pulumi.set(self, "group_key", value)

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        The labels that apply to the Group.
        Must not contain more than one entry. Must contain the entry
        'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
        'system/groups/external': '' if the Group is an external-identity-mapped group.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The resource name of the entity under which this Group resides in the
        Cloud Identity resource hierarchy.
        Must be of the form identitysources/{identity_source_id} for external-identity-mapped
        groups or customers/{customer_id} for Google Groups.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An extended description to help users determine the purpose of a Group.
        Must not be longer than 4,096 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the Group.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 group_key: Optional[pulumi.Input['GroupGroupKeyArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[str] create_time: The time when the Group was created.
        :param pulumi.Input[str] description: An extended description to help users determine the purpose of a Group.
               Must not be longer than 4,096 characters.
        :param pulumi.Input[str] display_name: The display name of the Group.
        :param pulumi.Input['GroupGroupKeyArgs'] group_key: EntityKey of the Group.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels that apply to the Group.
               Must not contain more than one entry. Must contain the entry
               'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
               'system/groups/external': '' if the Group is an external-identity-mapped group.
        :param pulumi.Input[str] name: Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
        :param pulumi.Input[str] parent: The resource name of the entity under which this Group resides in the
               Cloud Identity resource hierarchy.
               Must be of the form identitysources/{identity_source_id} for external-identity-mapped
               groups or customers/{customer_id} for Google Groups.
        :param pulumi.Input[str] update_time: The time when the Group was last updated.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if group_key is not None:
            pulumi.set(__self__, "group_key", group_key)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the Group was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An extended description to help users determine the purpose of a Group.
        Must not be longer than 4,096 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the Group.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="groupKey")
    def group_key(self) -> Optional[pulumi.Input['GroupGroupKeyArgs']]:
        """
        EntityKey of the Group.
        Structure is documented below.
        """
        return pulumi.get(self, "group_key")

    @group_key.setter
    def group_key(self, value: Optional[pulumi.Input['GroupGroupKeyArgs']]):
        pulumi.set(self, "group_key", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels that apply to the Group.
        Must not contain more than one entry. Must contain the entry
        'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
        'system/groups/external': '' if the Group is an external-identity-mapped group.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the entity under which this Group resides in the
        Cloud Identity resource hierarchy.
        Must be of the form identitysources/{identity_source_id} for external-identity-mapped
        groups or customers/{customer_id} for Google Groups.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the Group was last updated.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 group_key: Optional[pulumi.Input[pulumi.InputType['GroupGroupKeyArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A Cloud Identity resource representing a Group.

        To get more information about Group, see:

        * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/identity/docs/how-to/setup)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Cloud Identity Groups Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        cloud_identity_group_basic = gcp.cloudidentity.Group("cloudIdentityGroupBasic",
            display_name="my-identity-group",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            },
            parent="customers/A01b123xz")
        ```

        ## Import

        Group can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudidentity/group:Group default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An extended description to help users determine the purpose of a Group.
               Must not be longer than 4,096 characters.
        :param pulumi.Input[str] display_name: The display name of the Group.
        :param pulumi.Input[pulumi.InputType['GroupGroupKeyArgs']] group_key: EntityKey of the Group.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels that apply to the Group.
               Must not contain more than one entry. Must contain the entry
               'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
               'system/groups/external': '' if the Group is an external-identity-mapped group.
        :param pulumi.Input[str] parent: The resource name of the entity under which this Group resides in the
               Cloud Identity resource hierarchy.
               Must be of the form identitysources/{identity_source_id} for external-identity-mapped
               groups or customers/{customer_id} for Google Groups.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Cloud Identity resource representing a Group.

        To get more information about Group, see:

        * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1beta1/groups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/identity/docs/how-to/setup)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Cloud Identity Groups Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        cloud_identity_group_basic = gcp.cloudidentity.Group("cloudIdentityGroupBasic",
            display_name="my-identity-group",
            group_key=gcp.cloudidentity.GroupGroupKeyArgs(
                id="my-identity-group@example.com",
            ),
            labels={
                "cloudidentity.googleapis.com/groups.discussion_forum": "",
            },
            parent="customers/A01b123xz")
        ```

        ## Import

        Group can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudidentity/group:Group default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 group_key: Optional[pulumi.Input[pulumi.InputType['GroupGroupKeyArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
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
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if group_key is None and not opts.urn:
                raise TypeError("Missing required property 'group_key'")
            __props__.__dict__["group_key"] = group_key
            if labels is None and not opts.urn:
                raise TypeError("Missing required property 'labels'")
            __props__.__dict__["labels"] = labels
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__["parent"] = parent
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(Group, __self__).__init__(
            'gcp:cloudidentity/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            group_key: Optional[pulumi.Input[pulumi.InputType['GroupGroupKeyArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time when the Group was created.
        :param pulumi.Input[str] description: An extended description to help users determine the purpose of a Group.
               Must not be longer than 4,096 characters.
        :param pulumi.Input[str] display_name: The display name of the Group.
        :param pulumi.Input[pulumi.InputType['GroupGroupKeyArgs']] group_key: EntityKey of the Group.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels that apply to the Group.
               Must not contain more than one entry. Must contain the entry
               'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
               'system/groups/external': '' if the Group is an external-identity-mapped group.
        :param pulumi.Input[str] name: Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
        :param pulumi.Input[str] parent: The resource name of the entity under which this Group resides in the
               Cloud Identity resource hierarchy.
               Must be of the form identitysources/{identity_source_id} for external-identity-mapped
               groups or customers/{customer_id} for Google Groups.
        :param pulumi.Input[str] update_time: The time when the Group was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["group_key"] = group_key
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["parent"] = parent
        __props__.__dict__["update_time"] = update_time
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the Group was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An extended description to help users determine the purpose of a Group.
        Must not be longer than 4,096 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The display name of the Group.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="groupKey")
    def group_key(self) -> pulumi.Output['outputs.GroupGroupKey']:
        """
        EntityKey of the Group.
        Structure is documented below.
        """
        return pulumi.get(self, "group_key")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The labels that apply to the Group.
        Must not contain more than one entry. Must contain the entry
        'cloudidentity.googleapis.com/groups.discussion_forum': '' if the Group is a Google Group or
        'system/groups/external': '' if the Group is an external-identity-mapped group.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the Group in the format: groups/{group_id}, where group_id is the unique ID assigned to the Group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The resource name of the entity under which this Group resides in the
        Cloud Identity resource hierarchy.
        Must be of the form identitysources/{identity_source_id} for external-identity-mapped
        groups or customers/{customer_id} for Google Groups.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the Group was last updated.
        """
        return pulumi.get(self, "update_time")

