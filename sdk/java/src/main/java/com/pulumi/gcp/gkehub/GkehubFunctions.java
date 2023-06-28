// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkehub.inputs.GetFeatureIamPolicyArgs;
import com.pulumi.gcp.gkehub.inputs.GetFeatureIamPolicyPlainArgs;
import com.pulumi.gcp.gkehub.inputs.GetMembershipIamPolicyArgs;
import com.pulumi.gcp.gkehub.inputs.GetMembershipIamPolicyPlainArgs;
import com.pulumi.gcp.gkehub.outputs.GetFeatureIamPolicyResult;
import com.pulumi.gcp.gkehub.outputs.GetMembershipIamPolicyResult;
import java.util.concurrent.CompletableFuture;

public final class GkehubFunctions {
    /**
     * Retrieves the current IAM policy data for feature
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetFeatureIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getFeatureIamPolicy(GetFeatureIamPolicyArgs.builder()
     *             .project(google_gke_hub_feature.feature().project())
     *             .location(google_gke_hub_feature.feature().location())
     *             .name(google_gke_hub_feature.feature().name())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFeatureIamPolicyResult> getFeatureIamPolicy(GetFeatureIamPolicyArgs args) {
        return getFeatureIamPolicy(args, InvokeOptions.Empty);
    }
    /**
     * Retrieves the current IAM policy data for feature
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetFeatureIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getFeatureIamPolicy(GetFeatureIamPolicyArgs.builder()
     *             .project(google_gke_hub_feature.feature().project())
     *             .location(google_gke_hub_feature.feature().location())
     *             .name(google_gke_hub_feature.feature().name())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFeatureIamPolicyResult> getFeatureIamPolicyPlain(GetFeatureIamPolicyPlainArgs args) {
        return getFeatureIamPolicyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Retrieves the current IAM policy data for feature
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetFeatureIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getFeatureIamPolicy(GetFeatureIamPolicyArgs.builder()
     *             .project(google_gke_hub_feature.feature().project())
     *             .location(google_gke_hub_feature.feature().location())
     *             .name(google_gke_hub_feature.feature().name())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFeatureIamPolicyResult> getFeatureIamPolicy(GetFeatureIamPolicyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:gkehub/getFeatureIamPolicy:getFeatureIamPolicy", TypeShape.of(GetFeatureIamPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Retrieves the current IAM policy data for feature
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetFeatureIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getFeatureIamPolicy(GetFeatureIamPolicyArgs.builder()
     *             .project(google_gke_hub_feature.feature().project())
     *             .location(google_gke_hub_feature.feature().location())
     *             .name(google_gke_hub_feature.feature().name())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFeatureIamPolicyResult> getFeatureIamPolicyPlain(GetFeatureIamPolicyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:gkehub/getFeatureIamPolicy:getFeatureIamPolicy", TypeShape.of(GetFeatureIamPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Retrieves the current IAM policy data for membership
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetMembershipIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getMembershipIamPolicy(GetMembershipIamPolicyArgs.builder()
     *             .project(google_gke_hub_membership.membership().project())
     *             .membershipId(google_gke_hub_membership.membership().membership_id())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetMembershipIamPolicyResult> getMembershipIamPolicy(GetMembershipIamPolicyArgs args) {
        return getMembershipIamPolicy(args, InvokeOptions.Empty);
    }
    /**
     * Retrieves the current IAM policy data for membership
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetMembershipIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getMembershipIamPolicy(GetMembershipIamPolicyArgs.builder()
     *             .project(google_gke_hub_membership.membership().project())
     *             .membershipId(google_gke_hub_membership.membership().membership_id())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetMembershipIamPolicyResult> getMembershipIamPolicyPlain(GetMembershipIamPolicyPlainArgs args) {
        return getMembershipIamPolicyPlain(args, InvokeOptions.Empty);
    }
    /**
     * Retrieves the current IAM policy data for membership
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetMembershipIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getMembershipIamPolicy(GetMembershipIamPolicyArgs.builder()
     *             .project(google_gke_hub_membership.membership().project())
     *             .membershipId(google_gke_hub_membership.membership().membership_id())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetMembershipIamPolicyResult> getMembershipIamPolicy(GetMembershipIamPolicyArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:gkehub/getMembershipIamPolicy:getMembershipIamPolicy", TypeShape.of(GetMembershipIamPolicyResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Retrieves the current IAM policy data for membership
     * 
     * ## example
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.gcp.gkehub.GkehubFunctions;
     * import com.pulumi.gcp.gkehub.inputs.GetMembershipIamPolicyArgs;
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
     *         final var policy = GkehubFunctions.getMembershipIamPolicy(GetMembershipIamPolicyArgs.builder()
     *             .project(google_gke_hub_membership.membership().project())
     *             .membershipId(google_gke_hub_membership.membership().membership_id())
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetMembershipIamPolicyResult> getMembershipIamPolicyPlain(GetMembershipIamPolicyPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:gkehub/getMembershipIamPolicy:getMembershipIamPolicy", TypeShape.of(GetMembershipIamPolicyResult.class), args, Utilities.withVersion(options));
    }
}
