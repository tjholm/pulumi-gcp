// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Dataform Repository Release Config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gitRepository = new gcp.sourcerepo.Repository("gitRepository", {}, {
 *     provider: google_beta,
 * });
 * const secret = new gcp.secretmanager.Secret("secret", {
 *     secretId: "my_secret",
 *     replication: {
 *         auto: {},
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const secretVersion = new gcp.secretmanager.SecretVersion("secretVersion", {
 *     secret: secret.id,
 *     secretData: "secret-data",
 * }, {
 *     provider: google_beta,
 * });
 * const repository = new gcp.dataform.Repository("repository", {
 *     region: "us-central1",
 *     gitRemoteSettings: {
 *         url: gitRepository.url,
 *         defaultBranch: "main",
 *         authenticationTokenSecretVersion: secretVersion.id,
 *     },
 *     workspaceCompilationOverrides: {
 *         defaultDatabase: "database",
 *         schemaSuffix: "_suffix",
 *         tablePrefix: "prefix_",
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * const release = new gcp.dataform.RepositoryReleaseConfig("release", {
 *     project: repository.project,
 *     region: repository.region,
 *     repository: repository.name,
 *     gitCommitish: "main",
 *     cronSchedule: "0 7 * * *",
 *     timeZone: "America/New_York",
 *     codeCompilationConfig: {
 *         defaultDatabase: "gcp-example-project",
 *         defaultSchema: "example-dataset",
 *         defaultLocation: "us-central1",
 *         assertionSchema: "example-assertion-dataset",
 *         databaseSuffix: "",
 *         schemaSuffix: "",
 *         tablePrefix: "",
 *         vars: {
 *             var1: "value",
 *         },
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * RepositoryReleaseConfig can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/repositories/{{repository}}/releaseConfigs/{{name}}` * `{{project}}/{{region}}/{{repository}}/{{name}}` * `{{region}}/{{repository}}/{{name}}` * `{{repository}}/{{name}}` When using the `pulumi import` command, RepositoryReleaseConfig can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default projects/{{project}}/locations/{{region}}/repositories/{{repository}}/releaseConfigs/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default {{project}}/{{region}}/{{repository}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default {{region}}/{{repository}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig default {{repository}}/{{name}}
 * ```
 */
export class RepositoryReleaseConfig extends pulumi.CustomResource {
    /**
     * Get an existing RepositoryReleaseConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryReleaseConfigState, opts?: pulumi.CustomResourceOptions): RepositoryReleaseConfig {
        return new RepositoryReleaseConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataform/repositoryReleaseConfig:RepositoryReleaseConfig';

    /**
     * Returns true if the given object is an instance of RepositoryReleaseConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RepositoryReleaseConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RepositoryReleaseConfig.__pulumiType;
    }

    /**
     * Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
     * Structure is documented below.
     */
    public readonly codeCompilationConfig!: pulumi.Output<outputs.dataform.RepositoryReleaseConfigCodeCompilationConfig | undefined>;
    /**
     * Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     */
    public readonly cronSchedule!: pulumi.Output<string | undefined>;
    /**
     * Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
     *
     *
     * - - -
     */
    public readonly gitCommitish!: pulumi.Output<string>;
    /**
     * The release's name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Records of the 10 most recent scheduled release attempts, ordered in in descending order of releaseTime. Updated whenever automatic creation of a compilation result is triggered by cronSchedule.
     * Structure is documented below.
     */
    public /*out*/ readonly recentScheduledReleaseRecords!: pulumi.Output<outputs.dataform.RepositoryReleaseConfigRecentScheduledReleaseRecord[]>;
    /**
     * A reference to the region
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * A reference to the Dataform repository
     */
    public readonly repository!: pulumi.Output<string | undefined>;
    /**
     * Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     */
    public readonly timeZone!: pulumi.Output<string | undefined>;

    /**
     * Create a RepositoryReleaseConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryReleaseConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryReleaseConfigArgs | RepositoryReleaseConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryReleaseConfigState | undefined;
            resourceInputs["codeCompilationConfig"] = state ? state.codeCompilationConfig : undefined;
            resourceInputs["cronSchedule"] = state ? state.cronSchedule : undefined;
            resourceInputs["gitCommitish"] = state ? state.gitCommitish : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["recentScheduledReleaseRecords"] = state ? state.recentScheduledReleaseRecords : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["repository"] = state ? state.repository : undefined;
            resourceInputs["timeZone"] = state ? state.timeZone : undefined;
        } else {
            const args = argsOrState as RepositoryReleaseConfigArgs | undefined;
            if ((!args || args.gitCommitish === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gitCommitish'");
            }
            resourceInputs["codeCompilationConfig"] = args ? args.codeCompilationConfig : undefined;
            resourceInputs["cronSchedule"] = args ? args.cronSchedule : undefined;
            resourceInputs["gitCommitish"] = args ? args.gitCommitish : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["repository"] = args ? args.repository : undefined;
            resourceInputs["timeZone"] = args ? args.timeZone : undefined;
            resourceInputs["recentScheduledReleaseRecords"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RepositoryReleaseConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RepositoryReleaseConfig resources.
 */
export interface RepositoryReleaseConfigState {
    /**
     * Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
     * Structure is documented below.
     */
    codeCompilationConfig?: pulumi.Input<inputs.dataform.RepositoryReleaseConfigCodeCompilationConfig>;
    /**
     * Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     */
    cronSchedule?: pulumi.Input<string>;
    /**
     * Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
     *
     *
     * - - -
     */
    gitCommitish?: pulumi.Input<string>;
    /**
     * The release's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Records of the 10 most recent scheduled release attempts, ordered in in descending order of releaseTime. Updated whenever automatic creation of a compilation result is triggered by cronSchedule.
     * Structure is documented below.
     */
    recentScheduledReleaseRecords?: pulumi.Input<pulumi.Input<inputs.dataform.RepositoryReleaseConfigRecentScheduledReleaseRecord>[]>;
    /**
     * A reference to the region
     */
    region?: pulumi.Input<string>;
    /**
     * A reference to the Dataform repository
     */
    repository?: pulumi.Input<string>;
    /**
     * Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     */
    timeZone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RepositoryReleaseConfig resource.
 */
export interface RepositoryReleaseConfigArgs {
    /**
     * Optional. If set, fields of codeCompilationConfig override the default compilation settings that are specified in dataform.json.
     * Structure is documented below.
     */
    codeCompilationConfig?: pulumi.Input<inputs.dataform.RepositoryReleaseConfigCodeCompilationConfig>;
    /**
     * Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     */
    cronSchedule?: pulumi.Input<string>;
    /**
     * Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository.
     *
     *
     * - - -
     */
    gitCommitish: pulumi.Input<string>;
    /**
     * The release's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A reference to the region
     */
    region?: pulumi.Input<string>;
    /**
     * A reference to the Dataform repository
     */
    repository?: pulumi.Input<string>;
    /**
     * Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     */
    timeZone?: pulumi.Input<string>;
}
