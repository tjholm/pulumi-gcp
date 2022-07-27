# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GroupGroupKeyArgs',
    'GroupMembershipMemberKeyArgs',
    'GroupMembershipPreferredMemberKeyArgs',
    'GroupMembershipRoleArgs',
]

@pulumi.input_type
class GroupGroupKeyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of the entity.
               For Google-managed entities, the id must be the email address of an existing
               group or user.
               For external-identity-mapped entities, the id must be a string conforming
               to the Identity Source's requirements.
               Must be unique within a namespace.
        :param pulumi.Input[str] namespace: The namespace in which the entity exists.
               If not specified, the EntityKey represents a Google-managed entity
               such as a Google user or a Google Group.
               If specified, the EntityKey represents an external-identity-mapped group.
               The namespace must correspond to an identity source created in Admin Console
               and must be in the form of `identitysources/{identity_source_id}`.
        """
        pulumi.set(__self__, "id", id)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of the entity.
        For Google-managed entities, the id must be the email address of an existing
        group or user.
        For external-identity-mapped entities, the id must be a string conforming
        to the Identity Source's requirements.
        Must be unique within a namespace.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace in which the entity exists.
        If not specified, the EntityKey represents a Google-managed entity
        such as a Google user or a Google Group.
        If specified, the EntityKey represents an external-identity-mapped group.
        The namespace must correspond to an identity source created in Admin Console
        and must be in the form of `identitysources/{identity_source_id}`.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class GroupMembershipMemberKeyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of the entity.
               For Google-managed entities, the id must be the email address of an existing
               group or user.
               For external-identity-mapped entities, the id must be a string conforming
               to the Identity Source's requirements.
               Must be unique within a namespace.
        :param pulumi.Input[str] namespace: The namespace in which the entity exists.
               If not specified, the EntityKey represents a Google-managed entity
               such as a Google user or a Google Group.
               If specified, the EntityKey represents an external-identity-mapped group.
               The namespace must correspond to an identity source created in Admin Console
               and must be in the form of `identitysources/{identity_source_id}`.
        """
        pulumi.set(__self__, "id", id)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of the entity.
        For Google-managed entities, the id must be the email address of an existing
        group or user.
        For external-identity-mapped entities, the id must be a string conforming
        to the Identity Source's requirements.
        Must be unique within a namespace.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace in which the entity exists.
        If not specified, the EntityKey represents a Google-managed entity
        such as a Google user or a Google Group.
        If specified, the EntityKey represents an external-identity-mapped group.
        The namespace must correspond to an identity source created in Admin Console
        and must be in the form of `identitysources/{identity_source_id}`.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class GroupMembershipPreferredMemberKeyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The ID of the entity.
               For Google-managed entities, the id must be the email address of an existing
               group or user.
               For external-identity-mapped entities, the id must be a string conforming
               to the Identity Source's requirements.
               Must be unique within a namespace.
        :param pulumi.Input[str] namespace: The namespace in which the entity exists.
               If not specified, the EntityKey represents a Google-managed entity
               such as a Google user or a Google Group.
               If specified, the EntityKey represents an external-identity-mapped group.
               The namespace must correspond to an identity source created in Admin Console
               and must be in the form of `identitysources/{identity_source_id}`.
        """
        pulumi.set(__self__, "id", id)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of the entity.
        For Google-managed entities, the id must be the email address of an existing
        group or user.
        For external-identity-mapped entities, the id must be a string conforming
        to the Identity Source's requirements.
        Must be unique within a namespace.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace in which the entity exists.
        If not specified, the EntityKey represents a Google-managed entity
        such as a Google user or a Google Group.
        If specified, the EntityKey represents an external-identity-mapped group.
        The namespace must correspond to an identity source created in Admin Console
        and must be in the form of `identitysources/{identity_source_id}`.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class GroupMembershipRoleArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
               Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
        Possible values are `OWNER`, `MANAGER`, and `MEMBER`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


