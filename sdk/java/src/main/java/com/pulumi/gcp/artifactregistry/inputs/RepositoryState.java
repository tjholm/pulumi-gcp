// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.artifactregistry.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryCleanupPolicyArgs;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryDockerConfigArgs;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryMavenConfigArgs;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryRemoteRepositoryConfigArgs;
import com.pulumi.gcp.artifactregistry.inputs.RepositoryVirtualRepositoryConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryState Empty = new RepositoryState();

    /**
     * (Optional, Beta)
     * Cleanup policies for this repository. Cleanup policies indicate when
     * certain package versions can be automatically deleted.
     * Map keys are policy IDs supplied by users during policy creation. They must
     * unique within a repository and be under 128 characters in length.
     * Structure is documented below.
     * 
     */
    @Import(name="cleanupPolicies")
    private @Nullable Output<List<RepositoryCleanupPolicyArgs>> cleanupPolicies;

    /**
     * @return (Optional, Beta)
     * Cleanup policies for this repository. Cleanup policies indicate when
     * certain package versions can be automatically deleted.
     * Map keys are policy IDs supplied by users during policy creation. They must
     * unique within a repository and be under 128 characters in length.
     * Structure is documented below.
     * 
     */
    public Optional<Output<List<RepositoryCleanupPolicyArgs>>> cleanupPolicies() {
        return Optional.ofNullable(this.cleanupPolicies);
    }

    /**
     * (Optional, Beta)
     * If true, the cleanup pipeline is prevented from deleting versions in this
     * repository.
     * 
     */
    @Import(name="cleanupPolicyDryRun")
    private @Nullable Output<Boolean> cleanupPolicyDryRun;

    /**
     * @return (Optional, Beta)
     * If true, the cleanup pipeline is prevented from deleting versions in this
     * repository.
     * 
     */
    public Optional<Output<Boolean>> cleanupPolicyDryRun() {
        return Optional.ofNullable(this.cleanupPolicyDryRun);
    }

    /**
     * The time when the repository was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return The time when the repository was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * The user-provided description of the repository.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The user-provided description of the repository.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Docker repository config contains repository level configuration for the repositories of docker type.
     * Structure is documented below.
     * 
     */
    @Import(name="dockerConfig")
    private @Nullable Output<RepositoryDockerConfigArgs> dockerConfig;

    /**
     * @return Docker repository config contains repository level configuration for the repositories of docker type.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RepositoryDockerConfigArgs>> dockerConfig() {
        return Optional.ofNullable(this.dockerConfig);
    }

    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Import(name="effectiveLabels")
    private @Nullable Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Optional<Output<Map<String,String>>> effectiveLabels() {
        return Optional.ofNullable(this.effectiveLabels);
    }

    /**
     * The format of packages that are stored in the repository. Supported formats
     * can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
     * You can only create alpha formats if you are a member of the
     * [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * 
     * ***
     * 
     */
    @Import(name="format")
    private @Nullable Output<String> format;

    /**
     * @return The format of packages that are stored in the repository. Supported formats
     * can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
     * You can only create alpha formats if you are a member of the
     * [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
     * 
     * ***
     * 
     */
    public Optional<Output<String>> format() {
        return Optional.ofNullable(this.format);
    }

    /**
     * The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     * 
     */
    @Import(name="kmsKeyName")
    private @Nullable Output<String> kmsKeyName;

    /**
     * @return The Cloud KMS resource name of the customer managed encryption key that’s
     * used to encrypt the contents of the Repository. Has the form:
     * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
     * This value may not be changed after the Repository has been created.
     * 
     */
    public Optional<Output<String>> kmsKeyName() {
        return Optional.ofNullable(this.kmsKeyName);
    }

    /**
     * Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Labels with user-defined metadata.
     * This field may contain up to 64 entries. Label keys and values may be no
     * longer than 63 characters. Label keys must begin with a lowercase letter
     * and may only contain lowercase letters, numeric characters, underscores,
     * and dashes.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The name of the location this repository is located in.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The name of the location this repository is located in.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * MavenRepositoryConfig is maven related repository details.
     * Provides additional configuration details for repositories of the maven
     * format type.
     * Structure is documented below.
     * 
     */
    @Import(name="mavenConfig")
    private @Nullable Output<RepositoryMavenConfigArgs> mavenConfig;

    /**
     * @return MavenRepositoryConfig is maven related repository details.
     * Provides additional configuration details for repositories of the maven
     * format type.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RepositoryMavenConfigArgs>> mavenConfig() {
        return Optional.ofNullable(this.mavenConfig);
    }

    /**
     * The mode configures the repository to serve artifacts from different sources.
     * Default value is `STANDARD_REPOSITORY`.
     * Possible values are: `STANDARD_REPOSITORY`, `VIRTUAL_REPOSITORY`, `REMOTE_REPOSITORY`.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<String> mode;

    /**
     * @return The mode configures the repository to serve artifacts from different sources.
     * Default value is `STANDARD_REPOSITORY`.
     * Possible values are: `STANDARD_REPOSITORY`, `VIRTUAL_REPOSITORY`, `REMOTE_REPOSITORY`.
     * 
     */
    public Optional<Output<String>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The name of the repository, for example:
     * &#34;repo1&#34;
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the repository, for example:
     * &#34;repo1&#34;
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
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Import(name="pulumiLabels")
    private @Nullable Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Optional<Output<Map<String,String>>> pulumiLabels() {
        return Optional.ofNullable(this.pulumiLabels);
    }

    /**
     * Configuration specific for a Remote Repository.
     * Structure is documented below.
     * 
     */
    @Import(name="remoteRepositoryConfig")
    private @Nullable Output<RepositoryRemoteRepositoryConfigArgs> remoteRepositoryConfig;

    /**
     * @return Configuration specific for a Remote Repository.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RepositoryRemoteRepositoryConfigArgs>> remoteRepositoryConfig() {
        return Optional.ofNullable(this.remoteRepositoryConfig);
    }

    /**
     * The last part of the repository name, for example:
     * &#34;repo1&#34;
     * 
     */
    @Import(name="repositoryId")
    private @Nullable Output<String> repositoryId;

    /**
     * @return The last part of the repository name, for example:
     * &#34;repo1&#34;
     * 
     */
    public Optional<Output<String>> repositoryId() {
        return Optional.ofNullable(this.repositoryId);
    }

    /**
     * The time when the repository was last updated.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return The time when the repository was last updated.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    /**
     * Configuration specific for a Virtual Repository.
     * Structure is documented below.
     * 
     */
    @Import(name="virtualRepositoryConfig")
    private @Nullable Output<RepositoryVirtualRepositoryConfigArgs> virtualRepositoryConfig;

    /**
     * @return Configuration specific for a Virtual Repository.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RepositoryVirtualRepositoryConfigArgs>> virtualRepositoryConfig() {
        return Optional.ofNullable(this.virtualRepositoryConfig);
    }

    private RepositoryState() {}

    private RepositoryState(RepositoryState $) {
        this.cleanupPolicies = $.cleanupPolicies;
        this.cleanupPolicyDryRun = $.cleanupPolicyDryRun;
        this.createTime = $.createTime;
        this.description = $.description;
        this.dockerConfig = $.dockerConfig;
        this.effectiveLabels = $.effectiveLabels;
        this.format = $.format;
        this.kmsKeyName = $.kmsKeyName;
        this.labels = $.labels;
        this.location = $.location;
        this.mavenConfig = $.mavenConfig;
        this.mode = $.mode;
        this.name = $.name;
        this.project = $.project;
        this.pulumiLabels = $.pulumiLabels;
        this.remoteRepositoryConfig = $.remoteRepositoryConfig;
        this.repositoryId = $.repositoryId;
        this.updateTime = $.updateTime;
        this.virtualRepositoryConfig = $.virtualRepositoryConfig;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryState $;

        public Builder() {
            $ = new RepositoryState();
        }

        public Builder(RepositoryState defaults) {
            $ = new RepositoryState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cleanupPolicies (Optional, Beta)
         * Cleanup policies for this repository. Cleanup policies indicate when
         * certain package versions can be automatically deleted.
         * Map keys are policy IDs supplied by users during policy creation. They must
         * unique within a repository and be under 128 characters in length.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cleanupPolicies(@Nullable Output<List<RepositoryCleanupPolicyArgs>> cleanupPolicies) {
            $.cleanupPolicies = cleanupPolicies;
            return this;
        }

        /**
         * @param cleanupPolicies (Optional, Beta)
         * Cleanup policies for this repository. Cleanup policies indicate when
         * certain package versions can be automatically deleted.
         * Map keys are policy IDs supplied by users during policy creation. They must
         * unique within a repository and be under 128 characters in length.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cleanupPolicies(List<RepositoryCleanupPolicyArgs> cleanupPolicies) {
            return cleanupPolicies(Output.of(cleanupPolicies));
        }

        /**
         * @param cleanupPolicies (Optional, Beta)
         * Cleanup policies for this repository. Cleanup policies indicate when
         * certain package versions can be automatically deleted.
         * Map keys are policy IDs supplied by users during policy creation. They must
         * unique within a repository and be under 128 characters in length.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder cleanupPolicies(RepositoryCleanupPolicyArgs... cleanupPolicies) {
            return cleanupPolicies(List.of(cleanupPolicies));
        }

        /**
         * @param cleanupPolicyDryRun (Optional, Beta)
         * If true, the cleanup pipeline is prevented from deleting versions in this
         * repository.
         * 
         * @return builder
         * 
         */
        public Builder cleanupPolicyDryRun(@Nullable Output<Boolean> cleanupPolicyDryRun) {
            $.cleanupPolicyDryRun = cleanupPolicyDryRun;
            return this;
        }

        /**
         * @param cleanupPolicyDryRun (Optional, Beta)
         * If true, the cleanup pipeline is prevented from deleting versions in this
         * repository.
         * 
         * @return builder
         * 
         */
        public Builder cleanupPolicyDryRun(Boolean cleanupPolicyDryRun) {
            return cleanupPolicyDryRun(Output.of(cleanupPolicyDryRun));
        }

        /**
         * @param createTime The time when the repository was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime The time when the repository was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param description The user-provided description of the repository.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The user-provided description of the repository.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param dockerConfig Docker repository config contains repository level configuration for the repositories of docker type.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder dockerConfig(@Nullable Output<RepositoryDockerConfigArgs> dockerConfig) {
            $.dockerConfig = dockerConfig;
            return this;
        }

        /**
         * @param dockerConfig Docker repository config contains repository level configuration for the repositories of docker type.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder dockerConfig(RepositoryDockerConfigArgs dockerConfig) {
            return dockerConfig(Output.of(dockerConfig));
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(@Nullable Output<Map<String,String>> effectiveLabels) {
            $.effectiveLabels = effectiveLabels;
            return this;
        }

        /**
         * @param effectiveLabels All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
         * 
         * @return builder
         * 
         */
        public Builder effectiveLabels(Map<String,String> effectiveLabels) {
            return effectiveLabels(Output.of(effectiveLabels));
        }

        /**
         * @param format The format of packages that are stored in the repository. Supported formats
         * can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
         * You can only create alpha formats if you are a member of the
         * [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder format(@Nullable Output<String> format) {
            $.format = format;
            return this;
        }

        /**
         * @param format The format of packages that are stored in the repository. Supported formats
         * can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
         * You can only create alpha formats if you are a member of the
         * [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder format(String format) {
            return format(Output.of(format));
        }

        /**
         * @param kmsKeyName The Cloud KMS resource name of the customer managed encryption key that’s
         * used to encrypt the contents of the Repository. Has the form:
         * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
         * This value may not be changed after the Repository has been created.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyName(@Nullable Output<String> kmsKeyName) {
            $.kmsKeyName = kmsKeyName;
            return this;
        }

        /**
         * @param kmsKeyName The Cloud KMS resource name of the customer managed encryption key that’s
         * used to encrypt the contents of the Repository. Has the form:
         * `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
         * This value may not be changed after the Repository has been created.
         * 
         * @return builder
         * 
         */
        public Builder kmsKeyName(String kmsKeyName) {
            return kmsKeyName(Output.of(kmsKeyName));
        }

        /**
         * @param labels Labels with user-defined metadata.
         * This field may contain up to 64 entries. Label keys and values may be no
         * longer than 63 characters. Label keys must begin with a lowercase letter
         * and may only contain lowercase letters, numeric characters, underscores,
         * and dashes.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Labels with user-defined metadata.
         * This field may contain up to 64 entries. Label keys and values may be no
         * longer than 63 characters. Label keys must begin with a lowercase letter
         * and may only contain lowercase letters, numeric characters, underscores,
         * and dashes.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The name of the location this repository is located in.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The name of the location this repository is located in.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param mavenConfig MavenRepositoryConfig is maven related repository details.
         * Provides additional configuration details for repositories of the maven
         * format type.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder mavenConfig(@Nullable Output<RepositoryMavenConfigArgs> mavenConfig) {
            $.mavenConfig = mavenConfig;
            return this;
        }

        /**
         * @param mavenConfig MavenRepositoryConfig is maven related repository details.
         * Provides additional configuration details for repositories of the maven
         * format type.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder mavenConfig(RepositoryMavenConfigArgs mavenConfig) {
            return mavenConfig(Output.of(mavenConfig));
        }

        /**
         * @param mode The mode configures the repository to serve artifacts from different sources.
         * Default value is `STANDARD_REPOSITORY`.
         * Possible values are: `STANDARD_REPOSITORY`, `VIRTUAL_REPOSITORY`, `REMOTE_REPOSITORY`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<String> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The mode configures the repository to serve artifacts from different sources.
         * Default value is `STANDARD_REPOSITORY`.
         * Possible values are: `STANDARD_REPOSITORY`, `VIRTUAL_REPOSITORY`, `REMOTE_REPOSITORY`.
         * 
         * @return builder
         * 
         */
        public Builder mode(String mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name The name of the repository, for example:
         * &#34;repo1&#34;
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the repository, for example:
         * &#34;repo1&#34;
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
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(@Nullable Output<Map<String,String>> pulumiLabels) {
            $.pulumiLabels = pulumiLabels;
            return this;
        }

        /**
         * @param pulumiLabels The combination of labels configured directly on the resource
         * and default labels configured on the provider.
         * 
         * @return builder
         * 
         */
        public Builder pulumiLabels(Map<String,String> pulumiLabels) {
            return pulumiLabels(Output.of(pulumiLabels));
        }

        /**
         * @param remoteRepositoryConfig Configuration specific for a Remote Repository.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder remoteRepositoryConfig(@Nullable Output<RepositoryRemoteRepositoryConfigArgs> remoteRepositoryConfig) {
            $.remoteRepositoryConfig = remoteRepositoryConfig;
            return this;
        }

        /**
         * @param remoteRepositoryConfig Configuration specific for a Remote Repository.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder remoteRepositoryConfig(RepositoryRemoteRepositoryConfigArgs remoteRepositoryConfig) {
            return remoteRepositoryConfig(Output.of(remoteRepositoryConfig));
        }

        /**
         * @param repositoryId The last part of the repository name, for example:
         * &#34;repo1&#34;
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(@Nullable Output<String> repositoryId) {
            $.repositoryId = repositoryId;
            return this;
        }

        /**
         * @param repositoryId The last part of the repository name, for example:
         * &#34;repo1&#34;
         * 
         * @return builder
         * 
         */
        public Builder repositoryId(String repositoryId) {
            return repositoryId(Output.of(repositoryId));
        }

        /**
         * @param updateTime The time when the repository was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime The time when the repository was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        /**
         * @param virtualRepositoryConfig Configuration specific for a Virtual Repository.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder virtualRepositoryConfig(@Nullable Output<RepositoryVirtualRepositoryConfigArgs> virtualRepositoryConfig) {
            $.virtualRepositoryConfig = virtualRepositoryConfig;
            return this;
        }

        /**
         * @param virtualRepositoryConfig Configuration specific for a Virtual Repository.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder virtualRepositoryConfig(RepositoryVirtualRepositoryConfigArgs virtualRepositoryConfig) {
            return virtualRepositoryConfig(Output.of(virtualRepositoryConfig));
        }

        public RepositoryState build() {
            return $;
        }
    }

}
