# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetInstanceGroupResult',
    'AwaitableGetInstanceGroupResult',
    'get_instance_group',
]

@pulumi.output_type
class GetInstanceGroupResult:
    """
    A collection of values returned by getInstanceGroup.
    """
    def __init__(__self__, description=None, id=None, instances=None, name=None, named_ports=None, network=None, project=None, self_link=None, size=None, zone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if named_ports and not isinstance(named_ports, list):
            raise TypeError("Expected argument 'named_ports' to be a list")
        pulumi.set(__self__, "named_ports", named_ports)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Textual description of the instance group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence[str]:
        """
        List of instances in the group.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namedPorts")
    def named_ports(self) -> Sequence['outputs.GetInstanceGroupNamedPortResult']:
        """
        List of named ports in the group.
        """
        return pulumi.get(self, "named_ports")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The URL of the network the instance group is in.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The number of instances in the group.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetInstanceGroupResult(GetInstanceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceGroupResult(
            description=self.description,
            id=self.id,
            instances=self.instances,
            name=self.name,
            named_ports=self.named_ports,
            network=self.network,
            project=self.project,
            self_link=self.self_link,
            size=self.size,
            zone=self.zone)


def get_instance_group(name: Optional[str] = None,
                       project: Optional[str] = None,
                       self_link: Optional[str] = None,
                       zone: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceGroupResult:
    """
    Get a Compute Instance Group within GCE.
    For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
    and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)

    ```python
    import pulumi
    import pulumi_gcp as gcp

    all = gcp.compute.get_instance_group(name="instance-group-name",
        zone="us-central1-a")
    ```


    :param str name: The name of the instance group. Either `name` or `self_link` must be provided.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str self_link: The self link of the instance group. Either `name` or `self_link` must be provided.
    :param str zone: The zone of the instance group. If referencing the instance group by name
           and `zone` is not provided, the provider zone is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['selfLink'] = self_link
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getInstanceGroup:getInstanceGroup', __args__, opts=opts, typ=GetInstanceGroupResult).value

    return AwaitableGetInstanceGroupResult(
        description=__ret__.description,
        id=__ret__.id,
        instances=__ret__.instances,
        name=__ret__.name,
        named_ports=__ret__.named_ports,
        network=__ret__.network,
        project=__ret__.project,
        self_link=__ret__.self_link,
        size=__ret__.size,
        zone=__ret__.zone)
