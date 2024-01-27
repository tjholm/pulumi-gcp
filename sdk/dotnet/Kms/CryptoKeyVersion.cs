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
    /// A `CryptoKeyVersion` represents an individual cryptographic key, and the associated key material.
    /// 
    /// Destroying a cryptoKeyVersion will not delete the resource from the project.
    /// 
    /// To get more information about CryptoKeyVersion, see:
    /// 
    /// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions)
    /// * How-to Guides
    ///     * [Creating a key Version](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions/create)
    /// 
    /// ## Example Usage
    /// ### Kms Crypto Key Version Basic
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
    ///     var cryptokey = new Gcp.Kms.CryptoKey("cryptokey", new()
    ///     {
    ///         KeyRing = keyring.Id,
    ///         RotationPeriod = "7776000s",
    ///     });
    /// 
    ///     var example_key = new Gcp.Kms.CryptoKeyVersion("example-key", new()
    ///     {
    ///         CryptoKey = cryptokey.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CryptoKeyVersion can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, CryptoKeyVersion can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:kms/cryptoKeyVersion:CryptoKeyVersion default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:kms/cryptoKeyVersion:CryptoKeyVersion")]
    public partial class CryptoKeyVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
        /// Only provided for key versions with protectionLevel HSM.
        /// Structure is documented below.
        /// </summary>
        [Output("attestations")]
        public Output<ImmutableArray<Outputs.CryptoKeyVersionAttestation>> Attestations { get; private set; } = null!;

        /// <summary>
        /// The name of the cryptoKey associated with the CryptoKeyVersions.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("cryptoKey")]
        public Output<string> CryptoKey { get; private set; } = null!;

        /// <summary>
        /// The time this CryptoKeyVersion key material was generated
        /// </summary>
        [Output("generateTime")]
        public Output<string> GenerateTime { get; private set; } = null!;

        /// <summary>
        /// The resource name for this CryptoKeyVersion.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        /// </summary>
        [Output("protectionLevel")]
        public Output<string> ProtectionLevel { get; private set; } = null!;

        /// <summary>
        /// The current state of the CryptoKeyVersion.
        /// Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a CryptoKeyVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CryptoKeyVersion(string name, CryptoKeyVersionArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKeyVersion:CryptoKeyVersion", name, args ?? new CryptoKeyVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CryptoKeyVersion(string name, Input<string> id, CryptoKeyVersionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/cryptoKeyVersion:CryptoKeyVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CryptoKeyVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CryptoKeyVersion Get(string name, Input<string> id, CryptoKeyVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new CryptoKeyVersion(name, id, state, options);
        }
    }

    public sealed class CryptoKeyVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the cryptoKey associated with the CryptoKeyVersions.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("cryptoKey", required: true)]
        public Input<string> CryptoKey { get; set; } = null!;

        /// <summary>
        /// The current state of the CryptoKeyVersion.
        /// Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public CryptoKeyVersionArgs()
        {
        }
        public static new CryptoKeyVersionArgs Empty => new CryptoKeyVersionArgs();
    }

    public sealed class CryptoKeyVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        [Input("attestations")]
        private InputList<Inputs.CryptoKeyVersionAttestationGetArgs>? _attestations;

        /// <summary>
        /// Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
        /// Only provided for key versions with protectionLevel HSM.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.CryptoKeyVersionAttestationGetArgs> Attestations
        {
            get => _attestations ?? (_attestations = new InputList<Inputs.CryptoKeyVersionAttestationGetArgs>());
            set => _attestations = value;
        }

        /// <summary>
        /// The name of the cryptoKey associated with the CryptoKeyVersions.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("cryptoKey")]
        public Input<string>? CryptoKey { get; set; }

        /// <summary>
        /// The time this CryptoKeyVersion key material was generated
        /// </summary>
        [Input("generateTime")]
        public Input<string>? GenerateTime { get; set; }

        /// <summary>
        /// The resource name for this CryptoKeyVersion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.
        /// </summary>
        [Input("protectionLevel")]
        public Input<string>? ProtectionLevel { get; set; }

        /// <summary>
        /// The current state of the CryptoKeyVersion.
        /// Possible values are: `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, `IMPORT_FAILED`.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public CryptoKeyVersionState()
        {
        }
        public static new CryptoKeyVersionState Empty => new CryptoKeyVersionState();
    }
}
