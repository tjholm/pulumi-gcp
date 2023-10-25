// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetInstanceIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for instance
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
        ///     var policy = Gcp.Compute.GetInstanceIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_compute_instance.Default.Project,
        ///         Zone = google_compute_instance.Default.Zone,
        ///         InstanceName = google_compute_instance.Default.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetInstanceIamPolicyResult> InvokeAsync(GetInstanceIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstanceIamPolicyResult>("gcp:compute/getInstanceIamPolicy:getInstanceIamPolicy", args ?? new GetInstanceIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for instance
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
        ///     var policy = Gcp.Compute.GetInstanceIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_compute_instance.Default.Project,
        ///         Zone = google_compute_instance.Default.Zone,
        ///         InstanceName = google_compute_instance.Default.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstanceIamPolicyResult> Invoke(GetInstanceIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstanceIamPolicyResult>("gcp:compute/getInstanceIamPolicy:getInstanceIamPolicy", args ?? new GetInstanceIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
        /// zone is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("zone")]
        public string? Zone { get; set; }

        public GetInstanceIamPolicyArgs()
        {
        }
        public static new GetInstanceIamPolicyArgs Empty => new GetInstanceIamPolicyArgs();
    }

    public sealed class GetInstanceIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
        /// zone is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public GetInstanceIamPolicyInvokeArgs()
        {
        }
        public static new GetInstanceIamPolicyInvokeArgs Empty => new GetInstanceIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceIamPolicyResult
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceName;
        /// <summary>
        /// (Required only by `gcp.compute.InstanceIAMPolicy`) The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        public readonly string PolicyData;
        public readonly string Project;
        public readonly string Zone;

        [OutputConstructor]
        private GetInstanceIamPolicyResult(
            string etag,

            string id,

            string instanceName,

            string policyData,

            string project,

            string zone)
        {
            Etag = etag;
            Id = id;
            InstanceName = instanceName;
            PolicyData = policyData;
            Project = project;
            Zone = zone;
        }
    }
}
