// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.logging;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.logging.BillingAccountExclusionArgs;
import com.pulumi.gcp.logging.inputs.BillingAccountExclusionState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.logging.BillingAccountExclusion;
 * import com.pulumi.gcp.logging.BillingAccountExclusionArgs;
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
 *         var my_exclusion = new BillingAccountExclusion("my-exclusion", BillingAccountExclusionArgs.builder()
 *             .name("my-instance-debug-exclusion")
 *             .billingAccount("ABCDEF-012345-GHIJKL")
 *             .description("Exclude GCE instance debug logs")
 *             .filter("resource.type = gce_instance AND severity <= DEBUG")
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
 * Billing account logging exclusions can be imported using their URI, e.g.
 * 
 * * `billingAccounts/{{billing_account}}/exclusions/{{name}}`
 * 
 * When using the `pulumi import` command, billing account logging exclusions can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:logging/billingAccountExclusion:BillingAccountExclusion default billingAccounts/{{billing_account}}/exclusions/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:logging/billingAccountExclusion:BillingAccountExclusion")
public class BillingAccountExclusion extends com.pulumi.resources.CustomResource {
    /**
     * The billing account to create the exclusion for.
     * 
     */
    @Export(name="billingAccount", refs={String.class}, tree="[0]")
    private Output<String> billingAccount;

    /**
     * @return The billing account to create the exclusion for.
     * 
     */
    public Output<String> billingAccount() {
        return this.billingAccount;
    }
    /**
     * A human-readable description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-readable description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     * 
     */
    @Export(name="disabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disabled;

    /**
     * @return Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     * 
     */
    public Output<Optional<Boolean>> disabled() {
        return Codegen.optional(this.disabled);
    }
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     * 
     */
    @Export(name="filter", refs={String.class}, tree="[0]")
    private Output<String> filter;

    /**
     * @return The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     * 
     */
    public Output<String> filter() {
        return this.filter;
    }
    /**
     * The name of the logging exclusion.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the logging exclusion.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BillingAccountExclusion(String name) {
        this(name, BillingAccountExclusionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BillingAccountExclusion(String name, BillingAccountExclusionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BillingAccountExclusion(String name, BillingAccountExclusionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, args == null ? BillingAccountExclusionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BillingAccountExclusion(String name, Output<String> id, @Nullable BillingAccountExclusionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:logging/billingAccountExclusion:BillingAccountExclusion", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static BillingAccountExclusion get(String name, Output<String> id, @Nullable BillingAccountExclusionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BillingAccountExclusion(name, id, state, options);
    }
}
