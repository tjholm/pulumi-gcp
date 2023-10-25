# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetTriggerResult',
    'AwaitableGetTriggerResult',
    'get_trigger',
    'get_trigger_output',
]

@pulumi.output_type
class GetTriggerResult:
    """
    A collection of values returned by getTrigger.
    """
    def __init__(__self__, approval_configs=None, bitbucket_server_trigger_configs=None, builds=None, create_time=None, description=None, disabled=None, filename=None, filter=None, git_file_sources=None, githubs=None, id=None, ignored_files=None, include_build_logs=None, included_files=None, location=None, name=None, project=None, pubsub_configs=None, repository_event_configs=None, service_account=None, source_to_builds=None, substitutions=None, tags=None, trigger_id=None, trigger_templates=None, webhook_configs=None):
        if approval_configs and not isinstance(approval_configs, list):
            raise TypeError("Expected argument 'approval_configs' to be a list")
        pulumi.set(__self__, "approval_configs", approval_configs)
        if bitbucket_server_trigger_configs and not isinstance(bitbucket_server_trigger_configs, list):
            raise TypeError("Expected argument 'bitbucket_server_trigger_configs' to be a list")
        pulumi.set(__self__, "bitbucket_server_trigger_configs", bitbucket_server_trigger_configs)
        if builds and not isinstance(builds, list):
            raise TypeError("Expected argument 'builds' to be a list")
        pulumi.set(__self__, "builds", builds)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if filename and not isinstance(filename, str):
            raise TypeError("Expected argument 'filename' to be a str")
        pulumi.set(__self__, "filename", filename)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if git_file_sources and not isinstance(git_file_sources, list):
            raise TypeError("Expected argument 'git_file_sources' to be a list")
        pulumi.set(__self__, "git_file_sources", git_file_sources)
        if githubs and not isinstance(githubs, list):
            raise TypeError("Expected argument 'githubs' to be a list")
        pulumi.set(__self__, "githubs", githubs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ignored_files and not isinstance(ignored_files, list):
            raise TypeError("Expected argument 'ignored_files' to be a list")
        pulumi.set(__self__, "ignored_files", ignored_files)
        if include_build_logs and not isinstance(include_build_logs, str):
            raise TypeError("Expected argument 'include_build_logs' to be a str")
        pulumi.set(__self__, "include_build_logs", include_build_logs)
        if included_files and not isinstance(included_files, list):
            raise TypeError("Expected argument 'included_files' to be a list")
        pulumi.set(__self__, "included_files", included_files)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if pubsub_configs and not isinstance(pubsub_configs, list):
            raise TypeError("Expected argument 'pubsub_configs' to be a list")
        pulumi.set(__self__, "pubsub_configs", pubsub_configs)
        if repository_event_configs and not isinstance(repository_event_configs, list):
            raise TypeError("Expected argument 'repository_event_configs' to be a list")
        pulumi.set(__self__, "repository_event_configs", repository_event_configs)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if source_to_builds and not isinstance(source_to_builds, list):
            raise TypeError("Expected argument 'source_to_builds' to be a list")
        pulumi.set(__self__, "source_to_builds", source_to_builds)
        if substitutions and not isinstance(substitutions, dict):
            raise TypeError("Expected argument 'substitutions' to be a dict")
        pulumi.set(__self__, "substitutions", substitutions)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if trigger_id and not isinstance(trigger_id, str):
            raise TypeError("Expected argument 'trigger_id' to be a str")
        pulumi.set(__self__, "trigger_id", trigger_id)
        if trigger_templates and not isinstance(trigger_templates, list):
            raise TypeError("Expected argument 'trigger_templates' to be a list")
        pulumi.set(__self__, "trigger_templates", trigger_templates)
        if webhook_configs and not isinstance(webhook_configs, list):
            raise TypeError("Expected argument 'webhook_configs' to be a list")
        pulumi.set(__self__, "webhook_configs", webhook_configs)

    @property
    @pulumi.getter(name="approvalConfigs")
    def approval_configs(self) -> Sequence['outputs.GetTriggerApprovalConfigResult']:
        return pulumi.get(self, "approval_configs")

    @property
    @pulumi.getter(name="bitbucketServerTriggerConfigs")
    def bitbucket_server_trigger_configs(self) -> Sequence['outputs.GetTriggerBitbucketServerTriggerConfigResult']:
        return pulumi.get(self, "bitbucket_server_trigger_configs")

    @property
    @pulumi.getter
    def builds(self) -> Sequence['outputs.GetTriggerBuildResult']:
        return pulumi.get(self, "builds")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def filename(self) -> str:
        return pulumi.get(self, "filename")

    @property
    @pulumi.getter
    def filter(self) -> str:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter(name="gitFileSources")
    def git_file_sources(self) -> Sequence['outputs.GetTriggerGitFileSourceResult']:
        return pulumi.get(self, "git_file_sources")

    @property
    @pulumi.getter
    def githubs(self) -> Sequence['outputs.GetTriggerGithubResult']:
        return pulumi.get(self, "githubs")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ignoredFiles")
    def ignored_files(self) -> Sequence[str]:
        return pulumi.get(self, "ignored_files")

    @property
    @pulumi.getter(name="includeBuildLogs")
    def include_build_logs(self) -> str:
        return pulumi.get(self, "include_build_logs")

    @property
    @pulumi.getter(name="includedFiles")
    def included_files(self) -> Sequence[str]:
        return pulumi.get(self, "included_files")

    @property
    @pulumi.getter
    def location(self) -> str:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="pubsubConfigs")
    def pubsub_configs(self) -> Sequence['outputs.GetTriggerPubsubConfigResult']:
        return pulumi.get(self, "pubsub_configs")

    @property
    @pulumi.getter(name="repositoryEventConfigs")
    def repository_event_configs(self) -> Sequence['outputs.GetTriggerRepositoryEventConfigResult']:
        return pulumi.get(self, "repository_event_configs")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="sourceToBuilds")
    def source_to_builds(self) -> Sequence['outputs.GetTriggerSourceToBuildResult']:
        return pulumi.get(self, "source_to_builds")

    @property
    @pulumi.getter
    def substitutions(self) -> Mapping[str, str]:
        return pulumi.get(self, "substitutions")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="triggerId")
    def trigger_id(self) -> str:
        return pulumi.get(self, "trigger_id")

    @property
    @pulumi.getter(name="triggerTemplates")
    def trigger_templates(self) -> Sequence['outputs.GetTriggerTriggerTemplateResult']:
        return pulumi.get(self, "trigger_templates")

    @property
    @pulumi.getter(name="webhookConfigs")
    def webhook_configs(self) -> Sequence['outputs.GetTriggerWebhookConfigResult']:
        return pulumi.get(self, "webhook_configs")


