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

__all__ = ['IAMBindingArgs', 'IAMBinding']

@pulumi.input_type
class IAMBindingArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 project: pulumi.Input[str],
                 role: pulumi.Input[str],
                 condition: Optional[pulumi.Input['IAMBindingConditionArgs']] = None):
        """
        The set of arguments for constructing a IAMBinding resource.
        :param pulumi.Input[str] project: The project id of the target project. This is not
               inferred from the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input['IAMBindingConditionArgs'] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "role", role)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The project id of the target project. This is not
        inferred from the provider.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The role that should be applied. Only one
        `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['IAMBindingConditionArgs']]:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['IAMBindingConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _IAMBindingState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['IAMBindingConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IAMBinding resources.
        :param pulumi.Input['IAMBindingConditionArgs'] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] project: The project id of the target project. This is not
               inferred from the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['IAMBindingConditionArgs']]:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['IAMBindingConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the project's IAM policy.
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
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project id of the target project. This is not
        inferred from the provider.
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
        `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class IAMBinding(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `project_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project "your-project-id roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `project_id` and role, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project "your-project-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `project_id`.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project your-project-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project "your-project-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding to include the title of condition, e.g. `google_project_iam_binding.my_project "{{your-project-id}} roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] project: The project id of the target project. This is not
               inferred from the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IAMBindingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `project_id`, role, and member e.g.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project "your-project-id roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; the resource in question and the role.

        This binding resource can be imported using the `project_id` and role, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project "your-project-id roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `project_id`.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project your-project-id
        ```

         IAM audit config imports use the identifier of the resource in question and the service, e.g.

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding my_project "your-project-id foo.googleapis.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure

        ```sh
         $ pulumi import gcp:projects/iAMBinding:IAMBinding to include the title of condition, e.g. `google_project_iam_binding.my_project "{{your-project-id}} roles/{{role_id}} condition-title"`
        ```

        :param str resource_name: The name of the resource.
        :param IAMBindingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IAMBindingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IAMBindingArgs.__new__(IAMBindingArgs)

            __props__.__dict__["condition"] = condition
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["etag"] = None
        super(IAMBinding, __self__).__init__(
            'gcp:projects/iAMBinding:IAMBinding',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'IAMBinding':
        """
        Get an existing IAMBinding resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['IAMBindingConditionArgs']] condition: An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
               Structure is documented below.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] project: The project id of the target project. This is not
               inferred from the provider.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IAMBindingState.__new__(_IAMBindingState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["members"] = members
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        return IAMBinding(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.IAMBindingCondition']]:
        """
        An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the project's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project id of the target project. This is not
        inferred from the provider.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The role that should be applied. Only one
        `projects.IAMBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

