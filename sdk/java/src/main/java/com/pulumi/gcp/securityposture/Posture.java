// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.securityposture;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.securityposture.PostureArgs;
import com.pulumi.gcp.securityposture.inputs.PostureState;
import com.pulumi.gcp.securityposture.outputs.PosturePolicySet;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Posture represents a collection of policy set including its name, state, description
 * and policy sets. A policy set includes set of policies along with their definition.
 * A posture can be created at the organization level.
 * Every update to a deployed posture creates a new posture revision with an updated revision_id.
 * 
 * To get more information about Posture, see:
 * 
 * * How-to Guides
 *     * [Create and deploy a posture](https://cloud.google.com/security-command-center/docs/how-to-use-security-posture)
 * 
 * ## Example Usage
 * 
 * ### Securityposture Posture Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securityposture.Posture;
 * import com.pulumi.gcp.securityposture.PostureArgs;
 * import com.pulumi.gcp.securityposture.inputs.PosturePolicySetArgs;
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
 *         var posture1 = new Posture("posture1", PostureArgs.builder()
 *             .postureId("posture_example")
 *             .parent("organizations/123456789")
 *             .location("global")
 *             .state("ACTIVE")
 *             .description("a new posture")
 *             .policySets(            
 *                 PosturePolicySetArgs.builder()
 *                     .policySetId("org_policy_set")
 *                     .description("set of org policies")
 *                     .policies(                    
 *                         PosturePolicySetPolicyArgs.builder()
 *                             .policyId("canned_org_policy")
 *                             .constraint(PosturePolicySetPolicyConstraintArgs.builder()
 *                                 .orgPolicyConstraint(PosturePolicySetPolicyConstraintOrgPolicyConstraintArgs.builder()
 *                                     .cannedConstraintId("storage.uniformBucketLevelAccess")
 *                                     .policyRules(PosturePolicySetPolicyConstraintOrgPolicyConstraintPolicyRuleArgs.builder()
 *                                         .enforce(true)
 *                                         .condition(PosturePolicySetPolicyConstraintOrgPolicyConstraintPolicyRuleConditionArgs.builder()
 *                                             .description("condition description")
 *                                             .expression("resource.matchTag('org_id/tag_key_short_name,'tag_value_short_name')")
 *                                             .title("a CEL condition")
 *                                             .build())
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build(),
 *                         PosturePolicySetPolicyArgs.builder()
 *                             .policyId("custom_org_policy")
 *                             .constraint(PosturePolicySetPolicyConstraintArgs.builder()
 *                                 .orgPolicyConstraintCustom(PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomArgs.builder()
 *                                     .customConstraint(PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomCustomConstraintArgs.builder()
 *                                         .name("organizations/123456789/customConstraints/custom.disableGkeAutoUpgrade")
 *                                         .displayName("Disable GKE auto upgrade")
 *                                         .description("Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced.")
 *                                         .actionType("ALLOW")
 *                                         .condition("resource.management.autoUpgrade == false")
 *                                         .methodTypes(                                        
 *                                             "CREATE",
 *                                             "UPDATE")
 *                                         .resourceTypes("container.googleapis.com/NodePool")
 *                                         .build())
 *                                     .policyRules(PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRuleArgs.builder()
 *                                         .enforce(true)
 *                                         .condition(PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomPolicyRuleConditionArgs.builder()
 *                                             .description("condition description")
 *                                             .expression("resource.matchTagId('tagKeys/key_id','tagValues/value_id')")
 *                                             .title("a CEL condition")
 *                                             .build())
 *                                         .build())
 *                                     .build())
 *                                 .build())
 *                             .build())
 *                     .build(),
 *                 PosturePolicySetArgs.builder()
 *                     .policySetId("sha_policy_set")
 *                     .description("set of sha policies")
 *                     .policies(                    
 *                         PosturePolicySetPolicyArgs.builder()
 *                             .policyId("sha_builtin_module")
 *                             .constraint(PosturePolicySetPolicyConstraintArgs.builder()
 *                                 .securityHealthAnalyticsModule(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsModuleArgs.builder()
 *                                     .moduleName("BIGQUERY_TABLE_CMEK_DISABLED")
 *                                     .moduleEnablementState("ENABLED")
 *                                     .build())
 *                                 .build())
 *                             .description("enable BIGQUERY_TABLE_CMEK_DISABLED")
 *                             .build(),
 *                         PosturePolicySetPolicyArgs.builder()
 *                             .policyId("sha_custom_module")
 *                             .constraint(PosturePolicySetPolicyConstraintArgs.builder()
 *                                 .securityHealthAnalyticsCustomModule(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleArgs.builder()
 *                                     .displayName("custom_SHA_policy")
 *                                     .config(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigArgs.builder()
 *                                         .predicate(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigPredicateArgs.builder()
 *                                             .expression("resource.rotationPeriod > duration('2592000s')")
 *                                             .build())
 *                                         .customOutput(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputArgs.builder()
 *                                             .properties(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputPropertyArgs.builder()
 *                                                 .name("duration")
 *                                                 .valueExpression(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputPropertyValueExpressionArgs.builder()
 *                                                     .expression("resource.rotationPeriod")
 *                                                     .build())
 *                                                 .build())
 *                                             .build())
 *                                         .resourceSelector(PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigResourceSelectorArgs.builder()
 *                                             .resourceTypes("cloudkms.googleapis.com/CryptoKey")
 *                                             .build())
 *                                         .severity("LOW")
 *                                         .description("Custom Module")
 *                                         .recommendation("Testing custom modules")
 *                                         .build())
 *                                     .moduleEnablementState("ENABLED")
 *                                     .build())
 *                                 .build())
 *                             .build())
 *                     .build())
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
 * Posture can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/locations/{{location}}/postures/{{posture_id}}`
 * 
 * When using the `pulumi import` command, Posture can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:securityposture/posture:Posture default {{parent}}/locations/{{location}}/postures/{{posture_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:securityposture/posture:Posture")
public class Posture extends com.pulumi.resources.CustomResource {
    /**
     * Time the Posture was created in UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time the Posture was created in UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the posture.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the posture.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Location of the resource, eg: global.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Location of the resource, eg: global.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Name of the posture.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the posture.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * List of policy sets for the posture.
     * Structure is documented below.
     * 
     */
    @Export(name="policySets", refs={List.class,PosturePolicySet.class}, tree="[0,1]")
    private Output<List<PosturePolicySet>> policySets;

    /**
     * @return List of policy sets for the posture.
     * Structure is documented below.
     * 
     */
    public Output<List<PosturePolicySet>> policySets() {
        return this.policySets;
    }
    /**
     * Id of the posture. It is an immutable field.
     * 
     */
    @Export(name="postureId", refs={String.class}, tree="[0]")
    private Output<String> postureId;

    /**
     * @return Id of the posture. It is an immutable field.
     * 
     */
    public Output<String> postureId() {
        return this.postureId;
    }
    /**
     * If set, there are currently changes in flight to the posture.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return If set, there are currently changes in flight to the posture.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * Revision_id of the posture.
     * 
     */
    @Export(name="revisionId", refs={String.class}, tree="[0]")
    private Output<String> revisionId;

    /**
     * @return Revision_id of the posture.
     * 
     */
    public Output<String> revisionId() {
        return this.revisionId;
    }
    /**
     * State of the posture. Update to state field should not be triggered along with
     * with other field updates.
     * Possible values are: `DEPRECATED`, `DRAFT`, `ACTIVE`.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the posture. Update to state field should not be triggered along with
     * with other field updates.
     * Possible values are: `DEPRECATED`, `DRAFT`, `ACTIVE`.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Time the Posture was updated in UTC.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time the Posture was updated in UTC.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Posture(java.lang.String name) {
        this(name, PostureArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Posture(java.lang.String name, PostureArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Posture(java.lang.String name, PostureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securityposture/posture:Posture", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Posture(java.lang.String name, Output<java.lang.String> id, @Nullable PostureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securityposture/posture:Posture", name, state, makeResourceOptions(options, id), false);
    }

    private static PostureArgs makeArgs(PostureArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PostureArgs.Empty : args;
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
    public static Posture get(java.lang.String name, Output<java.lang.String> id, @Nullable PostureState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Posture(name, id, state, options);
    }
}
