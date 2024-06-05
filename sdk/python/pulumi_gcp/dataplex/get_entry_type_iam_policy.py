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
    'GetEntryTypeIamPolicyResult',
    'AwaitableGetEntryTypeIamPolicyResult',
    'get_entry_type_iam_policy',
    'get_entry_type_iam_policy_output',
]

@pulumi.output_type
class GetEntryTypeIamPolicyResult:
    """
    A collection of values returned by getEntryTypeIamPolicy.
    """
    def __init__(__self__, entry_type_id=None, etag=None, id=None, location=None, policy_data=None, project=None):
        if entry_type_id and not isinstance(entry_type_id, str):
            raise TypeError("Expected argument 'entry_type_id' to be a str")
        pulumi.set(__self__, "entry_type_id", entry_type_id)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="entryTypeId")
    def entry_type_id(self) -> str:
        return pulumi.get(self, "entry_type_id")

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        (Required only by `dataplex.EntryTypeIamPolicy`) The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")


class AwaitableGetEntryTypeIamPolicyResult(GetEntryTypeIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntryTypeIamPolicyResult(
            entry_type_id=self.entry_type_id,
            etag=self.etag,
            id=self.id,
            location=self.location,
            policy_data=self.policy_data,
            project=self.project)


def get_entry_type_iam_policy(entry_type_id: Optional[str] = None,
                              location: Optional[str] = None,
                              project: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntryTypeIamPolicyResult:
    """
    Retrieves the current IAM policy data for entrytype

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.dataplex.get_entry_type_iam_policy(project=test_entry_type_basic["project"],
        location=test_entry_type_basic["location"],
        entry_type_id=test_entry_type_basic["entryTypeId"])
    ```


    :param str location: The location where entry type will be created in.
           Used to find the parent resource to bind the IAM policy to. If not specified,
           the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
           location is specified, it is taken from the provider configuration.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    __args__ = dict()
    __args__['entryTypeId'] = entry_type_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:dataplex/getEntryTypeIamPolicy:getEntryTypeIamPolicy', __args__, opts=opts, typ=GetEntryTypeIamPolicyResult).value

    return AwaitableGetEntryTypeIamPolicyResult(
        entry_type_id=pulumi.get(__ret__, 'entry_type_id'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        policy_data=pulumi.get(__ret__, 'policy_data'),
        project=pulumi.get(__ret__, 'project'))


@_utilities.lift_output_func(get_entry_type_iam_policy)
def get_entry_type_iam_policy_output(entry_type_id: Optional[pulumi.Input[str]] = None,
                                     location: Optional[pulumi.Input[Optional[str]]] = None,
                                     project: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEntryTypeIamPolicyResult]:
    """
    Retrieves the current IAM policy data for entrytype

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.dataplex.get_entry_type_iam_policy(project=test_entry_type_basic["project"],
        location=test_entry_type_basic["location"],
        entry_type_id=test_entry_type_basic["entryTypeId"])
    ```


    :param str location: The location where entry type will be created in.
           Used to find the parent resource to bind the IAM policy to. If not specified,
           the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
           location is specified, it is taken from the provider configuration.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
    """
    ...
