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

__all__ = ['FolderFeedArgs', 'FolderFeed']

@pulumi.input_type
class FolderFeedArgs:
    def __init__(__self__, *,
                 billing_project: pulumi.Input[str],
                 feed_id: pulumi.Input[str],
                 feed_output_config: pulumi.Input['FolderFeedFeedOutputConfigArgs'],
                 folder: pulumi.Input[str],
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 condition: Optional[pulumi.Input['FolderFeedConditionArgs']] = None,
                 content_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FolderFeed resource.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input['FolderFeedFeedOutputConfigArgs'] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] folder: The folder this feed should be created in.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input['FolderFeedConditionArgs'] condition: A condition which determines whether an asset update should be published. If specified, an asset
               will be returned only when the expression evaluates to true. When set, expression field
               must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
               expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
               condition are optional.
               Structure is documented below.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        """
        pulumi.set(__self__, "billing_project", billing_project)
        pulumi.set(__self__, "feed_id", feed_id)
        pulumi.set(__self__, "feed_output_config", feed_output_config)
        pulumi.set(__self__, "folder", folder)
        if asset_names is not None:
            pulumi.set(__self__, "asset_names", asset_names)
        if asset_types is not None:
            pulumi.set(__self__, "asset_types", asset_types)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)

    @property
    @pulumi.getter(name="billingProject")
    def billing_project(self) -> pulumi.Input[str]:
        """
        The project whose identity will be used when sending messages to the
        destination pubsub topic. It also specifies the project for API
        enablement check, quota, and billing.
        """
        return pulumi.get(self, "billing_project")

    @billing_project.setter
    def billing_project(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_project", value)

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Input[str]:
        """
        This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "feed_id", value)

    @property
    @pulumi.getter(name="feedOutputConfig")
    def feed_output_config(self) -> pulumi.Input['FolderFeedFeedOutputConfigArgs']:
        """
        Output configuration for asset feed destination.
        Structure is documented below.
        """
        return pulumi.get(self, "feed_output_config")

    @feed_output_config.setter
    def feed_output_config(self, value: pulumi.Input['FolderFeedFeedOutputConfigArgs']):
        pulumi.set(self, "feed_output_config", value)

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Input[str]:
        """
        The folder this feed should be created in.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="assetNames")
    def asset_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the full names of the assets to receive updates. You must specify either or both of
        assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
        exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
        See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        """
        return pulumi.get(self, "asset_names")

    @asset_names.setter
    def asset_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "asset_names", value)

    @property
    @pulumi.getter(name="assetTypes")
    def asset_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of types of the assets to receive updates. You must specify either or both of assetNames
        and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
        the feed. For example: "compute.googleapis.com/Disk"
        See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
        supported asset types.
        """
        return pulumi.get(self, "asset_types")

    @asset_types.setter
    def asset_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "asset_types", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['FolderFeedConditionArgs']]:
        """
        A condition which determines whether an asset update should be published. If specified, an asset
        will be returned only when the expression evaluates to true. When set, expression field
        must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
        expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
        condition are optional.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['FolderFeedConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        Asset content type. If not specified, no content but the asset name and type will be returned.
        Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)


