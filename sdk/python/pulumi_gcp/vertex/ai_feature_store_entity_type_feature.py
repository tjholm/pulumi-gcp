# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AiFeatureStoreEntityTypeFeatureArgs', 'AiFeatureStoreEntityTypeFeature']

@pulumi.input_type
class AiFeatureStoreEntityTypeFeatureArgs:
    def __init__(__self__, *,
                 entitytype: pulumi.Input[str],
                 value_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AiFeatureStoreEntityTypeFeature resource.
        :param pulumi.Input[str] entitytype: The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
               
               
               - - -
        :param pulumi.Input[str] value_type: Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        :param pulumi.Input[str] description: Description of the feature.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the feature.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        """
        pulumi.set(__self__, "entitytype", entitytype)
        pulumi.set(__self__, "value_type", value_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def entitytype(self) -> pulumi.Input[str]:
        """
        The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.


        - - -
        """
        return pulumi.get(self, "entitytype")

    @entitytype.setter
    def entitytype(self, value: pulumi.Input[str]):
        pulumi.set(self, "entitytype", value)

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> pulumi.Input[str]:
        """
        Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        """
        return pulumi.get(self, "value_type")

    @value_type.setter
    def value_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "value_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the feature.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the feature.

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
        The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AiFeatureStoreEntityTypeFeatureState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 entitytype: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 value_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AiFeatureStoreEntityTypeFeature resources.
        :param pulumi.Input[str] create_time: The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        :param pulumi.Input[str] description: Description of the feature.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] entitytype: The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
               
               
               - - -
        :param pulumi.Input[str] etag: Used to perform consistent read-modify-write updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the feature.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] region: The region of the feature
        :param pulumi.Input[str] update_time: The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        :param pulumi.Input[str] value_type: Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if effective_labels is not None:
            pulumi.set(__self__, "effective_labels", effective_labels)
        if entitytype is not None:
            pulumi.set(__self__, "entitytype", entitytype)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pulumi_labels is not None:
            pulumi.set(__self__, "pulumi_labels", pulumi_labels)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if value_type is not None:
            pulumi.set(__self__, "value_type", value_type)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the feature.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    @pulumi.getter
    def entitytype(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.


        - - -
        """
        return pulumi.get(self, "entitytype")

    @entitytype.setter
    def entitytype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entitytype", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        Used to perform consistent read-modify-write updates.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to the feature.

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
        The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
        The region of the feature
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        """
        return pulumi.get(self, "value_type")

    @value_type.setter
    def value_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value_type", value)


class AiFeatureStoreEntityTypeFeature(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entitytype: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 value_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Feature Metadata information that describes an attribute of an entity type. For example, apple is an entity type, and color is a feature that describes apple.

        To get more information about FeaturestoreEntitytypeFeature, see:

        * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vertex-ai/docs)

        ## Example Usage
        ### Vertex Ai Featurestore Entitytype Feature

        ```python
        import pulumi
        import pulumi_gcp as gcp

        featurestore = gcp.vertex.AiFeatureStore("featurestore",
            labels={
                "foo": "bar",
            },
            region="us-central1",
            online_serving_config=gcp.vertex.AiFeatureStoreOnlineServingConfigArgs(
                fixed_node_count=2,
            ))
        entity = gcp.vertex.AiFeatureStoreEntityType("entity",
            labels={
                "foo": "bar",
            },
            featurestore=featurestore.id)
        feature = gcp.vertex.AiFeatureStoreEntityTypeFeature("feature",
            labels={
                "foo": "bar",
            },
            entitytype=entity.id,
            value_type="INT64_ARRAY")
        ```
        ### Vertex Ai Featurestore Entitytype Feature With Beta Fields

        ```python
        import pulumi
        import pulumi_gcp as gcp

        featurestore = gcp.vertex.AiFeatureStore("featurestore",
            labels={
                "foo": "bar",
            },
            region="us-central1",
            online_serving_config=gcp.vertex.AiFeatureStoreOnlineServingConfigArgs(
                fixed_node_count=2,
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        entity = gcp.vertex.AiFeatureStoreEntityType("entity",
            labels={
                "foo": "bar",
            },
            featurestore=featurestore.id,
            monitoring_config=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigArgs(
                snapshot_analysis=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs(
                    disabled=False,
                    monitoring_interval="86400s",
                ),
                categorical_threshold_config=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs(
                    value=0.3,
                ),
                numerical_threshold_config=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs(
                    value=0.3,
                ),
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        feature = gcp.vertex.AiFeatureStoreEntityTypeFeature("feature",
            labels={
                "foo": "bar",
            },
            entitytype=entity.id,
            value_type="INT64_ARRAY",
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        FeaturestoreEntitytypeFeature can be imported using any of these accepted formats:

        ```sh
         $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature default {{entitytype}}/features/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the feature.
        :param pulumi.Input[str] entitytype: The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
               
               
               - - -
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the feature.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        :param pulumi.Input[str] value_type: Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AiFeatureStoreEntityTypeFeatureArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Feature Metadata information that describes an attribute of an entity type. For example, apple is an entity type, and color is a feature that describes apple.

        To get more information about FeaturestoreEntitytypeFeature, see:

        * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/vertex-ai/docs)

        ## Example Usage
        ### Vertex Ai Featurestore Entitytype Feature

        ```python
        import pulumi
        import pulumi_gcp as gcp

        featurestore = gcp.vertex.AiFeatureStore("featurestore",
            labels={
                "foo": "bar",
            },
            region="us-central1",
            online_serving_config=gcp.vertex.AiFeatureStoreOnlineServingConfigArgs(
                fixed_node_count=2,
            ))
        entity = gcp.vertex.AiFeatureStoreEntityType("entity",
            labels={
                "foo": "bar",
            },
            featurestore=featurestore.id)
        feature = gcp.vertex.AiFeatureStoreEntityTypeFeature("feature",
            labels={
                "foo": "bar",
            },
            entitytype=entity.id,
            value_type="INT64_ARRAY")
        ```
        ### Vertex Ai Featurestore Entitytype Feature With Beta Fields

        ```python
        import pulumi
        import pulumi_gcp as gcp

        featurestore = gcp.vertex.AiFeatureStore("featurestore",
            labels={
                "foo": "bar",
            },
            region="us-central1",
            online_serving_config=gcp.vertex.AiFeatureStoreOnlineServingConfigArgs(
                fixed_node_count=2,
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        entity = gcp.vertex.AiFeatureStoreEntityType("entity",
            labels={
                "foo": "bar",
            },
            featurestore=featurestore.id,
            monitoring_config=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigArgs(
                snapshot_analysis=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs(
                    disabled=False,
                    monitoring_interval="86400s",
                ),
                categorical_threshold_config=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs(
                    value=0.3,
                ),
                numerical_threshold_config=gcp.vertex.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs(
                    value=0.3,
                ),
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        feature = gcp.vertex.AiFeatureStoreEntityTypeFeature("feature",
            labels={
                "foo": "bar",
            },
            entitytype=entity.id,
            value_type="INT64_ARRAY",
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        FeaturestoreEntitytypeFeature can be imported using any of these accepted formats:

        ```sh
         $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature default {{entitytype}}/features/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param AiFeatureStoreEntityTypeFeatureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AiFeatureStoreEntityTypeFeatureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entitytype: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 value_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AiFeatureStoreEntityTypeFeatureArgs.__new__(AiFeatureStoreEntityTypeFeatureArgs)

            __props__.__dict__["description"] = description
            if entitytype is None and not opts.urn:
                raise TypeError("Missing required property 'entitytype'")
            __props__.__dict__["entitytype"] = entitytype
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if value_type is None and not opts.urn:
                raise TypeError("Missing required property 'value_type'")
            __props__.__dict__["value_type"] = value_type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["effective_labels"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["pulumi_labels"] = None
            __props__.__dict__["region"] = None
            __props__.__dict__["update_time"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["effectiveLabels", "pulumiLabels"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AiFeatureStoreEntityTypeFeature, __self__).__init__(
            'gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            entitytype: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            value_type: Optional[pulumi.Input[str]] = None) -> 'AiFeatureStoreEntityTypeFeature':
        """
        Get an existing AiFeatureStoreEntityTypeFeature resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        :param pulumi.Input[str] description: Description of the feature.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] entitytype: The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
               
               
               - - -
        :param pulumi.Input[str] etag: Used to perform consistent read-modify-write updates.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to the feature.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] region: The region of the feature
        :param pulumi.Input[str] update_time: The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        :param pulumi.Input[str] value_type: Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AiFeatureStoreEntityTypeFeatureState.__new__(_AiFeatureStoreEntityTypeFeatureState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["effective_labels"] = effective_labels
        __props__.__dict__["entitytype"] = entitytype
        __props__.__dict__["etag"] = etag
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["pulumi_labels"] = pulumi_labels
        __props__.__dict__["region"] = region
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["value_type"] = value_type
        return AiFeatureStoreEntityTypeFeature(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the feature.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @property
    @pulumi.getter
    def entitytype(self) -> pulumi.Output[str]:
        """
        The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.


        - - -
        """
        return pulumi.get(self, "entitytype")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Used to perform consistent read-modify-write updates.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to the feature.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        """
        return pulumi.get(self, "name")

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
    def region(self) -> pulumi.Output[str]:
        """
        The region of the feature
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> pulumi.Output[str]:
        """
        Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        """
        return pulumi.get(self, "value_type")