class AwaitableGetTriggerResult(GetTriggerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTriggerResult(
            approval_configs=self.approval_configs,
            bitbucket_server_trigger_configs=self.bitbucket_server_trigger_configs,
            builds=self.builds,
            create_time=self.create_time,
            description=self.description,
            disabled=self.disabled,
            filename=self.filename,
            filter=self.filter,
            git_file_sources=self.git_file_sources,
            githubs=self.githubs,
            id=self.id,
            ignored_files=self.ignored_files,
            include_build_logs=self.include_build_logs,
            included_files=self.included_files,
            location=self.location,
            name=self.name,
            project=self.project,
            pubsub_configs=self.pubsub_configs,
            repository_event_configs=self.repository_event_configs,
            service_account=self.service_account,
            source_to_builds=self.source_to_builds,
            substitutions=self.substitutions,
            tags=self.tags,
            trigger_id=self.trigger_id,
            trigger_templates=self.trigger_templates,
            webhook_configs=self.webhook_configs)


def get_trigger(location: Optional[str] = None,
                project: Optional[str] = None,
                trigger_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTriggerResult:
    """
    To get more information about Cloudbuild Trigger, see:

    * [API documentation](https://cloud.google.com/build/docs/api/reference/rest/v1/projects.triggers)
    * How-to Guides
        * [Official Documentation](https://cloud.google.com/build/docs/automating-builds/create-manage-triggers)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    name = gcp.cloudbuild.get_trigger(project="your-project-id",
        trigger_id=google_cloudbuild_trigger["filename-trigger"]["trigger_id"],
        location="location of trigger build")
    ```


    :param str location: The Cloud Build location for the trigger.
           
           - - -
    :param str project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
    :param str trigger_id: The unique identifier for the trigger..
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['triggerId'] = trigger_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:cloudbuild/getTrigger:getTrigger', __args__, opts=opts, typ=GetTriggerResult).value

    return AwaitableGetTriggerResult(
        approval_configs=pulumi.get(__ret__, 'approval_configs'),
        bitbucket_server_trigger_configs=pulumi.get(__ret__, 'bitbucket_server_trigger_configs'),
        builds=pulumi.get(__ret__, 'builds'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        disabled=pulumi.get(__ret__, 'disabled'),
        filename=pulumi.get(__ret__, 'filename'),
        filter=pulumi.get(__ret__, 'filter'),
        git_file_sources=pulumi.get(__ret__, 'git_file_sources'),
        githubs=pulumi.get(__ret__, 'githubs'),
        id=pulumi.get(__ret__, 'id'),
        ignored_files=pulumi.get(__ret__, 'ignored_files'),
        include_build_logs=pulumi.get(__ret__, 'include_build_logs'),
        included_files=pulumi.get(__ret__, 'included_files'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        project=pulumi.get(__ret__, 'project'),
        pubsub_configs=pulumi.get(__ret__, 'pubsub_configs'),
        repository_event_configs=pulumi.get(__ret__, 'repository_event_configs'),
        service_account=pulumi.get(__ret__, 'service_account'),
        source_to_builds=pulumi.get(__ret__, 'source_to_builds'),
        substitutions=pulumi.get(__ret__, 'substitutions'),
        tags=pulumi.get(__ret__, 'tags'),
        trigger_id=pulumi.get(__ret__, 'trigger_id'),
        trigger_templates=pulumi.get(__ret__, 'trigger_templates'),
        webhook_configs=pulumi.get(__ret__, 'webhook_configs'))


@_utilities.lift_output_func(get_trigger)
def get_trigger_output(location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       trigger_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTriggerResult]:
    """
    To get more information about Cloudbuild Trigger, see:

    * [API documentation](https://cloud.google.com/build/docs/api/reference/rest/v1/projects.triggers)
    * How-to Guides
        * [Official Documentation](https://cloud.google.com/build/docs/automating-builds/create-manage-triggers)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    name = gcp.cloudbuild.get_trigger(project="your-project-id",
        trigger_id=google_cloudbuild_trigger["filename-trigger"]["trigger_id"],
        location="location of trigger build")
    ```


    :param str location: The Cloud Build location for the trigger.
           
           - - -
    :param str project: The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
    :param str trigger_id: The unique identifier for the trigger..
    """
    ...
