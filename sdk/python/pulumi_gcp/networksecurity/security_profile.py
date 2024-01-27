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

__all__ = ['SecurityProfileArgs', 'SecurityProfile']

@pulumi.input_type
class SecurityProfileArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 threat_prevention_profile: Optional[pulumi.Input['SecurityProfileThreatPreventionProfileArgs']] = None):
        """
        The set of arguments for constructing a SecurityProfile resource.
        :param pulumi.Input[str] type: The type of security profile.
               Possible values are: `THREAT_PREVENTION`.
        :param pulumi.Input[str] description: An optional description of the security profile. The Max length is 512 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value label pairs to assign to the resource.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] location: The location of the security profile.
               The default value is `global`.
        :param pulumi.Input[str] name: The name of the security profile resource.
               
               
               - - -
        :param pulumi.Input[str] parent: The name of the parent this security profile belongs to.
               Format: organizations/{organization_id}.
        :param pulumi.Input['SecurityProfileThreatPreventionProfileArgs'] threat_prevention_profile: The threat prevention configuration for the security profile.
               Structure is documented below.
        """
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if threat_prevention_profile is not None:
            pulumi.set(__self__, "threat_prevention_profile", threat_prevention_profile)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of security profile.
        Possible values are: `THREAT_PREVENTION`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the security profile. The Max length is 512 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key/value label pairs to assign to the resource.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the security profile.
        The default value is `global`.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security profile resource.


        - - -
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parent this security profile belongs to.
        Format: organizations/{organization_id}.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="threatPreventionProfile")
    def threat_prevention_profile(self) -> Optional[pulumi.Input['SecurityProfileThreatPreventionProfileArgs']]:
        """
        The threat prevention configuration for the security profile.
        Structure is documented below.
        """
        return pulumi.get(self, "threat_prevention_profile")

    @threat_prevention_profile.setter
    def threat_prevention_profile(self, value: Optional[pulumi.Input['SecurityProfileThreatPreventionProfileArgs']]):
        pulumi.set(self, "threat_prevention_profile", value)


