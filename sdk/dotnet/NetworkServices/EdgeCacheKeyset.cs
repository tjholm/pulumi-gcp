// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices
{
    /// <summary>
    /// ## Example Usage
    /// ### Network Services Edge Cache Keyset Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkServices.EdgeCacheKeyset("default", new()
    ///     {
    ///         Description = "The default keyset",
    ///         PublicKeys = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.EdgeCacheKeysetPublicKeyArgs
    ///             {
    ///                 Id = "my-public-key",
    ///                 Value = "FHsTyFHNmvNpw4o7-rp-M1yqMyBF8vXSBRkZtkQ0RKY",
    ///             },
    ///             new Gcp.NetworkServices.Inputs.EdgeCacheKeysetPublicKeyArgs
    ///             {
    ///                 Id = "my-public-key-2",
    ///                 Value = "hzd03llxB1u5FOLKFkZ6_wCJqC7jtN0bg7xlBqS6WVM",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Network Services Edge Cache Keyset Dual Token
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var secret_basic = new Gcp.SecretManager.Secret("secret-basic", new()
    ///     {
    ///         SecretId = "secret-name",
    ///         Replication = new Gcp.SecretManager.Inputs.SecretReplicationArgs
    ///         {
    ///             Automatic = true,
    ///         },
    ///     });
    /// 
    ///     var secret_version_basic = new Gcp.SecretManager.SecretVersion("secret-version-basic", new()
    ///     {
    ///         Secret = secret_basic.Id,
    ///         SecretData = "secret-data",
    ///     });
    /// 
    ///     var @default = new Gcp.NetworkServices.EdgeCacheKeyset("default", new()
    ///     {
    ///         Description = "The default keyset",
    ///         PublicKeys = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.EdgeCacheKeysetPublicKeyArgs
    ///             {
    ///                 Id = "my-public-key",
    ///                 Managed = true,
    ///             },
    ///         },
    ///         ValidationSharedKeys = new[]
    ///         {
    ///             new Gcp.NetworkServices.Inputs.EdgeCacheKeysetValidationSharedKeyArgs
    ///             {
    ///                 SecretVersion = secret_version_basic.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EdgeCacheKeyset can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset")]
    public partial class EdgeCacheKeyset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the EdgeCache resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created.
        /// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
        /// and all following characters must be a dash, underscore, letter or digit.
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
        /// An ordered list of Ed25519 public keys to use for validating signed requests.
        /// You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
        /// You may specify no more than one Google-managed public key.
        /// If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
        /// Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
        /// Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
        /// Structure is documented below.
        /// </summary>
        [Output("publicKeys")]
        public Output<ImmutableArray<Outputs.EdgeCacheKeysetPublicKey>> PublicKeys { get; private set; } = null!;

        /// <summary>
        /// An ordered list of shared keys to use for validating signed requests.
        /// Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
        /// You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
        /// You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
        /// Structure is documented below.
        /// </summary>
        [Output("validationSharedKeys")]
        public Output<ImmutableArray<Outputs.EdgeCacheKeysetValidationSharedKey>> ValidationSharedKeys { get; private set; } = null!;


        /// <summary>
        /// Create a EdgeCacheKeyset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EdgeCacheKeyset(string name, EdgeCacheKeysetArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset", name, args ?? new EdgeCacheKeysetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EdgeCacheKeyset(string name, Input<string> id, EdgeCacheKeysetState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EdgeCacheKeyset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EdgeCacheKeyset Get(string name, Input<string> id, EdgeCacheKeysetState? state = null, CustomResourceOptions? options = null)
        {
            return new EdgeCacheKeyset(name, id, state, options);
        }
    }

    public sealed class EdgeCacheKeysetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the EdgeCache resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created.
        /// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
        /// and all following characters must be a dash, underscore, letter or digit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("publicKeys")]
        private InputList<Inputs.EdgeCacheKeysetPublicKeyArgs>? _publicKeys;

        /// <summary>
        /// An ordered list of Ed25519 public keys to use for validating signed requests.
        /// You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
        /// You may specify no more than one Google-managed public key.
        /// If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
        /// Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
        /// Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.EdgeCacheKeysetPublicKeyArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<Inputs.EdgeCacheKeysetPublicKeyArgs>());
            set => _publicKeys = value;
        }

        [Input("validationSharedKeys")]
        private InputList<Inputs.EdgeCacheKeysetValidationSharedKeyArgs>? _validationSharedKeys;

        /// <summary>
        /// An ordered list of shared keys to use for validating signed requests.
        /// Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
        /// You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
        /// You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.EdgeCacheKeysetValidationSharedKeyArgs> ValidationSharedKeys
        {
            get => _validationSharedKeys ?? (_validationSharedKeys = new InputList<Inputs.EdgeCacheKeysetValidationSharedKeyArgs>());
            set => _validationSharedKeys = value;
        }

        public EdgeCacheKeysetArgs()
        {
        }
        public static new EdgeCacheKeysetArgs Empty => new EdgeCacheKeysetArgs();
    }

    public sealed class EdgeCacheKeysetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the EdgeCache resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is created.
        /// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
        /// and all following characters must be a dash, underscore, letter or digit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("publicKeys")]
        private InputList<Inputs.EdgeCacheKeysetPublicKeyGetArgs>? _publicKeys;

        /// <summary>
        /// An ordered list of Ed25519 public keys to use for validating signed requests.
        /// You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
        /// You may specify no more than one Google-managed public key.
        /// If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
        /// Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
        /// Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.EdgeCacheKeysetPublicKeyGetArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<Inputs.EdgeCacheKeysetPublicKeyGetArgs>());
            set => _publicKeys = value;
        }

        [Input("validationSharedKeys")]
        private InputList<Inputs.EdgeCacheKeysetValidationSharedKeyGetArgs>? _validationSharedKeys;

        /// <summary>
        /// An ordered list of shared keys to use for validating signed requests.
        /// Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
        /// You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
        /// You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.EdgeCacheKeysetValidationSharedKeyGetArgs> ValidationSharedKeys
        {
            get => _validationSharedKeys ?? (_validationSharedKeys = new InputList<Inputs.EdgeCacheKeysetValidationSharedKeyGetArgs>());
            set => _validationSharedKeys = value;
        }

        public EdgeCacheKeysetState()
        {
        }
        public static new EdgeCacheKeysetState Empty => new EdgeCacheKeysetState();
    }
}
