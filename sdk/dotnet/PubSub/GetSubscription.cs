// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub
{
    public static class GetSubscription
    {
        /// <summary>
        /// Get information about a Google Cloud Pub/Sub Subscription. For more information see
        /// the [official documentation](https://cloud.google.com/pubsub/docs/)
        /// and [API](https://cloud.google.com/pubsub/docs/apis).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_pubsub_subscription = Gcp.PubSub.GetSubscription.Invoke(new()
        ///     {
        ///         Name = "my-pubsub-subscription",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSubscriptionResult> InvokeAsync(GetSubscriptionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionResult>("gcp:pubsub/getSubscription:getSubscription", args ?? new GetSubscriptionArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Google Cloud Pub/Sub Subscription. For more information see
        /// the [official documentation](https://cloud.google.com/pubsub/docs/)
        /// and [API](https://cloud.google.com/pubsub/docs/apis).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_pubsub_subscription = Gcp.PubSub.GetSubscription.Invoke(new()
        ///     {
        ///         Name = "my-pubsub-subscription",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSubscriptionResult> Invoke(GetSubscriptionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubscriptionResult>("gcp:pubsub/getSubscription:getSubscription", args ?? new GetSubscriptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriptionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Cloud Pub/Sub Subscription.
        /// 
        /// - - -
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetSubscriptionArgs()
        {
        }
        public static new GetSubscriptionArgs Empty => new GetSubscriptionArgs();
    }

    public sealed class GetSubscriptionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Cloud Pub/Sub Subscription.
        /// 
        /// - - -
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetSubscriptionInvokeArgs()
        {
        }
        public static new GetSubscriptionInvokeArgs Empty => new GetSubscriptionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubscriptionResult
    {
        public readonly int AckDeadlineSeconds;
        public readonly ImmutableArray<Outputs.GetSubscriptionBigqueryConfigResult> BigqueryConfigs;
        public readonly ImmutableArray<Outputs.GetSubscriptionCloudStorageConfigResult> CloudStorageConfigs;
        public readonly ImmutableArray<Outputs.GetSubscriptionDeadLetterPolicyResult> DeadLetterPolicies;
        public readonly bool EnableExactlyOnceDelivery;
        public readonly bool EnableMessageOrdering;
        public readonly ImmutableArray<Outputs.GetSubscriptionExpirationPolicyResult> ExpirationPolicies;
        public readonly string Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string MessageRetentionDuration;
        public readonly string Name;
        public readonly string? Project;
        public readonly ImmutableArray<Outputs.GetSubscriptionPushConfigResult> PushConfigs;
        public readonly bool RetainAckedMessages;
        public readonly ImmutableArray<Outputs.GetSubscriptionRetryPolicyResult> RetryPolicies;
        public readonly string Topic;

        [OutputConstructor]
        private GetSubscriptionResult(
            int ackDeadlineSeconds,

            ImmutableArray<Outputs.GetSubscriptionBigqueryConfigResult> bigqueryConfigs,

            ImmutableArray<Outputs.GetSubscriptionCloudStorageConfigResult> cloudStorageConfigs,

            ImmutableArray<Outputs.GetSubscriptionDeadLetterPolicyResult> deadLetterPolicies,

            bool enableExactlyOnceDelivery,

            bool enableMessageOrdering,

            ImmutableArray<Outputs.GetSubscriptionExpirationPolicyResult> expirationPolicies,

            string filter,

            string id,

            ImmutableDictionary<string, string> labels,

            string messageRetentionDuration,

            string name,

            string? project,

            ImmutableArray<Outputs.GetSubscriptionPushConfigResult> pushConfigs,

            bool retainAckedMessages,

            ImmutableArray<Outputs.GetSubscriptionRetryPolicyResult> retryPolicies,

            string topic)
        {
            AckDeadlineSeconds = ackDeadlineSeconds;
            BigqueryConfigs = bigqueryConfigs;
            CloudStorageConfigs = cloudStorageConfigs;
            DeadLetterPolicies = deadLetterPolicies;
            EnableExactlyOnceDelivery = enableExactlyOnceDelivery;
            EnableMessageOrdering = enableMessageOrdering;
            ExpirationPolicies = expirationPolicies;
            Filter = filter;
            Id = id;
            Labels = labels;
            MessageRetentionDuration = messageRetentionDuration;
            Name = name;
            Project = project;
            PushConfigs = pushConfigs;
            RetainAckedMessages = retainAckedMessages;
            RetryPolicies = retryPolicies;
            Topic = topic;
        }
    }
}
