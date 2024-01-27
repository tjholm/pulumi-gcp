// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    /// <summary>
    /// This message configures which resources and services to monitor for availability.
    /// 
    /// To get more information about UptimeCheckConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.uptimeCheckConfigs)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/monitoring/uptime-checks/)
    /// 
    /// ## Example Usage
    /// ### Uptime Check Config Http
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var http = new Gcp.Monitoring.UptimeCheckConfig("http", new()
    ///     {
    ///         CheckerType = "STATIC_IP_CHECKERS",
    ///         ContentMatchers = new[]
    ///         {
    ///             new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherArgs
    ///             {
    ///                 Content = "\"example\"",
    ///                 JsonPathMatcher = new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherJsonPathMatcherArgs
    ///                 {
    ///                     JsonMatcher = "EXACT_MATCH",
    ///                     JsonPath = "$.path",
    ///                 },
    ///                 Matcher = "MATCHES_JSON_PATH",
    ///             },
    ///         },
    ///         DisplayName = "http-uptime-check",
    ///         HttpCheck = new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckArgs
    ///         {
    ///             Body = "Zm9vJTI1M0RiYXI=",
    ///             ContentType = "USER_PROVIDED",
    ///             CustomContentType = "application/json",
    ///             Path = "some-path",
    ///             PingConfig = new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckPingConfigArgs
    ///             {
    ///                 PingsCount = 1,
    ///             },
    ///             Port = 8010,
    ///             RequestMethod = "POST",
    ///         },
    ///         MonitoredResource = new Gcp.Monitoring.Inputs.UptimeCheckConfigMonitoredResourceArgs
    ///         {
    ///             Labels = 
    ///             {
    ///                 { "host", "192.168.1.1" },
    ///                 { "projectId", "my-project-name" },
    ///             },
    ///             Type = "uptime_url",
    ///         },
    ///         Timeout = "60s",
    ///         UserLabels = 
    ///         {
    ///             { "example-key", "example-value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Uptime Check Config Status Code
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var statusCode = new Gcp.Monitoring.UptimeCheckConfig("statusCode", new()
    ///     {
    ///         CheckerType = "STATIC_IP_CHECKERS",
    ///         ContentMatchers = new[]
    ///         {
    ///             new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherArgs
    ///             {
    ///                 Content = "\"example\"",
    ///                 JsonPathMatcher = new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherJsonPathMatcherArgs
    ///                 {
    ///                     JsonMatcher = "EXACT_MATCH",
    ///                     JsonPath = "$.path",
    ///                 },
    ///                 Matcher = "MATCHES_JSON_PATH",
    ///             },
    ///         },
    ///         DisplayName = "http-uptime-check",
    ///         HttpCheck = new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckArgs
    ///         {
    ///             AcceptedResponseStatusCodes = new[]
    ///             {
    ///                 new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckAcceptedResponseStatusCodeArgs
    ///                 {
    ///                     StatusClass = "STATUS_CLASS_2XX",
    ///                 },
    ///                 new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckAcceptedResponseStatusCodeArgs
    ///                 {
    ///                     StatusValue = 301,
    ///                 },
    ///                 new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckAcceptedResponseStatusCodeArgs
    ///                 {
    ///                     StatusValue = 302,
    ///                 },
    ///             },
    ///             Body = "Zm9vJTI1M0RiYXI=",
    ///             ContentType = "URL_ENCODED",
    ///             Path = "some-path",
    ///             Port = 8010,
    ///             RequestMethod = "POST",
    ///         },
    ///         MonitoredResource = new Gcp.Monitoring.Inputs.UptimeCheckConfigMonitoredResourceArgs
    ///         {
    ///             Labels = 
    ///             {
    ///                 { "host", "192.168.1.1" },
    ///                 { "projectId", "my-project-name" },
    ///             },
    ///             Type = "uptime_url",
    ///         },
    ///         Timeout = "60s",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Uptime Check Config Https
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var https = new Gcp.Monitoring.UptimeCheckConfig("https", new()
    ///     {
    ///         ContentMatchers = new[]
    ///         {
    ///             new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherArgs
    ///             {
    ///                 Content = "example",
    ///                 JsonPathMatcher = new Gcp.Monitoring.Inputs.UptimeCheckConfigContentMatcherJsonPathMatcherArgs
    ///                 {
    ///                     JsonMatcher = "REGEX_MATCH",
    ///                     JsonPath = "$.path",
    ///                 },
    ///                 Matcher = "MATCHES_JSON_PATH",
    ///             },
    ///         },
    ///         DisplayName = "https-uptime-check",
    ///         HttpCheck = new Gcp.Monitoring.Inputs.UptimeCheckConfigHttpCheckArgs
    ///         {
    ///             Path = "/some-path",
    ///             Port = 443,
    ///             UseSsl = true,
    ///             ValidateSsl = true,
    ///         },
    ///         MonitoredResource = new Gcp.Monitoring.Inputs.UptimeCheckConfigMonitoredResourceArgs
    ///         {
    ///             Labels = 
    ///             {
    ///                 { "host", "192.168.1.1" },
    ///                 { "projectId", "my-project-name" },
    ///             },
    ///             Type = "uptime_url",
    ///         },
    ///         Timeout = "60s",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Uptime Check Tcp
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var check = new Gcp.Monitoring.Group("check", new()
    ///     {
    ///         DisplayName = "uptime-check-group",
    ///         Filter = "resource.metadata.name=has_substring(\"foo\")",
    ///     });
    /// 
    ///     var tcpGroup = new Gcp.Monitoring.UptimeCheckConfig("tcpGroup", new()
    ///     {
    ///         DisplayName = "tcp-uptime-check",
    ///         Timeout = "60s",
    ///         TcpCheck = new Gcp.Monitoring.Inputs.UptimeCheckConfigTcpCheckArgs
    ///         {
    ///             Port = 888,
    ///             PingConfig = new Gcp.Monitoring.Inputs.UptimeCheckConfigTcpCheckPingConfigArgs
    ///             {
    ///                 PingsCount = 2,
    ///             },
    ///         },
    ///         ResourceGroup = new Gcp.Monitoring.Inputs.UptimeCheckConfigResourceGroupArgs
    ///         {
    ///             ResourceType = "INSTANCE",
    ///             GroupId = check.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Uptime Check Config Synthetic Monitor
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "US",
    ///         UniformBucketLevelAccess = true,
    ///     });
    /// 
    ///     var @object = new Gcp.Storage.BucketObject("object", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Source = new FileAsset("synthetic-fn-source.zip"),
    ///     });
    /// 
    ///     // Add path to the zipped function source code
    ///     var function = new Gcp.CloudFunctionsV2.Function("function", new()
    ///     {
    ///         Location = "us-central1",
    ///         BuildConfig = new Gcp.CloudFunctionsV2.Inputs.FunctionBuildConfigArgs
    ///         {
    ///             Runtime = "nodejs16",
    ///             EntryPoint = "SyntheticFunction",
    ///             Source = new Gcp.CloudFunctionsV2.Inputs.FunctionBuildConfigSourceArgs
    ///             {
    ///                 StorageSource = new Gcp.CloudFunctionsV2.Inputs.FunctionBuildConfigSourceStorageSourceArgs
    ///                 {
    ///                     Bucket = bucket.Name,
    ///                     Object = @object.Name,
    ///                 },
    ///             },
    ///         },
    ///         ServiceConfig = new Gcp.CloudFunctionsV2.Inputs.FunctionServiceConfigArgs
    ///         {
    ///             MaxInstanceCount = 1,
    ///             AvailableMemory = "256M",
    ///             TimeoutSeconds = 60,
    ///         },
    ///     });
    /// 
    ///     var syntheticMonitor = new Gcp.Monitoring.UptimeCheckConfig("syntheticMonitor", new()
    ///     {
    ///         DisplayName = "synthetic_monitor",
    ///         Timeout = "60s",
    ///         SyntheticMonitor = new Gcp.Monitoring.Inputs.UptimeCheckConfigSyntheticMonitorArgs
    ///         {
    ///             CloudFunctionV2 = new Gcp.Monitoring.Inputs.UptimeCheckConfigSyntheticMonitorCloudFunctionV2Args
    ///             {
    ///                 Name = function.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// UptimeCheckConfig can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, UptimeCheckConfig can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig")]
    public partial class UptimeCheckConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The checker type to use for the check. If the monitored resource type is `servicedirectory_service`, `checker_type` must be set to `VPC_CHECKERS`.
        /// Possible values are: `STATIC_IP_CHECKERS`, `VPC_CHECKERS`.
        /// </summary>
        [Output("checkerType")]
        public Output<string> CheckerType { get; private set; } = null!;

        /// <summary>
        /// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
        /// Structure is documented below.
        /// </summary>
        [Output("contentMatchers")]
        public Output<ImmutableArray<Outputs.UptimeCheckConfigContentMatcher>> ContentMatchers { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.
        /// Structure is documented below.
        /// </summary>
        [Output("httpCheck")]
        public Output<Outputs.UptimeCheckConfigHttpCheck?> HttpCheck { get; private set; } = null!;

        /// <summary>
        /// The [monitored resource]
        /// (https://cloud.google.com/monitoring/api/resources) associated with the
        /// configuration. The following monitored resource types are supported for
        /// uptime checks:
        /// </summary>
        [Output("monitoredResource")]
        public Output<Outputs.UptimeCheckConfigMonitoredResource?> MonitoredResource { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of the cloud function resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        /// </summary>
        [Output("period")]
        public Output<string?> Period { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The group resource associated with the configuration.
        /// Structure is documented below.
        /// </summary>
        [Output("resourceGroup")]
        public Output<Outputs.UptimeCheckConfigResourceGroup?> ResourceGroup { get; private set; } = null!;

        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        /// </summary>
        [Output("selectedRegions")]
        public Output<ImmutableArray<string>> SelectedRegions { get; private set; } = null!;

        /// <summary>
        /// A Synthetic Monitor deployed to a Cloud Functions V2 instance.
        /// Structure is documented below.
        /// </summary>
        [Output("syntheticMonitor")]
        public Output<Outputs.UptimeCheckConfigSyntheticMonitor?> SyntheticMonitor { get; private set; } = null!;

        /// <summary>
        /// Contains information needed to make a TCP check.
        /// Structure is documented below.
        /// </summary>
        [Output("tcpCheck")]
        public Output<Outputs.UptimeCheckConfigTcpCheck?> TcpCheck { get; private set; } = null!;

        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). See the accepted formats
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("timeout")]
        public Output<string> Timeout { get; private set; } = null!;

        /// <summary>
        /// The id of the uptime check
        /// </summary>
        [Output("uptimeCheckId")]
        public Output<string> UptimeCheckId { get; private set; } = null!;

        /// <summary>
        /// User-supplied key/value data to be used for organizing and identifying the `UptimeCheckConfig` objects. The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        /// </summary>
        [Output("userLabels")]
        public Output<ImmutableDictionary<string, string>?> UserLabels { get; private set; } = null!;


        /// <summary>
        /// Create a UptimeCheckConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UptimeCheckConfig(string name, UptimeCheckConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, args ?? new UptimeCheckConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UptimeCheckConfig(string name, Input<string> id, UptimeCheckConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UptimeCheckConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UptimeCheckConfig Get(string name, Input<string> id, UptimeCheckConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new UptimeCheckConfig(name, id, state, options);
        }
    }

    public sealed class UptimeCheckConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The checker type to use for the check. If the monitored resource type is `servicedirectory_service`, `checker_type` must be set to `VPC_CHECKERS`.
        /// Possible values are: `STATIC_IP_CHECKERS`, `VPC_CHECKERS`.
        /// </summary>
        [Input("checkerType")]
        public Input<string>? CheckerType { get; set; }

        [Input("contentMatchers")]
        private InputList<Inputs.UptimeCheckConfigContentMatcherArgs>? _contentMatchers;

        /// <summary>
        /// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.UptimeCheckConfigContentMatcherArgs> ContentMatchers
        {
            get => _contentMatchers ?? (_contentMatchers = new InputList<Inputs.UptimeCheckConfigContentMatcherArgs>());
            set => _contentMatchers = value;
        }

        /// <summary>
        /// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.
        /// Structure is documented below.
        /// </summary>
        [Input("httpCheck")]
        public Input<Inputs.UptimeCheckConfigHttpCheckArgs>? HttpCheck { get; set; }

        /// <summary>
        /// The [monitored resource]
        /// (https://cloud.google.com/monitoring/api/resources) associated with the
        /// configuration. The following monitored resource types are supported for
        /// uptime checks:
        /// </summary>
        [Input("monitoredResource")]
        public Input<Inputs.UptimeCheckConfigMonitoredResourceArgs>? MonitoredResource { get; set; }

        /// <summary>
        /// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The group resource associated with the configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("resourceGroup")]
        public Input<Inputs.UptimeCheckConfigResourceGroupArgs>? ResourceGroup { get; set; }

        [Input("selectedRegions")]
        private InputList<string>? _selectedRegions;

        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        /// </summary>
        public InputList<string> SelectedRegions
        {
            get => _selectedRegions ?? (_selectedRegions = new InputList<string>());
            set => _selectedRegions = value;
        }

        /// <summary>
        /// A Synthetic Monitor deployed to a Cloud Functions V2 instance.
        /// Structure is documented below.
        /// </summary>
        [Input("syntheticMonitor")]
        public Input<Inputs.UptimeCheckConfigSyntheticMonitorArgs>? SyntheticMonitor { get; set; }

        /// <summary>
        /// Contains information needed to make a TCP check.
        /// Structure is documented below.
        /// </summary>
        [Input("tcpCheck")]
        public Input<Inputs.UptimeCheckConfigTcpCheckArgs>? TcpCheck { get; set; }

        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). See the accepted formats
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("timeout", required: true)]
        public Input<string> Timeout { get; set; } = null!;

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// User-supplied key/value data to be used for organizing and identifying the `UptimeCheckConfig` objects. The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public UptimeCheckConfigArgs()
        {
        }
        public static new UptimeCheckConfigArgs Empty => new UptimeCheckConfigArgs();
    }

    public sealed class UptimeCheckConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The checker type to use for the check. If the monitored resource type is `servicedirectory_service`, `checker_type` must be set to `VPC_CHECKERS`.
        /// Possible values are: `STATIC_IP_CHECKERS`, `VPC_CHECKERS`.
        /// </summary>
        [Input("checkerType")]
        public Input<string>? CheckerType { get; set; }

        [Input("contentMatchers")]
        private InputList<Inputs.UptimeCheckConfigContentMatcherGetArgs>? _contentMatchers;

        /// <summary>
        /// The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.UptimeCheckConfigContentMatcherGetArgs> ContentMatchers
        {
            get => _contentMatchers ?? (_contentMatchers = new InputList<Inputs.UptimeCheckConfigContentMatcherGetArgs>());
            set => _contentMatchers = value;
        }

        /// <summary>
        /// A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Contains information needed to make an HTTP or HTTPS check.
        /// Structure is documented below.
        /// </summary>
        [Input("httpCheck")]
        public Input<Inputs.UptimeCheckConfigHttpCheckGetArgs>? HttpCheck { get; set; }

        /// <summary>
        /// The [monitored resource]
        /// (https://cloud.google.com/monitoring/api/resources) associated with the
        /// configuration. The following monitored resource types are supported for
        /// uptime checks:
        /// </summary>
        [Input("monitoredResource")]
        public Input<Inputs.UptimeCheckConfigMonitoredResourceGetArgs>? MonitoredResource { get; set; }

        /// <summary>
        /// The fully qualified name of the cloud function resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The group resource associated with the configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("resourceGroup")]
        public Input<Inputs.UptimeCheckConfigResourceGroupGetArgs>? ResourceGroup { get; set; }

        [Input("selectedRegions")]
        private InputList<string>? _selectedRegions;

        /// <summary>
        /// The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.
        /// </summary>
        public InputList<string> SelectedRegions
        {
            get => _selectedRegions ?? (_selectedRegions = new InputList<string>());
            set => _selectedRegions = value;
        }

        /// <summary>
        /// A Synthetic Monitor deployed to a Cloud Functions V2 instance.
        /// Structure is documented below.
        /// </summary>
        [Input("syntheticMonitor")]
        public Input<Inputs.UptimeCheckConfigSyntheticMonitorGetArgs>? SyntheticMonitor { get; set; }

        /// <summary>
        /// Contains information needed to make a TCP check.
        /// Structure is documented below.
        /// </summary>
        [Input("tcpCheck")]
        public Input<Inputs.UptimeCheckConfigTcpCheckGetArgs>? TcpCheck { get; set; }

        /// <summary>
        /// The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). See the accepted formats
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        /// <summary>
        /// The id of the uptime check
        /// </summary>
        [Input("uptimeCheckId")]
        public Input<string>? UptimeCheckId { get; set; }

        [Input("userLabels")]
        private InputMap<string>? _userLabels;

        /// <summary>
        /// User-supplied key/value data to be used for organizing and identifying the `UptimeCheckConfig` objects. The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
        /// </summary>
        public InputMap<string> UserLabels
        {
            get => _userLabels ?? (_userLabels = new InputMap<string>());
            set => _userLabels = value;
        }

        public UptimeCheckConfigState()
        {
        }
        public static new UptimeCheckConfigState Empty => new UptimeCheckConfigState();
    }
}
