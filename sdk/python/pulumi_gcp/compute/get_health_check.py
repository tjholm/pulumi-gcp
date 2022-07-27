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
    'GetHealthCheckResult',
    'AwaitableGetHealthCheckResult',
    'get_health_check',
    'get_health_check_output',
]

@pulumi.output_type
class GetHealthCheckResult:
    """
    A collection of values returned by getHealthCheck.
    """
    def __init__(__self__, check_interval_sec=None, creation_timestamp=None, description=None, grpc_health_checks=None, healthy_threshold=None, http2_health_checks=None, http_health_checks=None, https_health_checks=None, id=None, log_configs=None, name=None, project=None, self_link=None, ssl_health_checks=None, tcp_health_checks=None, timeout_sec=None, type=None, unhealthy_threshold=None):
        if check_interval_sec and not isinstance(check_interval_sec, int):
            raise TypeError("Expected argument 'check_interval_sec' to be a int")
        pulumi.set(__self__, "check_interval_sec", check_interval_sec)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if grpc_health_checks and not isinstance(grpc_health_checks, list):
            raise TypeError("Expected argument 'grpc_health_checks' to be a list")
        pulumi.set(__self__, "grpc_health_checks", grpc_health_checks)
        if healthy_threshold and not isinstance(healthy_threshold, int):
            raise TypeError("Expected argument 'healthy_threshold' to be a int")
        pulumi.set(__self__, "healthy_threshold", healthy_threshold)
        if http2_health_checks and not isinstance(http2_health_checks, list):
            raise TypeError("Expected argument 'http2_health_checks' to be a list")
        pulumi.set(__self__, "http2_health_checks", http2_health_checks)
        if http_health_checks and not isinstance(http_health_checks, list):
            raise TypeError("Expected argument 'http_health_checks' to be a list")
        pulumi.set(__self__, "http_health_checks", http_health_checks)
        if https_health_checks and not isinstance(https_health_checks, list):
            raise TypeError("Expected argument 'https_health_checks' to be a list")
        pulumi.set(__self__, "https_health_checks", https_health_checks)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_configs and not isinstance(log_configs, list):
            raise TypeError("Expected argument 'log_configs' to be a list")
        pulumi.set(__self__, "log_configs", log_configs)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if ssl_health_checks and not isinstance(ssl_health_checks, list):
            raise TypeError("Expected argument 'ssl_health_checks' to be a list")
        pulumi.set(__self__, "ssl_health_checks", ssl_health_checks)
        if tcp_health_checks and not isinstance(tcp_health_checks, list):
            raise TypeError("Expected argument 'tcp_health_checks' to be a list")
        pulumi.set(__self__, "tcp_health_checks", tcp_health_checks)
        if timeout_sec and not isinstance(timeout_sec, int):
            raise TypeError("Expected argument 'timeout_sec' to be a int")
        pulumi.set(__self__, "timeout_sec", timeout_sec)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unhealthy_threshold and not isinstance(unhealthy_threshold, int):
            raise TypeError("Expected argument 'unhealthy_threshold' to be a int")
        pulumi.set(__self__, "unhealthy_threshold", unhealthy_threshold)

    @property
    @pulumi.getter(name="checkIntervalSec")
    def check_interval_sec(self) -> int:
        return pulumi.get(self, "check_interval_sec")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="grpcHealthChecks")
    def grpc_health_checks(self) -> Sequence['outputs.GetHealthCheckGrpcHealthCheckResult']:
        return pulumi.get(self, "grpc_health_checks")

    @property
    @pulumi.getter(name="healthyThreshold")
    def healthy_threshold(self) -> int:
        return pulumi.get(self, "healthy_threshold")

    @property
    @pulumi.getter(name="http2HealthChecks")
    def http2_health_checks(self) -> Sequence['outputs.GetHealthCheckHttp2HealthCheckResult']:
        return pulumi.get(self, "http2_health_checks")

    @property
    @pulumi.getter(name="httpHealthChecks")
    def http_health_checks(self) -> Sequence['outputs.GetHealthCheckHttpHealthCheckResult']:
        return pulumi.get(self, "http_health_checks")

    @property
    @pulumi.getter(name="httpsHealthChecks")
    def https_health_checks(self) -> Sequence['outputs.GetHealthCheckHttpsHealthCheckResult']:
        return pulumi.get(self, "https_health_checks")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logConfigs")
    def log_configs(self) -> Sequence['outputs.GetHealthCheckLogConfigResult']:
        return pulumi.get(self, "log_configs")

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
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sslHealthChecks")
    def ssl_health_checks(self) -> Sequence['outputs.GetHealthCheckSslHealthCheckResult']:
        return pulumi.get(self, "ssl_health_checks")

    @property
    @pulumi.getter(name="tcpHealthChecks")
    def tcp_health_checks(self) -> Sequence['outputs.GetHealthCheckTcpHealthCheckResult']:
        return pulumi.get(self, "tcp_health_checks")

    @property
    @pulumi.getter(name="timeoutSec")
    def timeout_sec(self) -> int:
        return pulumi.get(self, "timeout_sec")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="unhealthyThreshold")
    def unhealthy_threshold(self) -> int:
        return pulumi.get(self, "unhealthy_threshold")


class AwaitableGetHealthCheckResult(GetHealthCheckResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHealthCheckResult(
            check_interval_sec=self.check_interval_sec,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            grpc_health_checks=self.grpc_health_checks,
            healthy_threshold=self.healthy_threshold,
            http2_health_checks=self.http2_health_checks,
            http_health_checks=self.http_health_checks,
            https_health_checks=self.https_health_checks,
            id=self.id,
            log_configs=self.log_configs,
            name=self.name,
            project=self.project,
            self_link=self.self_link,
            ssl_health_checks=self.ssl_health_checks,
            tcp_health_checks=self.tcp_health_checks,
            timeout_sec=self.timeout_sec,
            type=self.type,
            unhealthy_threshold=self.unhealthy_threshold)


def get_health_check(name: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHealthCheckResult:
    """
    Get information about a HealthCheck.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    health_check = gcp.compute.get_health_check(name="my-hc")
    ```


    :param str name: Name of the resource.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:compute/getHealthCheck:getHealthCheck', __args__, opts=opts, typ=GetHealthCheckResult).value

    return AwaitableGetHealthCheckResult(
        check_interval_sec=__ret__.check_interval_sec,
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        grpc_health_checks=__ret__.grpc_health_checks,
        healthy_threshold=__ret__.healthy_threshold,
        http2_health_checks=__ret__.http2_health_checks,
        http_health_checks=__ret__.http_health_checks,
        https_health_checks=__ret__.https_health_checks,
        id=__ret__.id,
        log_configs=__ret__.log_configs,
        name=__ret__.name,
        project=__ret__.project,
        self_link=__ret__.self_link,
        ssl_health_checks=__ret__.ssl_health_checks,
        tcp_health_checks=__ret__.tcp_health_checks,
        timeout_sec=__ret__.timeout_sec,
        type=__ret__.type,
        unhealthy_threshold=__ret__.unhealthy_threshold)


@_utilities.lift_output_func(get_health_check)
def get_health_check_output(name: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHealthCheckResult]:
    """
    Get information about a HealthCheck.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    health_check = gcp.compute.get_health_check(name="my-hc")
    ```


    :param str name: Name of the resource.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    ...
