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

__all__ = ['GameServerClusterArgs', 'GameServerCluster']

@pulumi.input_type
class GameServerClusterArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 connection_info: pulumi.Input['GameServerClusterConnectionInfoArgs'],
                 realm_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GameServerCluster resource.
        :param pulumi.Input[str] cluster_id: Required. The resource name of the game server cluster
        :param pulumi.Input['GameServerClusterConnectionInfoArgs'] connection_info: Game server cluster connection information. This information is used to
               manage game server clusters.
               Structure is documented below.
        :param pulumi.Input[str] realm_id: The realm id of the game server realm.
        :param pulumi.Input[str] description: Human readable description of the cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this game server cluster. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: Location of the Cluster.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "connection_info", connection_info)
        pulumi.set(__self__, "realm_id", realm_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Required. The resource name of the game server cluster
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="connectionInfo")
    def connection_info(self) -> pulumi.Input['GameServerClusterConnectionInfoArgs']:
        """
        Game server cluster connection information. This information is used to
        manage game server clusters.
        Structure is documented below.
        """
        return pulumi.get(self, "connection_info")

    @connection_info.setter
    def connection_info(self, value: pulumi.Input['GameServerClusterConnectionInfoArgs']):
        pulumi.set(self, "connection_info", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm id of the game server realm.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable description of the cluster.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels associated with this game server cluster. Each label is a
        key-value pair.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the Cluster.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

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
class _GameServerClusterState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 connection_info: Optional[pulumi.Input['GameServerClusterConnectionInfoArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GameServerCluster resources.
        :param pulumi.Input[str] cluster_id: Required. The resource name of the game server cluster
        :param pulumi.Input['GameServerClusterConnectionInfoArgs'] connection_info: Game server cluster connection information. This information is used to
               manage game server clusters.
               Structure is documented below.
        :param pulumi.Input[str] description: Human readable description of the cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this game server cluster. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: Location of the Cluster.
        :param pulumi.Input[str] name: The resource id of the game server cluster, eg:
               'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
               'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] realm_id: The realm id of the game server realm.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if connection_info is not None:
            pulumi.set(__self__, "connection_info", connection_info)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The resource name of the game server cluster
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="connectionInfo")
    def connection_info(self) -> Optional[pulumi.Input['GameServerClusterConnectionInfoArgs']]:
        """
        Game server cluster connection information. This information is used to
        manage game server clusters.
        Structure is documented below.
        """
        return pulumi.get(self, "connection_info")

    @connection_info.setter
    def connection_info(self, value: Optional[pulumi.Input['GameServerClusterConnectionInfoArgs']]):
        pulumi.set(self, "connection_info", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable description of the cluster.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels associated with this game server cluster. Each label is a
        key-value pair.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the Cluster.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of the game server cluster, eg:
        'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
        'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
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
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm id of the game server realm.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class GameServerCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 connection_info: Optional[pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A game server cluster resource.

        To get more information about GameServerCluster, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms.gameServerClusters)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage

        ## Import

        GameServerCluster can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{project}}/{{location}}/{{realm_id}}/{{cluster_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{location}}/{{realm_id}}/{{cluster_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Required. The resource name of the game server cluster
        :param pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']] connection_info: Game server cluster connection information. This information is used to
               manage game server clusters.
               Structure is documented below.
        :param pulumi.Input[str] description: Human readable description of the cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this game server cluster. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: Location of the Cluster.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] realm_id: The realm id of the game server realm.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GameServerClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A game server cluster resource.

        To get more information about GameServerCluster, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.realms.gameServerClusters)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage

        ## Import

        GameServerCluster can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{project}}/{{location}}/{{realm_id}}/{{cluster_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerCluster:GameServerCluster default {{location}}/{{realm_id}}/{{cluster_id}}
        ```

        :param str resource_name: The name of the resource.
        :param GameServerClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GameServerClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 connection_info: Optional[pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GameServerClusterArgs.__new__(GameServerClusterArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if connection_info is None and not opts.urn:
                raise TypeError("Missing required property 'connection_info'")
            __props__.__dict__["connection_info"] = connection_info
            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["name"] = None
        super(GameServerCluster, __self__).__init__(
            'gcp:gameservices/gameServerCluster:GameServerCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            connection_info: Optional[pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'GameServerCluster':
        """
        Get an existing GameServerCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Required. The resource name of the game server cluster
        :param pulumi.Input[pulumi.InputType['GameServerClusterConnectionInfoArgs']] connection_info: Game server cluster connection information. This information is used to
               manage game server clusters.
               Structure is documented below.
        :param pulumi.Input[str] description: Human readable description of the cluster.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels associated with this game server cluster. Each label is a
               key-value pair.
        :param pulumi.Input[str] location: Location of the Cluster.
        :param pulumi.Input[str] name: The resource id of the game server cluster, eg:
               'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
               'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] realm_id: The realm id of the game server realm.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GameServerClusterState.__new__(_GameServerClusterState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["connection_info"] = connection_info
        __props__.__dict__["description"] = description
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["realm_id"] = realm_id
        return GameServerCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Required. The resource name of the game server cluster
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="connectionInfo")
    def connection_info(self) -> pulumi.Output['outputs.GameServerClusterConnectionInfo']:
        """
        Game server cluster connection information. This information is used to
        manage game server clusters.
        Structure is documented below.
        """
        return pulumi.get(self, "connection_info")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human readable description of the cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The labels associated with this game server cluster. Each label is a
        key-value pair.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Location of the Cluster.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource id of the game server cluster, eg:
        'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'. For example,
        'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
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
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm id of the game server realm.
        """
        return pulumi.get(self, "realm_id")

