# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pubsub_configs: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[str] name: Resource name of the repository, of the form `{{repo}}`.
               The repo name may contain slashes. eg, `name/with/slash`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]] pubsub_configs: How this repository publishes a change in the repository through Cloud Pub/Sub.
               Keyed by the topic names.
               Structure is documented below.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if pubsub_configs is not None:
            pulumi.set(__self__, "pubsub_configs", pubsub_configs)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the repository, of the form `{{repo}}`.
        The repo name may contain slashes. eg, `name/with/slash`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="pubsubConfigs")
    def pubsub_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]]]:
        """
        How this repository publishes a change in the repository through Cloud Pub/Sub.
        Keyed by the topic names.
        Structure is documented below.
        """
        return pulumi.get(self, "pubsub_configs")

    @pubsub_configs.setter
    def pubsub_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]]]):
        pulumi.set(self, "pubsub_configs", value)


@pulumi.input_type
class _RepositoryState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pubsub_configs: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Repository resources.
        :param pulumi.Input[str] name: Resource name of the repository, of the form `{{repo}}`.
               The repo name may contain slashes. eg, `name/with/slash`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]] pubsub_configs: How this repository publishes a change in the repository through Cloud Pub/Sub.
               Keyed by the topic names.
               Structure is documented below.
        :param pulumi.Input[int] size: The disk usage of the repo, in bytes.
        :param pulumi.Input[str] url: URL to clone the repository from Google Cloud Source Repositories.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if pubsub_configs is not None:
            pulumi.set(__self__, "pubsub_configs", pubsub_configs)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Resource name of the repository, of the form `{{repo}}`.
        The repo name may contain slashes. eg, `name/with/slash`
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="pubsubConfigs")
    def pubsub_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]]]:
        """
        How this repository publishes a change in the repository through Cloud Pub/Sub.
        Keyed by the topic names.
        Structure is documented below.
        """
        return pulumi.get(self, "pubsub_configs")

    @pubsub_configs.setter
    def pubsub_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryPubsubConfigArgs']]]]):
        pulumi.set(self, "pubsub_configs", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[int]]:
        """
        The disk usage of the repo, in bytes.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        URL to clone the repository from Google Cloud Source Repositories.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pubsub_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryPubsubConfigArgs']]]]] = None,
                 __props__=None):
        """
        A repository (or repo) is a Git repository storing versioned source content.

        To get more information about Repository, see:

        * [API documentation](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/source-repositories/)

        ## Example Usage
        ### Sourcerepo Repository Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_repo = gcp.sourcerepo.Repository("my-repo")
        ```
        ### Sourcerepo Repository Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_account = gcp.service_account.Account("testAccount",
            account_id="my-account",
            display_name="Test Service Account")
        topic = gcp.pubsub.Topic("topic")
        my_repo = gcp.sourcerepo.Repository("my-repo", pubsub_configs=[gcp.sourcerepo.RepositoryPubsubConfigArgs(
            topic=topic.id,
            message_format="JSON",
            service_account_email=test_account.email,
        )])
        ```

        ## Import

        Repository can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:sourcerepo/repository:Repository default projects/{{project}}/repos/{{name}}
        ```

        ```sh
         $ pulumi import gcp:sourcerepo/repository:Repository default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Resource name of the repository, of the form `{{repo}}`.
               The repo name may contain slashes. eg, `name/with/slash`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryPubsubConfigArgs']]]] pubsub_configs: How this repository publishes a change in the repository through Cloud Pub/Sub.
               Keyed by the topic names.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RepositoryArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A repository (or repo) is a Git repository storing versioned source content.

        To get more information about Repository, see:

        * [API documentation](https://cloud.google.com/source-repositories/docs/reference/rest/v1/projects.repos)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/source-repositories/)

        ## Example Usage
        ### Sourcerepo Repository Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        my_repo = gcp.sourcerepo.Repository("my-repo")
        ```
        ### Sourcerepo Repository Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        test_account = gcp.service_account.Account("testAccount",
            account_id="my-account",
            display_name="Test Service Account")
        topic = gcp.pubsub.Topic("topic")
        my_repo = gcp.sourcerepo.Repository("my-repo", pubsub_configs=[gcp.sourcerepo.RepositoryPubsubConfigArgs(
            topic=topic.id,
            message_format="JSON",
            service_account_email=test_account.email,
        )])
        ```

        ## Import

        Repository can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:sourcerepo/repository:Repository default projects/{{project}}/repos/{{name}}
        ```

        ```sh
         $ pulumi import gcp:sourcerepo/repository:Repository default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 pubsub_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryPubsubConfigArgs']]]]] = None,
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
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["pubsub_configs"] = pubsub_configs
            __props__.__dict__["size"] = None
            __props__.__dict__["url"] = None
        super(Repository, __self__).__init__(
            'gcp:sourcerepo/repository:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            pubsub_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryPubsubConfigArgs']]]]] = None,
            size: Optional[pulumi.Input[int]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Resource name of the repository, of the form `{{repo}}`.
               The repo name may contain slashes. eg, `name/with/slash`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RepositoryPubsubConfigArgs']]]] pubsub_configs: How this repository publishes a change in the repository through Cloud Pub/Sub.
               Keyed by the topic names.
               Structure is documented below.
        :param pulumi.Input[int] size: The disk usage of the repo, in bytes.
        :param pulumi.Input[str] url: URL to clone the repository from Google Cloud Source Repositories.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryState.__new__(_RepositoryState)

        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["pubsub_configs"] = pubsub_configs
        __props__.__dict__["size"] = size
        __props__.__dict__["url"] = url
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the repository, of the form `{{repo}}`.
        The repo name may contain slashes. eg, `name/with/slash`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pubsubConfigs")
    def pubsub_configs(self) -> pulumi.Output[Optional[Sequence['outputs.RepositoryPubsubConfig']]]:
        """
        How this repository publishes a change in the repository through Cloud Pub/Sub.
        Keyed by the topic names.
        Structure is documented below.
        """
        return pulumi.get(self, "pubsub_configs")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The disk usage of the repo, in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        URL to clone the repository from Google Cloud Source Repositories.
        """
        return pulumi.get(self, "url")

