// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkSecurity
{
    /// <summary>
    /// ## Example Usage
    /// ### Network Security Client Tls Policy Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkSecurity.ClientTlsPolicy("default", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///         Sni = "secure.example.com",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Network Security Client Tls Policy Advanced
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkSecurity.ClientTlsPolicy("default", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Description = "my description",
    ///         ClientCertificate = new Gcp.NetworkSecurity.Inputs.ClientTlsPolicyClientCertificateArgs
    ///         {
    ///             CertificateProviderInstance = new Gcp.NetworkSecurity.Inputs.ClientTlsPolicyClientCertificateCertificateProviderInstanceArgs
    ///             {
    ///                 PluginInstance = "google_cloud_private_spiffe",
    ///             },
    ///         },
    ///         ServerValidationCas = new[]
    ///         {
    ///             new Gcp.NetworkSecurity.Inputs.ClientTlsPolicyServerValidationCaArgs
    ///             {
    ///                 GrpcEndpoint = new Gcp.NetworkSecurity.Inputs.ClientTlsPolicyServerValidationCaGrpcEndpointArgs
    ///                 {
    ///                     TargetUri = "unix:mypath",
    ///                 },
    ///             },
    ///             new Gcp.NetworkSecurity.Inputs.ClientTlsPolicyServerValidationCaArgs
    ///             {
    ///                 GrpcEndpoint = new Gcp.NetworkSecurity.Inputs.ClientTlsPolicyServerValidationCaGrpcEndpointArgs
    ///                 {
    ///                     TargetUri = "unix:mypath1",
    ///                 },
    ///             },
    ///         },
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
    /// ClientTlsPolicy can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, ClientTlsPolicy can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy default projects/{{project}}/locations/{{location}}/clientTlsPolicies/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy")]
    public partial class ClientTlsPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
        /// Structure is documented below.
        /// </summary>
        [Output("clientCertificate")]
        public Output<Outputs.ClientTlsPolicyClientCertificate?> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// Time the ClientTlsPolicy was created in UTC.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the ClientTlsPolicy resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location of the client tls policy.
        /// The default value is `global`.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the ClientTlsPolicy resource.
        /// 
        /// 
        /// - - -
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
        /// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
        /// Structure is documented below.
        /// </summary>
        [Output("serverValidationCas")]
        public Output<ImmutableArray<Outputs.ClientTlsPolicyServerValidationCa>> ServerValidationCas { get; private set; } = null!;

        /// <summary>
        /// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
        /// </summary>
        [Output("sni")]
        public Output<string?> Sni { get; private set; } = null!;

        /// <summary>
        /// Time the ClientTlsPolicy was updated in UTC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ClientTlsPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientTlsPolicy(string name, ClientTlsPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy", name, args ?? new ClientTlsPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientTlsPolicy(string name, Input<string> id, ClientTlsPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networksecurity/clientTlsPolicy:ClientTlsPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientTlsPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientTlsPolicy Get(string name, Input<string> id, ClientTlsPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientTlsPolicy(name, id, state, options);
        }
    }

    public sealed class ClientTlsPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
        /// Structure is documented below.
        /// </summary>
        [Input("clientCertificate")]
        public Input<Inputs.ClientTlsPolicyClientCertificateArgs>? ClientCertificate { get; set; }

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the ClientTlsPolicy resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location of the client tls policy.
        /// The default value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the ClientTlsPolicy resource.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serverValidationCas")]
        private InputList<Inputs.ClientTlsPolicyServerValidationCaArgs>? _serverValidationCas;

        /// <summary>
        /// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClientTlsPolicyServerValidationCaArgs> ServerValidationCas
        {
            get => _serverValidationCas ?? (_serverValidationCas = new InputList<Inputs.ClientTlsPolicyServerValidationCaArgs>());
            set => _serverValidationCas = value;
        }

        /// <summary>
        /// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
        /// </summary>
        [Input("sni")]
        public Input<string>? Sni { get; set; }

        public ClientTlsPolicyArgs()
        {
        }
        public static new ClientTlsPolicyArgs Empty => new ClientTlsPolicyArgs();
    }

    public sealed class ClientTlsPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
        /// Structure is documented below.
        /// </summary>
        [Input("clientCertificate")]
        public Input<Inputs.ClientTlsPolicyClientCertificateGetArgs>? ClientCertificate { get; set; }

        /// <summary>
        /// Time the ClientTlsPolicy was created in UTC.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
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

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the ClientTlsPolicy resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location of the client tls policy.
        /// The default value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the ClientTlsPolicy resource.
        /// 
        /// 
        /// - - -
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

        [Input("serverValidationCas")]
        private InputList<Inputs.ClientTlsPolicyServerValidationCaGetArgs>? _serverValidationCas;

        /// <summary>
        /// Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ClientTlsPolicyServerValidationCaGetArgs> ServerValidationCas
        {
            get => _serverValidationCas ?? (_serverValidationCas = new InputList<Inputs.ClientTlsPolicyServerValidationCaGetArgs>());
            set => _serverValidationCas = value;
        }

        /// <summary>
        /// Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
        /// </summary>
        [Input("sni")]
        public Input<string>? Sni { get; set; }

        /// <summary>
        /// Time the ClientTlsPolicy was updated in UTC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ClientTlsPolicyState()
        {
        }
        public static new ClientTlsPolicyState Empty => new ClientTlsPolicyState();
    }
}
