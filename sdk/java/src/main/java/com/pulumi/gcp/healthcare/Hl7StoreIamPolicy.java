// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.healthcare;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.healthcare.Hl7StoreIamPolicyArgs;
import com.pulumi.gcp.healthcare.inputs.Hl7StoreIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
 * 
 * * `gcp.healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
 * * `gcp.healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
 * * `gcp.healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
 * 
 * &gt; **Note:** `gcp.healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.Hl7StoreIamBinding` and `gcp.healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_healthcare\_hl7\_v2\_store\_iam\_policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamPolicy;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/editor")
 *                 .members("user:jane{@literal @}example.com")
 *                 .build())
 *             .build());
 * 
 *         var hl7V2Store = new Hl7StoreIamPolicy("hl7V2Store", Hl7StoreIamPolicyArgs.builder()
 *             .hl7V2StoreId("your-hl7-v2-store-id")
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_healthcare\_hl7\_v2\_store\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamBinding;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamBindingArgs;
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
 *         var hl7V2Store = new Hl7StoreIamBinding("hl7V2Store", Hl7StoreIamBindingArgs.builder()
 *             .hl7V2StoreId("your-hl7-v2-store-id")
 *             .role("roles/editor")
 *             .members("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_healthcare\_hl7\_v2\_store\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamMember;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamMemberArgs;
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
 *         var hl7V2Store = new Hl7StoreIamMember("hl7V2Store", Hl7StoreIamMemberArgs.builder()
 *             .hl7V2StoreId("your-hl7-v2-store-id")
 *             .role("roles/editor")
 *             .member("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_healthcare\_hl7\_v2\_store\_iam\_policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamPolicy;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/editor")
 *                 .members("user:jane{@literal @}example.com")
 *                 .build())
 *             .build());
 * 
 *         var hl7V2Store = new Hl7StoreIamPolicy("hl7V2Store", Hl7StoreIamPolicyArgs.builder()
 *             .hl7V2StoreId("your-hl7-v2-store-id")
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_healthcare\_hl7\_v2\_store\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamBinding;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamBindingArgs;
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
 *         var hl7V2Store = new Hl7StoreIamBinding("hl7V2Store", Hl7StoreIamBindingArgs.builder()
 *             .hl7V2StoreId("your-hl7-v2-store-id")
 *             .role("roles/editor")
 *             .members("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_healthcare\_hl7\_v2\_store\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamMember;
 * import com.pulumi.gcp.healthcare.Hl7StoreIamMemberArgs;
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
 *         var hl7V2Store = new Hl7StoreIamMember("hl7V2Store", Hl7StoreIamMemberArgs.builder()
 *             .hl7V2StoreId("your-hl7-v2-store-id")
 *             .role("roles/editor")
 *             .member("user:jane{@literal @}example.com")
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
 * ### Importing IAM policies
 * 
 * IAM policy imports use the identifier of the Google Cloud Healthcare HL7v2 store resource. For example:
 * 
 * * `&#34;{{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}}&#34;`
 * 
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 * 
 * tf
 * 
 * import {
 * 
 *   id = &#34;{{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}}&#34;
 * 
 *   to = google_healthcare_hl7_v2_store_iam_policy.default
 * 
 * }
 * 
 * The `pulumi import` command can also be used:
 * 
 * ```sh
 * $ pulumi import gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy default {{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}}
 * ```
 * 
 */
@ResourceType(type="gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy")
public class Hl7StoreIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * (Computed) The etag of the HL7v2 store&#39;s IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the HL7v2 store&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The HL7v2 store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
     * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Export(name="hl7V2StoreId", refs={String.class}, tree="[0]")
    private Output<String> hl7V2StoreId;

    /**
     * @return The HL7v2 store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
     * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    public Output<String> hl7V2StoreId() {
        return this.hl7V2StoreId;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Hl7StoreIamPolicy(String name) {
        this(name, Hl7StoreIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Hl7StoreIamPolicy(String name, Hl7StoreIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Hl7StoreIamPolicy(String name, Hl7StoreIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, args == null ? Hl7StoreIamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Hl7StoreIamPolicy(String name, Output<String> id, @Nullable Hl7StoreIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, state, makeResourceOptions(options, id));
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
    public static Hl7StoreIamPolicy get(String name, Output<String> id, @Nullable Hl7StoreIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Hl7StoreIamPolicy(name, id, state, options);
    }
}
