# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'RepositoryCleanupPolicyArgs',
    'RepositoryCleanupPolicyConditionArgs',
    'RepositoryCleanupPolicyMostRecentVersionsArgs',
    'RepositoryDockerConfigArgs',
    'RepositoryIamBindingConditionArgs',
    'RepositoryIamMemberConditionArgs',
    'RepositoryMavenConfigArgs',
    'RepositoryRemoteRepositoryConfigArgs',
    'RepositoryRemoteRepositoryConfigAptRepositoryArgs',
    'RepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryArgs',
    'RepositoryRemoteRepositoryConfigDockerRepositoryArgs',
    'RepositoryRemoteRepositoryConfigMavenRepositoryArgs',
    'RepositoryRemoteRepositoryConfigNpmRepositoryArgs',
    'RepositoryRemoteRepositoryConfigPythonRepositoryArgs',
    'RepositoryRemoteRepositoryConfigUpstreamCredentialsArgs',
    'RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsArgs',
    'RepositoryRemoteRepositoryConfigYumRepositoryArgs',
    'RepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryArgs',
    'RepositoryVirtualRepositoryConfigArgs',
    'RepositoryVirtualRepositoryConfigUpstreamPolicyArgs',
]

@pulumi.input_type
class RepositoryCleanupPolicyArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 condition: Optional[pulumi.Input['RepositoryCleanupPolicyConditionArgs']] = None,
                 most_recent_versions: Optional[pulumi.Input['RepositoryCleanupPolicyMostRecentVersionsArgs']] = None):
        """
        :param pulumi.Input[str] id: The identifier for this object. Format specified above.
        :param pulumi.Input[str] action: Policy action.
               Possible values are: `DELETE`, `KEEP`.
        :param pulumi.Input['RepositoryCleanupPolicyConditionArgs'] condition: Policy condition for matching versions.
               Structure is documented below.
        :param pulumi.Input['RepositoryCleanupPolicyMostRecentVersionsArgs'] most_recent_versions: Policy condition for retaining a minimum number of versions. May only be
               specified with a Keep action.
               Structure is documented below.
        """
        pulumi.set(__self__, "id", id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if most_recent_versions is not None:
            pulumi.set(__self__, "most_recent_versions", most_recent_versions)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The identifier for this object. Format specified above.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Policy action.
        Possible values are: `DELETE`, `KEEP`.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input['RepositoryCleanupPolicyConditionArgs']]:
        """
        Policy condition for matching versions.
        Structure is documented below.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input['RepositoryCleanupPolicyConditionArgs']]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="mostRecentVersions")
    def most_recent_versions(self) -> Optional[pulumi.Input['RepositoryCleanupPolicyMostRecentVersionsArgs']]:
        """
        Policy condition for retaining a minimum number of versions. May only be
        specified with a Keep action.
        Structure is documented below.
        """
        return pulumi.get(self, "most_recent_versions")

    @most_recent_versions.setter
    def most_recent_versions(self, value: Optional[pulumi.Input['RepositoryCleanupPolicyMostRecentVersionsArgs']]):
        pulumi.set(self, "most_recent_versions", value)


@pulumi.input_type
class RepositoryCleanupPolicyConditionArgs:
    def __init__(__self__, *,
                 newer_than: Optional[pulumi.Input[str]] = None,
                 older_than: Optional[pulumi.Input[str]] = None,
                 package_name_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tag_state: Optional[pulumi.Input[str]] = None,
                 version_name_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] newer_than: Match versions newer than a duration.
        :param pulumi.Input[str] older_than: Match versions older than a duration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] package_name_prefixes: Match versions by package prefix. Applied on any prefix match.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tag_prefixes: Match versions by tag prefix. Applied on any prefix match.
        :param pulumi.Input[str] tag_state: Match versions by tag status.
               Default value is `ANY`.
               Possible values are: `TAGGED`, `UNTAGGED`, `ANY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] version_name_prefixes: Match versions by version name prefix. Applied on any prefix match.
        """
        if newer_than is not None:
            pulumi.set(__self__, "newer_than", newer_than)
        if older_than is not None:
            pulumi.set(__self__, "older_than", older_than)
        if package_name_prefixes is not None:
            pulumi.set(__self__, "package_name_prefixes", package_name_prefixes)
        if tag_prefixes is not None:
            pulumi.set(__self__, "tag_prefixes", tag_prefixes)
        if tag_state is not None:
            pulumi.set(__self__, "tag_state", tag_state)
        if version_name_prefixes is not None:
            pulumi.set(__self__, "version_name_prefixes", version_name_prefixes)

    @property
    @pulumi.getter(name="newerThan")
    def newer_than(self) -> Optional[pulumi.Input[str]]:
        """
        Match versions newer than a duration.
        """
        return pulumi.get(self, "newer_than")

    @newer_than.setter
    def newer_than(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "newer_than", value)

    @property
    @pulumi.getter(name="olderThan")
    def older_than(self) -> Optional[pulumi.Input[str]]:
        """
        Match versions older than a duration.
        """
        return pulumi.get(self, "older_than")

    @older_than.setter
    def older_than(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "older_than", value)

    @property
    @pulumi.getter(name="packageNamePrefixes")
    def package_name_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Match versions by package prefix. Applied on any prefix match.
        """
        return pulumi.get(self, "package_name_prefixes")

    @package_name_prefixes.setter
    def package_name_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "package_name_prefixes", value)

    @property
    @pulumi.getter(name="tagPrefixes")
    def tag_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Match versions by tag prefix. Applied on any prefix match.
        """
        return pulumi.get(self, "tag_prefixes")

    @tag_prefixes.setter
    def tag_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tag_prefixes", value)

    @property
    @pulumi.getter(name="tagState")
    def tag_state(self) -> Optional[pulumi.Input[str]]:
        """
        Match versions by tag status.
        Default value is `ANY`.
        Possible values are: `TAGGED`, `UNTAGGED`, `ANY`.
        """
        return pulumi.get(self, "tag_state")

    @tag_state.setter
    def tag_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tag_state", value)

    @property
    @pulumi.getter(name="versionNamePrefixes")
    def version_name_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Match versions by version name prefix. Applied on any prefix match.
        """
        return pulumi.get(self, "version_name_prefixes")

    @version_name_prefixes.setter
    def version_name_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "version_name_prefixes", value)


@pulumi.input_type
class RepositoryCleanupPolicyMostRecentVersionsArgs:
    def __init__(__self__, *,
                 keep_count: Optional[pulumi.Input[int]] = None,
                 package_name_prefixes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[int] keep_count: Minimum number of versions to keep.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] package_name_prefixes: Match versions by package prefix. Applied on any prefix match.
        """
        if keep_count is not None:
            pulumi.set(__self__, "keep_count", keep_count)
        if package_name_prefixes is not None:
            pulumi.set(__self__, "package_name_prefixes", package_name_prefixes)

    @property
    @pulumi.getter(name="keepCount")
    def keep_count(self) -> Optional[pulumi.Input[int]]:
        """
        Minimum number of versions to keep.
        """
        return pulumi.get(self, "keep_count")

    @keep_count.setter
    def keep_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "keep_count", value)

    @property
    @pulumi.getter(name="packageNamePrefixes")
    def package_name_prefixes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Match versions by package prefix. Applied on any prefix match.
        """
        return pulumi.get(self, "package_name_prefixes")

    @package_name_prefixes.setter
    def package_name_prefixes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "package_name_prefixes", value)


@pulumi.input_type
class RepositoryDockerConfigArgs:
    def __init__(__self__, *,
                 immutable_tags: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] immutable_tags: The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
        """
        if immutable_tags is not None:
            pulumi.set(__self__, "immutable_tags", immutable_tags)

    @property
    @pulumi.getter(name="immutableTags")
    def immutable_tags(self) -> Optional[pulumi.Input[bool]]:
        """
        The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created.
        """
        return pulumi.get(self, "immutable_tags")

    @immutable_tags.setter
    def immutable_tags(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "immutable_tags", value)


