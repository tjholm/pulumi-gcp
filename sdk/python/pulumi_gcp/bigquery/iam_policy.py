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

__all__ = ['IamPolicyArgs', 'IamPolicy']

@pulumi.input_type
class IamPolicyArgs:
    def __init__(__self__, *,
                 dataset_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 table_id: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IamPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        pulumi.set(__self__, "dataset_id", dataset_id)
        pulumi.set(__self__, "policy_data", policy_data)
        pulumi.set(__self__, "table_id", table_id)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset_id", value)

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
    @pulumi.getter(name="tableId")
    def table_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "table_id")

    @table_id.setter
    def table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_id", value)

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
class _IamPolicyState:
    def __init__(__self__, *,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 table_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IamPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        if dataset_id is not None:
            pulumi.set(__self__, "dataset_id", dataset_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if table_id is not None:
            pulumi.set(__self__, "table_id", table_id)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dataset_id")

    @dataset_id.setter
    def dataset_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset_id", value)

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
    @pulumi.getter(name="tableId")
    def table_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "table_id")

    @table_id.setter
    def table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_id", value)


class IamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:

        * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
        * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `bigquery.IamPolicy`: Retrieves the IAM policy for the table

        > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.

        ## bigquery.IamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```
        ## bigquery.IamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## bigquery.IamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for BigQuery Table
        Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:

        * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
        * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `bigquery.IamPolicy`: Retrieves the IAM policy for the table

        > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.

        ## bigquery.IamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```
        ## bigquery.IamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## bigquery.IamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}

        * {{project}}/{{dataset_id}}/{{table_id}}

        * {{dataset_id}}/{{table_id}}

        * {{table_id}}

        Any variables not passed in the import command will be taken from the provider configuration.

        BigQuery table IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:

        * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
        * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `bigquery.IamPolicy`: Retrieves the IAM policy for the table

        > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.

        ## bigquery.IamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```
        ## bigquery.IamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## bigquery.IamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## This resource supports User Project Overrides.

        - 

        # IAM policy for BigQuery Table
        Three different resources help you manage your IAM policy for BigQuery Table. Each of these resources serves a different use case:

        * `bigquery.IamPolicy`: Authoritative. Sets the IAM policy for the table and replaces any existing policy already attached.
        * `bigquery.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
        * `bigquery.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `bigquery.IamPolicy`: Retrieves the IAM policy for the table

        > **Note:** `bigquery.IamPolicy` **cannot** be used in conjunction with `bigquery.IamBinding` and `bigquery.IamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.IamBinding` resources **can be** used in conjunction with `bigquery.IamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.

        ## bigquery.IamPolicy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[{
            "role": "roles/bigquery.dataOwner",
            "members": ["user:jane@example.com"],
            "condition": {
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            },
        }])
        policy = gcp.bigquery.IamPolicy("policy",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            policy_data=admin.policy_data)
        ```
        ## bigquery.IamBinding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.IamBinding("binding",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            members=["user:jane@example.com"],
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```
        ## bigquery.IamMember

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.IamMember("member",
            project=test["project"],
            dataset_id=test["datasetId"],
            table_id=test["tableId"],
            role="roles/bigquery.dataOwner",
            member="user:jane@example.com",
            condition={
                "title": "expires_after_2019_12_31",
                "description": "Expiring at midnight of 2019-12-31",
                "expression": "request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            })
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms:

        * projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}

        * {{project}}/{{dataset_id}}/{{table_id}}

        * {{dataset_id}}/{{table_id}}

        * {{table_id}}

        Any variables not passed in the import command will be taken from the provider configuration.

        BigQuery table IAM resources can be imported using the resource identifiers, role, and member.

        IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.

        ```sh
        $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
        ```

        IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.

        ```sh
        $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
        ```

        IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
        $ pulumi import gcp:bigquery/iamPolicy:IamPolicy editor projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
        ```

        -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the

         full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param IamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IamPolicyArgs.__new__(IamPolicyArgs)

            if dataset_id is None and not opts.urn:
                raise TypeError("Missing required property 'dataset_id'")
            __props__.__dict__["dataset_id"] = dataset_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            if table_id is None and not opts.urn:
                raise TypeError("Missing required property 'table_id'")
            __props__.__dict__["table_id"] = table_id
            __props__.__dict__["etag"] = None
        super(IamPolicy, __self__).__init__(
            'gcp:bigquery/iamPolicy:IamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dataset_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            table_id: Optional[pulumi.Input[str]] = None) -> 'IamPolicy':
        """
        Get an existing IamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IamPolicyState.__new__(_IamPolicyState)

        __props__.__dict__["dataset_id"] = dataset_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        __props__.__dict__["table_id"] = table_id
        return IamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="datasetId")
    def dataset_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dataset_id")

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
    @pulumi.getter(name="tableId")
    def table_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "table_id")

