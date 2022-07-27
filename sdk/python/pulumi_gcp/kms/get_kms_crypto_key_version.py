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

__all__ = [
    'GetKMSCryptoKeyVersionResult',
    'AwaitableGetKMSCryptoKeyVersionResult',
    'get_kms_crypto_key_version',
    'get_kms_crypto_key_version_output',
]

@pulumi.output_type
class GetKMSCryptoKeyVersionResult:
    """
    A collection of values returned by getKMSCryptoKeyVersion.
    """
    def __init__(__self__, algorithm=None, crypto_key=None, id=None, name=None, protection_level=None, public_keys=None, state=None, version=None):
        if algorithm and not isinstance(algorithm, str):
            raise TypeError("Expected argument 'algorithm' to be a str")
        pulumi.set(__self__, "algorithm", algorithm)
        if crypto_key and not isinstance(crypto_key, str):
            raise TypeError("Expected argument 'crypto_key' to be a str")
        pulumi.set(__self__, "crypto_key", crypto_key)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if protection_level and not isinstance(protection_level, str):
            raise TypeError("Expected argument 'protection_level' to be a str")
        pulumi.set(__self__, "protection_level", protection_level)
        if public_keys and not isinstance(public_keys, list):
            raise TypeError("Expected argument 'public_keys' to be a list")
        pulumi.set(__self__, "public_keys", public_keys)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if version and not isinstance(version, int):
            raise TypeError("Expected argument 'version' to be a int")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def algorithm(self) -> str:
        """
        The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter(name="cryptoKey")
    def crypto_key(self) -> str:
        return pulumi.get(self, "crypto_key")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for this CryptoKeyVersion in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> str:
        """
        The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion. See the [protection_level reference](https://cloud.google.com/kms/docs/reference/rest/v1/ProtectionLevel) for possible outputs.
        """
        return pulumi.get(self, "protection_level")

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Sequence['outputs.GetKMSCryptoKeyVersionPublicKeyResult']:
        """
        If the enclosing CryptoKey has purpose `ASYMMETRIC_SIGN` or `ASYMMETRIC_DECRYPT`, this block contains details about the public key associated to this CryptoKeyVersion. Structure is documented below.
        """
        return pulumi.get(self, "public_keys")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the CryptoKeyVersion. See the [state reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions#CryptoKeyVersion.CryptoKeyVersionState) for possible outputs.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def version(self) -> Optional[int]:
        return pulumi.get(self, "version")


class AwaitableGetKMSCryptoKeyVersionResult(GetKMSCryptoKeyVersionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKMSCryptoKeyVersionResult(
            algorithm=self.algorithm,
            crypto_key=self.crypto_key,
            id=self.id,
            name=self.name,
            protection_level=self.protection_level,
            public_keys=self.public_keys,
            state=self.state,
            version=self.version)


def get_kms_crypto_key_version(crypto_key: Optional[str] = None,
                               version: Optional[int] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKMSCryptoKeyVersionResult:
    """
    Provides access to a Google Cloud Platform KMS CryptoKeyVersion. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_version)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions).

    A CryptoKeyVersion represents an individual cryptographic key, and the associated key material.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_key_ring = gcp.kms.get_kms_key_ring(name="my-key-ring",
        location="us-central1")
    my_crypto_key = gcp.kms.get_kms_crypto_key(name="my-crypto-key",
        key_ring=my_key_ring.id)
    my_crypto_key_version = gcp.kms.get_kms_crypto_key_version(crypto_key=data["google_kms_crypto_key"]["my_key"]["id"])
    ```


    :param str crypto_key: The `id` of the Google Cloud Platform CryptoKey to which the key version belongs. This is also the `id` field of the 
           `kms.CryptoKey` resource/datasource.
    :param int version: The version number for this CryptoKeyVersion. Defaults to `1`.
    """
    __args__ = dict()
    __args__['cryptoKey'] = crypto_key
    __args__['version'] = version
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('gcp:kms/getKMSCryptoKeyVersion:getKMSCryptoKeyVersion', __args__, opts=opts, typ=GetKMSCryptoKeyVersionResult).value

    return AwaitableGetKMSCryptoKeyVersionResult(
        algorithm=__ret__.algorithm,
        crypto_key=__ret__.crypto_key,
        id=__ret__.id,
        name=__ret__.name,
        protection_level=__ret__.protection_level,
        public_keys=__ret__.public_keys,
        state=__ret__.state,
        version=__ret__.version)


@_utilities.lift_output_func(get_kms_crypto_key_version)
def get_kms_crypto_key_version_output(crypto_key: Optional[pulumi.Input[str]] = None,
                                      version: Optional[pulumi.Input[Optional[int]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKMSCryptoKeyVersionResult]:
    """
    Provides access to a Google Cloud Platform KMS CryptoKeyVersion. For more information see
    [the official documentation](https://cloud.google.com/kms/docs/object-hierarchy#key_version)
    and
    [API](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions).

    A CryptoKeyVersion represents an individual cryptographic key, and the associated key material.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_gcp as gcp

    my_key_ring = gcp.kms.get_kms_key_ring(name="my-key-ring",
        location="us-central1")
    my_crypto_key = gcp.kms.get_kms_crypto_key(name="my-crypto-key",
        key_ring=my_key_ring.id)
    my_crypto_key_version = gcp.kms.get_kms_crypto_key_version(crypto_key=data["google_kms_crypto_key"]["my_key"]["id"])
    ```


    :param str crypto_key: The `id` of the Google Cloud Platform CryptoKey to which the key version belongs. This is also the `id` field of the 
           `kms.CryptoKey` resource/datasource.
    :param int version: The version number for this CryptoKeyVersion. Defaults to `1`.
    """
    ...
