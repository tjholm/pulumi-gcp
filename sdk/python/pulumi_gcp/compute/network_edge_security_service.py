# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkEdgeSecurityServiceArgs', 'NetworkEdgeSecurityService']

@pulumi.input_type
class NetworkEdgeSecurityServiceArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_policy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetworkEdgeSecurityService resource.
        :param pulumi.Input[str] description: Free-text description of the resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway security policy.
        :param pulumi.Input[str] security_policy: The resource URL for the network edge security service associated with this network edge security service.
        """
        NetworkEdgeSecurityServiceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            name=name,
            project=project,
            region=region,
            security_policy=security_policy,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             security_policy: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if security_policy is None and 'securityPolicy' in kwargs:
            security_policy = kwargs['securityPolicy']

        if description is not None:
            _setter("description", description)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if security_policy is not None:
            _setter("security_policy", security_policy)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created.


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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the gateway security policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The resource URL for the network edge security service associated with this network edge security service.
        """
        return pulumi.get(self, "security_policy")

    @security_policy.setter
    def security_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy", value)


@pulumi.input_type
class _NetworkEdgeSecurityServiceState:
    def __init__(__self__, *,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_policy: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_service_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkEdgeSecurityService resources.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: Free-text description of the resource.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService.
               An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway security policy.
        :param pulumi.Input[str] security_policy: The resource URL for the network edge security service associated with this network edge security service.
        :param pulumi.Input[str] self_link: Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_service_id: Server-defined URL for this resource with the resource id.
        :param pulumi.Input[str] service_id: The unique identifier for the resource. This identifier is defined by the server.
        """
        _NetworkEdgeSecurityServiceState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            creation_timestamp=creation_timestamp,
            description=description,
            fingerprint=fingerprint,
            name=name,
            project=project,
            region=region,
            security_policy=security_policy,
            self_link=self_link,
            self_link_with_service_id=self_link_with_service_id,
            service_id=service_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             creation_timestamp: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             fingerprint: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             security_policy: Optional[pulumi.Input[str]] = None,
             self_link: Optional[pulumi.Input[str]] = None,
             self_link_with_service_id: Optional[pulumi.Input[str]] = None,
             service_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if creation_timestamp is None and 'creationTimestamp' in kwargs:
            creation_timestamp = kwargs['creationTimestamp']
        if security_policy is None and 'securityPolicy' in kwargs:
            security_policy = kwargs['securityPolicy']
        if self_link is None and 'selfLink' in kwargs:
            self_link = kwargs['selfLink']
        if self_link_with_service_id is None and 'selfLinkWithServiceId' in kwargs:
            self_link_with_service_id = kwargs['selfLinkWithServiceId']
        if service_id is None and 'serviceId' in kwargs:
            service_id = kwargs['serviceId']

        if creation_timestamp is not None:
            _setter("creation_timestamp", creation_timestamp)
        if description is not None:
            _setter("description", description)
        if fingerprint is not None:
            _setter("fingerprint", fingerprint)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if security_policy is not None:
            _setter("security_policy", security_policy)
        if self_link is not None:
            _setter("self_link", self_link)
        if self_link_with_service_id is not None:
            _setter("self_link_with_service_id", self_link_with_service_id)
        if service_id is not None:
            _setter("service_id", service_id)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @creation_timestamp.setter
    def creation_timestamp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_timestamp", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService.
        An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is created.


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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the gateway security policy.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The resource URL for the network edge security service associated with this network edge security service.
        """
        return pulumi.get(self, "security_policy")

    @security_policy.setter
    def security_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_policy", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="selfLinkWithServiceId")
    def self_link_with_service_id(self) -> Optional[pulumi.Input[str]]:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_service_id")

    @self_link_with_service_id.setter
    def self_link_with_service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link_with_service_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)


class NetworkEdgeSecurityService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Compute Network Edge Security Service Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.NetworkEdgeSecurityService("default",
            region="us-east1",
            description="My basic resource",
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        NetworkEdgeSecurityService can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default projects/{{project}}/regions/{{region}}/networkEdgeSecurityServices/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Free-text description of the resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway security policy.
        :param pulumi.Input[str] security_policy: The resource URL for the network edge security service associated with this network edge security service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkEdgeSecurityServiceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Compute Network Edge Security Service Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.compute.NetworkEdgeSecurityService("default",
            region="us-east1",
            description="My basic resource",
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        NetworkEdgeSecurityService can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default projects/{{project}}/regions/{{region}}/networkEdgeSecurityServices/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param NetworkEdgeSecurityServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkEdgeSecurityServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            NetworkEdgeSecurityServiceArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkEdgeSecurityServiceArgs.__new__(NetworkEdgeSecurityServiceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["security_policy"] = security_policy
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["self_link_with_service_id"] = None
            __props__.__dict__["service_id"] = None
        super(NetworkEdgeSecurityService, __self__).__init__(
            'gcp:compute/networkEdgeSecurityService:NetworkEdgeSecurityService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            security_policy: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            self_link_with_service_id: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'NetworkEdgeSecurityService':
        """
        Get an existing NetworkEdgeSecurityService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: Free-text description of the resource.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService.
               An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway security policy.
        :param pulumi.Input[str] security_policy: The resource URL for the network edge security service associated with this network edge security service.
        :param pulumi.Input[str] self_link: Server-defined URL for the resource.
        :param pulumi.Input[str] self_link_with_service_id: Server-defined URL for this resource with the resource id.
        :param pulumi.Input[str] service_id: The unique identifier for the resource. This identifier is defined by the server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkEdgeSecurityServiceState.__new__(_NetworkEdgeSecurityServiceState)

        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["description"] = description
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["security_policy"] = security_policy
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["self_link_with_service_id"] = self_link_with_service_id
        __props__.__dict__["service_id"] = service_id
        return NetworkEdgeSecurityService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Free-text description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a NetworkEdgeSecurityService.
        An up-to-date fingerprint must be provided in order to update the NetworkEdgeSecurityService, otherwise the request will fail with error 412 conditionNotMet.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is created.


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
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region of the gateway security policy.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The resource URL for the network edge security service associated with this network edge security service.
        """
        return pulumi.get(self, "security_policy")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithServiceId")
    def self_link_with_service_id(self) -> pulumi.Output[str]:
        """
        Server-defined URL for this resource with the resource id.
        """
        return pulumi.get(self, "self_link_with_service_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        The unique identifier for the resource. This identifier is defined by the server.
        """
        return pulumi.get(self, "service_id")

