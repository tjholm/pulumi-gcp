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

__all__ = ['LiteSubscriptionArgs', 'LiteSubscription']

@pulumi.input_type
class LiteSubscriptionArgs:
    def __init__(__self__, *,
                 topic: pulumi.Input[str],
                 delivery_config: Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LiteSubscription resource.
        :param pulumi.Input[str] topic: A reference to a Topic resource.
        :param pulumi.Input['LiteSubscriptionDeliveryConfigArgs'] delivery_config: The settings for this subscription's message delivery.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the subscription.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the pubsub lite topic.
        :param pulumi.Input[str] zone: The zone of the pubsub lite topic.
        """
        LiteSubscriptionArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            topic=topic,
            delivery_config=delivery_config,
            name=name,
            project=project,
            region=region,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             topic: Optional[pulumi.Input[str]] = None,
             delivery_config: Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             zone: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if topic is None:
            raise TypeError("Missing 'topic' argument")
        if delivery_config is None and 'deliveryConfig' in kwargs:
            delivery_config = kwargs['deliveryConfig']

        _setter("topic", topic)
        if delivery_config is not None:
            _setter("delivery_config", delivery_config)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        A reference to a Topic resource.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter(name="deliveryConfig")
    def delivery_config(self) -> Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']]:
        """
        The settings for this subscription's message delivery.
        Structure is documented below.
        """
        return pulumi.get(self, "delivery_config")

    @delivery_config.setter
    def delivery_config(self, value: Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']]):
        pulumi.set(self, "delivery_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the subscription.


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
        The region of the pubsub lite topic.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the pubsub lite topic.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


@pulumi.input_type
class _LiteSubscriptionState:
    def __init__(__self__, *,
                 delivery_config: Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LiteSubscription resources.
        :param pulumi.Input['LiteSubscriptionDeliveryConfigArgs'] delivery_config: The settings for this subscription's message delivery.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the subscription.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the pubsub lite topic.
        :param pulumi.Input[str] topic: A reference to a Topic resource.
        :param pulumi.Input[str] zone: The zone of the pubsub lite topic.
        """
        _LiteSubscriptionState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            delivery_config=delivery_config,
            name=name,
            project=project,
            region=region,
            topic=topic,
            zone=zone,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             delivery_config: Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             topic: Optional[pulumi.Input[str]] = None,
             zone: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if delivery_config is None and 'deliveryConfig' in kwargs:
            delivery_config = kwargs['deliveryConfig']

        if delivery_config is not None:
            _setter("delivery_config", delivery_config)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if topic is not None:
            _setter("topic", topic)
        if zone is not None:
            _setter("zone", zone)

    @property
    @pulumi.getter(name="deliveryConfig")
    def delivery_config(self) -> Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']]:
        """
        The settings for this subscription's message delivery.
        Structure is documented below.
        """
        return pulumi.get(self, "delivery_config")

    @delivery_config.setter
    def delivery_config(self, value: Optional[pulumi.Input['LiteSubscriptionDeliveryConfigArgs']]):
        pulumi.set(self, "delivery_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the subscription.


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
        The region of the pubsub lite topic.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        """
        A reference to a Topic resource.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone of the pubsub lite topic.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class LiteSubscription(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delivery_config: Optional[pulumi.Input[pulumi.InputType['LiteSubscriptionDeliveryConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A named resource representing the stream of messages from a single,
        specific topic, to be delivered to the subscribing application.

        To get more information about Subscription, see:

        * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.subscriptions)
        * How-to Guides
            * [Managing Subscriptions](https://cloud.google.com/pubsub/lite/docs/subscriptions)

        ## Example Usage
        ### Pubsub Lite Subscription Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.get_project()
        example_lite_topic = gcp.pubsub.LiteTopic("exampleLiteTopic",
            project=project.number,
            partition_config=gcp.pubsub.LiteTopicPartitionConfigArgs(
                count=1,
                capacity=gcp.pubsub.LiteTopicPartitionConfigCapacityArgs(
                    publish_mib_per_sec=4,
                    subscribe_mib_per_sec=8,
                ),
            ),
            retention_config=gcp.pubsub.LiteTopicRetentionConfigArgs(
                per_partition_bytes="32212254720",
            ))
        example_lite_subscription = gcp.pubsub.LiteSubscription("exampleLiteSubscription",
            topic=example_lite_topic.name,
            delivery_config=gcp.pubsub.LiteSubscriptionDeliveryConfigArgs(
                delivery_requirement="DELIVER_AFTER_STORED",
            ))
        ```

        ## Import

        Subscription can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LiteSubscriptionDeliveryConfigArgs']] delivery_config: The settings for this subscription's message delivery.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the subscription.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the pubsub lite topic.
        :param pulumi.Input[str] topic: A reference to a Topic resource.
        :param pulumi.Input[str] zone: The zone of the pubsub lite topic.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LiteSubscriptionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A named resource representing the stream of messages from a single,
        specific topic, to be delivered to the subscribing application.

        To get more information about Subscription, see:

        * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.subscriptions)
        * How-to Guides
            * [Managing Subscriptions](https://cloud.google.com/pubsub/lite/docs/subscriptions)

        ## Example Usage
        ### Pubsub Lite Subscription Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.get_project()
        example_lite_topic = gcp.pubsub.LiteTopic("exampleLiteTopic",
            project=project.number,
            partition_config=gcp.pubsub.LiteTopicPartitionConfigArgs(
                count=1,
                capacity=gcp.pubsub.LiteTopicPartitionConfigCapacityArgs(
                    publish_mib_per_sec=4,
                    subscribe_mib_per_sec=8,
                ),
            ),
            retention_config=gcp.pubsub.LiteTopicRetentionConfigArgs(
                per_partition_bytes="32212254720",
            ))
        example_lite_subscription = gcp.pubsub.LiteSubscription("exampleLiteSubscription",
            topic=example_lite_topic.name,
            delivery_config=gcp.pubsub.LiteSubscriptionDeliveryConfigArgs(
                delivery_requirement="DELIVER_AFTER_STORED",
            ))
        ```

        ## Import

        Subscription can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{project}}/{{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{zone}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param LiteSubscriptionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LiteSubscriptionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            LiteSubscriptionArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delivery_config: Optional[pulumi.Input[pulumi.InputType['LiteSubscriptionDeliveryConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LiteSubscriptionArgs.__new__(LiteSubscriptionArgs)

            delivery_config = _utilities.configure(delivery_config, LiteSubscriptionDeliveryConfigArgs, True)
            __props__.__dict__["delivery_config"] = delivery_config
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            if topic is None and not opts.urn:
                raise TypeError("Missing required property 'topic'")
            __props__.__dict__["topic"] = topic
            __props__.__dict__["zone"] = zone
        super(LiteSubscription, __self__).__init__(
            'gcp:pubsub/liteSubscription:LiteSubscription',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delivery_config: Optional[pulumi.Input[pulumi.InputType['LiteSubscriptionDeliveryConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            topic: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'LiteSubscription':
        """
        Get an existing LiteSubscription resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['LiteSubscriptionDeliveryConfigArgs']] delivery_config: The settings for this subscription's message delivery.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the subscription.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the pubsub lite topic.
        :param pulumi.Input[str] topic: A reference to a Topic resource.
        :param pulumi.Input[str] zone: The zone of the pubsub lite topic.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LiteSubscriptionState.__new__(_LiteSubscriptionState)

        __props__.__dict__["delivery_config"] = delivery_config
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["topic"] = topic
        __props__.__dict__["zone"] = zone
        return LiteSubscription(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deliveryConfig")
    def delivery_config(self) -> pulumi.Output[Optional['outputs.LiteSubscriptionDeliveryConfig']]:
        """
        The settings for this subscription's message delivery.
        Structure is documented below.
        """
        return pulumi.get(self, "delivery_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the subscription.


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
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region of the pubsub lite topic.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Output[str]:
        """
        A reference to a Topic resource.
        """
        return pulumi.get(self, "topic")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[Optional[str]]:
        """
        The zone of the pubsub lite topic.
        """
        return pulumi.get(self, "zone")

