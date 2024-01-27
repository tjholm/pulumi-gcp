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
    /// Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
    ///  For more information see
    /// [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications)
    /// and
    /// [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
    /// 
    /// In order to enable notifications, a special Google Cloud Storage service account unique to the project
    /// must exist and have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project.
    /// This service account is not created automatically when a project is created.
    /// To ensure the service account exists and obtain its email address for use in granting the correct IAM permission, use the
    /// [`gcp.storage.getProjectServiceAccount`](https://www.terraform.io/docs/providers/google/d/storage_project_service_account.html)
    /// datasource's `email_address` value, and see below for an example of enabling notifications by granting the correct IAM permission.
    /// See [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.
    /// 
    /// &gt; **NOTE**: This resource can affect your storage IAM policy. If you are using this in the same config as your storage IAM policy resources, consider
    /// making this resource dependent on those IAM resources via `depends_on`. This will safeguard against errors due to IAM race conditions.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var gcsAccount = Gcp.Storage.GetProjectServiceAccount.Invoke();
    /// 
    ///     var topic = new Gcp.PubSub.Topic("topic");
    /// 
    ///     var binding = new Gcp.PubSub.TopicIAMBinding("binding", new()
    ///     {
    ///         Topic = topic.Id,
    ///         Role = "roles/pubsub.publisher",
    ///         Members = new[]
    ///         {
    ///             $"serviceAccount:{gcsAccount.Apply(getProjectServiceAccountResult =&gt; getProjectServiceAccountResult.EmailAddress)}",
    ///         },
    ///     });
    /// 
    ///     // End enabling notifications
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "US",
    ///     });
    /// 
    ///     var notification = new Gcp.Storage.Notification("notification", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         PayloadFormat = "JSON_API_V1",
    ///         Topic = topic.Id,
    ///         EventTypes = new[]
    ///         {
    ///             "OBJECT_FINALIZE",
    ///             "OBJECT_METADATA_UPDATE",
    ///         },
    ///         CustomAttributes = 
    ///         {
    ///             { "new-attribute", "new-attribute-value" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             binding,
    ///         },
    ///     });
    /// 
    ///     // Enable notifications by giving the correct IAM permission to the unique service account.
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Storage notifications can be imported using any of these accepted formats* `{{bucket_name}}/notificationConfigs/{{id}}` When using the `pulumi import` command, Storage notifications can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:storage/notification:Notification default {{bucket_name}}/notificationConfigs/{{id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:storage/notification:Notification")]
    public partial class Notification : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        /// </summary>
        [Output("customAttributes")]
        public Output<ImmutableDictionary<string, string>?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        /// </summary>
        [Output("eventTypes")]
        public Output<ImmutableArray<string>> EventTypes { get; private set; } = null!;

        /// <summary>
        /// The ID of the created notification.
        /// </summary>
        [Output("notificationId")]
        public Output<string> NotificationId { get; private set; } = null!;

        /// <summary>
        /// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        /// </summary>
        [Output("objectNamePrefix")]
        public Output<string?> ObjectNamePrefix { get; private set; } = null!;

        /// <summary>
        /// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        /// </summary>
        [Output("payloadFormat")]
        public Output<string> PayloadFormat { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The Cloud PubSub topic to which this subscription publishes. Expects either the
        /// topic name, assumed to belong to the default GCP provider project, or the project-level name,
        /// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
        /// you will need to use the project-level name.
        /// 
        /// - - -
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;


        /// <summary>
        /// Create a Notification resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Notification(string name, NotificationArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/notification:Notification", name, args ?? new NotificationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Notification(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/notification:Notification", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Notification resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Notification Get(string name, Input<string> id, NotificationState? state = null, CustomResourceOptions? options = null)
        {
            return new Notification(name, id, state, options);
        }
    }

    public sealed class NotificationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        /// <summary>
        /// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        /// </summary>
        [Input("objectNamePrefix")]
        public Input<string>? ObjectNamePrefix { get; set; }

        /// <summary>
        /// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        /// </summary>
        [Input("payloadFormat", required: true)]
        public Input<string> PayloadFormat { get; set; } = null!;

        /// <summary>
        /// The Cloud PubSub topic to which this subscription publishes. Expects either the
        /// topic name, assumed to belong to the default GCP provider project, or the project-level name,
        /// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
        /// you will need to use the project-level name.
        /// 
        /// - - -
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public NotificationArgs()
        {
        }
        public static new NotificationArgs Empty => new NotificationArgs();
    }

    public sealed class NotificationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("customAttributes")]
        private InputMap<string>? _customAttributes;

        /// <summary>
        /// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
        /// </summary>
        public InputMap<string> CustomAttributes
        {
            get => _customAttributes ?? (_customAttributes = new InputMap<string>());
            set => _customAttributes = value;
        }

        [Input("eventTypes")]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        /// <summary>
        /// The ID of the created notification.
        /// </summary>
        [Input("notificationId")]
        public Input<string>? NotificationId { get; set; }

        /// <summary>
        /// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
        /// </summary>
        [Input("objectNamePrefix")]
        public Input<string>? ObjectNamePrefix { get; set; }

        /// <summary>
        /// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
        /// </summary>
        [Input("payloadFormat")]
        public Input<string>? PayloadFormat { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The Cloud PubSub topic to which this subscription publishes. Expects either the
        /// topic name, assumed to belong to the default GCP provider project, or the project-level name,
        /// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
        /// you will need to use the project-level name.
        /// 
        /// - - -
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        public NotificationState()
        {
        }
        public static new NotificationState Empty => new NotificationState();
    }
}
