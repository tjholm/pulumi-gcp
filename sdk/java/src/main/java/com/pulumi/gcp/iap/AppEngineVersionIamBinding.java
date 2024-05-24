// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.AppEngineVersionIamBindingArgs;
import com.pulumi.gcp.iap.inputs.AppEngineVersionIamBindingState;
import com.pulumi.gcp.iap.outputs.AppEngineVersionIamBindingCondition;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineVersion. Each of these resources serves a different use case:
 * 
 * * `gcp.iap.AppEngineVersionIamPolicy`: Authoritative. Sets the IAM policy for the appengineversion and replaces any existing policy already attached.
 * * `gcp.iap.AppEngineVersionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineversion are preserved.
 * * `gcp.iap.AppEngineVersionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineversion are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.iap.AppEngineVersionIamPolicy`: Retrieves the IAM policy for the appengineversion
 * 
 * &gt; **Note:** `gcp.iap.AppEngineVersionIamPolicy` **cannot** be used in conjunction with `gcp.iap.AppEngineVersionIamBinding` and `gcp.iap.AppEngineVersionIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.iap.AppEngineVersionIamBinding` resources **can be** used in conjunction with `gcp.iap.AppEngineVersionIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
 * 
 * ## google\_iap\_app\_engine\_version\_iam\_policy
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicy;
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicyArgs;
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
 *                 .role("roles/iap.httpsResourceAccessor")
 *                 .members("user:jane{@literal @}example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new AppEngineVersionIamPolicy("policy", AppEngineVersionIamPolicyArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicy;
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicyArgs;
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
 *                 .role("roles/iap.httpsResourceAccessor")
 *                 .members("user:jane{@literal @}example.com")
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title("expires_after_2019_12_31")
 *                     .description("Expiring at midnight of 2019-12-31")
 *                     .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new AppEngineVersionIamPolicy("policy", AppEngineVersionIamPolicyArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_iap\_app\_engine\_version\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.AppEngineVersionIamBinding;
 * import com.pulumi.gcp.iap.AppEngineVersionIamBindingArgs;
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
 *         var binding = new AppEngineVersionIamBinding("binding", AppEngineVersionIamBindingArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .members("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamBinding;
 * import com.pulumi.gcp.iap.AppEngineVersionIamBindingArgs;
 * import com.pulumi.gcp.iap.inputs.AppEngineVersionIamBindingConditionArgs;
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
 *         var binding = new AppEngineVersionIamBinding("binding", AppEngineVersionIamBindingArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .members("user:jane{@literal @}example.com")
 *             .condition(AppEngineVersionIamBindingConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_iap\_app\_engine\_version\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.AppEngineVersionIamMember;
 * import com.pulumi.gcp.iap.AppEngineVersionIamMemberArgs;
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
 *         var member = new AppEngineVersionIamMember("member", AppEngineVersionIamMemberArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .member("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamMember;
 * import com.pulumi.gcp.iap.AppEngineVersionIamMemberArgs;
 * import com.pulumi.gcp.iap.inputs.AppEngineVersionIamMemberConditionArgs;
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
 *         var member = new AppEngineVersionIamMember("member", AppEngineVersionIamMemberArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .member("user:jane{@literal @}example.com")
 *             .condition(AppEngineVersionIamMemberConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_iap\_app\_engine\_version\_iam\_policy
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicy;
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicyArgs;
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
 *                 .role("roles/iap.httpsResourceAccessor")
 *                 .members("user:jane{@literal @}example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new AppEngineVersionIamPolicy("policy", AppEngineVersionIamPolicyArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicy;
 * import com.pulumi.gcp.iap.AppEngineVersionIamPolicyArgs;
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
 *                 .role("roles/iap.httpsResourceAccessor")
 *                 .members("user:jane{@literal @}example.com")
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title("expires_after_2019_12_31")
 *                     .description("Expiring at midnight of 2019-12-31")
 *                     .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new AppEngineVersionIamPolicy("policy", AppEngineVersionIamPolicyArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_iap\_app\_engine\_version\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.AppEngineVersionIamBinding;
 * import com.pulumi.gcp.iap.AppEngineVersionIamBindingArgs;
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
 *         var binding = new AppEngineVersionIamBinding("binding", AppEngineVersionIamBindingArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .members("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamBinding;
 * import com.pulumi.gcp.iap.AppEngineVersionIamBindingArgs;
 * import com.pulumi.gcp.iap.inputs.AppEngineVersionIamBindingConditionArgs;
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
 *         var binding = new AppEngineVersionIamBinding("binding", AppEngineVersionIamBindingArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .members("user:jane{@literal @}example.com")
 *             .condition(AppEngineVersionIamBindingConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_iap\_app\_engine\_version\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.AppEngineVersionIamMember;
 * import com.pulumi.gcp.iap.AppEngineVersionIamMemberArgs;
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
 *         var member = new AppEngineVersionIamMember("member", AppEngineVersionIamMemberArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .member("user:jane{@literal @}example.com")
 *             .build());
 * 
 *     }
 * }
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
 * import com.pulumi.gcp.iap.AppEngineVersionIamMember;
 * import com.pulumi.gcp.iap.AppEngineVersionIamMemberArgs;
 * import com.pulumi.gcp.iap.inputs.AppEngineVersionIamMemberConditionArgs;
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
 *         var member = new AppEngineVersionIamMember("member", AppEngineVersionIamMemberArgs.builder()
 *             .project(version.project())
 *             .appId(version.project())
 *             .service(version.service())
 *             .versionId(version.versionId())
 *             .role("roles/iap.httpsResourceAccessor")
 *             .member("user:jane{@literal @}example.com")
 *             .condition(AppEngineVersionIamMemberConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
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
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms:
 * 
 * * projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}}
 * 
 * * {{project}}/{{appId}}/{{service}}/{{versionId}}
 * 
 * * {{appId}}/{{service}}/{{versionId}}
 * 
 * * {{version}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Identity-Aware Proxy appengineversion IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding editor &#34;projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane{@literal @}example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding editor &#34;projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding")
public class AppEngineVersionIamBinding extends com.pulumi.resources.CustomResource {
    /**
     * Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="appId", refs={String.class}, tree="[0]")
    private Output<String> appId;

    /**
     * @return Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Export(name="condition", refs={AppEngineVersionIamBindingCondition.class}, tree="[0]")
    private Output</* @Nullable */ AppEngineVersionIamBindingCondition> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Output<Optional<AppEngineVersionIamBindingCondition>> condition() {
        return Codegen.optional(this.condition);
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
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
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
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    public Output<List<String>> members() {
        return this.members;
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
     * The role that should be applied. Only one
     * `gcp.iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.iap.AppEngineVersionIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="service", refs={String.class}, tree="[0]")
    private Output<String> service;

    /**
     * @return Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> service() {
        return this.service;
    }
    /**
     * Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="versionId", refs={String.class}, tree="[0]")
    private Output<String> versionId;

    /**
     * @return Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> versionId() {
        return this.versionId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppEngineVersionIamBinding(String name) {
        this(name, AppEngineVersionIamBindingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppEngineVersionIamBinding(String name, AppEngineVersionIamBindingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppEngineVersionIamBinding(String name, AppEngineVersionIamBindingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding", name, args == null ? AppEngineVersionIamBindingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppEngineVersionIamBinding(String name, Output<String> id, @Nullable AppEngineVersionIamBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/appEngineVersionIamBinding:AppEngineVersionIamBinding", name, state, makeResourceOptions(options, id));
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
    public static AppEngineVersionIamBinding get(String name, Output<String> id, @Nullable AppEngineVersionIamBindingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppEngineVersionIamBinding(name, id, state, options);
    }
}
