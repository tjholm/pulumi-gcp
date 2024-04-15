// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InstanceIAMPolicyArgs;
import com.pulumi.gcp.compute.inputs.InstanceIAMPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Compute Engine Instance. Each of these resources serves a different use case:
 * 
 * * `gcp.compute.InstanceIAMPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
 * * `gcp.compute.InstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
 * * `gcp.compute.InstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.compute.InstanceIAMPolicy`: Retrieves the IAM policy for the instance
 * 
 * &gt; **Note:** `gcp.compute.InstanceIAMPolicy` **cannot** be used in conjunction with `gcp.compute.InstanceIAMBinding` and `gcp.compute.InstanceIAMMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.compute.InstanceIAMBinding` resources **can be** used in conjunction with `gcp.compute.InstanceIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
 * 
 * ## google\_compute\_instance\_iam\_policy
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
 * import com.pulumi.gcp.compute.InstanceIAMPolicy;
 * import com.pulumi.gcp.compute.InstanceIAMPolicyArgs;
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
 *                 .role(&#34;roles/compute.osLogin&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new InstanceIAMPolicy(&#34;policy&#34;, InstanceIAMPolicyArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
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
 * import com.pulumi.gcp.compute.InstanceIAMPolicy;
 * import com.pulumi.gcp.compute.InstanceIAMPolicyArgs;
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
 *                 .role(&#34;roles/compute.osLogin&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title(&#34;expires_after_2019_12_31&#34;)
 *                     .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                     .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new InstanceIAMPolicy(&#34;policy&#34;, InstanceIAMPolicyArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_compute\_instance\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMBinding;
 * import com.pulumi.gcp.compute.InstanceIAMBindingArgs;
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
 *         var binding = new InstanceIAMBinding(&#34;binding&#34;, InstanceIAMBindingArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMBinding;
 * import com.pulumi.gcp.compute.InstanceIAMBindingArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceIAMBindingConditionArgs;
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
 *         var binding = new InstanceIAMBinding(&#34;binding&#34;, InstanceIAMBindingArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .condition(InstanceIAMBindingConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_compute\_instance\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMMember;
 * import com.pulumi.gcp.compute.InstanceIAMMemberArgs;
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
 *         var member = new InstanceIAMMember(&#34;member&#34;, InstanceIAMMemberArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMMember;
 * import com.pulumi.gcp.compute.InstanceIAMMemberArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceIAMMemberConditionArgs;
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
 *         var member = new InstanceIAMMember(&#34;member&#34;, InstanceIAMMemberArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .condition(InstanceIAMMemberConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## google\_compute\_instance\_iam\_policy
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
 * import com.pulumi.gcp.compute.InstanceIAMPolicy;
 * import com.pulumi.gcp.compute.InstanceIAMPolicyArgs;
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
 *                 .role(&#34;roles/compute.osLogin&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .build())
 *             .build());
 * 
 *         var policy = new InstanceIAMPolicy(&#34;policy&#34;, InstanceIAMPolicyArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
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
 * import com.pulumi.gcp.compute.InstanceIAMPolicy;
 * import com.pulumi.gcp.compute.InstanceIAMPolicyArgs;
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
 *                 .role(&#34;roles/compute.osLogin&#34;)
 *                 .members(&#34;user:jane@example.com&#34;)
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title(&#34;expires_after_2019_12_31&#34;)
 *                     .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                     .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new InstanceIAMPolicy(&#34;policy&#34;, InstanceIAMPolicyArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -&gt; getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_compute\_instance\_iam\_binding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMBinding;
 * import com.pulumi.gcp.compute.InstanceIAMBindingArgs;
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
 *         var binding = new InstanceIAMBinding(&#34;binding&#34;, InstanceIAMBindingArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMBinding;
 * import com.pulumi.gcp.compute.InstanceIAMBindingArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceIAMBindingConditionArgs;
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
 *         var binding = new InstanceIAMBinding(&#34;binding&#34;, InstanceIAMBindingArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .members(&#34;user:jane@example.com&#34;)
 *             .condition(InstanceIAMBindingConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## google\_compute\_instance\_iam\_member
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMMember;
 * import com.pulumi.gcp.compute.InstanceIAMMemberArgs;
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
 *         var member = new InstanceIAMMember(&#34;member&#34;, InstanceIAMMemberArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.InstanceIAMMember;
 * import com.pulumi.gcp.compute.InstanceIAMMemberArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceIAMMemberConditionArgs;
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
 *         var member = new InstanceIAMMember(&#34;member&#34;, InstanceIAMMemberArgs.builder()        
 *             .project(default_.project())
 *             .zone(default_.zone())
 *             .instanceName(default_.name())
 *             .role(&#34;roles/compute.osLogin&#34;)
 *             .member(&#34;user:jane@example.com&#34;)
 *             .condition(InstanceIAMMemberConditionArgs.builder()
 *                 .title(&#34;expires_after_2019_12_31&#34;)
 *                 .description(&#34;Expiring at midnight of 2019-12-31&#34;)
 *                 .expression(&#34;request.time &lt; timestamp(\&#34;2020-01-01T00:00:00Z\&#34;)&#34;)
 *                 .build())
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
 * * projects/{{project}}/zones/{{zone}}/instances/{{name}}
 * 
 * * {{project}}/{{zone}}/{{name}}
 * 
 * * {{zone}}/{{name}}
 * 
 * * {{name}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Compute Engine instance IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/instanceIAMPolicy:InstanceIAMPolicy editor &#34;projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin user:jane@example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/instanceIAMPolicy:InstanceIAMPolicy editor &#34;projects/{{project}}/zones/{{zone}}/instances/{{instance}} roles/compute.osLogin&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:compute/instanceIAMPolicy:InstanceIAMPolicy editor projects/{{project}}/zones/{{zone}}/instances/{{instance}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:compute/instanceIAMPolicy:InstanceIAMPolicy")
public class InstanceIAMPolicy extends com.pulumi.resources.CustomResource {
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
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
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
     * A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
     * zone is specified, it is taken from the provider configuration.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
     * zone is specified, it is taken from the provider configuration.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceIAMPolicy(String name) {
        this(name, InstanceIAMPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceIAMPolicy(String name, InstanceIAMPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceIAMPolicy(String name, InstanceIAMPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceIAMPolicy:InstanceIAMPolicy", name, args == null ? InstanceIAMPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceIAMPolicy(String name, Output<String> id, @Nullable InstanceIAMPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/instanceIAMPolicy:InstanceIAMPolicy", name, state, makeResourceOptions(options, id));
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
    public static InstanceIAMPolicy get(String name, Output<String> id, @Nullable InstanceIAMPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceIAMPolicy(name, id, state, options);
    }
}
