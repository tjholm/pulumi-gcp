# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SshPublicKeyArgs', 'SshPublicKey']

@pulumi.input_type
class SshPublicKeyArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 user: pulumi.Input[str],
                 expiration_time_usec: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SshPublicKey resource.
        :param pulumi.Input[str] key: Public key text in SSH format, defined by RFC4253 section 6.6.
        :param pulumi.Input[str] user: The user email.
        :param pulumi.Input[str] expiration_time_usec: An expiration time in microseconds since epoch.
        :param pulumi.Input[str] project: The project ID of the Google Cloud Platform project.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "user", user)
        if expiration_time_usec is not None:
            pulumi.set(__self__, "expiration_time_usec", expiration_time_usec)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Public key text in SSH format, defined by RFC4253 section 6.6.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The user email.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter(name="expirationTimeUsec")
    def expiration_time_usec(self) -> Optional[pulumi.Input[str]]:
        """
        An expiration time in microseconds since epoch.
        """
        return pulumi.get(self, "expiration_time_usec")

    @expiration_time_usec.setter
    def expiration_time_usec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time_usec", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID of the Google Cloud Platform project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _SshPublicKeyState:
    def __init__(__self__, *,
                 expiration_time_usec: Optional[pulumi.Input[str]] = None,
                 fingerprint: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SshPublicKey resources.
        :param pulumi.Input[str] expiration_time_usec: An expiration time in microseconds since epoch.
        :param pulumi.Input[str] fingerprint: The SHA-256 fingerprint of the SSH public key.
        :param pulumi.Input[str] key: Public key text in SSH format, defined by RFC4253 section 6.6.
        :param pulumi.Input[str] project: The project ID of the Google Cloud Platform project.
        :param pulumi.Input[str] user: The user email.
        """
        if expiration_time_usec is not None:
            pulumi.set(__self__, "expiration_time_usec", expiration_time_usec)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="expirationTimeUsec")
    def expiration_time_usec(self) -> Optional[pulumi.Input[str]]:
        """
        An expiration time in microseconds since epoch.
        """
        return pulumi.get(self, "expiration_time_usec")

    @expiration_time_usec.setter
    def expiration_time_usec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time_usec", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        The SHA-256 fingerprint of the SSH public key.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Public key text in SSH format, defined by RFC4253 section 6.6.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project ID of the Google Cloud Platform project.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The user email.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class SshPublicKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_time_usec: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The SSH public key information associated with a Google account.

        To get more information about SSHPublicKey, see:

        * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest/v1/users.sshPublicKeys)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/oslogin)

        ## Example Usage
        ### Os Login Ssh Key Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        me = gcp.organizations.get_client_open_id_user_info()
        cache = gcp.oslogin.SshPublicKey("cache",
            user=me.email,
            key=(lambda path: open(path).read())("path/to/id_rsa.pub"))
        ```

        ## Import

        SSHPublicKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default users/{{user}}/sshPublicKeys/{{fingerprint}}
        ```

        ```sh
         $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default {{user}}/{{fingerprint}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_time_usec: An expiration time in microseconds since epoch.
        :param pulumi.Input[str] key: Public key text in SSH format, defined by RFC4253 section 6.6.
        :param pulumi.Input[str] project: The project ID of the Google Cloud Platform project.
        :param pulumi.Input[str] user: The user email.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SshPublicKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The SSH public key information associated with a Google account.

        To get more information about SSHPublicKey, see:

        * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest/v1/users.sshPublicKeys)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/compute/docs/oslogin)

        ## Example Usage
        ### Os Login Ssh Key Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        me = gcp.organizations.get_client_open_id_user_info()
        cache = gcp.oslogin.SshPublicKey("cache",
            user=me.email,
            key=(lambda path: open(path).read())("path/to/id_rsa.pub"))
        ```

        ## Import

        SSHPublicKey can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default users/{{user}}/sshPublicKeys/{{fingerprint}}
        ```

        ```sh
         $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default {{user}}/{{fingerprint}}
        ```

        :param str resource_name: The name of the resource.
        :param SshPublicKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SshPublicKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 expiration_time_usec: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SshPublicKeyArgs.__new__(SshPublicKeyArgs)

            __props__.__dict__["expiration_time_usec"] = expiration_time_usec
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["project"] = project
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
            __props__.__dict__["fingerprint"] = None
        super(SshPublicKey, __self__).__init__(
            'gcp:oslogin/sshPublicKey:SshPublicKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            expiration_time_usec: Optional[pulumi.Input[str]] = None,
            fingerprint: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'SshPublicKey':
        """
        Get an existing SshPublicKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] expiration_time_usec: An expiration time in microseconds since epoch.
        :param pulumi.Input[str] fingerprint: The SHA-256 fingerprint of the SSH public key.
        :param pulumi.Input[str] key: Public key text in SSH format, defined by RFC4253 section 6.6.
        :param pulumi.Input[str] project: The project ID of the Google Cloud Platform project.
        :param pulumi.Input[str] user: The user email.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SshPublicKeyState.__new__(_SshPublicKeyState)

        __props__.__dict__["expiration_time_usec"] = expiration_time_usec
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["key"] = key
        __props__.__dict__["project"] = project
        __props__.__dict__["user"] = user
        return SshPublicKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="expirationTimeUsec")
    def expiration_time_usec(self) -> pulumi.Output[Optional[str]]:
        """
        An expiration time in microseconds since epoch.
        """
        return pulumi.get(self, "expiration_time_usec")

    @property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[str]:
        """
        The SHA-256 fingerprint of the SSH public key.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        Public key text in SSH format, defined by RFC4253 section 6.6.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[Optional[str]]:
        """
        The project ID of the Google Cloud Platform project.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The user email.
        """
        return pulumi.get(self, "user")

