# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EndpointAttachmentArgs', 'EndpointAttachment']

@pulumi.input_type
class EndpointAttachmentArgs:
    def __init__(__self__, *,
                 endpoint_attachment_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 org_id: pulumi.Input[str],
                 service_attachment: pulumi.Input[str]):
        """
        The set of arguments for constructing a EndpointAttachment resource.
        :param pulumi.Input[str] endpoint_attachment_id: ID of the endpoint attachment.
               
               
               - - -
        :param pulumi.Input[str] location: Location of the endpoint attachment.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] service_attachment: Format: projects/*/regions/*/serviceAttachments/*
        """
        pulumi.set(__self__, "endpoint_attachment_id", endpoint_attachment_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "service_attachment", service_attachment)

    @property
    @pulumi.getter(name="endpointAttachmentId")
    def endpoint_attachment_id(self) -> pulumi.Input[str]:
        """
        ID of the endpoint attachment.


        - - -
        """
        return pulumi.get(self, "endpoint_attachment_id")

    @endpoint_attachment_id.setter
    def endpoint_attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_attachment_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        Location of the endpoint attachment.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        The Apigee Organization associated with the Apigee instance,
        in the format `organizations/{{org_name}}`.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> pulumi.Input[str]:
        """
        Format: projects/*/regions/*/serviceAttachments/*
        """
        return pulumi.get(self, "service_attachment")

    @service_attachment.setter
    def service_attachment(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_attachment", value)


@pulumi.input_type
class _EndpointAttachmentState:
    def __init__(__self__, *,
                 connection_state: Optional[pulumi.Input[str]] = None,
                 endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointAttachment resources.
        :param pulumi.Input[str] connection_state: State of the endpoint attachment connection to the service attachment.
        :param pulumi.Input[str] endpoint_attachment_id: ID of the endpoint attachment.
               
               
               - - -
        :param pulumi.Input[str] host: Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
        :param pulumi.Input[str] location: Location of the endpoint attachment.
        :param pulumi.Input[str] name: Name of the Endpoint Attachment in the following format:
               organizations/{organization}/endpointAttachments/{endpointAttachment}.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] service_attachment: Format: projects/*/regions/*/serviceAttachments/*
        """
        if connection_state is not None:
            pulumi.set(__self__, "connection_state", connection_state)
        if endpoint_attachment_id is not None:
            pulumi.set(__self__, "endpoint_attachment_id", endpoint_attachment_id)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if service_attachment is not None:
            pulumi.set(__self__, "service_attachment", service_attachment)

    @property
    @pulumi.getter(name="connectionState")
    def connection_state(self) -> Optional[pulumi.Input[str]]:
        """
        State of the endpoint attachment connection to the service attachment.
        """
        return pulumi.get(self, "connection_state")

    @connection_state.setter
    def connection_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_state", value)

    @property
    @pulumi.getter(name="endpointAttachmentId")
    def endpoint_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the endpoint attachment.


        - - -
        """
        return pulumi.get(self, "endpoint_attachment_id")

    @endpoint_attachment_id.setter
    def endpoint_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_attachment_id", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the endpoint attachment.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Endpoint Attachment in the following format:
        organizations/{organization}/endpointAttachments/{endpointAttachment}.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Apigee Organization associated with the Apigee instance,
        in the format `organizations/{{org_name}}`.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> Optional[pulumi.Input[str]]:
        """
        Format: projects/*/regions/*/serviceAttachments/*
        """
        return pulumi.get(self, "service_attachment")

    @service_attachment.setter
    def service_attachment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_attachment", value)


class EndpointAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Apigee Endpoint Attachment.

        To get more information about EndpointAttachment, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.endpointAttachments/create)
        * How-to Guides
            * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

        ## Example Usage
        ### Apigee Endpoint Attachment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        current = gcp.organizations.get_client_config()
        apigee_network = gcp.compute.Network("apigeeNetwork")
        apigee_range = gcp.compute.GlobalAddress("apigeeRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=apigee_network.id)
        apigee_vpc_connection = gcp.servicenetworking.Connection("apigeeVpcConnection",
            network=apigee_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[apigee_range.name])
        apigee_org = gcp.apigee.Organization("apigeeOrg",
            analytics_region="us-central1",
            project_id=current.project,
            authorized_network=apigee_network.id,
            opts=pulumi.ResourceOptions(depends_on=[apigee_vpc_connection]))
        apigee_endpoint_attachment = gcp.apigee.EndpointAttachment("apigeeEndpointAttachment",
            org_id=apigee_org.id,
            endpoint_attachment_id="test1",
            location="{google_compute_service_attachment location}",
            service_attachment="{google_compute_service_attachment id}")
        ```

        ## Import

        EndpointAttachment can be imported using any of these accepted formats* `{{org_id}}/endpointAttachments/{{endpoint_attachment_id}}` * `{{org_id}}/{{endpoint_attachment_id}}` When using the `pulumi import` command, EndpointAttachment can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:apigee/endpointAttachment:EndpointAttachment default {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
        ```

        ```sh
         $ pulumi import gcp:apigee/endpointAttachment:EndpointAttachment default {{org_id}}/{{endpoint_attachment_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_attachment_id: ID of the endpoint attachment.
               
               
               - - -
        :param pulumi.Input[str] location: Location of the endpoint attachment.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] service_attachment: Format: projects/*/regions/*/serviceAttachments/*
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Apigee Endpoint Attachment.

        To get more information about EndpointAttachment, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.endpointAttachments/create)
        * How-to Guides
            * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)

        ## Example Usage
        ### Apigee Endpoint Attachment Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        current = gcp.organizations.get_client_config()
        apigee_network = gcp.compute.Network("apigeeNetwork")
        apigee_range = gcp.compute.GlobalAddress("apigeeRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=apigee_network.id)
        apigee_vpc_connection = gcp.servicenetworking.Connection("apigeeVpcConnection",
            network=apigee_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[apigee_range.name])
        apigee_org = gcp.apigee.Organization("apigeeOrg",
            analytics_region="us-central1",
            project_id=current.project,
            authorized_network=apigee_network.id,
            opts=pulumi.ResourceOptions(depends_on=[apigee_vpc_connection]))
        apigee_endpoint_attachment = gcp.apigee.EndpointAttachment("apigeeEndpointAttachment",
            org_id=apigee_org.id,
            endpoint_attachment_id="test1",
            location="{google_compute_service_attachment location}",
            service_attachment="{google_compute_service_attachment id}")
        ```

        ## Import

        EndpointAttachment can be imported using any of these accepted formats* `{{org_id}}/endpointAttachments/{{endpoint_attachment_id}}` * `{{org_id}}/{{endpoint_attachment_id}}` When using the `pulumi import` command, EndpointAttachment can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:apigee/endpointAttachment:EndpointAttachment default {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
        ```

        ```sh
         $ pulumi import gcp:apigee/endpointAttachment:EndpointAttachment default {{org_id}}/{{endpoint_attachment_id}}
        ```

        :param str resource_name: The name of the resource.
        :param EndpointAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 service_attachment: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointAttachmentArgs.__new__(EndpointAttachmentArgs)

            if endpoint_attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_attachment_id'")
            __props__.__dict__["endpoint_attachment_id"] = endpoint_attachment_id
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            if service_attachment is None and not opts.urn:
                raise TypeError("Missing required property 'service_attachment'")
            __props__.__dict__["service_attachment"] = service_attachment
            __props__.__dict__["connection_state"] = None
            __props__.__dict__["host"] = None
            __props__.__dict__["name"] = None
        super(EndpointAttachment, __self__).__init__(
            'gcp:apigee/endpointAttachment:EndpointAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_state: Optional[pulumi.Input[str]] = None,
            endpoint_attachment_id: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            service_attachment: Optional[pulumi.Input[str]] = None) -> 'EndpointAttachment':
        """
        Get an existing EndpointAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_state: State of the endpoint attachment connection to the service attachment.
        :param pulumi.Input[str] endpoint_attachment_id: ID of the endpoint attachment.
               
               
               - - -
        :param pulumi.Input[str] host: Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
        :param pulumi.Input[str] location: Location of the endpoint attachment.
        :param pulumi.Input[str] name: Name of the Endpoint Attachment in the following format:
               organizations/{organization}/endpointAttachments/{endpointAttachment}.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] service_attachment: Format: projects/*/regions/*/serviceAttachments/*
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointAttachmentState.__new__(_EndpointAttachmentState)

        __props__.__dict__["connection_state"] = connection_state
        __props__.__dict__["endpoint_attachment_id"] = endpoint_attachment_id
        __props__.__dict__["host"] = host
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["service_attachment"] = service_attachment
        return EndpointAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionState")
    def connection_state(self) -> pulumi.Output[str]:
        """
        State of the endpoint attachment connection to the service attachment.
        """
        return pulumi.get(self, "connection_state")

    @property
    @pulumi.getter(name="endpointAttachmentId")
    def endpoint_attachment_id(self) -> pulumi.Output[str]:
        """
        ID of the endpoint attachment.


        - - -
        """
        return pulumi.get(self, "endpoint_attachment_id")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Location of the endpoint attachment.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Endpoint Attachment in the following format:
        organizations/{organization}/endpointAttachments/{endpointAttachment}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The Apigee Organization associated with the Apigee instance,
        in the format `organizations/{{org_name}}`.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="serviceAttachment")
    def service_attachment(self) -> pulumi.Output[str]:
        """
        Format: projects/*/regions/*/serviceAttachments/*
        """
        return pulumi.get(self, "service_attachment")

