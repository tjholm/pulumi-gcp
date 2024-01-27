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

__all__ = ['ReservationArgs', 'Reservation']

@pulumi.input_type
class ReservationArgs:
    def __init__(__self__, *,
                 slot_capacity: pulumi.Input[int],
                 autoscale: Optional[pulumi.Input['ReservationAutoscaleArgs']] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Reservation resource.
        :param pulumi.Input[int] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
               unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        :param pulumi.Input['ReservationAutoscaleArgs'] autoscale: The configuration parameters for the auto scaling feature.
               Structure is documented below.
        :param pulumi.Input[int] concurrency: Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        :param pulumi.Input[str] edition: The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query using this reservation will use idle slots from other reservations within
               the same admin project. If true, a query using this reservation will execute with the slot
               capacity specified above at most.
        :param pulumi.Input[str] location: The geographic location where the transfer config should reside.
               Examples: US, EU, asia-northeast1. The default value is US.
        :param pulumi.Input[bool] multi_region_auxiliary: Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
               If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        :param pulumi.Input[str] name: The name of the reservation. This field must only contain alphanumeric characters or dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "slot_capacity", slot_capacity)
        if autoscale is not None:
            pulumi.set(__self__, "autoscale", autoscale)
        if concurrency is not None:
            pulumi.set(__self__, "concurrency", concurrency)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if ignore_idle_slots is not None:
            pulumi.set(__self__, "ignore_idle_slots", ignore_idle_slots)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if multi_region_auxiliary is not None:
            pulumi.set(__self__, "multi_region_auxiliary", multi_region_auxiliary)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="slotCapacity")
    def slot_capacity(self) -> pulumi.Input[int]:
        """
        Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
        unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        return pulumi.get(self, "slot_capacity")

    @slot_capacity.setter
    def slot_capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "slot_capacity", value)

    @property
    @pulumi.getter
    def autoscale(self) -> Optional[pulumi.Input['ReservationAutoscaleArgs']]:
        """
        The configuration parameters for the auto scaling feature.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscale")

    @autoscale.setter
    def autoscale(self, value: Optional[pulumi.Input['ReservationAutoscaleArgs']]):
        pulumi.set(self, "autoscale", value)

    @property
    @pulumi.getter
    def concurrency(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="ignoreIdleSlots")
    def ignore_idle_slots(self) -> Optional[pulumi.Input[bool]]:
        """
        If false, any query using this reservation will use idle slots from other reservations within
        the same admin project. If true, a query using this reservation will execute with the slot
        capacity specified above at most.
        """
        return pulumi.get(self, "ignore_idle_slots")

    @ignore_idle_slots.setter
    def ignore_idle_slots(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_idle_slots", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geographic location where the transfer config should reside.
        Examples: US, EU, asia-northeast1. The default value is US.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="multiRegionAuxiliary")
    def multi_region_auxiliary(self) -> Optional[pulumi.Input[bool]]:
        """
        Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
        If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        """
        return pulumi.get(self, "multi_region_auxiliary")

    @multi_region_auxiliary.setter
    def multi_region_auxiliary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_region_auxiliary", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the reservation. This field must only contain alphanumeric characters or dash.


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


