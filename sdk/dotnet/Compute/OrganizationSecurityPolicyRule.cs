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
    /// A rule for the OrganizationSecurityPolicy.
    /// 
    /// To get more information about OrganizationSecurityPolicyRule, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/organizationSecurityPolicies/addRule)
    /// * How-to Guides
    ///     * [Creating firewall rules](https://cloud.google.com/vpc/docs/using-firewall-policies#create-rules)
    /// 
    /// ## Example Usage
    /// ### Organization Security Policy Rule Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var policyOrganizationSecurityPolicy = new Gcp.Compute.OrganizationSecurityPolicy("policyOrganizationSecurityPolicy", new()
    ///     {
    ///         DisplayName = "tf-test",
    ///         Parent = "organizations/123456789",
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
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OrganizationSecurityPolicyRule can be imported using any of these accepted formats:
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule default {{policy_id}}/priority/{{priority}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule")]
    public partial class OrganizationSecurityPolicyRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Action to perform when the client connection triggers the rule. Can currently be either
        /// "allow", "deny" or "goto_next".
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// A description of the rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The direction in which this rule applies. If unspecified an INGRESS rule is created.
        /// Possible values are: `INGRESS`, `EGRESS`.
        /// </summary>
        [Output("direction")]
        public Output<string?> Direction { get; private set; } = null!;

        /// <summary>
        /// Denotes whether to enable logging for a particular rule.
        /// If logging is enabled, logs will be exported to the
        /// configured export destination in Stackdriver.
        /// </summary>
        [Output("enableLogging")]
        public Output<bool?> EnableLogging { get; private set; } = null!;

        /// <summary>
        /// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        /// Structure is documented below.
        /// </summary>
        [Output("match")]
        public Output<Outputs.OrganizationSecurityPolicyRuleMatch> Match { get; private set; } = null!;

        /// <summary>
        /// The ID of the OrganizationSecurityPolicy this rule applies to.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// If set to true, the specified action is not enforced.
        /// </summary>
        [Output("preview")]
        public Output<bool?> Preview { get; private set; } = null!;

        /// <summary>
        /// An integer indicating the priority of a rule in the list. The priority must be a value
        /// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
        /// highest priority and 2147483647 is the lowest prority.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// A list of network resource URLs to which this rule applies.
        /// This field allows you to control which network's VMs get
        /// this rule. If this field is left blank, all VMs
        /// within the organization will receive the rule.
        /// </summary>
        [Output("targetResources")]
        public Output<ImmutableArray<string>> TargetResources { get; private set; } = null!;

        /// <summary>
        /// A list of service accounts indicating the sets of
        /// instances that are applied with this rule.
        /// </summary>
        [Output("targetServiceAccounts")]
        public Output<ImmutableArray<string>> TargetServiceAccounts { get; private set; } = null!;


        /// <summary>
        /// Create a OrganizationSecurityPolicyRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrganizationSecurityPolicyRule(string name, OrganizationSecurityPolicyRuleArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule", name, args ?? new OrganizationSecurityPolicyRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrganizationSecurityPolicyRule(string name, Input<string> id, OrganizationSecurityPolicyRuleState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/organizationSecurityPolicyRule:OrganizationSecurityPolicyRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrganizationSecurityPolicyRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrganizationSecurityPolicyRule Get(string name, Input<string> id, OrganizationSecurityPolicyRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new OrganizationSecurityPolicyRule(name, id, state, options);
        }
    }

    public sealed class OrganizationSecurityPolicyRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Action to perform when the client connection triggers the rule. Can currently be either
        /// "allow", "deny" or "goto_next".
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// A description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The direction in which this rule applies. If unspecified an INGRESS rule is created.
        /// Possible values are: `INGRESS`, `EGRESS`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Denotes whether to enable logging for a particular rule.
        /// If logging is enabled, logs will be exported to the
        /// configured export destination in Stackdriver.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        /// Structure is documented below.
        /// </summary>
        [Input("match", required: true)]
        public Input<Inputs.OrganizationSecurityPolicyRuleMatchArgs> Match { get; set; } = null!;

        /// <summary>
        /// The ID of the OrganizationSecurityPolicy this rule applies to.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// If set to true, the specified action is not enforced.
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// An integer indicating the priority of a rule in the list. The priority must be a value
        /// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
        /// highest priority and 2147483647 is the lowest prority.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("targetResources")]
        private InputList<string>? _targetResources;

        /// <summary>
        /// A list of network resource URLs to which this rule applies.
        /// This field allows you to control which network's VMs get
        /// this rule. If this field is left blank, all VMs
        /// within the organization will receive the rule.
        /// </summary>
        public InputList<string> TargetResources
        {
            get => _targetResources ?? (_targetResources = new InputList<string>());
            set => _targetResources = value;
        }

        [Input("targetServiceAccounts")]
        private InputList<string>? _targetServiceAccounts;

        /// <summary>
        /// A list of service accounts indicating the sets of
        /// instances that are applied with this rule.
        /// </summary>
        public InputList<string> TargetServiceAccounts
        {
            get => _targetServiceAccounts ?? (_targetServiceAccounts = new InputList<string>());
            set => _targetServiceAccounts = value;
        }

        public OrganizationSecurityPolicyRuleArgs()
        {
        }
        public static new OrganizationSecurityPolicyRuleArgs Empty => new OrganizationSecurityPolicyRuleArgs();
    }

    public sealed class OrganizationSecurityPolicyRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Action to perform when the client connection triggers the rule. Can currently be either
        /// "allow", "deny" or "goto_next".
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The direction in which this rule applies. If unspecified an INGRESS rule is created.
        /// Possible values are: `INGRESS`, `EGRESS`.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Denotes whether to enable logging for a particular rule.
        /// If logging is enabled, logs will be exported to the
        /// configured export destination in Stackdriver.
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
        /// Structure is documented below.
        /// </summary>
        [Input("match")]
        public Input<Inputs.OrganizationSecurityPolicyRuleMatchGetArgs>? Match { get; set; }

        /// <summary>
        /// The ID of the OrganizationSecurityPolicy this rule applies to.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// If set to true, the specified action is not enforced.
        /// </summary>
        [Input("preview")]
        public Input<bool>? Preview { get; set; }

        /// <summary>
        /// An integer indicating the priority of a rule in the list. The priority must be a value
        /// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
        /// highest priority and 2147483647 is the lowest prority.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("targetResources")]
        private InputList<string>? _targetResources;

        /// <summary>
        /// A list of network resource URLs to which this rule applies.
        /// This field allows you to control which network's VMs get
        /// this rule. If this field is left blank, all VMs
        /// within the organization will receive the rule.
        /// </summary>
        public InputList<string> TargetResources
        {
            get => _targetResources ?? (_targetResources = new InputList<string>());
            set => _targetResources = value;
        }

        [Input("targetServiceAccounts")]
        private InputList<string>? _targetServiceAccounts;

        /// <summary>
        /// A list of service accounts indicating the sets of
        /// instances that are applied with this rule.
        /// </summary>
        public InputList<string> TargetServiceAccounts
        {
            get => _targetServiceAccounts ?? (_targetServiceAccounts = new InputList<string>());
            set => _targetServiceAccounts = value;
        }

        public OrganizationSecurityPolicyRuleState()
        {
        }
        public static new OrganizationSecurityPolicyRuleState Empty => new OrganizationSecurityPolicyRuleState();
    }
}
