# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ServiceIamMemberArgs', 'ServiceIamMember']

@pulumi.input_type
class ServiceIamMemberArgs:
    def __init__(__self__, *,
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['ServiceIamMemberConditionArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceIamMember resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] location: The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
               * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
               * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
               * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        """
        ServiceIamMemberArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            member=member,
            role=role,
            condition=condition,
            location=location,
            name=name,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             member: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             condition: Optional[pulumi.Input['ServiceIamMemberConditionArgs']] = None,
             location: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if member is None:
            raise TypeError("Missing 'member' argument")
        if role is None:
            raise TypeError("Missing 'role' argument")

        _setter("member", member)
        _setter("role", role)
        if condition is not None:
            _setter("condition", condition)
        if location is not None:
            _setter("location", location)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)

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
        `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['ServiceIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['ServiceIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

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

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ServiceIamMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['ServiceIamMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceIamMember resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
               * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
               * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
               * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        _ServiceIamMemberState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            condition=condition,
            etag=etag,
            location=location,
            member=member,
            name=name,
            project=project,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             condition: Optional[pulumi.Input['ServiceIamMemberConditionArgs']] = None,
             etag: Optional[pulumi.Input[str]] = None,
             location: Optional[pulumi.Input[str]] = None,
             member: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if condition is not None:
            _setter("condition", condition)
        if etag is not None:
            _setter("etag", etag)
        if location is not None:
            _setter("location", location)
        if member is not None:
            _setter("member", member)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['ServiceIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['ServiceIamMemberConditionArgs']]):
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
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member", value)

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

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
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
        `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class ServiceIamMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['ServiceIamMemberConditionArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Cloud Run (v2 API) Service. Each of these resources serves a different use case:

        * `cloudrunv2.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `cloudrunv2.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `cloudrunv2.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `cloudrunv2.ServiceIamPolicy`: Retrieves the IAM policy for the service

        > **Note:** `cloudrunv2.ServiceIamPolicy` **cannot** be used in conjunction with `cloudrunv2.ServiceIamBinding` and `cloudrunv2.ServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `cloudrunv2.ServiceIamBinding` resources **can be** used in conjunction with `cloudrunv2.ServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_cloud\\_run\\_v2\\_service\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.cloudrunv2.ServiceIamPolicy("policy",
            project=google_cloud_run_v2_service["default"]["project"],
            location=google_cloud_run_v2_service["default"]["location"],
            policy_data=admin.policy_data)
        ```

        ## google\\_cloud\\_run\\_v2\\_service\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.cloudrunv2.ServiceIamBinding("binding",
            project=google_cloud_run_v2_service["default"]["project"],
            location=google_cloud_run_v2_service["default"]["location"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_cloud\\_run\\_v2\\_service\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.cloudrunv2.ServiceIamMember("member",
            project=google_cloud_run_v2_service["default"]["project"],
            location=google_cloud_run_v2_service["default"]["location"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/services/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Run (v2 API) service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor projects/{{project}}/locations/{{location}}/services/{{service}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
               * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
               * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
               * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceIamMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Cloud Run (v2 API) Service. Each of these resources serves a different use case:

        * `cloudrunv2.ServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `cloudrunv2.ServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `cloudrunv2.ServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `cloudrunv2.ServiceIamPolicy`: Retrieves the IAM policy for the service

        > **Note:** `cloudrunv2.ServiceIamPolicy` **cannot** be used in conjunction with `cloudrunv2.ServiceIamBinding` and `cloudrunv2.ServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `cloudrunv2.ServiceIamBinding` resources **can be** used in conjunction with `cloudrunv2.ServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_cloud\\_run\\_v2\\_service\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.cloudrunv2.ServiceIamPolicy("policy",
            project=google_cloud_run_v2_service["default"]["project"],
            location=google_cloud_run_v2_service["default"]["location"],
            policy_data=admin.policy_data)
        ```

        ## google\\_cloud\\_run\\_v2\\_service\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.cloudrunv2.ServiceIamBinding("binding",
            project=google_cloud_run_v2_service["default"]["project"],
            location=google_cloud_run_v2_service["default"]["location"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_cloud\\_run\\_v2\\_service\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.cloudrunv2.ServiceIamMember("member",
            project=google_cloud_run_v2_service["default"]["project"],
            location=google_cloud_run_v2_service["default"]["location"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/services/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Run (v2 API) service IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:cloudrunv2/serviceIamMember:ServiceIamMember editor projects/{{project}}/locations/{{location}}/services/{{service}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param ServiceIamMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceIamMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ServiceIamMemberArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['ServiceIamMemberConditionArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceIamMemberArgs.__new__(ServiceIamMemberArgs)

            condition = _utilities.configure(condition, ServiceIamMemberConditionArgs, True)
            __props__.__dict__["condition"] = condition
            __props__.__dict__["location"] = location
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(ServiceIamMember, __self__).__init__(
            'gcp:cloudrunv2/serviceIamMember:ServiceIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['ServiceIamMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'ServiceIamMember':
        """
        Get an existing ServiceIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] name: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
               Each entry can have one of the following values:
               * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
               * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
               * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
               * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
               * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
               * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
               * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
               * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
               * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceIamMemberState.__new__(_ServiceIamMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["location"] = location
        __props__.__dict__["member"] = member
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        return ServiceIamMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.ServiceIamMemberCondition']]:
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
    def location(self) -> pulumi.Output[str]:
        """
        The location of the cloud run service Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        Each entry can have one of the following values:
        * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `cloudrunv2.ServiceIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

