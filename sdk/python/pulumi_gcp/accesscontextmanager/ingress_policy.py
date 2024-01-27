# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IngressPolicyArgs', 'IngressPolicy']

@pulumi.input_type
class IngressPolicyArgs:
    def __init__(__self__, *,
                 ingress_policy_name: pulumi.Input[str],
                 resource: pulumi.Input[str]):
        """
        The set of arguments for constructing a IngressPolicy resource.
        :param pulumi.Input[str] ingress_policy_name: The name of the Service Perimeter to add this resource to.
               
               
               - - -
        :param pulumi.Input[str] resource: A GCP resource that is inside of the service perimeter.
        """
        pulumi.set(__self__, "ingress_policy_name", ingress_policy_name)
        pulumi.set(__self__, "resource", resource)

    @property
    @pulumi.getter(name="ingressPolicyName")
    def ingress_policy_name(self) -> pulumi.Input[str]:
        """
        The name of the Service Perimeter to add this resource to.


        - - -
        """
        return pulumi.get(self, "ingress_policy_name")

    @ingress_policy_name.setter
    def ingress_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ingress_policy_name", value)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[str]:
        """
        A GCP resource that is inside of the service perimeter.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource", value)


@pulumi.input_type
class _IngressPolicyState:
    def __init__(__self__, *,
                 ingress_policy_name: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IngressPolicy resources.
        :param pulumi.Input[str] ingress_policy_name: The name of the Service Perimeter to add this resource to.
               
               
               - - -
        :param pulumi.Input[str] resource: A GCP resource that is inside of the service perimeter.
        """
        if ingress_policy_name is not None:
            pulumi.set(__self__, "ingress_policy_name", ingress_policy_name)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)

    @property
    @pulumi.getter(name="ingressPolicyName")
    def ingress_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Service Perimeter to add this resource to.


        - - -
        """
        return pulumi.get(self, "ingress_policy_name")

    @ingress_policy_name.setter
    def ingress_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ingress_policy_name", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[str]]:
        """
        A GCP resource that is inside of the service perimeter.
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource", value)


class IngressPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ingress_policy_name: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource has been deprecated, please refer to ServicePerimeterIngressPolicy.

        To get more information about IngressPolicy, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters#ingresspolicy)

        ## Import

        IngressPolicy can be imported using any of these accepted formats* `{{ingress_policy_name}}/{{resource}}` When using the `pulumi import` command, IngressPolicy can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:accesscontextmanager/ingressPolicy:IngressPolicy default {{ingress_policy_name}}/{{resource}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ingress_policy_name: The name of the Service Perimeter to add this resource to.
               
               
               - - -
        :param pulumi.Input[str] resource: A GCP resource that is inside of the service perimeter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IngressPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource has been deprecated, please refer to ServicePerimeterIngressPolicy.

        To get more information about IngressPolicy, see:

        * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters#ingresspolicy)

        ## Import

        IngressPolicy can be imported using any of these accepted formats* `{{ingress_policy_name}}/{{resource}}` When using the `pulumi import` command, IngressPolicy can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:accesscontextmanager/ingressPolicy:IngressPolicy default {{ingress_policy_name}}/{{resource}}
        ```

        :param str resource_name: The name of the resource.
        :param IngressPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IngressPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ingress_policy_name: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IngressPolicyArgs.__new__(IngressPolicyArgs)

            if ingress_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'ingress_policy_name'")
            __props__.__dict__["ingress_policy_name"] = ingress_policy_name
            if resource is None and not opts.urn:
                raise TypeError("Missing required property 'resource'")
            __props__.__dict__["resource"] = resource
        super(IngressPolicy, __self__).__init__(
            'gcp:accesscontextmanager/ingressPolicy:IngressPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ingress_policy_name: Optional[pulumi.Input[str]] = None,
            resource: Optional[pulumi.Input[str]] = None) -> 'IngressPolicy':
        """
        Get an existing IngressPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ingress_policy_name: The name of the Service Perimeter to add this resource to.
               
               
               - - -
        :param pulumi.Input[str] resource: A GCP resource that is inside of the service perimeter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IngressPolicyState.__new__(_IngressPolicyState)

        __props__.__dict__["ingress_policy_name"] = ingress_policy_name
        __props__.__dict__["resource"] = resource
        return IngressPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ingressPolicyName")
    def ingress_policy_name(self) -> pulumi.Output[str]:
        """
        The name of the Service Perimeter to add this resource to.


        - - -
        """
        return pulumi.get(self, "ingress_policy_name")

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Output[str]:
        """
        A GCP resource that is inside of the service perimeter.
        """
        return pulumi.get(self, "resource")

