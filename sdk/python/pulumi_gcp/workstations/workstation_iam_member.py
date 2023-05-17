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

__all__ = ['WorkstationIamMemberArgs', 'WorkstationIamMember']

@pulumi.input_type
class WorkstationIamMemberArgs:
    def __init__(__self__, *,
                 member: pulumi.Input[str],
                 role: pulumi.Input[str],
                 workstation_cluster_id: pulumi.Input[str],
                 workstation_config_id: pulumi.Input[str],
                 workstation_id: pulumi.Input[str],
                 condition: Optional[pulumi.Input['WorkstationIamMemberConditionArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WorkstationIamMember resource.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        :param pulumi.Input[str] location: The location where the workstation parent resources reside.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "workstation_cluster_id", workstation_cluster_id)
        pulumi.set(__self__, "workstation_config_id", workstation_config_id)
        pulumi.set(__self__, "workstation_id", workstation_id)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

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
        `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="workstationClusterId")
    def workstation_cluster_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workstation_cluster_id")

    @workstation_cluster_id.setter
    def workstation_cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workstation_cluster_id", value)

    @property
    @pulumi.getter(name="workstationConfigId")
    def workstation_config_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workstation_config_id")

    @workstation_config_id.setter
    def workstation_config_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workstation_config_id", value)

    @property
    @pulumi.getter(name="workstationId")
    def workstation_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workstation_id")

    @workstation_id.setter
    def workstation_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workstation_id", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['WorkstationIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['WorkstationIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the workstation parent resources reside.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

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
class _WorkstationIamMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['WorkstationIamMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 workstation_cluster_id: Optional[pulumi.Input[str]] = None,
                 workstation_config_id: Optional[pulumi.Input[str]] = None,
                 workstation_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WorkstationIamMember resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location where the workstation parent resources reside.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if workstation_cluster_id is not None:
            pulumi.set(__self__, "workstation_cluster_id", workstation_cluster_id)
        if workstation_config_id is not None:
            pulumi.set(__self__, "workstation_config_id", workstation_config_id)
        if workstation_id is not None:
            pulumi.set(__self__, "workstation_id", workstation_id)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['WorkstationIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['WorkstationIamMemberConditionArgs']]):
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
        The location where the workstation parent resources reside.
        Used to find the parent resource to bind the IAM policy to
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
        `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="workstationClusterId")
    def workstation_cluster_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workstation_cluster_id")

    @workstation_cluster_id.setter
    def workstation_cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workstation_cluster_id", value)

    @property
    @pulumi.getter(name="workstationConfigId")
    def workstation_config_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workstation_config_id")

    @workstation_config_id.setter
    def workstation_config_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workstation_config_id", value)

    @property
    @pulumi.getter(name="workstationId")
    def workstation_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "workstation_id")

    @workstation_id.setter
    def workstation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workstation_id", value)


class WorkstationIamMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['WorkstationIamMemberConditionArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 workstation_cluster_id: Optional[pulumi.Input[str]] = None,
                 workstation_config_id: Optional[pulumi.Input[str]] = None,
                 workstation_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} * {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}} * {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}} * {{workstation_id}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Workstations workstation IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:workstations/workstationIamMember:WorkstationIamMember editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:workstations/workstationIamMember:WorkstationIamMember editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:workstations/workstationIamMember:WorkstationIamMember editor projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location where the workstation parent resources reside.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkstationIamMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} * {{project}}/{{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}} * {{location}}/{{workstation_cluster_id}}/{{workstation_config_id}}/{{workstation_id}} * {{workstation_id}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Workstations workstation IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:workstations/workstationIamMember:WorkstationIamMember editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:workstations/workstationIamMember:WorkstationIamMember editor "projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:workstations/workstationIamMember:WorkstationIamMember editor projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}/workstationConfigs/{{workstation_config_id}}/workstations/{{workstation_id}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param WorkstationIamMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkstationIamMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['WorkstationIamMemberConditionArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 workstation_cluster_id: Optional[pulumi.Input[str]] = None,
                 workstation_config_id: Optional[pulumi.Input[str]] = None,
                 workstation_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkstationIamMemberArgs.__new__(WorkstationIamMemberArgs)

            __props__.__dict__["condition"] = condition
            __props__.__dict__["location"] = location
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            __props__.__dict__["project"] = project
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if workstation_cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'workstation_cluster_id'")
            __props__.__dict__["workstation_cluster_id"] = workstation_cluster_id
            if workstation_config_id is None and not opts.urn:
                raise TypeError("Missing required property 'workstation_config_id'")
            __props__.__dict__["workstation_config_id"] = workstation_config_id
            if workstation_id is None and not opts.urn:
                raise TypeError("Missing required property 'workstation_id'")
            __props__.__dict__["workstation_id"] = workstation_id
            __props__.__dict__["etag"] = None
        super(WorkstationIamMember, __self__).__init__(
            'gcp:workstations/workstationIamMember:WorkstationIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['WorkstationIamMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            workstation_cluster_id: Optional[pulumi.Input[str]] = None,
            workstation_config_id: Optional[pulumi.Input[str]] = None,
            workstation_id: Optional[pulumi.Input[str]] = None) -> 'WorkstationIamMember':
        """
        Get an existing WorkstationIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location where the workstation parent resources reside.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] role: The role that should be applied. Only one
               `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
               `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkstationIamMemberState.__new__(_WorkstationIamMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["location"] = location
        __props__.__dict__["member"] = member
        __props__.__dict__["project"] = project
        __props__.__dict__["role"] = role
        __props__.__dict__["workstation_cluster_id"] = workstation_cluster_id
        __props__.__dict__["workstation_config_id"] = workstation_config_id
        __props__.__dict__["workstation_id"] = workstation_id
        return WorkstationIamMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.WorkstationIamMemberCondition']]:
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
        The location where the workstation parent resources reside.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

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
        `workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
        `[projects|organizations]/{parent-name}/roles/{role-name}`.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="workstationClusterId")
    def workstation_cluster_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workstation_cluster_id")

    @property
    @pulumi.getter(name="workstationConfigId")
    def workstation_config_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workstation_config_id")

    @property
    @pulumi.getter(name="workstationId")
    def workstation_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workstation_id")

