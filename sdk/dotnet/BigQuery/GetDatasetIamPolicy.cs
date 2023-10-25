// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    public static class GetDatasetIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for a BigQuery dataset.
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
        ///     var policy = Gcp.BigQuery.GetDatasetIamPolicy.Invoke(new()
        ///     {
        ///         DatasetId = google_bigquery_dataset.Dataset.Dataset_id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDatasetIamPolicyResult> InvokeAsync(GetDatasetIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatasetIamPolicyResult>("gcp:bigquery/getDatasetIamPolicy:getDatasetIamPolicy", args ?? new GetDatasetIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for a BigQuery dataset.
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
        ///     var policy = Gcp.BigQuery.GetDatasetIamPolicy.Invoke(new()
        ///     {
        ///         DatasetId = google_bigquery_dataset.Dataset.Dataset_id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDatasetIamPolicyResult> Invoke(GetDatasetIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatasetIamPolicyResult>("gcp:bigquery/getDatasetIamPolicy:getDatasetIamPolicy", args ?? new GetDatasetIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatasetIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The dataset ID.
        /// </summary>
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetDatasetIamPolicyArgs()
        {
        }
        public static new GetDatasetIamPolicyArgs Empty => new GetDatasetIamPolicyArgs();
    }

    public sealed class GetDatasetIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The dataset ID.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatasetIamPolicyInvokeArgs()
        {
        }
        public static new GetDatasetIamPolicyInvokeArgs Empty => new GetDatasetIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatasetIamPolicyResult
    {
        public readonly string DatasetId;
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The policy data
        /// </summary>
        public readonly string PolicyData;
        public readonly string Project;

        [OutputConstructor]
        private GetDatasetIamPolicyResult(
            string datasetId,

            string etag,

            string id,

            string policyData,

            string project)
        {
            DatasetId = datasetId;
            Etag = etag;
            Id = id;
            PolicyData = policyData;
            Project = project;
        }
    }
}
