// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Configuration for an automated build in response to source repository changes.
 *
 * To get more information about Trigger, see:
 *
 * * [API documentation](https://cloud.google.com/cloud-build/docs/api/reference/rest/v1/projects.triggers)
 * * How-to Guides
 *     * [Automating builds using build triggers](https://cloud.google.com/cloud-build/docs/running-builds/automate-builds)
 *
 * > **Note:** You can retrieve the email of the Cloud Build Service Account used in jobs by using the `gcp.projects.ServiceIdentity` resource.
 *
 * ## Example Usage
 * ### Cloudbuild Trigger Filename
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const filename_trigger = new gcp.cloudbuild.Trigger("filename-trigger", {
 *     filename: "cloudbuild.yaml",
 *     substitutions: {
 *         _BAZ: "qux",
 *         _FOO: "bar",
 *     },
 *     triggerTemplate: {
 *         branchName: "master",
 *         repoName: "my-repo",
 *     },
 * });
 * ```
 * ### Cloudbuild Trigger Build
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const build_trigger = new gcp.cloudbuild.Trigger("build-trigger", {
 *     build: {
 *         artifacts: {
 *             images: ["gcr.io/$PROJECT_ID/$REPO_NAME:$COMMIT_SHA"],
 *             objects: {
 *                 location: "gs://bucket/path/to/somewhere/",
 *                 paths: ["path"],
 *             },
 *         },
 *         logsBucket: "gs://mybucket/logs",
 *         options: {
 *             diskSizeGb: 100,
 *             dynamicSubstitutions: true,
 *             envs: ["ekey = evalue"],
 *             logStreamingOption: "STREAM_OFF",
 *             logging: "LEGACY",
 *             machineType: "N1_HIGHCPU_8",
 *             requestedVerifyOption: "VERIFIED",
 *             secretEnvs: ["secretenv = svalue"],
 *             sourceProvenanceHashes: ["MD5"],
 *             substitutionOption: "ALLOW_LOOSE",
 *             volumes: [{
 *                 name: "v1",
 *                 path: "v1",
 *             }],
 *             workerPool: "pool",
 *         },
 *         queueTtl: "20s",
 *         secrets: [{
 *             kmsKeyName: "projects/myProject/locations/global/keyRings/keyring-name/cryptoKeys/key-name",
 *             secretEnv: {
 *                 PASSWORD: "ZW5jcnlwdGVkLXBhc3N3b3JkCg==",
 *             },
 *         }],
 *         source: {
 *             storageSource: {
 *                 bucket: "mybucket",
 *                 object: "source_code.tar.gz",
 *             },
 *         },
 *         steps: [{
 *             args: [
 *                 "cp",
 *                 "gs://mybucket/remotefile.zip",
 *                 "localfile.zip",
 *             ],
 *             name: "gcr.io/cloud-builders/gsutil",
 *             timeout: "120s",
 *         }],
 *         substitutions: {
 *             _BAZ: "qux",
 *             _FOO: "bar",
 *         },
 *         tags: [
 *             "build",
 *             "newFeature",
 *         ],
 *     },
 *     triggerTemplate: {
 *         branchName: "master",
 *         repoName: "my-repo",
 *     },
 * });
 * ```
 * ### Cloudbuild Trigger Service Account
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = gcp.organizations.getProject({});
 * const cloudbuildServiceAccount = new gcp.serviceaccount.Account("cloudbuildServiceAccount", {accountId: "my-service-account"});
 * const actAs = new gcp.projects.IAMMember("actAs", {
 *     project: project.then(project => project.projectId),
 *     role: "roles/iam.serviceAccountUser",
 *     member: pulumi.interpolate`serviceAccount:${cloudbuildServiceAccount.email}`,
 * });
 * const logsWriter = new gcp.projects.IAMMember("logsWriter", {
 *     project: project.then(project => project.projectId),
 *     role: "roles/logging.logWriter",
 *     member: pulumi.interpolate`serviceAccount:${cloudbuildServiceAccount.email}`,
 * });
 * const service_account_trigger = new gcp.cloudbuild.Trigger("service-account-trigger", {
 *     triggerTemplate: {
 *         branchName: "master",
 *         repoName: "my-repo",
 *     },
 *     serviceAccount: cloudbuildServiceAccount.id,
 *     filename: "cloudbuild.yaml",
 * }, {
 *     dependsOn: [
 *         actAs,
 *         logsWriter,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Trigger can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:cloudbuild/trigger:Trigger default projects/{{project}}/triggers/{{trigger_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudbuild/trigger:Trigger default {{project}}/{{trigger_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudbuild/trigger:Trigger default {{trigger_id}}
 * ```
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudbuild/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * Contents of the build template. Either a filename or build template must be provided.
     * Structure is documented below.
     */
    public readonly build!: pulumi.Output<outputs.cloudbuild.TriggerBuild | undefined>;
    /**
     * Time when the trigger was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Human-readable description of the trigger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the trigger is disabled or not. If true, the trigger will never result in a build.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
     */
    public readonly filename!: pulumi.Output<string | undefined>;
    /**
     * Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    public readonly github!: pulumi.Output<outputs.cloudbuild.TriggerGithub | undefined>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If ignoredFiles and changed files are both empty, then they are not
     * used to determine whether or not to trigger a build.
     * If ignoredFiles is not empty, then we ignore any files that match any
     * of the ignoredFile globs. If the change has no files that are outside
     * of the ignoredFiles globs, then we do not trigger a build.
     */
    public readonly ignoredFiles!: pulumi.Output<string[] | undefined>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is empty, then as far as this filter is concerned, we
     * should trigger the build.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is not empty, then we make sure that at least one of
     * those files matches a includedFiles glob. If not, then we do not trigger
     * a build.
     */
    public readonly includedFiles!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the volume to mount.
     * Volume names must be unique per build step and must be valid names for Docker volumes.
     * Each named volume must be used by at least two build steps.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * PubsubConfig describes the configuration of a trigger that creates
     * a build whenever a Pub/Sub message is published.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    public readonly pubsubConfig!: pulumi.Output<outputs.cloudbuild.TriggerPubsubConfig | undefined>;
    /**
     * The service account used for all user-controlled operations including
     * triggers.patch, triggers.run, builds.create, and builds.cancel.
     * If no service account is set, then the standard Cloud Build service account
     * ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
     * Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
     */
    public readonly serviceAccount!: pulumi.Output<string | undefined>;
    /**
     * Substitutions to use in a triggered build. Should only be used with triggers.run
     */
    public readonly substitutions!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Tags for annotation of a Build. These are not docker tags.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The unique identifier for the trigger.
     */
    public /*out*/ readonly triggerId!: pulumi.Output<string>;
    /**
     * Template describing the types of source changes to trigger a build.
     * Branch and tag names in trigger templates are interpreted as regular
     * expressions. Any branch or tag change that matches that regular
     * expression will trigger a build.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    public readonly triggerTemplate!: pulumi.Output<outputs.cloudbuild.TriggerTriggerTemplate | undefined>;
    /**
     * WebhookConfig describes the configuration of a trigger that creates
     * a build whenever a webhook is sent to a trigger's webhook URL.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    public readonly webhookConfig!: pulumi.Output<outputs.cloudbuild.TriggerWebhookConfig | undefined>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            inputs["build"] = state ? state.build : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["disabled"] = state ? state.disabled : undefined;
            inputs["filename"] = state ? state.filename : undefined;
            inputs["github"] = state ? state.github : undefined;
            inputs["ignoredFiles"] = state ? state.ignoredFiles : undefined;
            inputs["includedFiles"] = state ? state.includedFiles : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["pubsubConfig"] = state ? state.pubsubConfig : undefined;
            inputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            inputs["substitutions"] = state ? state.substitutions : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["triggerId"] = state ? state.triggerId : undefined;
            inputs["triggerTemplate"] = state ? state.triggerTemplate : undefined;
            inputs["webhookConfig"] = state ? state.webhookConfig : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            inputs["build"] = args ? args.build : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disabled"] = args ? args.disabled : undefined;
            inputs["filename"] = args ? args.filename : undefined;
            inputs["github"] = args ? args.github : undefined;
            inputs["ignoredFiles"] = args ? args.ignoredFiles : undefined;
            inputs["includedFiles"] = args ? args.includedFiles : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["pubsubConfig"] = args ? args.pubsubConfig : undefined;
            inputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            inputs["substitutions"] = args ? args.substitutions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["triggerTemplate"] = args ? args.triggerTemplate : undefined;
            inputs["webhookConfig"] = args ? args.webhookConfig : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["triggerId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Trigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * Contents of the build template. Either a filename or build template must be provided.
     * Structure is documented below.
     */
    build?: pulumi.Input<inputs.cloudbuild.TriggerBuild>;
    /**
     * Time when the trigger was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Human-readable description of the trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the trigger is disabled or not. If true, the trigger will never result in a build.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
     */
    filename?: pulumi.Input<string>;
    /**
     * Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    github?: pulumi.Input<inputs.cloudbuild.TriggerGithub>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If ignoredFiles and changed files are both empty, then they are not
     * used to determine whether or not to trigger a build.
     * If ignoredFiles is not empty, then we ignore any files that match any
     * of the ignoredFile globs. If the change has no files that are outside
     * of the ignoredFiles globs, then we do not trigger a build.
     */
    ignoredFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is empty, then as far as this filter is concerned, we
     * should trigger the build.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is not empty, then we make sure that at least one of
     * those files matches a includedFiles glob. If not, then we do not trigger
     * a build.
     */
    includedFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the volume to mount.
     * Volume names must be unique per build step and must be valid names for Docker volumes.
     * Each named volume must be used by at least two build steps.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * PubsubConfig describes the configuration of a trigger that creates
     * a build whenever a Pub/Sub message is published.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    pubsubConfig?: pulumi.Input<inputs.cloudbuild.TriggerPubsubConfig>;
    /**
     * The service account used for all user-controlled operations including
     * triggers.patch, triggers.run, builds.create, and builds.cancel.
     * If no service account is set, then the standard Cloud Build service account
     * ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
     * Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Substitutions to use in a triggered build. Should only be used with triggers.run
     */
    substitutions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Tags for annotation of a Build. These are not docker tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique identifier for the trigger.
     */
    triggerId?: pulumi.Input<string>;
    /**
     * Template describing the types of source changes to trigger a build.
     * Branch and tag names in trigger templates are interpreted as regular
     * expressions. Any branch or tag change that matches that regular
     * expression will trigger a build.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    triggerTemplate?: pulumi.Input<inputs.cloudbuild.TriggerTriggerTemplate>;
    /**
     * WebhookConfig describes the configuration of a trigger that creates
     * a build whenever a webhook is sent to a trigger's webhook URL.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    webhookConfig?: pulumi.Input<inputs.cloudbuild.TriggerWebhookConfig>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * Contents of the build template. Either a filename or build template must be provided.
     * Structure is documented below.
     */
    build?: pulumi.Input<inputs.cloudbuild.TriggerBuild>;
    /**
     * Human-readable description of the trigger.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the trigger is disabled or not. If true, the trigger will never result in a build.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.
     */
    filename?: pulumi.Input<string>;
    /**
     * Describes the configuration of a trigger that creates a build whenever a GitHub event is received.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    github?: pulumi.Input<inputs.cloudbuild.TriggerGithub>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If ignoredFiles and changed files are both empty, then they are not
     * used to determine whether or not to trigger a build.
     * If ignoredFiles is not empty, then we ignore any files that match any
     * of the ignoredFile globs. If the change has no files that are outside
     * of the ignoredFiles globs, then we do not trigger a build.
     */
    ignoredFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
     * extended with support for `**`.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is empty, then as far as this filter is concerned, we
     * should trigger the build.
     * If any of the files altered in the commit pass the ignoredFiles filter
     * and includedFiles is not empty, then we make sure that at least one of
     * those files matches a includedFiles glob. If not, then we do not trigger
     * a build.
     */
    includedFiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the volume to mount.
     * Volume names must be unique per build step and must be valid names for Docker volumes.
     * Each named volume must be used by at least two build steps.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * PubsubConfig describes the configuration of a trigger that creates
     * a build whenever a Pub/Sub message is published.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    pubsubConfig?: pulumi.Input<inputs.cloudbuild.TriggerPubsubConfig>;
    /**
     * The service account used for all user-controlled operations including
     * triggers.patch, triggers.run, builds.create, and builds.cancel.
     * If no service account is set, then the standard Cloud Build service account
     * ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.
     * Format: projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT_ID_OR_EMAIL}
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Substitutions to use in a triggered build. Should only be used with triggers.run
     */
    substitutions?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Tags for annotation of a Build. These are not docker tags.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Template describing the types of source changes to trigger a build.
     * Branch and tag names in trigger templates are interpreted as regular
     * expressions. Any branch or tag change that matches that regular
     * expression will trigger a build.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    triggerTemplate?: pulumi.Input<inputs.cloudbuild.TriggerTriggerTemplate>;
    /**
     * WebhookConfig describes the configuration of a trigger that creates
     * a build whenever a webhook is sent to a trigger's webhook URL.
     * One of `triggerTemplate`, `github`, `pubsubConfig` or `webhookConfig` must be provided.
     * Structure is documented below.
     */
    webhookConfig?: pulumi.Input<inputs.cloudbuild.TriggerWebhookConfig>;
}
