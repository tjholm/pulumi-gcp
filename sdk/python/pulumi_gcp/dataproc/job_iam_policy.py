# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['JobIAMPolicyArgs', 'JobIAMPolicy']

@pulumi.input_type
class JobIAMPolicyArgs:
    def __init__(__self__, *,
                 job_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a JobIAMPolicy resource.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations_get_iam_policy` data source.
               
               - - -
        :param pulumi.Input[str] project: The project in which the job belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the job belongs. If it
               is not provided, the provider will use a default.
        """
        pulumi.set(__self__, "job_id", job_id)
        pulumi.set(__self__, "policy_data", policy_data)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by a `organizations_get_iam_policy` data source.

        - - -
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the job belongs. If it
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
        The region in which the job belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _JobIAMPolicyState:
    def __init__(__self__, *,
                 etag: Optional[pulumi.Input[str]] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering JobIAMPolicy resources.
        :param pulumi.Input[str] etag: (Computed) The etag of the jobs's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations_get_iam_policy` data source.
               
               - - -
        :param pulumi.Input[str] project: The project in which the job belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the job belongs. If it
               is not provided, the provider will use a default.
        """
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if job_id is not None:
            pulumi.set(__self__, "job_id", job_id)
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
        (Computed) The etag of the jobs's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "job_id")

    @job_id.setter
    def job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "job_id", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by a `organizations_get_iam_policy` data source.

        - - -
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project in which the job belongs. If it
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
        The region in which the job belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class JobIAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:

        * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
        * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
        * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.

        > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.

        > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_dataproc\\_job\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        editor = gcp.dataproc.JobIAMPolicy("editor",
            project="your-project",
            region="your-region",
            job_id="your-dataproc-job",
            policy_data=admin.policy_data)
        ```

        ## google\\_dataproc\\_job\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.JobIAMBinding("editor",
            job_id="your-dataproc-job",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_dataproc\\_job\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.JobIAMMember("editor",
            job_id="your-dataproc-job",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        ### Importing IAM policies IAM policy imports use the `job_id` identifier of the Dataproc Job resource only. For example* `projects/{project}/regions/{region}/jobs/{job_id}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {

         id = "projects/{project}/regions/{region}/jobs/{job_id}"

         to = google_dataproc_job_iam_policy.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:dataproc/jobIAMPolicy:JobIAMPolicy default "projects/{project}/regions/{region}/jobs/{job_id}"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations_get_iam_policy` data source.
               
               - - -
        :param pulumi.Input[str] project: The project in which the job belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the job belongs. If it
               is not provided, the provider will use a default.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobIAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:

        * `dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
        * `dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
        * `dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.

        > **Note:** `dataproc.JobIAMPolicy` **cannot** be used in conjunction with `dataproc.JobIAMBinding` and `dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `dataproc.JobIAMPolicy` replaces the entire policy.

        > **Note:** `dataproc.JobIAMBinding` resources **can be** used in conjunction with `dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.

        ## google\\_dataproc\\_job\\_iam\\_policy

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/editor",
            members=["user:jane@example.com"],
        )])
        editor = gcp.dataproc.JobIAMPolicy("editor",
            project="your-project",
            region="your-region",
            job_id="your-dataproc-job",
            policy_data=admin.policy_data)
        ```

        ## google\\_dataproc\\_job\\_iam\\_binding

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.JobIAMBinding("editor",
            job_id="your-dataproc-job",
            members=["user:jane@example.com"],
            role="roles/editor")
        ```

        ## google\\_dataproc\\_job\\_iam\\_member

        ```python
        import pulumi
        import pulumi_gcp as gcp

        editor = gcp.dataproc.JobIAMMember("editor",
            job_id="your-dataproc-job",
            member="user:jane@example.com",
            role="roles/editor")
        ```

        ## Import

        ### Importing IAM policies IAM policy imports use the `job_id` identifier of the Dataproc Job resource only. For example* `projects/{project}/regions/{region}/jobs/{job_id}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {

         id = "projects/{project}/regions/{region}/jobs/{job_id}"

         to = google_dataproc_job_iam_policy.default } The `pulumi import` command can also be used

        ```sh
         $ pulumi import gcp:dataproc/jobIAMPolicy:JobIAMPolicy default "projects/{project}/regions/{region}/jobs/{job_id}"
        ```

        :param str resource_name: The name of the resource.
        :param JobIAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobIAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = JobIAMPolicyArgs.__new__(JobIAMPolicyArgs)

            if job_id is None and not opts.urn:
                raise TypeError("Missing required property 'job_id'")
            __props__.__dict__["job_id"] = job_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["etag"] = None
        super(JobIAMPolicy, __self__).__init__(
            'gcp:dataproc/jobIAMPolicy:JobIAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            etag: Optional[pulumi.Input[str]] = None,
            job_id: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'JobIAMPolicy':
        """
        Get an existing JobIAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: (Computed) The etag of the jobs's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by a `organizations_get_iam_policy` data source.
               
               - - -
        :param pulumi.Input[str] project: The project in which the job belongs. If it
               is not provided, the provider will use a default.
        :param pulumi.Input[str] region: The region in which the job belongs. If it
               is not provided, the provider will use a default.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobIAMPolicyState.__new__(_JobIAMPolicyState)

        __props__.__dict__["etag"] = etag
        __props__.__dict__["job_id"] = job_id
        __props__.__dict__["policy_data"] = policy_data
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        return JobIAMPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the jobs's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="jobId")
    def job_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "job_id")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by a `organizations_get_iam_policy` data source.

        - - -
        """
        return pulumi.get(self, "policy_data")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project in which the job belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which the job belongs. If it
        is not provided, the provider will use a default.
        """
        return pulumi.get(self, "region")

