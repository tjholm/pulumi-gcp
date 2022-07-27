# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OrganizationSecurityPolicyAssociationArgs', 'OrganizationSecurityPolicyAssociation']

@pulumi.input_type
class OrganizationSecurityPolicyAssociationArgs:
    def __init__(__self__, *,
                 attachment_id: pulumi.Input[str],
                 policy_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrganizationSecurityPolicyAssociation resource.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        :param pulumi.Input[str] name: The name for an association.
        """
        pulumi.set(__self__, "attachment_id", attachment_id)
        pulumi.set(__self__, "policy_id", policy_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> pulumi.Input[str]:
        """
        The resource that the security policy is attached to.
        """
        return pulumi.get(self, "attachment_id")

    @attachment_id.setter
    def attachment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "attachment_id", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Input[str]:
        """
        The security policy ID of the association.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OrganizationSecurityPolicyAssociationState:
    def __init__(__self__, *,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrganizationSecurityPolicyAssociation resources.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] display_name: The display name of the security policy of the association.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        """
        if attachment_id is not None:
            pulumi.set(__self__, "attachment_id", attachment_id)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource that the security policy is attached to.
        """
        return pulumi.get(self, "attachment_id")

    @attachment_id.setter
    def attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attachment_id", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the security policy of the association.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The security policy ID of the association.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)


class OrganizationSecurityPolicyAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An association for the OrganizationSecurityPolicy.

        To get more information about OrganizationSecurityPolicyAssociation, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addAssociation)
        * How-to Guides
            * [Associating a policy with the organization or folder](https://cloud.google.com/vpc/docs/using-firewall-policies#associate)

        ## Example Usage
        ### Organization Security Policy Association Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        security_policy_target = gcp.organizations.Folder("securityPolicyTarget",
            display_name="tf-test-secpol",
            parent="organizations/123456789",
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy = gcp.compute.OrganizationSecurityPolicy("policyOrganizationSecurityPolicy",
            display_name="tf-test",
            parent=security_policy_target.name,
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy_rule = gcp.compute.OrganizationSecurityPolicyRule("policyOrganizationSecurityPolicyRule",
            policy_id=policy_organization_security_policy.id,
            action="allow",
            direction="INGRESS",
            enable_logging=True,
            match=gcp.compute.OrganizationSecurityPolicyRuleMatchArgs(
                config=gcp.compute.OrganizationSecurityPolicyRuleMatchConfigArgs(
                    src_ip_ranges=[
                        "192.168.0.0/16",
                        "10.0.0.0/8",
                    ],
                    layer4_configs=[
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="tcp",
                            ports=["22"],
                        ),
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="icmp",
                        ),
                    ],
                ),
            ),
            priority=100,
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy_association = gcp.compute.OrganizationSecurityPolicyAssociation("policyOrganizationSecurityPolicyAssociation",
            attachment_id=policy_organization_security_policy.parent,
            policy_id=policy_organization_security_policy.id,
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        OrganizationSecurityPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/association/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationSecurityPolicyAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An association for the OrganizationSecurityPolicy.

        To get more information about OrganizationSecurityPolicyAssociation, see:

        * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addAssociation)
        * How-to Guides
            * [Associating a policy with the organization or folder](https://cloud.google.com/vpc/docs/using-firewall-policies#associate)

        ## Example Usage
        ### Organization Security Policy Association Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        security_policy_target = gcp.organizations.Folder("securityPolicyTarget",
            display_name="tf-test-secpol",
            parent="organizations/123456789",
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy = gcp.compute.OrganizationSecurityPolicy("policyOrganizationSecurityPolicy",
            display_name="tf-test",
            parent=security_policy_target.name,
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy_rule = gcp.compute.OrganizationSecurityPolicyRule("policyOrganizationSecurityPolicyRule",
            policy_id=policy_organization_security_policy.id,
            action="allow",
            direction="INGRESS",
            enable_logging=True,
            match=gcp.compute.OrganizationSecurityPolicyRuleMatchArgs(
                config=gcp.compute.OrganizationSecurityPolicyRuleMatchConfigArgs(
                    src_ip_ranges=[
                        "192.168.0.0/16",
                        "10.0.0.0/8",
                    ],
                    layer4_configs=[
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="tcp",
                            ports=["22"],
                        ),
                        gcp.compute.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs(
                            ip_protocol="icmp",
                        ),
                    ],
                ),
            ),
            priority=100,
            opts=pulumi.ResourceOptions(provider=google_beta))
        policy_organization_security_policy_association = gcp.compute.OrganizationSecurityPolicyAssociation("policyOrganizationSecurityPolicyAssociation",
            attachment_id=policy_organization_security_policy.parent,
            policy_id=policy_organization_security_policy.id,
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        OrganizationSecurityPolicyAssociation can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/association/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param OrganizationSecurityPolicyAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationSecurityPolicyAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attachment_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrganizationSecurityPolicyAssociationArgs.__new__(OrganizationSecurityPolicyAssociationArgs)

            if attachment_id is None and not opts.urn:
                raise TypeError("Missing required property 'attachment_id'")
            __props__.__dict__["attachment_id"] = attachment_id
            __props__.__dict__["name"] = name
            if policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_id'")
            __props__.__dict__["policy_id"] = policy_id
            __props__.__dict__["display_name"] = None
        super(OrganizationSecurityPolicyAssociation, __self__).__init__(
            'gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attachment_id: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_id: Optional[pulumi.Input[str]] = None) -> 'OrganizationSecurityPolicyAssociation':
        """
        Get an existing OrganizationSecurityPolicyAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attachment_id: The resource that the security policy is attached to.
        :param pulumi.Input[str] display_name: The display name of the security policy of the association.
        :param pulumi.Input[str] name: The name for an association.
        :param pulumi.Input[str] policy_id: The security policy ID of the association.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrganizationSecurityPolicyAssociationState.__new__(_OrganizationSecurityPolicyAssociationState)

        __props__.__dict__["attachment_id"] = attachment_id
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["name"] = name
        __props__.__dict__["policy_id"] = policy_id
        return OrganizationSecurityPolicyAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachmentId")
    def attachment_id(self) -> pulumi.Output[str]:
        """
        The resource that the security policy is attached to.
        """
        return pulumi.get(self, "attachment_id")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the security policy of the association.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for an association.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        The security policy ID of the association.
        """
        return pulumi.get(self, "policy_id")

