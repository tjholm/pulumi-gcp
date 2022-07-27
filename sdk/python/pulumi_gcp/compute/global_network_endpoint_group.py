# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GlobalNetworkEndpointGroupArgs', 'GlobalNetworkEndpointGroup']

@pulumi.input_type
class GlobalNetworkEndpointGroupArgs:
    def __init__(__self__, *,
                 network_endpoint_type: pulumi.Input[str],
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GlobalNetworkEndpointGroup resource.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
               Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
               network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "network_endpoint_type", network_endpoint_type)
        if default_port is not None:
            pulumi.set(__self__, "default_port", default_port)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> pulumi.Input[str]:
        """
        Type of network endpoints in this network endpoint group.
        Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        """
        return pulumi.get(self, "network_endpoint_type")

    @network_endpoint_type.setter
    def network_endpoint_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_endpoint_type", value)

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> Optional[pulumi.Input[int]]:
        """
        The default port used if the port number is not specified in the
        network endpoint.
        """
        return pulumi.get(self, "default_port")

    @default_port.setter
    def default_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_port", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
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
class _GlobalNetworkEndpointGroupState:
    def __init__(__self__, *,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GlobalNetworkEndpointGroup resources.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
               network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
               Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        if default_port is not None:
            pulumi.set(__self__, "default_port", default_port)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_endpoint_type is not None:
            pulumi.set(__self__, "network_endpoint_type", network_endpoint_type)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> Optional[pulumi.Input[int]]:
        """
        The default port used if the port number is not specified in the
        network endpoint.
        """
        return pulumi.get(self, "default_port")

    @default_port.setter
    def default_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_port", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of network endpoints in this network endpoint group.
        Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        """
        return pulumi.get(self, "network_endpoint_type")

    @network_endpoint_type.setter
    def network_endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_endpoint_type", value)

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


class GlobalNetworkEndpointGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A global network endpoint group contains endpoints that reside outside of Google Cloud.
        Currently a global network endpoint group can only support a single endpoint.

        Recreating a global network endpoint group that's in use by another resource will give a
        `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
        to avoid this type of error.

        To get more information about GlobalNetworkEndpointGroup, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/internet-neg-concepts)

        ## Example Usage
        ### Global Network Endpoint Group

        ```python
        import pulumi
        import pulumi_gcp as gcp

        neg = gcp.compute.GlobalNetworkEndpointGroup("neg",
            default_port=90,
            network_endpoint_type="INTERNET_FQDN_PORT")
        ```
        ### Global Network Endpoint Group Ip Address

        ```python
        import pulumi
        import pulumi_gcp as gcp

        neg = gcp.compute.GlobalNetworkEndpointGroup("neg",
            default_port=90,
            network_endpoint_type="INTERNET_IP_PORT")
        ```

        ## Import

        GlobalNetworkEndpointGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default projects/{{project}}/global/networkEndpointGroups/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
               network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
               Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GlobalNetworkEndpointGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A global network endpoint group contains endpoints that reside outside of Google Cloud.
        Currently a global network endpoint group can only support a single endpoint.

        Recreating a global network endpoint group that's in use by another resource will give a
        `resourceInUseByAnotherResource` error. Use `lifecycle.create_before_destroy`
        to avoid this type of error.

        To get more information about GlobalNetworkEndpointGroup, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/internet-neg-concepts)

        ## Example Usage
        ### Global Network Endpoint Group

        ```python
        import pulumi
        import pulumi_gcp as gcp

        neg = gcp.compute.GlobalNetworkEndpointGroup("neg",
            default_port=90,
            network_endpoint_type="INTERNET_FQDN_PORT")
        ```
        ### Global Network Endpoint Group Ip Address

        ```python
        import pulumi
        import pulumi_gcp as gcp

        neg = gcp.compute.GlobalNetworkEndpointGroup("neg",
            default_port=90,
            network_endpoint_type="INTERNET_IP_PORT")
        ```

        ## Import

        GlobalNetworkEndpointGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default projects/{{project}}/global/networkEndpointGroups/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param GlobalNetworkEndpointGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GlobalNetworkEndpointGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_port: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_endpoint_type: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GlobalNetworkEndpointGroupArgs.__new__(GlobalNetworkEndpointGroupArgs)

            __props__.__dict__["default_port"] = default_port
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if network_endpoint_type is None and not opts.urn:
                raise TypeError("Missing required property 'network_endpoint_type'")
            __props__.__dict__["network_endpoint_type"] = network_endpoint_type
            __props__.__dict__["project"] = project
            __props__.__dict__["self_link"] = None
        super(GlobalNetworkEndpointGroup, __self__).__init__(
            'gcp:compute/globalNetworkEndpointGroup:GlobalNetworkEndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_port: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_endpoint_type: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'GlobalNetworkEndpointGroup':
        """
        Get an existing GlobalNetworkEndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_port: The default port used if the port number is not specified in the
               network endpoint.
        :param pulumi.Input[str] description: An optional description of this resource. Provide this property when
               you create the resource.
        :param pulumi.Input[str] name: Name of the resource; provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035. Specifically, the name must be 1-63 characters long and match
               the regular expression `a-z?` which means the
               first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last
               character, which cannot be a dash.
        :param pulumi.Input[str] network_endpoint_type: Type of network endpoints in this network endpoint group.
               Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GlobalNetworkEndpointGroupState.__new__(_GlobalNetworkEndpointGroupState)

        __props__.__dict__["default_port"] = default_port
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["network_endpoint_type"] = network_endpoint_type
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        return GlobalNetworkEndpointGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultPort")
    def default_port(self) -> pulumi.Output[Optional[int]]:
        """
        The default port used if the port number is not specified in the
        network endpoint.
        """
        return pulumi.get(self, "default_port")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of this resource. Provide this property when
        you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource; provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035. Specifically, the name must be 1-63 characters long and match
        the regular expression `a-z?` which means the
        first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the last
        character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkEndpointType")
    def network_endpoint_type(self) -> pulumi.Output[str]:
        """
        Type of network endpoints in this network endpoint group.
        Possible values are `INTERNET_IP_PORT` and `INTERNET_FQDN_PORT`.
        """
        return pulumi.get(self, "network_endpoint_type")

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

