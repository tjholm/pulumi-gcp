// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceDirectory
{
    /// <summary>
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} * {{project}}/{{location}}/{{namespace_id}} * {{location}}/{{namespace_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory namespace IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding")]
    public partial class NamespaceIamBinding : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.NamespaceIamBindingCondition?> Condition { get; private set; } = null!;

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
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a NamespaceIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NamespaceIamBinding(string name, NamespaceIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding", name, args ?? new NamespaceIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NamespaceIamBinding(string name, Input<string> id, NamespaceIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NamespaceIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NamespaceIamBinding Get(string name, Input<string> id, NamespaceIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new NamespaceIamBinding(name, id, state, options);
        }
    }

    public sealed class NamespaceIamBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.NamespaceIamBindingConditionArgs>? Condition { get; set; }

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
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public NamespaceIamBindingArgs()
        {
        }
        public static new NamespaceIamBindingArgs Empty => new NamespaceIamBindingArgs();
    }

    public sealed class NamespaceIamBindingState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.NamespaceIamBindingConditionGetArgs>? Condition { get; set; }

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
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public NamespaceIamBindingState()
        {
        }
        public static new NamespaceIamBindingState Empty => new NamespaceIamBindingState();
    }
}
