# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetKMSKeyRingResult',
    'AwaitableGetKMSKeyRingResult',
    'get_kms_key_ring',
]

@pulumi.output_type
class GetKMSKeyRingResult:
    """
    A collection of values returned by getKMSKeyRing.
    """
    def __init__(__self__, id=None, location=None, name=None, project=None, self_link=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The self link of the created KeyRing. Its format is `projects/{projectId}/locations/{location}/keyRings/{keyRingName}`.
        """
        return pulumi.get(self, "self_link")


class AwaitableGetKMSKeyRingResult(GetKMSKeyRingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSKeyRingResult(
            id=self.id,
            location=self.location,
            name=self.name,
            project=self.project,
            self_link=self.self_link)


def get_kms_key_ring(location: Optional[str] = None,
                     name: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKMSKeyRingResult:
    """
    Provides access to Google Cloud Platform KMS KeyRing. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_ring)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings).

    A KeyRing is a grouping of CryptoKeys for organizational purposes. A KeyRing belongs to a Google Cloud Platform Project
    and resides in a specific location.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_key_ring = gcp.kms.get_kms_key_ring(location="us-central1",
        name="my-key-ring")
    ```


    :param str location: The Google Cloud Platform location for the KeyRing.
           A full list of valid locations can be found by running `gcloud kms locations list`.
    :param str name: The KeyRing's name.
           A KeyRing name must exist within the provided location and match the regular expression `[a-zA-Z0-9_-]{1,63}`
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSKeyRing:getKMSKeyRing', __args__, opts=opts, typ=GetKMSKeyRingResult).value

    return AwaitableGetKMSKeyRingResult(
        id=__ret__.id,
        location=__ret__.location,
        name=__ret__.name,
        project=__ret__.project,
        self_link=__ret__.self_link)
