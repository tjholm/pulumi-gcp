// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * !> **Warning:** This data source is deprecated. Use the `gcp.kms.SecretCiphertext` **resource** instead.
 *
 * This data source allows you to encrypt data with Google Cloud KMS and use the
 * ciphertext within your resource definitions.
 *
 * For more information see
 * [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
 *
 * > **NOTE:** Using this data source will allow you to conceal secret data within your
 * resource definitions, but it does not take care of protecting that data in the
 * logging output, plan output, or state output.  Please take care to secure your secret
 * data outside of resource definitions.
 */
export function getKMSSecretCiphertext(args: GetKMSSecretCiphertextArgs, opts?: pulumi.InvokeOptions): Promise<GetKMSSecretCiphertextResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gcp:kms/getKMSSecretCiphertext:getKMSSecretCiphertext", {
        "cryptoKey": args.cryptoKey,
        "plaintext": args.plaintext,
    }, opts);
}

/**
 * A collection of arguments for invoking getKMSSecretCiphertext.
 */
export interface GetKMSSecretCiphertextArgs {
    /**
     * The id of the CryptoKey that will be used to
     * encrypt the provided plaintext. This is represented by the format
     * `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
     */
    cryptoKey: string;
    /**
     * The plaintext to be encrypted
     */
    plaintext: string;
}

/**
 * A collection of values returned by getKMSSecretCiphertext.
 */
export interface GetKMSSecretCiphertextResult {
    /**
     * Contains the result of encrypting the provided plaintext, encoded in base64.
     */
    readonly ciphertext: string;
    readonly cryptoKey: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly plaintext: string;
}

export function getKMSSecretCiphertextOutput(args: GetKMSSecretCiphertextOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKMSSecretCiphertextResult> {
    return pulumi.output(args).apply(a => getKMSSecretCiphertext(a, opts))
}

/**
 * A collection of arguments for invoking getKMSSecretCiphertext.
 */
export interface GetKMSSecretCiphertextOutputArgs {
    /**
     * The id of the CryptoKey that will be used to
     * encrypt the provided plaintext. This is represented by the format
     * `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
     */
    cryptoKey: pulumi.Input<string>;
    /**
     * The plaintext to be encrypted
     */
    plaintext: pulumi.Input<string>;
}
