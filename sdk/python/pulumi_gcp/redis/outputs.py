# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'InstanceServerCaCert',
    'GetInstanceServerCaCertResult',
]

@pulumi.output_type
class InstanceServerCaCert(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createTime":
            suggest = "create_time"
        elif key == "expireTime":
            suggest = "expire_time"
        elif key == "serialNumber":
            suggest = "serial_number"
        elif key == "sha1Fingerprint":
            suggest = "sha1_fingerprint"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstanceServerCaCert. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstanceServerCaCert.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstanceServerCaCert.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cert: Optional[str] = None,
                 create_time: Optional[str] = None,
                 expire_time: Optional[str] = None,
                 serial_number: Optional[str] = None,
                 sha1_fingerprint: Optional[str] = None):
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if sha1_fingerprint is not None:
            pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)

    @property
    @pulumi.getter
    def cert(self) -> Optional[str]:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[str]:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[str]:
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[str]:
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> Optional[str]:
        return pulumi.get(self, "sha1_fingerprint")


@pulumi.output_type
class GetInstanceServerCaCertResult(dict):
    def __init__(__self__, *,
                 cert: str,
                 create_time: str,
                 expire_time: str,
                 serial_number: str,
                 sha1_fingerprint: str):
        pulumi.set(__self__, "cert", cert)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "serial_number", serial_number)
        pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)

    @property
    @pulumi.getter
    def cert(self) -> str:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> str:
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> str:
        return pulumi.get(self, "sha1_fingerprint")


