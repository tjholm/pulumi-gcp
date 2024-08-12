// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.DatasetIamPolicyArgs;
import com.pulumi.gcp.bigquery.inputs.DatasetIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for BigQuery dataset. Each of these resources serves a different use case:
 * 
 * * `gcp.bigquery.DatasetIamPolicy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
 * * `gcp.bigquery.DatasetIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
 * * `gcp.bigquery.DatasetIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
 * 
 * These resources are intended to convert the permissions system for BigQuery datasets to the standard IAM interface. For advanced usages, including [creating authorized views](https://cloud.google.com/bigquery/docs/share-access-views), please use either `gcp.bigquery.DatasetAccess` or the `access` field on `gcp.bigquery.Dataset`.
 * 
 * &gt; **Note:** These resources **cannot** be used with `gcp.bigquery.DatasetAccess` resources or the `access` field on `gcp.bigquery.Dataset` or they will fight over what the policy should be.
 * 
 * &gt; **Note:** Using any of these resources will remove any authorized view permissions from the dataset. To assign and preserve authorized view permissions use the `gcp.bigquery.DatasetAccess` instead.
 * 
 * &gt; **Note:** Legacy BigQuery roles `OWNER` `WRITER` and `READER` **cannot** be used with any of these IAM resources. Instead use the full role form of: `roles/bigquery.dataOwner` `roles/bigquery.dataEditor` and `roles/bigquery.dataViewer`.
 * 
 * &gt; **Note:** `gcp.bigquery.DatasetIamPolicy` **cannot** be used in conjunction with `gcp.bigquery.DatasetIamBinding` and `gcp.bigquery.DatasetIamMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.bigquery.DatasetIamBinding` resources **can be** used in conjunction with `gcp.bigquery.DatasetIamMember` resources **only if** they do not grant privilege to the same role.
 * 
 * ## gcp.bigquery.DatasetIamPolicy
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
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.DatasetIamPolicy;
 * import com.pulumi.gcp.bigquery.DatasetIamPolicyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         final var owner = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/bigquery.dataOwner")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var datasetDataset = new Dataset("datasetDataset", DatasetArgs.builder()
 *             .datasetId("example_dataset")
 *             .build());
 * 
 *         var dataset = new DatasetIamPolicy("dataset", DatasetIamPolicyArgs.builder()
 *             .datasetId(datasetDataset.datasetId())
 *             .policyData(owner.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.bigquery.DatasetIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.DatasetIamBinding;
 * import com.pulumi.gcp.bigquery.DatasetIamBindingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var dataset = new Dataset("dataset", DatasetArgs.builder()
 *             .datasetId("example_dataset")
 *             .build());
 * 
 *         var reader = new DatasetIamBinding("reader", DatasetIamBindingArgs.builder()
 *             .datasetId(dataset.datasetId())
 *             .role("roles/bigquery.dataViewer")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.bigquery.DatasetIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.DatasetIamMember;
 * import com.pulumi.gcp.bigquery.DatasetIamMemberArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var dataset = new Dataset("dataset", DatasetArgs.builder()
 *             .datasetId("example_dataset")
 *             .build());
 * 
 *         var editor = new DatasetIamMember("editor", DatasetIamMemberArgs.builder()
 *             .datasetId(dataset.datasetId())
 *             .role("roles/bigquery.dataEditor")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.bigquery.DatasetIamBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.DatasetIamBinding;
 * import com.pulumi.gcp.bigquery.DatasetIamBindingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var dataset = new Dataset("dataset", DatasetArgs.builder()
 *             .datasetId("example_dataset")
 *             .build());
 * 
 *         var reader = new DatasetIamBinding("reader", DatasetIamBindingArgs.builder()
 *             .datasetId(dataset.datasetId())
 *             .role("roles/bigquery.dataViewer")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## gcp.bigquery.DatasetIamMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.DatasetIamMember;
 * import com.pulumi.gcp.bigquery.DatasetIamMemberArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var dataset = new Dataset("dataset", DatasetArgs.builder()
 *             .datasetId("example_dataset")
 *             .build());
 * 
 *         var editor = new DatasetIamMember("editor", DatasetIamMemberArgs.builder()
 *             .datasetId(dataset.datasetId())
 *             .role("roles/bigquery.dataEditor")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ### Importing IAM policies
 * 
 * IAM policy imports use the identifier of the BigQuery Dataset resource. For example:
 * 
 * * `projects/{{project_id}}/datasets/{{dataset_id}}`
 * 
 * An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
 * 
 * tf
 * 
 * import {
 * 
 *   id = projects/{{project_id}}/datasets/{{dataset_id}}
 * 
 *   to = google_bigquery_dataset_iam_policy.default
 * 
 * }
 * 
 * The `pulumi import` command can also be used:
 * 
 * ```sh
 * $ pulumi import gcp:bigquery/datasetIamPolicy:DatasetIamPolicy default projects/{{project_id}}/datasets/{{dataset_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:bigquery/datasetIamPolicy:DatasetIamPolicy")
public class DatasetIamPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The dataset ID.
     * 
     */
    @Export(name="datasetId", refs={String.class}, tree="[0]")
    private Output<String> datasetId;

    /**
     * @return The dataset ID.
     * 
     */
    public Output<String> datasetId() {
        return this.datasetId;
    }
    /**
     * (Computed) The etag of the dataset&#39;s IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the dataset&#39;s IAM policy.
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
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatasetIamPolicy(java.lang.String name) {
        this(name, DatasetIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatasetIamPolicy(java.lang.String name, DatasetIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatasetIamPolicy(java.lang.String name, DatasetIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DatasetIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable DatasetIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/datasetIamPolicy:DatasetIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static DatasetIamPolicyArgs makeArgs(DatasetIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DatasetIamPolicyArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static DatasetIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable DatasetIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatasetIamPolicy(name, id, state, options);
    }
}
