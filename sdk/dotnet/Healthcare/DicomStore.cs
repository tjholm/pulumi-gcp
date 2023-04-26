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
    /// A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM
    /// (https://www.dicomstandard.org/about/) standard for Healthcare information exchange
    /// 
    /// To get more information about DicomStore, see:
    /// 
    /// * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.dicomStores)
    /// * How-to Guides
    ///     * [Creating a DICOM store](https://cloud.google.com/healthcare/docs/how-tos/dicom)
    /// 
    /// ## Example Usage
    /// ### Healthcare Dicom Store Basic
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
    ///     var @default = new Gcp.Healthcare.DicomStore("default", new()
    ///     {
    ///         Dataset = dataset.Id,
    ///         NotificationConfig = new Gcp.Healthcare.Inputs.DicomStoreNotificationConfigArgs
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
    /// ### Healthcare Dicom Store Bq Stream
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
    ///     var bqDataset = new Gcp.BigQuery.Dataset("bqDataset", new()
    ///     {
    ///         DatasetId = "dicom_bq_ds",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///         DeleteContentsOnDestroy = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var bqTable = new Gcp.BigQuery.Table("bqTable", new()
    ///     {
    ///         DeletionProtection = false,
    ///         DatasetId = bqDataset.DatasetId,
    ///         TableId = "dicom_bq_tb",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var @default = new Gcp.Healthcare.DicomStore("default", new()
    ///     {
    ///         Dataset = dataset.Id,
    ///         NotificationConfig = new Gcp.Healthcare.Inputs.DicomStoreNotificationConfigArgs
    ///         {
    ///             PubsubTopic = topic.Id,
    ///         },
    ///         Labels = 
    ///         {
    ///             { "label1", "labelvalue1" },
    ///         },
    ///         StreamConfigs = new[]
    ///         {
    ///             new Gcp.Healthcare.Inputs.DicomStoreStreamConfigArgs
    ///             {
    ///                 BigqueryDestination = new Gcp.Healthcare.Inputs.DicomStoreStreamConfigBigqueryDestinationArgs
    ///                 {
    ///                     TableUri = Output.Tuple(bqDataset.Project, bqDataset.DatasetId, bqTable.TableId).Apply(values =&gt;
    ///                     {
    ///                         var project = values.Item1;
    ///                         var datasetId = values.Item2;
    ///                         var tableId = values.Item3;
    ///                         return $"bq://{project}.{datasetId}.{tableId}";
    ///                     }),
    ///                 },
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
    /// DicomStore can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/dicomStores/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/dicomStore:DicomStore default {{dataset}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:healthcare/dicomStore:DicomStore")]
    public partial class DicomStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Output("dataset")]
        public Output<string> Dataset { get; private set; } = null!;

        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for the DicomStore.
        /// ** Changing this property may recreate the Dicom store (removing all data) **
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Output("notificationConfig")]
        public Output<Outputs.DicomStoreNotificationConfig?> NotificationConfig { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store.
        /// streamConfigs is an array, so you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery tables in a BigQuery dataset.
        /// Structure is documented below.
        /// </summary>
        [Output("streamConfigs")]
        public Output<ImmutableArray<Outputs.DicomStoreStreamConfig>> StreamConfigs { get; private set; } = null!;


        /// <summary>
        /// Create a DicomStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DicomStore(string name, DicomStoreArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dicomStore:DicomStore", name, args ?? new DicomStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DicomStore(string name, Input<string> id, DicomStoreState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dicomStore:DicomStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DicomStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DicomStore Get(string name, Input<string> id, DicomStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new DicomStore(name, id, state, options);
        }
    }

    public sealed class DicomStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Input("dataset", required: true)]
        public Input<string> Dataset { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the DicomStore.
        /// ** Changing this property may recreate the Dicom store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.DicomStoreNotificationConfigArgs>? NotificationConfig { get; set; }

        [Input("streamConfigs")]
        private InputList<Inputs.DicomStoreStreamConfigArgs>? _streamConfigs;

        /// <summary>
        /// To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store.
        /// streamConfigs is an array, so you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery tables in a BigQuery dataset.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.DicomStoreStreamConfigArgs> StreamConfigs
        {
            get => _streamConfigs ?? (_streamConfigs = new InputList<Inputs.DicomStoreStreamConfigArgs>());
            set => _streamConfigs = value;
        }

        public DicomStoreArgs()
        {
        }
        public static new DicomStoreArgs Empty => new DicomStoreArgs();
    }

    public sealed class DicomStoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifies the dataset addressed by this request. Must be in the format
        /// 'projects/{project}/locations/{location}/datasets/{dataset}'
        /// </summary>
        [Input("dataset")]
        public Input<string>? Dataset { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-supplied key-value pairs used to organize DICOM stores.
        /// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
        /// conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
        /// Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
        /// bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
        /// No more than 64 labels can be associated with a given store.
        /// An object containing a list of "key": value pairs.
        /// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the DicomStore.
        /// ** Changing this property may recreate the Dicom store (removing all data) **
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A nested object resource
        /// Structure is documented below.
        /// </summary>
        [Input("notificationConfig")]
        public Input<Inputs.DicomStoreNotificationConfigGetArgs>? NotificationConfig { get; set; }

        /// <summary>
        /// The fully qualified name of this dataset
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        [Input("streamConfigs")]
        private InputList<Inputs.DicomStoreStreamConfigGetArgs>? _streamConfigs;

        /// <summary>
        /// To enable streaming to BigQuery, configure the streamConfigs object in your DICOM store.
        /// streamConfigs is an array, so you can specify multiple BigQuery destinations. You can stream metadata from a single DICOM store to up to five BigQuery tables in a BigQuery dataset.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.DicomStoreStreamConfigGetArgs> StreamConfigs
        {
            get => _streamConfigs ?? (_streamConfigs = new InputList<Inputs.DicomStoreStreamConfigGetArgs>());
            set => _streamConfigs = value;
        }

        public DicomStoreState()
        {
        }
        public static new DicomStoreState Empty => new DicomStoreState();
    }
}
