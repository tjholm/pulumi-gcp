// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority
{
    /// <summary>
    /// Certificate Authority Service provides reusable and parameterized templates that you can use for common certificate issuance scenarios. A certificate template represents a relatively static and well-defined certificate issuance schema within an organization. A certificate template can essentially become a full-fledged vertical certificate issuance framework.
    /// 
    /// To get more information about CertificateTemplate, see:
    /// 
    /// * [API documentation](https://cloud.google.com/certificate-authority-service/docs/reference/rest)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/certificate-authority-service)
    ///     * [Understanding Certificate Templates](https://cloud.google.com/certificate-authority-service/docs/certificate-template)
    ///     * [Common configurations and Certificate Profiles](https://cloud.google.com/certificate-authority-service/docs/certificate-profile)
    /// 
    /// ## Example Usage
    /// 
    /// ### Privateca Template Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.CertificateAuthority.CertificateTemplate("default", new()
    ///     {
    ///         Name = "my-template",
    ///         Location = "us-central1",
    ///         Description = "A sample certificate template",
    ///         IdentityConstraints = new Gcp.CertificateAuthority.Inputs.CertificateTemplateIdentityConstraintsArgs
    ///         {
    ///             AllowSubjectAltNamesPassthrough = true,
    ///             AllowSubjectPassthrough = true,
    ///             CelExpression = new Gcp.CertificateAuthority.Inputs.CertificateTemplateIdentityConstraintsCelExpressionArgs
    ///             {
    ///                 Description = "Always true",
    ///                 Expression = "true",
    ///                 Location = "any.file.anywhere",
    ///                 Title = "Sample expression",
    ///             },
    ///         },
    ///         MaximumLifetime = "86400s",
    ///         PassthroughExtensions = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePassthroughExtensionsArgs
    ///         {
    ///             AdditionalExtensions = new[]
    ///             {
    ///                 new Gcp.CertificateAuthority.Inputs.CertificateTemplatePassthroughExtensionsAdditionalExtensionArgs
    ///                 {
    ///                     ObjectIdPaths = new[]
    ///                     {
    ///                         1,
    ///                         6,
    ///                     },
    ///                 },
    ///             },
    ///             KnownExtensions = new[]
    ///             {
    ///                 "EXTENDED_KEY_USAGE",
    ///             },
    ///         },
    ///         PredefinedValues = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesArgs
    ///         {
    ///             AdditionalExtensions = new[]
    ///             {
    ///                 new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesAdditionalExtensionArgs
    ///                 {
    ///                     ObjectId = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesAdditionalExtensionObjectIdArgs
    ///                     {
    ///                         ObjectIdPaths = new[]
    ///                         {
    ///                             1,
    ///                             6,
    ///                         },
    ///                     },
    ///                     Value = "c3RyaW5nCg==",
    ///                     Critical = true,
    ///                 },
    ///             },
    ///             AiaOcspServers = new[]
    ///             {
    ///                 "string",
    ///             },
    ///             CaOptions = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesCaOptionsArgs
    ///             {
    ///                 IsCa = false,
    ///                 MaxIssuerPathLength = 6,
    ///             },
    ///             KeyUsage = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesKeyUsageArgs
    ///             {
    ///                 BaseKeyUsage = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageArgs
    ///                 {
    ///                     CertSign = false,
    ///                     ContentCommitment = true,
    ///                     CrlSign = false,
    ///                     DataEncipherment = true,
    ///                     DecipherOnly = true,
    ///                     DigitalSignature = true,
    ///                     EncipherOnly = true,
    ///                     KeyAgreement = true,
    ///                     KeyEncipherment = true,
    ///                 },
    ///                 ExtendedKeyUsage = new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageArgs
    ///                 {
    ///                     ClientAuth = true,
    ///                     CodeSigning = true,
    ///                     EmailProtection = true,
    ///                     OcspSigning = true,
    ///                     ServerAuth = true,
    ///                     TimeStamping = true,
    ///                 },
    ///                 UnknownExtendedKeyUsages = new[]
    ///                 {
    ///                     new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsageArgs
    ///                     {
    ///                         ObjectIdPaths = new[]
    ///                         {
    ///                             1,
    ///                             6,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///             PolicyIds = new[]
    ///             {
    ///                 new Gcp.CertificateAuthority.Inputs.CertificateTemplatePredefinedValuesPolicyIdArgs
    ///                 {
    ///                     ObjectIdPaths = new[]
    ///                     {
    ///                         1,
    ///                         6,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "label-one", "value-one" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CertificateTemplate can be imported using any of these accepted formats:
    /// 
    /// * `projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}`
    /// 
    /// * `{{project}}/{{location}}/{{name}}`
    /// 
    /// * `{{location}}/{{name}}`
    /// 
    /// When using the `pulumi import` command, CertificateTemplate can be imported using one of the formats above. For example:
    /// 
    /// ```sh
    /// $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gcp:certificateauthority/certificateTemplate:CertificateTemplate default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:certificateauthority/certificateTemplate:CertificateTemplate")]
    public partial class CertificateTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. The time at which this CertificateTemplate was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. A human-readable description of scenarios this template is intended for.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
        /// Structure is documented below.
        /// </summary>
        [Output("identityConstraints")]
        public Output<Outputs.CertificateTemplateIdentityConstraints?> IdentityConstraints { get; private set; } = null!;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool's IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it.
        /// </summary>
        [Output("maximumLifetime")]
        public Output<string?> MaximumLifetime { get; private set; } = null!;

        /// <summary>
        /// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
        /// Structure is documented below.
        /// </summary>
        [Output("passthroughExtensions")]
        public Output<Outputs.CertificateTemplatePassthroughExtensions?> PassthroughExtensions { get; private set; } = null!;

        /// <summary>
        /// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
        /// Structure is documented below.
        /// </summary>
        [Output("predefinedValues")]
        public Output<Outputs.CertificateTemplatePredefinedValues?> PredefinedValues { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Output only. The time at which this CertificateTemplate was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateTemplate(string name, CertificateTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:certificateauthority/certificateTemplate:CertificateTemplate", name, args ?? new CertificateTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateTemplate(string name, Input<string> id, CertificateTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:certificateauthority/certificateTemplate:CertificateTemplate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateTemplate Get(string name, Input<string> id, CertificateTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateTemplate(name, id, state, options);
        }
    }

    public sealed class CertificateTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A human-readable description of scenarios this template is intended for.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
        /// Structure is documented below.
        /// </summary>
        [Input("identityConstraints")]
        public Input<Inputs.CertificateTemplateIdentityConstraintsArgs>? IdentityConstraints { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Optional. The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool's IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it.
        /// </summary>
        [Input("maximumLifetime")]
        public Input<string>? MaximumLifetime { get; set; }

        /// <summary>
        /// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
        /// Structure is documented below.
        /// </summary>
        [Input("passthroughExtensions")]
        public Input<Inputs.CertificateTemplatePassthroughExtensionsArgs>? PassthroughExtensions { get; set; }

        /// <summary>
        /// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
        /// Structure is documented below.
        /// </summary>
        [Input("predefinedValues")]
        public Input<Inputs.CertificateTemplatePredefinedValuesArgs>? PredefinedValues { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public CertificateTemplateArgs()
        {
        }
        public static new CertificateTemplateArgs Empty => new CertificateTemplateArgs();
    }

    public sealed class CertificateTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. The time at which this CertificateTemplate was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Optional. A human-readable description of scenarios this template is intended for.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Optional. Describes constraints on identities that may be appear in Certificates issued using this template. If this is omitted, then this template will not add restrictions on a certificate's identity.
        /// Structure is documented below.
        /// </summary>
        [Input("identityConstraints")]
        public Input<Inputs.CertificateTemplateIdentityConstraintsGetArgs>? IdentityConstraints { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Labels with user-defined metadata.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. The maximum lifetime allowed for all issued certificates that use this template. If the issuing CaPool's IssuancePolicy specifies a maximum lifetime the minimum of the two durations will be the maximum lifetime for issued. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it.
        /// </summary>
        [Input("maximumLifetime")]
        public Input<string>? MaximumLifetime { get; set; }

        /// <summary>
        /// The resource name for this CertificateTemplate in the format `projects/*/locations/*/certificateTemplates/*`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. Describes the set of X.509 extensions that may appear in a Certificate issued using this CertificateTemplate. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If the issuing CaPool's IssuancePolicy defines baseline_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this template will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CertificateTemplate's predefined_values.
        /// Structure is documented below.
        /// </summary>
        [Input("passthroughExtensions")]
        public Input<Inputs.CertificateTemplatePassthroughExtensionsGetArgs>? PassthroughExtensions { get; set; }

        /// <summary>
        /// Optional. A set of X.509 values that will be applied to all issued certificates that use this template. If the certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If the issuing CaPool's IssuancePolicy defines conflicting baseline_values for the same properties, the certificate issuance request will fail.
        /// Structure is documented below.
        /// </summary>
        [Input("predefinedValues")]
        public Input<Inputs.CertificateTemplatePredefinedValuesGetArgs>? PredefinedValues { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Output only. The time at which this CertificateTemplate was updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public CertificateTemplateState()
        {
        }
        public static new CertificateTemplateState Empty => new CertificateTemplateState();
    }
}
