# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TargetGrpcProxyArgs', 'TargetGrpcProxy']

@pulumi.input_type
class TargetGrpcProxyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 validate_for_proxyless: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a TargetGrpcProxy resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource
               is created. The name must be 1-63 characters long, and comply
               with RFC1035. Specifically, the name must be 1-63 characters long
               and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] url_map: URL to the UrlMap resource that defines the mapping from URL to
               the BackendService. The protocol field in the BackendService
               must be set to GRPC.
        :param pulumi.Input[bool] validate_for_proxyless: If true, indicates that the BackendServices referenced by
               the urlMap may be accessed by gRPC applications without using
               a sidecar proxy. This will enable configuration checks on urlMap
               and its referenced BackendServices to not allow unsupported features.
               A gRPC application must use "xds:///" scheme in the target URI
               of the service it is connecting to. If false, indicates that the
               BackendServices referenced by the urlMap will be accessed by gRPC
               applications via a sidecar proxy. In this case, a gRPC application
               must not use "xds:///" scheme in the target URI of the service
               it is connecting to
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if url_map is not None:
            pulumi.set(__self__, "url_map", url_map)
        if validate_for_proxyless is not None:
            pulumi.set(__self__, "validate_for_proxyless", validate_for_proxyless)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource
        is created. The name must be 1-63 characters long, and comply
        with RFC1035. Specifically, the name must be 1-63 characters long
        and match the regular expression `a-z?` which
        means the first character must be a lowercase letter, and all
        following characters must be a dash, lowercase letter, or digit,
        except the last character, which cannot be a dash.


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
    @pulumi.getter(name="urlMap")
    def url_map(self) -> Optional[pulumi.Input[str]]:
        """
        URL to the UrlMap resource that defines the mapping from URL to
        the BackendService. The protocol field in the BackendService
        must be set to GRPC.
        """
        return pulumi.get(self, "url_map")

    @url_map.setter
    def url_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_map", value)

    @property
    @pulumi.getter(name="validateForProxyless")
    def validate_for_proxyless(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, indicates that the BackendServices referenced by
        the urlMap may be accessed by gRPC applications without using
        a sidecar proxy. This will enable configuration checks on urlMap
        and its referenced BackendServices to not allow unsupported features.
        A gRPC application must use "xds:///" scheme in the target URI
        of the service it is connecting to. If false, indicates that the
        BackendServices referenced by the urlMap will be accessed by gRPC
        applications via a sidecar proxy. In this case, a gRPC application
        must not use "xds:///" scheme in the target URI of the service
        it is connecting to
        """
        return pulumi.get(self, "validate_for_proxyless")

    @validate_for_proxyless.setter
    def validate_for_proxyless(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_for_proxyless", value)


@pulumi.input_type
class _TargetGrpcProxyState:
    def __init__(__self__, *,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 self_link_with_id: Optional[pulumi.Input[str]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 validate_for_proxyless: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering TargetGrpcProxy resources.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in
               this object. This field is used in optimistic locking. This field
               will be ignored when inserting a TargetGrpcProxy. An up-to-date
               fingerprint must be provided in order to patch/update the
               TargetGrpcProxy; otherwise, the request will fail with error
               412 conditionNotMet. To see the latest fingerprint, make a get()
               request to retrieve the TargetGrpcProxy. A base64-encoded string.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource
               is created. The name must be 1-63 characters long, and comply
               with RFC1035. Specifically, the name must be 1-63 characters long
               and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] self_link_with_id: Server-defined URL with id for the resource.
        :param pulumi.Input[str] url_map: URL to the UrlMap resource that defines the mapping from URL to
               the BackendService. The protocol field in the BackendService
               must be set to GRPC.
        :param pulumi.Input[bool] validate_for_proxyless: If true, indicates that the BackendServices referenced by
               the urlMap may be accessed by gRPC applications without using
               a sidecar proxy. This will enable configuration checks on urlMap
               and its referenced BackendServices to not allow unsupported features.
               A gRPC application must use "xds:///" scheme in the target URI
               of the service it is connecting to. If false, indicates that the
               BackendServices referenced by the urlMap will be accessed by gRPC
               applications via a sidecar proxy. In this case, a gRPC application
               must not use "xds:///" scheme in the target URI of the service
               it is connecting to
        """
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id is not None:
            pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if url_map is not None:
            pulumi.set(__self__, "url_map", url_map)
        if validate_for_proxyless is not None:
            pulumi.set(__self__, "validate_for_proxyless", validate_for_proxyless)

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
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint of this resource. A hash of the contents stored in
        this object. This field is used in optimistic locking. This field
        will be ignored when inserting a TargetGrpcProxy. An up-to-date
        fingerprint must be provided in order to patch/update the
        TargetGrpcProxy; otherwise, the request will fail with error
        412 conditionNotMet. To see the latest fingerprint, make a get()
        request to retrieve the TargetGrpcProxy. A base64-encoded string.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource
        is created. The name must be 1-63 characters long, and comply
        with RFC1035. Specifically, the name must be 1-63 characters long
        and match the regular expression `a-z?` which
        means the first character must be a lowercase letter, and all
        following characters must be a dash, lowercase letter, or digit,
        except the last character, which cannot be a dash.


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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> Optional[pulumi.Input[str]]:
        """
        Server-defined URL with id for the resource.
        """
        return pulumi.get(self, "self_link_with_id")

    @self_link_with_id.setter
    def self_link_with_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link_with_id", value)

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> Optional[pulumi.Input[str]]:
        """
        URL to the UrlMap resource that defines the mapping from URL to
        the BackendService. The protocol field in the BackendService
        must be set to GRPC.
        """
        return pulumi.get(self, "url_map")

    @url_map.setter
    def url_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_map", value)

    @property
    @pulumi.getter(name="validateForProxyless")
    def validate_for_proxyless(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, indicates that the BackendServices referenced by
        the urlMap may be accessed by gRPC applications without using
        a sidecar proxy. This will enable configuration checks on urlMap
        and its referenced BackendServices to not allow unsupported features.
        A gRPC application must use "xds:///" scheme in the target URI
        of the service it is connecting to. If false, indicates that the
        BackendServices referenced by the urlMap will be accessed by gRPC
        applications via a sidecar proxy. In this case, a gRPC application
        must not use "xds:///" scheme in the target URI of the service
        it is connecting to
        """
        return pulumi.get(self, "validate_for_proxyless")

    @validate_for_proxyless.setter
    def validate_for_proxyless(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_for_proxyless", value)


class TargetGrpcProxy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 validate_for_proxyless: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Represents a Target gRPC Proxy resource. A target gRPC proxy is a component
        of load balancers intended for load balancing gRPC traffic. Global forwarding
        rules reference a target gRPC proxy. The Target gRPC Proxy references
        a URL map which specifies how traffic routes to gRPC backend services.

        To get more information about TargetGrpcProxy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetGrpcProxies)
        * How-to Guides
            * [Using Target gRPC Proxies](https://cloud.google.com/traffic-director/docs/proxyless-overview)

        ## Example Usage

        ## Import

        TargetGrpcProxy can be imported using any of these accepted formats* `projects/{{project}}/global/targetGrpcProxies/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, TargetGrpcProxy can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default projects/{{project}}/global/targetGrpcProxies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource
               is created. The name must be 1-63 characters long, and comply
               with RFC1035. Specifically, the name must be 1-63 characters long
               and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] url_map: URL to the UrlMap resource that defines the mapping from URL to
               the BackendService. The protocol field in the BackendService
               must be set to GRPC.
        :param pulumi.Input[bool] validate_for_proxyless: If true, indicates that the BackendServices referenced by
               the urlMap may be accessed by gRPC applications without using
               a sidecar proxy. This will enable configuration checks on urlMap
               and its referenced BackendServices to not allow unsupported features.
               A gRPC application must use "xds:///" scheme in the target URI
               of the service it is connecting to. If false, indicates that the
               BackendServices referenced by the urlMap will be accessed by gRPC
               applications via a sidecar proxy. In this case, a gRPC application
               must not use "xds:///" scheme in the target URI of the service
               it is connecting to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TargetGrpcProxyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a Target gRPC Proxy resource. A target gRPC proxy is a component
        of load balancers intended for load balancing gRPC traffic. Global forwarding
        rules reference a target gRPC proxy. The Target gRPC Proxy references
        a URL map which specifies how traffic routes to gRPC backend services.

        To get more information about TargetGrpcProxy, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetGrpcProxies)
        * How-to Guides
            * [Using Target gRPC Proxies](https://cloud.google.com/traffic-director/docs/proxyless-overview)

        ## Example Usage

        ## Import

        TargetGrpcProxy can be imported using any of these accepted formats* `projects/{{project}}/global/targetGrpcProxies/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, TargetGrpcProxy can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default projects/{{project}}/global/targetGrpcProxies/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/targetGrpcProxy:TargetGrpcProxy default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param TargetGrpcProxyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetGrpcProxyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 url_map: Optional[pulumi.Input[str]] = None,
                 validate_for_proxyless: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TargetGrpcProxyArgs.__new__(TargetGrpcProxyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["url_map"] = url_map
            __props__.__dict__["validate_for_proxyless"] = validate_for_proxyless
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["fingerprint"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["self_link_with_id"] = None
        super(TargetGrpcProxy, __self__).__init__(
            'gcp:compute/targetGrpcProxy:TargetGrpcProxy',
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
            self_link: Optional[pulumi.Input[str]] = None,
            self_link_with_id: Optional[pulumi.Input[str]] = None,
            url_map: Optional[pulumi.Input[str]] = None,
            validate_for_proxyless: Optional[pulumi.Input[bool]] = None) -> 'TargetGrpcProxy':
        """
        Get an existing TargetGrpcProxy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] description: An optional description of this resource.
        :param pulumi.Input[str] fingerprint: Fingerprint of this resource. A hash of the contents stored in
               this object. This field is used in optimistic locking. This field
               will be ignored when inserting a TargetGrpcProxy. An up-to-date
               fingerprint must be provided in order to patch/update the
               TargetGrpcProxy; otherwise, the request will fail with error
               412 conditionNotMet. To see the latest fingerprint, make a get()
               request to retrieve the TargetGrpcProxy. A base64-encoded string.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource
               is created. The name must be 1-63 characters long, and comply
               with RFC1035. Specifically, the name must be 1-63 characters long
               and match the regular expression `a-z?` which
               means the first character must be a lowercase letter, and all
               following characters must be a dash, lowercase letter, or digit,
               except the last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[str] self_link_with_id: Server-defined URL with id for the resource.
        :param pulumi.Input[str] url_map: URL to the UrlMap resource that defines the mapping from URL to
               the BackendService. The protocol field in the BackendService
               must be set to GRPC.
        :param pulumi.Input[bool] validate_for_proxyless: If true, indicates that the BackendServices referenced by
               the urlMap may be accessed by gRPC applications without using
               a sidecar proxy. This will enable configuration checks on urlMap
               and its referenced BackendServices to not allow unsupported features.
               A gRPC application must use "xds:///" scheme in the target URI
               of the service it is connecting to. If false, indicates that the
               BackendServices referenced by the urlMap will be accessed by gRPC
               applications via a sidecar proxy. In this case, a gRPC application
               must not use "xds:///" scheme in the target URI of the service
               it is connecting to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TargetGrpcProxyState.__new__(_TargetGrpcProxyState)

        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["description"] = description
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["self_link_with_id"] = self_link_with_id
        __props__.__dict__["url_map"] = url_map
        __props__.__dict__["validate_for_proxyless"] = validate_for_proxyless
        return TargetGrpcProxy(resource_name, opts=opts, __props__=__props__)

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
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        Fingerprint of this resource. A hash of the contents stored in
        this object. This field is used in optimistic locking. This field
        will be ignored when inserting a TargetGrpcProxy. An up-to-date
        fingerprint must be provided in order to patch/update the
        TargetGrpcProxy; otherwise, the request will fail with error
        412 conditionNotMet. To see the latest fingerprint, make a get()
        request to retrieve the TargetGrpcProxy. A base64-encoded string.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource
        is created. The name must be 1-63 characters long, and comply
        with RFC1035. Specifically, the name must be 1-63 characters long
        and match the regular expression `a-z?` which
        means the first character must be a lowercase letter, and all
        following characters must be a dash, lowercase letter, or digit,
        except the last character, which cannot be a dash.


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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The URI of the created resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> pulumi.Output[str]:
        """
        Server-defined URL with id for the resource.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="urlMap")
    def url_map(self) -> pulumi.Output[Optional[str]]:
        """
        URL to the UrlMap resource that defines the mapping from URL to
        the BackendService. The protocol field in the BackendService
        must be set to GRPC.
        """
        return pulumi.get(self, "url_map")

    @property
    @pulumi.getter(name="validateForProxyless")
    def validate_for_proxyless(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, indicates that the BackendServices referenced by
        the urlMap may be accessed by gRPC applications without using
        a sidecar proxy. This will enable configuration checks on urlMap
        and its referenced BackendServices to not allow unsupported features.
        A gRPC application must use "xds:///" scheme in the target URI
        of the service it is connecting to. If false, indicates that the
        BackendServices referenced by the urlMap will be accessed by gRPC
        applications via a sidecar proxy. In this case, a gRPC application
        must not use "xds:///" scheme in the target URI of the service
        it is connecting to
        """
        return pulumi.get(self, "validate_for_proxyless")

