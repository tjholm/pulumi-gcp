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
    'GetGlobalForwardingRuleResult',
    'AwaitableGetGlobalForwardingRuleResult',
    'get_global_forwarding_rule',
    'get_global_forwarding_rule_output',
]

@pulumi.output_type
class GetGlobalForwardingRuleResult:
    """
    A collection of values returned by getGlobalForwardingRule.
    """
    def __init__(__self__, allow_psc_global_access=None, base_forwarding_rule=None, description=None, id=None, ip_address=None, ip_protocol=None, ip_version=None, label_fingerprint=None, labels=None, load_balancing_scheme=None, metadata_filters=None, name=None, network=None, port_range=None, project=None, psc_connection_id=None, psc_connection_status=None, self_link=None, source_ip_ranges=None, target=None):
        if allow_psc_global_access and not isinstance(allow_psc_global_access, bool):
            raise TypeError("Expected argument 'allow_psc_global_access' to be a bool")
        pulumi.set(__self__, "allow_psc_global_access", allow_psc_global_access)
        if base_forwarding_rule and not isinstance(base_forwarding_rule, str):
            raise TypeError("Expected argument 'base_forwarding_rule' to be a str")
        pulumi.set(__self__, "base_forwarding_rule", base_forwarding_rule)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if ip_protocol and not isinstance(ip_protocol, str):
            raise TypeError("Expected argument 'ip_protocol' to be a str")
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        if ip_version and not isinstance(ip_version, str):
            raise TypeError("Expected argument 'ip_version' to be a str")
        pulumi.set(__self__, "ip_version", ip_version)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if load_balancing_scheme and not isinstance(load_balancing_scheme, str):
            raise TypeError("Expected argument 'load_balancing_scheme' to be a str")
        pulumi.set(__self__, "load_balancing_scheme", load_balancing_scheme)
        if metadata_filters and not isinstance(metadata_filters, list):
            raise TypeError("Expected argument 'metadata_filters' to be a list")
        pulumi.set(__self__, "metadata_filters", metadata_filters)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        pulumi.set(__self__, "network", network)
        if port_range and not isinstance(port_range, str):
            raise TypeError("Expected argument 'port_range' to be a str")
        pulumi.set(__self__, "port_range", port_range)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if psc_connection_id and not isinstance(psc_connection_id, str):
            raise TypeError("Expected argument 'psc_connection_id' to be a str")
        pulumi.set(__self__, "psc_connection_id", psc_connection_id)
        if psc_connection_status and not isinstance(psc_connection_status, str):
            raise TypeError("Expected argument 'psc_connection_status' to be a str")
        pulumi.set(__self__, "psc_connection_status", psc_connection_status)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if source_ip_ranges and not isinstance(source_ip_ranges, list):
            raise TypeError("Expected argument 'source_ip_ranges' to be a list")
        pulumi.set(__self__, "source_ip_ranges", source_ip_ranges)
        if target and not isinstance(target, str):
            raise TypeError("Expected argument 'target' to be a str")
        pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter(name="allowPscGlobalAccess")
    def allow_psc_global_access(self) -> bool:
        return pulumi.get(self, "allow_psc_global_access")

    @property
    @pulumi.getter(name="baseForwardingRule")
    def base_forwarding_rule(self) -> str:
        return pulumi.get(self, "base_forwarding_rule")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> str:
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> str:
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="loadBalancingScheme")
    def load_balancing_scheme(self) -> str:
        return pulumi.get(self, "load_balancing_scheme")

    @property
    @pulumi.getter(name="metadataFilters")
    def metadata_filters(self) -> Sequence['outputs.GetGlobalForwardingRuleMetadataFilterResult']:
        return pulumi.get(self, "metadata_filters")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> str:
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="portRange")
    def port_range(self) -> str:
        return pulumi.get(self, "port_range")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pscConnectionId")
    def psc_connection_id(self) -> str:
        return pulumi.get(self, "psc_connection_id")

    @property
    @pulumi.getter(name="pscConnectionStatus")
    def psc_connection_status(self) -> str:
        return pulumi.get(self, "psc_connection_status")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sourceIpRanges")
    def source_ip_ranges(self) -> Sequence[str]:
        return pulumi.get(self, "source_ip_ranges")

    @property
    @pulumi.getter
    def target(self) -> str:
        return pulumi.get(self, "target")


class AwaitableGetGlobalForwardingRuleResult(GetGlobalForwardingRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlobalForwardingRuleResult(
            allow_psc_global_access=self.allow_psc_global_access,
            base_forwarding_rule=self.base_forwarding_rule,
            description=self.description,
            id=self.id,
            ip_address=self.ip_address,
            ip_protocol=self.ip_protocol,
            ip_version=self.ip_version,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            load_balancing_scheme=self.load_balancing_scheme,
            metadata_filters=self.metadata_filters,
            name=self.name,
            network=self.network,
            port_range=self.port_range,
            project=self.project,
            psc_connection_id=self.psc_connection_id,
            psc_connection_status=self.psc_connection_status,
            self_link=self.self_link,
            source_ip_ranges=self.source_ip_ranges,
            target=self.target)


def get_global_forwarding_rule(name: Optional[str] = None,
                               project: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGlobalForwardingRuleResult:
    """
    Get a global forwarding rule within GCE from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_forwarding_rule = gcp.compute.get_global_forwarding_rule(name="forwarding-rule-global")
    ```


    :param str name: The name of the global forwarding rule.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:compute/getGlobalForwardingRule:getGlobalForwardingRule', __args__, opts=opts, typ=GetGlobalForwardingRuleResult).value

    return AwaitableGetGlobalForwardingRuleResult(
        allow_psc_global_access=__ret__.allow_psc_global_access,
        base_forwarding_rule=__ret__.base_forwarding_rule,
        description=__ret__.description,
        id=__ret__.id,
        ip_address=__ret__.ip_address,
        ip_protocol=__ret__.ip_protocol,
        ip_version=__ret__.ip_version,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        load_balancing_scheme=__ret__.load_balancing_scheme,
        metadata_filters=__ret__.metadata_filters,
        name=__ret__.name,
        network=__ret__.network,
        port_range=__ret__.port_range,
        project=__ret__.project,
        psc_connection_id=__ret__.psc_connection_id,
        psc_connection_status=__ret__.psc_connection_status,
        self_link=__ret__.self_link,
        source_ip_ranges=__ret__.source_ip_ranges,
        target=__ret__.target)


@_utilities.lift_output_func(get_global_forwarding_rule)
def get_global_forwarding_rule_output(name: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGlobalForwardingRuleResult]:
    """
    Get a global forwarding rule within GCE from its name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_forwarding_rule = gcp.compute.get_global_forwarding_rule(name="forwarding-rule-global")
    ```


    :param str name: The name of the global forwarding rule.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    """
    ...
