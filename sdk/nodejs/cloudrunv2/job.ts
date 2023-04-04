// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Cloud Run Job resource that references a container image which is run to completion.
 *
 * To get more information about Job, see:
 *
 * * [API documentation](https://cloud.google.com/run/docs/reference/rest/v2/projects.locations.jobs)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/run/docs/)
 *
 * ## Example Usage
 * ### Cloudrunv2 Job Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.cloudrunv2.Job("default", {
 *     location: "us-central1",
 *     launchStage: "BETA",
 *     template: {
 *         template: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *             }],
 *         },
 *     },
 * });
 * ```
 * ### Cloudrunv2 Job Sql
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const secret = new gcp.secretmanager.Secret("secret", {
 *     secretId: "secret",
 *     replication: {
 *         automatic: true,
 *     },
 * });
 * const instance = new gcp.sql.DatabaseInstance("instance", {
 *     region: "us-central1",
 *     databaseVersion: "MYSQL_5_7",
 *     settings: {
 *         tier: "db-f1-micro",
 *     },
 *     deletionProtection: true,
 * });
 * const _default = new gcp.cloudrunv2.Job("default", {
 *     location: "us-central1",
 *     launchStage: "BETA",
 *     template: {
 *         template: {
 *             volumes: [{
 *                 name: "cloudsql",
 *                 cloudSqlInstance: {
 *                     instances: [instance.connectionName],
 *                 },
 *             }],
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *                 envs: [
 *                     {
 *                         name: "FOO",
 *                         value: "bar",
 *                     },
 *                     {
 *                         name: "latestdclsecret",
 *                         valueSource: {
 *                             secretKeyRef: {
 *                                 secret: secret.secretId,
 *                                 version: "1",
 *                             },
 *                         },
 *                     },
 *                 ],
 *                 volumeMounts: [{
 *                     name: "cloudsql",
 *                     mountPath: "/cloudsql",
 *                 }],
 *             }],
 *         },
 *     },
 * });
 * const project = gcp.organizations.getProject({});
 * const secret_version_data = new gcp.secretmanager.SecretVersion("secret-version-data", {
 *     secret: secret.name,
 *     secretData: "secret-data",
 * });
 * const secret_access = new gcp.secretmanager.SecretIamMember("secret-access", {
 *     secretId: secret.id,
 *     role: "roles/secretmanager.secretAccessor",
 *     member: project.then(project => `serviceAccount:${project.number}-compute@developer.gserviceaccount.com`),
 * }, {
 *     dependsOn: [secret],
 * });
 * ```
 * ### Cloudrunv2 Job Vpcaccess
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const customTestNetwork = new gcp.compute.Network("customTestNetwork", {autoCreateSubnetworks: false});
 * const customTestSubnetwork = new gcp.compute.Subnetwork("customTestSubnetwork", {
 *     ipCidrRange: "10.2.0.0/28",
 *     region: "us-central1",
 *     network: customTestNetwork.id,
 * });
 * const connector = new gcp.vpcaccess.Connector("connector", {
 *     subnet: {
 *         name: customTestSubnetwork.name,
 *     },
 *     machineType: "e2-standard-4",
 *     minInstances: 2,
 *     maxInstances: 3,
 *     region: "us-central1",
 * });
 * const _default = new gcp.cloudrunv2.Job("default", {
 *     location: "us-central1",
 *     launchStage: "BETA",
 *     template: {
 *         template: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *             }],
 *             vpcAccess: {
 *                 connector: connector.id,
 *                 egress: "ALL_TRAFFIC",
 *             },
 *         },
 *     },
 * });
 * ```
 * ### Cloudrunv2 Job Secret
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const secret = new gcp.secretmanager.Secret("secret", {
 *     secretId: "secret",
 *     replication: {
 *         automatic: true,
 *     },
 * });
 * const _default = new gcp.cloudrunv2.Job("default", {
 *     location: "us-central1",
 *     launchStage: "BETA",
 *     template: {
 *         template: {
 *             volumes: [{
 *                 name: "a-volume",
 *                 secret: {
 *                     secret: secret.secretId,
 *                     defaultMode: 292,
 *                     items: [{
 *                         version: "1",
 *                         path: "my-secret",
 *                         mode: 256,
 *                     }],
 *                 },
 *             }],
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *                 volumeMounts: [{
 *                     name: "a-volume",
 *                     mountPath: "/secrets",
 *                 }],
 *             }],
 *         },
 *     },
 * });
 * const project = gcp.organizations.getProject({});
 * const secret_version_data = new gcp.secretmanager.SecretVersion("secret-version-data", {
 *     secret: secret.name,
 *     secretData: "secret-data",
 * });
 * const secret_access = new gcp.secretmanager.SecretIamMember("secret-access", {
 *     secretId: secret.id,
 *     role: "roles/secretmanager.secretAccessor",
 *     member: project.then(project => `serviceAccount:${project.number}-compute@developer.gserviceaccount.com`),
 * }, {
 *     dependsOn: [secret],
 * });
 * ```
 *
 * ## Import
 *
 * Job can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:cloudrunv2/job:Job default projects/{{project}}/locations/{{location}}/jobs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudrunv2/job:Job default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudrunv2/job:Job default {{location}}/{{name}}
 * ```
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudrunv2/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Settings for the Binary Authorization feature.
     * Structure is documented below.
     */
    public readonly binaryAuthorization!: pulumi.Output<outputs.cloudrunv2.JobBinaryAuthorization | undefined>;
    /**
     * Arbitrary identifier for the API client.
     */
    public readonly client!: pulumi.Output<string | undefined>;
    /**
     * Arbitrary version identifier for the API client.
     */
    public readonly clientVersion!: pulumi.Output<string | undefined>;
    /**
     * The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on `reconciliation` process in Cloud Run.
     * Structure is documented below.
     */
    public /*out*/ readonly conditions!: pulumi.Output<outputs.cloudrunv2.JobCondition[]>;
    /**
     * A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Number of executions created for this job.
     */
    public /*out*/ readonly executionCount!: pulumi.Output<number>;
    /**
     * A number that monotonically increases every time the user modifies the desired state.
     */
    public /*out*/ readonly generation!: pulumi.Output<string>;
    /**
     * KRM-style labels for the resource.
     * (Optional)
     * KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Name of the last created execution.
     * Structure is documented below.
     */
    public /*out*/ readonly latestCreatedExecutions!: pulumi.Output<outputs.cloudrunv2.JobLatestCreatedExecution[]>;
    /**
     * The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
     * Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
     */
    public readonly launchStage!: pulumi.Output<string>;
    /**
     * The location of the cloud run job
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Name of the Job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.
     */
    public /*out*/ readonly observedGeneration!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Returns true if the Job is currently being acted upon by the system to bring it into the desired state.
     * When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.
     * If reconciliation succeeded, the following fields will match: observedGeneration and generation, latestSucceededExecution and latestCreatedExecution.
     * If reconciliation failed, observedGeneration and latestSucceededExecution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * The template used to create executions for this Job.
     * Structure is documented below.
     */
    public readonly template!: pulumi.Output<outputs.cloudrunv2.JobTemplate>;
    /**
     * The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state
     * Structure is documented below.
     */
    public /*out*/ readonly terminalConditions!: pulumi.Output<outputs.cloudrunv2.JobTerminalCondition[]>;
    /**
     * Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobState | undefined;
            resourceInputs["binaryAuthorization"] = state ? state.binaryAuthorization : undefined;
            resourceInputs["client"] = state ? state.client : undefined;
            resourceInputs["clientVersion"] = state ? state.clientVersion : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["executionCount"] = state ? state.executionCount : undefined;
            resourceInputs["generation"] = state ? state.generation : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["latestCreatedExecutions"] = state ? state.latestCreatedExecutions : undefined;
            resourceInputs["launchStage"] = state ? state.launchStage : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["observedGeneration"] = state ? state.observedGeneration : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["reconciling"] = state ? state.reconciling : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
            resourceInputs["terminalConditions"] = state ? state.terminalConditions : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.template === undefined) && !opts.urn) {
                throw new Error("Missing required property 'template'");
            }
            resourceInputs["binaryAuthorization"] = args ? args.binaryAuthorization : undefined;
            resourceInputs["client"] = args ? args.client : undefined;
            resourceInputs["clientVersion"] = args ? args.clientVersion : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["launchStage"] = args ? args.launchStage : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["conditions"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["executionCount"] = undefined /*out*/;
            resourceInputs["generation"] = undefined /*out*/;
            resourceInputs["latestCreatedExecutions"] = undefined /*out*/;
            resourceInputs["observedGeneration"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["terminalConditions"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Job.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * Settings for the Binary Authorization feature.
     * Structure is documented below.
     */
    binaryAuthorization?: pulumi.Input<inputs.cloudrunv2.JobBinaryAuthorization>;
    /**
     * Arbitrary identifier for the API client.
     */
    client?: pulumi.Input<string>;
    /**
     * Arbitrary version identifier for the API client.
     */
    clientVersion?: pulumi.Input<string>;
    /**
     * The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in reconciling for additional information on `reconciliation` process in Cloud Run.
     * Structure is documented below.
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.cloudrunv2.JobCondition>[]>;
    /**
     * A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
     */
    etag?: pulumi.Input<string>;
    /**
     * Number of executions created for this job.
     */
    executionCount?: pulumi.Input<number>;
    /**
     * A number that monotonically increases every time the user modifies the desired state.
     */
    generation?: pulumi.Input<string>;
    /**
     * KRM-style labels for the resource.
     * (Optional)
     * KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the last created execution.
     * Structure is documented below.
     */
    latestCreatedExecutions?: pulumi.Input<pulumi.Input<inputs.cloudrunv2.JobLatestCreatedExecution>[]>;
    /**
     * The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
     * Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
     */
    launchStage?: pulumi.Input<string>;
    /**
     * The location of the cloud run job
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the Job.
     */
    name?: pulumi.Input<string>;
    /**
     * The generation of this Job. See comments in reconciling for additional information on reconciliation process in Cloud Run.
     */
    observedGeneration?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Returns true if the Job is currently being acted upon by the system to bring it into the desired state.
     * When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.
     * If reconciliation succeeded, the following fields will match: observedGeneration and generation, latestSucceededExecution and latestCreatedExecution.
     * If reconciliation failed, observedGeneration and latestSucceededExecution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions
     */
    reconciling?: pulumi.Input<boolean>;
    /**
     * The template used to create executions for this Job.
     * Structure is documented below.
     */
    template?: pulumi.Input<inputs.cloudrunv2.JobTemplate>;
    /**
     * The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state
     * Structure is documented below.
     */
    terminalConditions?: pulumi.Input<pulumi.Input<inputs.cloudrunv2.JobTerminalCondition>[]>;
    /**
     * Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    uid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * Settings for the Binary Authorization feature.
     * Structure is documented below.
     */
    binaryAuthorization?: pulumi.Input<inputs.cloudrunv2.JobBinaryAuthorization>;
    /**
     * Arbitrary identifier for the API client.
     */
    client?: pulumi.Input<string>;
    /**
     * Arbitrary version identifier for the API client.
     */
    clientVersion?: pulumi.Input<string>;
    /**
     * KRM-style labels for the resource.
     * (Optional)
     * KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The launch stage as defined by Google Cloud Platform Launch Stages. Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed.
     * Possible values are `UNIMPLEMENTED`, `PRELAUNCH`, `EARLY_ACCESS`, `ALPHA`, `BETA`, `GA`, and `DEPRECATED`.
     */
    launchStage?: pulumi.Input<string>;
    /**
     * The location of the cloud run job
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the Job.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The template used to create executions for this Job.
     * Structure is documented below.
     */
    template: pulumi.Input<inputs.cloudrunv2.JobTemplate>;
}
