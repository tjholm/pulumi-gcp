// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigquery.DatasetAccessArgs;
import com.pulumi.gcp.bigquery.inputs.DatasetAccessState;
import com.pulumi.gcp.bigquery.outputs.DatasetAccessAuthorizedDataset;
import com.pulumi.gcp.bigquery.outputs.DatasetAccessView;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Bigquery Dataset Access Basic User
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.serviceAccount.Account;
 * import com.pulumi.gcp.serviceAccount.AccountArgs;
 * import com.pulumi.gcp.bigquery.DatasetAccess;
 * import com.pulumi.gcp.bigquery.DatasetAccessArgs;
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
 *         var dataset = new Dataset(&#34;dataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_dataset&#34;)
 *             .build());
 * 
 *         var bqowner = new Account(&#34;bqowner&#34;, AccountArgs.builder()        
 *             .accountId(&#34;bqowner&#34;)
 *             .build());
 * 
 *         var access = new DatasetAccess(&#34;access&#34;, DatasetAccessArgs.builder()        
 *             .datasetId(dataset.datasetId())
 *             .role(&#34;OWNER&#34;)
 *             .userByEmail(bqowner.email())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Bigquery Dataset Access View
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.Table;
 * import com.pulumi.gcp.bigquery.TableArgs;
 * import com.pulumi.gcp.bigquery.inputs.TableViewArgs;
 * import com.pulumi.gcp.bigquery.DatasetAccess;
 * import com.pulumi.gcp.bigquery.DatasetAccessArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessViewArgs;
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
 *         var private_ = new Dataset(&#34;private&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_dataset&#34;)
 *             .build());
 * 
 *         var publicDataset = new Dataset(&#34;publicDataset&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;example_dataset2&#34;)
 *             .build());
 * 
 *         var publicTable = new Table(&#34;publicTable&#34;, TableArgs.builder()        
 *             .deletionProtection(false)
 *             .datasetId(publicDataset.datasetId())
 *             .tableId(&#34;example_table&#34;)
 *             .view(TableViewArgs.builder()
 *                 .query(&#34;SELECT state FROM [lookerdata:cdc.project_tycho_reports]&#34;)
 *                 .useLegacySql(false)
 *                 .build())
 *             .build());
 * 
 *         var access = new DatasetAccess(&#34;access&#34;, DatasetAccessArgs.builder()        
 *             .datasetId(private_.datasetId())
 *             .view(DatasetAccessViewArgs.builder()
 *                 .projectId(publicTable.project())
 *                 .datasetId(publicDataset.datasetId())
 *                 .tableId(publicTable.tableId())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Bigquery Dataset Access Authorized Dataset
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigquery.DatasetAccess;
 * import com.pulumi.gcp.bigquery.DatasetAccessArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessAuthorizedDatasetArgs;
 * import com.pulumi.gcp.bigquery.inputs.DatasetAccessAuthorizedDatasetDatasetArgs;
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
 *         var private_ = new Dataset(&#34;private&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;private&#34;)
 *             .build());
 * 
 *         var public_ = new Dataset(&#34;public&#34;, DatasetArgs.builder()        
 *             .datasetId(&#34;public&#34;)
 *             .build());
 * 
 *         var access = new DatasetAccess(&#34;access&#34;, DatasetAccessArgs.builder()        
 *             .datasetId(private_.datasetId())
 *             .authorizedDataset(DatasetAccessAuthorizedDatasetArgs.builder()
 *                 .dataset(DatasetAccessAuthorizedDatasetDatasetArgs.builder()
 *                     .projectId(public_.project())
 *                     .datasetId(public_.datasetId())
 *                     .build())
 *                 .targetTypes(&#34;VIEWS&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:bigquery/datasetAccess:DatasetAccess")
public class DatasetAccess extends com.pulumi.resources.CustomResource {
    /**
     * If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
     * stored in state as a different member type
     * 
     */
    @Export(name="apiUpdatedMember", type=Boolean.class, parameters={})
    private Output<Boolean> apiUpdatedMember;

    /**
     * @return If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
     * stored in state as a different member type
     * 
     */
    public Output<Boolean> apiUpdatedMember() {
        return this.apiUpdatedMember;
    }
    /**
     * The dataset this entry applies to
     * Structure is documented below.
     * 
     */
    @Export(name="authorizedDataset", type=DatasetAccessAuthorizedDataset.class, parameters={})
    private Output</* @Nullable */ DatasetAccessAuthorizedDataset> authorizedDataset;

    /**
     * @return The dataset this entry applies to
     * Structure is documented below.
     * 
     */
    public Output<Optional<DatasetAccessAuthorizedDataset>> authorizedDataset() {
        return Codegen.optional(this.authorizedDataset);
    }
    /**
     * The ID of the dataset containing this table.
     * 
     */
    @Export(name="datasetId", type=String.class, parameters={})
    private Output<String> datasetId;

    /**
     * @return The ID of the dataset containing this table.
     * 
     */
    public Output<String> datasetId() {
        return this.datasetId;
    }
    /**
     * A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     * 
     */
    @Export(name="domain", type=String.class, parameters={})
    private Output</* @Nullable */ String> domain;

    /**
     * @return A domain to grant access to. Any users signed in with the
     * domain specified will be granted the specified access
     * 
     */
    public Output<Optional<String>> domain() {
        return Codegen.optional(this.domain);
    }
    /**
     * An email address of a Google Group to grant access to.
     * 
     */
    @Export(name="groupByEmail", type=String.class, parameters={})
    private Output</* @Nullable */ String> groupByEmail;

    /**
     * @return An email address of a Google Group to grant access to.
     * 
     */
    public Output<Optional<String>> groupByEmail() {
        return Codegen.optional(this.groupByEmail);
    }
    /**
     * Some other type of member that appears in the IAM Policy but isn&#39;t a user,
     * group, domain, or special group. For example: `allUsers`
     * 
     */
    @Export(name="iamMember", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamMember;

    /**
     * @return Some other type of member that appears in the IAM Policy but isn&#39;t a user,
     * group, domain, or special group. For example: `allUsers`
     * 
     */
    public Output<Optional<String>> iamMember() {
        return Codegen.optional(this.iamMember);
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
     * Describes the rights granted to the user specified by the other
     * member of the access object. Basic, predefined, and custom roles are
     * supported. Predefined roles that have equivalent basic roles are
     * swapped by the API to their basic counterparts, and will show a diff
     * post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     * 
     */
    @Export(name="role", type=String.class, parameters={})
    private Output</* @Nullable */ String> role;

    /**
     * @return Describes the rights granted to the user specified by the other
     * member of the access object. Basic, predefined, and custom roles are
     * supported. Predefined roles that have equivalent basic roles are
     * swapped by the API to their basic counterparts, and will show a diff
     * post-create. See
     * [official docs](https://cloud.google.com/bigquery/docs/access-control).
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }
    /**
     * A special group to grant access to. Possible values include:
     * 
     */
    @Export(name="specialGroup", type=String.class, parameters={})
    private Output</* @Nullable */ String> specialGroup;

    /**
     * @return A special group to grant access to. Possible values include:
     * 
     */
    public Output<Optional<String>> specialGroup() {
        return Codegen.optional(this.specialGroup);
    }
    /**
     * An email address of a user to grant access to. For example:
     * fred@example.com
     * 
     */
    @Export(name="userByEmail", type=String.class, parameters={})
    private Output</* @Nullable */ String> userByEmail;

    /**
     * @return An email address of a user to grant access to. For example:
     * fred@example.com
     * 
     */
    public Output<Optional<String>> userByEmail() {
        return Codegen.optional(this.userByEmail);
    }
    /**
     * A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.
     * Structure is documented below.
     * 
     */
    @Export(name="view", type=DatasetAccessView.class, parameters={})
    private Output</* @Nullable */ DatasetAccessView> view;

    /**
     * @return A view from a different dataset to grant access to. Queries
     * executed against that view will have read access to tables in
     * this dataset. The role field is not required when this field is
     * set. If that view is updated by any user, access to the view
     * needs to be granted again via an update operation.
     * Structure is documented below.
     * 
     */
    public Output<Optional<DatasetAccessView>> view() {
        return Codegen.optional(this.view);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DatasetAccess(String name) {
        this(name, DatasetAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DatasetAccess(String name, DatasetAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DatasetAccess(String name, DatasetAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/datasetAccess:DatasetAccess", name, args == null ? DatasetAccessArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DatasetAccess(String name, Output<String> id, @Nullable DatasetAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigquery/datasetAccess:DatasetAccess", name, state, makeResourceOptions(options, id));
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
    public static DatasetAccess get(String name, Output<String> id, @Nullable DatasetAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DatasetAccess(name, id, state, options);
    }
}
