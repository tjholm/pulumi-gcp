// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetRegionDiskIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for regiondisk
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
        ///     var policy = Gcp.Compute.GetRegionDiskIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_compute_region_disk.Regiondisk.Project,
        ///         Region = google_compute_region_disk.Regiondisk.Region,
        ///         Name = google_compute_region_disk.Regiondisk.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRegionDiskIamPolicyResult> InvokeAsync(GetRegionDiskIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionDiskIamPolicyResult>("gcp:compute/getRegionDiskIamPolicy:getRegionDiskIamPolicy", args ?? new GetRegionDiskIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for regiondisk
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
        ///     var policy = Gcp.Compute.GetRegionDiskIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_compute_region_disk.Regiondisk.Project,
        ///         Region = google_compute_region_disk.Regiondisk.Region,
        ///         Name = google_compute_region_disk.Regiondisk.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRegionDiskIamPolicyResult> Invoke(GetRegionDiskIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionDiskIamPolicyResult>("gcp:compute/getRegionDiskIamPolicy:getRegionDiskIamPolicy", args ?? new GetRegionDiskIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionDiskIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// A reference to the region where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        /// region is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetRegionDiskIamPolicyArgs()
        {
        }
        public static new GetRegionDiskIamPolicyArgs Empty => new GetRegionDiskIamPolicyArgs();
    }

    public sealed class GetRegionDiskIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the region where the disk resides. Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
        /// region is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetRegionDiskIamPolicyInvokeArgs()
        {
        }
        public static new GetRegionDiskIamPolicyInvokeArgs Empty => new GetRegionDiskIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionDiskIamPolicyResult
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// (Required only by `gcp.compute.RegionDiskIamPolicy`) The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        public readonly string PolicyData;
        public readonly string Project;
        public readonly string Region;

        [OutputConstructor]
        private GetRegionDiskIamPolicyResult(
            string etag,

            string id,

            string name,

            string policyData,

            string project,

            string region)
        {
            Etag = etag;
            Id = id;
            Name = name;
            PolicyData = policyData;
            Project = project;
            Region = region;
        }
    }
}
