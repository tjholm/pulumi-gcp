# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApiArgs', 'Api']

@pulumi.input_type
class ApiArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 managed_service: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Api resource.
        :param pulumi.Input[str] api_id: Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] managed_service: Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
               If not specified, a new Service will automatically be created in the same project as this API.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "api_id", api_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if managed_service is not None:
            pulumi.set(__self__, "managed_service", managed_service)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

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
    @pulumi.getter(name="managedService")
    def managed_service(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
        If not specified, a new Service will automatically be created in the same project as this API.
        """
        return pulumi.get(self, "managed_service")

    @managed_service.setter
    def managed_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_service", value)

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
class _ApiState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 managed_service: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Api resources.
        :param pulumi.Input[str] api_id: Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] managed_service: Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
               If not specified, a new Service will automatically be created in the same project as this API.
        :param pulumi.Input[str] name: The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if managed_service is not None:
            pulumi.set(__self__, "managed_service", managed_service)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

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
    @pulumi.getter(name="managedService")
    def managed_service(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
        If not specified, a new Service will automatically be created in the same project as this API.
        """
        return pulumi.get(self, "managed_service")

    @managed_service.setter
    def managed_service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_service", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
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


class Api(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 managed_service: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A consumable API that can be used by multiple Gateways.

        To get more information about Api, see:

        * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/api-gateway/docs/quickstart)

        ## Example Usage
        ### Apigateway Api Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        api = gcp.apigateway.Api("api", api_id="api",
        opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Api can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigateway/api:Api default projects/{{project}}/locations/global/apis/{{api_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/api:Api default {{project}}/{{api_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/api:Api default {{api_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] managed_service: Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
               If not specified, a new Service will automatically be created in the same project as this API.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A consumable API that can be used by multiple Gateways.

        To get more information about Api, see:

        * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/api-gateway/docs/quickstart)

        ## Example Usage
        ### Apigateway Api Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        api = gcp.apigateway.Api("api", api_id="api",
        opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Api can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigateway/api:Api default projects/{{project}}/locations/global/apis/{{api_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/api:Api default {{project}}/{{api_id}}
        ```

        ```sh
         $ pulumi import gcp:apigateway/api:Api default {{api_id}}
        ```

        :param str resource_name: The name of the resource.
        :param ApiArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 managed_service: Optional[pulumi.Input[str]] = None,
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
            __props__ = ApiArgs.__new__(ApiArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["labels"] = labels
            __props__.__dict__["managed_service"] = managed_service
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
        super(Api, __self__).__init__(
            'gcp:apigateway/api:Api',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            managed_service: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Api':
        """
        Get an existing Api resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        :param pulumi.Input[str] create_time: Creation timestamp in RFC3339 text format.
        :param pulumi.Input[str] display_name: A user-visible name for the API.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Resource labels to represent user-provided metadata.
        :param pulumi.Input[str] managed_service: Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
               If not specified, a new Service will automatically be created in the same project as this API.
        :param pulumi.Input[str] name: The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiState.__new__(_ApiState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["labels"] = labels
        __props__.__dict__["managed_service"] = managed_service
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        return Api(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        Identifier to assign to the API. Must be unique within scope of the parent resource(project)
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        A user-visible name for the API.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="managedService")
    def managed_service(self) -> pulumi.Output[str]:
        """
        Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
        If not specified, a new Service will automatically be created in the same project as this API.
        """
        return pulumi.get(self, "managed_service")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
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

