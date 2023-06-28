// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vmwareengine.inputs.GetClusterArgs;
import com.pulumi.gcp.vmwareengine.inputs.GetClusterPlainArgs;
import com.pulumi.gcp.vmwareengine.inputs.GetNetworkArgs;
import com.pulumi.gcp.vmwareengine.inputs.GetNetworkPlainArgs;
import com.pulumi.gcp.vmwareengine.inputs.GetPrivateCloudArgs;
import com.pulumi.gcp.vmwareengine.inputs.GetPrivateCloudPlainArgs;
import com.pulumi.gcp.vmwareengine.outputs.GetClusterResult;
import com.pulumi.gcp.vmwareengine.outputs.GetNetworkResult;
import com.pulumi.gcp.vmwareengine.outputs.GetPrivateCloudResult;
import java.util.concurrent.CompletableFuture;

public final class VmwareengineFunctions {
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetClusterArgs;
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
     *         final var myCluster = VmwareengineFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;my-cluster&#34;)
     *             .parent(&#34;project/locations/us-west1-a/privateClouds/my-cloud&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetClusterResult> getCluster(GetClusterArgs args) {
        return getCluster(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetClusterArgs;
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
     *         final var myCluster = VmwareengineFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;my-cluster&#34;)
     *             .parent(&#34;project/locations/us-west1-a/privateClouds/my-cloud&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetClusterResult> getClusterPlain(GetClusterPlainArgs args) {
        return getClusterPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetClusterArgs;
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
     *         final var myCluster = VmwareengineFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;my-cluster&#34;)
     *             .parent(&#34;project/locations/us-west1-a/privateClouds/my-cloud&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetClusterResult> getCluster(GetClusterArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:vmwareengine/getCluster:getCluster", TypeShape.of(GetClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetClusterArgs;
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
     *         final var myCluster = VmwareengineFunctions.getCluster(GetClusterArgs.builder()
     *             .name(&#34;my-cluster&#34;)
     *             .parent(&#34;project/locations/us-west1-a/privateClouds/my-cloud&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetClusterResult> getClusterPlain(GetClusterPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:vmwareengine/getCluster:getCluster", TypeShape.of(GetClusterResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetNetworkArgs;
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
     *         final var myNw = VmwareengineFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name(&#34;us-central1-default&#34;)
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetNetworkResult> getNetwork(GetNetworkArgs args) {
        return getNetwork(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetNetworkArgs;
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
     *         final var myNw = VmwareengineFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name(&#34;us-central1-default&#34;)
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetNetworkResult> getNetworkPlain(GetNetworkPlainArgs args) {
        return getNetworkPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetNetworkArgs;
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
     *         final var myNw = VmwareengineFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name(&#34;us-central1-default&#34;)
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetNetworkResult> getNetwork(GetNetworkArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:vmwareengine/getNetwork:getNetwork", TypeShape.of(GetNetworkResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetNetworkArgs;
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
     *         final var myNw = VmwareengineFunctions.getNetwork(GetNetworkArgs.builder()
     *             .name(&#34;us-central1-default&#34;)
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetNetworkResult> getNetworkPlain(GetNetworkPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:vmwareengine/getNetwork:getNetwork", TypeShape.of(GetNetworkResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetPrivateCloudArgs;
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
     *         final var myPc = VmwareengineFunctions.getPrivateCloud(GetPrivateCloudArgs.builder()
     *             .name(&#34;my-pc&#34;)
     *             .location(&#34;us-central1-a&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetPrivateCloudResult> getPrivateCloud(GetPrivateCloudArgs args) {
        return getPrivateCloud(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetPrivateCloudArgs;
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
     *         final var myPc = VmwareengineFunctions.getPrivateCloud(GetPrivateCloudArgs.builder()
     *             .name(&#34;my-pc&#34;)
     *             .location(&#34;us-central1-a&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetPrivateCloudResult> getPrivateCloudPlain(GetPrivateCloudPlainArgs args) {
        return getPrivateCloudPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetPrivateCloudArgs;
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
     *         final var myPc = VmwareengineFunctions.getPrivateCloud(GetPrivateCloudArgs.builder()
     *             .name(&#34;my-pc&#34;)
     *             .location(&#34;us-central1-a&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetPrivateCloudResult> getPrivateCloud(GetPrivateCloudArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:vmwareengine/getPrivateCloud:getPrivateCloud", TypeShape.of(GetPrivateCloudResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.vmwareengine.VmwareengineFunctions;
     * import com.pulumi.gcp.vmwareengine.inputs.GetPrivateCloudArgs;
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
     *         final var myPc = VmwareengineFunctions.getPrivateCloud(GetPrivateCloudArgs.builder()
     *             .name(&#34;my-pc&#34;)
     *             .location(&#34;us-central1-a&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetPrivateCloudResult> getPrivateCloudPlain(GetPrivateCloudPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:vmwareengine/getPrivateCloud:getPrivateCloud", TypeShape.of(GetPrivateCloudResult.class), args, Utilities.withVersion(options));
    }
}
