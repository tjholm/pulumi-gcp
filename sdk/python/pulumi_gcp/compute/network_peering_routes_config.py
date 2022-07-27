# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkPeeringRoutesConfigArgs', 'NetworkPeeringRoutesConfig']

@pulumi.input_type
class NetworkPeeringRoutesConfigArgs:
    def __init__(__self__, *,
                 export_custom_routes: pulumi.Input[bool],
                 import_custom_routes: pulumi.Input[bool],
                 network: pulumi.Input[str],
                 peering: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetworkPeeringRoutesConfig resource.
        :param pulumi.Input[bool] export_custom_routes: Whether to export the custom routes to the peer network.
        :param pulumi.Input[bool] import_custom_routes: Whether to import the custom routes to the peer network.
        :param pulumi.Input[str] network: The name of the primary network for the peering.
        :param pulumi.Input[str] peering: Name of the peering.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "export_custom_routes", export_custom_routes)
        pulumi.set(__self__, "import_custom_routes", import_custom_routes)
        pulumi.set(__self__, "network", network)
        pulumi.set(__self__, "peering", peering)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="exportCustomRoutes")
    def export_custom_routes(self) -> pulumi.Input[bool]:
        """
        Whether to export the custom routes to the peer network.
        """
        return pulumi.get(self, "export_custom_routes")

    @export_custom_routes.setter
    def export_custom_routes(self, value: pulumi.Input[bool]):
        pulumi.set(self, "export_custom_routes", value)

    @property
    @pulumi.getter(name="importCustomRoutes")
    def import_custom_routes(self) -> pulumi.Input[bool]:
        """
        Whether to import the custom routes to the peer network.
        """
        return pulumi.get(self, "import_custom_routes")

    @import_custom_routes.setter
    def import_custom_routes(self, value: pulumi.Input[bool]):
        pulumi.set(self, "import_custom_routes", value)

    @property
    @pulumi.getter
    def network(self) -> pulumi.Input[str]:
        """
        The name of the primary network for the peering.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: pulumi.Input[str]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def peering(self) -> pulumi.Input[str]:
        """
        Name of the peering.
        """
        return pulumi.get(self, "peering")

    @peering.setter
    def peering(self, value: pulumi.Input[str]):
        pulumi.set(self, "peering", value)

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


@pulumi.input_type
class _NetworkPeeringRoutesConfigState:
    def __init__(__self__, *,
                 export_custom_routes: Optional[pulumi.Input[bool]] = None,
                 import_custom_routes: Optional[pulumi.Input[bool]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 peering: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkPeeringRoutesConfig resources.
        :param pulumi.Input[bool] export_custom_routes: Whether to export the custom routes to the peer network.
        :param pulumi.Input[bool] import_custom_routes: Whether to import the custom routes to the peer network.
        :param pulumi.Input[str] network: The name of the primary network for the peering.
        :param pulumi.Input[str] peering: Name of the peering.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if export_custom_routes is not None:
            pulumi.set(__self__, "export_custom_routes", export_custom_routes)
        if import_custom_routes is not None:
            pulumi.set(__self__, "import_custom_routes", import_custom_routes)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if peering is not None:
            pulumi.set(__self__, "peering", peering)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="exportCustomRoutes")
    def export_custom_routes(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to export the custom routes to the peer network.
        """
        return pulumi.get(self, "export_custom_routes")

    @export_custom_routes.setter
    def export_custom_routes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "export_custom_routes", value)

    @property
    @pulumi.getter(name="importCustomRoutes")
    def import_custom_routes(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to import the custom routes to the peer network.
        """
        return pulumi.get(self, "import_custom_routes")

    @import_custom_routes.setter
    def import_custom_routes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "import_custom_routes", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the primary network for the peering.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def peering(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the peering.
        """
        return pulumi.get(self, "peering")

    @peering.setter
    def peering(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering", value)

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


class NetworkPeeringRoutesConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 export_custom_routes: Optional[pulumi.Input[bool]] = None,
                 import_custom_routes: Optional[pulumi.Input[bool]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 peering: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage a network peering's route settings without managing the peering as
        a whole. This resource is primarily intended for use with GCP-generated
        peerings that shouldn't otherwise be managed by other tools. Deleting this
        resource is a no-op and the peering will not be modified.

        To get more information about NetworkPeeringRoutesConfig, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vpc/docs/vpc-peering)

        ## Example Usage
        ### Network Peering Routes Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network_primary = gcp.compute.Network("networkPrimary", auto_create_subnetworks=False)
        network_secondary = gcp.compute.Network("networkSecondary", auto_create_subnetworks=False)
        peering_primary = gcp.compute.NetworkPeering("peeringPrimary",
            network=network_primary.id,
            peer_network=network_secondary.id,
            import_custom_routes=True,
            export_custom_routes=True)
        peering_primary_routes = gcp.compute.NetworkPeeringRoutesConfig("peeringPrimaryRoutes",
            peering=peering_primary.name,
            network=network_primary.name,
            import_custom_routes=True,
            export_custom_routes=True)
        peering_secondary = gcp.compute.NetworkPeering("peeringSecondary",
            network=network_secondary.id,
            peer_network=network_primary.id)
        ```
        ### Network Peering Routes Config Gke

        ```python
        import pulumi
        import pulumi_gcp as gcp

        container_network = gcp.compute.Network("containerNetwork", auto_create_subnetworks=False)
        container_subnetwork = gcp.compute.Subnetwork("containerSubnetwork",
            region="us-central1",
            network=container_network.name,
            ip_cidr_range="10.0.36.0/24",
            private_ip_google_access=True,
            secondary_ip_ranges=[
                gcp.compute.SubnetworkSecondaryIpRangeArgs(
                    range_name="pod",
                    ip_cidr_range="10.0.0.0/19",
                ),
                gcp.compute.SubnetworkSecondaryIpRangeArgs(
                    range_name="svc",
                    ip_cidr_range="10.0.32.0/22",
                ),
            ])
        private_cluster = gcp.container.Cluster("privateCluster",
            location="us-central1-a",
            initial_node_count=1,
            network=container_network.name,
            subnetwork=container_subnetwork.name,
            private_cluster_config=gcp.container.ClusterPrivateClusterConfigArgs(
                enable_private_endpoint=True,
                enable_private_nodes=True,
                master_ipv4_cidr_block="10.42.0.0/28",
            ),
            master_authorized_networks_config=gcp.container.ClusterMasterAuthorizedNetworksConfigArgs(),
            ip_allocation_policy=gcp.container.ClusterIpAllocationPolicyArgs(
                cluster_secondary_range_name=container_subnetwork.secondary_ip_ranges[0].range_name,
                services_secondary_range_name=container_subnetwork.secondary_ip_ranges[1].range_name,
            ))
        peering_gke_routes = gcp.compute.NetworkPeeringRoutesConfig("peeringGkeRoutes",
            peering=private_cluster.private_cluster_config.peering_name,
            network=container_network.name,
            import_custom_routes=True,
            export_custom_routes=True)
        ```

        ## Import

        NetworkPeeringRoutesConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{project}}/{{network}}/{{peering}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{network}}/{{peering}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] export_custom_routes: Whether to export the custom routes to the peer network.
        :param pulumi.Input[bool] import_custom_routes: Whether to import the custom routes to the peer network.
        :param pulumi.Input[str] network: The name of the primary network for the peering.
        :param pulumi.Input[str] peering: Name of the peering.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkPeeringRoutesConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage a network peering's route settings without managing the peering as
        a whole. This resource is primarily intended for use with GCP-generated
        peerings that shouldn't otherwise be managed by other tools. Deleting this
        resource is a no-op and the peering will not be modified.

        To get more information about NetworkPeeringRoutesConfig, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks/updatePeering)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vpc/docs/vpc-peering)

        ## Example Usage
        ### Network Peering Routes Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        network_primary = gcp.compute.Network("networkPrimary", auto_create_subnetworks=False)
        network_secondary = gcp.compute.Network("networkSecondary", auto_create_subnetworks=False)
        peering_primary = gcp.compute.NetworkPeering("peeringPrimary",
            network=network_primary.id,
            peer_network=network_secondary.id,
            import_custom_routes=True,
            export_custom_routes=True)
        peering_primary_routes = gcp.compute.NetworkPeeringRoutesConfig("peeringPrimaryRoutes",
            peering=peering_primary.name,
            network=network_primary.name,
            import_custom_routes=True,
            export_custom_routes=True)
        peering_secondary = gcp.compute.NetworkPeering("peeringSecondary",
            network=network_secondary.id,
            peer_network=network_primary.id)
        ```
        ### Network Peering Routes Config Gke

        ```python
        import pulumi
        import pulumi_gcp as gcp

        container_network = gcp.compute.Network("containerNetwork", auto_create_subnetworks=False)
        container_subnetwork = gcp.compute.Subnetwork("containerSubnetwork",
            region="us-central1",
            network=container_network.name,
            ip_cidr_range="10.0.36.0/24",
            private_ip_google_access=True,
            secondary_ip_ranges=[
                gcp.compute.SubnetworkSecondaryIpRangeArgs(
                    range_name="pod",
                    ip_cidr_range="10.0.0.0/19",
                ),
                gcp.compute.SubnetworkSecondaryIpRangeArgs(
                    range_name="svc",
                    ip_cidr_range="10.0.32.0/22",
                ),
            ])
        private_cluster = gcp.container.Cluster("privateCluster",
            location="us-central1-a",
            initial_node_count=1,
            network=container_network.name,
            subnetwork=container_subnetwork.name,
            private_cluster_config=gcp.container.ClusterPrivateClusterConfigArgs(
                enable_private_endpoint=True,
                enable_private_nodes=True,
                master_ipv4_cidr_block="10.42.0.0/28",
            ),
            master_authorized_networks_config=gcp.container.ClusterMasterAuthorizedNetworksConfigArgs(),
            ip_allocation_policy=gcp.container.ClusterIpAllocationPolicyArgs(
                cluster_secondary_range_name=container_subnetwork.secondary_ip_ranges[0].range_name,
                services_secondary_range_name=container_subnetwork.secondary_ip_ranges[1].range_name,
            ))
        peering_gke_routes = gcp.compute.NetworkPeeringRoutesConfig("peeringGkeRoutes",
            peering=private_cluster.private_cluster_config.peering_name,
            network=container_network.name,
            import_custom_routes=True,
            export_custom_routes=True)
        ```

        ## Import

        NetworkPeeringRoutesConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default projects/{{project}}/global/networks/{{network}}/networkPeerings/{{peering}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{project}}/{{network}}/{{peering}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig default {{network}}/{{peering}}
        ```

        :param str resource_name: The name of the resource.
        :param NetworkPeeringRoutesConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkPeeringRoutesConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 export_custom_routes: Optional[pulumi.Input[bool]] = None,
                 import_custom_routes: Optional[pulumi.Input[bool]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 peering: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkPeeringRoutesConfigArgs.__new__(NetworkPeeringRoutesConfigArgs)

            if export_custom_routes is None and not opts.urn:
                raise TypeError("Missing required property 'export_custom_routes'")
            __props__.__dict__["export_custom_routes"] = export_custom_routes
            if import_custom_routes is None and not opts.urn:
                raise TypeError("Missing required property 'import_custom_routes'")
            __props__.__dict__["import_custom_routes"] = import_custom_routes
            if network is None and not opts.urn:
                raise TypeError("Missing required property 'network'")
            __props__.__dict__["network"] = network
            if peering is None and not opts.urn:
                raise TypeError("Missing required property 'peering'")
            __props__.__dict__["peering"] = peering
            __props__.__dict__["project"] = project
        super(NetworkPeeringRoutesConfig, __self__).__init__(
            'gcp:compute/networkPeeringRoutesConfig:NetworkPeeringRoutesConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            export_custom_routes: Optional[pulumi.Input[bool]] = None,
            import_custom_routes: Optional[pulumi.Input[bool]] = None,
            network: Optional[pulumi.Input[str]] = None,
            peering: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'NetworkPeeringRoutesConfig':
        """
        Get an existing NetworkPeeringRoutesConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] export_custom_routes: Whether to export the custom routes to the peer network.
        :param pulumi.Input[bool] import_custom_routes: Whether to import the custom routes to the peer network.
        :param pulumi.Input[str] network: The name of the primary network for the peering.
        :param pulumi.Input[str] peering: Name of the peering.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkPeeringRoutesConfigState.__new__(_NetworkPeeringRoutesConfigState)

        __props__.__dict__["export_custom_routes"] = export_custom_routes
        __props__.__dict__["import_custom_routes"] = import_custom_routes
        __props__.__dict__["network"] = network
        __props__.__dict__["peering"] = peering
        __props__.__dict__["project"] = project
        return NetworkPeeringRoutesConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="exportCustomRoutes")
    def export_custom_routes(self) -> pulumi.Output[bool]:
        """
        Whether to export the custom routes to the peer network.
        """
        return pulumi.get(self, "export_custom_routes")

    @property
    @pulumi.getter(name="importCustomRoutes")
    def import_custom_routes(self) -> pulumi.Output[bool]:
        """
        Whether to import the custom routes to the peer network.
        """
        return pulumi.get(self, "import_custom_routes")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[str]:
        """
        The name of the primary network for the peering.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def peering(self) -> pulumi.Output[str]:
        """
        Name of the peering.
        """
        return pulumi.get(self, "peering")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

