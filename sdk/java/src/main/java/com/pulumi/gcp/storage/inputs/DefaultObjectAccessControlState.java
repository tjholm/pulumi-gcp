// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.storage.inputs.DefaultObjectAccessControlProjectTeamArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DefaultObjectAccessControlState extends com.pulumi.resources.ResourceArgs {

    public static final DefaultObjectAccessControlState Empty = new DefaultObjectAccessControlState();

    /**
     * The name of the bucket.
     * 
     */
    @Import(name="bucket")
    private @Nullable Output<String> bucket;

    /**
     * @return The name of the bucket.
     * 
     */
    public Optional<Output<String>> bucket() {
        return Optional.ofNullable(this.bucket);
    }

    /**
     * The domain associated with the entity.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return The domain associated with the entity.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * The email address associated with the entity.
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return The email address associated with the entity.
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * The entity holding the permission, in one of the following forms:
     * * user-{{userId}}
     * * user-{{email}} (such as &#34;user-liz@example.com&#34;)
     * * group-{{groupId}}
     * * group-{{email}} (such as &#34;group-example@googlegroups.com&#34;)
     * * domain-{{domain}} (such as &#34;domain-example.com&#34;)
     * * project-team-{{projectId}}
     * * allUsers
     * * allAuthenticatedUsers
     * 
     */
    @Import(name="entity")
    private @Nullable Output<String> entity;

    /**
     * @return The entity holding the permission, in one of the following forms:
     * * user-{{userId}}
     * * user-{{email}} (such as &#34;user-liz@example.com&#34;)
     * * group-{{groupId}}
     * * group-{{email}} (such as &#34;group-example@googlegroups.com&#34;)
     * * domain-{{domain}} (such as &#34;domain-example.com&#34;)
     * * project-team-{{projectId}}
     * * allUsers
     * * allAuthenticatedUsers
     * 
     */
    public Optional<Output<String>> entity() {
        return Optional.ofNullable(this.entity);
    }

    /**
     * The ID for the entity
     * 
     */
    @Import(name="entityId")
    private @Nullable Output<String> entityId;

    /**
     * @return The ID for the entity
     * 
     */
    public Optional<Output<String>> entityId() {
        return Optional.ofNullable(this.entityId);
    }

    /**
     * The content generation of the object, if applied to an object.
     * 
     */
    @Import(name="generation")
    private @Nullable Output<Integer> generation;

    /**
     * @return The content generation of the object, if applied to an object.
     * 
     */
    public Optional<Output<Integer>> generation() {
        return Optional.ofNullable(this.generation);
    }

    /**
     * The name of the object, if applied to an object.
     * 
     */
    @Import(name="object")
    private @Nullable Output<String> object;

    /**
     * @return The name of the object, if applied to an object.
     * 
     */
    public Optional<Output<String>> object() {
        return Optional.ofNullable(this.object);
    }

    /**
     * The project team associated with the entity
     * 
     */
    @Import(name="projectTeams")
    private @Nullable Output<List<DefaultObjectAccessControlProjectTeamArgs>> projectTeams;

    /**
     * @return The project team associated with the entity
     * 
     */
    public Optional<Output<List<DefaultObjectAccessControlProjectTeamArgs>>> projectTeams() {
        return Optional.ofNullable(this.projectTeams);
    }

    /**
     * The access permission for the entity.
     * Possible values are `OWNER` and `READER`.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return The access permission for the entity.
     * Possible values are `OWNER` and `READER`.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    private DefaultObjectAccessControlState() {}

    private DefaultObjectAccessControlState(DefaultObjectAccessControlState $) {
        this.bucket = $.bucket;
        this.domain = $.domain;
        this.email = $.email;
        this.entity = $.entity;
        this.entityId = $.entityId;
        this.generation = $.generation;
        this.object = $.object;
        this.projectTeams = $.projectTeams;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DefaultObjectAccessControlState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DefaultObjectAccessControlState $;

        public Builder() {
            $ = new DefaultObjectAccessControlState();
        }

        public Builder(DefaultObjectAccessControlState defaults) {
            $ = new DefaultObjectAccessControlState(Objects.requireNonNull(defaults));
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(@Nullable Output<String> bucket) {
            $.bucket = bucket;
            return this;
        }

        /**
         * @param bucket The name of the bucket.
         * 
         * @return builder
         * 
         */
        public Builder bucket(String bucket) {
            return bucket(Output.of(bucket));
        }

        /**
         * @param domain The domain associated with the entity.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The domain associated with the entity.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param email The email address associated with the entity.
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email The email address associated with the entity.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param entity The entity holding the permission, in one of the following forms:
         * * user-{{userId}}
         * * user-{{email}} (such as &#34;user-liz@example.com&#34;)
         * * group-{{groupId}}
         * * group-{{email}} (such as &#34;group-example@googlegroups.com&#34;)
         * * domain-{{domain}} (such as &#34;domain-example.com&#34;)
         * * project-team-{{projectId}}
         * * allUsers
         * * allAuthenticatedUsers
         * 
         * @return builder
         * 
         */
        public Builder entity(@Nullable Output<String> entity) {
            $.entity = entity;
            return this;
        }

        /**
         * @param entity The entity holding the permission, in one of the following forms:
         * * user-{{userId}}
         * * user-{{email}} (such as &#34;user-liz@example.com&#34;)
         * * group-{{groupId}}
         * * group-{{email}} (such as &#34;group-example@googlegroups.com&#34;)
         * * domain-{{domain}} (such as &#34;domain-example.com&#34;)
         * * project-team-{{projectId}}
         * * allUsers
         * * allAuthenticatedUsers
         * 
         * @return builder
         * 
         */
        public Builder entity(String entity) {
            return entity(Output.of(entity));
        }

        /**
         * @param entityId The ID for the entity
         * 
         * @return builder
         * 
         */
        public Builder entityId(@Nullable Output<String> entityId) {
            $.entityId = entityId;
            return this;
        }

        /**
         * @param entityId The ID for the entity
         * 
         * @return builder
         * 
         */
        public Builder entityId(String entityId) {
            return entityId(Output.of(entityId));
        }

        /**
         * @param generation The content generation of the object, if applied to an object.
         * 
         * @return builder
         * 
         */
        public Builder generation(@Nullable Output<Integer> generation) {
            $.generation = generation;
            return this;
        }

        /**
         * @param generation The content generation of the object, if applied to an object.
         * 
         * @return builder
         * 
         */
        public Builder generation(Integer generation) {
            return generation(Output.of(generation));
        }

        /**
         * @param object The name of the object, if applied to an object.
         * 
         * @return builder
         * 
         */
        public Builder object(@Nullable Output<String> object) {
            $.object = object;
            return this;
        }

        /**
         * @param object The name of the object, if applied to an object.
         * 
         * @return builder
         * 
         */
        public Builder object(String object) {
            return object(Output.of(object));
        }

        /**
         * @param projectTeams The project team associated with the entity
         * 
         * @return builder
         * 
         */
        public Builder projectTeams(@Nullable Output<List<DefaultObjectAccessControlProjectTeamArgs>> projectTeams) {
            $.projectTeams = projectTeams;
            return this;
        }

        /**
         * @param projectTeams The project team associated with the entity
         * 
         * @return builder
         * 
         */
        public Builder projectTeams(List<DefaultObjectAccessControlProjectTeamArgs> projectTeams) {
            return projectTeams(Output.of(projectTeams));
        }

        /**
         * @param projectTeams The project team associated with the entity
         * 
         * @return builder
         * 
         */
        public Builder projectTeams(DefaultObjectAccessControlProjectTeamArgs... projectTeams) {
            return projectTeams(List.of(projectTeams));
        }

        /**
         * @param role The access permission for the entity.
         * Possible values are `OWNER` and `READER`.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The access permission for the entity.
         * Possible values are `OWNER` and `READER`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public DefaultObjectAccessControlState build() {
            return $;
        }
    }

}
