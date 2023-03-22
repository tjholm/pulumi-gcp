# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConsumersIamPolicyArgs', 'ConsumersIamPolicy']

@pulumi.input_type
class ConsumersIamPolicyArgs:
    def __init__(__self__, *,
                 consumer_project: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a ConsumersIamPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a _organizations_get_iam_policy_ data source.
        """
        pulumi.set(__self__, "consumer_project", consumer_project)
        pulumi.set(__self__, "policy_data", policy_data)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="consumerProject")
    def consumer_project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "consumer_project")

    @consumer_project.setter
    def consumer_project(self, value: pulumi.Input[str]):
        pulumi.set(self, "consumer_project", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a _organizations_get_iam_policy_ data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _ConsumersIamPolicyState:
    def __init__(__self__, *,
                 consumer_project: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConsumersIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a _organizations_get_iam_policy_ data source.
        """
        if consumer_project is not None:
            pulumi.set(__self__, "consumer_project", consumer_project)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="consumerProject")
    def consumer_project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "consumer_project")

    @consumer_project.setter
    def consumer_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_project", value)

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
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a _organizations_get_iam_policy_ data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class ConsumersIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_project: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Cloud Endpoints ServiceConsumers. Each of these resources serves a different use case:

        * `endpoints.ConsumersIamPolicy`: Authoritative. Sets the IAM policy for the serviceconsumers and replaces any existing policy already attached.
        * `endpoints.ConsumersIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the serviceconsumers are preserved.
        * `endpoints.ConsumersIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the serviceconsumers are preserved.

        > **Note:** `endpoints.ConsumersIamPolicy` **cannot** be used in conjunction with `endpoints.ConsumersIamBinding` and `endpoints.ConsumersIamMember` or they will fight over what your policy should be.

        > **Note:** `endpoints.ConsumersIamBinding` resources **can be** used in conjunction with `endpoints.ConsumersIamMember` resources **only if** they do not grant privilege to the same role.

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* services/{{service_name}}/consumers/{{consumer_project}} * {{service_name}}/{{consumer_project}} * {{consumer_project}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Endpoints serviceconsumers IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor "services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor "services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor services/{{service_name}}/consumers/{{consumer_project}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a _organizations_get_iam_policy_ data source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConsumersIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Cloud Endpoints ServiceConsumers. Each of these resources serves a different use case:

        * `endpoints.ConsumersIamPolicy`: Authoritative. Sets the IAM policy for the serviceconsumers and replaces any existing policy already attached.
        * `endpoints.ConsumersIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the serviceconsumers are preserved.
        * `endpoints.ConsumersIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the serviceconsumers are preserved.

        > **Note:** `endpoints.ConsumersIamPolicy` **cannot** be used in conjunction with `endpoints.ConsumersIamBinding` and `endpoints.ConsumersIamMember` or they will fight over what your policy should be.

        > **Note:** `endpoints.ConsumersIamBinding` resources **can be** used in conjunction with `endpoints.ConsumersIamMember` resources **only if** they do not grant privilege to the same role.

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* services/{{service_name}}/consumers/{{consumer_project}} * {{service_name}}/{{consumer_project}} * {{consumer_project}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Endpoints serviceconsumers IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor "services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor "services/{{service_name}}/consumers/{{consumer_project}} roles/servicemanagement.serviceController"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy editor services/{{service_name}}/consumers/{{consumer_project}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param ConsumersIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConsumersIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_project: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConsumersIamPolicyArgs.__new__(ConsumersIamPolicyArgs)

            if consumer_project is None and not opts.urn:
                raise TypeError("Missing required property 'consumer_project'")
            __props__.__dict__["consumer_project"] = consumer_project
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["etag"] = None
        super(ConsumersIamPolicy, __self__).__init__(
            'gcp:endpoints/consumersIamPolicy:ConsumersIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consumer_project: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'ConsumersIamPolicy':
        """
        Get an existing ConsumersIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a _organizations_get_iam_policy_ data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConsumersIamPolicyState.__new__(_ConsumersIamPolicyState)

        __props__.__dict__["consumer_project"] = consumer_project
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["service_name"] = service_name
        return ConsumersIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consumerProject")
    def consumer_project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "consumer_project")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a _organizations_get_iam_policy_ data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_name")

