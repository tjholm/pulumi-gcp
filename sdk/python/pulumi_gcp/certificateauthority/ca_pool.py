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

__all__ = ['CaPoolArgs', 'CaPool']

@pulumi.input_type
class CaPoolArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 tier: pulumi.Input[str],
                 issuance_policy: Optional[pulumi.Input['CaPoolIssuancePolicyArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 publishing_options: Optional[pulumi.Input['CaPoolPublishingOptionsArgs']] = None):
        """
        The set of arguments for constructing a CaPool resource.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] tier: The Tier of this CaPool.
               Possible values are `ENTERPRISE` and `DEVOPS`.
        :param pulumi.Input['CaPoolIssuancePolicyArgs'] issuance_policy: The IssuancePolicy to control how Certificates will be issued from this CaPool.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata.
               An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
               "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The name for this CaPool.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['CaPoolPublishingOptionsArgs'] publishing_options: The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
               Structure is documented below.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "tier", tier)
        if issuance_policy is not None:
            pulumi.set(__self__, "issuance_policy", issuance_policy)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if publishing_options is not None:
            pulumi.set(__self__, "publishing_options", publishing_options)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Input[str]:
        """
        The Tier of this CaPool.
        Possible values are `ENTERPRISE` and `DEVOPS`.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: pulumi.Input[str]):
        pulumi.set(self, "tier", value)

    @property
    @pulumi.getter(name="issuancePolicy")
    def issuance_policy(self) -> Optional[pulumi.Input['CaPoolIssuancePolicyArgs']]:
        """
        The IssuancePolicy to control how Certificates will be issued from this CaPool.
        Structure is documented below.
        """
        return pulumi.get(self, "issuance_policy")

    @issuance_policy.setter
    def issuance_policy(self, value: Optional[pulumi.Input['CaPoolIssuancePolicyArgs']]):
        pulumi.set(self, "issuance_policy", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels with user-defined metadata.
        An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
        "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for this CaPool.
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
    @pulumi.getter(name="publishingOptions")
    def publishing_options(self) -> Optional[pulumi.Input['CaPoolPublishingOptionsArgs']]:
        """
        The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        Structure is documented below.
        """
        return pulumi.get(self, "publishing_options")

    @publishing_options.setter
    def publishing_options(self, value: Optional[pulumi.Input['CaPoolPublishingOptionsArgs']]):
        pulumi.set(self, "publishing_options", value)


@pulumi.input_type
class _CaPoolState:
    def __init__(__self__, *,
                 issuance_policy: Optional[pulumi.Input['CaPoolIssuancePolicyArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 publishing_options: Optional[pulumi.Input['CaPoolPublishingOptionsArgs']] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CaPool resources.
        :param pulumi.Input['CaPoolIssuancePolicyArgs'] issuance_policy: The IssuancePolicy to control how Certificates will be issued from this CaPool.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata.
               An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
               "1.3kg", "count": "3" }.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] name: The name for this CaPool.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['CaPoolPublishingOptionsArgs'] publishing_options: The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
               Structure is documented below.
        :param pulumi.Input[str] tier: The Tier of this CaPool.
               Possible values are `ENTERPRISE` and `DEVOPS`.
        """
        if issuance_policy is not None:
            pulumi.set(__self__, "issuance_policy", issuance_policy)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if publishing_options is not None:
            pulumi.set(__self__, "publishing_options", publishing_options)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter(name="issuancePolicy")
    def issuance_policy(self) -> Optional[pulumi.Input['CaPoolIssuancePolicyArgs']]:
        """
        The IssuancePolicy to control how Certificates will be issued from this CaPool.
        Structure is documented below.
        """
        return pulumi.get(self, "issuance_policy")

    @issuance_policy.setter
    def issuance_policy(self, value: Optional[pulumi.Input['CaPoolIssuancePolicyArgs']]):
        pulumi.set(self, "issuance_policy", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels with user-defined metadata.
        An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
        "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for this CaPool.
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
    @pulumi.getter(name="publishingOptions")
    def publishing_options(self) -> Optional[pulumi.Input['CaPoolPublishingOptionsArgs']]:
        """
        The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        Structure is documented below.
        """
        return pulumi.get(self, "publishing_options")

    @publishing_options.setter
    def publishing_options(self, value: Optional[pulumi.Input['CaPoolPublishingOptionsArgs']]):
        pulumi.set(self, "publishing_options", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The Tier of this CaPool.
        Possible values are `ENTERPRISE` and `DEVOPS`.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


class CaPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 issuance_policy: Optional[pulumi.Input[pulumi.InputType['CaPoolIssuancePolicyArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 publishing_options: Optional[pulumi.Input[pulumi.InputType['CaPoolPublishingOptionsArgs']]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A CaPool represents a group of CertificateAuthorities that form a trust anchor. A CaPool can be used to manage
        issuance policies for one or more CertificateAuthority resources and to rotate CA certificates in and out of the
        trust anchor.

        ## Example Usage
        ### Privateca Capool Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.certificateauthority.CaPool("default",
            labels={
                "foo": "bar",
            },
            location="us-central1",
            publishing_options=gcp.certificateauthority.CaPoolPublishingOptionsArgs(
                publish_ca_cert=True,
                publish_crl=True,
            ),
            tier="ENTERPRISE")
        ```

        ## Import

        CaPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
        ```

        ```sh
         $ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CaPoolIssuancePolicyArgs']] issuance_policy: The IssuancePolicy to control how Certificates will be issued from this CaPool.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata.
               An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
               "1.3kg", "count": "3" }.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] name: The name for this CaPool.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['CaPoolPublishingOptionsArgs']] publishing_options: The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
               Structure is documented below.
        :param pulumi.Input[str] tier: The Tier of this CaPool.
               Possible values are `ENTERPRISE` and `DEVOPS`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CaPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A CaPool represents a group of CertificateAuthorities that form a trust anchor. A CaPool can be used to manage
        issuance policies for one or more CertificateAuthority resources and to rotate CA certificates in and out of the
        trust anchor.

        ## Example Usage
        ### Privateca Capool Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.certificateauthority.CaPool("default",
            labels={
                "foo": "bar",
            },
            location="us-central1",
            publishing_options=gcp.certificateauthority.CaPoolPublishingOptionsArgs(
                publish_ca_cert=True,
                publish_crl=True,
            ),
            tier="ENTERPRISE")
        ```

        ## Import

        CaPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
        ```

        ```sh
         $ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param CaPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CaPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 issuance_policy: Optional[pulumi.Input[pulumi.InputType['CaPoolIssuancePolicyArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 publishing_options: Optional[pulumi.Input[pulumi.InputType['CaPoolPublishingOptionsArgs']]] = None,
                 tier: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CaPoolArgs.__new__(CaPoolArgs)

            __props__.__dict__["issuance_policy"] = issuance_policy
            __props__.__dict__["labels"] = labels
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["publishing_options"] = publishing_options
            if tier is None and not opts.urn:
                raise TypeError("Missing required property 'tier'")
            __props__.__dict__["tier"] = tier
        super(CaPool, __self__).__init__(
            'gcp:certificateauthority/caPool:CaPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            issuance_policy: Optional[pulumi.Input[pulumi.InputType['CaPoolIssuancePolicyArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            publishing_options: Optional[pulumi.Input[pulumi.InputType['CaPoolPublishingOptionsArgs']]] = None,
            tier: Optional[pulumi.Input[str]] = None) -> 'CaPool':
        """
        Get an existing CaPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CaPoolIssuancePolicyArgs']] issuance_policy: The IssuancePolicy to control how Certificates will be issued from this CaPool.
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata.
               An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
               "1.3kg", "count": "3" }.
        :param pulumi.Input[str] location: String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] name: The name for this CaPool.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['CaPoolPublishingOptionsArgs']] publishing_options: The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
               Structure is documented below.
        :param pulumi.Input[str] tier: The Tier of this CaPool.
               Possible values are `ENTERPRISE` and `DEVOPS`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CaPoolState.__new__(_CaPoolState)

        __props__.__dict__["issuance_policy"] = issuance_policy
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["publishing_options"] = publishing_options
        __props__.__dict__["tier"] = tier
        return CaPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="issuancePolicy")
    def issuance_policy(self) -> pulumi.Output[Optional['outputs.CaPoolIssuancePolicy']]:
        """
        The IssuancePolicy to control how Certificates will be issued from this CaPool.
        Structure is documented below.
        """
        return pulumi.get(self, "issuance_policy")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels with user-defined metadata.
        An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
        "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for this CaPool.
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
    @pulumi.getter(name="publishingOptions")
    def publishing_options(self) -> pulumi.Output[Optional['outputs.CaPoolPublishingOptions']]:
        """
        The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        Structure is documented below.
        """
        return pulumi.get(self, "publishing_options")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[str]:
        """
        The Tier of this CaPool.
        Possible values are `ENTERPRISE` and `DEVOPS`.
        """
        return pulumi.get(self, "tier")