@pulumi.input_type
class RepositoryIamBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class RepositoryIamMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class RepositoryMavenConfigArgs:
    def __init__(__self__, *,
                 allow_snapshot_overwrites: Optional[pulumi.Input[bool]] = None,
                 version_policy: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] allow_snapshot_overwrites: The repository with this flag will allow publishing the same
               snapshot versions.
        :param pulumi.Input[str] version_policy: Version policy defines the versions that the registry will accept.
               Default value is `VERSION_POLICY_UNSPECIFIED`.
               Possible values are: `VERSION_POLICY_UNSPECIFIED`, `RELEASE`, `SNAPSHOT`.
        """
        if allow_snapshot_overwrites is not None:
            pulumi.set(__self__, "allow_snapshot_overwrites", allow_snapshot_overwrites)
        if version_policy is not None:
            pulumi.set(__self__, "version_policy", version_policy)

    @property
    @pulumi.getter(name="allowSnapshotOverwrites")
    def allow_snapshot_overwrites(self) -> Optional[pulumi.Input[bool]]:
        """
        The repository with this flag will allow publishing the same
        snapshot versions.
        """
        return pulumi.get(self, "allow_snapshot_overwrites")

    @allow_snapshot_overwrites.setter
    def allow_snapshot_overwrites(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_snapshot_overwrites", value)

    @property
    @pulumi.getter(name="versionPolicy")
    def version_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Version policy defines the versions that the registry will accept.
        Default value is `VERSION_POLICY_UNSPECIFIED`.
        Possible values are: `VERSION_POLICY_UNSPECIFIED`, `RELEASE`, `SNAPSHOT`.
        """
        return pulumi.get(self, "version_policy")

    @version_policy.setter
    def version_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_policy", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigArgs:
    def __init__(__self__, *,
                 apt_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 docker_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigDockerRepositoryArgs']] = None,
                 maven_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigMavenRepositoryArgs']] = None,
                 npm_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigNpmRepositoryArgs']] = None,
                 python_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigPythonRepositoryArgs']] = None,
                 upstream_credentials: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsArgs']] = None,
                 yum_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryArgs']] = None):
        """
        :param pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryArgs'] apt_repository: Specific settings for an Apt remote repository.
               Structure is documented below.
        :param pulumi.Input[str] description: The description of the remote source.
        :param pulumi.Input['RepositoryRemoteRepositoryConfigDockerRepositoryArgs'] docker_repository: Specific settings for a Docker remote repository.
               Structure is documented below.
        :param pulumi.Input['RepositoryRemoteRepositoryConfigMavenRepositoryArgs'] maven_repository: Specific settings for a Maven remote repository.
               Structure is documented below.
        :param pulumi.Input['RepositoryRemoteRepositoryConfigNpmRepositoryArgs'] npm_repository: Specific settings for an Npm remote repository.
               Structure is documented below.
        :param pulumi.Input['RepositoryRemoteRepositoryConfigPythonRepositoryArgs'] python_repository: Specific settings for a Python remote repository.
               Structure is documented below.
        :param pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsArgs'] upstream_credentials: The credentials used to access the remote repository.
               Structure is documented below.
        :param pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryArgs'] yum_repository: Specific settings for an Yum remote repository.
               Structure is documented below.
        """
        if apt_repository is not None:
            pulumi.set(__self__, "apt_repository", apt_repository)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if docker_repository is not None:
            pulumi.set(__self__, "docker_repository", docker_repository)
        if maven_repository is not None:
            pulumi.set(__self__, "maven_repository", maven_repository)
        if npm_repository is not None:
            pulumi.set(__self__, "npm_repository", npm_repository)
        if python_repository is not None:
            pulumi.set(__self__, "python_repository", python_repository)
        if upstream_credentials is not None:
            pulumi.set(__self__, "upstream_credentials", upstream_credentials)
        if yum_repository is not None:
            pulumi.set(__self__, "yum_repository", yum_repository)

    @property
    @pulumi.getter(name="aptRepository")
    def apt_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryArgs']]:
        """
        Specific settings for an Apt remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "apt_repository")

    @apt_repository.setter
    def apt_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryArgs']]):
        pulumi.set(self, "apt_repository", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the remote source.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dockerRepository")
    def docker_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigDockerRepositoryArgs']]:
        """
        Specific settings for a Docker remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "docker_repository")

    @docker_repository.setter
    def docker_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigDockerRepositoryArgs']]):
        pulumi.set(self, "docker_repository", value)

    @property
    @pulumi.getter(name="mavenRepository")
    def maven_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigMavenRepositoryArgs']]:
        """
        Specific settings for a Maven remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "maven_repository")

    @maven_repository.setter
    def maven_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigMavenRepositoryArgs']]):
        pulumi.set(self, "maven_repository", value)

    @property
    @pulumi.getter(name="npmRepository")
    def npm_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigNpmRepositoryArgs']]:
        """
        Specific settings for an Npm remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "npm_repository")

    @npm_repository.setter
    def npm_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigNpmRepositoryArgs']]):
        pulumi.set(self, "npm_repository", value)

    @property
    @pulumi.getter(name="pythonRepository")
    def python_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigPythonRepositoryArgs']]:
        """
        Specific settings for a Python remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "python_repository")

    @python_repository.setter
    def python_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigPythonRepositoryArgs']]):
        pulumi.set(self, "python_repository", value)

    @property
    @pulumi.getter(name="upstreamCredentials")
    def upstream_credentials(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsArgs']]:
        """
        The credentials used to access the remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "upstream_credentials")

    @upstream_credentials.setter
    def upstream_credentials(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsArgs']]):
        pulumi.set(self, "upstream_credentials", value)

    @property
    @pulumi.getter(name="yumRepository")
    def yum_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryArgs']]:
        """
        Specific settings for an Yum remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "yum_repository")

    @yum_repository.setter
    def yum_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryArgs']]):
        pulumi.set(self, "yum_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigAptRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryArgs']] = None):
        """
        :param pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryArgs'] public_repository: One of the publicly available Apt repositories supported by Artifact Registry.
               Structure is documented below.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryArgs']]:
        """
        One of the publicly available Apt repositories supported by Artifact Registry.
        Structure is documented below.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryArgs']]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigAptRepositoryPublicRepositoryArgs:
    def __init__(__self__, *,
                 repository_base: pulumi.Input[str],
                 repository_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] repository_base: A common public repository base for Yum.
               Possible values are: `CENTOS`, `CENTOS_DEBUG`, `CENTOS_VAULT`, `CENTOS_STREAM`, `ROCKY`, `EPEL`.
        :param pulumi.Input[str] repository_path: Specific repository from the base, e.g. `"centos/8-stream/BaseOS/x86_64/os"`
        """
        pulumi.set(__self__, "repository_base", repository_base)
        pulumi.set(__self__, "repository_path", repository_path)

    @property
    @pulumi.getter(name="repositoryBase")
    def repository_base(self) -> pulumi.Input[str]:
        """
        A common public repository base for Yum.
        Possible values are: `CENTOS`, `CENTOS_DEBUG`, `CENTOS_VAULT`, `CENTOS_STREAM`, `ROCKY`, `EPEL`.
        """
        return pulumi.get(self, "repository_base")

    @repository_base.setter
    def repository_base(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_base", value)

    @property
    @pulumi.getter(name="repositoryPath")
    def repository_path(self) -> pulumi.Input[str]:
        """
        Specific repository from the base, e.g. `"centos/8-stream/BaseOS/x86_64/os"`
        """
        return pulumi.get(self, "repository_path")

    @repository_path.setter
    def repository_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_path", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigDockerRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] public_repository: Address of the remote repository.
               Default value is `DOCKER_HUB`.
               Possible values are: `DOCKER_HUB`.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the remote repository.
        Default value is `DOCKER_HUB`.
        Possible values are: `DOCKER_HUB`.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigMavenRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] public_repository: Address of the remote repository.
               Default value is `MAVEN_CENTRAL`.
               Possible values are: `MAVEN_CENTRAL`.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the remote repository.
        Default value is `MAVEN_CENTRAL`.
        Possible values are: `MAVEN_CENTRAL`.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigNpmRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] public_repository: Address of the remote repository.
               Default value is `NPMJS`.
               Possible values are: `NPMJS`.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the remote repository.
        Default value is `NPMJS`.
        Possible values are: `NPMJS`.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigPythonRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] public_repository: Address of the remote repository.
               Default value is `PYPI`.
               Possible values are: `PYPI`.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input[str]]:
        """
        Address of the remote repository.
        Default value is `PYPI`.
        Possible values are: `PYPI`.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigUpstreamCredentialsArgs:
    def __init__(__self__, *,
                 username_password_credentials: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsArgs']] = None):
        """
        :param pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsArgs'] username_password_credentials: Use username and password to access the remote repository.
               Structure is documented below.
        """
        if username_password_credentials is not None:
            pulumi.set(__self__, "username_password_credentials", username_password_credentials)

    @property
    @pulumi.getter(name="usernamePasswordCredentials")
    def username_password_credentials(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsArgs']]:
        """
        Use username and password to access the remote repository.
        Structure is documented below.
        """
        return pulumi.get(self, "username_password_credentials")

    @username_password_credentials.setter
    def username_password_credentials(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsArgs']]):
        pulumi.set(self, "username_password_credentials", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigUpstreamCredentialsUsernamePasswordCredentialsArgs:
    def __init__(__self__, *,
                 password_secret_version: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] password_secret_version: The Secret Manager key version that holds the password to access the
               remote repository. Must be in the format of
               `projects/{project}/secrets/{secret}/versions/{version}`.
        :param pulumi.Input[str] username: The username to access the remote repository.
        """
        if password_secret_version is not None:
            pulumi.set(__self__, "password_secret_version", password_secret_version)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="passwordSecretVersion")
    def password_secret_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Secret Manager key version that holds the password to access the
        remote repository. Must be in the format of
        `projects/{project}/secrets/{secret}/versions/{version}`.
        """
        return pulumi.get(self, "password_secret_version")

    @password_secret_version.setter
    def password_secret_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_secret_version", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username to access the remote repository.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigYumRepositoryArgs:
    def __init__(__self__, *,
                 public_repository: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryArgs']] = None):
        """
        :param pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryArgs'] public_repository: One of the publicly available Yum repositories supported by Artifact Registry.
               Structure is documented below.
        """
        if public_repository is not None:
            pulumi.set(__self__, "public_repository", public_repository)

    @property
    @pulumi.getter(name="publicRepository")
    def public_repository(self) -> Optional[pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryArgs']]:
        """
        One of the publicly available Yum repositories supported by Artifact Registry.
        Structure is documented below.
        """
        return pulumi.get(self, "public_repository")

    @public_repository.setter
    def public_repository(self, value: Optional[pulumi.Input['RepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryArgs']]):
        pulumi.set(self, "public_repository", value)


@pulumi.input_type
class RepositoryRemoteRepositoryConfigYumRepositoryPublicRepositoryArgs:
    def __init__(__self__, *,
                 repository_base: pulumi.Input[str],
                 repository_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] repository_base: A common public repository base for Yum.
               Possible values are: `CENTOS`, `CENTOS_DEBUG`, `CENTOS_VAULT`, `CENTOS_STREAM`, `ROCKY`, `EPEL`.
        :param pulumi.Input[str] repository_path: Specific repository from the base, e.g. `"centos/8-stream/BaseOS/x86_64/os"`
        """
        pulumi.set(__self__, "repository_base", repository_base)
        pulumi.set(__self__, "repository_path", repository_path)

    @property
    @pulumi.getter(name="repositoryBase")
    def repository_base(self) -> pulumi.Input[str]:
        """
        A common public repository base for Yum.
        Possible values are: `CENTOS`, `CENTOS_DEBUG`, `CENTOS_VAULT`, `CENTOS_STREAM`, `ROCKY`, `EPEL`.
        """
        return pulumi.get(self, "repository_base")

    @repository_base.setter
    def repository_base(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_base", value)

    @property
    @pulumi.getter(name="repositoryPath")
    def repository_path(self) -> pulumi.Input[str]:
        """
        Specific repository from the base, e.g. `"centos/8-stream/BaseOS/x86_64/os"`
        """
        return pulumi.get(self, "repository_path")

    @repository_path.setter
    def repository_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "repository_path", value)


