# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'DataExchangeIamBindingConditionArgs',
    'DataExchangeIamBindingConditionArgsDict',
    'DataExchangeIamMemberConditionArgs',
    'DataExchangeIamMemberConditionArgsDict',
    'ListingBigqueryDatasetArgs',
    'ListingBigqueryDatasetArgsDict',
    'ListingDataProviderArgs',
    'ListingDataProviderArgsDict',
    'ListingIamBindingConditionArgs',
    'ListingIamBindingConditionArgsDict',
    'ListingIamMemberConditionArgs',
    'ListingIamMemberConditionArgsDict',
    'ListingPublisherArgs',
    'ListingPublisherArgsDict',
    'ListingRestrictedExportConfigArgs',
    'ListingRestrictedExportConfigArgsDict',
]

MYPY = False

if not MYPY:
    class DataExchangeIamBindingConditionArgsDict(TypedDict):
        expression: pulumi.Input[str]
        title: pulumi.Input[str]
        description: NotRequired[pulumi.Input[str]]
elif False:
    DataExchangeIamBindingConditionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataExchangeIamBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


if not MYPY:
    class DataExchangeIamMemberConditionArgsDict(TypedDict):
        expression: pulumi.Input[str]
        title: pulumi.Input[str]
        description: NotRequired[pulumi.Input[str]]
elif False:
    DataExchangeIamMemberConditionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DataExchangeIamMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


if not MYPY:
    class ListingBigqueryDatasetArgsDict(TypedDict):
        dataset: pulumi.Input[str]
        """
        Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123

        - - -
        """
elif False:
    ListingBigqueryDatasetArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ListingBigqueryDatasetArgs:
    def __init__(__self__, *,
                 dataset: pulumi.Input[str]):
        """
        :param pulumi.Input[str] dataset: Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123
               
               - - -
        """
        pulumi.set(__self__, "dataset", dataset)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Input[str]:
        """
        Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123

        - - -
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset", value)


if not MYPY:
    class ListingDataProviderArgsDict(TypedDict):
        name: pulumi.Input[str]
        """
        Name of the data provider.
        """
        primary_contact: NotRequired[pulumi.Input[str]]
        """
        Email or URL of the data provider.
        """
elif False:
    ListingDataProviderArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ListingDataProviderArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 primary_contact: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Name of the data provider.
        :param pulumi.Input[str] primary_contact: Email or URL of the data provider.
        """
        pulumi.set(__self__, "name", name)
        if primary_contact is not None:
            pulumi.set(__self__, "primary_contact", primary_contact)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the data provider.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="primaryContact")
    def primary_contact(self) -> Optional[pulumi.Input[str]]:
        """
        Email or URL of the data provider.
        """
        return pulumi.get(self, "primary_contact")

    @primary_contact.setter
    def primary_contact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_contact", value)


if not MYPY:
    class ListingIamBindingConditionArgsDict(TypedDict):
        expression: pulumi.Input[str]
        title: pulumi.Input[str]
        description: NotRequired[pulumi.Input[str]]
elif False:
    ListingIamBindingConditionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ListingIamBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


if not MYPY:
    class ListingIamMemberConditionArgsDict(TypedDict):
        expression: pulumi.Input[str]
        title: pulumi.Input[str]
        description: NotRequired[pulumi.Input[str]]
elif False:
    ListingIamMemberConditionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ListingIamMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


if not MYPY:
    class ListingPublisherArgsDict(TypedDict):
        name: pulumi.Input[str]
        """
        Name of the listing publisher.
        """
        primary_contact: NotRequired[pulumi.Input[str]]
        """
        Email or URL of the listing publisher.
        """
elif False:
    ListingPublisherArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ListingPublisherArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 primary_contact: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] name: Name of the listing publisher.
        :param pulumi.Input[str] primary_contact: Email or URL of the listing publisher.
        """
        pulumi.set(__self__, "name", name)
        if primary_contact is not None:
            pulumi.set(__self__, "primary_contact", primary_contact)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Name of the listing publisher.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="primaryContact")
    def primary_contact(self) -> Optional[pulumi.Input[str]]:
        """
        Email or URL of the listing publisher.
        """
        return pulumi.get(self, "primary_contact")

    @primary_contact.setter
    def primary_contact(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_contact", value)


if not MYPY:
    class ListingRestrictedExportConfigArgsDict(TypedDict):
        enabled: NotRequired[pulumi.Input[bool]]
        """
        If true, enable restricted export.
        """
        restrict_query_result: NotRequired[pulumi.Input[bool]]
        """
        If true, restrict export of query result derived from restricted linked dataset table.
        """
elif False:
    ListingRestrictedExportConfigArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ListingRestrictedExportConfigArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 restrict_query_result: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] enabled: If true, enable restricted export.
        :param pulumi.Input[bool] restrict_query_result: If true, restrict export of query result derived from restricted linked dataset table.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if restrict_query_result is not None:
            pulumi.set(__self__, "restrict_query_result", restrict_query_result)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, enable restricted export.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="restrictQueryResult")
    def restrict_query_result(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, restrict export of query result derived from restricted linked dataset table.
        """
        return pulumi.get(self, "restrict_query_result")

    @restrict_query_result.setter
    def restrict_query_result(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "restrict_query_result", value)


