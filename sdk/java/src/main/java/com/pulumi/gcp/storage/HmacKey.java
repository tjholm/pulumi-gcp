// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.storage.HmacKeyArgs;
import com.pulumi.gcp.storage.inputs.HmacKeyState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The hmacKeys resource represents an HMAC key within Cloud Storage. The resource
 * consists of a secret and HMAC key metadata. HMAC keys can be used as credentials
 * for service accounts.
 * 
 * To get more information about HmacKey, see:
 * 
 * * [API documentation](https://cloud.google.com/storage/docs/json_api/v1/projects/hmacKeys)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/storage/docs/authentication/managing-hmackeys)
 * 
 * ## Example Usage
 * 
 * ### Storage Hmac Key
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.storage.HmacKey;
 * import com.pulumi.gcp.storage.HmacKeyArgs;
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
 *         // Create a new service account
 *         var serviceAccount = new Account("serviceAccount", AccountArgs.builder()
 *             .accountId("my-svc-acc")
 *             .build());
 * 
 *         //Create the HMAC key for the associated service account
 *         var key = new HmacKey("key", HmacKeyArgs.builder()
 *             .serviceAccountEmail(serviceAccount.email())
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
 * HmacKey can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/hmacKeys/{{access_id}}`
 * 
 * * `{{project}}/{{access_id}}`
 * 
 * * `{{access_id}}`
 * 
 * When using the `pulumi import` command, HmacKey can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:storage/hmacKey:HmacKey default projects/{{project}}/hmacKeys/{{access_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:storage/hmacKey:HmacKey default {{project}}/{{access_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:storage/hmacKey:HmacKey default {{access_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:storage/hmacKey:HmacKey")
public class HmacKey extends com.pulumi.resources.CustomResource {
    /**
     * The access ID of the HMAC Key.
     * 
     */
    @Export(name="accessId", refs={String.class}, tree="[0]")
    private Output<String> accessId;

    /**
     * @return The access ID of the HMAC Key.
     * 
     */
    public Output<String> accessId() {
        return this.accessId;
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
     * HMAC secret key material.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return HMAC secret key material.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }
    /**
     * The email address of the key&#39;s associated service account.
     * 
     * ***
     * 
     */
    @Export(name="serviceAccountEmail", refs={String.class}, tree="[0]")
    private Output<String> serviceAccountEmail;

    /**
     * @return The email address of the key&#39;s associated service account.
     * 
     * ***
     * 
     */
    public Output<String> serviceAccountEmail() {
        return this.serviceAccountEmail;
    }
    /**
     * The state of the key. Can be set to one of ACTIVE, INACTIVE.
     * Default value is `ACTIVE`.
     * Possible values are: `ACTIVE`, `INACTIVE`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state of the key. Can be set to one of ACTIVE, INACTIVE.
     * Default value is `ACTIVE`.
     * Possible values are: `ACTIVE`, `INACTIVE`.
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }
    /**
     * &#39;The creation time of the HMAC key in RFC 3339 format. &#39;
     * 
     */
    @Export(name="timeCreated", refs={String.class}, tree="[0]")
    private Output<String> timeCreated;

    /**
     * @return &#39;The creation time of the HMAC key in RFC 3339 format. &#39;
     * 
     */
    public Output<String> timeCreated() {
        return this.timeCreated;
    }
    /**
     * &#39;The last modification time of the HMAC key metadata in RFC 3339 format.&#39;
     * 
     */
    @Export(name="updated", refs={String.class}, tree="[0]")
    private Output<String> updated;

    /**
     * @return &#39;The last modification time of the HMAC key metadata in RFC 3339 format.&#39;
     * 
     */
    public Output<String> updated() {
        return this.updated;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HmacKey(String name) {
        this(name, HmacKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HmacKey(String name, HmacKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HmacKey(String name, HmacKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:storage/hmacKey:HmacKey", name, args == null ? HmacKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HmacKey(String name, Output<String> id, @Nullable HmacKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:storage/hmacKey:HmacKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secret"
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
    public static HmacKey get(String name, Output<String> id, @Nullable HmacKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HmacKey(name, id, state, options);
    }
}
