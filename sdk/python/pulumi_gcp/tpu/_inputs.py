# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'NodeNetworkEndpointArgs',
    'NodeSchedulingConfigArgs',
]

@pulumi.input_type
class NodeNetworkEndpointArgs:
    def __init__(__self__, *,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class NodeSchedulingConfigArgs:
    def __init__(__self__, *,
                 preemptible: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] preemptible: Defines whether the TPU instance is preemptible.
        """
        pulumi.set(__self__, "preemptible", preemptible)

    @property
    @pulumi.getter
    def preemptible(self) -> pulumi.Input[bool]:
        """
        Defines whether the TPU instance is preemptible.
        """
        return pulumi.get(self, "preemptible")

    @preemptible.setter
    def preemptible(self, value: pulumi.Input[bool]):
        pulumi.set(self, "preemptible", value)


