// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetBackendBucket
    {
        /// <summary>
        /// Get information about a BackendBucket.
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
        ///         var my_backend_bucket = Output.Create(Gcp.Compute.GetBackendBucket.InvokeAsync(new Gcp.Compute.GetBackendBucketArgs
        ///         {
        ///             Name = "my-backend",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBackendBucketResult> InvokeAsync(GetBackendBucketArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBackendBucketResult>("gcp:compute/getBackendBucket:getBackendBucket", args ?? new GetBackendBucketArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a BackendBucket.
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
        ///         var my_backend_bucket = Output.Create(Gcp.Compute.GetBackendBucket.InvokeAsync(new Gcp.Compute.GetBackendBucketArgs
        ///         {
        ///             Name = "my-backend",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBackendBucketResult> Invoke(GetBackendBucketInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetBackendBucketResult>("gcp:compute/getBackendBucket:getBackendBucket", args ?? new GetBackendBucketInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackendBucketArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetBackendBucketArgs()
        {
        }
    }

    public sealed class GetBackendBucketInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetBackendBucketInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBackendBucketResult
    {
        public readonly string BucketName;
        public readonly ImmutableArray<Outputs.GetBackendBucketCdnPolicyResult> CdnPolicies;
        public readonly string CreationTimestamp;
        public readonly ImmutableArray<string> CustomResponseHeaders;
        public readonly string Description;
        public readonly bool EnableCdn;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Project;
        public readonly string SelfLink;

        [OutputConstructor]
        private GetBackendBucketResult(
            string bucketName,

            ImmutableArray<Outputs.GetBackendBucketCdnPolicyResult> cdnPolicies,

            string creationTimestamp,

            ImmutableArray<string> customResponseHeaders,

            string description,

            bool enableCdn,

            string id,

            string name,

            string? project,

            string selfLink)
        {
            BucketName = bucketName;
            CdnPolicies = cdnPolicies;
            CreationTimestamp = creationTimestamp;
            CustomResponseHeaders = customResponseHeaders;
            Description = description;
            EnableCdn = enableCdn;
            Id = id;
            Name = name;
            Project = project;
            SelfLink = selfLink;
        }
    }
}
