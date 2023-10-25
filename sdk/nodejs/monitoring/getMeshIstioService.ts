// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Monitoring Service is the root resource under which operational aspects of a
 * generic service are accessible. A service is some discrete, autonomous, and
 * network-accessible unit, designed to solve an individual concern
 *
 * An Mesh Istio monitoring service is automatically created by GCP to monitor
 * Mesh Istio services.
 *
 * To get more information about Service, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
 * * How-to Guides
 *     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 *
 * ## Example Usage
 * ### Monitoring Mesh Istio Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.monitoring.getMeshIstioService({
 *     meshUid: "proj-573164786102",
 *     serviceName: "prometheus",
 *     serviceNamespace: "istio-system",
 * });
 * ```
 */
export function getMeshIstioService(args: GetMeshIstioServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetMeshIstioServiceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:monitoring/getMeshIstioService:getMeshIstioService", {
        "meshUid": args.meshUid,
        "project": args.project,
        "serviceName": args.serviceName,
        "serviceNamespace": args.serviceNamespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getMeshIstioService.
 */
export interface GetMeshIstioServiceArgs {
    /**
     * Identifier for the mesh in which this Istio service is defined.
     * Corresponds to the meshUid metric label in Istio metrics.
     */
    meshUid: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The name of the Istio service underlying this service.
     * Corresponds to the destinationServiceName metric label in Istio metrics.
     *
     * - - -
     *
     * Other optional fields include:
     */
    serviceName: string;
    /**
     * The namespace of the Istio service underlying this service.
     * Corresponds to the destinationServiceNamespace metric label in Istio metrics.
     */
    serviceNamespace: string;
}

/**
 * A collection of values returned by getMeshIstioService.
 */
export interface GetMeshIstioServiceResult {
    /**
     * Name used for UI elements listing this (Monitoring) Service.
     */
    readonly displayName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly meshUid: string;
    /**
     * The full REST resource name for this channel. The syntax is:
     * `projects/[PROJECT_ID]/services/[SERVICE_ID]`.
     */
    readonly name: string;
    readonly project?: string;
    readonly serviceId: string;
    readonly serviceName: string;
    readonly serviceNamespace: string;
    /**
     * Configuration for how to query telemetry on the Service. Structure is documented below.
     */
    readonly telemetries: outputs.monitoring.GetMeshIstioServiceTelemetry[];
    readonly userLabels: {[key: string]: string};
}
/**
 * A Monitoring Service is the root resource under which operational aspects of a
 * generic service are accessible. A service is some discrete, autonomous, and
 * network-accessible unit, designed to solve an individual concern
 *
 * An Mesh Istio monitoring service is automatically created by GCP to monitor
 * Mesh Istio services.
 *
 * To get more information about Service, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/services)
 * * How-to Guides
 *     * [Service Monitoring](https://cloud.google.com/monitoring/service-monitoring)
 *     * [Monitoring API Documentation](https://cloud.google.com/monitoring/api/v3/)
 *
 * ## Example Usage
 * ### Monitoring Mesh Istio Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const default = gcp.monitoring.getMeshIstioService({
 *     meshUid: "proj-573164786102",
 *     serviceName: "prometheus",
 *     serviceNamespace: "istio-system",
 * });
 * ```
 */
export function getMeshIstioServiceOutput(args: GetMeshIstioServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMeshIstioServiceResult> {
    return pulumi.output(args).apply((a: any) => getMeshIstioService(a, opts))
}

/**
 * A collection of arguments for invoking getMeshIstioService.
 */
export interface GetMeshIstioServiceOutputArgs {
    /**
     * Identifier for the mesh in which this Istio service is defined.
     * Corresponds to the meshUid metric label in Istio metrics.
     */
    meshUid: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the Istio service underlying this service.
     * Corresponds to the destinationServiceName metric label in Istio metrics.
     *
     * - - -
     *
     * Other optional fields include:
     */
    serviceName: pulumi.Input<string>;
    /**
     * The namespace of the Istio service underlying this service.
     * Corresponds to the destinationServiceNamespace metric label in Istio metrics.
     */
    serviceNamespace: pulumi.Input<string>;
}
