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

__all__ = ['BudgetArgs', 'Budget']

@pulumi.input_type
class BudgetArgs:
    def __init__(__self__, *,
                 amount: pulumi.Input['BudgetAmountArgs'],
                 billing_account: pulumi.Input[str],
                 all_updates_rule: Optional[pulumi.Input['BudgetAllUpdatesRuleArgs']] = None,
                 budget_filter: Optional[pulumi.Input['BudgetBudgetFilterArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]]] = None):
        """
        The set of arguments for constructing a Budget resource.
        :param pulumi.Input['BudgetAmountArgs'] amount: The budgeted amount for each usage period.
               Structure is documented below.
        :param pulumi.Input[str] billing_account: ID of the billing account to set a budget on.
        :param pulumi.Input['BudgetAllUpdatesRuleArgs'] all_updates_rule: Defines notifications that are sent on every update to the
               billing account's spend, regardless of the thresholds defined
               using threshold rules.
               Structure is documented below.
        :param pulumi.Input['BudgetBudgetFilterArgs'] budget_filter: Filters that define which resources are used to compute the actual
               spend against the budget.
               Structure is documented below.
        :param pulumi.Input[str] display_name: User data for display name in UI. Must be <= 60 chars.
        :param pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]] threshold_rules: Rules that trigger alerts (notifications of thresholds being
               crossed) when spend exceeds the specified percentages of the
               budget.
               Structure is documented below.
        """
        pulumi.set(__self__, "amount", amount)
        pulumi.set(__self__, "billing_account", billing_account)
        if all_updates_rule is not None:
            pulumi.set(__self__, "all_updates_rule", all_updates_rule)
        if budget_filter is not None:
            pulumi.set(__self__, "budget_filter", budget_filter)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if threshold_rules is not None:
            pulumi.set(__self__, "threshold_rules", threshold_rules)

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Input['BudgetAmountArgs']:
        """
        The budgeted amount for each usage period.
        Structure is documented below.
        """
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: pulumi.Input['BudgetAmountArgs']):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> pulumi.Input[str]:
        """
        ID of the billing account to set a budget on.
        """
        return pulumi.get(self, "billing_account")

    @billing_account.setter
    def billing_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_account", value)

    @property
    @pulumi.getter(name="allUpdatesRule")
    def all_updates_rule(self) -> Optional[pulumi.Input['BudgetAllUpdatesRuleArgs']]:
        """
        Defines notifications that are sent on every update to the
        billing account's spend, regardless of the thresholds defined
        using threshold rules.
        Structure is documented below.
        """
        return pulumi.get(self, "all_updates_rule")

    @all_updates_rule.setter
    def all_updates_rule(self, value: Optional[pulumi.Input['BudgetAllUpdatesRuleArgs']]):
        pulumi.set(self, "all_updates_rule", value)

    @property
    @pulumi.getter(name="budgetFilter")
    def budget_filter(self) -> Optional[pulumi.Input['BudgetBudgetFilterArgs']]:
        """
        Filters that define which resources are used to compute the actual
        spend against the budget.
        Structure is documented below.
        """
        return pulumi.get(self, "budget_filter")

    @budget_filter.setter
    def budget_filter(self, value: Optional[pulumi.Input['BudgetBudgetFilterArgs']]):
        pulumi.set(self, "budget_filter", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User data for display name in UI. Must be <= 60 chars.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="thresholdRules")
    def threshold_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]]]:
        """
        Rules that trigger alerts (notifications of thresholds being
        crossed) when spend exceeds the specified percentages of the
        budget.
        Structure is documented below.
        """
        return pulumi.get(self, "threshold_rules")

    @threshold_rules.setter
    def threshold_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]]]):
        pulumi.set(self, "threshold_rules", value)


