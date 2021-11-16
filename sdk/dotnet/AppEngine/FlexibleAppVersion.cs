// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine
{
    /// <summary>
    /// Flexible App Version resource to create a new version of flexible GAE Application. Based on Google Compute Engine,
    /// the App Engine flexible environment automatically scales your app up and down while also balancing the load.
    /// Learn about the differences between the standard environment and the flexible environment
    /// at https://cloud.google.com/appengine/docs/the-appengine-environments.
    /// 
    /// &gt; **Note:** The App Engine flexible environment service account uses the member ID `service-[YOUR_PROJECT_NUMBER]@gae-api-prod.google.com.iam.gserviceaccount.com`
    /// It should have the App Engine Flexible Environment Service Agent role, which will be applied when the `appengineflex.googleapis.com` service is enabled.
    /// 
    /// To get more information about FlexibleAppVersion, see:
    /// 
    /// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/appengine/docs/flexible)
    /// 
    /// ## Example Usage
    /// ### App Engine Flexible App Version
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myProject = new Gcp.Organizations.Project("myProject", new Gcp.Organizations.ProjectArgs
    ///         {
    ///             ProjectId = "appeng-flex",
    ///             OrgId = "123456789",
    ///             BillingAccount = "000000-0000000-0000000-000000",
    ///         });
    ///         var app = new Gcp.AppEngine.Application("app", new Gcp.AppEngine.ApplicationArgs
    ///         {
    ///             Project = myProject.ProjectId,
    ///             LocationId = "us-central",
    ///         });
    ///         var service = new Gcp.Projects.Service("service", new Gcp.Projects.ServiceArgs
    ///         {
    ///             Project = myProject.ProjectId,
    ///             Service = "appengineflex.googleapis.com",
    ///             DisableDependentServices = false,
    ///         });
    ///         var gaeApi = new Gcp.Projects.IAMMember("gaeApi", new Gcp.Projects.IAMMemberArgs
    ///         {
    ///             Project = service.Project,
    ///             Role = "roles/compute.networkUser",
    ///             Member = myProject.Number.Apply(number =&gt; $"serviceAccount:service-{number}@gae-api-prod.google.com.iam.gserviceaccount.com"),
    ///         });
    ///         var bucket = new Gcp.Storage.Bucket("bucket", new Gcp.Storage.BucketArgs
    ///         {
    ///             Project = myProject.ProjectId,
    ///             Location = "US",
    ///         });
    ///         var @object = new Gcp.Storage.BucketObject("object", new Gcp.Storage.BucketObjectArgs
    ///         {
    ///             Bucket = bucket.Name,
    ///             Source = new FileAsset("./test-fixtures/appengine/hello-world.zip"),
    ///         });
    ///         var myappV1 = new Gcp.AppEngine.FlexibleAppVersion("myappV1", new Gcp.AppEngine.FlexibleAppVersionArgs
    ///         {
    ///             VersionId = "v1",
    ///             Project = gaeApi.Project,
    ///             Service = "default",
    ///             Runtime = "nodejs",
    ///             Entrypoint = new Gcp.AppEngine.Inputs.FlexibleAppVersionEntrypointArgs
    ///             {
    ///                 Shell = "node ./app.js",
    ///             },
    ///             Deployment = new Gcp.AppEngine.Inputs.FlexibleAppVersionDeploymentArgs
    ///             {
    ///                 Zip = new Gcp.AppEngine.Inputs.FlexibleAppVersionDeploymentZipArgs
    ///                 {
    ///                     SourceUrl = Output.Tuple(bucket.Name, @object.Name).Apply(values =&gt;
    ///                     {
    ///                         var bucketName = values.Item1;
    ///                         var objectName = values.Item2;
    ///                         return $"https://storage.googleapis.com/{bucketName}/{objectName}";
    ///                     }),
    ///                 },
    ///             },
    ///             LivenessCheck = new Gcp.AppEngine.Inputs.FlexibleAppVersionLivenessCheckArgs
    ///             {
    ///                 Path = "/",
    ///             },
    ///             ReadinessCheck = new Gcp.AppEngine.Inputs.FlexibleAppVersionReadinessCheckArgs
    ///             {
    ///                 Path = "/",
    ///             },
    ///             EnvVariables = 
    ///             {
    ///                 { "port", "8080" },
    ///             },
    ///             Handlers = 
    ///             {
    ///                 new Gcp.AppEngine.Inputs.FlexibleAppVersionHandlerArgs
    ///                 {
    ///                     UrlRegex = ".*\\/my-path\\/*",
    ///                     SecurityLevel = "SECURE_ALWAYS",
    ///                     Login = "LOGIN_REQUIRED",
    ///                     AuthFailAction = "AUTH_FAIL_ACTION_REDIRECT",
    ///                     StaticFiles = new Gcp.AppEngine.Inputs.FlexibleAppVersionHandlerStaticFilesArgs
    ///                     {
    ///                         Path = "my-other-path",
    ///                         UploadPathRegex = ".*\\/my-path\\/*",
    ///                     },
    ///                 },
    ///             },
    ///             AutomaticScaling = new Gcp.AppEngine.Inputs.FlexibleAppVersionAutomaticScalingArgs
    ///             {
    ///                 CoolDownPeriod = "120s",
    ///                 CpuUtilization = new Gcp.AppEngine.Inputs.FlexibleAppVersionAutomaticScalingCpuUtilizationArgs
    ///                 {
    ///                     TargetUtilization = 0.5,
    ///                 },
    ///             },
    ///             NoopOnDestroy = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// FlexibleAppVersion can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/flexibleAppVersion:FlexibleAppVersion default apps/{{project}}/services/{{service}}/versions/{{version_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/flexibleAppVersion:FlexibleAppVersion default {{project}}/{{service}}/{{version_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/flexibleAppVersion:FlexibleAppVersion default {{service}}/{{version_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:appengine/flexibleAppVersion:FlexibleAppVersion")]
    public partial class FlexibleAppVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// Serving configuration for Google Cloud Endpoints.
        /// Structure is documented below.
        /// </summary>
        [Output("apiConfig")]
        public Output<Outputs.FlexibleAppVersionApiConfig?> ApiConfig { get; private set; } = null!;

        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics.
        /// Structure is documented below.
        /// </summary>
        [Output("automaticScaling")]
        public Output<Outputs.FlexibleAppVersionAutomaticScaling?> AutomaticScaling { get; private set; } = null!;

        /// <summary>
        /// Metadata settings that are supplied to this version to enable beta runtime features.
        /// </summary>
        [Output("betaSettings")]
        public Output<ImmutableDictionary<string, string>?> BetaSettings { get; private set; } = null!;

        /// <summary>
        /// Duration that static files should be cached by web proxies and browsers.
        /// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
        /// </summary>
        [Output("defaultExpiration")]
        public Output<string?> DefaultExpiration { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the service will be deleted if it is the last version.
        /// </summary>
        [Output("deleteServiceOnDestroy")]
        public Output<bool?> DeleteServiceOnDestroy { get; private set; } = null!;

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Output("deployment")]
        public Output<Outputs.FlexibleAppVersionDeployment?> Deployment { get; private set; } = null!;

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Output("endpointsApiService")]
        public Output<Outputs.FlexibleAppVersionEndpointsApiService?> EndpointsApiService { get; private set; } = null!;

        /// <summary>
        /// The entrypoint for the application.
        /// Structure is documented below.
        /// </summary>
        [Output("entrypoint")]
        public Output<Outputs.FlexibleAppVersionEntrypoint?> Entrypoint { get; private set; } = null!;

        /// <summary>
        /// Environment variables available to the application.  As these are not returned in the API request, the provider will not detect any changes made outside of the config.
        /// </summary>
        [Output("envVariables")]
        public Output<ImmutableDictionary<string, string>?> EnvVariables { get; private set; } = null!;

        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests.
        /// The first matching URL handles the request and other request handlers are not attempted.
        /// Structure is documented below.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.FlexibleAppVersionHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// A list of the types of messages that this application is able to receive.
        /// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
        /// </summary>
        [Output("inboundServices")]
        public Output<ImmutableArray<string>> InboundServices { get; private set; } = null!;

        /// <summary>
        /// Instance class that is used to run this version. Valid values are
        /// AutomaticScaling: F1, F2, F4, F4_1G
        /// ManualScaling: B1, B2, B4, B8, B4_1G
        /// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
        /// </summary>
        [Output("instanceClass")]
        public Output<string?> InstanceClass { get; private set; } = null!;

        /// <summary>
        /// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
        /// Structure is documented below.
        /// </summary>
        [Output("livenessCheck")]
        public Output<Outputs.FlexibleAppVersionLivenessCheck> LivenessCheck { get; private set; } = null!;

        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        /// Structure is documented below.
        /// </summary>
        [Output("manualScaling")]
        public Output<Outputs.FlexibleAppVersionManualScaling?> ManualScaling { get; private set; } = null!;

        /// <summary>
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Extra network settings
        /// Structure is documented below.
        /// </summary>
        [Output("network")]
        public Output<Outputs.FlexibleAppVersionNetwork?> Network { get; private set; } = null!;

        /// <summary>
        /// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
        /// </summary>
        [Output("nobuildFilesRegex")]
        public Output<string?> NobuildFilesRegex { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, the application version will not be deleted.
        /// </summary>
        [Output("noopOnDestroy")]
        public Output<bool?> NoopOnDestroy { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.
        /// Structure is documented below.
        /// </summary>
        [Output("readinessCheck")]
        public Output<Outputs.FlexibleAppVersionReadinessCheck> ReadinessCheck { get; private set; } = null!;

        /// <summary>
        /// Machine resources for a version.
        /// Structure is documented below.
        /// </summary>
        [Output("resources")]
        public Output<Outputs.FlexibleAppVersionResources?> Resources { get; private set; } = null!;

        /// <summary>
        /// Desired runtime. Example python27.
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// The version of the API in the given runtime environment.
        /// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        /// </summary>
        [Output("runtimeApiVersion")]
        public Output<string> RuntimeApiVersion { get; private set; } = null!;

        /// <summary>
        /// The channel of the runtime to use. Only available for some runtimes.
        /// </summary>
        [Output("runtimeChannel")]
        public Output<string?> RuntimeChannel { get; private set; } = null!;

        /// <summary>
        /// The path or name of the app's main executable.
        /// </summary>
        [Output("runtimeMainExecutablePath")]
        public Output<string?> RuntimeMainExecutablePath { get; private set; } = null!;

        /// <summary>
        /// AppEngine service resource. Can contain numbers, letters, and hyphens.
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
        /// Default value is `SERVING`.
        /// Possible values are `SERVING` and `STOPPED`.
        /// </summary>
        [Output("servingStatus")]
        public Output<string?> ServingStatus { get; private set; } = null!;

        /// <summary>
        /// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
        /// Reserved names,"default", "latest", and any name with the prefix "ah-".
        /// </summary>
        [Output("versionId")]
        public Output<string?> VersionId { get; private set; } = null!;

        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// Structure is documented below.
        /// </summary>
        [Output("vpcAccessConnector")]
        public Output<Outputs.FlexibleAppVersionVpcAccessConnector?> VpcAccessConnector { get; private set; } = null!;


        /// <summary>
        /// Create a FlexibleAppVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlexibleAppVersion(string name, FlexibleAppVersionArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/flexibleAppVersion:FlexibleAppVersion", name, args ?? new FlexibleAppVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlexibleAppVersion(string name, Input<string> id, FlexibleAppVersionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/flexibleAppVersion:FlexibleAppVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FlexibleAppVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlexibleAppVersion Get(string name, Input<string> id, FlexibleAppVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new FlexibleAppVersion(name, id, state, options);
        }
    }

    public sealed class FlexibleAppVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Serving configuration for Google Cloud Endpoints.
        /// Structure is documented below.
        /// </summary>
        [Input("apiConfig")]
        public Input<Inputs.FlexibleAppVersionApiConfigArgs>? ApiConfig { get; set; }

        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics.
        /// Structure is documented below.
        /// </summary>
        [Input("automaticScaling")]
        public Input<Inputs.FlexibleAppVersionAutomaticScalingArgs>? AutomaticScaling { get; set; }

        [Input("betaSettings")]
        private InputMap<string>? _betaSettings;

        /// <summary>
        /// Metadata settings that are supplied to this version to enable beta runtime features.
        /// </summary>
        public InputMap<string> BetaSettings
        {
            get => _betaSettings ?? (_betaSettings = new InputMap<string>());
            set => _betaSettings = value;
        }

        /// <summary>
        /// Duration that static files should be cached by web proxies and browsers.
        /// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
        /// </summary>
        [Input("defaultExpiration")]
        public Input<string>? DefaultExpiration { get; set; }

        /// <summary>
        /// If set to `true`, the service will be deleted if it is the last version.
        /// </summary>
        [Input("deleteServiceOnDestroy")]
        public Input<bool>? DeleteServiceOnDestroy { get; set; }

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.FlexibleAppVersionDeploymentArgs>? Deployment { get; set; }

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Input("endpointsApiService")]
        public Input<Inputs.FlexibleAppVersionEndpointsApiServiceArgs>? EndpointsApiService { get; set; }

        /// <summary>
        /// The entrypoint for the application.
        /// Structure is documented below.
        /// </summary>
        [Input("entrypoint")]
        public Input<Inputs.FlexibleAppVersionEntrypointArgs>? Entrypoint { get; set; }

        [Input("envVariables")]
        private InputMap<string>? _envVariables;

        /// <summary>
        /// Environment variables available to the application.  As these are not returned in the API request, the provider will not detect any changes made outside of the config.
        /// </summary>
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        [Input("handlers")]
        private InputList<Inputs.FlexibleAppVersionHandlerArgs>? _handlers;

        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests.
        /// The first matching URL handles the request and other request handlers are not attempted.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FlexibleAppVersionHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.FlexibleAppVersionHandlerArgs>());
            set => _handlers = value;
        }

        [Input("inboundServices")]
        private InputList<string>? _inboundServices;

        /// <summary>
        /// A list of the types of messages that this application is able to receive.
        /// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
        /// </summary>
        public InputList<string> InboundServices
        {
            get => _inboundServices ?? (_inboundServices = new InputList<string>());
            set => _inboundServices = value;
        }

        /// <summary>
        /// Instance class that is used to run this version. Valid values are
        /// AutomaticScaling: F1, F2, F4, F4_1G
        /// ManualScaling: B1, B2, B4, B8, B4_1G
        /// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
        /// Structure is documented below.
        /// </summary>
        [Input("livenessCheck", required: true)]
        public Input<Inputs.FlexibleAppVersionLivenessCheckArgs> LivenessCheck { get; set; } = null!;

        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        /// Structure is documented below.
        /// </summary>
        [Input("manualScaling")]
        public Input<Inputs.FlexibleAppVersionManualScalingArgs>? ManualScaling { get; set; }

        /// <summary>
        /// Extra network settings
        /// Structure is documented below.
        /// </summary>
        [Input("network")]
        public Input<Inputs.FlexibleAppVersionNetworkArgs>? Network { get; set; }

        /// <summary>
        /// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
        /// </summary>
        [Input("nobuildFilesRegex")]
        public Input<string>? NobuildFilesRegex { get; set; }

        /// <summary>
        /// If set to `true`, the application version will not be deleted.
        /// </summary>
        [Input("noopOnDestroy")]
        public Input<bool>? NoopOnDestroy { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.
        /// Structure is documented below.
        /// </summary>
        [Input("readinessCheck", required: true)]
        public Input<Inputs.FlexibleAppVersionReadinessCheckArgs> ReadinessCheck { get; set; } = null!;

        /// <summary>
        /// Machine resources for a version.
        /// Structure is documented below.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.FlexibleAppVersionResourcesArgs>? Resources { get; set; }

        /// <summary>
        /// Desired runtime. Example python27.
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        /// <summary>
        /// The version of the API in the given runtime environment.
        /// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        /// </summary>
        [Input("runtimeApiVersion")]
        public Input<string>? RuntimeApiVersion { get; set; }

        /// <summary>
        /// The channel of the runtime to use. Only available for some runtimes.
        /// </summary>
        [Input("runtimeChannel")]
        public Input<string>? RuntimeChannel { get; set; }

        /// <summary>
        /// The path or name of the app's main executable.
        /// </summary>
        [Input("runtimeMainExecutablePath")]
        public Input<string>? RuntimeMainExecutablePath { get; set; }

        /// <summary>
        /// AppEngine service resource. Can contain numbers, letters, and hyphens.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
        /// Default value is `SERVING`.
        /// Possible values are `SERVING` and `STOPPED`.
        /// </summary>
        [Input("servingStatus")]
        public Input<string>? ServingStatus { get; set; }

        /// <summary>
        /// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
        /// Reserved names,"default", "latest", and any name with the prefix "ah-".
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// Structure is documented below.
        /// </summary>
        [Input("vpcAccessConnector")]
        public Input<Inputs.FlexibleAppVersionVpcAccessConnectorArgs>? VpcAccessConnector { get; set; }

        public FlexibleAppVersionArgs()
        {
        }
    }

    public sealed class FlexibleAppVersionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Serving configuration for Google Cloud Endpoints.
        /// Structure is documented below.
        /// </summary>
        [Input("apiConfig")]
        public Input<Inputs.FlexibleAppVersionApiConfigGetArgs>? ApiConfig { get; set; }

        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics.
        /// Structure is documented below.
        /// </summary>
        [Input("automaticScaling")]
        public Input<Inputs.FlexibleAppVersionAutomaticScalingGetArgs>? AutomaticScaling { get; set; }

        [Input("betaSettings")]
        private InputMap<string>? _betaSettings;

        /// <summary>
        /// Metadata settings that are supplied to this version to enable beta runtime features.
        /// </summary>
        public InputMap<string> BetaSettings
        {
            get => _betaSettings ?? (_betaSettings = new InputMap<string>());
            set => _betaSettings = value;
        }

        /// <summary>
        /// Duration that static files should be cached by web proxies and browsers.
        /// Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
        /// </summary>
        [Input("defaultExpiration")]
        public Input<string>? DefaultExpiration { get; set; }

        /// <summary>
        /// If set to `true`, the service will be deleted if it is the last version.
        /// </summary>
        [Input("deleteServiceOnDestroy")]
        public Input<bool>? DeleteServiceOnDestroy { get; set; }

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Input("deployment")]
        public Input<Inputs.FlexibleAppVersionDeploymentGetArgs>? Deployment { get; set; }

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Input("endpointsApiService")]
        public Input<Inputs.FlexibleAppVersionEndpointsApiServiceGetArgs>? EndpointsApiService { get; set; }

        /// <summary>
        /// The entrypoint for the application.
        /// Structure is documented below.
        /// </summary>
        [Input("entrypoint")]
        public Input<Inputs.FlexibleAppVersionEntrypointGetArgs>? Entrypoint { get; set; }

        [Input("envVariables")]
        private InputMap<string>? _envVariables;

        /// <summary>
        /// Environment variables available to the application.  As these are not returned in the API request, the provider will not detect any changes made outside of the config.
        /// </summary>
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        [Input("handlers")]
        private InputList<Inputs.FlexibleAppVersionHandlerGetArgs>? _handlers;

        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests.
        /// The first matching URL handles the request and other request handlers are not attempted.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FlexibleAppVersionHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.FlexibleAppVersionHandlerGetArgs>());
            set => _handlers = value;
        }

        [Input("inboundServices")]
        private InputList<string>? _inboundServices;

        /// <summary>
        /// A list of the types of messages that this application is able to receive.
        /// Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
        /// </summary>
        public InputList<string> InboundServices
        {
            get => _inboundServices ?? (_inboundServices = new InputList<string>());
            set => _inboundServices = value;
        }

        /// <summary>
        /// Instance class that is used to run this version. Valid values are
        /// AutomaticScaling: F1, F2, F4, F4_1G
        /// ManualScaling: B1, B2, B4, B8, B4_1G
        /// Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        /// <summary>
        /// Health checking configuration for VM instances. Unhealthy instances are killed and replaced with new instances.
        /// Structure is documented below.
        /// </summary>
        [Input("livenessCheck")]
        public Input<Inputs.FlexibleAppVersionLivenessCheckGetArgs>? LivenessCheck { get; set; }

        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        /// Structure is documented below.
        /// </summary>
        [Input("manualScaling")]
        public Input<Inputs.FlexibleAppVersionManualScalingGetArgs>? ManualScaling { get; set; }

        /// <summary>
        /// Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Extra network settings
        /// Structure is documented below.
        /// </summary>
        [Input("network")]
        public Input<Inputs.FlexibleAppVersionNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// Files that match this pattern will not be built into this version. Only applicable for Go runtimes.
        /// </summary>
        [Input("nobuildFilesRegex")]
        public Input<string>? NobuildFilesRegex { get; set; }

        /// <summary>
        /// If set to `true`, the application version will not be deleted.
        /// </summary>
        [Input("noopOnDestroy")]
        public Input<bool>? NoopOnDestroy { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.
        /// Structure is documented below.
        /// </summary>
        [Input("readinessCheck")]
        public Input<Inputs.FlexibleAppVersionReadinessCheckGetArgs>? ReadinessCheck { get; set; }

        /// <summary>
        /// Machine resources for a version.
        /// Structure is documented below.
        /// </summary>
        [Input("resources")]
        public Input<Inputs.FlexibleAppVersionResourcesGetArgs>? Resources { get; set; }

        /// <summary>
        /// Desired runtime. Example python27.
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        /// <summary>
        /// The version of the API in the given runtime environment.
        /// Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
        /// </summary>
        [Input("runtimeApiVersion")]
        public Input<string>? RuntimeApiVersion { get; set; }

        /// <summary>
        /// The channel of the runtime to use. Only available for some runtimes.
        /// </summary>
        [Input("runtimeChannel")]
        public Input<string>? RuntimeChannel { get; set; }

        /// <summary>
        /// The path or name of the app's main executable.
        /// </summary>
        [Input("runtimeMainExecutablePath")]
        public Input<string>? RuntimeMainExecutablePath { get; set; }

        /// <summary>
        /// AppEngine service resource. Can contain numbers, letters, and hyphens.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.
        /// Default value is `SERVING`.
        /// Possible values are `SERVING` and `STOPPED`.
        /// </summary>
        [Input("servingStatus")]
        public Input<string>? ServingStatus { get; set; }

        /// <summary>
        /// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens.
        /// Reserved names,"default", "latest", and any name with the prefix "ah-".
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// Structure is documented below.
        /// </summary>
        [Input("vpcAccessConnector")]
        public Input<Inputs.FlexibleAppVersionVpcAccessConnectorGetArgs>? VpcAccessConnector { get; set; }

        public FlexibleAppVersionState()
        {
        }
    }
}
