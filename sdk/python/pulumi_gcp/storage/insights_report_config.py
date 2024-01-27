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

__all__ = ['InsightsReportConfigArgs', 'InsightsReportConfig']

@pulumi.input_type
class InsightsReportConfigArgs:
    def __init__(__self__, *,
                 csv_options: pulumi.Input['InsightsReportConfigCsvOptionsArgs'],
                 location: pulumi.Input[str],
                 display_name: Optional[pulumi.Input[str]] = None,
                 frequency_options: Optional[pulumi.Input['InsightsReportConfigFrequencyOptionsArgs']] = None,
                 object_metadata_report_options: Optional[pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InsightsReportConfig resource.
        :param pulumi.Input['InsightsReportConfigCsvOptionsArgs'] csv_options: Options for configuring the format of the inventory report CSV file.
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
               must be in the same location.
        :param pulumi.Input[str] display_name: The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        :param pulumi.Input['InsightsReportConfigFrequencyOptionsArgs'] frequency_options: Options for configuring how inventory reports are generated.
               Structure is documented below.
        :param pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs'] object_metadata_report_options: Options for including metadata in an inventory report.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "csv_options", csv_options)
        pulumi.set(__self__, "location", location)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if frequency_options is not None:
            pulumi.set(__self__, "frequency_options", frequency_options)
        if object_metadata_report_options is not None:
            pulumi.set(__self__, "object_metadata_report_options", object_metadata_report_options)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="csvOptions")
    def csv_options(self) -> pulumi.Input['InsightsReportConfigCsvOptionsArgs']:
        """
        Options for configuring the format of the inventory report CSV file.
        Structure is documented below.
        """
        return pulumi.get(self, "csv_options")

    @csv_options.setter
    def csv_options(self, value: pulumi.Input['InsightsReportConfigCsvOptionsArgs']):
        pulumi.set(self, "csv_options", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
        must be in the same location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="frequencyOptions")
    def frequency_options(self) -> Optional[pulumi.Input['InsightsReportConfigFrequencyOptionsArgs']]:
        """
        Options for configuring how inventory reports are generated.
        Structure is documented below.
        """
        return pulumi.get(self, "frequency_options")

    @frequency_options.setter
    def frequency_options(self, value: Optional[pulumi.Input['InsightsReportConfigFrequencyOptionsArgs']]):
        pulumi.set(self, "frequency_options", value)

    @property
    @pulumi.getter(name="objectMetadataReportOptions")
    def object_metadata_report_options(self) -> Optional[pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs']]:
        """
        Options for including metadata in an inventory report.
        Structure is documented below.
        """
        return pulumi.get(self, "object_metadata_report_options")

    @object_metadata_report_options.setter
    def object_metadata_report_options(self, value: Optional[pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs']]):
        pulumi.set(self, "object_metadata_report_options", value)

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
class _InsightsReportConfigState:
    def __init__(__self__, *,
                 csv_options: Optional[pulumi.Input['InsightsReportConfigCsvOptionsArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 frequency_options: Optional[pulumi.Input['InsightsReportConfigFrequencyOptionsArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_metadata_report_options: Optional[pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InsightsReportConfig resources.
        :param pulumi.Input['InsightsReportConfigCsvOptionsArgs'] csv_options: Options for configuring the format of the inventory report CSV file.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        :param pulumi.Input['InsightsReportConfigFrequencyOptionsArgs'] frequency_options: Options for configuring how inventory reports are generated.
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
               must be in the same location.
        :param pulumi.Input[str] name: The UUID of the inventory report configuration.
        :param pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs'] object_metadata_report_options: Options for including metadata in an inventory report.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        if csv_options is not None:
            pulumi.set(__self__, "csv_options", csv_options)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if frequency_options is not None:
            pulumi.set(__self__, "frequency_options", frequency_options)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_metadata_report_options is not None:
            pulumi.set(__self__, "object_metadata_report_options", object_metadata_report_options)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="csvOptions")
    def csv_options(self) -> Optional[pulumi.Input['InsightsReportConfigCsvOptionsArgs']]:
        """
        Options for configuring the format of the inventory report CSV file.
        Structure is documented below.
        """
        return pulumi.get(self, "csv_options")

    @csv_options.setter
    def csv_options(self, value: Optional[pulumi.Input['InsightsReportConfigCsvOptionsArgs']]):
        pulumi.set(self, "csv_options", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="frequencyOptions")
    def frequency_options(self) -> Optional[pulumi.Input['InsightsReportConfigFrequencyOptionsArgs']]:
        """
        Options for configuring how inventory reports are generated.
        Structure is documented below.
        """
        return pulumi.get(self, "frequency_options")

    @frequency_options.setter
    def frequency_options(self, value: Optional[pulumi.Input['InsightsReportConfigFrequencyOptionsArgs']]):
        pulumi.set(self, "frequency_options", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
        must be in the same location.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The UUID of the inventory report configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectMetadataReportOptions")
    def object_metadata_report_options(self) -> Optional[pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs']]:
        """
        Options for including metadata in an inventory report.
        Structure is documented below.
        """
        return pulumi.get(self, "object_metadata_report_options")

    @object_metadata_report_options.setter
    def object_metadata_report_options(self, value: Optional[pulumi.Input['InsightsReportConfigObjectMetadataReportOptionsArgs']]):
        pulumi.set(self, "object_metadata_report_options", value)

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


class InsightsReportConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csv_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigCsvOptionsArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 frequency_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigFrequencyOptionsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 object_metadata_report_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigObjectMetadataReportOptionsArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents an inventory report configuration.

        To get more information about ReportConfig, see:

        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/reportConfig)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/insights/using-storage-insights)

        ## Example Usage
        ### Storage Insights Report Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.get_project()
        report_bucket = gcp.storage.Bucket("reportBucket",
            location="us-central1",
            force_destroy=True,
            uniform_bucket_level_access=True)
        admin = gcp.storage.BucketIAMMember("admin",
            bucket=report_bucket.name,
            role="roles/storage.admin",
            member=f"serviceAccount:service-{project.number}@gcp-sa-storageinsights.iam.gserviceaccount.com")
        config = gcp.storage.InsightsReportConfig("config",
            display_name="Test Report Config",
            location="us-central1",
            frequency_options=gcp.storage.InsightsReportConfigFrequencyOptionsArgs(
                frequency="WEEKLY",
                start_date=gcp.storage.InsightsReportConfigFrequencyOptionsStartDateArgs(
                    day=15,
                    month=3,
                    year=2050,
                ),
                end_date=gcp.storage.InsightsReportConfigFrequencyOptionsEndDateArgs(
                    day=15,
                    month=4,
                    year=2050,
                ),
            ),
            csv_options=gcp.storage.InsightsReportConfigCsvOptionsArgs(
                record_separator="\\n",
                delimiter=",",
                header_required=False,
            ),
            object_metadata_report_options=gcp.storage.InsightsReportConfigObjectMetadataReportOptionsArgs(
                metadata_fields=[
                    "bucket",
                    "name",
                    "project",
                ],
                storage_filters=gcp.storage.InsightsReportConfigObjectMetadataReportOptionsStorageFiltersArgs(
                    bucket=report_bucket.name,
                ),
                storage_destination_options=gcp.storage.InsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsArgs(
                    bucket=report_bucket.name,
                    destination_path="test-insights-reports",
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[admin]))
        ```

        ## Import

        ReportConfig can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, ReportConfig can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}
        ```

        ```sh
         $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InsightsReportConfigCsvOptionsArgs']] csv_options: Options for configuring the format of the inventory report CSV file.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        :param pulumi.Input[pulumi.InputType['InsightsReportConfigFrequencyOptionsArgs']] frequency_options: Options for configuring how inventory reports are generated.
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
               must be in the same location.
        :param pulumi.Input[pulumi.InputType['InsightsReportConfigObjectMetadataReportOptionsArgs']] object_metadata_report_options: Options for including metadata in an inventory report.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InsightsReportConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents an inventory report configuration.

        To get more information about ReportConfig, see:

        * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/reportConfig)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/storage/docs/insights/using-storage-insights)

        ## Example Usage
        ### Storage Insights Report Config

        ```python
        import pulumi
        import pulumi_gcp as gcp

        project = gcp.organizations.get_project()
        report_bucket = gcp.storage.Bucket("reportBucket",
            location="us-central1",
            force_destroy=True,
            uniform_bucket_level_access=True)
        admin = gcp.storage.BucketIAMMember("admin",
            bucket=report_bucket.name,
            role="roles/storage.admin",
            member=f"serviceAccount:service-{project.number}@gcp-sa-storageinsights.iam.gserviceaccount.com")
        config = gcp.storage.InsightsReportConfig("config",
            display_name="Test Report Config",
            location="us-central1",
            frequency_options=gcp.storage.InsightsReportConfigFrequencyOptionsArgs(
                frequency="WEEKLY",
                start_date=gcp.storage.InsightsReportConfigFrequencyOptionsStartDateArgs(
                    day=15,
                    month=3,
                    year=2050,
                ),
                end_date=gcp.storage.InsightsReportConfigFrequencyOptionsEndDateArgs(
                    day=15,
                    month=4,
                    year=2050,
                ),
            ),
            csv_options=gcp.storage.InsightsReportConfigCsvOptionsArgs(
                record_separator="\\n",
                delimiter=",",
                header_required=False,
            ),
            object_metadata_report_options=gcp.storage.InsightsReportConfigObjectMetadataReportOptionsArgs(
                metadata_fields=[
                    "bucket",
                    "name",
                    "project",
                ],
                storage_filters=gcp.storage.InsightsReportConfigObjectMetadataReportOptionsStorageFiltersArgs(
                    bucket=report_bucket.name,
                ),
                storage_destination_options=gcp.storage.InsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsArgs(
                    bucket=report_bucket.name,
                    destination_path="test-insights-reports",
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[admin]))
        ```

        ## Import

        ReportConfig can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, ReportConfig can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}
        ```

        ```sh
         $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default {{project}}/{{location}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default {{location}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param InsightsReportConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InsightsReportConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 csv_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigCsvOptionsArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 frequency_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigFrequencyOptionsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 object_metadata_report_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigObjectMetadataReportOptionsArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InsightsReportConfigArgs.__new__(InsightsReportConfigArgs)

            if csv_options is None and not opts.urn:
                raise TypeError("Missing required property 'csv_options'")
            __props__.__dict__["csv_options"] = csv_options
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["frequency_options"] = frequency_options
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["object_metadata_report_options"] = object_metadata_report_options
            __props__.__dict__["project"] = project
            __props__.__dict__["name"] = None
        super(InsightsReportConfig, __self__).__init__(
            'gcp:storage/insightsReportConfig:InsightsReportConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            csv_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigCsvOptionsArgs']]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            frequency_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigFrequencyOptionsArgs']]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object_metadata_report_options: Optional[pulumi.Input[pulumi.InputType['InsightsReportConfigObjectMetadataReportOptionsArgs']]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'InsightsReportConfig':
        """
        Get an existing InsightsReportConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['InsightsReportConfigCsvOptionsArgs']] csv_options: Options for configuring the format of the inventory report CSV file.
               Structure is documented below.
        :param pulumi.Input[str] display_name: The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        :param pulumi.Input[pulumi.InputType['InsightsReportConfigFrequencyOptionsArgs']] frequency_options: Options for configuring how inventory reports are generated.
               Structure is documented below.
        :param pulumi.Input[str] location: The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
               must be in the same location.
        :param pulumi.Input[str] name: The UUID of the inventory report configuration.
        :param pulumi.Input[pulumi.InputType['InsightsReportConfigObjectMetadataReportOptionsArgs']] object_metadata_report_options: Options for including metadata in an inventory report.
               Structure is documented below.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InsightsReportConfigState.__new__(_InsightsReportConfigState)

        __props__.__dict__["csv_options"] = csv_options
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["frequency_options"] = frequency_options
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["object_metadata_report_options"] = object_metadata_report_options
        __props__.__dict__["project"] = project
        return InsightsReportConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="csvOptions")
    def csv_options(self) -> pulumi.Output['outputs.InsightsReportConfigCsvOptions']:
        """
        Options for configuring the format of the inventory report CSV file.
        Structure is documented below.
        """
        return pulumi.get(self, "csv_options")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="frequencyOptions")
    def frequency_options(self) -> pulumi.Output[Optional['outputs.InsightsReportConfigFrequencyOptions']]:
        """
        Options for configuring how inventory reports are generated.
        Structure is documented below.
        """
        return pulumi.get(self, "frequency_options")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
        must be in the same location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The UUID of the inventory report configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectMetadataReportOptions")
    def object_metadata_report_options(self) -> pulumi.Output[Optional['outputs.InsightsReportConfigObjectMetadataReportOptions']]:
        """
        Options for including metadata in an inventory report.
        Structure is documented below.
        """
        return pulumi.get(self, "object_metadata_report_options")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

