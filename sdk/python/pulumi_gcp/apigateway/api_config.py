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

__all__ = ['ApiConfigArgs', 'ApiConfig']

@pulumi.input_type
class ApiConfigArgs:
    def __init__(__self__, *,
                 api: pulumi.Input[str],
                 openapi_documents: pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]],
                 api_config_id: Optional[pulumi.Input[str]] = None,
                 api_config_id_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_config: Optional[pulumi.Input['ApiConfigGatewayConfigArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiConfig resource.
        :param pulumi.Input[str] api: The API to attach the config to.
        :param pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]] openapi_documents: An OpenAPI Specification Document describing an API.
               Structure is documented below.
        :param pulumi.Input[str] api_config_id: Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        :param pulumi.Input[str] api_config_id_prefix: Creates a unique name beginning with the
               specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input['ApiConfigGatewayConfigArgs'] gateway_config: Immutable. Gateway specific configuration.
               If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "api", api)
        pulumi.set(__self__, "openapi_documents", openapi_documents)
        if api_config_id is not None:
            pulumi.set(__self__, "api_config_id", api_config_id)
        if api_config_id_prefix is not None:
            pulumi.set(__self__, "api_config_id_prefix", api_config_id_prefix)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if gateway_config is not None:
            pulumi.set(__self__, "gateway_config", gateway_config)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def api(self) -> pulumi.Input[str]:
        """
        The API to attach the config to.
        """
        return pulumi.get(self, "api")

    @api.setter
    def api(self, value: pulumi.Input[str]):
        pulumi.set(self, "api", value)

    @property
    @pulumi.getter(name="openapiDocuments")
    def openapi_documents(self) -> pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]]:
        """
        An OpenAPI Specification Document describing an API.
        Structure is documented below.
        """
        return pulumi.get(self, "openapi_documents")

    @openapi_documents.setter
    def openapi_documents(self, value: pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]]):
        pulumi.set(self, "openapi_documents", value)

    @property
    @pulumi.getter(name="apiConfigId")
    def api_config_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        """
        return pulumi.get(self, "api_config_id")

    @api_config_id.setter
    def api_config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_config_id", value)

    @property
    @pulumi.getter(name="apiConfigIdPrefix")
    def api_config_id_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the
        specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        """
        return pulumi.get(self, "api_config_id_prefix")

    @api_config_id_prefix.setter
    def api_config_id_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_config_id_prefix", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-visible name for the API.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="gatewayConfig")
    def gateway_config(self) -> Optional[pulumi.Input['ApiConfigGatewayConfigArgs']]:
        """
        Immutable. Gateway specific configuration.
        If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        Structure is documented below.
        """
        return pulumi.get(self, "gateway_config")

    @gateway_config.setter
    def gateway_config(self, value: Optional[pulumi.Input['ApiConfigGatewayConfigArgs']]):
        pulumi.set(self, "gateway_config", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user-provided metadata.
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
class _ApiConfigState:
    def __init__(__self__, *,
                 api: Optional[pulumi.Input[str]] = None,
                 api_config_id: Optional[pulumi.Input[str]] = None,
                 api_config_id_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_config: Optional[pulumi.Input['ApiConfigGatewayConfigArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_config_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiConfig resources.
        :param pulumi.Input[str] api: The API to attach the config to.
        :param pulumi.Input[str] api_config_id: Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        :param pulumi.Input[str] api_config_id_prefix: Creates a unique name beginning with the
               specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input['ApiConfigGatewayConfigArgs'] gateway_config: Immutable. Gateway specific configuration.
               If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] name: The resource name of the API Config.
        :param pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]] openapi_documents: An OpenAPI Specification Document describing an API.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_config_id: The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        """
        if api is not None:
            pulumi.set(__self__, "api", api)
        if api_config_id is not None:
            pulumi.set(__self__, "api_config_id", api_config_id)
        if api_config_id_prefix is not None:
            pulumi.set(__self__, "api_config_id_prefix", api_config_id_prefix)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if gateway_config is not None:
            pulumi.set(__self__, "gateway_config", gateway_config)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if openapi_documents is not None:
            pulumi.set(__self__, "openapi_documents", openapi_documents)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service_config_id is not None:
            pulumi.set(__self__, "service_config_id", service_config_id)

    @property
    @pulumi.getter
    def api(self) -> Optional[pulumi.Input[str]]:
        """
        The API to attach the config to.
        """
        return pulumi.get(self, "api")

    @api.setter
    def api(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api", value)

    @property
    @pulumi.getter(name="apiConfigId")
    def api_config_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        """
        return pulumi.get(self, "api_config_id")

    @api_config_id.setter
    def api_config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_config_id", value)

    @property
    @pulumi.getter(name="apiConfigIdPrefix")
    def api_config_id_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the
        specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        """
        return pulumi.get(self, "api_config_id_prefix")

    @api_config_id_prefix.setter
    def api_config_id_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_config_id_prefix", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-visible name for the API.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="gatewayConfig")
    def gateway_config(self) -> Optional[pulumi.Input['ApiConfigGatewayConfigArgs']]:
        """
        Immutable. Gateway specific configuration.
        If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        Structure is documented below.
        """
        return pulumi.get(self, "gateway_config")

    @gateway_config.setter
    def gateway_config(self, value: Optional[pulumi.Input['ApiConfigGatewayConfigArgs']]):
        pulumi.set(self, "gateway_config", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the API Config.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="openapiDocuments")
    def openapi_documents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]]]:
        """
        An OpenAPI Specification Document describing an API.
        Structure is documented below.
        """
        return pulumi.get(self, "openapi_documents")

    @openapi_documents.setter
    def openapi_documents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApiConfigOpenapiDocumentArgs']]]]):
        pulumi.set(self, "openapi_documents", value)

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
    @pulumi.getter(name="serviceConfigId")
    def service_config_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        """
        return pulumi.get(self, "service_config_id")

    @service_config_id.setter
    def service_config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_config_id", value)


class ApiConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api: Optional[pulumi.Input[str]] = None,
                 api_config_id: Optional[pulumi.Input[str]] = None,
                 api_config_id_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_config: Optional[pulumi.Input[pulumi.InputType['ApiConfigGatewayConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiConfigOpenapiDocumentArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An API Configuration is an association of an API Controller Config and a Gateway Config

        To get more information about ApiConfig, see:

        * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis.configs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/api-gateway/docs/creating-api-config)

        ## Example Usage

        ## Import

        ApiConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigateway/apiConfig:ApiConfig default projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{project}}/{{api}}/{{api_config_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{api}}/{{api_config_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api: The API to attach the config to.
        :param pulumi.Input[str] api_config_id: Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        :param pulumi.Input[str] api_config_id_prefix: Creates a unique name beginning with the
               specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input[pulumi.InputType['ApiConfigGatewayConfigArgs']] gateway_config: Immutable. Gateway specific configuration.
               If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiConfigOpenapiDocumentArgs']]]] openapi_documents: An OpenAPI Specification Document describing an API.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An API Configuration is an association of an API Controller Config and a Gateway Config

        To get more information about ApiConfig, see:

        * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis.configs)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/api-gateway/docs/creating-api-config)

        ## Example Usage

        ## Import

        ApiConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigateway/apiConfig:ApiConfig default projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{project}}/{{api}}/{{api_config_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{api}}/{{api_config_id}}
        ```

        :param str resource_name: The name of the resource.
        :param ApiConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api: Optional[pulumi.Input[str]] = None,
                 api_config_id: Optional[pulumi.Input[str]] = None,
                 api_config_id_prefix: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 gateway_config: Optional[pulumi.Input[pulumi.InputType['ApiConfigGatewayConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiConfigOpenapiDocumentArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApiConfigArgs.__new__(ApiConfigArgs)

            if api is None and not opts.urn:
                raise TypeError("Missing required property 'api'")
            __props__.__dict__["api"] = api
            __props__.__dict__["api_config_id"] = api_config_id
            __props__.__dict__["api_config_id_prefix"] = api_config_id_prefix
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["gateway_config"] = gateway_config
            __props__.__dict__["labels"] = labels
            if openapi_documents is None and not opts.urn:
                raise TypeError("Missing required property 'openapi_documents'")
            __props__.__dict__["openapi_documents"] = openapi_documents
            __props__.__dict__["project"] = project
            __props__.__dict__["name"] = None
            __props__.__dict__["service_config_id"] = None
        super(ApiConfig, __self__).__init__(
            'gcp:apigateway/apiConfig:ApiConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api: Optional[pulumi.Input[str]] = None,
            api_config_id: Optional[pulumi.Input[str]] = None,
            api_config_id_prefix: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            gateway_config: Optional[pulumi.Input[pulumi.InputType['ApiConfigGatewayConfigArgs']]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            openapi_documents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiConfigOpenapiDocumentArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            service_config_id: Optional[pulumi.Input[str]] = None) -> 'ApiConfig':
        """
        Get an existing ApiConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api: The API to attach the config to.
        :param pulumi.Input[str] api_config_id: Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        :param pulumi.Input[str] api_config_id_prefix: Creates a unique name beginning with the
               specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input[pulumi.InputType['ApiConfigGatewayConfigArgs']] gateway_config: Immutable. Gateway specific configuration.
               If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] name: The resource name of the API Config.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiConfigOpenapiDocumentArgs']]]] openapi_documents: An OpenAPI Specification Document describing an API.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_config_id: The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiConfigState.__new__(_ApiConfigState)

        __props__.__dict__["api"] = api
        __props__.__dict__["api_config_id"] = api_config_id
        __props__.__dict__["api_config_id_prefix"] = api_config_id_prefix
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["gateway_config"] = gateway_config
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["openapi_documents"] = openapi_documents
        __props__.__dict__["project"] = project
        __props__.__dict__["service_config_id"] = service_config_id
        return ApiConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def api(self) -> pulumi.Output[str]:
        """
        The API to attach the config to.
        """
        return pulumi.get(self, "api")

    @property
    @pulumi.getter(name="apiConfigId")
    def api_config_id(self) -> pulumi.Output[str]:
        """
        Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
        """
        return pulumi.get(self, "api_config_id")

    @property
    @pulumi.getter(name="apiConfigIdPrefix")
    def api_config_id_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique name beginning with the
        specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
        """
        return pulumi.get(self, "api_config_id_prefix")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A user-visible name for the API.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="gatewayConfig")
    def gateway_config(self) -> pulumi.Output[Optional['outputs.ApiConfigGatewayConfig']]:
        """
        Immutable. Gateway specific configuration.
        If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
        Structure is documented below.
        """
        return pulumi.get(self, "gateway_config")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the API Config.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openapiDocuments")
    def openapi_documents(self) -> pulumi.Output[Sequence['outputs.ApiConfigOpenapiDocument']]:
        """
        An OpenAPI Specification Document describing an API.
        Structure is documented below.
        """
        return pulumi.get(self, "openapi_documents")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serviceConfigId")
    def service_config_id(self) -> pulumi.Output[str]:
        """
        The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
        """
        return pulumi.get(self, "service_config_id")

