// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.dataform.inputs.RepositoryWorkflowConfigInvocationConfigArgs;
import com.pulumi.gcp.dataform.inputs.RepositoryWorkflowConfigRecentScheduledExecutionRecordArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryWorkflowConfigState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryWorkflowConfigState Empty = new RepositoryWorkflowConfigState();

    /**
     * Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     * 
     */
    @Import(name="cronSchedule")
    private @Nullable Output<String> cronSchedule;

    /**
     * @return Optional. Optional schedule (in cron format) for automatic creation of compilation results.
     * 
     */
    public Optional<Output<String>> cronSchedule() {
        return Optional.ofNullable(this.cronSchedule);
    }

    /**
     * Optional. If left unset, a default InvocationConfig will be used.
     * Structure is documented below.
     * 
     */
    @Import(name="invocationConfig")
    private @Nullable Output<RepositoryWorkflowConfigInvocationConfigArgs> invocationConfig;

    /**
     * @return Optional. If left unset, a default InvocationConfig will be used.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RepositoryWorkflowConfigInvocationConfigArgs>> invocationConfig() {
        return Optional.ofNullable(this.invocationConfig);
    }

    /**
     * The workflow&#39;s name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The workflow&#39;s name.
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
     * Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
     * Structure is documented below.
     * 
     */
    @Import(name="recentScheduledExecutionRecords")
    private @Nullable Output<List<RepositoryWorkflowConfigRecentScheduledExecutionRecordArgs>> recentScheduledExecutionRecords;

    /**
     * @return Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<RepositoryWorkflowConfigRecentScheduledExecutionRecordArgs>>> recentScheduledExecutionRecords() {
        return Optional.ofNullable(this.recentScheduledExecutionRecords);
    }

    /**
     * A reference to the region
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return A reference to the region
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*&#47;locations/*&#47;repositories/*&#47;releaseConfigs/*.
     * 
     * ***
     * 
     */
    @Import(name="releaseConfig")
    private @Nullable Output<String> releaseConfig;

    /**
     * @return The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*&#47;locations/*&#47;repositories/*&#47;releaseConfigs/*.
     * 
     * ***
     * 
     */
    public Optional<Output<String>> releaseConfig() {
        return Optional.ofNullable(this.releaseConfig);
    }

    /**
     * A reference to the Dataform repository
     * 
     */
    @Import(name="repository")
    private @Nullable Output<String> repository;

    /**
     * @return A reference to the Dataform repository
     * 
     */
    public Optional<Output<String>> repository() {
        return Optional.ofNullable(this.repository);
    }

    /**
     * Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     * 
     */
    @Import(name="timeZone")
    private @Nullable Output<String> timeZone;

    /**
     * @return Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
     * 
     */
    public Optional<Output<String>> timeZone() {
        return Optional.ofNullable(this.timeZone);
    }

    private RepositoryWorkflowConfigState() {}

    private RepositoryWorkflowConfigState(RepositoryWorkflowConfigState $) {
        this.cronSchedule = $.cronSchedule;
        this.invocationConfig = $.invocationConfig;
        this.name = $.name;
        this.project = $.project;
        this.recentScheduledExecutionRecords = $.recentScheduledExecutionRecords;
        this.region = $.region;
        this.releaseConfig = $.releaseConfig;
        this.repository = $.repository;
        this.timeZone = $.timeZone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryWorkflowConfigState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryWorkflowConfigState $;

        public Builder() {
            $ = new RepositoryWorkflowConfigState();
        }

        public Builder(RepositoryWorkflowConfigState defaults) {
            $ = new RepositoryWorkflowConfigState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cronSchedule Optional. Optional schedule (in cron format) for automatic creation of compilation results.
         * 
         * @return builder
         * 
         */
        public Builder cronSchedule(@Nullable Output<String> cronSchedule) {
            $.cronSchedule = cronSchedule;
            return this;
        }

        /**
         * @param cronSchedule Optional. Optional schedule (in cron format) for automatic creation of compilation results.
         * 
         * @return builder
         * 
         */
        public Builder cronSchedule(String cronSchedule) {
            return cronSchedule(Output.of(cronSchedule));
        }

        /**
         * @param invocationConfig Optional. If left unset, a default InvocationConfig will be used.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder invocationConfig(@Nullable Output<RepositoryWorkflowConfigInvocationConfigArgs> invocationConfig) {
            $.invocationConfig = invocationConfig;
            return this;
        }

        /**
         * @param invocationConfig Optional. If left unset, a default InvocationConfig will be used.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder invocationConfig(RepositoryWorkflowConfigInvocationConfigArgs invocationConfig) {
            return invocationConfig(Output.of(invocationConfig));
        }

        /**
         * @param name The workflow&#39;s name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The workflow&#39;s name.
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
         * @param recentScheduledExecutionRecords Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder recentScheduledExecutionRecords(@Nullable Output<List<RepositoryWorkflowConfigRecentScheduledExecutionRecordArgs>> recentScheduledExecutionRecords) {
            $.recentScheduledExecutionRecords = recentScheduledExecutionRecords;
            return this;
        }

        /**
         * @param recentScheduledExecutionRecords Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder recentScheduledExecutionRecords(List<RepositoryWorkflowConfigRecentScheduledExecutionRecordArgs> recentScheduledExecutionRecords) {
            return recentScheduledExecutionRecords(Output.of(recentScheduledExecutionRecords));
        }

        /**
         * @param recentScheduledExecutionRecords Records of the 10 most recent scheduled execution attempts, ordered in in descending order of executionTime. Updated whenever automatic creation of a workflow invocation is triggered by cronSchedule.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder recentScheduledExecutionRecords(RepositoryWorkflowConfigRecentScheduledExecutionRecordArgs... recentScheduledExecutionRecords) {
            return recentScheduledExecutionRecords(List.of(recentScheduledExecutionRecords));
        }

        /**
         * @param region A reference to the region
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region A reference to the region
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param releaseConfig The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*&#47;locations/*&#47;repositories/*&#47;releaseConfigs/*.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder releaseConfig(@Nullable Output<String> releaseConfig) {
            $.releaseConfig = releaseConfig;
            return this;
        }

        /**
         * @param releaseConfig The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/*&#47;locations/*&#47;repositories/*&#47;releaseConfigs/*.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder releaseConfig(String releaseConfig) {
            return releaseConfig(Output.of(releaseConfig));
        }

        /**
         * @param repository A reference to the Dataform repository
         * 
         * @return builder
         * 
         */
        public Builder repository(@Nullable Output<String> repository) {
            $.repository = repository;
            return this;
        }

        /**
         * @param repository A reference to the Dataform repository
         * 
         * @return builder
         * 
         */
        public Builder repository(String repository) {
            return repository(Output.of(repository));
        }

        /**
         * @param timeZone Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(@Nullable Output<String> timeZone) {
            $.timeZone = timeZone;
            return this;
        }

        /**
         * @param timeZone Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
         * 
         * @return builder
         * 
         */
        public Builder timeZone(String timeZone) {
            return timeZone(Output.of(timeZone));
        }

        public RepositoryWorkflowConfigState build() {
            return $;
        }
    }

}
