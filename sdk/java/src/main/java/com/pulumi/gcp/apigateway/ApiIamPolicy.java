// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigateway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigateway.ApiIamPolicyArgs;
import com.pulumi.gcp.apigateway.inputs.ApiIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for API Gateway Api. Each of these resources serves a different use case:
 * 
 * * `gcp.apigateway.ApiIamPolicy`: Authoritative. Sets the IAM policy for the api and replaces any existing policy already attached.
 * * `gcp.apigateway.ApiIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the api are preserved.
 * * `gcp.apigateway.ApiIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the api are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.apigateway.ApiIamPolicy`: Retrieves the IAM policy for the api
 * 
 * &gt; **Note:** `gcp.apigateway.ApiIamPolicy` **cannot** be used in conjunction with `gcp.apigateway.ApiIamBinding` and `gcp.apigateway.ApiIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.apigateway.ApiIamBinding` resources **can be** used in conjunction with `gcp.apigateway.ApiIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_api\_gateway\_api\_iam\_policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.apigateway.ApiIamPolicy;
 * import com.pulumi.gcp.apigateway.ApiIamPolicyArgs;
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
 *                 .role(&#34;roles/apigateway.viewer&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new ApiIamPolicy(&#34;policy&#34;, ApiIamPolicyArgs.builder()        
 *             .project(api.project())
 *             .api(api.apiId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_api\_gateway\_api\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigateway.ApiIamBinding;
 * import com.pulumi.gcp.apigateway.ApiIamBindingArgs;
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
 *         var binding = new ApiIamBinding(&#34;binding&#34;, ApiIamBindingArgs.builder()        
 *             .project(api.project())
 *             .api(api.apiId())
 *             .role(&#34;roles/apigateway.viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_api\_gateway\_api\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigateway.ApiIamMember;
 * import com.pulumi.gcp.apigateway.ApiIamMemberArgs;
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
 *         var member = new ApiIamMember(&#34;member&#34;, ApiIamMemberArgs.builder()        
 *             .project(api.project())
 *             .api(api.apiId())
 *             .role(&#34;roles/apigateway.viewer&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_api\_gateway\_api\_iam\_policy
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.apigateway.ApiIamPolicy;
 * import com.pulumi.gcp.apigateway.ApiIamPolicyArgs;
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
 *                 .role(&#34;roles/apigateway.viewer&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new ApiIamPolicy(&#34;policy&#34;, ApiIamPolicyArgs.builder()        
 *             .project(api.project())
 *             .api(api.apiId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_api\_gateway\_api\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigateway.ApiIamBinding;
 * import com.pulumi.gcp.apigateway.ApiIamBindingArgs;
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
 *         var binding = new ApiIamBinding(&#34;binding&#34;, ApiIamBindingArgs.builder()        
 *             .project(api.project())
 *             .api(api.apiId())
 *             .role(&#34;roles/apigateway.viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_api\_gateway\_api\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigateway.ApiIamMember;
 * import com.pulumi.gcp.apigateway.ApiIamMemberArgs;
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
 *         var member = new ApiIamMember(&#34;member&#34;, ApiIamMemberArgs.builder()        
 *             .project(api.project())
 *             .api(api.apiId())
 *             .role(&#34;roles/apigateway.viewer&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms:
 * 
 * * projects/{{project}}/locations/global/apis/{{api}}
 * 
 * * {{project}}/{{api}}
 * 
 * * {{api}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * API Gateway api IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor &#34;projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer user:jane@example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor &#34;projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor projects/{{project}}/locations/global/apis/{{api}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:apigateway/apiIamPolicy:ApiIamPolicy")
public class ApiIamPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="api", refs={String.class}, tree="[0]")
    private Output<String> api;

    public Output<String> api() {
        return this.api;
    }
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApiIamPolicy(String name) {
        this(name, ApiIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApiIamPolicy(String name, ApiIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApiIamPolicy(String name, ApiIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/apiIamPolicy:ApiIamPolicy", name, args == null ? ApiIamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApiIamPolicy(String name, Output<String> id, @Nullable ApiIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/apiIamPolicy:ApiIamPolicy", name, state, makeResourceOptions(options, id));
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
    public static ApiIamPolicy get(String name, Output<String> id, @Nullable ApiIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApiIamPolicy(name, id, state, options);
    }
}
