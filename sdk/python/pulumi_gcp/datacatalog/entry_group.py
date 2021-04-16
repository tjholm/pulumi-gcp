# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EntryGroupArgs', 'EntryGroup']

@pulumi.input_type
class EntryGroupArgs:
    def __init__(__self__, *,
                 entry_group_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EntryGroup resource.
        :param pulumi.Input[str] entry_group_id: The id of the entry group to create. The id must begin with a letter or underscore,
               contain only English letters, numbers and underscores, and be at most 64 characters.
        :param pulumi.Input[str] description: Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        :param pulumi.Input[str] display_name: A short name to identify the entry group, for example, "analytics data - jan 2011".
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: EntryGroup location region.
        """
        pulumi.set(__self__, "entry_group_id", entry_group_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="entryGroupId")
    def entry_group_id(self) -> pulumi.Input[str]:
        """
        The id of the entry group to create. The id must begin with a letter or underscore,
        contain only English letters, numbers and underscores, and be at most 64 characters.
        """
        return pulumi.get(self, "entry_group_id")

    @entry_group_id.setter
    def entry_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entry_group_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A short name to identify the entry group, for example, "analytics data - jan 2011".
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

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
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        EntryGroup location region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _EntryGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 entry_group_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EntryGroup resources.
        :param pulumi.Input[str] description: Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        :param pulumi.Input[str] display_name: A short name to identify the entry group, for example, "analytics data - jan 2011".
        :param pulumi.Input[str] entry_group_id: The id of the entry group to create. The id must begin with a letter or underscore,
               contain only English letters, numbers and underscores, and be at most 64 characters.
        :param pulumi.Input[str] name: The resource name of the entry group in URL format. Example:
               projects/{project}/locations/{location}/entryGroups/{entryGroupId}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: EntryGroup location region.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if entry_group_id is not None:
            pulumi.set(__self__, "entry_group_id", entry_group_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        A short name to identify the entry group, for example, "analytics data - jan 2011".
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="entryGroupId")
    def entry_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the entry group to create. The id must begin with a letter or underscore,
        contain only English letters, numbers and underscores, and be at most 64 characters.
        """
        return pulumi.get(self, "entry_group_id")

    @entry_group_id.setter
    def entry_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entry_group_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the entry group in URL format. Example:
        projects/{project}/locations/{location}/entryGroups/{entryGroupId}
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
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        EntryGroup location region.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class EntryGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 entry_group_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.

        To get more information about EntryGroup, see:

        * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-catalog/docs)

        ## Example Usage
        ### Data Catalog Entry Group Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_entry_group = gcp.datacatalog.EntryGroup("basicEntryGroup", entry_group_id="my_group")
        ```
        ### Data Catalog Entry Group Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_entry_group = gcp.datacatalog.EntryGroup("basicEntryGroup",
            description="example entry group",
            display_name="entry group",
            entry_group_id="my_group")
        ```

        ## Import

        EntryGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datacatalog/entryGroup:EntryGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        :param pulumi.Input[str] display_name: A short name to identify the entry group, for example, "analytics data - jan 2011".
        :param pulumi.Input[str] entry_group_id: The id of the entry group to create. The id must begin with a letter or underscore,
               contain only English letters, numbers and underscores, and be at most 64 characters.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: EntryGroup location region.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EntryGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An EntryGroup resource represents a logical grouping of zero or more Data Catalog Entry resources.

        To get more information about EntryGroup, see:

        * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.entryGroups)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/data-catalog/docs)

        ## Example Usage
        ### Data Catalog Entry Group Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_entry_group = gcp.datacatalog.EntryGroup("basicEntryGroup", entry_group_id="my_group")
        ```
        ### Data Catalog Entry Group Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        basic_entry_group = gcp.datacatalog.EntryGroup("basicEntryGroup",
            description="example entry group",
            display_name="entry group",
            entry_group_id="my_group")
        ```

        ## Import

        EntryGroup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:datacatalog/entryGroup:EntryGroup default {{name}}
        ```

        :param str resource_name: The name of the resource.
        :param EntryGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EntryGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 entry_group_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
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
            __props__ = EntryGroupArgs.__new__(EntryGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if entry_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'entry_group_id'")
            __props__.__dict__["entry_group_id"] = entry_group_id
            __props__.__dict__["project"] = project
            __props__.__dict__["region"] = region
            __props__.__dict__["name"] = None
        super(EntryGroup, __self__).__init__(
            'gcp:datacatalog/entryGroup:EntryGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            entry_group_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'EntryGroup':
        """
        Get an existing EntryGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        :param pulumi.Input[str] display_name: A short name to identify the entry group, for example, "analytics data - jan 2011".
        :param pulumi.Input[str] entry_group_id: The id of the entry group to create. The id must begin with a letter or underscore,
               contain only English letters, numbers and underscores, and be at most 64 characters.
        :param pulumi.Input[str] name: The resource name of the entry group in URL format. Example:
               projects/{project}/locations/{location}/entryGroups/{entryGroupId}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: EntryGroup location region.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EntryGroupState.__new__(_EntryGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["entry_group_id"] = entry_group_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["region"] = region
        return EntryGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        A short name to identify the entry group, for example, "analytics data - jan 2011".
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="entryGroupId")
    def entry_group_id(self) -> pulumi.Output[str]:
        """
        The id of the entry group to create. The id must begin with a letter or underscore,
        contain only English letters, numbers and underscores, and be at most 64 characters.
        """
        return pulumi.get(self, "entry_group_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the entry group in URL format. Example:
        projects/{project}/locations/{location}/entryGroups/{entryGroupId}
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
    def region(self) -> pulumi.Output[str]:
        """
        EntryGroup location region.
        """
        return pulumi.get(self, "region")

