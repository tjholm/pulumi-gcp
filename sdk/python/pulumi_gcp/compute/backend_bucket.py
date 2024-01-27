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

__all__ = ['BackendBucketArgs', 'BackendBucket']

@pulumi.input_type
class BackendBucketArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 cdn_policy: Optional[pulumi.Input['BackendBucketCdnPolicyArgs']] = None,
                 compression_mode: Optional[pulumi.Input[str]] = None,
                 custom_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_security_policy: Optional[pulumi.Input[str]] = None,
                 enable_cdn: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BackendBucket resource.
        :param pulumi.Input[str] bucket_name: Cloud Storage bucket name.
        :param pulumi.Input['BackendBucketCdnPolicyArgs'] cdn_policy: Cloud CDN configuration for this Backend Bucket.
               Structure is documented below.
        :param pulumi.Input[str] compression_mode: Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
               Possible values are: `AUTOMATIC`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_response_headers: Headers that the HTTP/S load balancer should add to proxied responses.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the
               client when the resource is created.
        :param pulumi.Input[str] edge_security_policy: The security policy associated with this backend bucket.
        :param pulumi.Input[bool] enable_cdn: If true, enable Cloud CDN for this BackendBucket.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        if cdn_policy is not None:
            pulumi.set(__self__, "cdn_policy", cdn_policy)
        if compression_mode is not None:
            pulumi.set(__self__, "compression_mode", compression_mode)
        if custom_response_headers is not None:
            pulumi.set(__self__, "custom_response_headers", custom_response_headers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if edge_security_policy is not None:
            pulumi.set(__self__, "edge_security_policy", edge_security_policy)
        if enable_cdn is not None:
            pulumi.set(__self__, "enable_cdn", enable_cdn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        Cloud Storage bucket name.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="cdnPolicy")
    def cdn_policy(self) -> Optional[pulumi.Input['BackendBucketCdnPolicyArgs']]:
        """
        Cloud CDN configuration for this Backend Bucket.
        Structure is documented below.
        """
        return pulumi.get(self, "cdn_policy")

    @cdn_policy.setter
    def cdn_policy(self, value: Optional[pulumi.Input['BackendBucketCdnPolicyArgs']]):
        pulumi.set(self, "cdn_policy", value)

    @property
    @pulumi.getter(name="compressionMode")
    def compression_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        Possible values are: `AUTOMATIC`, `DISABLED`.
        """
        return pulumi.get(self, "compression_mode")

    @compression_mode.setter
    def compression_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compression_mode", value)

    @property
    @pulumi.getter(name="customResponseHeaders")
    def custom_response_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Headers that the HTTP/S load balancer should add to proxied responses.
        """
        return pulumi.get(self, "custom_response_headers")

    @custom_response_headers.setter
    def custom_response_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_response_headers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional textual description of the resource; provided by the
        client when the resource is created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="edgeSecurityPolicy")
    def edge_security_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The security policy associated with this backend bucket.
        """
        return pulumi.get(self, "edge_security_policy")

    @edge_security_policy.setter
    def edge_security_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_security_policy", value)

    @property
    @pulumi.getter(name="enableCdn")
    def enable_cdn(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, enable Cloud CDN for this BackendBucket.
        """
        return pulumi.get(self, "enable_cdn")

    @enable_cdn.setter
    def enable_cdn(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_cdn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the
        last character, which cannot be a dash.


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
class _BackendBucketState:
    def __init__(__self__, *,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 cdn_policy: Optional[pulumi.Input['BackendBucketCdnPolicyArgs']] = None,
                 compression_mode: Optional[pulumi.Input[str]] = None,
                 creation_timestamp: Optional[pulumi.Input[str]] = None,
                 custom_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_security_policy: Optional[pulumi.Input[str]] = None,
                 enable_cdn: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BackendBucket resources.
        :param pulumi.Input[str] bucket_name: Cloud Storage bucket name.
        :param pulumi.Input['BackendBucketCdnPolicyArgs'] cdn_policy: Cloud CDN configuration for this Backend Bucket.
               Structure is documented below.
        :param pulumi.Input[str] compression_mode: Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
               Possible values are: `AUTOMATIC`, `DISABLED`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_response_headers: Headers that the HTTP/S load balancer should add to proxied responses.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the
               client when the resource is created.
        :param pulumi.Input[str] edge_security_policy: The security policy associated with this backend bucket.
        :param pulumi.Input[bool] enable_cdn: If true, enable Cloud CDN for this BackendBucket.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if cdn_policy is not None:
            pulumi.set(__self__, "cdn_policy", cdn_policy)
        if compression_mode is not None:
            pulumi.set(__self__, "compression_mode", compression_mode)
        if creation_timestamp is not None:
            pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if custom_response_headers is not None:
            pulumi.set(__self__, "custom_response_headers", custom_response_headers)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if edge_security_policy is not None:
            pulumi.set(__self__, "edge_security_policy", edge_security_policy)
        if enable_cdn is not None:
            pulumi.set(__self__, "enable_cdn", enable_cdn)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud Storage bucket name.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="cdnPolicy")
    def cdn_policy(self) -> Optional[pulumi.Input['BackendBucketCdnPolicyArgs']]:
        """
        Cloud CDN configuration for this Backend Bucket.
        Structure is documented below.
        """
        return pulumi.get(self, "cdn_policy")

    @cdn_policy.setter
    def cdn_policy(self, value: Optional[pulumi.Input['BackendBucketCdnPolicyArgs']]):
        pulumi.set(self, "cdn_policy", value)

    @property
    @pulumi.getter(name="compressionMode")
    def compression_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        Possible values are: `AUTOMATIC`, `DISABLED`.
        """
        return pulumi.get(self, "compression_mode")

    @compression_mode.setter
    def compression_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compression_mode", value)

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
    @pulumi.getter(name="customResponseHeaders")
    def custom_response_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Headers that the HTTP/S load balancer should add to proxied responses.
        """
        return pulumi.get(self, "custom_response_headers")

    @custom_response_headers.setter
    def custom_response_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_response_headers", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional textual description of the resource; provided by the
        client when the resource is created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="edgeSecurityPolicy")
    def edge_security_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The security policy associated with this backend bucket.
        """
        return pulumi.get(self, "edge_security_policy")

    @edge_security_policy.setter
    def edge_security_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_security_policy", value)

    @property
    @pulumi.getter(name="enableCdn")
    def enable_cdn(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, enable Cloud CDN for this BackendBucket.
        """
        return pulumi.get(self, "enable_cdn")

    @enable_cdn.setter
    def enable_cdn(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_cdn", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the
        last character, which cannot be a dash.


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


class BackendBucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 cdn_policy: Optional[pulumi.Input[pulumi.InputType['BackendBucketCdnPolicyArgs']]] = None,
                 compression_mode: Optional[pulumi.Input[str]] = None,
                 custom_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_security_policy: Optional[pulumi.Input[str]] = None,
                 enable_cdn: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S)
        load balancing.

        An HTTP(S) load balancer can direct traffic to specified URLs to a
        backend bucket rather than a backend service. It can send requests for
        static content to a Cloud Storage bucket and requests for dynamic content
        to a virtual machine instance.

        To get more information about BackendBucket, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendBuckets)
        * How-to Guides
            * [Using a Cloud Storage bucket as a load balancer backend](https://cloud.google.com/compute/docs/load-balancing/http/backend-bucket)

        ## Example Usage
        ### Backend Bucket Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_bucket = gcp.storage.Bucket("imageBucket", location="EU")
        image_backend = gcp.compute.BackendBucket("imageBackend",
            description="Contains beautiful images",
            bucket_name=image_bucket.name,
            enable_cdn=True)
        ```
        ### Backend Bucket Security Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_backend_bucket = gcp.storage.Bucket("imageBackendBucket", location="EU")
        policy = gcp.compute.SecurityPolicy("policy",
            description="basic security policy",
            type="CLOUD_ARMOR_EDGE")
        image_backend_backend_bucket = gcp.compute.BackendBucket("imageBackendBackendBucket",
            description="Contains beautiful images",
            bucket_name=image_backend_bucket.name,
            enable_cdn=True,
            edge_security_policy=policy.id)
        ```
        ### Backend Bucket Query String Whitelist

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_bucket = gcp.storage.Bucket("imageBucket", location="EU")
        image_backend = gcp.compute.BackendBucket("imageBackend",
            description="Contains beautiful images",
            bucket_name=image_bucket.name,
            enable_cdn=True,
            cdn_policy=gcp.compute.BackendBucketCdnPolicyArgs(
                cache_key_policy=gcp.compute.BackendBucketCdnPolicyCacheKeyPolicyArgs(
                    query_string_whitelists=["image-version"],
                ),
            ))
        ```
        ### Backend Bucket Include Http Headers

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_bucket = gcp.storage.Bucket("imageBucket", location="EU")
        image_backend = gcp.compute.BackendBucket("imageBackend",
            description="Contains beautiful images",
            bucket_name=image_bucket.name,
            enable_cdn=True,
            cdn_policy=gcp.compute.BackendBucketCdnPolicyArgs(
                cache_key_policy=gcp.compute.BackendBucketCdnPolicyCacheKeyPolicyArgs(
                    include_http_headers=["X-My-Header-Field"],
                ),
            ))
        ```

        ## Import

        BackendBucket can be imported using any of these accepted formats* `projects/{{project}}/global/backendBuckets/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, BackendBucket can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:compute/backendBucket:BackendBucket default projects/{{project}}/global/backendBuckets/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/backendBucket:BackendBucket default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/backendBucket:BackendBucket default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: Cloud Storage bucket name.
        :param pulumi.Input[pulumi.InputType['BackendBucketCdnPolicyArgs']] cdn_policy: Cloud CDN configuration for this Backend Bucket.
               Structure is documented below.
        :param pulumi.Input[str] compression_mode: Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
               Possible values are: `AUTOMATIC`, `DISABLED`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_response_headers: Headers that the HTTP/S load balancer should add to proxied responses.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the
               client when the resource is created.
        :param pulumi.Input[str] edge_security_policy: The security policy associated with this backend bucket.
        :param pulumi.Input[bool] enable_cdn: If true, enable Cloud CDN for this BackendBucket.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackendBucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S)
        load balancing.

        An HTTP(S) load balancer can direct traffic to specified URLs to a
        backend bucket rather than a backend service. It can send requests for
        static content to a Cloud Storage bucket and requests for dynamic content
        to a virtual machine instance.

        To get more information about BackendBucket, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendBuckets)
        * How-to Guides
            * [Using a Cloud Storage bucket as a load balancer backend](https://cloud.google.com/compute/docs/load-balancing/http/backend-bucket)

        ## Example Usage
        ### Backend Bucket Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_bucket = gcp.storage.Bucket("imageBucket", location="EU")
        image_backend = gcp.compute.BackendBucket("imageBackend",
            description="Contains beautiful images",
            bucket_name=image_bucket.name,
            enable_cdn=True)
        ```
        ### Backend Bucket Security Policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_backend_bucket = gcp.storage.Bucket("imageBackendBucket", location="EU")
        policy = gcp.compute.SecurityPolicy("policy",
            description="basic security policy",
            type="CLOUD_ARMOR_EDGE")
        image_backend_backend_bucket = gcp.compute.BackendBucket("imageBackendBackendBucket",
            description="Contains beautiful images",
            bucket_name=image_backend_bucket.name,
            enable_cdn=True,
            edge_security_policy=policy.id)
        ```
        ### Backend Bucket Query String Whitelist

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_bucket = gcp.storage.Bucket("imageBucket", location="EU")
        image_backend = gcp.compute.BackendBucket("imageBackend",
            description="Contains beautiful images",
            bucket_name=image_bucket.name,
            enable_cdn=True,
            cdn_policy=gcp.compute.BackendBucketCdnPolicyArgs(
                cache_key_policy=gcp.compute.BackendBucketCdnPolicyCacheKeyPolicyArgs(
                    query_string_whitelists=["image-version"],
                ),
            ))
        ```
        ### Backend Bucket Include Http Headers

        ```python
        import pulumi
        import pulumi_gcp as gcp

        image_bucket = gcp.storage.Bucket("imageBucket", location="EU")
        image_backend = gcp.compute.BackendBucket("imageBackend",
            description="Contains beautiful images",
            bucket_name=image_bucket.name,
            enable_cdn=True,
            cdn_policy=gcp.compute.BackendBucketCdnPolicyArgs(
                cache_key_policy=gcp.compute.BackendBucketCdnPolicyCacheKeyPolicyArgs(
                    include_http_headers=["X-My-Header-Field"],
                ),
            ))
        ```

        ## Import

        BackendBucket can be imported using any of these accepted formats* `projects/{{project}}/global/backendBuckets/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, BackendBucket can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:compute/backendBucket:BackendBucket default projects/{{project}}/global/backendBuckets/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/backendBucket:BackendBucket default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/backendBucket:BackendBucket default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param BackendBucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackendBucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 cdn_policy: Optional[pulumi.Input[pulumi.InputType['BackendBucketCdnPolicyArgs']]] = None,
                 compression_mode: Optional[pulumi.Input[str]] = None,
                 custom_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_security_policy: Optional[pulumi.Input[str]] = None,
                 enable_cdn: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackendBucketArgs.__new__(BackendBucketArgs)

            if bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_name'")
            __props__.__dict__["bucket_name"] = bucket_name
            __props__.__dict__["cdn_policy"] = cdn_policy
            __props__.__dict__["compression_mode"] = compression_mode
            __props__.__dict__["custom_response_headers"] = custom_response_headers
            __props__.__dict__["description"] = description
            __props__.__dict__["edge_security_policy"] = edge_security_policy
            __props__.__dict__["enable_cdn"] = enable_cdn
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["creation_timestamp"] = None
            __props__.__dict__["self_link"] = None
        super(BackendBucket, __self__).__init__(
            'gcp:compute/backendBucket:BackendBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_name: Optional[pulumi.Input[str]] = None,
            cdn_policy: Optional[pulumi.Input[pulumi.InputType['BackendBucketCdnPolicyArgs']]] = None,
            compression_mode: Optional[pulumi.Input[str]] = None,
            creation_timestamp: Optional[pulumi.Input[str]] = None,
            custom_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            edge_security_policy: Optional[pulumi.Input[str]] = None,
            enable_cdn: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None) -> 'BackendBucket':
        """
        Get an existing BackendBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: Cloud Storage bucket name.
        :param pulumi.Input[pulumi.InputType['BackendBucketCdnPolicyArgs']] cdn_policy: Cloud CDN configuration for this Backend Bucket.
               Structure is documented below.
        :param pulumi.Input[str] compression_mode: Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
               Possible values are: `AUTOMATIC`, `DISABLED`.
        :param pulumi.Input[str] creation_timestamp: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] custom_response_headers: Headers that the HTTP/S load balancer should add to proxied responses.
        :param pulumi.Input[str] description: An optional textual description of the resource; provided by the
               client when the resource is created.
        :param pulumi.Input[str] edge_security_policy: The security policy associated with this backend bucket.
        :param pulumi.Input[bool] enable_cdn: If true, enable Cloud CDN for this BackendBucket.
        :param pulumi.Input[str] name: Name of the resource. Provided by the client when the resource is
               created. The name must be 1-63 characters long, and comply with
               RFC1035.  Specifically, the name must be 1-63 characters long and
               match the regular expression `a-z?` which means
               the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the
               last character, which cannot be a dash.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackendBucketState.__new__(_BackendBucketState)

        __props__.__dict__["bucket_name"] = bucket_name
        __props__.__dict__["cdn_policy"] = cdn_policy
        __props__.__dict__["compression_mode"] = compression_mode
        __props__.__dict__["creation_timestamp"] = creation_timestamp
        __props__.__dict__["custom_response_headers"] = custom_response_headers
        __props__.__dict__["description"] = description
        __props__.__dict__["edge_security_policy"] = edge_security_policy
        __props__.__dict__["enable_cdn"] = enable_cdn
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        return BackendBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[str]:
        """
        Cloud Storage bucket name.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="cdnPolicy")
    def cdn_policy(self) -> pulumi.Output['outputs.BackendBucketCdnPolicy']:
        """
        Cloud CDN configuration for this Backend Bucket.
        Structure is documented below.
        """
        return pulumi.get(self, "cdn_policy")

    @property
    @pulumi.getter(name="compressionMode")
    def compression_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        Possible values are: `AUTOMATIC`, `DISABLED`.
        """
        return pulumi.get(self, "compression_mode")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customResponseHeaders")
    def custom_response_headers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Headers that the HTTP/S load balancer should add to proxied responses.
        """
        return pulumi.get(self, "custom_response_headers")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional textual description of the resource; provided by the
        client when the resource is created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeSecurityPolicy")
    def edge_security_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The security policy associated with this backend bucket.
        """
        return pulumi.get(self, "edge_security_policy")

    @property
    @pulumi.getter(name="enableCdn")
    def enable_cdn(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, enable Cloud CDN for this BackendBucket.
        """
        return pulumi.get(self, "enable_cdn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the resource. Provided by the client when the resource is
        created. The name must be 1-63 characters long, and comply with
        RFC1035.  Specifically, the name must be 1-63 characters long and
        match the regular expression `a-z?` which means
        the first character must be a lowercase letter, and all following
        characters must be a dash, lowercase letter, or digit, except the
        last character, which cannot be a dash.


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

