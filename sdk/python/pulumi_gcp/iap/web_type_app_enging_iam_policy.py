# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WebTypeAppEngingIamPolicyArgs', 'WebTypeAppEngingIamPolicy']

@pulumi.input_type
class WebTypeAppEngingIamPolicyArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a WebTypeAppEngingIamPolicy resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
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
        WebTypeAppEngingIamPolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            app_id=app_id,
            policy_data=policy_data,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             app_id: Optional[pulumi.Input[str]] = None,
             policy_data: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if app_id is None and 'appId' in kwargs:
            app_id = kwargs['appId']
        if app_id is None:
            raise TypeError("Missing 'app_id' argument")
        if policy_data is None and 'policyData' in kwargs:
            policy_data = kwargs['policyData']
        if policy_data is None:
            raise TypeError("Missing 'policy_data' argument")

        _setter("app_id", app_id)
        _setter("policy_data", policy_data)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

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
class _WebTypeAppEngingIamPolicyState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering WebTypeAppEngingIamPolicy resources.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
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
        _WebTypeAppEngingIamPolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            app_id=app_id,
            etag=etag,
            policy_data=policy_data,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             app_id: Optional[pulumi.Input[str]] = None,
             etag: Optional[pulumi.Input[str]] = None,
             policy_data: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if app_id is None and 'appId' in kwargs:
            app_id = kwargs['appId']
        if policy_data is None and 'policyData' in kwargs:
            policy_data = kwargs['policyData']

        if app_id is not None:
            _setter("app_id", app_id)
        if etag is not None:
            _setter("etag", etag)
        if policy_data is not None:
            _setter("policy_data", policy_data)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

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


class WebTypeAppEngingIamPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:

        * `iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
        * `iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
        * `iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `iap.WebTypeAppEngingIamPolicy`: Retrieves the IAM policy for the webtypeappengine

        > **Note:** `iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `iap.WebTypeAppEngingIamBinding` and `iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.

        > **Note:** `iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.

        ## google\\_iap\\_web\\_type\\_app\\_engine\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
        )])
        policy = gcp.iap.WebTypeAppEngingIamPolicy("policy",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ),
        )])
        policy = gcp.iap.WebTypeAppEngingIamPolicy("policy",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            policy_data=admin.policy_data)
        ```
        ## google\\_iap\\_web\\_type\\_app\\_engine\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebTypeAppEngingIamBinding("binding",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebTypeAppEngingIamBinding("binding",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.iap.WebTypeAppEngingIamBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```
        ## google\\_iap\\_web\\_type\\_app\\_engine\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebTypeAppEngingIamMember("member",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebTypeAppEngingIamMember("member",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com",
            condition=gcp.iap.WebTypeAppEngingIamMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}} * {{project}}/{{appId}} * {{appId}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypeappengine IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
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
                 args: WebTypeAppEngingIamPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:

        * `iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
        * `iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
        * `iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.

        A data source can be used to retrieve policy data in advent you do not need creation

        * `iap.WebTypeAppEngingIamPolicy`: Retrieves the IAM policy for the webtypeappengine

        > **Note:** `iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `iap.WebTypeAppEngingIamBinding` and `iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.

        > **Note:** `iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.

        > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.

        ## google\\_iap\\_web\\_type\\_app\\_engine\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
        )])
        policy = gcp.iap.WebTypeAppEngingIamPolicy("policy",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ),
        )])
        policy = gcp.iap.WebTypeAppEngingIamPolicy("policy",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            policy_data=admin.policy_data)
        ```
        ## google\\_iap\\_web\\_type\\_app\\_engine\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebTypeAppEngingIamBinding("binding",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        binding = gcp.iap.WebTypeAppEngingIamBinding("binding",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            members=["user:jane@example.com"],
            condition=gcp.iap.WebTypeAppEngingIamBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```
        ## google\\_iap\\_web\\_type\\_app\\_engine\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebTypeAppEngingIamMember("member",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        member = gcp.iap.WebTypeAppEngingIamMember("member",
            project=google_app_engine_application["app"]["project"],
            app_id=google_app_engine_application["app"]["app_id"],
            role="roles/iap.httpsResourceAccessor",
            member="user:jane@example.com",
            condition=gcp.iap.WebTypeAppEngingIamMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\\"2020-01-01T00:00:00Z\\")",
            ))
        ```

        ## Import

        For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}} * {{project}}/{{appId}} * {{appId}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypeappengine IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.

        ```sh
         $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor user:jane@example.com"
        ```

         IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.

        ```sh
         $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}} roles/iap.httpsResourceAccessor"
        ```

         IAM policy imports use the identifier of the resource in question, e.g.

        ```sh
         $ pulumi import gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param WebTypeAppEngingIamPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WebTypeAppEngingIamPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            WebTypeAppEngingIamPolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WebTypeAppEngingIamPolicyArgs.__new__(WebTypeAppEngingIamPolicyArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["etag"] = None
        super(WebTypeAppEngingIamPolicy, __self__).__init__(
            'gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'WebTypeAppEngingIamPolicy':
        """
        Get an existing WebTypeAppEngingIamPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        :param pulumi.Input[str] etag: (Computed) The etag of the IAM policy.
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

        __props__ = _WebTypeAppEngingIamPolicyState.__new__(_WebTypeAppEngingIamPolicyState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        return WebTypeAppEngingIamPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        """
        return pulumi.get(self, "app_id")

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

