# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SslCertArgs', 'SslCert']

@pulumi.input_type
class SslCertArgs:
    def __init__(__self__, *,
                 common_name: pulumi.Input[str],
                 instance: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SslCert resource.
        :param pulumi.Input[str] common_name: The common name to be used in the certificate to identify the
               client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        pulumi.set(__self__, "common_name", common_name)
        pulumi.set(__self__, "instance", instance)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Input[str]:
        """
        The common name to be used in the certificate to identify the
        client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Input[str]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _SslCertState:
    def __init__(__self__, *,
                 cert: Optional[pulumi.Input[str]] = None,
                 cert_serial_number: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 server_ca_cert: Optional[pulumi.Input[str]] = None,
                 sha1_fingerprint: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SslCert resources.
        :param pulumi.Input[str] cert: The actual certificate data for this client certificate.
        :param pulumi.Input[str] cert_serial_number: The serial number extracted from the certificate data.
        :param pulumi.Input[str] common_name: The common name to be used in the certificate to identify the
               client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_time: The time when the certificate was created in RFC 3339 format,
               for example 2012-11-15T16:19:00.094Z.
        :param pulumi.Input[str] expiration_time: The time when the certificate expires in RFC 3339 format,
               for example 2012-11-15T16:19:00.094Z.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] private_key: The private key associated with the client certificate.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] server_ca_cert: The CA cert of the server this client cert was generated from.
        :param pulumi.Input[str] sha1_fingerprint: The SHA1 Fingerprint of the certificate.
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if cert_serial_number is not None:
            pulumi.set(__self__, "cert_serial_number", cert_serial_number)
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if expiration_time is not None:
            pulumi.set(__self__, "expiration_time", expiration_time)
        if instance is not None:
            pulumi.set(__self__, "instance", instance)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if server_ca_cert is not None:
            pulumi.set(__self__, "server_ca_cert", server_ca_cert)
        if sha1_fingerprint is not None:
            pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[str]]:
        """
        The actual certificate data for this client certificate.
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter(name="certSerialNumber")
    def cert_serial_number(self) -> Optional[pulumi.Input[str]]:
        """
        The serial number extracted from the certificate data.
        """
        return pulumi.get(self, "cert_serial_number")

    @cert_serial_number.setter
    def cert_serial_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cert_serial_number", value)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[pulumi.Input[str]]:
        """
        The common name to be used in the certificate to identify the
        client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the certificate was created in RFC 3339 format,
        for example 2012-11-15T16:19:00.094Z.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time when the certificate expires in RFC 3339 format,
        for example 2012-11-15T16:19:00.094Z.
        """
        return pulumi.get(self, "expiration_time")

    @expiration_time.setter
    def expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time", value)

    @property
    @pulumi.getter
    def instance(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @instance.setter
    def instance(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        The private key associated with the client certificate.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serverCaCert")
    def server_ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The CA cert of the server this client cert was generated from.
        """
        return pulumi.get(self, "server_ca_cert")

    @server_ca_cert.setter
    def server_ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_ca_cert", value)

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The SHA1 Fingerprint of the certificate.
        """
        return pulumi.get(self, "sha1_fingerprint")

    @sha1_fingerprint.setter
    def sha1_fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sha1_fingerprint", value)


class SslCert(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new Google SQL SSL Cert on a Google SQL Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/sslCerts).

        > **Note:** All arguments including the private key will be stored in the raw state as plain-text

        ## Example Usage

        Example creating a SQL Client Certificate.

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        main = gcp.sql.DatabaseInstance("main",
            database_version="MYSQL_5_7",
            settings=gcp.sql.DatabaseInstanceSettingsArgs(
                tier="db-f1-micro",
            ))
        client_cert = gcp.sql.SslCert("clientCert",
            common_name="client-name",
            instance=main.name)
        ```

        ## Import

        Since the contents of the certificate cannot be accessed after its creation, this resource cannot be imported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] common_name: The common name to be used in the certificate to identify the
               client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SslCertArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Google SQL SSL Cert on a Google SQL Instance. For more information, see the [official documentation](https://cloud.google.com/sql/), or the [JSON API](https://cloud.google.com/sql/docs/mysql/admin-api/v1beta4/sslCerts).

        > **Note:** All arguments including the private key will be stored in the raw state as plain-text

        ## Example Usage

        Example creating a SQL Client Certificate.

        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_random as random

        db_name_suffix = random.RandomId("dbNameSuffix", byte_length=4)
        main = gcp.sql.DatabaseInstance("main",
            database_version="MYSQL_5_7",
            settings=gcp.sql.DatabaseInstanceSettingsArgs(
                tier="db-f1-micro",
            ))
        client_cert = gcp.sql.SslCert("clientCert",
            common_name="client-name",
            instance=main.name)
        ```

        ## Import

        Since the contents of the certificate cannot be accessed after its creation, this resource cannot be imported.

        :param str resource_name: The name of the resource.
        :param SslCertArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslCertArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SslCertArgs.__new__(SslCertArgs)

            if common_name is None and not opts.urn:
                raise TypeError("Missing required property 'common_name'")
            __props__.__dict__["common_name"] = common_name
            if instance is None and not opts.urn:
                raise TypeError("Missing required property 'instance'")
            __props__.__dict__["instance"] = instance
            __props__.__dict__["project"] = project
            __props__.__dict__["cert"] = None
            __props__.__dict__["cert_serial_number"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["expiration_time"] = None
            __props__.__dict__["private_key"] = None
            __props__.__dict__["server_ca_cert"] = None
            __props__.__dict__["sha1_fingerprint"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SslCert, __self__).__init__(
            'gcp:sql/sslCert:SslCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cert: Optional[pulumi.Input[str]] = None,
            cert_serial_number: Optional[pulumi.Input[str]] = None,
            common_name: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            expiration_time: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            server_ca_cert: Optional[pulumi.Input[str]] = None,
            sha1_fingerprint: Optional[pulumi.Input[str]] = None) -> 'SslCert':
        """
        Get an existing SslCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cert: The actual certificate data for this client certificate.
        :param pulumi.Input[str] cert_serial_number: The serial number extracted from the certificate data.
        :param pulumi.Input[str] common_name: The common name to be used in the certificate to identify the
               client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_time: The time when the certificate was created in RFC 3339 format,
               for example 2012-11-15T16:19:00.094Z.
        :param pulumi.Input[str] expiration_time: The time when the certificate expires in RFC 3339 format,
               for example 2012-11-15T16:19:00.094Z.
        :param pulumi.Input[str] instance: The name of the Cloud SQL instance. Changing this
               forces a new resource to be created.
        :param pulumi.Input[str] private_key: The private key associated with the client certificate.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.
        :param pulumi.Input[str] server_ca_cert: The CA cert of the server this client cert was generated from.
        :param pulumi.Input[str] sha1_fingerprint: The SHA1 Fingerprint of the certificate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SslCertState.__new__(_SslCertState)

        __props__.__dict__["cert"] = cert
        __props__.__dict__["cert_serial_number"] = cert_serial_number
        __props__.__dict__["common_name"] = common_name
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["expiration_time"] = expiration_time
        __props__.__dict__["instance"] = instance
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["project"] = project
        __props__.__dict__["server_ca_cert"] = server_ca_cert
        __props__.__dict__["sha1_fingerprint"] = sha1_fingerprint
        return SslCert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Output[str]:
        """
        The actual certificate data for this client certificate.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="certSerialNumber")
    def cert_serial_number(self) -> pulumi.Output[str]:
        """
        The serial number extracted from the certificate data.
        """
        return pulumi.get(self, "cert_serial_number")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Output[str]:
        """
        The common name to be used in the certificate to identify the
        client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the certificate was created in RFC 3339 format,
        for example 2012-11-15T16:19:00.094Z.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[str]:
        """
        The time when the certificate expires in RFC 3339 format,
        for example 2012-11-15T16:19:00.094Z.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def instance(self) -> pulumi.Output[str]:
        """
        The name of the Cloud SQL instance. Changing this
        forces a new resource to be created.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The private key associated with the client certificate.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs. If it
        is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="serverCaCert")
    def server_ca_cert(self) -> pulumi.Output[str]:
        """
        The CA cert of the server this client cert was generated from.
        """
        return pulumi.get(self, "server_ca_cert")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> pulumi.Output[str]:
        """
        The SHA1 Fingerprint of the certificate.
        """
        return pulumi.get(self, "sha1_fingerprint")

