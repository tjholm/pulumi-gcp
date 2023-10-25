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
    'GetTransferProjectServiceAccountResult',
    'AwaitableGetTransferProjectServiceAccountResult',
    'get_transfer_project_service_account',
    'get_transfer_project_service_account_output',
]

@pulumi.output_type
class GetTransferProjectServiceAccountResult:
    """
    A collection of values returned by getTransferProjectServiceAccount.
    """
    def __init__(__self__, email=None, id=None, member=None, project=None, subject_id=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if member and not isinstance(member, str):
            raise TypeError("Expected argument 'member' to be a str")
        pulumi.set(__self__, "member", member)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if subject_id and not isinstance(subject_id, str):
            raise TypeError("Expected argument 'subject_id' to be a str")
        pulumi.set(__self__, "subject_id", subject_id)

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        Email address of the default service account used by Storage Transfer Jobs running in this project.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def member(self) -> str:
        """
        The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
        """
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="subjectId")
    def subject_id(self) -> str:
        """
        Unique identifier for the service account.
        """
        return pulumi.get(self, "subject_id")


class AwaitableGetTransferProjectServiceAccountResult(GetTransferProjectServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransferProjectServiceAccountResult(
            email=self.email,
            id=self.id,
            member=self.member,
            project=self.project,
            subject_id=self.subject_id)


def get_transfer_project_service_account(project: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTransferProjectServiceAccountResult:
    """
    Use this data source to retrieve Storage Transfer service account for this project

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    default = gcp.storage.get_transfer_project_service_account()
    pulumi.export("defaultAccount", default.email)
    ```


    :param str project: The project ID. If it is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:storage/getTransferProjectServiceAccount:getTransferProjectServiceAccount', __args__, opts=opts, typ=GetTransferProjectServiceAccountResult).value

    return AwaitableGetTransferProjectServiceAccountResult(
        email=pulumi.get(__ret__, 'email'),
        id=pulumi.get(__ret__, 'id'),
        member=pulumi.get(__ret__, 'member'),
        project=pulumi.get(__ret__, 'project'),
        subject_id=pulumi.get(__ret__, 'subject_id'))


@_utilities.lift_output_func(get_transfer_project_service_account)
def get_transfer_project_service_account_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTransferProjectServiceAccountResult]:
    """
    Use this data source to retrieve Storage Transfer service account for this project

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    default = gcp.storage.get_transfer_project_service_account()
    pulumi.export("defaultAccount", default.email)
    ```


    :param str project: The project ID. If it is not provided, the provider project is used.
    """
    ...
