// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataCatalog
{
    /// <summary>
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* {{policy_tag}} Any variables not passed in the import command will be taken from the provider configuration. Data catalog policytag IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding editor "{{policy_tag}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding editor "{{policy_tag}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding editor {{policy_tag}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding")]
    public partial class PolicyTagIamBinding : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.PolicyTagIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("policyTag")]
        public Output<string> PolicyTag { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyTagIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyTagIamBinding(string name, PolicyTagIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding", name, args ?? new PolicyTagIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyTagIamBinding(string name, Input<string> id, PolicyTagIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/policyTagIamBinding:PolicyTagIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyTagIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyTagIamBinding Get(string name, Input<string> id, PolicyTagIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyTagIamBinding(name, id, state, options);
        }
    }

    public sealed class PolicyTagIamBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.PolicyTagIamBindingConditionArgs>? Condition { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("policyTag", required: true)]
        public Input<string> PolicyTag { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public PolicyTagIamBindingArgs()
        {
        }
        public static new PolicyTagIamBindingArgs Empty => new PolicyTagIamBindingArgs();
    }

    public sealed class PolicyTagIamBindingState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.PolicyTagIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("policyTag")]
        public Input<string>? PolicyTag { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.datacatalog.PolicyTagIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public PolicyTagIamBindingState()
        {
        }
        public static new PolicyTagIamBindingState Empty => new PolicyTagIamBindingState();
    }
}
