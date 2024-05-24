// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.migrationcenter;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.migrationcenter.PreferenceSetArgs;
import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetState;
import com.pulumi.gcp.migrationcenter.outputs.PreferenceSetVirtualMachinePreferences;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages the PreferenceSet resource.
 * 
 * To get more information about PreferenceSet, see:
 * 
 * * [API documentation](https://cloud.google.com/migration-center/docs/reference/rest/v1)
 * * How-to Guides
 *     * [Managing Migration Preferences](https://cloud.google.com/migration-center/docs/migration-preferences)
 * 
 * ## Example Usage
 * 
 * ### Preference Set Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.migrationcenter.PreferenceSet;
 * import com.pulumi.gcp.migrationcenter.PreferenceSetArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesArgs;
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
 *         var default_ = new PreferenceSet("default", PreferenceSetArgs.builder()
 *             .location("us-central1")
 *             .preferenceSetId("preference-set-test")
 *             .description("Terraform integration test description")
 *             .displayName("Terraform integration test display")
 *             .virtualMachinePreferences(PreferenceSetVirtualMachinePreferencesArgs.builder()
 *                 .vmwareEnginePreferences(PreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesArgs.builder()
 *                     .cpuOvercommitRatio(1.5)
 *                     .build())
 *                 .sizingOptimizationStrategy("SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE")
 *                 .targetProduct("COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Preference Set Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.migrationcenter.PreferenceSet;
 * import com.pulumi.gcp.migrationcenter.PreferenceSetArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesRegionPreferencesArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesArgs;
 * import com.pulumi.gcp.migrationcenter.inputs.PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesArgs;
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
 *         var default_ = new PreferenceSet("default", PreferenceSetArgs.builder()
 *             .location("us-central1")
 *             .preferenceSetId("preference-set-test")
 *             .description("Terraform integration test description")
 *             .displayName("Terraform integration test display")
 *             .virtualMachinePreferences(PreferenceSetVirtualMachinePreferencesArgs.builder()
 *                 .vmwareEnginePreferences(PreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesArgs.builder()
 *                     .cpuOvercommitRatio(1.5)
 *                     .storageDeduplicationCompressionRatio(1.3)
 *                     .commitmentPlan("ON_DEMAND")
 *                     .build())
 *                 .sizingOptimizationStrategy("SIZING_OPTIMIZATION_STRATEGY_SAME_AS_SOURCE")
 *                 .targetProduct("COMPUTE_MIGRATION_TARGET_PRODUCT_COMPUTE_ENGINE")
 *                 .commitmentPlan("COMMITMENT_PLAN_ONE_YEAR")
 *                 .regionPreferences(PreferenceSetVirtualMachinePreferencesRegionPreferencesArgs.builder()
 *                     .preferredRegions("us-central1")
 *                     .build())
 *                 .soleTenancyPreferences(PreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesArgs.builder()
 *                     .commitmentPlan("ON_DEMAND")
 *                     .cpuOvercommitRatio(1.2)
 *                     .hostMaintenancePolicy("HOST_MAINTENANCE_POLICY_DEFAULT")
 *                     .nodeTypes(PreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesNodeTypeArgs.builder()
 *                         .nodeName("tf-test")
 *                         .build())
 *                     .build())
 *                 .computeEnginePreferences(PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesArgs.builder()
 *                     .licenseType("LICENSE_TYPE_BRING_YOUR_OWN_LICENSE")
 *                     .machinePreferences(PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesArgs.builder()
 *                         .allowedMachineSeries(PreferenceSetVirtualMachinePreferencesComputeEnginePreferencesMachinePreferencesAllowedMachineSeriesArgs.builder()
 *                             .code("C3")
 *                             .build())
 *                         .build())
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
 * PreferenceSet can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}`
 * 
 * * `{{project}}/{{location}}/{{preference_set_id}}`
 * 
 * * `{{location}}/{{preference_set_id}}`
 * 
 * When using the `pulumi import` command, PreferenceSet can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:migrationcenter/preferenceSet:PreferenceSet default projects/{{project}}/locations/{{location}}/preferenceSets/{{preference_set_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:migrationcenter/preferenceSet:PreferenceSet default {{project}}/{{location}}/{{preference_set_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:migrationcenter/preferenceSet:PreferenceSet default {{location}}/{{preference_set_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:migrationcenter/preferenceSet:PreferenceSet")
public class PreferenceSet extends com.pulumi.resources.CustomResource {
    /**
     * Output only. The timestamp when the preference set was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. The timestamp when the preference set was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A description of the preference set.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the preference set.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * User-friendly display name. Maximum length is 63 characters.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User-friendly display name. Maximum length is 63 characters.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Part of `parent`. See documentation of `projectsId`.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return Part of `parent`. See documentation of `projectsId`.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Output only. Name of the preference set.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Output only. Name of the preference set.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression `a-z?`.
     * 
     * ***
     * 
     */
    @Export(name="preferenceSetId", refs={String.class}, tree="[0]")
    private Output<String> preferenceSetId;

    /**
     * @return Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression `a-z?`.
     * 
     * ***
     * 
     */
    public Output<String> preferenceSetId() {
        return this.preferenceSetId;
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
     * Output only. The timestamp when the preference set was last updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. The timestamp when the preference set was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * VirtualMachinePreferences enables you to create sets of assumptions, for example, a geographical location and pricing track, for your migrated virtual machines. The set of preferences influence recommendations for migrating virtual machine assets.
     * Structure is documented below.
     * 
     */
    @Export(name="virtualMachinePreferences", refs={PreferenceSetVirtualMachinePreferences.class}, tree="[0]")
    private Output</* @Nullable */ PreferenceSetVirtualMachinePreferences> virtualMachinePreferences;

    /**
     * @return VirtualMachinePreferences enables you to create sets of assumptions, for example, a geographical location and pricing track, for your migrated virtual machines. The set of preferences influence recommendations for migrating virtual machine assets.
     * Structure is documented below.
     * 
     */
    public Output<Optional<PreferenceSetVirtualMachinePreferences>> virtualMachinePreferences() {
        return Codegen.optional(this.virtualMachinePreferences);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PreferenceSet(String name) {
        this(name, PreferenceSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PreferenceSet(String name, PreferenceSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PreferenceSet(String name, PreferenceSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:migrationcenter/preferenceSet:PreferenceSet", name, args == null ? PreferenceSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PreferenceSet(String name, Output<String> id, @Nullable PreferenceSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:migrationcenter/preferenceSet:PreferenceSet", name, state, makeResourceOptions(options, id));
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
    public static PreferenceSet get(String name, Output<String> id, @Nullable PreferenceSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PreferenceSet(name, id, state, options);
    }
}
