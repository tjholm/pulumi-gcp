// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    public static class GetNotificationChannel
    {
        /// <summary>
        /// A NotificationChannel is a medium through which an alert is delivered
        /// when a policy violation is detected. Examples of channels include email, SMS,
        /// and third-party messaging applications. Fields containing sensitive information
        /// like authentication tokens or contact info are only partially populated on retrieval.
        /// 
        /// 
        /// To get more information about NotificationChannel, see:
        /// 
        /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
        /// * How-to Guides
        ///     * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
        ///     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Notification Channel Basic
        /// 
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var basic = Gcp.Monitoring.GetNotificationChannel.Invoke(new()
        ///     {
        ///         DisplayName = "Test Notification Channel",
        ///     });
        /// 
        ///     var alertPolicy = new Gcp.Monitoring.AlertPolicy("alertPolicy", new()
        ///     {
        ///         DisplayName = "My Alert Policy",
        ///         NotificationChannels = new[]
        ///         {
        ///             basic.Apply(getNotificationChannelResult =&gt; getNotificationChannelResult.Name),
        ///         },
        ///         Combiner = "OR",
        ///         Conditions = new[]
        ///         {
        ///             new Gcp.Monitoring.Inputs.AlertPolicyConditionArgs
        ///             {
        ///                 DisplayName = "test condition",
        ///                 ConditionThreshold = new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdArgs
        ///                 {
        ///                     Filter = "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
        ///                     Duration = "60s",
        ///                     Comparison = "COMPARISON_GT",
        ///                     Aggregations = new[]
        ///                     {
        ///                         new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdAggregationArgs
        ///                         {
        ///                             AlignmentPeriod = "60s",
        ///                             PerSeriesAligner = "ALIGN_RATE",
        ///                         },
        ///                     },
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNotificationChannelResult> InvokeAsync(GetNotificationChannelArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNotificationChannelResult>("gcp:monitoring/getNotificationChannel:getNotificationChannel", args ?? new GetNotificationChannelArgs(), options.WithDefaults());

        /// <summary>
        /// A NotificationChannel is a medium through which an alert is delivered
        /// when a policy violation is detected. Examples of channels include email, SMS,
        /// and third-party messaging applications. Fields containing sensitive information
        /// like authentication tokens or contact info are only partially populated on retrieval.
        /// 
        /// 
        /// To get more information about NotificationChannel, see:
        /// 
        /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannels)
        /// * How-to Guides
        ///     * [Notification Options](https://cloud.google.com/monitoring/support/notification-options)
        ///     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Notification Channel Basic
        /// 
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var basic = Gcp.Monitoring.GetNotificationChannel.Invoke(new()
        ///     {
        ///         DisplayName = "Test Notification Channel",
        ///     });
        /// 
        ///     var alertPolicy = new Gcp.Monitoring.AlertPolicy("alertPolicy", new()
        ///     {
        ///         DisplayName = "My Alert Policy",
        ///         NotificationChannels = new[]
        ///         {
        ///             basic.Apply(getNotificationChannelResult =&gt; getNotificationChannelResult.Name),
        ///         },
        ///         Combiner = "OR",
        ///         Conditions = new[]
        ///         {
        ///             new Gcp.Monitoring.Inputs.AlertPolicyConditionArgs
        ///             {
        ///                 DisplayName = "test condition",
        ///                 ConditionThreshold = new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdArgs
        ///                 {
        ///                     Filter = "metric.type=\"compute.googleapis.com/instance/disk/write_bytes_count\" AND resource.type=\"gce_instance\"",
        ///                     Duration = "60s",
        ///                     Comparison = "COMPARISON_GT",
        ///                     Aggregations = new[]
        ///                     {
        ///                         new Gcp.Monitoring.Inputs.AlertPolicyConditionConditionThresholdAggregationArgs
        ///                         {
        ///                             AlignmentPeriod = "60s",
        ///                             PerSeriesAligner = "ALIGN_RATE",
        ///                         },
        ///                     },
        ///                 },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNotificationChannelResult> Invoke(GetNotificationChannelInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNotificationChannelResult>("gcp:monitoring/getNotificationChannel:getNotificationChannel", args ?? new GetNotificationChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNotificationChannelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name for this notification channel.
        /// </summary>
        [Input("displayName")]
        public string? DisplayName { get; set; }

        [Input("labels")]
        private Dictionary<string, string>? _labels;

        /// <summary>
        /// Labels (corresponding to the
        /// NotificationChannelDescriptor schema) to filter the notification channels by.
        /// </summary>
        public Dictionary<string, string> Labels
        {
            get => _labels ?? (_labels = new Dictionary<string, string>());
            set => _labels = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The type of the notification channel.
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        [Input("userLabels")]
        private Dictionary<string, string>? _userLabels;

        /// <summary>
        /// User-provided key-value labels to filter by.
        /// </summary>
        public Dictionary<string, string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new Dictionary<string, string>());
            set => _userLabels = value;
        }

        public GetNotificationChannelArgs()
        {
        }
        public static new GetNotificationChannelArgs Empty => new GetNotificationChannelArgs();
    }

    public sealed class GetNotificationChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The display name for this notification channel.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels (corresponding to the
        /// NotificationChannelDescriptor schema) to filter the notification channels by.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The type of the notification channel.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// User-provided key-value labels to filter by.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public GetNotificationChannelInvokeArgs()
        {
        }
        public static new GetNotificationChannelInvokeArgs Empty => new GetNotificationChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetNotificationChannelResult
    {
        public readonly string Description;
        public readonly string? DisplayName;
        public readonly bool Enabled;
        public readonly bool ForceDelete;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string>? Labels;
        public readonly string Name;
        public readonly string? Project;
        public readonly ImmutableArray<Outputs.GetNotificationChannelSensitiveLabelResult> SensitiveLabels;
        public readonly string? Type;
        public readonly ImmutableDictionary<string, string>? UserLabels;
        public readonly string VerificationStatus;

        [OutputConstructor]
        private GetNotificationChannelResult(
            string description,

            string? displayName,

            bool enabled,

            bool forceDelete,

            string id,

            ImmutableDictionary<string, string>? labels,

            string name,

            string? project,

            ImmutableArray<Outputs.GetNotificationChannelSensitiveLabelResult> sensitiveLabels,

            string? type,

            ImmutableDictionary<string, string>? userLabels,

            string verificationStatus)
        {
            Description = description;
            DisplayName = displayName;
            Enabled = enabled;
            ForceDelete = forceDelete;
            Id = id;
            Labels = labels;
            Name = name;
            Project = project;
            SensitiveLabels = sensitiveLabels;
            Type = type;
            UserLabels = userLabels;
            VerificationStatus = verificationStatus;
        }
    }
}
