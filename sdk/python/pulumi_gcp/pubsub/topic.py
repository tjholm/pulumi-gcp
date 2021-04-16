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

__all__ = ['TopicArgs', 'Topic']

@pulumi.input_type
class TopicArgs:
    def __init__(__self__, *,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 message_storage_policy: Optional[pulumi.Input['TopicMessageStoragePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Topic resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access
               to messages published on this topic. Your project's PubSub service account
               (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
               `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
               The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input['TopicMessageStoragePolicyArgs'] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where
               messages published to the topic may be stored. If not present, then no
               constraints are in effect.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if message_storage_policy is not None:
            pulumi.set(__self__, "message_storage_policy", message_storage_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the Cloud KMS CryptoKey to be used to protect access
        to messages published on this topic. Your project's PubSub service account
        (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
        `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to this Topic.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="messageStoragePolicy")
    def message_storage_policy(self) -> Optional[pulumi.Input['TopicMessageStoragePolicyArgs']]:
        """
        Policy constraining the set of Google Cloud Platform regions where
        messages published to the topic may be stored. If not present, then no
        constraints are in effect.
        Structure is documented below.
        """
        return pulumi.get(self, "message_storage_policy")

    @message_storage_policy.setter
    def message_storage_policy(self, value: Optional[pulumi.Input['TopicMessageStoragePolicyArgs']]):
        pulumi.set(self, "message_storage_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the topic.
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


@pulumi.input_type
class _TopicState:
    def __init__(__self__, *,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 message_storage_policy: Optional[pulumi.Input['TopicMessageStoragePolicyArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Topic resources.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access
               to messages published on this topic. Your project's PubSub service account
               (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
               `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
               The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input['TopicMessageStoragePolicyArgs'] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where
               messages published to the topic may be stored. If not present, then no
               constraints are in effect.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if kms_key_name is not None:
            pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if message_storage_policy is not None:
            pulumi.set(__self__, "message_storage_policy", message_storage_policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the Cloud KMS CryptoKey to be used to protect access
        to messages published on this topic. Your project's PubSub service account
        (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
        `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        """
        return pulumi.get(self, "kms_key_name")

    @kms_key_name.setter
    def kms_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of key/value label pairs to assign to this Topic.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="messageStoragePolicy")
    def message_storage_policy(self) -> Optional[pulumi.Input['TopicMessageStoragePolicyArgs']]:
        """
        Policy constraining the set of Google Cloud Platform regions where
        messages published to the topic may be stored. If not present, then no
        constraints are in effect.
        Structure is documented below.
        """
        return pulumi.get(self, "message_storage_policy")

    @message_storage_policy.setter
    def message_storage_policy(self, value: Optional[pulumi.Input['TopicMessageStoragePolicyArgs']]):
        pulumi.set(self, "message_storage_policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the topic.
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


class Topic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 message_storage_policy: Optional[pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A named resource to which messages are sent by publishers.

        To get more information about Topic, see:

        * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
        * How-to Guides
            * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)

        > **Note:** You can retrieve the email of the Google Managed Pub/Sub Service Account used for forwarding
        by using the `projects.ServiceIdentity` resource.

        ## Example Usage
        ### Pubsub Topic Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.pubsub.Topic("example", labels={
            "foo": "bar",
        })
        ```
        ### Pubsub Topic Cmek

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRing("keyRing", location="global")
        crypto_key = gcp.kms.CryptoKey("cryptoKey", key_ring=key_ring.id)
        example = gcp.pubsub.Topic("example", kms_key_name=crypto_key.id)
        ```
        ### Pubsub Topic Geo Restricted

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.pubsub.Topic("example", message_storage_policy=gcp.pubsub.TopicMessageStoragePolicyArgs(
            allowed_persistence_regions=["europe-west3"],
        ))
        ```

        ## Import

        Topic can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:pubsub/topic:Topic default projects/{{project}}/topics/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/topic:Topic default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/topic:Topic default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access
               to messages published on this topic. Your project's PubSub service account
               (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
               `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
               The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where
               messages published to the topic may be stored. If not present, then no
               constraints are in effect.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TopicArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A named resource to which messages are sent by publishers.

        To get more information about Topic, see:

        * [API documentation](https://cloud.google.com/pubsub/docs/reference/rest/v1/projects.topics)
        * How-to Guides
            * [Managing Topics](https://cloud.google.com/pubsub/docs/admin#managing_topics)

        > **Note:** You can retrieve the email of the Google Managed Pub/Sub Service Account used for forwarding
        by using the `projects.ServiceIdentity` resource.

        ## Example Usage
        ### Pubsub Topic Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.pubsub.Topic("example", labels={
            "foo": "bar",
        })
        ```
        ### Pubsub Topic Cmek

        ```python
        import pulumi
        import pulumi_gcp as gcp

        key_ring = gcp.kms.KeyRing("keyRing", location="global")
        crypto_key = gcp.kms.CryptoKey("cryptoKey", key_ring=key_ring.id)
        example = gcp.pubsub.Topic("example", kms_key_name=crypto_key.id)
        ```
        ### Pubsub Topic Geo Restricted

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.pubsub.Topic("example", message_storage_policy=gcp.pubsub.TopicMessageStoragePolicyArgs(
            allowed_persistence_regions=["europe-west3"],
        ))
        ```

        ## Import

        Topic can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:pubsub/topic:Topic default projects/{{project}}/topics/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/topic:Topic default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:pubsub/topic:Topic default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param TopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 kms_key_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 message_storage_policy: Optional[pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
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
            __props__ = TopicArgs.__new__(TopicArgs)

            __props__.__dict__["kms_key_name"] = kms_key_name
            __props__.__dict__["labels"] = labels
            __props__.__dict__["message_storage_policy"] = message_storage_policy
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
        super(Topic, __self__).__init__(
            'gcp:pubsub/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            kms_key_name: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            message_storage_policy: Optional[pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kms_key_name: The resource name of the Cloud KMS CryptoKey to be used to protect access
               to messages published on this topic. Your project's PubSub service account
               (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
               `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
               The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of key/value label pairs to assign to this Topic.
        :param pulumi.Input[pulumi.InputType['TopicMessageStoragePolicyArgs']] message_storage_policy: Policy constraining the set of Google Cloud Platform regions where
               messages published to the topic may be stored. If not present, then no
               constraints are in effect.
               Structure is documented below.
        :param pulumi.Input[str] name: Name of the topic.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TopicState.__new__(_TopicState)

        __props__.__dict__["kms_key_name"] = kms_key_name
        __props__.__dict__["labels"] = labels
        __props__.__dict__["message_storage_policy"] = message_storage_policy
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        return Topic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> pulumi.Output[Optional[str]]:
        """
        The resource name of the Cloud KMS CryptoKey to be used to protect access
        to messages published on this topic. Your project's PubSub service account
        (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have
        `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.
        The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*`
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A set of key/value label pairs to assign to this Topic.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="messageStoragePolicy")
    def message_storage_policy(self) -> pulumi.Output['outputs.TopicMessageStoragePolicy']:
        """
        Policy constraining the set of Google Cloud Platform regions where
        messages published to the topic may be stored. If not present, then no
        constraints are in effect.
        Structure is documented below.
        """
        return pulumi.get(self, "message_storage_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the topic.
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

