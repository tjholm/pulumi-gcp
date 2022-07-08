// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigateway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigateway.ApiConfigIamMemberArgs;
import com.pulumi.gcp.apigateway.inputs.ApiConfigIamMemberState;
import com.pulumi.gcp.apigateway.outputs.ApiConfigIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for API Gateway ApiConfig. Each of these resources serves a different use case:
 * 
 * * `gcp.apigateway.ApiConfigIamPolicy`: Authoritative. Sets the IAM policy for the apiconfig and replaces any existing policy already attached.
 * * `gcp.apigateway.ApiConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the apiconfig are preserved.
 * * `gcp.apigateway.ApiConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the apiconfig are preserved.
 * 
 * &gt; **Note:** `gcp.apigateway.ApiConfigIamPolicy` **cannot** be used in conjunction with `gcp.apigateway.ApiConfigIamBinding` and `gcp.apigateway.ApiConfigIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.apigateway.ApiConfigIamBinding` resources **can be** used in conjunction with `gcp.apigateway.ApiConfigIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_api\_gateway\_api\_config\_iam\_policy
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var admin = Output.of(OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role(&#34;roles/apigateway.viewer&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build()));
 * 
 *         var policy = new ApiConfigIamPolicy(&#34;policy&#34;, ApiConfigIamPolicyArgs.builder()        
 *             .api(google_api_gateway_api_config.api_cfg().api())
 *             .apiConfig(google_api_gateway_api_config.api_cfg().api_config_id())
 *             .policyData(admin.apply(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_api\_gateway\_api\_config\_iam\_binding
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var binding = new ApiConfigIamBinding(&#34;binding&#34;, ApiConfigIamBindingArgs.builder()        
 *             .api(google_api_gateway_api_config.api_cfg().api())
 *             .apiConfig(google_api_gateway_api_config.api_cfg().api_config_id())
 *             .role(&#34;roles/apigateway.viewer&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## google\_api\_gateway\_api\_config\_iam\_member
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var member = new ApiConfigIamMember(&#34;member&#34;, ApiConfigIamMemberArgs.builder()        
 *             .api(google_api_gateway_api_config.api_cfg().api())
 *             .apiConfig(google_api_gateway_api_config.api_cfg().api_config_id())
 *             .role(&#34;roles/apigateway.viewer&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
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
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} * {{project}}/{{api}}/{{api_config}} * {{api}}/{{api_config}} * {{api_config}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway apiconfig IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfigIamMember:ApiConfigIamMember editor &#34;projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfigIamMember:ApiConfigIamMember editor &#34;projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfigIamMember:ApiConfigIamMember editor projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:apigateway/apiConfigIamMember:ApiConfigIamMember")
public class ApiConfigIamMember extends com.pulumi.resources.CustomResource {
    /**
     * The API to attach the config to.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="api", type=String.class, parameters={})
    private Output<String> api;

    /**
     * @return The API to attach the config to.
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> api() {
        return this.api;
    }
    @Export(name="apiConfig", type=String.class, parameters={})
    private Output<String> apiConfig;

    public Output<String> apiConfig() {
        return this.apiConfig;
    }
    @Export(name="condition", type=ApiConfigIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ ApiConfigIamMemberCondition> condition;

    public Output<Optional<ApiConfigIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
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
    @Export(name="member", type=String.class, parameters={})
    private Output<String> member;

    public Output<String> member() {
        return this.member;
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
     * The role that should be applied. Only one
     * `gcp.apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.apigateway.ApiConfigIamBinding` can be used per role. Note that custom roles must be of the format
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
    public ApiConfigIamMember(String name) {
        this(name, ApiConfigIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApiConfigIamMember(String name, ApiConfigIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApiConfigIamMember(String name, ApiConfigIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/apiConfigIamMember:ApiConfigIamMember", name, args == null ? ApiConfigIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApiConfigIamMember(String name, Output<String> id, @Nullable ApiConfigIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/apiConfigIamMember:ApiConfigIamMember", name, state, makeResourceOptions(options, id));
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
    public static ApiConfigIamMember get(String name, Output<String> id, @Nullable ApiConfigIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApiConfigIamMember(name, id, state, options);
    }
}
