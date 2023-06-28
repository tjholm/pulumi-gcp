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

__all__ = ['PrivateCloudArgs', 'PrivateCloud']

@pulumi.input_type
class PrivateCloudArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 management_cluster: pulumi.Input['PrivateCloudManagementClusterArgs'],
                 network_config: pulumi.Input['PrivateCloudNetworkConfigArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PrivateCloud resource.
        :param pulumi.Input[str] location: The location where the PrivateCloud should reside.
        :param pulumi.Input['PrivateCloudManagementClusterArgs'] management_cluster: The management cluster for this private cloud. This used for creating and managing the default cluster.
               Structure is documented below.
        :param pulumi.Input['PrivateCloudNetworkConfigArgs'] network_config: Network configuration in the consumer project with which the peering has to be done.
               Structure is documented below.
        :param pulumi.Input[str] description: User-provided description for this private cloud.
        :param pulumi.Input[str] name: The ID of the PrivateCloud.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "management_cluster", management_cluster)
        pulumi.set(__self__, "network_config", network_config)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location where the PrivateCloud should reside.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Input['PrivateCloudManagementClusterArgs']:
        """
        The management cluster for this private cloud. This used for creating and managing the default cluster.
        Structure is documented below.
        """
        return pulumi.get(self, "management_cluster")

    @management_cluster.setter
    def management_cluster(self, value: pulumi.Input['PrivateCloudManagementClusterArgs']):
        pulumi.set(self, "management_cluster", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Input['PrivateCloudNetworkConfigArgs']:
        """
        Network configuration in the consumer project with which the peering has to be done.
        Structure is documented below.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: pulumi.Input['PrivateCloudNetworkConfigArgs']):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided description for this private cloud.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the PrivateCloud.
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


@pulumi.input_type
class _PrivateCloudState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 hcxes: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudHcxArgs']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input['PrivateCloudManagementClusterArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input['PrivateCloudNetworkConfigArgs']] = None,
                 nsxes: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudNsxArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 vcenters: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudVcenterArgs']]]] = None):
        """
        Input properties used for looking up and filtering PrivateCloud resources.
        :param pulumi.Input[str] description: User-provided description for this private cloud.
        :param pulumi.Input[Sequence[pulumi.Input['PrivateCloudHcxArgs']]] hcxes: Details about a HCX Cloud Manager appliance.
               Structure is documented below.
        :param pulumi.Input[str] location: The location where the PrivateCloud should reside.
        :param pulumi.Input['PrivateCloudManagementClusterArgs'] management_cluster: The management cluster for this private cloud. This used for creating and managing the default cluster.
               Structure is documented below.
        :param pulumi.Input[str] name: The ID of the PrivateCloud.
        :param pulumi.Input['PrivateCloudNetworkConfigArgs'] network_config: Network configuration in the consumer project with which the peering has to be done.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['PrivateCloudNsxArgs']]] nsxes: Details about a NSX Manager appliance.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: State of the appliance.
               Possible values are: `ACTIVE`, `CREATING`.
        :param pulumi.Input[str] uid: System-generated unique identifier for the resource.
        :param pulumi.Input[Sequence[pulumi.Input['PrivateCloudVcenterArgs']]] vcenters: Details about a vCenter Server management appliance.
               Structure is documented below.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if hcxes is not None:
            pulumi.set(__self__, "hcxes", hcxes)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if management_cluster is not None:
            pulumi.set(__self__, "management_cluster", management_cluster)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if nsxes is not None:
            pulumi.set(__self__, "nsxes", nsxes)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if vcenters is not None:
            pulumi.set(__self__, "vcenters", vcenters)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided description for this private cloud.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def hcxes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudHcxArgs']]]]:
        """
        Details about a HCX Cloud Manager appliance.
        Structure is documented below.
        """
        return pulumi.get(self, "hcxes")

    @hcxes.setter
    def hcxes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudHcxArgs']]]]):
        pulumi.set(self, "hcxes", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the PrivateCloud should reside.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> Optional[pulumi.Input['PrivateCloudManagementClusterArgs']]:
        """
        The management cluster for this private cloud. This used for creating and managing the default cluster.
        Structure is documented below.
        """
        return pulumi.get(self, "management_cluster")

    @management_cluster.setter
    def management_cluster(self, value: Optional[pulumi.Input['PrivateCloudManagementClusterArgs']]):
        pulumi.set(self, "management_cluster", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the PrivateCloud.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['PrivateCloudNetworkConfigArgs']]:
        """
        Network configuration in the consumer project with which the peering has to be done.
        Structure is documented below.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['PrivateCloudNetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter
    def nsxes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudNsxArgs']]]]:
        """
        Details about a NSX Manager appliance.
        Structure is documented below.
        """
        return pulumi.get(self, "nsxes")

    @nsxes.setter
    def nsxes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudNsxArgs']]]]):
        pulumi.set(self, "nsxes", value)

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
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the appliance.
        Possible values are: `ACTIVE`, `CREATING`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        System-generated unique identifier for the resource.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter
    def vcenters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudVcenterArgs']]]]:
        """
        Details about a vCenter Server management appliance.
        Structure is documented below.
        """
        return pulumi.get(self, "vcenters")

    @vcenters.setter
    def vcenters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PrivateCloudVcenterArgs']]]]):
        pulumi.set(self, "vcenters", value)


class PrivateCloud(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['PrivateCloudNetworkConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Vmware Engine Private Cloud Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        pc_nw = gcp.vmwareengine.Network("pc-nw",
            location="us-west1",
            type="LEGACY",
            description="PC network description.",
            opts=pulumi.ResourceOptions(provider=google_beta))
        vmw_engine_pc = gcp.vmwareengine.PrivateCloud("vmw-engine-pc",
            location="us-west1-a",
            description="Sample test PC.",
            network_config=gcp.vmwareengine.PrivateCloudNetworkConfigArgs(
                management_cidr="192.168.30.0/24",
                vmware_engine_network=pc_nw.id,
            ),
            management_cluster=gcp.vmwareengine.PrivateCloudManagementClusterArgs(
                cluster_id="sample-mgmt-cluster",
                node_type_configs=[gcp.vmwareengine.PrivateCloudManagementClusterNodeTypeConfigArgs(
                    node_type_id="standard-72",
                    node_count=3,
                )],
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Vmware Engine Private Cloud Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        pc_nw = gcp.vmwareengine.Network("pc-nw",
            location="us-west1",
            type="LEGACY",
            description="PC network description.",
            opts=pulumi.ResourceOptions(provider=google_beta))
        vmw_engine_pc = gcp.vmwareengine.PrivateCloud("vmw-engine-pc",
            location="us-west1-a",
            description="Sample test PC.",
            network_config=gcp.vmwareengine.PrivateCloudNetworkConfigArgs(
                management_cidr="192.168.30.0/24",
                vmware_engine_network=pc_nw.id,
            ),
            management_cluster=gcp.vmwareengine.PrivateCloudManagementClusterArgs(
                cluster_id="sample-mgmt-cluster",
                node_type_configs=[gcp.vmwareengine.PrivateCloudManagementClusterNodeTypeConfigArgs(
                    node_type_id="standard-72",
                    node_count=3,
                    custom_core_count=32,
                )],
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        PrivateCloud can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:vmwareengine/privateCloud:PrivateCloud default projects/{{project}}/locations/{{location}}/privateClouds/{{name}}
        ```

        ```sh
         $ pulumi import gcp:vmwareengine/privateCloud:PrivateCloud default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:vmwareengine/privateCloud:PrivateCloud default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User-provided description for this private cloud.
        :param pulumi.Input[str] location: The location where the PrivateCloud should reside.
        :param pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']] management_cluster: The management cluster for this private cloud. This used for creating and managing the default cluster.
               Structure is documented below.
        :param pulumi.Input[str] name: The ID of the PrivateCloud.
        :param pulumi.Input[pulumi.InputType['PrivateCloudNetworkConfigArgs']] network_config: Network configuration in the consumer project with which the peering has to be done.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateCloudArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Vmware Engine Private Cloud Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        pc_nw = gcp.vmwareengine.Network("pc-nw",
            location="us-west1",
            type="LEGACY",
            description="PC network description.",
            opts=pulumi.ResourceOptions(provider=google_beta))
        vmw_engine_pc = gcp.vmwareengine.PrivateCloud("vmw-engine-pc",
            location="us-west1-a",
            description="Sample test PC.",
            network_config=gcp.vmwareengine.PrivateCloudNetworkConfigArgs(
                management_cidr="192.168.30.0/24",
                vmware_engine_network=pc_nw.id,
            ),
            management_cluster=gcp.vmwareengine.PrivateCloudManagementClusterArgs(
                cluster_id="sample-mgmt-cluster",
                node_type_configs=[gcp.vmwareengine.PrivateCloudManagementClusterNodeTypeConfigArgs(
                    node_type_id="standard-72",
                    node_count=3,
                )],
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Vmware Engine Private Cloud Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        pc_nw = gcp.vmwareengine.Network("pc-nw",
            location="us-west1",
            type="LEGACY",
            description="PC network description.",
            opts=pulumi.ResourceOptions(provider=google_beta))
        vmw_engine_pc = gcp.vmwareengine.PrivateCloud("vmw-engine-pc",
            location="us-west1-a",
            description="Sample test PC.",
            network_config=gcp.vmwareengine.PrivateCloudNetworkConfigArgs(
                management_cidr="192.168.30.0/24",
                vmware_engine_network=pc_nw.id,
            ),
            management_cluster=gcp.vmwareengine.PrivateCloudManagementClusterArgs(
                cluster_id="sample-mgmt-cluster",
                node_type_configs=[gcp.vmwareengine.PrivateCloudManagementClusterNodeTypeConfigArgs(
                    node_type_id="standard-72",
                    node_count=3,
                    custom_core_count=32,
                )],
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        PrivateCloud can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:vmwareengine/privateCloud:PrivateCloud default projects/{{project}}/locations/{{location}}/privateClouds/{{name}}
        ```

        ```sh
         $ pulumi import gcp:vmwareengine/privateCloud:PrivateCloud default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:vmwareengine/privateCloud:PrivateCloud default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param PrivateCloudArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateCloudArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['PrivateCloudNetworkConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateCloudArgs.__new__(PrivateCloudArgs)

            __props__.__dict__["description"] = description
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if management_cluster is None and not opts.urn:
                raise TypeError("Missing required property 'management_cluster'")
            __props__.__dict__["management_cluster"] = management_cluster
            __props__.__dict__["name"] = name
            if network_config is None and not opts.urn:
                raise TypeError("Missing required property 'network_config'")
            __props__.__dict__["network_config"] = network_config
            __props__.__dict__["project"] = project
            __props__.__dict__["hcxes"] = None
            __props__.__dict__["nsxes"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["vcenters"] = None
        super(PrivateCloud, __self__).__init__(
            'gcp:vmwareengine/privateCloud:PrivateCloud',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            hcxes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudHcxArgs']]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            management_cluster: Optional[pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_config: Optional[pulumi.Input[pulumi.InputType['PrivateCloudNetworkConfigArgs']]] = None,
            nsxes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudNsxArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            uid: Optional[pulumi.Input[str]] = None,
            vcenters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudVcenterArgs']]]]] = None) -> 'PrivateCloud':
        """
        Get an existing PrivateCloud resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User-provided description for this private cloud.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudHcxArgs']]]] hcxes: Details about a HCX Cloud Manager appliance.
               Structure is documented below.
        :param pulumi.Input[str] location: The location where the PrivateCloud should reside.
        :param pulumi.Input[pulumi.InputType['PrivateCloudManagementClusterArgs']] management_cluster: The management cluster for this private cloud. This used for creating and managing the default cluster.
               Structure is documented below.
        :param pulumi.Input[str] name: The ID of the PrivateCloud.
        :param pulumi.Input[pulumi.InputType['PrivateCloudNetworkConfigArgs']] network_config: Network configuration in the consumer project with which the peering has to be done.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudNsxArgs']]]] nsxes: Details about a NSX Manager appliance.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: State of the appliance.
               Possible values are: `ACTIVE`, `CREATING`.
        :param pulumi.Input[str] uid: System-generated unique identifier for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PrivateCloudVcenterArgs']]]] vcenters: Details about a vCenter Server management appliance.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PrivateCloudState.__new__(_PrivateCloudState)

        __props__.__dict__["description"] = description
        __props__.__dict__["hcxes"] = hcxes
        __props__.__dict__["location"] = location
        __props__.__dict__["management_cluster"] = management_cluster
        __props__.__dict__["name"] = name
        __props__.__dict__["network_config"] = network_config
        __props__.__dict__["nsxes"] = nsxes
        __props__.__dict__["project"] = project
        __props__.__dict__["state"] = state
        __props__.__dict__["uid"] = uid
        __props__.__dict__["vcenters"] = vcenters
        return PrivateCloud(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        User-provided description for this private cloud.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def hcxes(self) -> pulumi.Output[Sequence['outputs.PrivateCloudHcx']]:
        """
        Details about a HCX Cloud Manager appliance.
        Structure is documented below.
        """
        return pulumi.get(self, "hcxes")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location where the PrivateCloud should reside.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Output['outputs.PrivateCloudManagementCluster']:
        """
        The management cluster for this private cloud. This used for creating and managing the default cluster.
        Structure is documented below.
        """
        return pulumi.get(self, "management_cluster")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The ID of the PrivateCloud.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output['outputs.PrivateCloudNetworkConfig']:
        """
        Network configuration in the consumer project with which the peering has to be done.
        Structure is documented below.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter
    def nsxes(self) -> pulumi.Output[Sequence['outputs.PrivateCloudNsx']]:
        """
        Details about a NSX Manager appliance.
        Structure is documented below.
        """
        return pulumi.get(self, "nsxes")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the appliance.
        Possible values are: `ACTIVE`, `CREATING`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        System-generated unique identifier for the resource.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def vcenters(self) -> pulumi.Output[Sequence['outputs.PrivateCloudVcenter']]:
        """
        Details about a vCenter Server management appliance.
        Structure is documented below.
        """
        return pulumi.get(self, "vcenters")

