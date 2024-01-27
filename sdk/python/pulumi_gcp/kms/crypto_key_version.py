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

__all__ = ['CryptoKeyVersionArgs', 'CryptoKeyVersion']

@pulumi.input_type
class CryptoKeyVersionArgs:
    def __init__(__self__, *,
                 crypto_key: pulumi.Input[str],
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CryptoKeyVersion resource.
        :param pulumi.Input[str] crypto_key: The name of the cryptoKey associated with the CryptoKeyVersions.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
               
               
               - - -
        :param pulumi.Input[str] state: The current state of the CryptoKeyVersion.
               Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        pulumi.set(__self__, "crypto_key", crypto_key)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="cryptoKey")
    def crypto_key(self) -> pulumi.Input[str]:
        """
        The name of the cryptoKey associated with the CryptoKeyVersions.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`


        - - -
        """
        return pulumi.get(self, "crypto_key")

    @crypto_key.setter
    def crypto_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "crypto_key", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the CryptoKeyVersion.
        Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _CryptoKeyVersionState:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 attestations: Optional[pulumi.Input[Sequence[pulumi.Input['CryptoKeyVersionAttestationArgs']]]] = None,
                 crypto_key: Optional[pulumi.Input[str]] = None,
                 generate_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CryptoKeyVersion resources.
        :param pulumi.Input[str] algorithm: The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        :param pulumi.Input[Sequence[pulumi.Input['CryptoKeyVersionAttestationArgs']]] attestations: Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
               Only provided for key versions with protectionLevel HSM.
               Structure is documented below.
        :param pulumi.Input[str] crypto_key: The name of the cryptoKey associated with the CryptoKeyVersions.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
               
               
               - - -
        :param pulumi.Input[str] generate_time: The time this CryptoKeyVersion key material was generated
        :param pulumi.Input[str] name: The resource name for this CryptoKeyVersion.
        :param pulumi.Input[str] protection_level: The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        :param pulumi.Input[str] state: The current state of the CryptoKeyVersion.
               Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if attestations is not None:
            pulumi.set(__self__, "attestations", attestations)
        if crypto_key is not None:
            pulumi.set(__self__, "crypto_key", crypto_key)
        if generate_time is not None:
            pulumi.set(__self__, "generate_time", generate_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protection_level is not None:
            pulumi.set(__self__, "protection_level", protection_level)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def attestations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CryptoKeyVersionAttestationArgs']]]]:
        """
        Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
        Only provided for key versions with protectionLevel HSM.
        Structure is documented below.
        """
        return pulumi.get(self, "attestations")

    @attestations.setter
    def attestations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CryptoKeyVersionAttestationArgs']]]]):
        pulumi.set(self, "attestations", value)

    @property
    @pulumi.getter(name="cryptoKey")
    def crypto_key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cryptoKey associated with the CryptoKeyVersions.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`


        - - -
        """
        return pulumi.get(self, "crypto_key")

    @crypto_key.setter
    def crypto_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crypto_key", value)

    @property
    @pulumi.getter(name="generateTime")
    def generate_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time this CryptoKeyVersion key material was generated
        """
        return pulumi.get(self, "generate_time")

    @generate_time.setter
    def generate_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "generate_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for this CryptoKeyVersion.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> Optional[pulumi.Input[str]]:
        """
        The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        """
        return pulumi.get(self, "protection_level")

    @protection_level.setter
    def protection_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protection_level", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the CryptoKeyVersion.
        Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class CryptoKeyVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crypto_key: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A `CryptoKeyVersion` represents an individual cryptographic key, and the associated key material.

        Destroying a cryptoKeyVersion will not delete the resource from the project.

        To get more information about CryptoKeyVersion, see:

        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions)
        * How-to Guides
            * [Creating a key Version](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions/create)

        ## Example Usage
        ### Kms Crypto Key Version Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        cryptokey = gcp.kms.CryptoKey("cryptokey",
            key_ring=keyring.id,
            rotation_period="7776000s")
        example_key = gcp.kms.CryptoKeyVersion("example-key", crypto_key=cryptokey.id)
        ```

        ## Import

        CryptoKeyVersion can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, CryptoKeyVersion can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:kms/cryptoKeyVersion:CryptoKeyVersion default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key: The name of the cryptoKey associated with the CryptoKeyVersions.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
               
               
               - - -
        :param pulumi.Input[str] state: The current state of the CryptoKeyVersion.
               Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CryptoKeyVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A `CryptoKeyVersion` represents an individual cryptographic key, and the associated key material.

        Destroying a cryptoKeyVersion will not delete the resource from the project.

        To get more information about CryptoKeyVersion, see:

        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions)
        * How-to Guides
            * [Creating a key Version](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions/create)

        ## Example Usage
        ### Kms Crypto Key Version Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        cryptokey = gcp.kms.CryptoKey("cryptokey",
            key_ring=keyring.id,
            rotation_period="7776000s")
        example_key = gcp.kms.CryptoKeyVersion("example-key", crypto_key=cryptokey.id)
        ```

        ## Import

        CryptoKeyVersion can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, CryptoKeyVersion can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:kms/cryptoKeyVersion:CryptoKeyVersion default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param CryptoKeyVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CryptoKeyVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crypto_key: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CryptoKeyVersionArgs.__new__(CryptoKeyVersionArgs)

            if crypto_key is None and not opts.urn:
                raise TypeError("Missing required property 'crypto_key'")
            __props__.__dict__["crypto_key"] = crypto_key
            __props__.__dict__["state"] = state
            __props__.__dict__["algorithm"] = None
            __props__.__dict__["attestations"] = None
            __props__.__dict__["generate_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["protection_level"] = None
        super(CryptoKeyVersion, __self__).__init__(
            'gcp:kms/cryptoKeyVersion:CryptoKeyVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm: Optional[pulumi.Input[str]] = None,
            attestations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CryptoKeyVersionAttestationArgs']]]]] = None,
            crypto_key: Optional[pulumi.Input[str]] = None,
            generate_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protection_level: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'CryptoKeyVersion':
        """
        Get an existing CryptoKeyVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CryptoKeyVersionAttestationArgs']]]] attestations: Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
               Only provided for key versions with protectionLevel HSM.
               Structure is documented below.
        :param pulumi.Input[str] crypto_key: The name of the cryptoKey associated with the CryptoKeyVersions.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
               
               
               - - -
        :param pulumi.Input[str] generate_time: The time this CryptoKeyVersion key material was generated
        :param pulumi.Input[str] name: The resource name for this CryptoKeyVersion.
        :param pulumi.Input[str] protection_level: The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        :param pulumi.Input[str] state: The current state of the CryptoKeyVersion.
               Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CryptoKeyVersionState.__new__(_CryptoKeyVersionState)

        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["attestations"] = attestations
        __props__.__dict__["crypto_key"] = crypto_key
        __props__.__dict__["generate_time"] = generate_time
        __props__.__dict__["name"] = name
        __props__.__dict__["protection_level"] = protection_level
        __props__.__dict__["state"] = state
        return CryptoKeyVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[str]:
        """
        The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def attestations(self) -> pulumi.Output[Sequence['outputs.CryptoKeyVersionAttestation']]:
        """
        Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
        Only provided for key versions with protectionLevel HSM.
        Structure is documented below.
        """
        return pulumi.get(self, "attestations")

    @property
    @pulumi.getter(name="cryptoKey")
    def crypto_key(self) -> pulumi.Output[str]:
        """
        The name of the cryptoKey associated with the CryptoKeyVersions.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`


        - - -
        """
        return pulumi.get(self, "crypto_key")

    @property
    @pulumi.getter(name="generateTime")
    def generate_time(self) -> pulumi.Output[str]:
        """
        The time this CryptoKeyVersion key material was generated
        """
        return pulumi.get(self, "generate_time")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for this CryptoKeyVersion.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> pulumi.Output[str]:
        """
        The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        """
        return pulumi.get(self, "protection_level")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the CryptoKeyVersion.
        Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        """
        return pulumi.get(self, "state")

