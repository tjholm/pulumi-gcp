// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctionsv2;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudfunctionsv2.FunctionIamPolicyArgs;
import com.pulumi.gcp.cloudfunctionsv2.inputs.FunctionIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/locations/{{location}}/functions/{{cloud_function}} * {{project}}/{{location}}/{{cloud_function}} * {{location}}/{{cloud_function}} * {{cloud_function}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Functions (2nd gen) function IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/functionIamPolicy:FunctionIamPolicy editor &#34;projects/{{project}}/locations/{{location}}/functions/{{cloud_function}} roles/viewer user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/functionIamPolicy:FunctionIamPolicy editor &#34;projects/{{project}}/locations/{{location}}/functions/{{cloud_function}} roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/functionIamPolicy:FunctionIamPolicy editor projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:cloudfunctionsv2/functionIamPolicy:FunctionIamPolicy")
public class FunctionIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="cloudFunction", type=String.class, parameters={})
    private Output<String> cloudFunction;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> cloudFunction() {
        return this.cloudFunction;
    }
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The location of this cloud function. Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The location of this cloud function. Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     * 
     */
    @Export(name="policyData", type=String.class, parameters={})
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
    @Export(name="project", type=String.class, parameters={})
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FunctionIamPolicy(String name) {
        this(name, FunctionIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FunctionIamPolicy(String name, FunctionIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FunctionIamPolicy(String name, FunctionIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudfunctionsv2/functionIamPolicy:FunctionIamPolicy", name, args == null ? FunctionIamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FunctionIamPolicy(String name, Output<String> id, @Nullable FunctionIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudfunctionsv2/functionIamPolicy:FunctionIamPolicy", name, state, makeResourceOptions(options, id));
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
    public static FunctionIamPolicy get(String name, Output<String> id, @Nullable FunctionIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FunctionIamPolicy(name, id, state, options);
    }
}
