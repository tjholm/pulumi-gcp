# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CryptoKeyIAMPolicyArgs', 'CryptoKeyIAMPolicy']

@pulumi.input_type
class CryptoKeyIAMPolicyArgs:
    def __init__(__self__, *,
                 crypto_key_id: pulumi.Input[str],
                 policy_data: pulumi.Input[str]):
        """
        The set of arguments for constructing a CryptoKeyIAMPolicy resource.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        pulumi.set(__self__, "crypto_key_id", crypto_key_id)
        pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="cryptoKeyId")
    def crypto_key_id(self) -> pulumi.Input[str]:
        """
        The crypto key ID, in the form
        `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        the provider's project setting will be used as a fallback.
        """
        return pulumi.get(self, "crypto_key_id")

    @crypto_key_id.setter
    def crypto_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "crypto_key_id", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Input[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_data", value)


@pulumi.input_type
class _CryptoKeyIAMPolicyState:
    def __init__(__self__, *,
                 crypto_key_id: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CryptoKeyIAMPolicy resources.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        if crypto_key_id is not None:
            pulumi.set(__self__, "crypto_key_id", crypto_key_id)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if policy_data is not None:
            pulumi.set(__self__, "policy_data", policy_data)

    @property
    @pulumi.getter(name="cryptoKeyId")
    def crypto_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The crypto key ID, in the form
        `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        the provider's project setting will be used as a fallback.
        """
        return pulumi.get(self, "crypto_key_id")

    @crypto_key_id.setter
    def crypto_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crypto_key_id", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) The etag of the project's IAM policy.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> Optional[pulumi.Input[str]]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

    @policy_data.setter
    def policy_data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_data", value)


class CryptoKeyIAMPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crypto_key_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:

        * `kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
        * `kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
        * `kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.

        > **Note:** `kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `kms.CryptoKeyIAMBinding` and `kms.CryptoKeyIAMMember` or they will fight over what your policy should be.

        > **Note:** `kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        key = gcp.kms.CryptoKey("key",
            key_ring=keyring.id,
            rotation_period="100000s")
        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"],
        )])
        crypto_key = gcp.kms.CryptoKeyIAMPolicy("cryptoKey",
            crypto_key_id=key.id,
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            role="roles/cloudkms.cryptoKeyEncrypter",
        )])
        ```

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMBinding("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMBinding("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"],
            condition=gcp.kms.CryptoKeyIAMBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMMember("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMMember("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            member="user:jane@example.com",
            condition=gcp.kms.CryptoKeyIAMMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.

        ```sh
         $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; first the resource in question and then the role.

        These bindings can be imported using the `crypto_key_id` and role, e.g.

        ```sh
         $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `crypto_key_id`, e.g.

        ```sh
         $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key your-project-id/location-name/key-ring-name/key-name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CryptoKeyIAMPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:

        * `kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
        * `kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
        * `kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.

        > **Note:** `kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `kms.CryptoKeyIAMBinding` and `kms.CryptoKeyIAMMember` or they will fight over what your policy should be.

        > **Note:** `kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.

        ```python
        import pulumi
        import pulumi_gcp as gcp

        keyring = gcp.kms.KeyRing("keyring", location="global")
        key = gcp.kms.CryptoKey("key",
            key_ring=keyring.id,
            rotation_period="100000s")
        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"],
        )])
        crypto_key = gcp.kms.CryptoKeyIAMPolicy("cryptoKey",
            crypto_key_id=key.id,
            policy_data=admin.policy_data)
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        admin = gcp.organizations.get_iam_policy(bindings=[gcp.organizations.GetIAMPolicyBindingArgs(
            condition=gcp.organizations.GetIAMPolicyBindingConditionArgs(
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
                title="expires_after_2019_12_31",
            ),
            members=["user:jane@example.com"],
            role="roles/cloudkms.cryptoKeyEncrypter",
        )])
        ```

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMBinding("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"])
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMBinding("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            members=["user:jane@example.com"],
            condition=gcp.kms.CryptoKeyIAMBindingConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMMember("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            member="user:jane@example.com")
        ```

        With IAM Conditions:

        ```python
        import pulumi
        import pulumi_gcp as gcp

        crypto_key = gcp.kms.CryptoKeyIAMMember("cryptoKey",
            crypto_key_id=google_kms_crypto_key["key"]["id"],
            role="roles/cloudkms.cryptoKeyEncrypter",
            member="user:jane@example.com",
            condition=gcp.kms.CryptoKeyIAMMemberConditionArgs(
                title="expires_after_2019_12_31",
                description="Expiring at midnight of 2019-12-31",
                expression="request.time < timestamp(\"2020-01-01T00:00:00Z\")",
            ))
        ```

        ## Import

        IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.

        This member resource can be imported using the `crypto_key_id`, role, and member identity e.g.

        ```sh
         $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com"
        ```

         IAM binding imports use space-delimited identifiers; first the resource in question and then the role.

        These bindings can be imported using the `crypto_key_id` and role, e.g.

        ```sh
         $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key "your-project-id/location-name/key-ring-name/key-name roles/editor"
        ```

         IAM policy imports use the identifier of the resource in question.

        This policy resource can be imported using the `crypto_key_id`, e.g.

        ```sh
         $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy crypto_key your-project-id/location-name/key-ring-name/key-name
        ```

        :param str resource_name: The name of the resource.
        :param CryptoKeyIAMPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CryptoKeyIAMPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 crypto_key_id: Optional[pulumi.Input[str]] = None,
                 policy_data: Optional[pulumi.Input[str]] = None,
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
            __props__ = CryptoKeyIAMPolicyArgs.__new__(CryptoKeyIAMPolicyArgs)

            if crypto_key_id is None and not opts.urn:
                raise TypeError("Missing required property 'crypto_key_id'")
            __props__.__dict__["crypto_key_id"] = crypto_key_id
            if policy_data is None and not opts.urn:
                raise TypeError("Missing required property 'policy_data'")
            __props__.__dict__["policy_data"] = policy_data
            __props__.__dict__["etag"] = None
        super(CryptoKeyIAMPolicy, __self__).__init__(
            'gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            crypto_key_id: Optional[pulumi.Input[str]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            policy_data: Optional[pulumi.Input[str]] = None) -> 'CryptoKeyIAMPolicy':
        """
        Get an existing CryptoKeyIAMPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] crypto_key_id: The crypto key ID, in the form
               `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
               `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
               the provider's project setting will be used as a fallback.
        :param pulumi.Input[str] etag: (Computed) The etag of the project's IAM policy.
        :param pulumi.Input[str] policy_data: The policy data generated by
               a `organizations.getIAMPolicy` data source.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CryptoKeyIAMPolicyState.__new__(_CryptoKeyIAMPolicyState)

        __props__.__dict__["crypto_key_id"] = crypto_key_id
        __props__.__dict__["etag"] = etag
        __props__.__dict__["policy_data"] = policy_data
        return CryptoKeyIAMPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cryptoKeyId")
    def crypto_key_id(self) -> pulumi.Output[str]:
        """
        The crypto key ID, in the form
        `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
        `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
        the provider's project setting will be used as a fallback.
        """
        return pulumi.get(self, "crypto_key_id")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        (Computed) The etag of the project's IAM policy.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="policyData")
    def policy_data(self) -> pulumi.Output[str]:
        """
        The policy data generated by
        a `organizations.getIAMPolicy` data source.
        """
        return pulumi.get(self, "policy_data")

