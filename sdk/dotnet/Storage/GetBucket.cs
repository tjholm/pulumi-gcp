// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static class GetBucket
    {
        /// <summary>
        /// Gets an existing bucket in Google Cloud Storage service (GCS).
        /// See [the official documentation](https://cloud.google.com/storage/docs/key-terms#buckets)
        /// and
        /// [API](https://cloud.google.com/storage/docs/json_api/v1/buckets).
        /// 
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_bucket = Output.Create(Gcp.Storage.GetBucket.InvokeAsync(new Gcp.Storage.GetBucketArgs
        ///         {
        ///             Name = "my-bucket",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBucketResult> InvokeAsync(GetBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBucketResult>("gcp:storage/getBucket:getBucket", args ?? new GetBucketArgs(), options.WithVersion());
    }


    public sealed class GetBucketArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetBucketArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBucketResult
    {
        public readonly ImmutableArray<Outputs.GetBucketCorResult> Cors;
        public readonly bool DefaultEventBasedHold;
        public readonly ImmutableArray<Outputs.GetBucketEncryptionResult> Encryptions;
        public readonly bool ForceDestroy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<Outputs.GetBucketLifecycleRuleResult> LifecycleRules;
        public readonly string Location;
        public readonly ImmutableArray<Outputs.GetBucketLoggingResult> Loggings;
        public readonly string Name;
        public readonly string Project;
        public readonly bool RequesterPays;
        public readonly ImmutableArray<Outputs.GetBucketRetentionPolicyResult> RetentionPolicies;
        public readonly string SelfLink;
        public readonly string StorageClass;
        public readonly bool UniformBucketLevelAccess;
        public readonly string Url;
        public readonly ImmutableArray<Outputs.GetBucketVersioningResult> Versionings;
        public readonly ImmutableArray<Outputs.GetBucketWebsiteResult> Websites;

        [OutputConstructor]
        private GetBucketResult(
            ImmutableArray<Outputs.GetBucketCorResult> cors,

            bool defaultEventBasedHold,

            ImmutableArray<Outputs.GetBucketEncryptionResult> encryptions,

            bool forceDestroy,

            string id,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GetBucketLifecycleRuleResult> lifecycleRules,

            string location,

            ImmutableArray<Outputs.GetBucketLoggingResult> loggings,

            string name,

            string project,

            bool requesterPays,

            ImmutableArray<Outputs.GetBucketRetentionPolicyResult> retentionPolicies,

            string selfLink,

            string storageClass,

            bool uniformBucketLevelAccess,

            string url,

            ImmutableArray<Outputs.GetBucketVersioningResult> versionings,

            ImmutableArray<Outputs.GetBucketWebsiteResult> websites)
        {
            Cors = cors;
            DefaultEventBasedHold = defaultEventBasedHold;
            Encryptions = encryptions;
            ForceDestroy = forceDestroy;
            Id = id;
            Labels = labels;
            LifecycleRules = lifecycleRules;
            Location = location;
            Loggings = loggings;
            Name = name;
            Project = project;
            RequesterPays = requesterPays;
            RetentionPolicies = retentionPolicies;
            SelfLink = selfLink;
            StorageClass = storageClass;
            UniformBucketLevelAccess = uniformBucketLevelAccess;
            Url = url;
            Versionings = versionings;
            Websites = websites;
        }
    }
}
