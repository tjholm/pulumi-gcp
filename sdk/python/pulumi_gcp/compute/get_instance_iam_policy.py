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
    'GetInstanceIamPolicyResult',
    'AwaitableGetInstanceIamPolicyResult',
    'get_instance_iam_policy',
    'get_instance_iam_policy_output',
]

@pulumi.output_type
class GetInstanceIamPolicyResult:
    """
    A collection of values returned by getInstanceIamPolicy.
    """
    def __init__(__self__, etag=None, id=None, instance_name=None, policy_data=None, project=None, zone=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

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
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        (Required only by `compute.InstanceIAMPolicy`) The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetInstanceIamPolicyResult(GetInstanceIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceIamPolicyResult(
            etag=self.etag,
            id=self.id,
            instance_name=self.instance_name,
            policy_data=self.policy_data,
            project=self.project,
            zone=self.zone)


def get_instance_iam_policy(instance_name: Optional[str] = None,
                            project: Optional[str] = None,
                            zone: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceIamPolicyResult:
    """
    Retrieves the current IAM policy data for instance

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.compute.get_instance_iam_policy(project=google_compute_instance["default"]["project"],
        zone=google_compute_instance["default"]["zone"],
        instance_name=google_compute_instance["default"]["name"])
    ```


    :param str instance_name: Used to find the parent resource to bind the IAM policy to
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    :param str zone: A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
           the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
           zone is specified, it is taken from the provider configuration.
    """
    __args__ = dict()
    __args__['instanceName'] = instance_name
    __args__['project'] = project
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:compute/getInstanceIamPolicy:getInstanceIamPolicy', __args__, opts=opts, typ=GetInstanceIamPolicyResult).value

    return AwaitableGetInstanceIamPolicyResult(
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        instance_name=pulumi.get(__ret__, 'instance_name'),
        policy_data=pulumi.get(__ret__, 'policy_data'),
        project=pulumi.get(__ret__, 'project'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_instance_iam_policy)
def get_instance_iam_policy_output(instance_name: Optional[pulumi.Input[str]] = None,
                                   project: Optional[pulumi.Input[Optional[str]]] = None,
                                   zone: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceIamPolicyResult]:
    """
    Retrieves the current IAM policy data for instance

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.compute.get_instance_iam_policy(project=google_compute_instance["default"]["project"],
        zone=google_compute_instance["default"]["zone"],
        instance_name=google_compute_instance["default"]["name"])
    ```


    :param str instance_name: Used to find the parent resource to bind the IAM policy to
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    :param str zone: A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
           the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
           zone is specified, it is taken from the provider configuration.
    """
    ...
