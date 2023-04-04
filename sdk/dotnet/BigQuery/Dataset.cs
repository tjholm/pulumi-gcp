// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// ## Example Usage
    /// ### Bigquery Dataset Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bqowner = new Gcp.ServiceAccount.Account("bqowner", new()
    ///     {
    ///         AccountId = "bqowner",
    ///     });
    /// 
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "EU",
    ///         DefaultTableExpirationMs = 3600000,
    ///         Labels = 
    ///         {
    ///             { "env", "default" },
    ///         },
    ///         Accesses = new[]
    ///         {
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "OWNER",
    ///                 UserByEmail = bqowner.Email,
    ///             },
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "READER",
    ///                 Domain = "hashicorp.com",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Dataset Cmek
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyRing = new Gcp.Kms.KeyRing("keyRing", new()
    ///     {
    ///         Location = "us",
    ///     });
    /// 
    ///     var cryptoKey = new Gcp.Kms.CryptoKey("cryptoKey", new()
    ///     {
    ///         KeyRing = keyRing.Id,
    ///     });
    /// 
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///         DefaultTableExpirationMs = 3600000,
    ///         DefaultEncryptionConfiguration = new Gcp.BigQuery.Inputs.DatasetDefaultEncryptionConfigurationArgs
    ///         {
    ///             KmsKeyName = cryptoKey.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Dataset Authorized Dataset
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bqowner = new Gcp.ServiceAccount.Account("bqowner", new()
    ///     {
    ///         AccountId = "bqowner",
    ///     });
    /// 
    ///     var @public = new Gcp.BigQuery.Dataset("public", new()
    ///     {
    ///         DatasetId = "public",
    ///         FriendlyName = "test",
    ///         Description = "This dataset is public",
    ///         Location = "EU",
    ///         DefaultTableExpirationMs = 3600000,
    ///         Labels = 
    ///         {
    ///             { "env", "default" },
    ///         },
    ///         Accesses = new[]
    ///         {
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "OWNER",
    ///                 UserByEmail = bqowner.Email,
    ///             },
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "READER",
    ///                 Domain = "hashicorp.com",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "private",
    ///         FriendlyName = "test",
    ///         Description = "This dataset is private",
    ///         Location = "EU",
    ///         DefaultTableExpirationMs = 3600000,
    ///         Labels = 
    ///         {
    ///             { "env", "default" },
    ///         },
    ///         Accesses = new[]
    ///         {
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "OWNER",
    ///                 UserByEmail = bqowner.Email,
    ///             },
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "READER",
    ///                 Domain = "hashicorp.com",
    ///             },
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Dataset = new Gcp.BigQuery.Inputs.DatasetAccessDatasetArgs
    ///                 {
    ///                     Dataset = new Gcp.BigQuery.Inputs.DatasetAccessDatasetDatasetArgs
    ///                     {
    ///                         ProjectId = @public.Project,
    ///                         DatasetId = @public.DatasetId,
    ///                     },
    ///                     TargetTypes = new[]
    ///                     {
    ///                         "VIEWS",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Dataset Authorized Routine
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var publicDataset = new Gcp.BigQuery.Dataset("publicDataset", new()
    ///     {
    ///         DatasetId = "public_dataset",
    ///         Description = "This dataset is public",
    ///     });
    /// 
    ///     var publicRoutine = new Gcp.BigQuery.Routine("publicRoutine", new()
    ///     {
    ///         DatasetId = publicDataset.DatasetId,
    ///         RoutineId = "public_routine",
    ///         RoutineType = "TABLE_VALUED_FUNCTION",
    ///         Language = "SQL",
    ///         DefinitionBody = @"SELECT 1 + value AS value
    /// ",
    ///         Arguments = new[]
    ///         {
    ///             new Gcp.BigQuery.Inputs.RoutineArgumentArgs
    ///             {
    ///                 Name = "value",
    ///                 ArgumentKind = "FIXED_TYPE",
    ///                 DataType = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["typeKind"] = "INT64",
    ///                 }),
    ///             },
    ///         },
    ///         ReturnTableType = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["columns"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["name"] = "value",
    ///                     ["type"] = new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["typeKind"] = "INT64",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var @private = new Gcp.BigQuery.Dataset("private", new()
    ///     {
    ///         DatasetId = "private_dataset",
    ///         Description = "This dataset is private",
    ///         Accesses = new[]
    ///         {
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Role = "OWNER",
    ///                 UserByEmail = "my@service-account.com",
    ///             },
    ///             new Gcp.BigQuery.Inputs.DatasetAccessArgs
    ///             {
    ///                 Routine = new Gcp.BigQuery.Inputs.DatasetAccessRoutineArgs
    ///                 {
    ///                     ProjectId = publicRoutine.Project,
    ///                     DatasetId = publicRoutine.DatasetId,
    ///                     RoutineId = publicRoutine.RoutineId,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Dataset can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/dataset:Dataset default projects/{{project}}/datasets/{{dataset_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/dataset:Dataset default {{project}}/{{dataset_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/dataset:Dataset default {{dataset_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigquery/dataset:Dataset")]
    public partial class Dataset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An array of objects that define dataset access for one or more entities.
        /// Structure is documented below.
        /// </summary>
        [Output("accesses")]
        public Output<ImmutableArray<Outputs.DatasetAccess>> Accesses { get; private set; } = null!;

        /// <summary>
        /// The time when this dataset was created, in milliseconds since the
        /// epoch.
        /// </summary>
        [Output("creationTime")]
        public Output<int> CreationTime { get; private set; } = null!;

        /// <summary>
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// </summary>
        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// Defines the default collation specification of future tables created
        /// in the dataset. If a table is created in this dataset without table-level
        /// default collation, then the table inherits the dataset default collation,
        /// which is applied to the string fields that do not have explicit collation
        /// specified. A change to this field affects only tables created afterwards,
        /// and does not alter the existing tables.
        /// The following values are supported:
        /// - 'und:ci': undetermined locale, case insensitive.
        /// - '': empty string. Default to case-sensitive behavior.
        /// </summary>
        [Output("defaultCollation")]
        public Output<string> DefaultCollation { get; private set; } = null!;

        /// <summary>
        /// The default encryption key for all tables in the dataset. Once this property is set,
        /// all newly-created partitioned tables in the dataset will have encryption key set to
        /// this value, unless table creation request (or query) overrides the key.
        /// Structure is documented below.
        /// </summary>
        [Output("defaultEncryptionConfiguration")]
        public Output<Outputs.DatasetDefaultEncryptionConfiguration?> DefaultEncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// The default partition expiration for all partitioned tables in
        /// the dataset, in milliseconds.
        /// </summary>
        [Output("defaultPartitionExpirationMs")]
        public Output<int?> DefaultPartitionExpirationMs { get; private set; } = null!;

        /// <summary>
        /// The default lifetime of all tables in the dataset, in milliseconds.
        /// The minimum value is 3600000 milliseconds (one hour).
        /// </summary>
        [Output("defaultTableExpirationMs")]
        public Output<int?> DefaultTableExpirationMs { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, delete all the tables in the
        /// dataset when destroying the resource; otherwise,
        /// destroying the resource will fail if tables are present.
        /// </summary>
        [Output("deleteContentsOnDestroy")]
        public Output<bool?> DeleteContentsOnDestroy { get; private set; } = null!;

        /// <summary>
        /// A user-friendly description of the dataset
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A hash of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A descriptive name for the dataset
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
        /// By default, this is FALSE, which means the dataset and its table names are
        /// case-sensitive. This field does not affect routine references.
        /// </summary>
        [Output("isCaseInsensitive")]
        public Output<bool> IsCaseInsensitive { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this dataset. You can use these to
        /// organize and group your datasets
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The date when this dataset or any of its tables was last modified, in
        /// milliseconds since the epoch.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<int> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The geographic location where the dataset should reside.
        /// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
        /// </summary>
        [Output("maxTimeTravelHours")]
        public Output<string> MaxTimeTravelHours { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Dataset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Dataset(string name, DatasetArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/dataset:Dataset", name, args ?? new DatasetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Dataset(string name, Input<string> id, DatasetState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/dataset:Dataset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Dataset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Dataset Get(string name, Input<string> id, DatasetState? state = null, CustomResourceOptions? options = null)
        {
            return new Dataset(name, id, state, options);
        }
    }

    public sealed class DatasetArgs : global::Pulumi.ResourceArgs
    {
        [Input("accesses")]
        private InputList<Inputs.DatasetAccessArgs>? _accesses;

        /// <summary>
        /// An array of objects that define dataset access for one or more entities.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.DatasetAccessArgs> Accesses
        {
            get => _accesses ?? (_accesses = new InputList<Inputs.DatasetAccessArgs>());
            set => _accesses = value;
        }

        /// <summary>
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// Defines the default collation specification of future tables created
        /// in the dataset. If a table is created in this dataset without table-level
        /// default collation, then the table inherits the dataset default collation,
        /// which is applied to the string fields that do not have explicit collation
        /// specified. A change to this field affects only tables created afterwards,
        /// and does not alter the existing tables.
        /// The following values are supported:
        /// - 'und:ci': undetermined locale, case insensitive.
        /// - '': empty string. Default to case-sensitive behavior.
        /// </summary>
        [Input("defaultCollation")]
        public Input<string>? DefaultCollation { get; set; }

        /// <summary>
        /// The default encryption key for all tables in the dataset. Once this property is set,
        /// all newly-created partitioned tables in the dataset will have encryption key set to
        /// this value, unless table creation request (or query) overrides the key.
        /// Structure is documented below.
        /// </summary>
        [Input("defaultEncryptionConfiguration")]
        public Input<Inputs.DatasetDefaultEncryptionConfigurationArgs>? DefaultEncryptionConfiguration { get; set; }

        /// <summary>
        /// The default partition expiration for all partitioned tables in
        /// the dataset, in milliseconds.
        /// </summary>
        [Input("defaultPartitionExpirationMs")]
        public Input<int>? DefaultPartitionExpirationMs { get; set; }

        /// <summary>
        /// The default lifetime of all tables in the dataset, in milliseconds.
        /// The minimum value is 3600000 milliseconds (one hour).
        /// </summary>
        [Input("defaultTableExpirationMs")]
        public Input<int>? DefaultTableExpirationMs { get; set; }

        /// <summary>
        /// If set to `true`, delete all the tables in the
        /// dataset when destroying the resource; otherwise,
        /// destroying the resource will fail if tables are present.
        /// </summary>
        [Input("deleteContentsOnDestroy")]
        public Input<bool>? DeleteContentsOnDestroy { get; set; }

        /// <summary>
        /// A user-friendly description of the dataset
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A descriptive name for the dataset
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
        /// By default, this is FALSE, which means the dataset and its table names are
        /// case-sensitive. This field does not affect routine references.
        /// </summary>
        [Input("isCaseInsensitive")]
        public Input<bool>? IsCaseInsensitive { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this dataset. You can use these to
        /// organize and group your datasets
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The geographic location where the dataset should reside.
        /// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
        /// </summary>
        [Input("maxTimeTravelHours")]
        public Input<string>? MaxTimeTravelHours { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatasetArgs()
        {
        }
        public static new DatasetArgs Empty => new DatasetArgs();
    }

    public sealed class DatasetState : global::Pulumi.ResourceArgs
    {
        [Input("accesses")]
        private InputList<Inputs.DatasetAccessGetArgs>? _accesses;

        /// <summary>
        /// An array of objects that define dataset access for one or more entities.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.DatasetAccessGetArgs> Accesses
        {
            get => _accesses ?? (_accesses = new InputList<Inputs.DatasetAccessGetArgs>());
            set => _accesses = value;
        }

        /// <summary>
        /// The time when this dataset was created, in milliseconds since the
        /// epoch.
        /// </summary>
        [Input("creationTime")]
        public Input<int>? CreationTime { get; set; }

        /// <summary>
        /// A unique ID for this dataset, without the project name. The ID
        /// must contain only letters (a-z, A-Z), numbers (0-9), or
        /// underscores (_). The maximum length is 1,024 characters.
        /// </summary>
        [Input("datasetId")]
        public Input<string>? DatasetId { get; set; }

        /// <summary>
        /// Defines the default collation specification of future tables created
        /// in the dataset. If a table is created in this dataset without table-level
        /// default collation, then the table inherits the dataset default collation,
        /// which is applied to the string fields that do not have explicit collation
        /// specified. A change to this field affects only tables created afterwards,
        /// and does not alter the existing tables.
        /// The following values are supported:
        /// - 'und:ci': undetermined locale, case insensitive.
        /// - '': empty string. Default to case-sensitive behavior.
        /// </summary>
        [Input("defaultCollation")]
        public Input<string>? DefaultCollation { get; set; }

        /// <summary>
        /// The default encryption key for all tables in the dataset. Once this property is set,
        /// all newly-created partitioned tables in the dataset will have encryption key set to
        /// this value, unless table creation request (or query) overrides the key.
        /// Structure is documented below.
        /// </summary>
        [Input("defaultEncryptionConfiguration")]
        public Input<Inputs.DatasetDefaultEncryptionConfigurationGetArgs>? DefaultEncryptionConfiguration { get; set; }

        /// <summary>
        /// The default partition expiration for all partitioned tables in
        /// the dataset, in milliseconds.
        /// </summary>
        [Input("defaultPartitionExpirationMs")]
        public Input<int>? DefaultPartitionExpirationMs { get; set; }

        /// <summary>
        /// The default lifetime of all tables in the dataset, in milliseconds.
        /// The minimum value is 3600000 milliseconds (one hour).
        /// </summary>
        [Input("defaultTableExpirationMs")]
        public Input<int>? DefaultTableExpirationMs { get; set; }

        /// <summary>
        /// If set to `true`, delete all the tables in the
        /// dataset when destroying the resource; otherwise,
        /// destroying the resource will fail if tables are present.
        /// </summary>
        [Input("deleteContentsOnDestroy")]
        public Input<bool>? DeleteContentsOnDestroy { get; set; }

        /// <summary>
        /// A user-friendly description of the dataset
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A hash of the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// A descriptive name for the dataset
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
        /// By default, this is FALSE, which means the dataset and its table names are
        /// case-sensitive. This field does not affect routine references.
        /// </summary>
        [Input("isCaseInsensitive")]
        public Input<bool>? IsCaseInsensitive { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this dataset. You can use these to
        /// organize and group your datasets
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The date when this dataset or any of its tables was last modified, in
        /// milliseconds since the epoch.
        /// </summary>
        [Input("lastModifiedTime")]
        public Input<int>? LastModifiedTime { get; set; }

        /// <summary>
        /// The geographic location where the dataset should reside.
        /// See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
        /// </summary>
        [Input("maxTimeTravelHours")]
        public Input<string>? MaxTimeTravelHours { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public DatasetState()
        {
        }
        public static new DatasetState Empty => new DatasetState();
    }
}
