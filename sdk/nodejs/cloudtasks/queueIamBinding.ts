// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Cloud Tasks Queue. Each of these resources serves a different use case:
 *
 * * `gcp.cloudtasks.QueueIamPolicy`: Authoritative. Sets the IAM policy for the queue and replaces any existing policy already attached.
 * * `gcp.cloudtasks.QueueIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the queue are preserved.
 * * `gcp.cloudtasks.QueueIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the queue are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.cloudtasks.QueueIamPolicy`: Retrieves the IAM policy for the queue
 *
 * > **Note:** `gcp.cloudtasks.QueueIamPolicy` **cannot** be used in conjunction with `gcp.cloudtasks.QueueIamBinding` and `gcp.cloudtasks.QueueIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.cloudtasks.QueueIamBinding` resources **can be** used in conjunction with `gcp.cloudtasks.QueueIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_cloud\_tasks\_queue\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.cloudtasks.QueueIamPolicy("policy", {
 *     project: google_cloud_tasks_queue["default"].project,
 *     location: google_cloud_tasks_queue["default"].location,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_cloud\_tasks\_queue\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.cloudtasks.QueueIamBinding("binding", {
 *     project: google_cloud_tasks_queue["default"].project,
 *     location: google_cloud_tasks_queue["default"].location,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_cloud\_tasks\_queue\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.cloudtasks.QueueIamMember("member", {
 *     project: google_cloud_tasks_queue["default"].project,
 *     location: google_cloud_tasks_queue["default"].location,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/queues/{{name}} * {{project}}/{{location}}/{{name}} * {{location}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Tasks queue IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:cloudtasks/queueIamBinding:QueueIamBinding editor "projects/{{project}}/locations/{{location}}/queues/{{queue}} roles/viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:cloudtasks/queueIamBinding:QueueIamBinding editor "projects/{{project}}/locations/{{location}}/queues/{{queue}} roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:cloudtasks/queueIamBinding:QueueIamBinding editor projects/{{project}}/locations/{{location}}/queues/{{queue}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class QueueIamBinding extends pulumi.CustomResource {
    /**
     * Get an existing QueueIamBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueIamBindingState, opts?: pulumi.CustomResourceOptions): QueueIamBinding {
        return new QueueIamBinding(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudtasks/queueIamBinding:QueueIamBinding';

    /**
     * Returns true if the given object is an instance of QueueIamBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueueIamBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueueIamBinding.__pulumiType;
    }

    public readonly condition!: pulumi.Output<outputs.cloudtasks.QueueIamBindingCondition | undefined>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The location of the queue Used to find the parent resource to bind the IAM policy to
     */
    public readonly location!: pulumi.Output<string>;
    public readonly members!: pulumi.Output<string[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     *
     * * `member/members` - (Required) Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
     * * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
     * * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudtasks.QueueIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    public readonly role!: pulumi.Output<string>;

    /**
     * Create a QueueIamBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueIamBindingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueIamBindingArgs | QueueIamBindingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueueIamBindingState | undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as QueueIamBindingArgs | undefined;
            if ((!args || args.members === undefined) && !opts.urn) {
                throw new Error("Missing required property 'members'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QueueIamBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QueueIamBinding resources.
 */
export interface QueueIamBindingState {
    condition?: pulumi.Input<inputs.cloudtasks.QueueIamBindingCondition>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The location of the queue Used to find the parent resource to bind the IAM policy to
     */
    location?: pulumi.Input<string>;
    members?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     *
     * * `member/members` - (Required) Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
     * * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
     * * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
     */
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudtasks.QueueIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QueueIamBinding resource.
 */
export interface QueueIamBindingArgs {
    condition?: pulumi.Input<inputs.cloudtasks.QueueIamBindingCondition>;
    /**
     * The location of the queue Used to find the parent resource to bind the IAM policy to
     */
    location?: pulumi.Input<string>;
    members: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     *
     * * `member/members` - (Required) Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
     * * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
     * * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
     */
    project?: pulumi.Input<string>;
    /**
     * The role that should be applied. Only one
     * `gcp.cloudtasks.QueueIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     */
    role: pulumi.Input<string>;
}
