# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 service_name: pulumi.Input[str],
                 grpc_config: Optional[pulumi.Input[str]] = None,
                 openapi_config: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protoc_output_base64: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input[str] service_name: The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        :param pulumi.Input[str] grpc_config: The full text of the Service Config YAML file (Example located here). If provided, must also provide
               protoc_output_base64. open_api config must not be provided.
        :param pulumi.Input[str] openapi_config: The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
               protoc_output_base64 must be specified.
        :param pulumi.Input[str] project: The project ID that the service belongs to. If not provided, provider project is used.
        :param pulumi.Input[str] protoc_output_base64: The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
               base64-encoded.
        """
        pulumi.set(__self__, "service_name", service_name)
        if grpc_config is not None:
            pulumi.set(__self__, "grpc_config", grpc_config)
        if openapi_config is not None:
            pulumi.set(__self__, "openapi_config", openapi_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if protoc_output_base64 is not None:
            pulumi.set(__self__, "protoc_output_base64", protoc_output_base64)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter(name="grpcConfig")
    def grpc_config(self) -> Optional[pulumi.Input[str]]:
        """
        The full text of the Service Config YAML file (Example located here). If provided, must also provide
        protoc_output_base64. open_api config must not be provided.
        """
        return pulumi.get(self, "grpc_config")

    @grpc_config.setter
    def grpc_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grpc_config", value)

    @property
    @pulumi.getter(name="openapiConfig")
    def openapi_config(self) -> Optional[pulumi.Input[str]]:
        """
        The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
        protoc_output_base64 must be specified.
        """
        return pulumi.get(self, "openapi_config")

    @openapi_config.setter
    def openapi_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "openapi_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID that the service belongs to. If not provided, provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="protocOutputBase64")
    def protoc_output_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
        base64-encoded.
        """
        return pulumi.get(self, "protoc_output_base64")

    @protoc_output_base64.setter
    def protoc_output_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protoc_output_base64", value)