@pulumi.input_type
class _SecurityProfileState:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 threat_prevention_profile: Optional[pulumi.Input['SecurityProfileThreatPreventionProfileArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityProfile resources.
        :param pulumi.Input[str] create_time: Time the security profile was created in UTC.
        :param pulumi.Input[str] description: An optional description of the security profile. The Max length is 512 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] etag: This checksum is computed by the server based on the value of other fields,
               and may be sent on update and delete requests to ensure the client has an up-to-date
               value before proceeding.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value label pairs to assign to the resource.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] location: The location of the security profile.
               The default value is `global`.
        :param pulumi.Input[str] name: The name of the security profile resource.
               
               
               - - -
        :param pulumi.Input[str] parent: The name of the parent this security profile belongs to.
               Format: organizations/{organization_id}.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] self_link: Server-defined URL of this resource.
        :param pulumi.Input['SecurityProfileThreatPreventionProfileArgs'] threat_prevention_profile: The threat prevention configuration for the security profile.
               Structure is documented below.
        :param pulumi.Input[str] type: The type of security profile.
               Possible values are: `THREAT_PREVENTION`.
        :param pulumi.Input[str] update_time: Time the security profile was updated in UTC.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if effective_labels is not None:
            pulumi.set(__self__, "effective_labels", effective_labels)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent is not None:
            pulumi.set(__self__, "parent", parent)
        if pulumi_labels is not None:
            pulumi.set(__self__, "pulumi_labels", pulumi_labels)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if threat_prevention_profile is not None:
            pulumi.set(__self__, "threat_prevention_profile", threat_prevention_profile)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the security profile was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the security profile. The Max length is 512 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @effective_labels.setter
    def effective_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "effective_labels", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        This checksum is computed by the server based on the value of other fields,
        and may be sent on update and delete requests to ensure the client has an up-to-date
        value before proceeding.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of key/value label pairs to assign to the resource.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the security profile.
        The default value is `global`.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the security profile resource.


        - - -
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parent(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the parent this security profile belongs to.
        Format: organizations/{organization_id}.
        """
        return pulumi.get(self, "parent")

    @parent.setter
    def parent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent", value)

    @property
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @pulumi_labels.setter
    def pulumi_labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "pulumi_labels", value)

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        Server-defined URL of this resource.
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="threatPreventionProfile")
    def threat_prevention_profile(self) -> Optional[pulumi.Input['SecurityProfileThreatPreventionProfileArgs']]:
        """
        The threat prevention configuration for the security profile.
        Structure is documented below.
        """
        return pulumi.get(self, "threat_prevention_profile")

    @threat_prevention_profile.setter
    def threat_prevention_profile(self, value: Optional[pulumi.Input['SecurityProfileThreatPreventionProfileArgs']]):
        pulumi.set(self, "threat_prevention_profile", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of security profile.
        Possible values are: `THREAT_PREVENTION`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the security profile was updated in UTC.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class SecurityProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 threat_prevention_profile: Optional[pulumi.Input[pulumi.InputType['SecurityProfileThreatPreventionProfileArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Network Security Security Profile Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networksecurity.SecurityProfile("default",
            parent="organizations/123456789",
            description="my description",
            type="THREAT_PREVENTION",
            labels={
                "foo": "bar",
            },
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Network Security Security Profile Overrides

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networksecurity.SecurityProfile("default",
            parent="organizations/123456789",
            description="my description",
            type="THREAT_PREVENTION",
            threat_prevention_profile=gcp.networksecurity.SecurityProfileThreatPreventionProfileArgs(
                severity_overrides=[
                    gcp.networksecurity.SecurityProfileThreatPreventionProfileSeverityOverrideArgs(
                        action="ALLOW",
                        severity="INFORMATIONAL",
                    ),
                    gcp.networksecurity.SecurityProfileThreatPreventionProfileSeverityOverrideArgs(
                        action="DENY",
                        severity="HIGH",
                    ),
                ],
                threat_overrides=[gcp.networksecurity.SecurityProfileThreatPreventionProfileThreatOverrideArgs(
                    action="ALLOW",
                    threat_id="280647",
                )],
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        SecurityProfile can be imported using any of these accepted formats* `{{parent}}/locations/{{location}}/securityProfiles/{{name}}` When using the `pulumi import` command, SecurityProfile can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:networksecurity/securityProfile:SecurityProfile default {{parent}}/locations/{{location}}/securityProfiles/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of the security profile. The Max length is 512 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value label pairs to assign to the resource.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] location: The location of the security profile.
               The default value is `global`.
        :param pulumi.Input[str] name: The name of the security profile resource.
               
               
               - - -
        :param pulumi.Input[str] parent: The name of the parent this security profile belongs to.
               Format: organizations/{organization_id}.
        :param pulumi.Input[pulumi.InputType['SecurityProfileThreatPreventionProfileArgs']] threat_prevention_profile: The threat prevention configuration for the security profile.
               Structure is documented below.
        :param pulumi.Input[str] type: The type of security profile.
               Possible values are: `THREAT_PREVENTION`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Network Security Security Profile Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networksecurity.SecurityProfile("default",
            parent="organizations/123456789",
            description="my description",
            type="THREAT_PREVENTION",
            labels={
                "foo": "bar",
            },
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```
        ### Network Security Security Profile Overrides

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.networksecurity.SecurityProfile("default",
            parent="organizations/123456789",
            description="my description",
            type="THREAT_PREVENTION",
            threat_prevention_profile=gcp.networksecurity.SecurityProfileThreatPreventionProfileArgs(
                severity_overrides=[
                    gcp.networksecurity.SecurityProfileThreatPreventionProfileSeverityOverrideArgs(
                        action="ALLOW",
                        severity="INFORMATIONAL",
                    ),
                    gcp.networksecurity.SecurityProfileThreatPreventionProfileSeverityOverrideArgs(
                        action="DENY",
                        severity="HIGH",
                    ),
                ],
                threat_overrides=[gcp.networksecurity.SecurityProfileThreatPreventionProfileThreatOverrideArgs(
                    action="ALLOW",
                    threat_id="280647",
                )],
            ),
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        SecurityProfile can be imported using any of these accepted formats* `{{parent}}/locations/{{location}}/securityProfiles/{{name}}` When using the `pulumi import` command, SecurityProfile can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:networksecurity/securityProfile:SecurityProfile default {{parent}}/locations/{{location}}/securityProfiles/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param SecurityProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 threat_prevention_profile: Optional[pulumi.Input[pulumi.InputType['SecurityProfileThreatPreventionProfileArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityProfileArgs.__new__(SecurityProfileArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["parent"] = parent
            __props__.__dict__["threat_prevention_profile"] = threat_prevention_profile
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["effective_labels"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["pulumi_labels"] = None
            __props__.__dict__["self_link"] = None
            __props__.__dict__["update_time"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["effectiveLabels", "pulumiLabels"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecurityProfile, __self__).__init__(
            'gcp:networksecurity/securityProfile:SecurityProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            effective_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent: Optional[pulumi.Input[str]] = None,
            pulumi_labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            threat_prevention_profile: Optional[pulumi.Input[pulumi.InputType['SecurityProfileThreatPreventionProfileArgs']]] = None,
            type: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'SecurityProfile':
        """
        Get an existing SecurityProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_time: Time the security profile was created in UTC.
        :param pulumi.Input[str] description: An optional description of the security profile. The Max length is 512 characters.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] effective_labels: All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        :param pulumi.Input[str] etag: This checksum is computed by the server based on the value of other fields,
               and may be sent on update and delete requests to ensure the client has an up-to-date
               value before proceeding.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A map of key/value label pairs to assign to the resource.
               
               **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
               Please refer to the field `effective_labels` for all of the labels present on the resource.
        :param pulumi.Input[str] location: The location of the security profile.
               The default value is `global`.
        :param pulumi.Input[str] name: The name of the security profile resource.
               
               
               - - -
        :param pulumi.Input[str] parent: The name of the parent this security profile belongs to.
               Format: organizations/{organization_id}.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] pulumi_labels: The combination of labels configured directly on the resource
               and default labels configured on the provider.
        :param pulumi.Input[str] self_link: Server-defined URL of this resource.
        :param pulumi.Input[pulumi.InputType['SecurityProfileThreatPreventionProfileArgs']] threat_prevention_profile: The threat prevention configuration for the security profile.
               Structure is documented below.
        :param pulumi.Input[str] type: The type of security profile.
               Possible values are: `THREAT_PREVENTION`.
        :param pulumi.Input[str] update_time: Time the security profile was updated in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityProfileState.__new__(_SecurityProfileState)

        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["effective_labels"] = effective_labels
        __props__.__dict__["etag"] = etag
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["parent"] = parent
        __props__.__dict__["pulumi_labels"] = pulumi_labels
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["threat_prevention_profile"] = threat_prevention_profile
        __props__.__dict__["type"] = type
        __props__.__dict__["update_time"] = update_time
        return SecurityProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time the security profile was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of the security profile. The Max length is 512 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="effectiveLabels")
    def effective_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        """
        return pulumi.get(self, "effective_labels")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        This checksum is computed by the server based on the value of other fields,
        and may be sent on update and delete requests to ensure the client has an up-to-date
        value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of key/value label pairs to assign to the resource.

        **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        Please refer to the field `effective_labels` for all of the labels present on the resource.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the security profile.
        The default value is `global`.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the security profile resource.


        - - -
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the parent this security profile belongs to.
        Format: organizations/{organization_id}.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="pulumiLabels")
    def pulumi_labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The combination of labels configured directly on the resource
        and default labels configured on the provider.
        """
        return pulumi.get(self, "pulumi_labels")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        Server-defined URL of this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="threatPreventionProfile")
    def threat_prevention_profile(self) -> pulumi.Output[Optional['outputs.SecurityProfileThreatPreventionProfile']]:
        """
        The threat prevention configuration for the security profile.
        Structure is documented below.
        """
        return pulumi.get(self, "threat_prevention_profile")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of security profile.
        Possible values are: `THREAT_PREVENTION`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time the security profile was updated in UTC.
        """
        return pulumi.get(self, "update_time")

