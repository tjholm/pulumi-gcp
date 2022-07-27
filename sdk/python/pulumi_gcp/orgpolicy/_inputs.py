# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'PolicySpecArgs',
    'PolicySpecRuleArgs',
    'PolicySpecRuleConditionArgs',
    'PolicySpecRuleValuesArgs',
]

@pulumi.input_type
class PolicySpecArgs:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 inherit_from_parent: Optional[pulumi.Input[bool]] = None,
                 reset: Optional[pulumi.Input[bool]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['PolicySpecRuleArgs']]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] etag: -
               An opaque tag indicating the current version of the `Policy`, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the `Policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current `Policy` to use when executing a read-modify-write loop. When the `Policy` is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        :param pulumi.Input[bool] inherit_from_parent: Determines the inheritance behavior for this `Policy`. If `inherit_from_parent` is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
        :param pulumi.Input[bool] reset: Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific `Constraint` at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        :param pulumi.Input[Sequence[pulumi.Input['PolicySpecRuleArgs']]] rules: Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set `enforced` to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
        :param pulumi.Input[str] update_time: -
               Output only. The time stamp this was previously updated. This represents the last time a call to `CreatePolicy` or `UpdatePolicy` was made for that `Policy`.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if inherit_from_parent is not None:
            pulumi.set(__self__, "inherit_from_parent", inherit_from_parent)
        if reset is not None:
            pulumi.set(__self__, "reset", reset)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        -
        An opaque tag indicating the current version of the `Policy`, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the `Policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current `Policy` to use when executing a read-modify-write loop. When the `Policy` is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="inheritFromParent")
    def inherit_from_parent(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines the inheritance behavior for this `Policy`. If `inherit_from_parent` is true, PolicyRules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this Policy becomes the new root for evaluation. This field can be set only for Policies which configure list constraints.
        """
        return pulumi.get(self, "inherit_from_parent")

    @inherit_from_parent.setter
    def inherit_from_parent(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "inherit_from_parent", value)

    @property
    @pulumi.getter
    def reset(self) -> Optional[pulumi.Input[bool]]:
        """
        Ignores policies set above this resource and restores the `constraint_default` enforcement behavior of the specific `Constraint` at this resource. This field can be set in policies for either list or boolean constraints. If set, `rules` must be empty and `inherit_from_parent` must be set to false.
        """
        return pulumi.get(self, "reset")

    @reset.setter
    def reset(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reset", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['PolicySpecRuleArgs']]]]:
        """
        Up to 10 PolicyRules are allowed. In Policies for boolean constraints, the following requirements apply: - There must be one and only one PolicyRule where condition is unset. - BooleanPolicyRules with conditions must set `enforced` to the opposite of the PolicyRule without a condition. - During policy evaluation, PolicyRules with conditions that are true for a target resource take precedence.
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['PolicySpecRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        -
        Output only. The time stamp this was previously updated. This represents the last time a call to `CreatePolicy` or `UpdatePolicy` was made for that `Policy`.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


@pulumi.input_type
class PolicySpecRuleArgs:
    def __init__(__self__, *,
                 allow_all: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['PolicySpecRuleConditionArgs']] = None,
                 deny_all: Optional[pulumi.Input[str]] = None,
                 enforce: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input['PolicySpecRuleValuesArgs']] = None):
        """
        :param pulumi.Input[str] allow_all: Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
        :param pulumi.Input['PolicySpecRuleConditionArgs'] condition: A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        :param pulumi.Input[str] deny_all: Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
        :param pulumi.Input[str] enforce: If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
        :param pulumi.Input['PolicySpecRuleValuesArgs'] values: List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.
        """
        if allow_all is not None:
            pulumi.set(__self__, "allow_all", allow_all)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if deny_all is not None:
            pulumi.set(__self__, "deny_all", deny_all)
        if enforce is not None:
            pulumi.set(__self__, "enforce", enforce)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter(name="allowAll")
    def allow_all(self) -> Optional[pulumi.Input[str]]:
        """
        Setting this to true means that all values are allowed. This field can be set only in Policies for list constraints.
        """
        return pulumi.get(self, "allow_all")

    @allow_all.setter
    def allow_all(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_all", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['PolicySpecRuleConditionArgs']]:
        """
        A condition which determines whether this rule is used in the evaluation of the policy. When set, the `expression` field in the `Expr' must include from 1 to 10 subexpressions, joined by the "||" or "&&" operators. Each subexpression must be of the form "resource.matchTag('/tag_key_short_name, 'tag_value_short_name')". or "resource.matchTagId('tagKeys/key_id', 'tagValues/value_id')". where key_name and value_name are the resource names for Label Keys and Values. These names are available from the Tag Manager Service. An example expression is: "resource.matchTag('123456789/environment, 'prod')". or "resource.matchTagId('tagKeys/123', 'tagValues/456')".
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['PolicySpecRuleConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="denyAll")
    def deny_all(self) -> Optional[pulumi.Input[str]]:
        """
        Setting this to true means that all values are denied. This field can be set only in Policies for list constraints.
        """
        return pulumi.get(self, "deny_all")

    @deny_all.setter
    def deny_all(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deny_all", value)

    @property
    @pulumi.getter
    def enforce(self) -> Optional[pulumi.Input[str]]:
        """
        If `true`, then the `Policy` is enforced. If `false`, then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
        """
        return pulumi.get(self, "enforce")

    @enforce.setter
    def enforce(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enforce", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input['PolicySpecRuleValuesArgs']]:
        """
        List of values to be used for this PolicyRule. This field can be set only in Policies for list constraints.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input['PolicySpecRuleValuesArgs']]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class PolicySpecRuleConditionArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 expression: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param pulumi.Input[str] expression: Textual representation of an expression in Common Expression Language syntax.
        :param pulumi.Input[str] location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param pulumi.Input[str] title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class PolicySpecRuleValuesArgs:
    def __init__(__self__, *,
                 allowed_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 denied_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_values: List of values allowed at this resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] denied_values: List of values denied at this resource.
        """
        if allowed_values is not None:
            pulumi.set(__self__, "allowed_values", allowed_values)
        if denied_values is not None:
            pulumi.set(__self__, "denied_values", denied_values)

    @property
    @pulumi.getter(name="allowedValues")
    def allowed_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of values allowed at this resource.
        """
        return pulumi.get(self, "allowed_values")

    @allowed_values.setter
    def allowed_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_values", value)

    @property
    @pulumi.getter(name="deniedValues")
    def denied_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of values denied at this resource.
        """
        return pulumi.get(self, "denied_values")

    @denied_values.setter
    def denied_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "denied_values", value)


