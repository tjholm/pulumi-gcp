// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
import com.pulumi.gcp.alloydb.inputs.GetLocationsPlainArgs;
import com.pulumi.gcp.alloydb.inputs.GetSupportedDatabaseFlagsArgs;
import com.pulumi.gcp.alloydb.inputs.GetSupportedDatabaseFlagsPlainArgs;
import com.pulumi.gcp.alloydb.outputs.GetLocationsResult;
import com.pulumi.gcp.alloydb.outputs.GetSupportedDatabaseFlagsResult;
import java.util.concurrent.CompletableFuture;

public final class AlloydbFunctions {
    /**
     * Use this data source to get information about the available locations. For more details refer the [API docs](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
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
     *         final var qa = AlloydbFunctions.getLocations();
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLocationsResult> getLocations() {
        return getLocations(GetLocationsArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information about the available locations. For more details refer the [API docs](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
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
     *         final var qa = AlloydbFunctions.getLocations();
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLocationsResult> getLocationsPlain() {
        return getLocationsPlain(GetLocationsPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information about the available locations. For more details refer the [API docs](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
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
     *         final var qa = AlloydbFunctions.getLocations();
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLocationsResult> getLocations(GetLocationsArgs args) {
        return getLocations(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information about the available locations. For more details refer the [API docs](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
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
     *         final var qa = AlloydbFunctions.getLocations();
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLocationsResult> getLocationsPlain(GetLocationsPlainArgs args) {
        return getLocationsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information about the available locations. For more details refer the [API docs](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
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
     *         final var qa = AlloydbFunctions.getLocations();
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetLocationsResult> getLocations(GetLocationsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:alloydb/getLocations:getLocations", TypeShape.of(GetLocationsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information about the available locations. For more details refer the [API docs](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations).
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetLocationsArgs;
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
     *         final var qa = AlloydbFunctions.getLocations();
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetLocationsResult> getLocationsPlain(GetLocationsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:alloydb/getLocations:getLocations", TypeShape.of(GetLocationsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information about the supported alloydb database flags in a location.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetSupportedDatabaseFlagsArgs;
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
     *         final var qa = AlloydbFunctions.getSupportedDatabaseFlags(GetSupportedDatabaseFlagsArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetSupportedDatabaseFlagsResult> getSupportedDatabaseFlags(GetSupportedDatabaseFlagsArgs args) {
        return getSupportedDatabaseFlags(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information about the supported alloydb database flags in a location.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetSupportedDatabaseFlagsArgs;
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
     *         final var qa = AlloydbFunctions.getSupportedDatabaseFlags(GetSupportedDatabaseFlagsArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetSupportedDatabaseFlagsResult> getSupportedDatabaseFlagsPlain(GetSupportedDatabaseFlagsPlainArgs args) {
        return getSupportedDatabaseFlagsPlain(args, InvokeOptions.Empty);
    }
    /**
     * Use this data source to get information about the supported alloydb database flags in a location.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetSupportedDatabaseFlagsArgs;
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
     *         final var qa = AlloydbFunctions.getSupportedDatabaseFlags(GetSupportedDatabaseFlagsArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetSupportedDatabaseFlagsResult> getSupportedDatabaseFlags(GetSupportedDatabaseFlagsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:alloydb/getSupportedDatabaseFlags:getSupportedDatabaseFlags", TypeShape.of(GetSupportedDatabaseFlagsResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Use this data source to get information about the supported alloydb database flags in a location.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.alloydb.AlloydbFunctions;
     * import com.pulumi.gcp.alloydb.inputs.GetSupportedDatabaseFlagsArgs;
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
     *         final var qa = AlloydbFunctions.getSupportedDatabaseFlags(GetSupportedDatabaseFlagsArgs.builder()
     *             .location(&#34;us-central1&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetSupportedDatabaseFlagsResult> getSupportedDatabaseFlagsPlain(GetSupportedDatabaseFlagsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:alloydb/getSupportedDatabaseFlags:getSupportedDatabaseFlags", TypeShape.of(GetSupportedDatabaseFlagsResult.class), args, Utilities.withVersion(options));
    }
}
