// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.iap.TunnelInstanceIAMMemberArgs;
import com.pulumi.gcp.iap.inputs.TunnelInstanceIAMMemberState;
import com.pulumi.gcp.iap.outputs.TunnelInstanceIAMMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelInstance. Each of these resources serves a different use case:
 * 
 * * `gcp.iap.TunnelInstanceIAMPolicy`: Authoritative. Sets the IAM policy for the tunnelinstance and replaces any existing policy already attached.
 * * `gcp.iap.TunnelInstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnelinstance are preserved.
 * * `gcp.iap.TunnelInstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnelinstance are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.iap.TunnelInstanceIAMPolicy`: Retrieves the IAM policy for the tunnelinstance
 * 
 * &gt; **Note:** `gcp.iap.TunnelInstanceIAMPolicy` **cannot** be used in conjunction with `gcp.iap.TunnelInstanceIAMBinding` and `gcp.iap.TunnelInstanceIAMMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.iap.TunnelInstanceIAMBinding` resources **can be** used in conjunction with `gcp.iap.TunnelInstanceIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
 * 
 * ## gcp.iap.TunnelInstanceIAMPolicy
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
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicy;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/iap.tunnelResourceAccessor")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new TunnelInstanceIAMPolicy("policy", TunnelInstanceIAMPolicyArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
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
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicy;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/iap.tunnelResourceAccessor")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title("expires_after_2019_12_31")
 *                     .description("Expiring at midnight of 2019-12-31")
 *                     .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new TunnelInstanceIAMPolicy("policy", TunnelInstanceIAMPolicyArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## gcp.iap.TunnelInstanceIAMBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBinding;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBindingArgs;
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
 *         var binding = new TunnelInstanceIAMBinding("binding", TunnelInstanceIAMBindingArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBinding;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBindingArgs;
 * import com.pulumi.gcp.iap.inputs.TunnelInstanceIAMBindingConditionArgs;
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
 *         var binding = new TunnelInstanceIAMBinding("binding", TunnelInstanceIAMBindingArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .condition(TunnelInstanceIAMBindingConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## gcp.iap.TunnelInstanceIAMMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMember;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMemberArgs;
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
 *         var member = new TunnelInstanceIAMMember("member", TunnelInstanceIAMMemberArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMember;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMemberArgs;
 * import com.pulumi.gcp.iap.inputs.TunnelInstanceIAMMemberConditionArgs;
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
 *         var member = new TunnelInstanceIAMMember("member", TunnelInstanceIAMMemberArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .condition(TunnelInstanceIAMMemberConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## This resource supports User Project Overrides.
 * 
 * - 
 * 
 * # IAM policy for Identity-Aware Proxy TunnelInstance
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelInstance. Each of these resources serves a different use case:
 * 
 * * `gcp.iap.TunnelInstanceIAMPolicy`: Authoritative. Sets the IAM policy for the tunnelinstance and replaces any existing policy already attached.
 * * `gcp.iap.TunnelInstanceIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunnelinstance are preserved.
 * * `gcp.iap.TunnelInstanceIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunnelinstance are preserved.
 * 
 * A data source can be used to retrieve policy data in advent you do not need creation
 * 
 * * `gcp.iap.TunnelInstanceIAMPolicy`: Retrieves the IAM policy for the tunnelinstance
 * 
 * &gt; **Note:** `gcp.iap.TunnelInstanceIAMPolicy` **cannot** be used in conjunction with `gcp.iap.TunnelInstanceIAMBinding` and `gcp.iap.TunnelInstanceIAMMember` or they will fight over what your policy should be.
 * 
 * &gt; **Note:** `gcp.iap.TunnelInstanceIAMBinding` resources **can be** used in conjunction with `gcp.iap.TunnelInstanceIAMMember` resources **only if** they do not grant privilege to the same role.
 * 
 * &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
 * 
 * ## gcp.iap.TunnelInstanceIAMPolicy
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
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicy;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/iap.tunnelResourceAccessor")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .build())
 *             .build());
 * 
 *         var policy = new TunnelInstanceIAMPolicy("policy", TunnelInstanceIAMPolicyArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
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
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicy;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMPolicyArgs;
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
 *         final var admin = OrganizationsFunctions.getIAMPolicy(GetIAMPolicyArgs.builder()
 *             .bindings(GetIAMPolicyBindingArgs.builder()
 *                 .role("roles/iap.tunnelResourceAccessor")
 *                 .members("user:jane}{@literal @}{@code example.com")
 *                 .condition(GetIAMPolicyBindingConditionArgs.builder()
 *                     .title("expires_after_2019_12_31")
 *                     .description("Expiring at midnight of 2019-12-31")
 *                     .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var policy = new TunnelInstanceIAMPolicy("policy", TunnelInstanceIAMPolicyArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .policyData(admin.applyValue(getIAMPolicyResult -> getIAMPolicyResult.policyData()))
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## gcp.iap.TunnelInstanceIAMBinding
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBinding;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBindingArgs;
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
 *         var binding = new TunnelInstanceIAMBinding("binding", TunnelInstanceIAMBindingArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBinding;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMBindingArgs;
 * import com.pulumi.gcp.iap.inputs.TunnelInstanceIAMBindingConditionArgs;
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
 *         var binding = new TunnelInstanceIAMBinding("binding", TunnelInstanceIAMBindingArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .members("user:jane}{@literal @}{@code example.com")
 *             .condition(TunnelInstanceIAMBindingConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## gcp.iap.TunnelInstanceIAMMember
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMember;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMemberArgs;
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
 *         var member = new TunnelInstanceIAMMember("member", TunnelInstanceIAMMemberArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * With IAM Conditions:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMember;
 * import com.pulumi.gcp.iap.TunnelInstanceIAMMemberArgs;
 * import com.pulumi.gcp.iap.inputs.TunnelInstanceIAMMemberConditionArgs;
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
 *         var member = new TunnelInstanceIAMMember("member", TunnelInstanceIAMMemberArgs.builder()
 *             .project(tunnelvm.project())
 *             .zone(tunnelvm.zone())
 *             .instance(tunnelvm.name())
 *             .role("roles/iap.tunnelResourceAccessor")
 *             .member("user:jane}{@literal @}{@code example.com")
 *             .condition(TunnelInstanceIAMMemberConditionArgs.builder()
 *                 .title("expires_after_2019_12_31")
 *                 .description("Expiring at midnight of 2019-12-31")
 *                 .expression("request.time < timestamp(\"2020-01-01T00:00:00Z\")")
 *                 .build())
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
 * For all import syntaxes, the &#34;resource in question&#34; can take any of the following forms:
 * 
 * * projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{name}}
 * 
 * * projects/{{project}}/zones/{{zone}}/instances/{{name}}
 * 
 * * {{project}}/{{zone}}/{{name}}
 * 
 * * {{zone}}/{{name}}
 * 
 * * {{name}}
 * 
 * Any variables not passed in the import command will be taken from the provider configuration.
 * 
 * Identity-Aware Proxy tunnelinstance IAM resources can be imported using the resource identifiers, role, and member.
 * 
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor &#34;projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor user:jane{@literal @}example.com&#34;
 * ```
 * 
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor &#34;projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}} roles/iap.tunnelResourceAccessor&#34;
 * ```
 * 
 * IAM policy imports use the identifier of the resource in question, e.g.
 * 
 * ```sh
 * $ pulumi import gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember editor projects/{{project}}/iap_tunnel/zones/{{zone}}/instances/{{tunnel_instance}}
 * ```
 * 
 * -&gt; **Custom Roles**: If you&#39;re importing a IAM resource with a custom role, make sure to use the
 * 
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 * 
 */
@ResourceType(type="gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember")
public class TunnelInstanceIAMMember extends com.pulumi.resources.CustomResource {
    /**
     * An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    @Export(name="condition", refs={TunnelInstanceIAMMemberCondition.class}, tree="[0]")
    private Output</* @Nullable */ TunnelInstanceIAMMemberCondition> condition;

    /**
     * @return An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
     * Structure is documented below.
     * 
     */
    public Output<Optional<TunnelInstanceIAMMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    /**
     * (Computed) The etag of the IAM policy.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return (Computed) The etag of the IAM policy.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Export(name="instance", refs={String.class}, tree="[0]")
    private Output<String> instance;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    @Export(name="member", refs={String.class}, tree="[0]")
    private Output<String> member;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    public Output<String> member() {
        return this.member;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The role that should be applied. Only one
     * `gcp.iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.iap.TunnelInstanceIAMBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TunnelInstanceIAMMember(java.lang.String name) {
        this(name, TunnelInstanceIAMMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TunnelInstanceIAMMember(java.lang.String name, TunnelInstanceIAMMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TunnelInstanceIAMMember(java.lang.String name, TunnelInstanceIAMMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TunnelInstanceIAMMember(java.lang.String name, Output<java.lang.String> id, @Nullable TunnelInstanceIAMMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:iap/tunnelInstanceIAMMember:TunnelInstanceIAMMember", name, state, makeResourceOptions(options, id), false);
    }

    private static TunnelInstanceIAMMemberArgs makeArgs(TunnelInstanceIAMMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TunnelInstanceIAMMemberArgs.Empty : args;
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
    public static TunnelInstanceIAMMember get(java.lang.String name, Output<java.lang.String> id, @Nullable TunnelInstanceIAMMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TunnelInstanceIAMMember(name, id, state, options);
    }
}
