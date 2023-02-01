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

__all__ = ['CryptoKeyArgs', 'CryptoKey']

@pulumi.input_type
class CryptoKeyArgs:
    def __init__(__self__, *,
                 key_ring: pulumi.Input[str],
                 destroy_scheduled_duration: Optional[pulumi.Input[str]] = None,
                 import_only: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[str]] = None,
                 skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
                 version_template: Optional[pulumi.Input['CryptoKeyVersionTemplateArgs']] = None):
        """
        The set of arguments for constructing a CryptoKey resource.
        :param pulumi.Input[str] key_ring: The KeyRing that this key belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[str] destroy_scheduled_duration: The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
               If not specified at creation time, the default duration is 24 hours.
        :param pulumi.Input[bool] import_only: Whether this key may contain imported versions only.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata to apply to this resource.
        :param pulumi.Input[str] name: The resource name for the CryptoKey.
        :param pulumi.Input[str] purpose: The immutable purpose of this CryptoKey. See the
               [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
               for possible inputs.
               Default value is `ENCRYPT_DECRYPT`.
               Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
               The first rotation will take place after the specified period. The rotation period has
               the format of a decimal number with up to 9 fractional digits, followed by the
               letter `s` (seconds). It must be greater than a day (ie, 86400).
        :param pulumi.Input[bool] skip_initial_version_creation: If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
               You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        :param pulumi.Input['CryptoKeyVersionTemplateArgs'] version_template: A template describing settings for new crypto key versions.
               Structure is documented below.
        """
        pulumi.set(__self__, "key_ring", key_ring)
        if destroy_scheduled_duration is not None:
            pulumi.set(__self__, "destroy_scheduled_duration", destroy_scheduled_duration)
        if import_only is not None:
            pulumi.set(__self__, "import_only", import_only)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if rotation_period is not None:
            pulumi.set(__self__, "rotation_period", rotation_period)
        if skip_initial_version_creation is not None:
            pulumi.set(__self__, "skip_initial_version_creation", skip_initial_version_creation)
        if version_template is not None:
            pulumi.set(__self__, "version_template", version_template)

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> pulumi.Input[str]:
        """
        The KeyRing that this key belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @key_ring.setter
    def key_ring(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_ring", value)

    @property
    @pulumi.getter(name="destroyScheduledDuration")
    def destroy_scheduled_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
        If not specified at creation time, the default duration is 24 hours.
        """
        return pulumi.get(self, "destroy_scheduled_duration")

    @destroy_scheduled_duration.setter
    def destroy_scheduled_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destroy_scheduled_duration", value)

    @property
    @pulumi.getter(name="importOnly")
    def import_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this key may contain imported versions only.
        """
        return pulumi.get(self, "import_only")

    @import_only.setter
    def import_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "import_only", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels with user-defined metadata to apply to this resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the CryptoKey.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        The immutable purpose of this CryptoKey. See the
        [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        for possible inputs.
        Default value is `ENCRYPT_DECRYPT`.
        Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> Optional[pulumi.Input[str]]:
        """
        Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        The first rotation will take place after the specified period. The rotation period has
        the format of a decimal number with up to 9 fractional digits, followed by the
        letter `s` (seconds). It must be greater than a day (ie, 86400).
        """
        return pulumi.get(self, "rotation_period")

    @rotation_period.setter
    def rotation_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rotation_period", value)

    @property
    @pulumi.getter(name="skipInitialVersionCreation")
    def skip_initial_version_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        """
        return pulumi.get(self, "skip_initial_version_creation")

    @skip_initial_version_creation.setter
    def skip_initial_version_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_initial_version_creation", value)

    @property
    @pulumi.getter(name="versionTemplate")
    def version_template(self) -> Optional[pulumi.Input['CryptoKeyVersionTemplateArgs']]:
        """
        A template describing settings for new crypto key versions.
        Structure is documented below.
        """
        return pulumi.get(self, "version_template")

    @version_template.setter
    def version_template(self, value: Optional[pulumi.Input['CryptoKeyVersionTemplateArgs']]):
        pulumi.set(self, "version_template", value)


