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

__all__ = ['KeyRingImportJobArgs', 'KeyRingImportJob']

@pulumi.input_type
class KeyRingImportJobArgs:
    def __init__(__self__, *,
                 import_job_id: pulumi.Input[str],
                 import_method: pulumi.Input[str],
                 key_ring: pulumi.Input[str],
                 protection_level: pulumi.Input[str]):
        """
        The set of arguments for constructing a KeyRingImportJob resource.
        :param pulumi.Input[str] import_job_id: It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        :param pulumi.Input[str] import_method: The wrapping method to be used for incoming key material.
               Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        :param pulumi.Input[str] key_ring: The KeyRing that this import job belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[str] protection_level: The protection level of the ImportJob. This must match the protectionLevel of the
               versionTemplate on the CryptoKey you attempt to import into.
               Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        """
        pulumi.set(__self__, "import_job_id", import_job_id)
        pulumi.set(__self__, "import_method", import_method)
        pulumi.set(__self__, "key_ring", key_ring)
        pulumi.set(__self__, "protection_level", protection_level)

    @property
    @pulumi.getter(name="importJobId")
    def import_job_id(self) -> pulumi.Input[str]:
        """
        It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        """
        return pulumi.get(self, "import_job_id")

    @import_job_id.setter
    def import_job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "import_job_id", value)

    @property
    @pulumi.getter(name="importMethod")
    def import_method(self) -> pulumi.Input[str]:
        """
        The wrapping method to be used for incoming key material.
        Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        """
        return pulumi.get(self, "import_method")

    @import_method.setter
    def import_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "import_method", value)

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> pulumi.Input[str]:
        """
        The KeyRing that this import job belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @key_ring.setter
    def key_ring(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_ring", value)

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> pulumi.Input[str]:
        """
        The protection level of the ImportJob. This must match the protectionLevel of the
        versionTemplate on the CryptoKey you attempt to import into.
        Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        """
        return pulumi.get(self, "protection_level")

    @protection_level.setter
    def protection_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "protection_level", value)


@pulumi.input_type
class _KeyRingImportJobState:
    def __init__(__self__, *,
                 attestations: Optional[pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobAttestationArgs']]]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 import_job_id: Optional[pulumi.Input[str]] = None,
                 import_method: Optional[pulumi.Input[str]] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None,
                 public_keys: Optional[pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobPublicKeyArgs']]]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KeyRingImportJob resources.
        :param pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobAttestationArgs']]] attestations: Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
               statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
               ImportMethod is one with a protection level of HSM.
        :param pulumi.Input[str] expire_time: The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
        :param pulumi.Input[str] import_job_id: It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        :param pulumi.Input[str] import_method: The wrapping method to be used for incoming key material.
               Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        :param pulumi.Input[str] key_ring: The KeyRing that this import job belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[str] name: The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
        :param pulumi.Input[str] protection_level: The protection level of the ImportJob. This must match the protectionLevel of the
               versionTemplate on the CryptoKey you attempt to import into.
               Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        :param pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobPublicKeyArgs']]] public_keys: The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
        :param pulumi.Input[str] state: The current state of the ImportJob, indicating if it can be used.
        """
        if attestations is not None:
            pulumi.set(__self__, "attestations", attestations)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if import_job_id is not None:
            pulumi.set(__self__, "import_job_id", import_job_id)
        if import_method is not None:
            pulumi.set(__self__, "import_method", import_method)
        if key_ring is not None:
            pulumi.set(__self__, "key_ring", key_ring)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if protection_level is not None:
            pulumi.set(__self__, "protection_level", protection_level)
        if public_keys is not None:
            pulumi.set(__self__, "public_keys", public_keys)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def attestations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobAttestationArgs']]]]:
        """
        Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
        statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
        ImportMethod is one with a protection level of HSM.
        """
        return pulumi.get(self, "attestations")

    @attestations.setter
    def attestations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobAttestationArgs']]]]):
        pulumi.set(self, "attestations", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter(name="importJobId")
    def import_job_id(self) -> Optional[pulumi.Input[str]]:
        """
        It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        """
        return pulumi.get(self, "import_job_id")

    @import_job_id.setter
    def import_job_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "import_job_id", value)

    @property
    @pulumi.getter(name="importMethod")
    def import_method(self) -> Optional[pulumi.Input[str]]:
        """
        The wrapping method to be used for incoming key material.
        Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        """
        return pulumi.get(self, "import_method")

    @import_method.setter
    def import_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "import_method", value)

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> Optional[pulumi.Input[str]]:
        """
        The KeyRing that this import job belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @key_ring.setter
    def key_ring(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_ring", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> Optional[pulumi.Input[str]]:
        """
        The protection level of the ImportJob. This must match the protectionLevel of the
        versionTemplate on the CryptoKey you attempt to import into.
        Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        """
        return pulumi.get(self, "protection_level")

    @protection_level.setter
    def protection_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protection_level", value)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobPublicKeyArgs']]]]:
        """
        The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
        """
        return pulumi.get(self, "public_keys")

    @public_keys.setter
    def public_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KeyRingImportJobPublicKeyArgs']]]]):
        pulumi.set(self, "public_keys", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the ImportJob, indicating if it can be used.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class KeyRingImportJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 import_job_id: Optional[pulumi.Input[str]] = None,
                 import_method: Optional[pulumi.Input[str]] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A `KeyRingImportJob` can be used to create `CryptoKeys` and `CryptoKeyVersions` using pre-existing
        key material, generated outside of Cloud KMS. A `KeyRingImportJob` expires 3 days after it is created.
        Once expired, Cloud KMS will no longer be able to import or unwrap any key material that
        was wrapped with the `KeyRingImportJob`'s public key.

        > **Note:** KeyRingImportJobs cannot be deleted from Google Cloud Platform.
        Destroying a provider-managed KeyRingImportJob will remove it from state but
        *will not delete the resource from the project.*

        To get more information about KeyRingImportJob, see:

        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.importJobs)
        * How-to Guides
            * [Importing a key](https://cloud.google.com/kms/docs/importing-a-key)

        ## Example Usage

        ## Import

        KeyRingImportJob can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:kms/keyRingImportJob:KeyRingImportJob default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] import_job_id: It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        :param pulumi.Input[str] import_method: The wrapping method to be used for incoming key material.
               Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        :param pulumi.Input[str] key_ring: The KeyRing that this import job belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[str] protection_level: The protection level of the ImportJob. This must match the protectionLevel of the
               versionTemplate on the CryptoKey you attempt to import into.
               Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyRingImportJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A `KeyRingImportJob` can be used to create `CryptoKeys` and `CryptoKeyVersions` using pre-existing
        key material, generated outside of Cloud KMS. A `KeyRingImportJob` expires 3 days after it is created.
        Once expired, Cloud KMS will no longer be able to import or unwrap any key material that
        was wrapped with the `KeyRingImportJob`'s public key.

        > **Note:** KeyRingImportJobs cannot be deleted from Google Cloud Platform.
        Destroying a provider-managed KeyRingImportJob will remove it from state but
        *will not delete the resource from the project.*

        To get more information about KeyRingImportJob, see:

        * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.importJobs)
        * How-to Guides
            * [Importing a key](https://cloud.google.com/kms/docs/importing-a-key)

        ## Example Usage

        ## Import

        KeyRingImportJob can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:kms/keyRingImportJob:KeyRingImportJob default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param KeyRingImportJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyRingImportJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 import_job_id: Optional[pulumi.Input[str]] = None,
                 import_method: Optional[pulumi.Input[str]] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 protection_level: Optional[pulumi.Input[str]] = None,
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
            __props__ = KeyRingImportJobArgs.__new__(KeyRingImportJobArgs)

            if import_job_id is None and not opts.urn:
                raise TypeError("Missing required property 'import_job_id'")
            __props__.__dict__["import_job_id"] = import_job_id
            if import_method is None and not opts.urn:
                raise TypeError("Missing required property 'import_method'")
            __props__.__dict__["import_method"] = import_method
            if key_ring is None and not opts.urn:
                raise TypeError("Missing required property 'key_ring'")
            __props__.__dict__["key_ring"] = key_ring
            if protection_level is None and not opts.urn:
                raise TypeError("Missing required property 'protection_level'")
            __props__.__dict__["protection_level"] = protection_level
            __props__.__dict__["attestations"] = None
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["public_keys"] = None
            __props__.__dict__["state"] = None
        super(KeyRingImportJob, __self__).__init__(
            'gcp:kms/keyRingImportJob:KeyRingImportJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attestations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyRingImportJobAttestationArgs']]]]] = None,
            expire_time: Optional[pulumi.Input[str]] = None,
            import_job_id: Optional[pulumi.Input[str]] = None,
            import_method: Optional[pulumi.Input[str]] = None,
            key_ring: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            protection_level: Optional[pulumi.Input[str]] = None,
            public_keys: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyRingImportJobPublicKeyArgs']]]]] = None,
            state: Optional[pulumi.Input[str]] = None) -> 'KeyRingImportJob':
        """
        Get an existing KeyRingImportJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyRingImportJobAttestationArgs']]]] attestations: Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
               statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
               ImportMethod is one with a protection level of HSM.
        :param pulumi.Input[str] expire_time: The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
        :param pulumi.Input[str] import_job_id: It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        :param pulumi.Input[str] import_method: The wrapping method to be used for incoming key material.
               Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        :param pulumi.Input[str] key_ring: The KeyRing that this import job belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[str] name: The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
        :param pulumi.Input[str] protection_level: The protection level of the ImportJob. This must match the protectionLevel of the
               versionTemplate on the CryptoKey you attempt to import into.
               Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyRingImportJobPublicKeyArgs']]]] public_keys: The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
        :param pulumi.Input[str] state: The current state of the ImportJob, indicating if it can be used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyRingImportJobState.__new__(_KeyRingImportJobState)

        __props__.__dict__["attestations"] = attestations
        __props__.__dict__["expire_time"] = expire_time
        __props__.__dict__["import_job_id"] = import_job_id
        __props__.__dict__["import_method"] = import_method
        __props__.__dict__["key_ring"] = key_ring
        __props__.__dict__["name"] = name
        __props__.__dict__["protection_level"] = protection_level
        __props__.__dict__["public_keys"] = public_keys
        __props__.__dict__["state"] = state
        return KeyRingImportJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attestations(self) -> pulumi.Output[Sequence['outputs.KeyRingImportJobAttestation']]:
        """
        Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
        statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
        ImportMethod is one with a protection level of HSM.
        """
        return pulumi.get(self, "attestations")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="importJobId")
    def import_job_id(self) -> pulumi.Output[str]:
        """
        It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
        """
        return pulumi.get(self, "import_job_id")

    @property
    @pulumi.getter(name="importMethod")
    def import_method(self) -> pulumi.Output[str]:
        """
        The wrapping method to be used for incoming key material.
        Possible values are `RSA_OAEP_3072_SHA1_AES_256` and `RSA_OAEP_4096_SHA1_AES_256`.
        """
        return pulumi.get(self, "import_method")

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> pulumi.Output[str]:
        """
        The KeyRing that this import job belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="protectionLevel")
    def protection_level(self) -> pulumi.Output[str]:
        """
        The protection level of the ImportJob. This must match the protectionLevel of the
        versionTemplate on the CryptoKey you attempt to import into.
        Possible values are `SOFTWARE`, `HSM`, and `EXTERNAL`.
        """
        return pulumi.get(self, "protection_level")

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> pulumi.Output[Sequence['outputs.KeyRingImportJobPublicKey']]:
        """
        The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
        """
        return pulumi.get(self, "public_keys")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the ImportJob, indicating if it can be used.
        """
        return pulumi.get(self, "state")

