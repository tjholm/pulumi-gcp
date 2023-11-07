// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataPlex
{
    /// <summary>
    /// Represents a user-visible job which provides the insights for the related data source.
    /// 
    /// To get more information about Datascan, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dataplex/docs/reference/rest)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/dataplex/docs)
    /// 
    /// ## Example Usage
    /// ### Dataplex Datascan Basic Profile
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basicProfile = new Gcp.DataPlex.Datascan("basicProfile", new()
    ///     {
    ///         Data = new Gcp.DataPlex.Inputs.DatascanDataArgs
    ///         {
    ///             Resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare",
    ///         },
    ///         DataProfileSpec = null,
    ///         DataScanId = "dataprofile-basic",
    ///         ExecutionSpec = new Gcp.DataPlex.Inputs.DatascanExecutionSpecArgs
    ///         {
    ///             Trigger = new Gcp.DataPlex.Inputs.DatascanExecutionSpecTriggerArgs
    ///             {
    ///                 OnDemand = null,
    ///             },
    ///         },
    ///         Location = "us-central1",
    ///         Project = "my-project-name",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Dataplex Datascan Full Profile
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var source = new Gcp.BigQuery.Dataset("source", new()
    ///     {
    ///         DatasetId = "dataplex_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///         DeleteContentsOnDestroy = true,
    ///     });
    /// 
    ///     var fullProfile = new Gcp.DataPlex.Datascan("fullProfile", new()
    ///     {
    ///         Location = "us-central1",
    ///         DisplayName = "Full Datascan Profile",
    ///         DataScanId = "dataprofile-full",
    ///         Description = "Example resource - Full Datascan Profile",
    ///         Labels = 
    ///         {
    ///             { "author", "billing" },
    ///         },
    ///         Data = new Gcp.DataPlex.Inputs.DatascanDataArgs
    ///         {
    ///             Resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare",
    ///         },
    ///         ExecutionSpec = new Gcp.DataPlex.Inputs.DatascanExecutionSpecArgs
    ///         {
    ///             Trigger = new Gcp.DataPlex.Inputs.DatascanExecutionSpecTriggerArgs
    ///             {
    ///                 Schedule = new Gcp.DataPlex.Inputs.DatascanExecutionSpecTriggerScheduleArgs
    ///                 {
    ///                     Cron = "TZ=America/New_York 1 1 * * *",
    ///                 },
    ///             },
    ///         },
    ///         DataProfileSpec = new Gcp.DataPlex.Inputs.DatascanDataProfileSpecArgs
    ///         {
    ///             SamplingPercent = 80,
    ///             RowFilter = "word_count &gt; 10",
    ///             IncludeFields = new Gcp.DataPlex.Inputs.DatascanDataProfileSpecIncludeFieldsArgs
    ///             {
    ///                 FieldNames = new[]
    ///                 {
    ///                     "word_count",
    ///                 },
    ///             },
    ///             ExcludeFields = new Gcp.DataPlex.Inputs.DatascanDataProfileSpecExcludeFieldsArgs
    ///             {
    ///                 FieldNames = new[]
    ///                 {
    ///                     "property_type",
    ///                 },
    ///             },
    ///             PostScanActions = new Gcp.DataPlex.Inputs.DatascanDataProfileSpecPostScanActionsArgs
    ///             {
    ///                 BigqueryExport = new Gcp.DataPlex.Inputs.DatascanDataProfileSpecPostScanActionsBigqueryExportArgs
    ///                 {
    ///                     ResultsTable = "//bigquery.googleapis.com/projects/my-project-name/datasets/dataplex_dataset/tables/profile_export",
    ///                 },
    ///             },
    ///         },
    ///         Project = "my-project-name",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             source,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Dataplex Datascan Basic Quality
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basicQuality = new Gcp.DataPlex.Datascan("basicQuality", new()
    ///     {
    ///         Data = new Gcp.DataPlex.Inputs.DatascanDataArgs
    ///         {
    ///             Resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare",
    ///         },
    ///         DataQualitySpec = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecArgs
    ///         {
    ///             Rules = new[]
    ///             {
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Description = "rule 1 for validity dimension",
    ///                     Dimension = "VALIDITY",
    ///                     Name = "rule1",
    ///                     TableConditionExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleTableConditionExpectationArgs
    ///                     {
    ///                         SqlExpression = "COUNT(*) &gt; 0",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         DataScanId = "dataquality-basic",
    ///         ExecutionSpec = new Gcp.DataPlex.Inputs.DatascanExecutionSpecArgs
    ///         {
    ///             Trigger = new Gcp.DataPlex.Inputs.DatascanExecutionSpecTriggerArgs
    ///             {
    ///                 OnDemand = null,
    ///             },
    ///         },
    ///         Location = "us-central1",
    ///         Project = "my-project-name",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Dataplex Datascan Full Quality
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fullQuality = new Gcp.DataPlex.Datascan("fullQuality", new()
    ///     {
    ///         Data = new Gcp.DataPlex.Inputs.DatascanDataArgs
    ///         {
    ///             Resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/austin_bikeshare/tables/bikeshare_stations",
    ///         },
    ///         DataQualitySpec = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecArgs
    ///         {
    ///             RowFilter = "station_id &gt; 1000",
    ///             Rules = new[]
    ///             {
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "address",
    ///                     Dimension = "VALIDITY",
    ///                     NonNullExpectation = null,
    ///                     Threshold = 0.99,
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "council_district",
    ///                     Dimension = "VALIDITY",
    ///                     IgnoreNull = true,
    ///                     RangeExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleRangeExpectationArgs
    ///                     {
    ///                         MaxValue = "10",
    ///                         MinValue = "1",
    ///                         StrictMaxEnabled = false,
    ///                         StrictMinEnabled = true,
    ///                     },
    ///                     Threshold = 0.9,
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "power_type",
    ///                     Dimension = "VALIDITY",
    ///                     IgnoreNull = false,
    ///                     RegexExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleRegexExpectationArgs
    ///                     {
    ///                         Regex = ".*solar.*",
    ///                     },
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "property_type",
    ///                     Dimension = "VALIDITY",
    ///                     IgnoreNull = false,
    ///                     SetExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleSetExpectationArgs
    ///                     {
    ///                         Values = new[]
    ///                         {
    ///                             "sidewalk",
    ///                             "parkland",
    ///                         },
    ///                     },
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "address",
    ///                     Dimension = "UNIQUENESS",
    ///                     UniquenessExpectation = null,
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "number_of_docks",
    ///                     Dimension = "VALIDITY",
    ///                     StatisticRangeExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleStatisticRangeExpectationArgs
    ///                     {
    ///                         MaxValue = "15",
    ///                         MinValue = "5",
    ///                         Statistic = "MEAN",
    ///                         StrictMaxEnabled = true,
    ///                         StrictMinEnabled = true,
    ///                     },
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Column = "footprint_length",
    ///                     Dimension = "VALIDITY",
    ///                     RowConditionExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleRowConditionExpectationArgs
    ///                     {
    ///                         SqlExpression = "footprint_length &gt; 0 AND footprint_length &lt;= 10",
    ///                     },
    ///                 },
    ///                 new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleArgs
    ///                 {
    ///                     Dimension = "VALIDITY",
    ///                     TableConditionExpectation = new Gcp.DataPlex.Inputs.DatascanDataQualitySpecRuleTableConditionExpectationArgs
    ///                     {
    ///                         SqlExpression = "COUNT(*) &gt; 0",
    ///                     },
    ///                 },
    ///             },
    ///             SamplingPercent = 5,
    ///         },
    ///         DataScanId = "dataquality-full",
    ///         Description = "Example resource - Full Datascan Quality",
    ///         DisplayName = "Full Datascan Quality",
    ///         ExecutionSpec = new Gcp.DataPlex.Inputs.DatascanExecutionSpecArgs
    ///         {
    ///             Field = "modified_date",
    ///             Trigger = new Gcp.DataPlex.Inputs.DatascanExecutionSpecTriggerArgs
    ///             {
    ///                 Schedule = new Gcp.DataPlex.Inputs.DatascanExecutionSpecTriggerScheduleArgs
    ///                 {
    ///                     Cron = "TZ=America/New_York 1 1 * * *",
    ///                 },
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "author", "billing" },
    ///         },
    ///         Location = "us-central1",
    ///         Project = "my-project-name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Datascan can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/datascan:Datascan default projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/datascan:Datascan default {{project}}/{{location}}/{{data_scan_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/datascan:Datascan default {{location}}/{{data_scan_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/datascan:Datascan default {{data_scan_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataplex/datascan:Datascan")]
    public partial class Datascan : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the scan was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The data source for DataScan.
        /// Structure is documented below.
        /// </summary>
        [Output("data")]
        public Output<Outputs.DatascanData> Data { get; private set; } = null!;

        /// <summary>
        /// DataProfileScan related setting.
        /// Structure is documented below.
        /// </summary>
        [Output("dataProfileSpec")]
        public Output<Outputs.DatascanDataProfileSpec?> DataProfileSpec { get; private set; } = null!;

        /// <summary>
        /// DataQualityScan related setting.
        /// Structure is documented below.
        /// </summary>
        [Output("dataQualitySpec")]
        public Output<Outputs.DatascanDataQualitySpec?> DataQualitySpec { get; private set; } = null!;

        /// <summary>
        /// DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
        /// </summary>
        [Output("dataScanId")]
        public Output<string> DataScanId { get; private set; } = null!;

        /// <summary>
        /// Description of the rule.
        /// The maximum length is 1,024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// DataScan execution settings.
        /// Structure is documented below.
        /// </summary>
        [Output("executionSpec")]
        public Output<Outputs.DatascanExecutionSpec> ExecutionSpec { get; private set; } = null!;

        /// <summary>
        /// Status of the data scan execution.
        /// Structure is documented below.
        /// </summary>
        [Output("executionStatuses")]
        public Output<ImmutableArray<Outputs.DatascanExecutionStatus>> ExecutionStatuses { get; private set; } = null!;

        /// <summary>
        /// User-defined labels for the scan. A list of key-&gt;value pairs.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location where the data scan should reside.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A mutable name for the rule.
        /// The name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).
        /// The maximum length is 63 characters.
        /// Must start with a letter.
        /// Must end with a number or a letter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Current state of the DataScan.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of DataScan.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the scan was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Datascan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Datascan(string name, DatascanArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataplex/datascan:Datascan", name, args ?? new DatascanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Datascan(string name, Input<string> id, DatascanState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataplex/datascan:Datascan", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Datascan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Datascan Get(string name, Input<string> id, DatascanState? state = null, CustomResourceOptions? options = null)
        {
            return new Datascan(name, id, state, options);
        }
    }

    public sealed class DatascanArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data source for DataScan.
        /// Structure is documented below.
        /// </summary>
        [Input("data", required: true)]
        public Input<Inputs.DatascanDataArgs> Data { get; set; } = null!;

        /// <summary>
        /// DataProfileScan related setting.
        /// Structure is documented below.
        /// </summary>
        [Input("dataProfileSpec")]
        public Input<Inputs.DatascanDataProfileSpecArgs>? DataProfileSpec { get; set; }

        /// <summary>
        /// DataQualityScan related setting.
        /// Structure is documented below.
        /// </summary>
        [Input("dataQualitySpec")]
        public Input<Inputs.DatascanDataQualitySpecArgs>? DataQualitySpec { get; set; }

        /// <summary>
        /// DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
        /// </summary>
        [Input("dataScanId", required: true)]
        public Input<string> DataScanId { get; set; } = null!;

        /// <summary>
        /// Description of the rule.
        /// The maximum length is 1,024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// DataScan execution settings.
        /// Structure is documented below.
        /// </summary>
        [Input("executionSpec", required: true)]
        public Input<Inputs.DatascanExecutionSpecArgs> ExecutionSpec { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the scan. A list of key-&gt;value pairs.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location where the data scan should reside.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatascanArgs()
        {
        }
        public static new DatascanArgs Empty => new DatascanArgs();
    }

    public sealed class DatascanState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the scan was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The data source for DataScan.
        /// Structure is documented below.
        /// </summary>
        [Input("data")]
        public Input<Inputs.DatascanDataGetArgs>? Data { get; set; }

        /// <summary>
        /// DataProfileScan related setting.
        /// Structure is documented below.
        /// </summary>
        [Input("dataProfileSpec")]
        public Input<Inputs.DatascanDataProfileSpecGetArgs>? DataProfileSpec { get; set; }

        /// <summary>
        /// DataQualityScan related setting.
        /// Structure is documented below.
        /// </summary>
        [Input("dataQualitySpec")]
        public Input<Inputs.DatascanDataQualitySpecGetArgs>? DataQualitySpec { get; set; }

        /// <summary>
        /// DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
        /// </summary>
        [Input("dataScanId")]
        public Input<string>? DataScanId { get; set; }

        /// <summary>
        /// Description of the rule.
        /// The maximum length is 1,024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// DataScan execution settings.
        /// Structure is documented below.
        /// </summary>
        [Input("executionSpec")]
        public Input<Inputs.DatascanExecutionSpecGetArgs>? ExecutionSpec { get; set; }

        [Input("executionStatuses")]
        private InputList<Inputs.DatascanExecutionStatusGetArgs>? _executionStatuses;

        /// <summary>
        /// Status of the data scan execution.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.DatascanExecutionStatusGetArgs> ExecutionStatuses
        {
            get => _executionStatuses ?? (_executionStatuses = new InputList<Inputs.DatascanExecutionStatusGetArgs>());
            set => _executionStatuses = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels for the scan. A list of key-&gt;value pairs.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location where the data scan should reside.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A mutable name for the rule.
        /// The name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).
        /// The maximum length is 63 characters.
        /// Must start with a letter.
        /// Must end with a number or a letter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Current state of the DataScan.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of DataScan.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The time when the scan was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public DatascanState()
        {
        }
        public static new DatascanState Empty => new DatascanState();
    }
}
