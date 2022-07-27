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

__all__ = ['SecurityScanConfigArgs', 'SecurityScanConfig']

@pulumi.input_type
class SecurityScanConfigArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 starting_urls: pulumi.Input[Sequence[pulumi.Input[str]]],
                 authentication: Optional[pulumi.Input['SecurityScanConfigAuthenticationArgs']] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 export_to_security_command_center: Optional[pulumi.Input[str]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input['SecurityScanConfigScheduleArgs']] = None,
                 target_platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecurityScanConfig resource.
        :param pulumi.Input[str] display_name: The user provider display name of the ScanConfig.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] starting_urls: The starting URLs from which the scanner finds site pages.
        :param pulumi.Input['SecurityScanConfigAuthenticationArgs'] authentication: The authentication configuration.
               If specified, service will use the authentication configuration during scanning.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_patterns: The blacklist URL patterns as described in
               https://cloud.google.com/security-scanner/docs/excluded-urls
        :param pulumi.Input[str] export_to_security_command_center: Controls export of scan configurations and results to Cloud Security Command Center.
               Default value is `ENABLED`.
               Possible values are `ENABLED` and `DISABLED`.
        :param pulumi.Input[int] max_qps: The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
               Defaults to 15.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['SecurityScanConfigScheduleArgs'] schedule: The schedule of the ScanConfig
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_platforms: Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
               Each value may be one of `APP_ENGINE` and `COMPUTE`.
        :param pulumi.Input[str] user_agent: Type of the user agents used for scanning
               Default value is `CHROME_LINUX`.
               Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "starting_urls", starting_urls)
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if blacklist_patterns is not None:
            pulumi.set(__self__, "blacklist_patterns", blacklist_patterns)
        if export_to_security_command_center is not None:
            pulumi.set(__self__, "export_to_security_command_center", export_to_security_command_center)
        if max_qps is not None:
            pulumi.set(__self__, "max_qps", max_qps)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if target_platforms is not None:
            pulumi.set(__self__, "target_platforms", target_platforms)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The user provider display name of the ScanConfig.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="startingUrls")
    def starting_urls(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The starting URLs from which the scanner finds site pages.
        """
        return pulumi.get(self, "starting_urls")

    @starting_urls.setter
    def starting_urls(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "starting_urls", value)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input['SecurityScanConfigAuthenticationArgs']]:
        """
        The authentication configuration.
        If specified, service will use the authentication configuration during scanning.
        Structure is documented below.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input['SecurityScanConfigAuthenticationArgs']]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter(name="blacklistPatterns")
    def blacklist_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The blacklist URL patterns as described in
        https://cloud.google.com/security-scanner/docs/excluded-urls
        """
        return pulumi.get(self, "blacklist_patterns")

    @blacklist_patterns.setter
    def blacklist_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blacklist_patterns", value)

    @property
    @pulumi.getter(name="exportToSecurityCommandCenter")
    def export_to_security_command_center(self) -> Optional[pulumi.Input[str]]:
        """
        Controls export of scan configurations and results to Cloud Security Command Center.
        Default value is `ENABLED`.
        Possible values are `ENABLED` and `DISABLED`.
        """
        return pulumi.get(self, "export_to_security_command_center")

    @export_to_security_command_center.setter
    def export_to_security_command_center(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "export_to_security_command_center", value)

    @property
    @pulumi.getter(name="maxQps")
    def max_qps(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
        Defaults to 15.
        """
        return pulumi.get(self, "max_qps")

    @max_qps.setter
    def max_qps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_qps", value)

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
    def schedule(self) -> Optional[pulumi.Input['SecurityScanConfigScheduleArgs']]:
        """
        The schedule of the ScanConfig
        Structure is documented below.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['SecurityScanConfigScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="targetPlatforms")
    def target_platforms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        Each value may be one of `APP_ENGINE` and `COMPUTE`.
        """
        return pulumi.get(self, "target_platforms")

    @target_platforms.setter
    def target_platforms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_platforms", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the user agents used for scanning
        Default value is `CHROME_LINUX`.
        Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)


