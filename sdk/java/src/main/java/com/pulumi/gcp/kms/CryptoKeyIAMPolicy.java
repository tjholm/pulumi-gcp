// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.kms.CryptoKeyIAMPolicyArgs;
import com.pulumi.gcp.kms.inputs.CryptoKeyIAMPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for KMS crypto key. Each of these resources serves a different use case:
 * 
 * * `gcp.kms.CryptoKeyIAMPolicy`: Authoritative. Sets the IAM policy for the crypto key and replaces any existing policy already attached.
 * * `gcp.kms.CryptoKeyIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the crypto key are preserved.
 * * `gcp.kms.CryptoKeyIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the crypto key are preserved.
 * 
 * &gt; **Note:** `gcp.kms.CryptoKeyIAMPolicy` **cannot** be used in conjunction with `gcp.kms.CryptoKeyIAMBinding` and `gcp.kms.CryptoKeyIAMMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.kms.CryptoKeyIAMBinding` resources **can be** used in conjunction with `gcp.kms.CryptoKeyIAMMember` resources **only if** they do not grant privilege to the same role.
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
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.kms.CryptoKeyIAMPolicy;
 * import com.pulumi.gcp.kms.CryptoKeyIAMPolicyArgs;
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
 *         var key = new CryptoKey("key", CryptoKeyArgs.builder()
 *             .name("crypto-key-example")
 *             .keyRing(keyring.id())
 *             .rotationPeriod("7776000s")
 *             .build());
 * 
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/cloudkms.cryptoKeyEncrypter")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var cryptoKey = new CryptoKeyIAMPolicy("cryptoKey", CryptoKeyIAMPolicyArgs.builder()
 *             .cryptoKeyId(key.id())
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
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/cloudkms.cryptoKeyEncrypter")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title("expires_after_2019_12_31")
 *                     .description("Expiring at midnight of 2019-12-31")
 *                     .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
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
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBinding;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBindingArgs;
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
 *         var cryptoKey = new CryptoKeyIAMBinding("cryptoKey", CryptoKeyIAMBindingArgs.builder()
 *             .cryptoKeyId(key.id())
 *             .role("roles/cloudkms.cryptoKeyEncrypter")
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
 * import com.pulumi.gcp.kms.CryptoKeyIAMBinding;
 * import com.pulumi.gcp.kms.CryptoKeyIAMBindingArgs;
 * import com.pulumi.gcp.kms.inputs.CryptoKeyIAMBindingConditionArgs;
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
 *         var cryptoKey = new CryptoKeyIAMBinding("cryptoKey", CryptoKeyIAMBindingArgs.builder()
 *             .cryptoKeyId(key.id())
 *             .role("roles/cloudkms.cryptoKeyEncrypter")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .condition(CryptoKeyIAMBindingConditionArgs.builder()
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
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
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
 *         var cryptoKey = new CryptoKeyIAMMember("cryptoKey", CryptoKeyIAMMemberArgs.builder()
 *             .cryptoKeyId(key.id())
 *             .role("roles/cloudkms.cryptoKeyEncrypter")
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
 * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
 * import com.pulumi.gcp.kms.inputs.CryptoKeyIAMMemberConditionArgs;
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
 *         var cryptoKey = new CryptoKeyIAMMember("cryptoKey", CryptoKeyIAMMemberArgs.builder()
 *             .cryptoKeyId(key.id())
 *             .role("roles/cloudkms.cryptoKeyEncrypter")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .condition(CryptoKeyIAMMemberConditionArgs.builder()
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
 * IAM policy imports use the identifier of the KMS crypto key only. For example:
 * 
 * * `{{project_id}}/{{location}}/{{key_ring_name}}/{{crypto_key_name}}`
 * 
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 * 
 * tf
 * 
 * import {
 * 
 *   id = &#34;{{project_id}}/{{location}}/{{key_ring_name}}/{{crypto_key_name}}&#34;
 * 
 *   to = google_kms_crypto_key_iam_policy.default
 * 
 * }
 * 
 * The `pulumi import` command can also be used:
 * 
 * ```sh
 * $ pulumi import gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy default {{project_id}}/{{location}}/{{key_ring_name}}/{{crypto_key_name}}
 * ```
 * 
 */
@ResourceType(type="gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy")
public class CryptoKeyIAMPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
     * the provider&#39;s project setting will be used as a fallback.
     * 
     */
    @Export(name="cryptoKeyId", refs={String.class}, tree="[0]")
    private Output<String> cryptoKeyId;

    /**
     * @return The crypto key ID, in the form
     * `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` or
     * `{location_name}/{key_ring_name}/{crypto_key_name}`. In the second form,
     * the provider&#39;s project setting will be used as a fallback.
     * 
     */
    public Output<String> cryptoKeyId() {
        return this.cryptoKeyId;
    }
    /**
     * (Computed) The etag of the project&#39;s IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the project&#39;s IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
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
    public CryptoKeyIAMPolicy(java.lang.String name) {
        this(name, CryptoKeyIAMPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CryptoKeyIAMPolicy(java.lang.String name, CryptoKeyIAMPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CryptoKeyIAMPolicy(java.lang.String name, CryptoKeyIAMPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CryptoKeyIAMPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable CryptoKeyIAMPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:kms/cryptoKeyIAMPolicy:CryptoKeyIAMPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static CryptoKeyIAMPolicyArgs makeArgs(CryptoKeyIAMPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CryptoKeyIAMPolicyArgs.Empty : args;
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
    public static CryptoKeyIAMPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable CryptoKeyIAMPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CryptoKeyIAMPolicy(name, id, state, options);
    }
}
