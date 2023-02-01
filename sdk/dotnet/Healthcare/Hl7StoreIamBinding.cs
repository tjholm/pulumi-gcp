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
    /// This member resource can be imported using the `hl7_v2_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `hl7_v2_store_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `hl7_v2_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding hl7_v2_store_iam your-project-id/location-name/dataset-name/hl7-v2-store-name
    /// ```
    /// </summary>
    [GcpResourceType("gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding")]
    public partial class Hl7StoreIamBinding : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.Hl7StoreIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("hl7V2StoreId")]
        public Output<string> Hl7V2StoreId { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a Hl7StoreIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hl7StoreIamBinding(string name, Hl7StoreIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding", name, args ?? new Hl7StoreIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hl7StoreIamBinding(string name, Input<string> id, Hl7StoreIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamBinding:Hl7StoreIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Hl7StoreIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hl7StoreIamBinding Get(string name, Input<string> id, Hl7StoreIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new Hl7StoreIamBinding(name, id, state, options);
        }
    }

    public sealed class Hl7StoreIamBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.Hl7StoreIamBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId", required: true)]
        public Input<string> Hl7V2StoreId { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public Hl7StoreIamBindingArgs()
        {
        }
        public static new Hl7StoreIamBindingArgs Empty => new Hl7StoreIamBindingArgs();
    }

    public sealed class Hl7StoreIamBindingState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.Hl7StoreIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId")]
        public Input<string>? Hl7V2StoreId { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.Hl7StoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public Hl7StoreIamBindingState()
        {
        }
        public static new Hl7StoreIamBindingState Empty => new Hl7StoreIamBindingState();
    }
}
