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

__all__ = ['DatasetIamBindingArgs', 'DatasetIamBinding']

@pulumi.input_type
class DatasetIamBindingArgs:
    def __init__(__self__, *,
                 dataset_id: pulumi.Input[str],
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['DatasetIamBindingConditionArgs']] = None):
        """
        The set of arguments for constructing a DatasetIamBinding resource.
        :param pulumi.Input[str] dataset_id: The dataset ID, in the form
               `{project_id}/{location_name}/{dataset_name}` or
               `{location_name}/{dataset_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        """
        The dataset ID, in the form
        `{project_id}/{location_name}/{dataset_name}` or
        `{location_name}/{dataset_name}`. In the second form, the provider's
        project setting will be used as a fallback.
        """
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

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
        `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['DatasetIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['DatasetIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _DatasetIamBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['DatasetIamBindingConditionArgs']] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatasetIamBinding resources.
        :param pulumi.Input[str] dataset_id: The dataset ID, in the form
               `{project_id}/{location_name}/{dataset_name}` or
               `{location_name}/{dataset_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the dataset's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if dataset_id is not None:
            pulumi.set(__self__, "dataset_id", dataset_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['DatasetIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['DatasetIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> Optional[pulumi.Input[str]]:
        """
        The dataset ID, in the form
        `{project_id}/{location_name}/{dataset_name}` or
        `{location_name}/{dataset_name}`. In the second form, the provider's
        project setting will be used as a fallback.
        """
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset_id", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the dataset's IAM policy.
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
        `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class DatasetIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DatasetIamBindingConditionArgs']]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:

        * `healthcare.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
        * `healthcare.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
        * `healthcare.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.

        > **Note:** `healthcare.DatasetIamPolicy` **cannot** be used in conjunction with `healthcare.DatasetIamBinding` and `healthcare.DatasetIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.DatasetIamBinding` resources **can be** used in conjunction with `healthcare.DatasetIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_healthcare\\_dataset\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        dataset = gcp.healthcare.DatasetIamPolicy("dataset",
            dataset_id="your-dataset-id",
            policy_data=admin.policy_data)
        ```

        ## google\\_healthcare\\_dataset\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.DatasetIamBinding("dataset",
            dataset_id="your-dataset-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_healthcare\\_dataset\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.DatasetIamMember("dataset",
            dataset_id="your-dataset-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `dataset_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/datasetIamBinding:DatasetIamBinding dataset_iam "your-project-id/location-name/dataset-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `dataset_id` and role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/datasetIamBinding:DatasetIamBinding dataset_iam "your-project-id/location-name/dataset-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `dataset_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/datasetIamBinding:DatasetIamBinding dataset_iam your-project-id/location-name/dataset-name
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_id: The dataset ID, in the form
               `{project_id}/{location_name}/{dataset_name}` or
               `{location_name}/{dataset_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:

        * `healthcare.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
        * `healthcare.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
        * `healthcare.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.

        > **Note:** `healthcare.DatasetIamPolicy` **cannot** be used in conjunction with `healthcare.DatasetIamBinding` and `healthcare.DatasetIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.DatasetIamBinding` resources **can be** used in conjunction with `healthcare.DatasetIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_healthcare\\_dataset\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        dataset = gcp.healthcare.DatasetIamPolicy("dataset",
            dataset_id="your-dataset-id",
            policy_data=admin.policy_data)
        ```

        ## google\\_healthcare\\_dataset\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.DatasetIamBinding("dataset",
            dataset_id="your-dataset-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_healthcare\\_dataset\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.DatasetIamMember("dataset",
            dataset_id="your-dataset-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `dataset_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/datasetIamBinding:DatasetIamBinding dataset_iam "your-project-id/location-name/dataset-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `dataset_id` and role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/datasetIamBinding:DatasetIamBinding dataset_iam "your-project-id/location-name/dataset-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `dataset_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/datasetIamBinding:DatasetIamBinding dataset_iam your-project-id/location-name/dataset-name
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param DatasetIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DatasetIamBindingConditionArgs']]] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatasetIamBindingArgs.__new__(DatasetIamBindingArgs)

            __props__.__dict__["condition"] = condition
            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__.__dict__["dataset_id"] = dataset_id
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(DatasetIamBinding, __self__).__init__(
            'gcp:healthcare/datasetIamBinding:DatasetIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['DatasetIamBindingConditionArgs']]] = None,
            dataset_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'DatasetIamBinding':
        """
        Get an existing DatasetIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset_id: The dataset ID, in the form
               `{project_id}/{location_name}/{dataset_name}` or
               `{location_name}/{dataset_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the dataset's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatasetIamBindingState.__new__(_DatasetIamBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["dataset_id"] = dataset_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["role"] = role
        return DatasetIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.DatasetIamBindingCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Output[str]:
        """
        The dataset ID, in the form
        `{project_id}/{location_name}/{dataset_name}` or
        `{location_name}/{dataset_name}`. In the second form, the provider's
        project setting will be used as a fallback.
        """
        return pulumi.get(self, "dataset_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the dataset's IAM policy.
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
        `healthcare.DatasetIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

