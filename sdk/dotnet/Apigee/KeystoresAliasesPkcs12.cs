// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee
{
    /// <summary>
    /// An alias from a pkcs12 file.
    /// 
    /// To get more information about KeystoresAliasesPkcs12, see:
    /// 
    /// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases)
    /// * How-to Guides
    ///     * [Keystores Aliases](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases)
    /// 
    /// ## Import
    /// 
    /// KeystoresAliasesPkcs12 can be imported using any of these accepted formats* `organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}` * `{{org_id}}/{{environment}}/{{keystore}}/{{alias}}` When using the `pulumi import` command, KeystoresAliasesPkcs12 can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/keystoresAliasesPkcs12:KeystoresAliasesPkcs12 default organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/keystoresAliasesPkcs12:KeystoresAliasesPkcs12 default {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigee/keystoresAliasesPkcs12:KeystoresAliasesPkcs12")]
    public partial class KeystoresAliasesPkcs12 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Chain of certificates under this alias.
        /// Structure is documented below.
        /// </summary>
        [Output("certsInfos")]
        public Output<ImmutableArray<Outputs.KeystoresAliasesPkcs12CertsInfo>> CertsInfos { get; private set; } = null!;

        /// <summary>
        /// Environment associated with the alias
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// PKCS12 file content
        /// 
        /// - - -
        /// </summary>
        [Output("file")]
        public Output<string> File { get; private set; } = null!;

        /// <summary>
        /// Hash of the pkcs file
        /// </summary>
        [Output("filehash")]
        public Output<string> Filehash { get; private set; } = null!;

        /// <summary>
        /// Keystore Name
        /// </summary>
        [Output("keystore")]
        public Output<string> Keystore { get; private set; } = null!;

        /// <summary>
        /// Organization ID associated with the alias, without organization/ prefix
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Password for the PKCS12 file if it's encrypted
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Optional.Type of Alias
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a KeystoresAliasesPkcs12 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeystoresAliasesPkcs12(string name, KeystoresAliasesPkcs12Args args, CustomResourceOptions? options = null)
            : base("gcp:apigee/keystoresAliasesPkcs12:KeystoresAliasesPkcs12", name, args ?? new KeystoresAliasesPkcs12Args(), MakeResourceOptions(options, ""))
        {
        }

        private KeystoresAliasesPkcs12(string name, Input<string> id, KeystoresAliasesPkcs12State? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/keystoresAliasesPkcs12:KeystoresAliasesPkcs12", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KeystoresAliasesPkcs12 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeystoresAliasesPkcs12 Get(string name, Input<string> id, KeystoresAliasesPkcs12State? state = null, CustomResourceOptions? options = null)
        {
            return new KeystoresAliasesPkcs12(name, id, state, options);
        }
    }

    public sealed class KeystoresAliasesPkcs12Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Environment associated with the alias
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// PKCS12 file content
        /// 
        /// - - -
        /// </summary>
        [Input("file", required: true)]
        public Input<string> File { get; set; } = null!;

        /// <summary>
        /// Hash of the pkcs file
        /// </summary>
        [Input("filehash", required: true)]
        public Input<string> Filehash { get; set; } = null!;

        /// <summary>
        /// Keystore Name
        /// </summary>
        [Input("keystore", required: true)]
        public Input<string> Keystore { get; set; } = null!;

        /// <summary>
        /// Organization ID associated with the alias, without organization/ prefix
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Password for the PKCS12 file if it's encrypted
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public KeystoresAliasesPkcs12Args()
        {
        }
        public static new KeystoresAliasesPkcs12Args Empty => new KeystoresAliasesPkcs12Args();
    }

    public sealed class KeystoresAliasesPkcs12State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        [Input("certsInfos")]
        private InputList<Inputs.KeystoresAliasesPkcs12CertsInfoGetArgs>? _certsInfos;

        /// <summary>
        /// Chain of certificates under this alias.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.KeystoresAliasesPkcs12CertsInfoGetArgs> CertsInfos
        {
            get => _certsInfos ?? (_certsInfos = new InputList<Inputs.KeystoresAliasesPkcs12CertsInfoGetArgs>());
            set => _certsInfos = value;
        }

        /// <summary>
        /// Environment associated with the alias
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        /// <summary>
        /// PKCS12 file content
        /// 
        /// - - -
        /// </summary>
        [Input("file")]
        public Input<string>? File { get; set; }

        /// <summary>
        /// Hash of the pkcs file
        /// </summary>
        [Input("filehash")]
        public Input<string>? Filehash { get; set; }

        /// <summary>
        /// Keystore Name
        /// </summary>
        [Input("keystore")]
        public Input<string>? Keystore { get; set; }

        /// <summary>
        /// Organization ID associated with the alias, without organization/ prefix
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Password for the PKCS12 file if it's encrypted
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Optional.Type of Alias
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public KeystoresAliasesPkcs12State()
        {
        }
        public static new KeystoresAliasesPkcs12State Empty => new KeystoresAliasesPkcs12State();
    }
}
