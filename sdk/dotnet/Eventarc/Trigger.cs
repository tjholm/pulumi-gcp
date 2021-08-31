// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Eventarc
{
    /// <summary>
    /// ## Example Usage
    /// ### Basic
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.CloudRun.Service("default", new Gcp.CloudRun.ServiceArgs
    ///         {
    ///             Location = "europe-west1",
    ///             Metadata = new Gcp.CloudRun.Inputs.ServiceMetadataArgs
    ///             {
    ///                 Namespace = "my-project-name",
    ///             },
    ///             Template = new Gcp.CloudRun.Inputs.ServiceTemplateArgs
    ///             {
    ///                 Spec = new Gcp.CloudRun.Inputs.ServiceTemplateSpecArgs
    ///                 {
    ///                     Containers = 
    ///                     {
    ///                         new Gcp.CloudRun.Inputs.ServiceTemplateSpecContainerArgs
    ///                         {
    ///                             Image = "gcr.io/cloudrun/hello",
    ///                             Args = 
    ///                             {
    ///                                 "arrgs",
    ///                             },
    ///                         },
    ///                     },
    ///                     ContainerConcurrency = 50,
    ///                 },
    ///             },
    ///             Traffics = 
    ///             {
    ///                 new Gcp.CloudRun.Inputs.ServiceTrafficArgs
    ///                 {
    ///                     Percent = 100,
    ///                     LatestRevision = true,
    ///                 },
    ///             },
    ///         });
    ///         var primary = new Gcp.Eventarc.Trigger("primary", new Gcp.Eventarc.TriggerArgs
    ///         {
    ///             Location = "europe-west1",
    ///             MatchingCriterias = 
    ///             {
    ///                 new Gcp.Eventarc.Inputs.TriggerMatchingCriteriaArgs
    ///                 {
    ///                     Attribute = "type",
    ///                     Value = "google.cloud.pubsub.topic.v1.messagePublished",
    ///                 },
    ///             },
    ///             Destination = new Gcp.Eventarc.Inputs.TriggerDestinationArgs
    ///             {
    ///                 CloudRunService = new Gcp.Eventarc.Inputs.TriggerDestinationCloudRunServiceArgs
    ///                 {
    ///                     Service = @default.Name,
    ///                     Region = "europe-west1",
    ///                 },
    ///             },
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///         });
    ///         var foo = new Gcp.PubSub.Topic("foo", new Gcp.PubSub.TopicArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Trigger can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:eventarc/trigger:Trigger default projects/{{project}}/locations/{{location}}/triggers/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:eventarc/trigger:Trigger default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:eventarc/trigger:Trigger default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:eventarc/trigger:Trigger")]
    public partial class Trigger : Pulumi.CustomResource
    {
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Required. Destination specifies where the events should be sent to.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.TriggerDestination> Destination { get; private set; } = null!;

        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Optional. User labels attached to the triggers that can be used to group resources.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        /// </summary>
        [Output("matchingCriterias")]
        public Output<ImmutableArray<Outputs.TriggerMatchingCriteria>> MatchingCriterias { get; private set; } = null!;

        /// <summary>
        /// Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string?> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        /// </summary>
        [Output("transports")]
        public Output<ImmutableArray<Outputs.TriggerTransport>> Transports { get; private set; } = null!;

        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs args, CustomResourceOptions? options = null)
            : base("gcp:eventarc/trigger:Trigger", name, args ?? new TriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
            : base("gcp:eventarc/trigger:Trigger", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, TriggerState? state = null, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, state, options);
        }
    }

    public sealed class TriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Destination specifies where the events should be sent to.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.TriggerDestinationArgs> Destination { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User labels attached to the triggers that can be used to group resources.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("matchingCriterias", required: true)]
        private InputList<Inputs.TriggerMatchingCriteriaArgs>? _matchingCriterias;

        /// <summary>
        /// Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        /// </summary>
        public InputList<Inputs.TriggerMatchingCriteriaArgs> MatchingCriterias
        {
            get => _matchingCriterias ?? (_matchingCriterias = new InputList<Inputs.TriggerMatchingCriteriaArgs>());
            set => _matchingCriterias = value;
        }

        /// <summary>
        /// Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("transports")]
        private InputList<Inputs.TriggerTransportArgs>? _transports;

        /// <summary>
        /// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        /// </summary>
        public InputList<Inputs.TriggerTransportArgs> Transports
        {
            get => _transports ?? (_transports = new InputList<Inputs.TriggerTransportArgs>());
            set => _transports = value;
        }

        public TriggerArgs()
        {
        }
    }

    public sealed class TriggerState : Pulumi.ResourceArgs
    {
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Required. Destination specifies where the events should be sent to.
        /// </summary>
        [Input("destination")]
        public Input<Inputs.TriggerDestinationGetArgs>? Destination { get; set; }

        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User labels attached to the triggers that can be used to group resources.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("matchingCriterias")]
        private InputList<Inputs.TriggerMatchingCriteriaGetArgs>? _matchingCriterias;

        /// <summary>
        /// Required. The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        /// </summary>
        public InputList<Inputs.TriggerMatchingCriteriaGetArgs> MatchingCriterias
        {
            get => _matchingCriterias ?? (_matchingCriterias = new InputList<Inputs.TriggerMatchingCriteriaGetArgs>());
            set => _matchingCriterias = value;
        }

        /// <summary>
        /// Required. The resource name of the trigger. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        [Input("transports")]
        private InputList<Inputs.TriggerTransportGetArgs>? _transports;

        /// <summary>
        /// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        /// </summary>
        public InputList<Inputs.TriggerTransportGetArgs> Transports
        {
            get => _transports ?? (_transports = new InputList<Inputs.TriggerTransportGetArgs>());
            set => _transports = value;
        }

        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public TriggerState()
        {
        }
    }
}
