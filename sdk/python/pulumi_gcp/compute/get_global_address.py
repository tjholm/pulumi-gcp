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
    'GetGlobalAddressResult',
    'AwaitableGetGlobalAddressResult',
    'get_global_address',
    'get_global_address_output',
]

@pulumi.output_type
class GetGlobalAddressResult:
    """
    A collection of values returned by getGlobalAddress.
    """
    def __init__(__self__, address=None, id=None, name=None, project=None, self_link=None, status=None):
        if address and not isinstance(address, str):
            raise TypeError("Expected argument 'address' to be a str")
        pulumi.set(__self__, "address", address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def address(self) -> str:
        """
        The IP of the created resource.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Indicates if the address is used. Possible values are: RESERVED or IN_USE.
        """
        return pulumi.get(self, "status")


class AwaitableGetGlobalAddressResult(GetGlobalAddressResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlobalAddressResult(
            address=self.address,
            id=self.id,
            name=self.name,
            project=self.project,
            self_link=self.self_link,
            status=self.status)


def get_global_address(name: Optional[str] = None,
                       project: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGlobalAddressResult:
    """
    Get the IP address from a static address reserved for a Global Forwarding Rule which are only used for HTTP load balancing. For more information see
    the official [API](https://cloud.google.com/compute/docs/reference/latest/globalAddresses) documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_address = gcp.compute.get_global_address(name="foobar")
    prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
    frontend = gcp.dns.RecordSet("frontend",
        name=prod.dns_name.apply(lambda dns_name: f"lb.{dns_name}"),
        type="A",
        ttl=300,
        managed_zone=prod.name,
        rrdatas=[my_address.address])
    ```


    :param str name: A unique name for the resource, required by GCE.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:compute/getGlobalAddress:getGlobalAddress', __args__, opts=opts, typ=GetGlobalAddressResult).value

    return AwaitableGetGlobalAddressResult(
        address=__ret__.address,
        id=__ret__.id,
        name=__ret__.name,
        project=__ret__.project,
        self_link=__ret__.self_link,
        status=__ret__.status)


@_utilities.lift_output_func(get_global_address)
def get_global_address_output(name: Optional[pulumi.Input[str]] = None,
                              project: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGlobalAddressResult]:
    """
    Get the IP address from a static address reserved for a Global Forwarding Rule which are only used for HTTP load balancing. For more information see
    the official [API](https://cloud.google.com/compute/docs/reference/latest/globalAddresses) documentation.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_address = gcp.compute.get_global_address(name="foobar")
    prod = gcp.dns.ManagedZone("prod", dns_name="prod.mydomain.com.")
    frontend = gcp.dns.RecordSet("frontend",
        name=prod.dns_name.apply(lambda dns_name: f"lb.{dns_name}"),
        type="A",
        ttl=300,
        managed_zone=prod.name,
        rrdatas=[my_address.address])
    ```


    :param str name: A unique name for the resource, required by GCE.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    ...
