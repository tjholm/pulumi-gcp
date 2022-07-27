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
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, filter=None, id=None, projects=None):
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)

    @property
    @pulumi.getter
    def filter(self) -> str:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def projects(self) -> Sequence['outputs.GetProjectProjectResult']:
        """
        A list of projects matching the provided filter. Structure is defined below.
        """
        return pulumi.get(self, "projects")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            filter=self.filter,
            id=self.id,
            projects=self.projects)


def get_project(filter: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Retrieve information about a set of projects based on a filter. See the
    [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
    for more details.

    ## Example Usage
    ### Searching For Projects About To Be Deleted In An Org

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_org_projects = gcp.projects.get_project(filter="parent.id:012345678910 lifecycleState:DELETE_REQUESTED")
    deletion_candidate = gcp.organizations.get_project(project_id=my_org_projects.projects[0].project_id)
    ```


    :param str filter: A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters).
    """
    __args__ = dict()
    __args__['filter'] = filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:projects/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        filter=__ret__.filter,
        id=__ret__.id,
        projects=__ret__.projects)


@_utilities.lift_output_func(get_project)
def get_project_output(filter: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Retrieve information about a set of projects based on a filter. See the
    [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
    for more details.

    ## Example Usage
    ### Searching For Projects About To Be Deleted In An Org

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_org_projects = gcp.projects.get_project(filter="parent.id:012345678910 lifecycleState:DELETE_REQUESTED")
    deletion_candidate = gcp.organizations.get_project(project_id=my_org_projects.projects[0].project_id)
    ```


    :param str filter: A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters).
    """
    ...
