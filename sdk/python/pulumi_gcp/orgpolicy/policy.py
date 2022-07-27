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

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 parent: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input['PolicySpecArgs']] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input[str] parent: The parent of the resource.
        :param pulumi.Input[str] name: Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        :param pulumi.Input['PolicySpecArgs'] spec: Basic information about the Organization Policy.
        """
        pulumi.set(__self__, "parent", parent)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Input[str]:
        """
        The parent of the resource.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['PolicySpecArgs']]:
        """
        Basic information about the Organization Policy.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['PolicySpecArgs']]):
        pulumi.set(self, "spec", value)


@pulumi.input_type
class _PolicyState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input['PolicySpecArgs']] = None):
        """
        Input properties used for looking up and filtering Policy resources.
        :param pulumi.Input[str] name: Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        :param pulumi.Input[str] parent: The parent of the resource.
        :param pulumi.Input['PolicySpecArgs'] spec: Basic information about the Organization Policy.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The parent of the resource.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['PolicySpecArgs']]:
        """
        Basic information about the Organization Policy.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['PolicySpecArgs']]):
        pulumi.set(self, "spec", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['PolicySpecArgs']]] = None,
                 __props__=None):
        """
        An organization policy gives you programmatic control over your organization's cloud resources.  Using Organization Policies, you will be able to configure constraints across your entire resource hierarchy.

        For more information, see:
        * [Understanding Org Policy concepts](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
        * [The resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy)
        * [All valid constraints](https://cloud.google.com/resource-manager/docs/organization-policy/org-policy-constraints)
        ## Example Usage
        ### Enforce_policy
        A test of an enforce orgpolicy policy for a project
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.organizations.Project("basic",
            org_id="123456789",
            project_id="id")
        primary = gcp.orgpolicy.Policy("primary",
            parent=basic.name.apply(lambda name: f"projects/{name}"),
            spec=gcp.orgpolicy.PolicySpecArgs(
                rules=[gcp.orgpolicy.PolicySpecRuleArgs(
                    enforce="FALSE",
                )],
            ))
        ```
        ### Folder_policy
        A test of an orgpolicy policy for a folder
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.organizations.Folder("basic",
            parent="organizations/123456789",
            display_name="folder")
        primary = gcp.orgpolicy.Policy("primary",
            parent=basic.name,
            spec=gcp.orgpolicy.PolicySpecArgs(
                inherit_from_parent=True,
                rules=[gcp.orgpolicy.PolicySpecRuleArgs(
                    deny_all="TRUE",
                )],
            ))
        ```
        ### Organization_policy
        A test of an orgpolicy policy for an organization
        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.orgpolicy.Policy("primary",
            parent="organizations/123456789",
            spec=gcp.orgpolicy.PolicySpecArgs(
                reset=True,
            ))
        ```
        ### Project_policy
        A test of an orgpolicy policy for a project
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.organizations.Project("basic",
            org_id="123456789",
            project_id="id")
        primary = gcp.orgpolicy.Policy("primary",
            parent=basic.name.apply(lambda name: f"projects/{name}"),
            spec=gcp.orgpolicy.PolicySpecArgs(
                rules=[
                    gcp.orgpolicy.PolicySpecRuleArgs(
                        condition=gcp.orgpolicy.PolicySpecRuleConditionArgs(
                            description="A sample condition for the policy",
                            expression="resource.matchLabels('labelKeys/123', 'labelValues/345')",
                            location="sample-location.log",
                            title="sample-condition",
                        ),
                        values=gcp.orgpolicy.PolicySpecRuleValuesArgs(
                            allowed_values=["projects/allowed-project"],
                            denied_values=["projects/denied-project"],
                        ),
                    ),
                    gcp.orgpolicy.PolicySpecRuleArgs(
                        allow_all="TRUE",
                    ),
                ],
            ))
        ```

        ## Import

        Policy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:orgpolicy/policy:Policy default {{parent}}/policies/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        :param pulumi.Input[str] parent: The parent of the resource.
        :param pulumi.Input[pulumi.InputType['PolicySpecArgs']] spec: Basic information about the Organization Policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An organization policy gives you programmatic control over your organization's cloud resources.  Using Organization Policies, you will be able to configure constraints across your entire resource hierarchy.

        For more information, see:
        * [Understanding Org Policy concepts](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
        * [The resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy)
        * [All valid constraints](https://cloud.google.com/resource-manager/docs/organization-policy/org-policy-constraints)
        ## Example Usage
        ### Enforce_policy
        A test of an enforce orgpolicy policy for a project
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.organizations.Project("basic",
            org_id="123456789",
            project_id="id")
        primary = gcp.orgpolicy.Policy("primary",
            parent=basic.name.apply(lambda name: f"projects/{name}"),
            spec=gcp.orgpolicy.PolicySpecArgs(
                rules=[gcp.orgpolicy.PolicySpecRuleArgs(
                    enforce="FALSE",
                )],
            ))
        ```
        ### Folder_policy
        A test of an orgpolicy policy for a folder
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.organizations.Folder("basic",
            parent="organizations/123456789",
            display_name="folder")
        primary = gcp.orgpolicy.Policy("primary",
            parent=basic.name,
            spec=gcp.orgpolicy.PolicySpecArgs(
                inherit_from_parent=True,
                rules=[gcp.orgpolicy.PolicySpecRuleArgs(
                    deny_all="TRUE",
                )],
            ))
        ```
        ### Organization_policy
        A test of an orgpolicy policy for an organization
        ```python
        import pulumi
        import pulumi_gcp as gcp

        primary = gcp.orgpolicy.Policy("primary",
            parent="organizations/123456789",
            spec=gcp.orgpolicy.PolicySpecArgs(
                reset=True,
            ))
        ```
        ### Project_policy
        A test of an orgpolicy policy for a project
        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic = gcp.organizations.Project("basic",
            org_id="123456789",
            project_id="id")
        primary = gcp.orgpolicy.Policy("primary",
            parent=basic.name.apply(lambda name: f"projects/{name}"),
            spec=gcp.orgpolicy.PolicySpecArgs(
                rules=[
                    gcp.orgpolicy.PolicySpecRuleArgs(
                        condition=gcp.orgpolicy.PolicySpecRuleConditionArgs(
                            description="A sample condition for the policy",
                            expression="resource.matchLabels('labelKeys/123', 'labelValues/345')",
                            location="sample-location.log",
                            title="sample-condition",
                        ),
                        values=gcp.orgpolicy.PolicySpecRuleValuesArgs(
                            allowed_values=["projects/allowed-project"],
                            denied_values=["projects/denied-project"],
                        ),
                    ),
                    gcp.orgpolicy.PolicySpecRuleArgs(
                        allow_all="TRUE",
                    ),
                ],
            ))
        ```

        ## Import

        Policy can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:orgpolicy/policy:Policy default {{parent}}/policies/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['PolicySpecArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyArgs.__new__(PolicyArgs)

            __props__.__dict__["name"] = name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__.__dict__["parent"] = parent
            __props__.__dict__["spec"] = spec
        super(Policy, __self__).__init__(
            'gcp:orgpolicy/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None,
            spec: Optional[pulumi.Input[pulumi.InputType['PolicySpecArgs']]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        :param pulumi.Input[str] parent: The parent of the resource.
        :param pulumi.Input[pulumi.InputType['PolicySpecArgs']] spec: Basic information about the Organization Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyState.__new__(_PolicyState)

        __props__.__dict__["name"] = name
        __props__.__dict__["parent"] = parent
        __props__.__dict__["spec"] = spec
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[str]:
        """
        The parent of the resource.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[Optional['outputs.PolicySpec']]:
        """
        Basic information about the Organization Policy.
        """
        return pulumi.get(self, "spec")