@pulumi.input_type
class _FolderFeedState:
    def __init__(__self__, *,
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 billing_project: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['FolderFeedConditionArgs']] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 feed_id: Optional[pulumi.Input[str]] = None,
                 feed_output_config: Optional[pulumi.Input['FolderFeedFeedOutputConfigArgs']] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FolderFeed resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing.
        :param pulumi.Input['FolderFeedConditionArgs'] condition: A condition which determines whether an asset update should be published. If specified, an asset
               will be returned only when the expression evaluates to true. When set, expression field
               must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
               expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
               condition are optional.
               Structure is documented below.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input['FolderFeedFeedOutputConfigArgs'] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] folder: The folder this feed should be created in.
        :param pulumi.Input[str] folder_id: The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
        :param pulumi.Input[str] name: The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
        """
        if asset_names is not None:
            pulumi.set(__self__, "asset_names", asset_names)
        if asset_types is not None:
            pulumi.set(__self__, "asset_types", asset_types)
        if billing_project is not None:
            pulumi.set(__self__, "billing_project", billing_project)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if feed_id is not None:
            pulumi.set(__self__, "feed_id", feed_id)
        if feed_output_config is not None:
            pulumi.set(__self__, "feed_output_config", feed_output_config)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if folder_id is not None:
            pulumi.set(__self__, "folder_id", folder_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="assetNames")
    def asset_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of the full names of the assets to receive updates. You must specify either or both of
        assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
        exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
        See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        """
        return pulumi.get(self, "asset_names")

    @asset_names.setter
    def asset_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "asset_names", value)

    @property
    @pulumi.getter(name="assetTypes")
    def asset_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of types of the assets to receive updates. You must specify either or both of assetNames
        and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
        the feed. For example: "compute.googleapis.com/Disk"
        See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
        supported asset types.
        """
        return pulumi.get(self, "asset_types")

    @asset_types.setter
    def asset_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "asset_types", value)

    @property
    @pulumi.getter(name="billingProject")
    def billing_project(self) -> Optional[pulumi.Input[str]]:
        """
        The project whose identity will be used when sending messages to the
        destination pubsub topic. It also specifies the project for API
        enablement check, quota, and billing.
        """
        return pulumi.get(self, "billing_project")

    @billing_project.setter
    def billing_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_project", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['FolderFeedConditionArgs']]:
        """
        A condition which determines whether an asset update should be published. If specified, an asset
        will be returned only when the expression evaluates to true. When set, expression field
        must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
        expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
        condition are optional.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['FolderFeedConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        Asset content type. If not specified, no content but the asset name and type will be returned.
        Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> Optional[pulumi.Input[str]]:
        """
        This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        """
        return pulumi.get(self, "feed_id")

    @feed_id.setter
    def feed_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feed_id", value)

    @property
    @pulumi.getter(name="feedOutputConfig")
    def feed_output_config(self) -> Optional[pulumi.Input['FolderFeedFeedOutputConfigArgs']]:
        """
        Output configuration for asset feed destination.
        Structure is documented below.
        """
        return pulumi.get(self, "feed_output_config")

    @feed_output_config.setter
    def feed_output_config(self, value: Optional[pulumi.Input['FolderFeedFeedOutputConfigArgs']]):
        pulumi.set(self, "feed_output_config", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The folder this feed should be created in.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
        """
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class FolderFeed(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 billing_project: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['FolderFeedConditionArgs']]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 feed_id: Optional[pulumi.Input[str]] = None,
                 feed_output_config: Optional[pulumi.Input[pulumi.InputType['FolderFeedFeedOutputConfigArgs']]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Describes a Cloud Asset Inventory feed used to to listen to asset updates.

        To get more information about FolderFeed, see:

        * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/asset-inventory/docs)

        ## Example Usage

        ## Import

        FolderFeed can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudasset/folderFeed:FolderFeed default folders/{{folder_id}}/feeds/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudasset/folderFeed:FolderFeed default {{folder_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing.
        :param pulumi.Input[pulumi.InputType['FolderFeedConditionArgs']] condition: A condition which determines whether an asset update should be published. If specified, an asset
               will be returned only when the expression evaluates to true. When set, expression field
               must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
               expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
               condition are optional.
               Structure is documented below.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input[pulumi.InputType['FolderFeedFeedOutputConfigArgs']] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] folder: The folder this feed should be created in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderFeedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Describes a Cloud Asset Inventory feed used to to listen to asset updates.

        To get more information about FolderFeed, see:

        * [API documentation](https://cloud.google.com/asset-inventory/docs/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/asset-inventory/docs)

        ## Example Usage

        ## Import

        FolderFeed can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudasset/folderFeed:FolderFeed default folders/{{folder_id}}/feeds/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudasset/folderFeed:FolderFeed default {{folder_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param FolderFeedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderFeedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 billing_project: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['FolderFeedConditionArgs']]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 feed_id: Optional[pulumi.Input[str]] = None,
                 feed_output_config: Optional[pulumi.Input[pulumi.InputType['FolderFeedFeedOutputConfigArgs']]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FolderFeedArgs.__new__(FolderFeedArgs)

            __props__.__dict__["asset_names"] = asset_names
            __props__.__dict__["asset_types"] = asset_types
            if billing_project is None and not opts.urn:
                raise TypeError("Missing required property 'billing_project'")
            __props__.__dict__["billing_project"] = billing_project
            __props__.__dict__["condition"] = condition
            __props__.__dict__["content_type"] = content_type
            if feed_id is None and not opts.urn:
                raise TypeError("Missing required property 'feed_id'")
            __props__.__dict__["feed_id"] = feed_id
            if feed_output_config is None and not opts.urn:
                raise TypeError("Missing required property 'feed_output_config'")
            __props__.__dict__["feed_output_config"] = feed_output_config
            if folder is None and not opts.urn:
                raise TypeError("Missing required property 'folder'")
            __props__.__dict__["folder"] = folder
            __props__.__dict__["folder_id"] = None
            __props__.__dict__["name"] = None
        super(FolderFeed, __self__).__init__(
            'gcp:cloudasset/folderFeed:FolderFeed',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asset_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            asset_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            billing_project: Optional[pulumi.Input[str]] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['FolderFeedConditionArgs']]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            feed_id: Optional[pulumi.Input[str]] = None,
            feed_output_config: Optional[pulumi.Input[pulumi.InputType['FolderFeedFeedOutputConfigArgs']]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            folder_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'FolderFeed':
        """
        Get an existing FolderFeed resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_names: A list of the full names of the assets to receive updates. You must specify either or both of
               assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
               exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
               See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] asset_types: A list of types of the assets to receive updates. You must specify either or both of assetNames
               and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
               the feed. For example: "compute.googleapis.com/Disk"
               See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
               supported asset types.
        :param pulumi.Input[str] billing_project: The project whose identity will be used when sending messages to the
               destination pubsub topic. It also specifies the project for API
               enablement check, quota, and billing.
        :param pulumi.Input[pulumi.InputType['FolderFeedConditionArgs']] condition: A condition which determines whether an asset update should be published. If specified, an asset
               will be returned only when the expression evaluates to true. When set, expression field
               must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
               expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
               condition are optional.
               Structure is documented below.
        :param pulumi.Input[str] content_type: Asset content type. If not specified, no content but the asset name and type will be returned.
               Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        :param pulumi.Input[str] feed_id: This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        :param pulumi.Input[pulumi.InputType['FolderFeedFeedOutputConfigArgs']] feed_output_config: Output configuration for asset feed destination.
               Structure is documented below.
        :param pulumi.Input[str] folder: The folder this feed should be created in.
        :param pulumi.Input[str] folder_id: The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
        :param pulumi.Input[str] name: The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FolderFeedState.__new__(_FolderFeedState)

        __props__.__dict__["asset_names"] = asset_names
        __props__.__dict__["asset_types"] = asset_types
        __props__.__dict__["billing_project"] = billing_project
        __props__.__dict__["condition"] = condition
        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["feed_id"] = feed_id
        __props__.__dict__["feed_output_config"] = feed_output_config
        __props__.__dict__["folder"] = folder
        __props__.__dict__["folder_id"] = folder_id
        __props__.__dict__["name"] = name
        return FolderFeed(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assetNames")
    def asset_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of the full names of the assets to receive updates. You must specify either or both of
        assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
        exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
        See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
        """
        return pulumi.get(self, "asset_names")

    @property
    @pulumi.getter(name="assetTypes")
    def asset_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of types of the assets to receive updates. You must specify either or both of assetNames
        and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
        the feed. For example: "compute.googleapis.com/Disk"
        See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
        supported asset types.
        """
        return pulumi.get(self, "asset_types")

    @property
    @pulumi.getter(name="billingProject")
    def billing_project(self) -> pulumi.Output[str]:
        """
        The project whose identity will be used when sending messages to the
        destination pubsub topic. It also specifies the project for API
        enablement check, quota, and billing.
        """
        return pulumi.get(self, "billing_project")

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.FolderFeedCondition']]:
        """
        A condition which determines whether an asset update should be published. If specified, an asset
        will be returned only when the expression evaluates to true. When set, expression field
        must be a valid CEL expression on a TemporalAsset with name temporal_asset. Example: a Feed with
        expression "temporal_asset.deleted == true" will only publish Asset deletions. Other fields of
        condition are optional.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[Optional[str]]:
        """
        Asset content type. If not specified, no content but the asset name and type will be returned.
        Possible values are `CONTENT_TYPE_UNSPECIFIED`, `RESOURCE`, `IAM_POLICY`, `ORG_POLICY`, and `ACCESS_POLICY`.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="feedId")
    def feed_id(self) -> pulumi.Output[str]:
        """
        This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
        """
        return pulumi.get(self, "feed_id")

    @property
    @pulumi.getter(name="feedOutputConfig")
    def feed_output_config(self) -> pulumi.Output['outputs.FolderFeedFeedOutputConfig']:
        """
        Output configuration for asset feed destination.
        Structure is documented below.
        """
        return pulumi.get(self, "feed_output_config")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The folder this feed should be created in.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Output[str]:
        """
        The ID of the folder where this feed has been created. Both [FOLDER_NUMBER] and folders/[FOLDER_NUMBER] are accepted.
        """
        return pulumi.get(self, "folder_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The format will be folders/{folder_number}/feeds/{client-assigned_feed_identifier}.
        """
        return pulumi.get(self, "name")

