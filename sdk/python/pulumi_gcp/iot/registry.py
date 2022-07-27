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

__all__ = ['RegistryArgs', 'Registry']

@pulumi.input_type
class RegistryArgs:
    def __init__(__self__, *,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]]] = None,
                 http_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 mqtt_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        The set of arguments for constructing a Registry resource.
        :param pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]] credentials: List of public key certificates to authenticate devices.
               The structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]] event_notification_configs: List of configurations for event notifications, such as PubSub topics
               to publish device events to.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, Any]] http_config: Activate or deactivate HTTP.
               The structure is documented below.
        :param pulumi.Input[str] log_level: The default logging verbosity for activity from devices in this
               registry. Specifies which events should be written to logs. For
               example, if the LogLevel is ERROR, only events that terminate in
               errors will be logged. LogLevel is inclusive; enabling INFO logging
               will also enable ERROR logging.
               Default value is `NONE`.
               Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        :param pulumi.Input[Mapping[str, Any]] mqtt_config: Activate or deactivate MQTT.
               The structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[Mapping[str, Any]] state_notification_config: A PubSub topic to publish device state updates.
               The structure is documented below.
        """
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if event_notification_configs is not None:
            pulumi.set(__self__, "event_notification_configs", event_notification_configs)
        if http_config is not None:
            pulumi.set(__self__, "http_config", http_config)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if mqtt_config is not None:
            pulumi.set(__self__, "mqtt_config", mqtt_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if state_notification_config is not None:
            pulumi.set(__self__, "state_notification_config", state_notification_config)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]]:
        """
        List of public key certificates to authenticate devices.
        The structure is documented below.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="eventNotificationConfigs")
    def event_notification_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]]]:
        """
        List of configurations for event notifications, such as PubSub topics
        to publish device events to.
        Structure is documented below.
        """
        return pulumi.get(self, "event_notification_configs")

    @event_notification_configs.setter
    def event_notification_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]]]):
        pulumi.set(self, "event_notification_configs", value)

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Activate or deactivate HTTP.
        The structure is documented below.
        """
        return pulumi.get(self, "http_config")

    @http_config.setter
    def http_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "http_config", value)

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input[str]]:
        """
        The default logging verbosity for activity from devices in this
        registry. Specifies which events should be written to logs. For
        example, if the LogLevel is ERROR, only events that terminate in
        errors will be logged. LogLevel is inclusive; enabling INFO logging
        will also enable ERROR logging.
        Default value is `NONE`.
        Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter(name="mqttConfig")
    def mqtt_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Activate or deactivate MQTT.
        The structure is documented below.
        """
        return pulumi.get(self, "mqtt_config")

    @mqtt_config.setter
    def mqtt_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "mqtt_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the resource, required by device registry.
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
        The region in which the created registry should reside.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="stateNotificationConfig")
    def state_notification_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A PubSub topic to publish device state updates.
        The structure is documented below.
        """
        return pulumi.get(self, "state_notification_config")

    @state_notification_config.setter
    def state_notification_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "state_notification_config", value)


