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

__all__ = ['AppGatewayArgs', 'AppGateway']

@pulumi.input_type
class AppGatewayArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AppGateway resource.
        :param pulumi.Input[str] display_name: An arbitrary user-provided name for the AppGateway.
        :param pulumi.Input[str] host_type: The type of hosting used by the AppGateway.
               Default value is `HOST_TYPE_UNSPECIFIED`.
               Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: ID of the AppGateway.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the AppGateway.
        :param pulumi.Input[str] type: The type of network connectivity used by the AppGateway.
               Default value is `TYPE_UNSPECIFIED`.
               Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if host_type is not None:
            pulumi.set(__self__, "host_type", host_type)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        An arbitrary user-provided name for the AppGateway.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of hosting used by the AppGateway.
        Default value is `HOST_TYPE_UNSPECIFIED`.
        Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        """
        return pulumi.get(self, "host_type")

    @host_type.setter
    def host_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_type", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user provided metadata.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the AppGateway.


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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the AppGateway.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of network connectivity used by the AppGateway.
        Default value is `TYPE_UNSPECIFIED`.
        Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _AppGatewayState:
    def __init__(__self__, *,
                 allocated_connections: Optional[pulumi.Input[Sequence[pulumi.Input['AppGatewayAllocatedConnectionArgs']]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uri: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AppGateway resources.
        :param pulumi.Input[Sequence[pulumi.Input['AppGatewayAllocatedConnectionArgs']]] allocated_connections: A list of connections allocated for the Gateway.
               Structure is documented below.
        :param pulumi.Input[str] display_name: An arbitrary user-provided name for the AppGateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] host_type: The type of hosting used by the AppGateway.
               Default value is `HOST_TYPE_UNSPECIFIED`.
               Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: ID of the AppGateway.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] region: The region of the AppGateway.
        :param pulumi.Input[str] state: Represents the different states of a AppGateway.
        :param pulumi.Input[str] type: The type of network connectivity used by the AppGateway.
               Default value is `TYPE_UNSPECIFIED`.
               Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        :param pulumi.Input[str] uri: Server-defined URI for this resource.
        """
        if allocated_connections is not None:
            pulumi.set(__self__, "allocated_connections", allocated_connections)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if effective_labels is not None:
            pulumi.set(__self__, "effective_labels", effective_labels)
        if host_type is not None:
            pulumi.set(__self__, "host_type", host_type)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if pulumi_labels is not None:
            pulumi.set(__self__, "pulumi_labels", pulumi_labels)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uri is not None:
            pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter(name="allocatedConnections")
    def allocated_connections(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AppGatewayAllocatedConnectionArgs']]]]:
        """
        A list of connections allocated for the Gateway.
        Structure is documented below.
        """
        return pulumi.get(self, "allocated_connections")

    @allocated_connections.setter
    def allocated_connections(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AppGatewayAllocatedConnectionArgs']]]]):
        pulumi.set(self, "allocated_connections", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        An arbitrary user-provided name for the AppGateway.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @effective_labels.setter
    def effective_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "effective_labels", value)

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of hosting used by the AppGateway.
        Default value is `HOST_TYPE_UNSPECIFIED`.
        Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        """
        return pulumi.get(self, "host_type")

    @host_type.setter
    def host_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_type", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user provided metadata.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the AppGateway.


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
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @pulumi_labels.setter
    def pulumi_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "pulumi_labels", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the AppGateway.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Represents the different states of a AppGateway.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of network connectivity used by the AppGateway.
        Default value is `TYPE_UNSPECIFIED`.
        Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def uri(self) -> Optional[pulumi.Input[str]]:
        """
        Server-defined URI for this resource.
        """
        return pulumi.get(self, "uri")

    @uri.setter
    def uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uri", value)


class AppGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A BeyondCorp AppGateway resource represents a BeyondCorp protected AppGateway to a remote application. It creates
        all the necessary GCP components needed for creating a BeyondCorp protected AppGateway. Multiple connectors can be
        authorised for a single AppGateway.

        To get more information about AppGateway, see:

        * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appgateways)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)

        ## Example Usage
        ### Beyondcorp App Gateway Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        app_gateway = gcp.beyondcorp.AppGateway("appGateway",
            host_type="GCP_REGIONAL_MIG",
            region="us-central1",
            type="TCP_PROXY")
        ```
        ### Beyondcorp App Gateway Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        app_gateway = gcp.beyondcorp.AppGateway("appGateway",
            display_name="some display name",
            host_type="GCP_REGIONAL_MIG",
            labels={
                "bar": "baz",
                "foo": "bar",
            },
            region="us-central1",
            type="TCP_PROXY")
        ```

        ## Import

        AppGateway can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default projects/{{project}}/locations/{{region}}/appGateways/{{name}}
        ```

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: An arbitrary user-provided name for the AppGateway.
        :param pulumi.Input[str] host_type: The type of hosting used by the AppGateway.
               Default value is `HOST_TYPE_UNSPECIFIED`.
               Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: ID of the AppGateway.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the AppGateway.
        :param pulumi.Input[str] type: The type of network connectivity used by the AppGateway.
               Default value is `TYPE_UNSPECIFIED`.
               Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AppGatewayArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A BeyondCorp AppGateway resource represents a BeyondCorp protected AppGateway to a remote application. It creates
        all the necessary GCP components needed for creating a BeyondCorp protected AppGateway. Multiple connectors can be
        authorised for a single AppGateway.

        To get more information about AppGateway, see:

        * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appgateways)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)

        ## Example Usage
        ### Beyondcorp App Gateway Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        app_gateway = gcp.beyondcorp.AppGateway("appGateway",
            host_type="GCP_REGIONAL_MIG",
            region="us-central1",
            type="TCP_PROXY")
        ```
        ### Beyondcorp App Gateway Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        app_gateway = gcp.beyondcorp.AppGateway("appGateway",
            display_name="some display name",
            host_type="GCP_REGIONAL_MIG",
            labels={
                "bar": "baz",
                "foo": "bar",
            },
            region="us-central1",
            type="TCP_PROXY")
        ```

        ## Import

        AppGateway can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default projects/{{project}}/locations/{{region}}/appGateways/{{name}}
        ```

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param AppGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AppGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 host_type: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AppGatewayArgs.__new__(AppGatewayArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["host_type"] = host_type
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["type"] = type
            __props__.__dict__["allocated_connections"] = None
            __props__.__dict__["effective_labels"] = None
            __props__.__dict__["pulumi_labels"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uri"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["effectiveLabels", "pulumiLabels"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AppGateway, __self__).__init__(
            'gcp:beyondcorp/appGateway:AppGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocated_connections: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AppGatewayAllocatedConnectionArgs']]]]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            host_type: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'AppGateway':
        """
        Get an existing AppGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AppGatewayAllocatedConnectionArgs']]]] allocated_connections: A list of connections allocated for the Gateway.
               Structure is documented below.
        :param pulumi.Input[str] display_name: An arbitrary user-provided name for the AppGateway.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] host_type: The type of hosting used by the AppGateway.
               Default value is `HOST_TYPE_UNSPECIFIED`.
               Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user provided metadata.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: ID of the AppGateway.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] region: The region of the AppGateway.
        :param pulumi.Input[str] state: Represents the different states of a AppGateway.
        :param pulumi.Input[str] type: The type of network connectivity used by the AppGateway.
               Default value is `TYPE_UNSPECIFIED`.
               Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        :param pulumi.Input[str] uri: Server-defined URI for this resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AppGatewayState.__new__(_AppGatewayState)

        __props__.__dict__["allocated_connections"] = allocated_connections
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["effective_labels"] = effective_labels
        __props__.__dict__["host_type"] = host_type
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["pulumi_labels"] = pulumi_labels
        __props__.__dict__["region"] = region
        __props__.__dict__["state"] = state
        __props__.__dict__["type"] = type
        __props__.__dict__["uri"] = uri
        return AppGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocatedConnections")
    def allocated_connections(self) -> pulumi.Output[Sequence['outputs.AppGatewayAllocatedConnection']]:
        """
        A list of connections allocated for the Gateway.
        Structure is documented below.
        """
        return pulumi.get(self, "allocated_connections")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        An arbitrary user-provided name for the AppGateway.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @property
    @pulumi.getter(name="hostType")
    def host_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of hosting used by the AppGateway.
        Default value is `HOST_TYPE_UNSPECIFIED`.
        Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
        """
        return pulumi.get(self, "host_type")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource labels to represent user provided metadata.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        ID of the AppGateway.


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
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region of the AppGateway.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Represents the different states of a AppGateway.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of network connectivity used by the AppGateway.
        Default value is `TYPE_UNSPECIFIED`.
        Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        Server-defined URI for this resource.
        """
        return pulumi.get(self, "uri")

