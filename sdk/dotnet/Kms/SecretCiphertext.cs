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
    /// Encrypts secret data with Google Cloud KMS and provides access to the ciphertext.
    /// 
    /// &gt; **NOTE:** Using this resource will allow you to conceal secret data within your
    /// resource definitions, but it does not take care of protecting that data in the
    /// logging output, plan output, or state output.  Please take care to secure your secret
    /// data outside of resource definitions.
    /// 
    /// To get more information about SecretCiphertext, see:
    /// 
    /// * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys/encrypt)
    /// * How-to Guides
    ///     * [Encrypting and decrypting data with a symmetric key](https://cloud.google.com/kms/docs/encrypt-decrypt)
    /// 
    /// &gt; **Warning:** All arguments including `plaintext` and `additional_authenticated_data` will be stored in the raw state as plain-text.
    /// 
    /// ## Example Usage
    /// ### Kms Secret Ciphertext Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
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
    ///         RotationPeriod = "100000s",
    ///     });
    /// 
    ///     var myPassword = new Gcp.Kms.SecretCiphertext("myPassword", new()
    ///     {
    ///         CryptoKey = cryptokey.Id,
    ///         Plaintext = "my-secret-password",
    ///     });
    /// 
    ///     var instance = new Gcp.Compute.Instance("instance", new()
    ///     {
    ///         MachineType = "e2-medium",
    ///         Zone = "us-central1-a",
    ///         BootDisk = new Gcp.Compute.Inputs.InstanceBootDiskArgs
    ///         {
    ///             InitializeParams = new Gcp.Compute.Inputs.InstanceBootDiskInitializeParamsArgs
    ///             {
    ///                 Image = "debian-cloud/debian-11",
    ///             },
    ///         },
    ///         NetworkInterfaces = new[]
    ///         {
    ///             new Gcp.Compute.Inputs.InstanceNetworkInterfaceArgs
    ///             {
    ///                 Network = "default",
    ///                 AccessConfigs = new[]
    ///                 {
    ///                     null,
    ///                 },
    ///             },
    ///         },
    ///         Metadata = 
    ///         {
    ///             { "password", myPassword.Ciphertext },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:kms/secretCiphertext:SecretCiphertext")]
    public partial class SecretCiphertext : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The additional authenticated data used for integrity checks during encryption and decryption.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("additionalAuthenticatedData")]
        public Output<string?> AdditionalAuthenticatedData { get; private set; } = null!;

        /// <summary>
        /// Contains the result of encrypting the provided plaintext, encoded in base64.
        /// </summary>
        [Output("ciphertext")]
        public Output<string> Ciphertext { get; private set; } = null!;

        /// <summary>
        /// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
        /// </summary>
        [Output("cryptoKey")]
        public Output<string> CryptoKey { get; private set; } = null!;

        /// <summary>
        /// The plaintext to be encrypted.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("plaintext")]
        public Output<string> Plaintext { get; private set; } = null!;


        /// <summary>
        /// Create a SecretCiphertext resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretCiphertext(string name, SecretCiphertextArgs args, CustomResourceOptions? options = null)
            : base("gcp:kms/secretCiphertext:SecretCiphertext", name, args ?? new SecretCiphertextArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretCiphertext(string name, Input<string> id, SecretCiphertextState? state = null, CustomResourceOptions? options = null)
            : base("gcp:kms/secretCiphertext:SecretCiphertext", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "additionalAuthenticatedData",
                    "plaintext",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretCiphertext resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretCiphertext Get(string name, Input<string> id, SecretCiphertextState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretCiphertext(name, id, state, options);
        }
    }

    public sealed class SecretCiphertextArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticatedData")]
        private Input<string>? _additionalAuthenticatedData;

        /// <summary>
        /// The additional authenticated data used for integrity checks during encryption and decryption.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? AdditionalAuthenticatedData
        {
            get => _additionalAuthenticatedData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _additionalAuthenticatedData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
        /// </summary>
        [Input("cryptoKey", required: true)]
        public Input<string> CryptoKey { get; set; } = null!;

        [Input("plaintext", required: true)]
        private Input<string>? _plaintext;

        /// <summary>
        /// The plaintext to be encrypted.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? Plaintext
        {
            get => _plaintext;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _plaintext = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretCiphertextArgs()
        {
        }
        public static new SecretCiphertextArgs Empty => new SecretCiphertextArgs();
    }

    public sealed class SecretCiphertextState : global::Pulumi.ResourceArgs
    {
        [Input("additionalAuthenticatedData")]
        private Input<string>? _additionalAuthenticatedData;

        /// <summary>
        /// The additional authenticated data used for integrity checks during encryption and decryption.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? AdditionalAuthenticatedData
        {
            get => _additionalAuthenticatedData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _additionalAuthenticatedData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Contains the result of encrypting the provided plaintext, encoded in base64.
        /// </summary>
        [Input("ciphertext")]
        public Input<string>? Ciphertext { get; set; }

        /// <summary>
        /// The full name of the CryptoKey that will be used to encrypt the provided plaintext.
        /// Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}'`
        /// </summary>
        [Input("cryptoKey")]
        public Input<string>? CryptoKey { get; set; }

        [Input("plaintext")]
        private Input<string>? _plaintext;

        /// <summary>
        /// The plaintext to be encrypted.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? Plaintext
        {
            get => _plaintext;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _plaintext = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretCiphertextState()
        {
        }
        public static new SecretCiphertextState Empty => new SecretCiphertextState();
    }
}
