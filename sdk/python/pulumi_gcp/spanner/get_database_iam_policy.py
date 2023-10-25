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
    'GetDatabaseIamPolicyResult',
    'AwaitableGetDatabaseIamPolicyResult',
    'get_database_iam_policy',
    'get_database_iam_policy_output',
]

@pulumi.output_type
class GetDatabaseIamPolicyResult:
    """
    A collection of values returned by getDatabaseIamPolicy.
    """
    def __init__(__self__, database=None, etag=None, id=None, instance=None, policy_data=None, project=None):
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance and not isinstance(instance, str):
            raise TypeError("Expected argument 'instance' to be a str")
        pulumi.set(__self__, "instance", instance)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def database(self) -> str:
        return pulumi.get(self, "database")

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
    def instance(self) -> str:
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        (Computed) The policy data
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")


class AwaitableGetDatabaseIamPolicyResult(GetDatabaseIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseIamPolicyResult(
            database=self.database,
            etag=self.etag,
            id=self.id,
            instance=self.instance,
            policy_data=self.policy_data,
            project=self.project)


def get_database_iam_policy(database: Optional[str] = None,
                            instance: Optional[str] = None,
                            project: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseIamPolicyResult:
    """
    Retrieves the current IAM policy data for a Spanner database.

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    foo = gcp.spanner.get_database_iam_policy(project=google_spanner_database["database"]["project"],
        database=google_spanner_database["database"]["name"],
        instance=google_spanner_database["database"]["instance"])
    ```


    :param str database: The name of the Spanner database.
    :param str instance: The name of the Spanner instance the database belongs to.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['database'] = database
    __args__['instance'] = instance
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:spanner/getDatabaseIamPolicy:getDatabaseIamPolicy', __args__, opts=opts, typ=GetDatabaseIamPolicyResult).value

    return AwaitableGetDatabaseIamPolicyResult(
        database=pulumi.get(__ret__, 'database'),
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        instance=pulumi.get(__ret__, 'instance'),
        policy_data=pulumi.get(__ret__, 'policy_data'),
        project=pulumi.get(__ret__, 'project'))


@_utilities.lift_output_func(get_database_iam_policy)
def get_database_iam_policy_output(database: Optional[pulumi.Input[str]] = None,
                                   instance: Optional[pulumi.Input[str]] = None,
                                   project: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseIamPolicyResult]:
    """
    Retrieves the current IAM policy data for a Spanner database.

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    foo = gcp.spanner.get_database_iam_policy(project=google_spanner_database["database"]["project"],
        database=google_spanner_database["database"]["name"],
        instance=google_spanner_database["database"]["instance"])
    ```


    :param str database: The name of the Spanner database.
    :param str instance: The name of the Spanner instance the database belongs to.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    ...
