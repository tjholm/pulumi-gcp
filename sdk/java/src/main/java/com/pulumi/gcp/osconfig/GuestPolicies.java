// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.osconfig;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.osconfig.GuestPoliciesArgs;
import com.pulumi.gcp.osconfig.inputs.GuestPoliciesState;
import com.pulumi.gcp.osconfig.outputs.GuestPoliciesAssignment;
import com.pulumi.gcp.osconfig.outputs.GuestPoliciesPackage;
import com.pulumi.gcp.osconfig.outputs.GuestPoliciesPackageRepository;
import com.pulumi.gcp.osconfig.outputs.GuestPoliciesRecipe;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Os Config Guest Policies Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
 * import com.pulumi.gcp.compute.Instance;
 * import com.pulumi.gcp.compute.InstanceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskInitializeParamsArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceNetworkInterfaceArgs;
 * import com.pulumi.gcp.osconfig.GuestPolicies;
 * import com.pulumi.gcp.osconfig.GuestPoliciesArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesAssignmentArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesPackageArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         final var myImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-11&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build());
 * 
 *         var foobar = new Instance(&#34;foobar&#34;, InstanceArgs.builder()        
 *             .machineType(&#34;e2-medium&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .canIpForward(false)
 *             .tags(            
 *                 &#34;foo&#34;,
 *                 &#34;bar&#34;)
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image(myImage.applyValue(getImageResult -&gt; getImageResult.selfLink()))
 *                     .build())
 *                 .build())
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .build())
 *             .metadata(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var guestPolicies = new GuestPolicies(&#34;guestPolicies&#34;, GuestPoliciesArgs.builder()        
 *             .guestPolicyId(&#34;guest-policy&#34;)
 *             .assignment(GuestPoliciesAssignmentArgs.builder()
 *                 .instances(foobar.id())
 *                 .build())
 *             .packages(GuestPoliciesPackageArgs.builder()
 *                 .name(&#34;my-package&#34;)
 *                 .desiredState(&#34;UPDATED&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Os Config Guest Policies Packages
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.osconfig.GuestPolicies;
 * import com.pulumi.gcp.osconfig.GuestPoliciesArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesAssignmentArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesPackageArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesPackageRepositoryArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesPackageRepositoryAptArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesPackageRepositoryYumArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var guestPolicies = new GuestPolicies(&#34;guestPolicies&#34;, GuestPoliciesArgs.builder()        
 *             .guestPolicyId(&#34;guest-policy&#34;)
 *             .assignment(GuestPoliciesAssignmentArgs.builder()
 *                 .groupLabels(                
 *                     GuestPoliciesAssignmentGroupLabelArgs.builder()
 *                         .labels(Map.ofEntries(
 *                             Map.entry(&#34;color&#34;, &#34;red&#34;),
 *                             Map.entry(&#34;env&#34;, &#34;test&#34;)
 *                         ))
 *                         .build(),
 *                     GuestPoliciesAssignmentGroupLabelArgs.builder()
 *                         .labels(Map.ofEntries(
 *                             Map.entry(&#34;color&#34;, &#34;blue&#34;),
 *                             Map.entry(&#34;env&#34;, &#34;test&#34;)
 *                         ))
 *                         .build())
 *                 .build())
 *             .packages(            
 *                 GuestPoliciesPackageArgs.builder()
 *                     .name(&#34;my-package&#34;)
 *                     .desiredState(&#34;INSTALLED&#34;)
 *                     .build(),
 *                 GuestPoliciesPackageArgs.builder()
 *                     .name(&#34;bad-package-1&#34;)
 *                     .desiredState(&#34;REMOVED&#34;)
 *                     .build(),
 *                 GuestPoliciesPackageArgs.builder()
 *                     .name(&#34;bad-package-2&#34;)
 *                     .desiredState(&#34;REMOVED&#34;)
 *                     .manager(&#34;APT&#34;)
 *                     .build())
 *             .packageRepositories(            
 *                 GuestPoliciesPackageRepositoryArgs.builder()
 *                     .apt(GuestPoliciesPackageRepositoryAptArgs.builder()
 *                         .uri(&#34;https://packages.cloud.google.com/apt&#34;)
 *                         .archiveType(&#34;DEB&#34;)
 *                         .distribution(&#34;cloud-sdk-stretch&#34;)
 *                         .components(&#34;main&#34;)
 *                         .build())
 *                     .build(),
 *                 GuestPoliciesPackageRepositoryArgs.builder()
 *                     .yum(GuestPoliciesPackageRepositoryYumArgs.builder()
 *                         .id(&#34;google-cloud-sdk&#34;)
 *                         .displayName(&#34;Google Cloud SDK&#34;)
 *                         .baseUrl(&#34;https://packages.cloud.google.com/yum/repos/cloud-sdk-el7-x86_64&#34;)
 *                         .gpgKeys(                        
 *                             &#34;https://packages.cloud.google.com/yum/doc/yum-key.gpg&#34;,
 *                             &#34;https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg&#34;)
 *                         .build())
 *                     .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Os Config Guest Policies Recipes
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.osconfig.GuestPolicies;
 * import com.pulumi.gcp.osconfig.GuestPoliciesArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesAssignmentArgs;
 * import com.pulumi.gcp.osconfig.inputs.GuestPoliciesRecipeArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var guestPolicies = new GuestPolicies(&#34;guestPolicies&#34;, GuestPoliciesArgs.builder()        
 *             .guestPolicyId(&#34;guest-policy&#34;)
 *             .assignment(GuestPoliciesAssignmentArgs.builder()
 *                 .zones(                
 *                     &#34;us-east1-b&#34;,
 *                     &#34;us-east1-d&#34;)
 *                 .build())
 *             .recipes(GuestPoliciesRecipeArgs.builder()
 *                 .name(&#34;guest-policy-recipe&#34;)
 *                 .desiredState(&#34;INSTALLED&#34;)
 *                 .artifacts(GuestPoliciesRecipeArtifactArgs.builder()
 *                     .id(&#34;guest-policy-artifact-id&#34;)
 *                     .gcs(GuestPoliciesRecipeArtifactGcsArgs.builder()
 *                         .bucket(&#34;my-bucket&#34;)
 *                         .object(&#34;executable.msi&#34;)
 *                         .generation(1546030865175603)
 *                         .build())
 *                     .build())
 *                 .installSteps(GuestPoliciesRecipeInstallStepArgs.builder()
 *                     .msiInstallation(GuestPoliciesRecipeInstallStepMsiInstallationArgs.builder()
 *                         .artifactId(&#34;guest-policy-artifact-id&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GuestPolicies can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default projects/{{project}}/guestPolicies/{{guest_policy_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default {{project}}/{{guest_policy_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:osconfig/guestPolicies:GuestPolicies default {{guest_policy_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:osconfig/guestPolicies:GuestPolicies")
public class GuestPolicies extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the VM instances that are assigned to this policy. This allows you to target sets
     * or groups of VM instances by different parameters such as labels, names, OS, or zones.
     * If left empty, all VM instances underneath this policy are targeted.
     * At the same level in the resource hierarchy (that is within a project), the service prevents
     * the creation of multiple policies that conflict with each other.
     * For more information, see how the service
     * [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
     * Structure is documented below.
     * 
     */
    @Export(name="assignment", type=GuestPoliciesAssignment.class, parameters={})
    private Output<GuestPoliciesAssignment> assignment;

    /**
     * @return Specifies the VM instances that are assigned to this policy. This allows you to target sets
     * or groups of VM instances by different parameters such as labels, names, OS, or zones.
     * If left empty, all VM instances underneath this policy are targeted.
     * At the same level in the resource hierarchy (that is within a project), the service prevents
     * the creation of multiple policies that conflict with each other.
     * For more information, see how the service
     * [handles assignment conflicts](https://cloud.google.com/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
     * Structure is documented below.
     * 
     */
    public Output<GuestPoliciesAssignment> assignment() {
        return this.assignment;
    }
    /**
     * Time this guest policy was created. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds.
     * Example: &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Time this guest policy was created. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds.
     * Example: &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the guest policy. Length of the description is limited to 1024 characters.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the guest policy. Length of the description is limited to 1024 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The etag for this guest policy. If this is provided on update, it must match the server&#39;s etag.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return The etag for this guest policy. If this is provided on update, it must match the server&#39;s etag.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The logical name of the guest policy in the project with the following restrictions:
     * * Must contain only lowercase letters, numbers, and hyphens.
     * * Must start with a letter.
     * * Must be between 1-63 characters.
     * * Must end with a number or a letter.
     * * Must be unique within the project.
     * 
     */
    @Export(name="guestPolicyId", type=String.class, parameters={})
    private Output<String> guestPolicyId;

    /**
     * @return The logical name of the guest policy in the project with the following restrictions:
     * * Must contain only lowercase letters, numbers, and hyphens.
     * * Must start with a letter.
     * * Must be between 1-63 characters.
     * * Must end with a number or a letter.
     * * Must be unique within the project.
     * 
     */
    public Output<String> guestPolicyId() {
        return this.guestPolicyId;
    }
    /**
     * The name of the package. A package is uniquely identified for conflict validation
     * by checking the package name and the manager(s) that the package targets.
     * (Required)
     * The name of the repository.
     * (Required)
     * Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
     * Names are also used to identify resources which helps to determine whether guest policies have conflicts.
     * This means that requests to create multiple recipes with the same name and version are rejected since they
     * could potentially have conflicting assignments.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the package. A package is uniquely identified for conflict validation
     * by checking the package name and the manager(s) that the package targets.
     * (Required)
     * The name of the repository.
     * (Required)
     * Unique identifier for the recipe. Only one recipe with a given name is installed on an instance.
     * Names are also used to identify resources which helps to determine whether guest policies have conflicts.
     * This means that requests to create multiple recipes with the same name and version are rejected since they
     * could potentially have conflicting assignments.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of package repositories to configure on the VM instance.
     * This is done before any other configs are applied so they can use these repos.
     * Package repositories are only configured if the corresponding package manager(s) are available.
     * Structure is documented below.
     * 
     */
    @Export(name="packageRepositories", type=List.class, parameters={GuestPoliciesPackageRepository.class})
    private Output</* @Nullable */ List<GuestPoliciesPackageRepository>> packageRepositories;

    /**
     * @return A list of package repositories to configure on the VM instance.
     * This is done before any other configs are applied so they can use these repos.
     * Package repositories are only configured if the corresponding package manager(s) are available.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<GuestPoliciesPackageRepository>>> packageRepositories() {
        return Codegen.optional(this.packageRepositories);
    }
    /**
     * The software packages to be managed by this policy.
     * Structure is documented below.
     * 
     */
    @Export(name="packages", type=List.class, parameters={GuestPoliciesPackage.class})
    private Output</* @Nullable */ List<GuestPoliciesPackage>> packages;

    /**
     * @return The software packages to be managed by this policy.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<GuestPoliciesPackage>>> packages() {
        return Codegen.optional(this.packages);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     * A list of Recipes to install on the VM instance.
     * Structure is documented below.
     * 
     */
    @Export(name="recipes", type=List.class, parameters={GuestPoliciesRecipe.class})
    private Output</* @Nullable */ List<GuestPoliciesRecipe>> recipes;

    /**
     * @return A list of Recipes to install on the VM instance.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<GuestPoliciesRecipe>>> recipes() {
        return Codegen.optional(this.recipes);
    }
    /**
     * Last time this guest policy was updated. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds.
     * Example: &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return Last time this guest policy was updated. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, accurate to nanoseconds.
     * Example: &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GuestPolicies(String name) {
        this(name, GuestPoliciesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GuestPolicies(String name, GuestPoliciesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GuestPolicies(String name, GuestPoliciesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:osconfig/guestPolicies:GuestPolicies", name, args == null ? GuestPoliciesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GuestPolicies(String name, Output<String> id, @Nullable GuestPoliciesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:osconfig/guestPolicies:GuestPolicies", name, state, makeResourceOptions(options, id));
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
    public static GuestPolicies get(String name, Output<String> id, @Nullable GuestPoliciesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GuestPolicies(name, id, state, options);
    }
}
