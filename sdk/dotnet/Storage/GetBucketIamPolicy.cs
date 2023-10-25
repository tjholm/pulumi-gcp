// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    public static class GetBucketIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for bucket
        /// 
        /// 
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var policy = Gcp.Storage.GetBucketIamPolicy.Invoke(new()
        ///     {
        ///         Bucket = google_storage_bucket.Default.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetBucketIamPolicyResult> InvokeAsync(GetBucketIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBucketIamPolicyResult>("gcp:storage/getBucketIamPolicy:getBucketIamPolicy", args ?? new GetBucketIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for bucket
        /// 
        /// 
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var policy = Gcp.Storage.GetBucketIamPolicy.Invoke(new()
        ///     {
        ///         Bucket = google_storage_bucket.Default.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetBucketIamPolicyResult> Invoke(GetBucketIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBucketIamPolicyResult>("gcp:storage/getBucketIamPolicy:getBucketIamPolicy", args ?? new GetBucketIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBucketIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("bucket", required: true)]
        public string Bucket { get; set; } = null!;

        public GetBucketIamPolicyArgs()
        {
        }
        public static new GetBucketIamPolicyArgs Empty => new GetBucketIamPolicyArgs();
    }

    public sealed class GetBucketIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        public GetBucketIamPolicyInvokeArgs()
        {
        }
        public static new GetBucketIamPolicyInvokeArgs Empty => new GetBucketIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetBucketIamPolicyResult
    {
        public readonly string Bucket;
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Required only by `gcp.storage.BucketIAMPolicy`) The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        public readonly string PolicyData;

        [OutputConstructor]
        private GetBucketIamPolicyResult(
            string bucket,

            string etag,

            string id,

            string policyData)
        {
            Bucket = bucket;
            Etag = etag;
            Id = id;
            PolicyData = policyData;
        }
    }
}
