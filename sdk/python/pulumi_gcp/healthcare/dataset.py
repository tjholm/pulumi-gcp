# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DatasetArgs', 'Dataset']

@pulumi.input_type
class DatasetArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Dataset resource.
        :param pulumi.Input[str] location: The location for the Dataset.
               
               
               - - -
        :param pulumi.Input[str] name: The resource name for the Dataset.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] time_zone: The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
               "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
               (e.g., HL7 messages) where no explicit timezone is specified.
        """
        pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location for the Dataset.


        - - -
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the Dataset.
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
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
        "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
        (e.g., HL7 messages) where no explicit timezone is specified.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)


@pulumi.input_type
class _DatasetState:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 self_link: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Dataset resources.
        :param pulumi.Input[str] location: The location for the Dataset.
               
               
               - - -
        :param pulumi.Input[str] name: The resource name for the Dataset.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        :param pulumi.Input[str] time_zone: The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
               "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
               (e.g., HL7 messages) where no explicit timezone is specified.
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if self_link is not None:
            pulumi.set(__self__, "self_link", self_link)
        if time_zone is not None:
            pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location for the Dataset.


        - - -
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the Dataset.
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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> Optional[pulumi.Input[str]]:
        """
        The fully qualified name of this dataset
        """
        return pulumi.get(self, "self_link")

    @self_link.setter
    def self_link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "self_link", value)

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
        "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
        (e.g., HL7 messages) where no explicit timezone is specified.
        """
        return pulumi.get(self, "time_zone")

    @time_zone.setter
    def time_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_zone", value)


class Dataset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A Healthcare `Dataset` is a toplevel logical grouping of `dicomStores`, `fhirStores` and `hl7V2Stores`.

        To get more information about Dataset, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets)
        * How-to Guides
            * [Creating a dataset](https://cloud.google.com/healthcare/docs/how-tos/datasets)

        ## Example Usage
        ### Healthcare Dataset Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.healthcare.Dataset("default",
            location="us-central1",
            time_zone="UTC")
        ```

        ## Import

        Dataset can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/datasets/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Dataset can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:healthcare/dataset:Dataset default projects/{{project}}/locations/{{location}}/datasets/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/dataset:Dataset default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/dataset:Dataset default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location for the Dataset.
               
               
               - - -
        :param pulumi.Input[str] name: The resource name for the Dataset.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] time_zone: The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
               "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
               (e.g., HL7 messages) where no explicit timezone is specified.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatasetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A Healthcare `Dataset` is a toplevel logical grouping of `dicomStores`, `fhirStores` and `hl7V2Stores`.

        To get more information about Dataset, see:

        * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets)
        * How-to Guides
            * [Creating a dataset](https://cloud.google.com/healthcare/docs/how-tos/datasets)

        ## Example Usage
        ### Healthcare Dataset Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default = gcp.healthcare.Dataset("default",
            location="us-central1",
            time_zone="UTC")
        ```

        ## Import

        Dataset can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/datasets/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Dataset can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:healthcare/dataset:Dataset default projects/{{project}}/locations/{{location}}/datasets/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/dataset:Dataset default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:healthcare/dataset:Dataset default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param DatasetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatasetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 time_zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatasetArgs.__new__(DatasetArgs)

            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["time_zone"] = time_zone
            __props__.__dict__["self_link"] = None
        super(Dataset, __self__).__init__(
            'gcp:healthcare/dataset:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            self_link: Optional[pulumi.Input[str]] = None,
            time_zone: Optional[pulumi.Input[str]] = None) -> 'Dataset':
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: The location for the Dataset.
               
               
               - - -
        :param pulumi.Input[str] name: The resource name for the Dataset.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The fully qualified name of this dataset
        :param pulumi.Input[str] time_zone: The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
               "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
               (e.g., HL7 messages) where no explicit timezone is specified.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatasetState.__new__(_DatasetState)

        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["self_link"] = self_link
        __props__.__dict__["time_zone"] = time_zone
        return Dataset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location for the Dataset.


        - - -
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the Dataset.
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
    @pulumi.getter(name="selfLink")
    def self_link(self) -> pulumi.Output[str]:
        """
        The fully qualified name of this dataset
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> pulumi.Output[str]:
        """
        The default timezone used by this dataset. Must be a either a valid IANA time zone name such as
        "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources
        (e.g., HL7 messages) where no explicit timezone is specified.
        """
        return pulumi.get(self, "time_zone")

