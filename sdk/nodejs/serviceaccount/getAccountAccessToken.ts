// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source provides a google `oauth2` `accessToken` for a different service account than the one initially running the script.
 *
 * For more information see
 * [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
 *
 * ## Example Usage
 *
 * To allow `service_A` to impersonate `service_B`, grant the [Service Account Token Creator](https://cloud.google.com/iam/docs/service-accounts#the_service_account_token_creator_role) on B to A.
 *
 * In the IAM policy below, `service_A` is given the Token Creator role impersonate `service_B`
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const token_creator_iam = new gcp.serviceaccount.IAMBinding("token-creator-iam", {
 *     members: ["serviceAccount:service_A@projectA.iam.gserviceaccount.com"],
 *     role: "roles/iam.serviceAccountTokenCreator",
 *     serviceAccountId: "projects/-/serviceAccounts/service_B@projectB.iam.gserviceaccount.com",
 * });
 * ```
 *
 * Once the IAM permissions are set, you can apply the new token to a provider bootstrapped with it.  Any resources that references the aliased provider will run as the new identity.
 *
 * In the example below, `gcp.organizations.Project` will run as `service_B`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * export = async () => {
 *     const defaultClientConfig = await gcp.organizations.getClientConfig({});
 *     const defaultAccountAccessToken = await gcp.serviceAccount.getAccountAccessToken({
 *         targetServiceAccount: "service_B@projectB.iam.gserviceaccount.com",
 *         scopes: [
 *             "userinfo-email",
 *             "cloud-platform",
 *         ],
 *         lifetime: "300s",
 *     });
 *     const impersonated = new pulumi.providers.Google("impersonated", {accessToken: defaultAccountAccessToken.accessToken});
 *     const me = await gcp.organizations.getClientOpenIdUserInfo({});
 *     return {
 *         "target-email": me.email,
 *     };
 * }
 * ```
 *
 * > *Note*: the generated token is non-refreshable and can have a maximum `lifetime` of `3600` seconds.
 */
export function getAccountAccessToken(args: GetAccountAccessTokenArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountAccessTokenResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:serviceAccount/getAccountAccessToken:getAccountAccessToken", {
        "delegates": args.delegates,
        "lifetime": args.lifetime,
        "scopes": args.scopes,
        "targetServiceAccount": args.targetServiceAccount,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountAccessToken.
 */
export interface GetAccountAccessTokenArgs {
    /**
     * Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
     */
    delegates?: string[];
    /**
     * Lifetime of the impersonated token (defaults to its max: `3600s`).
     */
    lifetime?: string;
    /**
     * The scopes the new credential should have (e.g. `["cloud-platform"]`)
     */
    scopes: string[];
    /**
     * The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
     */
    targetServiceAccount: string;
}

/**
 * A collection of values returned by getAccountAccessToken.
 */
export interface GetAccountAccessTokenResult {
    /**
     * The `accessToken` representing the new generated identity.
     */
    readonly accessToken: string;
    readonly delegates?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lifetime?: string;
    readonly scopes: string[];
    readonly targetServiceAccount: string;
}
/**
 * This data source provides a google `oauth2` `accessToken` for a different service account than the one initially running the script.
 *
 * For more information see
 * [the official documentation](https://cloud.google.com/iam/docs/creating-short-lived-service-account-credentials) as well as [iamcredentials.generateAccessToken()](https://cloud.google.com/iam/credentials/reference/rest/v1/projects.serviceAccounts/generateAccessToken)
 *
 * ## Example Usage
 *
 * To allow `service_A` to impersonate `service_B`, grant the [Service Account Token Creator](https://cloud.google.com/iam/docs/service-accounts#the_service_account_token_creator_role) on B to A.
 *
 * In the IAM policy below, `service_A` is given the Token Creator role impersonate `service_B`
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const token_creator_iam = new gcp.serviceaccount.IAMBinding("token-creator-iam", {
 *     members: ["serviceAccount:service_A@projectA.iam.gserviceaccount.com"],
 *     role: "roles/iam.serviceAccountTokenCreator",
 *     serviceAccountId: "projects/-/serviceAccounts/service_B@projectB.iam.gserviceaccount.com",
 * });
 * ```
 *
 * Once the IAM permissions are set, you can apply the new token to a provider bootstrapped with it.  Any resources that references the aliased provider will run as the new identity.
 *
 * In the example below, `gcp.organizations.Project` will run as `service_B`.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * export = async () => {
 *     const defaultClientConfig = await gcp.organizations.getClientConfig({});
 *     const defaultAccountAccessToken = await gcp.serviceAccount.getAccountAccessToken({
 *         targetServiceAccount: "service_B@projectB.iam.gserviceaccount.com",
 *         scopes: [
 *             "userinfo-email",
 *             "cloud-platform",
 *         ],
 *         lifetime: "300s",
 *     });
 *     const impersonated = new pulumi.providers.Google("impersonated", {accessToken: defaultAccountAccessToken.accessToken});
 *     const me = await gcp.organizations.getClientOpenIdUserInfo({});
 *     return {
 *         "target-email": me.email,
 *     };
 * }
 * ```
 *
 * > *Note*: the generated token is non-refreshable and can have a maximum `lifetime` of `3600` seconds.
 */
export function getAccountAccessTokenOutput(args: GetAccountAccessTokenOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountAccessTokenResult> {
    return pulumi.output(args).apply((a: any) => getAccountAccessToken(a, opts))
}

/**
 * A collection of arguments for invoking getAccountAccessToken.
 */
export interface GetAccountAccessTokenOutputArgs {
    /**
     * Delegate chain of approvals needed to perform full impersonation. Specify the fully qualified service account name.  (e.g. `["projects/-/serviceAccounts/delegate-svc-account@project-id.iam.gserviceaccount.com"]`)
     */
    delegates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lifetime of the impersonated token (defaults to its max: `3600s`).
     */
    lifetime?: pulumi.Input<string>;
    /**
     * The scopes the new credential should have (e.g. `["cloud-platform"]`)
     */
    scopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The service account _to_ impersonate (e.g. `service_B@your-project-id.iam.gserviceaccount.com`)
     */
    targetServiceAccount: pulumi.Input<string>;
}