@pulumi.input_type
class _BudgetState:
    def __init__(__self__, *,
                 all_updates_rule: Optional[pulumi.Input['BudgetAllUpdatesRuleArgs']] = None,
                 amount: Optional[pulumi.Input['BudgetAmountArgs']] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 budget_filter: Optional[pulumi.Input['BudgetBudgetFilterArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]]] = None):
        """
        Input properties used for looking up and filtering Budget resources.
        :param pulumi.Input['BudgetAllUpdatesRuleArgs'] all_updates_rule: Defines notifications that are sent on every update to the
               billing account's spend, regardless of the thresholds defined
               using threshold rules.
               Structure is documented below.
        :param pulumi.Input['BudgetAmountArgs'] amount: The budgeted amount for each usage period.
               Structure is documented below.
        :param pulumi.Input[str] billing_account: ID of the billing account to set a budget on.
        :param pulumi.Input['BudgetBudgetFilterArgs'] budget_filter: Filters that define which resources are used to compute the actual
               spend against the budget.
               Structure is documented below.
        :param pulumi.Input[str] display_name: User data for display name in UI. Must be <= 60 chars.
        :param pulumi.Input[str] name: Resource name of the budget. The resource name
               implies the scope of a budget. Values are of the form
               billingAccounts/{billingAccountId}/budgets/{budgetId}.
        :param pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]] threshold_rules: Rules that trigger alerts (notifications of thresholds being
               crossed) when spend exceeds the specified percentages of the
               budget.
               Structure is documented below.
        """
        if all_updates_rule is not None:
            pulumi.set(__self__, "all_updates_rule", all_updates_rule)
        if amount is not None:
            pulumi.set(__self__, "amount", amount)
        if billing_account is not None:
            pulumi.set(__self__, "billing_account", billing_account)
        if budget_filter is not None:
            pulumi.set(__self__, "budget_filter", budget_filter)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if threshold_rules is not None:
            pulumi.set(__self__, "threshold_rules", threshold_rules)

    @property
    @pulumi.getter(name="allUpdatesRule")
    def all_updates_rule(self) -> Optional[pulumi.Input['BudgetAllUpdatesRuleArgs']]:
        """
        Defines notifications that are sent on every update to the
        billing account's spend, regardless of the thresholds defined
        using threshold rules.
        Structure is documented below.
        """
        return pulumi.get(self, "all_updates_rule")

    @all_updates_rule.setter
    def all_updates_rule(self, value: Optional[pulumi.Input['BudgetAllUpdatesRuleArgs']]):
        pulumi.set(self, "all_updates_rule", value)

    @property
    @pulumi.getter
    def amount(self) -> Optional[pulumi.Input['BudgetAmountArgs']]:
        """
        The budgeted amount for each usage period.
        Structure is documented below.
        """
        return pulumi.get(self, "amount")

    @amount.setter
    def amount(self, value: Optional[pulumi.Input['BudgetAmountArgs']]):
        pulumi.set(self, "amount", value)

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the billing account to set a budget on.
        """
        return pulumi.get(self, "billing_account")

    @billing_account.setter
    def billing_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_account", value)

    @property
    @pulumi.getter(name="budgetFilter")
    def budget_filter(self) -> Optional[pulumi.Input['BudgetBudgetFilterArgs']]:
        """
        Filters that define which resources are used to compute the actual
        spend against the budget.
        Structure is documented below.
        """
        return pulumi.get(self, "budget_filter")

    @budget_filter.setter
    def budget_filter(self, value: Optional[pulumi.Input['BudgetBudgetFilterArgs']]):
        pulumi.set(self, "budget_filter", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        User data for display name in UI. Must be <= 60 chars.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the budget. The resource name
        implies the scope of a budget. Values are of the form
        billingAccounts/{billingAccountId}/budgets/{budgetId}.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="thresholdRules")
    def threshold_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]]]:
        """
        Rules that trigger alerts (notifications of thresholds being
        crossed) when spend exceeds the specified percentages of the
        budget.
        Structure is documented below.
        """
        return pulumi.get(self, "threshold_rules")

    @threshold_rules.setter
    def threshold_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BudgetThresholdRuleArgs']]]]):
        pulumi.set(self, "threshold_rules", value)


class Budget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_updates_rule: Optional[pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']]] = None,
                 amount: Optional[pulumi.Input[pulumi.InputType['BudgetAmountArgs']]] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 budget_filter: Optional[pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]]] = None,
                 __props__=None):
        """
        Budget configuration for a billing account.

        To get more information about Budget, see:

        * [API documentation](https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets)
        * How-to Guides
            * [Creating a budget](https://cloud.google.com/billing/docs/how-to/budgets)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Billing Budgets API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Billing Budget Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[gcp.billing.BudgetThresholdRuleArgs(
                threshold_percent=0.5,
            )])
        ```
        ### Billing Budget Lastperiod

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
            ),
            amount=gcp.billing.BudgetAmountArgs(
                last_period_amount=True,
            ),
            threshold_rules=[gcp.billing.BudgetThresholdRuleArgs(
                threshold_percent=10,
            )])
        ```
        ### Billing Budget Filter

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
                credit_types_treatment="INCLUDE_SPECIFIED_CREDITS",
                services=["services/24E6-581D-38E5"],
                credit_types=[
                    "PROMOTION",
                    "FREE_TIER",
                ],
                resource_ancestors=["organizations/123456789"],
            ),
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.5,
                ),
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.9,
                    spend_basis="FORECASTED_SPEND",
                ),
            ])
        ```
        ### Billing Budget Notify

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        notification_channel = gcp.monitoring.NotificationChannel("notificationChannel",
            display_name="Example Notification Channel",
            type="email",
            labels={
                "email_address": "address@example.com",
            })
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
            ),
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=1,
                ),
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=1,
                    spend_basis="FORECASTED_SPEND",
                ),
            ],
            all_updates_rule=gcp.billing.BudgetAllUpdatesRuleArgs(
                monitoring_notification_channels=[notification_channel.id],
                disable_default_iam_recipients=True,
            ))
        ```
        ### Billing Budget Customperiod

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
                credit_types_treatment="EXCLUDE_ALL_CREDITS",
                services=["services/24E6-581D-38E5"],
                custom_period=gcp.billing.BudgetBudgetFilterCustomPeriodArgs(
                    start_date=gcp.billing.BudgetBudgetFilterCustomPeriodStartDateArgs(
                        year=2022,
                        month=1,
                        day=1,
                    ),
                    end_date=gcp.billing.BudgetBudgetFilterCustomPeriodEndDateArgs(
                        year=2023,
                        month=12,
                        day=31,
                    ),
                ),
            ),
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.5,
                ),
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.9,
                ),
            ])
        ```

        ## Import

        Budget can be imported using any of these accepted formats* `billingAccounts/{{billing_account}}/budgets/{{name}}` * `{{billing_account}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Budget can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:billing/budget:Budget default billingAccounts/{{billing_account}}/budgets/{{name}}
        ```

        ```sh
         $ pulumi import gcp:billing/budget:Budget default {{billing_account}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:billing/budget:Budget default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']] all_updates_rule: Defines notifications that are sent on every update to the
               billing account's spend, regardless of the thresholds defined
               using threshold rules.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['BudgetAmountArgs']] amount: The budgeted amount for each usage period.
               Structure is documented below.
        :param pulumi.Input[str] billing_account: ID of the billing account to set a budget on.
        :param pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']] budget_filter: Filters that define which resources are used to compute the actual
               spend against the budget.
               Structure is documented below.
        :param pulumi.Input[str] display_name: User data for display name in UI. Must be <= 60 chars.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]] threshold_rules: Rules that trigger alerts (notifications of thresholds being
               crossed) when spend exceeds the specified percentages of the
               budget.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BudgetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Budget configuration for a billing account.

        To get more information about Budget, see:

        * [API documentation](https://cloud.google.com/billing/docs/reference/budget/rest/v1/billingAccounts.budgets)
        * How-to Guides
            * [Creating a budget](https://cloud.google.com/billing/docs/how-to/budgets)

        > **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
        you must specify a `billing_project` and set `user_project_override` to true
        in the provider configuration. Otherwise the Billing Budgets API will return a 403 error.
        Your account must have the `serviceusage.services.use` permission on the
        `billing_project` you defined.

        ## Example Usage
        ### Billing Budget Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[gcp.billing.BudgetThresholdRuleArgs(
                threshold_percent=0.5,
            )])
        ```
        ### Billing Budget Lastperiod

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
            ),
            amount=gcp.billing.BudgetAmountArgs(
                last_period_amount=True,
            ),
            threshold_rules=[gcp.billing.BudgetThresholdRuleArgs(
                threshold_percent=10,
            )])
        ```
        ### Billing Budget Filter

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
                credit_types_treatment="INCLUDE_SPECIFIED_CREDITS",
                services=["services/24E6-581D-38E5"],
                credit_types=[
                    "PROMOTION",
                    "FREE_TIER",
                ],
                resource_ancestors=["organizations/123456789"],
            ),
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.5,
                ),
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.9,
                    spend_basis="FORECASTED_SPEND",
                ),
            ])
        ```
        ### Billing Budget Notify

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        notification_channel = gcp.monitoring.NotificationChannel("notificationChannel",
            display_name="Example Notification Channel",
            type="email",
            labels={
                "email_address": "address@example.com",
            })
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
            ),
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=1,
                ),
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=1,
                    spend_basis="FORECASTED_SPEND",
                ),
            ],
            all_updates_rule=gcp.billing.BudgetAllUpdatesRuleArgs(
                monitoring_notification_channels=[notification_channel.id],
                disable_default_iam_recipients=True,
            ))
        ```
        ### Billing Budget Customperiod

        ```python
        import pulumi
        import pulumi_gcp as gcp

        account = gcp.organizations.get_billing_account(billing_account="000000-0000000-0000000-000000")
        project = gcp.organizations.get_project()
        budget = gcp.billing.Budget("budget",
            billing_account=account.id,
            display_name="Example Billing Budget",
            budget_filter=gcp.billing.BudgetBudgetFilterArgs(
                projects=[f"projects/{project.number}"],
                credit_types_treatment="EXCLUDE_ALL_CREDITS",
                services=["services/24E6-581D-38E5"],
                custom_period=gcp.billing.BudgetBudgetFilterCustomPeriodArgs(
                    start_date=gcp.billing.BudgetBudgetFilterCustomPeriodStartDateArgs(
                        year=2022,
                        month=1,
                        day=1,
                    ),
                    end_date=gcp.billing.BudgetBudgetFilterCustomPeriodEndDateArgs(
                        year=2023,
                        month=12,
                        day=31,
                    ),
                ),
            ),
            amount=gcp.billing.BudgetAmountArgs(
                specified_amount=gcp.billing.BudgetAmountSpecifiedAmountArgs(
                    currency_code="USD",
                    units="100000",
                ),
            ),
            threshold_rules=[
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.5,
                ),
                gcp.billing.BudgetThresholdRuleArgs(
                    threshold_percent=0.9,
                ),
            ])
        ```

        ## Import

        Budget can be imported using any of these accepted formats* `billingAccounts/{{billing_account}}/budgets/{{name}}` * `{{billing_account}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Budget can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:billing/budget:Budget default billingAccounts/{{billing_account}}/budgets/{{name}}
        ```

        ```sh
         $ pulumi import gcp:billing/budget:Budget default {{billing_account}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:billing/budget:Budget default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param BudgetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BudgetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 all_updates_rule: Optional[pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']]] = None,
                 amount: Optional[pulumi.Input[pulumi.InputType['BudgetAmountArgs']]] = None,
                 billing_account: Optional[pulumi.Input[str]] = None,
                 budget_filter: Optional[pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BudgetArgs.__new__(BudgetArgs)

            __props__.__dict__["all_updates_rule"] = all_updates_rule
            if amount is None and not opts.urn:
                raise TypeError("Missing required property 'amount'")
            __props__.__dict__["amount"] = amount
            if billing_account is None and not opts.urn:
                raise TypeError("Missing required property 'billing_account'")
            __props__.__dict__["billing_account"] = billing_account
            __props__.__dict__["budget_filter"] = budget_filter
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["threshold_rules"] = threshold_rules
            __props__.__dict__["name"] = None
        super(Budget, __self__).__init__(
            'gcp:billing/budget:Budget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            all_updates_rule: Optional[pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']]] = None,
            amount: Optional[pulumi.Input[pulumi.InputType['BudgetAmountArgs']]] = None,
            billing_account: Optional[pulumi.Input[str]] = None,
            budget_filter: Optional[pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            threshold_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]]] = None) -> 'Budget':
        """
        Get an existing Budget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BudgetAllUpdatesRuleArgs']] all_updates_rule: Defines notifications that are sent on every update to the
               billing account's spend, regardless of the thresholds defined
               using threshold rules.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['BudgetAmountArgs']] amount: The budgeted amount for each usage period.
               Structure is documented below.
        :param pulumi.Input[str] billing_account: ID of the billing account to set a budget on.
        :param pulumi.Input[pulumi.InputType['BudgetBudgetFilterArgs']] budget_filter: Filters that define which resources are used to compute the actual
               spend against the budget.
               Structure is documented below.
        :param pulumi.Input[str] display_name: User data for display name in UI. Must be <= 60 chars.
        :param pulumi.Input[str] name: Resource name of the budget. The resource name
               implies the scope of a budget. Values are of the form
               billingAccounts/{billingAccountId}/budgets/{budgetId}.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BudgetThresholdRuleArgs']]]] threshold_rules: Rules that trigger alerts (notifications of thresholds being
               crossed) when spend exceeds the specified percentages of the
               budget.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BudgetState.__new__(_BudgetState)

        __props__.__dict__["all_updates_rule"] = all_updates_rule
        __props__.__dict__["amount"] = amount
        __props__.__dict__["billing_account"] = billing_account
        __props__.__dict__["budget_filter"] = budget_filter
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["name"] = name
        __props__.__dict__["threshold_rules"] = threshold_rules
        return Budget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allUpdatesRule")
    def all_updates_rule(self) -> pulumi.Output[Optional['outputs.BudgetAllUpdatesRule']]:
        """
        Defines notifications that are sent on every update to the
        billing account's spend, regardless of the thresholds defined
        using threshold rules.
        Structure is documented below.
        """
        return pulumi.get(self, "all_updates_rule")

    @property
    @pulumi.getter
    def amount(self) -> pulumi.Output['outputs.BudgetAmount']:
        """
        The budgeted amount for each usage period.
        Structure is documented below.
        """
        return pulumi.get(self, "amount")

    @property
    @pulumi.getter(name="billingAccount")
    def billing_account(self) -> pulumi.Output[str]:
        """
        ID of the billing account to set a budget on.
        """
        return pulumi.get(self, "billing_account")

    @property
    @pulumi.getter(name="budgetFilter")
    def budget_filter(self) -> pulumi.Output['outputs.BudgetBudgetFilter']:
        """
        Filters that define which resources are used to compute the actual
        spend against the budget.
        Structure is documented below.
        """
        return pulumi.get(self, "budget_filter")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        User data for display name in UI. Must be <= 60 chars.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the budget. The resource name
        implies the scope of a budget. Values are of the form
        billingAccounts/{billingAccountId}/budgets/{budgetId}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="thresholdRules")
    def threshold_rules(self) -> pulumi.Output[Optional[Sequence['outputs.BudgetThresholdRule']]]:
        """
        Rules that trigger alerts (notifications of thresholds being
        crossed) when spend exceeds the specified percentages of the
        budget.
        Structure is documented below.
        """
        return pulumi.get(self, "threshold_rules")

