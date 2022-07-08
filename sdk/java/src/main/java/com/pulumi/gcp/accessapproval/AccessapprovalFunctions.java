// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accessapproval;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.accessapproval.inputs.GetFolderServiceAccountArgs;
import com.pulumi.gcp.accessapproval.inputs.GetFolderServiceAccountPlainArgs;
import com.pulumi.gcp.accessapproval.inputs.GetOrganizationServiceAccountArgs;
import com.pulumi.gcp.accessapproval.inputs.GetOrganizationServiceAccountPlainArgs;
import com.pulumi.gcp.accessapproval.inputs.GetProjectServiceAccountArgs;
import com.pulumi.gcp.accessapproval.inputs.GetProjectServiceAccountPlainArgs;
import com.pulumi.gcp.accessapproval.outputs.GetFolderServiceAccountResult;
import com.pulumi.gcp.accessapproval.outputs.GetOrganizationServiceAccountResult;
import com.pulumi.gcp.accessapproval.outputs.GetProjectServiceAccountResult;
import java.util.concurrent.CompletableFuture;

public final class AccessapprovalFunctions {
    /**
     * Get the email address of a folder&#39;s Access Approval service account.
     * 
     * Each Google Cloud folder has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getFolderServiceAccount(GetFolderServiceAccountArgs.builder()
     *             .folderId(&#34;my-folder&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getFolderServiceAccountResult -&gt; getFolderServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFolderServiceAccountResult> getFolderServiceAccount(GetFolderServiceAccountArgs args) {
        return getFolderServiceAccount(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a folder&#39;s Access Approval service account.
     * 
     * Each Google Cloud folder has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getFolderServiceAccount(GetFolderServiceAccountArgs.builder()
     *             .folderId(&#34;my-folder&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getFolderServiceAccountResult -&gt; getFolderServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFolderServiceAccountResult> getFolderServiceAccountPlain(GetFolderServiceAccountPlainArgs args) {
        return getFolderServiceAccountPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a folder&#39;s Access Approval service account.
     * 
     * Each Google Cloud folder has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getFolderServiceAccount(GetFolderServiceAccountArgs.builder()
     *             .folderId(&#34;my-folder&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getFolderServiceAccountResult -&gt; getFolderServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetFolderServiceAccountResult> getFolderServiceAccount(GetFolderServiceAccountArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:accessapproval/getFolderServiceAccount:getFolderServiceAccount", TypeShape.of(GetFolderServiceAccountResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the email address of a folder&#39;s Access Approval service account.
     * 
     * Each Google Cloud folder has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getFolderServiceAccount(GetFolderServiceAccountArgs.builder()
     *             .folderId(&#34;my-folder&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getFolderServiceAccountResult -&gt; getFolderServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetFolderServiceAccountResult> getFolderServiceAccountPlain(GetFolderServiceAccountPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:accessapproval/getFolderServiceAccount:getFolderServiceAccount", TypeShape.of(GetFolderServiceAccountResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the email address of an organization&#39;s Access Approval service account.
     * 
     * Each Google Cloud organization has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getOrganizationServiceAccount(GetOrganizationServiceAccountArgs.builder()
     *             .organizationId(&#34;my-organization&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getOrganizationServiceAccountResult -&gt; getOrganizationServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetOrganizationServiceAccountResult> getOrganizationServiceAccount(GetOrganizationServiceAccountArgs args) {
        return getOrganizationServiceAccount(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of an organization&#39;s Access Approval service account.
     * 
     * Each Google Cloud organization has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getOrganizationServiceAccount(GetOrganizationServiceAccountArgs.builder()
     *             .organizationId(&#34;my-organization&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getOrganizationServiceAccountResult -&gt; getOrganizationServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetOrganizationServiceAccountResult> getOrganizationServiceAccountPlain(GetOrganizationServiceAccountPlainArgs args) {
        return getOrganizationServiceAccountPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of an organization&#39;s Access Approval service account.
     * 
     * Each Google Cloud organization has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getOrganizationServiceAccount(GetOrganizationServiceAccountArgs.builder()
     *             .organizationId(&#34;my-organization&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getOrganizationServiceAccountResult -&gt; getOrganizationServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetOrganizationServiceAccountResult> getOrganizationServiceAccount(GetOrganizationServiceAccountArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:accessapproval/getOrganizationServiceAccount:getOrganizationServiceAccount", TypeShape.of(GetOrganizationServiceAccountResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the email address of an organization&#39;s Access Approval service account.
     * 
     * Each Google Cloud organization has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getOrganizationServiceAccount(GetOrganizationServiceAccountArgs.builder()
     *             .organizationId(&#34;my-organization&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getOrganizationServiceAccountResult -&gt; getOrganizationServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetOrganizationServiceAccountResult> getOrganizationServiceAccountPlain(GetOrganizationServiceAccountPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:accessapproval/getOrganizationServiceAccount:getOrganizationServiceAccount", TypeShape.of(GetOrganizationServiceAccountResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the email address of a project&#39;s Access Approval service account.
     * 
     * Each Google Cloud project has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getProjectServiceAccount(GetProjectServiceAccountArgs.builder()
     *             .projectId(&#34;my-project&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getProjectServiceAccountResult -&gt; getProjectServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetProjectServiceAccountResult> getProjectServiceAccount(GetProjectServiceAccountArgs args) {
        return getProjectServiceAccount(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a project&#39;s Access Approval service account.
     * 
     * Each Google Cloud project has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getProjectServiceAccount(GetProjectServiceAccountArgs.builder()
     *             .projectId(&#34;my-project&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getProjectServiceAccountResult -&gt; getProjectServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetProjectServiceAccountResult> getProjectServiceAccountPlain(GetProjectServiceAccountPlainArgs args) {
        return getProjectServiceAccountPlain(args, InvokeOptions.Empty);
    }
    /**
     * Get the email address of a project&#39;s Access Approval service account.
     * 
     * Each Google Cloud project has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getProjectServiceAccount(GetProjectServiceAccountArgs.builder()
     *             .projectId(&#34;my-project&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getProjectServiceAccountResult -&gt; getProjectServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static Output<GetProjectServiceAccountResult> getProjectServiceAccount(GetProjectServiceAccountArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("gcp:accessapproval/getProjectServiceAccount:getProjectServiceAccount", TypeShape.of(GetProjectServiceAccountResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Get the email address of a project&#39;s Access Approval service account.
     * 
     * Each Google Cloud project has a unique service account used by Access Approval.
     * When using Access Approval with a
     * [custom signing key](https://cloud.google.com/cloud-provider-access-management/access-approval/docs/review-approve-access-requests-custom-keys),
     * this account needs to be granted the `cloudkms.signerVerifier` IAM role on the
     * Cloud KMS key used to sign approvals.
     * 
     * ## Example Usage
     * ```java
     * package generated_program;
     * 
     * import java.util.*;
     * import java.io.*;
     * import java.nio.*;
     * import com.pulumi.*;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         final var serviceAccount = Output.of(AccessapprovalFunctions.getProjectServiceAccount(GetProjectServiceAccountArgs.builder()
     *             .projectId(&#34;my-project&#34;)
     *             .build()));
     * 
     *         var iam = new CryptoKeyIAMMember(&#34;iam&#34;, CryptoKeyIAMMemberArgs.builder()        
     *             .cryptoKeyId(google_kms_crypto_key.crypto_key().id())
     *             .role(&#34;roles/cloudkms.signerVerifier&#34;)
     *             .member(String.format(&#34;serviceAccount:%s&#34;, serviceAccount.apply(getProjectServiceAccountResult -&gt; getProjectServiceAccountResult.accountEmail())))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public static CompletableFuture<GetProjectServiceAccountResult> getProjectServiceAccountPlain(GetProjectServiceAccountPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("gcp:accessapproval/getProjectServiceAccount:getProjectServiceAccount", TypeShape.of(GetProjectServiceAccountResult.class), args, Utilities.withVersion(options));
    }
}
