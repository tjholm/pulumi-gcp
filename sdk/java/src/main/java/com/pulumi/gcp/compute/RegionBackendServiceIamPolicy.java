// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.RegionBackendServiceIamPolicyArgs;
import com.pulumi.gcp.compute.inputs.RegionBackendServiceIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms:
 * 
 * * projects/{{project}}/regions/{{region}}/backendServices/{{name}}
 * 
 * * {{project}}/{{region}}/{{name}}
 * 
 * * {{region}}/{{name}}
 * 
 * * {{name}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Compute Engine regionbackendservice IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionBackendServiceIamPolicy:RegionBackendServiceIamPolicy editor &#34;projects/{{project}}/regions/{{region}}/backendServices/{{region_backend_service}} roles/compute.admin user:jane{@literal @}example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionBackendServiceIamPolicy:RegionBackendServiceIamPolicy editor &#34;projects/{{project}}/regions/{{region}}/backendServices/{{region_backend_service}} roles/compute.admin&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/regionBackendServiceIamPolicy:RegionBackendServiceIamPolicy editor projects/{{project}}/regions/{{region}}/backendServices/{{region_backend_service}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:compute/regionBackendServiceIamPolicy:RegionBackendServiceIamPolicy")
public class RegionBackendServiceIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Export(name="policyData", refs={String.class}, tree="[0]")
    private Output<String> policyData;

    /**
     * @return The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    public Output<String> policyData() {
        return this.policyData;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The Region in which the created backend service should reside.
     * If it is not provided, the provider region is used.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionBackendServiceIamPolicy(java.lang.String name) {
        this(name, RegionBackendServiceIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionBackendServiceIamPolicy(java.lang.String name, RegionBackendServiceIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionBackendServiceIamPolicy(java.lang.String name, RegionBackendServiceIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionBackendServiceIamPolicy:RegionBackendServiceIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegionBackendServiceIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable RegionBackendServiceIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/regionBackendServiceIamPolicy:RegionBackendServiceIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static RegionBackendServiceIamPolicyArgs makeArgs(RegionBackendServiceIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegionBackendServiceIamPolicyArgs.Empty : args;
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
    public static RegionBackendServiceIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable RegionBackendServiceIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionBackendServiceIamPolicy(name, id, state, options);
    }
}
