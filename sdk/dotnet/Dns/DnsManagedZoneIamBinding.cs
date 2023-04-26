// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Cloud DNS ManagedZone. Each of these resources serves a different use case:
    /// 
    /// * `gcp.dns.DnsManagedZoneIamPolicy`: Authoritative. Sets the IAM policy for the managedzone and replaces any existing policy already attached.
    /// * `gcp.dns.DnsManagedZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedzone are preserved.
    /// * `gcp.dns.DnsManagedZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedzone are preserved.
    /// 
    /// &gt; **Note:** `gcp.dns.DnsManagedZoneIamPolicy` **cannot** be used in conjunction with `gcp.dns.DnsManagedZoneIamBinding` and `gcp.dns.DnsManagedZoneIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.dns.DnsManagedZoneIamBinding` resources **can be** used in conjunction with `gcp.dns.DnsManagedZoneIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_dns\_managed\_zone\_iam\_policy
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
    ///                 Role = "roles/viewer",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.Dns.DnsManagedZoneIamPolicy("policy", new()
    ///     {
    ///         Project = google_dns_managed_zone.Default.Project,
    ///         ManagedZone = google_dns_managed_zone.Default.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_dns\_managed\_zone\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.Dns.DnsManagedZoneIamBinding("binding", new()
    ///     {
    ///         Project = google_dns_managed_zone.Default.Project,
    ///         ManagedZone = google_dns_managed_zone.Default.Name,
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
    /// ## google\_dns\_managed\_zone\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.Dns.DnsManagedZoneIamMember("member", new()
    ///     {
    ///         Project = google_dns_managed_zone.Default.Project,
    ///         ManagedZone = google_dns_managed_zone.Default.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/managedZones/{{managed_zone}} * {{project}}/{{managed_zone}} * {{managed_zone}} Any variables not passed in the import command will be taken from the provider configuration. Cloud DNS managedzone IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dns/dnsManagedZoneIamBinding:DnsManagedZoneIamBinding editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dns/dnsManagedZoneIamBinding:DnsManagedZoneIamBinding editor "projects/{{project}}/managedZones/{{managed_zone}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dns/dnsManagedZoneIamBinding:DnsManagedZoneIamBinding editor projects/{{project}}/managedZones/{{managed_zone}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:dns/dnsManagedZoneIamBinding:DnsManagedZoneIamBinding")]
    public partial class DnsManagedZoneIamBinding : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.DnsManagedZoneIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("managedZone")]
        public Output<string> ManagedZone { get; private set; } = null!;

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
        /// `gcp.dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a DnsManagedZoneIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsManagedZoneIamBinding(string name, DnsManagedZoneIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:dns/dnsManagedZoneIamBinding:DnsManagedZoneIamBinding", name, args ?? new DnsManagedZoneIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsManagedZoneIamBinding(string name, Input<string> id, DnsManagedZoneIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dns/dnsManagedZoneIamBinding:DnsManagedZoneIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DnsManagedZoneIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsManagedZoneIamBinding Get(string name, Input<string> id, DnsManagedZoneIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsManagedZoneIamBinding(name, id, state, options);
        }
    }

    public sealed class DnsManagedZoneIamBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DnsManagedZoneIamBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("managedZone", required: true)]
        public Input<string> ManagedZone { get; set; } = null!;

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
        /// `gcp.dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public DnsManagedZoneIamBindingArgs()
        {
        }
        public static new DnsManagedZoneIamBindingArgs Empty => new DnsManagedZoneIamBindingArgs();
    }

    public sealed class DnsManagedZoneIamBindingState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DnsManagedZoneIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("managedZone")]
        public Input<string>? ManagedZone { get; set; }

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
        /// `gcp.dns.DnsManagedZoneIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public DnsManagedZoneIamBindingState()
        {
        }
        public static new DnsManagedZoneIamBindingState Empty => new DnsManagedZoneIamBindingState();
    }
}
