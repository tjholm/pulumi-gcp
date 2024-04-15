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
    /// Three different resources help you manage your IAM policy for Dataplex Zone. Each of these resources serves a different use case:
    /// 
    /// * `gcp.dataplex.ZoneIamPolicy`: Authoritative. Sets the IAM policy for the zone and replaces any existing policy already attached.
    /// * `gcp.dataplex.ZoneIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the zone are preserved.
    /// * `gcp.dataplex.ZoneIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the zone are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.dataplex.ZoneIamPolicy`: Retrieves the IAM policy for the zone
    /// 
    /// &gt; **Note:** `gcp.dataplex.ZoneIamPolicy` **cannot** be used in conjunction with `gcp.dataplex.ZoneIamBinding` and `gcp.dataplex.ZoneIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.dataplex.ZoneIamBinding` resources **can be** used in conjunction with `gcp.dataplex.ZoneIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_dataplex\_zone\_iam\_policy
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     var policy = new Gcp.DataPlex.ZoneIamPolicy("policy", new()
    ///     {
    ///         Project = example.Project,
    ///         Location = example.Location,
    ///         Lake = example.Lake,
    ///         DataplexZone = example.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_dataplex\_zone\_iam\_binding
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.DataPlex.ZoneIamBinding("binding", new()
    ///     {
    ///         Project = example.Project,
    ///         Location = example.Location,
    ///         Lake = example.Lake,
    ///         DataplexZone = example.Name,
    ///         Role = "roles/viewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_dataplex\_zone\_iam\_member
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.DataPlex.ZoneIamMember("member", new()
    ///     {
    ///         Project = example.Project,
    ///         Location = example.Location,
    ///         Lake = example.Lake,
    ///         DataplexZone = example.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_dataplex\_zone\_iam\_policy
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     var policy = new Gcp.DataPlex.ZoneIamPolicy("policy", new()
    ///     {
    ///         Project = example.Project,
    ///         Location = example.Location,
    ///         Lake = example.Lake,
    ///         DataplexZone = example.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_dataplex\_zone\_iam\_binding
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.DataPlex.ZoneIamBinding("binding", new()
    ///     {
    ///         Project = example.Project,
    ///         Location = example.Location,
    ///         Lake = example.Lake,
    ///         DataplexZone = example.Name,
    ///         Role = "roles/viewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_dataplex\_zone\_iam\_member
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.DataPlex.ZoneIamMember("member", new()
    ///     {
    ///         Project = example.Project,
    ///         Location = example.Location,
    ///         Lake = example.Lake,
    ///         DataplexZone = example.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms:
    /// 
    /// * projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
    /// 
    /// * {{project}}/{{location}}/{{lake}}/{{name}}
    /// 
    /// * {{location}}/{{lake}}/{{name}}
    /// 
    /// * {{name}}
    /// 
    /// Any variables not passed in the import command will be taken from the provider configuration.
    /// 
    /// Dataplex zone IAM resources can be imported using the resource identifiers, role, and member.
    /// 
    /// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:dataplex/zoneIamPolicy:ZoneIamPolicy editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{zone}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    /// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:dataplex/zoneIamPolicy:ZoneIamPolicy editor "projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{zone}} roles/viewer"
    /// ```
    /// 
    /// IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:dataplex/zoneIamPolicy:ZoneIamPolicy editor projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{zone}}
    /// ```
    /// 
    /// -&gt; **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    ///  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:dataplex/zoneIamPolicy:ZoneIamPolicy")]
    public partial class ZoneIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
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

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ZoneIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZoneIamPolicy(string name, ZoneIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataplex/zoneIamPolicy:ZoneIamPolicy", name, args ?? new ZoneIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZoneIamPolicy(string name, Input<string> id, ZoneIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataplex/zoneIamPolicy:ZoneIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ZoneIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZoneIamPolicy Get(string name, Input<string> id, ZoneIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ZoneIamPolicy(name, id, state, options);
        }
    }

    public sealed class ZoneIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("dataplexZone", required: true)]
        public Input<string> DataplexZone { get; set; } = null!;

        [Input("lake", required: true)]
        public Input<string> Lake { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ZoneIamPolicyArgs()
        {
        }
        public static new ZoneIamPolicyArgs Empty => new ZoneIamPolicyArgs();
    }

    public sealed class ZoneIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
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

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ZoneIamPolicyState()
        {
        }
        public static new ZoneIamPolicyState Empty => new ZoneIamPolicyState();
    }
}
