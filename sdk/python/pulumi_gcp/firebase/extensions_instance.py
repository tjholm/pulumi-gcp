# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ExtensionsInstanceArgs', 'ExtensionsInstance']

@pulumi.input_type
class ExtensionsInstanceArgs:
    def __init__(__self__, *,
                 config: pulumi.Input['ExtensionsInstanceConfigArgs'],
                 instance_id: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ExtensionsInstance resource.
        :param pulumi.Input['ExtensionsInstanceConfigArgs'] config: The current Config of the Extension Instance.
               Structure is documented below.
        :param pulumi.Input[str] instance_id: The ID to use for the Extension Instance, which will become the final
               component of the instance's name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ExtensionsInstanceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            config=config,
            instance_id=instance_id,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             config: Optional[pulumi.Input['ExtensionsInstanceConfigArgs']] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if config is None:
            raise TypeError("Missing 'config' argument")
        if instance_id is None and 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if instance_id is None:
            raise TypeError("Missing 'instance_id' argument")

        _setter("config", config)
        _setter("instance_id", instance_id)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input['ExtensionsInstanceConfigArgs']:
        """
        The current Config of the Extension Instance.
        Structure is documented below.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input['ExtensionsInstanceConfigArgs']):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID to use for the Extension Instance, which will become the final
        component of the instance's name.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

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
class _ExtensionsInstanceState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input['ExtensionsInstanceConfigArgs']] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 error_statuses: Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceErrorStatusArgs']]]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 last_operation_name: Optional[pulumi.Input[str]] = None,
                 last_operation_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 runtime_datas: Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceRuntimeDataArgs']]]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExtensionsInstance resources.
        :param pulumi.Input['ExtensionsInstanceConfigArgs'] config: The current Config of the Extension Instance.
               Structure is documented below.
        :param pulumi.Input[str] create_time: (Output)
               The time at which the Extension Instance Config was created.
        :param pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceErrorStatusArgs']]] error_statuses: If this Instance has `state: ERRORED`, the error messages
               will be found here.
               Structure is documented below.
        :param pulumi.Input[str] etag: A weak etag that is computed by the server based on other configuration
               values and may be sent on update and delete requests to ensure the
               client has an up-to-date value before proceeding.
        :param pulumi.Input[str] instance_id: The ID to use for the Extension Instance, which will become the final
               component of the instance's name.
        :param pulumi.Input[str] last_operation_name: The name of the last operation that acted on this Extension
               Instance
        :param pulumi.Input[str] last_operation_type: The type of the last operation that acted on the Extension Instance.
        :param pulumi.Input[str] name: (Output)
               The unique identifier for this configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceRuntimeDataArgs']]] runtime_datas: Data set by the extension instance at runtime.
               Structure is documented below.
        :param pulumi.Input[str] service_account_email: The email of the service account to be used at runtime by compute resources
               created for the operation of the Extension instance.
        :param pulumi.Input[str] state: The processing state of the extension instance.
        :param pulumi.Input[str] update_time: The time at which the Extension Instance was updated.
        """
        _ExtensionsInstanceState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            config=config,
            create_time=create_time,
            error_statuses=error_statuses,
            etag=etag,
            instance_id=instance_id,
            last_operation_name=last_operation_name,
            last_operation_type=last_operation_type,
            name=name,
            project=project,
            runtime_datas=runtime_datas,
            service_account_email=service_account_email,
            state=state,
            update_time=update_time,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             config: Optional[pulumi.Input['ExtensionsInstanceConfigArgs']] = None,
             create_time: Optional[pulumi.Input[str]] = None,
             error_statuses: Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceErrorStatusArgs']]]] = None,
             etag: Optional[pulumi.Input[str]] = None,
             instance_id: Optional[pulumi.Input[str]] = None,
             last_operation_name: Optional[pulumi.Input[str]] = None,
             last_operation_type: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             runtime_datas: Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceRuntimeDataArgs']]]] = None,
             service_account_email: Optional[pulumi.Input[str]] = None,
             state: Optional[pulumi.Input[str]] = None,
             update_time: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if create_time is None and 'createTime' in kwargs:
            create_time = kwargs['createTime']
        if error_statuses is None and 'errorStatuses' in kwargs:
            error_statuses = kwargs['errorStatuses']
        if instance_id is None and 'instanceId' in kwargs:
            instance_id = kwargs['instanceId']
        if last_operation_name is None and 'lastOperationName' in kwargs:
            last_operation_name = kwargs['lastOperationName']
        if last_operation_type is None and 'lastOperationType' in kwargs:
            last_operation_type = kwargs['lastOperationType']
        if runtime_datas is None and 'runtimeDatas' in kwargs:
            runtime_datas = kwargs['runtimeDatas']
        if service_account_email is None and 'serviceAccountEmail' in kwargs:
            service_account_email = kwargs['serviceAccountEmail']
        if update_time is None and 'updateTime' in kwargs:
            update_time = kwargs['updateTime']

        if config is not None:
            _setter("config", config)
        if create_time is not None:
            _setter("create_time", create_time)
        if error_statuses is not None:
            _setter("error_statuses", error_statuses)
        if etag is not None:
            _setter("etag", etag)
        if instance_id is not None:
            _setter("instance_id", instance_id)
        if last_operation_name is not None:
            _setter("last_operation_name", last_operation_name)
        if last_operation_type is not None:
            _setter("last_operation_type", last_operation_type)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if runtime_datas is not None:
            _setter("runtime_datas", runtime_datas)
        if service_account_email is not None:
            _setter("service_account_email", service_account_email)
        if state is not None:
            _setter("state", state)
        if update_time is not None:
            _setter("update_time", update_time)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['ExtensionsInstanceConfigArgs']]:
        """
        The current Config of the Extension Instance.
        Structure is documented below.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['ExtensionsInstanceConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        (Output)
        The time at which the Extension Instance Config was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="errorStatuses")
    def error_statuses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceErrorStatusArgs']]]]:
        """
        If this Instance has `state: ERRORED`, the error messages
        will be found here.
        Structure is documented below.
        """
        return pulumi.get(self, "error_statuses")

    @error_statuses.setter
    def error_statuses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceErrorStatusArgs']]]]):
        pulumi.set(self, "error_statuses", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        A weak etag that is computed by the server based on other configuration
        values and may be sent on update and delete requests to ensure the
        client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for the Extension Instance, which will become the final
        component of the instance's name.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="lastOperationName")
    def last_operation_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the last operation that acted on this Extension
        Instance
        """
        return pulumi.get(self, "last_operation_name")

    @last_operation_name.setter
    def last_operation_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_operation_name", value)

    @property
    @pulumi.getter(name="lastOperationType")
    def last_operation_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the last operation that acted on the Extension Instance.
        """
        return pulumi.get(self, "last_operation_type")

    @last_operation_type.setter
    def last_operation_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_operation_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        (Output)
        The unique identifier for this configuration.
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
    @pulumi.getter(name="runtimeDatas")
    def runtime_datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceRuntimeDataArgs']]]]:
        """
        Data set by the extension instance at runtime.
        Structure is documented below.
        """
        return pulumi.get(self, "runtime_datas")

    @runtime_datas.setter
    def runtime_datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExtensionsInstanceRuntimeDataArgs']]]]):
        pulumi.set(self, "runtime_datas", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        The email of the service account to be used at runtime by compute resources
        created for the operation of the Extension instance.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The processing state of the extension instance.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the Extension Instance was updated.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class ExtensionsInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['ExtensionsInstanceConfigArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Firebase Extentions Instance Resize Image

        ```python
        import pulumi
        import pulumi_gcp as gcp

        images = gcp.storage.Bucket("images",
            project="my-project-name",
            location="US",
            uniform_bucket_level_access=True,
            force_destroy=True,
            opts=pulumi.ResourceOptions(provider=google_beta))
        resize_image = gcp.firebase.ExtensionsInstance("resizeImage",
            project="my-project-name",
            instance_id="storage-resize-images",
            config=gcp.firebase.ExtensionsInstanceConfigArgs(
                extension_ref="firebase/storage-resize-images",
                extension_version="0.1.37",
                params={
                    "DELETE_ORIGINAL_FILE": "false",
                    "MAKE_PUBLIC": "false",
                    "IMAGE_TYPE": "false",
                    "IS_ANIMATED": "true",
                    "FUNCTION_MEMORY": "1024",
                    "DO_BACKFILL": "false",
                    "IMG_SIZES": "200x200",
                    "IMG_BUCKET": images.name,
                    "LOCATION": "",
                },
                system_params={
                    "firebaseextensions.v1beta.function/maxInstances": "3000",
                    "firebaseextensions.v1beta.function/memory": "256",
                    "firebaseextensions.v1beta.function/minInstances": "0",
                    "firebaseextensions.v1beta.function/vpcConnectorEgressSettings": "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED",
                },
                allowed_event_types=["firebase.extensions.storage-resize-images.v1.complete"],
                eventarc_channel="projects/my-project-name/locations//channels/firebase",
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:firebase/extensionsInstance:ExtensionsInstance default projects/{{project}}/instances/{{instance_id}}
        ```

        ```sh
         $ pulumi import gcp:firebase/extensionsInstance:ExtensionsInstance default {{project}}/{{instance_id}}
        ```

        ```sh
         $ pulumi import gcp:firebase/extensionsInstance:ExtensionsInstance default {{instance_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ExtensionsInstanceConfigArgs']] config: The current Config of the Extension Instance.
               Structure is documented below.
        :param pulumi.Input[str] instance_id: The ID to use for the Extension Instance, which will become the final
               component of the instance's name.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExtensionsInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Firebase Extentions Instance Resize Image

        ```python
        import pulumi
        import pulumi_gcp as gcp

        images = gcp.storage.Bucket("images",
            project="my-project-name",
            location="US",
            uniform_bucket_level_access=True,
            force_destroy=True,
            opts=pulumi.ResourceOptions(provider=google_beta))
        resize_image = gcp.firebase.ExtensionsInstance("resizeImage",
            project="my-project-name",
            instance_id="storage-resize-images",
            config=gcp.firebase.ExtensionsInstanceConfigArgs(
                extension_ref="firebase/storage-resize-images",
                extension_version="0.1.37",
                params={
                    "DELETE_ORIGINAL_FILE": "false",
                    "MAKE_PUBLIC": "false",
                    "IMAGE_TYPE": "false",
                    "IS_ANIMATED": "true",
                    "FUNCTION_MEMORY": "1024",
                    "DO_BACKFILL": "false",
                    "IMG_SIZES": "200x200",
                    "IMG_BUCKET": images.name,
                    "LOCATION": "",
                },
                system_params={
                    "firebaseextensions.v1beta.function/maxInstances": "3000",
                    "firebaseextensions.v1beta.function/memory": "256",
                    "firebaseextensions.v1beta.function/minInstances": "0",
                    "firebaseextensions.v1beta.function/vpcConnectorEgressSettings": "VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED",
                },
                allowed_event_types=["firebase.extensions.storage-resize-images.v1.complete"],
                eventarc_channel="projects/my-project-name/locations//channels/firebase",
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:firebase/extensionsInstance:ExtensionsInstance default projects/{{project}}/instances/{{instance_id}}
        ```

        ```sh
         $ pulumi import gcp:firebase/extensionsInstance:ExtensionsInstance default {{project}}/{{instance_id}}
        ```

        ```sh
         $ pulumi import gcp:firebase/extensionsInstance:ExtensionsInstance default {{instance_id}}
        ```

        :param str resource_name: The name of the resource.
        :param ExtensionsInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExtensionsInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ExtensionsInstanceArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['ExtensionsInstanceConfigArgs']]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExtensionsInstanceArgs.__new__(ExtensionsInstanceArgs)

            config = _utilities.configure(config, ExtensionsInstanceConfigArgs, True)
            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["error_statuses"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["last_operation_name"] = None
            __props__.__dict__["last_operation_type"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["runtime_datas"] = None
            __props__.__dict__["service_account_email"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        super(ExtensionsInstance, __self__).__init__(
            'gcp:firebase/extensionsInstance:ExtensionsInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[pulumi.InputType['ExtensionsInstanceConfigArgs']]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            error_statuses: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExtensionsInstanceErrorStatusArgs']]]]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            last_operation_name: Optional[pulumi.Input[str]] = None,
            last_operation_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            runtime_datas: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExtensionsInstanceRuntimeDataArgs']]]]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'ExtensionsInstance':
        """
        Get an existing ExtensionsInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ExtensionsInstanceConfigArgs']] config: The current Config of the Extension Instance.
               Structure is documented below.
        :param pulumi.Input[str] create_time: (Output)
               The time at which the Extension Instance Config was created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExtensionsInstanceErrorStatusArgs']]]] error_statuses: If this Instance has `state: ERRORED`, the error messages
               will be found here.
               Structure is documented below.
        :param pulumi.Input[str] etag: A weak etag that is computed by the server based on other configuration
               values and may be sent on update and delete requests to ensure the
               client has an up-to-date value before proceeding.
        :param pulumi.Input[str] instance_id: The ID to use for the Extension Instance, which will become the final
               component of the instance's name.
        :param pulumi.Input[str] last_operation_name: The name of the last operation that acted on this Extension
               Instance
        :param pulumi.Input[str] last_operation_type: The type of the last operation that acted on the Extension Instance.
        :param pulumi.Input[str] name: (Output)
               The unique identifier for this configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExtensionsInstanceRuntimeDataArgs']]]] runtime_datas: Data set by the extension instance at runtime.
               Structure is documented below.
        :param pulumi.Input[str] service_account_email: The email of the service account to be used at runtime by compute resources
               created for the operation of the Extension instance.
        :param pulumi.Input[str] state: The processing state of the extension instance.
        :param pulumi.Input[str] update_time: The time at which the Extension Instance was updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExtensionsInstanceState.__new__(_ExtensionsInstanceState)

        __props__.__dict__["config"] = config
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["error_statuses"] = error_statuses
        __props__.__dict__["etag"] = etag
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["last_operation_name"] = last_operation_name
        __props__.__dict__["last_operation_type"] = last_operation_type
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["runtime_datas"] = runtime_datas
        __props__.__dict__["service_account_email"] = service_account_email
        __props__.__dict__["state"] = state
        __props__.__dict__["update_time"] = update_time
        return ExtensionsInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.ExtensionsInstanceConfig']:
        """
        The current Config of the Extension Instance.
        Structure is documented below.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        (Output)
        The time at which the Extension Instance Config was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="errorStatuses")
    def error_statuses(self) -> pulumi.Output[Sequence['outputs.ExtensionsInstanceErrorStatus']]:
        """
        If this Instance has `state: ERRORED`, the error messages
        will be found here.
        Structure is documented below.
        """
        return pulumi.get(self, "error_statuses")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A weak etag that is computed by the server based on other configuration
        values and may be sent on update and delete requests to ensure the
        client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID to use for the Extension Instance, which will become the final
        component of the instance's name.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="lastOperationName")
    def last_operation_name(self) -> pulumi.Output[str]:
        """
        The name of the last operation that acted on this Extension
        Instance
        """
        return pulumi.get(self, "last_operation_name")

    @property
    @pulumi.getter(name="lastOperationType")
    def last_operation_type(self) -> pulumi.Output[str]:
        """
        The type of the last operation that acted on the Extension Instance.
        """
        return pulumi.get(self, "last_operation_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        (Output)
        The unique identifier for this configuration.
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
    @pulumi.getter(name="runtimeDatas")
    def runtime_datas(self) -> pulumi.Output[Sequence['outputs.ExtensionsInstanceRuntimeData']]:
        """
        Data set by the extension instance at runtime.
        Structure is documented below.
        """
        return pulumi.get(self, "runtime_datas")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        The email of the service account to be used at runtime by compute resources
        created for the operation of the Extension instance.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The processing state of the extension instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time at which the Extension Instance was updated.
        """
        return pulumi.get(self, "update_time")

