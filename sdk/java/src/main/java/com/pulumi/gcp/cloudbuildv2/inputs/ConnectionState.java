// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuildv2.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionGithubConfigArgs;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionGithubEnterpriseConfigArgs;
import com.pulumi.gcp.cloudbuildv2.inputs.ConnectionInstallationStateArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionState extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionState Empty = new ConnectionState();

    /**
     * Allows clients to store small amounts of arbitrary data.
     * 
     */
    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    /**
     * @return Allows clients to store small amounts of arbitrary data.
     * 
     */
    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    /**
     * Output only. Server assigned timestamp for when the connection was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Output only. Server assigned timestamp for when the connection was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    @Import(name="etag")
    private @Nullable Output<String> etag;

    /**
     * @return This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     * 
     */
    public Optional<Output<String>> etag() {
        return Optional.ofNullable(this.etag);
    }

    /**
     * Configuration for connections to github.com.
     * 
     */
    @Import(name="githubConfig")
    private @Nullable Output<ConnectionGithubConfigArgs> githubConfig;

    /**
     * @return Configuration for connections to github.com.
     * 
     */
    public Optional<Output<ConnectionGithubConfigArgs>> githubConfig() {
        return Optional.ofNullable(this.githubConfig);
    }

    /**
     * Configuration for connections to an instance of GitHub Enterprise.
     * 
     */
    @Import(name="githubEnterpriseConfig")
    private @Nullable Output<ConnectionGithubEnterpriseConfigArgs> githubEnterpriseConfig;

    /**
     * @return Configuration for connections to an instance of GitHub Enterprise.
     * 
     */
    public Optional<Output<ConnectionGithubEnterpriseConfigArgs>> githubEnterpriseConfig() {
        return Optional.ofNullable(this.githubEnterpriseConfig);
    }

    /**
     * Output only. Installation state of the Connection.
     * 
     */
    @Import(name="installationStates")
    private @Nullable Output<List<ConnectionInstallationStateArgs>> installationStates;

    /**
     * @return Output only. Installation state of the Connection.
     * 
     */
    public Optional<Output<List<ConnectionInstallationStateArgs>>> installationStates() {
        return Optional.ofNullable(this.installationStates);
    }

    /**
     * The location for the resource
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The project for the resource
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * Output only. Set to true when the connection is being set up or updated in the background.
     * 
     */
    @Import(name="reconciling")
    private @Nullable Output<Boolean> reconciling;

    /**
     * @return Output only. Set to true when the connection is being set up or updated in the background.
     * 
     */
    public Optional<Output<Boolean>> reconciling() {
        return Optional.ofNullable(this.reconciling);
    }

    /**
     * Output only. Server assigned timestamp for when the connection was updated.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Output only. Server assigned timestamp for when the connection was updated.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private ConnectionState() {}

    private ConnectionState(ConnectionState $) {
        this.annotations = $.annotations;
        this.createTime = $.createTime;
        this.disabled = $.disabled;
        this.etag = $.etag;
        this.githubConfig = $.githubConfig;
        this.githubEnterpriseConfig = $.githubEnterpriseConfig;
        this.installationStates = $.installationStates;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.reconciling = $.reconciling;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionState $;

        public Builder() {
            $ = new ConnectionState();
        }

        public Builder(ConnectionState defaults) {
            $ = new ConnectionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param annotations Allows clients to store small amounts of arbitrary data.
         * 
         * @return builder
         * 
         */
        public Builder annotations(@Nullable Output<Map<String,String>> annotations) {
            $.annotations = annotations;
            return this;
        }

        /**
         * @param annotations Allows clients to store small amounts of arbitrary data.
         * 
         * @return builder
         * 
         */
        public Builder annotations(Map<String,String> annotations) {
            return annotations(Output.of(annotations));
        }

        /**
         * @param createTime Output only. Server assigned timestamp for when the connection was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Output only. Server assigned timestamp for when the connection was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param disabled If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param etag This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
         * 
         * @return builder
         * 
         */
        public Builder etag(@Nullable Output<String> etag) {
            $.etag = etag;
            return this;
        }

        /**
         * @param etag This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
         * 
         * @return builder
         * 
         */
        public Builder etag(String etag) {
            return etag(Output.of(etag));
        }

        /**
         * @param githubConfig Configuration for connections to github.com.
         * 
         * @return builder
         * 
         */
        public Builder githubConfig(@Nullable Output<ConnectionGithubConfigArgs> githubConfig) {
            $.githubConfig = githubConfig;
            return this;
        }

        /**
         * @param githubConfig Configuration for connections to github.com.
         * 
         * @return builder
         * 
         */
        public Builder githubConfig(ConnectionGithubConfigArgs githubConfig) {
            return githubConfig(Output.of(githubConfig));
        }

        /**
         * @param githubEnterpriseConfig Configuration for connections to an instance of GitHub Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder githubEnterpriseConfig(@Nullable Output<ConnectionGithubEnterpriseConfigArgs> githubEnterpriseConfig) {
            $.githubEnterpriseConfig = githubEnterpriseConfig;
            return this;
        }

        /**
         * @param githubEnterpriseConfig Configuration for connections to an instance of GitHub Enterprise.
         * 
         * @return builder
         * 
         */
        public Builder githubEnterpriseConfig(ConnectionGithubEnterpriseConfigArgs githubEnterpriseConfig) {
            return githubEnterpriseConfig(Output.of(githubEnterpriseConfig));
        }

        /**
         * @param installationStates Output only. Installation state of the Connection.
         * 
         * @return builder
         * 
         */
        public Builder installationStates(@Nullable Output<List<ConnectionInstallationStateArgs>> installationStates) {
            $.installationStates = installationStates;
            return this;
        }

        /**
         * @param installationStates Output only. Installation state of the Connection.
         * 
         * @return builder
         * 
         */
        public Builder installationStates(List<ConnectionInstallationStateArgs> installationStates) {
            return installationStates(Output.of(installationStates));
        }

        /**
         * @param installationStates Output only. Installation state of the Connection.
         * 
         * @return builder
         * 
         */
        public Builder installationStates(ConnectionInstallationStateArgs... installationStates) {
            return installationStates(List.of(installationStates));
        }

        /**
         * @param location The location for the resource
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location for the resource
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param reconciling Output only. Set to true when the connection is being set up or updated in the background.
         * 
         * @return builder
         * 
         */
        public Builder reconciling(@Nullable Output<Boolean> reconciling) {
            $.reconciling = reconciling;
            return this;
        }

        /**
         * @param reconciling Output only. Set to true when the connection is being set up or updated in the background.
         * 
         * @return builder
         * 
         */
        public Builder reconciling(Boolean reconciling) {
            return reconciling(Output.of(reconciling));
        }

        /**
         * @param updateTime Output only. Server assigned timestamp for when the connection was updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Output only. Server assigned timestamp for when the connection was updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public ConnectionState build() {
            return $;
        }
    }

}
