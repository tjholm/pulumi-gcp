// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.WebBackendServiceIamPolicyArgs;
import com.pulumi.gcp.iap.inputs.WebBackendServiceIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy WebBackendService. Each of these resources serves a different use case:
 * 
 * * `gcp.iap.WebBackendServiceIamPolicy`: Authoritative. Sets the IAM policy for the webbackendservice and replaces any existing policy already attached.
 * * `gcp.iap.WebBackendServiceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webbackendservice are preserved.
 * * `gcp.iap.WebBackendServiceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webbackendservice are preserved.
 * 
 * &gt; **Note:** `gcp.iap.WebBackendServiceIamPolicy` **cannot** be used in conjunction with `gcp.iap.WebBackendServiceIamBinding` and `gcp.iap.WebBackendServiceIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.iap.WebBackendServiceIamBinding` resources **can be** used in conjunction with `gcp.iap.WebBackendServiceIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## google\_iap\_web\_backend\_service\_iam\_policy
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.iap.WebBackendServiceIamPolicy;
 * import com.pulumi.gcp.iap.WebBackendServiceIamPolicyArgs;
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
 *                 .role(&#34;roles/iap.httpsResourceAccessor&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new WebBackendServiceIamPolicy(&#34;policy&#34;, WebBackendServiceIamPolicyArgs.builder()        
 *             .project(google_compute_backend_service.default().project())
 *             .webBackendService(google_compute_backend_service.default().name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With IAM Conditions:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetIAMPolicyArgs;
 * import com.pulumi.gcp.iap.WebBackendServiceIamPolicy;
 * import com.pulumi.gcp.iap.WebBackendServiceIamPolicyArgs;
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
 *                 .role(&#34;roles/iap.httpsResourceAccessor&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title(&#34;expires_after_2019_12_31&#34;)
 *                     .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                     .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new WebBackendServiceIamPolicy(&#34;policy&#34;, WebBackendServiceIamPolicyArgs.builder()        
 *             .project(google_compute_backend_service.default().project())
 *             .webBackendService(google_compute_backend_service.default().name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## google\_iap\_web\_backend\_service\_iam\_binding
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.WebBackendServiceIamBinding;
 * import com.pulumi.gcp.iap.WebBackendServiceIamBindingArgs;
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
 *         var binding = new WebBackendServiceIamBinding(&#34;binding&#34;, WebBackendServiceIamBindingArgs.builder()        
 *             .project(google_compute_backend_service.default().project())
 *             .webBackendService(google_compute_backend_service.default().name())
 *             .role(&#34;roles/iap.httpsResourceAccessor&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With IAM Conditions:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.WebBackendServiceIamBinding;
 * import com.pulumi.gcp.iap.WebBackendServiceIamBindingArgs;
 * import com.pulumi.gcp.iap.inputs.WebBackendServiceIamBindingConditionArgs;
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
 *         var binding = new WebBackendServiceIamBinding(&#34;binding&#34;, WebBackendServiceIamBindingArgs.builder()        
 *             .project(google_compute_backend_service.default().project())
 *             .webBackendService(google_compute_backend_service.default().name())
 *             .role(&#34;roles/iap.httpsResourceAccessor&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .condition(WebBackendServiceIamBindingConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## google\_iap\_web\_backend\_service\_iam\_member
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.WebBackendServiceIamMember;
 * import com.pulumi.gcp.iap.WebBackendServiceIamMemberArgs;
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
 *         var member = new WebBackendServiceIamMember(&#34;member&#34;, WebBackendServiceIamMemberArgs.builder()        
 *             .project(google_compute_backend_service.default().project())
 *             .webBackendService(google_compute_backend_service.default().name())
 *             .role(&#34;roles/iap.httpsResourceAccessor&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * With IAM Conditions:
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.WebBackendServiceIamMember;
 * import com.pulumi.gcp.iap.WebBackendServiceIamMemberArgs;
 * import com.pulumi.gcp.iap.inputs.WebBackendServiceIamMemberConditionArgs;
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
 *         var member = new WebBackendServiceIamMember(&#34;member&#34;, WebBackendServiceIamMemberArgs.builder()        
 *             .project(google_compute_backend_service.default().project())
 *             .webBackendService(google_compute_backend_service.default().name())
 *             .role(&#34;roles/iap.httpsResourceAccessor&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .condition(WebBackendServiceIamMemberConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms* projects/{{project}}/iap_web/compute/services/{{name}} * {{project}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webbackendservice IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor &#34;projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor user:jane@example.com&#34;
 * ```
 * 
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor &#34;projects/{{project}}/iap_web/compute/services/{{web_backend_service}} roles/iap.httpsResourceAccessor&#34;
 * ```
 * 
 *  IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy editor projects/{{project}}/iap_web/compute/services/{{web_backend_service}}
 * ```
 * 
 *  -&gt; **Custom Roles**If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy")
public class WebBackendServiceIamPolicy extends com.pulumi.resources.CustomResource {
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
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="webBackendService", type=String.class, parameters={})
    private Output<String> webBackendService;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> webBackendService() {
        return this.webBackendService;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WebBackendServiceIamPolicy(String name) {
        this(name, WebBackendServiceIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WebBackendServiceIamPolicy(String name, WebBackendServiceIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WebBackendServiceIamPolicy(String name, WebBackendServiceIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, args == null ? WebBackendServiceIamPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private WebBackendServiceIamPolicy(String name, Output<String> id, @Nullable WebBackendServiceIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/webBackendServiceIamPolicy:WebBackendServiceIamPolicy", name, state, makeResourceOptions(options, id));
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
    public static WebBackendServiceIamPolicy get(String name, Output<String> id, @Nullable WebBackendServiceIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WebBackendServiceIamPolicy(name, id, state, options);
    }
}
