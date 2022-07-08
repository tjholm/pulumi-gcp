// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.artifactregistry;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.artifactregistry.RepositoryArgs;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryState;
import com.pulumi.gcp.artifactregistry.outputs.RepositoryMavenConfig;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A repository for storing artifacts
 * 
 * To get more information about Repository, see:
 * 
 * * [API documentation](https://cloud.google.com/artifact-registry/docs/reference/rest/v1beta2/projects.locations.repositories)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/artifact-registry/docs/overview)
 * 
 * ## Example Usage
 * ### Artifact Registry Repository Basic
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
 *         var my_repo = new Repository(&#34;my-repo&#34;, RepositoryArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .repositoryId(&#34;my-repository&#34;)
 *             .description(&#34;example docker repository&#34;)
 *             .format(&#34;DOCKER&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Artifact Registry Repository Cmek
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
 *         var my_repo = new Repository(&#34;my-repo&#34;, RepositoryArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .repositoryId(&#34;my-repository&#34;)
 *             .description(&#34;example docker repository with cmek&#34;)
 *             .format(&#34;DOCKER&#34;)
 *             .kmsKeyName(&#34;kms-key&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Artifact Registry Repository Iam
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
 *         var my_repo = new Repository(&#34;my-repo&#34;, RepositoryArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .repositoryId(&#34;my-repository&#34;)
 *             .description(&#34;example docker repository with iam&#34;)
 *             .format(&#34;DOCKER&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var test_account = new Account(&#34;test-account&#34;, AccountArgs.builder()        
 *             .accountId(&#34;my-account&#34;)
 *             .displayName(&#34;Test Service Account&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var test_iam = new RepositoryIamMember(&#34;test-iam&#34;, RepositoryIamMemberArgs.builder()        
 *             .location(my_repo.location())
 *             .repository(my_repo.name())
 *             .role(&#34;roles/artifactregistry.reader&#34;)
 *             .member(test_account.email().apply(email -&gt; String.format(&#34;serviceAccount:%s&#34;, email)))
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
 * Repository can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default {{project}}/{{location}}/{{repository_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default {{location}}/{{repository_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:artifactregistry/repository:Repository default {{repository_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:artifactregistry/repository:Repository")
public class Repository extends com.pulumi.resources.CustomResource {
    /**
     * The time when the repository was created.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return The time when the repository was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The user-provided description of the repository.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The user-provided description of the repository.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The format of packages that are stored in the repository. Supported formats
     * can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
     * You can only create alpha formats if you are a member of the
     * [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * 
     */
    @Export(name="format", type=String.class, parameters={})
    private Output<String> format;

    /**
     * @return The format of packages that are stored in the repository. Supported formats
     * can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
     * You can only create alpha formats if you are a member of the
     * [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * 
     */
    public Output<String> format() {
        return this.format;
    }
    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     * 
     */
    @Export(name="kmsKeyName", type=String.class, parameters={})
    private Output</* @Nullable */ String> kmsKeyName;

    /**
     * @return The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     * 
     */
    public Output<Optional<String>> kmsKeyName() {
        return Codegen.optional(this.kmsKeyName);
    }
    /**
     * Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name of the location this repository is located in.
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return The name of the location this repository is located in.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * MavenRepositoryConfig is maven related repository details.
     * Provides additional configuration details for repositories of the maven
     * format type.
     * Structure is documented below.
     * 
     */
    @Export(name="mavenConfig", type=RepositoryMavenConfig.class, parameters={})
    private Output</* @Nullable */ RepositoryMavenConfig> mavenConfig;

    /**
     * @return MavenRepositoryConfig is maven related repository details.
     * Provides additional configuration details for repositories of the maven
     * format type.
     * Structure is documented below.
     * 
     */
    public Output<Optional<RepositoryMavenConfig>> mavenConfig() {
        return Codegen.optional(this.mavenConfig);
    }
    /**
     * The name of the repository, for example: &#34;projects/p1/locations/us-central1/repositories/repo1&#34;
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The name of the repository, for example: &#34;projects/p1/locations/us-central1/repositories/repo1&#34;
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The last part of the repository name, for example:
     * &#34;repo1&#34;
     * 
     */
    @Export(name="repositoryId", type=String.class, parameters={})
    private Output<String> repositoryId;

    /**
     * @return The last part of the repository name, for example:
     * &#34;repo1&#34;
     * 
     */
    public Output<String> repositoryId() {
        return this.repositoryId;
    }
    /**
     * The time when the repository was last updated.
     * 
     */
    @Export(name="updateTime", type=String.class, parameters={})
    private Output<String> updateTime;

    /**
     * @return The time when the repository was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Repository(String name) {
        this(name, RepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Repository(String name, RepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Repository(String name, RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:artifactregistry/repository:Repository", name, args == null ? RepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Repository(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:artifactregistry/repository:Repository", name, state, makeResourceOptions(options, id));
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
    public static Repository get(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Repository(name, id, state, options);
    }
}
