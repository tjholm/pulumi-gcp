// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.kms.CryptoKeyArgs;
import com.pulumi.gcp.kms.inputs.CryptoKeyState;
import com.pulumi.gcp.kms.outputs.CryptoKeyPrimary;
import com.pulumi.gcp.kms.outputs.CryptoKeyVersionTemplate;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A `CryptoKey` represents a logical key that can be used for cryptographic operations.
 * 
 * &gt; **Note:** CryptoKeys cannot be deleted from Google Cloud Platform.
 * Destroying a provider-managed CryptoKey will remove it from state
 * and delete all CryptoKeyVersions, rendering the key unusable, but *will
 * not delete the resource from the project.* When the provider destroys these keys,
 * any data previously encrypted with these keys will be irrecoverable.
 * For this reason, it is strongly recommended that you add
 * lifecycle
 * hooks to the resource to prevent accidental destruction.
 * 
 * To get more information about CryptoKey, see:
 * 
 * * [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys)
 * * How-to Guides
 *     * [Creating a key](https://cloud.google.com/kms/docs/creating-keys#create_a_key)
 * 
 * ## Example Usage
 * ### Kms Crypto Key Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
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
 *         var keyring = new KeyRing(&#34;keyring&#34;, KeyRingArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .build());
 * 
 *         var example_key = new CryptoKey(&#34;example-key&#34;, CryptoKeyArgs.builder()        
 *             .keyRing(keyring.id())
 *             .rotationPeriod(&#34;7776000s&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Kms Crypto Key Asymmetric Sign
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.kms.inputs.CryptoKeyVersionTemplateArgs;
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
 *         var keyring = new KeyRing(&#34;keyring&#34;, KeyRingArgs.builder()        
 *             .location(&#34;global&#34;)
 *             .build());
 * 
 *         var example_asymmetric_sign_key = new CryptoKey(&#34;example-asymmetric-sign-key&#34;, CryptoKeyArgs.builder()        
 *             .keyRing(keyring.id())
 *             .purpose(&#34;ASYMMETRIC_SIGN&#34;)
 *             .versionTemplate(CryptoKeyVersionTemplateArgs.builder()
 *                 .algorithm(&#34;EC_SIGN_P384_SHA384&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * CryptoKey can be imported using any of these accepted formats* `{{key_ring}}/cryptoKeys/{{name}}` * `{{key_ring}}/{{name}}` When using the `pulumi import` command, CryptoKey can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/cryptoKeys/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:kms/cryptoKey:CryptoKey default {{key_ring}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:kms/cryptoKey:CryptoKey")
public class CryptoKey extends com.pulumi.resources.CustomResource {
    /**
     * The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
     * If not specified at creation time, the default duration is 24 hours.
     * 
     */
    @Export(name="destroyScheduledDuration", refs={String.class}, tree="[0]")
    private Output<String> destroyScheduledDuration;

    /**
     * @return The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
     * If not specified at creation time, the default duration is 24 hours.
     * 
     */
    public Output<String> destroyScheduledDuration() {
        return this.destroyScheduledDuration;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Whether this key may contain imported versions only.
     * 
     */
    @Export(name="importOnly", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> importOnly;

    /**
     * @return Whether this key may contain imported versions only.
     * 
     */
    public Output<Boolean> importOnly() {
        return this.importOnly;
    }
    /**
     * The KeyRing that this key belongs to.
     * Format: `&#39;projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}&#39;`.
     * 
     * ***
     * 
     */
    @Export(name="keyRing", refs={String.class}, tree="[0]")
    private Output<String> keyRing;

    /**
     * @return The KeyRing that this key belongs to.
     * Format: `&#39;projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}&#39;`.
     * 
     * ***
     * 
     */
    public Output<String> keyRing() {
        return this.keyRing;
    }
    /**
     * Labels with user-defined metadata to apply to this resource.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels with user-defined metadata to apply to this resource.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The resource name for the CryptoKey.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name for the CryptoKey.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A copy of the primary CryptoKeyVersion that will be used by cryptoKeys.encrypt when this CryptoKey is given in EncryptRequest.name.
     * Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys, this field will be unset.
     * Structure is documented below.
     * 
     */
    @Export(name="primaries", refs={List.class,CryptoKeyPrimary.class}, tree="[0,1]")
    private Output<List<CryptoKeyPrimary>> primaries;

    /**
     * @return A copy of the primary CryptoKeyVersion that will be used by cryptoKeys.encrypt when this CryptoKey is given in EncryptRequest.name.
     * Keys with purpose ENCRYPT_DECRYPT may have a primary. For other keys, this field will be unset.
     * Structure is documented below.
     * 
     */
    public Output<List<CryptoKeyPrimary>> primaries() {
        return this.primaries;
    }
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * The immutable purpose of this CryptoKey. See the
     * [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
     * for possible inputs.
     * Default value is &#34;ENCRYPT_DECRYPT&#34;.
     * 
     */
    @Export(name="purpose", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> purpose;

    /**
     * @return The immutable purpose of this CryptoKey. See the
     * [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose)
     * for possible inputs.
     * Default value is &#34;ENCRYPT_DECRYPT&#34;.
     * 
     */
    public Output<Optional<String>> purpose() {
        return Codegen.optional(this.purpose);
    }
    /**
     * Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
     * The first rotation will take place after the specified period. The rotation period has
     * the format of a decimal number with up to 9 fractional digits, followed by the
     * letter `s` (seconds). It must be greater than a day (ie, 86400).
     * 
     */
    @Export(name="rotationPeriod", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rotationPeriod;

    /**
     * @return Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
     * The first rotation will take place after the specified period. The rotation period has
     * the format of a decimal number with up to 9 fractional digits, followed by the
     * letter `s` (seconds). It must be greater than a day (ie, 86400).
     * 
     */
    public Output<Optional<String>> rotationPeriod() {
        return Codegen.optional(this.rotationPeriod);
    }
    /**
     * If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
     * You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
     * 
     */
    @Export(name="skipInitialVersionCreation", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> skipInitialVersionCreation;

    /**
     * @return If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
     * You must use the `gcp.kms.KeyRingImportJob` resource to import the CryptoKeyVersion.
     * 
     */
    public Output<Optional<Boolean>> skipInitialVersionCreation() {
        return Codegen.optional(this.skipInitialVersionCreation);
    }
    /**
     * A template describing settings for new crypto key versions.
     * Structure is documented below.
     * 
     */
    @Export(name="versionTemplate", refs={CryptoKeyVersionTemplate.class}, tree="[0]")
    private Output<CryptoKeyVersionTemplate> versionTemplate;

    /**
     * @return A template describing settings for new crypto key versions.
     * Structure is documented below.
     * 
     */
    public Output<CryptoKeyVersionTemplate> versionTemplate() {
        return this.versionTemplate;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CryptoKey(String name) {
        this(name, CryptoKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CryptoKey(String name, CryptoKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CryptoKey(String name, CryptoKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/cryptoKey:CryptoKey", name, args == null ? CryptoKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CryptoKey(String name, Output<String> id, @Nullable CryptoKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/cryptoKey:CryptoKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
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
    public static CryptoKey get(String name, Output<String> id, @Nullable CryptoKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CryptoKey(name, id, state, options);
    }
}
