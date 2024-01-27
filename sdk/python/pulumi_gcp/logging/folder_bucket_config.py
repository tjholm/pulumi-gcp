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
from ._inputs import *

__all__ = ['FolderBucketConfigArgs', 'FolderBucketConfig']

@pulumi.input_type
class FolderBucketConfigArgs:
    def __init__(__self__, *,
                 bucket_id: pulumi.Input[str],
                 folder: pulumi.Input[str],
                 location: pulumi.Input[str],
                 cmek_settings: Optional[pulumi.Input['FolderBucketConfigCmekSettingsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a FolderBucketConfig resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[str] folder: The parent resource that contains the logging bucket.
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input['FolderBucketConfigCmekSettingsArgs'] cmek_settings: The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
               key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
               updating the log bucket. Changing the KMS key is allowed.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]] index_configs: A list of indexed fields and related configuration data. Structure is documented below.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        pulumi.set(__self__, "bucket_id", bucket_id)
        pulumi.set(__self__, "folder", folder)
        pulumi.set(__self__, "location", location)
        if cmek_settings is not None:
            pulumi.set(__self__, "cmek_settings", cmek_settings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if index_configs is not None:
            pulumi.set(__self__, "index_configs", index_configs)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> pulumi.Input[str]:
        """
        The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        """
        return pulumi.get(self, "bucket_id")

    @bucket_id.setter
    def bucket_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_id", value)

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Input[str]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location of the bucket.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="cmekSettings")
    def cmek_settings(self) -> Optional[pulumi.Input['FolderBucketConfigCmekSettingsArgs']]:
        """
        The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
        key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
        updating the log bucket. Changing the KMS key is allowed.
        """
        return pulumi.get(self, "cmek_settings")

    @cmek_settings.setter
    def cmek_settings(self, value: Optional[pulumi.Input['FolderBucketConfigCmekSettingsArgs']]):
        pulumi.set(self, "cmek_settings", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="indexConfigs")
    def index_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]]]:
        """
        A list of indexed fields and related configuration data. Structure is documented below.
        """
        return pulumi.get(self, "index_configs")

    @index_configs.setter
    def index_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]]]):
        pulumi.set(self, "index_configs", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


@pulumi.input_type
class _FolderBucketConfigState:
    def __init__(__self__, *,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 cmek_settings: Optional[pulumi.Input['FolderBucketConfigCmekSettingsArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]]] = None,
                 lifecycle_state: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering FolderBucketConfig resources.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input['FolderBucketConfigCmekSettingsArgs'] cmek_settings: The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
               key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
               updating the log bucket. Changing the KMS key is allowed.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] folder: The parent resource that contains the logging bucket.
        :param pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]] index_configs: A list of indexed fields and related configuration data. Structure is documented below.
        :param pulumi.Input[str] lifecycle_state: The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] name: The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        if bucket_id is not None:
            pulumi.set(__self__, "bucket_id", bucket_id)
        if cmek_settings is not None:
            pulumi.set(__self__, "cmek_settings", cmek_settings)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if index_configs is not None:
            pulumi.set(__self__, "index_configs", index_configs)
        if lifecycle_state is not None:
            pulumi.set(__self__, "lifecycle_state", lifecycle_state)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        """
        return pulumi.get(self, "bucket_id")

    @bucket_id.setter
    def bucket_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_id", value)

    @property
    @pulumi.getter(name="cmekSettings")
    def cmek_settings(self) -> Optional[pulumi.Input['FolderBucketConfigCmekSettingsArgs']]:
        """
        The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
        key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
        updating the log bucket. Changing the KMS key is allowed.
        """
        return pulumi.get(self, "cmek_settings")

    @cmek_settings.setter
    def cmek_settings(self, value: Optional[pulumi.Input['FolderBucketConfigCmekSettingsArgs']]):
        pulumi.set(self, "cmek_settings", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="indexConfigs")
    def index_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]]]:
        """
        A list of indexed fields and related configuration data. Structure is documented below.
        """
        return pulumi.get(self, "index_configs")

    @index_configs.setter
    def index_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FolderBucketConfigIndexConfigArgs']]]]):
        pulumi.set(self, "index_configs", value)

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> Optional[pulumi.Input[str]]:
        """
        The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        """
        return pulumi.get(self, "lifecycle_state")

    @lifecycle_state.setter
    def lifecycle_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lifecycle_state", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the bucket.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)


class FolderBucketConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 cmek_settings: Optional[pulumi.Input[pulumi.InputType['FolderBucketConfigCmekSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderBucketConfigIndexConfigArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Manages a folder-level logging bucket config. For more information see
        [the official logging documentation](https://cloud.google.com/logging/docs/) and
        [Storing Logs](https://cloud.google.com/logging/docs/storage).

        > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.organizations.Folder("default",
            display_name="some-folder-name",
            parent="organizations/123456789")
        basic = gcp.logging.FolderBucketConfig("basic",
            folder=default.name,
            location="global",
            retention_days=30,
            bucket_id="_Default",
            index_configs={
                "filePath": "jsonPayload.request.status",
                "type": "INDEX_TYPE_STRING",
            })
        ```

        ## Import

        This resource can be imported using the following format* `folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}` When using the `pulumi import` command, this resource can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:logging/folderBucketConfig:FolderBucketConfig default folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[pulumi.InputType['FolderBucketConfigCmekSettingsArgs']] cmek_settings: The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
               key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
               updating the log bucket. Changing the KMS key is allowed.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] folder: The parent resource that contains the logging bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderBucketConfigIndexConfigArgs']]]] index_configs: A list of indexed fields and related configuration data. Structure is documented below.
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderBucketConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a folder-level logging bucket config. For more information see
        [the official logging documentation](https://cloud.google.com/logging/docs/) and
        [Storing Logs](https://cloud.google.com/logging/docs/storage).

        > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.organizations.Folder("default",
            display_name="some-folder-name",
            parent="organizations/123456789")
        basic = gcp.logging.FolderBucketConfig("basic",
            folder=default.name,
            location="global",
            retention_days=30,
            bucket_id="_Default",
            index_configs={
                "filePath": "jsonPayload.request.status",
                "type": "INDEX_TYPE_STRING",
            })
        ```

        ## Import

        This resource can be imported using the following format* `folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}` When using the `pulumi import` command, this resource can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:logging/folderBucketConfig:FolderBucketConfig default folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}
        ```

        :param str resource_name: The name of the resource.
        :param FolderBucketConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderBucketConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_id: Optional[pulumi.Input[str]] = None,
                 cmek_settings: Optional[pulumi.Input[pulumi.InputType['FolderBucketConfigCmekSettingsArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 index_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderBucketConfigIndexConfigArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderBucketConfigArgs.__new__(FolderBucketConfigArgs)

            if bucket_id is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_id'")
            __props__.__dict__["bucket_id"] = bucket_id
            __props__.__dict__["cmek_settings"] = cmek_settings
            __props__.__dict__["description"] = description
            if folder is None and not opts.urn:
                raise TypeError("Missing required property 'folder'")
            __props__.__dict__["folder"] = folder
            __props__.__dict__["index_configs"] = index_configs
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["retention_days"] = retention_days
            __props__.__dict__["lifecycle_state"] = None
            __props__.__dict__["name"] = None
        super(FolderBucketConfig, __self__).__init__(
            'gcp:logging/folderBucketConfig:FolderBucketConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_id: Optional[pulumi.Input[str]] = None,
            cmek_settings: Optional[pulumi.Input[pulumi.InputType['FolderBucketConfigCmekSettingsArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            index_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderBucketConfigIndexConfigArgs']]]]] = None,
            lifecycle_state: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retention_days: Optional[pulumi.Input[int]] = None) -> 'FolderBucketConfig':
        """
        Get an existing FolderBucketConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_id: The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        :param pulumi.Input[pulumi.InputType['FolderBucketConfigCmekSettingsArgs']] cmek_settings: The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
               key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
               updating the log bucket. Changing the KMS key is allowed.
        :param pulumi.Input[str] description: Describes this bucket.
        :param pulumi.Input[str] folder: The parent resource that contains the logging bucket.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FolderBucketConfigIndexConfigArgs']]]] index_configs: A list of indexed fields and related configuration data. Structure is documented below.
        :param pulumi.Input[str] lifecycle_state: The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        :param pulumi.Input[str] location: The location of the bucket.
        :param pulumi.Input[str] name: The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        :param pulumi.Input[int] retention_days: Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderBucketConfigState.__new__(_FolderBucketConfigState)

        __props__.__dict__["bucket_id"] = bucket_id
        __props__.__dict__["cmek_settings"] = cmek_settings
        __props__.__dict__["description"] = description
        __props__.__dict__["folder"] = folder
        __props__.__dict__["index_configs"] = index_configs
        __props__.__dict__["lifecycle_state"] = lifecycle_state
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["retention_days"] = retention_days
        return FolderBucketConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketId")
    def bucket_id(self) -> pulumi.Output[str]:
        """
        The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        """
        return pulumi.get(self, "bucket_id")

    @property
    @pulumi.getter(name="cmekSettings")
    def cmek_settings(self) -> pulumi.Output[Optional['outputs.FolderBucketConfigCmekSettings']]:
        """
        The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
        key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
        updating the log bucket. Changing the KMS key is allowed.
        """
        return pulumi.get(self, "cmek_settings")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Describes this bucket.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The parent resource that contains the logging bucket.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="indexConfigs")
    def index_configs(self) -> pulumi.Output[Optional[Sequence['outputs.FolderBucketConfigIndexConfig']]]:
        """
        A list of indexed fields and related configuration data. Structure is documented below.
        """
        return pulumi.get(self, "index_configs")

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> pulumi.Output[str]:
        """
        The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        """
        return pulumi.get(self, "lifecycle_state")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the bucket.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        """
        return pulumi.get(self, "retention_days")

