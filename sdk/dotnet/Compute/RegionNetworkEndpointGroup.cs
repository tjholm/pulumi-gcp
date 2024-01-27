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
    /// A regional NEG that can support Serverless Products.
    /// 
    /// To get more information about RegionNetworkEndpointGroup, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/regionNetworkEndpointGroups)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/serverless-neg-concepts)
    /// 
    /// ## Example Usage
    /// ### Region Network Endpoint Group Functions
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
    ///     });
    /// 
    ///     var archive = new Gcp.Storage.BucketObject("archive", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Source = new FileAsset("path/to/index.zip"),
    ///     });
    /// 
    ///     var functionNegFunction = new Gcp.CloudFunctions.Function("functionNegFunction", new()
    ///     {
    ///         Description = "My function",
    ///         Runtime = "nodejs10",
    ///         AvailableMemoryMb = 128,
    ///         SourceArchiveBucket = bucket.Name,
    ///         SourceArchiveObject = archive.Name,
    ///         TriggerHttp = true,
    ///         Timeout = 60,
    ///         EntryPoint = "helloGET",
    ///     });
    /// 
    ///     // Cloud Functions Example
    ///     var functionNegRegionNetworkEndpointGroup = new Gcp.Compute.RegionNetworkEndpointGroup("functionNegRegionNetworkEndpointGroup", new()
    ///     {
    ///         NetworkEndpointType = "SERVERLESS",
    ///         Region = "us-central1",
    ///         CloudFunction = new Gcp.Compute.Inputs.RegionNetworkEndpointGroupCloudFunctionArgs
    ///         {
    ///             Function = functionNegFunction.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Region Network Endpoint Group Cloudrun
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cloudrunNegService = new Gcp.CloudRun.Service("cloudrunNegService", new()
    ///     {
    ///         Location = "us-central1",
    ///         Template = new Gcp.CloudRun.Inputs.ServiceTemplateArgs
    ///         {
    ///             Spec = new Gcp.CloudRun.Inputs.ServiceTemplateSpecArgs
    ///             {
    ///                 Containers = new[]
    ///                 {
    ///                     new Gcp.CloudRun.Inputs.ServiceTemplateSpecContainerArgs
    ///                     {
    ///                         Image = "us-docker.pkg.dev/cloudrun/container/hello",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         Traffics = new[]
    ///         {
    ///             new Gcp.CloudRun.Inputs.ServiceTrafficArgs
    ///             {
    ///                 Percent = 100,
    ///                 LatestRevision = true,
    ///             },
    ///         },
    ///     });
    /// 
    ///     // Cloud Run Example
    ///     var cloudrunNegRegionNetworkEndpointGroup = new Gcp.Compute.RegionNetworkEndpointGroup("cloudrunNegRegionNetworkEndpointGroup", new()
    ///     {
    ///         NetworkEndpointType = "SERVERLESS",
    ///         Region = "us-central1",
    ///         CloudRun = new Gcp.Compute.Inputs.RegionNetworkEndpointGroupCloudRunArgs
    ///         {
    ///             Service = cloudrunNegService.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Region Network Endpoint Group Appengine
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var appengineNegBucket = new Gcp.Storage.Bucket("appengineNegBucket", new()
    ///     {
    ///         Location = "US",
    ///     });
    /// 
    ///     var appengineNegBucketObject = new Gcp.Storage.BucketObject("appengineNegBucketObject", new()
    ///     {
    ///         Bucket = appengineNegBucket.Name,
    ///         Source = new FileAsset("./test-fixtures/hello-world.zip"),
    ///     });
    /// 
    ///     var appengineNegFlexibleAppVersion = new Gcp.AppEngine.FlexibleAppVersion("appengineNegFlexibleAppVersion", new()
    ///     {
    ///         VersionId = "v1",
    ///         Service = "appengine-network-endpoint-group",
    ///         Runtime = "nodejs",
    ///         Entrypoint = new Gcp.AppEngine.Inputs.FlexibleAppVersionEntrypointArgs
    ///         {
    ///             Shell = "node ./app.js",
    ///         },
    ///         Deployment = new Gcp.AppEngine.Inputs.FlexibleAppVersionDeploymentArgs
    ///         {
    ///             Zip = new Gcp.AppEngine.Inputs.FlexibleAppVersionDeploymentZipArgs
    ///             {
    ///                 SourceUrl = Output.Tuple(appengineNegBucket.Name, appengineNegBucketObject.Name).Apply(values =&gt;
    ///                 {
    ///                     var appengineNegBucketName = values.Item1;
    ///                     var appengineNegBucketObjectName = values.Item2;
    ///                     return $"https://storage.googleapis.com/{appengineNegBucketName}/{appengineNegBucketObjectName}";
    ///                 }),
    ///             },
    ///         },
    ///         LivenessCheck = new Gcp.AppEngine.Inputs.FlexibleAppVersionLivenessCheckArgs
    ///         {
    ///             Path = "/",
    ///         },
    ///         ReadinessCheck = new Gcp.AppEngine.Inputs.FlexibleAppVersionReadinessCheckArgs
    ///         {
    ///             Path = "/",
    ///         },
    ///         EnvVariables = 
    ///         {
    ///             { "port", "8080" },
    ///         },
    ///         Handlers = new[]
    ///         {
    ///             new Gcp.AppEngine.Inputs.FlexibleAppVersionHandlerArgs
    ///             {
    ///                 UrlRegex = ".*\\/my-path\\/*",
    ///                 SecurityLevel = "SECURE_ALWAYS",
    ///                 Login = "LOGIN_REQUIRED",
    ///                 AuthFailAction = "AUTH_FAIL_ACTION_REDIRECT",
    ///                 StaticFiles = new Gcp.AppEngine.Inputs.FlexibleAppVersionHandlerStaticFilesArgs
    ///                 {
    ///                     Path = "my-other-path",
    ///                     UploadPathRegex = ".*\\/my-path\\/*",
    ///                 },
    ///             },
    ///         },
    ///         AutomaticScaling = new Gcp.AppEngine.Inputs.FlexibleAppVersionAutomaticScalingArgs
    ///         {
    ///             CoolDownPeriod = "120s",
    ///             CpuUtilization = new Gcp.AppEngine.Inputs.FlexibleAppVersionAutomaticScalingCpuUtilizationArgs
    ///             {
    ///                 TargetUtilization = 0.5,
    ///             },
    ///         },
    ///         DeleteServiceOnDestroy = true,
    ///     });
    /// 
    ///     // App Engine Example
    ///     var appengineNegRegionNetworkEndpointGroup = new Gcp.Compute.RegionNetworkEndpointGroup("appengineNegRegionNetworkEndpointGroup", new()
    ///     {
    ///         NetworkEndpointType = "SERVERLESS",
    ///         Region = "us-central1",
    ///         AppEngine = new Gcp.Compute.Inputs.RegionNetworkEndpointGroupAppEngineArgs
    ///         {
    ///             Service = appengineNegFlexibleAppVersion.Service,
    ///             Version = appengineNegFlexibleAppVersion.VersionId,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Region Network Endpoint Group Psc
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pscNeg = new Gcp.Compute.RegionNetworkEndpointGroup("pscNeg", new()
    ///     {
    ///         NetworkEndpointType = "PRIVATE_SERVICE_CONNECT",
    ///         PscTargetService = "asia-northeast3-cloudkms.googleapis.com",
    ///         Region = "asia-northeast3",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RegionNetworkEndpointGroup can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, RegionNetworkEndpointGroup can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default projects/{{project}}/regions/{{region}}/networkEndpointGroups/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup")]
    public partial class RegionNetworkEndpointGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Output("appEngine")]
        public Output<Outputs.RegionNetworkEndpointGroupAppEngine?> AppEngine { get; private set; } = null!;

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Output("cloudFunction")]
        public Output<Outputs.RegionNetworkEndpointGroupCloudFunction?> CloudFunction { get; private set; } = null!;

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Output("cloudRun")]
        public Output<Outputs.RegionNetworkEndpointGroupCloudRun?> CloudRun { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource; provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// This field is only used for PSC.
        /// The URL of the network to which all network endpoints in the NEG belong. Uses
        /// "default" project network if unspecified.
        /// </summary>
        [Output("network")]
        public Output<string?> Network { get; private set; } = null!;

        /// <summary>
        /// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
        /// Default value is `SERVERLESS`.
        /// Possible values are: `SERVERLESS`, `PRIVATE_SERVICE_CONNECT`.
        /// </summary>
        [Output("networkEndpointType")]
        public Output<string?> NetworkEndpointType { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The target service url used to set up private service connection to
        /// a Google API or a PSC Producer Service Attachment.
        /// </summary>
        [Output("pscTargetService")]
        public Output<string?> PscTargetService { get; private set; } = null!;

        /// <summary>
        /// A reference to the region where the Serverless NEGs Reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// (Optional, Beta)
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Output("serverlessDeployment")]
        public Output<Outputs.RegionNetworkEndpointGroupServerlessDeployment?> ServerlessDeployment { get; private set; } = null!;

        /// <summary>
        /// This field is only used for PSC.
        /// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        [Output("subnetwork")]
        public Output<string?> Subnetwork { get; private set; } = null!;


        /// <summary>
        /// Create a RegionNetworkEndpointGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionNetworkEndpointGroup(string name, RegionNetworkEndpointGroupArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup", name, args ?? new RegionNetworkEndpointGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionNetworkEndpointGroup(string name, Input<string> id, RegionNetworkEndpointGroupState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/regionNetworkEndpointGroup:RegionNetworkEndpointGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegionNetworkEndpointGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionNetworkEndpointGroup Get(string name, Input<string> id, RegionNetworkEndpointGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionNetworkEndpointGroup(name, id, state, options);
        }
    }

    public sealed class RegionNetworkEndpointGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("appEngine")]
        public Input<Inputs.RegionNetworkEndpointGroupAppEngineArgs>? AppEngine { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("cloudFunction")]
        public Input<Inputs.RegionNetworkEndpointGroupCloudFunctionArgs>? CloudFunction { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("cloudRun")]
        public Input<Inputs.RegionNetworkEndpointGroupCloudRunArgs>? CloudRun { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is only used for PSC.
        /// The URL of the network to which all network endpoints in the NEG belong. Uses
        /// "default" project network if unspecified.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
        /// Default value is `SERVERLESS`.
        /// Possible values are: `SERVERLESS`, `PRIVATE_SERVICE_CONNECT`.
        /// </summary>
        [Input("networkEndpointType")]
        public Input<string>? NetworkEndpointType { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The target service url used to set up private service connection to
        /// a Google API or a PSC Producer Service Attachment.
        /// </summary>
        [Input("pscTargetService")]
        public Input<string>? PscTargetService { get; set; }

        /// <summary>
        /// A reference to the region where the Serverless NEGs Reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// (Optional, Beta)
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("serverlessDeployment")]
        public Input<Inputs.RegionNetworkEndpointGroupServerlessDeploymentArgs>? ServerlessDeployment { get; set; }

        /// <summary>
        /// This field is only used for PSC.
        /// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        public RegionNetworkEndpointGroupArgs()
        {
        }
        public static new RegionNetworkEndpointGroupArgs Empty => new RegionNetworkEndpointGroupArgs();
    }

    public sealed class RegionNetworkEndpointGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("appEngine")]
        public Input<Inputs.RegionNetworkEndpointGroupAppEngineGetArgs>? AppEngine { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("cloudFunction")]
        public Input<Inputs.RegionNetworkEndpointGroupCloudFunctionGetArgs>? CloudFunction { get; set; }

        /// <summary>
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("cloudRun")]
        public Input<Inputs.RegionNetworkEndpointGroupCloudRunGetArgs>? CloudRun { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when
        /// you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource; provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035. Specifically, the name must be 1-63 characters long and match
        /// the regular expression `a-z?` which means the
        /// first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the last
        /// character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// This field is only used for PSC.
        /// The URL of the network to which all network endpoints in the NEG belong. Uses
        /// "default" project network if unspecified.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// Type of network endpoints in this network endpoint group. Defaults to SERVERLESS
        /// Default value is `SERVERLESS`.
        /// Possible values are: `SERVERLESS`, `PRIVATE_SERVICE_CONNECT`.
        /// </summary>
        [Input("networkEndpointType")]
        public Input<string>? NetworkEndpointType { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The target service url used to set up private service connection to
        /// a Google API or a PSC Producer Service Attachment.
        /// </summary>
        [Input("pscTargetService")]
        public Input<string>? PscTargetService { get; set; }

        /// <summary>
        /// A reference to the region where the Serverless NEGs Reside.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// Only valid when networkEndpointType is "SERVERLESS".
        /// Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
        /// Structure is documented below.
        /// </summary>
        [Input("serverlessDeployment")]
        public Input<Inputs.RegionNetworkEndpointGroupServerlessDeploymentGetArgs>? ServerlessDeployment { get; set; }

        /// <summary>
        /// This field is only used for PSC.
        /// Optional URL of the subnetwork to which all network endpoints in the NEG belong.
        /// </summary>
        [Input("subnetwork")]
        public Input<string>? Subnetwork { get; set; }

        public RegionNetworkEndpointGroupState()
        {
        }
        public static new RegionNetworkEndpointGroupState Empty => new RegionNetworkEndpointGroupState();
    }
}
