// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataPlex
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Cloud Dataplex Asset. Each of these resources serves a different use case:
    /// 
    /// * `gcp.dataplex.AssetIamPolicy`: Authoritative. Sets the IAM policy for the asset and replaces any existing policy already attached.
    /// * `gcp.dataplex.AssetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the asset are preserved.
    /// * `gcp.dataplex.AssetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the asset are preserved.
    /// 
    /// &gt; **Note:** `gcp.dataplex.AssetIamPolicy` **cannot** be used in conjunction with `gcp.dataplex.AssetIamBinding` and `gcp.dataplex.AssetIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.dataplex.AssetIamBinding` resources **can be** used in conjunction with `gcp.dataplex.AssetIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_dataplex\_asset\_iam\_policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
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
    ///                 Role = "roles/viewer",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.DataPlex.AssetIamPolicy("policy", new()
    ///     {
    ///         Project = google_dataplex_asset.Example.Project,
    ///         Location = google_dataplex_asset.Example.Location,
    ///         Lake = google_dataplex_asset.Example.Lake,
    ///         DataplexZone = google_dataplex_asset.Example.Dataplex_zone,
    ///         Asset = google_dataplex_asset.Example.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_dataplex\_asset\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.DataPlex.AssetIamBinding("binding", new()
    ///     {
    ///         Project = google_dataplex_asset.Example.Project,
    ///         Location = google_dataplex_asset.Example.Location,
    ///         Lake = google_dataplex_asset.Example.Lake,
    ///         DataplexZone = google_dataplex_asset.Example.Dataplex_zone,
    ///         Asset = google_dataplex_asset.Example.Name,
    ///         Role = "roles/viewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_dataplex\_asset\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.DataPlex.AssetIamMember("member", new()
    ///     {
    ///         Project = google_dataplex_asset.Example.Project,
    ///         Location = google_dataplex_asset.Example.Location,
    ///         Lake = google_dataplex_asset.Example.Lake,
    ///         DataplexZone = google_dataplex_asset.Example.Dataplex_zone,
    ///         Asset = google_dataplex_asset.Example.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{name}} * {{project}}/{{location}}/{{lake}}/{{dataplex_zone}}/{{name}} * {{location}}/{{lake}}/{{dataplex_zone}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Dataplex asset IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/assetIamBinding:AssetIamBinding editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{asset}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/assetIamBinding:AssetIamBinding editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{asset}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/assetIamBinding:AssetIamBinding editor projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}/assets/{{asset}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:dataplex/assetIamBinding:AssetIamBinding")]
    public partial class AssetIamBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("asset")]
        public Output<string> Asset { get; private set; } = null!;

        [Output("condition")]
        public Output<Outputs.AssetIamBindingCondition?> Condition { get; private set; } = null!;

        [Output("dataplexZone")]
        public Output<string> DataplexZone { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("lake")]
        public Output<string> Lake { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataplex.AssetIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a AssetIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssetIamBinding(string name, AssetIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataplex/assetIamBinding:AssetIamBinding", name, args ?? new AssetIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssetIamBinding(string name, Input<string> id, AssetIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataplex/assetIamBinding:AssetIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AssetIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssetIamBinding Get(string name, Input<string> id, AssetIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new AssetIamBinding(name, id, state, options);
        }
    }

    public sealed class AssetIamBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("asset", required: true)]
        public Input<string> Asset { get; set; } = null!;

        [Input("condition")]
        public Input<Inputs.AssetIamBindingConditionArgs>? Condition { get; set; }

        [Input("dataplexZone", required: true)]
        public Input<string> DataplexZone { get; set; } = null!;

        [Input("lake", required: true)]
        public Input<string> Lake { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataplex.AssetIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public AssetIamBindingArgs()
        {
        }
        public static new AssetIamBindingArgs Empty => new AssetIamBindingArgs();
    }

    public sealed class AssetIamBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("asset")]
        public Input<string>? Asset { get; set; }

        [Input("condition")]
        public Input<Inputs.AssetIamBindingConditionGetArgs>? Condition { get; set; }

        [Input("dataplexZone")]
        public Input<string>? DataplexZone { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("lake")]
        public Input<string>? Lake { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.dataplex.AssetIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public AssetIamBindingState()
        {
        }
        public static new AssetIamBindingState Empty => new AssetIamBindingState();
    }
}
