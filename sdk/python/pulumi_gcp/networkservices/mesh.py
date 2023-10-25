# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MeshArgs', 'Mesh']

@pulumi.input_type
class MeshArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 interception_port: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Mesh resource.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[int] interception_port: Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
               specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
               be redirected to this port regardless of its actual ip:port destination. If unset, a port
               '15001' is used as the interception port. This will is applicable only for sidecar proxy
               deployments.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of label tags associated with the Mesh resource.
        :param pulumi.Input[str] name: Short name of the Mesh resource to be created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        MeshArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            interception_port=interception_port,
            labels=labels,
            name=name,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             interception_port: Optional[pulumi.Input[int]] = None,
             labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if interception_port is None and 'interceptionPort' in kwargs:
            interception_port = kwargs['interceptionPort']

        if description is not None:
            _setter("description", description)
        if interception_port is not None:
            _setter("interception_port", interception_port)
        if labels is not None:
            _setter("labels", labels)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="interceptionPort")
    def interception_port(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
        specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
        be redirected to this port regardless of its actual ip:port destination. If unset, a port
        '15001' is used as the interception port. This will is applicable only for sidecar proxy
        deployments.
        """
        return pulumi.get(self, "interception_port")

    @interception_port.setter
    def interception_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interception_port", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of label tags associated with the Mesh resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Short name of the Mesh resource to be created.


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
class _MeshState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interception_port: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Mesh resources.
        :param pulumi.Input[str] create_time: Time the Mesh was created in UTC.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[int] interception_port: Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
               specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
               be redirected to this port regardless of its actual ip:port destination. If unset, a port
               '15001' is used as the interception port. This will is applicable only for sidecar proxy
               deployments.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of label tags associated with the Mesh resource.
        :param pulumi.Input[str] name: Short name of the Mesh resource to be created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: Server-defined URL of this resource.
        :param pulumi.Input[str] update_time: Time the Mesh was updated in UTC.
        """
        _MeshState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            create_time=create_time,
            description=description,
            interception_port=interception_port,
            labels=labels,
            name=name,
            project=project,
            self_link=self_link,
            update_time=update_time,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             create_time: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             interception_port: Optional[pulumi.Input[int]] = None,
             labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             self_link: Optional[pulumi.Input[str]] = None,
             update_time: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if create_time is None and 'createTime' in kwargs:
            create_time = kwargs['createTime']
        if interception_port is None and 'interceptionPort' in kwargs:
            interception_port = kwargs['interceptionPort']
        if self_link is None and 'selfLink' in kwargs:
            self_link = kwargs['selfLink']
        if update_time is None and 'updateTime' in kwargs:
            update_time = kwargs['updateTime']

        if create_time is not None:
            _setter("create_time", create_time)
        if description is not None:
            _setter("description", description)
        if interception_port is not None:
            _setter("interception_port", interception_port)
        if labels is not None:
            _setter("labels", labels)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if self_link is not None:
            _setter("self_link", self_link)
        if update_time is not None:
            _setter("update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the Mesh was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="interceptionPort")
    def interception_port(self) -> Optional[pulumi.Input[int]]:
        """
        Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
        specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
        be redirected to this port regardless of its actual ip:port destination. If unset, a port
        '15001' is used as the interception port. This will is applicable only for sidecar proxy
        deployments.
        """
        return pulumi.get(self, "interception_port")

    @interception_port.setter
    def interception_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interception_port", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of label tags associated with the Mesh resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Short name of the Mesh resource to be created.


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
        Server-defined URL of this resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the Mesh was updated in UTC.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Mesh(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interception_port: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Network Services Mesh Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networkservices.Mesh("default",
            labels={
                "foo": "bar",
            },
            description="my description",
            interception_port=443,
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Network Services Mesh No Port

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networkservices.Mesh("default",
            labels={
                "foo": "bar",
            },
            description="my description",
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Mesh can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:networkservices/mesh:Mesh default projects/{{project}}/locations/global/meshes/{{name}}
        ```

        ```sh
         $ pulumi import gcp:networkservices/mesh:Mesh default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:networkservices/mesh:Mesh default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[int] interception_port: Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
               specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
               be redirected to this port regardless of its actual ip:port destination. If unset, a port
               '15001' is used as the interception port. This will is applicable only for sidecar proxy
               deployments.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of label tags associated with the Mesh resource.
        :param pulumi.Input[str] name: Short name of the Mesh resource to be created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MeshArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Network Services Mesh Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networkservices.Mesh("default",
            labels={
                "foo": "bar",
            },
            description="my description",
            interception_port=443,
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Network Services Mesh No Port

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networkservices.Mesh("default",
            labels={
                "foo": "bar",
            },
            description="my description",
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        Mesh can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:networkservices/mesh:Mesh default projects/{{project}}/locations/global/meshes/{{name}}
        ```

        ```sh
         $ pulumi import gcp:networkservices/mesh:Mesh default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:networkservices/mesh:Mesh default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param MeshArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MeshArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            MeshArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 interception_port: Optional[pulumi.Input[int]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MeshArgs.__new__(MeshArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["interception_port"] = interception_port
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["update_time"] = None
        super(Mesh, __self__).__init__(
            'gcp:networkservices/mesh:Mesh',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            interception_port: Optional[pulumi.Input[int]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Mesh':
        """
        Get an existing Mesh resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Time the Mesh was created in UTC.
        :param pulumi.Input[str] description: A free-text description of the resource. Max length 1024 characters.
        :param pulumi.Input[int] interception_port: Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
               specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
               be redirected to this port regardless of its actual ip:port destination. If unset, a port
               '15001' is used as the interception port. This will is applicable only for sidecar proxy
               deployments.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of label tags associated with the Mesh resource.
        :param pulumi.Input[str] name: Short name of the Mesh resource to be created.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: Server-defined URL of this resource.
        :param pulumi.Input[str] update_time: Time the Mesh was updated in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MeshState.__new__(_MeshState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["interception_port"] = interception_port
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["update_time"] = update_time
        return Mesh(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time the Mesh was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A free-text description of the resource. Max length 1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="interceptionPort")
    def interception_port(self) -> pulumi.Output[Optional[int]]:
        """
        Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
        specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
        be redirected to this port regardless of its actual ip:port destination. If unset, a port
        '15001' is used as the interception port. This will is applicable only for sidecar proxy
        deployments.
        """
        return pulumi.get(self, "interception_port")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Set of label tags associated with the Mesh resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Short name of the Mesh resource to be created.


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
        Server-defined URL of this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time the Mesh was updated in UTC.
        """
        return pulumi.get(self, "update_time")

