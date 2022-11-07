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
    'GetDiskResult',
    'AwaitableGetDiskResult',
    'get_disk',
    'get_disk_output',
]

@pulumi.output_type
class GetDiskResult:
    """
    A collection of values returned by getDisk.
    """
    def __init__(__self__, creation_timestamp=None, description=None, disk_encryption_keys=None, id=None, image=None, interface=None, label_fingerprint=None, labels=None, last_attach_timestamp=None, last_detach_timestamp=None, multi_writer=None, name=None, physical_block_size_bytes=None, project=None, provisioned_iops=None, resource_policies=None, self_link=None, size=None, snapshot=None, source_disk=None, source_disk_id=None, source_image_encryption_keys=None, source_image_id=None, source_snapshot_encryption_keys=None, source_snapshot_id=None, type=None, users=None, zone=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disk_encryption_keys and not isinstance(disk_encryption_keys, list):
            raise TypeError("Expected argument 'disk_encryption_keys' to be a list")
        pulumi.set(__self__, "disk_encryption_keys", disk_encryption_keys)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if image and not isinstance(image, str):
            raise TypeError("Expected argument 'image' to be a str")
        pulumi.set(__self__, "image", image)
        if interface and not isinstance(interface, str):
            raise TypeError("Expected argument 'interface' to be a str")
        pulumi.set(__self__, "interface", interface)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_attach_timestamp and not isinstance(last_attach_timestamp, str):
            raise TypeError("Expected argument 'last_attach_timestamp' to be a str")
        pulumi.set(__self__, "last_attach_timestamp", last_attach_timestamp)
        if last_detach_timestamp and not isinstance(last_detach_timestamp, str):
            raise TypeError("Expected argument 'last_detach_timestamp' to be a str")
        pulumi.set(__self__, "last_detach_timestamp", last_detach_timestamp)
        if multi_writer and not isinstance(multi_writer, bool):
            raise TypeError("Expected argument 'multi_writer' to be a bool")
        pulumi.set(__self__, "multi_writer", multi_writer)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if physical_block_size_bytes and not isinstance(physical_block_size_bytes, int):
            raise TypeError("Expected argument 'physical_block_size_bytes' to be a int")
        pulumi.set(__self__, "physical_block_size_bytes", physical_block_size_bytes)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if provisioned_iops and not isinstance(provisioned_iops, int):
            raise TypeError("Expected argument 'provisioned_iops' to be a int")
        pulumi.set(__self__, "provisioned_iops", provisioned_iops)
        if resource_policies and not isinstance(resource_policies, list):
            raise TypeError("Expected argument 'resource_policies' to be a list")
        pulumi.set(__self__, "resource_policies", resource_policies)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if snapshot and not isinstance(snapshot, str):
            raise TypeError("Expected argument 'snapshot' to be a str")
        pulumi.set(__self__, "snapshot", snapshot)
        if source_disk and not isinstance(source_disk, str):
            raise TypeError("Expected argument 'source_disk' to be a str")
        pulumi.set(__self__, "source_disk", source_disk)
        if source_disk_id and not isinstance(source_disk_id, str):
            raise TypeError("Expected argument 'source_disk_id' to be a str")
        pulumi.set(__self__, "source_disk_id", source_disk_id)
        if source_image_encryption_keys and not isinstance(source_image_encryption_keys, list):
            raise TypeError("Expected argument 'source_image_encryption_keys' to be a list")
        pulumi.set(__self__, "source_image_encryption_keys", source_image_encryption_keys)
        if source_image_id and not isinstance(source_image_id, str):
            raise TypeError("Expected argument 'source_image_id' to be a str")
        pulumi.set(__self__, "source_image_id", source_image_id)
        if source_snapshot_encryption_keys and not isinstance(source_snapshot_encryption_keys, list):
            raise TypeError("Expected argument 'source_snapshot_encryption_keys' to be a list")
        pulumi.set(__self__, "source_snapshot_encryption_keys", source_snapshot_encryption_keys)
        if source_snapshot_id and not isinstance(source_snapshot_id, str):
            raise TypeError("Expected argument 'source_snapshot_id' to be a str")
        pulumi.set(__self__, "source_snapshot_id", source_snapshot_id)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskEncryptionKeys")
    def disk_encryption_keys(self) -> Sequence['outputs.GetDiskDiskEncryptionKeyResult']:
        return pulumi.get(self, "disk_encryption_keys")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def image(self) -> str:
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def interface(self) -> str:
        return pulumi.get(self, "interface")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        A map of labels applied to this disk.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastAttachTimestamp")
    def last_attach_timestamp(self) -> str:
        return pulumi.get(self, "last_attach_timestamp")

    @property
    @pulumi.getter(name="lastDetachTimestamp")
    def last_detach_timestamp(self) -> str:
        return pulumi.get(self, "last_detach_timestamp")

    @property
    @pulumi.getter(name="multiWriter")
    def multi_writer(self) -> bool:
        return pulumi.get(self, "multi_writer")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="physicalBlockSizeBytes")
    def physical_block_size_bytes(self) -> int:
        return pulumi.get(self, "physical_block_size_bytes")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="provisionedIops")
    def provisioned_iops(self) -> int:
        return pulumi.get(self, "provisioned_iops")

    @property
    @pulumi.getter(name="resourcePolicies")
    def resource_policies(self) -> Sequence[str]:
        return pulumi.get(self, "resource_policies")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> int:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def snapshot(self) -> str:
        return pulumi.get(self, "snapshot")

    @property
    @pulumi.getter(name="sourceDisk")
    def source_disk(self) -> str:
        return pulumi.get(self, "source_disk")

    @property
    @pulumi.getter(name="sourceDiskId")
    def source_disk_id(self) -> str:
        return pulumi.get(self, "source_disk_id")

    @property
    @pulumi.getter(name="sourceImageEncryptionKeys")
    def source_image_encryption_keys(self) -> Sequence['outputs.GetDiskSourceImageEncryptionKeyResult']:
        return pulumi.get(self, "source_image_encryption_keys")

    @property
    @pulumi.getter(name="sourceImageId")
    def source_image_id(self) -> str:
        return pulumi.get(self, "source_image_id")

    @property
    @pulumi.getter(name="sourceSnapshotEncryptionKeys")
    def source_snapshot_encryption_keys(self) -> Sequence['outputs.GetDiskSourceSnapshotEncryptionKeyResult']:
        return pulumi.get(self, "source_snapshot_encryption_keys")

    @property
    @pulumi.getter(name="sourceSnapshotId")
    def source_snapshot_id(self) -> str:
        return pulumi.get(self, "source_snapshot_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def users(self) -> Sequence[str]:
        return pulumi.get(self, "users")

    @property
    @pulumi.getter
    def zone(self) -> Optional[str]:
        return pulumi.get(self, "zone")


class AwaitableGetDiskResult(GetDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiskResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            disk_encryption_keys=self.disk_encryption_keys,
            id=self.id,
            image=self.image,
            interface=self.interface,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            last_attach_timestamp=self.last_attach_timestamp,
            last_detach_timestamp=self.last_detach_timestamp,
            multi_writer=self.multi_writer,
            name=self.name,
            physical_block_size_bytes=self.physical_block_size_bytes,
            project=self.project,
            provisioned_iops=self.provisioned_iops,
            resource_policies=self.resource_policies,
            self_link=self.self_link,
            size=self.size,
            snapshot=self.snapshot,
            source_disk=self.source_disk,
            source_disk_id=self.source_disk_id,
            source_image_encryption_keys=self.source_image_encryption_keys,
            source_image_id=self.source_image_id,
            source_snapshot_encryption_keys=self.source_snapshot_encryption_keys,
            source_snapshot_id=self.source_snapshot_id,
            type=self.type,
            users=self.users,
            zone=self.zone)


def get_disk(name: Optional[str] = None,
             project: Optional[str] = None,
             zone: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiskResult:
    """
    Get information about a Google Compute Persistent disks.

    [the official documentation](https://cloud.google.com/compute/docs/disks) and its [API](https://cloud.google.com/compute/docs/reference/latest/disks).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    persistent_boot_disk = gcp.compute.get_disk(name="persistent-boot-disk",
        project="example")
    # ...
    default = gcp.compute.Instance("default", boot_disk=gcp.compute.InstanceBootDiskArgs(
        source=persistent_boot_disk.self_link,
        auto_delete=False,
    ))
    ```


    :param str name: The name of a specific disk.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the provider project is used.
    :param str zone: A reference to the zone where the disk resides.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:compute/getDisk:getDisk', __args__, opts=opts, typ=GetDiskResult).value

    return AwaitableGetDiskResult(
        creation_timestamp=__ret__.creation_timestamp,
        description=__ret__.description,
        disk_encryption_keys=__ret__.disk_encryption_keys,
        id=__ret__.id,
        image=__ret__.image,
        interface=__ret__.interface,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        last_attach_timestamp=__ret__.last_attach_timestamp,
        last_detach_timestamp=__ret__.last_detach_timestamp,
        multi_writer=__ret__.multi_writer,
        name=__ret__.name,
        physical_block_size_bytes=__ret__.physical_block_size_bytes,
        project=__ret__.project,
        provisioned_iops=__ret__.provisioned_iops,
        resource_policies=__ret__.resource_policies,
        self_link=__ret__.self_link,
        size=__ret__.size,
        snapshot=__ret__.snapshot,
        source_disk=__ret__.source_disk,
        source_disk_id=__ret__.source_disk_id,
        source_image_encryption_keys=__ret__.source_image_encryption_keys,
        source_image_id=__ret__.source_image_id,
        source_snapshot_encryption_keys=__ret__.source_snapshot_encryption_keys,
        source_snapshot_id=__ret__.source_snapshot_id,
        type=__ret__.type,
        users=__ret__.users,
        zone=__ret__.zone)


@_utilities.lift_output_func(get_disk)
def get_disk_output(name: Optional[pulumi.Input[str]] = None,
                    project: Optional[pulumi.Input[Optional[str]]] = None,
                    zone: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDiskResult]:
    """
    Get information about a Google Compute Persistent disks.

    [the official documentation](https://cloud.google.com/compute/docs/disks) and its [API](https://cloud.google.com/compute/docs/reference/latest/disks).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    persistent_boot_disk = gcp.compute.get_disk(name="persistent-boot-disk",
        project="example")
    # ...
    default = gcp.compute.Instance("default", boot_disk=gcp.compute.InstanceBootDiskArgs(
        source=persistent_boot_disk.self_link,
        auto_delete=False,
    ))
    ```


    :param str name: The name of a specific disk.
    :param str project: The ID of the project in which the resource belongs.
           If it is not provided, the provider project is used.
    :param str zone: A reference to the zone where the disk resides.
    """
    ...
