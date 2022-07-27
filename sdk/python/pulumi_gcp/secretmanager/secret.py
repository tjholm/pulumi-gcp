# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecretArgs', 'Secret']

@pulumi.input_type
class SecretArgs:
    def __init__(__self__, *,
                 replication: pulumi.Input['SecretReplicationArgs'],
                 secret_id: pulumi.Input[str],
                 expire_time: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 rotation: Optional[pulumi.Input['SecretRotationArgs']] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Secret resource.
        :param pulumi.Input['SecretReplicationArgs'] replication: The replication policy of the secret data attached to the Secret. It cannot be changed
               after the Secret has been created.
               Structure is documented below.
        :param pulumi.Input[str] secret_id: This must be unique within the project.
        :param pulumi.Input[str] expire_time: Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this Secret.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
               Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
               No more than 64 labels can be assigned to a given resource.
               An object containing a list of "key": value pairs. Example:
               { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['SecretRotationArgs'] rotation: The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]] topics: A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
               Structure is documented below.
        :param pulumi.Input[str] ttl: The TTL for the Secret.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        pulumi.set(__self__, "replication", replication)
        pulumi.set(__self__, "secret_id", secret_id)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if rotation is not None:
            pulumi.set(__self__, "rotation", rotation)
        if topics is not None:
            pulumi.set(__self__, "topics", topics)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def replication(self) -> pulumi.Input['SecretReplicationArgs']:
        """
        The replication policy of the secret data attached to the Secret. It cannot be changed
        after the Secret has been created.
        Structure is documented below.
        """
        return pulumi.get(self, "replication")

    @replication.setter
    def replication(self, value: pulumi.Input['SecretReplicationArgs']):
        pulumi.set(self, "replication", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Input[str]:
        """
        This must be unique within the project.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels assigned to this Secret.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
        and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
        Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
        and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
        No more than 64 labels can be assigned to a given resource.
        An object containing a list of "key": value pairs. Example:
        { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

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
    @pulumi.getter
    def rotation(self) -> Optional[pulumi.Input['SecretRotationArgs']]:
        """
        The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
        Structure is documented below.
        """
        return pulumi.get(self, "rotation")

    @rotation.setter
    def rotation(self, value: Optional[pulumi.Input['SecretRotationArgs']]):
        pulumi.set(self, "rotation", value)

    @property
    @pulumi.getter
    def topics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]]]:
        """
        A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
        Structure is documented below.
        """
        return pulumi.get(self, "topics")

    @topics.setter
    def topics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]]]):
        pulumi.set(self, "topics", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The TTL for the Secret.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class _SecretState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input['SecretReplicationArgs']] = None,
                 rotation: Optional[pulumi.Input['SecretRotationArgs']] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Secret resources.
        :param pulumi.Input[str] create_time: The time at which the Secret was created.
        :param pulumi.Input[str] expire_time: Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this Secret.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
               Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
               No more than 64 labels can be assigned to a given resource.
               An object containing a list of "key": value pairs. Example:
               { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.
               For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['SecretReplicationArgs'] replication: The replication policy of the secret data attached to the Secret. It cannot be changed
               after the Secret has been created.
               Structure is documented below.
        :param pulumi.Input['SecretRotationArgs'] rotation: The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
               Structure is documented below.
        :param pulumi.Input[str] secret_id: This must be unique within the project.
        :param pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]] topics: A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
               Structure is documented below.
        :param pulumi.Input[str] ttl: The TTL for the Secret.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if replication is not None:
            pulumi.set(__self__, "replication", replication)
        if rotation is not None:
            pulumi.set(__self__, "rotation", rotation)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if topics is not None:
            pulumi.set(__self__, "topics", topics)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the Secret was created.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The labels assigned to this Secret.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
        and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
        Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
        and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
        No more than 64 labels can be assigned to a given resource.
        An object containing a list of "key": value pairs. Example:
        { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.
        For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
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
    @pulumi.getter
    def replication(self) -> Optional[pulumi.Input['SecretReplicationArgs']]:
        """
        The replication policy of the secret data attached to the Secret. It cannot be changed
        after the Secret has been created.
        Structure is documented below.
        """
        return pulumi.get(self, "replication")

    @replication.setter
    def replication(self, value: Optional[pulumi.Input['SecretReplicationArgs']]):
        pulumi.set(self, "replication", value)

    @property
    @pulumi.getter
    def rotation(self) -> Optional[pulumi.Input['SecretRotationArgs']]:
        """
        The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
        Structure is documented below.
        """
        return pulumi.get(self, "rotation")

    @rotation.setter
    def rotation(self, value: Optional[pulumi.Input['SecretRotationArgs']]):
        pulumi.set(self, "rotation", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        This must be unique within the project.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter
    def topics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]]]:
        """
        A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
        Structure is documented below.
        """
        return pulumi.get(self, "topics")

    @topics.setter
    def topics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretTopicArgs']]]]):
        pulumi.set(self, "topics", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The TTL for the Secret.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


class Secret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input[pulumi.InputType['SecretReplicationArgs']]] = None,
                 rotation: Optional[pulumi.Input[pulumi.InputType['SecretRotationArgs']]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretTopicArgs']]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A Secret is a logical secret whose value and versions can be accessed.

        To get more information about Secret, see:

        * [API documentation](https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets)

        ## Example Usage
        ### Secret Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        secret_basic = gcp.secretmanager.Secret("secret-basic",
            labels={
                "label": "my-label",
            },
            replication=gcp.secretmanager.SecretReplicationArgs(
                user_managed=gcp.secretmanager.SecretReplicationUserManagedArgs(
                    replicas=[
                        gcp.secretmanager.SecretReplicationUserManagedReplicaArgs(
                            location="us-central1",
                        ),
                        gcp.secretmanager.SecretReplicationUserManagedReplicaArgs(
                            location="us-east1",
                        ),
                    ],
                ),
            ),
            secret_id="secret")
        ```

        ## Import

        Secret can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:secretmanager/secret:Secret default projects/{{project}}/secrets/{{secret_id}}
        ```

        ```sh
         $ pulumi import gcp:secretmanager/secret:Secret default {{project}}/{{secret_id}}
        ```

        ```sh
         $ pulumi import gcp:secretmanager/secret:Secret default {{secret_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expire_time: Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this Secret.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
               Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
               No more than 64 labels can be assigned to a given resource.
               An object containing a list of "key": value pairs. Example:
               { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SecretReplicationArgs']] replication: The replication policy of the secret data attached to the Secret. It cannot be changed
               after the Secret has been created.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['SecretRotationArgs']] rotation: The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
               Structure is documented below.
        :param pulumi.Input[str] secret_id: This must be unique within the project.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretTopicArgs']]]] topics: A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
               Structure is documented below.
        :param pulumi.Input[str] ttl: The TTL for the Secret.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Secret is a logical secret whose value and versions can be accessed.

        To get more information about Secret, see:

        * [API documentation](https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets)

        ## Example Usage
        ### Secret Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        secret_basic = gcp.secretmanager.Secret("secret-basic",
            labels={
                "label": "my-label",
            },
            replication=gcp.secretmanager.SecretReplicationArgs(
                user_managed=gcp.secretmanager.SecretReplicationUserManagedArgs(
                    replicas=[
                        gcp.secretmanager.SecretReplicationUserManagedReplicaArgs(
                            location="us-central1",
                        ),
                        gcp.secretmanager.SecretReplicationUserManagedReplicaArgs(
                            location="us-east1",
                        ),
                    ],
                ),
            ),
            secret_id="secret")
        ```

        ## Import

        Secret can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:secretmanager/secret:Secret default projects/{{project}}/secrets/{{secret_id}}
        ```

        ```sh
         $ pulumi import gcp:secretmanager/secret:Secret default {{project}}/{{secret_id}}
        ```

        ```sh
         $ pulumi import gcp:secretmanager/secret:Secret default {{secret_id}}
        ```

        :param str resource_name: The name of the resource.
        :param SecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 replication: Optional[pulumi.Input[pulumi.InputType['SecretReplicationArgs']]] = None,
                 rotation: Optional[pulumi.Input[pulumi.InputType['SecretRotationArgs']]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 topics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretTopicArgs']]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretArgs.__new__(SecretArgs)

            __props__.__dict__["expire_time"] = expire_time
            __props__.__dict__["labels"] = labels
            __props__.__dict__["project"] = project
            if replication is None and not opts.urn:
                raise TypeError("Missing required property 'replication'")
            __props__.__dict__["replication"] = replication
            __props__.__dict__["rotation"] = rotation
            if secret_id is None and not opts.urn:
                raise TypeError("Missing required property 'secret_id'")
            __props__.__dict__["secret_id"] = secret_id
            __props__.__dict__["topics"] = topics
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
        super(Secret, __self__).__init__(
            'gcp:secretmanager/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            replication: Optional[pulumi.Input[pulumi.InputType['SecretReplicationArgs']]] = None,
            rotation: Optional[pulumi.Input[pulumi.InputType['SecretRotationArgs']]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            topics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretTopicArgs']]]]] = None,
            ttl: Optional[pulumi.Input[str]] = None) -> 'Secret':
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: The time at which the Secret was created.
        :param pulumi.Input[str] expire_time: Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The labels assigned to this Secret.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
               Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
               and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
               No more than 64 labels can be assigned to a given resource.
               An object containing a list of "key": value pairs. Example:
               { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        :param pulumi.Input[str] name: The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.
               For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SecretReplicationArgs']] replication: The replication policy of the secret data attached to the Secret. It cannot be changed
               after the Secret has been created.
               Structure is documented below.
        :param pulumi.Input[pulumi.InputType['SecretRotationArgs']] rotation: The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
               Structure is documented below.
        :param pulumi.Input[str] secret_id: This must be unique within the project.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretTopicArgs']]]] topics: A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
               Structure is documented below.
        :param pulumi.Input[str] ttl: The TTL for the Secret.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretState.__new__(_SecretState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["replication"] = replication
        __props__.__dict__["rotation"] = rotation
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["topics"] = topics
        __props__.__dict__["ttl"] = ttl
        return Secret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time at which the Secret was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        Timestamp in UTC when the Secret is scheduled to expire. This is always provided on output, regardless of what was sent on input.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The labels assigned to this Secret.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
        and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}
        Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
        and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}
        No more than 64 labels can be assigned to a given resource.
        An object containing a list of "key": value pairs. Example:
        { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*/topics/*.
        For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
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
    @pulumi.getter
    def replication(self) -> pulumi.Output['outputs.SecretReplication']:
        """
        The replication policy of the secret data attached to the Secret. It cannot be changed
        after the Secret has been created.
        Structure is documented below.
        """
        return pulumi.get(self, "replication")

    @property
    @pulumi.getter
    def rotation(self) -> pulumi.Output[Optional['outputs.SecretRotation']]:
        """
        The rotation time and period for a Secret. At `next_rotation_time`, Secret Manager will send a Pub/Sub notification to the topics configured on the Secret. `topics` must be set to configure rotation.
        Structure is documented below.
        """
        return pulumi.get(self, "rotation")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        This must be unique within the project.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter
    def topics(self) -> pulumi.Output[Optional[Sequence['outputs.SecretTopic']]]:
        """
        A list of up to 10 Pub/Sub topics to which messages are published when control plane operations are called on the secret or its versions.
        Structure is documented below.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The TTL for the Secret.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "ttl")

