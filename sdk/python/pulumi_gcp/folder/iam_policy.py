# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IAMPolicyArgs', 'IAMPolicy']

@pulumi.input_type
class IAMPolicyArgs:
    def __init__(__self__, *,
                 folder: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a IAMPolicy resource.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        pulumi.set(__self__, "folder", folder)
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Input[str]:
        """
        The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The `organizations.get_iam_policy` data source that represents
        the IAM policy that will be applied to the folder. The policy will be
        merged with any existing policy applied to the folder.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _IAMPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if folder is not None:
            pulumi.set(__self__, "folder", folder)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the folder's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def folder(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @folder.setter
    def folder(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "folder", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The `organizations.get_iam_policy` data source that represents
        the IAM policy that will be applied to the folder. The policy will be
        merged with any existing policy applied to the folder.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)


class IAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `folder`, role, and member e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `folder` and role, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `folder`.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder folder
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `folder`, role, and member e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `folder` and role, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `folder`.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder folder
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:folder/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param IAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 folder: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMPolicyArgs.__new__(IAMPolicyArgs)

            if folder is None and not opts.urn:
                raise TypeError("Missing required property 'folder'")
            __props__.__dict__["folder"] = folder
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["etag"] = None
        super(IAMPolicy, __self__).__init__(
            'gcp:folder/iAMPolicy:IAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            folder: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'IAMPolicy':
        """
        Get an existing IAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the folder's IAM policy.
        :param pulumi.Input[str] folder: The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        :param pulumi.Input[str] policy_data: The `organizations.get_iam_policy` data source that represents
               the IAM policy that will be applied to the folder. The policy will be
               merged with any existing policy applied to the folder.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMPolicyState.__new__(_IAMPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["folder"] = folder
        __props__.__dict__["policy_data"] = policy_data
        return IAMPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the folder's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[str]:
        """
        The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The `organizations.get_iam_policy` data source that represents
        the IAM policy that will be applied to the folder. The policy will be
        merged with any existing policy applied to the folder.
        """
        return pulumi.get(self, "policy_data")

