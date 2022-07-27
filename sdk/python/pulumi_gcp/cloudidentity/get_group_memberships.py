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

__all__ = [
    'GetGroupMembershipsResult',
    'AwaitableGetGroupMembershipsResult',
    'get_group_memberships',
    'get_group_memberships_output',
]

@pulumi.output_type
class GetGroupMembershipsResult:
    """
    A collection of values returned by getGroupMemberships.
    """
    def __init__(__self__, group=None, id=None, memberships=None):
        if group and not isinstance(group, str):
            raise TypeError("Expected argument 'group' to be a str")
        pulumi.set(__self__, "group", group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if memberships and not isinstance(memberships, list):
            raise TypeError("Expected argument 'memberships' to be a list")
        pulumi.set(__self__, "memberships", memberships)

    @property
    @pulumi.getter
    def group(self) -> str:
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def memberships(self) -> Sequence['outputs.GetGroupMembershipsMembershipResult']:
        """
        The list of memberships under the given group. Structure is documented below.
        """
        return pulumi.get(self, "memberships")


class AwaitableGetGroupMembershipsResult(GetGroupMembershipsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupMembershipsResult(
            group=self.group,
            id=self.id,
            memberships=self.memberships)


def get_group_memberships(group: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupMembershipsResult:
    """
    Use this data source to get list of the Cloud Identity Group Memberships within a given Group.

    https://cloud.google.com/identity/docs/concepts/overview#memberships

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    members = gcp.cloudidentity.get_group_memberships(group="groups/123eab45c6defghi")
    ```


    :param str group: The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
    """
    __args__ = dict()
    __args__['group'] = group
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:cloudidentity/getGroupMemberships:getGroupMemberships', __args__, opts=opts, typ=GetGroupMembershipsResult).value

    return AwaitableGetGroupMembershipsResult(
        group=__ret__.group,
        id=__ret__.id,
        memberships=__ret__.memberships)


@_utilities.lift_output_func(get_group_memberships)
def get_group_memberships_output(group: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupMembershipsResult]:
    """
    Use this data source to get list of the Cloud Identity Group Memberships within a given Group.

    https://cloud.google.com/identity/docs/concepts/overview#memberships

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    members = gcp.cloudidentity.get_group_memberships(group="groups/123eab45c6defghi")
    ```


    :param str group: The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
    """
    ...
