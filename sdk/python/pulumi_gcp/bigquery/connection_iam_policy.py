# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConnectionIamPolicyArgs', 'ConnectionIamPolicy']

@pulumi.input_type
class ConnectionIamPolicyArgs:
    def __init__(__self__, *,
                 connection_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConnectionIamPolicy resource.
        :param pulumi.Input[str] connection_id: Optional connection id that should be assigned to the created connection.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] location: The geographic location where the connection should reside.
               Cloud SQL instance must be in the same location as the connection
               with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
               Examples: US, EU, asia-northeast1, us-central1, europe-west1.
               Spanner Connections same as spanner region
               AWS allowed regions are aws-us-east-1
               Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        ConnectionIamPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            connection_id=connection_id,
            policy_data=policy_data,
            location=location,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             connection_id: Optional[pulumi.Input[str]] = None,
             policy_data: Optional[pulumi.Input[str]] = None,
             location: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if connection_id is None and 'connectionId' in kwargs:
            connection_id = kwargs['connectionId']
        if connection_id is None:
            raise TypeError("Missing 'connection_id' argument")
        if policy_data is None and 'policyData' in kwargs:
            policy_data = kwargs['policyData']
        if policy_data is None:
            raise TypeError("Missing 'policy_data' argument")

        _setter("connection_id", connection_id)
        _setter("policy_data", policy_data)
        if location is not None:
            _setter("location", location)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Input[str]:
        """
        Optional connection id that should be assigned to the created connection.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "connection_id", value)

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
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The geographic location where the connection should reside.
        Cloud SQL instance must be in the same location as the connection
        with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
        Examples: US, EU, asia-northeast1, us-central1, europe-west1.
        Spanner Connections same as spanner region
        AWS allowed regions are aws-us-east-1
        Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
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

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ConnectionIamPolicyState:
    def __init__(__self__, *,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConnectionIamPolicy resources.
        :param pulumi.Input[str] connection_id: Optional connection id that should be assigned to the created connection.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The geographic location where the connection should reside.
               Cloud SQL instance must be in the same location as the connection
               with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
               Examples: US, EU, asia-northeast1, us-central1, europe-west1.
               Spanner Connections same as spanner region
               AWS allowed regions are aws-us-east-1
               Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        _ConnectionIamPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            connection_id=connection_id,
            etag=etag,
            location=location,
            policy_data=policy_data,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             connection_id: Optional[pulumi.Input[str]] = None,
             etag: Optional[pulumi.Input[str]] = None,
             location: Optional[pulumi.Input[str]] = None,
             policy_data: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if connection_id is None and 'connectionId' in kwargs:
            connection_id = kwargs['connectionId']
        if policy_data is None and 'policyData' in kwargs:
            policy_data = kwargs['policyData']

        if connection_id is not None:
            _setter("connection_id", connection_id)
        if etag is not None:
            _setter("etag", etag)
        if location is not None:
            _setter("location", location)
        if policy_data is not None:
            _setter("policy_data", policy_data)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional connection id that should be assigned to the created connection.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "connection_id")

    @connection_id.setter
    def connection_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_id", value)

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
        The geographic location where the connection should reside.
        Cloud SQL instance must be in the same location as the connection
        with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
        Examples: US, EU, asia-northeast1, us-central1, europe-west1.
        Spanner Connections same as spanner region
        AWS allowed regions are aws-us-east-1
        Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
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

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class ConnectionIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for BigQuery Connection Connection. Each of these resources serves a different use case:

        * `bigquery.ConnectionIamPolicy`: Authoritative. Sets the IAM policy for the connection and replaces any existing policy already attached.
        * `bigquery.ConnectionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the connection are preserved.
        * `bigquery.ConnectionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the connection are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `bigquery.ConnectionIamPolicy`: Retrieves the IAM policy for the connection

        > **Note:** `bigquery.ConnectionIamPolicy` **cannot** be used in conjunction with `bigquery.ConnectionIamBinding` and `bigquery.ConnectionIamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.ConnectionIamBinding` resources **can be** used in conjunction with `bigquery.ConnectionIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_bigquery\\_connection\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.bigquery.ConnectionIamPolicy("policy",
            project=google_bigquery_connection["connection"]["project"],
            location=google_bigquery_connection["connection"]["location"],
            connection_id=google_bigquery_connection["connection"]["connection_id"],
            policy_data=admin.policy_data)
        ```

        ## google\\_bigquery\\_connection\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.ConnectionIamBinding("binding",
            project=google_bigquery_connection["connection"]["project"],
            location=google_bigquery_connection["connection"]["location"],
            connection_id=google_bigquery_connection["connection"]["connection_id"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_bigquery\\_connection\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.ConnectionIamMember("member",
            project=google_bigquery_connection["connection"]["project"],
            location=google_bigquery_connection["connection"]["location"],
            connection_id=google_bigquery_connection["connection"]["connection_id"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/connections/{{connection_id}} * {{project}}/{{location}}/{{connection_id}} * {{location}}/{{connection_id}} * {{connection_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Connection connection IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor "projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor "projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_id: Optional connection id that should be assigned to the created connection.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] location: The geographic location where the connection should reside.
               Cloud SQL instance must be in the same location as the connection
               with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
               Examples: US, EU, asia-northeast1, us-central1, europe-west1.
               Spanner Connections same as spanner region
               AWS allowed regions are aws-us-east-1
               Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConnectionIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for BigQuery Connection Connection. Each of these resources serves a different use case:

        * `bigquery.ConnectionIamPolicy`: Authoritative. Sets the IAM policy for the connection and replaces any existing policy already attached.
        * `bigquery.ConnectionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the connection are preserved.
        * `bigquery.ConnectionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the connection are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `bigquery.ConnectionIamPolicy`: Retrieves the IAM policy for the connection

        > **Note:** `bigquery.ConnectionIamPolicy` **cannot** be used in conjunction with `bigquery.ConnectionIamBinding` and `bigquery.ConnectionIamMember` or they will fight over what your policy should be.

        > **Note:** `bigquery.ConnectionIamBinding` resources **can be** used in conjunction with `bigquery.ConnectionIamMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_bigquery\\_connection\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/viewer",
            members=["user:jane@example.com"],
        )])
        policy = gcp.bigquery.ConnectionIamPolicy("policy",
            project=google_bigquery_connection["connection"]["project"],
            location=google_bigquery_connection["connection"]["location"],
            connection_id=google_bigquery_connection["connection"]["connection_id"],
            policy_data=admin.policy_data)
        ```

        ## google\\_bigquery\\_connection\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.bigquery.ConnectionIamBinding("binding",
            project=google_bigquery_connection["connection"]["project"],
            location=google_bigquery_connection["connection"]["location"],
            connection_id=google_bigquery_connection["connection"]["connection_id"],
            role="roles/viewer",
            members=["user:jane@example.com"])
        ```

        ## google\\_bigquery\\_connection\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.bigquery.ConnectionIamMember("member",
            project=google_bigquery_connection["connection"]["project"],
            location=google_bigquery_connection["connection"]["location"],
            connection_id=google_bigquery_connection["connection"]["connection_id"],
            role="roles/viewer",
            member="user:jane@example.com")
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/connections/{{connection_id}} * {{project}}/{{location}}/{{connection_id}} * {{location}}/{{connection_id}} * {{connection_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Connection connection IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor "projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor "projects/{{project}}/locations/{{location}}/connections/{{connection_id}} roles/viewer"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy editor projects/{{project}}/locations/{{location}}/connections/{{connection_id}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param ConnectionIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConnectionIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ConnectionIamPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConnectionIamPolicyArgs.__new__(ConnectionIamPolicyArgs)

            if connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'connection_id'")
            __props__.__dict__["connection_id"] = connection_id
            __props__.__dict__["location"] = location
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["etag"] = None
        super(ConnectionIamPolicy, __self__).__init__(
            'gcp:bigquery/connectionIamPolicy:ConnectionIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            location: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'ConnectionIamPolicy':
        """
        Get an existing ConnectionIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_id: Optional connection id that should be assigned to the created connection.
               Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
        :param pulumi.Input[str] location: The geographic location where the connection should reside.
               Cloud SQL instance must be in the same location as the connection
               with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
               Examples: US, EU, asia-northeast1, us-central1, europe-west1.
               Spanner Connections same as spanner region
               AWS allowed regions are aws-us-east-1
               Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations_get_iam_policy` data source.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
               
               * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConnectionIamPolicyState.__new__(_ConnectionIamPolicyState)

        __props__.__dict__["connection_id"] = connection_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["location"] = location
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        return ConnectionIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionId")
    def connection_id(self) -> pulumi.Output[str]:
        """
        Optional connection id that should be assigned to the created connection.
        Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "connection_id")

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
        The geographic location where the connection should reside.
        Cloud SQL instance must be in the same location as the connection
        with following exceptions: Cloud SQL us-central1 maps to BigQuery US, Cloud SQL europe-west1 maps to BigQuery EU.
        Examples: US, EU, asia-northeast1, us-central1, europe-west1.
        Spanner Connections same as spanner region
        AWS allowed regions are aws-us-east-1
        Azure allowed regions are azure-eastus2 Used to find the parent resource to bind the IAM policy to
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

        * `member/members` - (Required) Identities that will be granted the privilege in `role`.
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
        return pulumi.get(self, "project")

