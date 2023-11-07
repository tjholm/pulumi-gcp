// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.assuredworkloads;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.assuredworkloads.WorkloadArgs;
import com.pulumi.gcp.assuredworkloads.inputs.WorkloadState;
import com.pulumi.gcp.assuredworkloads.outputs.WorkloadKmsSettings;
import com.pulumi.gcp.assuredworkloads.outputs.WorkloadResource;
import com.pulumi.gcp.assuredworkloads.outputs.WorkloadResourceSetting;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The AssuredWorkloads Workload resource
 * 
 * ## Example Usage
 * ### Basic_workload
 * A basic test of a assuredworkloads api
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.assuredworkloads.Workload;
 * import com.pulumi.gcp.assuredworkloads.WorkloadArgs;
 * import com.pulumi.gcp.assuredworkloads.inputs.WorkloadKmsSettingsArgs;
 * import com.pulumi.gcp.assuredworkloads.inputs.WorkloadResourceSettingArgs;
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
 *         var primary = new Workload(&#34;primary&#34;, WorkloadArgs.builder()        
 *             .billingAccount(&#34;billingAccounts/000000-0000000-0000000-000000&#34;)
 *             .complianceRegime(&#34;FEDRAMP_MODERATE&#34;)
 *             .displayName(&#34;Workload Example&#34;)
 *             .kmsSettings(WorkloadKmsSettingsArgs.builder()
 *                 .nextRotationTime(&#34;9999-10-02T15:01:23Z&#34;)
 *                 .rotationPeriod(&#34;10368000s&#34;)
 *                 .build())
 *             .labels(Map.of(&#34;label-one&#34;, &#34;value-one&#34;))
 *             .location(&#34;us-west1&#34;)
 *             .organization(&#34;123456789&#34;)
 *             .provisionedResourcesParent(&#34;folders/519620126891&#34;)
 *             .resourceSettings(            
 *                 WorkloadResourceSettingArgs.builder()
 *                     .resourceType(&#34;CONSUMER_PROJECT&#34;)
 *                     .build(),
 *                 WorkloadResourceSettingArgs.builder()
 *                     .resourceType(&#34;ENCRYPTION_KEYS_PROJECT&#34;)
 *                     .build(),
 *                 WorkloadResourceSettingArgs.builder()
 *                     .resourceId(&#34;ring&#34;)
 *                     .resourceType(&#34;KEYRING&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Workload can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:assuredworkloads/workload:Workload default organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:assuredworkloads/workload:Workload default {{organization}}/{{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:assuredworkloads/workload:Workload")
public class Workload extends com.pulumi.resources.CustomResource {
    /**
     * Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, &#39;billingAccounts/012345-567890-ABCDEF`.
     * 
     */
    @Export(name="billingAccount", refs={String.class}, tree="[0]")
    private Output<String> billingAccount;

    /**
     * @return Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, &#39;billingAccounts/012345-567890-ABCDEF`.
     * 
     */
    public Output<String> billingAccount() {
        return this.billingAccount;
    }
    /**
     * Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS
     * 
     */
    @Export(name="complianceRegime", refs={String.class}, tree="[0]")
    private Output<String> complianceRegime;

    /**
     * @return Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS
     * 
     */
    public Output<String> complianceRegime() {
        return this.complianceRegime;
    }
    /**
     * Output only. Immutable. The Workload creation timestamp.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Immutable. The Workload creation timestamp.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,Object>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
     * 
     */
    @Export(name="kmsSettings", refs={WorkloadKmsSettings.class}, tree="[0]")
    private Output</* @Nullable */ WorkloadKmsSettings> kmsSettings;

    /**
     * @return Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
     * 
     */
    public Output<Optional<WorkloadKmsSettings>> kmsSettings() {
        return Codegen.optional(this.kmsSettings);
    }
    /**
     * Optional. Labels applied to the workload.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Optional. Labels applied to the workload.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Output only. The resource name of the workload.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Output only. The resource name of the workload.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The organization for the resource
     * 
     * ***
     * 
     */
    @Export(name="organization", refs={String.class}, tree="[0]")
    private Output<String> organization;

    /**
     * @return The organization for the resource
     * 
     * ***
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    /**
     * Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
     * 
     */
    @Export(name="provisionedResourcesParent", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> provisionedResourcesParent;

    /**
     * @return Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
     * 
     */
    public Output<Optional<String>> provisionedResourcesParent() {
        return Codegen.optional(this.provisionedResourcesParent);
    }
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,Object>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     * 
     */
    @Export(name="resourceSettings", refs={List.class,WorkloadResourceSetting.class}, tree="[0,1]")
    private Output</* @Nullable */ List<WorkloadResourceSetting>> resourceSettings;

    /**
     * @return Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     * 
     */
    public Output<Optional<List<WorkloadResourceSetting>>> resourceSettings() {
        return Codegen.optional(this.resourceSettings);
    }
    /**
     * Output only. The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
     * 
     */
    @Export(name="resources", refs={List.class,WorkloadResource.class}, tree="[0,1]")
    private Output<List<WorkloadResource>> resources;

    /**
     * @return Output only. The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
     * 
     */
    public Output<List<WorkloadResource>> resources() {
        return this.resources;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Workload(String name) {
        this(name, WorkloadArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Workload(String name, WorkloadArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Workload(String name, WorkloadArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:assuredworkloads/workload:Workload", name, args == null ? WorkloadArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Workload(String name, Output<String> id, @Nullable WorkloadState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:assuredworkloads/workload:Workload", name, state, makeResourceOptions(options, id));
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
    public static Workload get(String name, Output<String> id, @Nullable WorkloadState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Workload(name, id, state, options);
    }
}
