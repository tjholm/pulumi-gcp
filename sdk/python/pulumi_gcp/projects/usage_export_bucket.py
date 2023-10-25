# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UsageExportBucketArgs', 'UsageExportBucket']

@pulumi.input_type
class UsageExportBucketArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UsageExportBucket resource.
        :param pulumi.Input[str] bucket_name: The bucket to store reports in.
        :param pulumi.Input[str] prefix: A prefix for the reports, for instance, the project name.
        :param pulumi.Input[str] project: The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        UsageExportBucketArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            bucket_name=bucket_name,
            prefix=prefix,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             bucket_name: Optional[pulumi.Input[str]] = None,
             prefix: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if bucket_name is None and 'bucketName' in kwargs:
            bucket_name = kwargs['bucketName']
        if bucket_name is None:
            raise TypeError("Missing 'bucket_name' argument")

        _setter("bucket_name", bucket_name)
        if prefix is not None:
            _setter("prefix", prefix)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        The bucket to store reports in.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A prefix for the reports, for instance, the project name.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _UsageExportBucketState:
    def __init__(__self__, *,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UsageExportBucket resources.
        :param pulumi.Input[str] bucket_name: The bucket to store reports in.
        :param pulumi.Input[str] prefix: A prefix for the reports, for instance, the project name.
        :param pulumi.Input[str] project: The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        _UsageExportBucketState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            bucket_name=bucket_name,
            prefix=prefix,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             bucket_name: Optional[pulumi.Input[str]] = None,
             prefix: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if bucket_name is None and 'bucketName' in kwargs:
            bucket_name = kwargs['bucketName']

        if bucket_name is not None:
            _setter("bucket_name", bucket_name)
        if prefix is not None:
            _setter("prefix", prefix)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The bucket to store reports in.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        A prefix for the reports, for instance, the project name.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class UsageExportBucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows creation and management of a Google Cloud Platform project.

        Projects created with this resource must be associated with an Organization.
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.

        The user or service account that is running this provider when creating a `organizations.Project`
        resource must have `roles/resourcemanager.projectCreator` on the specified organization. See the
        [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
        doc for more information.

        > This resource reads the specified billing account on every pulumi up and plan operation so you must have permissions on the specified billing account.

        To get more information about projects, see:

        * [API documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects)
        * How-to Guides
            * [Creating and managing projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            org_id="1234567",
            project_id="your-project-id")
        ```

        To create a project under a specific folder

        ```python
        import pulumi
        import pulumi_gcp as gcp

        department1 = gcp.organizations.Folder("department1",
            display_name="Department 1",
            parent="organizations/1234567")
        my_project_in_a_folder = gcp.organizations.Project("myProject-in-a-folder",
            project_id="your-project-id",
            folder_id=department1.name)
        ```

        ## Import

        Projects can be imported using the `project_id`, e.g.

        ```sh
         $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: The bucket to store reports in.
        :param pulumi.Input[str] prefix: A prefix for the reports, for instance, the project name.
        :param pulumi.Input[str] project: The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UsageExportBucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows creation and management of a Google Cloud Platform project.

        Projects created with this resource must be associated with an Organization.
        See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.

        The user or service account that is running this provider when creating a `organizations.Project`
        resource must have `roles/resourcemanager.projectCreator` on the specified organization. See the
        [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
        doc for more information.

        > This resource reads the specified billing account on every pulumi up and plan operation so you must have permissions on the specified billing account.

        To get more information about projects, see:

        * [API documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects)
        * How-to Guides
            * [Creating and managing projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects)

        ## Example Usage

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_project = gcp.organizations.Project("myProject",
            org_id="1234567",
            project_id="your-project-id")
        ```

        To create a project under a specific folder

        ```python
        import pulumi
        import pulumi_gcp as gcp

        department1 = gcp.organizations.Folder("department1",
            display_name="Department 1",
            parent="organizations/1234567")
        my_project_in_a_folder = gcp.organizations.Project("myProject-in-a-folder",
            project_id="your-project-id",
            folder_id=department1.name)
        ```

        ## Import

        Projects can be imported using the `project_id`, e.g.

        ```sh
         $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
        ```

        :param str resource_name: The name of the resource.
        :param UsageExportBucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UsageExportBucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            UsageExportBucketArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UsageExportBucketArgs.__new__(UsageExportBucketArgs)

            if bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_name'")
            __props__.__dict__["bucket_name"] = bucket_name
            __props__.__dict__["prefix"] = prefix
            __props__.__dict__["project"] = project
        super(UsageExportBucket, __self__).__init__(
            'gcp:projects/usageExportBucket:UsageExportBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket_name: Optional[pulumi.Input[str]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'UsageExportBucket':
        """
        Get an existing UsageExportBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: The bucket to store reports in.
        :param pulumi.Input[str] prefix: A prefix for the reports, for instance, the project name.
        :param pulumi.Input[str] project: The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UsageExportBucketState.__new__(_UsageExportBucketState)

        __props__.__dict__["bucket_name"] = bucket_name
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["project"] = project
        return UsageExportBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[str]:
        """
        The bucket to store reports in.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[Optional[str]]:
        """
        A prefix for the reports, for instance, the project name.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project to set the export bucket on. If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

