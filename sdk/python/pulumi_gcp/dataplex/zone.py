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

__all__ = ['ZoneArgs', 'Zone']

@pulumi.input_type
class ZoneArgs:
    def __init__(__self__, *,
                 discovery_spec: pulumi.Input['ZoneDiscoverySpecArgs'],
                 lake: pulumi.Input[str],
                 location: pulumi.Input[str],
                 resource_spec: pulumi.Input['ZoneResourceSpecArgs'],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Zone resource.
        :param pulumi.Input['ZoneDiscoverySpecArgs'] discovery_spec: Required. Specification of the discovery feature applied to data in this zone.
        :param pulumi.Input[str] lake: The lake for the resource
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input['ZoneResourceSpecArgs'] resource_spec: Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        :param pulumi.Input[str] type: Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        :param pulumi.Input[str] description: Optional. Description of the zone.
        :param pulumi.Input[str] display_name: Optional. User friendly display name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User defined labels for the zone.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of the zone.
        :param pulumi.Input[str] project: The project for the resource
        """
        pulumi.set(__self__, "discovery_spec", discovery_spec)
        pulumi.set(__self__, "lake", lake)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "resource_spec", resource_spec)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="discoverySpec")
    def discovery_spec(self) -> pulumi.Input['ZoneDiscoverySpecArgs']:
        """
        Required. Specification of the discovery feature applied to data in this zone.
        """
        return pulumi.get(self, "discovery_spec")

    @discovery_spec.setter
    def discovery_spec(self, value: pulumi.Input['ZoneDiscoverySpecArgs']):
        pulumi.set(self, "discovery_spec", value)

    @property
    @pulumi.getter
    def lake(self) -> pulumi.Input[str]:
        """
        The lake for the resource
        """
        return pulumi.get(self, "lake")

    @lake.setter
    def lake(self, value: pulumi.Input[str]):
        pulumi.set(self, "lake", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="resourceSpec")
    def resource_spec(self) -> pulumi.Input['ZoneResourceSpecArgs']:
        """
        Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        """
        return pulumi.get(self, "resource_spec")

    @resource_spec.setter
    def resource_spec(self, value: pulumi.Input['ZoneResourceSpecArgs']):
        pulumi.set(self, "resource_spec", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the zone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User defined labels for the zone.

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
        The name of the zone.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ZoneState:
    def __init__(__self__, *,
                 asset_statuses: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAssetStatusArgs']]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discovery_spec: Optional[pulumi.Input['ZoneDiscoverySpecArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 effective_labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lake: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pulumi_labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 resource_spec: Optional[pulumi.Input['ZoneResourceSpecArgs']] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Zone resources.
        :param pulumi.Input[Sequence[pulumi.Input['ZoneAssetStatusArgs']]] asset_statuses: Output only. Aggregated status of the underlying assets of the zone.
        :param pulumi.Input[str] create_time: Output only. The time when the zone was created.
        :param pulumi.Input[str] description: Optional. Description of the zone.
        :param pulumi.Input['ZoneDiscoverySpecArgs'] discovery_spec: Required. Specification of the discovery feature applied to data in this zone.
        :param pulumi.Input[str] display_name: Optional. User friendly display name.
        :param pulumi.Input[Mapping[str, Any]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User defined labels for the zone.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] lake: The lake for the resource
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[str] name: The name of the zone.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[Mapping[str, Any]] pulumi_labels: The combination of labels configured directly on the resource and default labels configured on the provider.
        :param pulumi.Input['ZoneResourceSpecArgs'] resource_spec: Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        :param pulumi.Input[str] state: Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        :param pulumi.Input[str] type: Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        :param pulumi.Input[str] uid: Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        :param pulumi.Input[str] update_time: Output only. The time when the zone was last updated.
        """
        if asset_statuses is not None:
            pulumi.set(__self__, "asset_statuses", asset_statuses)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if discovery_spec is not None:
            pulumi.set(__self__, "discovery_spec", discovery_spec)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if effective_labels is not None:
            pulumi.set(__self__, "effective_labels", effective_labels)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if lake is not None:
            pulumi.set(__self__, "lake", lake)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if pulumi_labels is not None:
            pulumi.set(__self__, "pulumi_labels", pulumi_labels)
        if resource_spec is not None:
            pulumi.set(__self__, "resource_spec", resource_spec)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="assetStatuses")
    def asset_statuses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAssetStatusArgs']]]]:
        """
        Output only. Aggregated status of the underlying assets of the zone.
        """
        return pulumi.get(self, "asset_statuses")

    @asset_statuses.setter
    def asset_statuses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ZoneAssetStatusArgs']]]]):
        pulumi.set(self, "asset_statuses", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. The time when the zone was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the zone.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="discoverySpec")
    def discovery_spec(self) -> Optional[pulumi.Input['ZoneDiscoverySpecArgs']]:
        """
        Required. Specification of the discovery feature applied to data in this zone.
        """
        return pulumi.get(self, "discovery_spec")

    @discovery_spec.setter
    def discovery_spec(self, value: Optional[pulumi.Input['ZoneDiscoverySpecArgs']]):
        pulumi.set(self, "discovery_spec", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @effective_labels.setter
    def effective_labels(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "effective_labels", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User defined labels for the zone.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def lake(self) -> Optional[pulumi.Input[str]]:
        """
        The lake for the resource
        """
        return pulumi.get(self, "lake")

    @lake.setter
    def lake(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lake", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the zone.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The combination of labels configured directly on the resource and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @pulumi_labels.setter
    def pulumi_labels(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "pulumi_labels", value)

    @property
    @pulumi.getter(name="resourceSpec")
    def resource_spec(self) -> Optional[pulumi.Input['ZoneResourceSpecArgs']]:
        """
        Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        """
        return pulumi.get(self, "resource_spec")

    @resource_spec.setter
    def resource_spec(self, value: Optional[pulumi.Input['ZoneResourceSpecArgs']]):
        pulumi.set(self, "resource_spec", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. The time when the zone was last updated.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Zone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discovery_spec: Optional[pulumi.Input[pulumi.InputType['ZoneDiscoverySpecArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lake: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_spec: Optional[pulumi.Input[pulumi.InputType['ZoneResourceSpecArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The Dataplex Zone resource

        ## Example Usage
        ### Basic_zone
        A basic example of a dataplex zone
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.dataplex.Lake("basic",
            location="us-west1",
            description="Lake for DCL",
            display_name="Lake for DCL",
            project="my-project-name",
            labels={
                "my-lake": "exists",
            })
        primary = gcp.dataplex.Zone("primary",
            discovery_spec=gcp.dataplex.ZoneDiscoverySpecArgs(
                enabled=False,
            ),
            lake=basic.name,
            location="us-west1",
            resource_spec=gcp.dataplex.ZoneResourceSpecArgs(
                location_type="MULTI_REGION",
            ),
            type="RAW",
            description="Zone for DCL",
            display_name="Zone for DCL",
            project="my-project-name",
            labels={})
        ```

        ## Import

        Zone can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}` * `{{project}}/{{location}}/{{lake}}/{{name}}` * `{{location}}/{{lake}}/{{name}}` When using the `pulumi import` command, Zone can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:dataplex/zone:Zone default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
        ```

        ```sh
         $ pulumi import gcp:dataplex/zone:Zone default {{project}}/{{location}}/{{lake}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:dataplex/zone:Zone default {{location}}/{{lake}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. Description of the zone.
        :param pulumi.Input[pulumi.InputType['ZoneDiscoverySpecArgs']] discovery_spec: Required. Specification of the discovery feature applied to data in this zone.
        :param pulumi.Input[str] display_name: Optional. User friendly display name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User defined labels for the zone.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] lake: The lake for the resource
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[str] name: The name of the zone.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[pulumi.InputType['ZoneResourceSpecArgs']] resource_spec: Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        :param pulumi.Input[str] type: Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Dataplex Zone resource

        ## Example Usage
        ### Basic_zone
        A basic example of a dataplex zone
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.dataplex.Lake("basic",
            location="us-west1",
            description="Lake for DCL",
            display_name="Lake for DCL",
            project="my-project-name",
            labels={
                "my-lake": "exists",
            })
        primary = gcp.dataplex.Zone("primary",
            discovery_spec=gcp.dataplex.ZoneDiscoverySpecArgs(
                enabled=False,
            ),
            lake=basic.name,
            location="us-west1",
            resource_spec=gcp.dataplex.ZoneResourceSpecArgs(
                location_type="MULTI_REGION",
            ),
            type="RAW",
            description="Zone for DCL",
            display_name="Zone for DCL",
            project="my-project-name",
            labels={})
        ```

        ## Import

        Zone can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}` * `{{project}}/{{location}}/{{lake}}/{{name}}` * `{{location}}/{{lake}}/{{name}}` When using the `pulumi import` command, Zone can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:dataplex/zone:Zone default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
        ```

        ```sh
         $ pulumi import gcp:dataplex/zone:Zone default {{project}}/{{location}}/{{lake}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:dataplex/zone:Zone default {{location}}/{{lake}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 discovery_spec: Optional[pulumi.Input[pulumi.InputType['ZoneDiscoverySpecArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lake: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 resource_spec: Optional[pulumi.Input[pulumi.InputType['ZoneResourceSpecArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ZoneArgs.__new__(ZoneArgs)

            __props__.__dict__["description"] = description
            if discovery_spec is None and not opts.urn:
                raise TypeError("Missing required property 'discovery_spec'")
            __props__.__dict__["discovery_spec"] = discovery_spec
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["labels"] = labels
            if lake is None and not opts.urn:
                raise TypeError("Missing required property 'lake'")
            __props__.__dict__["lake"] = lake
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if resource_spec is None and not opts.urn:
                raise TypeError("Missing required property 'resource_spec'")
            __props__.__dict__["resource_spec"] = resource_spec
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["asset_statuses"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["effective_labels"] = None
            __props__.__dict__["pulumi_labels"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["effectiveLabels", "pulumiLabels"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Zone, __self__).__init__(
            'gcp:dataplex/zone:Zone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asset_statuses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAssetStatusArgs']]]]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            discovery_spec: Optional[pulumi.Input[pulumi.InputType['ZoneDiscoverySpecArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            effective_labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            lake: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            pulumi_labels: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            resource_spec: Optional[pulumi.Input[pulumi.InputType['ZoneResourceSpecArgs']]] = None,
            state: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uid: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Zone':
        """
        Get an existing Zone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ZoneAssetStatusArgs']]]] asset_statuses: Output only. Aggregated status of the underlying assets of the zone.
        :param pulumi.Input[str] create_time: Output only. The time when the zone was created.
        :param pulumi.Input[str] description: Optional. Description of the zone.
        :param pulumi.Input[pulumi.InputType['ZoneDiscoverySpecArgs']] discovery_spec: Required. Specification of the discovery feature applied to data in this zone.
        :param pulumi.Input[str] display_name: Optional. User friendly display name.
        :param pulumi.Input[Mapping[str, Any]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User defined labels for the zone.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] lake: The lake for the resource
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[str] name: The name of the zone.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[Mapping[str, Any]] pulumi_labels: The combination of labels configured directly on the resource and default labels configured on the provider.
        :param pulumi.Input[pulumi.InputType['ZoneResourceSpecArgs']] resource_spec: Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        :param pulumi.Input[str] state: Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        :param pulumi.Input[str] type: Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        :param pulumi.Input[str] uid: Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        :param pulumi.Input[str] update_time: Output only. The time when the zone was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ZoneState.__new__(_ZoneState)

        __props__.__dict__["asset_statuses"] = asset_statuses
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["discovery_spec"] = discovery_spec
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["effective_labels"] = effective_labels
        __props__.__dict__["labels"] = labels
        __props__.__dict__["lake"] = lake
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["pulumi_labels"] = pulumi_labels
        __props__.__dict__["resource_spec"] = resource_spec
        __props__.__dict__["state"] = state
        __props__.__dict__["type"] = type
        __props__.__dict__["uid"] = uid
        __props__.__dict__["update_time"] = update_time
        return Zone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="assetStatuses")
    def asset_statuses(self) -> pulumi.Output[Sequence['outputs.ZoneAssetStatus']]:
        """
        Output only. Aggregated status of the underlying assets of the zone.
        """
        return pulumi.get(self, "asset_statuses")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Output only. The time when the zone was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. Description of the zone.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discoverySpec")
    def discovery_spec(self) -> pulumi.Output['outputs.ZoneDiscoverySpec']:
        """
        Required. Specification of the discovery feature applied to data in this zone.
        """
        return pulumi.get(self, "discovery_spec")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Optional. User defined labels for the zone.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def lake(self) -> pulumi.Output[str]:
        """
        The lake for the resource
        """
        return pulumi.get(self, "lake")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the zone.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        The combination of labels configured directly on the resource and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @property
    @pulumi.getter(name="resourceSpec")
    def resource_spec(self) -> pulumi.Output['outputs.ZoneResourceSpec']:
        """
        Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        """
        return pulumi.get(self, "resource_spec")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Output only. The time when the zone was last updated.
        """
        return pulumi.get(self, "update_time")

