// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigquery.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.bigquery.inputs.AppProfileSingleClusterRoutingArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppProfileState extends com.pulumi.resources.ResourceArgs {

    public static final AppProfileState Empty = new AppProfileState();

    /**
     * The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
     * 
     */
    @Import(name="appProfileId")
    private @Nullable Output<String> appProfileId;

    /**
     * @return The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
     * 
     */
    public Optional<Output<String>> appProfileId() {
        return Optional.ofNullable(this.appProfileId);
    }

    /**
     * Long form description of the use case for this app profile.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Long form description of the use case for this app profile.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * If true, ignore safety checks when deleting/updating the app profile.
     * 
     */
    @Import(name="ignoreWarnings")
    private @Nullable Output<Boolean> ignoreWarnings;

    /**
     * @return If true, ignore safety checks when deleting/updating the app profile.
     * 
     */
    public Optional<Output<Boolean>> ignoreWarnings() {
        return Optional.ofNullable(this.ignoreWarnings);
    }

    /**
     * The name of the instance to create the app profile within.
     * 
     */
    @Import(name="instance")
    private @Nullable Output<String> instance;

    /**
     * @return The name of the instance to create the app profile within.
     * 
     */
    public Optional<Output<String>> instance() {
        return Optional.ofNullable(this.instance);
    }

    /**
     * The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
     * clusters are eligible.
     * 
     */
    @Import(name="multiClusterRoutingClusterIds")
    private @Nullable Output<List<String>> multiClusterRoutingClusterIds;

    /**
     * @return The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
     * clusters are eligible.
     * 
     */
    public Optional<Output<List<String>>> multiClusterRoutingClusterIds() {
        return Optional.ofNullable(this.multiClusterRoutingClusterIds);
    }

    /**
     * If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
     * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
     * consistency to improve availability.
     * 
     */
    @Import(name="multiClusterRoutingUseAny")
    private @Nullable Output<Boolean> multiClusterRoutingUseAny;

    /**
     * @return If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
     * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
     * consistency to improve availability.
     * 
     */
    public Optional<Output<Boolean>> multiClusterRoutingUseAny() {
        return Optional.ofNullable(this.multiClusterRoutingUseAny);
    }

    /**
     * The unique name of the requested app profile. Values are of the form
     * &#39;projects/&lt;project&gt;/instances/&lt;instance&gt;/appProfiles/&lt;appProfileId&gt;&#39;.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The unique name of the requested app profile. Values are of the form
     * &#39;projects/&lt;project&gt;/instances/&lt;instance&gt;/appProfiles/&lt;appProfileId&gt;&#39;.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Use a single-cluster routing policy.
     * Structure is documented below.
     * 
     */
    @Import(name="singleClusterRouting")
    private @Nullable Output<AppProfileSingleClusterRoutingArgs> singleClusterRouting;

    /**
     * @return Use a single-cluster routing policy.
     * Structure is documented below.
     * 
     */
    public Optional<Output<AppProfileSingleClusterRoutingArgs>> singleClusterRouting() {
        return Optional.ofNullable(this.singleClusterRouting);
    }

    private AppProfileState() {}

    private AppProfileState(AppProfileState $) {
        this.appProfileId = $.appProfileId;
        this.description = $.description;
        this.ignoreWarnings = $.ignoreWarnings;
        this.instance = $.instance;
        this.multiClusterRoutingClusterIds = $.multiClusterRoutingClusterIds;
        this.multiClusterRoutingUseAny = $.multiClusterRoutingUseAny;
        this.name = $.name;
        this.project = $.project;
        this.singleClusterRouting = $.singleClusterRouting;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppProfileState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppProfileState $;

        public Builder() {
            $ = new AppProfileState();
        }

        public Builder(AppProfileState defaults) {
            $ = new AppProfileState(Objects.requireNonNull(defaults));
        }

        /**
         * @param appProfileId The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
         * 
         * @return builder
         * 
         */
        public Builder appProfileId(@Nullable Output<String> appProfileId) {
            $.appProfileId = appProfileId;
            return this;
        }

        /**
         * @param appProfileId The unique name of the app profile in the form `[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
         * 
         * @return builder
         * 
         */
        public Builder appProfileId(String appProfileId) {
            return appProfileId(Output.of(appProfileId));
        }

        /**
         * @param description Long form description of the use case for this app profile.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Long form description of the use case for this app profile.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param ignoreWarnings If true, ignore safety checks when deleting/updating the app profile.
         * 
         * @return builder
         * 
         */
        public Builder ignoreWarnings(@Nullable Output<Boolean> ignoreWarnings) {
            $.ignoreWarnings = ignoreWarnings;
            return this;
        }

        /**
         * @param ignoreWarnings If true, ignore safety checks when deleting/updating the app profile.
         * 
         * @return builder
         * 
         */
        public Builder ignoreWarnings(Boolean ignoreWarnings) {
            return ignoreWarnings(Output.of(ignoreWarnings));
        }

        /**
         * @param instance The name of the instance to create the app profile within.
         * 
         * @return builder
         * 
         */
        public Builder instance(@Nullable Output<String> instance) {
            $.instance = instance;
            return this;
        }

        /**
         * @param instance The name of the instance to create the app profile within.
         * 
         * @return builder
         * 
         */
        public Builder instance(String instance) {
            return instance(Output.of(instance));
        }

        /**
         * @param multiClusterRoutingClusterIds The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
         * clusters are eligible.
         * 
         * @return builder
         * 
         */
        public Builder multiClusterRoutingClusterIds(@Nullable Output<List<String>> multiClusterRoutingClusterIds) {
            $.multiClusterRoutingClusterIds = multiClusterRoutingClusterIds;
            return this;
        }

        /**
         * @param multiClusterRoutingClusterIds The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
         * clusters are eligible.
         * 
         * @return builder
         * 
         */
        public Builder multiClusterRoutingClusterIds(List<String> multiClusterRoutingClusterIds) {
            return multiClusterRoutingClusterIds(Output.of(multiClusterRoutingClusterIds));
        }

        /**
         * @param multiClusterRoutingClusterIds The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all
         * clusters are eligible.
         * 
         * @return builder
         * 
         */
        public Builder multiClusterRoutingClusterIds(String... multiClusterRoutingClusterIds) {
            return multiClusterRoutingClusterIds(List.of(multiClusterRoutingClusterIds));
        }

        /**
         * @param multiClusterRoutingUseAny If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
         * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
         * consistency to improve availability.
         * 
         * @return builder
         * 
         */
        public Builder multiClusterRoutingUseAny(@Nullable Output<Boolean> multiClusterRoutingUseAny) {
            $.multiClusterRoutingUseAny = multiClusterRoutingUseAny;
            return this;
        }

        /**
         * @param multiClusterRoutingUseAny If true, read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available
         * in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes
         * consistency to improve availability.
         * 
         * @return builder
         * 
         */
        public Builder multiClusterRoutingUseAny(Boolean multiClusterRoutingUseAny) {
            return multiClusterRoutingUseAny(Output.of(multiClusterRoutingUseAny));
        }

        /**
         * @param name The unique name of the requested app profile. Values are of the form
         * &#39;projects/&lt;project&gt;/instances/&lt;instance&gt;/appProfiles/&lt;appProfileId&gt;&#39;.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The unique name of the requested app profile. Values are of the form
         * &#39;projects/&lt;project&gt;/instances/&lt;instance&gt;/appProfiles/&lt;appProfileId&gt;&#39;.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param singleClusterRouting Use a single-cluster routing policy.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder singleClusterRouting(@Nullable Output<AppProfileSingleClusterRoutingArgs> singleClusterRouting) {
            $.singleClusterRouting = singleClusterRouting;
            return this;
        }

        /**
         * @param singleClusterRouting Use a single-cluster routing policy.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder singleClusterRouting(AppProfileSingleClusterRoutingArgs singleClusterRouting) {
            return singleClusterRouting(Output.of(singleClusterRouting));
        }

        public AppProfileState build() {
            return $;
        }
    }

}