@pulumi.input_type
class _ReservationState:
    def __init__(__self__, *,
                 autoscale: Optional[pulumi.Input['ReservationAutoscaleArgs']] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 slot_capacity: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Reservation resources.
        :param pulumi.Input['ReservationAutoscaleArgs'] autoscale: The configuration parameters for the auto scaling feature.
               Structure is documented below.
        :param pulumi.Input[int] concurrency: Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        :param pulumi.Input[str] edition: The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query using this reservation will use idle slots from other reservations within
               the same admin project. If true, a query using this reservation will execute with the slot
               capacity specified above at most.
        :param pulumi.Input[str] location: The geographic location where the transfer config should reside.
               Examples: US, EU, asia-northeast1. The default value is US.
        :param pulumi.Input[bool] multi_region_auxiliary: Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
               If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        :param pulumi.Input[str] name: The name of the reservation. This field must only contain alphanumeric characters or dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[int] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
               unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        if autoscale is not None:
            pulumi.set(__self__, "autoscale", autoscale)
        if concurrency is not None:
            pulumi.set(__self__, "concurrency", concurrency)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if ignore_idle_slots is not None:
            pulumi.set(__self__, "ignore_idle_slots", ignore_idle_slots)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if multi_region_auxiliary is not None:
            pulumi.set(__self__, "multi_region_auxiliary", multi_region_auxiliary)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if slot_capacity is not None:
            pulumi.set(__self__, "slot_capacity", slot_capacity)

    @property
    @pulumi.getter
    def autoscale(self) -> Optional[pulumi.Input['ReservationAutoscaleArgs']]:
        """
        The configuration parameters for the auto scaling feature.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscale")

    @autoscale.setter
    def autoscale(self, value: Optional[pulumi.Input['ReservationAutoscaleArgs']]):
        pulumi.set(self, "autoscale", value)

    @property
    @pulumi.getter
    def concurrency(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input[str]]:
        """
        The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="ignoreIdleSlots")
    def ignore_idle_slots(self) -> Optional[pulumi.Input[bool]]:
        """
        If false, any query using this reservation will use idle slots from other reservations within
        the same admin project. If true, a query using this reservation will execute with the slot
        capacity specified above at most.
        """
        return pulumi.get(self, "ignore_idle_slots")

    @ignore_idle_slots.setter
    def ignore_idle_slots(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_idle_slots", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geographic location where the transfer config should reside.
        Examples: US, EU, asia-northeast1. The default value is US.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="multiRegionAuxiliary")
    def multi_region_auxiliary(self) -> Optional[pulumi.Input[bool]]:
        """
        Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
        If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        """
        return pulumi.get(self, "multi_region_auxiliary")

    @multi_region_auxiliary.setter
    def multi_region_auxiliary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_region_auxiliary", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the reservation. This field must only contain alphanumeric characters or dash.


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
    @pulumi.getter(name="slotCapacity")
    def slot_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
        unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        return pulumi.get(self, "slot_capacity")

    @slot_capacity.setter
    def slot_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "slot_capacity", value)


