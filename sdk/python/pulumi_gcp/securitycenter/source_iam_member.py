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

__all__ = ['SourceIamMemberArgs', 'SourceIamMember']

@pulumi.input_type
class SourceIamMemberArgs:
    def __init__(__self__, *,
                 member: pulumi.Input[str],
                 organization: pulumi.Input[str],
                 role: pulumi.Input[str],
                 source: pulumi.Input[str],
                 condition: Optional[pulumi.Input['SourceIamMemberConditionArgs']] = None):
        """
        The set of arguments for constructing a SourceIamMember resource.
        :param pulumi.Input[str] organization: The organization whose Cloud Security Command Center the Source
               lives in.
               
               
               - - -
        """
        pulumi.set(__self__, "member", member)
        pulumi.set(__self__, "organization", organization)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "source", source)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)

    @property
    @pulumi.getter
    def member(self) -> pulumi.Input[str]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: pulumi.Input[str]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Input[str]:
        """
        The organization whose Cloud Security Command Center the Source
        lives in.


        - - -
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['SourceIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['SourceIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)


@pulumi.input_type
class _SourceIamMemberState:
    def __init__(__self__, *,
                 condition: Optional[pulumi.Input['SourceIamMemberConditionArgs']] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SourceIamMember resources.
        :param pulumi.Input[str] organization: The organization whose Cloud Security Command Center the Source
               lives in.
               
               
               - - -
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if member is not None:
            pulumi.set(__self__, "member", member)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if source is not None:
            pulumi.set(__self__, "source", source)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['SourceIamMemberConditionArgs']]:
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['SourceIamMemberConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def member(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "member")

    @member.setter
    def member(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "member", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        """
        The organization whose Cloud Security Command Center the Source
        lives in.


        - - -
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)


class SourceIamMember(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['SourceIamMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A Cloud Security Command Center's (Cloud SCC) finding source. A finding
        source is an entity or a mechanism that can produce a finding. A source is
        like a container of findings that come from the same scanner, logger,
        monitor, etc.

        To get more information about Source, see:

        * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.sources)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/security-command-center/docs)

        ## Example Usage
        ### Scc Source Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        custom_source = gcp.securitycenter.Source("customSource",
            description="My custom Cloud Security Command Center Finding Source",
            display_name="My Source",
            organization="123456789")
        ```

        ## Import

        Source can be imported using any of these accepted formats* `organizations/{{organization}}/sources/{{name}}` * `{{organization}}/{{name}}` When using the `pulumi import` command, Source can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:securitycenter/sourceIamMember:SourceIamMember default organizations/{{organization}}/sources/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/sourceIamMember:SourceIamMember default {{organization}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] organization: The organization whose Cloud Security Command Center the Source
               lives in.
               
               
               - - -
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SourceIamMemberArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Cloud Security Command Center's (Cloud SCC) finding source. A finding
        source is an entity or a mechanism that can produce a finding. A source is
        like a container of findings that come from the same scanner, logger,
        monitor, etc.

        To get more information about Source, see:

        * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.sources)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/security-command-center/docs)

        ## Example Usage
        ### Scc Source Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        custom_source = gcp.securitycenter.Source("customSource",
            description="My custom Cloud Security Command Center Finding Source",
            display_name="My Source",
            organization="123456789")
        ```

        ## Import

        Source can be imported using any of these accepted formats* `organizations/{{organization}}/sources/{{name}}` * `{{organization}}/{{name}}` When using the `pulumi import` command, Source can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:securitycenter/sourceIamMember:SourceIamMember default organizations/{{organization}}/sources/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/sourceIamMember:SourceIamMember default {{organization}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param SourceIamMemberArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SourceIamMemberArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 condition: Optional[pulumi.Input[pulumi.InputType['SourceIamMemberConditionArgs']]] = None,
                 member: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SourceIamMemberArgs.__new__(SourceIamMemberArgs)

            __props__.__dict__["condition"] = condition
            if member is None and not opts.urn:
                raise TypeError("Missing required property 'member'")
            __props__.__dict__["member"] = member
            if organization is None and not opts.urn:
                raise TypeError("Missing required property 'organization'")
            __props__.__dict__["organization"] = organization
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["etag"] = None
        super(SourceIamMember, __self__).__init__(
            'gcp:securitycenter/sourceIamMember:SourceIamMember',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            condition: Optional[pulumi.Input[pulumi.InputType['SourceIamMemberConditionArgs']]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            member: Optional[pulumi.Input[str]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            source: Optional[pulumi.Input[str]] = None) -> 'SourceIamMember':
        """
        Get an existing SourceIamMember resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] organization: The organization whose Cloud Security Command Center the Source
               lives in.
               
               
               - - -
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SourceIamMemberState.__new__(_SourceIamMemberState)

        __props__.__dict__["condition"] = condition
        __props__.__dict__["etag"] = etag
        __props__.__dict__["member"] = member
        __props__.__dict__["organization"] = organization
        __props__.__dict__["role"] = role
        __props__.__dict__["source"] = source
        return SourceIamMember(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[Optional['outputs.SourceIamMemberCondition']]:
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def member(self) -> pulumi.Output[str]:
        return pulumi.get(self, "member")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[str]:
        """
        The organization whose Cloud Security Command Center the Source
        lives in.


        - - -
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source")

