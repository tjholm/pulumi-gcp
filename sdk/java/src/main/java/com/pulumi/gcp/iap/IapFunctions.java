// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.inputs.GetClientArgs;
import com.pulumi.gcp.iap.inputs.GetClientPlainArgs;
import com.pulumi.gcp.iap.outputs.GetClientResult;
import java.util.concurrent.CompletableFuture;

public final class IapFunctions {
    /**
     * Get info about a Google Cloud IAP Client.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.organizations.OrganizationsFunctions;
     * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
     * import com.pulumi.gcp.iap.IapFunctions;
     * import com.pulumi.gcp.iap.inputs.GetClientArgs;
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
     *         final var project = OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(&#34;foobar&#34;)
     *             .build());
     * 
     *         final var projectClient = IapFunctions.getClient(GetClientArgs.builder()
     *             .brand(String.format(&#34;projects/%s/brands/[BRAND_NUMBER]&#34;, project.applyValue(getProjectResult -&gt; getProjectResult.number())))
     *             .clientId(FOO.apps().googleusercontent().com())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetClientResult> getClient(GetClientArgs args) {
        return getClient(args, InvokeOptions.Empty);
    }
    /**
     * Get info about a Google Cloud IAP Client.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.organizations.OrganizationsFunctions;
     * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
     * import com.pulumi.gcp.iap.IapFunctions;
     * import com.pulumi.gcp.iap.inputs.GetClientArgs;
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
     *         final var project = OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(&#34;foobar&#34;)
     *             .build());
     * 
     *         final var projectClient = IapFunctions.getClient(GetClientArgs.builder()
     *             .brand(String.format(&#34;projects/%s/brands/[BRAND_NUMBER]&#34;, project.applyValue(getProjectResult -&gt; getProjectResult.number())))
     *             .clientId(FOO.apps().googleusercontent().com())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetClientResult> getClientPlain(GetClientPlainArgs args) {
        return getClientPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get info about a Google Cloud IAP Client.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.organizations.OrganizationsFunctions;
     * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
     * import com.pulumi.gcp.iap.IapFunctions;
     * import com.pulumi.gcp.iap.inputs.GetClientArgs;
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
     *         final var project = OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(&#34;foobar&#34;)
     *             .build());
     * 
     *         final var projectClient = IapFunctions.getClient(GetClientArgs.builder()
     *             .brand(String.format(&#34;projects/%s/brands/[BRAND_NUMBER]&#34;, project.applyValue(getProjectResult -&gt; getProjectResult.number())))
     *             .clientId(FOO.apps().googleusercontent().com())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetClientResult> getClient(GetClientArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:iap/getClient:getClient", TypeShape.of(GetClientResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get info about a Google Cloud IAP Client.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.organizations.OrganizationsFunctions;
     * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
     * import com.pulumi.gcp.iap.IapFunctions;
     * import com.pulumi.gcp.iap.inputs.GetClientArgs;
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
     *         final var project = OrganizationsFunctions.getProject(GetProjectArgs.builder()
     *             .projectId(&#34;foobar&#34;)
     *             .build());
     * 
     *         final var projectClient = IapFunctions.getClient(GetClientArgs.builder()
     *             .brand(String.format(&#34;projects/%s/brands/[BRAND_NUMBER]&#34;, project.applyValue(getProjectResult -&gt; getProjectResult.number())))
     *             .clientId(FOO.apps().googleusercontent().com())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetClientResult> getClientPlain(GetClientPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:iap/getClient:getClient", TypeShape.of(GetClientResult.class), args, Utilities.withVersion(options));
    }
}
