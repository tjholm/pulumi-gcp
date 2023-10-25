// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get the email address of a project's unique [automatic Google Cloud Storage service account](https://cloud.google.com/storage/docs/projects#service-accounts).
 *
 * For each Google Cloud project, Google maintains a unique service account which
 * is used as the identity for various Google Cloud Storage operations, including
 * operations involving
 * [customer-managed encryption keys](https://cloud.google.com/storage/docs/encryption/customer-managed-keys)
 * and those involving
 * [storage notifications to pub/sub](https://cloud.google.com/storage/docs/gsutil/commands/notification).
 * This automatic Google service account requires access to the relevant Cloud KMS keys or pub/sub topics, respectively, in order for Cloud Storage to use
 * these customer-managed resources.
 *
 * The service account has a well-known, documented naming format which is parameterised on the numeric Google project ID.
 * However, as noted in [the docs](https://cloud.google.com/storage/docs/projects#service-accounts), it is only created when certain relevant actions occur which
 * presuppose its existence.
 * These actions include calling a [Cloud Storage API endpoint](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount/get) to yield the
 * service account's identity, or performing some operations in the UI which must use the service account's identity, such as attempting to list Cloud KMS keys
 * on the bucket creation page.
 *
 * Use of this data source calls the relevant API endpoint to obtain the service account's identity and thus ensures it exists prior to any API operations
 * which demand its existence, such as specifying it in Cloud IAM policy.
 * Always prefer to use this data source over interpolating the project ID into the well-known format for this service account, as the latter approach may cause
 * provider update errors in cases where the service account does not yet exist.
 *
 * >  When you write provider code which uses features depending on this service account *and* your provider code adds the service account in IAM policy on other resources,
 *    you must take care for race conditions between the establishment of the IAM policy and creation of the relevant Cloud Storage resource.
 *    Cloud Storage APIs will require permissions on resources such as pub/sub topics or Cloud KMS keys to exist *before* the attempt to utilise them in a
 *    bucket configuration, otherwise the API calls will fail.
 *    You may need to use `dependsOn` to create an explicit dependency between the IAM policy resource and the Cloud Storage resource which depends on it.
 *    See the examples here and in the `gcp.storage.Notification` resource.
 *
 * For more information see
 * [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).
 *
 * ## Example Usage
 * ### Pub/Sub Notifications
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gcsAccount = gcp.storage.getProjectServiceAccount({});
 * const binding = new gcp.pubsub.TopicIAMBinding("binding", {
 *     topic: google_pubsub_topic.topic.name,
 *     role: "roles/pubsub.publisher",
 *     members: [gcsAccount.then(gcsAccount => `serviceAccount:${gcsAccount.emailAddress}`)],
 * });
 * ```
 * ### Cloud KMS Keys
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gcsAccount = gcp.storage.getProjectServiceAccount({});
 * const binding = new gcp.kms.CryptoKeyIAMBinding("binding", {
 *     cryptoKeyId: "your-crypto-key-id",
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     members: [gcsAccount.then(gcsAccount => `serviceAccount:${gcsAccount.emailAddress}`)],
 * });
 * const bucket = new gcp.storage.Bucket("bucket", {
 *     location: "US",
 *     encryption: {
 *         defaultKmsKeyName: "your-crypto-key-id",
 *     },
 * }, {
 *     dependsOn: [binding],
 * });
 * ```
 */
export function getProjectServiceAccount(args?: GetProjectServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectServiceAccountResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:storage/getProjectServiceAccount:getProjectServiceAccount", {
        "project": args.project,
        "userProject": args.userProject,
    }, opts);
}

/**
 * A collection of arguments for invoking getProjectServiceAccount.
 */
export interface GetProjectServiceAccountArgs {
    /**
     * The project the unique service account was created for. If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The project the lookup originates from. This field is used if you are making the request
     * from a different account than the one you are finding the service account for.
     */
    userProject?: string;
}

/**
 * A collection of values returned by getProjectServiceAccount.
 */
export interface GetProjectServiceAccountResult {
    /**
     * The email address of the service account. This value is often used to refer to the service account
     * in order to grant IAM permissions.
     */
    readonly emailAddress: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The Identity of the service account in the form `serviceAccount:{email_address}`. This value is often used to refer to the service account in order to grant IAM permissions.
     */
    readonly member: string;
    readonly project: string;
    readonly userProject?: string;
}
/**
 * Get the email address of a project's unique [automatic Google Cloud Storage service account](https://cloud.google.com/storage/docs/projects#service-accounts).
 *
 * For each Google Cloud project, Google maintains a unique service account which
 * is used as the identity for various Google Cloud Storage operations, including
 * operations involving
 * [customer-managed encryption keys](https://cloud.google.com/storage/docs/encryption/customer-managed-keys)
 * and those involving
 * [storage notifications to pub/sub](https://cloud.google.com/storage/docs/gsutil/commands/notification).
 * This automatic Google service account requires access to the relevant Cloud KMS keys or pub/sub topics, respectively, in order for Cloud Storage to use
 * these customer-managed resources.
 *
 * The service account has a well-known, documented naming format which is parameterised on the numeric Google project ID.
 * However, as noted in [the docs](https://cloud.google.com/storage/docs/projects#service-accounts), it is only created when certain relevant actions occur which
 * presuppose its existence.
 * These actions include calling a [Cloud Storage API endpoint](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount/get) to yield the
 * service account's identity, or performing some operations in the UI which must use the service account's identity, such as attempting to list Cloud KMS keys
 * on the bucket creation page.
 *
 * Use of this data source calls the relevant API endpoint to obtain the service account's identity and thus ensures it exists prior to any API operations
 * which demand its existence, such as specifying it in Cloud IAM policy.
 * Always prefer to use this data source over interpolating the project ID into the well-known format for this service account, as the latter approach may cause
 * provider update errors in cases where the service account does not yet exist.
 *
 * >  When you write provider code which uses features depending on this service account *and* your provider code adds the service account in IAM policy on other resources,
 *    you must take care for race conditions between the establishment of the IAM policy and creation of the relevant Cloud Storage resource.
 *    Cloud Storage APIs will require permissions on resources such as pub/sub topics or Cloud KMS keys to exist *before* the attempt to utilise them in a
 *    bucket configuration, otherwise the API calls will fail.
 *    You may need to use `dependsOn` to create an explicit dependency between the IAM policy resource and the Cloud Storage resource which depends on it.
 *    See the examples here and in the `gcp.storage.Notification` resource.
 *
 * For more information see
 * [the API reference](https://cloud.google.com/storage/docs/json_api/v1/projects/serviceAccount).
 *
 * ## Example Usage
 * ### Pub/Sub Notifications
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gcsAccount = gcp.storage.getProjectServiceAccount({});
 * const binding = new gcp.pubsub.TopicIAMBinding("binding", {
 *     topic: google_pubsub_topic.topic.name,
 *     role: "roles/pubsub.publisher",
 *     members: [gcsAccount.then(gcsAccount => `serviceAccount:${gcsAccount.emailAddress}`)],
 * });
 * ```
 * ### Cloud KMS Keys
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gcsAccount = gcp.storage.getProjectServiceAccount({});
 * const binding = new gcp.kms.CryptoKeyIAMBinding("binding", {
 *     cryptoKeyId: "your-crypto-key-id",
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     members: [gcsAccount.then(gcsAccount => `serviceAccount:${gcsAccount.emailAddress}`)],
 * });
 * const bucket = new gcp.storage.Bucket("bucket", {
 *     location: "US",
 *     encryption: {
 *         defaultKmsKeyName: "your-crypto-key-id",
 *     },
 * }, {
 *     dependsOn: [binding],
 * });
 * ```
 */
export function getProjectServiceAccountOutput(args?: GetProjectServiceAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProjectServiceAccountResult> {
    return pulumi.output(args).apply((a: any) => getProjectServiceAccount(a, opts))
}

/**
 * A collection of arguments for invoking getProjectServiceAccount.
 */
export interface GetProjectServiceAccountOutputArgs {
    /**
     * The project the unique service account was created for. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The project the lookup originates from. This field is used if you are making the request
     * from a different account than the one you are finding the service account for.
     */
    userProject?: pulumi.Input<string>;
}
