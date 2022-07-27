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
    'InboundSamlConfigIdpConfig',
    'InboundSamlConfigIdpConfigIdpCertificate',
    'InboundSamlConfigSpConfig',
    'InboundSamlConfigSpConfigSpCertificate',
    'TenantInboundSamlConfigIdpConfig',
    'TenantInboundSamlConfigIdpConfigIdpCertificate',
    'TenantInboundSamlConfigSpConfig',
    'TenantInboundSamlConfigSpConfigSpCertificate',
]

@pulumi.output_type
class InboundSamlConfigIdpConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "idpCertificates":
            suggest = "idp_certificates"
        elif key == "idpEntityId":
            suggest = "idp_entity_id"
        elif key == "ssoUrl":
            suggest = "sso_url"
        elif key == "signRequest":
            suggest = "sign_request"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InboundSamlConfigIdpConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InboundSamlConfigIdpConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InboundSamlConfigIdpConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 idp_certificates: Sequence['outputs.InboundSamlConfigIdpConfigIdpCertificate'],
                 idp_entity_id: str,
                 sso_url: str,
                 sign_request: Optional[bool] = None):
        """
        :param Sequence['InboundSamlConfigIdpConfigIdpCertificateArgs'] idp_certificates: The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
               Structure is documented below.
        :param str idp_entity_id: Unique identifier for all SAML entities
        :param str sso_url: URL to send Authentication request to.
        :param bool sign_request: Indicates if outbounding SAMLRequest should be signed.
        """
        pulumi.set(__self__, "idp_certificates", idp_certificates)
        pulumi.set(__self__, "idp_entity_id", idp_entity_id)
        pulumi.set(__self__, "sso_url", sso_url)
        if sign_request is not None:
            pulumi.set(__self__, "sign_request", sign_request)

    @property
    @pulumi.getter(name="idpCertificates")
    def idp_certificates(self) -> Sequence['outputs.InboundSamlConfigIdpConfigIdpCertificate']:
        """
        The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
        Structure is documented below.
        """
        return pulumi.get(self, "idp_certificates")

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> str:
        """
        Unique identifier for all SAML entities
        """
        return pulumi.get(self, "idp_entity_id")

    @property
    @pulumi.getter(name="ssoUrl")
    def sso_url(self) -> str:
        """
        URL to send Authentication request to.
        """
        return pulumi.get(self, "sso_url")

    @property
    @pulumi.getter(name="signRequest")
    def sign_request(self) -> Optional[bool]:
        """
        Indicates if outbounding SAMLRequest should be signed.
        """
        return pulumi.get(self, "sign_request")


@pulumi.output_type
class InboundSamlConfigIdpConfigIdpCertificate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "x509Certificate":
            suggest = "x509_certificate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InboundSamlConfigIdpConfigIdpCertificate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InboundSamlConfigIdpConfigIdpCertificate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InboundSamlConfigIdpConfigIdpCertificate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 x509_certificate: Optional[str] = None):
        """
        :param str x509_certificate: -
               The x509 certificate
        """
        if x509_certificate is not None:
            pulumi.set(__self__, "x509_certificate", x509_certificate)

    @property
    @pulumi.getter(name="x509Certificate")
    def x509_certificate(self) -> Optional[str]:
        """
        -
        The x509 certificate
        """
        return pulumi.get(self, "x509_certificate")


@pulumi.output_type
class InboundSamlConfigSpConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "callbackUri":
            suggest = "callback_uri"
        elif key == "spCertificates":
            suggest = "sp_certificates"
        elif key == "spEntityId":
            suggest = "sp_entity_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InboundSamlConfigSpConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InboundSamlConfigSpConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InboundSamlConfigSpConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 callback_uri: Optional[str] = None,
                 sp_certificates: Optional[Sequence['outputs.InboundSamlConfigSpConfigSpCertificate']] = None,
                 sp_entity_id: Optional[str] = None):
        """
        :param str callback_uri: Callback URI where responses from IDP are handled. Must start with `https://`.
        :param Sequence['InboundSamlConfigSpConfigSpCertificateArgs'] sp_certificates: -
               The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
               Structure is documented below.
        :param str sp_entity_id: Unique identifier for all SAML entities.
        """
        if callback_uri is not None:
            pulumi.set(__self__, "callback_uri", callback_uri)
        if sp_certificates is not None:
            pulumi.set(__self__, "sp_certificates", sp_certificates)
        if sp_entity_id is not None:
            pulumi.set(__self__, "sp_entity_id", sp_entity_id)

    @property
    @pulumi.getter(name="callbackUri")
    def callback_uri(self) -> Optional[str]:
        """
        Callback URI where responses from IDP are handled. Must start with `https://`.
        """
        return pulumi.get(self, "callback_uri")

    @property
    @pulumi.getter(name="spCertificates")
    def sp_certificates(self) -> Optional[Sequence['outputs.InboundSamlConfigSpConfigSpCertificate']]:
        """
        -
        The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
        Structure is documented below.
        """
        return pulumi.get(self, "sp_certificates")

    @property
    @pulumi.getter(name="spEntityId")
    def sp_entity_id(self) -> Optional[str]:
        """
        Unique identifier for all SAML entities.
        """
        return pulumi.get(self, "sp_entity_id")


@pulumi.output_type
class InboundSamlConfigSpConfigSpCertificate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "x509Certificate":
            suggest = "x509_certificate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InboundSamlConfigSpConfigSpCertificate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InboundSamlConfigSpConfigSpCertificate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InboundSamlConfigSpConfigSpCertificate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 x509_certificate: Optional[str] = None):
        """
        :param str x509_certificate: -
               The x509 certificate
        """
        if x509_certificate is not None:
            pulumi.set(__self__, "x509_certificate", x509_certificate)

    @property
    @pulumi.getter(name="x509Certificate")
    def x509_certificate(self) -> Optional[str]:
        """
        -
        The x509 certificate
        """
        return pulumi.get(self, "x509_certificate")


