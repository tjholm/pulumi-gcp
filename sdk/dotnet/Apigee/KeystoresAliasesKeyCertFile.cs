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
    /// An alias from a key/certificate pair.
    /// 
    /// To get more information about KeystoresAliasesKeyCertFile, see:
    /// 
    /// * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases)
    /// * How-to Guides
    ///     * [Keystores Aliases](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.environments.keystores.aliases)
    /// 
    /// ## Import
    /// 
    /// KeystoresAliasesKeyCertFile can be imported using any of these accepted formats* `organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}` * `{{org_id}}/{{environment}}/{{keystore}}/{{alias}}` When using the `pulumi import` command, KeystoresAliasesKeyCertFile can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile default organizations/{{org_id}}/environments/{{environment}}/keystores/{{keystore}}/aliases/{{alias}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile default {{org_id}}/{{environment}}/{{keystore}}/{{alias}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile")]
    public partial class KeystoresAliasesKeyCertFile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Cert content
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("cert")]
        public Output<string> Cert { get; private set; } = null!;

        /// <summary>
        /// Chain of certificates under this alias.
        /// Structure is documented below.
        /// </summary>
        [Output("certsInfo")]
        public Output<Outputs.KeystoresAliasesKeyCertFileCertsInfo> CertsInfo { get; private set; } = null!;

        /// <summary>
        /// Environment associated with the alias
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// Private Key content, omit if uploading to truststore
        /// </summary>
        [Output("key")]
        public Output<string?> Key { get; private set; } = null!;

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
        /// Password for the Private Key if it's encrypted
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Optional.Type of Alias
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a KeystoresAliasesKeyCertFile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeystoresAliasesKeyCertFile(string name, KeystoresAliasesKeyCertFileArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile", name, args ?? new KeystoresAliasesKeyCertFileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeystoresAliasesKeyCertFile(string name, Input<string> id, KeystoresAliasesKeyCertFileState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "key",
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KeystoresAliasesKeyCertFile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeystoresAliasesKeyCertFile Get(string name, Input<string> id, KeystoresAliasesKeyCertFileState? state = null, CustomResourceOptions? options = null)
        {
            return new KeystoresAliasesKeyCertFile(name, id, state, options);
        }
    }

    public sealed class KeystoresAliasesKeyCertFileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// Cert content
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("cert", required: true)]
        public Input<string> Cert { get; set; } = null!;

        /// <summary>
        /// Chain of certificates under this alias.
        /// Structure is documented below.
        /// </summary>
        [Input("certsInfo")]
        public Input<Inputs.KeystoresAliasesKeyCertFileCertsInfoArgs>? CertsInfo { get; set; }

        /// <summary>
        /// Environment associated with the alias
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// Private Key content, omit if uploading to truststore
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the Private Key if it's encrypted
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public KeystoresAliasesKeyCertFileArgs()
        {
        }
        public static new KeystoresAliasesKeyCertFileArgs Empty => new KeystoresAliasesKeyCertFileArgs();
    }

    public sealed class KeystoresAliasesKeyCertFileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alias Name
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Cert content
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("cert")]
        public Input<string>? Cert { get; set; }

        /// <summary>
        /// Chain of certificates under this alias.
        /// Structure is documented below.
        /// </summary>
        [Input("certsInfo")]
        public Input<Inputs.KeystoresAliasesKeyCertFileCertsInfoGetArgs>? CertsInfo { get; set; }

        /// <summary>
        /// Environment associated with the alias
        /// </summary>
        [Input("environment")]
        public Input<string>? Environment { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// Private Key content, omit if uploading to truststore
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password for the Private Key if it's encrypted
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Optional.Type of Alias
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public KeystoresAliasesKeyCertFileState()
        {
        }
        public static new KeystoresAliasesKeyCertFileState Empty => new KeystoresAliasesKeyCertFileState();
    }
}
