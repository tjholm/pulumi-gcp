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

__all__ = ['ServicePerimetersArgs', 'ServicePerimeters']

@pulumi.input_type
class ServicePerimetersArgs:
    def __init__(__self__, *,
                 parent: pulumi.Input[str],
                 service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]] = None):
        """
        The set of arguments for constructing a ServicePerimeters resource.
        :param pulumi.Input[str] parent: The AccessPolicy this ServicePerimeter lives in.
               Format: accessPolicies/{policy_id}
               
               
               - - -
        :param pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]] service_perimeters: The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
               Structure is documented below.
        """
        ServicePerimetersArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            parent=parent,
            service_perimeters=service_perimeters,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             parent: Optional[pulumi.Input[str]] = None,
             service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if parent is None:
            raise TypeError("Missing 'parent' argument")
        if service_perimeters is None and 'servicePerimeters' in kwargs:
            service_perimeters = kwargs['servicePerimeters']

        _setter("parent", parent)
        if service_perimeters is not None:
            _setter("service_perimeters", service_perimeters)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The AccessPolicy this ServicePerimeter lives in.
        Format: accessPolicies/{policy_id}


        - - -
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="servicePerimeters")
    def service_perimeters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]]:
        """
        The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
        Structure is documented below.
        """
        return pulumi.get(self, "service_perimeters")

    @service_perimeters.setter
    def service_perimeters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]]):
        pulumi.set(self, "service_perimeters", value)


@pulumi.input_type
class _ServicePerimetersState:
    def __init__(__self__, *,
                 parent: Optional[pulumi.Input[str]] = None,
                 service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]] = None):
        """
        Input properties used for looking up and filtering ServicePerimeters resources.
        :param pulumi.Input[str] parent: The AccessPolicy this ServicePerimeter lives in.
               Format: accessPolicies/{policy_id}
               
               
               - - -
        :param pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]] service_perimeters: The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
               Structure is documented below.
        """
        _ServicePerimetersState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            parent=parent,
            service_perimeters=service_perimeters,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             parent: Optional[pulumi.Input[str]] = None,
             service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if service_perimeters is None and 'servicePerimeters' in kwargs:
            service_perimeters = kwargs['servicePerimeters']

        if parent is not None:
            _setter("parent", parent)
        if service_perimeters is not None:
            _setter("service_perimeters", service_perimeters)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The AccessPolicy this ServicePerimeter lives in.
        Format: accessPolicies/{policy_id}


        - - -
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="servicePerimeters")
    def service_perimeters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]]:
        """
        The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
        Structure is documented below.
        """
        return pulumi.get(self, "service_perimeters")

    @service_perimeters.setter
    def service_perimeters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServicePerimetersServicePerimeterArgs']]]]):
        pulumi.set(self, "service_perimeters", value)


