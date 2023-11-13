// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    /// <summary>
    /// A `CryptoKey` represents a logical key that can be used for cryptographic operations.
    /// 
    /// &gt; **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
    /// Destroying a provider-managed CryptoKey will remove it from state
    /// and delete all CryptoKeyVersions, rendering the key unusable, but *will
    /// not delete the resource from the project.* When the provider destroys these keys,
    /// any data previously encrypted with these keys will be irrecoverable.
    /// For this reason, it is strongly recommended that you add lifecycle hooks
    /// to the resource to prevent accidental destruction.
    /// 
    /// To get more information about CryptoKey, see:
    /// 
    /// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
    /// * How-to Guides
    ///     * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)
    /// 
    /// ## Example Usage
    /// ### Kms Crypto Key Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyring = new Gcp.Kms.KeyRing("keyring", new()
    ///     {
    ///         Location = "global",
    ///     });
    /// 
    ///     var example_key = new Gcp.Kms.CryptoKey("example-key", new()
    ///     {
    ///         KeyRing = keyring.Id,
    ///         RotationPeriod = "100000s",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Kms Crypto Key Asymmetric Sign
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var keyring = new Gcp.Kms.KeyRing("keyring", new()
    ///     {
    ///         Location = "global",
    ///     });
    /// 
    ///     var example_asymmetric_sign_key = new Gcp.Kms.CryptoKey("example-asymmetric-sign-key", new()
    ///     {
    ///         KeyRing = keyring.Id,
    ///         Purpose = "ASYMMETRIC_SIGN",
    ///         VersionTemplate = new Gcp.Kms.Inputs.CryptoKeyVersionTemplateArgs
    ///         {
    ///             Algorithm = "EC_SIGN_P384_SHA384",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CryptoKey can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:kms/cryptoKey:CryptoKey")]
    public partial class CryptoKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
        /// If not specified at creation time, the default duration is 24 hours.
        /// </summary>
        [Output("destroyScheduledDuration")]
        public Output<string> DestroyScheduledDuration { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Whether this key may contain imported versions only.
        /// </summary>
        [Output("importOnly")]
        public Output<bool> ImportOnly { get; private set; } = null!;

        /// <summary>
        /// The KeyRing that this key belongs to.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("keyRing")]
        public Output<string> KeyRing { get; private set; } = null!;

        /// <summary>
        /// Labels with user-defined metadata to apply to this resource.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name for the CryptoKey.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The immutable purpose of this CryptoKey. See the
        /// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        /// for possible inputs.
        /// Default value is "ENCRYPT_DECRYPT".
        /// </summary>
        [Output("purpose")]
        public Output<string?> Purpose { get; private set; } = null!;

        /// <summary>
        /// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        /// The first rotation will take place after the specified period. The rotation period has
        /// the format of a decimal number with up to 9 fractional digits, followed by the
        /// letter `s` (seconds). It must be greater than a day (ie, 86400).
        /// </summary>
        [Output("rotationPeriod")]
        public Output<string?> RotationPeriod { get; private set; } = null!;

        /// <summary>
        /// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        /// You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        /// </summary>
        [Output("skipInitialVersionCreation")]
        public Output<bool?> SkipInitialVersionCreation { get; private set; } = null!;

        /// <summary>
        /// A template describing settings for new crypto key versions.
        /// Structure is documented below.
        /// </summary>
        [Output("versionTemplate")]
        public Output<Outputs.CryptoKeyVersionTemplate> VersionTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a CryptoKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CryptoKey(string name, CryptoKeyArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKey:CryptoKey", name, args ?? new CryptoKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CryptoKey(string name, Input<string> id, CryptoKeyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKey:CryptoKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CryptoKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CryptoKey Get(string name, Input<string> id, CryptoKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new CryptoKey(name, id, state, options);
        }
    }

    public sealed class CryptoKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
        /// If not specified at creation time, the default duration is 24 hours.
        /// </summary>
        [Input("destroyScheduledDuration")]
        public Input<string>? DestroyScheduledDuration { get; set; }

        /// <summary>
        /// Whether this key may contain imported versions only.
        /// </summary>
        [Input("importOnly")]
        public Input<bool>? ImportOnly { get; set; }

        /// <summary>
        /// The KeyRing that this key belongs to.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("keyRing", required: true)]
        public Input<string> KeyRing { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels with user-defined metadata to apply to this resource.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the CryptoKey.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The immutable purpose of this CryptoKey. See the
        /// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        /// for possible inputs.
        /// Default value is "ENCRYPT_DECRYPT".
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        /// The first rotation will take place after the specified period. The rotation period has
        /// the format of a decimal number with up to 9 fractional digits, followed by the
        /// letter `s` (seconds). It must be greater than a day (ie, 86400).
        /// </summary>
        [Input("rotationPeriod")]
        public Input<string>? RotationPeriod { get; set; }

        /// <summary>
        /// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        /// You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        /// </summary>
        [Input("skipInitialVersionCreation")]
        public Input<bool>? SkipInitialVersionCreation { get; set; }

        /// <summary>
        /// A template describing settings for new crypto key versions.
        /// Structure is documented below.
        /// </summary>
        [Input("versionTemplate")]
        public Input<Inputs.CryptoKeyVersionTemplateArgs>? VersionTemplate { get; set; }

        public CryptoKeyArgs()
        {
        }
        public static new CryptoKeyArgs Empty => new CryptoKeyArgs();
    }

    public sealed class CryptoKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
        /// If not specified at creation time, the default duration is 24 hours.
        /// </summary>
        [Input("destroyScheduledDuration")]
        public Input<string>? DestroyScheduledDuration { get; set; }

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
        /// Whether this key may contain imported versions only.
        /// </summary>
        [Input("importOnly")]
        public Input<bool>? ImportOnly { get; set; }

        /// <summary>
        /// The KeyRing that this key belongs to.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("keyRing")]
        public Input<string>? KeyRing { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels with user-defined metadata to apply to this resource.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name for the CryptoKey.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// The immutable purpose of this CryptoKey. See the
        /// [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
        /// for possible inputs.
        /// Default value is "ENCRYPT_DECRYPT".
        /// </summary>
        [Input("purpose")]
        public Input<string>? Purpose { get; set; }

        /// <summary>
        /// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
        /// The first rotation will take place after the specified period. The rotation period has
        /// the format of a decimal number with up to 9 fractional digits, followed by the
        /// letter `s` (seconds). It must be greater than a day (ie, 86400).
        /// </summary>
        [Input("rotationPeriod")]
        public Input<string>? RotationPeriod { get; set; }

        /// <summary>
        /// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
        /// You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
        /// </summary>
        [Input("skipInitialVersionCreation")]
        public Input<bool>? SkipInitialVersionCreation { get; set; }

        /// <summary>
        /// A template describing settings for new crypto key versions.
        /// Structure is documented below.
        /// </summary>
        [Input("versionTemplate")]
        public Input<Inputs.CryptoKeyVersionTemplateGetArgs>? VersionTemplate { get; set; }

        public CryptoKeyState()
        {
        }
        public static new CryptoKeyState Empty => new CryptoKeyState();
    }
}
