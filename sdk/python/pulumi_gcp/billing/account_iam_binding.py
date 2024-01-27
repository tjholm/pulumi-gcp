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

__all__ = ['AccountIamBindingArgs', 'AccountIamBinding']

@pulumi.input_type
class AccountIamBindingArgs:
    def __init__(__self__, *,
                 billing_account_id: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['AccountIamBindingConditionArgs']] = None):
        """
        The set of arguments for constructing a AccountIamBinding resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
               
               For `billing.AccountIamMember` or `billing.AccountIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `billing.AccountIamPolicy` only:
        """
        pulumi.set(__self__, "billing_account_id", billing_account_id)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> pulumi.Input[str]:
        """
        The billing account id.

        For `billing.AccountIamMember` or `billing.AccountIamBinding`:

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "billing_account_id")

    @billing_account_id.setter
    def billing_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_account_id", value)

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
        """
        The role that should be applied. Only one
        `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

        `billing.AccountIamPolicy` only:
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['AccountIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['AccountIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _AccountIamBindingState:
    def __init__(__self__, *,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['AccountIamBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AccountIamBinding resources.
        :param pulumi.Input[str] billing_account_id: The billing account id.
               
               For `billing.AccountIamMember` or `billing.AccountIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] etag: (Computed) The etag of the billing account's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `billing.AccountIamPolicy` only:
        """
        if billing_account_id is not None:
            pulumi.set(__self__, "billing_account_id", billing_account_id)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The billing account id.

        For `billing.AccountIamMember` or `billing.AccountIamBinding`:

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "billing_account_id")

    @billing_account_id.setter
    def billing_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_account_id", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['AccountIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['AccountIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the billing account's IAM policy.
        """
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
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

        `billing.AccountIamPolicy` only:
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class AccountIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['AccountIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage IAM policies on billing accounts. Each of these resources serves a different use case:

        * `billing.AccountIamPolicy`: Authoritative. Sets the IAM policy for the billing accounts and replaces any existing policy already attached.
        * `billing.AccountIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `billing.AccountIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role of the billing accounts are preserved.

        > **Note:** `billing.AccountIamPolicy` **cannot** be used in conjunction with `billing.AccountIamBinding` and `billing.AccountIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the billing account as `billing.AccountIamPolicy` replaces the entire policy.

        > **Note:** `billing.AccountIamBinding` resources **can be** used in conjunction with `billing.AccountIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_billing\\_account\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/billing.viewer",
            members=["user:jane@example.com"],
        )])
        editor = gcp.billing.AccountIamPolicy("editor",
            billing_account_id="00AA00-000AAA-00AA0A",
            policy_data=admin.policy_data)
        ```

        ## google\\_billing\\_account\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.billing.AccountIamBinding("editor",
            billing_account_id="00AA00-000AAA-00AA0A",
            members=["user:jane@example.com"],
            role="roles/billing.viewer")
        ```

        ## google\\_billing\\_account\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.billing.AccountIamMember("editor",
            billing_account_id="00AA00-000AAA-00AA0A",
            member="user:jane@example.com",
            role="roles/billing.viewer")
        ```

        ## Import

        ### Importing IAM policies IAM policy imports use the `billing_account_id` identifier of the Billing Account resource only. For example* `{{billing_account_id}}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {

         id = {{billing_account_id}}

         to = google_billing_account_iam_policy.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:billing/accountIamBinding:AccountIamBinding default {{billing_account_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
               
               For `billing.AccountIamMember` or `billing.AccountIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `billing.AccountIamPolicy` only:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage IAM policies on billing accounts. Each of these resources serves a different use case:

        * `billing.AccountIamPolicy`: Authoritative. Sets the IAM policy for the billing accounts and replaces any existing policy already attached.
        * `billing.AccountIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `billing.AccountIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role of the billing accounts are preserved.

        > **Note:** `billing.AccountIamPolicy` **cannot** be used in conjunction with `billing.AccountIamBinding` and `billing.AccountIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the billing account as `billing.AccountIamPolicy` replaces the entire policy.

        > **Note:** `billing.AccountIamBinding` resources **can be** used in conjunction with `billing.AccountIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_billing\\_account\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/billing.viewer",
            members=["user:jane@example.com"],
        )])
        editor = gcp.billing.AccountIamPolicy("editor",
            billing_account_id="00AA00-000AAA-00AA0A",
            policy_data=admin.policy_data)
        ```

        ## google\\_billing\\_account\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.billing.AccountIamBinding("editor",
            billing_account_id="00AA00-000AAA-00AA0A",
            members=["user:jane@example.com"],
            role="roles/billing.viewer")
        ```

        ## google\\_billing\\_account\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.billing.AccountIamMember("editor",
            billing_account_id="00AA00-000AAA-00AA0A",
            member="user:jane@example.com",
            role="roles/billing.viewer")
        ```

        ## Import

        ### Importing IAM policies IAM policy imports use the `billing_account_id` identifier of the Billing Account resource only. For example* `{{billing_account_id}}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {

         id = {{billing_account_id}}

         to = google_billing_account_iam_policy.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:billing/accountIamBinding:AccountIamBinding default {{billing_account_id}}
        ```

        :param str resource_name: The name of the resource.
        :param AccountIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_account_id: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['AccountIamBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountIamBindingArgs.__new__(AccountIamBindingArgs)

            if billing_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'billing_account_id'")
            __props__.__dict__["billing_account_id"] = billing_account_id
            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(AccountIamBinding, __self__).__init__(
            'gcp:billing/accountIamBinding:AccountIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_account_id: Optional[pulumi.Input[str]] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['AccountIamBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'AccountIamBinding':
        """
        Get an existing AccountIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_account_id: The billing account id.
               
               For `billing.AccountIamMember` or `billing.AccountIamBinding`:
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        :param pulumi.Input[str] etag: (Computed) The etag of the billing account's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
               
               `billing.AccountIamPolicy` only:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountIamBindingState.__new__(_AccountIamBindingState)

        __props__.__dict__["billing_account_id"] = billing_account_id
        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["role"] = role
        return AccountIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingAccountId")
    def billing_account_id(self) -> pulumi.Output[str]:
        """
        The billing account id.

        For `billing.AccountIamMember` or `billing.AccountIamBinding`:

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        """
        return pulumi.get(self, "billing_account_id")

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.AccountIamBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the billing account's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `billing.AccountIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).

        `billing.AccountIamPolicy` only:
        """
        return pulumi.get(self, "role")

