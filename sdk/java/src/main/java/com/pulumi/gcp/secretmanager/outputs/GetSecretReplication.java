// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.secretmanager.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.secretmanager.outputs.GetSecretReplicationUserManaged;
import java.lang.Boolean;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetSecretReplication {
    private final Boolean automatic;
    private final List<GetSecretReplicationUserManaged> userManageds;

    @CustomType.Constructor
    private GetSecretReplication(
        @CustomType.Parameter("automatic") Boolean automatic,
        @CustomType.Parameter("userManageds") List<GetSecretReplicationUserManaged> userManageds) {
        this.automatic = automatic;
        this.userManageds = userManageds;
    }

    public Boolean automatic() {
        return this.automatic;
    }
    public List<GetSecretReplicationUserManaged> userManageds() {
        return this.userManageds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretReplication defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean automatic;
        private List<GetSecretReplicationUserManaged> userManageds;

        public Builder() {
    	      // Empty
        }

        public Builder(GetSecretReplication defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.automatic = defaults.automatic;
    	      this.userManageds = defaults.userManageds;
        }

        public Builder automatic(Boolean automatic) {
            this.automatic = Objects.requireNonNull(automatic);
            return this;
        }
        public Builder userManageds(List<GetSecretReplicationUserManaged> userManageds) {
            this.userManageds = Objects.requireNonNull(userManageds);
            return this;
        }
        public Builder userManageds(GetSecretReplicationUserManaged... userManageds) {
            return userManageds(List.of(userManageds));
        }        public GetSecretReplication build() {
            return new GetSecretReplication(automatic, userManageds);
        }
    }
}
