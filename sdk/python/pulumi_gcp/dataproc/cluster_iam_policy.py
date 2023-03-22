# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClusterIAMPolicyArgs', 'ClusterIAMPolicy']

@pulumi.input_type
class ClusterIAMPolicyArgs:
    def __init__(__self__, *,
                 cluster: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ClusterIAMPolicy resource.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a _organizations_get_iam_policy_ data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, the provider will use a default.
        """
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Input[str]:
        """
        The name or relative resource id of the cluster to manage IAM policies for.
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by a _organizations_get_iam_policy_ data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the cluster belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the cluster belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _ClusterIAMPolicyState:
    def __init__(__self__, *,
                 cluster: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClusterIAMPolicy resources.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] etag: (Computed) The etag of the clusters's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by a _organizations_get_iam_policy_ data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, the provider will use a default.
        """
        if cluster is not None:
            pulumi.set(__self__, "cluster", cluster)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def cluster(self) -> Optional[pulumi.Input[str]]:
        """
        The name or relative resource id of the cluster to manage IAM policies for.
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the clusters's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by a _organizations_get_iam_policy_ data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the cluster belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the cluster belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class ClusterIAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:

        * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
        * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
        * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.

        > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.

        > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_dataproc\\_cluster\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        editor = gcp.dataproc.ClusterIAMPolicy("editor",
            project="your-project",
            region="your-region",
            cluster="your-dataproc-cluster",
            policy_data=admin.policy_data)
        ```

        ## google\\_dataproc\\_cluster\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.ClusterIAMBinding("editor",
            cluster="your-dataproc-cluster",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_dataproc\\_cluster\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.ClusterIAMMember("editor",
            cluster="your-dataproc-cluster",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        Cluster IAM resources can be imported using the project, region, cluster name, role and/or member.

        ```sh
         $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster}"
        ```

        ```sh
         $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor"
        ```

        ```sh
         $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] policy_data: The policy data generated by a _organizations_get_iam_policy_ data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, the provider will use a default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterIAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage IAM policies on dataproc clusters. Each of these resources serves a different use case:

        * `dataproc.ClusterIAMPolicy`: Authoritative. Sets the IAM policy for the cluster and replaces any existing policy already attached.
        * `dataproc.ClusterIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the cluster are preserved.
        * `dataproc.ClusterIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the cluster are preserved.

        > **Note:** `dataproc.ClusterIAMPolicy` **cannot** be used in conjunction with `dataproc.ClusterIAMBinding` and `dataproc.ClusterIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the cluster as `dataproc.ClusterIAMPolicy` replaces the entire policy.

        > **Note:** `dataproc.ClusterIAMBinding` resources **can be** used in conjunction with `dataproc.ClusterIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_dataproc\\_cluster\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        editor = gcp.dataproc.ClusterIAMPolicy("editor",
            project="your-project",
            region="your-region",
            cluster="your-dataproc-cluster",
            policy_data=admin.policy_data)
        ```

        ## google\\_dataproc\\_cluster\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.ClusterIAMBinding("editor",
            cluster="your-dataproc-cluster",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_dataproc\\_cluster\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.ClusterIAMMember("editor",
            cluster="your-dataproc-cluster",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        Cluster IAM resources can be imported using the project, region, cluster name, role and/or member.

        ```sh
         $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster}"
        ```

        ```sh
         $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor"
        ```

        ```sh
         $ pulumi import gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com"
        ```

         -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the

        full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.

        :param str resource_name: The name of the resource.
        :param ClusterIAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterIAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterIAMPolicyArgs.__new__(ClusterIAMPolicyArgs)

            if cluster is None and not opts.urn:
                raise TypeError("Missing required property 'cluster'")
            __props__.__dict__["cluster"] = cluster
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["etag"] = None
        super(ClusterIAMPolicy, __self__).__init__(
            'gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'ClusterIAMPolicy':
        """
        Get an existing ClusterIAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster: The name or relative resource id of the cluster to manage IAM policies for.
        :param pulumi.Input[str] etag: (Computed) The etag of the clusters's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by a _organizations_get_iam_policy_ data source.
        :param pulumi.Input[str] project: The project in which the cluster belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the cluster belongs. If it
               is not provided, the provider will use a default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterIAMPolicyState.__new__(_ClusterIAMPolicyState)

        __props__.__dict__["cluster"] = cluster
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        return ClusterIAMPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[str]:
        """
        The name or relative resource id of the cluster to manage IAM policies for.
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the clusters's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by a _organizations_get_iam_policy_ data source.
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project in which the cluster belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which the cluster belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "region")

