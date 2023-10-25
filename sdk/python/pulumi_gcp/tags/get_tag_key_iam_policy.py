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
    'GetTagKeyIamPolicyResult',
    'AwaitableGetTagKeyIamPolicyResult',
    'get_tag_key_iam_policy',
    'get_tag_key_iam_policy_output',
]

@pulumi.output_type
class GetTagKeyIamPolicyResult:
    """
    A collection of values returned by getTagKeyIamPolicy.
    """
    def __init__(__self__, etag=None, id=None, policy_data=None, tag_key=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_data and not isinstance(policy_data, str):
            raise TypeError("Expected argument 'policy_data' to be a str")
        pulumi.set(__self__, "policy_data", policy_data)
        if tag_key and not isinstance(tag_key, str):
            raise TypeError("Expected argument 'tag_key' to be a str")
        pulumi.set(__self__, "tag_key", tag_key)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> str:
        """
        (Required only by `tags.TagKeyIamPolicy`) The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter(name="tagKey")
    def tag_key(self) -> str:
        return pulumi.get(self, "tag_key")


class AwaitableGetTagKeyIamPolicyResult(GetTagKeyIamPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagKeyIamPolicyResult(
            etag=self.etag,
            id=self.id,
            policy_data=self.policy_data,
            tag_key=self.tag_key)


def get_tag_key_iam_policy(tag_key: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagKeyIamPolicyResult:
    """
    Retrieves the current IAM policy data for tagkey

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.tags.get_tag_key_iam_policy(tag_key=google_tags_tag_key["key"]["name"])
    ```


    :param str tag_key: Used to find the parent resource to bind the IAM policy to
    """
    __args__ = dict()
    __args__['tagKey'] = tag_key
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:tags/getTagKeyIamPolicy:getTagKeyIamPolicy', __args__, opts=opts, typ=GetTagKeyIamPolicyResult).value

    return AwaitableGetTagKeyIamPolicyResult(
        etag=pulumi.get(__ret__, 'etag'),
        id=pulumi.get(__ret__, 'id'),
        policy_data=pulumi.get(__ret__, 'policy_data'),
        tag_key=pulumi.get(__ret__, 'tag_key'))


@_utilities.lift_output_func(get_tag_key_iam_policy)
def get_tag_key_iam_policy_output(tag_key: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTagKeyIamPolicyResult]:
    """
    Retrieves the current IAM policy data for tagkey

    ## example

    ```python
    import pulumi
    import pulumi_gcp as gcp

    policy = gcp.tags.get_tag_key_iam_policy(tag_key=google_tags_tag_key["key"]["name"])
    ```


    :param str tag_key: Used to find the parent resource to bind the IAM policy to
    """
    ...
