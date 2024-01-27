// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.binaryauthorization;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.binaryauthorization.PolicyArgs;
import com.pulumi.gcp.binaryauthorization.inputs.PolicyState;
import com.pulumi.gcp.binaryauthorization.outputs.PolicyAdmissionWhitelistPattern;
import com.pulumi.gcp.binaryauthorization.outputs.PolicyClusterAdmissionRule;
import com.pulumi.gcp.binaryauthorization.outputs.PolicyDefaultAdmissionRule;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A policy for container image binary authorization.
 * 
 * To get more information about Policy, see:
 * 
 * * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/binary-authorization/)
 * 
 * ## Example Usage
 * ### Binary Authorization Policy Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.containeranalysis.Note;
 * import com.pulumi.gcp.containeranalysis.NoteArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityHintArgs;
 * import com.pulumi.gcp.binaryauthorization.Attestor;
 * import com.pulumi.gcp.binaryauthorization.AttestorArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.AttestorAttestationAuthorityNoteArgs;
 * import com.pulumi.gcp.binaryauthorization.Policy;
 * import com.pulumi.gcp.binaryauthorization.PolicyArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.PolicyAdmissionWhitelistPatternArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.PolicyDefaultAdmissionRuleArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.PolicyClusterAdmissionRuleArgs;
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
 *         var note = new Note(&#34;note&#34;, NoteArgs.builder()        
 *             .attestationAuthority(NoteAttestationAuthorityArgs.builder()
 *                 .hint(NoteAttestationAuthorityHintArgs.builder()
 *                     .humanReadableName(&#34;My attestor&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var attestor = new Attestor(&#34;attestor&#34;, AttestorArgs.builder()        
 *             .attestationAuthorityNote(AttestorAttestationAuthorityNoteArgs.builder()
 *                 .noteReference(note.name())
 *                 .build())
 *             .build());
 * 
 *         var policy = new Policy(&#34;policy&#34;, PolicyArgs.builder()        
 *             .admissionWhitelistPatterns(PolicyAdmissionWhitelistPatternArgs.builder()
 *                 .namePattern(&#34;gcr.io/google_containers/*&#34;)
 *                 .build())
 *             .defaultAdmissionRule(PolicyDefaultAdmissionRuleArgs.builder()
 *                 .evaluationMode(&#34;ALWAYS_ALLOW&#34;)
 *                 .enforcementMode(&#34;ENFORCED_BLOCK_AND_AUDIT_LOG&#34;)
 *                 .build())
 *             .clusterAdmissionRules(PolicyClusterAdmissionRuleArgs.builder()
 *                 .cluster(&#34;us-central1-a.prod-cluster&#34;)
 *                 .evaluationMode(&#34;REQUIRE_ATTESTATION&#34;)
 *                 .enforcementMode(&#34;ENFORCED_BLOCK_AND_AUDIT_LOG&#34;)
 *                 .requireAttestationsBies(attestor.name())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Binary Authorization Policy Global Evaluation
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.containeranalysis.Note;
 * import com.pulumi.gcp.containeranalysis.NoteArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityArgs;
 * import com.pulumi.gcp.containeranalysis.inputs.NoteAttestationAuthorityHintArgs;
 * import com.pulumi.gcp.binaryauthorization.Attestor;
 * import com.pulumi.gcp.binaryauthorization.AttestorArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.AttestorAttestationAuthorityNoteArgs;
 * import com.pulumi.gcp.binaryauthorization.Policy;
 * import com.pulumi.gcp.binaryauthorization.PolicyArgs;
 * import com.pulumi.gcp.binaryauthorization.inputs.PolicyDefaultAdmissionRuleArgs;
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
 *         var note = new Note(&#34;note&#34;, NoteArgs.builder()        
 *             .attestationAuthority(NoteAttestationAuthorityArgs.builder()
 *                 .hint(NoteAttestationAuthorityHintArgs.builder()
 *                     .humanReadableName(&#34;My attestor&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var attestor = new Attestor(&#34;attestor&#34;, AttestorArgs.builder()        
 *             .attestationAuthorityNote(AttestorAttestationAuthorityNoteArgs.builder()
 *                 .noteReference(note.name())
 *                 .build())
 *             .build());
 * 
 *         var policy = new Policy(&#34;policy&#34;, PolicyArgs.builder()        
 *             .defaultAdmissionRule(PolicyDefaultAdmissionRuleArgs.builder()
 *                 .evaluationMode(&#34;REQUIRE_ATTESTATION&#34;)
 *                 .enforcementMode(&#34;ENFORCED_BLOCK_AND_AUDIT_LOG&#34;)
 *                 .requireAttestationsBies(attestor.name())
 *                 .build())
 *             .globalPolicyEvaluationMode(&#34;ENABLE&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Policy can be imported using any of these accepted formats* `projects/{{project}}` * `{{project}}` When using the `pulumi import` command, Policy can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/policy:Policy default projects/{{project}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:binaryauthorization/policy:Policy default {{project}}
 * ```
 * 
 */
@ResourceType(type="gcp:binaryauthorization/policy:Policy")
public class Policy extends com.pulumi.resources.CustomResource {
    /**
     * A whitelist of image patterns to exclude from admission rules. If an
     * image&#39;s name matches a whitelist pattern, the image&#39;s admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     * 
     */
    @Export(name="admissionWhitelistPatterns", refs={List.class,PolicyAdmissionWhitelistPattern.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PolicyAdmissionWhitelistPattern>> admissionWhitelistPatterns;

    /**
     * @return A whitelist of image patterns to exclude from admission rules. If an
     * image&#39;s name matches a whitelist pattern, the image&#39;s admission
     * requests will always be permitted regardless of your admission rules.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<PolicyAdmissionWhitelistPattern>>> admissionWhitelistPatterns() {
        return Codegen.optional(this.admissionWhitelistPatterns);
    }
    /**
     * Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     * 
     * Identifier format: `{{location}}.{{clusterId}}`.
     * A location is either a compute zone (e.g. `us-central1-a`) or a region
     * (e.g. `us-central1`).
     * Structure is documented below.
     * 
     */
    @Export(name="clusterAdmissionRules", refs={List.class,PolicyClusterAdmissionRule.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PolicyClusterAdmissionRule>> clusterAdmissionRules;

    /**
     * @return Per-cluster admission rules. An admission rule specifies either that
     * all container images used in a pod creation request must be attested
     * to by one or more attestors, that all pod creations will be allowed,
     * or that all pod creations will be denied. There can be at most one
     * admission rule per cluster spec.
     * 
     * Identifier format: `{{location}}.{{clusterId}}`.
     * A location is either a compute zone (e.g. `us-central1-a`) or a region
     * (e.g. `us-central1`).
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<PolicyClusterAdmissionRule>>> clusterAdmissionRules() {
        return Codegen.optional(this.clusterAdmissionRules);
    }
    /**
     * Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     * 
     */
    @Export(name="defaultAdmissionRule", refs={PolicyDefaultAdmissionRule.class}, tree="[0]")
    private Output<PolicyDefaultAdmissionRule> defaultAdmissionRule;

    /**
     * @return Default admission rule for a cluster without a per-cluster admission
     * rule.
     * Structure is documented below.
     * 
     */
    public Output<PolicyDefaultAdmissionRule> defaultAdmissionRule() {
        return this.defaultAdmissionRule;
    }
    /**
     * A descriptive comment.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A descriptive comment.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are: `ENABLE`, `DISABLE`.
     * 
     */
    @Export(name="globalPolicyEvaluationMode", refs={String.class}, tree="[0]")
    private Output<String> globalPolicyEvaluationMode;

    /**
     * @return Controls the evaluation of a Google-maintained global admission policy
     * for common system-level images. Images not covered by the global
     * policy will be subject to the project admission policy.
     * Possible values are: `ENABLE`, `DISABLE`.
     * 
     */
    public Output<String> globalPolicyEvaluationMode() {
        return this.globalPolicyEvaluationMode;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Policy(String name) {
        this(name, PolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Policy(String name, PolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Policy(String name, PolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:binaryauthorization/policy:Policy", name, args == null ? PolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Policy(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:binaryauthorization/policy:Policy", name, state, makeResourceOptions(options, id));
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
    public static Policy get(String name, Output<String> id, @Nullable PolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Policy(name, id, state, options);
    }
}
