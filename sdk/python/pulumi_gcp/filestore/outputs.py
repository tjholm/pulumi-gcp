# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'InstanceFileShares',
    'InstanceFileSharesNfsExportOption',
    'InstanceNetwork',
]

@pulumi.output_type
class InstanceFileShares(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "capacityGb":
            suggest = "capacity_gb"
        elif key == "nfsExportOptions":
            suggest = "nfs_export_options"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceFileShares. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceFileShares.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceFileShares.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 capacity_gb: int,
                 name: str,
                 nfs_export_options: Optional[Sequence['outputs.InstanceFileSharesNfsExportOption']] = None):
        """
        :param int capacity_gb: File share capacity in GiB. This must be at least 1024 GiB
               for the standard tier, or 2560 GiB for the premium tier.
        :param str name: The name of the fileshare (16 characters or less)
        :param Sequence['InstanceFileSharesNfsExportOptionArgs'] nfs_export_options: Nfs Export Options. There is a limit of 10 export options per file share.
               Structure is documented below.
        """
        pulumi.set(__self__, "capacity_gb", capacity_gb)
        pulumi.set(__self__, "name", name)
        if nfs_export_options is not None:
            pulumi.set(__self__, "nfs_export_options", nfs_export_options)

    @property
    @pulumi.getter(name="capacityGb")
    def capacity_gb(self) -> int:
        """
        File share capacity in GiB. This must be at least 1024 GiB
        for the standard tier, or 2560 GiB for the premium tier.
        """
        return pulumi.get(self, "capacity_gb")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the fileshare (16 characters or less)
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nfsExportOptions")
    def nfs_export_options(self) -> Optional[Sequence['outputs.InstanceFileSharesNfsExportOption']]:
        """
        Nfs Export Options. There is a limit of 10 export options per file share.
        Structure is documented below.
        """
        return pulumi.get(self, "nfs_export_options")


@pulumi.output_type
class InstanceFileSharesNfsExportOption(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accessMode":
            suggest = "access_mode"
        elif key == "anonGid":
            suggest = "anon_gid"
        elif key == "anonUid":
            suggest = "anon_uid"
        elif key == "ipRanges":
            suggest = "ip_ranges"
        elif key == "squashMode":
            suggest = "squash_mode"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceFileSharesNfsExportOption. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceFileSharesNfsExportOption.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceFileSharesNfsExportOption.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 access_mode: Optional[str] = None,
                 anon_gid: Optional[int] = None,
                 anon_uid: Optional[int] = None,
                 ip_ranges: Optional[Sequence[str]] = None,
                 squash_mode: Optional[str] = None):
        """
        :param str access_mode: Either READ_ONLY, for allowing only read requests on the exported directory,
               or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
               Default value is `READ_WRITE`.
               Possible values are `READ_ONLY` and `READ_WRITE`.
        :param int anon_gid: An integer representing the anonymous group id with a default value of 65534.
               Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
               if this field is specified for other squashMode settings.
        :param int anon_uid: An integer representing the anonymous user id with a default value of 65534.
               Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
               if this field is specified for other squashMode settings.
        :param Sequence[str] ip_ranges: List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
               Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
               The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        :param str squash_mode: Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
               for not allowing root access. The default is NO_ROOT_SQUASH.
               Default value is `NO_ROOT_SQUASH`.
               Possible values are `NO_ROOT_SQUASH` and `ROOT_SQUASH`.
        """
        if access_mode is not None:
            pulumi.set(__self__, "access_mode", access_mode)
        if anon_gid is not None:
            pulumi.set(__self__, "anon_gid", anon_gid)
        if anon_uid is not None:
            pulumi.set(__self__, "anon_uid", anon_uid)
        if ip_ranges is not None:
            pulumi.set(__self__, "ip_ranges", ip_ranges)
        if squash_mode is not None:
            pulumi.set(__self__, "squash_mode", squash_mode)

    @property
    @pulumi.getter(name="accessMode")
    def access_mode(self) -> Optional[str]:
        """
        Either READ_ONLY, for allowing only read requests on the exported directory,
        or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
        Default value is `READ_WRITE`.
        Possible values are `READ_ONLY` and `READ_WRITE`.
        """
        return pulumi.get(self, "access_mode")

    @property
    @pulumi.getter(name="anonGid")
    def anon_gid(self) -> Optional[int]:
        """
        An integer representing the anonymous group id with a default value of 65534.
        Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
        if this field is specified for other squashMode settings.
        """
        return pulumi.get(self, "anon_gid")

    @property
    @pulumi.getter(name="anonUid")
    def anon_uid(self) -> Optional[int]:
        """
        An integer representing the anonymous user id with a default value of 65534.
        Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
        if this field is specified for other squashMode settings.
        """
        return pulumi.get(self, "anon_uid")

    @property
    @pulumi.getter(name="ipRanges")
    def ip_ranges(self) -> Optional[Sequence[str]]:
        """
        List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
        Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
        The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
        """
        return pulumi.get(self, "ip_ranges")

    @property
    @pulumi.getter(name="squashMode")
    def squash_mode(self) -> Optional[str]:
        """
        Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH,
        for not allowing root access. The default is NO_ROOT_SQUASH.
        Default value is `NO_ROOT_SQUASH`.
        Possible values are `NO_ROOT_SQUASH` and `ROOT_SQUASH`.
        """
        return pulumi.get(self, "squash_mode")


@pulumi.output_type
class InstanceNetwork(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "ipAddresses":
            suggest = "ip_addresses"
        elif key == "reservedIpRange":
            suggest = "reserved_ip_range"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceNetwork. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceNetwork.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceNetwork.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 modes: Sequence[str],
                 network: str,
                 ip_addresses: Optional[Sequence[str]] = None,
                 reserved_ip_range: Optional[str] = None):
        """
        :param Sequence[str] modes: IP versions for which the instance has
               IP addresses assigned.
               Each value may be one of `ADDRESS_MODE_UNSPECIFIED`, `MODE_IPV4`, and `MODE_IPV6`.
        :param str network: The name of the GCE VPC network to which the
               instance is connected.
        :param Sequence[str] ip_addresses: -
               A list of IPv4 or IPv6 addresses.
        :param str reserved_ip_range: A /29 CIDR block that identifies the range of IP
               addresses reserved for this instance.
        """
        pulumi.set(__self__, "modes", modes)
        pulumi.set(__self__, "network", network)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if reserved_ip_range is not None:
            pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)

    @property
    @pulumi.getter
    def modes(self) -> Sequence[str]:
        """
        IP versions for which the instance has
        IP addresses assigned.
        Each value may be one of `ADDRESS_MODE_UNSPECIFIED`, `MODE_IPV4`, and `MODE_IPV6`.
        """
        return pulumi.get(self, "modes")

    @property
    @pulumi.getter
    def network(self) -> str:
        """
        The name of the GCE VPC network to which the
        instance is connected.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[Sequence[str]]:
        """
        -
        A list of IPv4 or IPv6 addresses.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> Optional[str]:
        """
        A /29 CIDR block that identifies the range of IP
        addresses reserved for this instance.
        """
        return pulumi.get(self, "reserved_ip_range")