@pulumi.input_type
class _ServiceState:
    def __init__(__self__, *,
                 apis: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceApiArgs']]]] = None,
                 config_id: Optional[pulumi.Input[str]] = None,
                 dns_address: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointArgs']]]] = None,
                 grpc_config: Optional[pulumi.Input[str]] = None,
                 openapi_config: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protoc_output_base64: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Service resources.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceApiArgs']]] apis: A list of API objects.
        :param pulumi.Input[str] config_id: The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
               to compute engine instances as a tag.
        :param pulumi.Input[str] dns_address: The address at which the service can be found - usually the same as the service name.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceEndpointArgs']]] endpoints: A list of Endpoint objects.
        :param pulumi.Input[str] grpc_config: The full text of the Service Config YAML file (Example located here). If provided, must also provide
               protoc_output_base64. open_api config must not be provided.
        :param pulumi.Input[str] openapi_config: The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
               protoc_output_base64 must be specified.
        :param pulumi.Input[str] project: The project ID that the service belongs to. If not provided, provider project is used.
        :param pulumi.Input[str] protoc_output_base64: The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
               base64-encoded.
        :param pulumi.Input[str] service_name: The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        if apis is not None:
            pulumi.set(__self__, "apis", apis)
        if config_id is not None:
            pulumi.set(__self__, "config_id", config_id)
        if dns_address is not None:
            pulumi.set(__self__, "dns_address", dns_address)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if grpc_config is not None:
            pulumi.set(__self__, "grpc_config", grpc_config)
        if openapi_config is not None:
            pulumi.set(__self__, "openapi_config", openapi_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if protoc_output_base64 is not None:
            pulumi.set(__self__, "protoc_output_base64", protoc_output_base64)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def apis(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceApiArgs']]]]:
        """
        A list of API objects.
        """
        return pulumi.get(self, "apis")

    @apis.setter
    def apis(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceApiArgs']]]]):
        pulumi.set(self, "apis", value)

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> Optional[pulumi.Input[str]]:
        """
        The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
        to compute engine instances as a tag.
        """
        return pulumi.get(self, "config_id")

    @config_id.setter
    def config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config_id", value)

    @property
    @pulumi.getter(name="dnsAddress")
    def dns_address(self) -> Optional[pulumi.Input[str]]:
        """
        The address at which the service can be found - usually the same as the service name.
        """
        return pulumi.get(self, "dns_address")

    @dns_address.setter
    def dns_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_address", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointArgs']]]]:
        """
        A list of Endpoint objects.
        """
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceEndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter(name="grpcConfig")
    def grpc_config(self) -> Optional[pulumi.Input[str]]:
        """
        The full text of the Service Config YAML file (Example located here). If provided, must also provide
        protoc_output_base64. open_api config must not be provided.
        """
        return pulumi.get(self, "grpc_config")

    @grpc_config.setter
    def grpc_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grpc_config", value)

    @property
    @pulumi.getter(name="openapiConfig")
    def openapi_config(self) -> Optional[pulumi.Input[str]]:
        """
        The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
        protoc_output_base64 must be specified.
        """
        return pulumi.get(self, "openapi_config")

    @openapi_config.setter
    def openapi_config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "openapi_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID that the service belongs to. If not provided, provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="protocOutputBase64")
    def protoc_output_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
        base64-encoded.
        """
        return pulumi.get(self, "protoc_output_base64")

    @protoc_output_base64.setter
    def protoc_output_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protoc_output_base64", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grpc_config: Optional[pulumi.Input[str]] = None,
                 openapi_config: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protoc_output_base64: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grpc_config: The full text of the Service Config YAML file (Example located here). If provided, must also provide
               protoc_output_base64. open_api config must not be provided.
        :param pulumi.Input[str] openapi_config: The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
               protoc_output_base64 must be specified.
        :param pulumi.Input[str] project: The project ID that the service belongs to. If not provided, provider project is used.
        :param pulumi.Input[str] protoc_output_base64: The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
               base64-encoded.
        :param pulumi.Input[str] service_name: The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param ServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grpc_config: Optional[pulumi.Input[str]] = None,
                 openapi_config: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 protoc_output_base64: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["grpc_config"] = grpc_config
            __props__.__dict__["openapi_config"] = openapi_config
            __props__.__dict__["project"] = project
            __props__.__dict__["protoc_output_base64"] = protoc_output_base64
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["apis"] = None
            __props__.__dict__["config_id"] = None
            __props__.__dict__["dns_address"] = None
            __props__.__dict__["endpoints"] = None
        super(Service, __self__).__init__(
            'gcp:endpoints/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apis: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceApiArgs']]]]] = None,
            config_id: Optional[pulumi.Input[str]] = None,
            dns_address: Optional[pulumi.Input[str]] = None,
            endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceEndpointArgs']]]]] = None,
            grpc_config: Optional[pulumi.Input[str]] = None,
            openapi_config: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            protoc_output_base64: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceApiArgs']]]] apis: A list of API objects.
        :param pulumi.Input[str] config_id: The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
               to compute engine instances as a tag.
        :param pulumi.Input[str] dns_address: The address at which the service can be found - usually the same as the service name.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceEndpointArgs']]]] endpoints: A list of Endpoint objects.
        :param pulumi.Input[str] grpc_config: The full text of the Service Config YAML file (Example located here). If provided, must also provide
               protoc_output_base64. open_api config must not be provided.
        :param pulumi.Input[str] openapi_config: The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
               protoc_output_base64 must be specified.
        :param pulumi.Input[str] project: The project ID that the service belongs to. If not provided, provider project is used.
        :param pulumi.Input[str] protoc_output_base64: The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
               base64-encoded.
        :param pulumi.Input[str] service_name: The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceState.__new__(_ServiceState)

        __props__.__dict__["apis"] = apis
        __props__.__dict__["config_id"] = config_id
        __props__.__dict__["dns_address"] = dns_address
        __props__.__dict__["endpoints"] = endpoints
        __props__.__dict__["grpc_config"] = grpc_config
        __props__.__dict__["openapi_config"] = openapi_config
        __props__.__dict__["project"] = project
        __props__.__dict__["protoc_output_base64"] = protoc_output_base64
        __props__.__dict__["service_name"] = service_name
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def apis(self) -> pulumi.Output[Sequence['outputs.ServiceApi']]:
        """
        A list of API objects.
        """
        return pulumi.get(self, "apis")

    @property
    @pulumi.getter(name="configId")
    def config_id(self) -> pulumi.Output[str]:
        """
        The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
        to compute engine instances as a tag.
        """
        return pulumi.get(self, "config_id")

    @property
    @pulumi.getter(name="dnsAddress")
    def dns_address(self) -> pulumi.Output[str]:
        """
        The address at which the service can be found - usually the same as the service name.
        """
        return pulumi.get(self, "dns_address")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Sequence['outputs.ServiceEndpoint']]:
        """
        A list of Endpoint objects.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="grpcConfig")
    def grpc_config(self) -> pulumi.Output[Optional[str]]:
        """
        The full text of the Service Config YAML file (Example located here). If provided, must also provide
        protoc_output_base64. open_api config must not be provided.
        """
        return pulumi.get(self, "grpc_config")

    @property
    @pulumi.getter(name="openapiConfig")
    def openapi_config(self) -> pulumi.Output[Optional[str]]:
        """
        The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
        protoc_output_base64 must be specified.
        """
        return pulumi.get(self, "openapi_config")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID that the service belongs to. If not provided, provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="protocOutputBase64")
    def protoc_output_base64(self) -> pulumi.Output[Optional[str]]:
        """
        The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
        base64-encoded.
        """
        return pulumi.get(self, "protoc_output_base64")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
        """
        return pulumi.get(self, "service_name")

