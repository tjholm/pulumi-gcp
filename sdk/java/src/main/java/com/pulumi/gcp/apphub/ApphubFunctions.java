// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apphub;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apphub.inputs.GetDiscoveredServiceArgs;
import com.pulumi.gcp.apphub.inputs.GetDiscoveredServicePlainArgs;
import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadPlainArgs;
import com.pulumi.gcp.apphub.outputs.GetDiscoveredServiceResult;
import com.pulumi.gcp.apphub.outputs.GetDiscoveredWorkloadResult;
import java.util.concurrent.CompletableFuture;

public final class ApphubFunctions {
    /**
     * Get information about a discovered service from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredServiceArgs;
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
     *         final var my-service = ApphubFunctions.getDiscoveredService(GetDiscoveredServiceArgs.builder()
     *             .location(&#34;my-location&#34;)
     *             .serviceUri(&#34;my-service-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDiscoveredServiceResult> getDiscoveredService(GetDiscoveredServiceArgs args) {
        return getDiscoveredService(args, InvokeOptions.Empty);
    }
    /**
     * Get information about a discovered service from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredServiceArgs;
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
     *         final var my-service = ApphubFunctions.getDiscoveredService(GetDiscoveredServiceArgs.builder()
     *             .location(&#34;my-location&#34;)
     *             .serviceUri(&#34;my-service-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDiscoveredServiceResult> getDiscoveredServicePlain(GetDiscoveredServicePlainArgs args) {
        return getDiscoveredServicePlain(args, InvokeOptions.Empty);
    }
    /**
     * Get information about a discovered service from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredServiceArgs;
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
     *         final var my-service = ApphubFunctions.getDiscoveredService(GetDiscoveredServiceArgs.builder()
     *             .location(&#34;my-location&#34;)
     *             .serviceUri(&#34;my-service-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDiscoveredServiceResult> getDiscoveredService(GetDiscoveredServiceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:apphub/getDiscoveredService:getDiscoveredService", TypeShape.of(GetDiscoveredServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get information about a discovered service from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredServiceArgs;
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
     *         final var my-service = ApphubFunctions.getDiscoveredService(GetDiscoveredServiceArgs.builder()
     *             .location(&#34;my-location&#34;)
     *             .serviceUri(&#34;my-service-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDiscoveredServiceResult> getDiscoveredServicePlain(GetDiscoveredServicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:apphub/getDiscoveredService:getDiscoveredService", TypeShape.of(GetDiscoveredServiceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get information about a discovered workload from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
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
     *         final var my-workload = ApphubFunctions.getDiscoveredWorkload(GetDiscoveredWorkloadArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .workloadUri(&#34;my-workload-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDiscoveredWorkloadResult> getDiscoveredWorkload(GetDiscoveredWorkloadArgs args) {
        return getDiscoveredWorkload(args, InvokeOptions.Empty);
    }
    /**
     * Get information about a discovered workload from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
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
     *         final var my-workload = ApphubFunctions.getDiscoveredWorkload(GetDiscoveredWorkloadArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .workloadUri(&#34;my-workload-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDiscoveredWorkloadResult> getDiscoveredWorkloadPlain(GetDiscoveredWorkloadPlainArgs args) {
        return getDiscoveredWorkloadPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get information about a discovered workload from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
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
     *         final var my-workload = ApphubFunctions.getDiscoveredWorkload(GetDiscoveredWorkloadArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .workloadUri(&#34;my-workload-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDiscoveredWorkloadResult> getDiscoveredWorkload(GetDiscoveredWorkloadArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:apphub/getDiscoveredWorkload:getDiscoveredWorkload", TypeShape.of(GetDiscoveredWorkloadResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get information about a discovered workload from its uri.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.apphub.ApphubFunctions;
     * import com.pulumi.gcp.apphub.inputs.GetDiscoveredWorkloadArgs;
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
     *         final var my-workload = ApphubFunctions.getDiscoveredWorkload(GetDiscoveredWorkloadArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .workloadUri(&#34;my-workload-uri&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDiscoveredWorkloadResult> getDiscoveredWorkloadPlain(GetDiscoveredWorkloadPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:apphub/getDiscoveredWorkload:getDiscoveredWorkload", TypeShape.of(GetDiscoveredWorkloadResult.class), args, Utilities.withVersion(options));
    }
}
