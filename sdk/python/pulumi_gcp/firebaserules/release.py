# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReleaseArgs', 'Release']

@pulumi.input_type
class ReleaseArgs:
    def __init__(__self__, *,
                 ruleset_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Release resource.
        :param pulumi.Input[str] ruleset_name: Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        :param pulumi.Input[str] name: Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        :param pulumi.Input[str] project: The project for the resource
        """
        pulumi.set(__self__, "ruleset_name", ruleset_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="rulesetName")
    def ruleset_name(self) -> pulumi.Input[str]:
        """
        Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        """
        return pulumi.get(self, "ruleset_name")

    @ruleset_name.setter
    def ruleset_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "ruleset_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _ReleaseState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ruleset_name: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Release resources.
        :param pulumi.Input[str] create_time: Output only. Time the release was created.
        :param pulumi.Input[bool] disabled: Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.
        :param pulumi.Input[str] name: Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] ruleset_name: Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        :param pulumi.Input[str] update_time: Output only. Time the release was updated.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if ruleset_name is not None:
            pulumi.set(__self__, "ruleset_name", ruleset_name)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Time the release was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="rulesetName")
    def ruleset_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        """
        return pulumi.get(self, "ruleset_name")

    @ruleset_name.setter
    def ruleset_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ruleset_name", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. Time the release was updated.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Release(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ruleset_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        For more information, see:
        * [Get started with Firebase Security Rules](https://firebase.google.com/docs/rules/get-started)
        ## Example Usage
        ### Firestore_release
        Creates a Firebase Rules Release to Cloud Firestore
        ```python
        import pulumi
        import pulumi_gcp as gcp

        firestore = gcp.firebaserules.Ruleset("firestore",
            source=gcp.firebaserules.RulesetSourceArgs(
                files=[gcp.firebaserules.RulesetSourceFileArgs(
                    content="service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }",
                    name="firestore.rules",
                )],
            ),
            project="my-project-name")
        primary = gcp.firebaserules.Release("primary",
            ruleset_name=firestore.name.apply(lambda name: f"projects/my-project-name/rulesets/{name}"),
            project="my-project-name")
        ```
        ### Storage_release
        Creates a Firebase Rules Release for a Storage bucket
        ```python
        import pulumi
        import pulumi_gcp as gcp

        # Provision a non-default Cloud Storage bucket.
        bucket_bucket = gcp.storage.Bucket("bucketBucket",
            project="my-project-name",
            location="us-west1")
        # Make the Storage bucket accessible for Firebase SDKs, authentication, and Firebase Security Rules.
        bucket_storage_bucket = gcp.firebase.StorageBucket("bucketStorageBucket",
            project="my-project-name",
            bucket_id=bucket_bucket.name)
        # Create a ruleset of Firebase Security Rules from a local file.
        storage = gcp.firebaserules.Ruleset("storage",
            project="my-project-name",
            source=gcp.firebaserules.RulesetSourceArgs(
                files=[gcp.firebaserules.RulesetSourceFileArgs(
                    name="storage.rules",
                    content="service firebase.storage {match /b/{bucket}/o {match /{allPaths=**} {allow read, write: if request.auth != null;}}}",
                )],
            ),
            opts=pulumi.ResourceOptions(depends_on=[bucket_storage_bucket]))
        primary = gcp.firebaserules.Release("primary",
            ruleset_name=storage.name.apply(lambda name: f"projects/my-project-name/rulesets/{name}"),
            project="my-project-name")
        ```

        ## Import

        Release can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:firebaserules/release:Release default projects/{{project}}/releases/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] ruleset_name: Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ReleaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        For more information, see:
        * [Get started with Firebase Security Rules](https://firebase.google.com/docs/rules/get-started)
        ## Example Usage
        ### Firestore_release
        Creates a Firebase Rules Release to Cloud Firestore
        ```python
        import pulumi
        import pulumi_gcp as gcp

        firestore = gcp.firebaserules.Ruleset("firestore",
            source=gcp.firebaserules.RulesetSourceArgs(
                files=[gcp.firebaserules.RulesetSourceFileArgs(
                    content="service cloud.firestore {match /databases/{database}/documents { match /{document=**} { allow read, write: if false; } } }",
                    name="firestore.rules",
                )],
            ),
            project="my-project-name")
        primary = gcp.firebaserules.Release("primary",
            ruleset_name=firestore.name.apply(lambda name: f"projects/my-project-name/rulesets/{name}"),
            project="my-project-name")
        ```
        ### Storage_release
        Creates a Firebase Rules Release for a Storage bucket
        ```python
        import pulumi
        import pulumi_gcp as gcp

        # Provision a non-default Cloud Storage bucket.
        bucket_bucket = gcp.storage.Bucket("bucketBucket",
            project="my-project-name",
            location="us-west1")
        # Make the Storage bucket accessible for Firebase SDKs, authentication, and Firebase Security Rules.
        bucket_storage_bucket = gcp.firebase.StorageBucket("bucketStorageBucket",
            project="my-project-name",
            bucket_id=bucket_bucket.name)
        # Create a ruleset of Firebase Security Rules from a local file.
        storage = gcp.firebaserules.Ruleset("storage",
            project="my-project-name",
            source=gcp.firebaserules.RulesetSourceArgs(
                files=[gcp.firebaserules.RulesetSourceFileArgs(
                    name="storage.rules",
                    content="service firebase.storage {match /b/{bucket}/o {match /{allPaths=**} {allow read, write: if request.auth != null;}}}",
                )],
            ),
            opts=pulumi.ResourceOptions(depends_on=[bucket_storage_bucket]))
        primary = gcp.firebaserules.Release("primary",
            ruleset_name=storage.name.apply(lambda name: f"projects/my-project-name/rulesets/{name}"),
            project="my-project-name")
        ```

        ## Import

        Release can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:firebaserules/release:Release default projects/{{project}}/releases/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ReleaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReleaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ruleset_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReleaseArgs.__new__(ReleaseArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if ruleset_name is None and not opts.urn:
                raise TypeError("Missing required property 'ruleset_name'")
            __props__.__dict__["ruleset_name"] = ruleset_name
            __props__.__dict__["create_time"] = None
            __props__.__dict__["disabled"] = None
            __props__.__dict__["update_time"] = None
        super(Release, __self__).__init__(
            'gcp:firebaserules/release:Release',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            ruleset_name: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Release':
        """
        Get an existing Release resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Output only. Time the release was created.
        :param pulumi.Input[bool] disabled: Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.
        :param pulumi.Input[str] name: Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        :param pulumi.Input[str] project: The project for the resource
        :param pulumi.Input[str] ruleset_name: Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        :param pulumi.Input[str] update_time: Output only. Time the release was updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReleaseState.__new__(_ReleaseState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["ruleset_name"] = ruleset_name
        __props__.__dict__["update_time"] = update_time
        return Release(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Output only. Time the release was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[bool]:
        """
        Disable the release to keep it from being served. The response code of NOT_FOUND will be given for executables generated from this Release.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Format: `projects/{project_id}/releases/{release_id}`\\Firestore Rules Releases will **always** have the name 'cloud.firestore'
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project for the resource
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="rulesetName")
    def ruleset_name(self) -> pulumi.Output[str]:
        """
        Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist for the `Release` to be created.
        """
        return pulumi.get(self, "ruleset_name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Output only. Time the release was updated.
        """
        return pulumi.get(self, "update_time")

