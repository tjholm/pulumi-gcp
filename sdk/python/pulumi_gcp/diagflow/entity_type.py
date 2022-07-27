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

__all__ = ['EntityTypeArgs', 'EntityType']

@pulumi.input_type
class EntityTypeArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 kind: pulumi.Input[str],
                 enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EntityType resource.
        :param pulumi.Input[str] display_name: The name of this entity type to be displayed on the console.
        :param pulumi.Input[str] kind: Indicates the kind of entity type.
               * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
               * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
               types can contain references to other entity types (with or without aliases).
               * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
               Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        :param pulumi.Input[bool] enable_fuzzy_extraction: Enables fuzzy entity extraction during classification.
        :param pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]] entities: The collection of entity entries associated with the entity type.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "kind", kind)
        if enable_fuzzy_extraction is not None:
            pulumi.set(__self__, "enable_fuzzy_extraction", enable_fuzzy_extraction)
        if entities is not None:
            pulumi.set(__self__, "entities", entities)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The name of this entity type to be displayed on the console.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[str]:
        """
        Indicates the kind of entity type.
        * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
        * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
        types can contain references to other entity types (with or without aliases).
        * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
        Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter(name="enableFuzzyExtraction")
    def enable_fuzzy_extraction(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables fuzzy entity extraction during classification.
        """
        return pulumi.get(self, "enable_fuzzy_extraction")

    @enable_fuzzy_extraction.setter
    def enable_fuzzy_extraction(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_fuzzy_extraction", value)

    @property
    @pulumi.getter
    def entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]]]:
        """
        The collection of entity entries associated with the entity type.
        Structure is documented below.
        """
        return pulumi.get(self, "entities")

    @entities.setter
    def entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]]]):
        pulumi.set(self, "entities", value)

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
class _EntityTypeState:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EntityType resources.
        :param pulumi.Input[str] display_name: The name of this entity type to be displayed on the console.
        :param pulumi.Input[bool] enable_fuzzy_extraction: Enables fuzzy entity extraction during classification.
        :param pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]] entities: The collection of entity entries associated with the entity type.
               Structure is documented below.
        :param pulumi.Input[str] kind: Indicates the kind of entity type.
               * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
               * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
               types can contain references to other entity types (with or without aliases).
               * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
               Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        :param pulumi.Input[str] name: The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if enable_fuzzy_extraction is not None:
            pulumi.set(__self__, "enable_fuzzy_extraction", enable_fuzzy_extraction)
        if entities is not None:
            pulumi.set(__self__, "entities", entities)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this entity type to be displayed on the console.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="enableFuzzyExtraction")
    def enable_fuzzy_extraction(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables fuzzy entity extraction during classification.
        """
        return pulumi.get(self, "enable_fuzzy_extraction")

    @enable_fuzzy_extraction.setter
    def enable_fuzzy_extraction(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_fuzzy_extraction", value)

    @property
    @pulumi.getter
    def entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]]]:
        """
        The collection of entity entries associated with the entity type.
        Structure is documented below.
        """
        return pulumi.get(self, "entities")

    @entities.setter
    def entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EntityTypeEntityArgs']]]]):
        pulumi.set(self, "entities", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the kind of entity type.
        * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
        * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
        types can contain references to other entity types (with or without aliases).
        * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
        Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
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


class EntityType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents an entity type. Entity types serve as a tool for extracting parameter values from natural language queries.

        To get more information about EntityType, see:

        * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.entityTypes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dialogflow/docs/)

        ## Example Usage
        ### Dialogflow Entity Type Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_agent = gcp.diagflow.Agent("basicAgent",
            display_name="example_agent",
            default_language_code="en",
            time_zone="America/New_York")
        basic_entity_type = gcp.diagflow.EntityType("basicEntityType",
            display_name="",
            kind="KIND_MAP",
            entities=[
                gcp.diagflow.EntityTypeEntityArgs(
                    value="value1",
                    synonyms=[
                        "synonym1",
                        "synonym2",
                    ],
                ),
                gcp.diagflow.EntityTypeEntityArgs(
                    value="value2",
                    synonyms=[
                        "synonym3",
                        "synonym4",
                    ],
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[basic_agent]))
        ```

        ## Import

        EntityType can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:diagflow/entityType:EntityType default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The name of this entity type to be displayed on the console.
        :param pulumi.Input[bool] enable_fuzzy_extraction: Enables fuzzy entity extraction during classification.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]] entities: The collection of entity entries associated with the entity type.
               Structure is documented below.
        :param pulumi.Input[str] kind: Indicates the kind of entity type.
               * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
               * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
               types can contain references to other entity types (with or without aliases).
               * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
               Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EntityTypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents an entity type. Entity types serve as a tool for extracting parameter values from natural language queries.

        To get more information about EntityType, see:

        * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.entityTypes)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/dialogflow/docs/)

        ## Example Usage
        ### Dialogflow Entity Type Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_agent = gcp.diagflow.Agent("basicAgent",
            display_name="example_agent",
            default_language_code="en",
            time_zone="America/New_York")
        basic_entity_type = gcp.diagflow.EntityType("basicEntityType",
            display_name="",
            kind="KIND_MAP",
            entities=[
                gcp.diagflow.EntityTypeEntityArgs(
                    value="value1",
                    synonyms=[
                        "synonym1",
                        "synonym2",
                    ],
                ),
                gcp.diagflow.EntityTypeEntityArgs(
                    value="value2",
                    synonyms=[
                        "synonym3",
                        "synonym4",
                    ],
                ),
            ],
            opts=pulumi.ResourceOptions(depends_on=[basic_agent]))
        ```

        ## Import

        EntityType can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:diagflow/entityType:EntityType default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param EntityTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EntityTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EntityTypeArgs.__new__(EntityTypeArgs)

            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["enable_fuzzy_extraction"] = enable_fuzzy_extraction
            __props__.__dict__["entities"] = entities
            if kind is None and not opts.urn:
                raise TypeError("Missing required property 'kind'")
            __props__.__dict__["kind"] = kind
            __props__.__dict__["project"] = project
            __props__.__dict__["name"] = None
        super(EntityType, __self__).__init__(
            'gcp:diagflow/entityType:EntityType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enable_fuzzy_extraction: Optional[pulumi.Input[bool]] = None,
            entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'EntityType':
        """
        Get an existing EntityType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The name of this entity type to be displayed on the console.
        :param pulumi.Input[bool] enable_fuzzy_extraction: Enables fuzzy entity extraction during classification.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EntityTypeEntityArgs']]]] entities: The collection of entity entries associated with the entity type.
               Structure is documented below.
        :param pulumi.Input[str] kind: Indicates the kind of entity type.
               * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
               * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
               types can contain references to other entity types (with or without aliases).
               * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
               Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        :param pulumi.Input[str] name: The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EntityTypeState.__new__(_EntityTypeState)

        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["enable_fuzzy_extraction"] = enable_fuzzy_extraction
        __props__.__dict__["entities"] = entities
        __props__.__dict__["kind"] = kind
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        return EntityType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The name of this entity type to be displayed on the console.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enableFuzzyExtraction")
    def enable_fuzzy_extraction(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables fuzzy entity extraction during classification.
        """
        return pulumi.get(self, "enable_fuzzy_extraction")

    @property
    @pulumi.getter
    def entities(self) -> pulumi.Output[Optional[Sequence['outputs.EntityTypeEntity']]]:
        """
        The collection of entity entries associated with the entity type.
        Structure is documented below.
        """
        return pulumi.get(self, "entities")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Indicates the kind of entity type.
        * KIND_MAP: Map entity types allow mapping of a group of synonyms to a reference value.
        * KIND_LIST: List entity types contain a set of entries that do not map to reference values. However, list entity
        types can contain references to other entity types (with or without aliases).
        * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
        Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the entity type. Format: projects/<Project ID>/agent/entityTypes/<Entity type ID>.
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

