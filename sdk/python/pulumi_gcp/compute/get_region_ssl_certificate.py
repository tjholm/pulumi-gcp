# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetRegionSslCertificateResult',
    'AwaitableGetRegionSslCertificateResult',
    'get_region_ssl_certificate',
    'get_region_ssl_certificate_output',
]

@pulumi.output_type
class GetRegionSslCertificateResult:
    """
    A collection of values returned by getRegionSslCertificate.
    """
    def __init__(__self__, certificate=None, certificate_id=None, creation_timestamp=None, description=None, expire_time=None, id=None, name=None, name_prefix=None, private_key=None, project=None, region=None, self_link=None):
        if certificate and not isinstance(certificate, str):
            raise TypeError("Expected argument 'certificate' to be a str")
        pulumi.set(__self__, "certificate", certificate)
        if certificate_id and not isinstance(certificate_id, int):
            raise TypeError("Expected argument 'certificate_id' to be a int")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)
        if private_key and not isinstance(private_key, str):
            raise TypeError("Expected argument 'private_key' to be a str")
        pulumi.set(__self__, "private_key", private_key)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter
    def certificate(self) -> str:
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> int:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        return pulumi.get(self, "expire_time")

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
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> str:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> str:
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        return pulumi.get(self, "self_link")


class AwaitableGetRegionSslCertificateResult(GetRegionSslCertificateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionSslCertificateResult(
            certificate=self.certificate,
            certificate_id=self.certificate_id,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            expire_time=self.expire_time,
            id=self.id,
            name=self.name,
            name_prefix=self.name_prefix,
            private_key=self.private_key,
            project=self.project,
            region=self.region,
            self_link=self.self_link)


def get_region_ssl_certificate(name: Optional[str] = None,
                               project: Optional[str] = None,
                               region: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionSslCertificateResult:
    """
    Get info about a Region Google Compute SSL Certificate from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_cert = gcp.compute.get_region_ssl_certificate(name="my-cert")
    pulumi.export("certificate", my_cert.certificate)
    pulumi.export("certificateId", my_cert.certificate_id)
    pulumi.export("selfLink", my_cert.self_link)
    ```


    :param str name: The name of the certificate.
           
           - - -
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region in which the resource belongs. If it
           is not provided, the provider region is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:compute/getRegionSslCertificate:getRegionSslCertificate', __args__, opts=opts, typ=GetRegionSslCertificateResult).value

    return AwaitableGetRegionSslCertificateResult(
        certificate=pulumi.get(__ret__, 'certificate'),
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        creation_timestamp=pulumi.get(__ret__, 'creation_timestamp'),
        description=pulumi.get(__ret__, 'description'),
        expire_time=pulumi.get(__ret__, 'expire_time'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_prefix=pulumi.get(__ret__, 'name_prefix'),
        private_key=pulumi.get(__ret__, 'private_key'),
        project=pulumi.get(__ret__, 'project'),
        region=pulumi.get(__ret__, 'region'),
        self_link=pulumi.get(__ret__, 'self_link'))


@_utilities.lift_output_func(get_region_ssl_certificate)
def get_region_ssl_certificate_output(name: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      region: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionSslCertificateResult]:
    """
    Get info about a Region Google Compute SSL Certificate from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_cert = gcp.compute.get_region_ssl_certificate(name="my-cert")
    pulumi.export("certificate", my_cert.certificate)
    pulumi.export("certificateId", my_cert.certificate_id)
    pulumi.export("selfLink", my_cert.self_link)
    ```


    :param str name: The name of the certificate.
           
           - - -
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region in which the resource belongs. If it
           is not provided, the provider region is used.
    """
    ...
