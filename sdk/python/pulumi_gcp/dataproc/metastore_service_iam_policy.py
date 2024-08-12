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

__all__ = ['MetastoreServiceIamPolicyArgs', 'MetastoreServiceIamPolicy']

@pulumi.input_type
class MetastoreServiceIamPolicyArgs:
    def __init__(__self__, *,
                 policy_data: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MetastoreServiceIamPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] location: The location where the metastore service should reside.
               The default value is `global`.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
               location is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "policy_data", policy_data)
        pulumi.set(__self__, "service_id", service_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

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
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the metastore service should reside.
        The default value is `global`.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        location is specified, it is taken from the provider configuration.
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
class _MetastoreServiceIamPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MetastoreServiceIamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location where the metastore service should reside.
               The default value is `global`.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
               location is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

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
        The location where the metastore service should reside.
        The default value is `global`.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        location is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

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

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)


class MetastoreServiceIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Dataproc metastore Service. Each of these resources serves a different use case:

        * `dataproc.MetastoreServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `dataproc.MetastoreServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `dataproc.MetastoreServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dataproc.MetastoreServiceIamPolicy`: Retrieves the IAM policy for the service

        > **Note:** `dataproc.MetastoreServiceIamPolicy` **cannot** be used in conjunction with `dataproc.MetastoreServiceIamBinding` and `dataproc.MetastoreServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `dataproc.MetastoreServiceIamBinding` resources **can be** used in conjunction with `dataproc.MetastoreServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## dataproc.MetastoreServiceIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dataproc.MetastoreServiceIamPolicy("policy",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            policy_data=admin.policy_data)
        ```

        ## dataproc.MetastoreServiceIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dataproc.MetastoreServiceIamBinding("binding",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dataproc.MetastoreServiceIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dataproc.MetastoreServiceIamMember("member",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Dataproc metastore Service
        Three different resources help you manage your IAM policy for Dataproc metastore Service. Each of these resources serves a different use case:

        * `dataproc.MetastoreServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `dataproc.MetastoreServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `dataproc.MetastoreServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dataproc.MetastoreServiceIamPolicy`: Retrieves the IAM policy for the service

        > **Note:** `dataproc.MetastoreServiceIamPolicy` **cannot** be used in conjunction with `dataproc.MetastoreServiceIamBinding` and `dataproc.MetastoreServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `dataproc.MetastoreServiceIamBinding` resources **can be** used in conjunction with `dataproc.MetastoreServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## dataproc.MetastoreServiceIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dataproc.MetastoreServiceIamPolicy("policy",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            policy_data=admin.policy_data)
        ```

        ## dataproc.MetastoreServiceIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dataproc.MetastoreServiceIamBinding("binding",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dataproc.MetastoreServiceIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dataproc.MetastoreServiceIamMember("member",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/locations/{{location}}/services/{{service_id}}

        * {{project}}/{{location}}/{{service_id}}

        * {{location}}/{{service_id}}

        * {{service_id}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Dataproc metastore service IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy editor projects/{{project}}/locations/{{location}}/services/{{service_id}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location where the metastore service should reside.
               The default value is `global`.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
               location is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MetastoreServiceIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Dataproc metastore Service. Each of these resources serves a different use case:

        * `dataproc.MetastoreServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `dataproc.MetastoreServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `dataproc.MetastoreServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dataproc.MetastoreServiceIamPolicy`: Retrieves the IAM policy for the service

        > **Note:** `dataproc.MetastoreServiceIamPolicy` **cannot** be used in conjunction with `dataproc.MetastoreServiceIamBinding` and `dataproc.MetastoreServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `dataproc.MetastoreServiceIamBinding` resources **can be** used in conjunction with `dataproc.MetastoreServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## dataproc.MetastoreServiceIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dataproc.MetastoreServiceIamPolicy("policy",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            policy_data=admin.policy_data)
        ```

        ## dataproc.MetastoreServiceIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dataproc.MetastoreServiceIamBinding("binding",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dataproc.MetastoreServiceIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dataproc.MetastoreServiceIamMember("member",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for Dataproc metastore Service
        Three different resources help you manage your IAM policy for Dataproc metastore Service. Each of these resources serves a different use case:

        * `dataproc.MetastoreServiceIamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
        * `dataproc.MetastoreServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
        * `dataproc.MetastoreServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `dataproc.MetastoreServiceIamPolicy`: Retrieves the IAM policy for the service

        > **Note:** `dataproc.MetastoreServiceIamPolicy` **cannot** be used in conjunction with `dataproc.MetastoreServiceIamBinding` and `dataproc.MetastoreServiceIamMember` or they will fight over what your policy should be.

        > **Note:** `dataproc.MetastoreServiceIamBinding` resources **can be** used in conjunction with `dataproc.MetastoreServiceIamMember` resources **only if** they do not grant privilege to the same role.

        ## dataproc.MetastoreServiceIamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/viewer",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.dataproc.MetastoreServiceIamPolicy("policy",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            policy_data=admin.policy_data)
        ```

        ## dataproc.MetastoreServiceIamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.dataproc.MetastoreServiceIamBinding("binding",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## dataproc.MetastoreServiceIamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.dataproc.MetastoreServiceIamMember("member",
            project=default["project"],
            location=default["location"],
            service_id=default["serviceId"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/locations/{{location}}/services/{{service_id}}

        * {{project}}/{{location}}/{{service_id}}

        * {{location}}/{{service_id}}

        * {{service_id}}

        Any variables not passed in the import command will be taken from the provider configuration.

        Dataproc metastore service IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy editor "projects/{{project}}/locations/{{location}}/services/{{service_id}} roles/viewer"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy editor projects/{{project}}/locations/{{location}}/services/{{service_id}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param MetastoreServiceIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MetastoreServiceIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MetastoreServiceIamPolicyArgs.__new__(MetastoreServiceIamPolicyArgs)

            __props__.__dict__["location"] = location
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["etag"] = None
        super(MetastoreServiceIamPolicy, __self__).__init__(
            'gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'MetastoreServiceIamPolicy':
        """
        Get an existing MetastoreServiceIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The location where the metastore service should reside.
               The default value is `global`.
               Used to find the parent resource to bind the IAM policy to. If not specified,
               the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
               location is specified, it is taken from the provider configuration.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MetastoreServiceIamPolicyState.__new__(_MetastoreServiceIamPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["location"] = location
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        __props__.__dict__["service_id"] = service_id
        return MetastoreServiceIamPolicy(resource_name, opts=opts, __props__=__props__)

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
        The location where the metastore service should reside.
        The default value is `global`.
        Used to find the parent resource to bind the IAM policy to. If not specified,
        the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        location is specified, it is taken from the provider configuration.
        """
        return pulumi.get(self, "location")

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

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "service_id")

