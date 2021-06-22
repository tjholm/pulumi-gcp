# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'TriggerDestinationArgs',
    'TriggerDestinationCloudRunServiceArgs',
    'TriggerMatchingCriteriaArgs',
    'TriggerTransportArgs',
    'TriggerTransportPubsubArgs',
]

@pulumi.input_type
class TriggerDestinationArgs:
    def __init__(__self__, *,
                 cloud_run_service: Optional[pulumi.Input['TriggerDestinationCloudRunServiceArgs']] = None):
        """
        :param pulumi.Input['TriggerDestinationCloudRunServiceArgs'] cloud_run_service: Cloud Run fully-managed service that receives the events. The service should be running in the same project as the trigger.
        """
        if cloud_run_service is not None:
            pulumi.set(__self__, "cloud_run_service", cloud_run_service)

    @property
    @pulumi.getter(name="cloudRunService")
    def cloud_run_service(self) -> Optional[pulumi.Input['TriggerDestinationCloudRunServiceArgs']]:
        """
        Cloud Run fully-managed service that receives the events. The service should be running in the same project as the trigger.
        """
        return pulumi.get(self, "cloud_run_service")

    @cloud_run_service.setter
    def cloud_run_service(self, value: Optional[pulumi.Input['TriggerDestinationCloudRunServiceArgs']]):
        pulumi.set(self, "cloud_run_service", value)


@pulumi.input_type
class TriggerDestinationCloudRunServiceArgs:
    def __init__(__self__, *,
                 service: pulumi.Input[str],
                 path: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] service: Required. The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
        :param pulumi.Input[str] path: Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
        :param pulumi.Input[str] region: Required. The region the Cloud Run service is deployed in.
        """
        pulumi.set(__self__, "service", service)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def service(self) -> pulumi.Input[str]:
        """
        Required. The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: pulumi.Input[str]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The region the Cloud Run service is deployed in.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class TriggerMatchingCriteriaArgs:
    def __init__(__self__, *,
                 attribute: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] attribute: Required. The name of a CloudEvents attribute. Currently, only a subset of attributes can be specified. All triggers MUST provide a matching criteria for the 'type' attribute.
        :param pulumi.Input[str] value: Required. The value for the attribute.
        """
        pulumi.set(__self__, "attribute", attribute)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def attribute(self) -> pulumi.Input[str]:
        """
        Required. The name of a CloudEvents attribute. Currently, only a subset of attributes can be specified. All triggers MUST provide a matching criteria for the 'type' attribute.
        """
        return pulumi.get(self, "attribute")

    @attribute.setter
    def attribute(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Required. The value for the attribute.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class TriggerTransportArgs:
    def __init__(__self__, *,
                 pubsubs: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportPubsubArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['TriggerTransportPubsubArgs']]] pubsubs: The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
        """
        if pubsubs is not None:
            pulumi.set(__self__, "pubsubs", pubsubs)

    @property
    @pulumi.getter
    def pubsubs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportPubsubArgs']]]]:
        """
        The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
        """
        return pulumi.get(self, "pubsubs")

    @pubsubs.setter
    def pubsubs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TriggerTransportPubsubArgs']]]]):
        pulumi.set(self, "pubsubs", value)


@pulumi.input_type
class TriggerTransportPubsubArgs:
    def __init__(__self__, *,
                 subscription: Optional[pulumi.Input[str]] = None,
                 topic: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] subscription: -
               Output only. The name of the Pub/Sub subscription created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}`.
        :param pulumi.Input[str] topic: Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/topics/{TOPIC_NAME}`. You may set an existing topic for triggers of the type `google.cloud.pubsub.topic.v1.messagePublished` only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
        """
        if subscription is not None:
            pulumi.set(__self__, "subscription", subscription)
        if topic is not None:
            pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter
    def subscription(self) -> Optional[pulumi.Input[str]]:
        """
        -
        Output only. The name of the Pub/Sub subscription created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}`.
        """
        return pulumi.get(self, "subscription")

    @subscription.setter
    def subscription(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subscription", value)

    @property
    @pulumi.getter
    def topic(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/topics/{TOPIC_NAME}`. You may set an existing topic for triggers of the type `google.cloud.pubsub.topic.v1.messagePublished` only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic", value)


