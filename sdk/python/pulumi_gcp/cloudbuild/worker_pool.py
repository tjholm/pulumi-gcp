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

__all__ = ['WorkerPoolArgs', 'WorkerPool']

@pulumi.input_type
class WorkerPoolArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input['WorkerPoolNetworkConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input['WorkerPoolWorkerConfigArgs']] = None):
        """
        The set of arguments for constructing a WorkerPool resource.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
               limitations.
        :param pulumi.Input[str] display_name: A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        :param pulumi.Input[str] name: User-defined name of the `WorkerPool`.
        :param pulumi.Input['WorkerPoolNetworkConfigArgs'] network_config: Network configuration for the `WorkerPool`. Structure is documented below.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input['WorkerPoolWorkerConfigArgs'] worker_config: Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        pulumi.set(__self__, "location", location)
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if worker_config is not None:
            pulumi.set(__self__, "worker_config", worker_config)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
        limitations.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined name of the `WorkerPool`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['WorkerPoolNetworkConfigArgs']]:
        """
        Network configuration for the `WorkerPool`. Structure is documented below.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['WorkerPoolNetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> Optional[pulumi.Input['WorkerPoolWorkerConfigArgs']]:
        """
        Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        return pulumi.get(self, "worker_config")

    @worker_config.setter
    def worker_config(self, value: Optional[pulumi.Input['WorkerPoolWorkerConfigArgs']]):
        pulumi.set(self, "worker_config", value)


@pulumi.input_type
class _WorkerPoolState:
    def __init__(__self__, *,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 delete_time: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input['WorkerPoolNetworkConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input['WorkerPoolWorkerConfigArgs']] = None):
        """
        Input properties used for looking up and filtering WorkerPool resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
               limitations.
        :param pulumi.Input[str] create_time: Output only. Time at which the request to create the `WorkerPool` was received.
        :param pulumi.Input[str] delete_time: Output only. Time at which the request to delete the `WorkerPool` was received.
        :param pulumi.Input[str] display_name: A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[str] name: User-defined name of the `WorkerPool`.
        :param pulumi.Input['WorkerPoolNetworkConfigArgs'] network_config: Network configuration for the `WorkerPool`. Structure is documented below.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] state: Output only. `WorkerPool` state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
        :param pulumi.Input[str] uid: Output only. A unique identifier for the `WorkerPool`.
        :param pulumi.Input[str] update_time: Output only. Time at which the request to update the `WorkerPool` was received.
        :param pulumi.Input['WorkerPoolWorkerConfigArgs'] worker_config: Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        if annotations is not None:
            pulumi.set(__self__, "annotations", annotations)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if delete_time is not None:
            pulumi.set(__self__, "delete_time", delete_time)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_config is not None:
            pulumi.set(__self__, "network_config", network_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if worker_config is not None:
            pulumi.set(__self__, "worker_config", worker_config)

    @property
    @pulumi.getter
    def annotations(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
        limitations.
        """
        return pulumi.get(self, "annotations")

    @annotations.setter
    def annotations(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "annotations", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Time at which the request to create the `WorkerPool` was received.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Time at which the request to delete the `WorkerPool` was received.
        """
        return pulumi.get(self, "delete_time")

    @delete_time.setter
    def delete_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_time", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        User-defined name of the `WorkerPool`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> Optional[pulumi.Input['WorkerPoolNetworkConfigArgs']]:
        """
        Network configuration for the `WorkerPool`. Structure is documented below.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: Optional[pulumi.Input['WorkerPoolNetworkConfigArgs']]):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. `WorkerPool` state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. A unique identifier for the `WorkerPool`.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Time at which the request to update the `WorkerPool` was received.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> Optional[pulumi.Input['WorkerPoolWorkerConfigArgs']]:
        """
        Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        return pulumi.get(self, "worker_config")

    @worker_config.setter
    def worker_config(self, value: Optional[pulumi.Input['WorkerPoolWorkerConfigArgs']]):
        pulumi.set(self, "worker_config", value)


class WorkerPool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['WorkerPoolNetworkConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input[pulumi.InputType['WorkerPoolWorkerConfigArgs']]] = None,
                 __props__=None):
        """
        Definition of custom Cloud Build WorkerPools for running jobs with custom configuration and custom networking.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        pool = gcp.cloudbuild.WorkerPool("pool",
            location="europe-west1",
            worker_config=gcp.cloudbuild.WorkerPoolWorkerConfigArgs(
                disk_size_gb=100,
                machine_type="e2-standard-4",
                no_external_ip=False,
            ))
        ```
        ### Network Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        servicenetworking = gcp.projects.Service("servicenetworking",
            service="servicenetworking.googleapis.com",
            disable_on_destroy=False)
        network = gcp.compute.Network("network", auto_create_subnetworks=False,
        opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        worker_range = gcp.compute.GlobalAddress("workerRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=network.id)
        worker_pool_conn = gcp.servicenetworking.Connection("workerPoolConn",
            network=network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[worker_range.name],
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        pool = gcp.cloudbuild.WorkerPool("pool",
            location="europe-west1",
            worker_config=gcp.cloudbuild.WorkerPoolWorkerConfigArgs(
                disk_size_gb=100,
                machine_type="e2-standard-4",
                no_external_ip=False,
            ),
            network_config=gcp.cloudbuild.WorkerPoolNetworkConfigArgs(
                peered_network=network.id,
            ),
            opts=pulumi.ResourceOptions(depends_on=[worker_pool_conn]))
        ```

        ## Import

        WorkerPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default projects/{{project}}/locations/{{location}}/workerPools/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
               limitations.
        :param pulumi.Input[str] display_name: A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[str] name: User-defined name of the `WorkerPool`.
        :param pulumi.Input[pulumi.InputType['WorkerPoolNetworkConfigArgs']] network_config: Network configuration for the `WorkerPool`. Structure is documented below.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[pulumi.InputType['WorkerPoolWorkerConfigArgs']] worker_config: Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkerPoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of custom Cloud Build WorkerPools for running jobs with custom configuration and custom networking.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        pool = gcp.cloudbuild.WorkerPool("pool",
            location="europe-west1",
            worker_config=gcp.cloudbuild.WorkerPoolWorkerConfigArgs(
                disk_size_gb=100,
                machine_type="e2-standard-4",
                no_external_ip=False,
            ))
        ```
        ### Network Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        servicenetworking = gcp.projects.Service("servicenetworking",
            service="servicenetworking.googleapis.com",
            disable_on_destroy=False)
        network = gcp.compute.Network("network", auto_create_subnetworks=False,
        opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        worker_range = gcp.compute.GlobalAddress("workerRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=network.id)
        worker_pool_conn = gcp.servicenetworking.Connection("workerPoolConn",
            network=network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[worker_range.name],
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        pool = gcp.cloudbuild.WorkerPool("pool",
            location="europe-west1",
            worker_config=gcp.cloudbuild.WorkerPoolWorkerConfigArgs(
                disk_size_gb=100,
                machine_type="e2-standard-4",
                no_external_ip=False,
            ),
            network_config=gcp.cloudbuild.WorkerPoolNetworkConfigArgs(
                peered_network=network.id,
            ),
            opts=pulumi.ResourceOptions(depends_on=[worker_pool_conn]))
        ```

        ## Import

        WorkerPool can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default projects/{{project}}/locations/{{location}}/workerPools/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:cloudbuild/workerPool:WorkerPool default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param WorkerPoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkerPoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['WorkerPoolNetworkConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 worker_config: Optional[pulumi.Input[pulumi.InputType['WorkerPoolWorkerConfigArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkerPoolArgs.__new__(WorkerPoolArgs)

            __props__.__dict__["annotations"] = annotations
            __props__.__dict__["display_name"] = display_name
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["network_config"] = network_config
            __props__.__dict__["project"] = project
            __props__.__dict__["worker_config"] = worker_config
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        super(WorkerPool, __self__).__init__(
            'gcp:cloudbuild/workerPool:WorkerPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            annotations: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            delete_time: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_config: Optional[pulumi.Input[pulumi.InputType['WorkerPoolNetworkConfigArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            uid: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            worker_config: Optional[pulumi.Input[pulumi.InputType['WorkerPoolWorkerConfigArgs']]] = None) -> 'WorkerPool':
        """
        Get an existing WorkerPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] annotations: User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
               limitations.
        :param pulumi.Input[str] create_time: Output only. Time at which the request to create the `WorkerPool` was received.
        :param pulumi.Input[str] delete_time: Output only. Time at which the request to delete the `WorkerPool` was received.
        :param pulumi.Input[str] display_name: A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        :param pulumi.Input[str] location: The location for the resource
        :param pulumi.Input[str] name: User-defined name of the `WorkerPool`.
        :param pulumi.Input[pulumi.InputType['WorkerPoolNetworkConfigArgs']] network_config: Network configuration for the `WorkerPool`. Structure is documented below.
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] state: Output only. `WorkerPool` state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
        :param pulumi.Input[str] uid: Output only. A unique identifier for the `WorkerPool`.
        :param pulumi.Input[str] update_time: Output only. Time at which the request to update the `WorkerPool` was received.
        :param pulumi.Input[pulumi.InputType['WorkerPoolWorkerConfigArgs']] worker_config: Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkerPoolState.__new__(_WorkerPoolState)

        __props__.__dict__["annotations"] = annotations
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["delete_time"] = delete_time
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["network_config"] = network_config
        __props__.__dict__["project"] = project
        __props__.__dict__["state"] = state
        __props__.__dict__["uid"] = uid
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["worker_config"] = worker_config
        return WorkerPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User specified annotations. See https://google.aip.dev/128#annotations for more details such as format and size
        limitations.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Output only. Time at which the request to create the `WorkerPool` was received.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        Output only. Time at which the request to delete the `WorkerPool` was received.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A user-specified, human-readable name for the `WorkerPool`. If provided, this value must be 1-63 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location for the resource
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        User-defined name of the `WorkerPool`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output[Optional['outputs.WorkerPoolNetworkConfig']]:
        """
        Network configuration for the `WorkerPool`. Structure is documented below.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Output only. `WorkerPool` state. Possible values: STATE_UNSPECIFIED, PENDING, APPROVED, REJECTED, CANCELLED
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Output only. A unique identifier for the `WorkerPool`.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Output only. Time at which the request to update the `WorkerPool` was received.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="workerConfig")
    def worker_config(self) -> pulumi.Output['outputs.WorkerPoolWorkerConfig']:
        """
        Configuration to be used for a creating workers in the `WorkerPool`. Structure is documented below.
        """
        return pulumi.get(self, "worker_config")

