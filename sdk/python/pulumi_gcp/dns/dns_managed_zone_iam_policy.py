# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['DnsManagedZoneIamPolicyArgs', 'DnsManagedZoneIamPolicy']

@pulumi.input_type
class DnsManagedZoneIamPolicyArgs:
    def __init__(__self__, *,
                 managed_zone: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DnsManagedZoneIamPolicy resource.
        :param pulumi.Input[str] managed_zone: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "managed_zone", managed_zone)
        pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="managedZone")
    def managed_zone(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "managed_zone")

    @managed_zone.setter
    def managed_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "managed_zone", value)

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
class _DnsManagedZoneIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 managed_zone: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DnsManagedZoneIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] managed_zone: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if managed_zone is not None:
            pulumi.set(__self__, "managed_zone", managed_zone)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)

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
    @pulumi.getter(name="managedZone")
    def managed_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "managed_zone")

    @managed_zone.setter
    def managed_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_zone", value)

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


class DnsManagedZoneIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 managed_zone: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:

        * `dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
        * `dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
        * `dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dns.DnsManagedZoneIamPolicy`: Retrieves the IAM policy for the managedzone

        > **Note:** `dns.DnsManagedZoneIamPolicy` **cannot** be used in conjunction with `dns.DnsManagedZoneIamBinding` and `dns.DnsManagedZoneIamMember` or they will fight over what your policy should be.

        > **Note:** `dns.DnsManagedZoneIamBinding` resources **can be** used in conjunction with `dns.DnsManagedZoneIamMember` resources **only if** they do not grant privilege to the same role.

        ## dns.DnsManagedZoneIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dns.DnsManagedZoneIamPolicy("policy",
            project=default["project"],
            managed_zone=default["name"],
            policy_data=admin.policy_data)
        ```

        ## dns.DnsManagedZoneIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dns.DnsManagedZoneIamBinding("binding",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dns.DnsManagedZoneIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dns.DnsManagedZoneIamMember("member",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Cloud DNS ManagedZone
        Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:

        * `dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
        * `dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
        * `dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dns.DnsManagedZoneIamPolicy`: Retrieves the IAM policy for the managedzone

        > **Note:** `dns.DnsManagedZoneIamPolicy` **cannot** be used in conjunction with `dns.DnsManagedZoneIamBinding` and `dns.DnsManagedZoneIamMember` or they will fight over what your policy should be.

        > **Note:** `dns.DnsManagedZoneIamBinding` resources **can be** used in conjunction with `dns.DnsManagedZoneIamMember` resources **only if** they do not grant privilege to the same role.

        ## dns.DnsManagedZoneIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dns.DnsManagedZoneIamPolicy("policy",
            project=default["project"],
            managed_zone=default["name"],
            policy_data=admin.policy_data)
        ```

        ## dns.DnsManagedZoneIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dns.DnsManagedZoneIamBinding("binding",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dns.DnsManagedZoneIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dns.DnsManagedZoneIamMember("member",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/managedZones/{{managed_zone}}

        * {{project}}/{{managed_zone}}

        * {{managed_zone}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Cloud DNS managedzone IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor projects/{{project}}/managedZones/{{managed_zone}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] managed_zone: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DnsManagedZoneIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:

        * `dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
        * `dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
        * `dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dns.DnsManagedZoneIamPolicy`: Retrieves the IAM policy for the managedzone

        > **Note:** `dns.DnsManagedZoneIamPolicy` **cannot** be used in conjunction with `dns.DnsManagedZoneIamBinding` and `dns.DnsManagedZoneIamMember` or they will fight over what your policy should be.

        > **Note:** `dns.DnsManagedZoneIamBinding` resources **can be** used in conjunction with `dns.DnsManagedZoneIamMember` resources **only if** they do not grant privilege to the same role.

        ## dns.DnsManagedZoneIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dns.DnsManagedZoneIamPolicy("policy",
            project=default["project"],
            managed_zone=default["name"],
            policy_data=admin.policy_data)
        ```

        ## dns.DnsManagedZoneIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dns.DnsManagedZoneIamBinding("binding",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dns.DnsManagedZoneIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dns.DnsManagedZoneIamMember("member",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Cloud DNS ManagedZone
        Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:

        * `dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
        * `dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
        * `dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dns.DnsManagedZoneIamPolicy`: Retrieves the IAM policy for the managedzone

        > **Note:** `dns.DnsManagedZoneIamPolicy` **cannot** be used in conjunction with `dns.DnsManagedZoneIamBinding` and `dns.DnsManagedZoneIamMember` or they will fight over what your policy should be.

        > **Note:** `dns.DnsManagedZoneIamBinding` resources **can be** used in conjunction with `dns.DnsManagedZoneIamMember` resources **only if** they do not grant privilege to the same role.

        ## dns.DnsManagedZoneIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dns.DnsManagedZoneIamPolicy("policy",
            project=default["project"],
            managed_zone=default["name"],
            policy_data=admin.policy_data)
        ```

        ## dns.DnsManagedZoneIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dns.DnsManagedZoneIamBinding("binding",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dns.DnsManagedZoneIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dns.DnsManagedZoneIamMember("member",
            project=default["project"],
            managed_zone=default["name"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/managedZones/{{managed_zone}}

        * {{project}}/{{managed_zone}}

        * {{managed_zone}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Cloud DNS managedzone IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy editor projects/{{project}}/managedZones/{{managed_zone}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param DnsManagedZoneIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DnsManagedZoneIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 managed_zone: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DnsManagedZoneIamPolicyArgs.__new__(DnsManagedZoneIamPolicyArgs)

            if managed_zone is None and not opts.urn:
                raise TypeError("Missing required property 'managed_zone'")
            __props__.__dict__["managed_zone"] = managed_zone
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["etag"] = None
        super(DnsManagedZoneIamPolicy, __self__).__init__(
            'gcp:dns/dnsManagedZoneIamPolicy:DnsManagedZoneIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            managed_zone: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'DnsManagedZoneIamPolicy':
        """
        Get an existing DnsManagedZoneIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] managed_zone: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DnsManagedZoneIamPolicyState.__new__(_DnsManagedZoneIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["managed_zone"] = managed_zone
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        return DnsManagedZoneIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="managedZone")
    def managed_zone(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "managed_zone")

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

