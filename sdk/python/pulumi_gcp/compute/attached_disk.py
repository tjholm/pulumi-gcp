# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AttachedDiskArgs', 'AttachedDisk']

@pulumi.input_type
class AttachedDiskArgs:
    def __init__(__self__, *,
                 disk: pulumi.Input[str],
                 instance: pulumi.Input[str],
                 device_name: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AttachedDisk resource.
        :param pulumi.Input[str] disk: `name` or `self_link` of the disk that will be attached.
        :param pulumi.Input[str] instance: `name` or `self_link` of the compute instance that the disk will be attached to.
               If the `self_link` is provided then `zone` and `project` are extracted from the
               self link. If only the name is used then `zone` and `project` must be defined
               as properties on the resource or provider.
        :param pulumi.Input[str] device_name: Specifies a unique device name of your choice that is
               reflected into the /dev/disk/by-id/google-* tree of a Linux operating
               system running within the instance. This name can be used to
               reference the device for mounting, resizing, and so on, from within
               the instance.
        :param pulumi.Input[str] mode: The mode in which to attach this disk, either READ_WRITE or
               READ_ONLY. If not specified, the default is to attach the disk in
               READ_WRITE mode.
        :param pulumi.Input[str] project: The project that the referenced compute instance is a part of. If `instance` is referenced by its
               `self_link` the project defined in the link will take precedence.
        :param pulumi.Input[str] zone: The zone that the referenced compute instance is located within. If `instance` is referenced by its
               `self_link` the zone defined in the link will take precedence.
        """
        pulumi.set(__self__, "disk", disk)
        pulumi.set(__self__, "instance", instance)
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def disk(self) -> pulumi.Input[str]:
        """
        `name` or `self_link` of the disk that will be attached.
        """
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: pulumi.Input[str]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        `name` or `self_link` of the compute instance that the disk will be attached to.
        If the `self_link` is provided then `zone` and `project` are extracted from the
        self link. If only the name is used then `zone` and `project` must be defined
        as properties on the resource or provider.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique device name of your choice that is
        reflected into the /dev/disk/by-id/google-* tree of a Linux operating
        system running within the instance. This name can be used to
        reference the device for mounting, resizing, and so on, from within
        the instance.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode in which to attach this disk, either READ_WRITE or
        READ_ONLY. If not specified, the default is to attach the disk in
        READ_WRITE mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project that the referenced compute instance is a part of. If `instance` is referenced by its
        `self_link` the project defined in the link will take precedence.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone that the referenced compute instance is located within. If `instance` is referenced by its
        `self_link` the zone defined in the link will take precedence.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _AttachedDiskState:
    def __init__(__self__, *,
                 device_name: Optional[pulumi.Input[str]] = None,
                 disk: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AttachedDisk resources.
        :param pulumi.Input[str] device_name: Specifies a unique device name of your choice that is
               reflected into the /dev/disk/by-id/google-* tree of a Linux operating
               system running within the instance. This name can be used to
               reference the device for mounting, resizing, and so on, from within
               the instance.
        :param pulumi.Input[str] disk: `name` or `self_link` of the disk that will be attached.
        :param pulumi.Input[str] instance: `name` or `self_link` of the compute instance that the disk will be attached to.
               If the `self_link` is provided then `zone` and `project` are extracted from the
               self link. If only the name is used then `zone` and `project` must be defined
               as properties on the resource or provider.
        :param pulumi.Input[str] mode: The mode in which to attach this disk, either READ_WRITE or
               READ_ONLY. If not specified, the default is to attach the disk in
               READ_WRITE mode.
        :param pulumi.Input[str] project: The project that the referenced compute instance is a part of. If `instance` is referenced by its
               `self_link` the project defined in the link will take precedence.
        :param pulumi.Input[str] zone: The zone that the referenced compute instance is located within. If `instance` is referenced by its
               `self_link` the zone defined in the link will take precedence.
        """
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if disk is not None:
            pulumi.set(__self__, "disk", disk)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a unique device name of your choice that is
        reflected into the /dev/disk/by-id/google-* tree of a Linux operating
        system running within the instance. This name can be used to
        reference the device for mounting, resizing, and so on, from within
        the instance.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def disk(self) -> Optional[pulumi.Input[str]]:
        """
        `name` or `self_link` of the disk that will be attached.
        """
        return pulumi.get(self, "disk")

    @disk.setter
    def disk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        `name` or `self_link` of the compute instance that the disk will be attached to.
        If the `self_link` is provided then `zone` and `project` are extracted from the
        self link. If only the name is used then `zone` and `project` must be defined
        as properties on the resource or provider.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        The mode in which to attach this disk, either READ_WRITE or
        READ_ONLY. If not specified, the default is to attach the disk in
        READ_WRITE mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project that the referenced compute instance is a part of. If `instance` is referenced by its
        `self_link` the project defined in the link will take precedence.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone that the referenced compute instance is located within. If `instance` is referenced by its
        `self_link` the zone defined in the link will take precedence.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class AttachedDisk(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 disk: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_instance = gcp.compute.Instance("defaultInstance",
            machine_type="e2-medium",
            zone="us-west1-a",
            boot_disk=gcp.compute.InstanceBootDiskArgs(
                initialize_params=gcp.compute.InstanceBootDiskInitializeParamsArgs(
                    image="debian-cloud/debian-11",
                ),
            ),
            network_interfaces=[gcp.compute.InstanceNetworkInterfaceArgs(
                network="default",
            )])
        default_attached_disk = gcp.compute.AttachedDisk("defaultAttachedDisk",
            disk=google_compute_disk["default"]["id"],
            instance=default_instance.id)
        ```

        ## Import

        Attached Disk can be imported the following ways

        ```sh
         $ pulumi import gcp:compute/attachedDisk:AttachedDisk default projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
        ```

        ```sh
         $ pulumi import gcp:compute/attachedDisk:AttachedDisk default {{project}}/{{zone}}/{{instance.name}}/{{disk.name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: Specifies a unique device name of your choice that is
               reflected into the /dev/disk/by-id/google-* tree of a Linux operating
               system running within the instance. This name can be used to
               reference the device for mounting, resizing, and so on, from within
               the instance.
        :param pulumi.Input[str] disk: `name` or `self_link` of the disk that will be attached.
        :param pulumi.Input[str] instance: `name` or `self_link` of the compute instance that the disk will be attached to.
               If the `self_link` is provided then `zone` and `project` are extracted from the
               self link. If only the name is used then `zone` and `project` must be defined
               as properties on the resource or provider.
        :param pulumi.Input[str] mode: The mode in which to attach this disk, either READ_WRITE or
               READ_ONLY. If not specified, the default is to attach the disk in
               READ_WRITE mode.
        :param pulumi.Input[str] project: The project that the referenced compute instance is a part of. If `instance` is referenced by its
               `self_link` the project defined in the link will take precedence.
        :param pulumi.Input[str] zone: The zone that the referenced compute instance is located within. If `instance` is referenced by its
               `self_link` the zone defined in the link will take precedence.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AttachedDiskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_instance = gcp.compute.Instance("defaultInstance",
            machine_type="e2-medium",
            zone="us-west1-a",
            boot_disk=gcp.compute.InstanceBootDiskArgs(
                initialize_params=gcp.compute.InstanceBootDiskInitializeParamsArgs(
                    image="debian-cloud/debian-11",
                ),
            ),
            network_interfaces=[gcp.compute.InstanceNetworkInterfaceArgs(
                network="default",
            )])
        default_attached_disk = gcp.compute.AttachedDisk("defaultAttachedDisk",
            disk=google_compute_disk["default"]["id"],
            instance=default_instance.id)
        ```

        ## Import

        Attached Disk can be imported the following ways

        ```sh
         $ pulumi import gcp:compute/attachedDisk:AttachedDisk default projects/{{project}}/zones/{{zone}}/instances/{{instance.name}}/{{disk.name}}
        ```

        ```sh
         $ pulumi import gcp:compute/attachedDisk:AttachedDisk default {{project}}/{{zone}}/{{instance.name}}/{{disk.name}}
        ```

        :param str resource_name: The name of the resource.
        :param AttachedDiskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AttachedDiskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 disk: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AttachedDiskArgs.__new__(AttachedDiskArgs)

            __props__.__dict__["device_name"] = device_name
            if disk is None and not opts.urn:
                raise TypeError("Missing required property 'disk'")
            __props__.__dict__["disk"] = disk
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__.__dict__["instance"] = instance
            __props__.__dict__["mode"] = mode
            __props__.__dict__["project"] = project
            __props__.__dict__["zone"] = zone
        super(AttachedDisk, __self__).__init__(
            'gcp:compute/attachedDisk:AttachedDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_name: Optional[pulumi.Input[str]] = None,
            disk: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            mode: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'AttachedDisk':
        """
        Get an existing AttachedDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_name: Specifies a unique device name of your choice that is
               reflected into the /dev/disk/by-id/google-* tree of a Linux operating
               system running within the instance. This name can be used to
               reference the device for mounting, resizing, and so on, from within
               the instance.
        :param pulumi.Input[str] disk: `name` or `self_link` of the disk that will be attached.
        :param pulumi.Input[str] instance: `name` or `self_link` of the compute instance that the disk will be attached to.
               If the `self_link` is provided then `zone` and `project` are extracted from the
               self link. If only the name is used then `zone` and `project` must be defined
               as properties on the resource or provider.
        :param pulumi.Input[str] mode: The mode in which to attach this disk, either READ_WRITE or
               READ_ONLY. If not specified, the default is to attach the disk in
               READ_WRITE mode.
        :param pulumi.Input[str] project: The project that the referenced compute instance is a part of. If `instance` is referenced by its
               `self_link` the project defined in the link will take precedence.
        :param pulumi.Input[str] zone: The zone that the referenced compute instance is located within. If `instance` is referenced by its
               `self_link` the zone defined in the link will take precedence.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AttachedDiskState.__new__(_AttachedDiskState)

        __props__.__dict__["device_name"] = device_name
        __props__.__dict__["disk"] = disk
        __props__.__dict__["instance"] = instance
        __props__.__dict__["mode"] = mode
        __props__.__dict__["project"] = project
        __props__.__dict__["zone"] = zone
        return AttachedDisk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Output[str]:
        """
        Specifies a unique device name of your choice that is
        reflected into the /dev/disk/by-id/google-* tree of a Linux operating
        system running within the instance. This name can be used to
        reference the device for mounting, resizing, and so on, from within
        the instance.
        """
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter
    def disk(self) -> pulumi.Output[str]:
        """
        `name` or `self_link` of the disk that will be attached.
        """
        return pulumi.get(self, "disk")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        `name` or `self_link` of the compute instance that the disk will be attached to.
        If the `self_link` is provided then `zone` and `project` are extracted from the
        self link. If only the name is used then `zone` and `project` must be defined
        as properties on the resource or provider.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Output[Optional[str]]:
        """
        The mode in which to attach this disk, either READ_WRITE or
        READ_ONLY. If not specified, the default is to attach the disk in
        READ_WRITE mode.
        """
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project that the referenced compute instance is a part of. If `instance` is referenced by its
        `self_link` the project defined in the link will take precedence.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone that the referenced compute instance is located within. If `instance` is referenced by its
        `self_link` the zone defined in the link will take precedence.
        """
        return pulumi.get(self, "zone")

