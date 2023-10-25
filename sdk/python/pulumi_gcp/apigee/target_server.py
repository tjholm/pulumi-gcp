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

__all__ = ['TargetServerArgs', 'TargetServer']

@pulumi.input_type
class TargetServerArgs:
    def __init__(__self__, *,
                 env_id: pulumi.Input[str],
                 host: pulumi.Input[str],
                 port: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 s_sl_info: Optional[pulumi.Input['TargetServerSSlInfoArgs']] = None):
        """
        The set of arguments for constructing a TargetServer resource.
        :param pulumi.Input[str] env_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/environments/{{env_name}}`.
               
               
               - - -
        :param pulumi.Input[str] host: The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        :param pulumi.Input[int] port: The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        :param pulumi.Input[str] description: A human-readable description of this TargetServer.
        :param pulumi.Input[bool] is_enabled: Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        :param pulumi.Input[str] name: The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        :param pulumi.Input[str] protocol: Immutable. The protocol used by this TargetServer.
               Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        :param pulumi.Input['TargetServerSSlInfoArgs'] s_sl_info: Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
               Structure is documented below.
        """
        TargetServerArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            env_id=env_id,
            host=host,
            port=port,
            description=description,
            is_enabled=is_enabled,
            name=name,
            protocol=protocol,
            s_sl_info=s_sl_info,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             env_id: Optional[pulumi.Input[str]] = None,
             host: Optional[pulumi.Input[str]] = None,
             port: Optional[pulumi.Input[int]] = None,
             description: Optional[pulumi.Input[str]] = None,
             is_enabled: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             protocol: Optional[pulumi.Input[str]] = None,
             s_sl_info: Optional[pulumi.Input['TargetServerSSlInfoArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if env_id is None and 'envId' in kwargs:
            env_id = kwargs['envId']
        if env_id is None:
            raise TypeError("Missing 'env_id' argument")
        if host is None:
            raise TypeError("Missing 'host' argument")
        if port is None:
            raise TypeError("Missing 'port' argument")
        if is_enabled is None and 'isEnabled' in kwargs:
            is_enabled = kwargs['isEnabled']
        if s_sl_info is None and 'sSlInfo' in kwargs:
            s_sl_info = kwargs['sSlInfo']

        _setter("env_id", env_id)
        _setter("host", host)
        _setter("port", port)
        if description is not None:
            _setter("description", description)
        if is_enabled is not None:
            _setter("is_enabled", is_enabled)
        if name is not None:
            _setter("name", name)
        if protocol is not None:
            _setter("protocol", protocol)
        if s_sl_info is not None:
            _setter("s_sl_info", s_sl_info)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> pulumi.Input[str]:
        """
        The Apigee environment group associated with the Apigee environment,
        in the format `organizations/{{org_name}}/environments/{{env_name}}`.


        - - -
        """
        return pulumi.get(self, "env_id")

    @env_id.setter
    def env_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "env_id", value)

    @property
    @pulumi.getter
    def host(self) -> pulumi.Input[str]:
        """
        The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: pulumi.Input[str]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable description of this TargetServer.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The protocol used by this TargetServer.
        Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="sSlInfo")
    def s_sl_info(self) -> Optional[pulumi.Input['TargetServerSSlInfoArgs']]:
        """
        Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        Structure is documented below.
        """
        return pulumi.get(self, "s_sl_info")

    @s_sl_info.setter
    def s_sl_info(self, value: Optional[pulumi.Input['TargetServerSSlInfoArgs']]):
        pulumi.set(self, "s_sl_info", value)


@pulumi.input_type
class _TargetServerState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 env_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 s_sl_info: Optional[pulumi.Input['TargetServerSSlInfoArgs']] = None):
        """
        Input properties used for looking up and filtering TargetServer resources.
        :param pulumi.Input[str] description: A human-readable description of this TargetServer.
        :param pulumi.Input[str] env_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/environments/{{env_name}}`.
               
               
               - - -
        :param pulumi.Input[str] host: The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        :param pulumi.Input[bool] is_enabled: Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        :param pulumi.Input[str] name: The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        :param pulumi.Input[int] port: The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        :param pulumi.Input[str] protocol: Immutable. The protocol used by this TargetServer.
               Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        :param pulumi.Input['TargetServerSSlInfoArgs'] s_sl_info: Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
               Structure is documented below.
        """
        _TargetServerState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            description=description,
            env_id=env_id,
            host=host,
            is_enabled=is_enabled,
            name=name,
            port=port,
            protocol=protocol,
            s_sl_info=s_sl_info,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             description: Optional[pulumi.Input[str]] = None,
             env_id: Optional[pulumi.Input[str]] = None,
             host: Optional[pulumi.Input[str]] = None,
             is_enabled: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             port: Optional[pulumi.Input[int]] = None,
             protocol: Optional[pulumi.Input[str]] = None,
             s_sl_info: Optional[pulumi.Input['TargetServerSSlInfoArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if env_id is None and 'envId' in kwargs:
            env_id = kwargs['envId']
        if is_enabled is None and 'isEnabled' in kwargs:
            is_enabled = kwargs['isEnabled']
        if s_sl_info is None and 'sSlInfo' in kwargs:
            s_sl_info = kwargs['sSlInfo']

        if description is not None:
            _setter("description", description)
        if env_id is not None:
            _setter("env_id", env_id)
        if host is not None:
            _setter("host", host)
        if is_enabled is not None:
            _setter("is_enabled", is_enabled)
        if name is not None:
            _setter("name", name)
        if port is not None:
            _setter("port", port)
        if protocol is not None:
            _setter("protocol", protocol)
        if s_sl_info is not None:
            _setter("s_sl_info", s_sl_info)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-readable description of this TargetServer.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Apigee environment group associated with the Apigee environment,
        in the format `organizations/{{org_name}}/environments/{{env_name}}`.


        - - -
        """
        return pulumi.get(self, "env_id")

    @env_id.setter
    def env_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "env_id", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        """
        The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        """
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        """
        return pulumi.get(self, "is_enabled")

    @is_enabled.setter
    def is_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The protocol used by this TargetServer.
        Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="sSlInfo")
    def s_sl_info(self) -> Optional[pulumi.Input['TargetServerSSlInfoArgs']]:
        """
        Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        Structure is documented below.
        """
        return pulumi.get(self, "s_sl_info")

    @s_sl_info.setter
    def s_sl_info(self, value: Optional[pulumi.Input['TargetServerSSlInfoArgs']]):
        pulumi.set(self, "s_sl_info", value)


class TargetServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 s_sl_info: Optional[pulumi.Input[pulumi.InputType['TargetServerSSlInfoArgs']]] = None,
                 __props__=None):
        """
        TargetServer configuration. TargetServers are used to decouple a proxy TargetEndpoint HTTPTargetConnections from concrete URLs for backend services.

        To get more information about TargetServer, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.targetservers/create)
        * How-to Guides
            * [Load balancing across backend servers](https://cloud.google.com/apigee/docs/api-platform/deploy/load-balancing-across-backend-servers)

        ## Example Usage
        ### Apigee Target Server Test Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project",
            project_id="my-project",
            org_id="123456789",
            billing_account="000000-0000000-0000000-000000")
        apigee = gcp.projects.Service("apigee",
            project=project.project_id,
            service="apigee.googleapis.com")
        servicenetworking = gcp.projects.Service("servicenetworking",
            project=project.project_id,
            service="servicenetworking.googleapis.com",
            opts=pulumi.ResourceOptions(depends_on=[apigee]))
        compute = gcp.projects.Service("compute",
            project=project.project_id,
            service="compute.googleapis.com",
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        apigee_network = gcp.compute.Network("apigeeNetwork", project=project.project_id,
        opts=pulumi.ResourceOptions(depends_on=[compute]))
        apigee_range = gcp.compute.GlobalAddress("apigeeRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=apigee_network.id,
            project=project.project_id)
        apigee_vpc_connection = gcp.servicenetworking.Connection("apigeeVpcConnection",
            network=apigee_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[apigee_range.name],
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        apigee_org = gcp.apigee.Organization("apigeeOrg",
            analytics_region="us-central1",
            project_id=project.project_id,
            authorized_network=apigee_network.id,
            opts=pulumi.ResourceOptions(depends_on=[
                    apigee_vpc_connection,
                    apigee,
                ]))
        apigee_environment = gcp.apigee.Environment("apigeeEnvironment",
            org_id=apigee_org.id,
            description="Apigee Environment",
            display_name="environment-1")
        apigee_target_server = gcp.apigee.TargetServer("apigeeTargetServer",
            description="Apigee Target Server",
            protocol="HTTP",
            host="abc.foo.com",
            port=8080,
            env_id=apigee_environment.id)
        ```

        ## Import

        TargetServer can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/targetservers/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description of this TargetServer.
        :param pulumi.Input[str] env_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/environments/{{env_name}}`.
               
               
               - - -
        :param pulumi.Input[str] host: The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        :param pulumi.Input[bool] is_enabled: Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        :param pulumi.Input[str] name: The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        :param pulumi.Input[int] port: The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        :param pulumi.Input[str] protocol: Immutable. The protocol used by this TargetServer.
               Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        :param pulumi.Input[pulumi.InputType['TargetServerSSlInfoArgs']] s_sl_info: Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TargetServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        TargetServer configuration. TargetServers are used to decouple a proxy TargetEndpoint HTTPTargetConnections from concrete URLs for backend services.

        To get more information about TargetServer, see:

        * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.targetservers/create)
        * How-to Guides
            * [Load balancing across backend servers](https://cloud.google.com/apigee/docs/api-platform/deploy/load-balancing-across-backend-servers)

        ## Example Usage
        ### Apigee Target Server Test Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.Project("project",
            project_id="my-project",
            org_id="123456789",
            billing_account="000000-0000000-0000000-000000")
        apigee = gcp.projects.Service("apigee",
            project=project.project_id,
            service="apigee.googleapis.com")
        servicenetworking = gcp.projects.Service("servicenetworking",
            project=project.project_id,
            service="servicenetworking.googleapis.com",
            opts=pulumi.ResourceOptions(depends_on=[apigee]))
        compute = gcp.projects.Service("compute",
            project=project.project_id,
            service="compute.googleapis.com",
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        apigee_network = gcp.compute.Network("apigeeNetwork", project=project.project_id,
        opts=pulumi.ResourceOptions(depends_on=[compute]))
        apigee_range = gcp.compute.GlobalAddress("apigeeRange",
            purpose="VPC_PEERING",
            address_type="INTERNAL",
            prefix_length=16,
            network=apigee_network.id,
            project=project.project_id)
        apigee_vpc_connection = gcp.servicenetworking.Connection("apigeeVpcConnection",
            network=apigee_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[apigee_range.name],
            opts=pulumi.ResourceOptions(depends_on=[servicenetworking]))
        apigee_org = gcp.apigee.Organization("apigeeOrg",
            analytics_region="us-central1",
            project_id=project.project_id,
            authorized_network=apigee_network.id,
            opts=pulumi.ResourceOptions(depends_on=[
                    apigee_vpc_connection,
                    apigee,
                ]))
        apigee_environment = gcp.apigee.Environment("apigeeEnvironment",
            org_id=apigee_org.id,
            description="Apigee Environment",
            display_name="environment-1")
        apigee_target_server = gcp.apigee.TargetServer("apigeeTargetServer",
            description="Apigee Target Server",
            protocol="HTTP",
            host="abc.foo.com",
            port=8080,
            env_id=apigee_environment.id)
        ```

        ## Import

        TargetServer can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/targetservers/{{name}}
        ```

        ```sh
         $ pulumi import gcp:apigee/targetServer:TargetServer default {{env_id}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param TargetServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            TargetServerArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env_id: Optional[pulumi.Input[str]] = None,
                 host: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 s_sl_info: Optional[pulumi.Input[pulumi.InputType['TargetServerSSlInfoArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TargetServerArgs.__new__(TargetServerArgs)

            __props__.__dict__["description"] = description
            if env_id is None and not opts.urn:
                raise TypeError("Missing required property 'env_id'")
            __props__.__dict__["env_id"] = env_id
            if host is None and not opts.urn:
                raise TypeError("Missing required property 'host'")
            __props__.__dict__["host"] = host
            __props__.__dict__["is_enabled"] = is_enabled
            __props__.__dict__["name"] = name
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            __props__.__dict__["protocol"] = protocol
            s_sl_info = _utilities.configure(s_sl_info, TargetServerSSlInfoArgs, True)
            __props__.__dict__["s_sl_info"] = s_sl_info
        super(TargetServer, __self__).__init__(
            'gcp:apigee/targetServer:TargetServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            env_id: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            s_sl_info: Optional[pulumi.Input[pulumi.InputType['TargetServerSSlInfoArgs']]] = None) -> 'TargetServer':
        """
        Get an existing TargetServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A human-readable description of this TargetServer.
        :param pulumi.Input[str] env_id: The Apigee environment group associated with the Apigee environment,
               in the format `organizations/{{org_name}}/environments/{{env_name}}`.
               
               
               - - -
        :param pulumi.Input[str] host: The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        :param pulumi.Input[bool] is_enabled: Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        :param pulumi.Input[str] name: The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        :param pulumi.Input[int] port: The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        :param pulumi.Input[str] protocol: Immutable. The protocol used by this TargetServer.
               Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        :param pulumi.Input[pulumi.InputType['TargetServerSSlInfoArgs']] s_sl_info: Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TargetServerState.__new__(_TargetServerState)

        __props__.__dict__["description"] = description
        __props__.__dict__["env_id"] = env_id
        __props__.__dict__["host"] = host
        __props__.__dict__["is_enabled"] = is_enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["port"] = port
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["s_sl_info"] = s_sl_info
        return TargetServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-readable description of this TargetServer.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> pulumi.Output[str]:
        """
        The Apigee environment group associated with the Apigee environment,
        in the format `organizations/{{org_name}}/environments/{{env_name}}`.


        - - -
        """
        return pulumi.get(self, "env_id")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or more TargetServers need to taken out of rotation periodically. Defaults to true.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource id of this reference. Values must match the regular expression [\\w\\s-.]+.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Immutable. The protocol used by this TargetServer.
        Possible values are: `HTTP`, `HTTP2`, `GRPC_TARGET`, `GRPC`, `EXTERNAL_CALLOUT`.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sSlInfo")
    def s_sl_info(self) -> pulumi.Output[Optional['outputs.TargetServerSSlInfo']]:
        """
        Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
        Structure is documented below.
        """
        return pulumi.get(self, "s_sl_info")

