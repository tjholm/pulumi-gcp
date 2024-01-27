// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkservices;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.networkservices.EdgeCacheKeysetArgs;
import com.pulumi.gcp.networkservices.inputs.EdgeCacheKeysetState;
import com.pulumi.gcp.networkservices.outputs.EdgeCacheKeysetPublicKey;
import com.pulumi.gcp.networkservices.outputs.EdgeCacheKeysetValidationSharedKey;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * EdgeCacheKeyset represents a collection of public keys used for validating signed requests.
 * 
 * To get more information about EdgeCacheKeyset, see:
 * 
 * * [API documentation](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheKeysets)
 * * How-to Guides
 *     * [Create keysets](https://cloud.google.com/media-cdn/docs/create-keyset)
 * 
 * ## Example Usage
 * ### Network Services Edge Cache Keyset Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.networkservices.EdgeCacheKeyset;
 * import com.pulumi.gcp.networkservices.EdgeCacheKeysetArgs;
 * import com.pulumi.gcp.networkservices.inputs.EdgeCacheKeysetPublicKeyArgs;
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
 *         var default_ = new EdgeCacheKeyset(&#34;default&#34;, EdgeCacheKeysetArgs.builder()        
 *             .description(&#34;The default keyset&#34;)
 *             .publicKeys(            
 *                 EdgeCacheKeysetPublicKeyArgs.builder()
 *                     .id(&#34;my-public-key&#34;)
 *                     .value(&#34;FHsTyFHNmvNpw4o7-rp-M1yqMyBF8vXSBRkZtkQ0RKY&#34;)
 *                     .build(),
 *                 EdgeCacheKeysetPublicKeyArgs.builder()
 *                     .id(&#34;my-public-key-2&#34;)
 *                     .value(&#34;hzd03llxB1u5FOLKFkZ6_wCJqC7jtN0bg7xlBqS6WVM&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Network Services Edge Cache Keyset Dual Token
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.secretmanager.Secret;
 * import com.pulumi.gcp.secretmanager.SecretArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationArgs;
 * import com.pulumi.gcp.secretmanager.inputs.SecretReplicationAutoArgs;
 * import com.pulumi.gcp.secretmanager.SecretVersion;
 * import com.pulumi.gcp.secretmanager.SecretVersionArgs;
 * import com.pulumi.gcp.networkservices.EdgeCacheKeyset;
 * import com.pulumi.gcp.networkservices.EdgeCacheKeysetArgs;
 * import com.pulumi.gcp.networkservices.inputs.EdgeCacheKeysetPublicKeyArgs;
 * import com.pulumi.gcp.networkservices.inputs.EdgeCacheKeysetValidationSharedKeyArgs;
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
 *         var secret_basic = new Secret(&#34;secret-basic&#34;, SecretArgs.builder()        
 *             .secretId(&#34;secret-name&#34;)
 *             .replication(SecretReplicationArgs.builder()
 *                 .auto()
 *                 .build())
 *             .build());
 * 
 *         var secret_version_basic = new SecretVersion(&#34;secret-version-basic&#34;, SecretVersionArgs.builder()        
 *             .secret(secret_basic.id())
 *             .secretData(&#34;secret-data&#34;)
 *             .build());
 * 
 *         var default_ = new EdgeCacheKeyset(&#34;default&#34;, EdgeCacheKeysetArgs.builder()        
 *             .description(&#34;The default keyset&#34;)
 *             .publicKeys(EdgeCacheKeysetPublicKeyArgs.builder()
 *                 .id(&#34;my-public-key&#34;)
 *                 .managed(true)
 *                 .build())
 *             .validationSharedKeys(EdgeCacheKeysetValidationSharedKeyArgs.builder()
 *                 .secretVersion(secret_version_basic.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * EdgeCacheKeyset can be imported using any of these accepted formats* `projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, EdgeCacheKeyset can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default projects/{{project}}/locations/global/edgeCacheKeysets/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset")
public class EdgeCacheKeyset extends com.pulumi.resources.CustomResource {
    /**
     * A human-readable description of the resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-readable description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Set of label tags associated with the EdgeCache resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Set of label tags associated with the EdgeCache resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource; provided by the client when the resource is created.
     * The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
     * and all following characters must be a dash, underscore, letter or digit.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * An ordered list of Ed25519 public keys to use for validating signed requests.
     * You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
     * You may specify no more than one Google-managed public key.
     * If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
     * Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
     * Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
     * Structure is documented below.
     * 
     */
    @Export(name="publicKeys", refs={List.class,EdgeCacheKeysetPublicKey.class}, tree="[0,1]")
    private Output</* @Nullable */ List<EdgeCacheKeysetPublicKey>> publicKeys;

    /**
     * @return An ordered list of Ed25519 public keys to use for validating signed requests.
     * You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
     * You may specify no more than one Google-managed public key.
     * If you specify `public_keys`, you must specify at least one (1) key and may specify up to three (3) keys.
     * Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your corresponding private key.
     * Ensure that the private key is kept secret, and that only authorized users can add public keys to a keyset.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<EdgeCacheKeysetPublicKey>>> publicKeys() {
        return Codegen.optional(this.publicKeys);
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
     * An ordered list of shared keys to use for validating signed requests.
     * Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
     * You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
     * You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
     * Structure is documented below.
     * 
     */
    @Export(name="validationSharedKeys", refs={List.class,EdgeCacheKeysetValidationSharedKey.class}, tree="[0,1]")
    private Output</* @Nullable */ List<EdgeCacheKeysetValidationSharedKey>> validationSharedKeys;

    /**
     * @return An ordered list of shared keys to use for validating signed requests.
     * Shared keys are secret.  Ensure that only authorized users can add `validation_shared_keys` to a keyset.
     * You can rotate keys by appending (pushing) a new key to the list of `validation_shared_keys` and removing any superseded keys.
     * You must specify `public_keys` or `validation_shared_keys` (or both). The keys in `public_keys` are checked first.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<EdgeCacheKeysetValidationSharedKey>>> validationSharedKeys() {
        return Codegen.optional(this.validationSharedKeys);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EdgeCacheKeyset(String name) {
        this(name, EdgeCacheKeysetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EdgeCacheKeyset(String name, @Nullable EdgeCacheKeysetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EdgeCacheKeyset(String name, @Nullable EdgeCacheKeysetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset", name, args == null ? EdgeCacheKeysetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EdgeCacheKeyset(String name, Output<String> id, @Nullable EdgeCacheKeysetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:networkservices/edgeCacheKeyset:EdgeCacheKeyset", name, state, makeResourceOptions(options, id));
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
    public static EdgeCacheKeyset get(String name, Output<String> id, @Nullable EdgeCacheKeysetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EdgeCacheKeyset(name, id, state, options);
    }
}
