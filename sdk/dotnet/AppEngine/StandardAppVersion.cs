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
    /// Standard App Version resource to create a new version of standard GAE Application.
    /// Learn about the differences between the standard environment and the flexible environment
    /// at https://cloud.google.com/appengine/docs/the-appengine-environments.
    /// Currently supporting Zip and File Containers.
    /// 
    /// To get more information about StandardAppVersion, see:
    /// 
    /// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/appengine/docs/standard)
    /// 
    /// ## Example Usage
    /// ### App Engine Standard App Version
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var customServiceAccount = new Gcp.ServiceAccount.Account("customServiceAccount", new()
    ///     {
    ///         AccountId = "my-account",
    ///         DisplayName = "Custom Service Account",
    ///     });
    /// 
    ///     var gaeApi = new Gcp.Projects.IAMMember("gaeApi", new()
    ///     {
    ///         Project = customServiceAccount.Project,
    ///         Role = "roles/compute.networkUser",
    ///         Member = customServiceAccount.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///     });
    /// 
    ///     var storageViewer = new Gcp.Projects.IAMMember("storageViewer", new()
    ///     {
    ///         Project = customServiceAccount.Project,
    ///         Role = "roles/storage.objectViewer",
    ///         Member = customServiceAccount.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///     });
    /// 
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "US",
    ///     });
    /// 
    ///     var @object = new Gcp.Storage.BucketObject("object", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Source = new FileAsset("./test-fixtures/hello-world.zip"),
    ///     });
    /// 
    ///     var myappV1 = new Gcp.AppEngine.StandardAppVersion("myappV1", new()
    ///     {
    ///         VersionId = "v1",
    ///         Service = "myapp",
    ///         Runtime = "nodejs10",
    ///         Entrypoint = new Gcp.AppEngine.Inputs.StandardAppVersionEntrypointArgs
    ///         {
    ///             Shell = "node ./app.js",
    ///         },
    ///         Deployment = new Gcp.AppEngine.Inputs.StandardAppVersionDeploymentArgs
    ///         {
    ///             Zip = new Gcp.AppEngine.Inputs.StandardAppVersionDeploymentZipArgs
    ///             {
    ///                 SourceUrl = Output.Tuple(bucket.Name, @object.Name).Apply(values =&gt;
    ///                 {
    ///                     var bucketName = values.Item1;
    ///                     var objectName = values.Item2;
    ///                     return $"https://storage.googleapis.com/{bucketName}/{objectName}";
    ///                 }),
    ///             },
    ///         },
    ///         EnvVariables = 
    ///         {
    ///             { "port", "8080" },
    ///         },
    ///         AutomaticScaling = new Gcp.AppEngine.Inputs.StandardAppVersionAutomaticScalingArgs
    ///         {
    ///             MaxConcurrentRequests = 10,
    ///             MinIdleInstances = 1,
    ///             MaxIdleInstances = 3,
    ///             MinPendingLatency = "1s",
    ///             MaxPendingLatency = "5s",
    ///             StandardSchedulerSettings = new Gcp.AppEngine.Inputs.StandardAppVersionAutomaticScalingStandardSchedulerSettingsArgs
    ///             {
    ///                 TargetCpuUtilization = 0.5,
    ///                 TargetThroughputUtilization = 0.75,
    ///                 MinInstances = 2,
    ///                 MaxInstances = 10,
    ///             },
    ///         },
    ///         DeleteServiceOnDestroy = true,
    ///         ServiceAccount = customServiceAccount.Email,
    ///     });
    /// 
    ///     var myappV2 = new Gcp.AppEngine.StandardAppVersion("myappV2", new()
    ///     {
    ///         VersionId = "v2",
    ///         Service = "myapp",
    ///         Runtime = "nodejs10",
    ///         AppEngineApis = true,
    ///         Entrypoint = new Gcp.AppEngine.Inputs.StandardAppVersionEntrypointArgs
    ///         {
    ///             Shell = "node ./app.js",
    ///         },
    ///         Deployment = new Gcp.AppEngine.Inputs.StandardAppVersionDeploymentArgs
    ///         {
    ///             Zip = new Gcp.AppEngine.Inputs.StandardAppVersionDeploymentZipArgs
    ///             {
    ///                 SourceUrl = Output.Tuple(bucket.Name, @object.Name).Apply(values =&gt;
    ///                 {
    ///                     var bucketName = values.Item1;
    ///                     var objectName = values.Item2;
    ///                     return $"https://storage.googleapis.com/{bucketName}/{objectName}";
    ///                 }),
    ///             },
    ///         },
    ///         EnvVariables = 
    ///         {
    ///             { "port", "8080" },
    ///         },
    ///         BasicScaling = new Gcp.AppEngine.Inputs.StandardAppVersionBasicScalingArgs
    ///         {
    ///             MaxInstances = 5,
    ///         },
    ///         NoopOnDestroy = true,
    ///         ServiceAccount = customServiceAccount.Email,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// StandardAppVersion can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/standardAppVersion:StandardAppVersion default apps/{{project}}/services/{{service}}/versions/{{version_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/standardAppVersion:StandardAppVersion default {{project}}/{{service}}/{{version_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:appengine/standardAppVersion:StandardAppVersion default {{service}}/{{version_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:appengine/standardAppVersion:StandardAppVersion")]
    public partial class StandardAppVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allows App Engine second generation runtimes to access the legacy bundled services.
        /// </summary>
        [Output("appEngineApis")]
        public Output<bool?> AppEngineApis { get; private set; } = null!;

        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics.
        /// Structure is documented below.
        /// </summary>
        [Output("automaticScaling")]
        public Output<Outputs.StandardAppVersionAutomaticScaling?> AutomaticScaling { get; private set; } = null!;

        /// <summary>
        /// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
        /// Structure is documented below.
        /// </summary>
        [Output("basicScaling")]
        public Output<Outputs.StandardAppVersionBasicScaling?> BasicScaling { get; private set; } = null!;

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
        public Output<Outputs.StandardAppVersionDeployment> Deployment { get; private set; } = null!;

        /// <summary>
        /// The entrypoint for the application.
        /// Structure is documented below.
        /// </summary>
        [Output("entrypoint")]
        public Output<Outputs.StandardAppVersionEntrypoint> Entrypoint { get; private set; } = null!;

        /// <summary>
        /// Environment variables available to the application.
        /// </summary>
        [Output("envVariables")]
        public Output<ImmutableDictionary<string, string>?> EnvVariables { get; private set; } = null!;

        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests.
        /// The first matching URL handles the request and other request handlers are not attempted.
        /// Structure is documented below.
        /// </summary>
        [Output("handlers")]
        public Output<ImmutableArray<Outputs.StandardAppVersionHandler>> Handlers { get; private set; } = null!;

        /// <summary>
        /// A list of the types of messages that this application is able to receive.
        /// Each value may be one of: `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, `INBOUND_SERVICE_WARMUP`.
        /// </summary>
        [Output("inboundServices")]
        public Output<ImmutableArray<string>> InboundServices { get; private set; } = null!;

        /// <summary>
        /// Instance class that is used to run this version. Valid values are
        /// AutomaticScaling: F1, F2, F4, F4_1G
        /// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
        /// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        /// </summary>
        [Output("instanceClass")]
        public Output<string> InstanceClass { get; private set; } = null!;

        /// <summary>
        /// Configuration for third-party Python runtime libraries that are required by the application.
        /// Structure is documented below.
        /// </summary>
        [Output("libraries")]
        public Output<ImmutableArray<Outputs.StandardAppVersionLibrary>> Libraries { get; private set; } = null!;

        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        /// Structure is documented below.
        /// </summary>
        [Output("manualScaling")]
        public Output<Outputs.StandardAppVersionManualScaling?> ManualScaling { get; private set; } = null!;

        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// Desired runtime. Example python27.
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// The version of the API in the given runtime environment.
        /// Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/&lt;language&gt;/config/appref`\
        /// Substitute `&lt;language&gt;` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
        /// </summary>
        [Output("runtimeApiVersion")]
        public Output<string?> RuntimeApiVersion { get; private set; } = null!;

        /// <summary>
        /// AppEngine service resource
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Whether multiple requests can be dispatched to this version at once.
        /// </summary>
        [Output("threadsafe")]
        public Output<bool?> Threadsafe { get; private set; } = null!;

        /// <summary>
        /// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
        /// </summary>
        [Output("versionId")]
        public Output<string?> VersionId { get; private set; } = null!;

        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// Structure is documented below.
        /// </summary>
        [Output("vpcAccessConnector")]
        public Output<Outputs.StandardAppVersionVpcAccessConnector?> VpcAccessConnector { get; private set; } = null!;


        /// <summary>
        /// Create a StandardAppVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StandardAppVersion(string name, StandardAppVersionArgs args, CustomResourceOptions? options = null)
            : base("gcp:appengine/standardAppVersion:StandardAppVersion", name, args ?? new StandardAppVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StandardAppVersion(string name, Input<string> id, StandardAppVersionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:appengine/standardAppVersion:StandardAppVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StandardAppVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StandardAppVersion Get(string name, Input<string> id, StandardAppVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new StandardAppVersion(name, id, state, options);
        }
    }

    public sealed class StandardAppVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows App Engine second generation runtimes to access the legacy bundled services.
        /// </summary>
        [Input("appEngineApis")]
        public Input<bool>? AppEngineApis { get; set; }

        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics.
        /// Structure is documented below.
        /// </summary>
        [Input("automaticScaling")]
        public Input<Inputs.StandardAppVersionAutomaticScalingArgs>? AutomaticScaling { get; set; }

        /// <summary>
        /// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
        /// Structure is documented below.
        /// </summary>
        [Input("basicScaling")]
        public Input<Inputs.StandardAppVersionBasicScalingArgs>? BasicScaling { get; set; }

        /// <summary>
        /// If set to `true`, the service will be deleted if it is the last version.
        /// </summary>
        [Input("deleteServiceOnDestroy")]
        public Input<bool>? DeleteServiceOnDestroy { get; set; }

        /// <summary>
        /// Code and application artifacts that make up this version.
        /// Structure is documented below.
        /// </summary>
        [Input("deployment", required: true)]
        public Input<Inputs.StandardAppVersionDeploymentArgs> Deployment { get; set; } = null!;

        /// <summary>
        /// The entrypoint for the application.
        /// Structure is documented below.
        /// </summary>
        [Input("entrypoint", required: true)]
        public Input<Inputs.StandardAppVersionEntrypointArgs> Entrypoint { get; set; } = null!;

        [Input("envVariables")]
        private InputMap<string>? _envVariables;

        /// <summary>
        /// Environment variables available to the application.
        /// </summary>
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        [Input("handlers")]
        private InputList<Inputs.StandardAppVersionHandlerArgs>? _handlers;

        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests.
        /// The first matching URL handles the request and other request handlers are not attempted.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.StandardAppVersionHandlerArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.StandardAppVersionHandlerArgs>());
            set => _handlers = value;
        }

        [Input("inboundServices")]
        private InputList<string>? _inboundServices;

        /// <summary>
        /// A list of the types of messages that this application is able to receive.
        /// Each value may be one of: `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, `INBOUND_SERVICE_WARMUP`.
        /// </summary>
        public InputList<string> InboundServices
        {
            get => _inboundServices ?? (_inboundServices = new InputList<string>());
            set => _inboundServices = value;
        }

        /// <summary>
        /// Instance class that is used to run this version. Valid values are
        /// AutomaticScaling: F1, F2, F4, F4_1G
        /// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
        /// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        [Input("libraries")]
        private InputList<Inputs.StandardAppVersionLibraryArgs>? _libraries;

        /// <summary>
        /// Configuration for third-party Python runtime libraries that are required by the application.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.StandardAppVersionLibraryArgs> Libraries
        {
            get => _libraries ?? (_libraries = new InputList<Inputs.StandardAppVersionLibraryArgs>());
            set => _libraries = value;
        }

        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        /// Structure is documented below.
        /// </summary>
        [Input("manualScaling")]
        public Input<Inputs.StandardAppVersionManualScalingArgs>? ManualScaling { get; set; }

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
        /// Desired runtime. Example python27.
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        /// <summary>
        /// The version of the API in the given runtime environment.
        /// Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/&lt;language&gt;/config/appref`\
        /// Substitute `&lt;language&gt;` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
        /// </summary>
        [Input("runtimeApiVersion")]
        public Input<string>? RuntimeApiVersion { get; set; }

        /// <summary>
        /// AppEngine service resource
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// Whether multiple requests can be dispatched to this version at once.
        /// </summary>
        [Input("threadsafe")]
        public Input<bool>? Threadsafe { get; set; }

        /// <summary>
        /// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// Structure is documented below.
        /// </summary>
        [Input("vpcAccessConnector")]
        public Input<Inputs.StandardAppVersionVpcAccessConnectorArgs>? VpcAccessConnector { get; set; }

        public StandardAppVersionArgs()
        {
        }
        public static new StandardAppVersionArgs Empty => new StandardAppVersionArgs();
    }

    public sealed class StandardAppVersionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allows App Engine second generation runtimes to access the legacy bundled services.
        /// </summary>
        [Input("appEngineApis")]
        public Input<bool>? AppEngineApis { get; set; }

        /// <summary>
        /// Automatic scaling is based on request rate, response latencies, and other application metrics.
        /// Structure is documented below.
        /// </summary>
        [Input("automaticScaling")]
        public Input<Inputs.StandardAppVersionAutomaticScalingGetArgs>? AutomaticScaling { get; set; }

        /// <summary>
        /// Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
        /// Structure is documented below.
        /// </summary>
        [Input("basicScaling")]
        public Input<Inputs.StandardAppVersionBasicScalingGetArgs>? BasicScaling { get; set; }

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
        public Input<Inputs.StandardAppVersionDeploymentGetArgs>? Deployment { get; set; }

        /// <summary>
        /// The entrypoint for the application.
        /// Structure is documented below.
        /// </summary>
        [Input("entrypoint")]
        public Input<Inputs.StandardAppVersionEntrypointGetArgs>? Entrypoint { get; set; }

        [Input("envVariables")]
        private InputMap<string>? _envVariables;

        /// <summary>
        /// Environment variables available to the application.
        /// </summary>
        public InputMap<string> EnvVariables
        {
            get => _envVariables ?? (_envVariables = new InputMap<string>());
            set => _envVariables = value;
        }

        [Input("handlers")]
        private InputList<Inputs.StandardAppVersionHandlerGetArgs>? _handlers;

        /// <summary>
        /// An ordered list of URL-matching patterns that should be applied to incoming requests.
        /// The first matching URL handles the request and other request handlers are not attempted.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.StandardAppVersionHandlerGetArgs> Handlers
        {
            get => _handlers ?? (_handlers = new InputList<Inputs.StandardAppVersionHandlerGetArgs>());
            set => _handlers = value;
        }

        [Input("inboundServices")]
        private InputList<string>? _inboundServices;

        /// <summary>
        /// A list of the types of messages that this application is able to receive.
        /// Each value may be one of: `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, `INBOUND_SERVICE_WARMUP`.
        /// </summary>
        public InputList<string> InboundServices
        {
            get => _inboundServices ?? (_inboundServices = new InputList<string>());
            set => _inboundServices = value;
        }

        /// <summary>
        /// Instance class that is used to run this version. Valid values are
        /// AutomaticScaling: F1, F2, F4, F4_1G
        /// BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
        /// Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
        /// </summary>
        [Input("instanceClass")]
        public Input<string>? InstanceClass { get; set; }

        [Input("libraries")]
        private InputList<Inputs.StandardAppVersionLibraryGetArgs>? _libraries;

        /// <summary>
        /// Configuration for third-party Python runtime libraries that are required by the application.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.StandardAppVersionLibraryGetArgs> Libraries
        {
            get => _libraries ?? (_libraries = new InputList<Inputs.StandardAppVersionLibraryGetArgs>());
            set => _libraries = value;
        }

        /// <summary>
        /// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
        /// Structure is documented below.
        /// </summary>
        [Input("manualScaling")]
        public Input<Inputs.StandardAppVersionManualScalingGetArgs>? ManualScaling { get; set; }

        /// <summary>
        /// The identifier for this object. Format specified above.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// Desired runtime. Example python27.
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        /// <summary>
        /// The version of the API in the given runtime environment.
        /// Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/&lt;language&gt;/config/appref`\
        /// Substitute `&lt;language&gt;` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
        /// </summary>
        [Input("runtimeApiVersion")]
        public Input<string>? RuntimeApiVersion { get; set; }

        /// <summary>
        /// AppEngine service resource
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// Whether multiple requests can be dispatched to this version at once.
        /// </summary>
        [Input("threadsafe")]
        public Input<bool>? Threadsafe { get; set; }

        /// <summary>
        /// Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        /// <summary>
        /// Enables VPC connectivity for standard apps.
        /// Structure is documented below.
        /// </summary>
        [Input("vpcAccessConnector")]
        public Input<Inputs.StandardAppVersionVpcAccessConnectorGetArgs>? VpcAccessConnector { get; set; }

        public StandardAppVersionState()
        {
        }
        public static new StandardAppVersionState Empty => new StandardAppVersionState();
    }
}
