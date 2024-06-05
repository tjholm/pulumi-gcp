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
    'GetBackupPlanIamPolicyResult',
    'AwaitableGetBackupPlanIamPolicyResult',
    'get_backup_plan_iam_policy',
    'get_backup_plan_iam_policy_output',
]

@pulumi.output_type
class GetBackupPlanIamPolicyResult:
    """
    A collection of values returned by getBackupPlanIamPolicy.
    """
    def __init__(__self__, etag=None, id=None, location=None, name=None, policy_data=None, project=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        (Required only by `gkebackup.BackupPlanIamPolicy`) The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")


class AwaitableGetBackupPlanIamPolicyResult(GetBackupPlanIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupPlanIamPolicyResult(
            etag=self.etag,
            id=self.id,
            location=self.location,
            name=self.name,
            policy_data=self.policy_data,
            project=self.project)


def get_backup_plan_iam_policy(location: Optional[str] = None,
                               name: Optional[str] = None,
                               project: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupPlanIamPolicyResult:
    """
    Retrieves the current IAM policy data for backupplan

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.gkebackup.get_backup_plan_iam_policy(project=basic["project"],
        location=basic["location"],
        name=basic["name"])
    ```


    :param str location: The region of the Backup Plan.
           Used to find the parent resource to bind the IAM policy to. If not specified,
           the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
           location is specified, it is taken from the provider configuration.
    :param str name: Used to find the parent resource to bind the IAM policy to
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:gkebackup/getBackupPlanIamPolicy:getBackupPlanIamPolicy', __args__, opts=opts, typ=GetBackupPlanIamPolicyResult).value

    return AwaitableGetBackupPlanIamPolicyResult(
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        policy_data=pulumi.get(__ret__, 'policy_data'),
        project=pulumi.get(__ret__, 'project'))


@_utilities.lift_output_func(get_backup_plan_iam_policy)
def get_backup_plan_iam_policy_output(location: Optional[pulumi.Input[Optional[str]]] = None,
                                      name: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupPlanIamPolicyResult]:
    """
    Retrieves the current IAM policy data for backupplan

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.gkebackup.get_backup_plan_iam_policy(project=basic["project"],
        location=basic["location"],
        name=basic["name"])
    ```


    :param str location: The region of the Backup Plan.
           Used to find the parent resource to bind the IAM policy to. If not specified,
           the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
           location is specified, it is taken from the provider configuration.
    :param str name: Used to find the parent resource to bind the IAM policy to
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    ...
