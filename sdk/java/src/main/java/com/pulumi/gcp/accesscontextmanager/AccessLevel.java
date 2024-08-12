// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accesscontextmanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.accesscontextmanager.AccessLevelArgs;
import com.pulumi.gcp.accesscontextmanager.inputs.AccessLevelState;
import com.pulumi.gcp.accesscontextmanager.outputs.AccessLevelBasic;
import com.pulumi.gcp.accesscontextmanager.outputs.AccessLevelCustom;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An AccessLevel is a label that can be applied to requests to GCP services,
 * along with a list of requirements necessary for the label to be applied.
 * 
 * To get more information about AccessLevel, see:
 * 
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
 * * How-to Guides
 *     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
 * 
 * &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
 * you must specify a `billing_project` and set `user_project_override` to true
 * in the provider configuration. Otherwise the ACM API will return a 403 error.
 * Your account must have the `serviceusage.services.use` permission on the
 * `billing_project` you defined.
 * 
 * ## Example Usage
 * 
 * ### Access Context Manager Access Level Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicy;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicyArgs;
 * import com.pulumi.gcp.accesscontextmanager.AccessLevel;
 * import com.pulumi.gcp.accesscontextmanager.AccessLevelArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.AccessLevelBasicArgs;
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
 *         var access_policy = new AccessPolicy("access-policy", AccessPolicyArgs.builder()
 *             .parent("organizations/123456789")
 *             .title("my policy")
 *             .build());
 * 
 *         var access_level = new AccessLevel("access-level", AccessLevelArgs.builder()
 *             .parent(access_policy.name().applyValue(name -> String.format("accessPolicies/%s", name)))
 *             .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/accessLevels/chromeos_no_lock", name)))
 *             .title("chromeos_no_lock")
 *             .basic(AccessLevelBasicArgs.builder()
 *                 .conditions(AccessLevelBasicConditionArgs.builder()
 *                     .devicePolicy(AccessLevelBasicConditionDevicePolicyArgs.builder()
 *                         .requireScreenLock(true)
 *                         .osConstraints(AccessLevelBasicConditionDevicePolicyOsConstraintArgs.builder()
 *                             .osType("DESKTOP_CHROME_OS")
 *                             .build())
 *                         .build())
 *                     .regions(                    
 *                         "CH",
 *                         "IT",
 *                         "US")
 *                     .build())
 *                 .build())
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
 * AccessLevel can be imported using any of these accepted formats:
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, AccessLevel can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:accesscontextmanager/accessLevel:AccessLevel default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:accesscontextmanager/accessLevel:AccessLevel")
public class AccessLevel extends com.pulumi.resources.CustomResource {
    /**
     * A set of predefined conditions for the access level and a combining function.
     * Structure is documented below.
     * 
     */
    @Export(name="basic", refs={AccessLevelBasic.class}, tree="[0]")
    private Output</* @Nullable */ AccessLevelBasic> basic;

    /**
     * @return A set of predefined conditions for the access level and a combining function.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AccessLevelBasic>> basic() {
        return Codegen.optional(this.basic);
    }
    /**
     * Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
     * See CEL spec at: https://github.com/google/cel-spec.
     * Structure is documented below.
     * 
     */
    @Export(name="custom", refs={AccessLevelCustom.class}, tree="[0]")
    private Output</* @Nullable */ AccessLevelCustom> custom;

    /**
     * @return Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
     * See CEL spec at: https://github.com/google/cel-spec.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AccessLevelCustom>> custom() {
        return Codegen.optional(this.custom);
    }
    /**
     * Description of the AccessLevel and its use. Does not affect behavior.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the AccessLevel and its use. Does not affect behavior.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Resource name for the Access Level. The short_name component must begin
     * with a letter and only include alphanumeric and &#39;_&#39;.
     * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Resource name for the Access Level. The short_name component must begin
     * with a letter and only include alphanumeric and &#39;_&#39;.
     * Format: accessPolicies/{policy_id}/accessLevels/{short_name}
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The AccessPolicy this AccessLevel lives in.
     * Format: accessPolicies/{policy_id}
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The AccessPolicy this AccessLevel lives in.
     * Format: accessPolicies/{policy_id}
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Human readable title. Must be unique within the Policy.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Human readable title. Must be unique within the Policy.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccessLevel(java.lang.String name) {
        this(name, AccessLevelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccessLevel(java.lang.String name, AccessLevelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccessLevel(java.lang.String name, AccessLevelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:accesscontextmanager/accessLevel:AccessLevel", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AccessLevel(java.lang.String name, Output<java.lang.String> id, @Nullable AccessLevelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:accesscontextmanager/accessLevel:AccessLevel", name, state, makeResourceOptions(options, id), false);
    }

    private static AccessLevelArgs makeArgs(AccessLevelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AccessLevelArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static AccessLevel get(java.lang.String name, Output<java.lang.String> id, @Nullable AccessLevelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccessLevel(name, id, state, options);
    }
}
