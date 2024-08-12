// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for BigQuery dataset. Each of these resources serves a different use case:
    /// 
    /// * `gcp.bigquery.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
    /// * `gcp.bigquery.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
    /// * `gcp.bigquery.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
    /// 
    /// These resources are intended to convert the permissions system for BigQuery datasets to the standard IAM interface. For advanced usages, including [creating authorized views](https://cloud.google.com/bigquery/docs/share-access-views), please use either `gcp.bigquery.DatasetAccess` or the `access` field on `gcp.bigquery.Dataset`.
    /// 
    /// &gt; **Note:** These resources **cannot** be used with `gcp.bigquery.DatasetAccess` resources or the `access` field on `gcp.bigquery.Dataset` or they will fight over what the policy should be.
    /// 
    /// &gt; **Note:** Using any of these resources will remove any authorized view permissions from the dataset. To assign and preserve authorized view permissions use the `gcp.bigquery.DatasetAccess` instead.
    /// 
    /// &gt; **Note:** Legacy BigQuery roles `OWNER` `WRITER` and `READER` **cannot** be used with any of these IAM resources. Instead use the full role form of: `roles/bigquery.dataOwner` `roles/bigquery.dataEditor` and `roles/bigquery.dataViewer`.
    /// 
    /// &gt; **Note:** `gcp.bigquery.DatasetIamPolicy` **cannot** be used in conjunction with `gcp.bigquery.DatasetIamBinding` and `gcp.bigquery.DatasetIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.bigquery.DatasetIamBinding` resources **can be** used in conjunction with `gcp.bigquery.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## gcp.bigquery.DatasetIamPolicy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var owner = Gcp.Organizations.GetIAMPolicy.Invoke(new()
    ///     {
    ///         Bindings = new[]
    ///         {
    ///             new Gcp.Organizations.Inputs.GetIAMPolicyBindingInputArgs
    ///             {
    ///                 Role = "roles/bigquery.dataOwner",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var datasetDataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var dataset = new Gcp.BigQuery.DatasetIamPolicy("dataset", new()
    ///     {
    ///         DatasetId = datasetDataset.DatasetId,
    ///         PolicyData = owner.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.bigquery.DatasetIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var reader = new Gcp.BigQuery.DatasetIamBinding("reader", new()
    ///     {
    ///         DatasetId = dataset.DatasetId,
    ///         Role = "roles/bigquery.dataViewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.bigquery.DatasetIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var editor = new Gcp.BigQuery.DatasetIamMember("editor", new()
    ///     {
    ///         DatasetId = dataset.DatasetId,
    ///         Role = "roles/bigquery.dataEditor",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.bigquery.DatasetIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var reader = new Gcp.BigQuery.DatasetIamBinding("reader", new()
    ///     {
    ///         DatasetId = dataset.DatasetId,
    ///         Role = "roles/bigquery.dataViewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.bigquery.DatasetIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dataset = new Gcp.BigQuery.Dataset("dataset", new()
    ///     {
    ///         DatasetId = "example_dataset",
    ///     });
    /// 
    ///     var editor = new Gcp.BigQuery.DatasetIamMember("editor", new()
    ///     {
    ///         DatasetId = dataset.DatasetId,
    ///         Role = "roles/bigquery.dataEditor",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies
    /// 
    /// IAM policy imports use the identifier of the BigQuery Dataset resource. For example:
    /// 
    /// * `projects/{{project_id}}/datasets/{{dataset_id}}`
    /// 
    /// An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
    /// 
    /// tf
    /// 
    /// import {
    /// 
    ///   id = projects/{{project_id}}/datasets/{{dataset_id}}
    /// 
    ///   to = google_bigquery_dataset_iam_policy.default
    /// 
    /// }
    /// 
    /// The `pulumi import` command can also be used:
    /// 
    /// ```sh
    /// $ pulumi import gcp:bigquery/datasetIamPolicy:DatasetIamPolicy default projects/{{project_id}}/datasets/{{dataset_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy")]
    public partial class DatasetIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The dataset ID.
        /// </summary>
        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the dataset's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a DatasetIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetIamPolicy(string name, DatasetIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy", name, args ?? new DatasetIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetIamPolicy(string name, Input<string> id, DatasetIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatasetIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatasetIamPolicy Get(string name, Input<string> id, DatasetIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new DatasetIamPolicy(name, id, state, options);
        }
    }

    public sealed class DatasetIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dataset ID.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatasetIamPolicyArgs()
        {
        }
        public static new DatasetIamPolicyArgs Empty => new DatasetIamPolicyArgs();
    }

    public sealed class DatasetIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dataset ID.
        /// </summary>
        [Input("datasetId")]
        public Input<string>? DatasetId { get; set; }

        /// <summary>
        /// (Computed) The etag of the dataset's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public DatasetIamPolicyState()
        {
        }
        public static new DatasetIamPolicyState Empty => new DatasetIamPolicyState();
    }
}
