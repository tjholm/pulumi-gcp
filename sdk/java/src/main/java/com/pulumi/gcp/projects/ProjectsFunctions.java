// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.projects;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.projects.inputs.GetOrganizationPolicyArgs;
import com.pulumi.gcp.projects.inputs.GetOrganizationPolicyPlainArgs;
import com.pulumi.gcp.projects.inputs.GetProjectArgs;
import com.pulumi.gcp.projects.inputs.GetProjectPlainArgs;
import com.pulumi.gcp.projects.outputs.GetOrganizationPolicyResult;
import com.pulumi.gcp.projects.outputs.GetProjectResult;
import java.util.concurrent.CompletableFuture;

public final class ProjectsFunctions {
    /**
     * Allows management of Organization policies for a Google Project. For more information see
     * [the official
     * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
     * 
     * ## Example Usage
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
     *         final var policy = Output.of(ProjectsFunctions.getOrganizationPolicy(GetOrganizationPolicyArgs.builder()
     *             .project(&#34;project-id&#34;)
     *             .constraint(&#34;constraints/serviceuser.services&#34;)
     *             .build()));
     * 
     *         ctx.export(&#34;version&#34;, policy.apply(getOrganizationPolicyResult -&gt; getOrganizationPolicyResult.version()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetOrganizationPolicyResult> getOrganizationPolicy(GetOrganizationPolicyArgs args) {
        return getOrganizationPolicy(args, InvokeOptions.Empty);
    }
    /**
     * Allows management of Organization policies for a Google Project. For more information see
     * [the official
     * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
     * 
     * ## Example Usage
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
     *         final var policy = Output.of(ProjectsFunctions.getOrganizationPolicy(GetOrganizationPolicyArgs.builder()
     *             .project(&#34;project-id&#34;)
     *             .constraint(&#34;constraints/serviceuser.services&#34;)
     *             .build()));
     * 
     *         ctx.export(&#34;version&#34;, policy.apply(getOrganizationPolicyResult -&gt; getOrganizationPolicyResult.version()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetOrganizationPolicyResult> getOrganizationPolicyPlain(GetOrganizationPolicyPlainArgs args) {
        return getOrganizationPolicyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Allows management of Organization policies for a Google Project. For more information see
     * [the official
     * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
     * 
     * ## Example Usage
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
     *         final var policy = Output.of(ProjectsFunctions.getOrganizationPolicy(GetOrganizationPolicyArgs.builder()
     *             .project(&#34;project-id&#34;)
     *             .constraint(&#34;constraints/serviceuser.services&#34;)
     *             .build()));
     * 
     *         ctx.export(&#34;version&#34;, policy.apply(getOrganizationPolicyResult -&gt; getOrganizationPolicyResult.version()));
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetOrganizationPolicyResult> getOrganizationPolicy(GetOrganizationPolicyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:projects/getOrganizationPolicy:getOrganizationPolicy", TypeShape.of(GetOrganizationPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Allows management of Organization policies for a Google Project. For more information see
     * [the official
     * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
     * 
     * ## Example Usage
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
     *         final var policy = Output.of(ProjectsFunctions.getOrganizationPolicy(GetOrganizationPolicyArgs.builder()
     *             .project(&#34;project-id&#34;)
     *             .constraint(&#34;constraints/serviceuser.services&#34;)
     *             .build()));
     * 
     *         ctx.export(&#34;version&#34;, policy.apply(getOrganizationPolicyResult -&gt; getOrganizationPolicyResult.version()));
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetOrganizationPolicyResult> getOrganizationPolicyPlain(GetOrganizationPolicyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:projects/getOrganizationPolicy:getOrganizationPolicy", TypeShape.of(GetOrganizationPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Retrieve information about a set of projects based on a filter. See the
     * [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
     * for more details.
     * 
     * ## Example Usage
     * ### Searching For Projects About To Be Deleted In An Org
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
     *         final var my-org-projects = Output.of(ProjectsFunctions.getProject(GetProjectArgs.builder()
     *             .filter(&#34;parent.id:012345678910 lifecycleState:DELETE_REQUESTED&#34;)
     *             .build()));
     * 
     *         final var deletion-candidate = Output.of(OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(my_org_projects.projects()[0].projectId())
     *             .build()));
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetProjectResult> getProject(GetProjectArgs args) {
        return getProject(args, InvokeOptions.Empty);
    }
    /**
     * Retrieve information about a set of projects based on a filter. See the
     * [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
     * for more details.
     * 
     * ## Example Usage
     * ### Searching For Projects About To Be Deleted In An Org
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
     *         final var my-org-projects = Output.of(ProjectsFunctions.getProject(GetProjectArgs.builder()
     *             .filter(&#34;parent.id:012345678910 lifecycleState:DELETE_REQUESTED&#34;)
     *             .build()));
     * 
     *         final var deletion-candidate = Output.of(OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(my_org_projects.projects()[0].projectId())
     *             .build()));
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetProjectResult> getProjectPlain(GetProjectPlainArgs args) {
        return getProjectPlain(args, InvokeOptions.Empty);
    }
    /**
     * Retrieve information about a set of projects based on a filter. See the
     * [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
     * for more details.
     * 
     * ## Example Usage
     * ### Searching For Projects About To Be Deleted In An Org
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
     *         final var my-org-projects = Output.of(ProjectsFunctions.getProject(GetProjectArgs.builder()
     *             .filter(&#34;parent.id:012345678910 lifecycleState:DELETE_REQUESTED&#34;)
     *             .build()));
     * 
     *         final var deletion-candidate = Output.of(OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(my_org_projects.projects()[0].projectId())
     *             .build()));
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetProjectResult> getProject(GetProjectArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:projects/getProject:getProject", TypeShape.of(GetProjectResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Retrieve information about a set of projects based on a filter. See the
     * [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
     * for more details.
     * 
     * ## Example Usage
     * ### Searching For Projects About To Be Deleted In An Org
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
     *         final var my-org-projects = Output.of(ProjectsFunctions.getProject(GetProjectArgs.builder()
     *             .filter(&#34;parent.id:012345678910 lifecycleState:DELETE_REQUESTED&#34;)
     *             .build()));
     * 
     *         final var deletion-candidate = Output.of(OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(my_org_projects.projects()[0].projectId())
     *             .build()));
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetProjectResult> getProjectPlain(GetProjectPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:projects/getProject:getProject", TypeShape.of(GetProjectResult.class), args, Utilities.withVersion(options));
    }
}
