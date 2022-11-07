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

__all__ = ['GenericServiceArgs', 'GenericService']

@pulumi.input_type
class GenericServiceArgs:
    def __init__(__self__, *,
                 service_id: pulumi.Input[str],
                 basic_service: Optional[pulumi.Input['GenericServiceBasicServiceArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a GenericService resource.
        :param pulumi.Input[str] service_id: An optional service ID to use. If not given, the server will generate a
               service ID.
        :param pulumi.Input['GenericServiceBasicServiceArgs'] basic_service: A well-known service type, defined by its service type and service labels.
               Valid values are described at
               https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
               Structure is documented below.
        :param pulumi.Input[str] display_name: Name used for UI elements listing this Service.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: Labels which have been used to annotate the service. Label keys must start
               with a letter. Label keys and values may contain lowercase letters,
               numbers, underscores, and dashes. Label keys and values have a maximum
               length of 63 characters, and must be less than 128 bytes in size. Up to 64
               label entries may be stored. For labels which do not have a semantic value,
               the empty string may be supplied for the label value.
        """
        pulumi.set(__self__, "service_id", service_id)
        if basic_service is not None:
            pulumi.set(__self__, "basic_service", basic_service)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if user_labels is not None:
            pulumi.set(__self__, "user_labels", user_labels)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        An optional service ID to use. If not given, the server will generate a
        service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="basicService")
    def basic_service(self) -> Optional[pulumi.Input['GenericServiceBasicServiceArgs']]:
        """
        A well-known service type, defined by its service type and service labels.
        Valid values are described at
        https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
        Structure is documented below.
        """
        return pulumi.get(self, "basic_service")

    @basic_service.setter
    def basic_service(self, value: Optional[pulumi.Input['GenericServiceBasicServiceArgs']]):
        pulumi.set(self, "basic_service", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name used for UI elements listing this Service.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

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
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels which have been used to annotate the service. Label keys must start
        with a letter. Label keys and values may contain lowercase letters,
        numbers, underscores, and dashes. Label keys and values have a maximum
        length of 63 characters, and must be less than 128 bytes in size. Up to 64
        label entries may be stored. For labels which do not have a semantic value,
        the empty string may be supplied for the label value.
        """
        return pulumi.get(self, "user_labels")

    @user_labels.setter
    def user_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "user_labels", value)


@pulumi.input_type
class _GenericServiceState:
    def __init__(__self__, *,
                 basic_service: Optional[pulumi.Input['GenericServiceBasicServiceArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 telemetries: Optional[pulumi.Input[Sequence[pulumi.Input['GenericServiceTelemetryArgs']]]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering GenericService resources.
        :param pulumi.Input['GenericServiceBasicServiceArgs'] basic_service: A well-known service type, defined by its service type and service labels.
               Valid values are described at
               https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
               Structure is documented below.
        :param pulumi.Input[str] display_name: Name used for UI elements listing this Service.
        :param pulumi.Input[str] name: The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_id: An optional service ID to use. If not given, the server will generate a
               service ID.
        :param pulumi.Input[Sequence[pulumi.Input['GenericServiceTelemetryArgs']]] telemetries: Configuration for how to query telemetry on a Service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: Labels which have been used to annotate the service. Label keys must start
               with a letter. Label keys and values may contain lowercase letters,
               numbers, underscores, and dashes. Label keys and values have a maximum
               length of 63 characters, and must be less than 128 bytes in size. Up to 64
               label entries may be stored. For labels which do not have a semantic value,
               the empty string may be supplied for the label value.
        """
        if basic_service is not None:
            pulumi.set(__self__, "basic_service", basic_service)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if telemetries is not None:
            pulumi.set(__self__, "telemetries", telemetries)
        if user_labels is not None:
            pulumi.set(__self__, "user_labels", user_labels)

    @property
    @pulumi.getter(name="basicService")
    def basic_service(self) -> Optional[pulumi.Input['GenericServiceBasicServiceArgs']]:
        """
        A well-known service type, defined by its service type and service labels.
        Valid values are described at
        https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
        Structure is documented below.
        """
        return pulumi.get(self, "basic_service")

    @basic_service.setter
    def basic_service(self, value: Optional[pulumi.Input['GenericServiceBasicServiceArgs']]):
        pulumi.set(self, "basic_service", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name used for UI elements listing this Service.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
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
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        An optional service ID to use. If not given, the server will generate a
        service ID.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def telemetries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GenericServiceTelemetryArgs']]]]:
        """
        Configuration for how to query telemetry on a Service.
        """
        return pulumi.get(self, "telemetries")

    @telemetries.setter
    def telemetries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GenericServiceTelemetryArgs']]]]):
        pulumi.set(self, "telemetries", value)

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels which have been used to annotate the service. Label keys must start
        with a letter. Label keys and values may contain lowercase letters,
        numbers, underscores, and dashes. Label keys and values have a maximum
        length of 63 characters, and must be less than 128 bytes in size. Up to 64
        label entries may be stored. For labels which do not have a semantic value,
        the empty string may be supplied for the label value.
        """
        return pulumi.get(self, "user_labels")

    @user_labels.setter
    def user_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "user_labels", value)


class GenericService(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 basic_service: Optional[pulumi.Input[pulumi.InputType['GenericServiceBasicServiceArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        A Service is a discrete, autonomous, and network-accessible unit,
        designed to solve an individual concern (Wikipedia). In Cloud Monitoring,
        a Service acts as the root resource under which operational aspects of
        the service are accessible

        To get more information about GenericService, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
        * How-to Guides
            * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
            * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

        ## Example Usage
        ### Monitoring Service Example

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_service = gcp.monitoring.GenericService("myService",
            basic_service=gcp.monitoring.GenericServiceBasicServiceArgs(
                service_labels={
                    "moduleId": "another-module-id",
                },
                service_type="APP_ENGINE",
            ),
            display_name="My Service my-service",
            service_id="my-service",
            user_labels={
                "my_key": "my_value",
                "my_other_key": "my_other_value",
            })
        ```

        ## Import

        GenericService can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:monitoring/genericService:GenericService default projects/{{project}}/services/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:monitoring/genericService:GenericService default {{project}}/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:monitoring/genericService:GenericService default {{service_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GenericServiceBasicServiceArgs']] basic_service: A well-known service type, defined by its service type and service labels.
               Valid values are described at
               https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
               Structure is documented below.
        :param pulumi.Input[str] display_name: Name used for UI elements listing this Service.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_id: An optional service ID to use. If not given, the server will generate a
               service ID.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: Labels which have been used to annotate the service. Label keys must start
               with a letter. Label keys and values may contain lowercase letters,
               numbers, underscores, and dashes. Label keys and values have a maximum
               length of 63 characters, and must be less than 128 bytes in size. Up to 64
               label entries may be stored. For labels which do not have a semantic value,
               the empty string may be supplied for the label value.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GenericServiceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Service is a discrete, autonomous, and network-accessible unit,
        designed to solve an individual concern (Wikipedia). In Cloud Monitoring,
        a Service acts as the root resource under which operational aspects of
        the service are accessible

        To get more information about GenericService, see:

        * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
        * How-to Guides
            * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
            * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)

        ## Example Usage
        ### Monitoring Service Example

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_service = gcp.monitoring.GenericService("myService",
            basic_service=gcp.monitoring.GenericServiceBasicServiceArgs(
                service_labels={
                    "moduleId": "another-module-id",
                },
                service_type="APP_ENGINE",
            ),
            display_name="My Service my-service",
            service_id="my-service",
            user_labels={
                "my_key": "my_value",
                "my_other_key": "my_other_value",
            })
        ```

        ## Import

        GenericService can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:monitoring/genericService:GenericService default projects/{{project}}/services/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:monitoring/genericService:GenericService default {{project}}/{{service_id}}
        ```

        ```sh
         $ pulumi import gcp:monitoring/genericService:GenericService default {{service_id}}
        ```

        :param str resource_name: The name of the resource.
        :param GenericServiceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GenericServiceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 basic_service: Optional[pulumi.Input[pulumi.InputType['GenericServiceBasicServiceArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GenericServiceArgs.__new__(GenericServiceArgs)

            __props__.__dict__["basic_service"] = basic_service
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["project"] = project
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["user_labels"] = user_labels
            __props__.__dict__["name"] = None
            __props__.__dict__["telemetries"] = None
        super(GenericService, __self__).__init__(
            'gcp:monitoring/genericService:GenericService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            basic_service: Optional[pulumi.Input[pulumi.InputType['GenericServiceBasicServiceArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None,
            telemetries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GenericServiceTelemetryArgs']]]]] = None,
            user_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'GenericService':
        """
        Get an existing GenericService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GenericServiceBasicServiceArgs']] basic_service: A well-known service type, defined by its service type and service labels.
               Valid values are described at
               https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
               Structure is documented below.
        :param pulumi.Input[str] display_name: Name used for UI elements listing this Service.
        :param pulumi.Input[str] name: The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_id: An optional service ID to use. If not given, the server will generate a
               service ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GenericServiceTelemetryArgs']]]] telemetries: Configuration for how to query telemetry on a Service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] user_labels: Labels which have been used to annotate the service. Label keys must start
               with a letter. Label keys and values may contain lowercase letters,
               numbers, underscores, and dashes. Label keys and values have a maximum
               length of 63 characters, and must be less than 128 bytes in size. Up to 64
               label entries may be stored. For labels which do not have a semantic value,
               the empty string may be supplied for the label value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GenericServiceState.__new__(_GenericServiceState)

        __props__.__dict__["basic_service"] = basic_service
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["telemetries"] = telemetries
        __props__.__dict__["user_labels"] = user_labels
        return GenericService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="basicService")
    def basic_service(self) -> pulumi.Output[Optional['outputs.GenericServiceBasicService']]:
        """
        A well-known service type, defined by its service type and service labels.
        Valid values are described at
        https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
        Structure is documented below.
        """
        return pulumi.get(self, "basic_service")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name used for UI elements listing this Service.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The full resource name for this service. The syntax is: projects/[PROJECT_ID]/services/[SERVICE_ID].
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
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        An optional service ID to use. If not given, the server will generate a
        service ID.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def telemetries(self) -> pulumi.Output[Sequence['outputs.GenericServiceTelemetry']]:
        """
        Configuration for how to query telemetry on a Service.
        """
        return pulumi.get(self, "telemetries")

    @property
    @pulumi.getter(name="userLabels")
    def user_labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels which have been used to annotate the service. Label keys must start
        with a letter. Label keys and values may contain lowercase letters,
        numbers, underscores, and dashes. Label keys and values have a maximum
        length of 63 characters, and must be less than 128 bytes in size. Up to 64
        label entries may be stored. For labels which do not have a semantic value,
        the empty string may be supplied for the label value.
        """
        return pulumi.get(self, "user_labels")

