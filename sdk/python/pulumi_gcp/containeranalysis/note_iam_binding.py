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
from . import outputs
from ._inputs import *

__all__ = ['NoteIamBindingArgs', 'NoteIamBinding']

@pulumi.input_type
class NoteIamBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 note: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['NoteIamBindingConditionArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NoteIamBinding resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Identities that will be granted the privilege in `role`.
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
        :param pulumi.Input[str] note: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "note", note)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Identities that will be granted the privilege in `role`.
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
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def note(self) -> pulumi.Input[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: pulumi.Input[str]):
        pulumi.set(self, "note", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['NoteIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['NoteIamBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

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
class _NoteIamBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['NoteIamBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NoteIamBinding resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Identities that will be granted the privilege in `role`.
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
        :param pulumi.Input[str] note: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['NoteIamBindingConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['NoteIamBindingConditionArgs']]):
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
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Identities that will be granted the privilege in `role`.
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
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "note")

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note", value)

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
        `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class NoteIamBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[Union['NoteIamBindingConditionArgs', 'NoteIamBindingConditionArgsDict']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Container Registry Note. Each of these resources serves a different use case:

        * `containeranalysis.NoteIamPolicy`: Authoritative. Sets the IAM policy for the note and replaces any existing policy already attached.
        * `containeranalysis.NoteIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the note are preserved.
        * `containeranalysis.NoteIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the note are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `containeranalysis.NoteIamPolicy`: Retrieves the IAM policy for the note

        > **Note:** `containeranalysis.NoteIamPolicy` **cannot** be used in conjunction with `containeranalysis.NoteIamBinding` and `containeranalysis.NoteIamMember` or they will fight over what your policy should be.

        > **Note:** `containeranalysis.NoteIamBinding` resources **can be** used in conjunction with `containeranalysis.NoteIamMember` resources **only if** they do not grant privilege to the same role.

        ## containeranalysis.NoteIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/containeranalysis.notes.occurrences.viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.containeranalysis.NoteIamPolicy("policy",
            project=note["project"],
            note=note["name"],
            policy_data=admin.policy_data)
        ```

        ## containeranalysis.NoteIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.containeranalysis.NoteIamBinding("binding",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            members=["user:jane@example.com"])
        ```

        ## containeranalysis.NoteIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.containeranalysis.NoteIamMember("member",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            member="user:jane@example.com")
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Container Registry Note
        Three different resources help you manage your IAM policy for Container Registry Note. Each of these resources serves a different use case:

        * `containeranalysis.NoteIamPolicy`: Authoritative. Sets the IAM policy for the note and replaces any existing policy already attached.
        * `containeranalysis.NoteIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the note are preserved.
        * `containeranalysis.NoteIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the note are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `containeranalysis.NoteIamPolicy`: Retrieves the IAM policy for the note

        > **Note:** `containeranalysis.NoteIamPolicy` **cannot** be used in conjunction with `containeranalysis.NoteIamBinding` and `containeranalysis.NoteIamMember` or they will fight over what your policy should be.

        > **Note:** `containeranalysis.NoteIamBinding` resources **can be** used in conjunction with `containeranalysis.NoteIamMember` resources **only if** they do not grant privilege to the same role.

        ## containeranalysis.NoteIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/containeranalysis.notes.occurrences.viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.containeranalysis.NoteIamPolicy("policy",
            project=note["project"],
            note=note["name"],
            policy_data=admin.policy_data)
        ```

        ## containeranalysis.NoteIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.containeranalysis.NoteIamBinding("binding",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            members=["user:jane@example.com"])
        ```

        ## containeranalysis.NoteIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.containeranalysis.NoteIamMember("member",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/notes/{{name}}

        * {{project}}/{{name}}

        * {{name}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Container Registry note IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor projects/{{project}}/notes/{{note}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Identities that will be granted the privilege in `role`.
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
        :param pulumi.Input[str] note: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NoteIamBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Container Registry Note. Each of these resources serves a different use case:

        * `containeranalysis.NoteIamPolicy`: Authoritative. Sets the IAM policy for the note and replaces any existing policy already attached.
        * `containeranalysis.NoteIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the note are preserved.
        * `containeranalysis.NoteIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the note are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `containeranalysis.NoteIamPolicy`: Retrieves the IAM policy for the note

        > **Note:** `containeranalysis.NoteIamPolicy` **cannot** be used in conjunction with `containeranalysis.NoteIamBinding` and `containeranalysis.NoteIamMember` or they will fight over what your policy should be.

        > **Note:** `containeranalysis.NoteIamBinding` resources **can be** used in conjunction with `containeranalysis.NoteIamMember` resources **only if** they do not grant privilege to the same role.

        ## containeranalysis.NoteIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/containeranalysis.notes.occurrences.viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.containeranalysis.NoteIamPolicy("policy",
            project=note["project"],
            note=note["name"],
            policy_data=admin.policy_data)
        ```

        ## containeranalysis.NoteIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.containeranalysis.NoteIamBinding("binding",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            members=["user:jane@example.com"])
        ```

        ## containeranalysis.NoteIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.containeranalysis.NoteIamMember("member",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            member="user:jane@example.com")
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Container Registry Note
        Three different resources help you manage your IAM policy for Container Registry Note. Each of these resources serves a different use case:

        * `containeranalysis.NoteIamPolicy`: Authoritative. Sets the IAM policy for the note and replaces any existing policy already attached.
        * `containeranalysis.NoteIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the note are preserved.
        * `containeranalysis.NoteIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the note are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `containeranalysis.NoteIamPolicy`: Retrieves the IAM policy for the note

        > **Note:** `containeranalysis.NoteIamPolicy` **cannot** be used in conjunction with `containeranalysis.NoteIamBinding` and `containeranalysis.NoteIamMember` or they will fight over what your policy should be.

        > **Note:** `containeranalysis.NoteIamBinding` resources **can be** used in conjunction with `containeranalysis.NoteIamMember` resources **only if** they do not grant privilege to the same role.

        ## containeranalysis.NoteIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/containeranalysis.notes.occurrences.viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.containeranalysis.NoteIamPolicy("policy",
            project=note["project"],
            note=note["name"],
            policy_data=admin.policy_data)
        ```

        ## containeranalysis.NoteIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.containeranalysis.NoteIamBinding("binding",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            members=["user:jane@example.com"])
        ```

        ## containeranalysis.NoteIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.containeranalysis.NoteIamMember("member",
            project=note["project"],
            note=note["name"],
            role="roles/containeranalysis.notes.occurrences.viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/notes/{{name}}

        * {{project}}/{{name}}

        * {{name}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Container Registry note IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor "projects/{{project}}/notes/{{note}} roles/containeranalysis.notes.occurrences.viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:containeranalysis/noteIamBinding:NoteIamBinding editor projects/{{project}}/notes/{{note}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param NoteIamBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NoteIamBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[Union['NoteIamBindingConditionArgs', 'NoteIamBindingConditionArgsDict']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NoteIamBindingArgs.__new__(NoteIamBindingArgs)

            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if note is None and not opts.urn:
                raise TypeError("Missing required property 'note'")
            __props__.__dict__["note"] = note
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(NoteIamBinding, __self__).__init__(
            'gcp:containeranalysis/noteIamBinding:NoteIamBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[Union['NoteIamBindingConditionArgs', 'NoteIamBindingConditionArgsDict']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            note: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'NoteIamBinding':
        """
        Get an existing NoteIamBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: Identities that will be granted the privilege in `role`.
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
        :param pulumi.Input[str] note: Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NoteIamBindingState.__new__(_NoteIamBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["note"] = note
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        return NoteIamBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.NoteIamBindingCondition']]:
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
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        Identities that will be granted the privilege in `role`.
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
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def note(self) -> pulumi.Output[str]:
        """
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "note")

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
        `containeranalysis.NoteIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

