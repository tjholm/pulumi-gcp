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

__all__ = ['SubnetworkIAMMemberArgs', 'SubnetworkIAMMember']

@pulumi.input_type
class SubnetworkIAMMemberArgs:
    def __init__(__self__, *,
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 subnetwork: pulumi.Input[str],
                 condition: Optional[pulumi.Input['SubnetworkIAMMemberConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubnetworkIAMMember resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input['SubnetworkIAMMemberConditionArgs'] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "subnetwork", subnetwork)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def member(self) -> pulumi.Input[str]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: pulumi.Input[str]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def subnetwork(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "subnetwork")

    @subnetwork.setter
    def subnetwork(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnetwork", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['SubnetworkIAMMemberConditionArgs']]:
        """
        ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['SubnetworkIAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The GCP region for this subnetwork.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _SubnetworkIAMMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['SubnetworkIAMMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SubnetworkIAMMember resources.
        :param pulumi.Input['SubnetworkIAMMemberConditionArgs'] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: Used to find the parent resource to bind the IAM policy to
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if subnetwork is not None:
            pulumi.set(__self__, "subnetwork", subnetwork)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['SubnetworkIAMMemberConditionArgs']]:
        """
        ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['SubnetworkIAMMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The GCP region for this subnetwork.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def subnetwork(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "subnetwork")

    @subnetwork.setter
    def subnetwork(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnetwork", value)


class SubnetworkIAMMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['SubnetworkIAMMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Compute Engine Subnetwork. Each of these resources serves a different use case:

        * `compute.SubnetworkIAMPolicy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
        * `compute.SubnetworkIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
        * `compute.SubnetworkIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.

        > **Note:** `compute.SubnetworkIAMPolicy` **cannot** be used in conjunction with `compute.SubnetworkIAMBinding` and `compute.SubnetworkIAMMember` or they will fight over what your policy should be.

        > **Note:** `compute.SubnetworkIAMBinding` resources **can be** used in conjunction with `compute.SubnetworkIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_compute\\_subnetwork\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
        )])
        policy = gcp.compute.SubnetworkIAMPolicy("policy",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ),
        )])
        policy = gcp.compute.SubnetworkIAMPolicy("policy",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            policy_data=admin.policy_data)
        ```
        ## google\\_compute\\_subnetwork\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.SubnetworkIAMBinding("binding",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.SubnetworkIAMBinding("binding",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
            condition=gcp.compute.SubnetworkIAMBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```
        ## google\\_compute\\_subnetwork\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.SubnetworkIAMMember("member",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.SubnetworkIAMMember("member",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            member="user:jane@example.com",
            condition=gcp.compute.SubnetworkIAMMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/regions/{{region}}/subnetworks/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine subnetwork IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SubnetworkIAMMemberConditionArgs']] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: Used to find the parent resource to bind the IAM policy to
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetworkIAMMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Compute Engine Subnetwork. Each of these resources serves a different use case:

        * `compute.SubnetworkIAMPolicy`: Authoritative. Sets the IAM policy for the subnetwork and replaces any existing policy already attached.
        * `compute.SubnetworkIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the subnetwork are preserved.
        * `compute.SubnetworkIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the subnetwork are preserved.

        > **Note:** `compute.SubnetworkIAMPolicy` **cannot** be used in conjunction with `compute.SubnetworkIAMBinding` and `compute.SubnetworkIAMMember` or they will fight over what your policy should be.

        > **Note:** `compute.SubnetworkIAMBinding` resources **can be** used in conjunction with `compute.SubnetworkIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_compute\\_subnetwork\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
        )])
        policy = gcp.compute.SubnetworkIAMPolicy("policy",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ),
        )])
        policy = gcp.compute.SubnetworkIAMPolicy("policy",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            policy_data=admin.policy_data)
        ```
        ## google\\_compute\\_subnetwork\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.SubnetworkIAMBinding("binding",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.SubnetworkIAMBinding("binding",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            members=["user:jane@example.com"],
            condition=gcp.compute.SubnetworkIAMBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```
        ## google\\_compute\\_subnetwork\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.SubnetworkIAMMember("member",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.SubnetworkIAMMember("member",
            project=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["project"],
            region=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["region"],
            subnetwork=google_compute_subnetwork["network-with-private-secondary-ip-ranges"]["name"],
            role="roles/compute.networkUser",
            member="user:jane@example.com",
            condition=gcp.compute.SubnetworkIAMMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/regions/{{region}}/subnetworks/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine subnetwork IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param SubnetworkIAMMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetworkIAMMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['SubnetworkIAMMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 subnetwork: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetworkIAMMemberArgs.__new__(SubnetworkIAMMemberArgs)

            __props__.__dict__["condition"] = condition
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if subnetwork is None and not opts.urn:
                raise TypeError("Missing required property 'subnetwork'")
            __props__.__dict__["subnetwork"] = subnetwork
            __props__.__dict__["etag"] = None
        super(SubnetworkIAMMember, __self__).__init__(
            'gcp:compute/subnetworkIAMMember:SubnetworkIAMMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['SubnetworkIAMMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            subnetwork: Optional[pulumi.Input[str]] = None) -> 'SubnetworkIAMMember':
        """
        Get an existing SubnetworkIAMMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SubnetworkIAMMemberConditionArgs']] condition: ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The GCP region for this subnetwork.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] subnetwork: Used to find the parent resource to bind the IAM policy to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubnetworkIAMMemberState.__new__(_SubnetworkIAMMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["member"] = member
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        __props__.__dict__["role"] = role
        __props__.__dict__["subnetwork"] = subnetwork
        return SubnetworkIAMMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.SubnetworkIAMMemberCondition']]:
        """
        ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The GCP region for this subnetwork.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def subnetwork(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "subnetwork")

