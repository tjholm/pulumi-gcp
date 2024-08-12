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

__all__ = ['MachineImageIamPolicyArgs', 'MachineImageIamPolicy']

@pulumi.input_type
class MachineImageIamPolicyArgs:
    def __init__(__self__, *,
                 machine_image: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MachineImageIamPolicy resource.
        :param pulumi.Input[str] machine_image: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "machine_image", machine_image)
        pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="machineImage")
    def machine_image(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "machine_image")

    @machine_image.setter
    def machine_image(self, value: pulumi.Input[str]):
        pulumi.set(self, "machine_image", value)

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
class _MachineImageIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 machine_image: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MachineImageIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] machine_image: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if machine_image is not None:
            pulumi.set(__self__, "machine_image", machine_image)
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
    @pulumi.getter(name="machineImage")
    def machine_image(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "machine_image")

    @machine_image.setter
    def machine_image(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "machine_image", value)

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


class MachineImageIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 machine_image: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:

        * `compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
        * `compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
        * `compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `compute.MachineImageIamPolicy`: Retrieves the IAM policy for the machineimage

        > **Note:** `compute.MachineImageIamPolicy` **cannot** be used in conjunction with `compute.MachineImageIamBinding` and `compute.MachineImageIamMember` or they will fight over what your policy should be.

        > **Note:** `compute.MachineImageIamBinding` resources **can be** used in conjunction with `compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
        ## google\\_compute\\_machine\\_image\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```
        ## compute.MachineImageIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## compute.MachineImageIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Compute Engine MachineImage
        Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:

        * `compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
        * `compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
        * `compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `compute.MachineImageIamPolicy`: Retrieves the IAM policy for the machineimage

        > **Note:** `compute.MachineImageIamPolicy` **cannot** be used in conjunction with `compute.MachineImageIamBinding` and `compute.MachineImageIamMember` or they will fight over what your policy should be.

        > **Note:** `compute.MachineImageIamBinding` resources **can be** used in conjunction with `compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
        ## google\\_compute\\_machine\\_image\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```
        ## compute.MachineImageIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## compute.MachineImageIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/global/machineImages/{{name}}

        * {{project}}/{{name}}

        * {{name}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Compute Engine machineimage IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor projects/{{project}}/global/machineImages/{{machine_image}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] machine_image: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MachineImageIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:

        * `compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
        * `compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
        * `compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `compute.MachineImageIamPolicy`: Retrieves the IAM policy for the machineimage

        > **Note:** `compute.MachineImageIamPolicy` **cannot** be used in conjunction with `compute.MachineImageIamBinding` and `compute.MachineImageIamMember` or they will fight over what your policy should be.

        > **Note:** `compute.MachineImageIamBinding` resources **can be** used in conjunction with `compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
        ## google\\_compute\\_machine\\_image\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```
        ## compute.MachineImageIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## compute.MachineImageIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Compute Engine MachineImage
        Three different resources help you manage your IAM policy for Compute Engine MachineImage. Each of these resources serves a different use case:

        * `compute.MachineImageIamPolicy`: Authoritative. Sets the IAM policy for the machineimage and replaces any existing policy already attached.
        * `compute.MachineImageIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the machineimage are preserved.
        * `compute.MachineImageIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the machineimage are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `compute.MachineImageIamPolicy`: Retrieves the IAM policy for the machineimage

        > **Note:** `compute.MachineImageIamPolicy` **cannot** be used in conjunction with `compute.MachineImageIamBinding` and `compute.MachineImageIamMember` or they will fight over what your policy should be.

        > **Note:** `compute.MachineImageIamBinding` resources **can be** used in conjunction with `compute.MachineImageIamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
        ## google\\_compute\\_machine\\_image\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/compute.admin",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.compute.MachineImageIamPolicy("policy",
            project=image["project"],
            machine_image=image["name"],
            policy_data=admin.policy_data)
        ```
        ## compute.MachineImageIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.compute.MachineImageIamBinding("binding",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## compute.MachineImageIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.compute.MachineImageIamMember("member",
            project=image["project"],
            machine_image=image["name"],
            role="roles/compute.admin",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/global/machineImages/{{name}}

        * {{project}}/{{name}}

        * {{name}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Compute Engine machineimage IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor "projects/{{project}}/global/machineImages/{{machine_image}} roles/compute.admin"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:compute/machineImageIamPolicy:MachineImageIamPolicy editor projects/{{project}}/global/machineImages/{{machine_image}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param MachineImageIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MachineImageIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 machine_image: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MachineImageIamPolicyArgs.__new__(MachineImageIamPolicyArgs)

            if machine_image is None and not opts.urn:
                raise TypeError("Missing required property 'machine_image'")
            __props__.__dict__["machine_image"] = machine_image
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["etag"] = None
        super(MachineImageIamPolicy, __self__).__init__(
            'gcp:compute/machineImageIamPolicy:MachineImageIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            machine_image: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'MachineImageIamPolicy':
        """
        Get an existing MachineImageIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] machine_image: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MachineImageIamPolicyState.__new__(_MachineImageIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["machine_image"] = machine_image
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        return MachineImageIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="machineImage")
    def machine_image(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "machine_image")

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

