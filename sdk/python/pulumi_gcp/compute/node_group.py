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

__all__ = ['NodeGroupArgs', 'NodeGroup']

@pulumi.input_type
class NodeGroupArgs:
    def __init__(__self__, *,
                 node_template: pulumi.Input[str],
                 autoscaling_policy: Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 initial_size: Optional[pulumi.Input[int]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NodeGroup resource.
        :param pulumi.Input[str] node_template: The URL of the node template to which this node group belongs.
        :param pulumi.Input['NodeGroupAutoscalingPolicyArgs'] autoscaling_policy: If you use sole-tenant nodes for your workloads, you can use the node
               group autoscaler to automatically manage the sizes of your node groups.
               Structure is documented below.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[int] initial_size: The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] maintenance_policy: Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        :param pulumi.Input['NodeGroupMaintenanceWindowArgs'] maintenance_window: contains properties for the timeframe of maintenance
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[int] size: The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] zone: Zone where this node group is located
        """
        pulumi.set(__self__, "node_template", node_template)
        if autoscaling_policy is not None:
            pulumi.set(__self__, "autoscaling_policy", autoscaling_policy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if initial_size is not None:
            pulumi.set(__self__, "initial_size", initial_size)
        if maintenance_policy is not None:
            pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="nodeTemplate")
    def node_template(self) -> pulumi.Input[str]:
        """
        The URL of the node template to which this node group belongs.
        """
        return pulumi.get(self, "node_template")

    @node_template.setter
    def node_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_template", value)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']]:
        """
        If you use sole-tenant nodes for your workloads, you can use the node
        group autoscaler to automatically manage the sizes of your node groups.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscaling_policy")

    @autoscaling_policy.setter
    def autoscaling_policy(self, value: Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']]):
        pulumi.set(self, "autoscaling_policy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional textual description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="initialSize")
    def initial_size(self) -> Optional[pulumi.Input[int]]:
        """
        The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        """
        return pulumi.get(self, "initial_size")

    @initial_size.setter
    def initial_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "initial_size", value)

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        """
        return pulumi.get(self, "maintenance_policy")

    @maintenance_policy.setter
    def maintenance_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maintenance_policy", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']]:
        """
        contains properties for the timeframe of maintenance
        Structure is documented below.
        """
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource.
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
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where this node group is located
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _NodeGroupState:
    def __init__(__self__, *,
                 autoscaling_policy: Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 initial_size: Optional[pulumi.Input[int]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_template: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NodeGroup resources.
        :param pulumi.Input['NodeGroupAutoscalingPolicyArgs'] autoscaling_policy: If you use sole-tenant nodes for your workloads, you can use the node
               group autoscaler to automatically manage the sizes of your node groups.
               Structure is documented below.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[int] initial_size: The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] maintenance_policy: Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        :param pulumi.Input['NodeGroupMaintenanceWindowArgs'] maintenance_window: contains properties for the timeframe of maintenance
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] node_template: The URL of the node template to which this node group belongs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[int] size: The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] zone: Zone where this node group is located
        """
        if autoscaling_policy is not None:
            pulumi.set(__self__, "autoscaling_policy", autoscaling_policy)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if initial_size is not None:
            pulumi.set(__self__, "initial_size", initial_size)
        if maintenance_policy is not None:
            pulumi.set(__self__, "maintenance_policy", maintenance_policy)
        if maintenance_window is not None:
            pulumi.set(__self__, "maintenance_window", maintenance_window)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if node_template is not None:
            pulumi.set(__self__, "node_template", node_template)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']]:
        """
        If you use sole-tenant nodes for your workloads, you can use the node
        group autoscaler to automatically manage the sizes of your node groups.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscaling_policy")

    @autoscaling_policy.setter
    def autoscaling_policy(self, value: Optional[pulumi.Input['NodeGroupAutoscalingPolicyArgs']]):
        pulumi.set(self, "autoscaling_policy", value)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional textual description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="initialSize")
    def initial_size(self) -> Optional[pulumi.Input[int]]:
        """
        The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        """
        return pulumi.get(self, "initial_size")

    @initial_size.setter
    def initial_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "initial_size", value)

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        """
        return pulumi.get(self, "maintenance_policy")

    @maintenance_policy.setter
    def maintenance_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "maintenance_policy", value)

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']]:
        """
        contains properties for the timeframe of maintenance
        Structure is documented below.
        """
        return pulumi.get(self, "maintenance_window")

    @maintenance_window.setter
    def maintenance_window(self, value: Optional[pulumi.Input['NodeGroupMaintenanceWindowArgs']]):
        pulumi.set(self, "maintenance_window", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nodeTemplate")
    def node_template(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the node template to which this node group belongs.
        """
        return pulumi.get(self, "node_template")

    @node_template.setter
    def node_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_template", value)

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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Zone where this node group is located
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class NodeGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 initial_size: Optional[pulumi.Input[int]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_template: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a NodeGroup resource to manage a group of sole-tenant nodes.

        To get more information about NodeGroup, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
        * How-to Guides
            * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)

        > **Warning:** Due to limitations of the API, this provider cannot update the
        number of nodes in a node group and changes to node group size either
        through provider config or through external changes will cause
        the provider to delete and recreate the node group.

        ## Example Usage
        ### Node Group Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        soletenant_tmpl = gcp.compute.NodeTemplate("soletenant-tmpl",
            region="us-central1",
            node_type="n1-node-96-624")
        nodes = gcp.compute.NodeGroup("nodes",
            zone="us-central1-a",
            description="example google_compute_node_group for the Google Provider",
            size=1,
            node_template=soletenant_tmpl.id)
        ```
        ### Node Group Autoscaling Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        soletenant_tmpl = gcp.compute.NodeTemplate("soletenant-tmpl",
            region="us-central1",
            node_type="n1-node-96-624")
        nodes = gcp.compute.NodeGroup("nodes",
            zone="us-central1-a",
            description="example google_compute_node_group for Google Provider",
            maintenance_policy="RESTART_IN_PLACE",
            maintenance_window=gcp.compute.NodeGroupMaintenanceWindowArgs(
                start_time="08:00",
            ),
            initial_size=1,
            node_template=soletenant_tmpl.id,
            autoscaling_policy=gcp.compute.NodeGroupAutoscalingPolicyArgs(
                mode="ONLY_SCALE_OUT",
                min_nodes=1,
                max_nodes=10,
            ))
        ```

        ## Import

        NodeGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']] autoscaling_policy: If you use sole-tenant nodes for your workloads, you can use the node
               group autoscaler to automatically manage the sizes of your node groups.
               Structure is documented below.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[int] initial_size: The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] maintenance_policy: Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        :param pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']] maintenance_window: contains properties for the timeframe of maintenance
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] node_template: The URL of the node template to which this node group belongs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[int] size: The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] zone: Zone where this node group is located
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a NodeGroup resource to manage a group of sole-tenant nodes.

        To get more information about NodeGroup, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
        * How-to Guides
            * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)

        > **Warning:** Due to limitations of the API, this provider cannot update the
        number of nodes in a node group and changes to node group size either
        through provider config or through external changes will cause
        the provider to delete and recreate the node group.

        ## Example Usage
        ### Node Group Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        soletenant_tmpl = gcp.compute.NodeTemplate("soletenant-tmpl",
            region="us-central1",
            node_type="n1-node-96-624")
        nodes = gcp.compute.NodeGroup("nodes",
            zone="us-central1-a",
            description="example google_compute_node_group for the Google Provider",
            size=1,
            node_template=soletenant_tmpl.id)
        ```
        ### Node Group Autoscaling Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        soletenant_tmpl = gcp.compute.NodeTemplate("soletenant-tmpl",
            region="us-central1",
            node_type="n1-node-96-624")
        nodes = gcp.compute.NodeGroup("nodes",
            zone="us-central1-a",
            description="example google_compute_node_group for Google Provider",
            maintenance_policy="RESTART_IN_PLACE",
            maintenance_window=gcp.compute.NodeGroupMaintenanceWindowArgs(
                start_time="08:00",
            ),
            initial_size=1,
            node_template=soletenant_tmpl.id,
            autoscaling_policy=gcp.compute.NodeGroupAutoscalingPolicyArgs(
                mode="ONLY_SCALE_OUT",
                min_nodes=1,
                max_nodes=10,
            ))
        ```

        ## Import

        NodeGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeGroup:NodeGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param NodeGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 initial_size: Optional[pulumi.Input[int]] = None,
                 maintenance_policy: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_template: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodeGroupArgs.__new__(NodeGroupArgs)

            __props__.__dict__["autoscaling_policy"] = autoscaling_policy
            __props__.__dict__["description"] = description
            __props__.__dict__["initial_size"] = initial_size
            __props__.__dict__["maintenance_policy"] = maintenance_policy
            __props__.__dict__["maintenance_window"] = maintenance_window
            __props__.__dict__["name"] = name
            if node_template is None and not opts.urn:
                raise TypeError("Missing required property 'node_template'")
            __props__.__dict__["node_template"] = node_template
            __props__.__dict__["project"] = project
            __props__.__dict__["size"] = size
            __props__.__dict__["zone"] = zone
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["self_link"] = None
        super(NodeGroup, __self__).__init__(
            'gcp:compute/nodeGroup:NodeGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling_policy: Optional[pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            initial_size: Optional[pulumi.Input[int]] = None,
            maintenance_policy: Optional[pulumi.Input[str]] = None,
            maintenance_window: Optional[pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_template: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'NodeGroup':
        """
        Get an existing NodeGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NodeGroupAutoscalingPolicyArgs']] autoscaling_policy: If you use sole-tenant nodes for your workloads, you can use the node
               group autoscaler to automatically manage the sizes of your node groups.
               Structure is documented below.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[int] initial_size: The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] maintenance_policy: Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        :param pulumi.Input[pulumi.InputType['NodeGroupMaintenanceWindowArgs']] maintenance_window: contains properties for the timeframe of maintenance
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[str] node_template: The URL of the node template to which this node group belongs.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[int] size: The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        :param pulumi.Input[str] zone: Zone where this node group is located
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NodeGroupState.__new__(_NodeGroupState)

        __props__.__dict__["autoscaling_policy"] = autoscaling_policy
        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["description"] = description
        __props__.__dict__["initial_size"] = initial_size
        __props__.__dict__["maintenance_policy"] = maintenance_policy
        __props__.__dict__["maintenance_window"] = maintenance_window
        __props__.__dict__["name"] = name
        __props__.__dict__["node_template"] = node_template
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["size"] = size
        __props__.__dict__["zone"] = zone
        return NodeGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> pulumi.Output['outputs.NodeGroupAutoscalingPolicy']:
        """
        If you use sole-tenant nodes for your workloads, you can use the node
        group autoscaler to automatically manage the sizes of your node groups.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscaling_policy")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional textual description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="initialSize")
    def initial_size(self) -> pulumi.Output[Optional[int]]:
        """
        The initial number of nodes in the node group. One of `initial_size` or `size` must be specified.
        """
        return pulumi.get(self, "initial_size")

    @property
    @pulumi.getter(name="maintenancePolicy")
    def maintenance_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
        """
        return pulumi.get(self, "maintenance_policy")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> pulumi.Output[Optional['outputs.NodeGroupMaintenanceWindow']]:
        """
        contains properties for the timeframe of maintenance
        Structure is documented below.
        """
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeTemplate")
    def node_template(self) -> pulumi.Output[str]:
        """
        The URL of the node template to which this node group belongs.
        """
        return pulumi.get(self, "node_template")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The total number of nodes in the node group. One of `initial_size` or `size` must be specified.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Zone where this node group is located
        """
        return pulumi.get(self, "zone")

