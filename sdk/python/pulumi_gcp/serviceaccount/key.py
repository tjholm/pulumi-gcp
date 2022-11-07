# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 service_account_id: pulumi.Input[str],
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 public_key_data: Optional[pulumi.Input[str]] = None,
                 public_key_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[str] service_account_id: The Service account id of the Key. This can be a string in the format
               `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
               the **full** email address of the service account or its name can be specified as a value, in which case the project will
               automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
               syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
               unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger a new key to be generated.
        :param pulumi.Input[str] key_algorithm: The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
               Valid values are listed at
               [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
               (only used on create)
        :param pulumi.Input[str] private_key_type: The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        :param pulumi.Input[str] public_key_data: Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        :param pulumi.Input[str] public_key_type: The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        """
        pulumi.set(__self__, "service_account_id", service_account_id)
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if key_algorithm is not None:
            pulumi.set(__self__, "key_algorithm", key_algorithm)
        if private_key_type is not None:
            pulumi.set(__self__, "private_key_type", private_key_type)
        if public_key_data is not None:
            pulumi.set(__self__, "public_key_data", public_key_data)
        if public_key_type is not None:
            pulumi.set(__self__, "public_key_type", public_key_type)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Input[str]:
        """
        The Service account id of the Key. This can be a string in the format
        `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
        the **full** email address of the service account or its name can be specified as a value, in which case the project will
        automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
        syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
        unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger a new key to be generated.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
        Valid values are listed at
        [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
        (only used on create)
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter(name="privateKeyType")
    def private_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        """
        return pulumi.get(self, "private_key_type")

    @private_key_type.setter
    def private_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_type", value)

    @property
    @pulumi.getter(name="publicKeyData")
    def public_key_data(self) -> Optional[pulumi.Input[str]]:
        """
        Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        """
        return pulumi.get(self, "public_key_data")

    @public_key_data.setter
    def public_key_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key_data", value)

    @property
    @pulumi.getter(name="publicKeyType")
    def public_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        """
        return pulumi.get(self, "public_key_type")

    @public_key_type.setter
    def public_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key_type", value)


@pulumi.input_type
class _KeyState:
    def __init__(__self__, *,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 public_key_data: Optional[pulumi.Input[str]] = None,
                 public_key_type: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 valid_after: Optional[pulumi.Input[str]] = None,
                 valid_before: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Key resources.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger a new key to be generated.
        :param pulumi.Input[str] key_algorithm: The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
               Valid values are listed at
               [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
               (only used on create)
        :param pulumi.Input[str] name: The name used for this key pair
        :param pulumi.Input[str] private_key: The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
               service account keys through the CLI or web console. This is only populated when creating a new key.
        :param pulumi.Input[str] private_key_type: The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        :param pulumi.Input[str] public_key: The public key, base64 encoded
        :param pulumi.Input[str] public_key_data: Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        :param pulumi.Input[str] public_key_type: The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        :param pulumi.Input[str] service_account_id: The Service account id of the Key. This can be a string in the format
               `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
               the **full** email address of the service account or its name can be specified as a value, in which case the project will
               automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
               syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
               unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        :param pulumi.Input[str] valid_after: The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[str] valid_before: The key can be used before this timestamp.
               A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        if keepers is not None:
            pulumi.set(__self__, "keepers", keepers)
        if key_algorithm is not None:
            pulumi.set(__self__, "key_algorithm", key_algorithm)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if private_key_type is not None:
            pulumi.set(__self__, "private_key_type", private_key_type)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)
        if public_key_data is not None:
            pulumi.set(__self__, "public_key_data", public_key_data)
        if public_key_type is not None:
            pulumi.set(__self__, "public_key_type", public_key_type)
        if service_account_id is not None:
            pulumi.set(__self__, "service_account_id", service_account_id)
        if valid_after is not None:
            pulumi.set(__self__, "valid_after", valid_after)
        if valid_before is not None:
            pulumi.set(__self__, "valid_before", valid_before)

    @property
    @pulumi.getter
    def keepers(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger a new key to be generated.
        """
        return pulumi.get(self, "keepers")

    @keepers.setter
    def keepers(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "keepers", value)

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
        Valid values are listed at
        [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
        (only used on create)
        """
        return pulumi.get(self, "key_algorithm")

    @key_algorithm.setter
    def key_algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_algorithm", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name used for this key pair
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
        service account keys through the CLI or web console. This is only populated when creating a new key.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyType")
    def private_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        """
        return pulumi.get(self, "private_key_type")

    @private_key_type.setter
    def private_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_type", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public key, base64 encoded
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter(name="publicKeyData")
    def public_key_data(self) -> Optional[pulumi.Input[str]]:
        """
        Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        """
        return pulumi.get(self, "public_key_data")

    @public_key_data.setter
    def public_key_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key_data", value)

    @property
    @pulumi.getter(name="publicKeyType")
    def public_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        """
        return pulumi.get(self, "public_key_type")

    @public_key_type.setter
    def public_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key_type", value)

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Service account id of the Key. This can be a string in the format
        `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
        the **full** email address of the service account or its name can be specified as a value, in which case the project will
        automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
        syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
        unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        """
        return pulumi.get(self, "service_account_id")

    @service_account_id.setter
    def service_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_id", value)

    @property
    @pulumi.getter(name="validAfter")
    def valid_after(self) -> Optional[pulumi.Input[str]]:
        """
        The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "valid_after")

    @valid_after.setter
    def valid_after(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_after", value)

    @property
    @pulumi.getter(name="validBefore")
    def valid_before(self) -> Optional[pulumi.Input[str]]:
        """
        The key can be used before this timestamp.
        A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "valid_before")

    @valid_before.setter
    def valid_before(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "valid_before", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 public_key_data: Optional[pulumi.Input[str]] = None,
                 public_key_type: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Creating A New Key

        ```python
        import pulumi
        import pulumi_gcp as gcp

        myaccount = gcp.service_account.Account("myaccount",
            account_id="myaccount",
            display_name="My Service Account")
        mykey = gcp.service_account.Key("mykey",
            service_account_id=myaccount.name,
            public_key_type="TYPE_X509_PEM_FILE")
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger a new key to be generated.
        :param pulumi.Input[str] key_algorithm: The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
               Valid values are listed at
               [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
               (only used on create)
        :param pulumi.Input[str] private_key_type: The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        :param pulumi.Input[str] public_key_data: Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        :param pulumi.Input[str] public_key_type: The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        :param pulumi.Input[str] service_account_id: The Service account id of the Key. This can be a string in the format
               `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
               the **full** email address of the service account or its name can be specified as a value, in which case the project will
               automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
               syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
               unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Creating A New Key

        ```python
        import pulumi
        import pulumi_gcp as gcp

        myaccount = gcp.service_account.Account("myaccount",
            account_id="myaccount",
            display_name="My Service Account")
        mykey = gcp.service_account.Key("mykey",
            service_account_id=myaccount.name,
            public_key_type="TYPE_X509_PEM_FILE")
        ```

        ## Import

        This resource does not support import.

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 key_algorithm: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 public_key_data: Optional[pulumi.Input[str]] = None,
                 public_key_type: Optional[pulumi.Input[str]] = None,
                 service_account_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["keepers"] = keepers
            __props__.__dict__["key_algorithm"] = key_algorithm
            __props__.__dict__["private_key_type"] = private_key_type
            __props__.__dict__["public_key_data"] = public_key_data
            __props__.__dict__["public_key_type"] = public_key_type
            if service_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_id'")
            __props__.__dict__["service_account_id"] = service_account_id
            __props__.__dict__["name"] = None
            __props__.__dict__["private_key"] = None
            __props__.__dict__["public_key"] = None
            __props__.__dict__["valid_after"] = None
            __props__.__dict__["valid_before"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Key, __self__).__init__(
            'gcp:serviceAccount/key:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            keepers: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            key_algorithm: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            private_key_type: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None,
            public_key_data: Optional[pulumi.Input[str]] = None,
            public_key_type: Optional[pulumi.Input[str]] = None,
            service_account_id: Optional[pulumi.Input[str]] = None,
            valid_after: Optional[pulumi.Input[str]] = None,
            valid_before: Optional[pulumi.Input[str]] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] keepers: Arbitrary map of values that, when changed, will trigger a new key to be generated.
        :param pulumi.Input[str] key_algorithm: The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
               Valid values are listed at
               [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
               (only used on create)
        :param pulumi.Input[str] name: The name used for this key pair
        :param pulumi.Input[str] private_key: The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
               service account keys through the CLI or web console. This is only populated when creating a new key.
        :param pulumi.Input[str] private_key_type: The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        :param pulumi.Input[str] public_key: The public key, base64 encoded
        :param pulumi.Input[str] public_key_data: Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        :param pulumi.Input[str] public_key_type: The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        :param pulumi.Input[str] service_account_id: The Service account id of the Key. This can be a string in the format
               `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
               the **full** email address of the service account or its name can be specified as a value, in which case the project will
               automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
               syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
               unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        :param pulumi.Input[str] valid_after: The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        :param pulumi.Input[str] valid_before: The key can be used before this timestamp.
               A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyState.__new__(_KeyState)

        __props__.__dict__["keepers"] = keepers
        __props__.__dict__["key_algorithm"] = key_algorithm
        __props__.__dict__["name"] = name
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["private_key_type"] = private_key_type
        __props__.__dict__["public_key"] = public_key
        __props__.__dict__["public_key_data"] = public_key_data
        __props__.__dict__["public_key_type"] = public_key_type
        __props__.__dict__["service_account_id"] = service_account_id
        __props__.__dict__["valid_after"] = valid_after
        __props__.__dict__["valid_before"] = valid_before
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def keepers(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Arbitrary map of values that, when changed, will trigger a new key to be generated.
        """
        return pulumi.get(self, "keepers")

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> pulumi.Output[Optional[str]]:
        """
        The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
        Valid values are listed at
        [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
        (only used on create)
        """
        return pulumi.get(self, "key_algorithm")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name used for this key pair
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The private key in JSON format, base64 encoded. This is what you normally get as a file when creating
        service account keys through the CLI or web console. This is only populated when creating a new key.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyType")
    def private_key_type(self) -> pulumi.Output[Optional[str]]:
        """
        The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
        """
        return pulumi.get(self, "private_key_type")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        The public key, base64 encoded
        """
        return pulumi.get(self, "public_key")

    @property
    @pulumi.getter(name="publicKeyData")
    def public_key_data(self) -> pulumi.Output[Optional[str]]:
        """
        Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
        """
        return pulumi.get(self, "public_key_data")

    @property
    @pulumi.getter(name="publicKeyType")
    def public_key_type(self) -> pulumi.Output[Optional[str]]:
        """
        The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
        """
        return pulumi.get(self, "public_key_type")

    @property
    @pulumi.getter(name="serviceAccountId")
    def service_account_id(self) -> pulumi.Output[str]:
        """
        The Service account id of the Key. This can be a string in the format
        `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
        the **full** email address of the service account or its name can be specified as a value, in which case the project will
        automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
        syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
        unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
        """
        return pulumi.get(self, "service_account_id")

    @property
    @pulumi.getter(name="validAfter")
    def valid_after(self) -> pulumi.Output[str]:
        """
        The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "valid_after")

    @property
    @pulumi.getter(name="validBefore")
    def valid_before(self) -> pulumi.Output[str]:
        """
        The key can be used before this timestamp.
        A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "valid_before")

