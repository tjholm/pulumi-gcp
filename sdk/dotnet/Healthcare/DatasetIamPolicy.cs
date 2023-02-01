// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `dataset_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/datasetIamPolicy:DatasetIamPolicy dataset_iam "your-project-id/location-name/dataset-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `dataset_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/datasetIamPolicy:DatasetIamPolicy dataset_iam "your-project-id/location-name/dataset-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `dataset_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/datasetIamPolicy:DatasetIamPolicy dataset_iam your-project-id/location-name/dataset-name
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy")]
    public partial class DatasetIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The dataset ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}` or
        /// `{location_name}/{dataset_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
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
        /// Create a DatasetIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatasetIamPolicy(string name, DatasetIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy", name, args ?? new DatasetIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatasetIamPolicy(string name, Input<string> id, DatasetIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// The dataset ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}` or
        /// `{location_name}/{dataset_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public DatasetIamPolicyArgs()
        {
        }
        public static new DatasetIamPolicyArgs Empty => new DatasetIamPolicyArgs();
    }

    public sealed class DatasetIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The dataset ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}` or
        /// `{location_name}/{dataset_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
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

        public DatasetIamPolicyState()
        {
        }
        public static new DatasetIamPolicyState Empty => new DatasetIamPolicyState();
    }
}
