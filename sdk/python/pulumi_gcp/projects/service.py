# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceArgs', 'Service']

@pulumi.input_type
class ServiceArgs:
    def __init__(__self__, *,
                 service: pulumi.Input[str],
                 disable_dependent_services: Optional[pulumi.Input[bool]] = None,
                 disable_on_destroy: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Service resource.
        :param pulumi.Input[str] service: The service to enable.
        :param pulumi.Input[bool] disable_dependent_services: If `true`, services that are enabled
               and which depend on this service should also be disabled when this service is
               destroyed. If `false` or unset, an error will be generated if any enabled
               services depend on this service when destroying it.
        :param pulumi.Input[bool] disable_on_destroy: If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        :param pulumi.Input[str] project: The project ID. If not provided, the provider project
               is used.
        """
        pulumi.set(__self__, "service", service)
        if disable_dependent_services is not None:
            pulumi.set(__self__, "disable_dependent_services", disable_dependent_services)
        if disable_on_destroy is not None:
            pulumi.set(__self__, "disable_on_destroy", disable_on_destroy)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def service(self) -> pulumi.Input[str]:
        """
        The service to enable.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: pulumi.Input[str]):
        pulumi.set(self, "service", value)

    @property
    @pulumi.getter(name="disableDependentServices")
    def disable_dependent_services(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, services that are enabled
        and which depend on this service should also be disabled when this service is
        destroyed. If `false` or unset, an error will be generated if any enabled
        services depend on this service when destroying it.
        """
        return pulumi.get(self, "disable_dependent_services")

    @disable_dependent_services.setter
    def disable_dependent_services(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_dependent_services", value)

    @property
    @pulumi.getter(name="disableOnDestroy")
    def disable_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        """
        return pulumi.get(self, "disable_on_destroy")

    @disable_on_destroy.setter
    def disable_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_on_destroy", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID. If not provided, the provider project
        is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ServiceState:
    def __init__(__self__, *,
                 disable_dependent_services: Optional[pulumi.Input[bool]] = None,
                 disable_on_destroy: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Service resources.
        :param pulumi.Input[bool] disable_dependent_services: If `true`, services that are enabled
               and which depend on this service should also be disabled when this service is
               destroyed. If `false` or unset, an error will be generated if any enabled
               services depend on this service when destroying it.
        :param pulumi.Input[bool] disable_on_destroy: If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        :param pulumi.Input[str] project: The project ID. If not provided, the provider project
               is used.
        :param pulumi.Input[str] service: The service to enable.
        """
        if disable_dependent_services is not None:
            pulumi.set(__self__, "disable_dependent_services", disable_dependent_services)
        if disable_on_destroy is not None:
            pulumi.set(__self__, "disable_on_destroy", disable_on_destroy)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service is not None:
            pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="disableDependentServices")
    def disable_dependent_services(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, services that are enabled
        and which depend on this service should also be disabled when this service is
        destroyed. If `false` or unset, an error will be generated if any enabled
        services depend on this service when destroying it.
        """
        return pulumi.get(self, "disable_dependent_services")

    @disable_dependent_services.setter
    def disable_dependent_services(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_dependent_services", value)

    @property
    @pulumi.getter(name="disableOnDestroy")
    def disable_on_destroy(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        """
        return pulumi.get(self, "disable_on_destroy")

    @disable_on_destroy.setter
    def disable_on_destroy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_on_destroy", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID. If not provided, the provider project
        is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def service(self) -> Optional[pulumi.Input[str]]:
        """
        The service to enable.
        """
        return pulumi.get(self, "service")

    @service.setter
    def service(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service", value)


class Service(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_dependent_services: Optional[pulumi.Input[bool]] = None,
                 disable_on_destroy: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows management of a single API service for a Google Cloud Platform project.

        For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
        or run `gcloud services list --available`.

        This resource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
        to use.

        To get more information about `projects.Service`, see:

        * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
        * How-to Guides
            * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.Service("project",
            disable_dependent_services=True,
            project="your-project-id",
            service="iam.googleapis.com")
        ```

        ## Import

        Project services can be imported using the `project_id` and `service`, e.g.

        ```sh
         $ pulumi import gcp:projects/service:Service my_project your-project-id/iam.googleapis.com
        ```

         Note that unlike other resources that fail if they already exist, `terraform apply` can be successfully used to verify already enabled services. This means that when importing existing resources into Terraform, you can either import the `google_project_service` resources or treat them as new infrastructure and run `terraform apply` to add them to state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_dependent_services: If `true`, services that are enabled
               and which depend on this service should also be disabled when this service is
               destroyed. If `false` or unset, an error will be generated if any enabled
               services depend on this service when destroying it.
        :param pulumi.Input[bool] disable_on_destroy: If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        :param pulumi.Input[str] project: The project ID. If not provided, the provider project
               is used.
        :param pulumi.Input[str] service: The service to enable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows management of a single API service for a Google Cloud Platform project.

        For a list of services available, visit the [API library page](https://console.cloud.google.com/apis/library)
        or run `gcloud services list --available`.

        This resource requires the [Service Usage API](https://console.cloud.google.com/apis/library/serviceusage.googleapis.com)
        to use.

        To get more information about `projects.Service`, see:

        * [API documentation](https://cloud.google.com/service-usage/docs/reference/rest/v1/services)
        * How-to Guides
            * [Enabling and Disabling Services](https://cloud.google.com/service-usage/docs/enable-disable)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.projects.Service("project",
            disable_dependent_services=True,
            project="your-project-id",
            service="iam.googleapis.com")
        ```

        ## Import

        Project services can be imported using the `project_id` and `service`, e.g.

        ```sh
         $ pulumi import gcp:projects/service:Service my_project your-project-id/iam.googleapis.com
        ```

         Note that unlike other resources that fail if they already exist, `terraform apply` can be successfully used to verify already enabled services. This means that when importing existing resources into Terraform, you can either import the `google_project_service` resources or treat them as new infrastructure and run `terraform apply` to add them to state.

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
                 disable_dependent_services: Optional[pulumi.Input[bool]] = None,
                 disable_on_destroy: Optional[pulumi.Input[bool]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceArgs.__new__(ServiceArgs)

            __props__.__dict__["disable_dependent_services"] = disable_dependent_services
            __props__.__dict__["disable_on_destroy"] = disable_on_destroy
            __props__.__dict__["project"] = project
            if service is None and not opts.urn:
                raise TypeError("Missing required property 'service'")
            __props__.__dict__["service"] = service
        super(Service, __self__).__init__(
            'gcp:projects/service:Service',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_dependent_services: Optional[pulumi.Input[bool]] = None,
            disable_on_destroy: Optional[pulumi.Input[bool]] = None,
            project: Optional[pulumi.Input[str]] = None,
            service: Optional[pulumi.Input[str]] = None) -> 'Service':
        """
        Get an existing Service resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_dependent_services: If `true`, services that are enabled
               and which depend on this service should also be disabled when this service is
               destroyed. If `false` or unset, an error will be generated if any enabled
               services depend on this service when destroying it.
        :param pulumi.Input[bool] disable_on_destroy: If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        :param pulumi.Input[str] project: The project ID. If not provided, the provider project
               is used.
        :param pulumi.Input[str] service: The service to enable.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceState.__new__(_ServiceState)

        __props__.__dict__["disable_dependent_services"] = disable_dependent_services
        __props__.__dict__["disable_on_destroy"] = disable_on_destroy
        __props__.__dict__["project"] = project
        __props__.__dict__["service"] = service
        return Service(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableDependentServices")
    def disable_dependent_services(self) -> pulumi.Output[Optional[bool]]:
        """
        If `true`, services that are enabled
        and which depend on this service should also be disabled when this service is
        destroyed. If `false` or unset, an error will be generated if any enabled
        services depend on this service when destroying it.
        """
        return pulumi.get(self, "disable_dependent_services")

    @property
    @pulumi.getter(name="disableOnDestroy")
    def disable_on_destroy(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, disable the service when the resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
        """
        return pulumi.get(self, "disable_on_destroy")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project ID. If not provided, the provider project
        is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def service(self) -> pulumi.Output[str]:
        """
        The service to enable.
        """
        return pulumi.get(self, "service")

