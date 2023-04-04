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

__all__ = ['DataPolicyIamBindingArgs', 'DataPolicyIamBinding']

@pulumi.input_type
class DataPolicyIamBindingArgs:
    def __init__(__self__, *,
                 data_policy_id: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['DataPolicyIamBindingConditionArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataPolicyIamBinding resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] location: The name of the location of the data policy.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "data_policy_id", data_policy_id)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="dataPolicyId")
    def data_policy_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "data_policy_id")

    @data_policy_id.setter
    def data_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_policy_id", value)

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
        `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['DataPolicyIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['DataPolicyIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the location of the data policy.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

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


@pulumi.input_type
class _DataPolicyIamBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['DataPolicyIamBindingConditionArgs']] = None,
                 data_policy_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataPolicyIamBinding resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The name of the location of the data policy.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if data_policy_id is not None:
            pulumi.set(__self__, "data_policy_id", data_policy_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['DataPolicyIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['DataPolicyIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="dataPolicyId")
    def data_policy_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "data_policy_id")

    @data_policy_id.setter
    def data_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_policy_id", value)

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
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the location of the data policy.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

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
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class DataPolicyIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DataPolicyIamBindingConditionArgs']]] = None,
                 data_policy_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for BigQuery Data Policy DataPolicy. Each of these resources serves a different use case:

        * `bigquerydatapolicy.DataPolicyIamPolicy`: Authoritative. Sets the IAM policy for the datapolicy and replaces any existing policy already attached.
        * `bigquerydatapolicy.DataPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datapolicy are preserved.
        * `bigquerydatapolicy.DataPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datapolicy are preserved.

        > **Note:** `bigquerydatapolicy.DataPolicyIamPolicy` **cannot** be used in conjunction with `bigquerydatapolicy.DataPolicyIamBinding` and `bigquerydatapolicy.DataPolicyIamMember` or they will fight over what your policy should be.

        > **Note:** `bigquerydatapolicy.DataPolicyIamBinding` resources **can be** used in conjunction with `bigquerydatapolicy.DataPolicyIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_bigquery\\_datapolicy\\_data\\_policy\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.bigquerydatapolicy.DataPolicyIamPolicy("policy",
            project=google_bigquery_datapolicy_data_policy["data_policy"]["project"],
            location=google_bigquery_datapolicy_data_policy["data_policy"]["location"],
            data_policy_id=google_bigquery_datapolicy_data_policy["data_policy"]["data_policy_id"],
            policy_data=admin.policy_data)
        ```

        ## google\\_bigquery\\_datapolicy\\_data\\_policy\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquerydatapolicy.DataPolicyIamBinding("binding",
            project=google_bigquery_datapolicy_data_policy["data_policy"]["project"],
            location=google_bigquery_datapolicy_data_policy["data_policy"]["location"],
            data_policy_id=google_bigquery_datapolicy_data_policy["data_policy"]["data_policy_id"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_bigquery\\_datapolicy\\_data\\_policy\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquerydatapolicy.DataPolicyIamMember("member",
            project=google_bigquery_datapolicy_data_policy["data_policy"]["project"],
            location=google_bigquery_datapolicy_data_policy["data_policy"]["location"],
            data_policy_id=google_bigquery_datapolicy_data_policy["data_policy"]["data_policy_id"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} * {{project}}/{{location}}/{{data_policy_id}} * {{location}}/{{data_policy_id}} * {{data_policy_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Data Policy datapolicy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The name of the location of the data policy.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataPolicyIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for BigQuery Data Policy DataPolicy. Each of these resources serves a different use case:

        * `bigquerydatapolicy.DataPolicyIamPolicy`: Authoritative. Sets the IAM policy for the datapolicy and replaces any existing policy already attached.
        * `bigquerydatapolicy.DataPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datapolicy are preserved.
        * `bigquerydatapolicy.DataPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datapolicy are preserved.

        > **Note:** `bigquerydatapolicy.DataPolicyIamPolicy` **cannot** be used in conjunction with `bigquerydatapolicy.DataPolicyIamBinding` and `bigquerydatapolicy.DataPolicyIamMember` or they will fight over what your policy should be.

        > **Note:** `bigquerydatapolicy.DataPolicyIamBinding` resources **can be** used in conjunction with `bigquerydatapolicy.DataPolicyIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_bigquery\\_datapolicy\\_data\\_policy\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.bigquerydatapolicy.DataPolicyIamPolicy("policy",
            project=google_bigquery_datapolicy_data_policy["data_policy"]["project"],
            location=google_bigquery_datapolicy_data_policy["data_policy"]["location"],
            data_policy_id=google_bigquery_datapolicy_data_policy["data_policy"]["data_policy_id"],
            policy_data=admin.policy_data)
        ```

        ## google\\_bigquery\\_datapolicy\\_data\\_policy\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquerydatapolicy.DataPolicyIamBinding("binding",
            project=google_bigquery_datapolicy_data_policy["data_policy"]["project"],
            location=google_bigquery_datapolicy_data_policy["data_policy"]["location"],
            data_policy_id=google_bigquery_datapolicy_data_policy["data_policy"]["data_policy_id"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_bigquery\\_datapolicy\\_data\\_policy\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquerydatapolicy.DataPolicyIamMember("member",
            project=google_bigquery_datapolicy_data_policy["data_policy"]["project"],
            location=google_bigquery_datapolicy_data_policy["data_policy"]["location"],
            data_policy_id=google_bigquery_datapolicy_data_policy["data_policy"]["data_policy_id"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} * {{project}}/{{location}}/{{data_policy_id}} * {{location}}/{{data_policy_id}} * {{data_policy_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Data Policy datapolicy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param DataPolicyIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataPolicyIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DataPolicyIamBindingConditionArgs']]] = None,
                 data_policy_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataPolicyIamBindingArgs.__new__(DataPolicyIamBindingArgs)

            __props__.__dict__["condition"] = condition
            if data_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'data_policy_id'")
            __props__.__dict__["data_policy_id"] = data_policy_id
            __props__.__dict__["location"] = location
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(DataPolicyIamBinding, __self__).__init__(
            'gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['DataPolicyIamBindingConditionArgs']]] = None,
            data_policy_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'DataPolicyIamBinding':
        """
        Get an existing DataPolicyIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The name of the location of the data policy.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataPolicyIamBindingState.__new__(_DataPolicyIamBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["data_policy_id"] = data_policy_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["location"] = location
        __props__.__dict__["members"] = members
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        return DataPolicyIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.DataPolicyIamBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="dataPolicyId")
    def data_policy_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "data_policy_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The name of the location of the data policy.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

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
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

