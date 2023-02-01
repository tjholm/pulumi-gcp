// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class BackendServiceSignedUrlKey extends pulumi.CustomResource {
    /**
     * Get an existing BackendServiceSignedUrlKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendServiceSignedUrlKeyState, opts?: pulumi.CustomResourceOptions): BackendServiceSignedUrlKey {
        return new BackendServiceSignedUrlKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey';

    /**
     * Returns true if the given object is an instance of BackendServiceSignedUrlKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendServiceSignedUrlKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendServiceSignedUrlKey.__pulumiType;
    }

    /**
     * The backend service this signed URL key belongs.
     */
    public readonly backendService!: pulumi.Output<string>;
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    public readonly keyValue!: pulumi.Output<string>;
    /**
     * Name of the signed URL key.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a BackendServiceSignedUrlKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendServiceSignedUrlKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendServiceSignedUrlKeyArgs | BackendServiceSignedUrlKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendServiceSignedUrlKeyState | undefined;
            resourceInputs["backendService"] = state ? state.backendService : undefined;
            resourceInputs["keyValue"] = state ? state.keyValue : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as BackendServiceSignedUrlKeyArgs | undefined;
            if ((!args || args.backendService === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backendService'");
            }
            if ((!args || args.keyValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyValue'");
            }
            resourceInputs["backendService"] = args ? args.backendService : undefined;
            resourceInputs["keyValue"] = args?.keyValue ? pulumi.secret(args.keyValue) : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["keyValue"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(BackendServiceSignedUrlKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendServiceSignedUrlKey resources.
 */
export interface BackendServiceSignedUrlKeyState {
    /**
     * The backend service this signed URL key belongs.
     */
    backendService?: pulumi.Input<string>;
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    keyValue?: pulumi.Input<string>;
    /**
     * Name of the signed URL key.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendServiceSignedUrlKey resource.
 */
export interface BackendServiceSignedUrlKeyArgs {
    /**
     * The backend service this signed URL key belongs.
     */
    backendService: pulumi.Input<string>;
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     */
    keyValue: pulumi.Input<string>;
    /**
     * Name of the signed URL key.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