@pulumi.input_type
class RepositoryVirtualRepositoryConfigArgs:
    def __init__(__self__, *,
                 upstream_policies: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryVirtualRepositoryConfigUpstreamPolicyArgs']]]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['RepositoryVirtualRepositoryConfigUpstreamPolicyArgs']]] upstream_policies: Policies that configure the upstream artifacts distributed by the Virtual
               Repository. Upstream policies cannot be set on a standard repository.
               Structure is documented below.
        """
        if upstream_policies is not None:
            pulumi.set(__self__, "upstream_policies", upstream_policies)

    @property
    @pulumi.getter(name="upstreamPolicies")
    def upstream_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryVirtualRepositoryConfigUpstreamPolicyArgs']]]]:
        """
        Policies that configure the upstream artifacts distributed by the Virtual
        Repository. Upstream policies cannot be set on a standard repository.
        Structure is documented below.
        """
        return pulumi.get(self, "upstream_policies")

    @upstream_policies.setter
    def upstream_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RepositoryVirtualRepositoryConfigUpstreamPolicyArgs']]]]):
        pulumi.set(self, "upstream_policies", value)


@pulumi.input_type
class RepositoryVirtualRepositoryConfigUpstreamPolicyArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 repository: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: The user-provided ID of the upstream policy.
        :param pulumi.Input[int] priority: Entries with a greater priority value take precedence in the pull order.
        :param pulumi.Input[str] repository: A reference to the repository resource, for example:
               "projects/p1/locations/us-central1/repository/repo1".
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The user-provided ID of the upstream policy.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Entries with a greater priority value take precedence in the pull order.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        A reference to the repository resource, for example:
        "projects/p1/locations/us-central1/repository/repo1".
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)


