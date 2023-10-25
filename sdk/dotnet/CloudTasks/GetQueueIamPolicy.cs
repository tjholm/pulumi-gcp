// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudTasks
{
    public static class GetQueueIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for queue
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
        ///     var policy = Gcp.CloudTasks.GetQueueIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_cloud_tasks_queue.Default.Project,
        ///         Location = google_cloud_tasks_queue.Default.Location,
        ///         Name = google_cloud_tasks_queue.Default.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetQueueIamPolicyResult> InvokeAsync(GetQueueIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueueIamPolicyResult>("gcp:cloudtasks/getQueueIamPolicy:getQueueIamPolicy", args ?? new GetQueueIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for queue
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
        ///     var policy = Gcp.CloudTasks.GetQueueIamPolicy.Invoke(new()
        ///     {
        ///         Project = google_cloud_tasks_queue.Default.Project,
        ///         Location = google_cloud_tasks_queue.Default.Location,
        ///         Name = google_cloud_tasks_queue.Default.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetQueueIamPolicyResult> Invoke(GetQueueIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueueIamPolicyResult>("gcp:cloudtasks/getQueueIamPolicy:getQueueIamPolicy", args ?? new GetQueueIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueueIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location of the queue Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public string? Location { get; set; }

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

        public GetQueueIamPolicyArgs()
        {
        }
        public static new GetQueueIamPolicyArgs Empty => new GetQueueIamPolicyArgs();
    }

    public sealed class GetQueueIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location of the queue Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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

        public GetQueueIamPolicyInvokeArgs()
        {
        }
        public static new GetQueueIamPolicyInvokeArgs Empty => new GetQueueIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueueIamPolicyResult
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly string Name;
        /// <summary>
        /// (Required only by `gcp.cloudtasks.QueueIamPolicy`) The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        public readonly string PolicyData;
        public readonly string Project;

        [OutputConstructor]
        private GetQueueIamPolicyResult(
            string etag,

            string id,

            string location,

            string name,

            string policyData,

            string project)
        {
            Etag = etag;
            Id = id;
            Location = location;
            Name = name;
            PolicyData = policyData;
            Project = project;
        }
    }
}