class ServicePerimeters(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePerimetersServicePerimeterArgs']]]]] = None,
                 __props__=None):
        """
        Replace all existing Service Perimeters in an Access Policy with the Service Perimeters provided. This is done atomically.
        This is a bulk edit of all Service Perimeters and may override existing Service Perimeters created by `accesscontextmanager.ServicePerimeter`,
        thus causing a permadiff if used alongside `accesscontextmanager.ServicePerimeter` on the same parent.

        To get more information about ServicePerimeters, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
        * How-to Guides
            * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)

        ## Example Usage
        ### Access Context Manager Service Perimeters Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        access_policy = gcp.accesscontextmanager.AccessPolicy("access-policy",
            parent="organizations/123456789",
            title="my policy")
        service_perimeter = gcp.accesscontextmanager.ServicePerimeters("service-perimeter",
            parent=access_policy.name.apply(lambda name: f"accessPolicies/{name}"),
            service_perimeters=[
                gcp.accesscontextmanager.ServicePerimetersServicePerimeterArgs(
                    name=access_policy.name.apply(lambda name: f"accessPolicies/{name}/servicePerimeters/"),
                    status=gcp.accesscontextmanager.ServicePerimetersServicePerimeterStatusArgs(
                        restricted_services=["storage.googleapis.com"],
                    ),
                    title="",
                ),
                gcp.accesscontextmanager.ServicePerimetersServicePerimeterArgs(
                    name=access_policy.name.apply(lambda name: f"accessPolicies/{name}/servicePerimeters/"),
                    status=gcp.accesscontextmanager.ServicePerimetersServicePerimeterStatusArgs(
                        restricted_services=["bigtable.googleapis.com"],
                    ),
                    title="",
                ),
            ])
        access_level = gcp.accesscontextmanager.AccessLevel("access-level",
            basic=gcp.accesscontextmanager.AccessLevelBasicArgs(
                conditions=[gcp.accesscontextmanager.AccessLevelBasicConditionArgs(
                    device_policy=gcp.accesscontextmanager.AccessLevelBasicConditionDevicePolicyArgs(
                        os_constraints=[gcp.accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArgs(
                            os_type="DESKTOP_CHROME_OS",
                        )],
                        require_screen_lock=False,
                    ),
                    regions=[
                        "CH",
                        "IT",
                        "US",
                    ],
                )],
            ),
            parent=access_policy.name.apply(lambda name: f"accessPolicies/{name}"),
            title="chromeos_no_lock")
        ```

        ## Import

        ServicePerimeters can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}/servicePerimeters
        ```

        ```sh
         $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent: The AccessPolicy this ServicePerimeter lives in.
               Format: accessPolicies/{policy_id}
               
               
               - - -
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePerimetersServicePerimeterArgs']]]] service_perimeters: The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServicePerimetersArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Replace all existing Service Perimeters in an Access Policy with the Service Perimeters provided. This is done atomically.
        This is a bulk edit of all Service Perimeters and may override existing Service Perimeters created by `accesscontextmanager.ServicePerimeter`,
        thus causing a permadiff if used alongside `accesscontextmanager.ServicePerimeter` on the same parent.

        To get more information about ServicePerimeters, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
        * How-to Guides
            * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)

        ## Example Usage
        ### Access Context Manager Service Perimeters Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        access_policy = gcp.accesscontextmanager.AccessPolicy("access-policy",
            parent="organizations/123456789",
            title="my policy")
        service_perimeter = gcp.accesscontextmanager.ServicePerimeters("service-perimeter",
            parent=access_policy.name.apply(lambda name: f"accessPolicies/{name}"),
            service_perimeters=[
                gcp.accesscontextmanager.ServicePerimetersServicePerimeterArgs(
                    name=access_policy.name.apply(lambda name: f"accessPolicies/{name}/servicePerimeters/"),
                    status=gcp.accesscontextmanager.ServicePerimetersServicePerimeterStatusArgs(
                        restricted_services=["storage.googleapis.com"],
                    ),
                    title="",
                ),
                gcp.accesscontextmanager.ServicePerimetersServicePerimeterArgs(
                    name=access_policy.name.apply(lambda name: f"accessPolicies/{name}/servicePerimeters/"),
                    status=gcp.accesscontextmanager.ServicePerimetersServicePerimeterStatusArgs(
                        restricted_services=["bigtable.googleapis.com"],
                    ),
                    title="",
                ),
            ])
        access_level = gcp.accesscontextmanager.AccessLevel("access-level",
            basic=gcp.accesscontextmanager.AccessLevelBasicArgs(
                conditions=[gcp.accesscontextmanager.AccessLevelBasicConditionArgs(
                    device_policy=gcp.accesscontextmanager.AccessLevelBasicConditionDevicePolicyArgs(
                        os_constraints=[gcp.accesscontextmanager.AccessLevelBasicConditionDevicePolicyOsConstraintArgs(
                            os_type="DESKTOP_CHROME_OS",
                        )],
                        require_screen_lock=False,
                    ),
                    regions=[
                        "CH",
                        "IT",
                        "US",
                    ],
                )],
            ),
            parent=access_policy.name.apply(lambda name: f"accessPolicies/{name}"),
            title="chromeos_no_lock")
        ```

        ## Import

        ServicePerimeters can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}/servicePerimeters
        ```

        ```sh
         $ pulumi import gcp:accesscontextmanager/servicePerimeters:ServicePerimeters default {{parent}}
        ```

        :param str resource_name: The name of the resource.
        :param ServicePerimetersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServicePerimetersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ServicePerimetersArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePerimetersServicePerimeterArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServicePerimetersArgs.__new__(ServicePerimetersArgs)

            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__["parent"] = parent
            __props__.__dict__["service_perimeters"] = service_perimeters
        super(ServicePerimeters, __self__).__init__(
            'gcp:accesscontextmanager/servicePerimeters:ServicePerimeters',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            parent: Optional[pulumi.Input[str]] = None,
            service_perimeters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePerimetersServicePerimeterArgs']]]]] = None) -> 'ServicePerimeters':
        """
        Get an existing ServicePerimeters resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] parent: The AccessPolicy this ServicePerimeter lives in.
               Format: accessPolicies/{policy_id}
               
               
               - - -
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServicePerimetersServicePerimeterArgs']]]] service_perimeters: The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServicePerimetersState.__new__(_ServicePerimetersState)

        __props__.__dict__["parent"] = parent
        __props__.__dict__["service_perimeters"] = service_perimeters
        return ServicePerimeters(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The AccessPolicy this ServicePerimeter lives in.
        Format: accessPolicies/{policy_id}


        - - -
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="servicePerimeters")
    def service_perimeters(self) -> pulumi.Output[Optional[Sequence['outputs.ServicePerimetersServicePerimeter']]]:
        """
        The desired Service Perimeters that should replace all existing Service Perimeters in the Access Policy.
        Structure is documented below.
        """
        return pulumi.get(self, "service_perimeters")