@pulumi.input_type
class _CryptoKeyState:
    def __init__(__self__, *,
                 destroy_scheduled_duration: Optional[pulumi.Input[str]] = None,
                 import_only: Optional[pulumi.Input[bool]] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[str]] = None,
                 skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
                 version_template: Optional[pulumi.Input['CryptoKeyVersionTemplateArgs']] = None):
        """
        Input properties used for looking up and filtering CryptoKey resources.
        :param pulumi.Input[str] destroy_scheduled_duration: The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
               If not specified at creation time, the default duration is 24 hours.
        :param pulumi.Input[bool] import_only: Whether this key may contain imported versions only.
        :param pulumi.Input[str] key_ring: The KeyRing that this key belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata to apply to this resource.
        :param pulumi.Input[str] name: The resource name for the CryptoKey.
        :param pulumi.Input[str] purpose: The immutable purpose of this CryptoKey. See the
               [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
               for possible inputs.
               Default value is `ENCRYPT_DECRYPT`.
               Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
               The first rotation will take place after the specified period. The rotation period has
               the format of a decimal number with up to 9 fractional digits, followed by the
               letter `s` (seconds). It must be greater than a day (ie, 86400).
        :param pulumi.Input[bool] skip_initial_version_creation: If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
               You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        :param pulumi.Input['CryptoKeyVersionTemplateArgs'] version_template: A template describing settings for new crypto key versions.
               Structure is documented below.
        """
        if destroy_scheduled_duration is not None:
            pulumi.set(__self__, "destroy_scheduled_duration", destroy_scheduled_duration)
        if import_only is not None:
            pulumi.set(__self__, "import_only", import_only)
        if key_ring is not None:
            pulumi.set(__self__, "key_ring", key_ring)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if purpose is not None:
            pulumi.set(__self__, "purpose", purpose)
        if rotation_period is not None:
            pulumi.set(__self__, "rotation_period", rotation_period)
        if skip_initial_version_creation is not None:
            pulumi.set(__self__, "skip_initial_version_creation", skip_initial_version_creation)
        if version_template is not None:
            pulumi.set(__self__, "version_template", version_template)

    @property
    @pulumi.getter(name="destroyScheduledDuration")
    def destroy_scheduled_duration(self) -> Optional[pulumi.Input[str]]:
        """
        The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
        If not specified at creation time, the default duration is 24 hours.
        """
        return pulumi.get(self, "destroy_scheduled_duration")

    @destroy_scheduled_duration.setter
    def destroy_scheduled_duration(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destroy_scheduled_duration", value)

    @property
    @pulumi.getter(name="importOnly")
    def import_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this key may contain imported versions only.
        """
        return pulumi.get(self, "import_only")

    @import_only.setter
    def import_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "import_only", value)

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> Optional[pulumi.Input[str]]:
        """
        The KeyRing that this key belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @key_ring.setter
    def key_ring(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_ring", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels with user-defined metadata to apply to this resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the CryptoKey.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def purpose(self) -> Optional[pulumi.Input[str]]:
        """
        The immutable purpose of this CryptoKey. See the
        [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        for possible inputs.
        Default value is `ENCRYPT_DECRYPT`.
        Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        """
        return pulumi.get(self, "purpose")

    @purpose.setter
    def purpose(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "purpose", value)

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> Optional[pulumi.Input[str]]:
        """
        Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        The first rotation will take place after the specified period. The rotation period has
        the format of a decimal number with up to 9 fractional digits, followed by the
        letter `s` (seconds). It must be greater than a day (ie, 86400).
        """
        return pulumi.get(self, "rotation_period")

    @rotation_period.setter
    def rotation_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rotation_period", value)

    @property
    @pulumi.getter(name="skipInitialVersionCreation")
    def skip_initial_version_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        """
        return pulumi.get(self, "skip_initial_version_creation")

    @skip_initial_version_creation.setter
    def skip_initial_version_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_initial_version_creation", value)

    @property
    @pulumi.getter(name="versionTemplate")
    def version_template(self) -> Optional[pulumi.Input['CryptoKeyVersionTemplateArgs']]:
        """
        A template describing settings for new crypto key versions.
        Structure is documented below.
        """
        return pulumi.get(self, "version_template")

    @version_template.setter
    def version_template(self, value: Optional[pulumi.Input['CryptoKeyVersionTemplateArgs']]):
        pulumi.set(self, "version_template", value)


class CryptoKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destroy_scheduled_duration: Optional[pulumi.Input[str]] = None,
                 import_only: Optional[pulumi.Input[bool]] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[str]] = None,
                 skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
                 version_template: Optional[pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Kms Crypto Key Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        example_key = gcp.kms.CryptoKey("example-key",
            key_ring=keyring.id,
            rotation_period="100000s")
        ```
        ### Kms Crypto Key Asymmetric Sign

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        example_asymmetric_sign_key = gcp.kms.CryptoKey("example-asymmetric-sign-key",
            key_ring=keyring.id,
            purpose="ASYMMETRIC_SIGN",
            version_template=gcp.kms.CryptoKeyVersionTemplateArgs(
                algorithm="EC_SIGN_P384_SHA384",
            ))
        ```

        ## Import

        CryptoKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
        ```

        ```sh
         $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destroy_scheduled_duration: The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
               If not specified at creation time, the default duration is 24 hours.
        :param pulumi.Input[bool] import_only: Whether this key may contain imported versions only.
        :param pulumi.Input[str] key_ring: The KeyRing that this key belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata to apply to this resource.
        :param pulumi.Input[str] name: The resource name for the CryptoKey.
        :param pulumi.Input[str] purpose: The immutable purpose of this CryptoKey. See the
               [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
               for possible inputs.
               Default value is `ENCRYPT_DECRYPT`.
               Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
               The first rotation will take place after the specified period. The rotation period has
               the format of a decimal number with up to 9 fractional digits, followed by the
               letter `s` (seconds). It must be greater than a day (ie, 86400).
        :param pulumi.Input[bool] skip_initial_version_creation: If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
               You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        :param pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']] version_template: A template describing settings for new crypto key versions.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CryptoKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Kms Crypto Key Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        example_key = gcp.kms.CryptoKey("example-key",
            key_ring=keyring.id,
            rotation_period="100000s")
        ```
        ### Kms Crypto Key Asymmetric Sign

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        example_asymmetric_sign_key = gcp.kms.CryptoKey("example-asymmetric-sign-key",
            key_ring=keyring.id,
            purpose="ASYMMETRIC_SIGN",
            version_template=gcp.kms.CryptoKeyVersionTemplateArgs(
                algorithm="EC_SIGN_P384_SHA384",
            ))
        ```

        ## Import

        CryptoKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
        ```

        ```sh
         $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param CryptoKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CryptoKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destroy_scheduled_duration: Optional[pulumi.Input[str]] = None,
                 import_only: Optional[pulumi.Input[bool]] = None,
                 key_ring: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 purpose: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[str]] = None,
                 skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
                 version_template: Optional[pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CryptoKeyArgs.__new__(CryptoKeyArgs)

            __props__.__dict__["destroy_scheduled_duration"] = destroy_scheduled_duration
            __props__.__dict__["import_only"] = import_only
            if key_ring is None and not opts.urn:
                raise TypeError("Missing required property 'key_ring'")
            __props__.__dict__["key_ring"] = key_ring
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            __props__.__dict__["purpose"] = purpose
            __props__.__dict__["rotation_period"] = rotation_period
            __props__.__dict__["skip_initial_version_creation"] = skip_initial_version_creation
            __props__.__dict__["version_template"] = version_template
        super(CryptoKey, __self__).__init__(
            'gcp:kms/cryptoKey:CryptoKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destroy_scheduled_duration: Optional[pulumi.Input[str]] = None,
            import_only: Optional[pulumi.Input[bool]] = None,
            key_ring: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            purpose: Optional[pulumi.Input[str]] = None,
            rotation_period: Optional[pulumi.Input[str]] = None,
            skip_initial_version_creation: Optional[pulumi.Input[bool]] = None,
            version_template: Optional[pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']]] = None) -> 'CryptoKey':
        """
        Get an existing CryptoKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destroy_scheduled_duration: The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
               If not specified at creation time, the default duration is 24 hours.
        :param pulumi.Input[bool] import_only: Whether this key may contain imported versions only.
        :param pulumi.Input[str] key_ring: The KeyRing that this key belongs to.
               Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels with user-defined metadata to apply to this resource.
        :param pulumi.Input[str] name: The resource name for the CryptoKey.
        :param pulumi.Input[str] purpose: The immutable purpose of this CryptoKey. See the
               [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
               for possible inputs.
               Default value is `ENCRYPT_DECRYPT`.
               Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        :param pulumi.Input[str] rotation_period: Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
               The first rotation will take place after the specified period. The rotation period has
               the format of a decimal number with up to 9 fractional digits, followed by the
               letter `s` (seconds). It must be greater than a day (ie, 86400).
        :param pulumi.Input[bool] skip_initial_version_creation: If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
               You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        :param pulumi.Input[pulumi.InputType['CryptoKeyVersionTemplateArgs']] version_template: A template describing settings for new crypto key versions.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CryptoKeyState.__new__(_CryptoKeyState)

        __props__.__dict__["destroy_scheduled_duration"] = destroy_scheduled_duration
        __props__.__dict__["import_only"] = import_only
        __props__.__dict__["key_ring"] = key_ring
        __props__.__dict__["labels"] = labels
        __props__.__dict__["name"] = name
        __props__.__dict__["purpose"] = purpose
        __props__.__dict__["rotation_period"] = rotation_period
        __props__.__dict__["skip_initial_version_creation"] = skip_initial_version_creation
        __props__.__dict__["version_template"] = version_template
        return CryptoKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destroyScheduledDuration")
    def destroy_scheduled_duration(self) -> pulumi.Output[str]:
        """
        The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
        If not specified at creation time, the default duration is 24 hours.
        """
        return pulumi.get(self, "destroy_scheduled_duration")

    @property
    @pulumi.getter(name="importOnly")
    def import_only(self) -> pulumi.Output[bool]:
        """
        Whether this key may contain imported versions only.
        """
        return pulumi.get(self, "import_only")

    @property
    @pulumi.getter(name="keyRing")
    def key_ring(self) -> pulumi.Output[str]:
        """
        The KeyRing that this key belongs to.
        Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        """
        return pulumi.get(self, "key_ring")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Labels with user-defined metadata to apply to this resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the CryptoKey.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def purpose(self) -> pulumi.Output[Optional[str]]:
        """
        The immutable purpose of this CryptoKey. See the
        [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        for possible inputs.
        Default value is `ENCRYPT_DECRYPT`.
        Possible values are `ENCRYPT_DECRYPT`, `ASYMMETRIC_SIGN`, `ASYMMETRIC_DECRYPT`, and `MAC`.
        """
        return pulumi.get(self, "purpose")

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Output[Optional[str]]:
        """
        Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        The first rotation will take place after the specified period. The rotation period has
        the format of a decimal number with up to 9 fractional digits, followed by the
        letter `s` (seconds). It must be greater than a day (ie, 86400).
        """
        return pulumi.get(self, "rotation_period")

    @property
    @pulumi.getter(name="skipInitialVersionCreation")
    def skip_initial_version_creation(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        You must use the `kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        """
        return pulumi.get(self, "skip_initial_version_creation")

    @property
    @pulumi.getter(name="versionTemplate")
    def version_template(self) -> pulumi.Output['outputs.CryptoKeyVersionTemplate']:
        """
        A template describing settings for new crypto key versions.
        Structure is documented below.
        """
        return pulumi.get(self, "version_template")

