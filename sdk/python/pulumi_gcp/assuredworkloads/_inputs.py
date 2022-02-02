# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'WorkloadKmsSettingsArgs',
    'WorkloadResourceArgs',
    'WorkloadResourceSettingArgs',
]

@pulumi.input_type
class WorkloadKmsSettingsArgs:
    def __init__(__self__, *,
                 next_rotation_time: pulumi.Input[str],
                 rotation_period: pulumi.Input[str]):
        """
        :param pulumi.Input[str] next_rotation_time: Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
        :param pulumi.Input[str] rotation_period: Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
        """
        pulumi.set(__self__, "next_rotation_time", next_rotation_time)
        pulumi.set(__self__, "rotation_period", rotation_period)

    @property
    @pulumi.getter(name="nextRotationTime")
    def next_rotation_time(self) -> pulumi.Input[str]:
        """
        Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
        """
        return pulumi.get(self, "next_rotation_time")

    @next_rotation_time.setter
    def next_rotation_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_rotation_time", value)

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Input[str]:
        """
        Required. Input only. Immutable. will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
        """
        return pulumi.get(self, "rotation_period")

    @rotation_period.setter
    def rotation_period(self, value: pulumi.Input[str]):
        pulumi.set(self, "rotation_period", value)


@pulumi.input_type
class WorkloadResourceArgs:
    def __init__(__self__, *,
                 resource_id: Optional[pulumi.Input[int]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] resource_id: Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
        :param pulumi.Input[str] resource_type: Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[int]]:
        """
        Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


@pulumi.input_type
class WorkloadResourceSettingArgs:
    def __init__(__self__, *,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 resource_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] resource_id: Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
        :param pulumi.Input[str] resource_type: Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_type", value)


