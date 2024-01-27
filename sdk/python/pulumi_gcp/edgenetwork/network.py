# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkArgs', 'Network']

@pulumi.input_type
class NetworkArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 network_id: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Network resource.
        :param pulumi.Input[str] location: The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        :param pulumi.Input[str] network_id: A unique ID that identifies this network.
               
               
               - - -
        :param pulumi.Input[str] zone: The name of the target Distributed Cloud Edge zone.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels associated with this resource.
        :param pulumi.Input[int] mtu: IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "zone", zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        A unique ID that identifies this network.


        - - -
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        The name of the target Distributed Cloud Edge zone.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels associated with this resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

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
class _NetworkState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Network resources.
        :param pulumi.Input[str] create_time: The time when the subnet was created.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
               fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels associated with this resource.
        :param pulumi.Input[str] location: The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        :param pulumi.Input[int] mtu: IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        :param pulumi.Input[str] name: The canonical name of this resource, with format
               `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}`
        :param pulumi.Input[str] network_id: A unique ID that identifies this network.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: The time when the subnet was last updated.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
               fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        :param pulumi.Input[str] zone: The name of the target Distributed Cloud Edge zone.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the subnet was created.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels associated with this resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The canonical name of this resource, with format
        `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique ID that identifies this network.


        - - -
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

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
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the subnet was last updated.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the target Distributed Cloud Edge zone.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Network(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A Distributed Cloud Edge network, which provides L3 isolation within a zone.

        To get more information about Network, see:

        * [API documentation](https://cloud.google.com/distributed-cloud/edge/latest/docs/reference/network/rest/v1/projects.locations.zones.networks)
        * How-to Guides
            * [Create and manage networks](https://cloud.google.com/distributed-cloud/edge/latest/docs/networks#api)

        ## Example Usage
        ### Edgenetwork Network

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_network = gcp.edgenetwork.Network("exampleNetwork",
            network_id="example-network",
            location="us-west1",
            zone="",
            description="Example network.",
            mtu=9000,
            labels={
                "environment": "dev",
            })
        ```

        ## Import

        Network can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}` * `{{project}}/{{location}}/{{zone}}/{{network_id}}` * `{{location}}/{{zone}}/{{network_id}}` * `{{location}}/{{network_id}}` * `{{name}}` When using the `pulumi import` command, Network can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{project}}/{{location}}/{{zone}}/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{location}}/{{zone}}/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{location}}/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels associated with this resource.
        :param pulumi.Input[str] location: The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        :param pulumi.Input[int] mtu: IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        :param pulumi.Input[str] network_id: A unique ID that identifies this network.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] zone: The name of the target Distributed Cloud Edge zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Distributed Cloud Edge network, which provides L3 isolation within a zone.

        To get more information about Network, see:

        * [API documentation](https://cloud.google.com/distributed-cloud/edge/latest/docs/reference/network/rest/v1/projects.locations.zones.networks)
        * How-to Guides
            * [Create and manage networks](https://cloud.google.com/distributed-cloud/edge/latest/docs/networks#api)

        ## Example Usage
        ### Edgenetwork Network

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example_network = gcp.edgenetwork.Network("exampleNetwork",
            network_id="example-network",
            location="us-west1",
            zone="",
            description="Example network.",
            mtu=9000,
            labels={
                "environment": "dev",
            })
        ```

        ## Import

        Network can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}` * `{{project}}/{{location}}/{{zone}}/{{network_id}}` * `{{location}}/{{zone}}/{{network_id}}` * `{{location}}/{{network_id}}` * `{{name}}` When using the `pulumi import` command, Network can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{project}}/{{location}}/{{zone}}/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{location}}/{{zone}}/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{location}}/{{network_id}}
        ```

        ```sh
         $ pulumi import gcp:edgenetwork/network:Network default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param NetworkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkArgs.__new__(NetworkArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["mtu"] = mtu
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["project"] = project
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(Network, __self__).__init__(
            'gcp:edgenetwork/network:Network',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            mtu: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Network':
        """
        Get an existing Network resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time when the subnet was created.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
               fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels associated with this resource.
        :param pulumi.Input[str] location: The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        :param pulumi.Input[int] mtu: IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        :param pulumi.Input[str] name: The canonical name of this resource, with format
               `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}`
        :param pulumi.Input[str] network_id: A unique ID that identifies this network.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: The time when the subnet was last updated.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
               fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        :param pulumi.Input[str] zone: The name of the target Distributed Cloud Edge zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkState.__new__(_NetworkState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["mtu"] = mtu
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["project"] = project
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["zone"] = zone
        return Network(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the subnet was created.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels associated with this resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The Google Cloud region to which the target Distributed Cloud Edge zone belongs.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[Optional[int]]:
        """
        IP (L3) MTU value of the network. Default value is `1500`. Possible values are: `1500`, `9000`.
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The canonical name of this resource, with format
        `projects/{{project}}/locations/{{location}}/zones/{{zone}}/networks/{{network_id}}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        A unique ID that identifies this network.


        - - -
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the subnet was last updated.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        fractional digits. Examples: `2014-10-02T15:01:23Z` and `2014-10-02T15:01:23.045123456Z`.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The name of the target Distributed Cloud Edge zone.
        """
        return pulumi.get(self, "zone")

