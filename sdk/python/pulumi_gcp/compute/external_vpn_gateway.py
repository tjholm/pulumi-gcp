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

__all__ = ['ExternalVpnGatewayArgs', 'ExternalVpnGateway']

@pulumi.input_type
class ExternalVpnGatewayArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 redundancy_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ExternalVpnGateway resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]] interfaces: A list of interfaces on this external VPN gateway.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels for the external VPN gateway resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] redundancy_type: Indicates the redundancy type of this external VPN gateway
               Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        """
        ExternalVpnGatewayArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            interfaces=interfaces,
            labels=labels,
            name=name,
            project=project,
            redundancy_type=redundancy_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]] = None,
             labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             redundancy_type: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if redundancy_type is None and 'redundancyType' in kwargs:
            redundancy_type = kwargs['redundancyType']

        if description is not None:
            _setter("description", description)
        if interfaces is not None:
            _setter("interfaces", interfaces)
        if labels is not None:
            _setter("labels", labels)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if redundancy_type is not None:
            _setter("redundancy_type", redundancy_type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]]:
        """
        A list of interfaces on this external VPN gateway.
        Structure is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels for the external VPN gateway resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.


        - - -
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="redundancyType")
    def redundancy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the redundancy type of this external VPN gateway
        Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        """
        return pulumi.get(self, "redundancy_type")

    @redundancy_type.setter
    def redundancy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redundancy_type", value)


@pulumi.input_type
class _ExternalVpnGatewayState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]] = None,
                 label_fingerprint: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 redundancy_type: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExternalVpnGateway resources.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]] interfaces: A list of interfaces on this external VPN gateway.
               Structure is documented below.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource.  Used
               internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels for the external VPN gateway resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] redundancy_type: Indicates the redundancy type of this external VPN gateway
               Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        _ExternalVpnGatewayState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            interfaces=interfaces,
            label_fingerprint=label_fingerprint,
            labels=labels,
            name=name,
            project=project,
            redundancy_type=redundancy_type,
            self_link=self_link,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             interfaces: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]] = None,
             label_fingerprint: Optional[pulumi.Input[str]] = None,
             labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             redundancy_type: Optional[pulumi.Input[str]] = None,
             self_link: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if label_fingerprint is None and 'labelFingerprint' in kwargs:
            label_fingerprint = kwargs['labelFingerprint']
        if redundancy_type is None and 'redundancyType' in kwargs:
            redundancy_type = kwargs['redundancyType']
        if self_link is None and 'selfLink' in kwargs:
            self_link = kwargs['selfLink']

        if description is not None:
            _setter("description", description)
        if interfaces is not None:
            _setter("interfaces", interfaces)
        if label_fingerprint is not None:
            _setter("label_fingerprint", label_fingerprint)
        if labels is not None:
            _setter("labels", labels)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if redundancy_type is not None:
            _setter("redundancy_type", redundancy_type)
        if self_link is not None:
            _setter("self_link", self_link)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def interfaces(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]]:
        """
        A list of interfaces on this external VPN gateway.
        Structure is documented below.
        """
        return pulumi.get(self, "interfaces")

    @interfaces.setter
    def interfaces(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalVpnGatewayInterfaceArgs']]]]):
        pulumi.set(self, "interfaces", value)

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The fingerprint used for optimistic locking of this resource.  Used
        internally during updates.
        """
        return pulumi.get(self, "label_fingerprint")

    @label_fingerprint.setter
    def label_fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label_fingerprint", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels for the external VPN gateway resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.


        - - -
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="redundancyType")
    def redundancy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the redundancy type of this external VPN gateway
        Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        """
        return pulumi.get(self, "redundancy_type")

    @redundancy_type.setter
    def redundancy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "redundancy_type", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)


class ExternalVpnGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalVpnGatewayInterfaceArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 redundancy_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a VPN gateway managed outside of GCP.

        To get more information about ExternalVpnGateway, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)

        ## Example Usage
        ### External Vpn Gateway

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network = gcp.compute.Network("network",
            routing_mode="GLOBAL",
            auto_create_subnetworks=False)
        ha_gateway = gcp.compute.HaVpnGateway("haGateway",
            region="us-central1",
            network=network.id)
        external_gateway = gcp.compute.ExternalVpnGateway("externalGateway",
            redundancy_type="SINGLE_IP_INTERNALLY_REDUNDANT",
            description="An externally managed VPN gateway",
            interfaces=[gcp.compute.ExternalVpnGatewayInterfaceArgs(
                id=0,
                ip_address="8.8.8.8",
            )])
        network_subnet1 = gcp.compute.Subnetwork("networkSubnet1",
            ip_cidr_range="10.0.1.0/24",
            region="us-central1",
            network=network.id)
        network_subnet2 = gcp.compute.Subnetwork("networkSubnet2",
            ip_cidr_range="10.0.2.0/24",
            region="us-west1",
            network=network.id)
        router1 = gcp.compute.Router("router1",
            network=network.name,
            bgp=gcp.compute.RouterBgpArgs(
                asn=64514,
            ))
        tunnel1 = gcp.compute.VPNTunnel("tunnel1",
            region="us-central1",
            vpn_gateway=ha_gateway.id,
            peer_external_gateway=external_gateway.id,
            peer_external_gateway_interface=0,
            shared_secret="a secret message",
            router=router1.id,
            vpn_gateway_interface=0)
        tunnel2 = gcp.compute.VPNTunnel("tunnel2",
            region="us-central1",
            vpn_gateway=ha_gateway.id,
            peer_external_gateway=external_gateway.id,
            peer_external_gateway_interface=0,
            shared_secret="a secret message",
            router=router1.id.apply(lambda id: f" {id}"),
            vpn_gateway_interface=1)
        router1_interface1 = gcp.compute.RouterInterface("router1Interface1",
            router=router1.name,
            region="us-central1",
            ip_range="169.254.0.1/30",
            vpn_tunnel=tunnel1.name)
        router1_peer1 = gcp.compute.RouterPeer("router1Peer1",
            router=router1.name,
            region="us-central1",
            peer_ip_address="169.254.0.2",
            peer_asn=64515,
            advertised_route_priority=100,
            interface=router1_interface1.name)
        router1_interface2 = gcp.compute.RouterInterface("router1Interface2",
            router=router1.name,
            region="us-central1",
            ip_range="169.254.1.1/30",
            vpn_tunnel=tunnel2.name)
        router1_peer2 = gcp.compute.RouterPeer("router1Peer2",
            router=router1.name,
            region="us-central1",
            peer_ip_address="169.254.1.2",
            peer_asn=64515,
            advertised_route_priority=100,
            interface=router1_interface2.name)
        ```

        ## Import

        ExternalVpnGateway can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalVpnGatewayInterfaceArgs']]]] interfaces: A list of interfaces on this external VPN gateway.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels for the external VPN gateway resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] redundancy_type: Indicates the redundancy type of this external VPN gateway
               Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ExternalVpnGatewayArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a VPN gateway managed outside of GCP.

        To get more information about ExternalVpnGateway, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways)

        ## Example Usage
        ### External Vpn Gateway

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network = gcp.compute.Network("network",
            routing_mode="GLOBAL",
            auto_create_subnetworks=False)
        ha_gateway = gcp.compute.HaVpnGateway("haGateway",
            region="us-central1",
            network=network.id)
        external_gateway = gcp.compute.ExternalVpnGateway("externalGateway",
            redundancy_type="SINGLE_IP_INTERNALLY_REDUNDANT",
            description="An externally managed VPN gateway",
            interfaces=[gcp.compute.ExternalVpnGatewayInterfaceArgs(
                id=0,
                ip_address="8.8.8.8",
            )])
        network_subnet1 = gcp.compute.Subnetwork("networkSubnet1",
            ip_cidr_range="10.0.1.0/24",
            region="us-central1",
            network=network.id)
        network_subnet2 = gcp.compute.Subnetwork("networkSubnet2",
            ip_cidr_range="10.0.2.0/24",
            region="us-west1",
            network=network.id)
        router1 = gcp.compute.Router("router1",
            network=network.name,
            bgp=gcp.compute.RouterBgpArgs(
                asn=64514,
            ))
        tunnel1 = gcp.compute.VPNTunnel("tunnel1",
            region="us-central1",
            vpn_gateway=ha_gateway.id,
            peer_external_gateway=external_gateway.id,
            peer_external_gateway_interface=0,
            shared_secret="a secret message",
            router=router1.id,
            vpn_gateway_interface=0)
        tunnel2 = gcp.compute.VPNTunnel("tunnel2",
            region="us-central1",
            vpn_gateway=ha_gateway.id,
            peer_external_gateway=external_gateway.id,
            peer_external_gateway_interface=0,
            shared_secret="a secret message",
            router=router1.id.apply(lambda id: f" {id}"),
            vpn_gateway_interface=1)
        router1_interface1 = gcp.compute.RouterInterface("router1Interface1",
            router=router1.name,
            region="us-central1",
            ip_range="169.254.0.1/30",
            vpn_tunnel=tunnel1.name)
        router1_peer1 = gcp.compute.RouterPeer("router1Peer1",
            router=router1.name,
            region="us-central1",
            peer_ip_address="169.254.0.2",
            peer_asn=64515,
            advertised_route_priority=100,
            interface=router1_interface1.name)
        router1_interface2 = gcp.compute.RouterInterface("router1Interface2",
            router=router1.name,
            region="us-central1",
            ip_range="169.254.1.1/30",
            vpn_tunnel=tunnel2.name)
        router1_peer2 = gcp.compute.RouterPeer("router1Peer2",
            router=router1.name,
            region="us-central1",
            peer_ip_address="169.254.1.2",
            peer_asn=64515,
            advertised_route_priority=100,
            interface=router1_interface2.name)
        ```

        ## Import

        ExternalVpnGateway can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default projects/{{project}}/global/externalVpnGateways/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/externalVpnGateway:ExternalVpnGateway default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ExternalVpnGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExternalVpnGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ExternalVpnGatewayArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalVpnGatewayInterfaceArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 redundancy_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExternalVpnGatewayArgs.__new__(ExternalVpnGatewayArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["interfaces"] = interfaces
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["redundancy_type"] = redundancy_type
            __props__.__dict__["label_fingerprint"] = None
            __props__.__dict__["self_link"] = None
        super(ExternalVpnGateway, __self__).__init__(
            'gcp:compute/externalVpnGateway:ExternalVpnGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            interfaces: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalVpnGatewayInterfaceArgs']]]]] = None,
            label_fingerprint: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            redundancy_type: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'ExternalVpnGateway':
        """
        Get an existing ExternalVpnGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalVpnGatewayInterfaceArgs']]]] interfaces: A list of interfaces on this external VPN gateway.
               Structure is documented below.
        :param pulumi.Input[str] label_fingerprint: The fingerprint used for optimistic locking of this resource.  Used
               internally during updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels for the external VPN gateway resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] redundancy_type: Indicates the redundancy type of this external VPN gateway
               Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExternalVpnGatewayState.__new__(_ExternalVpnGatewayState)

        __props__.__dict__["description"] = description
        __props__.__dict__["interfaces"] = interfaces
        __props__.__dict__["label_fingerprint"] = label_fingerprint
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["redundancy_type"] = redundancy_type
        __props__.__dict__["self_link"] = self_link
        return ExternalVpnGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def interfaces(self) -> pulumi.Output[Optional[Sequence['outputs.ExternalVpnGatewayInterface']]]:
        """
        A list of interfaces on this external VPN gateway.
        Structure is documented below.
        """
        return pulumi.get(self, "interfaces")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> pulumi.Output[str]:
        """
        The fingerprint used for optimistic locking of this resource.  Used
        internally during updates.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels for the external VPN gateway resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.


        - - -
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="redundancyType")
    def redundancy_type(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates the redundancy type of this external VPN gateway
        Possible values are: `FOUR_IPS_REDUNDANCY`, `SINGLE_IP_INTERNALLY_REDUNDANT`, `TWO_IPS_REDUNDANCY`.
        """
        return pulumi.get(self, "redundancy_type")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

