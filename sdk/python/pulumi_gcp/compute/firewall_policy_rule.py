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

__all__ = ['FirewallPolicyRuleArgs', 'FirewallPolicyRule']

@pulumi.input_type
class FirewallPolicyRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 direction: pulumi.Input[str],
                 firewall_policy: pulumi.Input[str],
                 match: pulumi.Input['FirewallPolicyRuleMatchArgs'],
                 priority: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a FirewallPolicyRule resource.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        :param pulumi.Input[str] direction: The direction in which this rule applies. Possible values: INGRESS, EGRESS
        :param pulumi.Input[str] firewall_policy: The firewall policy of the resource.
        :param pulumi.Input['FirewallPolicyRuleMatchArgs'] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[str] description: An optional description for this resource.
        :param pulumi.Input[bool] disabled: Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "direction", direction)
        pulumi.set(__self__, "firewall_policy", firewall_policy)
        pulumi.set(__self__, "match", match)
        pulumi.set(__self__, "priority", priority)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if enable_logging is not None:
            pulumi.set(__self__, "enable_logging", enable_logging)
        if target_resources is not None:
            pulumi.set(__self__, "target_resources", target_resources)
        if target_service_accounts is not None:
            pulumi.set(__self__, "target_service_accounts", target_service_accounts)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Input[str]:
        """
        The direction in which this rule applies. Possible values: INGRESS, EGRESS
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: pulumi.Input[str]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Input[str]:
        """
        The firewall policy of the resource.
        """
        return pulumi.get(self, "firewall_policy")

    @firewall_policy.setter
    def firewall_policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "firewall_policy", value)

    @property
    @pulumi.getter
    def match(self) -> pulumi.Input['FirewallPolicyRuleMatchArgs']:
        """
        A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: pulumi.Input['FirewallPolicyRuleMatchArgs']):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description for this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        """
        return pulumi.get(self, "enable_logging")

    @enable_logging.setter
    def enable_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_logging", value)

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        """
        return pulumi.get(self, "target_resources")

    @target_resources.setter
    def target_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_resources", value)

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        return pulumi.get(self, "target_service_accounts")

    @target_service_accounts.setter
    def target_service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_service_accounts", value)


