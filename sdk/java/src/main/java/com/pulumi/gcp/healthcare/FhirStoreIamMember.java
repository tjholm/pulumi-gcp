// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.healthcare.FhirStoreIamMemberArgs;
import com.pulumi.gcp.healthcare.inputs.FhirStoreIamMemberState;
import com.pulumi.gcp.healthcare.outputs.FhirStoreIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
 * 
 * * `gcp.healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
 * * `gcp.healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
 * * `gcp.healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
 * 
 * &gt; **Note:** `gcp.healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.FhirStoreIamBinding` and `gcp.healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_healthcare\_fhir\_store\_iam\_policy
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var admin = Output.of(OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role(&#34;roles/editor&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var fhirStore = new FhirStoreIamPolicy(&#34;fhirStore&#34;, FhirStoreIamPolicyArgs.builder()        
 *             .fhirStoreId(&#34;your-fhir-store-id&#34;)
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_healthcare\_fhir\_store\_iam\_binding
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var fhirStore = new FhirStoreIamBinding(&#34;fhirStore&#34;, FhirStoreIamBindingArgs.builder()        
 *             .fhirStoreId(&#34;your-fhir-store-id&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .role(&#34;roles/editor&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_healthcare\_fhir\_store\_iam\_member
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var fhirStore = new FhirStoreIamMember(&#34;fhirStore&#34;, FhirStoreIamMemberArgs.builder()        
 *             .fhirStoreId(&#34;your-fhir-store-id&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .role(&#34;roles/editor&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 * 
 * This member resource can be imported using the `fhir_store_id`, role, and account e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember fhir_store_iam &#34;your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 * 
 * This binding resource can be imported using the `fhir_store_id` and role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember fhir_store_iam &#34;your-project-id/location-name/dataset-name/fhir-store-name roles/viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question.
 * 
 * This policy resource can be imported using the `fhir_store_id`, role, and account e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name
 * ```
 * 
 */
@ResourceType(type="gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember")
public class FhirStoreIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=FhirStoreIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ FhirStoreIamMemberCondition> condition;

    public Output<Optional<FhirStoreIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the FHIR store&#39;s IAM policy.
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the FHIR store&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The FHIR store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
     * `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Export(name="fhirStoreId", type=String.class, parameters={})
    private Output<String> fhirStoreId;

    /**
     * @return The FHIR store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
     * `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    public Output<String> fhirStoreId() {
        return this.fhirStoreId;
    }
    @Export(name="member", type=String.class, parameters={})
    private Output<String> member;

    public Output<String> member() {
        return this.member;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FhirStoreIamMember(String name) {
        this(name, FhirStoreIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FhirStoreIamMember(String name, FhirStoreIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FhirStoreIamMember(String name, FhirStoreIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember", name, args == null ? FhirStoreIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FhirStoreIamMember(String name, Output<String> id, @Nullable FhirStoreIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember", name, state, makeResourceOptions(options, id));
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
    public static FhirStoreIamMember get(String name, Output<String> id, @Nullable FhirStoreIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FhirStoreIamMember(name, id, state, options);
    }
}
