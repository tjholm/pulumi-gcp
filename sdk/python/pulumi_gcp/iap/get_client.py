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
    'GetClientResult',
    'AwaitableGetClientResult',
    'get_client',
    'get_client_output',
]

@pulumi.output_type
class GetClientResult:
    """
    A collection of values returned by getClient.
    """
    def __init__(__self__, brand=None, client_id=None, display_name=None, id=None, secret=None):
        if brand and not isinstance(brand, str):
            raise TypeError("Expected argument 'brand' to be a str")
        pulumi.set(__self__, "brand", brand)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if secret and not isinstance(secret, str):
            raise TypeError("Expected argument 'secret' to be a str")
        pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter
    def brand(self) -> str:
        return pulumi.get(self, "brand")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def secret(self) -> str:
        return pulumi.get(self, "secret")


class AwaitableGetClientResult(GetClientResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientResult(
            brand=self.brand,
            client_id=self.client_id,
            display_name=self.display_name,
            id=self.id,
            secret=self.secret)


def get_client(brand: Optional[str] = None,
               client_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientResult:
    """
    Get info about a Google Cloud IAP Client.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    project = gcp.organizations.get_project(project_id="foobar")
    project_client = gcp.iap.get_client(brand=f"projects/{project.number}/brands/[BRAND_NUMBER]",
        client_id=foo["apps"]["googleusercontent"]["com"])
    ```


    :param str brand: The name of the brand.
    :param str client_id: The client_id of the brand.
    """
    __args__ = dict()
    __args__['brand'] = brand
    __args__['clientId'] = client_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:iap/getClient:getClient', __args__, opts=opts, typ=GetClientResult).value

    return AwaitableGetClientResult(
        brand=__ret__.brand,
        client_id=__ret__.client_id,
        display_name=__ret__.display_name,
        id=__ret__.id,
        secret=__ret__.secret)


@_utilities.lift_output_func(get_client)
def get_client_output(brand: Optional[pulumi.Input[str]] = None,
                      client_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientResult]:
    """
    Get info about a Google Cloud IAP Client.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    project = gcp.organizations.get_project(project_id="foobar")
    project_client = gcp.iap.get_client(brand=f"projects/{project.number}/brands/[BRAND_NUMBER]",
        client_id=foo["apps"]["googleusercontent"]["com"])
    ```


    :param str brand: The name of the brand.
    :param str client_id: The client_id of the brand.
    """
    ...
