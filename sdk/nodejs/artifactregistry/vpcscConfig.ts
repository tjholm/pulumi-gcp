// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Artifact Registry Vpcsc Config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_config = new gcp.artifactregistry.VpcscConfig("my-config", {
 *     location: "us-central1",
 *     vpcscPolicy: "ALLOW",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * VPCSCConfig can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/vpcscConfig:VpcscConfig default projects/{{project}}/locations/{{location}}/vpcscConfig/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/vpcscConfig:VpcscConfig default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:artifactregistry/vpcscConfig:VpcscConfig default {{location}}/{{name}}
 * ```
 */
export class VpcscConfig extends pulumi.CustomResource {
    /**
     * Get an existing VpcscConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcscConfigState, opts?: pulumi.CustomResourceOptions): VpcscConfig {
        return new VpcscConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:artifactregistry/vpcscConfig:VpcscConfig';

    /**
     * Returns true if the given object is an instance of VpcscConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcscConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcscConfig.__pulumiType;
    }

    /**
     * The name of the location this config is located in.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the project's VPC SC Config.
     * Always of the form: projects/{project}/location/{location}/vpcscConfig
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The VPC SC policy for project and location.
     * Possible values are: `DENY`, `ALLOW`.
     */
    public readonly vpcscPolicy!: pulumi.Output<string | undefined>;

    /**
     * Create a VpcscConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpcscConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcscConfigArgs | VpcscConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcscConfigState | undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["vpcscPolicy"] = state ? state.vpcscPolicy : undefined;
        } else {
            const args = argsOrState as VpcscConfigArgs | undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["vpcscPolicy"] = args ? args.vpcscPolicy : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcscConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcscConfig resources.
 */
export interface VpcscConfigState {
    /**
     * The name of the location this config is located in.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the project's VPC SC Config.
     * Always of the form: projects/{project}/location/{location}/vpcscConfig
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The VPC SC policy for project and location.
     * Possible values are: `DENY`, `ALLOW`.
     */
    vpcscPolicy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcscConfig resource.
 */
export interface VpcscConfigArgs {
    /**
     * The name of the location this config is located in.
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The VPC SC policy for project and location.
     * Possible values are: `DENY`, `ALLOW`.
     */
    vpcscPolicy?: pulumi.Input<string>;
}
