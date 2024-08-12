// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.privilegedaccessmanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.privilegedaccessmanager.EntitlementArgs;
import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementState;
import com.pulumi.gcp.privilegedaccessmanager.outputs.EntitlementAdditionalNotificationTargets;
import com.pulumi.gcp.privilegedaccessmanager.outputs.EntitlementApprovalWorkflow;
import com.pulumi.gcp.privilegedaccessmanager.outputs.EntitlementEligibleUser;
import com.pulumi.gcp.privilegedaccessmanager.outputs.EntitlementPrivilegedAccess;
import com.pulumi.gcp.privilegedaccessmanager.outputs.EntitlementRequesterJustificationConfig;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An Entitlement defines the eligibility of a set of users to obtain a predefined access for some time possibly after going through an approval workflow.
 * 
 * ## Example Usage
 * 
 * ### Privileged Access Manager Entitlement Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.privilegedaccessmanager.entitlement;
 * import com.pulumi.gcp.privilegedaccessmanager.EntitlementArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementRequesterJustificationConfigArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementRequesterJustificationConfigUnstructuredArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementEligibleUserArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementPrivilegedAccessArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementPrivilegedAccessGcpIamAccessArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementAdditionalNotificationTargetsArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementApprovalWorkflowArgs;
 * import com.pulumi.gcp.privilegedaccessmanager.inputs.EntitlementApprovalWorkflowManualApprovalsArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var tfentitlement = new Entitlement("tfentitlement", EntitlementArgs.builder()
 *             .entitlementId("example-entitlement")
 *             .location("global")
 *             .maxRequestDuration("43200s")
 *             .parent("projects/my-project-name")
 *             .requesterJustificationConfig(EntitlementRequesterJustificationConfigArgs.builder()
 *                 .unstructured()
 *                 .build())
 *             .eligibleUsers(EntitlementEligibleUserArgs.builder()
 *                 .principals("group:test}{@literal @}{@code google.com")
 *                 .build())
 *             .privilegedAccess(EntitlementPrivilegedAccessArgs.builder()
 *                 .gcpIamAccess(EntitlementPrivilegedAccessGcpIamAccessArgs.builder()
 *                     .roleBindings(EntitlementPrivilegedAccessGcpIamAccessRoleBindingArgs.builder()
 *                         .role("roles/storage.admin")
 *                         .conditionExpression("request.time < timestamp(\"2024-04-23T18:30:00.000Z\")")
 *                         .build())
 *                     .resource("//cloudresourcemanager.googleapis.com/projects/my-project-name")
 *                     .resourceType("cloudresourcemanager.googleapis.com/Project")
 *                     .build())
 *                 .build())
 *             .additionalNotificationTargets(EntitlementAdditionalNotificationTargetsArgs.builder()
 *                 .adminEmailRecipients("user}{@literal @}{@code example.com")
 *                 .requesterEmailRecipients("user}{@literal @}{@code example.com")
 *                 .build())
 *             .approvalWorkflow(EntitlementApprovalWorkflowArgs.builder()
 *                 .manualApprovals(EntitlementApprovalWorkflowManualApprovalsArgs.builder()
 *                     .requireApproverJustification(true)
 *                     .steps(EntitlementApprovalWorkflowManualApprovalsStepArgs.builder()
 *                         .approvalsNeeded(1)
 *                         .approverEmailRecipients("user}{@literal @}{@code example.com")
 *                         .approvers(EntitlementApprovalWorkflowManualApprovalsStepApproversArgs.builder()
 *                             .principals("group:test}{@literal @}{@code google.com")
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Entitlement can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/locations/{{location}}/entitlements/{{entitlement_id}}`
 * 
 * When using the `pulumi import` command, Entitlement can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:privilegedaccessmanager/entitlement:entitlement default {{parent}}/locations/{{location}}/entitlements/{{entitlement_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:privilegedaccessmanager/entitlement:entitlement")
public class Entitlement extends com.pulumi.resources.CustomResource {
    /**
     * AdditionalNotificationTargets includes email addresses to be notified.
     * 
     */
    @Export(name="additionalNotificationTargets", refs={EntitlementAdditionalNotificationTargets.class}, tree="[0]")
    private Output</* @Nullable */ EntitlementAdditionalNotificationTargets> additionalNotificationTargets;

    /**
     * @return AdditionalNotificationTargets includes email addresses to be notified.
     * 
     */
    public Output<Optional<EntitlementAdditionalNotificationTargets>> additionalNotificationTargets() {
        return Codegen.optional(this.additionalNotificationTargets);
    }
    /**
     * The approvals needed before access will be granted to a requester. No approvals will be needed if this field is null.
     * Different types of approval workflows that can be used to gate privileged access granting.
     * 
     */
    @Export(name="approvalWorkflow", refs={EntitlementApprovalWorkflow.class}, tree="[0]")
    private Output</* @Nullable */ EntitlementApprovalWorkflow> approvalWorkflow;

    /**
     * @return The approvals needed before access will be granted to a requester. No approvals will be needed if this field is null.
     * Different types of approval workflows that can be used to gate privileged access granting.
     * 
     */
    public Output<Optional<EntitlementApprovalWorkflow>> approvalWorkflow() {
        return Codegen.optional(this.approvalWorkflow);
    }
    /**
     * Output only. Create time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Create time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Who can create Grants using Entitlement. This list should contain at most one entry
     * Structure is documented below.
     * 
     */
    @Export(name="eligibleUsers", refs={List.class,EntitlementEligibleUser.class}, tree="[0,1]")
    private Output<List<EntitlementEligibleUser>> eligibleUsers;

    /**
     * @return Who can create Grants using Entitlement. This list should contain at most one entry
     * Structure is documented below.
     * 
     */
    public Output<List<EntitlementEligibleUser>> eligibleUsers() {
        return this.eligibleUsers;
    }
    /**
     * The ID to use for this Entitlement. This will become the last part of the resource name.
     * This value should be 4-63 characters, and valid characters are &#34;[a-z]&#34;, &#34;[0-9]&#34;, and &#34;-&#34;. The first character should be from [a-z].
     * This value should be unique among all other Entitlements under the specified `parent`.
     * 
     */
    @Export(name="entitlementId", refs={String.class}, tree="[0]")
    private Output<String> entitlementId;

    /**
     * @return The ID to use for this Entitlement. This will become the last part of the resource name.
     * This value should be 4-63 characters, and valid characters are &#34;[a-z]&#34;, &#34;[0-9]&#34;, and &#34;-&#34;. The first character should be from [a-z].
     * This value should be unique among all other Entitlements under the specified `parent`.
     * 
     */
    public Output<String> entitlementId() {
        return this.entitlementId;
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
     * The region of the Entitlement resource.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The region of the Entitlement resource.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The maximum amount of time for which access would be granted for a request.
     * A requester can choose to ask for access for less than this duration but never more.
     * Format: calculate the time in seconds and concatenate it with &#39;s&#39; i.e. 2 hours = &#34;7200s&#34;, 45 minutes = &#34;2700s&#34;
     * 
     */
    @Export(name="maxRequestDuration", refs={String.class}, tree="[0]")
    private Output<String> maxRequestDuration;

    /**
     * @return The maximum amount of time for which access would be granted for a request.
     * A requester can choose to ask for access for less than this duration but never more.
     * Format: calculate the time in seconds and concatenate it with &#39;s&#39; i.e. 2 hours = &#34;7200s&#34;, 45 minutes = &#34;2700s&#34;
     * 
     */
    public Output<String> maxRequestDuration() {
        return this.maxRequestDuration;
    }
    /**
     * Output Only. The entitlement&#39;s name follows a hierarchical structure, comprising the organization, folder, or project, alongside the region and a unique entitlement ID.
     * Formats: organizations/{organization-number}/locations/{region}/entitlements/{entitlement-id}, folders/{folder-number}/locations/{region}/entitlements/{entitlement-id}, and projects/{project-id|project-number}/locations/{region}/entitlements/{entitlement-id}.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Output Only. The entitlement&#39;s name follows a hierarchical structure, comprising the organization, folder, or project, alongside the region and a unique entitlement ID.
     * Formats: organizations/{organization-number}/locations/{region}/entitlements/{entitlement-id}, folders/{folder-number}/locations/{region}/entitlements/{entitlement-id}, and projects/{project-id|project-number}/locations/{region}/entitlements/{entitlement-id}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Format: projects/{project-id|project-number} or organizations/{organization-number} or folders/{folder-number}
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return Format: projects/{project-id|project-number} or organizations/{organization-number} or folders/{folder-number}
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Privileged access that this service can be used to gate.
     * Structure is documented below.
     * 
     */
    @Export(name="privilegedAccess", refs={EntitlementPrivilegedAccess.class}, tree="[0]")
    private Output<EntitlementPrivilegedAccess> privilegedAccess;

    /**
     * @return Privileged access that this service can be used to gate.
     * Structure is documented below.
     * 
     */
    public Output<EntitlementPrivilegedAccess> privilegedAccess() {
        return this.privilegedAccess;
    }
    /**
     * Defines the ways in which a requester should provide the justification while requesting for access.
     * Structure is documented below.
     * 
     */
    @Export(name="requesterJustificationConfig", refs={EntitlementRequesterJustificationConfig.class}, tree="[0]")
    private Output<EntitlementRequesterJustificationConfig> requesterJustificationConfig;

    /**
     * @return Defines the ways in which a requester should provide the justification while requesting for access.
     * Structure is documented below.
     * 
     */
    public Output<EntitlementRequesterJustificationConfig> requesterJustificationConfig() {
        return this.requesterJustificationConfig;
    }
    /**
     * Output only. The current state of the Entitlement.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. The current state of the Entitlement.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Output only. Update time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. Update time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Entitlement(java.lang.String name) {
        this(name, EntitlementArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Entitlement(java.lang.String name, EntitlementArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Entitlement(java.lang.String name, EntitlementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:privilegedaccessmanager/entitlement:entitlement", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Entitlement(java.lang.String name, Output<java.lang.String> id, @Nullable EntitlementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:privilegedaccessmanager/entitlement:entitlement", name, state, makeResourceOptions(options, id), false);
    }

    private static EntitlementArgs makeArgs(EntitlementArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EntitlementArgs.Empty : args;
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
    public static Entitlement get(java.lang.String name, Output<java.lang.String> id, @Nullable EntitlementState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Entitlement(name, id, state, options);
    }
}