@pulumi.input_type
class _RegistryState:
    def __init__(__self__, *,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]]] = None,
                 http_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 mqtt_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input[Mapping[str, Any]]] = None):
        """
        Input properties used for looking up and filtering Registry resources.
        :param pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]] credentials: List of public key certificates to authenticate devices.
               The structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]] event_notification_configs: List of configurations for event notifications, such as PubSub topics
               to publish device events to.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, Any]] http_config: Activate or deactivate HTTP.
               The structure is documented below.
        :param pulumi.Input[str] log_level: The default logging verbosity for activity from devices in this
               registry. Specifies which events should be written to logs. For
               example, if the LogLevel is ERROR, only events that terminate in
               errors will be logged. LogLevel is inclusive; enabling INFO logging
               will also enable ERROR logging.
               Default value is `NONE`.
               Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        :param pulumi.Input[Mapping[str, Any]] mqtt_config: Activate or deactivate MQTT.
               The structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[Mapping[str, Any]] state_notification_config: A PubSub topic to publish device state updates.
               The structure is documented below.
        """
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if event_notification_configs is not None:
            pulumi.set(__self__, "event_notification_configs", event_notification_configs)
        if http_config is not None:
            pulumi.set(__self__, "http_config", http_config)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if mqtt_config is not None:
            pulumi.set(__self__, "mqtt_config", mqtt_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if state_notification_config is not None:
            pulumi.set(__self__, "state_notification_config", state_notification_config)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]]:
        """
        List of public key certificates to authenticate devices.
        The structure is documented below.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="eventNotificationConfigs")
    def event_notification_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]]]:
        """
        List of configurations for event notifications, such as PubSub topics
        to publish device events to.
        Structure is documented below.
        """
        return pulumi.get(self, "event_notification_configs")

    @event_notification_configs.setter
    def event_notification_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryEventNotificationConfigItemArgs']]]]):
        pulumi.set(self, "event_notification_configs", value)

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Activate or deactivate HTTP.
        The structure is documented below.
        """
        return pulumi.get(self, "http_config")

    @http_config.setter
    def http_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "http_config", value)

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input[str]]:
        """
        The default logging verbosity for activity from devices in this
        registry. Specifies which events should be written to logs. For
        example, if the LogLevel is ERROR, only events that terminate in
        errors will be logged. LogLevel is inclusive; enabling INFO logging
        will also enable ERROR logging.
        Default value is `NONE`.
        Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter(name="mqttConfig")
    def mqtt_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Activate or deactivate MQTT.
        The structure is documented below.
        """
        return pulumi.get(self, "mqtt_config")

    @mqtt_config.setter
    def mqtt_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "mqtt_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for the resource, required by device registry.
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
        The region in which the created registry should reside.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="stateNotificationConfig")
    def state_notification_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A PubSub topic to publish device state updates.
        The structure is documented below.
        """
        return pulumi.get(self, "state_notification_config")

    @state_notification_config.setter
    def state_notification_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "state_notification_config", value)


class Registry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]]] = None,
                 http_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 mqtt_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        """
        A Google Cloud IoT Core device registry.

        To get more information about DeviceRegistry, see:

        * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/iot/docs/)

        ## Example Usage
        ### Cloudiot Device Registry Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_registry = gcp.iot.Registry("test-registry")
        ```
        ### Cloudiot Device Registry Single Event Notification Configs

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_telemetry = gcp.pubsub.Topic("default-telemetry")
        test_registry = gcp.iot.Registry("test-registry", event_notification_configs=[gcp.iot.RegistryEventNotificationConfigItemArgs(
            pubsub_topic_name=default_telemetry.id,
            subfolder_matches="",
        )])
        ```
        ### Cloudiot Device Registry Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_devicestatus = gcp.pubsub.Topic("default-devicestatus")
        default_telemetry = gcp.pubsub.Topic("default-telemetry")
        additional_telemetry = gcp.pubsub.Topic("additional-telemetry")
        test_registry = gcp.iot.Registry("test-registry",
            event_notification_configs=[
                gcp.iot.RegistryEventNotificationConfigItemArgs(
                    pubsub_topic_name=additional_telemetry.id,
                    subfolder_matches="test/path",
                ),
                gcp.iot.RegistryEventNotificationConfigItemArgs(
                    pubsub_topic_name=default_telemetry.id,
                    subfolder_matches="",
                ),
            ],
            state_notification_config={
                "pubsub_topic_name": default_devicestatus.id,
            },
            mqtt_config={
                "mqtt_enabled_state": "MQTT_ENABLED",
            },
            http_config={
                "http_enabled_state": "HTTP_ENABLED",
            },
            log_level="INFO",
            credentials=[gcp.iot.RegistryCredentialArgs(
                public_key_certificate={
                    "format": "X509_CERTIFICATE_PEM",
                    "certificate": (lambda path: open(path).read())("test-fixtures/rsa_cert.pem"),
                },
            )])
        ```

        ## Import

        DeviceRegistry can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{project}}/locations/{{region}}/registries/{{name}}
        ```

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]] credentials: List of public key certificates to authenticate devices.
               The structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]] event_notification_configs: List of configurations for event notifications, such as PubSub topics
               to publish device events to.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, Any]] http_config: Activate or deactivate HTTP.
               The structure is documented below.
        :param pulumi.Input[str] log_level: The default logging verbosity for activity from devices in this
               registry. Specifies which events should be written to logs. For
               example, if the LogLevel is ERROR, only events that terminate in
               errors will be logged. LogLevel is inclusive; enabling INFO logging
               will also enable ERROR logging.
               Default value is `NONE`.
               Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        :param pulumi.Input[Mapping[str, Any]] mqtt_config: Activate or deactivate MQTT.
               The structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[Mapping[str, Any]] state_notification_config: A PubSub topic to publish device state updates.
               The structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RegistryArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Google Cloud IoT Core device registry.

        To get more information about DeviceRegistry, see:

        * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/iot/docs/)

        ## Example Usage
        ### Cloudiot Device Registry Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_registry = gcp.iot.Registry("test-registry")
        ```
        ### Cloudiot Device Registry Single Event Notification Configs

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_telemetry = gcp.pubsub.Topic("default-telemetry")
        test_registry = gcp.iot.Registry("test-registry", event_notification_configs=[gcp.iot.RegistryEventNotificationConfigItemArgs(
            pubsub_topic_name=default_telemetry.id,
            subfolder_matches="",
        )])
        ```
        ### Cloudiot Device Registry Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_devicestatus = gcp.pubsub.Topic("default-devicestatus")
        default_telemetry = gcp.pubsub.Topic("default-telemetry")
        additional_telemetry = gcp.pubsub.Topic("additional-telemetry")
        test_registry = gcp.iot.Registry("test-registry",
            event_notification_configs=[
                gcp.iot.RegistryEventNotificationConfigItemArgs(
                    pubsub_topic_name=additional_telemetry.id,
                    subfolder_matches="test/path",
                ),
                gcp.iot.RegistryEventNotificationConfigItemArgs(
                    pubsub_topic_name=default_telemetry.id,
                    subfolder_matches="",
                ),
            ],
            state_notification_config={
                "pubsub_topic_name": default_devicestatus.id,
            },
            mqtt_config={
                "mqtt_enabled_state": "MQTT_ENABLED",
            },
            http_config={
                "http_enabled_state": "HTTP_ENABLED",
            },
            log_level="INFO",
            credentials=[gcp.iot.RegistryCredentialArgs(
                public_key_certificate={
                    "format": "X509_CERTIFICATE_PEM",
                    "certificate": (lambda path: open(path).read())("test-fixtures/rsa_cert.pem"),
                },
            )])
        ```

        ## Import

        DeviceRegistry can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{project}}/locations/{{region}}/registries/{{name}}
        ```

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:iot/registry:Registry default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param RegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]]] = None,
                 http_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 log_level: Optional[pulumi.Input[str]] = None,
                 mqtt_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryArgs.__new__(RegistryArgs)

            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["event_notification_configs"] = event_notification_configs
            __props__.__dict__["http_config"] = http_config
            __props__.__dict__["log_level"] = log_level
            __props__.__dict__["mqtt_config"] = mqtt_config
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["state_notification_config"] = state_notification_config
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="gcp:kms/registry:Registry")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Registry, __self__).__init__(
            'gcp:iot/registry:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None,
            event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]]] = None,
            http_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            log_level: Optional[pulumi.Input[str]] = None,
            mqtt_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            state_notification_config: Optional[pulumi.Input[Mapping[str, Any]]] = None) -> 'Registry':
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]] credentials: List of public key certificates to authenticate devices.
               The structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryEventNotificationConfigItemArgs']]]] event_notification_configs: List of configurations for event notifications, such as PubSub topics
               to publish device events to.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, Any]] http_config: Activate or deactivate HTTP.
               The structure is documented below.
        :param pulumi.Input[str] log_level: The default logging verbosity for activity from devices in this
               registry. Specifies which events should be written to logs. For
               example, if the LogLevel is ERROR, only events that terminate in
               errors will be logged. LogLevel is inclusive; enabling INFO logging
               will also enable ERROR logging.
               Default value is `NONE`.
               Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        :param pulumi.Input[Mapping[str, Any]] mqtt_config: Activate or deactivate MQTT.
               The structure is documented below.
        :param pulumi.Input[str] name: A unique name for the resource, required by device registry.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
        :param pulumi.Input[Mapping[str, Any]] state_notification_config: A PubSub topic to publish device state updates.
               The structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistryState.__new__(_RegistryState)

        __props__.__dict__["credentials"] = credentials
        __props__.__dict__["event_notification_configs"] = event_notification_configs
        __props__.__dict__["http_config"] = http_config
        __props__.__dict__["log_level"] = log_level
        __props__.__dict__["mqtt_config"] = mqtt_config
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["state_notification_config"] = state_notification_config
        return Registry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional[Sequence['outputs.RegistryCredential']]]:
        """
        List of public key certificates to authenticate devices.
        The structure is documented below.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="eventNotificationConfigs")
    def event_notification_configs(self) -> pulumi.Output[Sequence['outputs.RegistryEventNotificationConfigItem']]:
        """
        List of configurations for event notifications, such as PubSub topics
        to publish device events to.
        Structure is documented below.
        """
        return pulumi.get(self, "event_notification_configs")

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Activate or deactivate HTTP.
        The structure is documented below.
        """
        return pulumi.get(self, "http_config")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> pulumi.Output[Optional[str]]:
        """
        The default logging verbosity for activity from devices in this
        registry. Specifies which events should be written to logs. For
        example, if the LogLevel is ERROR, only events that terminate in
        errors will be logged. LogLevel is inclusive; enabling INFO logging
        will also enable ERROR logging.
        Default value is `NONE`.
        Possible values are `NONE`, `ERROR`, `INFO`, and `DEBUG`.
        """
        return pulumi.get(self, "log_level")

    @property
    @pulumi.getter(name="mqttConfig")
    def mqtt_config(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Activate or deactivate MQTT.
        The structure is documented below.
        """
        return pulumi.get(self, "mqtt_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for the resource, required by device registry.
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
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which the created registry should reside.
        If it is not provided, the provider region is used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="stateNotificationConfig")
    def state_notification_config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A PubSub topic to publish device state updates.
        The structure is documented below.
        """
        return pulumi.get(self, "state_notification_config")

