// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iam;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iam.DenyPolicyArgs;
import com.pulumi.gcp.iam.inputs.DenyPolicyState;
import com.pulumi.gcp.iam.outputs.DenyPolicyRule;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a collection of denial policies to apply to a given resource.
 * 
 * To get more information about DenyPolicy, see:
 * 
 * * [API documentation](https://cloud.google.com/iam/docs/reference/rest/v2/policies)
 * * How-to Guides
 *     * [Permissions supported in deny policies](https://cloud.google.com/iam/docs/deny-permissions-support)
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * DenyPolicy can be imported using any of these accepted formats* `{{parent}}/{{name}}` When using the `pulumi import` command, DenyPolicy can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:iam/denyPolicy:DenyPolicy default {{parent}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:iam/denyPolicy:DenyPolicy")
public class DenyPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The display name of the rule.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name of the rule.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * The hash of the resource. Used internally during updates.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return The hash of the resource. Used internally during updates.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The name of the policy.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the policy.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The attachment point is identified by its URL-encoded full resource name.
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The attachment point is identified by its URL-encoded full resource name.
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Rules to be applied.
     * Structure is documented below.
     * 
     */
    @Export(name="rules", refs={List.class,DenyPolicyRule.class}, tree="[0,1]")
    private Output<List<DenyPolicyRule>> rules;

    /**
     * @return Rules to be applied.
     * Structure is documented below.
     * 
     */
    public Output<List<DenyPolicyRule>> rules() {
        return this.rules;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DenyPolicy(String name) {
        this(name, DenyPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DenyPolicy(String name, DenyPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DenyPolicy(String name, DenyPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iam/denyPolicy:DenyPolicy", name, args == null ? DenyPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DenyPolicy(String name, Output<String> id, @Nullable DenyPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iam/denyPolicy:DenyPolicy", name, state, makeResourceOptions(options, id));
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
    public static DenyPolicy get(String name, Output<String> id, @Nullable DenyPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DenyPolicy(name, id, state, options);
    }
}
