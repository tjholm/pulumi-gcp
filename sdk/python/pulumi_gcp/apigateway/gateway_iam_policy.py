# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GatewayIamPolicyArgs', 'GatewayIamPolicy']

@pulumi.input_type
class GatewayIamPolicyArgs:
    def __init__(__self__, *,
                 gateway: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GatewayIamPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway for the API.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        pulumi.set(__self__, "gateway", gateway)
        pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Input[str]:
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
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

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the gateway for the API.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _GatewayIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewayIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway for the API.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
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
    def gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
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
        The region of the gateway for the API.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class GatewayIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for API Gateway Gateway. Each of these resources serves a different use case:

        * `apigateway.GatewayIamPolicy`: Authoritative. Sets the IAM policy for the gateway and replaces any existing policy already attached.
        * `apigateway.GatewayIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the gateway are preserved.
        * `apigateway.GatewayIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the gateway are preserved.

        > **Note:** `apigateway.GatewayIamPolicy` **cannot** be used in conjunction with `apigateway.GatewayIamBinding` and `apigateway.GatewayIamMember` or they will fight over what your policy should be.

        > **Note:** `apigateway.GatewayIamBinding` resources **can be** used in conjunction with `apigateway.GatewayIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_api\_gateway\_gateway\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/apigateway.viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.apigateway.GatewayIamPolicy("policy",
            project=google_api_gateway_gateway["api_gw"]["project"],
            region=google_api_gateway_gateway["api_gw"]["region"],
            gateway=google_api_gateway_gateway["api_gw"]["gateway_id"],
            policy_data=admin.policy_data)
        ```

        ## google\_api\_gateway\_gateway\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.apigateway.GatewayIamBinding("binding",
            project=google_api_gateway_gateway["api_gw"]["project"],
            region=google_api_gateway_gateway["api_gw"]["region"],
            gateway=google_api_gateway_gateway["api_gw"]["gateway_id"],
            role="roles/apigateway.viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_api\_gateway\_gateway\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.apigateway.GatewayIamMember("member",
            project=google_api_gateway_gateway["api_gw"]["project"],
            region=google_api_gateway_gateway["api_gw"]["region"],
            gateway=google_api_gateway_gateway["api_gw"]["gateway_id"],
            role="roles/apigateway.viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/gateways/{{gateway}} * {{project}}/{{region}}/{{gateway}} * {{region}}/{{gateway}} * {{gateway}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway gateway IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway for the API.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for API Gateway Gateway. Each of these resources serves a different use case:

        * `apigateway.GatewayIamPolicy`: Authoritative. Sets the IAM policy for the gateway and replaces any existing policy already attached.
        * `apigateway.GatewayIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the gateway are preserved.
        * `apigateway.GatewayIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the gateway are preserved.

        > **Note:** `apigateway.GatewayIamPolicy` **cannot** be used in conjunction with `apigateway.GatewayIamBinding` and `apigateway.GatewayIamMember` or they will fight over what your policy should be.

        > **Note:** `apigateway.GatewayIamBinding` resources **can be** used in conjunction with `apigateway.GatewayIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\_api\_gateway\_gateway\_iam\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/apigateway.viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.apigateway.GatewayIamPolicy("policy",
            project=google_api_gateway_gateway["api_gw"]["project"],
            region=google_api_gateway_gateway["api_gw"]["region"],
            gateway=google_api_gateway_gateway["api_gw"]["gateway_id"],
            policy_data=admin.policy_data)
        ```

        ## google\_api\_gateway\_gateway\_iam\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.apigateway.GatewayIamBinding("binding",
            project=google_api_gateway_gateway["api_gw"]["project"],
            region=google_api_gateway_gateway["api_gw"]["region"],
            gateway=google_api_gateway_gateway["api_gw"]["gateway_id"],
            role="roles/apigateway.viewer",
            members=["user:jane@example.com"])
        ```

        ## google\_api\_gateway\_gateway\_iam\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.apigateway.GatewayIamMember("member",
            project=google_api_gateway_gateway["api_gw"]["project"],
            region=google_api_gateway_gateway["api_gw"]["region"],
            gateway=google_api_gateway_gateway["api_gw"]["gateway_id"],
            role="roles/apigateway.viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{region}}/gateways/{{gateway}} * {{project}}/{{region}}/{{gateway}} * {{region}}/{{gateway}} * {{gateway}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway gateway IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor "projects/{{project}}/locations/{{region}}/gateways/{{gateway}} roles/apigateway.viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy editor projects/{{project}}/locations/{{region}}/gateways/{{gateway}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param GatewayIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayIamPolicyArgs.__new__(GatewayIamPolicyArgs)

            if gateway is None and not opts.urn:
                raise TypeError("Missing required property 'gateway'")
            __props__.__dict__["gateway"] = gateway
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["etag"] = None
        super(GatewayIamPolicy, __self__).__init__(
            'gcp:apigateway/gatewayIamPolicy:GatewayIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'GatewayIamPolicy':
        """
        Get an existing GatewayIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        :param pulumi.Input[str] region: The region of the gateway for the API.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
               region is specified, it is taken from the provider configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayIamPolicyState.__new__(_GatewayIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        return GatewayIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
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
        The region of the gateway for the API.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        region is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "region")

