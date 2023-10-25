# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegionNetworkFirewallPolicyAssociationArgs', 'RegionNetworkFirewallPolicyAssociation']

@pulumi.input_type
class RegionNetworkFirewallPolicyAssociationArgs:
    def __init__(__self__, *,
                 attachment_target: pulumi.Input[str],
                 firewall_policy: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegionNetworkFirewallPolicyAssociation resource.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
               
               
               
               - - -
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] region: The location of this resource.
        """
        RegionNetworkFirewallPolicyAssociationArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            attachment_target=attachment_target,
            firewall_policy=firewall_policy,
            name=name,
            project=project,
            region=region,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             attachment_target: Optional[pulumi.Input[str]] = None,
             firewall_policy: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if attachment_target is None and 'attachmentTarget' in kwargs:
            attachment_target = kwargs['attachmentTarget']
        if attachment_target is None:
            raise TypeError("Missing 'attachment_target' argument")
        if firewall_policy is None and 'firewallPolicy' in kwargs:
            firewall_policy = kwargs['firewallPolicy']
        if firewall_policy is None:
            raise TypeError("Missing 'firewall_policy' argument")

        _setter("attachment_target", attachment_target)
        _setter("firewall_policy", firewall_policy)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)

    @property
    @pulumi.getter(name="attachmentTarget")
    def attachment_target(self) -> pulumi.Input[str]:
        """
        The target that the firewall policy is attached to.
        """
        return pulumi.get(self, "attachment_target")

    @attachment_target.setter
    def attachment_target(self, value: pulumi.Input[str]):
        pulumi.set(self, "attachment_target", value)

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Input[str]:
        """
        The firewall policy ID of the association.
        """
        return pulumi.get(self, "firewall_policy")

    @firewall_policy.setter
    def firewall_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "firewall_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.



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
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The location of this resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _RegionNetworkFirewallPolicyAssociationState:
    def __init__(__self__, *,
                 attachment_target: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 short_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RegionNetworkFirewallPolicyAssociation resources.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
               
               
               
               - - -
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] region: The location of this resource.
        :param pulumi.Input[str] short_name: The short name of the firewall policy of the association.
        """
        _RegionNetworkFirewallPolicyAssociationState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            attachment_target=attachment_target,
            firewall_policy=firewall_policy,
            name=name,
            project=project,
            region=region,
            short_name=short_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             attachment_target: Optional[pulumi.Input[str]] = None,
             firewall_policy: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             region: Optional[pulumi.Input[str]] = None,
             short_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if attachment_target is None and 'attachmentTarget' in kwargs:
            attachment_target = kwargs['attachmentTarget']
        if firewall_policy is None and 'firewallPolicy' in kwargs:
            firewall_policy = kwargs['firewallPolicy']
        if short_name is None and 'shortName' in kwargs:
            short_name = kwargs['shortName']

        if attachment_target is not None:
            _setter("attachment_target", attachment_target)
        if firewall_policy is not None:
            _setter("firewall_policy", firewall_policy)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if region is not None:
            _setter("region", region)
        if short_name is not None:
            _setter("short_name", short_name)

    @property
    @pulumi.getter(name="attachmentTarget")
    def attachment_target(self) -> Optional[pulumi.Input[str]]:
        """
        The target that the firewall policy is attached to.
        """
        return pulumi.get(self, "attachment_target")

    @attachment_target.setter
    def attachment_target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attachment_target", value)

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The firewall policy ID of the association.
        """
        return pulumi.get(self, "firewall_policy")

    @firewall_policy.setter
    def firewall_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.



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
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The location of this resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> Optional[pulumi.Input[str]]:
        """
        The short name of the firewall policy of the association.
        """
        return pulumi.get(self, "short_name")

    @short_name.setter
    def short_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_name", value)


class RegionNetworkFirewallPolicyAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_target: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The Compute NetworkFirewallPolicyAssociation resource

        ## Example Usage
        ### Regional
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_regional_network_firewall_policy = gcp.compute.RegionNetworkFirewallPolicy("basicRegionalNetworkFirewallPolicy",
            project="my-project-name",
            description="Sample global network firewall policy",
            region="us-west1")
        basic_network = gcp.compute.Network("basicNetwork")
        primary = gcp.compute.RegionNetworkFirewallPolicyAssociation("primary",
            attachment_target=basic_network.id,
            firewall_policy=basic_regional_network_firewall_policy.name,
            project="my-project-name",
            region="us-west1")
        ```

        ## Import

        NetworkFirewallPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/regionNetworkFirewallPolicyAssociation:RegionNetworkFirewallPolicyAssociation default projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/associations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/regionNetworkFirewallPolicyAssociation:RegionNetworkFirewallPolicyAssociation default {{project}}/{{region}}/{{firewall_policy}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
               
               
               
               - - -
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] region: The location of this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionNetworkFirewallPolicyAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Compute NetworkFirewallPolicyAssociation resource

        ## Example Usage
        ### Regional
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_regional_network_firewall_policy = gcp.compute.RegionNetworkFirewallPolicy("basicRegionalNetworkFirewallPolicy",
            project="my-project-name",
            description="Sample global network firewall policy",
            region="us-west1")
        basic_network = gcp.compute.Network("basicNetwork")
        primary = gcp.compute.RegionNetworkFirewallPolicyAssociation("primary",
            attachment_target=basic_network.id,
            firewall_policy=basic_regional_network_firewall_policy.name,
            project="my-project-name",
            region="us-west1")
        ```

        ## Import

        NetworkFirewallPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/regionNetworkFirewallPolicyAssociation:RegionNetworkFirewallPolicyAssociation default projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/associations/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/regionNetworkFirewallPolicyAssociation:RegionNetworkFirewallPolicyAssociation default {{project}}/{{region}}/{{firewall_policy}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param RegionNetworkFirewallPolicyAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionNetworkFirewallPolicyAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RegionNetworkFirewallPolicyAssociationArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_target: Optional[pulumi.Input[str]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegionNetworkFirewallPolicyAssociationArgs.__new__(RegionNetworkFirewallPolicyAssociationArgs)

            if attachment_target is None and not opts.urn:
                raise TypeError("Missing required property 'attachment_target'")
            __props__.__dict__["attachment_target"] = attachment_target
            if firewall_policy is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_policy'")
            __props__.__dict__["firewall_policy"] = firewall_policy
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["short_name"] = None
        super(RegionNetworkFirewallPolicyAssociation, __self__).__init__(
            'gcp:compute/regionNetworkFirewallPolicyAssociation:RegionNetworkFirewallPolicyAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachment_target: Optional[pulumi.Input[str]] = None,
            firewall_policy: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            short_name: Optional[pulumi.Input[str]] = None) -> 'RegionNetworkFirewallPolicyAssociation':
        """
        Get an existing RegionNetworkFirewallPolicyAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_target: The target that the firewall policy is attached to.
        :param pulumi.Input[str] firewall_policy: The firewall policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
               
               
               
               - - -
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] region: The location of this resource.
        :param pulumi.Input[str] short_name: The short name of the firewall policy of the association.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegionNetworkFirewallPolicyAssociationState.__new__(_RegionNetworkFirewallPolicyAssociationState)

        __props__.__dict__["attachment_target"] = attachment_target
        __props__.__dict__["firewall_policy"] = firewall_policy
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["short_name"] = short_name
        return RegionNetworkFirewallPolicyAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachmentTarget")
    def attachment_target(self) -> pulumi.Output[str]:
        """
        The target that the firewall policy is attached to.
        """
        return pulumi.get(self, "attachment_target")

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Output[str]:
        """
        The firewall policy ID of the association.
        """
        return pulumi.get(self, "firewall_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for an association.



        - - -
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The location of this resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="shortName")
    def short_name(self) -> pulumi.Output[str]:
        """
        The short name of the firewall policy of the association.
        """
        return pulumi.get(self, "short_name")

