// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/fhir/STU3/)
    /// standard for Healthcare information exchange
    /// 
    /// To get more information about FhirStore, see:
    /// 
    /// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.fhirStores)
    /// * How-to Guides
    ///     * [Creating a FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir)
    /// 
    /// ## Example Usage
    /// ### Healthcare Fhir Store Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var topic = new Gcp.PubSub.Topic("topic");
    /// 
    ///     var dataset = new Gcp.Healthcare.Dataset("dataset", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var @default = new Gcp.Healthcare.FhirStore("default", new()
    ///     {
    ///         Dataset = dataset.Id,
    ///         Version = "R4",
    ///         ComplexDataTypeReferenceParsing = "DISABLED",
    ///         EnableUpdateCreate = false,
    ///         DisableReferentialIntegrity = false,
    ///         DisableResourceVersioning = false,
    ///         EnableHistoryImport = false,
    ///         DefaultSearchHandlingStrict = false,
    ///         NotificationConfig = new Gcp.Healthcare.Inputs.FhirStoreNotificationConfigArgs
    ///         {
    ///             PubsubTopic = topic.Id,
    ///         },
    ///         Labels = 
    ///         {
    ///             { "label1", "labelvalue1" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Healthcare Fhir Store Streaming Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dataset = new Gcp.Healthcare.Dataset("dataset", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var bqDataset = new Gcp.BigQuery.Dataset("bqDataset", new()
    ///     {
    ///         DatasetId = "bq_example_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///         DeleteContentsOnDestroy = true,
    ///     });
    /// 
    ///     var @default = new Gcp.Healthcare.FhirStore("default", new()
    ///     {
    ///         Dataset = dataset.Id,
    ///         Version = "R4",
    ///         EnableUpdateCreate = false,
    ///         DisableReferentialIntegrity = false,
    ///         DisableResourceVersioning = false,
    ///         EnableHistoryImport = false,
    ///         Labels = 
    ///         {
    ///             { "label1", "labelvalue1" },
    ///         },
    ///         StreamConfigs = new[]
    ///         {
    ///             new Gcp.Healthcare.Inputs.FhirStoreStreamConfigArgs
    ///             {
    ///                 ResourceTypes = new[]
    ///                 {
    ///                     "Observation",
    ///                 },
    ///                 BigqueryDestination = new Gcp.Healthcare.Inputs.FhirStoreStreamConfigBigqueryDestinationArgs
    ///                 {
    ///                     DatasetUri = Output.Tuple(bqDataset.Project, bqDataset.DatasetId).Apply(values =&gt;
    ///                     {
    ///                         var project = values.Item1;
    ///                         var datasetId = values.Item2;
    ///                         return $"bq://{project}.{datasetId}";
    ///                     }),
    ///                     SchemaConfig = new Gcp.Healthcare.Inputs.FhirStoreStreamConfigBigqueryDestinationSchemaConfigArgs
    ///                     {
    ///                         RecursiveStructureDepth = 3,
    ///                         LastUpdatedPartitionConfig = new Gcp.Healthcare.Inputs.FhirStoreStreamConfigBigqueryDestinationSchemaConfigLastUpdatedPartitionConfigArgs
    ///                         {
    ///                             Type = "HOUR",
    ///                             ExpirationMs = "1000000",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var topic = new Gcp.PubSub.Topic("topic");
    /// 
    /// });
    /// ```
    /// ### Healthcare Fhir Store Notification Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var topic = new Gcp.PubSub.Topic("topic");
    /// 
    ///     var dataset = new Gcp.Healthcare.Dataset("dataset", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var @default = new Gcp.Healthcare.FhirStore("default", new()
    ///     {
    ///         Dataset = dataset.Id,
    ///         Version = "R4",
    ///         EnableUpdateCreate = false,
    ///         DisableReferentialIntegrity = false,
    ///         DisableResourceVersioning = false,
    ///         EnableHistoryImport = false,
    ///         Labels = 
    ///         {
    ///             { "label1", "labelvalue1" },
    ///         },
    ///         NotificationConfig = new Gcp.Healthcare.Inputs.FhirStoreNotificationConfigArgs
    ///         {
    ///             PubsubTopic = topic.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Healthcare Fhir Store Notification Configs
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var topic = new Gcp.PubSub.Topic("topic", new()
    ///     {
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var dataset = new Gcp.Healthcare.Dataset("dataset", new()
    ///     {
    ///         Location = "us-central1",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var @default = new Gcp.Healthcare.FhirStore("default", new()
    ///     {
    ///         Dataset = dataset.Id,
    ///         Version = "R4",
    ///         EnableUpdateCreate = false,
    ///         DisableReferentialIntegrity = false,
    ///         DisableResourceVersioning = false,
    ///         EnableHistoryImport = false,
    ///         Labels = 
    ///         {
    ///             { "label1", "labelvalue1" },
    ///         },
    ///         NotificationConfigs = new[]
    ///         {
    ///             new Gcp.Healthcare.Inputs.FhirStoreNotificationConfigArgs
    ///             {
    ///                 PubsubTopic = topic.Id,
    ///                 SendFullResource = true,
    ///                 SendPreviousResourceOnDelete = true,
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FhirStore can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/fhirStores/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:healthcare/fhirStore:FhirStore")]
    public partial class FhirStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
        /// Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
        /// </summary>
        [Output("complexDataTypeReferenceParsing")]
        public Output<string> ComplexDataTypeReferenceParsing { get; private set; } = null!;

        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("dataset")]
        public Output<string> Dataset { get; private set; } = null!;

        /// <summary>
        /// If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
        /// If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
        /// The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
        /// </summary>
        [Output("defaultSearchHandlingStrict")]
        public Output<bool?> DefaultSearchHandlingStrict { get; private set; } = null!;

        /// <summary>
        /// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
        /// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
        /// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
        /// will skip referential integrity check. Consequently, operations that rely on references, such as
        /// Patient.get$everything, will not return all the results if broken references exist.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Output("disableReferentialIntegrity")]
        public Output<bool?> DisableReferentialIntegrity { get; private set; } = null!;

        /// <summary>
        /// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
        /// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
        /// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
        /// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
        /// attempts to read the historical versions.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Output("disableResourceVersioning")]
        public Output<bool?> DisableResourceVersioning { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
        /// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
        /// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
        /// will fail with an error.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
        /// </summary>
        [Output("enableHistoryImport")]
        public Output<bool?> EnableHistoryImport { get; private set; } = null!;

        /// <summary>
        /// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
        /// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
        /// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
        /// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
        /// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
        /// notifications.
        /// </summary>
        [Output("enableUpdateCreate")]
        public Output<bool?> EnableUpdateCreate { get; private set; } = null!;

        /// <summary>
        /// User-supplied key-value pairs used to organize FHIR stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for the FhirStore.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.FhirStoreNotificationConfig?> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
        /// </summary>
        [Output("notificationConfigs")]
        public Output<ImmutableArray<Outputs.FhirStoreNotificationConfig>> NotificationConfigs { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
        /// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
        /// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
        /// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
        /// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
        /// the order of dozens of seconds) is expected before the results show up in the streaming destination.
        /// Structure is documented below.
        /// </summary>
        [Output("streamConfigs")]
        public Output<ImmutableArray<Outputs.FhirStoreStreamConfig>> StreamConfigs { get; private set; } = null!;

        /// <summary>
        /// The FHIR specification version.
        /// Default value is `STU3`.
        /// Possible values are: `DSTU2`, `STU3`, `R4`.
        /// </summary>
        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a FhirStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FhirStore(string name, FhirStoreArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/fhirStore:FhirStore", name, args ?? new FhirStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FhirStore(string name, Input<string> id, FhirStoreState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/fhirStore:FhirStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FhirStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FhirStore Get(string name, Input<string> id, FhirStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new FhirStore(name, id, state, options);
        }
    }

    public sealed class FhirStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
        /// Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
        /// </summary>
        [Input("complexDataTypeReferenceParsing")]
        public Input<string>? ComplexDataTypeReferenceParsing { get; set; }

        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        /// <summary>
        /// If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
        /// If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
        /// The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
        /// </summary>
        [Input("defaultSearchHandlingStrict")]
        public Input<bool>? DefaultSearchHandlingStrict { get; set; }

        /// <summary>
        /// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
        /// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
        /// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
        /// will skip referential integrity check. Consequently, operations that rely on references, such as
        /// Patient.get$everything, will not return all the results if broken references exist.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Input("disableReferentialIntegrity")]
        public Input<bool>? DisableReferentialIntegrity { get; set; }

        /// <summary>
        /// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
        /// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
        /// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
        /// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
        /// attempts to read the historical versions.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Input("disableResourceVersioning")]
        public Input<bool>? DisableResourceVersioning { get; set; }

        /// <summary>
        /// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
        /// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
        /// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
        /// will fail with an error.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
        /// </summary>
        [Input("enableHistoryImport")]
        public Input<bool>? EnableHistoryImport { get; set; }

        /// <summary>
        /// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
        /// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
        /// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
        /// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
        /// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
        /// notifications.
        /// </summary>
        [Input("enableUpdateCreate")]
        public Input<bool>? EnableUpdateCreate { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize FHIR stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
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
        /// The resource name for the FhirStore.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.FhirStoreNotificationConfigArgs>? NotificationConfig { get; set; }

        [Input("notificationConfigs")]
        private InputList<Inputs.FhirStoreNotificationConfigArgs>? _notificationConfigs;

        /// <summary>
        /// A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
        /// </summary>
        public InputList<Inputs.FhirStoreNotificationConfigArgs> NotificationConfigs
        {
            get => _notificationConfigs ?? (_notificationConfigs = new InputList<Inputs.FhirStoreNotificationConfigArgs>());
            set => _notificationConfigs = value;
        }

        [Input("streamConfigs")]
        private InputList<Inputs.FhirStoreStreamConfigArgs>? _streamConfigs;

        /// <summary>
        /// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
        /// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
        /// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
        /// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
        /// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
        /// the order of dozens of seconds) is expected before the results show up in the streaming destination.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FhirStoreStreamConfigArgs> StreamConfigs
        {
            get => _streamConfigs ?? (_streamConfigs = new InputList<Inputs.FhirStoreStreamConfigArgs>());
            set => _streamConfigs = value;
        }

        /// <summary>
        /// The FHIR specification version.
        /// Default value is `STU3`.
        /// Possible values are: `DSTU2`, `STU3`, `R4`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public FhirStoreArgs()
        {
        }
        public static new FhirStoreArgs Empty => new FhirStoreArgs();
    }

    public sealed class FhirStoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
        /// Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
        /// </summary>
        [Input("complexDataTypeReferenceParsing")]
        public Input<string>? ComplexDataTypeReferenceParsing { get; set; }

        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("dataset")]
        public Input<string>? Dataset { get; set; }

        /// <summary>
        /// If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
        /// If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
        /// The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
        /// </summary>
        [Input("defaultSearchHandlingStrict")]
        public Input<bool>? DefaultSearchHandlingStrict { get; set; }

        /// <summary>
        /// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
        /// creation. The default value is false, meaning that the API will enforce referential integrity and fail the
        /// requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
        /// will skip referential integrity check. Consequently, operations that rely on references, such as
        /// Patient.get$everything, will not return all the results if broken references exist.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Input("disableReferentialIntegrity")]
        public Input<bool>? DisableReferentialIntegrity { get; set; }

        /// <summary>
        /// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
        /// of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
        /// versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
        /// cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
        /// attempts to read the historical versions.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Input("disableResourceVersioning")]
        public Input<bool>? DisableResourceVersioning { get; set; }

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
        /// Whether to allow the bulk import API to accept history bundles and directly insert historical resource
        /// versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
        /// occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
        /// will fail with an error.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
        /// </summary>
        [Input("enableHistoryImport")]
        public Input<bool>? EnableHistoryImport { get; set; }

        /// <summary>
        /// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
        /// operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
        /// the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
        /// logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
        /// identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
        /// notifications.
        /// </summary>
        [Input("enableUpdateCreate")]
        public Input<bool>? EnableUpdateCreate { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize FHIR stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
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
        /// The resource name for the FhirStore.
        /// ** Changing this property may recreate the FHIR store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.FhirStoreNotificationConfigGetArgs>? NotificationConfig { get; set; }

        [Input("notificationConfigs")]
        private InputList<Inputs.FhirStoreNotificationConfigGetArgs>? _notificationConfigs;

        /// <summary>
        /// A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
        /// </summary>
        public InputList<Inputs.FhirStoreNotificationConfigGetArgs> NotificationConfigs
        {
            get => _notificationConfigs ?? (_notificationConfigs = new InputList<Inputs.FhirStoreNotificationConfigGetArgs>());
            set => _notificationConfigs = value;
        }

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
        /// The fully qualified name of this dataset
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("streamConfigs")]
        private InputList<Inputs.FhirStoreStreamConfigGetArgs>? _streamConfigs;

        /// <summary>
        /// A list of streaming configs that configure the destinations of streaming export for every resource mutation in
        /// this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
        /// resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
        /// from the list, the server stops streaming to that location. Before adding a new config, you must add the required
        /// bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
        /// the order of dozens of seconds) is expected before the results show up in the streaming destination.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FhirStoreStreamConfigGetArgs> StreamConfigs
        {
            get => _streamConfigs ?? (_streamConfigs = new InputList<Inputs.FhirStoreStreamConfigGetArgs>());
            set => _streamConfigs = value;
        }

        /// <summary>
        /// The FHIR specification version.
        /// Default value is `STU3`.
        /// Possible values are: `DSTU2`, `STU3`, `R4`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public FhirStoreState()
        {
        }
        public static new FhirStoreState Empty => new FhirStoreState();
    }
}
