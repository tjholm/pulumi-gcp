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
    'GetSubnetworkResult',
    'AwaitableGetSubnetworkResult',
    'get_subnetwork',
]

@pulumi.output_type
class GetSubnetworkResult:
    """
    A collection of values returned by getSubnetwork.
    """
    def __init__(__self__, description=None, gateway_address=None, id=None, ip_cidr_range=None, name=None, network=None, private_ip_google_access=None, project=None, region=None, secondary_ip_ranges=None, self_link=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if gateway_address and not isinstance(gateway_address, str):
            raise TypeError("Expected argument 'gateway_address' to be a str")
        pulumi.set(__self__, "gateway_address", gateway_address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_cidr_range and not isinstance(ip_cidr_range, str):
            raise TypeError("Expected argument 'ip_cidr_range' to be a str")
        pulumi.set(__self__, "ip_cidr_range", ip_cidr_range)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if private_ip_google_access and not isinstance(private_ip_google_access, bool):
            raise TypeError("Expected argument 'private_ip_google_access' to be a bool")
        pulumi.set(__self__, "private_ip_google_access", private_ip_google_access)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if secondary_ip_ranges and not isinstance(secondary_ip_ranges, list):
            raise TypeError("Expected argument 'secondary_ip_ranges' to be a list")
        pulumi.set(__self__, "secondary_ip_ranges", secondary_ip_ranges)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of this subnetwork.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="gatewayAddress")
    def gateway_address(self) -> str:
        """
        The IP address of the gateway.
        """
        return pulumi.get(self, "gateway_address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipCidrRange")
    def ip_cidr_range(self) -> str:
        """
        The range of IP addresses belonging to this subnetwork
        secondary range.
        """
        return pulumi.get(self, "ip_cidr_range")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The network name or resource link to the parent
        network of this subnetwork.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="privateIpGoogleAccess")
    def private_ip_google_access(self) -> bool:
        """
        Whether the VMs in this subnet
        can access Google services without assigned external IP
        addresses.
        """
        return pulumi.get(self, "private_ip_google_access")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secondaryIpRanges")
    def secondary_ip_ranges(self) -> Sequence['outputs.GetSubnetworkSecondaryIpRangeResult']:
        """
        An array of configurations for secondary IP ranges for
        VM instances contained in this subnetwork. Structure is documented below.
        """
        return pulumi.get(self, "secondary_ip_ranges")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        return pulumi.get(self, "self_link")


class AwaitableGetSubnetworkResult(GetSubnetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubnetworkResult(
            description=self.description,
            gateway_address=self.gateway_address,
            id=self.id,
            ip_cidr_range=self.ip_cidr_range,
            name=self.name,
            network=self.network,
            private_ip_google_access=self.private_ip_google_access,
            project=self.project,
            region=self.region,
            secondary_ip_ranges=self.secondary_ip_ranges,
            self_link=self.self_link)


def get_subnetwork(name: Optional[str] = None,
                   project: Optional[str] = None,
                   region: Optional[str] = None,
                   self_link: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubnetworkResult:
    """
    Get a subnetwork within GCE from its name and region.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_subnetwork = gcp.compute.get_subnetwork(name="default-us-east1",
        region="us-east1")
    ```


    :param str name: The name of the subnetwork. One of `name` or `self_link`
           must be specified.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region this subnetwork has been created in. If
           unspecified, this defaults to the region configured in the provider.
    :param str self_link: The self link of the subnetwork. If `self_link` is
           specified, `name`, `project`, and `region` are ignored.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    __args__['selfLink'] = self_link
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getSubnetwork:getSubnetwork', __args__, opts=opts, typ=GetSubnetworkResult).value

    return AwaitableGetSubnetworkResult(
        description=__ret__.description,
        gateway_address=__ret__.gateway_address,
        id=__ret__.id,
        ip_cidr_range=__ret__.ip_cidr_range,
        name=__ret__.name,
        network=__ret__.network,
        private_ip_google_access=__ret__.private_ip_google_access,
        project=__ret__.project,
        region=__ret__.region,
        secondary_ip_ranges=__ret__.secondary_ip_ranges,
        self_link=__ret__.self_link)
