# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 locations: pulumi.Input[Sequence[pulumi.Input[str]]],
                 reserved_ip_range: pulumi.Input[str],
                 admin: Optional[pulumi.Input[str]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[str] domain_name: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
               
               
               - - -
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
               e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] reserved_ip_range: The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
               Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        :param pulumi.Input[str] admin: The name of delegated administrator account used to perform Active Directory operations.
               If not specified, setupadmin will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
               If CIDR subnets overlap between networks, domain creation will fail.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels that can contain user-provided metadata
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "locations", locations)
        pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)
        if admin is not None:
            pulumi.set(__self__, "admin", admin)
        if authorized_networks is not None:
            pulumi.set(__self__, "authorized_networks", authorized_networks)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
        https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.


        - - -
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
        e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> pulumi.Input[str]:
        """
        The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
        Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        """
        return pulumi.get(self, "reserved_ip_range")

    @reserved_ip_range.setter
    def reserved_ip_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "reserved_ip_range", value)

    @property
    @pulumi.getter
    def admin(self) -> Optional[pulumi.Input[str]]:
        """
        The name of delegated administrator account used to perform Active Directory operations.
        If not specified, setupadmin will be used.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
        If CIDR subnets overlap between networks, domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @authorized_networks.setter
    def authorized_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_networks", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels that can contain user-provided metadata
        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

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
class _DomainState:
    def __init__(__self__, *,
                 admin: Optional[pulumi.Input[str]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 fqdn: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Domain resources.
        :param pulumi.Input[str] admin: The name of delegated administrator account used to perform Active Directory operations.
               If not specified, setupadmin will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
               If CIDR subnets overlap between networks, domain creation will fail.
        :param pulumi.Input[str] domain_name: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
               
               
               - - -
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] fqdn: The fully-qualified domain name of the exposed domain used by clients to connect to the service.
               Similar to what would be chosen for an Active Directory set up on an internal network.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels that can contain user-provided metadata
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
               e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] name: The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] reserved_ip_range: The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
               Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        """
        if admin is not None:
            pulumi.set(__self__, "admin", admin)
        if authorized_networks is not None:
            pulumi.set(__self__, "authorized_networks", authorized_networks)
        if domain_name is not None:
            pulumi.set(__self__, "domain_name", domain_name)
        if effective_labels is not None:
            pulumi.set(__self__, "effective_labels", effective_labels)
        if fqdn is not None:
            pulumi.set(__self__, "fqdn", fqdn)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if locations is not None:
            pulumi.set(__self__, "locations", locations)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if pulumi_labels is not None:
            pulumi.set(__self__, "pulumi_labels", pulumi_labels)
        if reserved_ip_range is not None:
            pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)

    @property
    @pulumi.getter
    def admin(self) -> Optional[pulumi.Input[str]]:
        """
        The name of delegated administrator account used to perform Active Directory operations.
        If not specified, setupadmin will be used.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
        If CIDR subnets overlap between networks, domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @authorized_networks.setter
    def authorized_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_networks", value)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
        https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.


        - - -
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain_name", value)

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
    def fqdn(self) -> Optional[pulumi.Input[str]]:
        """
        The fully-qualified domain name of the exposed domain used by clients to connect to the service.
        Similar to what would be chosen for an Active Directory set up on an internal network.
        """
        return pulumi.get(self, "fqdn")

    @fqdn.setter
    def fqdn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fqdn", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels that can contain user-provided metadata
        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def locations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
        e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
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
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> Optional[pulumi.Input[str]]:
        """
        The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
        Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        """
        return pulumi.get(self, "reserved_ip_range")

    @reserved_ip_range.setter
    def reserved_ip_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reserved_ip_range", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[str]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a Microsoft AD domain

        To get more information about Domain, see:

        * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains)
        * How-to Guides
            * [Managed Microsoft Active Directory Quickstart](https://cloud.google.com/managed-microsoft-ad/docs/quickstarts)

        ## Example Usage
        ### Active Directory Domain Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ad_domain = gcp.activedirectory.Domain("ad-domain",
            domain_name="tfgen.org.com",
            locations=["us-central1"],
            reserved_ip_range="192.168.255.0/24")
        ```

        ## Import

        Domain can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Domain can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:activedirectory/domain:Domain default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin: The name of delegated administrator account used to perform Active Directory operations.
               If not specified, setupadmin will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
               If CIDR subnets overlap between networks, domain creation will fail.
        :param pulumi.Input[str] domain_name: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
               
               
               - - -
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels that can contain user-provided metadata
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
               e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] reserved_ip_range: The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
               Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Microsoft AD domain

        To get more information about Domain, see:

        * [API documentation](https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains)
        * How-to Guides
            * [Managed Microsoft Active Directory Quickstart](https://cloud.google.com/managed-microsoft-ad/docs/quickstarts)

        ## Example Usage
        ### Active Directory Domain Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        ad_domain = gcp.activedirectory.Domain("ad-domain",
            domain_name="tfgen.org.com",
            locations=["us-central1"],
            reserved_ip_range="192.168.255.0/24")
        ```

        ## Import

        Domain can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Domain can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:activedirectory/domain:Domain default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[str]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["admin"] = admin
            __props__.__dict__["authorized_networks"] = authorized_networks
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["labels"] = labels
            if locations is None and not opts.urn:
                raise TypeError("Missing required property 'locations'")
            __props__.__dict__["locations"] = locations
            __props__.__dict__["project"] = project
            if reserved_ip_range is None and not opts.urn:
                raise TypeError("Missing required property 'reserved_ip_range'")
            __props__.__dict__["reserved_ip_range"] = reserved_ip_range
            __props__.__dict__["effective_labels"] = None
            __props__.__dict__["fqdn"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["pulumi_labels"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["effectiveLabels", "pulumiLabels"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Domain, __self__).__init__(
            'gcp:activedirectory/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin: Optional[pulumi.Input[str]] = None,
            authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            fqdn: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            reserved_ip_range: Optional[pulumi.Input[str]] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin: The name of delegated administrator account used to perform Active Directory operations.
               If not specified, setupadmin will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
               If CIDR subnets overlap between networks, domain creation will fail.
        :param pulumi.Input[str] domain_name: The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
               https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
               
               
               - - -
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] fqdn: The fully-qualified domain name of the exposed domain used by clients to connect to the service.
               Similar to what would be chosen for an Active Directory set up on an internal network.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels that can contain user-provided metadata
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
               e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] name: The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] reserved_ip_range: The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
               Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainState.__new__(_DomainState)

        __props__.__dict__["admin"] = admin
        __props__.__dict__["authorized_networks"] = authorized_networks
        __props__.__dict__["domain_name"] = domain_name
        __props__.__dict__["effective_labels"] = effective_labels
        __props__.__dict__["fqdn"] = fqdn
        __props__.__dict__["labels"] = labels
        __props__.__dict__["locations"] = locations
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["pulumi_labels"] = pulumi_labels
        __props__.__dict__["reserved_ip_range"] = reserved_ip_range
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admin(self) -> pulumi.Output[Optional[str]]:
        """
        The name of delegated administrator account used to perform Active Directory operations.
        If not specified, setupadmin will be used.
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
        If CIDR subnets overlap between networks, domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
        https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.


        - - -
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The fully-qualified domain name of the exposed domain used by clients to connect to the service.
        Similar to what would be chosen for an Active Directory set up on an internal network.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource labels that can contain user-provided metadata
        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence[str]]:
        """
        Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
        e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the domain using the format: `projects/{project}/locations/global/domains/{domainName}`.
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
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> pulumi.Output[str]:
        """
        The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
        Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
        """
        return pulumi.get(self, "reserved_ip_range")

