// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.SecretManager
{
    /// <summary>
    /// A secret version resource.
    /// 
    /// &gt; **Warning:** All arguments including `payload.secret_data` will be stored in the raw
    /// state as plain-text.
    /// 
    /// ## Example Usage
    /// ### Secret Version Basic
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
    ///         SecretId = "secret-version",
    ///         Labels = 
    ///         {
    ///             { "label", "my-label" },
    ///         },
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
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SecretVersion can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:secretmanager/secretVersion:SecretVersion default projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:secretmanager/secretVersion:SecretVersion")]
    public partial class SecretVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the Secret was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time at which the Secret was destroyed. Only present if state is DESTROYED.
        /// </summary>
        [Output("destroyTime")]
        public Output<string> DestroyTime { get; private set; } = null!;

        /// <summary>
        /// The current state of the SecretVersion.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The resource name of the SecretVersion. Format:
        /// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Secret Manager secret resource
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// The secret data. Must be no larger than 64KiB.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        [Output("secretData")]
        public Output<string> SecretData { get; private set; } = null!;

        /// <summary>
        /// The version of the Secret.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a SecretVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretVersion(string name, SecretVersionArgs args, CustomResourceOptions? options = null)
            : base("gcp:secretmanager/secretVersion:SecretVersion", name, args ?? new SecretVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretVersion(string name, Input<string> id, SecretVersionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:secretmanager/secretVersion:SecretVersion", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "secretData",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretVersion Get(string name, Input<string> id, SecretVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretVersion(name, id, state, options);
        }
    }

    public sealed class SecretVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The current state of the SecretVersion.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Secret Manager secret resource
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        [Input("secretData", required: true)]
        private Input<string>? _secretData;

        /// <summary>
        /// The secret data. Must be no larger than 64KiB.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? SecretData
        {
            get => _secretData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public SecretVersionArgs()
        {
        }
        public static new SecretVersionArgs Empty => new SecretVersionArgs();
    }

    public sealed class SecretVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time at which the Secret was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The time at which the Secret was destroyed. Only present if state is DESTROYED.
        /// </summary>
        [Input("destroyTime")]
        public Input<string>? DestroyTime { get; set; }

        /// <summary>
        /// The current state of the SecretVersion.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The resource name of the SecretVersion. Format:
        /// `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Secret Manager secret resource
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        [Input("secretData")]
        private Input<string>? _secretData;

        /// <summary>
        /// The secret data. Must be no larger than 64KiB.
        /// **Note**: This property is sensitive and will not be displayed in the plan.
        /// </summary>
        public Input<string>? SecretData
        {
            get => _secretData;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretData = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The version of the Secret.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public SecretVersionState()
        {
        }
        public static new SecretVersionState Empty => new SecretVersionState();
    }
}
