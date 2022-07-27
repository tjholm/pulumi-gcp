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

__all__ = ['DicomStoreIamMemberArgs', 'DicomStoreIamMember']

@pulumi.input_type
class DicomStoreIamMemberArgs:
    def __init__(__self__, *,
                 dicom_store_id: pulumi.Input[str],
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['DicomStoreIamMemberConditionArgs']] = None):
        """
        The set of arguments for constructing a DicomStoreIamMember resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        pulumi.set(__self__, "dicom_store_id", dicom_store_id)
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter(name="dicomStoreId")
    def dicom_store_id(self) -> pulumi.Input[str]:
        """
        The DICOM store ID, in the form
        `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        project setting will be used as a fallback.
        """
        return pulumi.get(self, "dicom_store_id")

    @dicom_store_id.setter
    def dicom_store_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dicom_store_id", value)

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
        `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['DicomStoreIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['DicomStoreIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _DicomStoreIamMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['DicomStoreIamMemberConditionArgs']] = None,
                 dicom_store_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DicomStoreIamMember resources.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the DICOM store's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if dicom_store_id is not None:
            pulumi.set(__self__, "dicom_store_id", dicom_store_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['DicomStoreIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['DicomStoreIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="dicomStoreId")
    def dicom_store_id(self) -> Optional[pulumi.Input[str]]:
        """
        The DICOM store ID, in the form
        `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        project setting will be used as a fallback.
        """
        return pulumi.get(self, "dicom_store_id")

    @dicom_store_id.setter
    def dicom_store_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dicom_store_id", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the DICOM store's IAM policy.
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
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The role that should be applied. Only one
        `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class DicomStoreIamMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DicomStoreIamMemberConditionArgs']]] = None,
                 dicom_store_id: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:

        * `healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
        * `healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
        * `healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.

        > **Note:** `healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `healthcare.DicomStoreIamBinding` and `healthcare.DicomStoreIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_healthcare\\_dicom\\_store\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        dicom_store = gcp.healthcare.DicomStoreIamPolicy("dicomStore",
            dicom_store_id="your-dicom-store-id",
            policy_data=admin.policy_data)
        ```

        ## google\\_healthcare\\_dicom\\_store\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dicom_store = gcp.healthcare.DicomStoreIamBinding("dicomStore",
            dicom_store_id="your-dicom-store-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_healthcare\\_dicom\\_store\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dicom_store = gcp.healthcare.DicomStoreIamMember("dicomStore",
            dicom_store_id="your-dicom-store-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `dicom_store_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `dicom_store_id` and role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `dicom_store_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DicomStoreIamMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:

        * `healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
        * `healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
        * `healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.

        > **Note:** `healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `healthcare.DicomStoreIamBinding` and `healthcare.DicomStoreIamMember` or they will fight over what your policy should be.

        > **Note:** `healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_healthcare\\_dicom\\_store\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        dicom_store = gcp.healthcare.DicomStoreIamPolicy("dicomStore",
            dicom_store_id="your-dicom-store-id",
            policy_data=admin.policy_data)
        ```

        ## google\\_healthcare\\_dicom\\_store\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dicom_store = gcp.healthcare.DicomStoreIamBinding("dicomStore",
            dicom_store_id="your-dicom-store-id",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_healthcare\\_dicom\\_store\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dicom_store = gcp.healthcare.DicomStoreIamMember("dicomStore",
            dicom_store_id="your-dicom-store-id",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `dicom_store_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `dicom_store_id` and role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `dicom_store_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
        ```

        :param str resource_name: The name of the resource.
        :param DicomStoreIamMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DicomStoreIamMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['DicomStoreIamMemberConditionArgs']]] = None,
                 dicom_store_id: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DicomStoreIamMemberArgs.__new__(DicomStoreIamMemberArgs)

            __props__.__dict__["condition"] = condition
            if dicom_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'dicom_store_id'")
            __props__.__dict__["dicom_store_id"] = dicom_store_id
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(DicomStoreIamMember, __self__).__init__(
            'gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['DicomStoreIamMemberConditionArgs']]] = None,
            dicom_store_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'DicomStoreIamMember':
        """
        Get an existing DicomStoreIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the DICOM store's IAM policy.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DicomStoreIamMemberState.__new__(_DicomStoreIamMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["dicom_store_id"] = dicom_store_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["member"] = member
        __props__.__dict__["role"] = role
        return DicomStoreIamMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.DicomStoreIamMemberCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="dicomStoreId")
    def dicom_store_id(self) -> pulumi.Output[str]:
        """
        The DICOM store ID, in the form
        `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        project setting will be used as a fallback.
        """
        return pulumi.get(self, "dicom_store_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the DICOM store's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

