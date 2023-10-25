# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ProjectCustomModuleArgs', 'ProjectCustomModule']

@pulumi.input_type
class ProjectCustomModuleArgs:
    def __init__(__self__, *,
                 custom_config: pulumi.Input['ProjectCustomModuleCustomConfigArgs'],
                 display_name: pulumi.Input[str],
                 enablement_state: pulumi.Input[str],
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ProjectCustomModule resource.
        :param pulumi.Input['ProjectCustomModuleCustomConfigArgs'] custom_config: The user specified custom configuration for the module.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The display name of the Security Health Analytics custom module. This
               display name becomes the finding category for all findings that are
               returned by this custom module. The display name must be between 1 and
               128 characters, start with a lowercase letter, and contain alphanumeric
               characters or underscores only.
        :param pulumi.Input[str] enablement_state: The enablement state of the custom module.
               Possible values are: `ENABLED`, `DISABLED`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ProjectCustomModuleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            custom_config=custom_config,
            display_name=display_name,
            enablement_state=enablement_state,
            project=project,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             custom_config: Optional[pulumi.Input['ProjectCustomModuleCustomConfigArgs']] = None,
             display_name: Optional[pulumi.Input[str]] = None,
             enablement_state: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if custom_config is None and 'customConfig' in kwargs:
            custom_config = kwargs['customConfig']
        if custom_config is None:
            raise TypeError("Missing 'custom_config' argument")
        if display_name is None and 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if display_name is None:
            raise TypeError("Missing 'display_name' argument")
        if enablement_state is None and 'enablementState' in kwargs:
            enablement_state = kwargs['enablementState']
        if enablement_state is None:
            raise TypeError("Missing 'enablement_state' argument")

        _setter("custom_config", custom_config)
        _setter("display_name", display_name)
        _setter("enablement_state", enablement_state)
        if project is not None:
            _setter("project", project)

    @property
    @pulumi.getter(name="customConfig")
    def custom_config(self) -> pulumi.Input['ProjectCustomModuleCustomConfigArgs']:
        """
        The user specified custom configuration for the module.
        Structure is documented below.
        """
        return pulumi.get(self, "custom_config")

    @custom_config.setter
    def custom_config(self, value: pulumi.Input['ProjectCustomModuleCustomConfigArgs']):
        pulumi.set(self, "custom_config", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the Security Health Analytics custom module. This
        display name becomes the finding category for all findings that are
        returned by this custom module. The display name must be between 1 and
        128 characters, start with a lowercase letter, and contain alphanumeric
        characters or underscores only.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="enablementState")
    def enablement_state(self) -> pulumi.Input[str]:
        """
        The enablement state of the custom module.
        Possible values are: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "enablement_state")

    @enablement_state.setter
    def enablement_state(self, value: pulumi.Input[str]):
        pulumi.set(self, "enablement_state", value)

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


@pulumi.input_type
class _ProjectCustomModuleState:
    def __init__(__self__, *,
                 ancestor_module: Optional[pulumi.Input[str]] = None,
                 custom_config: Optional[pulumi.Input['ProjectCustomModuleCustomConfigArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enablement_state: Optional[pulumi.Input[str]] = None,
                 last_editor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ProjectCustomModule resources.
        :param pulumi.Input[str] ancestor_module: If empty, indicates that the custom module was created in the organization,folder,
               or project in which you are viewing the custom module. Otherwise, ancestor_module
               specifies the organization or folder from which the custom module is inherited.
        :param pulumi.Input['ProjectCustomModuleCustomConfigArgs'] custom_config: The user specified custom configuration for the module.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The display name of the Security Health Analytics custom module. This
               display name becomes the finding category for all findings that are
               returned by this custom module. The display name must be between 1 and
               128 characters, start with a lowercase letter, and contain alphanumeric
               characters or underscores only.
        :param pulumi.Input[str] enablement_state: The enablement state of the custom module.
               Possible values are: `ENABLED`, `DISABLED`.
        :param pulumi.Input[str] last_editor: The editor that last updated the custom module.
        :param pulumi.Input[str] name: Name of the property for the custom output.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: The time at which the custom module was last updated.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
               up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        _ProjectCustomModuleState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            ancestor_module=ancestor_module,
            custom_config=custom_config,
            display_name=display_name,
            enablement_state=enablement_state,
            last_editor=last_editor,
            name=name,
            project=project,
            update_time=update_time,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             ancestor_module: Optional[pulumi.Input[str]] = None,
             custom_config: Optional[pulumi.Input['ProjectCustomModuleCustomConfigArgs']] = None,
             display_name: Optional[pulumi.Input[str]] = None,
             enablement_state: Optional[pulumi.Input[str]] = None,
             last_editor: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             project: Optional[pulumi.Input[str]] = None,
             update_time: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if ancestor_module is None and 'ancestorModule' in kwargs:
            ancestor_module = kwargs['ancestorModule']
        if custom_config is None and 'customConfig' in kwargs:
            custom_config = kwargs['customConfig']
        if display_name is None and 'displayName' in kwargs:
            display_name = kwargs['displayName']
        if enablement_state is None and 'enablementState' in kwargs:
            enablement_state = kwargs['enablementState']
        if last_editor is None and 'lastEditor' in kwargs:
            last_editor = kwargs['lastEditor']
        if update_time is None and 'updateTime' in kwargs:
            update_time = kwargs['updateTime']

        if ancestor_module is not None:
            _setter("ancestor_module", ancestor_module)
        if custom_config is not None:
            _setter("custom_config", custom_config)
        if display_name is not None:
            _setter("display_name", display_name)
        if enablement_state is not None:
            _setter("enablement_state", enablement_state)
        if last_editor is not None:
            _setter("last_editor", last_editor)
        if name is not None:
            _setter("name", name)
        if project is not None:
            _setter("project", project)
        if update_time is not None:
            _setter("update_time", update_time)

    @property
    @pulumi.getter(name="ancestorModule")
    def ancestor_module(self) -> Optional[pulumi.Input[str]]:
        """
        If empty, indicates that the custom module was created in the organization,folder,
        or project in which you are viewing the custom module. Otherwise, ancestor_module
        specifies the organization or folder from which the custom module is inherited.
        """
        return pulumi.get(self, "ancestor_module")

    @ancestor_module.setter
    def ancestor_module(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ancestor_module", value)

    @property
    @pulumi.getter(name="customConfig")
    def custom_config(self) -> Optional[pulumi.Input['ProjectCustomModuleCustomConfigArgs']]:
        """
        The user specified custom configuration for the module.
        Structure is documented below.
        """
        return pulumi.get(self, "custom_config")

    @custom_config.setter
    def custom_config(self, value: Optional[pulumi.Input['ProjectCustomModuleCustomConfigArgs']]):
        pulumi.set(self, "custom_config", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the Security Health Analytics custom module. This
        display name becomes the finding category for all findings that are
        returned by this custom module. The display name must be between 1 and
        128 characters, start with a lowercase letter, and contain alphanumeric
        characters or underscores only.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="enablementState")
    def enablement_state(self) -> Optional[pulumi.Input[str]]:
        """
        The enablement state of the custom module.
        Possible values are: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "enablement_state")

    @enablement_state.setter
    def enablement_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enablement_state", value)

    @property
    @pulumi.getter(name="lastEditor")
    def last_editor(self) -> Optional[pulumi.Input[str]]:
        """
        The editor that last updated the custom module.
        """
        return pulumi.get(self, "last_editor")

    @last_editor.setter
    def last_editor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_editor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the property for the custom output.
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
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the custom module was last updated.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
        up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class ProjectCustomModule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_config: Optional[pulumi.Input[pulumi.InputType['ProjectCustomModuleCustomConfigArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enablement_state: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents an instance of a Security Health Analytics custom module, including
        its full module name, display name, enablement state, and last updated time.
        You can create a custom module at the organization, folder, or project level.
        Custom modules that you create at the organization or folder level are inherited
        by the child folders and projects.

        To get more information about ProjectCustomModule, see:

        * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/projects.securityHealthAnalyticsSettings.customModules)
        * How-to Guides
            * [Overview of custom modules for Security Health Analytics](https://cloud.google.com/security-command-center/docs/custom-modules-sha-overview)

        ## Example Usage
        ### Scc Project Custom Module Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.securitycenter.ProjectCustomModule("example",
            custom_config=gcp.securitycenter.ProjectCustomModuleCustomConfigArgs(
                description="The rotation period of the identified cryptokey resource exceeds 30 days.",
                predicate=gcp.securitycenter.ProjectCustomModuleCustomConfigPredicateArgs(
                    expression="resource.rotationPeriod > duration(\\"2592000s\\")",
                ),
                recommendation="Set the rotation period to at most 30 days.",
                resource_selector=gcp.securitycenter.ProjectCustomModuleCustomConfigResourceSelectorArgs(
                    resource_types=["cloudkms.googleapis.com/CryptoKey"],
                ),
                severity="MEDIUM",
            ),
            display_name="basic_custom_module",
            enablement_state="ENABLED")
        ```
        ### Scc Project Custom Module Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.securitycenter.ProjectCustomModule("example",
            custom_config=gcp.securitycenter.ProjectCustomModuleCustomConfigArgs(
                custom_output=gcp.securitycenter.ProjectCustomModuleCustomConfigCustomOutputArgs(
                    properties=[gcp.securitycenter.ProjectCustomModuleCustomConfigCustomOutputPropertyArgs(
                        name="duration",
                        value_expression=gcp.securitycenter.ProjectCustomModuleCustomConfigCustomOutputPropertyValueExpressionArgs(
                            description="description of the expression",
                            expression="resource.rotationPeriod",
                            location="location of the expression",
                            title="Purpose of the expression",
                        ),
                    )],
                ),
                description="Description of the custom module",
                predicate=gcp.securitycenter.ProjectCustomModuleCustomConfigPredicateArgs(
                    description="description of the expression",
                    expression="resource.rotationPeriod > duration(\\"2592000s\\")",
                    location="location of the expression",
                    title="Purpose of the expression",
                ),
                recommendation="Steps to resolve violation",
                resource_selector=gcp.securitycenter.ProjectCustomModuleCustomConfigResourceSelectorArgs(
                    resource_types=["cloudkms.googleapis.com/CryptoKey"],
                ),
                severity="LOW",
            ),
            display_name="full_custom_module",
            enablement_state="ENABLED")
        ```

        ## Import

        ProjectCustomModule can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default projects/{{project}}/securityHealthAnalyticsSettings/customModules/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ProjectCustomModuleCustomConfigArgs']] custom_config: The user specified custom configuration for the module.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The display name of the Security Health Analytics custom module. This
               display name becomes the finding category for all findings that are
               returned by this custom module. The display name must be between 1 and
               128 characters, start with a lowercase letter, and contain alphanumeric
               characters or underscores only.
        :param pulumi.Input[str] enablement_state: The enablement state of the custom module.
               Possible values are: `ENABLED`, `DISABLED`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectCustomModuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents an instance of a Security Health Analytics custom module, including
        its full module name, display name, enablement state, and last updated time.
        You can create a custom module at the organization, folder, or project level.
        Custom modules that you create at the organization or folder level are inherited
        by the child folders and projects.

        To get more information about ProjectCustomModule, see:

        * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/projects.securityHealthAnalyticsSettings.customModules)
        * How-to Guides
            * [Overview of custom modules for Security Health Analytics](https://cloud.google.com/security-command-center/docs/custom-modules-sha-overview)

        ## Example Usage
        ### Scc Project Custom Module Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.securitycenter.ProjectCustomModule("example",
            custom_config=gcp.securitycenter.ProjectCustomModuleCustomConfigArgs(
                description="The rotation period of the identified cryptokey resource exceeds 30 days.",
                predicate=gcp.securitycenter.ProjectCustomModuleCustomConfigPredicateArgs(
                    expression="resource.rotationPeriod > duration(\\"2592000s\\")",
                ),
                recommendation="Set the rotation period to at most 30 days.",
                resource_selector=gcp.securitycenter.ProjectCustomModuleCustomConfigResourceSelectorArgs(
                    resource_types=["cloudkms.googleapis.com/CryptoKey"],
                ),
                severity="MEDIUM",
            ),
            display_name="basic_custom_module",
            enablement_state="ENABLED")
        ```
        ### Scc Project Custom Module Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        example = gcp.securitycenter.ProjectCustomModule("example",
            custom_config=gcp.securitycenter.ProjectCustomModuleCustomConfigArgs(
                custom_output=gcp.securitycenter.ProjectCustomModuleCustomConfigCustomOutputArgs(
                    properties=[gcp.securitycenter.ProjectCustomModuleCustomConfigCustomOutputPropertyArgs(
                        name="duration",
                        value_expression=gcp.securitycenter.ProjectCustomModuleCustomConfigCustomOutputPropertyValueExpressionArgs(
                            description="description of the expression",
                            expression="resource.rotationPeriod",
                            location="location of the expression",
                            title="Purpose of the expression",
                        ),
                    )],
                ),
                description="Description of the custom module",
                predicate=gcp.securitycenter.ProjectCustomModuleCustomConfigPredicateArgs(
                    description="description of the expression",
                    expression="resource.rotationPeriod > duration(\\"2592000s\\")",
                    location="location of the expression",
                    title="Purpose of the expression",
                ),
                recommendation="Steps to resolve violation",
                resource_selector=gcp.securitycenter.ProjectCustomModuleCustomConfigResourceSelectorArgs(
                    resource_types=["cloudkms.googleapis.com/CryptoKey"],
                ),
                severity="LOW",
            ),
            display_name="full_custom_module",
            enablement_state="ENABLED")
        ```

        ## Import

        ProjectCustomModule can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default projects/{{project}}/securityHealthAnalyticsSettings/customModules/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default {{project}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:securitycenter/projectCustomModule:ProjectCustomModule default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param ProjectCustomModuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectCustomModuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ProjectCustomModuleArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_config: Optional[pulumi.Input[pulumi.InputType['ProjectCustomModuleCustomConfigArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enablement_state: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectCustomModuleArgs.__new__(ProjectCustomModuleArgs)

            custom_config = _utilities.configure(custom_config, ProjectCustomModuleCustomConfigArgs, True)
            if custom_config is None and not opts.urn:
                raise TypeError("Missing required property 'custom_config'")
            __props__.__dict__["custom_config"] = custom_config
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            if enablement_state is None and not opts.urn:
                raise TypeError("Missing required property 'enablement_state'")
            __props__.__dict__["enablement_state"] = enablement_state
            __props__.__dict__["project"] = project
            __props__.__dict__["ancestor_module"] = None
            __props__.__dict__["last_editor"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        super(ProjectCustomModule, __self__).__init__(
            'gcp:securitycenter/projectCustomModule:ProjectCustomModule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ancestor_module: Optional[pulumi.Input[str]] = None,
            custom_config: Optional[pulumi.Input[pulumi.InputType['ProjectCustomModuleCustomConfigArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            enablement_state: Optional[pulumi.Input[str]] = None,
            last_editor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'ProjectCustomModule':
        """
        Get an existing ProjectCustomModule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ancestor_module: If empty, indicates that the custom module was created in the organization,folder,
               or project in which you are viewing the custom module. Otherwise, ancestor_module
               specifies the organization or folder from which the custom module is inherited.
        :param pulumi.Input[pulumi.InputType['ProjectCustomModuleCustomConfigArgs']] custom_config: The user specified custom configuration for the module.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The display name of the Security Health Analytics custom module. This
               display name becomes the finding category for all findings that are
               returned by this custom module. The display name must be between 1 and
               128 characters, start with a lowercase letter, and contain alphanumeric
               characters or underscores only.
        :param pulumi.Input[str] enablement_state: The enablement state of the custom module.
               Possible values are: `ENABLED`, `DISABLED`.
        :param pulumi.Input[str] last_editor: The editor that last updated the custom module.
        :param pulumi.Input[str] name: Name of the property for the custom output.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] update_time: The time at which the custom module was last updated.
               A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
               up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectCustomModuleState.__new__(_ProjectCustomModuleState)

        __props__.__dict__["ancestor_module"] = ancestor_module
        __props__.__dict__["custom_config"] = custom_config
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["enablement_state"] = enablement_state
        __props__.__dict__["last_editor"] = last_editor
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["update_time"] = update_time
        return ProjectCustomModule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ancestorModule")
    def ancestor_module(self) -> pulumi.Output[str]:
        """
        If empty, indicates that the custom module was created in the organization,folder,
        or project in which you are viewing the custom module. Otherwise, ancestor_module
        specifies the organization or folder from which the custom module is inherited.
        """
        return pulumi.get(self, "ancestor_module")

    @property
    @pulumi.getter(name="customConfig")
    def custom_config(self) -> pulumi.Output['outputs.ProjectCustomModuleCustomConfig']:
        """
        The user specified custom configuration for the module.
        Structure is documented below.
        """
        return pulumi.get(self, "custom_config")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the Security Health Analytics custom module. This
        display name becomes the finding category for all findings that are
        returned by this custom module. The display name must be between 1 and
        128 characters, start with a lowercase letter, and contain alphanumeric
        characters or underscores only.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enablementState")
    def enablement_state(self) -> pulumi.Output[str]:
        """
        The enablement state of the custom module.
        Possible values are: `ENABLED`, `DISABLED`.
        """
        return pulumi.get(self, "enablement_state")

    @property
    @pulumi.getter(name="lastEditor")
    def last_editor(self) -> pulumi.Output[str]:
        """
        The editor that last updated the custom module.
        """
        return pulumi.get(self, "last_editor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the property for the custom output.
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
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time at which the custom module was last updated.
        A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
        up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "update_time")

