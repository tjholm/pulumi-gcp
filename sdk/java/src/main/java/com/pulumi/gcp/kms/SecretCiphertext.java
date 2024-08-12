// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.kms.SecretCiphertextArgs;
import com.pulumi.gcp.kms.inputs.SecretCiphertextState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Encrypts secret data with Google Cloud KMS and provides access to the ciphertext.
 * 
 * &gt; **NOTE:** Using this resource will allow you to conceal secret data within your
 * resource definitions, but it does not take care of protecting that data in the
 * logging output, plan output, or state output.  Please take care to secure your secret
 * data outside of resource definitions.
 * 
 * To get more information about SecretCiphertext, see:
 * 
 * * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys/encrypt)
 * * How-to Guides
 *     * [Encrypting and decrypting data with a symmetric key](https://cloud.google.com/kms/docs/encrypt-decrypt)
 * 
 * ## Example Usage
 * 
 * ### Kms Secret Ciphertext Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.kms.SecretCiphertext;
 * import com.pulumi.gcp.kms.SecretCiphertextArgs;
 * import com.pulumi.gcp.compute.Instance;
 * import com.pulumi.gcp.compute.InstanceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskInitializeParamsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var keyring = new KeyRing("keyring", KeyRingArgs.builder()
 *             .name("keyring-example")
 *             .location("global")
 *             .build());
 * 
 *         var cryptokey = new CryptoKey("cryptokey", CryptoKeyArgs.builder()
 *             .name("crypto-key-example")
 *             .keyRing(keyring.id())
 *             .rotationPeriod("7776000s")
 *             .build());
 * 
 *         var myPassword = new SecretCiphertext("myPassword", SecretCiphertextArgs.builder()
 *             .cryptoKey(cryptokey.id())
 *             .plaintext("my-secret-password")
 *             .build());
 * 
 *         var instance = new Instance("instance", InstanceArgs.builder()
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .accessConfigs()
 *                 .network("default")
 *                 .build())
 *             .name("my-instance")
 *             .machineType("e2-medium")
 *             .zone("us-central1-a")
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image("debian-cloud/debian-11")
 *                     .build())
 *                 .build())
 *             .metadata(Map.of("password", myPassword.ciphertext()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:kms/secretCiphertext:SecretCiphertext")
public class SecretCiphertext extends com.pulumi.resources.CustomResource {
    /**
     * The additional authenticated data used for integrity checks during encryption and decryption.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Export(name="additionalAuthenticatedData", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> additionalAuthenticatedData;

    /**
     * @return The additional authenticated data used for integrity checks during encryption and decryption.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<Optional<String>> additionalAuthenticatedData() {
        return Codegen.optional(this.additionalAuthenticatedData);
    }
    /**
     * Contains the result of encrypting the provided plaintext, encoded in base64.
     * 
     */
    @Export(name="ciphertext", refs={String.class}, tree="[0]")
    private Output<String> ciphertext;

    /**
     * @return Contains the result of encrypting the provided plaintext, encoded in base64.
     * 
     */
    public Output<String> ciphertext() {
        return this.ciphertext;
    }
    /**
     * The full name of the CryptoKey that will be used to encrypt the provided plaintext.
     * Format: `&#39;projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}&#39;`
     * 
     * ***
     * 
     */
    @Export(name="cryptoKey", refs={String.class}, tree="[0]")
    private Output<String> cryptoKey;

    /**
     * @return The full name of the CryptoKey that will be used to encrypt the provided plaintext.
     * Format: `&#39;projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}&#39;`
     * 
     * ***
     * 
     */
    public Output<String> cryptoKey() {
        return this.cryptoKey;
    }
    /**
     * The plaintext to be encrypted.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Export(name="plaintext", refs={String.class}, tree="[0]")
    private Output<String> plaintext;

    /**
     * @return The plaintext to be encrypted.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<String> plaintext() {
        return this.plaintext;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretCiphertext(java.lang.String name) {
        this(name, SecretCiphertextArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretCiphertext(java.lang.String name, SecretCiphertextArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretCiphertext(java.lang.String name, SecretCiphertextArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/secretCiphertext:SecretCiphertext", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecretCiphertext(java.lang.String name, Output<java.lang.String> id, @Nullable SecretCiphertextState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/secretCiphertext:SecretCiphertext", name, state, makeResourceOptions(options, id), false);
    }

    private static SecretCiphertextArgs makeArgs(SecretCiphertextArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecretCiphertextArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "additionalAuthenticatedData",
                "plaintext"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static SecretCiphertext get(java.lang.String name, Output<java.lang.String> id, @Nullable SecretCiphertextState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretCiphertext(name, id, state, options);
    }
}
