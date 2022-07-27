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
    'GetBucketResult',
    'AwaitableGetBucketResult',
    'get_bucket',
    'get_bucket_output',
]

@pulumi.output_type
class GetBucketResult:
    """
    A collection of values returned by getBucket.
    """
    def __init__(__self__, cors=None, default_event_based_hold=None, encryptions=None, force_destroy=None, id=None, labels=None, lifecycle_rules=None, location=None, loggings=None, name=None, project=None, public_access_prevention=None, requester_pays=None, retention_policies=None, self_link=None, storage_class=None, uniform_bucket_level_access=None, url=None, versionings=None, websites=None):
        if cors and not isinstance(cors, list):
            raise TypeError("Expected argument 'cors' to be a list")
        pulumi.set(__self__, "cors", cors)
        if default_event_based_hold and not isinstance(default_event_based_hold, bool):
            raise TypeError("Expected argument 'default_event_based_hold' to be a bool")
        pulumi.set(__self__, "default_event_based_hold", default_event_based_hold)
        if encryptions and not isinstance(encryptions, list):
            raise TypeError("Expected argument 'encryptions' to be a list")
        pulumi.set(__self__, "encryptions", encryptions)
        if force_destroy and not isinstance(force_destroy, bool):
            raise TypeError("Expected argument 'force_destroy' to be a bool")
        pulumi.set(__self__, "force_destroy", force_destroy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if lifecycle_rules and not isinstance(lifecycle_rules, list):
            raise TypeError("Expected argument 'lifecycle_rules' to be a list")
        pulumi.set(__self__, "lifecycle_rules", lifecycle_rules)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if loggings and not isinstance(loggings, list):
            raise TypeError("Expected argument 'loggings' to be a list")
        pulumi.set(__self__, "loggings", loggings)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if public_access_prevention and not isinstance(public_access_prevention, str):
            raise TypeError("Expected argument 'public_access_prevention' to be a str")
        pulumi.set(__self__, "public_access_prevention", public_access_prevention)
        if requester_pays and not isinstance(requester_pays, bool):
            raise TypeError("Expected argument 'requester_pays' to be a bool")
        pulumi.set(__self__, "requester_pays", requester_pays)
        if retention_policies and not isinstance(retention_policies, list):
            raise TypeError("Expected argument 'retention_policies' to be a list")
        pulumi.set(__self__, "retention_policies", retention_policies)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if storage_class and not isinstance(storage_class, str):
            raise TypeError("Expected argument 'storage_class' to be a str")
        pulumi.set(__self__, "storage_class", storage_class)
        if uniform_bucket_level_access and not isinstance(uniform_bucket_level_access, bool):
            raise TypeError("Expected argument 'uniform_bucket_level_access' to be a bool")
        pulumi.set(__self__, "uniform_bucket_level_access", uniform_bucket_level_access)
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        pulumi.set(__self__, "url", url)
        if versionings and not isinstance(versionings, list):
            raise TypeError("Expected argument 'versionings' to be a list")
        pulumi.set(__self__, "versionings", versionings)
        if websites and not isinstance(websites, list):
            raise TypeError("Expected argument 'websites' to be a list")
        pulumi.set(__self__, "websites", websites)

    @property
    @pulumi.getter
    def cors(self) -> Sequence['outputs.GetBucketCorResult']:
        return pulumi.get(self, "cors")

    @property
    @pulumi.getter(name="defaultEventBasedHold")
    def default_event_based_hold(self) -> bool:
        return pulumi.get(self, "default_event_based_hold")

    @property
    @pulumi.getter
    def encryptions(self) -> Sequence['outputs.GetBucketEncryptionResult']:
        return pulumi.get(self, "encryptions")

    @property
    @pulumi.getter(name="forceDestroy")
    def force_destroy(self) -> bool:
        return pulumi.get(self, "force_destroy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lifecycleRules")
    def lifecycle_rules(self) -> Sequence['outputs.GetBucketLifecycleRuleResult']:
        return pulumi.get(self, "lifecycle_rules")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def loggings(self) -> Sequence['outputs.GetBucketLoggingResult']:
        return pulumi.get(self, "loggings")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="publicAccessPrevention")
    def public_access_prevention(self) -> str:
        return pulumi.get(self, "public_access_prevention")

    @property
    @pulumi.getter(name="requesterPays")
    def requester_pays(self) -> bool:
        return pulumi.get(self, "requester_pays")

    @property
    @pulumi.getter(name="retentionPolicies")
    def retention_policies(self) -> Sequence['outputs.GetBucketRetentionPolicyResult']:
        return pulumi.get(self, "retention_policies")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> str:
        return pulumi.get(self, "storage_class")

    @property
    @pulumi.getter(name="uniformBucketLevelAccess")
    def uniform_bucket_level_access(self) -> bool:
        return pulumi.get(self, "uniform_bucket_level_access")

    @property
    @pulumi.getter
    def url(self) -> str:
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def versionings(self) -> Sequence['outputs.GetBucketVersioningResult']:
        return pulumi.get(self, "versionings")

    @property
    @pulumi.getter
    def websites(self) -> Sequence['outputs.GetBucketWebsiteResult']:
        return pulumi.get(self, "websites")


class AwaitableGetBucketResult(GetBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketResult(
            cors=self.cors,
            default_event_based_hold=self.default_event_based_hold,
            encryptions=self.encryptions,
            force_destroy=self.force_destroy,
            id=self.id,
            labels=self.labels,
            lifecycle_rules=self.lifecycle_rules,
            location=self.location,
            loggings=self.loggings,
            name=self.name,
            project=self.project,
            public_access_prevention=self.public_access_prevention,
            requester_pays=self.requester_pays,
            retention_policies=self.retention_policies,
            self_link=self.self_link,
            storage_class=self.storage_class,
            uniform_bucket_level_access=self.uniform_bucket_level_access,
            url=self.url,
            versionings=self.versionings,
            websites=self.websites)


def get_bucket(name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBucketResult:
    """
    Gets an existing bucket in Google Cloud Storage service (GCS).
    See [the official documentation](https://cloud.google.com/storage/docs/key-terms#buckets)
    and
    [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_bucket = gcp.storage.get_bucket(name="my-bucket")
    ```


    :param str name: The name of the bucket.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:storage/getBucket:getBucket', __args__, opts=opts, typ=GetBucketResult).value

    return AwaitableGetBucketResult(
        cors=__ret__.cors,
        default_event_based_hold=__ret__.default_event_based_hold,
        encryptions=__ret__.encryptions,
        force_destroy=__ret__.force_destroy,
        id=__ret__.id,
        labels=__ret__.labels,
        lifecycle_rules=__ret__.lifecycle_rules,
        location=__ret__.location,
        loggings=__ret__.loggings,
        name=__ret__.name,
        project=__ret__.project,
        public_access_prevention=__ret__.public_access_prevention,
        requester_pays=__ret__.requester_pays,
        retention_policies=__ret__.retention_policies,
        self_link=__ret__.self_link,
        storage_class=__ret__.storage_class,
        uniform_bucket_level_access=__ret__.uniform_bucket_level_access,
        url=__ret__.url,
        versionings=__ret__.versionings,
        websites=__ret__.websites)


@_utilities.lift_output_func(get_bucket)
def get_bucket_output(name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBucketResult]:
    """
    Gets an existing bucket in Google Cloud Storage service (GCS).
    See [the official documentation](https://cloud.google.com/storage/docs/key-terms#buckets)
    and
    [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_bucket = gcp.storage.get_bucket(name="my-bucket")
    ```


    :param str name: The name of the bucket.
    """
    ...
