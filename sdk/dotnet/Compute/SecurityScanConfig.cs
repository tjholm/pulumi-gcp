// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// A ScanConfig resource contains the configurations to launch a scan.
    /// 
    /// To get more information about ScanConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/security-scanner/docs/reference/rest/v1beta/projects.scanConfigs)
    /// * How-to Guides
    ///     * [Using Cloud Security Scanner](https://cloud.google.com/security-scanner/docs/scanning)
    /// 
    /// ## Example Usage
    /// ### Scan Config Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var scannerStaticIp = new Gcp.Compute.Address("scannerStaticIp", new()
    ///     {
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var scan_config = new Gcp.Compute.SecurityScanConfig("scan-config", new()
    ///     {
    ///         DisplayName = "scan-config",
    ///         StartingUrls = new[]
    ///         {
    ///             scannerStaticIp.IPAddress.Apply(address =&gt; $"http://{address}"),
    ///         },
    ///         TargetPlatforms = new[]
    ///         {
    ///             "COMPUTE",
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
    /// ScanConfig can be imported using any of these accepted formats* `projects/{{project}}/scanConfigs/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, ScanConfig can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default projects/{{project}}/scanConfigs/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/securityScanConfig:SecurityScanConfig default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/securityScanConfig:SecurityScanConfig")]
    public partial class SecurityScanConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The authentication configuration.
        /// If specified, service will use the authentication configuration during scanning.
        /// Structure is documented below.
        /// </summary>
        [Output("authentication")]
        public Output<Outputs.SecurityScanConfigAuthentication?> Authentication { get; private set; } = null!;

        /// <summary>
        /// The blacklist URL patterns as described in
        /// https://cloud.google.com/security-scanner/docs/excluded-urls
        /// </summary>
        [Output("blacklistPatterns")]
        public Output<ImmutableArray<string>> BlacklistPatterns { get; private set; } = null!;

        /// <summary>
        /// The user provider display name of the ScanConfig.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Controls export of scan configurations and results to Cloud Security Command Center.
        /// Default value is `ENABLED`.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Output("exportToSecurityCommandCenter")]
        public Output<string?> ExportToSecurityCommandCenter { get; private set; } = null!;

        /// <summary>
        /// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
        /// Defaults to 15.
        /// </summary>
        [Output("maxQps")]
        public Output<int?> MaxQps { get; private set; } = null!;

        /// <summary>
        /// A server defined name for this index. Format:
        /// `projects/{{project}}/scanConfigs/{{server_generated_id}}`
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
        /// The schedule of the ScanConfig
        /// Structure is documented below.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.SecurityScanConfigSchedule?> Schedule { get; private set; } = null!;

        /// <summary>
        /// The starting URLs from which the scanner finds site pages.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("startingUrls")]
        public Output<ImmutableArray<string>> StartingUrls { get; private set; } = null!;

        /// <summary>
        /// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        /// Each value may be one of: `APP_ENGINE`, `COMPUTE`.
        /// </summary>
        [Output("targetPlatforms")]
        public Output<ImmutableArray<string>> TargetPlatforms { get; private set; } = null!;

        /// <summary>
        /// Type of the user agents used for scanning
        /// Default value is `CHROME_LINUX`.
        /// Possible values are: `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, `SAFARI_IPHONE`.
        /// </summary>
        [Output("userAgent")]
        public Output<string?> UserAgent { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityScanConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityScanConfig(string name, SecurityScanConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/securityScanConfig:SecurityScanConfig", name, args ?? new SecurityScanConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityScanConfig(string name, Input<string> id, SecurityScanConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/securityScanConfig:SecurityScanConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityScanConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityScanConfig Get(string name, Input<string> id, SecurityScanConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityScanConfig(name, id, state, options);
        }
    }

    public sealed class SecurityScanConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication configuration.
        /// If specified, service will use the authentication configuration during scanning.
        /// Structure is documented below.
        /// </summary>
        [Input("authentication")]
        public Input<Inputs.SecurityScanConfigAuthenticationArgs>? Authentication { get; set; }

        [Input("blacklistPatterns")]
        private InputList<string>? _blacklistPatterns;

        /// <summary>
        /// The blacklist URL patterns as described in
        /// https://cloud.google.com/security-scanner/docs/excluded-urls
        /// </summary>
        public InputList<string> BlacklistPatterns
        {
            get => _blacklistPatterns ?? (_blacklistPatterns = new InputList<string>());
            set => _blacklistPatterns = value;
        }

        /// <summary>
        /// The user provider display name of the ScanConfig.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Controls export of scan configurations and results to Cloud Security Command Center.
        /// Default value is `ENABLED`.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("exportToSecurityCommandCenter")]
        public Input<string>? ExportToSecurityCommandCenter { get; set; }

        /// <summary>
        /// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
        /// Defaults to 15.
        /// </summary>
        [Input("maxQps")]
        public Input<int>? MaxQps { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The schedule of the ScanConfig
        /// Structure is documented below.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.SecurityScanConfigScheduleArgs>? Schedule { get; set; }

        [Input("startingUrls", required: true)]
        private InputList<string>? _startingUrls;

        /// <summary>
        /// The starting URLs from which the scanner finds site pages.
        /// 
        /// 
        /// - - -
        /// </summary>
        public InputList<string> StartingUrls
        {
            get => _startingUrls ?? (_startingUrls = new InputList<string>());
            set => _startingUrls = value;
        }

        [Input("targetPlatforms")]
        private InputList<string>? _targetPlatforms;

        /// <summary>
        /// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        /// Each value may be one of: `APP_ENGINE`, `COMPUTE`.
        /// </summary>
        public InputList<string> TargetPlatforms
        {
            get => _targetPlatforms ?? (_targetPlatforms = new InputList<string>());
            set => _targetPlatforms = value;
        }

        /// <summary>
        /// Type of the user agents used for scanning
        /// Default value is `CHROME_LINUX`.
        /// Possible values are: `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, `SAFARI_IPHONE`.
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        public SecurityScanConfigArgs()
        {
        }
        public static new SecurityScanConfigArgs Empty => new SecurityScanConfigArgs();
    }

    public sealed class SecurityScanConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication configuration.
        /// If specified, service will use the authentication configuration during scanning.
        /// Structure is documented below.
        /// </summary>
        [Input("authentication")]
        public Input<Inputs.SecurityScanConfigAuthenticationGetArgs>? Authentication { get; set; }

        [Input("blacklistPatterns")]
        private InputList<string>? _blacklistPatterns;

        /// <summary>
        /// The blacklist URL patterns as described in
        /// https://cloud.google.com/security-scanner/docs/excluded-urls
        /// </summary>
        public InputList<string> BlacklistPatterns
        {
            get => _blacklistPatterns ?? (_blacklistPatterns = new InputList<string>());
            set => _blacklistPatterns = value;
        }

        /// <summary>
        /// The user provider display name of the ScanConfig.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Controls export of scan configurations and results to Cloud Security Command Center.
        /// Default value is `ENABLED`.
        /// Possible values are: `ENABLED`, `DISABLED`.
        /// </summary>
        [Input("exportToSecurityCommandCenter")]
        public Input<string>? ExportToSecurityCommandCenter { get; set; }

        /// <summary>
        /// The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively.
        /// Defaults to 15.
        /// </summary>
        [Input("maxQps")]
        public Input<int>? MaxQps { get; set; }

        /// <summary>
        /// A server defined name for this index. Format:
        /// `projects/{{project}}/scanConfigs/{{server_generated_id}}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The schedule of the ScanConfig
        /// Structure is documented below.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.SecurityScanConfigScheduleGetArgs>? Schedule { get; set; }

        [Input("startingUrls")]
        private InputList<string>? _startingUrls;

        /// <summary>
        /// The starting URLs from which the scanner finds site pages.
        /// 
        /// 
        /// - - -
        /// </summary>
        public InputList<string> StartingUrls
        {
            get => _startingUrls ?? (_startingUrls = new InputList<string>());
            set => _startingUrls = value;
        }

        [Input("targetPlatforms")]
        private InputList<string>? _targetPlatforms;

        /// <summary>
        /// Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.
        /// Each value may be one of: `APP_ENGINE`, `COMPUTE`.
        /// </summary>
        public InputList<string> TargetPlatforms
        {
            get => _targetPlatforms ?? (_targetPlatforms = new InputList<string>());
            set => _targetPlatforms = value;
        }

        /// <summary>
        /// Type of the user agents used for scanning
        /// Default value is `CHROME_LINUX`.
        /// Possible values are: `USER_AGENT_UNSPECIFIED`, `CHROME_LINUX`, `CHROME_ANDROID`, `SAFARI_IPHONE`.
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        public SecurityScanConfigState()
        {
        }
        public static new SecurityScanConfigState Empty => new SecurityScanConfigState();
    }
}
