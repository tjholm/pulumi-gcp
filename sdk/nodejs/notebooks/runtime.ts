// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Cloud AI Platform Notebook runtime.
 *
 * > **Note:** Due to limitations of the Notebooks Runtime API, many fields
 * in this resource do not properly detect drift. These fields will also not
 * appear in state once imported.
 *
 * To get more information about Runtime, see:
 *
 * * [API documentation](https://cloud.google.com/ai-platform/notebooks/docs/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/ai-platform-notebooks)
 *
 * ## Example Usage
 * ### Notebook Runtime Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const runtime = new gcp.notebooks.Runtime("runtime", {
 *     accessConfig: {
 *         accessType: "SINGLE_USER",
 *         runtimeOwner: "admin@hashicorptest.com",
 *     },
 *     location: "us-central1",
 *     virtualMachine: {
 *         virtualMachineConfig: {
 *             dataDisk: {
 *                 initializeParams: {
 *                     diskSizeGb: 100,
 *                     diskType: "PD_STANDARD",
 *                 },
 *             },
 *             machineType: "n1-standard-4",
 *         },
 *     },
 * });
 * ```
 * ### Notebook Runtime Basic Gpu
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const runtimeGpu = new gcp.notebooks.Runtime("runtimeGpu", {
 *     accessConfig: {
 *         accessType: "SINGLE_USER",
 *         runtimeOwner: "admin@hashicorptest.com",
 *     },
 *     location: "us-central1",
 *     softwareConfig: {
 *         installGpuDriver: true,
 *     },
 *     virtualMachine: {
 *         virtualMachineConfig: {
 *             acceleratorConfig: {
 *                 coreCount: 1,
 *                 type: "NVIDIA_TESLA_V100",
 *             },
 *             dataDisk: {
 *                 initializeParams: {
 *                     diskSizeGb: 100,
 *                     diskType: "PD_STANDARD",
 *                 },
 *             },
 *             machineType: "n1-standard-4",
 *         },
 *     },
 * });
 * ```
 * ### Notebook Runtime Basic Container
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const runtimeContainer = new gcp.notebooks.Runtime("runtimeContainer", {
 *     accessConfig: {
 *         accessType: "SINGLE_USER",
 *         runtimeOwner: "admin@hashicorptest.com",
 *     },
 *     location: "us-central1",
 *     virtualMachine: {
 *         virtualMachineConfig: {
 *             containerImages: [
 *                 {
 *                     repository: "gcr.io/deeplearning-platform-release/base-cpu",
 *                     tag: "latest",
 *                 },
 *                 {
 *                     repository: "gcr.io/deeplearning-platform-release/beam-notebooks",
 *                     tag: "latest",
 *                 },
 *             ],
 *             dataDisk: {
 *                 initializeParams: {
 *                     diskSizeGb: 100,
 *                     diskType: "PD_STANDARD",
 *                 },
 *             },
 *             machineType: "n1-standard-4",
 *         },
 *     },
 * });
 * ```
 * ### Notebook Runtime Kernels
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const runtimeContainer = new gcp.notebooks.Runtime("runtimeContainer", {
 *     accessConfig: {
 *         accessType: "SINGLE_USER",
 *         runtimeOwner: "admin@hashicorptest.com",
 *     },
 *     labels: {
 *         k: "val",
 *     },
 *     location: "us-central1",
 *     softwareConfig: {
 *         kernels: [{
 *             repository: "gcr.io/deeplearning-platform-release/base-cpu",
 *             tag: "latest",
 *         }],
 *     },
 *     virtualMachine: {
 *         virtualMachineConfig: {
 *             dataDisk: {
 *                 initializeParams: {
 *                     diskSizeGb: 100,
 *                     diskType: "PD_STANDARD",
 *                 },
 *             },
 *             machineType: "n1-standard-4",
 *         },
 *     },
 * });
 * ```
 * ### Notebook Runtime Script
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const runtimeContainer = new gcp.notebooks.Runtime("runtimeContainer", {
 *     accessConfig: {
 *         accessType: "SINGLE_USER",
 *         runtimeOwner: "admin@hashicorptest.com",
 *     },
 *     labels: {
 *         k: "val",
 *     },
 *     location: "us-central1",
 *     softwareConfig: {
 *         postStartupScriptBehavior: "RUN_EVERY_START",
 *     },
 *     virtualMachine: {
 *         virtualMachineConfig: {
 *             dataDisk: {
 *                 initializeParams: {
 *                     diskSizeGb: 100,
 *                     diskType: "PD_STANDARD",
 *                 },
 *             },
 *             machineType: "n1-standard-4",
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Runtime can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/runtimes/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Runtime can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:notebooks/runtime:Runtime default projects/{{project}}/locations/{{location}}/runtimes/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:notebooks/runtime:Runtime default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:notebooks/runtime:Runtime default {{location}}/{{name}}
 * ```
 */
export class Runtime extends pulumi.CustomResource {
    /**
     * Get an existing Runtime resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuntimeState, opts?: pulumi.CustomResourceOptions): Runtime {
        return new Runtime(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:notebooks/runtime:Runtime';

    /**
     * Returns true if the given object is an instance of Runtime.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Runtime {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Runtime.__pulumiType;
    }

    /**
     * The config settings for accessing runtime.
     * Structure is documented below.
     */
    public readonly accessConfig!: pulumi.Output<outputs.notebooks.RuntimeAccessConfig | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The health state of this runtime. For a list of possible output
     * values, see `https://cloud.google.com/vertex-ai/docs/workbench/
     * reference/rest/v1/projects.locations.runtimes#healthstate`.
     */
    public /*out*/ readonly healthState!: pulumi.Output<string>;
    /**
     * The labels to associate with this runtime. Label **keys** must
     * contain 1 to 63 characters, and must conform to [RFC 1035]
     * (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
     * empty, but, if present, must contain 1 to 63 characters, and must
     * conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
     * more than 32 labels can be associated with a cluster.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A reference to the zone where the machine resides.
     *
     *
     * - - -
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Contains Runtime daemon metrics such as Service status and JupyterLab
     * status
     * Structure is documented below.
     */
    public /*out*/ readonly metrics!: pulumi.Output<outputs.notebooks.RuntimeMetric[]>;
    /**
     * The name specified for the Notebook runtime.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The config settings for software inside the runtime.
     * Structure is documented below.
     */
    public readonly softwareConfig!: pulumi.Output<outputs.notebooks.RuntimeSoftwareConfig>;
    /**
     * The state of this runtime.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Use a Compute Engine VM image to start the managed notebook instance.
     * Structure is documented below.
     */
    public readonly virtualMachine!: pulumi.Output<outputs.notebooks.RuntimeVirtualMachine | undefined>;

    /**
     * Create a Runtime resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuntimeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuntimeArgs | RuntimeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RuntimeState | undefined;
            resourceInputs["accessConfig"] = state ? state.accessConfig : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["healthState"] = state ? state.healthState : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["metrics"] = state ? state.metrics : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["softwareConfig"] = state ? state.softwareConfig : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["virtualMachine"] = state ? state.virtualMachine : undefined;
        } else {
            const args = argsOrState as RuntimeArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["accessConfig"] = args ? args.accessConfig : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["softwareConfig"] = args ? args.softwareConfig : undefined;
            resourceInputs["virtualMachine"] = args ? args.virtualMachine : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["healthState"] = undefined /*out*/;
            resourceInputs["metrics"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Runtime.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Runtime resources.
 */
export interface RuntimeState {
    /**
     * The config settings for accessing runtime.
     * Structure is documented below.
     */
    accessConfig?: pulumi.Input<inputs.notebooks.RuntimeAccessConfig>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The health state of this runtime. For a list of possible output
     * values, see `https://cloud.google.com/vertex-ai/docs/workbench/
     * reference/rest/v1/projects.locations.runtimes#healthstate`.
     */
    healthState?: pulumi.Input<string>;
    /**
     * The labels to associate with this runtime. Label **keys** must
     * contain 1 to 63 characters, and must conform to [RFC 1035]
     * (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
     * empty, but, if present, must contain 1 to 63 characters, and must
     * conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
     * more than 32 labels can be associated with a cluster.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A reference to the zone where the machine resides.
     *
     *
     * - - -
     */
    location?: pulumi.Input<string>;
    /**
     * Contains Runtime daemon metrics such as Service status and JupyterLab
     * status
     * Structure is documented below.
     */
    metrics?: pulumi.Input<pulumi.Input<inputs.notebooks.RuntimeMetric>[]>;
    /**
     * The name specified for the Notebook runtime.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The config settings for software inside the runtime.
     * Structure is documented below.
     */
    softwareConfig?: pulumi.Input<inputs.notebooks.RuntimeSoftwareConfig>;
    /**
     * The state of this runtime.
     */
    state?: pulumi.Input<string>;
    /**
     * Use a Compute Engine VM image to start the managed notebook instance.
     * Structure is documented below.
     */
    virtualMachine?: pulumi.Input<inputs.notebooks.RuntimeVirtualMachine>;
}

/**
 * The set of arguments for constructing a Runtime resource.
 */
export interface RuntimeArgs {
    /**
     * The config settings for accessing runtime.
     * Structure is documented below.
     */
    accessConfig?: pulumi.Input<inputs.notebooks.RuntimeAccessConfig>;
    /**
     * The labels to associate with this runtime. Label **keys** must
     * contain 1 to 63 characters, and must conform to [RFC 1035]
     * (https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be
     * empty, but, if present, must contain 1 to 63 characters, and must
     * conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No
     * more than 32 labels can be associated with a cluster.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A reference to the zone where the machine resides.
     *
     *
     * - - -
     */
    location: pulumi.Input<string>;
    /**
     * The name specified for the Notebook runtime.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The config settings for software inside the runtime.
     * Structure is documented below.
     */
    softwareConfig?: pulumi.Input<inputs.notebooks.RuntimeSoftwareConfig>;
    /**
     * Use a Compute Engine VM image to start the managed notebook instance.
     * Structure is documented below.
     */
    virtualMachine?: pulumi.Input<inputs.notebooks.RuntimeVirtualMachine>;
}
