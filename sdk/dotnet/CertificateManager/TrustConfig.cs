// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateManager
{
    /// <summary>
    /// TrustConfig represents a resource that represents your Public Key Infrastructure (PKI) configuration in Certificate Manager for use in mutual TLS authentication scenarios.
    /// 
    /// To get more information about TrustConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/certificate-manager/docs/reference/certificate-manager/rest/v1/projects.locations.trustConfigs/create)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/certificate-manager/docs)
    /// 
    /// &gt; **Warning:** All arguments including the following potentially sensitive
    /// values will be stored in the raw state as plain text: `trust_stores.trust_stores.trust_anchors.trust_anchors.pem_certificate`, `trust_stores.trust_stores.intermediate_cas.intermediate_cas.pem_certificate`.
    /// Read more about sensitive data in state.
    /// 
    /// ## Example Usage
    /// ### Certificate Manager Trust Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.CertificateManager.TrustConfig("default", new()
    ///     {
    ///         Description = "sample description for the trust config",
    ///         Location = "us-central1",
    ///         TrustStores = new[]
    ///         {
    ///             new Gcp.CertificateManager.Inputs.TrustConfigTrustStoreArgs
    ///             {
    ///                 TrustAnchors = new[]
    ///                 {
    ///                     new Gcp.CertificateManager.Inputs.TrustConfigTrustStoreTrustAnchorArgs
    ///                     {
    ///                         PemCertificate = File.ReadAllText("test-fixtures/cert.pem"),
    ///                     },
    ///                 },
    ///                 IntermediateCas = new[]
    ///                 {
    ///                     new Gcp.CertificateManager.Inputs.TrustConfigTrustStoreIntermediateCaArgs
    ///                     {
    ///                         PemCertificate = File.ReadAllText("test-fixtures/cert.pem"),
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TrustConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/trustConfig:TrustConfig default projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/trustConfig:TrustConfig default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/trustConfig:TrustConfig default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:certificatemanager/trustConfig:TrustConfig")]
    public partial class TrustConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation timestamp of a TrustConfig.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// One or more paragraphs of text description of a trust config.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the trust config.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The trust config location.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// A user-defined name of the trust config. Trust config names must be unique globally.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Set of trust stores to perform validation against.
        /// This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.
        /// Structure is documented below.
        /// </summary>
        [Output("trustStores")]
        public Output<ImmutableArray<Outputs.TrustConfigTrustStore>> TrustStores { get; private set; } = null!;

        /// <summary>
        /// The last update timestamp of a TrustConfig.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a TrustConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrustConfig(string name, TrustConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:certificatemanager/trustConfig:TrustConfig", name, args ?? new TrustConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrustConfig(string name, Input<string> id, TrustConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:certificatemanager/trustConfig:TrustConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TrustConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrustConfig Get(string name, Input<string> id, TrustConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TrustConfig(name, id, state, options);
        }
    }

    public sealed class TrustConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One or more paragraphs of text description of a trust config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the trust config.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The trust config location.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// A user-defined name of the trust config. Trust config names must be unique globally.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("trustStores")]
        private InputList<Inputs.TrustConfigTrustStoreArgs>? _trustStores;

        /// <summary>
        /// Set of trust stores to perform validation against.
        /// This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.TrustConfigTrustStoreArgs> TrustStores
        {
            get => _trustStores ?? (_trustStores = new InputList<Inputs.TrustConfigTrustStoreArgs>());
            set => _trustStores = value;
        }

        public TrustConfigArgs()
        {
        }
        public static new TrustConfigArgs Empty => new TrustConfigArgs();
    }

    public sealed class TrustConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation timestamp of a TrustConfig.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// One or more paragraphs of text description of a trust config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
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

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the trust config.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The trust config location.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A user-defined name of the trust config. Trust config names must be unique globally.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        [Input("trustStores")]
        private InputList<Inputs.TrustConfigTrustStoreGetArgs>? _trustStores;

        /// <summary>
        /// Set of trust stores to perform validation against.
        /// This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.TrustConfigTrustStoreGetArgs> TrustStores
        {
            get => _trustStores ?? (_trustStores = new InputList<Inputs.TrustConfigTrustStoreGetArgs>());
            set => _trustStores = value;
        }

        /// <summary>
        /// The last update timestamp of a TrustConfig.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public TrustConfigState()
        {
        }
        public static new TrustConfigState Empty => new TrustConfigState();
    }
}
