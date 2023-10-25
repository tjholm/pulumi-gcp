// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get the service account from a project. For more information see
 * the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const objectViewer = gcp.serviceAccount.getAccount({
 *     accountId: "object-viewer",
 * });
 * ```
 */
export function getAccount(args: GetAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:serviceAccount/getAccount:getAccount", {
        "accountId": args.accountId,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountArgs {
    /**
     * The Google service account ID. This be one of:
     *
     * * The name of the service account within the project (e.g. `my-service`)
     *
     * * The fully-qualified path to a service account resource (e.g.
     * `projects/my-project/serviceAccounts/...`)
     *
     * * The email address of the service account (e.g.
     * `my-service@my-project.iam.gserviceaccount.com`)
     */
    accountId: string;
    /**
     * The ID of the project that the service account is present in.
     * Defaults to the provider project configuration.
     */
    project?: string;
}

/**
 * A collection of values returned by getAccount.
 */
export interface GetAccountResult {
    readonly accountId: string;
    /**
     * The display name for the service account.
     */
    readonly displayName: string;
    /**
     * The e-mail address of the service account. This value
     * should be referenced from any `gcp.organizations.getIAMPolicy` data sources
     * that would grant the service account privileges.
     */
    readonly email: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
     */
    readonly member: string;
    /**
     * The fully-qualified name of the service account.
     */
    readonly name: string;
    readonly project?: string;
    /**
     * The unique id of the service account.
     */
    readonly uniqueId: string;
}
/**
 * Get the service account from a project. For more information see
 * the official [API](https://cloud.google.com/compute/docs/access/service-accounts) documentation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const objectViewer = gcp.serviceAccount.getAccount({
 *     accountId: "object-viewer",
 * });
 * ```
 */
export function getAccountOutput(args: GetAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountResult> {
    return pulumi.output(args).apply((a: any) => getAccount(a, opts))
}

/**
 * A collection of arguments for invoking getAccount.
 */
export interface GetAccountOutputArgs {
    /**
     * The Google service account ID. This be one of:
     *
     * * The name of the service account within the project (e.g. `my-service`)
     *
     * * The fully-qualified path to a service account resource (e.g.
     * `projects/my-project/serviceAccounts/...`)
     *
     * * The email address of the service account (e.g.
     * `my-service@my-project.iam.gserviceaccount.com`)
     */
    accountId: pulumi.Input<string>;
    /**
     * The ID of the project that the service account is present in.
     * Defaults to the provider project configuration.
     */
    project?: pulumi.Input<string>;
}
