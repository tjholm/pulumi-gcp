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

__all__ = [
    'FieldIndexConfig',
    'FieldIndexConfigIndex',
    'FieldTtlConfig',
    'IndexField',
]

@pulumi.output_type
class FieldIndexConfig(dict):
    def __init__(__self__, *,
                 indexes: Optional[Sequence['outputs.FieldIndexConfigIndex']] = None):
        """
        :param Sequence['FieldIndexConfigIndexArgs'] indexes: The indexes to configure on the field. Order or array contains must be specified.
               Structure is documented below.
        """
        if indexes is not None:
            pulumi.set(__self__, "indexes", indexes)

    @property
    @pulumi.getter
    def indexes(self) -> Optional[Sequence['outputs.FieldIndexConfigIndex']]:
        """
        The indexes to configure on the field. Order or array contains must be specified.
        Structure is documented below.
        """
        return pulumi.get(self, "indexes")


@pulumi.output_type
class FieldIndexConfigIndex(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "arrayConfig":
            suggest = "array_config"
        elif key == "queryScope":
            suggest = "query_scope"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in FieldIndexConfigIndex. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        FieldIndexConfigIndex.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        FieldIndexConfigIndex.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 array_config: Optional[str] = None,
                 order: Optional[str] = None,
                 query_scope: Optional[str] = None):
        """
        :param str array_config: Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
               be specified.
               Possible values are: `CONTAINS`.
        :param str order: Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=, !=.
               Only one of `order` and `arrayConfig` can be specified.
               Possible values are: `ASCENDING`, `DESCENDING`.
        :param str query_scope: The scope at which a query is run. Collection scoped queries require you specify
               the collection at query time. Collection group scope allows queries across all
               collections with the same id.
               Default value is `COLLECTION`.
               Possible values are: `COLLECTION`, `COLLECTION_GROUP`.
        """
        if array_config is not None:
            pulumi.set(__self__, "array_config", array_config)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if query_scope is not None:
            pulumi.set(__self__, "query_scope", query_scope)

    @property
    @pulumi.getter(name="arrayConfig")
    def array_config(self) -> Optional[str]:
        """
        Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
        be specified.
        Possible values are: `CONTAINS`.
        """
        return pulumi.get(self, "array_config")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        """
        Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=, !=.
        Only one of `order` and `arrayConfig` can be specified.
        Possible values are: `ASCENDING`, `DESCENDING`.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> Optional[str]:
        """
        The scope at which a query is run. Collection scoped queries require you specify
        the collection at query time. Collection group scope allows queries across all
        collections with the same id.
        Default value is `COLLECTION`.
        Possible values are: `COLLECTION`, `COLLECTION_GROUP`.
        """
        return pulumi.get(self, "query_scope")


@pulumi.output_type
class FieldTtlConfig(dict):
    def __init__(__self__, *,
                 state: Optional[str] = None):
        """
        :param str state: (Output)
               The state of the TTL configuration.
        """
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        (Output)
        The state of the TTL configuration.
        """
        return pulumi.get(self, "state")


@pulumi.output_type
class IndexField(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "arrayConfig":
            suggest = "array_config"
        elif key == "fieldPath":
            suggest = "field_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in IndexField. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        IndexField.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        IndexField.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 array_config: Optional[str] = None,
                 field_path: Optional[str] = None,
                 order: Optional[str] = None):
        """
        :param str array_config: Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
               be specified.
               Possible values are: `CONTAINS`.
        :param str field_path: Name of the field.
        :param str order: Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
               Only one of `order` and `arrayConfig` can be specified.
               Possible values are: `ASCENDING`, `DESCENDING`.
        """
        if array_config is not None:
            pulumi.set(__self__, "array_config", array_config)
        if field_path is not None:
            pulumi.set(__self__, "field_path", field_path)
        if order is not None:
            pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter(name="arrayConfig")
    def array_config(self) -> Optional[str]:
        """
        Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
        be specified.
        Possible values are: `CONTAINS`.
        """
        return pulumi.get(self, "array_config")

    @property
    @pulumi.getter(name="fieldPath")
    def field_path(self) -> Optional[str]:
        """
        Name of the field.
        """
        return pulumi.get(self, "field_path")

    @property
    @pulumi.getter
    def order(self) -> Optional[str]:
        """
        Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
        Only one of `order` and `arrayConfig` can be specified.
        Possible values are: `ASCENDING`, `DESCENDING`.
        """
        return pulumi.get(self, "order")


