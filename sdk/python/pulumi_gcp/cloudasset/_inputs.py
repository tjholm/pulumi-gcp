# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'FolderFeedConditionArgs',
    'FolderFeedFeedOutputConfigArgs',
    'FolderFeedFeedOutputConfigPubsubDestinationArgs',
    'OrganizationFeedConditionArgs',
    'OrganizationFeedFeedOutputConfigArgs',
    'OrganizationFeedFeedOutputConfigPubsubDestinationArgs',
    'ProjectFeedConditionArgs',
    'ProjectFeedFeedOutputConfigArgs',
    'ProjectFeedFeedOutputConfigPubsubDestinationArgs',
]

@pulumi.input_type
class FolderFeedConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] description: Description of the expression. This is a longer text which describes the expression,
               e.g. when hovered over it in a UI.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file
               name and a position in the file.
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.
               This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "expression", expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the expression. This is a longer text which describes the expression,
        e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        String indicating the location of the expression for error reporting, e.g. a file
        name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Title for the expression, i.e. a short string describing its purpose.
        This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class FolderFeedFeedOutputConfigArgs:
    def __init__(__self__, *,
                 pubsub_destination: pulumi.Input['FolderFeedFeedOutputConfigPubsubDestinationArgs']):
        """
        :param pulumi.Input['FolderFeedFeedOutputConfigPubsubDestinationArgs'] pubsub_destination: Destination on Cloud Pubsub.
               Structure is documented below.
        """
        pulumi.set(__self__, "pubsub_destination", pubsub_destination)

    @property
    @pulumi.getter(name="pubsubDestination")
    def pubsub_destination(self) -> pulumi.Input['FolderFeedFeedOutputConfigPubsubDestinationArgs']:
        """
        Destination on Cloud Pubsub.
        Structure is documented below.
        """
        return pulumi.get(self, "pubsub_destination")

    @pubsub_destination.setter
    def pubsub_destination(self, value: pulumi.Input['FolderFeedFeedOutputConfigPubsubDestinationArgs']):
        pulumi.set(self, "pubsub_destination", value)


@pulumi.input_type
class FolderFeedFeedOutputConfigPubsubDestinationArgs:
    def __init__(__self__, *,
                 topic: pulumi.Input[str]):
        """
        :param pulumi.Input[str] topic: Destination on Cloud Pubsub topic.
        """
        pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        Destination on Cloud Pubsub topic.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)


@pulumi.input_type
class OrganizationFeedConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] description: Description of the expression. This is a longer text which describes the expression,
               e.g. when hovered over it in a UI.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file
               name and a position in the file.
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.
               This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "expression", expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the expression. This is a longer text which describes the expression,
        e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        String indicating the location of the expression for error reporting, e.g. a file
        name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Title for the expression, i.e. a short string describing its purpose.
        This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class OrganizationFeedFeedOutputConfigArgs:
    def __init__(__self__, *,
                 pubsub_destination: pulumi.Input['OrganizationFeedFeedOutputConfigPubsubDestinationArgs']):
        """
        :param pulumi.Input['OrganizationFeedFeedOutputConfigPubsubDestinationArgs'] pubsub_destination: Destination on Cloud Pubsub.
               Structure is documented below.
        """
        pulumi.set(__self__, "pubsub_destination", pubsub_destination)

    @property
    @pulumi.getter(name="pubsubDestination")
    def pubsub_destination(self) -> pulumi.Input['OrganizationFeedFeedOutputConfigPubsubDestinationArgs']:
        """
        Destination on Cloud Pubsub.
        Structure is documented below.
        """
        return pulumi.get(self, "pubsub_destination")

    @pubsub_destination.setter
    def pubsub_destination(self, value: pulumi.Input['OrganizationFeedFeedOutputConfigPubsubDestinationArgs']):
        pulumi.set(self, "pubsub_destination", value)


@pulumi.input_type
class OrganizationFeedFeedOutputConfigPubsubDestinationArgs:
    def __init__(__self__, *,
                 topic: pulumi.Input[str]):
        """
        :param pulumi.Input[str] topic: Destination on Cloud Pubsub topic.
        """
        pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        Destination on Cloud Pubsub topic.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)


@pulumi.input_type
class ProjectFeedConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] description: Description of the expression. This is a longer text which describes the expression,
               e.g. when hovered over it in a UI.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file
               name and a position in the file.
        :param pulumi.Input[str] title: Title for the expression, i.e. a short string describing its purpose.
               This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "expression", expression)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the expression. This is a longer text which describes the expression,
        e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        String indicating the location of the expression for error reporting, e.g. a file
        name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Title for the expression, i.e. a short string describing its purpose.
        This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class ProjectFeedFeedOutputConfigArgs:
    def __init__(__self__, *,
                 pubsub_destination: pulumi.Input['ProjectFeedFeedOutputConfigPubsubDestinationArgs']):
        """
        :param pulumi.Input['ProjectFeedFeedOutputConfigPubsubDestinationArgs'] pubsub_destination: Destination on Cloud Pubsub.
               Structure is documented below.
        """
        pulumi.set(__self__, "pubsub_destination", pubsub_destination)

    @property
    @pulumi.getter(name="pubsubDestination")
    def pubsub_destination(self) -> pulumi.Input['ProjectFeedFeedOutputConfigPubsubDestinationArgs']:
        """
        Destination on Cloud Pubsub.
        Structure is documented below.
        """
        return pulumi.get(self, "pubsub_destination")

    @pubsub_destination.setter
    def pubsub_destination(self, value: pulumi.Input['ProjectFeedFeedOutputConfigPubsubDestinationArgs']):
        pulumi.set(self, "pubsub_destination", value)


@pulumi.input_type
class ProjectFeedFeedOutputConfigPubsubDestinationArgs:
    def __init__(__self__, *,
                 topic: pulumi.Input[str]):
        """
        :param pulumi.Input[str] topic: Destination on Cloud Pubsub topic.
        """
        pulumi.set(__self__, "topic", topic)

    @property
    @pulumi.getter
    def topic(self) -> pulumi.Input[str]:
        """
        Destination on Cloud Pubsub topic.
        """
        return pulumi.get(self, "topic")

    @topic.setter
    def topic(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic", value)


