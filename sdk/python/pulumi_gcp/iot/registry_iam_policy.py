# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RegistryIamPolicyArgs', 'RegistryIamPolicy']

@pulumi.input_type
class RegistryIamPolicyArgs:
    def __init__(__self__, *,
                 policy_data: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegistryIamPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        pulumi.set(__self__, "policy_data", policy_data)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

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

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
        The region in which the created registry should reside.
        If it is not provided, the provider region is used.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _RegistryIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RegistryIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
        The region in which the created registry should reside.
        If it is not provided, the provider region is used.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class RegistryIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Cloud IoT Core DeviceRegistry. Each of these resources serves a different use case:

        * `iot.RegistryIamPolicy`: Authoritative. Sets the IAM policy for the deviceregistry and replaces any existing policy already attached.
        * `iot.RegistryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the deviceregistry are preserved.
        * `iot.RegistryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the deviceregistry are preserved.

        > **Note:** `iot.RegistryIamPolicy` **cannot** be used in conjunction with `iot.RegistryIamBinding` and `iot.RegistryIamMember` or they will fight over what your policy should be.

        > **Note:** `iot.RegistryIamBinding` resources **can be** used in conjunction with `iot.RegistryIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_cloudiot\\_registry\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.iot.RegistryIamPolicy("policy",
            project=google_cloudiot_registry["test-registry"]["project"],
            region=google_cloudiot_registry["test-registry"]["region"],
            policy_data=admin.policy_data)
        ```

        ## google\\_cloudiot\\_registry\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iot.RegistryIamBinding("binding",
            project=google_cloudiot_registry["test-registry"]["project"],
            region=google_cloudiot_registry["test-registry"]["region"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_cloudiot\\_registry\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iot.RegistryIamMember("member",
            project=google_cloudiot_registry["test-registry"]["project"],
            region=google_cloudiot_registry["test-registry"]["region"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/registries/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud IoT Core deviceregistry IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:iot/registryIamPolicy:RegistryIamPolicy editor "projects/{{project}}/locations/{{location}}/registries/{{device_registry}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:iot/registryIamPolicy:RegistryIamPolicy editor "projects/{{project}}/locations/{{location}}/registries/{{device_registry}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:iot/registryIamPolicy:RegistryIamPolicy editor projects/{{project}}/locations/{{location}}/registries/{{device_registry}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Cloud IoT Core DeviceRegistry. Each of these resources serves a different use case:

        * `iot.RegistryIamPolicy`: Authoritative. Sets the IAM policy for the deviceregistry and replaces any existing policy already attached.
        * `iot.RegistryIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the deviceregistry are preserved.
        * `iot.RegistryIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the deviceregistry are preserved.

        > **Note:** `iot.RegistryIamPolicy` **cannot** be used in conjunction with `iot.RegistryIamBinding` and `iot.RegistryIamMember` or they will fight over what your policy should be.

        > **Note:** `iot.RegistryIamBinding` resources **can be** used in conjunction with `iot.RegistryIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_cloudiot\\_registry\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.iot.RegistryIamPolicy("policy",
            project=google_cloudiot_registry["test-registry"]["project"],
            region=google_cloudiot_registry["test-registry"]["region"],
            policy_data=admin.policy_data)
        ```

        ## google\\_cloudiot\\_registry\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iot.RegistryIamBinding("binding",
            project=google_cloudiot_registry["test-registry"]["project"],
            region=google_cloudiot_registry["test-registry"]["region"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_cloudiot\\_registry\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iot.RegistryIamMember("member",
            project=google_cloudiot_registry["test-registry"]["project"],
            region=google_cloudiot_registry["test-registry"]["region"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/registries/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud IoT Core deviceregistry IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:iot/registryIamPolicy:RegistryIamPolicy editor "projects/{{project}}/locations/{{location}}/registries/{{device_registry}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:iot/registryIamPolicy:RegistryIamPolicy editor "projects/{{project}}/locations/{{location}}/registries/{{device_registry}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:iot/registryIamPolicy:RegistryIamPolicy editor projects/{{project}}/locations/{{location}}/registries/{{device_registry}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param RegistryIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryIamPolicyArgs.__new__(RegistryIamPolicyArgs)

            __props__.__dict__["name"] = name
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["etag"] = None
        super(RegistryIamPolicy, __self__).__init__(
            'gcp:iot/registryIamPolicy:RegistryIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'RegistryIamPolicy':
        """
        Get an existing RegistryIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region in which the created registry should reside.
               If it is not provided, the provider region is used.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistryIamPolicyState.__new__(_RegistryIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        return RegistryIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations_get_iam_policy` data source.
        """
        return pulumi.get(self, "policy_data")

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
        The region in which the created registry should reside.
        If it is not provided, the provider region is used.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

