// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc
{
    /// <summary>
    /// Three different resources help you manage IAM policies on dataproc jobs. Each of these resources serves a different use case:
    /// 
    /// * `gcp.dataproc.JobIAMPolicy`: Authoritative. Sets the IAM policy for the job and replaces any existing policy already attached.
    /// * `gcp.dataproc.JobIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the job are preserved.
    /// * `gcp.dataproc.JobIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the job are preserved.
    /// 
    /// &gt; **Note:** `gcp.dataproc.JobIAMPolicy` **cannot** be used in conjunction with `gcp.dataproc.JobIAMBinding` and `gcp.dataproc.JobIAMMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the job as `gcp.dataproc.JobIAMPolicy` replaces the entire policy.
    /// 
    /// &gt; **Note:** `gcp.dataproc.JobIAMBinding` resources **can be** used in conjunction with `gcp.dataproc.JobIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_dataproc\_job\_iam\_policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var admin = Gcp.Organizations.GetIAMPolicy.Invoke(new()
    ///     {
    ///         Bindings = new[]
    ///         {
    ///             new Gcp.Organizations.Inputs.GetIAMPolicyBindingInputArgs
    ///             {
    ///                 Role = "roles/editor",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var editor = new Gcp.Dataproc.JobIAMPolicy("editor", new()
    ///     {
    ///         Project = "your-project",
    ///         Region = "your-region",
    ///         JobId = "your-dataproc-job",
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_dataproc\_job\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var editor = new Gcp.Dataproc.JobIAMBinding("editor", new()
    ///     {
    ///         JobId = "your-dataproc-job",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Role = "roles/editor",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_dataproc\_job\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var editor = new Gcp.Dataproc.JobIAMMember("editor", new()
    ///     {
    ///         JobId = "your-dataproc-job",
    ///         Member = "user:jane@example.com",
    ///         Role = "roles/editor",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies IAM policy imports use the `job_id` identifier of the Dataproc Job resource only. For example* `projects/{project}/regions/{region}/jobs/{job_id}` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
    /// 
    ///  id = "projects/{project}/regions/{region}/jobs/{job_id}"
    /// 
    ///  to = google_dataproc_job_iam_policy.default } The `pulumi import` command can also be used
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/jobIAMPolicy:JobIAMPolicy default "projects/{project}/regions/{region}/jobs/{job_id}"
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataproc/jobIAMPolicy:JobIAMPolicy")]
    public partial class JobIAMPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the jobs's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
        /// 
        /// - - -
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The project in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The region in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a JobIAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobIAMPolicy(string name, JobIAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataproc/jobIAMPolicy:JobIAMPolicy", name, args ?? new JobIAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobIAMPolicy(string name, Input<string> id, JobIAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/jobIAMPolicy:JobIAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing JobIAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobIAMPolicy Get(string name, Input<string> id, JobIAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new JobIAMPolicy(name, id, state, options);
        }
    }

    public sealed class JobIAMPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
        /// 
        /// - - -
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The project in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public JobIAMPolicyArgs()
        {
        }
        public static new JobIAMPolicyArgs Empty => new JobIAMPolicyArgs();
    }

    public sealed class JobIAMPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the jobs's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// The policy data generated by a `gcp.organizations.getIAMPolicy` data source.
        /// 
        /// - - -
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The project in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the job belongs. If it
        /// is not provided, the provider will use a default.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public JobIAMPolicyState()
        {
        }
        public static new JobIAMPolicyState Empty => new JobIAMPolicyState();
    }
}
