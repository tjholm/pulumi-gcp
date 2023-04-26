// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring
{
    public static class GetIstioCanonicalService
    {
        /// <summary>
        /// A Monitoring Service is the root resource under which operational aspects of a
        /// generic service are accessible. A service is some discrete, autonomous, and
        /// network-accessible unit, designed to solve an individual concern
        /// 
        /// A monitoring Istio Canonical Service is automatically created by GCP to monitor
        /// Istio Canonical Services.
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
        /// ### Monitoring Istio Canonical Service
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
        ///     var @default = Gcp.Monitoring.GetIstioCanonicalService.Invoke(new()
        ///     {
        ///         CanonicalService = "prometheus",
        ///         CanonicalServiceNamespace = "istio-system",
        ///         MeshUid = "proj-573164786102",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIstioCanonicalServiceResult> InvokeAsync(GetIstioCanonicalServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIstioCanonicalServiceResult>("gcp:monitoring/getIstioCanonicalService:getIstioCanonicalService", args ?? new GetIstioCanonicalServiceArgs(), options.WithDefaults());

        /// <summary>
        /// A Monitoring Service is the root resource under which operational aspects of a
        /// generic service are accessible. A service is some discrete, autonomous, and
        /// network-accessible unit, designed to solve an individual concern
        /// 
        /// A monitoring Istio Canonical Service is automatically created by GCP to monitor
        /// Istio Canonical Services.
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
        /// ### Monitoring Istio Canonical Service
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
        ///     var @default = Gcp.Monitoring.GetIstioCanonicalService.Invoke(new()
        ///     {
        ///         CanonicalService = "prometheus",
        ///         CanonicalServiceNamespace = "istio-system",
        ///         MeshUid = "proj-573164786102",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIstioCanonicalServiceResult> Invoke(GetIstioCanonicalServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIstioCanonicalServiceResult>("gcp:monitoring/getIstioCanonicalService:getIstioCanonicalService", args ?? new GetIstioCanonicalServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIstioCanonicalServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the canonical service underlying this service.
        /// Corresponds to the destination_canonical_service_name metric label in label in Istio metrics.
        /// </summary>
        [Input("canonicalService", required: true)]
        public string CanonicalService { get; set; } = null!;

        /// <summary>
        /// The namespace of the canonical service underlying this service.
        /// Corresponds to the destination_canonical_service_namespace metric label in Istio metrics.
        /// </summary>
        [Input("canonicalServiceNamespace", required: true)]
        public string CanonicalServiceNamespace { get; set; } = null!;

        /// <summary>
        /// Identifier for the mesh in which this Istio service is defined.
        /// Corresponds to the meshUid metric label in Istio metrics.
        /// </summary>
        [Input("meshUid", required: true)]
        public string MeshUid { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetIstioCanonicalServiceArgs()
        {
        }
        public static new GetIstioCanonicalServiceArgs Empty => new GetIstioCanonicalServiceArgs();
    }

    public sealed class GetIstioCanonicalServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the canonical service underlying this service.
        /// Corresponds to the destination_canonical_service_name metric label in label in Istio metrics.
        /// </summary>
        [Input("canonicalService", required: true)]
        public Input<string> CanonicalService { get; set; } = null!;

        /// <summary>
        /// The namespace of the canonical service underlying this service.
        /// Corresponds to the destination_canonical_service_namespace metric label in Istio metrics.
        /// </summary>
        [Input("canonicalServiceNamespace", required: true)]
        public Input<string> CanonicalServiceNamespace { get; set; } = null!;

        /// <summary>
        /// Identifier for the mesh in which this Istio service is defined.
        /// Corresponds to the meshUid metric label in Istio metrics.
        /// </summary>
        [Input("meshUid", required: true)]
        public Input<string> MeshUid { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetIstioCanonicalServiceInvokeArgs()
        {
        }
        public static new GetIstioCanonicalServiceInvokeArgs Empty => new GetIstioCanonicalServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIstioCanonicalServiceResult
    {
        public readonly string CanonicalService;
        public readonly string CanonicalServiceNamespace;
        /// <summary>
        /// Name used for UI elements listing this (Monitoring) Service.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string MeshUid;
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
        public readonly ImmutableArray<Outputs.GetIstioCanonicalServiceTelemetryResult> Telemetries;
        public readonly ImmutableDictionary<string, string> UserLabels;

        [OutputConstructor]
        private GetIstioCanonicalServiceResult(
            string canonicalService,

            string canonicalServiceNamespace,

            string displayName,

            string id,

            string meshUid,

            string name,

            string? project,

            string serviceId,

            ImmutableArray<Outputs.GetIstioCanonicalServiceTelemetryResult> telemetries,

            ImmutableDictionary<string, string> userLabels)
        {
            CanonicalService = canonicalService;
            CanonicalServiceNamespace = canonicalServiceNamespace;
            DisplayName = displayName;
            Id = id;
            MeshUid = meshUid;
            Name = name;
            Project = project;
            ServiceId = serviceId;
            Telemetries = telemetries;
            UserLabels = userLabels;
        }
    }
}