@pulumi.output_type
class TenantInboundSamlConfigIdpConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "idpCertificates":
            suggest = "idp_certificates"
        elif key == "idpEntityId":
            suggest = "idp_entity_id"
        elif key == "ssoUrl":
            suggest = "sso_url"
        elif key == "signRequest":
            suggest = "sign_request"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TenantInboundSamlConfigIdpConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TenantInboundSamlConfigIdpConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TenantInboundSamlConfigIdpConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 idp_certificates: Sequence['outputs.TenantInboundSamlConfigIdpConfigIdpCertificate'],
                 idp_entity_id: str,
                 sso_url: str,
                 sign_request: Optional[bool] = None):
        """
        :param Sequence['TenantInboundSamlConfigIdpConfigIdpCertificateArgs'] idp_certificates: The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
               Structure is documented below.
        :param str idp_entity_id: Unique identifier for all SAML entities
        :param str sso_url: URL to send Authentication request to.
        :param bool sign_request: Indicates if outbounding SAMLRequest should be signed.
        """
        pulumi.set(__self__, "idp_certificates", idp_certificates)
        pulumi.set(__self__, "idp_entity_id", idp_entity_id)
        pulumi.set(__self__, "sso_url", sso_url)
        if sign_request is not None:
            pulumi.set(__self__, "sign_request", sign_request)

    @property
    @pulumi.getter(name="idpCertificates")
    def idp_certificates(self) -> Sequence['outputs.TenantInboundSamlConfigIdpConfigIdpCertificate']:
        """
        The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
        Structure is documented below.
        """
        return pulumi.get(self, "idp_certificates")

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> str:
        """
        Unique identifier for all SAML entities
        """
        return pulumi.get(self, "idp_entity_id")

    @property
    @pulumi.getter(name="ssoUrl")
    def sso_url(self) -> str:
        """
        URL to send Authentication request to.
        """
        return pulumi.get(self, "sso_url")

    @property
    @pulumi.getter(name="signRequest")
    def sign_request(self) -> Optional[bool]:
        """
        Indicates if outbounding SAMLRequest should be signed.
        """
        return pulumi.get(self, "sign_request")


@pulumi.output_type
class TenantInboundSamlConfigIdpConfigIdpCertificate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "x509Certificate":
            suggest = "x509_certificate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TenantInboundSamlConfigIdpConfigIdpCertificate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TenantInboundSamlConfigIdpConfigIdpCertificate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TenantInboundSamlConfigIdpConfigIdpCertificate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 x509_certificate: Optional[str] = None):
        """
        :param str x509_certificate: -
               The x509 certificate
        """
        if x509_certificate is not None:
            pulumi.set(__self__, "x509_certificate", x509_certificate)

    @property
    @pulumi.getter(name="x509Certificate")
    def x509_certificate(self) -> Optional[str]:
        """
        -
        The x509 certificate
        """
        return pulumi.get(self, "x509_certificate")


@pulumi.output_type
class TenantInboundSamlConfigSpConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "callbackUri":
            suggest = "callback_uri"
        elif key == "spEntityId":
            suggest = "sp_entity_id"
        elif key == "spCertificates":
            suggest = "sp_certificates"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TenantInboundSamlConfigSpConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TenantInboundSamlConfigSpConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TenantInboundSamlConfigSpConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 callback_uri: str,
                 sp_entity_id: str,
                 sp_certificates: Optional[Sequence['outputs.TenantInboundSamlConfigSpConfigSpCertificate']] = None):
        """
        :param str callback_uri: Callback URI where responses from IDP are handled. Must start with `https://`.
        :param str sp_entity_id: Unique identifier for all SAML entities.
        :param Sequence['TenantInboundSamlConfigSpConfigSpCertificateArgs'] sp_certificates: -
               The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
               Structure is documented below.
        """
        pulumi.set(__self__, "callback_uri", callback_uri)
        pulumi.set(__self__, "sp_entity_id", sp_entity_id)
        if sp_certificates is not None:
            pulumi.set(__self__, "sp_certificates", sp_certificates)

    @property
    @pulumi.getter(name="callbackUri")
    def callback_uri(self) -> str:
        """
        Callback URI where responses from IDP are handled. Must start with `https://`.
        """
        return pulumi.get(self, "callback_uri")

    @property
    @pulumi.getter(name="spEntityId")
    def sp_entity_id(self) -> str:
        """
        Unique identifier for all SAML entities.
        """
        return pulumi.get(self, "sp_entity_id")

    @property
    @pulumi.getter(name="spCertificates")
    def sp_certificates(self) -> Optional[Sequence['outputs.TenantInboundSamlConfigSpConfigSpCertificate']]:
        """
        -
        The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP.
        Structure is documented below.
        """
        return pulumi.get(self, "sp_certificates")


@pulumi.output_type
class TenantInboundSamlConfigSpConfigSpCertificate(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "x509Certificate":
            suggest = "x509_certificate"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TenantInboundSamlConfigSpConfigSpCertificate. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TenantInboundSamlConfigSpConfigSpCertificate.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TenantInboundSamlConfigSpConfigSpCertificate.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 x509_certificate: Optional[str] = None):
        """
        :param str x509_certificate: -
               The x509 certificate
        """
        if x509_certificate is not None:
            pulumi.set(__self__, "x509_certificate", x509_certificate)

    @property
    @pulumi.getter(name="x509Certificate")
    def x509_certificate(self) -> Optional[str]:
        """
        -
        The x509 certificate
        """
        return pulumi.get(self, "x509_certificate")


