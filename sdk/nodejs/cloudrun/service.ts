// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Service acts as a top-level container that manages a set of Routes and
 * Configurations which implement a network service. Service exists to provide a
 * singular abstraction which can be access controlled, reasoned about, and
 * which encapsulates software lifecycle decisions such as rollout policy and
 * team resource ownership. Service acts only as an orchestrator of the
 * underlying Routes and Configurations (much as a kubernetes Deployment
 * orchestrates ReplicaSets).
 *
 * The Service's controller will track the statuses of its owned Configuration
 * and Route, reflecting their statuses and conditions as its own.
 *
 * See also:
 * https://github.com/knative/specs/blob/main/specs/serving/overview.md
 *
 * To get more information about Service, see:
 *
 * * [API documentation](https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/run/docs/)
 *
 * > **Warning:** `googleCloudrunService` creates a Managed Google Cloud Run Service. If you need to create
 * a Cloud Run Service on Anthos(GKE/VMWare) then you will need to create it using the kubernetes alpha provider.
 * Have a look at the Cloud Run Anthos example below.
 *
 * ## Example Usage
 * ### Cloud Run Service Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.cloudrun.Service("default", {
 *     location: "us-central1",
 *     template: {
 *         spec: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *             }],
 *         },
 *     },
 *     traffics: [{
 *         latestRevision: true,
 *         percent: 100,
 *     }],
 * });
 * ```
 * ### Cloud Run Service Sql
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.sql.DatabaseInstance("instance", {
 *     region: "us-east1",
 *     databaseVersion: "MYSQL_5_7",
 *     settings: {
 *         tier: "db-f1-micro",
 *     },
 *     deletionProtection: true,
 * });
 * const _default = new gcp.cloudrun.Service("default", {
 *     location: "us-central1",
 *     template: {
 *         spec: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *             }],
 *         },
 *         metadata: {
 *             annotations: {
 *                 "autoscaling.knative.dev/maxScale": "1000",
 *                 "run.googleapis.com/cloudsql-instances": instance.connectionName,
 *                 "run.googleapis.com/client-name": "terraform",
 *             },
 *         },
 *     },
 *     autogenerateRevisionName: true,
 * });
 * ```
 * ### Cloud Run Service Noauth
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.cloudrun.Service("default", {
 *     location: "us-central1",
 *     template: {
 *         spec: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *             }],
 *         },
 *     },
 * });
 * const noauthIAMPolicy = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/run.invoker",
 *         members: ["allUsers"],
 *     }],
 * });
 * const noauthIamPolicy = new gcp.cloudrun.IamPolicy("noauthIamPolicy", {
 *     location: _default.location,
 *     project: _default.project,
 *     service: _default.name,
 *     policyData: noauthIAMPolicy.then(noauthIAMPolicy => noauthIAMPolicy.policyData),
 * });
 * ```
 * ### Cloud Run Service Probes
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.cloudrun.Service("default", {
 *     location: "us-central1",
 *     metadata: {
 *         annotations: {
 *             "run.googleapis.com/launch-stage": "BETA",
 *         },
 *     },
 *     template: {
 *         spec: {
 *             containers: [{
 *                 image: "us-docker.pkg.dev/cloudrun/container/hello",
 *                 startupProbe: {
 *                     initialDelaySeconds: 0,
 *                     timeoutSeconds: 1,
 *                     periodSeconds: 3,
 *                     failureThreshold: 1,
 *                     tcpSocket: {
 *                         port: 8080,
 *                     },
 *                 },
 *                 livenessProbe: {
 *                     httpGet: {
 *                         path: "/",
 *                     },
 *                 },
 *             }],
 *         },
 *     },
 *     traffics: [{
 *         percent: 100,
 *         latestRevision: true,
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Service can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:cloudrun/service:Service default locations/{{location}}/namespaces/{{project}}/services/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudrun/service:Service default {{location}}/{{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudrun/service:Service default {{location}}/{{name}}
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudrun/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * If set to `true`, the revision name (template.metadata.name) will be omitted and
     * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
     * is also set.
     * (For legacy support, if `template.metadata.name` is unset in state while
     * this field is set to false, the revision name will still autogenerate.)
     */
    public readonly autogenerateRevisionName!: pulumi.Output<boolean | undefined>;
    /**
     * The location of the cloud run instance. eg us-central1
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional metadata for this Revision, including labels and annotations.
     * Name will be generated by the Configuration. To set minimum instances
     * for this revision, use the "autoscaling.knative.dev/minScale" annotation
     * key. To set maximum instances for this revision, use the
     * "autoscaling.knative.dev/maxScale" annotation key. To set Cloud SQL
     * connections for the revision, use the "run.googleapis.com/cloudsql-instances"
     * annotation key.
     * Structure is documented below.
     * (Optional)
     * Metadata associated with this Service, including name, namespace, labels,
     * and annotations.
     * Structure is documented below.
     */
    public readonly metadata!: pulumi.Output<outputs.cloudrun.ServiceMetadata>;
    /**
     * Name must be unique within a namespace, within a Cloud Run region.
     * Is required when creating resources. Name is primarily intended
     * for creation idempotence and configuration definition. Cannot be updated.
     * More info: http://kubernetes.io/docs/user-guide/identifiers#names
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Status of the condition, one of True, False, Unknown.
     */
    public /*out*/ readonly statuses!: pulumi.Output<outputs.cloudrun.ServiceStatus[]>;
    /**
     * template holds the latest specification for the Revision to be stamped out. The template references the container image,
     * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
     * force a Revision to be created when the spec doesn't otherwise change, a nonce label may be provided in the template
     * metadata. For more details, see:
     * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
     * does not currently support referencing a build that is responsible for materializing the container image from source.
     */
    public readonly template!: pulumi.Output<outputs.cloudrun.ServiceTemplate | undefined>;
    /**
     * Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
     */
    public readonly traffics!: pulumi.Output<outputs.cloudrun.ServiceTraffic[]>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["autogenerateRevisionName"] = state ? state.autogenerateRevisionName : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["statuses"] = state ? state.statuses : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
            resourceInputs["traffics"] = state ? state.traffics : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["autogenerateRevisionName"] = args ? args.autogenerateRevisionName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["template"] = args ? args.template : undefined;
            resourceInputs["traffics"] = args ? args.traffics : undefined;
            resourceInputs["statuses"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * If set to `true`, the revision name (template.metadata.name) will be omitted and
     * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
     * is also set.
     * (For legacy support, if `template.metadata.name` is unset in state while
     * this field is set to false, the revision name will still autogenerate.)
     */
    autogenerateRevisionName?: pulumi.Input<boolean>;
    /**
     * The location of the cloud run instance. eg us-central1
     */
    location?: pulumi.Input<string>;
    /**
     * Optional metadata for this Revision, including labels and annotations.
     * Name will be generated by the Configuration. To set minimum instances
     * for this revision, use the "autoscaling.knative.dev/minScale" annotation
     * key. To set maximum instances for this revision, use the
     * "autoscaling.knative.dev/maxScale" annotation key. To set Cloud SQL
     * connections for the revision, use the "run.googleapis.com/cloudsql-instances"
     * annotation key.
     * Structure is documented below.
     * (Optional)
     * Metadata associated with this Service, including name, namespace, labels,
     * and annotations.
     * Structure is documented below.
     */
    metadata?: pulumi.Input<inputs.cloudrun.ServiceMetadata>;
    /**
     * Name must be unique within a namespace, within a Cloud Run region.
     * Is required when creating resources. Name is primarily intended
     * for creation idempotence and configuration definition. Cannot be updated.
     * More info: http://kubernetes.io/docs/user-guide/identifiers#names
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Status of the condition, one of True, False, Unknown.
     */
    statuses?: pulumi.Input<pulumi.Input<inputs.cloudrun.ServiceStatus>[]>;
    /**
     * template holds the latest specification for the Revision to be stamped out. The template references the container image,
     * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
     * force a Revision to be created when the spec doesn't otherwise change, a nonce label may be provided in the template
     * metadata. For more details, see:
     * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
     * does not currently support referencing a build that is responsible for materializing the container image from source.
     */
    template?: pulumi.Input<inputs.cloudrun.ServiceTemplate>;
    /**
     * Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
     */
    traffics?: pulumi.Input<pulumi.Input<inputs.cloudrun.ServiceTraffic>[]>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * If set to `true`, the revision name (template.metadata.name) will be omitted and
     * autogenerated by Cloud Run. This cannot be set to `true` while `template.metadata.name`
     * is also set.
     * (For legacy support, if `template.metadata.name` is unset in state while
     * this field is set to false, the revision name will still autogenerate.)
     */
    autogenerateRevisionName?: pulumi.Input<boolean>;
    /**
     * The location of the cloud run instance. eg us-central1
     */
    location: pulumi.Input<string>;
    /**
     * Optional metadata for this Revision, including labels and annotations.
     * Name will be generated by the Configuration. To set minimum instances
     * for this revision, use the "autoscaling.knative.dev/minScale" annotation
     * key. To set maximum instances for this revision, use the
     * "autoscaling.knative.dev/maxScale" annotation key. To set Cloud SQL
     * connections for the revision, use the "run.googleapis.com/cloudsql-instances"
     * annotation key.
     * Structure is documented below.
     * (Optional)
     * Metadata associated with this Service, including name, namespace, labels,
     * and annotations.
     * Structure is documented below.
     */
    metadata?: pulumi.Input<inputs.cloudrun.ServiceMetadata>;
    /**
     * Name must be unique within a namespace, within a Cloud Run region.
     * Is required when creating resources. Name is primarily intended
     * for creation idempotence and configuration definition. Cannot be updated.
     * More info: http://kubernetes.io/docs/user-guide/identifiers#names
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * template holds the latest specification for the Revision to be stamped out. The template references the container image,
     * and may also include labels and annotations that should be attached to the Revision. To correlate a Revision, and/or to
     * force a Revision to be created when the spec doesn't otherwise change, a nonce label may be provided in the template
     * metadata. For more details, see:
     * https://github.com/knative/serving/blob/main/docs/client-conventions.md#associate-modifications-with-revisions Cloud Run
     * does not currently support referencing a build that is responsible for materializing the container image from source.
     */
    template?: pulumi.Input<inputs.cloudrun.ServiceTemplate>;
    /**
     * Traffic specifies how to distribute traffic over a collection of Knative Revisions and Configurations
     */
    traffics?: pulumi.Input<pulumi.Input<inputs.cloudrun.ServiceTraffic>[]>;
}
