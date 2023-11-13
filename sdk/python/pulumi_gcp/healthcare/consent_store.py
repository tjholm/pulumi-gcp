# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ConsentStoreArgs', 'ConsentStore']

@pulumi.input_type
class ConsentStoreArgs:
    def __init__(__self__, *,
                 dataset: pulumi.Input[str],
                 default_consent_ttl: Optional[pulumi.Input[str]] = None,
                 enable_consent_create_on_update: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConsentStore resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               
               
               - - -
        :param pulumi.Input[str] default_consent_ttl: Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        :param pulumi.Input[bool] enable_consent_create_on_update: If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize Consent stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of this ConsentStore, for example:
               "consent1"
        """
        pulumi.set(__self__, "dataset", dataset)
        if default_consent_ttl is not None:
            pulumi.set(__self__, "default_consent_ttl", default_consent_ttl)
        if enable_consent_create_on_update is not None:
            pulumi.set(__self__, "enable_consent_create_on_update", enable_consent_create_on_update)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Input[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'


        - - -
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: pulumi.Input[str]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter(name="defaultConsentTtl")
    def default_consent_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "default_consent_ttl")

    @default_consent_ttl.setter
    def default_consent_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_consent_ttl", value)

    @property
    @pulumi.getter(name="enableConsentCreateOnUpdate")
    def enable_consent_create_on_update(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        """
        return pulumi.get(self, "enable_consent_create_on_update")

    @enable_consent_create_on_update.setter
    def enable_consent_create_on_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_consent_create_on_update", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key-value pairs used to organize Consent stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this ConsentStore, for example:
        "consent1"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ConsentStoreState:
    def __init__(__self__, *,
                 dataset: Optional[pulumi.Input[str]] = None,
                 default_consent_ttl: Optional[pulumi.Input[str]] = None,
                 effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 enable_consent_create_on_update: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering ConsentStore resources.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               
               
               - - -
        :param pulumi.Input[str] default_consent_ttl: Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[bool] enable_consent_create_on_update: If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize Consent stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of this ConsentStore, for example:
               "consent1"
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        """
        if dataset is not None:
            pulumi.set(__self__, "dataset", dataset)
        if default_consent_ttl is not None:
            pulumi.set(__self__, "default_consent_ttl", default_consent_ttl)
        if effective_labels is not None:
            pulumi.set(__self__, "effective_labels", effective_labels)
        if enable_consent_create_on_update is not None:
            pulumi.set(__self__, "enable_consent_create_on_update", enable_consent_create_on_update)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pulumi_labels is not None:
            pulumi.set(__self__, "pulumi_labels", pulumi_labels)

    @property
    @pulumi.getter
    def dataset(self) -> Optional[pulumi.Input[str]]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'


        - - -
        """
        return pulumi.get(self, "dataset")

    @dataset.setter
    def dataset(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dataset", value)

    @property
    @pulumi.getter(name="defaultConsentTtl")
    def default_consent_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "default_consent_ttl")

    @default_consent_ttl.setter
    def default_consent_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_consent_ttl", value)

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @effective_labels.setter
    def effective_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "effective_labels", value)

    @property
    @pulumi.getter(name="enableConsentCreateOnUpdate")
    def enable_consent_create_on_update(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        """
        return pulumi.get(self, "enable_consent_create_on_update")

    @enable_consent_create_on_update.setter
    def enable_consent_create_on_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_consent_create_on_update", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-supplied key-value pairs used to organize Consent stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this ConsentStore, for example:
        "consent1"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @pulumi_labels.setter
    def pulumi_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "pulumi_labels", value)


class ConsentStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 default_consent_ttl: Optional[pulumi.Input[str]] = None,
                 enable_consent_create_on_update: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The Consent Management API is a tool for tracking user consents and the documentation associated with the consents.

        To get more information about ConsentStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.consentStores)
        * How-to Guides
            * [Creating a Consent store](https://cloud.google.com/healthcare/docs/how-tos/consent)

        ## Example Usage
        ### Healthcare Consent Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        my_consent = gcp.healthcare.ConsentStore("my-consent", dataset=dataset.id)
        ```
        ### Healthcare Consent Store Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        my_consent = gcp.healthcare.ConsentStore("my-consent",
            dataset=dataset.id,
            enable_consent_create_on_update=True,
            default_consent_ttl="90000s",
            labels={
                "label1": "labelvalue1",
            })
        ```
        ### Healthcare Consent Store Iam

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        my_consent = gcp.healthcare.ConsentStore("my-consent", dataset=dataset.id)
        test_account = gcp.serviceaccount.Account("test-account",
            account_id="my-account",
            display_name="Test Service Account")
        test_iam = gcp.healthcare.ConsentStoreIamMember("test-iam",
            dataset=dataset.id,
            consent_store_id=my_consent.name,
            role="roles/editor",
            member=test_account.email.apply(lambda email: f"serviceAccount:{email}"))
        ```

        ## Import

        ConsentStore can be imported using any of these accepted formats:

        ```sh
         $ pulumi import gcp:healthcare/consentStore:ConsentStore default {{dataset}}/consentStores/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               
               
               - - -
        :param pulumi.Input[str] default_consent_ttl: Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        :param pulumi.Input[bool] enable_consent_create_on_update: If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize Consent stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of this ConsentStore, for example:
               "consent1"
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConsentStoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The Consent Management API is a tool for tracking user consents and the documentation associated with the consents.

        To get more information about ConsentStore, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.consentStores)
        * How-to Guides
            * [Creating a Consent store](https://cloud.google.com/healthcare/docs/how-tos/consent)

        ## Example Usage
        ### Healthcare Consent Store Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        my_consent = gcp.healthcare.ConsentStore("my-consent", dataset=dataset.id)
        ```
        ### Healthcare Consent Store Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        my_consent = gcp.healthcare.ConsentStore("my-consent",
            dataset=dataset.id,
            enable_consent_create_on_update=True,
            default_consent_ttl="90000s",
            labels={
                "label1": "labelvalue1",
            })
        ```
        ### Healthcare Consent Store Iam

        ```python
        import pulumi
        import pulumi_gcp as gcp

        dataset = gcp.healthcare.Dataset("dataset", location="us-central1")
        my_consent = gcp.healthcare.ConsentStore("my-consent", dataset=dataset.id)
        test_account = gcp.serviceaccount.Account("test-account",
            account_id="my-account",
            display_name="Test Service Account")
        test_iam = gcp.healthcare.ConsentStoreIamMember("test-iam",
            dataset=dataset.id,
            consent_store_id=my_consent.name,
            role="roles/editor",
            member=test_account.email.apply(lambda email: f"serviceAccount:{email}"))
        ```

        ## Import

        ConsentStore can be imported using any of these accepted formats:

        ```sh
         $ pulumi import gcp:healthcare/consentStore:ConsentStore default {{dataset}}/consentStores/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ConsentStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConsentStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dataset: Optional[pulumi.Input[str]] = None,
                 default_consent_ttl: Optional[pulumi.Input[str]] = None,
                 enable_consent_create_on_update: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConsentStoreArgs.__new__(ConsentStoreArgs)

            if dataset is None and not opts.urn:
                raise TypeError("Missing required property 'dataset'")
            __props__.__dict__["dataset"] = dataset
            __props__.__dict__["default_consent_ttl"] = default_consent_ttl
            __props__.__dict__["enable_consent_create_on_update"] = enable_consent_create_on_update
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["effective_labels"] = None
            __props__.__dict__["pulumi_labels"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["effectiveLabels", "pulumiLabels"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(ConsentStore, __self__).__init__(
            'gcp:healthcare/consentStore:ConsentStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dataset: Optional[pulumi.Input[str]] = None,
            default_consent_ttl: Optional[pulumi.Input[str]] = None,
            effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            enable_consent_create_on_update: Optional[pulumi.Input[bool]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ConsentStore':
        """
        Get an existing ConsentStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dataset: Identifies the dataset addressed by this request. Must be in the format
               'projects/{project}/locations/{location}/datasets/{dataset}'
               
               
               - - -
        :param pulumi.Input[str] default_consent_ttl: Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
               A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[bool] enable_consent_create_on_update: If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-supplied key-value pairs used to organize Consent stores.
               Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
               conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
               Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
               bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
               No more than 64 labels can be associated with a given store.
               An object containing a list of "key": value pairs.
               Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] name: The name of this ConsentStore, for example:
               "consent1"
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConsentStoreState.__new__(_ConsentStoreState)

        __props__.__dict__["dataset"] = dataset
        __props__.__dict__["default_consent_ttl"] = default_consent_ttl
        __props__.__dict__["effective_labels"] = effective_labels
        __props__.__dict__["enable_consent_create_on_update"] = enable_consent_create_on_update
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["pulumi_labels"] = pulumi_labels
        return ConsentStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dataset(self) -> pulumi.Output[str]:
        """
        Identifies the dataset addressed by this request. Must be in the format
        'projects/{project}/locations/{location}/datasets/{dataset}'


        - - -
        """
        return pulumi.get(self, "dataset")

    @property
    @pulumi.getter(name="defaultConsentTtl")
    def default_consent_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Default time to live for consents in this store. Must be at least 24 hours. Updating this field will not affect the expiration time of existing consents.
        A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "default_consent_ttl")

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @property
    @pulumi.getter(name="enableConsentCreateOnUpdate")
    def enable_consent_create_on_update(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, [consents.patch] [google.cloud.healthcare.v1.consent.UpdateConsent] creates the consent if it does not already exist.
        """
        return pulumi.get(self, "enable_consent_create_on_update")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User-supplied key-value pairs used to organize Consent stores.
        Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}][\\p{Ll}\\p{Lo}\\p{N}_-]{0,62}`
        Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        bytes, and must conform to the following PCRE regular expression: `[\\p{Ll}\\p{Lo}\\p{N}_-]{0,63}`
        No more than 64 labels can be associated with a given store.
        An object containing a list of "key": value pairs.
        Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this ConsentStore, for example:
        "consent1"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

