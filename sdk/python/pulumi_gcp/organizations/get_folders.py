# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetFoldersResult',
    'AwaitableGetFoldersResult',
    'get_folders',
    'get_folders_output',
]

@pulumi.output_type
class GetFoldersResult:
    """
    A collection of values returned by getFolders.
    """
    def __init__(__self__, folders=None, id=None, parent_id=None):
        if folders and not isinstance(folders, list):
            raise TypeError("Expected argument 'folders' to be a list")
        pulumi.set(__self__, "folders", folders)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)

    @property
    @pulumi.getter
    def folders(self) -> Sequence['outputs.GetFoldersFolderResult']:
        """
        A list of projects matching the provided filter. Structure is defined below.
        """
        return pulumi.get(self, "folders")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> str:
        return pulumi.get(self, "parent_id")


class AwaitableGetFoldersResult(GetFoldersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFoldersResult(
            folders=self.folders,
            id=self.id,
            parent_id=self.parent_id)


def get_folders(parent_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFoldersResult:
    """
    Retrieve information about a set of folders based on a parent ID. See the
    [REST API](https://cloud.google.com/resource-manager/reference/rest/v3/folders/list)
    for more details.

    ## Example Usage
    ### Searching For Folders At The Root Of An Org

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_org_folders = gcp.organizations.get_folders(parent_id=f"organizations/{var['organization_id']}")
    first_folder = gcp.organizations.get_folder(folder=my_org_folders.folders[0].name)
    ```


    :param str parent_id: A string parent as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v3/folders/list#query-parameters).
    """
    __args__ = dict()
    __args__['parentId'] = parent_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:organizations/getFolders:getFolders', __args__, opts=opts, typ=GetFoldersResult).value

    return AwaitableGetFoldersResult(
        folders=pulumi.get(__ret__, 'folders'),
        id=pulumi.get(__ret__, 'id'),
        parent_id=pulumi.get(__ret__, 'parent_id'))


@_utilities.lift_output_func(get_folders)
def get_folders_output(parent_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFoldersResult]:
    """
    Retrieve information about a set of folders based on a parent ID. See the
    [REST API](https://cloud.google.com/resource-manager/reference/rest/v3/folders/list)
    for more details.

    ## Example Usage
    ### Searching For Folders At The Root Of An Org

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_org_folders = gcp.organizations.get_folders(parent_id=f"organizations/{var['organization_id']}")
    first_folder = gcp.organizations.get_folder(folder=my_org_folders.folders[0].name)
    ```


    :param str parent_id: A string parent as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v3/folders/list#query-parameters).
    """
    ...
