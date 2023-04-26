// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// An association for the OrganizationSecurityPolicy.
    /// 
    /// To get more information about OrganizationSecurityPolicyAssociation, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addAssociation)
    /// * How-to Guides
    ///     * [Associating a policy with the organization or folder](https://cloud.google.com/vpc/docs/using-firewall-policies#associate)
    /// 
    /// ## Example Usage
    /// ### Organization Security Policy Association Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var securityPolicyTarget = new Gcp.Organizations.Folder("securityPolicyTarget", new()
    ///     {
    ///         DisplayName = "tf-test-secpol",
    ///         Parent = "organizations/123456789",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var policyOrganizationSecurityPolicy = new Gcp.Compute.OrganizationSecurityPolicy("policyOrganizationSecurityPolicy", new()
    ///     {
    ///         DisplayName = "tf-test",
    ///         Parent = securityPolicyTarget.Name,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var policyOrganizationSecurityPolicyRule = new Gcp.Compute.OrganizationSecurityPolicyRule("policyOrganizationSecurityPolicyRule", new()
    ///     {
    ///         PolicyId = policyOrganizationSecurityPolicy.Id,
    ///         Action = "allow",
    ///         Direction = "INGRESS",
    ///         EnableLogging = true,
    ///         Match = new Gcp.Compute.Inputs.OrganizationSecurityPolicyRuleMatchArgs
    ///         {
    ///             Config = new Gcp.Compute.Inputs.OrganizationSecurityPolicyRuleMatchConfigArgs
    ///             {
    ///                 SrcIpRanges = new[]
    ///                 {
    ///                     "192.168.0.0/16",
    ///                     "10.0.0.0/8",
    ///                 },
    ///                 Layer4Configs = new[]
    ///                 {
    ///                     new Gcp.Compute.Inputs.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs
    ///                     {
    ///                         IpProtocol = "tcp",
    ///                         Ports = new[]
    ///                         {
    ///                             "22",
    ///                         },
    ///                     },
    ///                     new Gcp.Compute.Inputs.OrganizationSecurityPolicyRuleMatchConfigLayer4ConfigArgs
    ///                     {
    ///                         IpProtocol = "icmp",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Priority = 100,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var policyOrganizationSecurityPolicyAssociation = new Gcp.Compute.OrganizationSecurityPolicyAssociation("policyOrganizationSecurityPolicyAssociation", new()
    ///     {
    ///         AttachmentId = policyOrganizationSecurityPolicy.Parent,
    ///         PolicyId = policyOrganizationSecurityPolicy.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OrganizationSecurityPolicyAssociation can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation default {{policy_id}}/association/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation")]
    public partial class OrganizationSecurityPolicyAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The resource that the security policy is attached to.
        /// </summary>
        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// The display name of the security policy of the association.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The security policy ID of the association.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSecurityPolicyAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSecurityPolicyAssociation(string name, OrganizationSecurityPolicyAssociationArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation", name, args ?? new OrganizationSecurityPolicyAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSecurityPolicyAssociation(string name, Input<string> id, OrganizationSecurityPolicyAssociationState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/organizationSecurityPolicyAssociation:OrganizationSecurityPolicyAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationSecurityPolicyAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSecurityPolicyAssociation Get(string name, Input<string> id, OrganizationSecurityPolicyAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationSecurityPolicyAssociation(name, id, state, options);
        }
    }

    public sealed class OrganizationSecurityPolicyAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource that the security policy is attached to.
        /// </summary>
        [Input("attachmentId", required: true)]
        public Input<string> AttachmentId { get; set; } = null!;

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The security policy ID of the association.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        public OrganizationSecurityPolicyAssociationArgs()
        {
        }
        public static new OrganizationSecurityPolicyAssociationArgs Empty => new OrganizationSecurityPolicyAssociationArgs();
    }

    public sealed class OrganizationSecurityPolicyAssociationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource that the security policy is attached to.
        /// </summary>
        [Input("attachmentId")]
        public Input<string>? AttachmentId { get; set; }

        /// <summary>
        /// The display name of the security policy of the association.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name for an association.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The security policy ID of the association.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        public OrganizationSecurityPolicyAssociationState()
        {
        }
        public static new OrganizationSecurityPolicyAssociationState Empty => new OrganizationSecurityPolicyAssociationState();
    }
}
