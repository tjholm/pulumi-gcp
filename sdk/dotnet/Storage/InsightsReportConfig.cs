// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// Represents an inventory report configuration.
    /// 
    /// To get more information about ReportConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/reportConfig)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/storage/docs/insights/using-storage-insights)
    /// 
    /// ## Example Usage
    /// ### Storage Insights Report Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    ///     var reportBucket = new Gcp.Storage.Bucket("reportBucket", new()
    ///     {
    ///         Location = "us-central1",
    ///         ForceDestroy = true,
    ///         UniformBucketLevelAccess = true,
    ///     });
    /// 
    ///     var config = new Gcp.Storage.InsightsReportConfig("config", new()
    ///     {
    ///         DisplayName = "Test Report Config",
    ///         Location = "us-central1",
    ///         FrequencyOptions = new Gcp.Storage.Inputs.InsightsReportConfigFrequencyOptionsArgs
    ///         {
    ///             Frequency = "WEEKLY",
    ///             StartDate = new Gcp.Storage.Inputs.InsightsReportConfigFrequencyOptionsStartDateArgs
    ///             {
    ///                 Day = 15,
    ///                 Month = 3,
    ///                 Year = 2050,
    ///             },
    ///             EndDate = new Gcp.Storage.Inputs.InsightsReportConfigFrequencyOptionsEndDateArgs
    ///             {
    ///                 Day = 15,
    ///                 Month = 4,
    ///                 Year = 2050,
    ///             },
    ///         },
    ///         CsvOptions = new Gcp.Storage.Inputs.InsightsReportConfigCsvOptionsArgs
    ///         {
    ///             RecordSeparator = @"
    /// ",
    ///             Delimiter = ",",
    ///             HeaderRequired = false,
    ///         },
    ///         ObjectMetadataReportOptions = new Gcp.Storage.Inputs.InsightsReportConfigObjectMetadataReportOptionsArgs
    ///         {
    ///             MetadataFields = new[]
    ///             {
    ///                 "bucket",
    ///                 "name",
    ///                 "project",
    ///             },
    ///             StorageFilters = new Gcp.Storage.Inputs.InsightsReportConfigObjectMetadataReportOptionsStorageFiltersArgs
    ///             {
    ///                 Bucket = reportBucket.Name,
    ///             },
    ///             StorageDestinationOptions = new Gcp.Storage.Inputs.InsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsArgs
    ///             {
    ///                 Bucket = reportBucket.Name,
    ///                 DestinationPath = "test-insights-reports",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var admin = new Gcp.Storage.BucketIAMMember("admin", new()
    ///     {
    ///         Bucket = reportBucket.Name,
    ///         Role = "roles/storage.admin",
    ///         Member = $"serviceAccount:service-{project.Apply(getProjectResult =&gt; getProjectResult.Number)}@gcp-sa-storageinsights.iam.gserviceaccount.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ReportConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/insightsReportConfig:InsightsReportConfig default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:storage/insightsReportConfig:InsightsReportConfig")]
    public partial class InsightsReportConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Options for configuring the format of the inventory report CSV file.
        /// Structure is documented below.
        /// </summary>
        [Output("csvOptions")]
        public Output<Outputs.InsightsReportConfigCsvOptions> CsvOptions { get; private set; } = null!;

        /// <summary>
        /// The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Options for configuring how inventory reports are generated.
        /// Structure is documented below.
        /// </summary>
        [Output("frequencyOptions")]
        public Output<Outputs.InsightsReportConfigFrequencyOptions?> FrequencyOptions { get; private set; } = null!;

        /// <summary>
        /// The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
        /// must be in the same location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The UUID of the inventory report configuration.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Options for including metadata in an inventory report.
        /// Structure is documented below.
        /// </summary>
        [Output("objectMetadataReportOptions")]
        public Output<Outputs.InsightsReportConfigObjectMetadataReportOptions?> ObjectMetadataReportOptions { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a InsightsReportConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InsightsReportConfig(string name, InsightsReportConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/insightsReportConfig:InsightsReportConfig", name, args ?? new InsightsReportConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InsightsReportConfig(string name, Input<string> id, InsightsReportConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/insightsReportConfig:InsightsReportConfig", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InsightsReportConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InsightsReportConfig Get(string name, Input<string> id, InsightsReportConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new InsightsReportConfig(name, id, state, options);
        }
    }

    public sealed class InsightsReportConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options for configuring the format of the inventory report CSV file.
        /// Structure is documented below.
        /// </summary>
        [Input("csvOptions", required: true)]
        public Input<Inputs.InsightsReportConfigCsvOptionsArgs> CsvOptions { get; set; } = null!;

        /// <summary>
        /// The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Options for configuring how inventory reports are generated.
        /// Structure is documented below.
        /// </summary>
        [Input("frequencyOptions")]
        public Input<Inputs.InsightsReportConfigFrequencyOptionsArgs>? FrequencyOptions { get; set; }

        /// <summary>
        /// The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
        /// must be in the same location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Options for including metadata in an inventory report.
        /// Structure is documented below.
        /// </summary>
        [Input("objectMetadataReportOptions")]
        public Input<Inputs.InsightsReportConfigObjectMetadataReportOptionsArgs>? ObjectMetadataReportOptions { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public InsightsReportConfigArgs()
        {
        }
        public static new InsightsReportConfigArgs Empty => new InsightsReportConfigArgs();
    }

    public sealed class InsightsReportConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Options for configuring the format of the inventory report CSV file.
        /// Structure is documented below.
        /// </summary>
        [Input("csvOptions")]
        public Input<Inputs.InsightsReportConfigCsvOptionsGetArgs>? CsvOptions { get; set; }

        /// <summary>
        /// The editable display name of the inventory report configuration. Has a limit of 256 characters. Can be empty.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Options for configuring how inventory reports are generated.
        /// Structure is documented below.
        /// </summary>
        [Input("frequencyOptions")]
        public Input<Inputs.InsightsReportConfigFrequencyOptionsGetArgs>? FrequencyOptions { get; set; }

        /// <summary>
        /// The location of the ReportConfig. The source and destination buckets specified in the ReportConfig
        /// must be in the same location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The UUID of the inventory report configuration.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Options for including metadata in an inventory report.
        /// Structure is documented below.
        /// </summary>
        [Input("objectMetadataReportOptions")]
        public Input<Inputs.InsightsReportConfigObjectMetadataReportOptionsGetArgs>? ObjectMetadataReportOptions { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public InsightsReportConfigState()
        {
        }
        public static new InsightsReportConfigState Empty => new InsightsReportConfigState();
    }
}
