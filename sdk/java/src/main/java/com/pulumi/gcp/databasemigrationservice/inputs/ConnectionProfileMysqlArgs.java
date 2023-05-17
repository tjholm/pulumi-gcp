// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.databasemigrationservice.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.databasemigrationservice.inputs.ConnectionProfileMysqlSslArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ConnectionProfileMysqlArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionProfileMysqlArgs Empty = new ConnectionProfileMysqlArgs();

    /**
     * If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
     * 
     */
    @Import(name="cloudSqlId")
    private @Nullable Output<String> cloudSqlId;

    /**
     * @return If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
     * 
     */
    public Optional<Output<String>> cloudSqlId() {
        return Optional.ofNullable(this.cloudSqlId);
    }

    /**
     * Required. The IP or hostname of the source MySQL database.
     * 
     */
    @Import(name="host", required=true)
    private Output<String> host;

    /**
     * @return Required. The IP or hostname of the source MySQL database.
     * 
     */
    public Output<String> host() {
        return this.host;
    }

    /**
     * Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.
     * This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.
     * This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    /**
     * (Output)
     * Output only. Indicates If this connection profile password is stored.
     * 
     */
    @Import(name="passwordSet")
    private @Nullable Output<Boolean> passwordSet;

    /**
     * @return (Output)
     * Output only. Indicates If this connection profile password is stored.
     * 
     */
    public Optional<Output<Boolean>> passwordSet() {
        return Optional.ofNullable(this.passwordSet);
    }

    /**
     * Required. The network port of the source MySQL database.
     * 
     */
    @Import(name="port", required=true)
    private Output<Integer> port;

    /**
     * @return Required. The network port of the source MySQL database.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }

    /**
     * SSL configuration for the destination to connect to the source database.
     * Structure is documented below.
     * 
     */
    @Import(name="ssl")
    private @Nullable Output<ConnectionProfileMysqlSslArgs> ssl;

    /**
     * @return SSL configuration for the destination to connect to the source database.
     * Structure is documented below.
     * 
     */
    public Optional<Output<ConnectionProfileMysqlSslArgs>> ssl() {
        return Optional.ofNullable(this.ssl);
    }

    /**
     * Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private ConnectionProfileMysqlArgs() {}

    private ConnectionProfileMysqlArgs(ConnectionProfileMysqlArgs $) {
        this.cloudSqlId = $.cloudSqlId;
        this.host = $.host;
        this.password = $.password;
        this.passwordSet = $.passwordSet;
        this.port = $.port;
        this.ssl = $.ssl;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionProfileMysqlArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionProfileMysqlArgs $;

        public Builder() {
            $ = new ConnectionProfileMysqlArgs();
        }

        public Builder(ConnectionProfileMysqlArgs defaults) {
            $ = new ConnectionProfileMysqlArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cloudSqlId If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
         * 
         * @return builder
         * 
         */
        public Builder cloudSqlId(@Nullable Output<String> cloudSqlId) {
            $.cloudSqlId = cloudSqlId;
            return this;
        }

        /**
         * @param cloudSqlId If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
         * 
         * @return builder
         * 
         */
        public Builder cloudSqlId(String cloudSqlId) {
            return cloudSqlId(Output.of(cloudSqlId));
        }

        /**
         * @param host Required. The IP or hostname of the source MySQL database.
         * 
         * @return builder
         * 
         */
        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host Required. The IP or hostname of the source MySQL database.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param password Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.
         * This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
         * **Note**: This property is sensitive and will not be displayed in the plan.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.
         * This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
         * **Note**: This property is sensitive and will not be displayed in the plan.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param passwordSet (Output)
         * Output only. Indicates If this connection profile password is stored.
         * 
         * @return builder
         * 
         */
        public Builder passwordSet(@Nullable Output<Boolean> passwordSet) {
            $.passwordSet = passwordSet;
            return this;
        }

        /**
         * @param passwordSet (Output)
         * Output only. Indicates If this connection profile password is stored.
         * 
         * @return builder
         * 
         */
        public Builder passwordSet(Boolean passwordSet) {
            return passwordSet(Output.of(passwordSet));
        }

        /**
         * @param port Required. The network port of the source MySQL database.
         * 
         * @return builder
         * 
         */
        public Builder port(Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Required. The network port of the source MySQL database.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param ssl SSL configuration for the destination to connect to the source database.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ssl(@Nullable Output<ConnectionProfileMysqlSslArgs> ssl) {
            $.ssl = ssl;
            return this;
        }

        /**
         * @param ssl SSL configuration for the destination to connect to the source database.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder ssl(ConnectionProfileMysqlSslArgs ssl) {
            return ssl(Output.of(ssl));
        }

        /**
         * @param username Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ConnectionProfileMysqlArgs build() {
            $.host = Objects.requireNonNull($.host, "expected parameter 'host' to be non-null");
            $.password = Objects.requireNonNull($.password, "expected parameter 'password' to be non-null");
            $.port = Objects.requireNonNull($.port, "expected parameter 'port' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}
