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
    'GetConfigResult',
    'AwaitableGetConfigResult',
    'get_config',
    'get_config_output',
]

@pulumi.output_type
class GetConfigResult:
    """
    A collection of values returned by getConfig.
    """
    def __init__(__self__, description=None, id=None, name=None, project=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def description(self) -> str:
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
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")


class AwaitableGetConfigResult(GetConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigResult(
            description=self.description,
            id=self.id,
            name=self.name,
            project=self.project)


def get_config(name: Optional[str] = None,
               project: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    run_service = gcp.runtimeconfig.get_config(name="my-service")
    ```


    :param str name: The name of the Runtime Configurator configuration.
           
           - - -
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:runtimeconfig/getConfig:getConfig', __args__, opts=opts, typ=GetConfigResult).value

    return AwaitableGetConfigResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project=pulumi.get(__ret__, 'project'))


@_utilities.lift_output_func(get_config)
def get_config_output(name: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    run_service = gcp.runtimeconfig.get_config(name="my-service")
    ```


    :param str name: The name of the Runtime Configurator configuration.
           
           - - -
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    ...