class Reservation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale: Optional[pulumi.Input[pulumi.InputType['ReservationAutoscaleArgs']]] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 slot_capacity: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        A reservation is a mechanism used to guarantee BigQuery slots to users.

        To get more information about Reservation, see:

        * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1/projects.locations.reservations/create)
        * How-to Guides
            * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)

        ## Example Usage
        ### Bigquery Reservation Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        reservation = gcp.bigquery.Reservation("reservation",
            autoscale=gcp.bigquery.ReservationAutoscaleArgs(
                max_slots=100,
            ),
            concurrency=0,
            edition="STANDARD",
            ignore_idle_slots=True,
            location="us-west2",
            slot_capacity=0)
        ```

        ## Import

        Reservation can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/reservations/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Reservation can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:bigquery/reservation:Reservation default projects/{{project}}/locations/{{location}}/reservations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/reservation:Reservation default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/reservation:Reservation default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ReservationAutoscaleArgs']] autoscale: The configuration parameters for the auto scaling feature.
               Structure is documented below.
        :param pulumi.Input[int] concurrency: Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        :param pulumi.Input[str] edition: The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query using this reservation will use idle slots from other reservations within
               the same admin project. If true, a query using this reservation will execute with the slot
               capacity specified above at most.
        :param pulumi.Input[str] location: The geographic location where the transfer config should reside.
               Examples: US, EU, asia-northeast1. The default value is US.
        :param pulumi.Input[bool] multi_region_auxiliary: Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
               If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        :param pulumi.Input[str] name: The name of the reservation. This field must only contain alphanumeric characters or dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[int] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
               unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReservationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A reservation is a mechanism used to guarantee BigQuery slots to users.

        To get more information about Reservation, see:

        * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1/projects.locations.reservations/create)
        * How-to Guides
            * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)

        ## Example Usage
        ### Bigquery Reservation Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        reservation = gcp.bigquery.Reservation("reservation",
            autoscale=gcp.bigquery.ReservationAutoscaleArgs(
                max_slots=100,
            ),
            concurrency=0,
            edition="STANDARD",
            ignore_idle_slots=True,
            location="us-west2",
            slot_capacity=0)
        ```

        ## Import

        Reservation can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/reservations/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Reservation can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:bigquery/reservation:Reservation default projects/{{project}}/locations/{{location}}/reservations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/reservation:Reservation default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:bigquery/reservation:Reservation default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ReservationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReservationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale: Optional[pulumi.Input[pulumi.InputType['ReservationAutoscaleArgs']]] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 edition: Optional[pulumi.Input[str]] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 slot_capacity: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReservationArgs.__new__(ReservationArgs)

            __props__.__dict__["autoscale"] = autoscale
            __props__.__dict__["concurrency"] = concurrency
            __props__.__dict__["edition"] = edition
            __props__.__dict__["ignore_idle_slots"] = ignore_idle_slots
            __props__.__dict__["location"] = location
            __props__.__dict__["multi_region_auxiliary"] = multi_region_auxiliary
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if slot_capacity is None and not opts.urn:
                raise TypeError("Missing required property 'slot_capacity'")
            __props__.__dict__["slot_capacity"] = slot_capacity
        super(Reservation, __self__).__init__(
            'gcp:bigquery/reservation:Reservation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscale: Optional[pulumi.Input[pulumi.InputType['ReservationAutoscaleArgs']]] = None,
            concurrency: Optional[pulumi.Input[int]] = None,
            edition: Optional[pulumi.Input[str]] = None,
            ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
            location: Optional[pulumi.Input[str]] = None,
            multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            slot_capacity: Optional[pulumi.Input[int]] = None) -> 'Reservation':
        """
        Get an existing Reservation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ReservationAutoscaleArgs']] autoscale: The configuration parameters for the auto scaling feature.
               Structure is documented below.
        :param pulumi.Input[int] concurrency: Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        :param pulumi.Input[str] edition: The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query using this reservation will use idle slots from other reservations within
               the same admin project. If true, a query using this reservation will execute with the slot
               capacity specified above at most.
        :param pulumi.Input[str] location: The geographic location where the transfer config should reside.
               Examples: US, EU, asia-northeast1. The default value is US.
        :param pulumi.Input[bool] multi_region_auxiliary: Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
               If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        :param pulumi.Input[str] name: The name of the reservation. This field must only contain alphanumeric characters or dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[int] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
               unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReservationState.__new__(_ReservationState)

        __props__.__dict__["autoscale"] = autoscale
        __props__.__dict__["concurrency"] = concurrency
        __props__.__dict__["edition"] = edition
        __props__.__dict__["ignore_idle_slots"] = ignore_idle_slots
        __props__.__dict__["location"] = location
        __props__.__dict__["multi_region_auxiliary"] = multi_region_auxiliary
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["slot_capacity"] = slot_capacity
        return Reservation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def autoscale(self) -> pulumi.Output[Optional['outputs.ReservationAutoscale']]:
        """
        The configuration parameters for the auto scaling feature.
        Structure is documented below.
        """
        return pulumi.get(self, "autoscale")

    @property
    @pulumi.getter
    def concurrency(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
        """
        return pulumi.get(self, "concurrency")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="ignoreIdleSlots")
    def ignore_idle_slots(self) -> pulumi.Output[Optional[bool]]:
        """
        If false, any query using this reservation will use idle slots from other reservations within
        the same admin project. If true, a query using this reservation will execute with the slot
        capacity specified above at most.
        """
        return pulumi.get(self, "ignore_idle_slots")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The geographic location where the transfer config should reside.
        Examples: US, EU, asia-northeast1. The default value is US.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="multiRegionAuxiliary")
    def multi_region_auxiliary(self) -> pulumi.Output[Optional[bool]]:
        """
        Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
        If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
        """
        return pulumi.get(self, "multi_region_auxiliary")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the reservation. This field must only contain alphanumeric characters or dash.


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
    @pulumi.getter(name="slotCapacity")
    def slot_capacity(self) -> pulumi.Output[int]:
        """
        Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
        unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
        """
        return pulumi.get(self, "slot_capacity")

