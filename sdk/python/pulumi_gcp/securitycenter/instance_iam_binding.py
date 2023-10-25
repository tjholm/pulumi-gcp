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

__all__ = ['InstanceIamBindingArgs', 'InstanceIamBinding']

@pulumi.input_type
class InstanceIamBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['InstanceIamBindingConditionArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceIamBinding resource.
        :param pulumi.Input[str] name: The ID of the instance or a fully qualified identifier for the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the Data Fusion instance.
        """
        InstanceIamBindingArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            members=members,
            role=role,
            condition=condition,
            name=name,
            project=project,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             role: Optional[pulumi.Input[str]] = None,
             condition: Optional[pulumi.Input['InstanceIamBindingConditionArgs']] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if members is None:
            raise TypeError("Missing 'members' argument")
        if role is None:
            raise TypeError("Missing 'role' argument")

        _setter("members", members)
        _setter("role", role)
        if condition is not None:
            _setter("condition", condition)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['InstanceIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['InstanceIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance or a fully qualified identifier for the instance.
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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the Data Fusion instance.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _InstanceIamBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['InstanceIamBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceIamBinding resources.
        :param pulumi.Input[str] name: The ID of the instance or a fully qualified identifier for the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the Data Fusion instance.
        """
        _InstanceIamBindingState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            condition=condition,
            etag=etag,
            members=members,
            name=name,
            project=project,
            region=region,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             condition: Optional[pulumi.Input['InstanceIamBindingConditionArgs']] = None,
             etag: Optional[pulumi.Input[str]] = None,
             members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if condition is not None:
            _setter("condition", condition)
        if etag is not None:
            _setter("etag", etag)
        if members is not None:
            _setter("members", members)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['InstanceIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['InstanceIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance or a fully qualified identifier for the instance.
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
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the Data Fusion instance.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class InstanceIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['InstanceIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a Data Fusion instance.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/data-fusion/docs/reference/rest/v1beta1/projects.locations.instances)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-fusion/docs/)

        ## Example Usage
        ### Data Fusion Instance Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_instance = gcp.datafusion.Instance("basicInstance",
            region="us-central1",
            type="BASIC")
        ```
        ### Data Fusion Instance Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.appengine.get_default_service_account()
        network = gcp.compute.Network("network")
        private_ip_alloc = gcp.compute.GlobalAddress("privateIpAlloc",
            address_type="INTERNAL",
            purpose="VPC_PEERING",
            prefix_length=22,
            network=network.id)
        extended_instance = gcp.datafusion.Instance("extendedInstance",
            description="My Data Fusion instance",
            display_name="My Data Fusion instance",
            region="us-central1",
            type="BASIC",
            enable_stackdriver_logging=True,
            enable_stackdriver_monitoring=True,
            private_instance=True,
            dataproc_service_account=default.email,
            labels={
                "example_key": "example_value",
            },
            network_config=gcp.datafusion.InstanceNetworkConfigArgs(
                network="default",
                ip_allocation=pulumi.Output.all(private_ip_alloc.address, private_ip_alloc.prefix_length).apply(lambda address, prefix_length: f"{address}/{prefix_length}"),
            ),
            accelerators=[gcp.datafusion.InstanceAcceleratorArgs(
                accelerator_type="CDC",
                state="ENABLED",
            )])
        ```
        ### Data Fusion Instance Cmek

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRing("keyRing", location="us-central1")
        crypto_key = gcp.kms.CryptoKey("cryptoKey", key_ring=key_ring.id)
        project = gcp.organizations.get_project()
        crypto_key_binding = gcp.kms.CryptoKeyIAMBinding("cryptoKeyBinding",
            crypto_key_id=crypto_key.id,
            role="roles/cloudkms.cryptoKeyEncrypterDecrypter",
            members=[f"serviceAccount:service-{project.number}@gcp-sa-datafusion.iam.gserviceaccount.com"])
        cmek = gcp.datafusion.Instance("cmek",
            region="us-central1",
            type="BASIC",
            crypto_key_config=gcp.datafusion.InstanceCryptoKeyConfigArgs(
                key_reference=crypto_key.id,
            ),
            opts=pulumi.ResourceOptions(depends_on=[crypto_key_binding]))
        ```
        ### Data Fusion Instance Enterprise

        ```python
        import pulumi
        import pulumi_gcp as gcp

        enterprise_instance = gcp.datafusion.Instance("enterpriseInstance",
            enable_rbac=True,
            region="us-central1",
            type="ENTERPRISE")
        ```
        ### Data Fusion Instance Event

        ```python
        import pulumi
        import pulumi_gcp as gcp

        event_topic = gcp.pubsub.Topic("eventTopic")
        event_instance = gcp.datafusion.Instance("eventInstance",
            region="us-central1",
            type="BASIC",
            event_publish_config=gcp.datafusion.InstanceEventPublishConfigArgs(
                enabled=True,
                topic=event_topic.id,
            ))
        ```
        ### Data Fusion Instance Zone

        ```python
        import pulumi
        import pulumi_gcp as gcp

        zone = gcp.datafusion.Instance("zone",
            region="us-central1",
            type="DEVELOPER",
            zone="us-central1-a")
        ```

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default projects/{{project}}/locations/{{region}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The ID of the instance or a fully qualified identifier for the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the Data Fusion instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a Data Fusion instance.

        To get more information about Instance, see:

        * [API documentation](https://cloud.google.com/data-fusion/docs/reference/rest/v1beta1/projects.locations.instances)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-fusion/docs/)

        ## Example Usage
        ### Data Fusion Instance Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_instance = gcp.datafusion.Instance("basicInstance",
            region="us-central1",
            type="BASIC")
        ```
        ### Data Fusion Instance Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.appengine.get_default_service_account()
        network = gcp.compute.Network("network")
        private_ip_alloc = gcp.compute.GlobalAddress("privateIpAlloc",
            address_type="INTERNAL",
            purpose="VPC_PEERING",
            prefix_length=22,
            network=network.id)
        extended_instance = gcp.datafusion.Instance("extendedInstance",
            description="My Data Fusion instance",
            display_name="My Data Fusion instance",
            region="us-central1",
            type="BASIC",
            enable_stackdriver_logging=True,
            enable_stackdriver_monitoring=True,
            private_instance=True,
            dataproc_service_account=default.email,
            labels={
                "example_key": "example_value",
            },
            network_config=gcp.datafusion.InstanceNetworkConfigArgs(
                network="default",
                ip_allocation=pulumi.Output.all(private_ip_alloc.address, private_ip_alloc.prefix_length).apply(lambda address, prefix_length: f"{address}/{prefix_length}"),
            ),
            accelerators=[gcp.datafusion.InstanceAcceleratorArgs(
                accelerator_type="CDC",
                state="ENABLED",
            )])
        ```
        ### Data Fusion Instance Cmek

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRing("keyRing", location="us-central1")
        crypto_key = gcp.kms.CryptoKey("cryptoKey", key_ring=key_ring.id)
        project = gcp.organizations.get_project()
        crypto_key_binding = gcp.kms.CryptoKeyIAMBinding("cryptoKeyBinding",
            crypto_key_id=crypto_key.id,
            role="roles/cloudkms.cryptoKeyEncrypterDecrypter",
            members=[f"serviceAccount:service-{project.number}@gcp-sa-datafusion.iam.gserviceaccount.com"])
        cmek = gcp.datafusion.Instance("cmek",
            region="us-central1",
            type="BASIC",
            crypto_key_config=gcp.datafusion.InstanceCryptoKeyConfigArgs(
                key_reference=crypto_key.id,
            ),
            opts=pulumi.ResourceOptions(depends_on=[crypto_key_binding]))
        ```
        ### Data Fusion Instance Enterprise

        ```python
        import pulumi
        import pulumi_gcp as gcp

        enterprise_instance = gcp.datafusion.Instance("enterpriseInstance",
            enable_rbac=True,
            region="us-central1",
            type="ENTERPRISE")
        ```
        ### Data Fusion Instance Event

        ```python
        import pulumi
        import pulumi_gcp as gcp

        event_topic = gcp.pubsub.Topic("eventTopic")
        event_instance = gcp.datafusion.Instance("eventInstance",
            region="us-central1",
            type="BASIC",
            event_publish_config=gcp.datafusion.InstanceEventPublishConfigArgs(
                enabled=True,
                topic=event_topic.id,
            ))
        ```
        ### Data Fusion Instance Zone

        ```python
        import pulumi
        import pulumi_gcp as gcp

        zone = gcp.datafusion.Instance("zone",
            region="us-central1",
            type="DEVELOPER",
            zone="us-central1-a")
        ```

        ## Import

        Instance can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default projects/{{project}}/locations/{{region}}/instances/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default {{project}}/{{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default {{region}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/instanceIamBinding:InstanceIamBinding default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param InstanceIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            InstanceIamBindingArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['InstanceIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceIamBindingArgs.__new__(InstanceIamBindingArgs)

            condition = _utilities.configure(condition, InstanceIamBindingConditionArgs, True)
            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(InstanceIamBinding, __self__).__init__(
            'gcp:securitycenter/instanceIamBinding:InstanceIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['InstanceIamBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'InstanceIamBinding':
        """
        Get an existing InstanceIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The ID of the instance or a fully qualified identifier for the instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: The region of the Data Fusion instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceIamBindingState.__new__(_InstanceIamBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["role"] = role
        return InstanceIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.InstanceIamBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The ID of the instance or a fully qualified identifier for the instance.
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
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region of the Data Fusion instance.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role")

