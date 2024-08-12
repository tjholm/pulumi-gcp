// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.kms.KeyRingIAMBindingArgs;
import com.pulumi.gcp.kms.inputs.KeyRingIAMBindingState;
import com.pulumi.gcp.kms.outputs.KeyRingIAMBindingCondition;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
 * 
 * * `gcp.kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
 * * `gcp.kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
 * * `gcp.kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
 * 
 * &gt; **Note:** `gcp.kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `gcp.kms.KeyRingIAMBinding` and `gcp.kms.KeyRingIAMMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.kms.KeyRingIAMBinding` resources **can be** used in conjunction with `gcp.kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.kms.KeyRingIAMPolicy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.kms.KeyRingIAMPolicy;
 * import com.pulumi.gcp.kms.KeyRingIAMPolicyArgs;
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
 *         var keyring = new KeyRing("keyring", KeyRingArgs.builder()
 *             .name("keyring-example")
 *             .location("global")
 *             .build());
 * 
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/editor")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var keyRing = new KeyRingIAMPolicy("keyRing", KeyRingIAMPolicyArgs.builder()
 *             .keyRingId(keyring.id())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.kms.KeyRingIAMPolicy;
 * import com.pulumi.gcp.kms.KeyRingIAMPolicyArgs;
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
 *         var keyring = new KeyRing("keyring", KeyRingArgs.builder()
 *             .name("keyring-example")
 *             .location("global")
 *             .build());
 * 
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/editor")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title("expires_after_2019_12_31")
 *                     .description("Expiring at midnight of 2019-12-31")
 *                     .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var keyRing = new KeyRingIAMPolicy("keyRing", KeyRingIAMPolicyArgs.builder()
 *             .keyRingId(keyring.id())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.kms.KeyRingIAMBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMBinding;
 * import com.pulumi.gcp.kms.KeyRingIAMBindingArgs;
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
 *         var keyRing = new KeyRingIAMBinding("keyRing", KeyRingIAMBindingArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMBinding;
 * import com.pulumi.gcp.kms.KeyRingIAMBindingArgs;
 * import com.pulumi.gcp.kms.inputs.KeyRingIAMBindingConditionArgs;
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
 *         var keyRing = new KeyRingIAMBinding("keyRing", KeyRingIAMBindingArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .condition(KeyRingIAMBindingConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.kms.KeyRingIAMMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMMember;
 * import com.pulumi.gcp.kms.KeyRingIAMMemberArgs;
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
 *         var keyRing = new KeyRingIAMMember("keyRing", KeyRingIAMMemberArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMMember;
 * import com.pulumi.gcp.kms.KeyRingIAMMemberArgs;
 * import com.pulumi.gcp.kms.inputs.KeyRingIAMMemberConditionArgs;
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
 *         var keyRing = new KeyRingIAMMember("keyRing", KeyRingIAMMemberArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .condition(KeyRingIAMMemberConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.kms.KeyRingIAMBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMBinding;
 * import com.pulumi.gcp.kms.KeyRingIAMBindingArgs;
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
 *         var keyRing = new KeyRingIAMBinding("keyRing", KeyRingIAMBindingArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMBinding;
 * import com.pulumi.gcp.kms.KeyRingIAMBindingArgs;
 * import com.pulumi.gcp.kms.inputs.KeyRingIAMBindingConditionArgs;
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
 *         var keyRing = new KeyRingIAMBinding("keyRing", KeyRingIAMBindingArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .condition(KeyRingIAMBindingConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.kms.KeyRingIAMMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMMember;
 * import com.pulumi.gcp.kms.KeyRingIAMMemberArgs;
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
 *         var keyRing = new KeyRingIAMMember("keyRing", KeyRingIAMMemberArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRingIAMMember;
 * import com.pulumi.gcp.kms.KeyRingIAMMemberArgs;
 * import com.pulumi.gcp.kms.inputs.KeyRingIAMMemberConditionArgs;
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
 *         var keyRing = new KeyRingIAMMember("keyRing", KeyRingIAMMemberArgs.builder()
 *             .keyRingId("your-key-ring-id")
 *             .role("roles/cloudkms.admin")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .condition(KeyRingIAMMemberConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
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
 * ### Importing IAM policies
 * 
 * IAM policy imports use the identifier of the Cloud KMS key ring only. For example:
 * 
 * * `{{project_id}}/{{location}}/{{key_ring_name}}`
 * 
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 * 
 * tf
 * 
 * import {
 * 
 *   id = &#34;{{project_id}}/{{location}}/{{key_ring_name}}&#34;
 * 
 *   to = google_kms_key_ring_iam_policy.default
 * 
 * }
 * 
 * The `pulumi import` command can also be used:
 * 
 * ```sh
 * $ pulumi import gcp:kms/keyRingIAMBinding:KeyRingIAMBinding default {{project_id}}/{{location}}/{{key_ring_name}}
 * ```
 * 
 */
@ResourceType(type="gcp:kms/keyRingIAMBinding:KeyRingIAMBinding")
public class KeyRingIAMBinding extends com.pulumi.resources.CustomResource {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Export(name="condition", refs={KeyRingIAMBindingCondition.class}, tree="[0]")
    private Output</* @Nullable */ KeyRingIAMBindingCondition> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Output<Optional<KeyRingIAMBindingCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the key ring&#39;s IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the key ring&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    @Export(name="keyRingId", refs={String.class}, tree="[0]")
    private Output<String> keyRingId;

    /**
     * @return The key ring ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}` or
     * `{location_name}/{key_ring_name}`. In the second form, the provider&#39;s
     * project setting will be used as a fallback.
     * 
     */
    public Output<String> keyRingId() {
        return this.keyRingId;
    }
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    @Export(name="members", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> members;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.kms.KeyRingIAMBinding` can be used per role. Note that custom roles must be of the format
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
    public KeyRingIAMBinding(java.lang.String name) {
        this(name, KeyRingIAMBindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KeyRingIAMBinding(java.lang.String name, KeyRingIAMBindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KeyRingIAMBinding(java.lang.String name, KeyRingIAMBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private KeyRingIAMBinding(java.lang.String name, Output<java.lang.String> id, @Nullable KeyRingIAMBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/keyRingIAMBinding:KeyRingIAMBinding", name, state, makeResourceOptions(options, id), false);
    }

    private static KeyRingIAMBindingArgs makeArgs(KeyRingIAMBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? KeyRingIAMBindingArgs.Empty : args;
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
    public static KeyRingIAMBinding get(java.lang.String name, Output<java.lang.String> id, @Nullable KeyRingIAMBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KeyRingIAMBinding(name, id, state, options);
    }
}
