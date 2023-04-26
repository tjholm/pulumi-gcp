// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// Manages a project-level logging bucket config. For more information see
    /// [the official logging documentation](https://cloud.google.com/logging/docs/) and
    /// [Storing Logs](https://cloud.google.com/logging/docs/storage).
    /// 
    /// &gt; **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
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
    ///     var @default = new Gcp.Organizations.Project("default", new()
    ///     {
    ///         ProjectId = "your-project-id",
    ///         OrgId = "123456789",
    ///     });
    /// 
    ///     var basic = new Gcp.Logging.ProjectBucketConfig("basic", new()
    ///     {
    ///         Project = @default.Id,
    ///         Location = "global",
    ///         RetentionDays = 30,
    ///         BucketId = "_Default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create logging bucket with customId
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basic = new Gcp.Logging.ProjectBucketConfig("basic", new()
    ///     {
    ///         BucketId = "custom-bucket",
    ///         Location = "global",
    ///         Project = "project_id",
    ///         RetentionDays = 30,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create logging bucket with Log Analytics enabled
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var analytics_enabled_bucket = new Gcp.Logging.ProjectBucketConfig("analytics-enabled-bucket", new()
    ///     {
    ///         BucketId = "custom-bucket",
    ///         EnableAnalytics = true,
    ///         Location = "global",
    ///         Project = "project_id",
    ///         RetentionDays = 30,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// Create logging bucket with customId and cmekSettings
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cmekSettings = Gcp.Logging.GetProjectCmekSettings.Invoke(new()
    ///     {
    ///         Project = "project_id",
    ///     });
    /// 
    ///     var keyring = new Gcp.Kms.KeyRing("keyring", new()
    ///     {
    ///         Location = "us-central1",
    ///     });
    /// 
    ///     var key = new Gcp.Kms.CryptoKey("key", new()
    ///     {
    ///         KeyRing = keyring.Id,
    ///         RotationPeriod = "100000s",
    ///     });
    /// 
    ///     var cryptoKeyBinding = new Gcp.Kms.CryptoKeyIAMBinding("cryptoKeyBinding", new()
    ///     {
    ///         CryptoKeyId = key.Id,
    ///         Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
    ///         Members = new[]
    ///         {
    ///             $"serviceAccount:{cmekSettings.Apply(getProjectCmekSettingsResult =&gt; getProjectCmekSettingsResult.ServiceAccountId)}",
    ///         },
    ///     });
    /// 
    ///     var example_project_bucket_cmek_settings = new Gcp.Logging.ProjectBucketConfig("example-project-bucket-cmek-settings", new()
    ///     {
    ///         Project = "project_id",
    ///         Location = "us-central1",
    ///         RetentionDays = 30,
    ///         BucketId = "custom-bucket",
    ///         CmekSettings = new Gcp.Logging.Inputs.ProjectBucketConfigCmekSettingsArgs
    ///         {
    ///             KmsKeyName = key.Id,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             cryptoKeyBinding,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the following format
    /// 
    /// ```sh
    ///  $ pulumi import gcp:logging/projectBucketConfig:ProjectBucketConfig default projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:logging/projectBucketConfig:ProjectBucketConfig")]
    public partial class ProjectBucketConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Output("bucketId")]
        public Output<string> BucketId { get; private set; } = null!;

        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
        /// </summary>
        [Output("cmekSettings")]
        public Output<Outputs.ProjectBucketConfigCmekSettings?> CmekSettings { get; private set; } = null!;

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
        /// </summary>
        [Output("enableAnalytics")]
        public Output<bool?> EnableAnalytics { get; private set; } = null!;

        /// <summary>
        /// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        /// </summary>
        [Output("lifecycleState")]
        public Output<string> LifecycleState { get; private set; } = null!;

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the CMEK settings.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectBucketConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectBucketConfig(string name, ProjectBucketConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, args ?? new ProjectBucketConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectBucketConfig(string name, Input<string> id, ProjectBucketConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectBucketConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectBucketConfig Get(string name, Input<string> id, ProjectBucketConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectBucketConfig(name, id, state, options);
        }
    }

    public sealed class ProjectBucketConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
        /// </summary>
        [Input("cmekSettings")]
        public Input<Inputs.ProjectBucketConfigCmekSettingsArgs>? CmekSettings { get; set; }

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
        /// </summary>
        [Input("enableAnalytics")]
        public Input<bool>? EnableAnalytics { get; set; }

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public ProjectBucketConfigArgs()
        {
        }
        public static new ProjectBucketConfigArgs Empty => new ProjectBucketConfigArgs();
    }

    public sealed class ProjectBucketConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Input("bucketId")]
        public Input<string>? BucketId { get; set; }

        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
        /// </summary>
        [Input("cmekSettings")]
        public Input<Inputs.ProjectBucketConfigCmekSettingsGetArgs>? CmekSettings { get; set; }

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
        /// </summary>
        [Input("enableAnalytics")]
        public Input<bool>? EnableAnalytics { get; set; }

        /// <summary>
        /// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
        /// </summary>
        [Input("lifecycleState")]
        public Input<string>? LifecycleState { get; set; }

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the CMEK settings.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public ProjectBucketConfigState()
        {
        }
        public static new ProjectBucketConfigState Empty => new ProjectBucketConfigState();
    }
}
