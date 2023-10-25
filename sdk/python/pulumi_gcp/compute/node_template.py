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

__all__ = ['NodeTemplateArgs', 'NodeTemplate']

@pulumi.input_type
class NodeTemplateArgs:
    def __init__(__self__, *,
                 cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 node_type_flexibility: Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 server_binding: Optional[pulumi.Input['NodeTemplateServerBindingArgs']] = None):
        """
        The set of arguments for constructing a NodeTemplate resource.
        :param pulumi.Input[str] cpu_overcommit_type: CPU overcommit.
               Default value is `NONE`.
               Possible values are: `ENABLED`, `NONE`.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] node_affinity_labels: Labels to use for node affinity, which will be used in
               instance scheduling.
        :param pulumi.Input[str] node_type: Node type to use for nodes group that are created from this template.
               Only one of nodeTypeFlexibility and nodeType can be specified.
        :param pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs'] node_type_flexibility: Flexible properties for the desired node type. Node groups that
               use this node template will create nodes of a type that matches
               these properties. Only one of nodeTypeFlexibility and nodeType can
               be specified.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where nodes using the node template will be created.
               If it is not provided, the provider region is used.
        :param pulumi.Input['NodeTemplateServerBindingArgs'] server_binding: The server binding policy for nodes using this template. Determines
               where the nodes should restart following a maintenance event.
               Structure is documented below.
        """
        NodeTemplateArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            cpu_overcommit_type=cpu_overcommit_type,
            description=description,
            name=name,
            node_affinity_labels=node_affinity_labels,
            node_type=node_type,
            node_type_flexibility=node_type_flexibility,
            project=project,
            region=region,
            server_binding=server_binding,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             node_type: Optional[pulumi.Input[str]] = None,
             node_type_flexibility: Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             server_binding: Optional[pulumi.Input['NodeTemplateServerBindingArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if cpu_overcommit_type is None and 'cpuOvercommitType' in kwargs:
            cpu_overcommit_type = kwargs['cpuOvercommitType']
        if node_affinity_labels is None and 'nodeAffinityLabels' in kwargs:
            node_affinity_labels = kwargs['nodeAffinityLabels']
        if node_type is None and 'nodeType' in kwargs:
            node_type = kwargs['nodeType']
        if node_type_flexibility is None and 'nodeTypeFlexibility' in kwargs:
            node_type_flexibility = kwargs['nodeTypeFlexibility']
        if server_binding is None and 'serverBinding' in kwargs:
            server_binding = kwargs['serverBinding']

        if cpu_overcommit_type is not None:
            _setter("cpu_overcommit_type", cpu_overcommit_type)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if node_affinity_labels is not None:
            _setter("node_affinity_labels", node_affinity_labels)
        if node_type is not None:
            _setter("node_type", node_type)
        if node_type_flexibility is not None:
            _setter("node_type_flexibility", node_type_flexibility)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if server_binding is not None:
            _setter("server_binding", server_binding)

    @property
    @pulumi.getter(name="cpuOvercommitType")
    def cpu_overcommit_type(self) -> Optional[pulumi.Input[str]]:
        """
        CPU overcommit.
        Default value is `NONE`.
        Possible values are: `ENABLED`, `NONE`.
        """
        return pulumi.get(self, "cpu_overcommit_type")

    @cpu_overcommit_type.setter
    def cpu_overcommit_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cpu_overcommit_type", value)

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
    @pulumi.getter(name="nodeAffinityLabels")
    def node_affinity_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to use for node affinity, which will be used in
        instance scheduling.
        """
        return pulumi.get(self, "node_affinity_labels")

    @node_affinity_labels.setter
    def node_affinity_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "node_affinity_labels", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        Node type to use for nodes group that are created from this template.
        Only one of nodeTypeFlexibility and nodeType can be specified.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="nodeTypeFlexibility")
    def node_type_flexibility(self) -> Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']]:
        """
        Flexible properties for the desired node type. Node groups that
        use this node template will create nodes of a type that matches
        these properties. Only one of nodeTypeFlexibility and nodeType can
        be specified.
        Structure is documented below.
        """
        return pulumi.get(self, "node_type_flexibility")

    @node_type_flexibility.setter
    def node_type_flexibility(self, value: Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']]):
        pulumi.set(self, "node_type_flexibility", value)

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
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region where nodes using the node template will be created.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serverBinding")
    def server_binding(self) -> Optional[pulumi.Input['NodeTemplateServerBindingArgs']]:
        """
        The server binding policy for nodes using this template. Determines
        where the nodes should restart following a maintenance event.
        Structure is documented below.
        """
        return pulumi.get(self, "server_binding")

    @server_binding.setter
    def server_binding(self, value: Optional[pulumi.Input['NodeTemplateServerBindingArgs']]):
        pulumi.set(self, "server_binding", value)


@pulumi.input_type
class _NodeTemplateState:
    def __init__(__self__, *,
                 cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 node_type_flexibility: Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 server_binding: Optional[pulumi.Input['NodeTemplateServerBindingArgs']] = None):
        """
        Input properties used for looking up and filtering NodeTemplate resources.
        :param pulumi.Input[str] cpu_overcommit_type: CPU overcommit.
               Default value is `NONE`.
               Possible values are: `ENABLED`, `NONE`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] node_affinity_labels: Labels to use for node affinity, which will be used in
               instance scheduling.
        :param pulumi.Input[str] node_type: Node type to use for nodes group that are created from this template.
               Only one of nodeTypeFlexibility and nodeType can be specified.
        :param pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs'] node_type_flexibility: Flexible properties for the desired node type. Node groups that
               use this node template will create nodes of a type that matches
               these properties. Only one of nodeTypeFlexibility and nodeType can
               be specified.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where nodes using the node template will be created.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input['NodeTemplateServerBindingArgs'] server_binding: The server binding policy for nodes using this template. Determines
               where the nodes should restart following a maintenance event.
               Structure is documented below.
        """
        _NodeTemplateState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            cpu_overcommit_type=cpu_overcommit_type,
            creation_timestamp=creation_timestamp,
            description=description,
            name=name,
            node_affinity_labels=node_affinity_labels,
            node_type=node_type,
            node_type_flexibility=node_type_flexibility,
            project=project,
            region=region,
            self_link=self_link,
            server_binding=server_binding,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
             creation_timestamp: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             node_type: Optional[pulumi.Input[str]] = None,
             node_type_flexibility: Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             self_link: Optional[pulumi.Input[str]] = None,
             server_binding: Optional[pulumi.Input['NodeTemplateServerBindingArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if cpu_overcommit_type is None and 'cpuOvercommitType' in kwargs:
            cpu_overcommit_type = kwargs['cpuOvercommitType']
        if creation_timestamp is None and 'creationTimestamp' in kwargs:
            creation_timestamp = kwargs['creationTimestamp']
        if node_affinity_labels is None and 'nodeAffinityLabels' in kwargs:
            node_affinity_labels = kwargs['nodeAffinityLabels']
        if node_type is None and 'nodeType' in kwargs:
            node_type = kwargs['nodeType']
        if node_type_flexibility is None and 'nodeTypeFlexibility' in kwargs:
            node_type_flexibility = kwargs['nodeTypeFlexibility']
        if self_link is None and 'selfLink' in kwargs:
            self_link = kwargs['selfLink']
        if server_binding is None and 'serverBinding' in kwargs:
            server_binding = kwargs['serverBinding']

        if cpu_overcommit_type is not None:
            _setter("cpu_overcommit_type", cpu_overcommit_type)
        if creation_timestamp is not None:
            _setter("creation_timestamp", creation_timestamp)
        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if node_affinity_labels is not None:
            _setter("node_affinity_labels", node_affinity_labels)
        if node_type is not None:
            _setter("node_type", node_type)
        if node_type_flexibility is not None:
            _setter("node_type_flexibility", node_type_flexibility)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if self_link is not None:
            _setter("self_link", self_link)
        if server_binding is not None:
            _setter("server_binding", server_binding)

    @property
    @pulumi.getter(name="cpuOvercommitType")
    def cpu_overcommit_type(self) -> Optional[pulumi.Input[str]]:
        """
        CPU overcommit.
        Default value is `NONE`.
        Possible values are: `ENABLED`, `NONE`.
        """
        return pulumi.get(self, "cpu_overcommit_type")

    @cpu_overcommit_type.setter
    def cpu_overcommit_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cpu_overcommit_type", value)

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
    @pulumi.getter(name="nodeAffinityLabels")
    def node_affinity_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels to use for node affinity, which will be used in
        instance scheduling.
        """
        return pulumi.get(self, "node_affinity_labels")

    @node_affinity_labels.setter
    def node_affinity_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "node_affinity_labels", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        Node type to use for nodes group that are created from this template.
        Only one of nodeTypeFlexibility and nodeType can be specified.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="nodeTypeFlexibility")
    def node_type_flexibility(self) -> Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']]:
        """
        Flexible properties for the desired node type. Node groups that
        use this node template will create nodes of a type that matches
        these properties. Only one of nodeTypeFlexibility and nodeType can
        be specified.
        Structure is documented below.
        """
        return pulumi.get(self, "node_type_flexibility")

    @node_type_flexibility.setter
    def node_type_flexibility(self, value: Optional[pulumi.Input['NodeTemplateNodeTypeFlexibilityArgs']]):
        pulumi.set(self, "node_type_flexibility", value)

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
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region where nodes using the node template will be created.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

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
    @pulumi.getter(name="serverBinding")
    def server_binding(self) -> Optional[pulumi.Input['NodeTemplateServerBindingArgs']]:
        """
        The server binding policy for nodes using this template. Determines
        where the nodes should restart following a maintenance event.
        Structure is documented below.
        """
        return pulumi.get(self, "server_binding")

    @server_binding.setter
    def server_binding(self, value: Optional[pulumi.Input['NodeTemplateServerBindingArgs']]):
        pulumi.set(self, "server_binding", value)


class NodeTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 node_type_flexibility: Optional[pulumi.Input[pulumi.InputType['NodeTemplateNodeTypeFlexibilityArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 server_binding: Optional[pulumi.Input[pulumi.InputType['NodeTemplateServerBindingArgs']]] = None,
                 __props__=None):
        """
        Represents a NodeTemplate resource. Node templates specify properties
        for creating sole-tenant nodes, such as node type, vCPU and memory
        requirements, node affinity labels, and region.

        To get more information about NodeTemplate, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
        * How-to Guides
            * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)

        ## Example Usage
        ### Node Template Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        template = gcp.compute.NodeTemplate("template",
            node_type="n1-node-96-624",
            region="us-central1")
        ```
        ### Node Template Server Binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        central1a = gcp.compute.get_node_types(zone="us-central1-a")
        template = gcp.compute.NodeTemplate("template",
            node_affinity_labels={
                "foo": "baz",
            },
            node_type="n1-node-96-624",
            region="us-central1",
            server_binding=gcp.compute.NodeTemplateServerBindingArgs(
                type="RESTART_NODE_ON_MINIMAL_SERVERS",
            ))
        ```

        ## Import

        NodeTemplate can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cpu_overcommit_type: CPU overcommit.
               Default value is `NONE`.
               Possible values are: `ENABLED`, `NONE`.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] node_affinity_labels: Labels to use for node affinity, which will be used in
               instance scheduling.
        :param pulumi.Input[str] node_type: Node type to use for nodes group that are created from this template.
               Only one of nodeTypeFlexibility and nodeType can be specified.
        :param pulumi.Input[pulumi.InputType['NodeTemplateNodeTypeFlexibilityArgs']] node_type_flexibility: Flexible properties for the desired node type. Node groups that
               use this node template will create nodes of a type that matches
               these properties. Only one of nodeTypeFlexibility and nodeType can
               be specified.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where nodes using the node template will be created.
               If it is not provided, the provider region is used.
        :param pulumi.Input[pulumi.InputType['NodeTemplateServerBindingArgs']] server_binding: The server binding policy for nodes using this template. Determines
               where the nodes should restart following a maintenance event.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NodeTemplateArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a NodeTemplate resource. Node templates specify properties
        for creating sole-tenant nodes, such as node type, vCPU and memory
        requirements, node affinity labels, and region.

        To get more information about NodeTemplate, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
        * How-to Guides
            * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)

        ## Example Usage
        ### Node Template Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        template = gcp.compute.NodeTemplate("template",
            node_type="n1-node-96-624",
            region="us-central1")
        ```
        ### Node Template Server Binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        central1a = gcp.compute.get_node_types(zone="us-central1-a")
        template = gcp.compute.NodeTemplate("template",
            node_affinity_labels={
                "foo": "baz",
            },
            node_type="n1-node-96-624",
            region="us-central1",
            server_binding=gcp.compute.NodeTemplateServerBindingArgs(
                type="RESTART_NODE_ON_MINIMAL_SERVERS",
            ))
        ```

        ## Import

        NodeTemplate can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/nodeTemplate:NodeTemplate default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param NodeTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            NodeTemplateArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 node_type_flexibility: Optional[pulumi.Input[pulumi.InputType['NodeTemplateNodeTypeFlexibilityArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 server_binding: Optional[pulumi.Input[pulumi.InputType['NodeTemplateServerBindingArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NodeTemplateArgs.__new__(NodeTemplateArgs)

            __props__.__dict__["cpu_overcommit_type"] = cpu_overcommit_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["node_affinity_labels"] = node_affinity_labels
            __props__.__dict__["node_type"] = node_type
            node_type_flexibility = _utilities.configure(node_type_flexibility, NodeTemplateNodeTypeFlexibilityArgs, True)
            __props__.__dict__["node_type_flexibility"] = node_type_flexibility
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            server_binding = _utilities.configure(server_binding, NodeTemplateServerBindingArgs, True)
            __props__.__dict__["server_binding"] = server_binding
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["self_link"] = None
        super(NodeTemplate, __self__).__init__(
            'gcp:compute/nodeTemplate:NodeTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cpu_overcommit_type: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            node_affinity_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            node_type: Optional[pulumi.Input[str]] = None,
            node_type_flexibility: Optional[pulumi.Input[pulumi.InputType['NodeTemplateNodeTypeFlexibilityArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            server_binding: Optional[pulumi.Input[pulumi.InputType['NodeTemplateServerBindingArgs']]] = None) -> 'NodeTemplate':
        """
        Get an existing NodeTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cpu_overcommit_type: CPU overcommit.
               Default value is `NONE`.
               Possible values are: `ENABLED`, `NONE`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional textual description of the resource.
        :param pulumi.Input[str] name: Name of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] node_affinity_labels: Labels to use for node affinity, which will be used in
               instance scheduling.
        :param pulumi.Input[str] node_type: Node type to use for nodes group that are created from this template.
               Only one of nodeTypeFlexibility and nodeType can be specified.
        :param pulumi.Input[pulumi.InputType['NodeTemplateNodeTypeFlexibilityArgs']] node_type_flexibility: Flexible properties for the desired node type. Node groups that
               use this node template will create nodes of a type that matches
               these properties. Only one of nodeTypeFlexibility and nodeType can
               be specified.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where nodes using the node template will be created.
               If it is not provided, the provider region is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[pulumi.InputType['NodeTemplateServerBindingArgs']] server_binding: The server binding policy for nodes using this template. Determines
               where the nodes should restart following a maintenance event.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NodeTemplateState.__new__(_NodeTemplateState)

        __props__.__dict__["cpu_overcommit_type"] = cpu_overcommit_type
        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["node_affinity_labels"] = node_affinity_labels
        __props__.__dict__["node_type"] = node_type
        __props__.__dict__["node_type_flexibility"] = node_type_flexibility
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["server_binding"] = server_binding
        return NodeTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cpuOvercommitType")
    def cpu_overcommit_type(self) -> pulumi.Output[Optional[str]]:
        """
        CPU overcommit.
        Default value is `NONE`.
        Possible values are: `ENABLED`, `NONE`.
        """
        return pulumi.get(self, "cpu_overcommit_type")

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
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeAffinityLabels")
    def node_affinity_labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels to use for node affinity, which will be used in
        instance scheduling.
        """
        return pulumi.get(self, "node_affinity_labels")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> pulumi.Output[Optional[str]]:
        """
        Node type to use for nodes group that are created from this template.
        Only one of nodeTypeFlexibility and nodeType can be specified.
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter(name="nodeTypeFlexibility")
    def node_type_flexibility(self) -> pulumi.Output[Optional['outputs.NodeTemplateNodeTypeFlexibility']]:
        """
        Flexible properties for the desired node type. Node groups that
        use this node template will create nodes of a type that matches
        these properties. Only one of nodeTypeFlexibility and nodeType can
        be specified.
        Structure is documented below.
        """
        return pulumi.get(self, "node_type_flexibility")

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
    def region(self) -> pulumi.Output[str]:
        """
        Region where nodes using the node template will be created.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="serverBinding")
    def server_binding(self) -> pulumi.Output['outputs.NodeTemplateServerBinding']:
        """
        The server binding policy for nodes using this template. Determines
        where the nodes should restart following a maintenance event.
        Structure is documented below.
        """
        return pulumi.get(self, "server_binding")

