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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.Organizations.Folder("default", new()
    ///     {
    ///         DisplayName = "some-folder-name",
    ///         Parent = "organizations/123456789",
    ///     });
    /// 
    ///     var basic = new Gcp.Logging.FolderBucketConfig("basic", new()
    ///     {
    ///         Folder = @default.Name,
    ///         Location = "global",
    ///         RetentionDays = 30,
    ///         BucketId = "_Default",
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
    ///  $ pulumi import gcp:logging/folderBucketConfig:FolderBucketConfig default folders/{{folder}}/locations/{{location}}/buckets/{{bucket_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:logging/folderBucketConfig:FolderBucketConfig")]
    public partial class FolderBucketConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Output("bucketId")]
        public Output<string> BucketId { get; private set; } = null!;

        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
        /// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
        /// updating the log bucket. Changing the KMS key is allowed.
        /// </summary>
        [Output("cmekSettings")]
        public Output<Outputs.FolderBucketConfigCmekSettings?> CmekSettings { get; private set; } = null!;

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Output("folder")]
        public Output<string> Folder { get; private set; } = null!;

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
        /// The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        /// </summary>
        [Output("retentionDays")]
        public Output<int?> RetentionDays { get; private set; } = null!;


        /// <summary>
        /// Create a FolderBucketConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FolderBucketConfig(string name, FolderBucketConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/folderBucketConfig:FolderBucketConfig", name, args ?? new FolderBucketConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FolderBucketConfig(string name, Input<string> id, FolderBucketConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/folderBucketConfig:FolderBucketConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FolderBucketConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FolderBucketConfig Get(string name, Input<string> id, FolderBucketConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new FolderBucketConfig(name, id, state, options);
        }
    }

    public sealed class FolderBucketConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
        /// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
        /// updating the log bucket. Changing the KMS key is allowed.
        /// </summary>
        [Input("cmekSettings")]
        public Input<Inputs.FolderBucketConfigCmekSettingsArgs>? CmekSettings { get; set; }

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        /// <summary>
        /// The location of the bucket.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public FolderBucketConfigArgs()
        {
        }
        public static new FolderBucketConfigArgs Empty => new FolderBucketConfigArgs();
    }

    public sealed class FolderBucketConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
        /// </summary>
        [Input("bucketId")]
        public Input<string>? BucketId { get; set; }

        /// <summary>
        /// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
        /// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
        /// updating the log bucket. Changing the KMS key is allowed.
        /// </summary>
        [Input("cmekSettings")]
        public Input<Inputs.FolderBucketConfigCmekSettingsGetArgs>? CmekSettings { get; set; }

        /// <summary>
        /// Describes this bucket.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The parent resource that contains the logging bucket.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

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
        /// The resource name of the bucket. For example: "folders/my-folder-id/locations/my-location/buckets/my-bucket-id"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
        /// </summary>
        [Input("retentionDays")]
        public Input<int>? RetentionDays { get; set; }

        public FolderBucketConfigState()
        {
        }
        public static new FolderBucketConfigState Empty => new FolderBucketConfigState();
    }
}
