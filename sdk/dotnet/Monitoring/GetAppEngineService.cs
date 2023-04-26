// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    public static class GetAppEngineService
    {
        /// <summary>
        /// A Monitoring Service is the root resource under which operational aspects of a
        /// generic service are accessible. A service is some discrete, autonomous, and
        /// network-accessible unit, designed to solve an individual concern
        /// 
        /// An App Engine monitoring service is automatically created by GCP to monitor
        /// App Engine services.
        /// 
        /// 
        /// To get more information about Service, see:
        /// 
        /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
        /// * How-to Guides
        ///     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
        ///     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Monitoring App Engine Service
        /// 
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
        ///     var @object = new Gcp.Storage.BucketObject("object", new()
        ///     {
        ///         Bucket = bucket.Name,
        ///         Source = new FileAsset("./test-fixtures/appengine/hello-world.zip"),
        ///     });
        /// 
        ///     var myapp = new Gcp.AppEngine.StandardAppVersion("myapp", new()
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
        ///         DeleteServiceOnDestroy = false,
        ///     });
        /// 
        ///     var srv = Gcp.Monitoring.GetAppEngineService.Invoke(new()
        ///     {
        ///         ModuleId = myapp.Service,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAppEngineServiceResult> InvokeAsync(GetAppEngineServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppEngineServiceResult>("gcp:monitoring/getAppEngineService:getAppEngineService", args ?? new GetAppEngineServiceArgs(), options.WithDefaults());

        /// <summary>
        /// A Monitoring Service is the root resource under which operational aspects of a
        /// generic service are accessible. A service is some discrete, autonomous, and
        /// network-accessible unit, designed to solve an individual concern
        /// 
        /// An App Engine monitoring service is automatically created by GCP to monitor
        /// App Engine services.
        /// 
        /// 
        /// To get more information about Service, see:
        /// 
        /// * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
        /// * How-to Guides
        ///     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
        ///     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Monitoring App Engine Service
        /// 
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
        ///     var @object = new Gcp.Storage.BucketObject("object", new()
        ///     {
        ///         Bucket = bucket.Name,
        ///         Source = new FileAsset("./test-fixtures/appengine/hello-world.zip"),
        ///     });
        /// 
        ///     var myapp = new Gcp.AppEngine.StandardAppVersion("myapp", new()
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
        ///         DeleteServiceOnDestroy = false,
        ///     });
        /// 
        ///     var srv = Gcp.Monitoring.GetAppEngineService.Invoke(new()
        ///     {
        ///         ModuleId = myapp.Service,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAppEngineServiceResult> Invoke(GetAppEngineServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppEngineServiceResult>("gcp:monitoring/getAppEngineService:getAppEngineService", args ?? new GetAppEngineServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppEngineServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the App Engine module underlying this
        /// service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
        /// </summary>
        [Input("moduleId", required: true)]
        public string ModuleId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetAppEngineServiceArgs()
        {
        }
        public static new GetAppEngineServiceArgs Empty => new GetAppEngineServiceArgs();
    }

    public sealed class GetAppEngineServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the App Engine module underlying this
        /// service. Corresponds to the moduleId resource label in the [gae_app](https://cloud.google.com/monitoring/api/resources#tag_gae_app) monitored resource, or the service/module name.
        /// </summary>
        [Input("moduleId", required: true)]
        public Input<string> ModuleId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAppEngineServiceInvokeArgs()
        {
        }
        public static new GetAppEngineServiceInvokeArgs Empty => new GetAppEngineServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppEngineServiceResult
    {
        /// <summary>
        /// Name used for UI elements listing this (Monitoring) Service.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ModuleId;
        /// <summary>
        /// The full REST resource name for this channel. The syntax is:
        /// `projects/[PROJECT_ID]/services/[SERVICE_ID]`.
        /// </summary>
        public readonly string Name;
        public readonly string? Project;
        public readonly string ServiceId;
        /// <summary>
        /// Configuration for how to query telemetry on the Service. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAppEngineServiceTelemetryResult> Telemetries;
        public readonly ImmutableDictionary<string, string> UserLabels;

        [OutputConstructor]
        private GetAppEngineServiceResult(
            string displayName,

            string id,

            string moduleId,

            string name,

            string? project,

            string serviceId,

            ImmutableArray<Outputs.GetAppEngineServiceTelemetryResult> telemetries,

            ImmutableDictionary<string, string> userLabels)
        {
            DisplayName = displayName;
            Id = id;
            ModuleId = moduleId;
            Name = name;
            Project = project;
            ServiceId = serviceId;
            Telemetries = telemetries;
            UserLabels = userLabels;
        }
    }
}
