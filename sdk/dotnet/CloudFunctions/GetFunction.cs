// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudFunctions
{
    public static class GetFunction
    {
        /// <summary>
        /// Get information about a Google Cloud Function. For more information see
        /// the [official documentation](https://cloud.google.com/functions/docs/)
        /// and [API](https://cloud.google.com/functions/docs/apis).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_function = Output.Create(Gcp.CloudFunctions.GetFunction.InvokeAsync(new Gcp.CloudFunctions.GetFunctionArgs
        ///         {
        ///             Name = "function",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("gcp:cloudfunctions/getFunction:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// Get information about a Google Cloud Function. For more information see
        /// the [official documentation](https://cloud.google.com/functions/docs/)
        /// and [API](https://cloud.google.com/functions/docs/apis).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var my_function = Output.Create(Gcp.CloudFunctions.GetFunction.InvokeAsync(new Gcp.CloudFunctions.GetFunctionArgs
        ///         {
        ///             Name = "function",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("gcp:cloudfunctions/getFunction:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a Cloud Function.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetFunctionArgs()
        {
        }
    }

    public sealed class GetFunctionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a Cloud Function.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public GetFunctionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// Available memory (in MB) to the function.
        /// </summary>
        public readonly int AvailableMemoryMb;
        public readonly ImmutableDictionary<string, object> BuildEnvironmentVariables;
        /// <summary>
        /// Description of the function.
        /// </summary>
        public readonly string Description;
        public readonly string DockerRegistry;
        public readonly string DockerRepository;
        /// <summary>
        /// Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
        /// </summary>
        public readonly string EntryPoint;
        public readonly ImmutableDictionary<string, object> EnvironmentVariables;
        /// <summary>
        /// A source that fires events in response to a condition in another service. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionEventTriggerResult> EventTriggers;
        public readonly string HttpsTriggerSecurityLevel;
        /// <summary>
        /// If function is triggered by HTTP, trigger URL is set here.
        /// </summary>
        public readonly string HttpsTriggerUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Controls what traffic can reach the function.
        /// </summary>
        public readonly string IngressSettings;
        public readonly string KmsKeyName;
        /// <summary>
        /// A map of labels applied to this function.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Labels;
        /// <summary>
        /// The limit on the maximum number of function instances that may coexist at a given time.
        /// </summary>
        public readonly int MaxInstances;
        public readonly int MinInstances;
        /// <summary>
        /// The name of the Cloud Function.
        /// </summary>
        public readonly string Name;
        public readonly string? Project;
        public readonly string? Region;
        /// <summary>
        /// The runtime in which the function is running.
        /// </summary>
        public readonly string Runtime;
        public readonly ImmutableArray<Outputs.GetFunctionSecretEnvironmentVariableResult> SecretEnvironmentVariables;
        public readonly ImmutableArray<Outputs.GetFunctionSecretVolumeResult> SecretVolumes;
        /// <summary>
        /// The service account email to be assumed by the cloud function.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// The GCS bucket containing the zip archive which contains the function.
        /// </summary>
        public readonly string SourceArchiveBucket;
        /// <summary>
        /// The source archive object (file) in archive bucket.
        /// </summary>
        public readonly string SourceArchiveObject;
        /// <summary>
        /// The URL of the Cloud Source Repository that the function is deployed from. Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetFunctionSourceRepositoryResult> SourceRepositories;
        /// <summary>
        /// Function execution timeout (in seconds).
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// If function is triggered by HTTP, this boolean is set.
        /// </summary>
        public readonly bool TriggerHttp;
        /// <summary>
        /// The VPC Network Connector that this cloud function can connect to.
        /// </summary>
        public readonly string VpcConnector;
        /// <summary>
        /// The egress settings for the connector, controlling what traffic is diverted through it.
        /// </summary>
        public readonly string VpcConnectorEgressSettings;

        [OutputConstructor]
        private GetFunctionResult(
            int availableMemoryMb,

            ImmutableDictionary<string, object> buildEnvironmentVariables,

            string description,

            string dockerRegistry,

            string dockerRepository,

            string entryPoint,

            ImmutableDictionary<string, object> environmentVariables,

            ImmutableArray<Outputs.GetFunctionEventTriggerResult> eventTriggers,

            string httpsTriggerSecurityLevel,

            string httpsTriggerUrl,

            string id,

            string ingressSettings,

            string kmsKeyName,

            ImmutableDictionary<string, object> labels,

            int maxInstances,

            int minInstances,

            string name,

            string? project,

            string? region,

            string runtime,

            ImmutableArray<Outputs.GetFunctionSecretEnvironmentVariableResult> secretEnvironmentVariables,

            ImmutableArray<Outputs.GetFunctionSecretVolumeResult> secretVolumes,

            string serviceAccountEmail,

            string sourceArchiveBucket,

            string sourceArchiveObject,

            ImmutableArray<Outputs.GetFunctionSourceRepositoryResult> sourceRepositories,

            int timeout,

            bool triggerHttp,

            string vpcConnector,

            string vpcConnectorEgressSettings)
        {
            AvailableMemoryMb = availableMemoryMb;
            BuildEnvironmentVariables = buildEnvironmentVariables;
            Description = description;
            DockerRegistry = dockerRegistry;
            DockerRepository = dockerRepository;
            EntryPoint = entryPoint;
            EnvironmentVariables = environmentVariables;
            EventTriggers = eventTriggers;
            HttpsTriggerSecurityLevel = httpsTriggerSecurityLevel;
            HttpsTriggerUrl = httpsTriggerUrl;
            Id = id;
            IngressSettings = ingressSettings;
            KmsKeyName = kmsKeyName;
            Labels = labels;
            MaxInstances = maxInstances;
            MinInstances = minInstances;
            Name = name;
            Project = project;
            Region = region;
            Runtime = runtime;
            SecretEnvironmentVariables = secretEnvironmentVariables;
            SecretVolumes = secretVolumes;
            ServiceAccountEmail = serviceAccountEmail;
            SourceArchiveBucket = sourceArchiveBucket;
            SourceArchiveObject = sourceArchiveObject;
            SourceRepositories = sourceRepositories;
            Timeout = timeout;
            TriggerHttp = triggerHttp;
            VpcConnector = vpcConnector;
            VpcConnectorEgressSettings = vpcConnectorEgressSettings;
        }
    }
}
