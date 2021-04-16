# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceIamPolicyArgs', 'InstanceIamPolicy']

@pulumi.input_type
class InstanceIamPolicyArgs:
    def __init__(__self__, *,
                 instance: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceIamPolicy resource.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, a default will be supplied.
        """
        pulumi.set(__self__, "instance", instance)
        pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The name or relative resource id of the instance to manage IAM policies for.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the instance belongs. If it
        is not provided, a default will be supplied.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _InstanceIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the instances's IAM policy.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, a default will be supplied.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the instances's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        The name or relative resource id of the instance to manage IAM policies for.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the instance belongs. If it
        is not provided, a default will be supplied.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class InstanceIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:

        * `bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
        * `bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `bigtable.InstanceIamBinding` and `bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `bigtable.InstanceIamPolicy` replaces the entire policy.

        > **Note:** `bigtable.InstanceIamBinding` resources **can be** used in conjunction with `bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_bigtable\_instance\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/bigtable.user",
            members=["user:jane@example.com"],
        )])
        editor = gcp.bigtable.InstanceIamPolicy("editor",
            project="your-project",
            instance="your-bigtable-instance",
            policy_data=admin.policy_data)
        ```

        ## google\_bigtable\_instance\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.InstanceIamBinding("editor",
            instance="your-bigtable-instance",
            members=["user:jane@example.com"],
            role="roles/bigtable.user")
        ```

        ## google\_bigtable\_instance\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.InstanceIamMember("editor",
            instance="your-bigtable-instance",
            member="user:jane@example.com",
            role="roles/bigtable.user")
        ```

        ## Import

        Instance IAM resources can be imported using the project, instance name, role and/or member.

        ```sh
         $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance}"
        ```

        ```sh
         $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance} roles/editor"
        ```

        ```sh
         $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, a default will be supplied.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage IAM policies on bigtable instances. Each of these resources serves a different use case:

        * `bigtable.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
        * `bigtable.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
        * `bigtable.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.

        > **Note:** `bigtable.InstanceIamPolicy` **cannot** be used in conjunction with `bigtable.InstanceIamBinding` and `bigtable.InstanceIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the instance as `bigtable.InstanceIamPolicy` replaces the entire policy.

        > **Note:** `bigtable.InstanceIamBinding` resources **can be** used in conjunction with `bigtable.InstanceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_bigtable\_instance\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/bigtable.user",
            members=["user:jane@example.com"],
        )])
        editor = gcp.bigtable.InstanceIamPolicy("editor",
            project="your-project",
            instance="your-bigtable-instance",
            policy_data=admin.policy_data)
        ```

        ## google\_bigtable\_instance\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.InstanceIamBinding("editor",
            instance="your-bigtable-instance",
            members=["user:jane@example.com"],
            role="roles/bigtable.user")
        ```

        ## google\_bigtable\_instance\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.bigtable.InstanceIamMember("editor",
            instance="your-bigtable-instance",
            member="user:jane@example.com",
            role="roles/bigtable.user")
        ```

        ## Import

        Instance IAM resources can be imported using the project, instance name, role and/or member.

        ```sh
         $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance}"
        ```

        ```sh
         $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance} roles/editor"
        ```

        ```sh
         $ pulumi import gcp:bigtable/instanceIamPolicy:InstanceIamPolicy editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param InstanceIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceIamPolicyArgs.__new__(InstanceIamPolicyArgs)

            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__.__dict__["instance"] = instance
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["etag"] = None
        super(InstanceIamPolicy, __self__).__init__(
            'gcp:bigtable/instanceIamPolicy:InstanceIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'InstanceIamPolicy':
        """
        Get an existing InstanceIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the instances's IAM policy.
        :param pulumi.Input[str] instance: The name or relative resource id of the instance to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The project in which the instance belongs. If it
               is not provided, a default will be supplied.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceIamPolicyState.__new__(_InstanceIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["instance"] = instance
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        return InstanceIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the instances's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name or relative resource id of the instance to manage IAM policies for.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project in which the instance belongs. If it
        is not provided, a default will be supplied.
        """
        return pulumi.get(self, "project")

