# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DiskAsyncReplicationArgs', 'DiskAsyncReplication']

@pulumi.input_type
class DiskAsyncReplicationArgs:
    def __init__(__self__, *,
                 primary_disk: pulumi.Input[str],
                 secondary_disk: pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']):
        """
        The set of arguments for constructing a DiskAsyncReplication resource.
        :param pulumi.Input[str] primary_disk: The primary disk (source of replication).
        :param pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs'] secondary_disk: The secondary disk (target of replication). You can specify only one value. Structure is documented below.
               
               The `secondary_disk` block includes:
        """
        DiskAsyncReplicationArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            primary_disk=primary_disk,
            secondary_disk=secondary_disk,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             primary_disk: Optional[pulumi.Input[str]] = None,
             secondary_disk: Optional[pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if primary_disk is None and 'primaryDisk' in kwargs:
            primary_disk = kwargs['primaryDisk']
        if primary_disk is None:
            raise TypeError("Missing 'primary_disk' argument")
        if secondary_disk is None and 'secondaryDisk' in kwargs:
            secondary_disk = kwargs['secondaryDisk']
        if secondary_disk is None:
            raise TypeError("Missing 'secondary_disk' argument")

        _setter("primary_disk", primary_disk)
        _setter("secondary_disk", secondary_disk)

    @property
    @pulumi.getter(name="primaryDisk")
    def primary_disk(self) -> pulumi.Input[str]:
        """
        The primary disk (source of replication).
        """
        return pulumi.get(self, "primary_disk")

    @primary_disk.setter
    def primary_disk(self, value: pulumi.Input[str]):
        pulumi.set(self, "primary_disk", value)

    @property
    @pulumi.getter(name="secondaryDisk")
    def secondary_disk(self) -> pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']:
        """
        The secondary disk (target of replication). You can specify only one value. Structure is documented below.

        The `secondary_disk` block includes:
        """
        return pulumi.get(self, "secondary_disk")

    @secondary_disk.setter
    def secondary_disk(self, value: pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']):
        pulumi.set(self, "secondary_disk", value)


@pulumi.input_type
class _DiskAsyncReplicationState:
    def __init__(__self__, *,
                 primary_disk: Optional[pulumi.Input[str]] = None,
                 secondary_disk: Optional[pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']] = None):
        """
        Input properties used for looking up and filtering DiskAsyncReplication resources.
        :param pulumi.Input[str] primary_disk: The primary disk (source of replication).
        :param pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs'] secondary_disk: The secondary disk (target of replication). You can specify only one value. Structure is documented below.
               
               The `secondary_disk` block includes:
        """
        _DiskAsyncReplicationState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            primary_disk=primary_disk,
            secondary_disk=secondary_disk,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             primary_disk: Optional[pulumi.Input[str]] = None,
             secondary_disk: Optional[pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if primary_disk is None and 'primaryDisk' in kwargs:
            primary_disk = kwargs['primaryDisk']
        if secondary_disk is None and 'secondaryDisk' in kwargs:
            secondary_disk = kwargs['secondaryDisk']

        if primary_disk is not None:
            _setter("primary_disk", primary_disk)
        if secondary_disk is not None:
            _setter("secondary_disk", secondary_disk)

    @property
    @pulumi.getter(name="primaryDisk")
    def primary_disk(self) -> Optional[pulumi.Input[str]]:
        """
        The primary disk (source of replication).
        """
        return pulumi.get(self, "primary_disk")

    @primary_disk.setter
    def primary_disk(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_disk", value)

    @property
    @pulumi.getter(name="secondaryDisk")
    def secondary_disk(self) -> Optional[pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']]:
        """
        The secondary disk (target of replication). You can specify only one value. Structure is documented below.

        The `secondary_disk` block includes:
        """
        return pulumi.get(self, "secondary_disk")

    @secondary_disk.setter
    def secondary_disk(self, value: Optional[pulumi.Input['DiskAsyncReplicationSecondaryDiskArgs']]):
        pulumi.set(self, "secondary_disk", value)


class DiskAsyncReplication(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 primary_disk: Optional[pulumi.Input[str]] = None,
                 secondary_disk: Optional[pulumi.Input[pulumi.InputType['DiskAsyncReplicationSecondaryDiskArgs']]] = None,
                 __props__=None):
        """
        Starts and stops asynchronous persistent disk replication. For more information
        see [the official documentation](https://cloud.google.com/compute/docs/disks/async-pd/about)
        and the [API](https://cloud.google.com/compute/docs/reference/rest/v1/disks).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary_disk = gcp.compute.Disk("primary-disk",
            type="pd-ssd",
            zone="europe-west4-a",
            physical_block_size_bytes=4096)
        secondary_disk = gcp.compute.Disk("secondary-disk",
            type="pd-ssd",
            zone="europe-west3-a",
            async_primary_disk=gcp.compute.DiskAsyncPrimaryDiskArgs(
                disk=primary_disk.id,
            ),
            physical_block_size_bytes=4096)
        replication = gcp.compute.DiskAsyncReplication("replication",
            primary_disk=primary_disk.id,
            secondary_disk=gcp.compute.DiskAsyncReplicationSecondaryDiskArgs(
                disk=secondary_disk.id,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] primary_disk: The primary disk (source of replication).
        :param pulumi.Input[pulumi.InputType['DiskAsyncReplicationSecondaryDiskArgs']] secondary_disk: The secondary disk (target of replication). You can specify only one value. Structure is documented below.
               
               The `secondary_disk` block includes:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DiskAsyncReplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Starts and stops asynchronous persistent disk replication. For more information
        see [the official documentation](https://cloud.google.com/compute/docs/disks/async-pd/about)
        and the [API](https://cloud.google.com/compute/docs/reference/rest/v1/disks).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary_disk = gcp.compute.Disk("primary-disk",
            type="pd-ssd",
            zone="europe-west4-a",
            physical_block_size_bytes=4096)
        secondary_disk = gcp.compute.Disk("secondary-disk",
            type="pd-ssd",
            zone="europe-west3-a",
            async_primary_disk=gcp.compute.DiskAsyncPrimaryDiskArgs(
                disk=primary_disk.id,
            ),
            physical_block_size_bytes=4096)
        replication = gcp.compute.DiskAsyncReplication("replication",
            primary_disk=primary_disk.id,
            secondary_disk=gcp.compute.DiskAsyncReplicationSecondaryDiskArgs(
                disk=secondary_disk.id,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param DiskAsyncReplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DiskAsyncReplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DiskAsyncReplicationArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 primary_disk: Optional[pulumi.Input[str]] = None,
                 secondary_disk: Optional[pulumi.Input[pulumi.InputType['DiskAsyncReplicationSecondaryDiskArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DiskAsyncReplicationArgs.__new__(DiskAsyncReplicationArgs)

            if primary_disk is None and not opts.urn:
                raise TypeError("Missing required property 'primary_disk'")
            __props__.__dict__["primary_disk"] = primary_disk
            secondary_disk = _utilities.configure(secondary_disk, DiskAsyncReplicationSecondaryDiskArgs, True)
            if secondary_disk is None and not opts.urn:
                raise TypeError("Missing required property 'secondary_disk'")
            __props__.__dict__["secondary_disk"] = secondary_disk
        super(DiskAsyncReplication, __self__).__init__(
            'gcp:compute/diskAsyncReplication:DiskAsyncReplication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            primary_disk: Optional[pulumi.Input[str]] = None,
            secondary_disk: Optional[pulumi.Input[pulumi.InputType['DiskAsyncReplicationSecondaryDiskArgs']]] = None) -> 'DiskAsyncReplication':
        """
        Get an existing DiskAsyncReplication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] primary_disk: The primary disk (source of replication).
        :param pulumi.Input[pulumi.InputType['DiskAsyncReplicationSecondaryDiskArgs']] secondary_disk: The secondary disk (target of replication). You can specify only one value. Structure is documented below.
               
               The `secondary_disk` block includes:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DiskAsyncReplicationState.__new__(_DiskAsyncReplicationState)

        __props__.__dict__["primary_disk"] = primary_disk
        __props__.__dict__["secondary_disk"] = secondary_disk
        return DiskAsyncReplication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="primaryDisk")
    def primary_disk(self) -> pulumi.Output[str]:
        """
        The primary disk (source of replication).
        """
        return pulumi.get(self, "primary_disk")

    @property
    @pulumi.getter(name="secondaryDisk")
    def secondary_disk(self) -> pulumi.Output['outputs.DiskAsyncReplicationSecondaryDisk']:
        """
        The secondary disk (target of replication). You can specify only one value. Structure is documented below.

        The `secondary_disk` block includes:
        """
        return pulumi.get(self, "secondary_disk")

