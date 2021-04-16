# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['GameServerConfigArgs', 'GameServerConfig']

@pulumi.input_type
class GameServerConfigArgs:
    def __init__(__self__, *,
                 config_id: pulumi.Input[str],
                 deployment_id: pulumi.Input[str],
                 fleet_configs: pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scaling_configs: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]]] = None):
        """
        The set of arguments for constructing a GameServerConfig resource.
        :param pulumi.Input[str] config_id: A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]] fleet_configs: The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.
               Structure is documented below.
        :param pulumi.Input[str] description: The description of the game server config.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels to group by.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]] scaling_configs: Optional. This contains the autoscaling settings.
               Structure is documented below.
        """
        pulumi.set(__self__, "config_id", config_id)
        pulumi.set(__self__, "deployment_id", deployment_id)
        pulumi.set(__self__, "fleet_configs", fleet_configs)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if scaling_configs is not None:
            pulumi.set(__self__, "scaling_configs", scaling_configs)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Input[str]:
        """
        A unique id for the deployment config.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Input[str]:
        """
        A unique id for the deployment.
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter(name="fleetConfigs")
    def fleet_configs(self) -> pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]]:
        """
        The fleet config contains list of fleet specs. In the Single Cloud, there
        will be only one.
        Structure is documented below.
        """
        return pulumi.get(self, "fleet_configs")

    @fleet_configs.setter
    def fleet_configs(self, value: pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]]):
        pulumi.set(self, "fleet_configs", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the game server config.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of labels to group by.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the Deployment.
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

    @property
    @pulumi.getter(name="scalingConfigs")
    def scaling_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]]]:
        """
        Optional. This contains the autoscaling settings.
        Structure is documented below.
        """
        return pulumi.get(self, "scaling_configs")

    @scaling_configs.setter
    def scaling_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]]]):
        pulumi.set(self, "scaling_configs", value)


@pulumi.input_type
class _GameServerConfigState:
    def __init__(__self__, *,
                 config_id: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fleet_configs: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scaling_configs: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]]] = None):
        """
        Input properties used for looking up and filtering GameServerConfig resources.
        :param pulumi.Input[str] config_id: A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: The description of the game server config.
        :param pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]] fleet_configs: The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels to group by.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] name: The name of the ScalingConfig
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]] scaling_configs: Optional. This contains the autoscaling settings.
               Structure is documented below.
        """
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if deployment_id is not None:
            pulumi.set(__self__, "deployment_id", deployment_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fleet_configs is not None:
            pulumi.set(__self__, "fleet_configs", fleet_configs)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if scaling_configs is not None:
            pulumi.set(__self__, "scaling_configs", scaling_configs)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique id for the deployment config.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> Optional[pulumi.Input[str]]:
        """
        A unique id for the deployment.
        """
        return pulumi.get(self, "deployment_id")

    @deployment_id.setter
    def deployment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deployment_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the game server config.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="fleetConfigs")
    def fleet_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]]]:
        """
        The fleet config contains list of fleet specs. In the Single Cloud, there
        will be only one.
        Structure is documented below.
        """
        return pulumi.get(self, "fleet_configs")

    @fleet_configs.setter
    def fleet_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigFleetConfigArgs']]]]):
        pulumi.set(self, "fleet_configs", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of labels to group by.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the Deployment.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ScalingConfig
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
    @pulumi.getter(name="scalingConfigs")
    def scaling_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]]]:
        """
        Optional. This contains the autoscaling settings.
        Structure is documented below.
        """
        return pulumi.get(self, "scaling_configs")

    @scaling_configs.setter
    def scaling_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GameServerConfigScalingConfigArgs']]]]):
        pulumi.set(self, "scaling_configs", value)


class GameServerConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fleet_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigFleetConfigArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scaling_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigScalingConfigArgs']]]]] = None,
                 __props__=None):
        """
        A game server config resource. Configs are global and immutable.

        To get more information about GameServerConfig, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage
        ### Game Service Config Basic

        ```python
        import pulumi
        import json
        import pulumi_gcp as gcp

        default_game_server_deployment = gcp.gameservices.GameServerDeployment("defaultGameServerDeployment",
            deployment_id="tf-test-deployment",
            description="a deployment description")
        default_game_server_config = gcp.gameservices.GameServerConfig("defaultGameServerConfig",
            config_id="tf-test-config",
            deployment_id=default_game_server_deployment.deployment_id,
            description="a config description",
            fleet_configs=[gcp.gameservices.GameServerConfigFleetConfigArgs(
                name="something-unique",
                fleet_spec=json.dumps({
                    "replicas": 1,
                    "scheduling": "Packed",
                    "template": {
                        "metadata": {
                            "name": "tf-test-game-server-template",
                        },
                        "spec": {
                            "ports": [{
                                "name": "default",
                                "portPolicy": "Dynamic",
                                "containerPort": 7654,
                                "protocol": "UDP",
                            }],
                            "template": {
                                "spec": {
                                    "containers": [{
                                        "name": "simple-udp-server",
                                        "image": "gcr.io/agones-images/udp-server:0.14",
                                    }],
                                },
                            },
                        },
                    },
                }),
            )],
            scaling_configs=[gcp.gameservices.GameServerConfigScalingConfigArgs(
                name="scaling-config-name",
                fleet_autoscaler_spec=json.dumps({
                    "policy": {
                        "type": "Webhook",
                        "webhook": {
                            "service": {
                                "name": "autoscaler-webhook-service",
                                "namespace": "default",
                                "path": "scale",
                            },
                        },
                    },
                }),
                selectors=[gcp.gameservices.GameServerConfigScalingConfigSelectorArgs(
                    labels={
                        "one": "two",
                    },
                )],
                schedules=[gcp.gameservices.GameServerConfigScalingConfigScheduleArgs(
                    cron_job_duration="3.500s",
                    cron_spec="0 0 * * 0",
                )],
            )])
        ```

        ## Import

        GameServerConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default {{project}}/{{location}}/{{deployment_id}}/{{config_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default {{location}}/{{deployment_id}}/{{config_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: The description of the game server config.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigFleetConfigArgs']]]] fleet_configs: The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels to group by.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigScalingConfigArgs']]]] scaling_configs: Optional. This contains the autoscaling settings.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GameServerConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A game server config resource. Configs are global and immutable.

        To get more information about GameServerConfig, see:

        * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/game-servers/docs)

        ## Example Usage
        ### Game Service Config Basic

        ```python
        import pulumi
        import json
        import pulumi_gcp as gcp

        default_game_server_deployment = gcp.gameservices.GameServerDeployment("defaultGameServerDeployment",
            deployment_id="tf-test-deployment",
            description="a deployment description")
        default_game_server_config = gcp.gameservices.GameServerConfig("defaultGameServerConfig",
            config_id="tf-test-config",
            deployment_id=default_game_server_deployment.deployment_id,
            description="a config description",
            fleet_configs=[gcp.gameservices.GameServerConfigFleetConfigArgs(
                name="something-unique",
                fleet_spec=json.dumps({
                    "replicas": 1,
                    "scheduling": "Packed",
                    "template": {
                        "metadata": {
                            "name": "tf-test-game-server-template",
                        },
                        "spec": {
                            "ports": [{
                                "name": "default",
                                "portPolicy": "Dynamic",
                                "containerPort": 7654,
                                "protocol": "UDP",
                            }],
                            "template": {
                                "spec": {
                                    "containers": [{
                                        "name": "simple-udp-server",
                                        "image": "gcr.io/agones-images/udp-server:0.14",
                                    }],
                                },
                            },
                        },
                    },
                }),
            )],
            scaling_configs=[gcp.gameservices.GameServerConfigScalingConfigArgs(
                name="scaling-config-name",
                fleet_autoscaler_spec=json.dumps({
                    "policy": {
                        "type": "Webhook",
                        "webhook": {
                            "service": {
                                "name": "autoscaler-webhook-service",
                                "namespace": "default",
                                "path": "scale",
                            },
                        },
                    },
                }),
                selectors=[gcp.gameservices.GameServerConfigScalingConfigSelectorArgs(
                    labels={
                        "one": "two",
                    },
                )],
                schedules=[gcp.gameservices.GameServerConfigScalingConfigScheduleArgs(
                    cron_job_duration="3.500s",
                    cron_spec="0 0 * * 0",
                )],
            )])
        ```

        ## Import

        GameServerConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default {{project}}/{{location}}/{{deployment_id}}/{{config_id}}
        ```

        ```sh
         $ pulumi import gcp:gameservices/gameServerConfig:GameServerConfig default {{location}}/{{deployment_id}}/{{config_id}}
        ```

        :param str resource_name: The name of the resource.
        :param GameServerConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GameServerConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 deployment_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fleet_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigFleetConfigArgs']]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scaling_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigScalingConfigArgs']]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GameServerConfigArgs.__new__(GameServerConfigArgs)

            if config_id is None and not opts.urn:
                raise TypeError("Missing required property 'config_id'")
            __props__.__dict__["config_id"] = config_id
            if deployment_id is None and not opts.urn:
                raise TypeError("Missing required property 'deployment_id'")
            __props__.__dict__["deployment_id"] = deployment_id
            __props__.__dict__["description"] = description
            if fleet_configs is None and not opts.urn:
                raise TypeError("Missing required property 'fleet_configs'")
            __props__.__dict__["fleet_configs"] = fleet_configs
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["scaling_configs"] = scaling_configs
            __props__.__dict__["name"] = None
        super(GameServerConfig, __self__).__init__(
            'gcp:gameservices/gameServerConfig:GameServerConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            deployment_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            fleet_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigFleetConfigArgs']]]]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            scaling_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigScalingConfigArgs']]]]] = None) -> 'GameServerConfig':
        """
        Get an existing GameServerConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] config_id: A unique id for the deployment config.
        :param pulumi.Input[str] deployment_id: A unique id for the deployment.
        :param pulumi.Input[str] description: The description of the game server config.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigFleetConfigArgs']]]] fleet_configs: The fleet config contains list of fleet specs. In the Single Cloud, there
               will be only one.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels to group by.
        :param pulumi.Input[str] location: Location of the Deployment.
        :param pulumi.Input[str] name: The name of the ScalingConfig
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GameServerConfigScalingConfigArgs']]]] scaling_configs: Optional. This contains the autoscaling settings.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GameServerConfigState.__new__(_GameServerConfigState)

        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["deployment_id"] = deployment_id
        __props__.__dict__["description"] = description
        __props__.__dict__["fleet_configs"] = fleet_configs
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["scaling_configs"] = scaling_configs
        return GameServerConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[str]:
        """
        A unique id for the deployment config.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="deploymentId")
    def deployment_id(self) -> pulumi.Output[str]:
        """
        A unique id for the deployment.
        """
        return pulumi.get(self, "deployment_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the game server config.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fleetConfigs")
    def fleet_configs(self) -> pulumi.Output[Sequence['outputs.GameServerConfigFleetConfig']]:
        """
        The fleet config contains list of fleet specs. In the Single Cloud, there
        will be only one.
        Structure is documented below.
        """
        return pulumi.get(self, "fleet_configs")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Set of labels to group by.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Location of the Deployment.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the ScalingConfig
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
    @pulumi.getter(name="scalingConfigs")
    def scaling_configs(self) -> pulumi.Output[Optional[Sequence['outputs.GameServerConfigScalingConfig']]]:
        """
        Optional. This contains the autoscaling settings.
        Structure is documented below.
        """
        return pulumi.get(self, "scaling_configs")

