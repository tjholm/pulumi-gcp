// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsLogin
{
    /// <summary>
    /// The SSH public key information associated with a Google account.
    /// 
    /// To get more information about SSHPublicKey, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/oslogin/rest/v1/users.sshPublicKeys)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/compute/docs/oslogin)
    /// 
    /// ## Example Usage
    /// ### Os Login Ssh Key Basic
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
    ///     var me = Gcp.Organizations.GetClientOpenIdUserInfo.Invoke();
    /// 
    ///     var cache = new Gcp.OsLogin.SshPublicKey("cache", new()
    ///     {
    ///         User = me.Apply(getClientOpenIdUserInfoResult =&gt; getClientOpenIdUserInfoResult.Email),
    ///         Key = File.ReadAllText("path/to/id_rsa.pub"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SSHPublicKey can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default users/{{user}}/sshPublicKeys/{{fingerprint}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:oslogin/sshPublicKey:SshPublicKey default {{user}}/{{fingerprint}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:oslogin/sshPublicKey:SshPublicKey")]
    public partial class SshPublicKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An expiration time in microseconds since epoch.
        /// </summary>
        [Output("expirationTimeUsec")]
        public Output<string?> ExpirationTimeUsec { get; private set; } = null!;

        /// <summary>
        /// The SHA-256 fingerprint of the SSH public key.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// Public key text in SSH format, defined by RFC4253 section 6.6.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// The project ID of the Google Cloud Platform project.
        /// </summary>
        [Output("project")]
        public Output<string?> Project { get; private set; } = null!;

        /// <summary>
        /// The user email.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a SshPublicKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SshPublicKey(string name, SshPublicKeyArgs args, CustomResourceOptions? options = null)
            : base("gcp:oslogin/sshPublicKey:SshPublicKey", name, args ?? new SshPublicKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SshPublicKey(string name, Input<string> id, SshPublicKeyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:oslogin/sshPublicKey:SshPublicKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SshPublicKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SshPublicKey Get(string name, Input<string> id, SshPublicKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new SshPublicKey(name, id, state, options);
        }
    }

    public sealed class SshPublicKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An expiration time in microseconds since epoch.
        /// </summary>
        [Input("expirationTimeUsec")]
        public Input<string>? ExpirationTimeUsec { get; set; }

        /// <summary>
        /// Public key text in SSH format, defined by RFC4253 section 6.6.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The project ID of the Google Cloud Platform project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The user email.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public SshPublicKeyArgs()
        {
        }
        public static new SshPublicKeyArgs Empty => new SshPublicKeyArgs();
    }

    public sealed class SshPublicKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An expiration time in microseconds since epoch.
        /// </summary>
        [Input("expirationTimeUsec")]
        public Input<string>? ExpirationTimeUsec { get; set; }

        /// <summary>
        /// The SHA-256 fingerprint of the SSH public key.
        /// </summary>
        [Input("fingerprint")]
        public Input<string>? Fingerprint { get; set; }

        /// <summary>
        /// Public key text in SSH format, defined by RFC4253 section 6.6.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The project ID of the Google Cloud Platform project.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The user email.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SshPublicKeyState()
        {
        }
        public static new SshPublicKeyState Empty => new SshPublicKeyState();
    }
}
