// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Deployment Manager Deployment Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const deployment = new gcp.deploymentmanager.Deployment("deployment", {
 *     target: {
 *         config: {
 *             content: fs.readFileSync("path/to/config.yml"),
 *         },
 *     },
 *     labels: [{
 *         key: "foo",
 *         value: "bar",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Deployment can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:deploymentmanager/deployment:Deployment default projects/{{project}}/deployments/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:deploymentmanager/deployment:Deployment default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:deploymentmanager/deployment:Deployment default {{name}}
 * ```
 */
export class Deployment extends pulumi.CustomResource {
    /**
     * Get an existing Deployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentState, opts?: pulumi.CustomResourceOptions): Deployment {
        return new Deployment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:deploymentmanager/deployment:Deployment';

    /**
     * Returns true if the given object is an instance of Deployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Deployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Deployment.__pulumiType;
    }

    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     * Default value is `CREATE_OR_ACQUIRE`.
     * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
     */
    public readonly createPolicy!: pulumi.Output<string | undefined>;
    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     * Default value is `DELETE`.
     * Possible values are `ABANDON` and `DELETE`.
     */
    public readonly deletePolicy!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier for deployment. Output only.
     */
    public /*out*/ readonly deploymentId!: pulumi.Output<string>;
    /**
     * Optional user-provided description of deployment.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Key-value pairs to apply to this labels.
     * Structure is documented below.
     */
    public readonly labels!: pulumi.Output<outputs.deploymentmanager.DeploymentLabel[] | undefined>;
    /**
     * Output only. URL of the manifest representing the last manifest that
     * was successfully deployed.
     */
    public /*out*/ readonly manifest!: pulumi.Output<string>;
    /**
     * Unique name for the deployment
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If set to true, a deployment is created with "shell" resources that are not actually instantiated. This allows you to
     * preview a deployment. It can be updated to false to actually deploy with real resources. ~>**NOTE:** Deployment Manager
     * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
     * deployments if either preview is updated to true or if other fields are updated while preview is true.
     */
    public readonly preview!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Output only. Server defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     */
    public readonly target!: pulumi.Output<outputs.deploymentmanager.DeploymentTarget>;

    /**
     * Create a Deployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentArgs | DeploymentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeploymentState | undefined;
            resourceInputs["createPolicy"] = state ? state.createPolicy : undefined;
            resourceInputs["deletePolicy"] = state ? state.deletePolicy : undefined;
            resourceInputs["deploymentId"] = state ? state.deploymentId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["manifest"] = state ? state.manifest : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["preview"] = state ? state.preview : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as DeploymentArgs | undefined;
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["createPolicy"] = args ? args.createPolicy : undefined;
            resourceInputs["deletePolicy"] = args ? args.deletePolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["preview"] = args ? args.preview : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
            resourceInputs["deploymentId"] = undefined /*out*/;
            resourceInputs["manifest"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Deployment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Deployment resources.
 */
export interface DeploymentState {
    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     * Default value is `CREATE_OR_ACQUIRE`.
     * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
     */
    createPolicy?: pulumi.Input<string>;
    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     * Default value is `DELETE`.
     * Possible values are `ABANDON` and `DELETE`.
     */
    deletePolicy?: pulumi.Input<string>;
    /**
     * Unique identifier for deployment. Output only.
     */
    deploymentId?: pulumi.Input<string>;
    /**
     * Optional user-provided description of deployment.
     */
    description?: pulumi.Input<string>;
    /**
     * Key-value pairs to apply to this labels.
     * Structure is documented below.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.deploymentmanager.DeploymentLabel>[]>;
    /**
     * Output only. URL of the manifest representing the last manifest that
     * was successfully deployed.
     */
    manifest?: pulumi.Input<string>;
    /**
     * Unique name for the deployment
     */
    name?: pulumi.Input<string>;
    /**
     * If set to true, a deployment is created with "shell" resources that are not actually instantiated. This allows you to
     * preview a deployment. It can be updated to false to actually deploy with real resources. ~>**NOTE:** Deployment Manager
     * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
     * deployments if either preview is updated to true or if other fields are updated while preview is true.
     */
    preview?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Output only. Server defined URL for the resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     */
    target?: pulumi.Input<inputs.deploymentmanager.DeploymentTarget>;
}

/**
 * The set of arguments for constructing a Deployment resource.
 */
export interface DeploymentArgs {
    /**
     * Set the policy to use for creating new resources. Only used on
     * create and update. Valid values are `CREATE_OR_ACQUIRE` (default) or
     * `ACQUIRE`. If set to `ACQUIRE` and resources do not already exist,
     * the deployment will fail. Note that updating this field does not
     * actually affect the deployment, just how it is updated.
     * Default value is `CREATE_OR_ACQUIRE`.
     * Possible values are `ACQUIRE` and `CREATE_OR_ACQUIRE`.
     */
    createPolicy?: pulumi.Input<string>;
    /**
     * Set the policy to use for deleting new resources on update/delete.
     * Valid values are `DELETE` (default) or `ABANDON`. If `DELETE`,
     * resource is deleted after removal from Deployment Manager. If
     * `ABANDON`, the resource is only removed from Deployment Manager
     * and is not actually deleted. Note that updating this field does not
     * actually change the deployment, just how it is updated.
     * Default value is `DELETE`.
     * Possible values are `ABANDON` and `DELETE`.
     */
    deletePolicy?: pulumi.Input<string>;
    /**
     * Optional user-provided description of deployment.
     */
    description?: pulumi.Input<string>;
    /**
     * Key-value pairs to apply to this labels.
     * Structure is documented below.
     */
    labels?: pulumi.Input<pulumi.Input<inputs.deploymentmanager.DeploymentLabel>[]>;
    /**
     * Unique name for the deployment
     */
    name?: pulumi.Input<string>;
    /**
     * If set to true, a deployment is created with "shell" resources that are not actually instantiated. This allows you to
     * preview a deployment. It can be updated to false to actually deploy with real resources. ~>**NOTE:** Deployment Manager
     * does not allow update of a deployment in preview (unless updating to preview=false). Thus, Terraform will force-recreate
     * deployments if either preview is updated to true or if other fields are updated while preview is true.
     */
    preview?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Parameters that define your deployment, including the deployment
     * configuration and relevant templates.
     * Structure is documented below.
     */
    target: pulumi.Input<inputs.deploymentmanager.DeploymentTarget>;
}
