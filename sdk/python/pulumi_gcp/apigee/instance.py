# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceArgs', 'Instance']

@pulumi.input_type
class InstanceArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 org_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Instance resource.
        :param pulumi.Input[str] location: Compute Engine location where the instance resides. For trial organization
               subscriptions, the location must be a Compute Engine zone. For paid organization
               subscriptions, it should correspond to a Compute Engine region.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] description: Description of the instance.
        :param pulumi.Input[str] disk_encryption_key_name: Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
               Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        :param pulumi.Input[str] display_name: Display name of the instance.
        :param pulumi.Input[str] name: Resource ID of the instance.
        :param pulumi.Input[str] peering_cidr_range: The size of the CIDR block range that will be reserved by the instance.
               Possible values are `SLASH_16` and `SLASH_20`.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "org_id", org_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_encryption_key_name is not None:
            pulumi.set(__self__, "disk_encryption_key_name", disk_encryption_key_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if peering_cidr_range is not None:
            pulumi.set(__self__, "peering_cidr_range", peering_cidr_range)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        Compute Engine location where the instance resides. For trial organization
        subscriptions, the location must be a Compute Engine zone. For paid organization
        subscriptions, it should correspond to a Compute Engine region.
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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskEncryptionKeyName")
    def disk_encryption_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
        Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        """
        return pulumi.get(self, "disk_encryption_key_name")

    @disk_encryption_key_name.setter
    def disk_encryption_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_encryption_key_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the instance.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource ID of the instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="peeringCidrRange")
    def peering_cidr_range(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the CIDR block range that will be reserved by the instance.
        Possible values are `SLASH_16` and `SLASH_20`.
        """
        return pulumi.get(self, "peering_cidr_range")

    @peering_cidr_range.setter
    def peering_cidr_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_cidr_range", value)


@pulumi.input_type
class _InstanceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Instance resources.
        :param pulumi.Input[str] description: Description of the instance.
        :param pulumi.Input[str] disk_encryption_key_name: Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
               Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        :param pulumi.Input[str] display_name: Display name of the instance.
        :param pulumi.Input[str] host: Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
        :param pulumi.Input[str] location: Compute Engine location where the instance resides. For trial organization
               subscriptions, the location must be a Compute Engine zone. For paid organization
               subscriptions, it should correspond to a Compute Engine region.
        :param pulumi.Input[str] name: Resource ID of the instance.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] peering_cidr_range: The size of the CIDR block range that will be reserved by the instance.
               Possible values are `SLASH_16` and `SLASH_20`.
        :param pulumi.Input[str] port: Output only. Port number of the exposed Apigee endpoint.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disk_encryption_key_name is not None:
            pulumi.set(__self__, "disk_encryption_key_name", disk_encryption_key_name)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if host is not None:
            pulumi.set(__self__, "host", host)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if peering_cidr_range is not None:
            pulumi.set(__self__, "peering_cidr_range", peering_cidr_range)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="diskEncryptionKeyName")
    def disk_encryption_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
        Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        """
        return pulumi.get(self, "disk_encryption_key_name")

    @disk_encryption_key_name.setter
    def disk_encryption_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "disk_encryption_key_name", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the instance.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Compute Engine location where the instance resides. For trial organization
        subscriptions, the location must be a Compute Engine zone. For paid organization
        subscriptions, it should correspond to a Compute Engine region.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource ID of the instance.
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
    @pulumi.getter(name="peeringCidrRange")
    def peering_cidr_range(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the CIDR block range that will be reserved by the instance.
        Possible values are `SLASH_16` and `SLASH_20`.
        """
        return pulumi.get(self, "peering_cidr_range")

    @peering_cidr_range.setter
    def peering_cidr_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peering_cidr_range", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Port number of the exposed Apigee endpoint.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)


class Instance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An `Instance` is the runtime dataplane in Apigee.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances/create)
        * How-to Guides
            * [Creating a runtime instance](https://cloud.google.com/apigee/docs/api-platform/get-started/create-instance)

        ## Example Usage

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance.
        :param pulumi.Input[str] disk_encryption_key_name: Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
               Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        :param pulumi.Input[str] display_name: Display name of the instance.
        :param pulumi.Input[str] location: Compute Engine location where the instance resides. For trial organization
               subscriptions, the location must be a Compute Engine zone. For paid organization
               subscriptions, it should correspond to a Compute Engine region.
        :param pulumi.Input[str] name: Resource ID of the instance.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] peering_cidr_range: The size of the CIDR block range that will be reserved by the instance.
               Possible values are `SLASH_16` and `SLASH_20`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An `Instance` is the runtime dataplane in Apigee.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.instances/create)
        * How-to Guides
            * [Creating a runtime instance](https://cloud.google.com/apigee/docs/api-platform/get-started/create-instance)

        ## Example Usage

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/instance:Instance default {{org_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param InstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 peering_cidr_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceArgs.__new__(InstanceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disk_encryption_key_name"] = disk_encryption_key_name
            __props__.__dict__["display_name"] = display_name
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["peering_cidr_range"] = peering_cidr_range
            __props__.__dict__["host"] = None
            __props__.__dict__["port"] = None
        super(Instance, __self__).__init__(
            'gcp:apigee/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disk_encryption_key_name: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            peering_cidr_range: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[str]] = None) -> 'Instance':
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the instance.
        :param pulumi.Input[str] disk_encryption_key_name: Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
               Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        :param pulumi.Input[str] display_name: Display name of the instance.
        :param pulumi.Input[str] host: Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
        :param pulumi.Input[str] location: Compute Engine location where the instance resides. For trial organization
               subscriptions, the location must be a Compute Engine zone. For paid organization
               subscriptions, it should correspond to a Compute Engine region.
        :param pulumi.Input[str] name: Resource ID of the instance.
        :param pulumi.Input[str] org_id: The Apigee Organization associated with the Apigee instance,
               in the format `organizations/{{org_name}}`.
        :param pulumi.Input[str] peering_cidr_range: The size of the CIDR block range that will be reserved by the instance.
               Possible values are `SLASH_16` and `SLASH_20`.
        :param pulumi.Input[str] port: Output only. Port number of the exposed Apigee endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceState.__new__(_InstanceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disk_encryption_key_name"] = disk_encryption_key_name
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["host"] = host
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["peering_cidr_range"] = peering_cidr_range
        __props__.__dict__["port"] = port
        return Instance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="diskEncryptionKeyName")
    def disk_encryption_key_name(self) -> pulumi.Output[Optional[str]]:
        """
        Customer Managed Encryption Key (CMEK) used for disk and volume encryption. Required for Apigee paid subscriptions only.
        Use the following format: `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`
        """
        return pulumi.get(self, "disk_encryption_key_name")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Display name of the instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Output only. Hostname or IP address of the exposed Apigee endpoint used by clients to connect to the service.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Compute Engine location where the instance resides. For trial organization
        subscriptions, the location must be a Compute Engine zone. For paid organization
        subscriptions, it should correspond to a Compute Engine region.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource ID of the instance.
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
    @pulumi.getter(name="peeringCidrRange")
    def peering_cidr_range(self) -> pulumi.Output[Optional[str]]:
        """
        The size of the CIDR block range that will be reserved by the instance.
        Possible values are `SLASH_16` and `SLASH_20`.
        """
        return pulumi.get(self, "peering_cidr_range")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[str]:
        """
        Output only. Port number of the exposed Apigee endpoint.
        """
        return pulumi.get(self, "port")

