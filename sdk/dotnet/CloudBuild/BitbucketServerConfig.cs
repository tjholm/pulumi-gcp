// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild
{
    /// <summary>
    /// BitbucketServerConfig represents the configuration for a Bitbucket Server.
    /// 
    /// To get more information about BitbucketServerConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/build/docs/api/reference/rest/v1/projects.locations.bitbucketServerConfigs)
    /// * How-to Guides
    ///     * [Connect to a Bitbucket Server host](https://cloud.google.com/build/docs/automating-builds/bitbucket/connect-host-bitbucket-server)
    /// 
    /// ## Example Usage
    /// ### Cloudbuild Bitbucket Server Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bbs_config = new Gcp.CloudBuild.BitbucketServerConfig("bbs-config", new()
    ///     {
    ///         ApiKey = "&lt;api-key&gt;",
    ///         ConfigId = "bbs-config",
    ///         HostUri = "https://bbs.com",
    ///         Location = "us-central1",
    ///         Secrets = new Gcp.CloudBuild.Inputs.BitbucketServerConfigSecretsArgs
    ///         {
    ///             AdminAccessTokenVersionName = "projects/myProject/secrets/mybbspat/versions/1",
    ///             ReadAccessTokenVersionName = "projects/myProject/secrets/mybbspat/versions/1",
    ///             WebhookSecretVersionName = "projects/myProject/secrets/mybbspat/versions/1",
    ///         },
    ///         Username = "test",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Cloudbuild Bitbucket Server Config Repositories
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bbs_config_with_repos = new Gcp.CloudBuild.BitbucketServerConfig("bbs-config-with-repos", new()
    ///     {
    ///         ApiKey = "&lt;api-key&gt;",
    ///         ConfigId = "bbs-config",
    ///         ConnectedRepositories = new[]
    ///         {
    ///             new Gcp.CloudBuild.Inputs.BitbucketServerConfigConnectedRepositoryArgs
    ///             {
    ///                 ProjectKey = "DEV",
    ///                 RepoSlug = "repo1",
    ///             },
    ///             new Gcp.CloudBuild.Inputs.BitbucketServerConfigConnectedRepositoryArgs
    ///             {
    ///                 ProjectKey = "PROD",
    ///                 RepoSlug = "repo1",
    ///             },
    ///         },
    ///         HostUri = "https://bbs.com",
    ///         Location = "us-central1",
    ///         Secrets = new Gcp.CloudBuild.Inputs.BitbucketServerConfigSecretsArgs
    ///         {
    ///             AdminAccessTokenVersionName = "projects/myProject/secrets/mybbspat/versions/1",
    ///             ReadAccessTokenVersionName = "projects/myProject/secrets/mybbspat/versions/1",
    ///             WebhookSecretVersionName = "projects/myProject/secrets/mybbspat/versions/1",
    ///         },
    ///         Username = "test",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BitbucketServerConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudbuild/bitbucketServerConfig:BitbucketServerConfig default projects/{{project}}/locations/{{location}}/bitbucketServerConfigs/{{config_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudbuild/bitbucketServerConfig:BitbucketServerConfig default {{project}}/{{location}}/{{config_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:cloudbuild/bitbucketServerConfig:BitbucketServerConfig default {{location}}/{{config_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:cloudbuild/bitbucketServerConfig:BitbucketServerConfig")]
    public partial class BitbucketServerConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed.
        /// Changing this field will result in deleting/ recreating the resource.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.
        /// </summary>
        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        /// <summary>
        /// Connected Bitbucket Server repositories for this config.
        /// Structure is documented below.
        /// </summary>
        [Output("connectedRepositories")]
        public Output<ImmutableArray<Outputs.BitbucketServerConfigConnectedRepository>> ConnectedRepositories { get; private set; } = null!;

        /// <summary>
        /// Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.
        /// If you need to change it, please create another BitbucketServerConfig.
        /// </summary>
        [Output("hostUri")]
        public Output<string> HostUri { get; private set; } = null!;

        /// <summary>
        /// The location of this bitbucket server config.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name for the config.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection.
        /// This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,
        /// no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format
        /// projects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.
        /// </summary>
        [Output("peeredNetwork")]
        public Output<string?> PeeredNetwork { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Secret Manager secrets needed by the config.
        /// Structure is documented below.
        /// </summary>
        [Output("secrets")]
        public Output<Outputs.BitbucketServerConfigSecrets> Secrets { get; private set; } = null!;

        /// <summary>
        /// SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
        /// </summary>
        [Output("sslCa")]
        public Output<string?> SslCa { get; private set; } = null!;

        /// <summary>
        /// Username of the account Cloud Build will use on Bitbucket Server.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// Output only. UUID included in webhook requests. The UUID is used to look up the corresponding config.
        /// </summary>
        [Output("webhookKey")]
        public Output<string> WebhookKey { get; private set; } = null!;


        /// <summary>
        /// Create a BitbucketServerConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BitbucketServerConfig(string name, BitbucketServerConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:cloudbuild/bitbucketServerConfig:BitbucketServerConfig", name, args ?? new BitbucketServerConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BitbucketServerConfig(string name, Input<string> id, BitbucketServerConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudbuild/bitbucketServerConfig:BitbucketServerConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BitbucketServerConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BitbucketServerConfig Get(string name, Input<string> id, BitbucketServerConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new BitbucketServerConfig(name, id, state, options);
        }
    }

    public sealed class BitbucketServerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed.
        /// Changing this field will result in deleting/ recreating the resource.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.
        /// </summary>
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        [Input("connectedRepositories")]
        private InputList<Inputs.BitbucketServerConfigConnectedRepositoryArgs>? _connectedRepositories;

        /// <summary>
        /// Connected Bitbucket Server repositories for this config.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BitbucketServerConfigConnectedRepositoryArgs> ConnectedRepositories
        {
            get => _connectedRepositories ?? (_connectedRepositories = new InputList<Inputs.BitbucketServerConfigConnectedRepositoryArgs>());
            set => _connectedRepositories = value;
        }

        /// <summary>
        /// Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.
        /// If you need to change it, please create another BitbucketServerConfig.
        /// </summary>
        [Input("hostUri", required: true)]
        public Input<string> HostUri { get; set; } = null!;

        /// <summary>
        /// The location of this bitbucket server config.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection.
        /// This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,
        /// no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format
        /// projects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.
        /// </summary>
        [Input("peeredNetwork")]
        public Input<string>? PeeredNetwork { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Secret Manager secrets needed by the config.
        /// Structure is documented below.
        /// </summary>
        [Input("secrets", required: true)]
        public Input<Inputs.BitbucketServerConfigSecretsArgs> Secrets { get; set; } = null!;

        /// <summary>
        /// SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
        /// </summary>
        [Input("sslCa")]
        public Input<string>? SslCa { get; set; }

        /// <summary>
        /// Username of the account Cloud Build will use on Bitbucket Server.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public BitbucketServerConfigArgs()
        {
        }
        public static new BitbucketServerConfigArgs Empty => new BitbucketServerConfigArgs();
    }

    public sealed class BitbucketServerConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. API Key that will be attached to webhook. Once this field has been set, it cannot be changed.
        /// Changing this field will result in deleting/ recreating the resource.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// The ID to use for the BitbucketServerConfig, which will become the final component of the BitbucketServerConfig's resource name.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        [Input("connectedRepositories")]
        private InputList<Inputs.BitbucketServerConfigConnectedRepositoryGetArgs>? _connectedRepositories;

        /// <summary>
        /// Connected Bitbucket Server repositories for this config.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.BitbucketServerConfigConnectedRepositoryGetArgs> ConnectedRepositories
        {
            get => _connectedRepositories ?? (_connectedRepositories = new InputList<Inputs.BitbucketServerConfigConnectedRepositoryGetArgs>());
            set => _connectedRepositories = value;
        }

        /// <summary>
        /// Immutable. The URI of the Bitbucket Server host. Once this field has been set, it cannot be changed.
        /// If you need to change it, please create another BitbucketServerConfig.
        /// </summary>
        [Input("hostUri")]
        public Input<string>? HostUri { get; set; }

        /// <summary>
        /// The location of this bitbucket server config.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name for the config.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The network to be used when reaching out to the Bitbucket Server instance. The VPC network must be enabled for private service connection.
        /// This should be set if the Bitbucket Server instance is hosted on-premises and not reachable by public internet. If this field is left empty,
        /// no network peering will occur and calls to the Bitbucket Server instance will be made over the public internet. Must be in the format
        /// projects/{project}/global/networks/{network}, where {project} is a project number or id and {network} is the name of a VPC network in the project.
        /// </summary>
        [Input("peeredNetwork")]
        public Input<string>? PeeredNetwork { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Secret Manager secrets needed by the config.
        /// Structure is documented below.
        /// </summary>
        [Input("secrets")]
        public Input<Inputs.BitbucketServerConfigSecretsGetArgs>? Secrets { get; set; }

        /// <summary>
        /// SSL certificate to use for requests to Bitbucket Server. The format should be PEM format but the extension can be one of .pem, .cer, or .crt.
        /// </summary>
        [Input("sslCa")]
        public Input<string>? SslCa { get; set; }

        /// <summary>
        /// Username of the account Cloud Build will use on Bitbucket Server.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Output only. UUID included in webhook requests. The UUID is used to look up the corresponding config.
        /// </summary>
        [Input("webhookKey")]
        public Input<string>? WebhookKey { get; set; }

        public BitbucketServerConfigState()
        {
        }
        public static new BitbucketServerConfigState Empty => new BitbucketServerConfigState();
    }
}
