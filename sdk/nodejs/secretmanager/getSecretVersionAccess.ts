// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get a payload of Secret Manager secret's version. It only requires the [Secret Manager Secret Accessor](https://cloud.google.com/secret-manager/docs/access-control#secretmanager.secretAccessor) role. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets.versions/access).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = gcp.secretmanager.getSecretVersionAccess({
 *     secret: "my-secret",
 * });
 * ```
 */
export function getSecretVersionAccess(args: GetSecretVersionAccessArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretVersionAccessResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:secretmanager/getSecretVersionAccess:getSecretVersionAccess", {
        "project": args.project,
        "secret": args.secret,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretVersionAccess.
 */
export interface GetSecretVersionAccessArgs {
    /**
     * The project to get the secret version for. If it
     * is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The secret to get the secret version for.
     */
    secret: string;
    /**
     * The version of the secret to get. If it
     * is not provided, the latest version is retrieved.
     */
    version?: string;
}

/**
 * A collection of values returned by getSecretVersionAccess.
 */
export interface GetSecretVersionAccessResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The resource name of the SecretVersion. Format:
     * `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
     */
    readonly name: string;
    readonly project: string;
    readonly secret: string;
    /**
     * The secret data. No larger than 64KiB.
     */
    readonly secretData: string;
    readonly version: string;
}
/**
 * Get a payload of Secret Manager secret's version. It only requires the [Secret Manager Secret Accessor](https://cloud.google.com/secret-manager/docs/access-control#secretmanager.secretAccessor) role. For more information see the [official documentation](https://cloud.google.com/secret-manager/docs/) and [API](https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets.versions/access).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = gcp.secretmanager.getSecretVersionAccess({
 *     secret: "my-secret",
 * });
 * ```
 */
export function getSecretVersionAccessOutput(args: GetSecretVersionAccessOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretVersionAccessResult> {
    return pulumi.output(args).apply((a: any) => getSecretVersionAccess(a, opts))
}

/**
 * A collection of arguments for invoking getSecretVersionAccess.
 */
export interface GetSecretVersionAccessOutputArgs {
    /**
     * The project to get the secret version for. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The secret to get the secret version for.
     */
    secret: pulumi.Input<string>;
    /**
     * The version of the secret to get. If it
     * is not provided, the latest version is retrieved.
     */
    version?: pulumi.Input<string>;
}