@pulumi.input_type
class _SecurityScanConfigState:
    def __init__(__self__, *,
                 authentication: Optional[pulumi.Input['SecurityScanConfigAuthenticationArgs']] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 export_to_security_command_center: Optional[pulumi.Input[str]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input['SecurityScanConfigScheduleArgs']] = None,
                 starting_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityScanConfig resources.
        :param pulumi.Input['SecurityScanConfigAuthenticationArgs'] authentication: The authentication configuration.
               If specified, service will use the authentication configuration during scanning.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_patterns: The blacklist URL patterns as described in
               https://cloud.google.com/security-scanner/docs/excluded-urls
        :param pulumi.Input[str] display_name: The user provider display name of the ScanConfig.
        :param pulumi.Input[str] export_to_security_command_center: Controls export of scan configurations and results to Cloud Security Command Center.
               Default value is `ENABLED`.
               Possible values are `ENABLED` and `DISABLED`.
        :param pulumi.Input[int] max_qps: The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
               Defaults to 15.
        :param pulumi.Input[str] name: A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['SecurityScanConfigScheduleArgs'] schedule: The schedule of the ScanConfig
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] starting_urls: The starting URLs from which the scanner finds site pages.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_platforms: Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
               Each value may be one of `APP_ENGINE` and `COMPUTE`.
        :param pulumi.Input[str] user_agent: Type of the user agents used for scanning
               Default value is `CHROME_LINUX`.
               Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if blacklist_patterns is not None:
            pulumi.set(__self__, "blacklist_patterns", blacklist_patterns)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if export_to_security_command_center is not None:
            pulumi.set(__self__, "export_to_security_command_center", export_to_security_command_center)
        if max_qps is not None:
            pulumi.set(__self__, "max_qps", max_qps)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if starting_urls is not None:
            pulumi.set(__self__, "starting_urls", starting_urls)
        if target_platforms is not None:
            pulumi.set(__self__, "target_platforms", target_platforms)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input['SecurityScanConfigAuthenticationArgs']]:
        """
        The authentication configuration.
        If specified, service will use the authentication configuration during scanning.
        Structure is documented below.
        """
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input['SecurityScanConfigAuthenticationArgs']]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter(name="blacklistPatterns")
    def blacklist_patterns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The blacklist URL patterns as described in
        https://cloud.google.com/security-scanner/docs/excluded-urls
        """
        return pulumi.get(self, "blacklist_patterns")

    @blacklist_patterns.setter
    def blacklist_patterns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blacklist_patterns", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The user provider display name of the ScanConfig.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="exportToSecurityCommandCenter")
    def export_to_security_command_center(self) -> Optional[pulumi.Input[str]]:
        """
        Controls export of scan configurations and results to Cloud Security Command Center.
        Default value is `ENABLED`.
        Possible values are `ENABLED` and `DISABLED`.
        """
        return pulumi.get(self, "export_to_security_command_center")

    @export_to_security_command_center.setter
    def export_to_security_command_center(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "export_to_security_command_center", value)

    @property
    @pulumi.getter(name="maxQps")
    def max_qps(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
        Defaults to 15.
        """
        return pulumi.get(self, "max_qps")

    @max_qps.setter
    def max_qps(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_qps", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    def schedule(self) -> Optional[pulumi.Input['SecurityScanConfigScheduleArgs']]:
        """
        The schedule of the ScanConfig
        Structure is documented below.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input['SecurityScanConfigScheduleArgs']]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter(name="startingUrls")
    def starting_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The starting URLs from which the scanner finds site pages.
        """
        return pulumi.get(self, "starting_urls")

    @starting_urls.setter
    def starting_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "starting_urls", value)

    @property
    @pulumi.getter(name="targetPlatforms")
    def target_platforms(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        Each value may be one of `APP_ENGINE` and `COMPUTE`.
        """
        return pulumi.get(self, "target_platforms")

    @target_platforms.setter
    def target_platforms(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "target_platforms", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the user agents used for scanning
        Default value is `CHROME_LINUX`.
        Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)


class SecurityScanConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[pulumi.InputType['SecurityScanConfigAuthenticationArgs']]] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 export_to_security_command_center: Optional[pulumi.Input[str]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['SecurityScanConfigScheduleArgs']]] = None,
                 starting_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A ScanConfig resource contains the configurations to launch a scan.

        To get more information about ScanConfig, see:

        * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
        * How-to Guides
            * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)

        > **Warning:** All arguments including `authentication.google_account.password` and `authentication.custom_account.password` will be stored in the raw state as plain-text.

        ## Example Usage
        ### Scan Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        scanner_static_ip = gcp.compute.Address("scannerStaticIp", opts=pulumi.ResourceOptions(provider=google_beta))
        scan_config = gcp.compute.SecurityScanConfig("scan-config",
            display_name="scan-config",
            starting_urls=[scanner_static_ip.address.apply(lambda address: f"http://{address}")],
            target_platforms=["COMPUTE"],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        ScanConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default projects/{{project}}/scanConfigs/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SecurityScanConfigAuthenticationArgs']] authentication: The authentication configuration.
               If specified, service will use the authentication configuration during scanning.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_patterns: The blacklist URL patterns as described in
               https://cloud.google.com/security-scanner/docs/excluded-urls
        :param pulumi.Input[str] display_name: The user provider display name of the ScanConfig.
        :param pulumi.Input[str] export_to_security_command_center: Controls export of scan configurations and results to Cloud Security Command Center.
               Default value is `ENABLED`.
               Possible values are `ENABLED` and `DISABLED`.
        :param pulumi.Input[int] max_qps: The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
               Defaults to 15.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SecurityScanConfigScheduleArgs']] schedule: The schedule of the ScanConfig
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] starting_urls: The starting URLs from which the scanner finds site pages.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_platforms: Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
               Each value may be one of `APP_ENGINE` and `COMPUTE`.
        :param pulumi.Input[str] user_agent: Type of the user agents used for scanning
               Default value is `CHROME_LINUX`.
               Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityScanConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A ScanConfig resource contains the configurations to launch a scan.

        To get more information about ScanConfig, see:

        * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
        * How-to Guides
            * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)

        > **Warning:** All arguments including `authentication.google_account.password` and `authentication.custom_account.password` will be stored in the raw state as plain-text.

        ## Example Usage
        ### Scan Config Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        scanner_static_ip = gcp.compute.Address("scannerStaticIp", opts=pulumi.ResourceOptions(provider=google_beta))
        scan_config = gcp.compute.SecurityScanConfig("scan-config",
            display_name="scan-config",
            starting_urls=[scanner_static_ip.address.apply(lambda address: f"http://{address}")],
            target_platforms=["COMPUTE"],
            opts=pulumi.ResourceOptions(provider=google_beta))
        ```

        ## Import

        ScanConfig can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default projects/{{project}}/scanConfigs/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param SecurityScanConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityScanConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authentication: Optional[pulumi.Input[pulumi.InputType['SecurityScanConfigAuthenticationArgs']]] = None,
                 blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 export_to_security_command_center: Optional[pulumi.Input[str]] = None,
                 max_qps: Optional[pulumi.Input[int]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['SecurityScanConfigScheduleArgs']]] = None,
                 starting_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityScanConfigArgs.__new__(SecurityScanConfigArgs)

            __props__.__dict__["authentication"] = authentication
            __props__.__dict__["blacklist_patterns"] = blacklist_patterns
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["export_to_security_command_center"] = export_to_security_command_center
            __props__.__dict__["max_qps"] = max_qps
            __props__.__dict__["project"] = project
            __props__.__dict__["schedule"] = schedule
            if starting_urls is None and not opts.urn:
                raise TypeError("Missing required property 'starting_urls'")
            __props__.__dict__["starting_urls"] = starting_urls
            __props__.__dict__["target_platforms"] = target_platforms
            __props__.__dict__["user_agent"] = user_agent
            __props__.__dict__["name"] = None
        super(SecurityScanConfig, __self__).__init__(
            'gcp:compute/securityScanConfig:SecurityScanConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authentication: Optional[pulumi.Input[pulumi.InputType['SecurityScanConfigAuthenticationArgs']]] = None,
            blacklist_patterns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            export_to_security_command_center: Optional[pulumi.Input[str]] = None,
            max_qps: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            schedule: Optional[pulumi.Input[pulumi.InputType['SecurityScanConfigScheduleArgs']]] = None,
            starting_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_platforms: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user_agent: Optional[pulumi.Input[str]] = None) -> 'SecurityScanConfig':
        """
        Get an existing SecurityScanConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SecurityScanConfigAuthenticationArgs']] authentication: The authentication configuration.
               If specified, service will use the authentication configuration during scanning.
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] blacklist_patterns: The blacklist URL patterns as described in
               https://cloud.google.com/security-scanner/docs/excluded-urls
        :param pulumi.Input[str] display_name: The user provider display name of the ScanConfig.
        :param pulumi.Input[str] export_to_security_command_center: Controls export of scan configurations and results to Cloud Security Command Center.
               Default value is `ENABLED`.
               Possible values are `ENABLED` and `DISABLED`.
        :param pulumi.Input[int] max_qps: The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
               Defaults to 15.
        :param pulumi.Input[str] name: A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[pulumi.InputType['SecurityScanConfigScheduleArgs']] schedule: The schedule of the ScanConfig
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] starting_urls: The starting URLs from which the scanner finds site pages.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_platforms: Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
               Each value may be one of `APP_ENGINE` and `COMPUTE`.
        :param pulumi.Input[str] user_agent: Type of the user agents used for scanning
               Default value is `CHROME_LINUX`.
               Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityScanConfigState.__new__(_SecurityScanConfigState)

        __props__.__dict__["authentication"] = authentication
        __props__.__dict__["blacklist_patterns"] = blacklist_patterns
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["export_to_security_command_center"] = export_to_security_command_center
        __props__.__dict__["max_qps"] = max_qps
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["starting_urls"] = starting_urls
        __props__.__dict__["target_platforms"] = target_platforms
        __props__.__dict__["user_agent"] = user_agent
        return SecurityScanConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authentication(self) -> pulumi.Output[Optional['outputs.SecurityScanConfigAuthentication']]:
        """
        The authentication configuration.
        If specified, service will use the authentication configuration during scanning.
        Structure is documented below.
        """
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter(name="blacklistPatterns")
    def blacklist_patterns(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The blacklist URL patterns as described in
        https://cloud.google.com/security-scanner/docs/excluded-urls
        """
        return pulumi.get(self, "blacklist_patterns")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The user provider display name of the ScanConfig.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="exportToSecurityCommandCenter")
    def export_to_security_command_center(self) -> pulumi.Output[Optional[str]]:
        """
        Controls export of scan configurations and results to Cloud Security Command Center.
        Default value is `ENABLED`.
        Possible values are `ENABLED` and `DISABLED`.
        """
        return pulumi.get(self, "export_to_security_command_center")

    @property
    @pulumi.getter(name="maxQps")
    def max_qps(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
        Defaults to 15.
        """
        return pulumi.get(self, "max_qps")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A server defined name for this index. Format: 'projects/{{project}}/scanConfigs/{{server_generated_id}}'
        """
        return pulumi.get(self, "name")

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
    def schedule(self) -> pulumi.Output[Optional['outputs.SecurityScanConfigSchedule']]:
        """
        The schedule of the ScanConfig
        Structure is documented below.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter(name="startingUrls")
    def starting_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        The starting URLs from which the scanner finds site pages.
        """
        return pulumi.get(self, "starting_urls")

    @property
    @pulumi.getter(name="targetPlatforms")
    def target_platforms(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        Each value may be one of `APP_ENGINE` and `COMPUTE`.
        """
        return pulumi.get(self, "target_platforms")

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> pulumi.Output[Optional[str]]:
        """
        Type of the user agents used for scanning
        Default value is `CHROME_LINUX`.
        Possible values are `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, and `SAFARI_IPHONE`.
        """
        return pulumi.get(self, "user_agent")

