# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HmacKeyArgs', 'HmacKey']

@pulumi.input_type
class HmacKeyArgs:
    def __init__(__self__, *,
                 service_account_email: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HmacKey resource.
        :param pulumi.Input[str] service_account_email: The email address of the key's associated service account.
               
               
               - - -
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] state: The state of the key. Can be set to one of ACTIVE, INACTIVE.
               Default value is `ACTIVE`.
               Possible values are: `ACTIVE`, `INACTIVE`.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        """
        The email address of the key's associated service account.


        - - -
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the key. Can be set to one of ACTIVE, INACTIVE.
        Default value is `ACTIVE`.
        Possible values are: `ACTIVE`, `INACTIVE`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


@pulumi.input_type
class _HmacKeyState:
    def __init__(__self__, *,
                 access_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 time_created: Optional[pulumi.Input[str]] = None,
                 updated: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HmacKey resources.
        :param pulumi.Input[str] access_id: The access ID of the HMAC Key.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] secret: HMAC secret key material.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] service_account_email: The email address of the key's associated service account.
               
               
               - - -
        :param pulumi.Input[str] state: The state of the key. Can be set to one of ACTIVE, INACTIVE.
               Default value is `ACTIVE`.
               Possible values are: `ACTIVE`, `INACTIVE`.
        :param pulumi.Input[str] time_created: 'The creation time of the HMAC key in RFC 3339 format. '
        :param pulumi.Input[str] updated: 'The last modification time of the HMAC key metadata in RFC 3339 format.'
        """
        if access_id is not None:
            pulumi.set(__self__, "access_id", access_id)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)
        if service_account_email is not None:
            pulumi.set(__self__, "service_account_email", service_account_email)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if time_created is not None:
            pulumi.set(__self__, "time_created", time_created)
        if updated is not None:
            pulumi.set(__self__, "updated", updated)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> Optional[pulumi.Input[str]]:
        """
        The access ID of the HMAC Key.
        """
        return pulumi.get(self, "access_id")

    @access_id.setter
    def access_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        HMAC secret key material.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address of the key's associated service account.


        - - -
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the key. Can be set to one of ACTIVE, INACTIVE.
        Default value is `ACTIVE`.
        Possible values are: `ACTIVE`, `INACTIVE`.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> Optional[pulumi.Input[str]]:
        """
        'The creation time of the HMAC key in RFC 3339 format. '
        """
        return pulumi.get(self, "time_created")

    @time_created.setter
    def time_created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_created", value)

    @property
    @pulumi.getter
    def updated(self) -> Optional[pulumi.Input[str]]:
        """
        'The last modification time of the HMAC key metadata in RFC 3339 format.'
        """
        return pulumi.get(self, "updated")

    @updated.setter
    def updated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated", value)


class HmacKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
        consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
        for service accounts.

        To get more information about HmacKey, see:

        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)

        ## Example Usage
        ### Storage Hmac Key

        ```python
        import pulumi
        import pulumi_gcp as gcp

        # Create a new service account
        service_account = gcp.serviceaccount.Account("serviceAccount", account_id="my-svc-acc")
        #Create the HMAC key for the associated service account
        key = gcp.storage.HmacKey("key", service_account_email=service_account.email)
        ```

        ## Import

        HmacKey can be imported using any of these accepted formats* `projects/{{project}}/hmacKeys/{{access_id}}` * `{{project}}/{{access_id}}` * `{{access_id}}` When using the `pulumi import` command, HmacKey can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:storage/hmacKey:HmacKey default projects/{{project}}/hmacKeys/{{access_id}}
        ```

        ```sh
         $ pulumi import gcp:storage/hmacKey:HmacKey default {{project}}/{{access_id}}
        ```

        ```sh
         $ pulumi import gcp:storage/hmacKey:HmacKey default {{access_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] service_account_email: The email address of the key's associated service account.
               
               
               - - -
        :param pulumi.Input[str] state: The state of the key. Can be set to one of ACTIVE, INACTIVE.
               Default value is `ACTIVE`.
               Possible values are: `ACTIVE`, `INACTIVE`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HmacKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
        consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
        for service accounts.

        To get more information about HmacKey, see:

        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)

        ## Example Usage
        ### Storage Hmac Key

        ```python
        import pulumi
        import pulumi_gcp as gcp

        # Create a new service account
        service_account = gcp.serviceaccount.Account("serviceAccount", account_id="my-svc-acc")
        #Create the HMAC key for the associated service account
        key = gcp.storage.HmacKey("key", service_account_email=service_account.email)
        ```

        ## Import

        HmacKey can be imported using any of these accepted formats* `projects/{{project}}/hmacKeys/{{access_id}}` * `{{project}}/{{access_id}}` * `{{access_id}}` When using the `pulumi import` command, HmacKey can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:storage/hmacKey:HmacKey default projects/{{project}}/hmacKeys/{{access_id}}
        ```

        ```sh
         $ pulumi import gcp:storage/hmacKey:HmacKey default {{project}}/{{access_id}}
        ```

        ```sh
         $ pulumi import gcp:storage/hmacKey:HmacKey default {{access_id}}
        ```

        :param str resource_name: The name of the resource.
        :param HmacKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HmacKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HmacKeyArgs.__new__(HmacKeyArgs)

            __props__.__dict__["project"] = project
            if service_account_email is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_email'")
            __props__.__dict__["service_account_email"] = service_account_email
            __props__.__dict__["state"] = state
            __props__.__dict__["access_id"] = None
            __props__.__dict__["secret"] = None
            __props__.__dict__["time_created"] = None
            __props__.__dict__["updated"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(HmacKey, __self__).__init__(
            'gcp:storage/hmacKey:HmacKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_id: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            time_created: Optional[pulumi.Input[str]] = None,
            updated: Optional[pulumi.Input[str]] = None) -> 'HmacKey':
        """
        Get an existing HmacKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_id: The access ID of the HMAC Key.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] secret: HMAC secret key material.
               **Note**: This property is sensitive and will not be displayed in the plan.
        :param pulumi.Input[str] service_account_email: The email address of the key's associated service account.
               
               
               - - -
        :param pulumi.Input[str] state: The state of the key. Can be set to one of ACTIVE, INACTIVE.
               Default value is `ACTIVE`.
               Possible values are: `ACTIVE`, `INACTIVE`.
        :param pulumi.Input[str] time_created: 'The creation time of the HMAC key in RFC 3339 format. '
        :param pulumi.Input[str] updated: 'The last modification time of the HMAC key metadata in RFC 3339 format.'
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HmacKeyState.__new__(_HmacKeyState)

        __props__.__dict__["access_id"] = access_id
        __props__.__dict__["project"] = project
        __props__.__dict__["secret"] = secret
        __props__.__dict__["service_account_email"] = service_account_email
        __props__.__dict__["state"] = state
        __props__.__dict__["time_created"] = time_created
        __props__.__dict__["updated"] = updated
        return HmacKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> pulumi.Output[str]:
        """
        The access ID of the HMAC Key.
        """
        return pulumi.get(self, "access_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        HMAC secret key material.
        **Note**: This property is sensitive and will not be displayed in the plan.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        The email address of the key's associated service account.


        - - -
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of the key. Can be set to one of ACTIVE, INACTIVE.
        Default value is `ACTIVE`.
        Possible values are: `ACTIVE`, `INACTIVE`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> pulumi.Output[str]:
        """
        'The creation time of the HMAC key in RFC 3339 format. '
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[str]:
        """
        'The last modification time of the HMAC key metadata in RFC 3339 format.'
        """
        return pulumi.get(self, "updated")

