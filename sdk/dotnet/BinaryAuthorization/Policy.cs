// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization
{
    /// <summary>
    /// A policy for container image binary authorization.
    /// 
    /// To get more information about Policy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/binary-authorization/)
    /// 
    /// ## Example Usage
    /// ### Binary Authorization Policy Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var note = new Gcp.ContainerAnalysis.Note("note", new()
    ///     {
    ///         AttestationAuthority = new Gcp.ContainerAnalysis.Inputs.NoteAttestationAuthorityArgs
    ///         {
    ///             Hint = new Gcp.ContainerAnalysis.Inputs.NoteAttestationAuthorityHintArgs
    ///             {
    ///                 HumanReadableName = "My attestor",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var attestor = new Gcp.BinaryAuthorization.Attestor("attestor", new()
    ///     {
    ///         AttestationAuthorityNote = new Gcp.BinaryAuthorization.Inputs.AttestorAttestationAuthorityNoteArgs
    ///         {
    ///             NoteReference = note.Name,
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.BinaryAuthorization.Policy("policy", new()
    ///     {
    ///         AdmissionWhitelistPatterns = new[]
    ///         {
    ///             new Gcp.BinaryAuthorization.Inputs.PolicyAdmissionWhitelistPatternArgs
    ///             {
    ///                 NamePattern = "gcr.io/google_containers/*",
    ///             },
    ///         },
    ///         DefaultAdmissionRule = new Gcp.BinaryAuthorization.Inputs.PolicyDefaultAdmissionRuleArgs
    ///         {
    ///             EvaluationMode = "ALWAYS_ALLOW",
    ///             EnforcementMode = "ENFORCED_BLOCK_AND_AUDIT_LOG",
    ///         },
    ///         ClusterAdmissionRules = new[]
    ///         {
    ///             new Gcp.BinaryAuthorization.Inputs.PolicyClusterAdmissionRuleArgs
    ///             {
    ///                 Cluster = "us-central1-a.prod-cluster",
    ///                 EvaluationMode = "REQUIRE_ATTESTATION",
    ///                 EnforcementMode = "ENFORCED_BLOCK_AND_AUDIT_LOG",
    ///                 RequireAttestationsBies = new[]
    ///                 {
    ///                     attestor.Name,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Binary Authorization Policy Global Evaluation
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var note = new Gcp.ContainerAnalysis.Note("note", new()
    ///     {
    ///         AttestationAuthority = new Gcp.ContainerAnalysis.Inputs.NoteAttestationAuthorityArgs
    ///         {
    ///             Hint = new Gcp.ContainerAnalysis.Inputs.NoteAttestationAuthorityHintArgs
    ///             {
    ///                 HumanReadableName = "My attestor",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var attestor = new Gcp.BinaryAuthorization.Attestor("attestor", new()
    ///     {
    ///         AttestationAuthorityNote = new Gcp.BinaryAuthorization.Inputs.AttestorAttestationAuthorityNoteArgs
    ///         {
    ///             NoteReference = note.Name,
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.BinaryAuthorization.Policy("policy", new()
    ///     {
    ///         DefaultAdmissionRule = new Gcp.BinaryAuthorization.Inputs.PolicyDefaultAdmissionRuleArgs
    ///         {
    ///             EvaluationMode = "REQUIRE_ATTESTATION",
    ///             EnforcementMode = "ENFORCED_BLOCK_AND_AUDIT_LOG",
    ///             RequireAttestationsBies = new[]
    ///             {
    ///                 attestor.Name,
    ///             },
    ///         },
    ///         GlobalPolicyEvaluationMode = "ENABLE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Policy can be imported using any of these accepted formats* `projects/{{project}}` * `{{project}}` When using the `pulumi import` command, Policy can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:binaryauthorization/policy:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A whitelist of image patterns to exclude from admission rules. If an
        /// image's name matches a whitelist pattern, the image's admission
        /// requests will always be permitted regardless of your admission rules.
        /// Structure is documented below.
        /// </summary>
        [Output("admissionWhitelistPatterns")]
        public Output<ImmutableArray<Outputs.PolicyAdmissionWhitelistPattern>> AdmissionWhitelistPatterns { get; private set; } = null!;

        /// <summary>
        /// Per-cluster admission rules. An admission rule specifies either that
        /// all container images used in a pod creation request must be attested
        /// to by one or more attestors, that all pod creations will be allowed,
        /// or that all pod creations will be denied. There can be at most one
        /// admission rule per cluster spec.
        /// 
        /// Identifier format: `{{location}}.{{clusterId}}`.
        /// A location is either a compute zone (e.g. `us-central1-a`) or a region
        /// (e.g. `us-central1`).
        /// Structure is documented below.
        /// </summary>
        [Output("clusterAdmissionRules")]
        public Output<ImmutableArray<Outputs.PolicyClusterAdmissionRule>> ClusterAdmissionRules { get; private set; } = null!;

        /// <summary>
        /// Default admission rule for a cluster without a per-cluster admission
        /// rule.
        /// Structure is documented below.
        /// </summary>
        [Output("defaultAdmissionRule")]
        public Output<Outputs.PolicyDefaultAdmissionRule> DefaultAdmissionRule { get; private set; } = null!;

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Controls the evaluation of a Google-maintained global admission policy
        /// for common system-level images. Images not covered by the global
        /// policy will be subject to the project admission policy.
        /// Possible values are: `ENABLE`, `DISABLE`.
        /// </summary>
        [Output("globalPolicyEvaluationMode")]
        public Output<string> GlobalPolicyEvaluationMode { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:binaryauthorization/policy:Policy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("admissionWhitelistPatterns")]
        private InputList<Inputs.PolicyAdmissionWhitelistPatternArgs>? _admissionWhitelistPatterns;

        /// <summary>
        /// A whitelist of image patterns to exclude from admission rules. If an
        /// image's name matches a whitelist pattern, the image's admission
        /// requests will always be permitted regardless of your admission rules.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PolicyAdmissionWhitelistPatternArgs> AdmissionWhitelistPatterns
        {
            get => _admissionWhitelistPatterns ?? (_admissionWhitelistPatterns = new InputList<Inputs.PolicyAdmissionWhitelistPatternArgs>());
            set => _admissionWhitelistPatterns = value;
        }

        [Input("clusterAdmissionRules")]
        private InputList<Inputs.PolicyClusterAdmissionRuleArgs>? _clusterAdmissionRules;

        /// <summary>
        /// Per-cluster admission rules. An admission rule specifies either that
        /// all container images used in a pod creation request must be attested
        /// to by one or more attestors, that all pod creations will be allowed,
        /// or that all pod creations will be denied. There can be at most one
        /// admission rule per cluster spec.
        /// 
        /// Identifier format: `{{location}}.{{clusterId}}`.
        /// A location is either a compute zone (e.g. `us-central1-a`) or a region
        /// (e.g. `us-central1`).
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PolicyClusterAdmissionRuleArgs> ClusterAdmissionRules
        {
            get => _clusterAdmissionRules ?? (_clusterAdmissionRules = new InputList<Inputs.PolicyClusterAdmissionRuleArgs>());
            set => _clusterAdmissionRules = value;
        }

        /// <summary>
        /// Default admission rule for a cluster without a per-cluster admission
        /// rule.
        /// Structure is documented below.
        /// </summary>
        [Input("defaultAdmissionRule", required: true)]
        public Input<Inputs.PolicyDefaultAdmissionRuleArgs> DefaultAdmissionRule { get; set; } = null!;

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Controls the evaluation of a Google-maintained global admission policy
        /// for common system-level images. Images not covered by the global
        /// policy will be subject to the project admission policy.
        /// Possible values are: `ENABLE`, `DISABLE`.
        /// </summary>
        [Input("globalPolicyEvaluationMode")]
        public Input<string>? GlobalPolicyEvaluationMode { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }

    public sealed class PolicyState : global::Pulumi.ResourceArgs
    {
        [Input("admissionWhitelistPatterns")]
        private InputList<Inputs.PolicyAdmissionWhitelistPatternGetArgs>? _admissionWhitelistPatterns;

        /// <summary>
        /// A whitelist of image patterns to exclude from admission rules. If an
        /// image's name matches a whitelist pattern, the image's admission
        /// requests will always be permitted regardless of your admission rules.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PolicyAdmissionWhitelistPatternGetArgs> AdmissionWhitelistPatterns
        {
            get => _admissionWhitelistPatterns ?? (_admissionWhitelistPatterns = new InputList<Inputs.PolicyAdmissionWhitelistPatternGetArgs>());
            set => _admissionWhitelistPatterns = value;
        }

        [Input("clusterAdmissionRules")]
        private InputList<Inputs.PolicyClusterAdmissionRuleGetArgs>? _clusterAdmissionRules;

        /// <summary>
        /// Per-cluster admission rules. An admission rule specifies either that
        /// all container images used in a pod creation request must be attested
        /// to by one or more attestors, that all pod creations will be allowed,
        /// or that all pod creations will be denied. There can be at most one
        /// admission rule per cluster spec.
        /// 
        /// Identifier format: `{{location}}.{{clusterId}}`.
        /// A location is either a compute zone (e.g. `us-central1-a`) or a region
        /// (e.g. `us-central1`).
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.PolicyClusterAdmissionRuleGetArgs> ClusterAdmissionRules
        {
            get => _clusterAdmissionRules ?? (_clusterAdmissionRules = new InputList<Inputs.PolicyClusterAdmissionRuleGetArgs>());
            set => _clusterAdmissionRules = value;
        }

        /// <summary>
        /// Default admission rule for a cluster without a per-cluster admission
        /// rule.
        /// Structure is documented below.
        /// </summary>
        [Input("defaultAdmissionRule")]
        public Input<Inputs.PolicyDefaultAdmissionRuleGetArgs>? DefaultAdmissionRule { get; set; }

        /// <summary>
        /// A descriptive comment.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Controls the evaluation of a Google-maintained global admission policy
        /// for common system-level images. Images not covered by the global
        /// policy will be subject to the project admission policy.
        /// Possible values are: `ENABLE`, `DISABLE`.
        /// </summary>
        [Input("globalPolicyEvaluationMode")]
        public Input<string>? GlobalPolicyEvaluationMode { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public PolicyState()
        {
        }
        public static new PolicyState Empty => new PolicyState();
    }
}
