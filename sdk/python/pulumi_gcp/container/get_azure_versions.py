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
    'GetAzureVersionsResult',
    'AwaitableGetAzureVersionsResult',
    'get_azure_versions',
    'get_azure_versions_output',
]

@pulumi.output_type
class GetAzureVersionsResult:
    """
    A collection of values returned by getAzureVersions.
    """
    def __init__(__self__, id=None, location=None, project=None, supported_regions=None, valid_versions=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if supported_regions and not isinstance(supported_regions, list):
            raise TypeError("Expected argument 'supported_regions' to be a list")
        pulumi.set(__self__, "supported_regions", supported_regions)
        if valid_versions and not isinstance(valid_versions, list):
            raise TypeError("Expected argument 'valid_versions' to be a list")
        pulumi.set(__self__, "valid_versions", valid_versions)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="supportedRegions")
    def supported_regions(self) -> Sequence[str]:
        """
        A list of Azure regions that are available for use with this project and GCP location.
        """
        return pulumi.get(self, "supported_regions")

    @property
    @pulumi.getter(name="validVersions")
    def valid_versions(self) -> Sequence[str]:
        """
        A list of versions available for use with this project and location.
        """
        return pulumi.get(self, "valid_versions")


class AwaitableGetAzureVersionsResult(GetAzureVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAzureVersionsResult(
            id=self.id,
            location=self.location,
            project=self.project,
            supported_regions=self.supported_regions,
            valid_versions=self.valid_versions)


def get_azure_versions(location: Optional[str] = None,
                       project: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAzureVersionsResult:
    """
    Provides access to available Kubernetes versions in a location for a given project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    central1b = gcp.container.get_azure_versions(location="us-west1",
        project="my-project")
    pulumi.export("firstAvailableVersion", data["google_container_azure_versions"]["versions"]["valid_versions"])
    ```


    :param str location: The location to list versions for.
    :param str project: ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
           Defaults to the project that the provider is authenticated with.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:container/getAzureVersions:getAzureVersions', __args__, opts=opts, typ=GetAzureVersionsResult).value

    return AwaitableGetAzureVersionsResult(
        id=__ret__.id,
        location=__ret__.location,
        project=__ret__.project,
        supported_regions=__ret__.supported_regions,
        valid_versions=__ret__.valid_versions)


@_utilities.lift_output_func(get_azure_versions)
def get_azure_versions_output(location: Optional[pulumi.Input[Optional[str]]] = None,
                              project: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAzureVersionsResult]:
    """
    Provides access to available Kubernetes versions in a location for a given project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    central1b = gcp.container.get_azure_versions(location="us-west1",
        project="my-project")
    pulumi.export("firstAvailableVersion", data["google_container_azure_versions"]["versions"]["valid_versions"])
    ```


    :param str location: The location to list versions for.
    :param str project: ID of the project to list available cluster versions for. Should match the project the cluster will be deployed to.
           Defaults to the project that the provider is authenticated with.
    """
    ...
