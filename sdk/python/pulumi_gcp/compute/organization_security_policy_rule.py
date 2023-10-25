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

__all__ = ['OrganizationSecurityPolicyRuleArgs', 'OrganizationSecurityPolicyRule']

@pulumi.input_type
class OrganizationSecurityPolicyRuleArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 match: pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs'],
                 policy_id: pulumi.Input[str],
                 priority: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 preview: Optional[pulumi.Input[bool]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OrganizationSecurityPolicyRule resource.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Can currently be either
               "allow", "deny" or "goto_next".
        :param pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs'] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
               Structure is documented below.
        :param pulumi.Input[str] policy_id: The ID of the OrganizationSecurityPolicy this rule applies to.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a value
               between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
               highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[str] description: A description of the rule.
        :param pulumi.Input[str] direction: The direction in which this rule applies. If unspecified an INGRESS rule is created.
               Possible values are: `INGRESS`, `EGRESS`.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule.
               If logging is enabled, logs will be exported to the
               configured export destination in Stackdriver.
        :param pulumi.Input[bool] preview: If set to true, the specified action is not enforced.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies.
               This field allows you to control which network's VMs get
               this rule. If this field is left blank, all VMs
               within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of
               instances that are applied with this rule.
        """
        OrganizationSecurityPolicyRuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            match=match,
            policy_id=policy_id,
            priority=priority,
            description=description,
            direction=direction,
            enable_logging=enable_logging,
            preview=preview,
            target_resources=target_resources,
            target_service_accounts=target_service_accounts,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional[pulumi.Input[str]] = None,
             match: Optional[pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']] = None,
             policy_id: Optional[pulumi.Input[str]] = None,
             priority: Optional[pulumi.Input[int]] = None,
             description: Optional[pulumi.Input[str]] = None,
             direction: Optional[pulumi.Input[str]] = None,
             enable_logging: Optional[pulumi.Input[bool]] = None,
             preview: Optional[pulumi.Input[bool]] = None,
             target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if action is None:
            raise TypeError("Missing 'action' argument")
        if match is None:
            raise TypeError("Missing 'match' argument")
        if policy_id is None and 'policyId' in kwargs:
            policy_id = kwargs['policyId']
        if policy_id is None:
            raise TypeError("Missing 'policy_id' argument")
        if priority is None:
            raise TypeError("Missing 'priority' argument")
        if enable_logging is None and 'enableLogging' in kwargs:
            enable_logging = kwargs['enableLogging']
        if target_resources is None and 'targetResources' in kwargs:
            target_resources = kwargs['targetResources']
        if target_service_accounts is None and 'targetServiceAccounts' in kwargs:
            target_service_accounts = kwargs['targetServiceAccounts']

        _setter("action", action)
        _setter("match", match)
        _setter("policy_id", policy_id)
        _setter("priority", priority)
        if description is not None:
            _setter("description", description)
        if direction is not None:
            _setter("direction", direction)
        if enable_logging is not None:
            _setter("enable_logging", enable_logging)
        if preview is not None:
            _setter("preview", preview)
        if target_resources is not None:
            _setter("target_resources", target_resources)
        if target_service_accounts is not None:
            _setter("target_service_accounts", target_service_accounts)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        The Action to perform when the client connection triggers the rule. Can currently be either
        "allow", "deny" or "goto_next".
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def match(self) -> pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']:
        """
        A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        Structure is documented below.
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The ID of the OrganizationSecurityPolicy this rule applies to.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Input[int]:
        """
        An integer indicating the priority of a rule in the list. The priority must be a value
        between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
        highest priority and 2147483647 is the lowest prority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: pulumi.Input[int]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction in which this rule applies. If unspecified an INGRESS rule is created.
        Possible values are: `INGRESS`, `EGRESS`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes whether to enable logging for a particular rule.
        If logging is enabled, logs will be exported to the
        configured export destination in Stackdriver.
        """
        return pulumi.get(self, "enable_logging")

    @enable_logging.setter
    def enable_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_logging", value)

    @property
    @pulumi.getter
    def preview(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the specified action is not enforced.
        """
        return pulumi.get(self, "preview")

    @preview.setter
    def preview(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preview", value)

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of network resource URLs to which this rule applies.
        This field allows you to control which network's VMs get
        this rule. If this field is left blank, all VMs
        within the organization will receive the rule.
        """
        return pulumi.get(self, "target_resources")

    @target_resources.setter
    def target_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_resources", value)

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of service accounts indicating the sets of
        instances that are applied with this rule.
        """
        return pulumi.get(self, "target_service_accounts")

    @target_service_accounts.setter
    def target_service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_service_accounts", value)


@pulumi.input_type
class _OrganizationSecurityPolicyRuleState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 match: Optional[pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 preview: Optional[pulumi.Input[bool]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering OrganizationSecurityPolicyRule resources.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Can currently be either
               "allow", "deny" or "goto_next".
        :param pulumi.Input[str] description: A description of the rule.
        :param pulumi.Input[str] direction: The direction in which this rule applies. If unspecified an INGRESS rule is created.
               Possible values are: `INGRESS`, `EGRESS`.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule.
               If logging is enabled, logs will be exported to the
               configured export destination in Stackdriver.
        :param pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs'] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
               Structure is documented below.
        :param pulumi.Input[str] policy_id: The ID of the OrganizationSecurityPolicy this rule applies to.
        :param pulumi.Input[bool] preview: If set to true, the specified action is not enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a value
               between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
               highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies.
               This field allows you to control which network's VMs get
               this rule. If this field is left blank, all VMs
               within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of
               instances that are applied with this rule.
        """
        _OrganizationSecurityPolicyRuleState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            action=action,
            description=description,
            direction=direction,
            enable_logging=enable_logging,
            match=match,
            policy_id=policy_id,
            preview=preview,
            priority=priority,
            target_resources=target_resources,
            target_service_accounts=target_service_accounts,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             action: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             direction: Optional[pulumi.Input[str]] = None,
             enable_logging: Optional[pulumi.Input[bool]] = None,
             match: Optional[pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']] = None,
             policy_id: Optional[pulumi.Input[str]] = None,
             preview: Optional[pulumi.Input[bool]] = None,
             priority: Optional[pulumi.Input[int]] = None,
             target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if enable_logging is None and 'enableLogging' in kwargs:
            enable_logging = kwargs['enableLogging']
        if policy_id is None and 'policyId' in kwargs:
            policy_id = kwargs['policyId']
        if target_resources is None and 'targetResources' in kwargs:
            target_resources = kwargs['targetResources']
        if target_service_accounts is None and 'targetServiceAccounts' in kwargs:
            target_service_accounts = kwargs['targetServiceAccounts']

        if action is not None:
            _setter("action", action)
        if description is not None:
            _setter("description", description)
        if direction is not None:
            _setter("direction", direction)
        if enable_logging is not None:
            _setter("enable_logging", enable_logging)
        if match is not None:
            _setter("match", match)
        if policy_id is not None:
            _setter("policy_id", policy_id)
        if preview is not None:
            _setter("preview", preview)
        if priority is not None:
            _setter("priority", priority)
        if target_resources is not None:
            _setter("target_resources", target_resources)
        if target_service_accounts is not None:
            _setter("target_service_accounts", target_service_accounts)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        The Action to perform when the client connection triggers the rule. Can currently be either
        "allow", "deny" or "goto_next".
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        The direction in which this rule applies. If unspecified an INGRESS rule is created.
        Possible values are: `INGRESS`, `EGRESS`.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes whether to enable logging for a particular rule.
        If logging is enabled, logs will be exported to the
        configured export destination in Stackdriver.
        """
        return pulumi.get(self, "enable_logging")

    @enable_logging.setter
    def enable_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_logging", value)

    @property
    @pulumi.getter
    def match(self) -> Optional[pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']]:
        """
        A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        Structure is documented below.
        """
        return pulumi.get(self, "match")

    @match.setter
    def match(self, value: Optional[pulumi.Input['OrganizationSecurityPolicyRuleMatchArgs']]):
        pulumi.set(self, "match", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the OrganizationSecurityPolicy this rule applies to.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def preview(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the specified action is not enforced.
        """
        return pulumi.get(self, "preview")

    @preview.setter
    def preview(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preview", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        An integer indicating the priority of a rule in the list. The priority must be a value
        between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
        highest priority and 2147483647 is the lowest prority.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of network resource URLs to which this rule applies.
        This field allows you to control which network's VMs get
        this rule. If this field is left blank, all VMs
        within the organization will receive the rule.
        """
        return pulumi.get(self, "target_resources")

    @target_resources.setter
    def target_resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_resources", value)

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of service accounts indicating the sets of
        instances that are applied with this rule.
        """
        return pulumi.get(self, "target_service_accounts")

    @target_service_accounts.setter
    def target_service_accounts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_service_accounts", value)


class OrganizationSecurityPolicyRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 match: Optional[pulumi.Input[pulumi.InputType['OrganizationSecurityPolicyRuleMatchArgs']]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 preview: Optional[pulumi.Input[bool]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        A rule for the OrganizationSecurityPolicy.

        To get more information about OrganizationSecurityPolicyRule, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addRule)
        * How-to Guides
            * [Creating firewall rules](https://cloud.google.com/vpc/docs/using-firewall-policies#create-rules)

        ## Example Usage
        ### Organization Security Policy Rule Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        policy_organization_security_policy = gcp.compute.OrganizationSecurityPolicy("policyOrganizationSecurityPolicy",
            display_name="tf-test",
            parent="organizations/123456789",
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy_rule = gcp.compute.OrganizationSecurityPolicyRule("policyOrganizationSecurityPolicyRule",
            policy_id=policy_organization_security_policy.id,
            action="allow",
            direction="INGRESS",
            enable_logging=True,
            match=gcp.compute.OrganizationSecurityPolicyRuleMatchArgs(
                config=gcp.compute.OrganizationSecurityPolicyRuleMatchConfigArgs(
                    src_ip_ranges=[
                        "192.168.0.0/16",
                        "10.0.0.0/8",
                    ],
                    layer4_configs=[
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="tcp",
                            ports=["22"],
                        ),
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="icmp",
                        ),
                    ],
                ),
            ),
            priority=100,
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        OrganizationSecurityPolicyRule can be imported using any of these accepted formats:

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule default {{policy_id}}/priority/{{priority}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Can currently be either
               "allow", "deny" or "goto_next".
        :param pulumi.Input[str] description: A description of the rule.
        :param pulumi.Input[str] direction: The direction in which this rule applies. If unspecified an INGRESS rule is created.
               Possible values are: `INGRESS`, `EGRESS`.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule.
               If logging is enabled, logs will be exported to the
               configured export destination in Stackdriver.
        :param pulumi.Input[pulumi.InputType['OrganizationSecurityPolicyRuleMatchArgs']] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
               Structure is documented below.
        :param pulumi.Input[str] policy_id: The ID of the OrganizationSecurityPolicy this rule applies to.
        :param pulumi.Input[bool] preview: If set to true, the specified action is not enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a value
               between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
               highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies.
               This field allows you to control which network's VMs get
               this rule. If this field is left blank, all VMs
               within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of
               instances that are applied with this rule.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationSecurityPolicyRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A rule for the OrganizationSecurityPolicy.

        To get more information about OrganizationSecurityPolicyRule, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addRule)
        * How-to Guides
            * [Creating firewall rules](https://cloud.google.com/vpc/docs/using-firewall-policies#create-rules)

        ## Example Usage
        ### Organization Security Policy Rule Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        policy_organization_security_policy = gcp.compute.OrganizationSecurityPolicy("policyOrganizationSecurityPolicy",
            display_name="tf-test",
            parent="organizations/123456789",
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy_rule = gcp.compute.OrganizationSecurityPolicyRule("policyOrganizationSecurityPolicyRule",
            policy_id=policy_organization_security_policy.id,
            action="allow",
            direction="INGRESS",
            enable_logging=True,
            match=gcp.compute.OrganizationSecurityPolicyRuleMatchArgs(
                config=gcp.compute.OrganizationSecurityPolicyRuleMatchConfigArgs(
                    src_ip_ranges=[
                        "192.168.0.0/16",
                        "10.0.0.0/8",
                    ],
                    layer4_configs=[
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="tcp",
                            ports=["22"],
                        ),
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="icmp",
                        ),
                    ],
                ),
            ),
            priority=100,
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        OrganizationSecurityPolicyRule can be imported using any of these accepted formats:

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule default {{policy_id}}/priority/{{priority}}
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationSecurityPolicyRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationSecurityPolicyRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            OrganizationSecurityPolicyRuleArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 match: Optional[pulumi.Input[pulumi.InputType['OrganizationSecurityPolicyRuleMatchArgs']]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 preview: Optional[pulumi.Input[bool]] = None,
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
            __props__ = OrganizationSecurityPolicyRuleArgs.__new__(OrganizationSecurityPolicyRuleArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["description"] = description
            __props__.__dict__["direction"] = direction
            __props__.__dict__["enable_logging"] = enable_logging
            match = _utilities.configure(match, OrganizationSecurityPolicyRuleMatchArgs, True)
            if match is None and not opts.urn:
                raise TypeError("Missing required property 'match'")
            __props__.__dict__["match"] = match
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["preview"] = preview
            if priority is None and not opts.urn:
                raise TypeError("Missing required property 'priority'")
            __props__.__dict__["priority"] = priority
            __props__.__dict__["target_resources"] = target_resources
            __props__.__dict__["target_service_accounts"] = target_service_accounts
        super(OrganizationSecurityPolicyRule, __self__).__init__(
            'gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule',
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
            enable_logging: Optional[pulumi.Input[bool]] = None,
            match: Optional[pulumi.Input[pulumi.InputType['OrganizationSecurityPolicyRuleMatchArgs']]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            preview: Optional[pulumi.Input[bool]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            target_resources: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'OrganizationSecurityPolicyRule':
        """
        Get an existing OrganizationSecurityPolicyRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The Action to perform when the client connection triggers the rule. Can currently be either
               "allow", "deny" or "goto_next".
        :param pulumi.Input[str] description: A description of the rule.
        :param pulumi.Input[str] direction: The direction in which this rule applies. If unspecified an INGRESS rule is created.
               Possible values are: `INGRESS`, `EGRESS`.
        :param pulumi.Input[bool] enable_logging: Denotes whether to enable logging for a particular rule.
               If logging is enabled, logs will be exported to the
               configured export destination in Stackdriver.
        :param pulumi.Input[pulumi.InputType['OrganizationSecurityPolicyRuleMatchArgs']] match: A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
               Structure is documented below.
        :param pulumi.Input[str] policy_id: The ID of the OrganizationSecurityPolicy this rule applies to.
        :param pulumi.Input[bool] preview: If set to true, the specified action is not enforced.
        :param pulumi.Input[int] priority: An integer indicating the priority of a rule in the list. The priority must be a value
               between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
               highest priority and 2147483647 is the lowest prority.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_resources: A list of network resource URLs to which this rule applies.
               This field allows you to control which network's VMs get
               this rule. If this field is left blank, all VMs
               within the organization will receive the rule.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_service_accounts: A list of service accounts indicating the sets of
               instances that are applied with this rule.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationSecurityPolicyRuleState.__new__(_OrganizationSecurityPolicyRuleState)

        __props__.__dict__["action"] = action
        __props__.__dict__["description"] = description
        __props__.__dict__["direction"] = direction
        __props__.__dict__["enable_logging"] = enable_logging
        __props__.__dict__["match"] = match
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["preview"] = preview
        __props__.__dict__["priority"] = priority
        __props__.__dict__["target_resources"] = target_resources
        __props__.__dict__["target_service_accounts"] = target_service_accounts
        return OrganizationSecurityPolicyRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        The Action to perform when the client connection triggers the rule. Can currently be either
        "allow", "deny" or "goto_next".
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[Optional[str]]:
        """
        The direction in which this rule applies. If unspecified an INGRESS rule is created.
        Possible values are: `INGRESS`, `EGRESS`.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="enableLogging")
    def enable_logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Denotes whether to enable logging for a particular rule.
        If logging is enabled, logs will be exported to the
        configured export destination in Stackdriver.
        """
        return pulumi.get(self, "enable_logging")

    @property
    @pulumi.getter
    def match(self) -> pulumi.Output['outputs.OrganizationSecurityPolicyRuleMatch']:
        """
        A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        Structure is documented below.
        """
        return pulumi.get(self, "match")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The ID of the OrganizationSecurityPolicy this rule applies to.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter
    def preview(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to true, the specified action is not enforced.
        """
        return pulumi.get(self, "preview")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        An integer indicating the priority of a rule in the list. The priority must be a value
        between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
        highest priority and 2147483647 is the lowest prority.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="targetResources")
    def target_resources(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of network resource URLs to which this rule applies.
        This field allows you to control which network's VMs get
        this rule. If this field is left blank, all VMs
        within the organization will receive the rule.
        """
        return pulumi.get(self, "target_resources")

    @property
    @pulumi.getter(name="targetServiceAccounts")
    def target_service_accounts(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of service accounts indicating the sets of
        instances that are applied with this rule.
        """
        return pulumi.get(self, "target_service_accounts")