@pulumi.input_type
class _FirewallPolicyRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input['FirewallPolicyRuleMatchArgs']] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 rule_tuple_count: Optional[pulumi.Input[int]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering FirewallPolicyRule resources.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        :param pulumi.Input[str] description: An optional description for this resource.
        :param pulumi.Input[str] direction: The direction in which this rule applies. Possible values: INGRESS, EGRESS
        :param pulumi.Input[bool] disabled: Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        :param pulumi.Input[str] firewall_policy: The firewall policy of the resource.
        :param pulumi.Input[str] kind: Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
        :param pulumi.Input['FirewallPolicyRuleMatchArgs'] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[int] rule_tuple_count: Calculation of the complexity of a single firewall policy rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if enable_logging is not None:
            pulumi.set(__self__, "enable_logging", enable_logging)
        if firewall_policy is not None:
            pulumi.set(__self__, "firewall_policy", firewall_policy)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if match is not None:
            pulumi.set(__self__, "match", match)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if rule_tuple_count is not None:
            pulumi.set(__self__, "rule_tuple_count", rule_tuple_count)
        if target_resources is not None:
            pulumi.set(__self__, "target_resources", target_resources)
        if target_service_accounts is not None:
            pulumi.set(__self__, "target_service_accounts", target_service_accounts)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description for this resource.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction in which this rule applies. Possible values: INGRESS, EGRESS
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        """
        return pulumi.get(self, "enable_logging")

    @enable_logging.setter
    def enable_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_logging", value)

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The firewall policy of the resource.
        """
        return pulumi.get(self, "firewall_policy")

    @firewall_policy.setter
    def firewall_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_policy", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def match(self) -> Optional[pulumi.Input['FirewallPolicyRuleMatchArgs']]:
        """
        A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: Optional[pulumi.Input['FirewallPolicyRuleMatchArgs']]):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="ruleTupleCount")
    def rule_tuple_count(self) -> Optional[pulumi.Input[int]]:
        """
        Calculation of the complexity of a single firewall policy rule.
        """
        return pulumi.get(self, "rule_tuple_count")

    @rule_tuple_count.setter
    def rule_tuple_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_tuple_count", value)

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        """
        return pulumi.get(self, "target_resources")

    @target_resources.setter
    def target_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_resources", value)

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        return pulumi.get(self, "target_service_accounts")

    @target_service_accounts.setter
    def target_service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_service_accounts", value)


class FirewallPolicyRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input[pulumi.InputType['FirewallPolicyRuleMatchArgs']]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The Compute FirewallPolicyRule resource

        ## Example Usage
        ### Basic_fir_sec_rule
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_global_networksecurity_address_group = gcp.networksecurity.AddressGroup("basicGlobalNetworksecurityAddressGroup",
            parent="organizations/123456789",
            description="Sample global networksecurity_address_group",
            location="global",
            items=["208.80.154.224/32"],
            type="IPV4",
            capacity=100)
        folder = gcp.organizations.Folder("folder",
            display_name="policy",
            parent="organizations/123456789")
        default = gcp.compute.FirewallPolicy("default",
            parent=folder.id,
            short_name="policy",
            description="Resource created for Terraform acceptance testing")
        primary = gcp.compute.FirewallPolicyRule("primary",
            firewall_policy=default.name,
            description="Resource created for Terraform acceptance testing",
            priority=9000,
            enable_logging=True,
            action="allow",
            direction="EGRESS",
            disabled=False,
            match=gcp.compute.FirewallPolicyRuleMatchArgs(
                layer4_configs=[
                    gcp.compute.FirewallPolicyRuleMatchLayer4ConfigArgs(
                        ip_protocol="tcp",
                        ports=["8080"],
                    ),
                    gcp.compute.FirewallPolicyRuleMatchLayer4ConfigArgs(
                        ip_protocol="udp",
                        ports=["22"],
                    ),
                ],
                dest_ip_ranges=["11.100.0.1/32"],
                dest_fqdns=[],
                dest_region_codes=["US"],
                dest_threat_intelligences=["iplist-known-malicious-ips"],
                src_address_groups=[],
                dest_address_groups=[basic_global_networksecurity_address_group.id],
            ),
            target_service_accounts=["my@service-account.com"])
        ```

        ## Import

        FirewallPolicyRule can be imported using any of these accepted formats* `locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}` * `{{firewall_policy}}/{{priority}}` When using the `pulumi import` command, FirewallPolicyRule can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
        ```

        ```sh
         $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default {{firewall_policy}}/{{priority}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        :param pulumi.Input[str] description: An optional description for this resource.
        :param pulumi.Input[str] direction: The direction in which this rule applies. Possible values: INGRESS, EGRESS
        :param pulumi.Input[bool] disabled: Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        :param pulumi.Input[str] firewall_policy: The firewall policy of the resource.
        :param pulumi.Input[pulumi.InputType['FirewallPolicyRuleMatchArgs']] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallPolicyRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Compute FirewallPolicyRule resource

        ## Example Usage
        ### Basic_fir_sec_rule
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_global_networksecurity_address_group = gcp.networksecurity.AddressGroup("basicGlobalNetworksecurityAddressGroup",
            parent="organizations/123456789",
            description="Sample global networksecurity_address_group",
            location="global",
            items=["208.80.154.224/32"],
            type="IPV4",
            capacity=100)
        folder = gcp.organizations.Folder("folder",
            display_name="policy",
            parent="organizations/123456789")
        default = gcp.compute.FirewallPolicy("default",
            parent=folder.id,
            short_name="policy",
            description="Resource created for Terraform acceptance testing")
        primary = gcp.compute.FirewallPolicyRule("primary",
            firewall_policy=default.name,
            description="Resource created for Terraform acceptance testing",
            priority=9000,
            enable_logging=True,
            action="allow",
            direction="EGRESS",
            disabled=False,
            match=gcp.compute.FirewallPolicyRuleMatchArgs(
                layer4_configs=[
                    gcp.compute.FirewallPolicyRuleMatchLayer4ConfigArgs(
                        ip_protocol="tcp",
                        ports=["8080"],
                    ),
                    gcp.compute.FirewallPolicyRuleMatchLayer4ConfigArgs(
                        ip_protocol="udp",
                        ports=["22"],
                    ),
                ],
                dest_ip_ranges=["11.100.0.1/32"],
                dest_fqdns=[],
                dest_region_codes=["US"],
                dest_threat_intelligences=["iplist-known-malicious-ips"],
                src_address_groups=[],
                dest_address_groups=[basic_global_networksecurity_address_group.id],
            ),
            target_service_accounts=["my@service-account.com"])
        ```

        ## Import

        FirewallPolicyRule can be imported using any of these accepted formats* `locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}` * `{{firewall_policy}}/{{priority}}` When using the `pulumi import` command, FirewallPolicyRule can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}
        ```

        ```sh
         $ pulumi import gcp:compute/firewallPolicyRule:FirewallPolicyRule default {{firewall_policy}}/{{priority}}
        ```

        :param str resource_name: The name of the resource.
        :param FirewallPolicyRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallPolicyRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 firewall_policy: Optional[pulumi.Input[str]] = None,
                 match: Optional[pulumi.Input[pulumi.InputType['FirewallPolicyRuleMatchArgs']]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallPolicyRuleArgs.__new__(FirewallPolicyRuleArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["description"] = description
            if direction is None and not opts.urn:
                raise TypeError("Missing required property 'direction'")
            __props__.__dict__["direction"] = direction
            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["enable_logging"] = enable_logging
            if firewall_policy is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_policy'")
            __props__.__dict__["firewall_policy"] = firewall_policy
            if match is None and not opts.urn:
                raise TypeError("Missing required property 'match'")
            __props__.__dict__["match"] = match
            if priority is None and not opts.urn:
                raise TypeError("Missing required property 'priority'")
            __props__.__dict__["priority"] = priority
            __props__.__dict__["target_resources"] = target_resources
            __props__.__dict__["target_service_accounts"] = target_service_accounts
            __props__.__dict__["kind"] = None
            __props__.__dict__["rule_tuple_count"] = None
        super(FirewallPolicyRule, __self__).__init__(
            'gcp:compute/firewallPolicyRule:FirewallPolicyRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            direction: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            enable_logging: Optional[pulumi.Input[bool]] = None,
            firewall_policy: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[str]] = None,
            match: Optional[pulumi.Input[pulumi.InputType['FirewallPolicyRuleMatchArgs']]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            rule_tuple_count: Optional[pulumi.Input[int]] = None,
            target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'FirewallPolicyRule':
        """
        Get an existing FirewallPolicyRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        :param pulumi.Input[str] description: An optional description for this resource.
        :param pulumi.Input[str] direction: The direction in which this rule applies. Possible values: INGRESS, EGRESS
        :param pulumi.Input[bool] disabled: Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        :param pulumi.Input[str] firewall_policy: The firewall policy of the resource.
        :param pulumi.Input[str] kind: Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
        :param pulumi.Input[pulumi.InputType['FirewallPolicyRuleMatchArgs']] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[int] rule_tuple_count: Calculation of the complexity of a single firewall policy rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FirewallPolicyRuleState.__new__(_FirewallPolicyRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["description"] = description
        __props__.__dict__["direction"] = direction
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["enable_logging"] = enable_logging
        __props__.__dict__["firewall_policy"] = firewall_policy
        __props__.__dict__["kind"] = kind
        __props__.__dict__["match"] = match
        __props__.__dict__["priority"] = priority
        __props__.__dict__["rule_tuple_count"] = rule_tuple_count
        __props__.__dict__["target_resources"] = target_resources
        __props__.__dict__["target_service_accounts"] = target_service_accounts
        return FirewallPolicyRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "goto_next".
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description for this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[str]:
        """
        The direction in which this rule applies. Possible values: INGRESS, EGRESS
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
        """
        return pulumi.get(self, "enable_logging")

    @property
    @pulumi.getter(name="firewallPolicy")
    def firewall_policy(self) -> pulumi.Output[str]:
        """
        The firewall policy of the resource.
        """
        return pulumi.get(self, "firewall_policy")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def match(self) -> pulumi.Output['outputs.FirewallPolicyRuleMatch']:
        """
        A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        """
        return pulumi.get(self, "match")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="ruleTupleCount")
    def rule_tuple_count(self) -> pulumi.Output[int]:
        """
        Calculation of the complexity of a single firewall policy rule.
        """
        return pulumi.get(self, "rule_tuple_count")

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of network resource URLs to which this rule applies. This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
        """
        return pulumi.get(self, "target_resources")

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of service accounts indicating the sets of instances that are applied with this rule.
        """
        return pulumi.get(self, "target_service_accounts")

