# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DicomStoreIamPolicyArgs', 'DicomStoreIamPolicy']

@pulumi.input_type
class DicomStoreIamPolicyArgs:
    def __init__(__self__, *,
                 dicom_store_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a DicomStoreIamPolicy resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        """
        pulumi.set(__self__, "dicom_store_id", dicom_store_id)
        pulumi.set(__self__, "policy_data", policy_data)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _DicomStoreIamPolicyState:
    def __init__(__self__, *,
                 dicom_store_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DicomStoreIamPolicy resources.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the DICOM store's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        """
        if dicom_store_id is not None:
            pulumi.set(__self__, "dicom_store_id", dicom_store_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)


class DicomStoreIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dicom_store_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
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
         $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `dicom_store_id` and role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `dicom_store_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DicomStoreIamPolicyArgs,
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
         $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `dicom_store_id` and role, e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `dicom_store_id`, role, and account e.g.

        ```sh
         $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
        ```

        :param str resource_name: The name of the resource.
        :param DicomStoreIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DicomStoreIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dicom_store_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DicomStoreIamPolicyArgs.__new__(DicomStoreIamPolicyArgs)

            if dicom_store_id is None and not opts.urn:
                raise TypeError("Missing required property 'dicom_store_id'")
            __props__.__dict__["dicom_store_id"] = dicom_store_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["etag"] = None
        super(DicomStoreIamPolicy, __self__).__init__(
            'gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dicom_store_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'DicomStoreIamPolicy':
        """
        Get an existing DicomStoreIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dicom_store_id: The DICOM store ID, in the form
               `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
               `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
               project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the DICOM store's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DicomStoreIamPolicyState.__new__(_DicomStoreIamPolicyState)

        __props__.__dict__["dicom_store_id"] = dicom_store_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        return DicomStoreIamPolicy(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

