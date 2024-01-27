// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * A Google Cloud Firebase instance. This enables Firebase resources on a given google project.
 * Since a FirebaseProject is actually also a GCP Project, a FirebaseProject uses underlying GCP
 * identifiers (most importantly, the projectId) as its own for easy interop with GCP APIs.
 * Once Firebase has been added to a Google Project it cannot be removed.
 *
 * To get more information about Project, see:
 *
 * * [API documentation](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects)
 * * How-to Guides
 *     * [Official Documentation](https://firebase.google.com/)
 *
 * ## Example Usage
 * ### Firebase Project Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultProject = new gcp.organizations.Project("defaultProject", {
 *     orgId: "123456789",
 *     labels: {
 *         firebase: "enabled",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const defaultFirebase_projectProject = new gcp.firebase.Project("defaultFirebase/projectProject", {project: defaultProject.projectId}, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * Project can be imported using any of these accepted formats* `projects/{{project}}` * `{{project}}` When using the `pulumi import` command, Project can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:firebase/project:Project default projects/{{project}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:firebase/project:Project default {{project}}
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:firebase/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * The GCP project display name
     */
    public /*out*/ readonly displayName!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The number of the google project that firebase is enabled on.
     */
    public /*out*/ readonly projectNumber!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["projectNumber"] = state ? state.projectNumber : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["projectNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * The GCP project display name
     */
    displayName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The number of the google project that firebase is enabled on.
     */
    projectNumber?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
