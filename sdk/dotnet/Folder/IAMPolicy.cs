// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Folder
{
    /// <summary>
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `folder`, role, and member e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `folder` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `folder`.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder folder
    /// ```
    /// 
    ///  IAM audit config imports use the identifier of the resource in question and the service, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy my_folder "folder foo.googleapis.com"
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -&gt; **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
    /// 
    /// ```sh
    ///  $ pulumi import gcp:folder/iAMPolicy:IAMPolicy to include the title of condition, e.g. `google_folder_iam_binding.my_folder "folder roles/{{role_id}} condition-title"`
    /// ```
    /// </summary>
    [GcpResourceType("gcp:folder/iAMPolicy:IAMPolicy")]
    public partial class IAMPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the folder's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Output("folder")]
        public Output<string> Folder { get; private set; } = null!;

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the folder. The policy will be
        /// merged with any existing policy applied to the folder.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a IAMPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMPolicy(string name, IAMPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:folder/iAMPolicy:IAMPolicy", name, args ?? new IAMPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMPolicy(string name, Input<string> id, IAMPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:folder/iAMPolicy:IAMPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IAMPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMPolicy Get(string name, Input<string> id, IAMPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMPolicy(name, id, state, options);
        }
    }

    public sealed class IAMPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the folder. The policy will be
        /// merged with any existing policy applied to the folder.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public IAMPolicyArgs()
        {
        }
        public static new IAMPolicyArgs Empty => new IAMPolicyArgs();
    }

    public sealed class IAMPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the folder's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Input("folder")]
        public Input<string>? Folder { get; set; }

        /// <summary>
        /// The `gcp.organizations.getIAMPolicy` data source that represents
        /// the IAM policy that will be applied to the folder. The policy will be
        /// merged with any existing policy applied to the folder.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public IAMPolicyState()
        {
        }
        public static new IAMPolicyState Empty => new IAMPolicyState();
    }
}
